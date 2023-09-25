# HyperscaleTaskEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the task (Unload, Masking, Load, Post-Load) | [optional] 
**Progress** | Pointer to **float32** | progress of the task (between 0 and 1) | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ProcessedObjects** | Pointer to **int64** | The number of objects (tables) already processed by this task. | [optional] 
**ProcessedRows** | Pointer to **int64** | The number of rows already processed by this task. | [optional] 
**ProcessedConstraints** | Pointer to **int64** | The number of constraints processed by this task (Post-load task only) | [optional] 
**TotalConstraints** | Pointer to **int64** | The total number of constraints to be processed by this task (Post-load task only) | [optional] 
**ProcessedIndexes** | Pointer to **int64** | The number of indexes processed by this task (Post-load task only) | [optional] 
**TotalIndexes** | Pointer to **int64** | The total number of indexes to be processed by this task (Post-load task only) | [optional] 
**ProcessedTriggers** | Pointer to **int64** | The number of triggered processed by this task (Post-load task only) | [optional] 
**TotalTriggers** | Pointer to **int64** | The total number of triggers to be processed by this task (Post-load task only) | [optional] 
**StartTime** | Pointer to **time.Time** | The date and time that this task was started. | [optional] 
**EndTime** | Pointer to **time.Time** | The date and time that this task completed. | [optional] 
**Errors** | Pointer to [**[]HyperscaleTaskError**](HyperscaleTaskError.md) |  | [optional] 

## Methods

### NewHyperscaleTaskEvent

`func NewHyperscaleTaskEvent() *HyperscaleTaskEvent`

NewHyperscaleTaskEvent instantiates a new HyperscaleTaskEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleTaskEventWithDefaults

`func NewHyperscaleTaskEventWithDefaults() *HyperscaleTaskEvent`

NewHyperscaleTaskEventWithDefaults instantiates a new HyperscaleTaskEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HyperscaleTaskEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperscaleTaskEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperscaleTaskEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperscaleTaskEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProgress

`func (o *HyperscaleTaskEvent) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *HyperscaleTaskEvent) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *HyperscaleTaskEvent) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *HyperscaleTaskEvent) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *HyperscaleTaskEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperscaleTaskEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperscaleTaskEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperscaleTaskEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProcessedObjects

`func (o *HyperscaleTaskEvent) GetProcessedObjects() int64`

GetProcessedObjects returns the ProcessedObjects field if non-nil, zero value otherwise.

### GetProcessedObjectsOk

`func (o *HyperscaleTaskEvent) GetProcessedObjectsOk() (*int64, bool)`

GetProcessedObjectsOk returns a tuple with the ProcessedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedObjects

`func (o *HyperscaleTaskEvent) SetProcessedObjects(v int64)`

SetProcessedObjects sets ProcessedObjects field to given value.

### HasProcessedObjects

`func (o *HyperscaleTaskEvent) HasProcessedObjects() bool`

HasProcessedObjects returns a boolean if a field has been set.

### GetProcessedRows

`func (o *HyperscaleTaskEvent) GetProcessedRows() int64`

GetProcessedRows returns the ProcessedRows field if non-nil, zero value otherwise.

### GetProcessedRowsOk

`func (o *HyperscaleTaskEvent) GetProcessedRowsOk() (*int64, bool)`

GetProcessedRowsOk returns a tuple with the ProcessedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRows

`func (o *HyperscaleTaskEvent) SetProcessedRows(v int64)`

SetProcessedRows sets ProcessedRows field to given value.

### HasProcessedRows

`func (o *HyperscaleTaskEvent) HasProcessedRows() bool`

HasProcessedRows returns a boolean if a field has been set.

### GetProcessedConstraints

`func (o *HyperscaleTaskEvent) GetProcessedConstraints() int64`

GetProcessedConstraints returns the ProcessedConstraints field if non-nil, zero value otherwise.

### GetProcessedConstraintsOk

`func (o *HyperscaleTaskEvent) GetProcessedConstraintsOk() (*int64, bool)`

GetProcessedConstraintsOk returns a tuple with the ProcessedConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedConstraints

`func (o *HyperscaleTaskEvent) SetProcessedConstraints(v int64)`

SetProcessedConstraints sets ProcessedConstraints field to given value.

### HasProcessedConstraints

`func (o *HyperscaleTaskEvent) HasProcessedConstraints() bool`

HasProcessedConstraints returns a boolean if a field has been set.

### GetTotalConstraints

`func (o *HyperscaleTaskEvent) GetTotalConstraints() int64`

GetTotalConstraints returns the TotalConstraints field if non-nil, zero value otherwise.

### GetTotalConstraintsOk

`func (o *HyperscaleTaskEvent) GetTotalConstraintsOk() (*int64, bool)`

GetTotalConstraintsOk returns a tuple with the TotalConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConstraints

`func (o *HyperscaleTaskEvent) SetTotalConstraints(v int64)`

SetTotalConstraints sets TotalConstraints field to given value.

### HasTotalConstraints

`func (o *HyperscaleTaskEvent) HasTotalConstraints() bool`

HasTotalConstraints returns a boolean if a field has been set.

### GetProcessedIndexes

`func (o *HyperscaleTaskEvent) GetProcessedIndexes() int64`

GetProcessedIndexes returns the ProcessedIndexes field if non-nil, zero value otherwise.

### GetProcessedIndexesOk

`func (o *HyperscaleTaskEvent) GetProcessedIndexesOk() (*int64, bool)`

GetProcessedIndexesOk returns a tuple with the ProcessedIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedIndexes

`func (o *HyperscaleTaskEvent) SetProcessedIndexes(v int64)`

SetProcessedIndexes sets ProcessedIndexes field to given value.

### HasProcessedIndexes

`func (o *HyperscaleTaskEvent) HasProcessedIndexes() bool`

HasProcessedIndexes returns a boolean if a field has been set.

### GetTotalIndexes

`func (o *HyperscaleTaskEvent) GetTotalIndexes() int64`

GetTotalIndexes returns the TotalIndexes field if non-nil, zero value otherwise.

### GetTotalIndexesOk

`func (o *HyperscaleTaskEvent) GetTotalIndexesOk() (*int64, bool)`

GetTotalIndexesOk returns a tuple with the TotalIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIndexes

`func (o *HyperscaleTaskEvent) SetTotalIndexes(v int64)`

SetTotalIndexes sets TotalIndexes field to given value.

### HasTotalIndexes

`func (o *HyperscaleTaskEvent) HasTotalIndexes() bool`

HasTotalIndexes returns a boolean if a field has been set.

### GetProcessedTriggers

`func (o *HyperscaleTaskEvent) GetProcessedTriggers() int64`

GetProcessedTriggers returns the ProcessedTriggers field if non-nil, zero value otherwise.

### GetProcessedTriggersOk

`func (o *HyperscaleTaskEvent) GetProcessedTriggersOk() (*int64, bool)`

GetProcessedTriggersOk returns a tuple with the ProcessedTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTriggers

`func (o *HyperscaleTaskEvent) SetProcessedTriggers(v int64)`

SetProcessedTriggers sets ProcessedTriggers field to given value.

### HasProcessedTriggers

`func (o *HyperscaleTaskEvent) HasProcessedTriggers() bool`

HasProcessedTriggers returns a boolean if a field has been set.

### GetTotalTriggers

`func (o *HyperscaleTaskEvent) GetTotalTriggers() int64`

GetTotalTriggers returns the TotalTriggers field if non-nil, zero value otherwise.

### GetTotalTriggersOk

`func (o *HyperscaleTaskEvent) GetTotalTriggersOk() (*int64, bool)`

GetTotalTriggersOk returns a tuple with the TotalTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTriggers

`func (o *HyperscaleTaskEvent) SetTotalTriggers(v int64)`

SetTotalTriggers sets TotalTriggers field to given value.

### HasTotalTriggers

`func (o *HyperscaleTaskEvent) HasTotalTriggers() bool`

HasTotalTriggers returns a boolean if a field has been set.

### GetStartTime

`func (o *HyperscaleTaskEvent) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HyperscaleTaskEvent) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HyperscaleTaskEvent) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HyperscaleTaskEvent) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *HyperscaleTaskEvent) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HyperscaleTaskEvent) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HyperscaleTaskEvent) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HyperscaleTaskEvent) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetErrors

`func (o *HyperscaleTaskEvent) GetErrors() []HyperscaleTaskError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HyperscaleTaskEvent) GetErrorsOk() (*[]HyperscaleTaskError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HyperscaleTaskEvent) SetErrors(v []HyperscaleTaskError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HyperscaleTaskEvent) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


