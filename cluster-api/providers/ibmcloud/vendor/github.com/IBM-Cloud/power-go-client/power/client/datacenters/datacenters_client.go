// Code generated by go-swagger; DO NOT EDIT.

package datacenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new datacenters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new datacenters API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new datacenters API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for datacenters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	V1DatacentersGet(params *V1DatacentersGetParams, opts ...ClientOption) (*V1DatacentersGetOK, error)

	V1DatacentersGetall(params *V1DatacentersGetallParams, opts ...ClientOption) (*V1DatacentersGetallOK, error)

	V1DatacentersPrivateGet(params *V1DatacentersPrivateGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1DatacentersPrivateGetOK, error)

	V1DatacentersPrivateGetall(params *V1DatacentersPrivateGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1DatacentersPrivateGetallOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
V1DatacentersGet gets a datacenter s information and capabilities
*/
func (a *Client) V1DatacentersGet(params *V1DatacentersGetParams, opts ...ClientOption) (*V1DatacentersGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1DatacentersGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.datacenters.get",
		Method:             "GET",
		PathPattern:        "/v1/datacenters/{datacenter_region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1DatacentersGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1DatacentersGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.datacenters.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1DatacentersGetall gets all datacenters information and capabilities
*/
func (a *Client) V1DatacentersGetall(params *V1DatacentersGetallParams, opts ...ClientOption) (*V1DatacentersGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1DatacentersGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.datacenters.getall",
		Method:             "GET",
		PathPattern:        "/v1/datacenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1DatacentersGetallReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1DatacentersGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.datacenters.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1DatacentersPrivateGet gets a private datacenter s information and capabilities
*/
func (a *Client) V1DatacentersPrivateGet(params *V1DatacentersPrivateGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1DatacentersPrivateGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1DatacentersPrivateGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.datacentersPrivate.get",
		Method:             "GET",
		PathPattern:        "/v1/datacenters/private/{datacenter_region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1DatacentersPrivateGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1DatacentersPrivateGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.datacentersPrivate.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1DatacentersPrivateGetall gets private datacenter information and capabilities
*/
func (a *Client) V1DatacentersPrivateGetall(params *V1DatacentersPrivateGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1DatacentersPrivateGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1DatacentersPrivateGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.datacentersPrivate.getall",
		Method:             "GET",
		PathPattern:        "/v1/datacenters/private",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1DatacentersPrivateGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1DatacentersPrivateGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.datacentersPrivate.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
