# ConnectorUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Connector name. | [optional] 
**Hostname** | Pointer to **string** | The network hostname or IP address of the database server. | [optional] 
**Port** | Pointer to **int32** | The TCP port of the server. | [optional] 
**Username** | Pointer to **string** | The username this Connector will use to connect to the database. | [optional] 
**Password** | Pointer to **string** | The password this Connector will use to connect to the database. | [optional] 

## Methods

### NewConnectorUpdateParameters

`func NewConnectorUpdateParameters() *ConnectorUpdateParameters`

NewConnectorUpdateParameters instantiates a new ConnectorUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorUpdateParametersWithDefaults

`func NewConnectorUpdateParametersWithDefaults() *ConnectorUpdateParameters`

NewConnectorUpdateParametersWithDefaults instantiates a new ConnectorUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorUpdateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorUpdateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorUpdateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorUpdateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *ConnectorUpdateParameters) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ConnectorUpdateParameters) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ConnectorUpdateParameters) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ConnectorUpdateParameters) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *ConnectorUpdateParameters) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnectorUpdateParameters) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnectorUpdateParameters) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ConnectorUpdateParameters) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *ConnectorUpdateParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConnectorUpdateParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConnectorUpdateParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConnectorUpdateParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ConnectorUpdateParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConnectorUpdateParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConnectorUpdateParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConnectorUpdateParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


