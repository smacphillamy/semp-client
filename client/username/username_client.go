// Code generated by go-swagger; DO NOT EDIT.

package username

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new username API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for username API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateUsername creates a username object

Create a Username object. Any attribute missing from the request will be set to its default value.


Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
password||||x|
userName|x|x|||



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation. Requests which include the following attributes require greater access scope/level:


Attribute|Access Scope/Level
:---|:---:
globalAccessLevel|global/admin
msgVpnDefaultAccessLevel|global/read-write



This has been available since 2.12.
*/
func (a *Client) CreateUsername(params *CreateUsernameParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUsernameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUsernameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUsername",
		Method:             "POST",
		PathPattern:        "/usernames",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUsernameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUsernameOK), nil

}

/*
CreateUsernameMsgVpnAccessLevelException creates a message v p n access level exception object

Create a Message VPN Access Level Exception object. Any attribute missing from the request will be set to its default value.


Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
msgVpnName|x|x|||
userName|x||x||



A SEMP client authorized with a minimum access scope/level of "global/read-write" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) CreateUsernameMsgVpnAccessLevelException(params *CreateUsernameMsgVpnAccessLevelExceptionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUsernameMsgVpnAccessLevelExceptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUsernameMsgVpnAccessLevelExceptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUsernameMsgVpnAccessLevelException",
		Method:             "POST",
		PathPattern:        "/usernames/{userName}/msgVpnAccessLevelExceptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUsernameMsgVpnAccessLevelExceptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUsernameMsgVpnAccessLevelExceptionOK), nil

}

/*
DeleteUsername deletes a username object

Delete a Username object.

A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) DeleteUsername(params *DeleteUsernameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsernameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsernameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUsername",
		Method:             "DELETE",
		PathPattern:        "/usernames/{userName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUsernameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUsernameOK), nil

}

/*
DeleteUsernameMsgVpnAccessLevelException deletes a message v p n access level exception object

Delete a Message VPN Access Level Exception object.

A SEMP client authorized with a minimum access scope/level of "global/read-write" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) DeleteUsernameMsgVpnAccessLevelException(params *DeleteUsernameMsgVpnAccessLevelExceptionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsernameMsgVpnAccessLevelExceptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsernameMsgVpnAccessLevelExceptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUsernameMsgVpnAccessLevelException",
		Method:             "DELETE",
		PathPattern:        "/usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUsernameMsgVpnAccessLevelExceptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUsernameMsgVpnAccessLevelExceptionOK), nil

}

/*
GetUsername gets a username object

Get a Username object.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
password||x|
userName|x||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) GetUsername(params *GetUsernameParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsernameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsernameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsername",
		Method:             "GET",
		PathPattern:        "/usernames/{userName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUsernameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsernameOK), nil

}

/*
GetUsernameMsgVpnAccessLevelException gets a message v p n access level exception object

Get a Message VPN Access Level Exception object.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
msgVpnName|x||
userName|x||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) GetUsernameMsgVpnAccessLevelException(params *GetUsernameMsgVpnAccessLevelExceptionParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsernameMsgVpnAccessLevelExceptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsernameMsgVpnAccessLevelExceptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsernameMsgVpnAccessLevelException",
		Method:             "GET",
		PathPattern:        "/usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUsernameMsgVpnAccessLevelExceptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsernameMsgVpnAccessLevelExceptionOK), nil

}

/*
GetUsernameMsgVpnAccessLevelExceptions gets a list of message v p n access level exception objects

Get a list of Message VPN Access Level Exception objects.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
msgVpnName|x||
userName|x||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) GetUsernameMsgVpnAccessLevelExceptions(params *GetUsernameMsgVpnAccessLevelExceptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsernameMsgVpnAccessLevelExceptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsernameMsgVpnAccessLevelExceptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsernameMsgVpnAccessLevelExceptions",
		Method:             "GET",
		PathPattern:        "/usernames/{userName}/msgVpnAccessLevelExceptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUsernameMsgVpnAccessLevelExceptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsernameMsgVpnAccessLevelExceptionsOK), nil

}

/*
GetUsernames gets a list of username objects

Get a list of Username objects.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
password||x|
userName|x||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) GetUsernames(params *GetUsernamesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsernamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsernamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsernames",
		Method:             "GET",
		PathPattern:        "/usernames",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUsernamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsernamesOK), nil

}

/*
ReplaceUsername replaces a username object

Replace a Username object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
password|||x||
userName|x|x|||



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation. Requests which include the following attributes require greater access scope/level:


Attribute|Access Scope/Level
:---|:---:
globalAccessLevel|global/admin
msgVpnDefaultAccessLevel|global/read-write



This has been available since 2.12.
*/
func (a *Client) ReplaceUsername(params *ReplaceUsernameParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceUsernameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceUsernameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceUsername",
		Method:             "PUT",
		PathPattern:        "/usernames/{userName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceUsernameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceUsernameOK), nil

}

/*
ReplaceUsernameMsgVpnAccessLevelException replaces a message v p n access level exception object

Replace a Message VPN Access Level Exception object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
msgVpnName|x|x|||
userName|x|x|||



A SEMP client authorized with a minimum access scope/level of "global/read-write" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) ReplaceUsernameMsgVpnAccessLevelException(params *ReplaceUsernameMsgVpnAccessLevelExceptionParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceUsernameMsgVpnAccessLevelExceptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceUsernameMsgVpnAccessLevelExceptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceUsernameMsgVpnAccessLevelException",
		Method:             "PUT",
		PathPattern:        "/usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceUsernameMsgVpnAccessLevelExceptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceUsernameMsgVpnAccessLevelExceptionOK), nil

}

/*
UpdateUsername updates a username object

Update a Username object. Any attribute missing from the request will be left unchanged.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
password|||x||
userName|x|x|||



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation. Requests which include the following attributes require greater access scope/level:


Attribute|Access Scope/Level
:---|:---:
globalAccessLevel|global/admin
msgVpnDefaultAccessLevel|global/read-write



This has been available since 2.12.
*/
func (a *Client) UpdateUsername(params *UpdateUsernameParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUsernameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUsernameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUsername",
		Method:             "PATCH",
		PathPattern:        "/usernames/{userName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUsernameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUsernameOK), nil

}

/*
UpdateUsernameMsgVpnAccessLevelException updates a message v p n access level exception object

Update a Message VPN Access Level Exception object. Any attribute missing from the request will be left unchanged.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
msgVpnName|x|x|||
userName|x|x|||



A SEMP client authorized with a minimum access scope/level of "global/read-write" is required to perform this operation.

This has been available since 2.12.
*/
func (a *Client) UpdateUsernameMsgVpnAccessLevelException(params *UpdateUsernameMsgVpnAccessLevelExceptionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUsernameMsgVpnAccessLevelExceptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUsernameMsgVpnAccessLevelExceptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUsernameMsgVpnAccessLevelException",
		Method:             "PATCH",
		PathPattern:        "/usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUsernameMsgVpnAccessLevelExceptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUsernameMsgVpnAccessLevelExceptionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
