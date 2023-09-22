# SMTPConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | True if outbound email is enabled. | [optional] 
**Server** | Pointer to **NullableString** | IP address or hostname of SMTP relay server. | [optional] 
**Port** | Pointer to **NullableInt32** | Port number to use. A value of -1 indicates the default (25 or 587 for TLS). | [optional] 
**AuthenticationEnabled** | Pointer to **bool** | True if username/password authentication should be used. | [optional] 
**TlsEnabled** | Pointer to **bool** | True if TLS (transport layer security) should be used. | [optional] 
**Username** | Pointer to **NullableString** | If authentication is enabled, username to use when authenticating to the server. | [optional] 
**Password** | Pointer to **NullableString** | If authentication is enabled, password to use when authenticating to the server. | [optional] 
**FromAddress** | Pointer to **NullableString** | From address to use when sending mail. If unspecified, &#39;noreply@delphix.com&#39; is used. | [optional] 
**SendTimeout** | Pointer to **NullableInt32** | Maximum timeout to wait, in seconds, when sending mail. | [optional] 

## Methods

### NewSMTPConfigParams

`func NewSMTPConfigParams() *SMTPConfigParams`

NewSMTPConfigParams instantiates a new SMTPConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPConfigParamsWithDefaults

`func NewSMTPConfigParamsWithDefaults() *SMTPConfigParams`

NewSMTPConfigParamsWithDefaults instantiates a new SMTPConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SMTPConfigParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SMTPConfigParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SMTPConfigParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SMTPConfigParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServer

`func (o *SMTPConfigParams) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *SMTPConfigParams) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *SMTPConfigParams) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *SMTPConfigParams) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *SMTPConfigParams) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *SMTPConfigParams) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetPort

`func (o *SMTPConfigParams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPConfigParams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPConfigParams) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPConfigParams) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SMTPConfigParams) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SMTPConfigParams) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetAuthenticationEnabled

`func (o *SMTPConfigParams) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *SMTPConfigParams) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *SMTPConfigParams) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *SMTPConfigParams) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *SMTPConfigParams) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *SMTPConfigParams) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *SMTPConfigParams) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *SMTPConfigParams) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetUsername

`func (o *SMTPConfigParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SMTPConfigParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SMTPConfigParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SMTPConfigParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SMTPConfigParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SMTPConfigParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *SMTPConfigParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SMTPConfigParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SMTPConfigParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SMTPConfigParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SMTPConfigParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SMTPConfigParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetFromAddress

`func (o *SMTPConfigParams) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *SMTPConfigParams) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *SMTPConfigParams) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *SMTPConfigParams) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### SetFromAddressNil

`func (o *SMTPConfigParams) SetFromAddressNil(b bool)`

 SetFromAddressNil sets the value for FromAddress to be an explicit nil

### UnsetFromAddress
`func (o *SMTPConfigParams) UnsetFromAddress()`

UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
### GetSendTimeout

`func (o *SMTPConfigParams) GetSendTimeout() int32`

GetSendTimeout returns the SendTimeout field if non-nil, zero value otherwise.

### GetSendTimeoutOk

`func (o *SMTPConfigParams) GetSendTimeoutOk() (*int32, bool)`

GetSendTimeoutOk returns a tuple with the SendTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTimeout

`func (o *SMTPConfigParams) SetSendTimeout(v int32)`

SetSendTimeout sets SendTimeout field to given value.

### HasSendTimeout

`func (o *SMTPConfigParams) HasSendTimeout() bool`

HasSendTimeout returns a boolean if a field has been set.

### SetSendTimeoutNil

`func (o *SMTPConfigParams) SetSendTimeoutNil(b bool)`

 SetSendTimeoutNil sets the value for SendTimeout to be an explicit nil

### UnsetSendTimeout
`func (o *SMTPConfigParams) UnsetSendTimeout()`

UnsetSendTimeout ensures that no value is present for SendTimeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


