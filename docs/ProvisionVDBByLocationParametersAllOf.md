# ProvisionVDBByLocationParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | Pointer to **string** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account provisioning this VDB must be configured as owner of the VDB. | [optional] [default to true]

## Methods

### NewProvisionVDBByLocationParametersAllOf

`func NewProvisionVDBByLocationParametersAllOf() *ProvisionVDBByLocationParametersAllOf`

NewProvisionVDBByLocationParametersAllOf instantiates a new ProvisionVDBByLocationParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBByLocationParametersAllOfWithDefaults

`func NewProvisionVDBByLocationParametersAllOfWithDefaults() *ProvisionVDBByLocationParametersAllOf`

NewProvisionVDBByLocationParametersAllOfWithDefaults instantiates a new ProvisionVDBByLocationParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *ProvisionVDBByLocationParametersAllOf) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBByLocationParametersAllOf) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBByLocationParametersAllOf) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBByLocationParametersAllOf) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSourceDataId

`func (o *ProvisionVDBByLocationParametersAllOf) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBByLocationParametersAllOf) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBByLocationParametersAllOf) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *ProvisionVDBByLocationParametersAllOf) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *ProvisionVDBByLocationParametersAllOf) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ProvisionVDBByLocationParametersAllOf) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ProvisionVDBByLocationParametersAllOf) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ProvisionVDBByLocationParametersAllOf) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


