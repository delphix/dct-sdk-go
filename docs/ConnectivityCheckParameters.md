# ConnectivityCheckParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | **string** | The ID of the engine to check. | 
**Host** | **string** | The hostname of the remote host machine to check. | 
**Port** | **NullableInt32** | The port of the remote host machine to check. | 
**Username** | Pointer to **string** | The username of the remote host machine to check. | [optional] 
**Password** | Pointer to **string** | The password of the remote host machine to check. | [optional] 

## Methods

### NewConnectivityCheckParameters

`func NewConnectivityCheckParameters(engineId string, host string, port NullableInt32, ) *ConnectivityCheckParameters`

NewConnectivityCheckParameters instantiates a new ConnectivityCheckParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityCheckParametersWithDefaults

`func NewConnectivityCheckParametersWithDefaults() *ConnectivityCheckParameters`

NewConnectivityCheckParametersWithDefaults instantiates a new ConnectivityCheckParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *ConnectivityCheckParameters) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ConnectivityCheckParameters) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ConnectivityCheckParameters) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetHost

`func (o *ConnectivityCheckParameters) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ConnectivityCheckParameters) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ConnectivityCheckParameters) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ConnectivityCheckParameters) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnectivityCheckParameters) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnectivityCheckParameters) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *ConnectivityCheckParameters) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ConnectivityCheckParameters) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *ConnectivityCheckParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConnectivityCheckParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConnectivityCheckParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConnectivityCheckParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ConnectivityCheckParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConnectivityCheckParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConnectivityCheckParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConnectivityCheckParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


