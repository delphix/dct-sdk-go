# OracleListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of this listener. | [optional] 
**Name** | Pointer to **string** | Name of this listener. | [optional] 
**Type** | Pointer to **string** | Type of this listener. | [optional] 
**ProtocolAddresses** | Pointer to **[]string** | The list of protocol addresses for this listener. | [optional] 
**ClientEndpoints** | Pointer to **[]string** | The list of client endpoints for this listener. | [optional] 
**IsDiscovered** | Pointer to **bool** | Whether this listener was automatically discovered or not. | [optional] 
**HostId** | Pointer to **string** | Id to the host this listener is associated with. | [optional] 

## Methods

### NewOracleListener

`func NewOracleListener() *OracleListener`

NewOracleListener instantiates a new OracleListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleListenerWithDefaults

`func NewOracleListenerWithDefaults() *OracleListener`

NewOracleListenerWithDefaults instantiates a new OracleListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OracleListener) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OracleListener) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OracleListener) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OracleListener) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OracleListener) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OracleListener) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OracleListener) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OracleListener) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *OracleListener) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OracleListener) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OracleListener) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OracleListener) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProtocolAddresses

`func (o *OracleListener) GetProtocolAddresses() []string`

GetProtocolAddresses returns the ProtocolAddresses field if non-nil, zero value otherwise.

### GetProtocolAddressesOk

`func (o *OracleListener) GetProtocolAddressesOk() (*[]string, bool)`

GetProtocolAddressesOk returns a tuple with the ProtocolAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddresses

`func (o *OracleListener) SetProtocolAddresses(v []string)`

SetProtocolAddresses sets ProtocolAddresses field to given value.

### HasProtocolAddresses

`func (o *OracleListener) HasProtocolAddresses() bool`

HasProtocolAddresses returns a boolean if a field has been set.

### GetClientEndpoints

`func (o *OracleListener) GetClientEndpoints() []string`

GetClientEndpoints returns the ClientEndpoints field if non-nil, zero value otherwise.

### GetClientEndpointsOk

`func (o *OracleListener) GetClientEndpointsOk() (*[]string, bool)`

GetClientEndpointsOk returns a tuple with the ClientEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEndpoints

`func (o *OracleListener) SetClientEndpoints(v []string)`

SetClientEndpoints sets ClientEndpoints field to given value.

### HasClientEndpoints

`func (o *OracleListener) HasClientEndpoints() bool`

HasClientEndpoints returns a boolean if a field has been set.

### GetIsDiscovered

`func (o *OracleListener) GetIsDiscovered() bool`

GetIsDiscovered returns the IsDiscovered field if non-nil, zero value otherwise.

### GetIsDiscoveredOk

`func (o *OracleListener) GetIsDiscoveredOk() (*bool, bool)`

GetIsDiscoveredOk returns a tuple with the IsDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiscovered

`func (o *OracleListener) SetIsDiscovered(v bool)`

SetIsDiscovered sets IsDiscovered field to given value.

### HasIsDiscovered

`func (o *OracleListener) HasIsDiscovered() bool`

HasIsDiscovered returns a boolean if a field has been set.

### GetHostId

`func (o *OracleListener) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *OracleListener) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *OracleListener) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *OracleListener) HasHostId() bool`

HasHostId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


