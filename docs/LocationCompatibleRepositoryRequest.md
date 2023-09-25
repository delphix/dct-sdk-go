# LocationCompatibleRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to get the compatible repos. All other objects referenced by the parameters must live on the same engine as the source. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine from where to get the compatible repos. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**Location** | Pointer to **string** | location from where compatible repo to be fetched. | [optional] 
**TimeflowId** | Pointer to **string** | ID of the timeflow from which compatible repos need to be fetched. | [optional] 
**EnvironmentId** | Pointer to **string** | The ID or name of the target environment. | [optional] 

## Methods

### NewLocationCompatibleRepositoryRequest

`func NewLocationCompatibleRepositoryRequest() *LocationCompatibleRepositoryRequest`

NewLocationCompatibleRepositoryRequest instantiates a new LocationCompatibleRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationCompatibleRepositoryRequestWithDefaults

`func NewLocationCompatibleRepositoryRequestWithDefaults() *LocationCompatibleRepositoryRequest`

NewLocationCompatibleRepositoryRequestWithDefaults instantiates a new LocationCompatibleRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDataId

`func (o *LocationCompatibleRepositoryRequest) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *LocationCompatibleRepositoryRequest) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *LocationCompatibleRepositoryRequest) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *LocationCompatibleRepositoryRequest) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetEngineId

`func (o *LocationCompatibleRepositoryRequest) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *LocationCompatibleRepositoryRequest) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *LocationCompatibleRepositoryRequest) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *LocationCompatibleRepositoryRequest) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetLocation

`func (o *LocationCompatibleRepositoryRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationCompatibleRepositoryRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationCompatibleRepositoryRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LocationCompatibleRepositoryRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTimeflowId

`func (o *LocationCompatibleRepositoryRequest) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *LocationCompatibleRepositoryRequest) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *LocationCompatibleRepositoryRequest) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *LocationCompatibleRepositoryRequest) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *LocationCompatibleRepositoryRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *LocationCompatibleRepositoryRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *LocationCompatibleRepositoryRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *LocationCompatibleRepositoryRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


