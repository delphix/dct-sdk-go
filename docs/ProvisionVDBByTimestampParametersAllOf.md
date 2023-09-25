# ProvisionVDBByTimestampParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | Pointer to **string** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**SourceDataId** | **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. | 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account provisioning this VDB must be configured as owner of the VDB. | [optional] [default to true]

## Methods

### NewProvisionVDBByTimestampParametersAllOf

`func NewProvisionVDBByTimestampParametersAllOf(sourceDataId string, ) *ProvisionVDBByTimestampParametersAllOf`

NewProvisionVDBByTimestampParametersAllOf instantiates a new ProvisionVDBByTimestampParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBByTimestampParametersAllOfWithDefaults

`func NewProvisionVDBByTimestampParametersAllOfWithDefaults() *ProvisionVDBByTimestampParametersAllOf`

NewProvisionVDBByTimestampParametersAllOfWithDefaults instantiates a new ProvisionVDBByTimestampParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *ProvisionVDBByTimestampParametersAllOf) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBByTimestampParametersAllOf) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBByTimestampParametersAllOf) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBByTimestampParametersAllOf) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSourceDataId

`func (o *ProvisionVDBByTimestampParametersAllOf) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBByTimestampParametersAllOf) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBByTimestampParametersAllOf) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.


### GetMakeCurrentAccountOwner

`func (o *ProvisionVDBByTimestampParametersAllOf) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ProvisionVDBByTimestampParametersAllOf) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ProvisionVDBByTimestampParametersAllOf) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ProvisionVDBByTimestampParametersAllOf) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


