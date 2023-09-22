# ProvisionVDBByLocationDefaultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**Location** | Pointer to **string** | The location to get the defaults from. | [optional] 
**TimeflowId** | Pointer to **string** | ID of the timeflow to provision from. | [optional] 

## Methods

### NewProvisionVDBByLocationDefaultsRequest

`func NewProvisionVDBByLocationDefaultsRequest() *ProvisionVDBByLocationDefaultsRequest`

NewProvisionVDBByLocationDefaultsRequest instantiates a new ProvisionVDBByLocationDefaultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBByLocationDefaultsRequestWithDefaults

`func NewProvisionVDBByLocationDefaultsRequestWithDefaults() *ProvisionVDBByLocationDefaultsRequest`

NewProvisionVDBByLocationDefaultsRequestWithDefaults instantiates a new ProvisionVDBByLocationDefaultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDataId

`func (o *ProvisionVDBByLocationDefaultsRequest) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBByLocationDefaultsRequest) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBByLocationDefaultsRequest) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *ProvisionVDBByLocationDefaultsRequest) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetEngineId

`func (o *ProvisionVDBByLocationDefaultsRequest) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBByLocationDefaultsRequest) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBByLocationDefaultsRequest) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBByLocationDefaultsRequest) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetLocation

`func (o *ProvisionVDBByLocationDefaultsRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProvisionVDBByLocationDefaultsRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProvisionVDBByLocationDefaultsRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProvisionVDBByLocationDefaultsRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTimeflowId

`func (o *ProvisionVDBByLocationDefaultsRequest) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *ProvisionVDBByLocationDefaultsRequest) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *ProvisionVDBByLocationDefaultsRequest) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *ProvisionVDBByLocationDefaultsRequest) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


