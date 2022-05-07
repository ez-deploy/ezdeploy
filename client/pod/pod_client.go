// Code generated by go-swagger; DO NOT EDIT.

package pod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pod API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pod API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CheckPodTicket(params *CheckPodTicketParams, opts ...ClientOption) (*CheckPodTicketOK, error)

	CreatePodTicket(params *CreatePodTicketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePodTicketCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CheckPodTicket Check Pod Ticket, and delete it
*/
func (a *Client) CheckPodTicket(params *CheckPodTicketParams, opts ...ClientOption) (*CheckPodTicketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckPodTicketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CheckPodTicket",
		Method:             "GET",
		PathPattern:        "/visit/pod/ticket/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckPodTicketReader{formats: a.formats},
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
	success, ok := result.(*CheckPodTicketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CheckPodTicket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreatePodTicket Create Visit Pod Once Ticket
*/
func (a *Client) CreatePodTicket(params *CreatePodTicketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePodTicketCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePodTicketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreatePodTicket",
		Method:             "POST",
		PathPattern:        "/visit/pod/ticket/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePodTicketReader{formats: a.formats},
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
	success, ok := result.(*CreatePodTicketCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreatePodTicket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
