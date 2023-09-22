# TimeToUpdateSourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnginesCount** | **int32** | no. of times same engine needs to be registered | 
**EnginesList** | **[]string** | list of engine hostnames to be registered engines_count times | 
**ExistingNoOfSources** | **int32** | no. of actual sources on 1 engine - this no. needs to be same for all engines | 
**SourcesCount** | **int32** | no. of times every source needs to be saved in data library | 

## Methods

### NewTimeToUpdateSourcesRequest

`func NewTimeToUpdateSourcesRequest(enginesCount int32, enginesList []string, existingNoOfSources int32, sourcesCount int32, ) *TimeToUpdateSourcesRequest`

NewTimeToUpdateSourcesRequest instantiates a new TimeToUpdateSourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeToUpdateSourcesRequestWithDefaults

`func NewTimeToUpdateSourcesRequestWithDefaults() *TimeToUpdateSourcesRequest`

NewTimeToUpdateSourcesRequestWithDefaults instantiates a new TimeToUpdateSourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnginesCount

`func (o *TimeToUpdateSourcesRequest) GetEnginesCount() int32`

GetEnginesCount returns the EnginesCount field if non-nil, zero value otherwise.

### GetEnginesCountOk

`func (o *TimeToUpdateSourcesRequest) GetEnginesCountOk() (*int32, bool)`

GetEnginesCountOk returns a tuple with the EnginesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnginesCount

`func (o *TimeToUpdateSourcesRequest) SetEnginesCount(v int32)`

SetEnginesCount sets EnginesCount field to given value.


### GetEnginesList

`func (o *TimeToUpdateSourcesRequest) GetEnginesList() []string`

GetEnginesList returns the EnginesList field if non-nil, zero value otherwise.

### GetEnginesListOk

`func (o *TimeToUpdateSourcesRequest) GetEnginesListOk() (*[]string, bool)`

GetEnginesListOk returns a tuple with the EnginesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnginesList

`func (o *TimeToUpdateSourcesRequest) SetEnginesList(v []string)`

SetEnginesList sets EnginesList field to given value.


### GetExistingNoOfSources

`func (o *TimeToUpdateSourcesRequest) GetExistingNoOfSources() int32`

GetExistingNoOfSources returns the ExistingNoOfSources field if non-nil, zero value otherwise.

### GetExistingNoOfSourcesOk

`func (o *TimeToUpdateSourcesRequest) GetExistingNoOfSourcesOk() (*int32, bool)`

GetExistingNoOfSourcesOk returns a tuple with the ExistingNoOfSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingNoOfSources

`func (o *TimeToUpdateSourcesRequest) SetExistingNoOfSources(v int32)`

SetExistingNoOfSources sets ExistingNoOfSources field to given value.


### GetSourcesCount

`func (o *TimeToUpdateSourcesRequest) GetSourcesCount() int32`

GetSourcesCount returns the SourcesCount field if non-nil, zero value otherwise.

### GetSourcesCountOk

`func (o *TimeToUpdateSourcesRequest) GetSourcesCountOk() (*int32, bool)`

GetSourcesCountOk returns a tuple with the SourcesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesCount

`func (o *TimeToUpdateSourcesRequest) SetSourcesCount(v int32)`

SetSourcesCount sets SourcesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


