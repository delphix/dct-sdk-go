# EnableScaleTestingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnginesCount** | **int32** | no. of times same engine needs to be registered | 
**EnginesList** | **[]string** | list of engine hostnames to be registered engines_count times | 
**VirtObjectsCount** | **int32** | no. of times to duplicate sources, containers, and timeflows | 
**SnapshotsCount** | **int32** | no. of times to duplicate snapshots | 

## Methods

### NewEnableScaleTestingRequest

`func NewEnableScaleTestingRequest(enginesCount int32, enginesList []string, virtObjectsCount int32, snapshotsCount int32, ) *EnableScaleTestingRequest`

NewEnableScaleTestingRequest instantiates a new EnableScaleTestingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableScaleTestingRequestWithDefaults

`func NewEnableScaleTestingRequestWithDefaults() *EnableScaleTestingRequest`

NewEnableScaleTestingRequestWithDefaults instantiates a new EnableScaleTestingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnginesCount

`func (o *EnableScaleTestingRequest) GetEnginesCount() int32`

GetEnginesCount returns the EnginesCount field if non-nil, zero value otherwise.

### GetEnginesCountOk

`func (o *EnableScaleTestingRequest) GetEnginesCountOk() (*int32, bool)`

GetEnginesCountOk returns a tuple with the EnginesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnginesCount

`func (o *EnableScaleTestingRequest) SetEnginesCount(v int32)`

SetEnginesCount sets EnginesCount field to given value.


### GetEnginesList

`func (o *EnableScaleTestingRequest) GetEnginesList() []string`

GetEnginesList returns the EnginesList field if non-nil, zero value otherwise.

### GetEnginesListOk

`func (o *EnableScaleTestingRequest) GetEnginesListOk() (*[]string, bool)`

GetEnginesListOk returns a tuple with the EnginesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnginesList

`func (o *EnableScaleTestingRequest) SetEnginesList(v []string)`

SetEnginesList sets EnginesList field to given value.


### GetVirtObjectsCount

`func (o *EnableScaleTestingRequest) GetVirtObjectsCount() int32`

GetVirtObjectsCount returns the VirtObjectsCount field if non-nil, zero value otherwise.

### GetVirtObjectsCountOk

`func (o *EnableScaleTestingRequest) GetVirtObjectsCountOk() (*int32, bool)`

GetVirtObjectsCountOk returns a tuple with the VirtObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtObjectsCount

`func (o *EnableScaleTestingRequest) SetVirtObjectsCount(v int32)`

SetVirtObjectsCount sets VirtObjectsCount field to given value.


### GetSnapshotsCount

`func (o *EnableScaleTestingRequest) GetSnapshotsCount() int32`

GetSnapshotsCount returns the SnapshotsCount field if non-nil, zero value otherwise.

### GetSnapshotsCountOk

`func (o *EnableScaleTestingRequest) GetSnapshotsCountOk() (*int32, bool)`

GetSnapshotsCountOk returns a tuple with the SnapshotsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotsCount

`func (o *EnableScaleTestingRequest) SetSnapshotsCount(v int32)`

SetSnapshotsCount sets SnapshotsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


