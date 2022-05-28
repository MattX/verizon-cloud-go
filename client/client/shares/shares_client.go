// Code generated by go-swagger; DO NOT EDIT.

package shares

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new shares API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for shares API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateShare(params *CreateShareParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateShareCreated, error)

	DeleteShare(params *DeleteShareParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteShareNoContent, error)

	ListShares(params *ListSharesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSharesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateShare creates share

  Allows a user to create a share.
*/
func (a *Client) CreateShare(params *CreateShareParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateShareCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateShareParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_share",
		Method:             "POST",
		PathPattern:        "/shares",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateShareReader{formats: a.formats},
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
	success, ok := result.(*CreateShareCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_share: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteShare deletes a share

  Allows a share owner to delete a share by passing shareUid.
*/
func (a *Client) DeleteShare(params *DeleteShareParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteShareNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteShareParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_share",
		Method:             "DELETE",
		PathPattern:        "/shares/{shareUid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteShareReader{formats: a.formats},
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
	success, ok := result.(*DeleteShareNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_share: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListShares lists shares

  Allows a user to get a list of shares.
*/
func (a *Client) ListShares(params *ListSharesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSharesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSharesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list_shares",
		Method:             "GET",
		PathPattern:        "/shares",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSharesReader{formats: a.formats},
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
	success, ok := result.(*ListSharesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_shares: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}