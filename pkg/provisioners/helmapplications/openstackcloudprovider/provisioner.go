/*
Copyright 2022-2023 EscherCloud.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstackcloudprovider

import (
	"errors"
	"fmt"

	"github.com/gophercloud/utils/openstack/clientconfig"

	unikornv1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	"github.com/eschercloudai/unikorn/pkg/provisioners"
	"github.com/eschercloudai/unikorn/pkg/provisioners/application"
	"github.com/eschercloudai/unikorn/pkg/provisioners/util"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	// applicationName is the unique name of the application.
	applicationName = "openstack-cloud-provider"
)

var (
	// ErrCloudConfiguration is returned when the cloud configuration is not
	// correctly formatted.
	ErrCloudConfiguration = errors.New("invalid cloud configuration")
)

// Provisioner encapsulates control plane provisioning.
type Provisioner struct {
	// cluster is the Kubernetes cluster we're provisioning.
	cluster *unikornv1.KubernetesCluster
}

// New returns a new initialized provisioner object.
func New(client client.Client, cluster *unikornv1.KubernetesCluster, helm *unikornv1.HelmApplication) provisioners.Provisioner {
	provisioner := &Provisioner{
		cluster: cluster,
	}

	return application.New(client, applicationName, cluster, helm).WithGenerator(provisioner).InNamespace("ocp-system")
}

// Ensure the Provisioner interface is implemented.
var _ application.ValuesGenerator = &Provisioner{}

// generateGlobalValues does the horrific translation
// between the myriad ways that OpenStack deems necessary to authenticate to the
// cloud configuration format.  See:
// https://github.com/kubernetes/cloud-provider-openstack/blob/master/docs/openstack-cloud-controller-manager/using-openstack-cloud-controller-manager.md#config-openstack-cloud-controller-manager
//
//nolint:cyclop
func (p *Provisioner) generateGlobalValues() (map[string]interface{}, error) {
	var clouds clientconfig.Clouds

	if err := yaml.Unmarshal(*p.cluster.Spec.Openstack.CloudConfig, &clouds); err != nil {
		return nil, err
	}

	cloud, ok := clouds.Clouds[*p.cluster.Spec.Openstack.Cloud]
	if !ok {
		return nil, fmt.Errorf("%w: cloud '%s' not found in clouds.yaml", ErrCloudConfiguration, *p.cluster.Spec.Openstack.Cloud)
	}

	// Like the "openstack" command we require application credentials to
	// be marked with the correct authentication type, passwords can be
	// explicitly or implictly typed, because that's just the way of the
	// world...  Password auth is just a convenience thing for ease of
	// development.  Production deployments will want to use application
	// credentials so (possibly external) credentials aren't leaked.  That
	// said given the whit show that is this code, it may be better to just
	// kill passwords.
	global := map[string]interface{}{
		"auth-url": cloud.AuthInfo.AuthURL,
	}

	//nolint:exhaustive
	switch cloud.AuthType {
	case "", clientconfig.AuthV3Password:
		// The user_id field is NOT supported by the provider.
		if cloud.AuthInfo.Username == "" {
			return nil, fmt.Errorf("%w: username must be specified in clouds.yaml", ErrCloudConfiguration)
		}

		global["username"] = cloud.AuthInfo.Username
		global["password"] = cloud.AuthInfo.Password

		// Try a flat, single domain first, then -- failing that -- look
		// for a more hierarchical topology.
		switch {
		case cloud.AuthInfo.DomainID != "":
			global["domain-id"] = cloud.AuthInfo.DomainID
		case cloud.AuthInfo.DomainName != "":
			global["domain-name"] = cloud.AuthInfo.DomainName
		default:
			switch {
			case cloud.AuthInfo.UserDomainID != "":
				global["user-domain-id"] = cloud.AuthInfo.UserDomainID
			case cloud.AuthInfo.UserDomainName != "":
				global["user-domain-name"] = cloud.AuthInfo.UserDomainName
			default:
				return nil, fmt.Errorf("%w: domain_name, domain_id, user_domain_name or user_domain_id must be specified in clouds.yaml", ErrCloudConfiguration)
			}

			switch {
			case cloud.AuthInfo.ProjectDomainID != "":
				global["tenant-domain-id"] = cloud.AuthInfo.ProjectDomainID
			case cloud.AuthInfo.ProjectDomainName != "":
				global["tenant-domain-name"] = cloud.AuthInfo.ProjectDomainName
			default:
				return nil, fmt.Errorf("%w: domain_name, domain_id, project_domain_name or project_domain_id must be specified in clouds.yaml", ErrCloudConfiguration)
			}
		}

		switch {
		case cloud.AuthInfo.ProjectID != "":
			global["tenant-id"] = cloud.AuthInfo.ProjectID
		case cloud.AuthInfo.ProjectName != "":
			global["tenant-name"] = cloud.AuthInfo.ProjectName
		default:
			return nil, fmt.Errorf("%w: project_name or project_id must be specified in clouds.yaml", ErrCloudConfiguration)
		}

	case clientconfig.AuthV3ApplicationCredential:
		global["application-credential-id"] = cloud.AuthInfo.ApplicationCredentialID
		global["application-credential-secret"] = cloud.AuthInfo.ApplicationCredentialSecret

	default:
		return nil, fmt.Errorf("%w: v3password or v3applicationcredential auth_type must be specified in clouds.yaml", ErrCloudConfiguration)
	}

	return global, nil
}

// Generate implements the application.Generator interface.
// Note there is an option, to just pass through the clouds.yaml file, however
// the chart doesn't allow it to be exposed so we need to translate between formats.
func (p *Provisioner) Values(version *string) (interface{}, error) {
	cloudConfigGlobal, err := p.generateGlobalValues()
	if err != nil {
		return nil, err
	}

	tolerations := util.ControlPlaneTolerations()
	tolerations = append(tolerations, util.ControlPlaneInitTolerations()...)

	values := map[string]interface{}{
		"cloudConfig": map[string]interface{}{
			"global": cloudConfigGlobal,
			"loadBalancer": map[string]interface{}{
				"floating-network-id": *p.cluster.Spec.Openstack.ExternalNetworkID,
			},
			"blockStorage": map[string]interface{}{
				"ignore-volume-az": true,
			},
		},
		"tolerations": tolerations,
		// See https://github.com/kubernetes/cloud-provider-openstack/issues/2049 for
		// more details, and no-one doing anything about it.
		"controllerExtraArgs": `{{list "--use-service-account-credentials=false" | toYaml}}`,
	}

	return values, nil
}