// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetWellKnownOpenidConfiguration request
	GetWellKnownOpenidConfiguration(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOauth2V2Authorization request
	GetOauth2V2Authorization(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOauth2V2Jwks request
	GetOauth2V2Jwks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostOauth2V2Token request with any body
	PostOauth2V2TokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostOauth2V2TokenWithFormdataBody(ctx context.Context, body PostOauth2V2TokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOidcCallback request
	GetOidcCallback(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetWellKnownOpenidConfiguration(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetWellKnownOpenidConfigurationRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetOauth2V2Authorization(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOauth2V2AuthorizationRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetOauth2V2Jwks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOauth2V2JwksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostOauth2V2TokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostOauth2V2TokenRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostOauth2V2TokenWithFormdataBody(ctx context.Context, body PostOauth2V2TokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostOauth2V2TokenRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetOidcCallback(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOidcCallbackRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetWellKnownOpenidConfigurationRequest generates requests for GetWellKnownOpenidConfiguration
func NewGetWellKnownOpenidConfigurationRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/.well-known/openid-configuration")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetOauth2V2AuthorizationRequest generates requests for GetOauth2V2Authorization
func NewGetOauth2V2AuthorizationRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/oauth2/v2/authorization")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetOauth2V2JwksRequest generates requests for GetOauth2V2Jwks
func NewGetOauth2V2JwksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/oauth2/v2/jwks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostOauth2V2TokenRequestWithFormdataBody calls the generic PostOauth2V2Token builder with application/x-www-form-urlencoded body
func NewPostOauth2V2TokenRequestWithFormdataBody(server string, body PostOauth2V2TokenFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewPostOauth2V2TokenRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewPostOauth2V2TokenRequestWithBody generates requests for PostOauth2V2Token with any type of body
func NewPostOauth2V2TokenRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/oauth2/v2/token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetOidcCallbackRequest generates requests for GetOidcCallback
func NewGetOidcCallbackRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/oidc/callback")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetWellKnownOpenidConfiguration request
	GetWellKnownOpenidConfigurationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWellKnownOpenidConfigurationResponse, error)

	// GetOauth2V2Authorization request
	GetOauth2V2AuthorizationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOauth2V2AuthorizationResponse, error)

	// GetOauth2V2Jwks request
	GetOauth2V2JwksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOauth2V2JwksResponse, error)

	// PostOauth2V2Token request with any body
	PostOauth2V2TokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostOauth2V2TokenResponse, error)

	PostOauth2V2TokenWithFormdataBodyWithResponse(ctx context.Context, body PostOauth2V2TokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*PostOauth2V2TokenResponse, error)

	// GetOidcCallback request
	GetOidcCallbackWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOidcCallbackResponse, error)
}

type GetWellKnownOpenidConfigurationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenidConfiguration
}

// Status returns HTTPResponse.Status
func (r GetWellKnownOpenidConfigurationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetWellKnownOpenidConfigurationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOauth2V2AuthorizationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetOauth2V2AuthorizationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOauth2V2AuthorizationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOauth2V2JwksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *JsonWebKeySet
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetOauth2V2JwksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOauth2V2JwksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostOauth2V2TokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Token
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PostOauth2V2TokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostOauth2V2TokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOidcCallbackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetOidcCallbackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOidcCallbackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetWellKnownOpenidConfigurationWithResponse request returning *GetWellKnownOpenidConfigurationResponse
func (c *ClientWithResponses) GetWellKnownOpenidConfigurationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWellKnownOpenidConfigurationResponse, error) {
	rsp, err := c.GetWellKnownOpenidConfiguration(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetWellKnownOpenidConfigurationResponse(rsp)
}

// GetOauth2V2AuthorizationWithResponse request returning *GetOauth2V2AuthorizationResponse
func (c *ClientWithResponses) GetOauth2V2AuthorizationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOauth2V2AuthorizationResponse, error) {
	rsp, err := c.GetOauth2V2Authorization(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOauth2V2AuthorizationResponse(rsp)
}

// GetOauth2V2JwksWithResponse request returning *GetOauth2V2JwksResponse
func (c *ClientWithResponses) GetOauth2V2JwksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOauth2V2JwksResponse, error) {
	rsp, err := c.GetOauth2V2Jwks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOauth2V2JwksResponse(rsp)
}

// PostOauth2V2TokenWithBodyWithResponse request with arbitrary body returning *PostOauth2V2TokenResponse
func (c *ClientWithResponses) PostOauth2V2TokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostOauth2V2TokenResponse, error) {
	rsp, err := c.PostOauth2V2TokenWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostOauth2V2TokenResponse(rsp)
}

func (c *ClientWithResponses) PostOauth2V2TokenWithFormdataBodyWithResponse(ctx context.Context, body PostOauth2V2TokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*PostOauth2V2TokenResponse, error) {
	rsp, err := c.PostOauth2V2TokenWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostOauth2V2TokenResponse(rsp)
}

// GetOidcCallbackWithResponse request returning *GetOidcCallbackResponse
func (c *ClientWithResponses) GetOidcCallbackWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOidcCallbackResponse, error) {
	rsp, err := c.GetOidcCallback(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOidcCallbackResponse(rsp)
}

// ParseGetWellKnownOpenidConfigurationResponse parses an HTTP response from a GetWellKnownOpenidConfigurationWithResponse call
func ParseGetWellKnownOpenidConfigurationResponse(rsp *http.Response) (*GetWellKnownOpenidConfigurationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetWellKnownOpenidConfigurationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenidConfiguration
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetOauth2V2AuthorizationResponse parses an HTTP response from a GetOauth2V2AuthorizationWithResponse call
func ParseGetOauth2V2AuthorizationResponse(rsp *http.Response) (*GetOauth2V2AuthorizationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOauth2V2AuthorizationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetOauth2V2JwksResponse parses an HTTP response from a GetOauth2V2JwksWithResponse call
func ParseGetOauth2V2JwksResponse(rsp *http.Response) (*GetOauth2V2JwksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOauth2V2JwksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest JsonWebKeySet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostOauth2V2TokenResponse parses an HTTP response from a PostOauth2V2TokenWithResponse call
func ParsePostOauth2V2TokenResponse(rsp *http.Response) (*PostOauth2V2TokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostOauth2V2TokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Token
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetOidcCallbackResponse parses an HTTP response from a GetOidcCallbackWithResponse call
func ParseGetOidcCallbackResponse(rsp *http.Response) (*GetOidcCallbackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOidcCallbackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
