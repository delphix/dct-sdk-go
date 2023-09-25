# KDCServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | One of more KDC servers. | [optional] 
**Port** | Pointer to **int32** | One of more KDC servers. | [optional] 

## Methods

### NewKDCServers

`func NewKDCServers() *KDCServers`

NewKDCServers instantiates a new KDCServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKDCServersWithDefaults

`func NewKDCServersWithDefaults() *KDCServers`

NewKDCServersWithDefaults instantiates a new KDCServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *KDCServers) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *KDCServers) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *KDCServers) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *KDCServers) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *KDCServers) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KDCServers) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KDCServers) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KDCServers) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


