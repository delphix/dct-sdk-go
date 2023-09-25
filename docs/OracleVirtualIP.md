# OracleVirtualIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | The domain name for the VirtualIP. | [optional] 
**Ip** | Pointer to **string** | The IP address for this VirtualIP. | [optional] 
**Discovered** | Pointer to **bool** | Whether this VirtualIP was discovered. | [optional] 

## Methods

### NewOracleVirtualIP

`func NewOracleVirtualIP() *OracleVirtualIP`

NewOracleVirtualIP instantiates a new OracleVirtualIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleVirtualIPWithDefaults

`func NewOracleVirtualIPWithDefaults() *OracleVirtualIP`

NewOracleVirtualIPWithDefaults instantiates a new OracleVirtualIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *OracleVirtualIP) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *OracleVirtualIP) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *OracleVirtualIP) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *OracleVirtualIP) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetIp

`func (o *OracleVirtualIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OracleVirtualIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OracleVirtualIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OracleVirtualIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDiscovered

`func (o *OracleVirtualIP) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *OracleVirtualIP) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *OracleVirtualIP) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *OracleVirtualIP) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


