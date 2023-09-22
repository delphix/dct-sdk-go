# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Execution entity ID. | [optional] 
**EngineId** | Pointer to **string** | The ID of the engine where this execution ran. | [optional] 
**EngineName** | Pointer to **string** | The name of the engine where this execution ran. | [optional] 
**MaskingJobId** | Pointer to **string** | The ID of the masking job that is being executed. | [optional] 
**SourceConnectorId** | Pointer to **string** | The ID of the source connector. This field is only used for multi-tenant jobs that are also on-the-fly. | [optional] 
**TargetConnectorId** | Pointer to **string** | The ID of the target connector. This field is only used for multi-tenant jobs. | [optional] 
**Status** | Pointer to **string** | The status of the execution regarding its completion. | [optional] 
**RowsMasked** | Pointer to **int64** | The number of rows masked or profiled so far by this execution. This is not applicable for JSON file type. | [optional] 
**RowsTotal** | Pointer to **int64** | The total number of rows that this execution should mask. This value is set to -1 while the total row count is being calculated. This is not applicable for JSON file type. | [optional] 
**BytesProcessed** | Pointer to **int64** | The number of bytes masked so far by this execution. This is only applicable for JSON file type. | [optional] 
**BytesTotal** | Pointer to **int64** | The total number of bytes that this execution should mask. This value is set to -1 while the total byte count is being calculated. This is only applicable for JSON file type. | [optional] 
**StartTime** | Pointer to **time.Time** | The date and time that this execution was started. | [optional] 
**SubmitTime** | Pointer to **time.Time** | The date and time that this execution was submitted. | [optional] 
**EndTime** | Pointer to **time.Time** | The date and time that this execution completed. | [optional] 
**TaskEvents** | Pointer to [**[]TaskEvent**](TaskEvent.md) | The progression of steps or events performed by this execution. Only available for executions on masking engines that are version 6.0.14.0 and higher. | [optional] 

## Methods

### NewExecution

`func NewExecution() *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Execution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Execution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Execution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Execution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEngineId

`func (o *Execution) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Execution) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Execution) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *Execution) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEngineName

`func (o *Execution) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *Execution) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *Execution) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *Execution) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetMaskingJobId

`func (o *Execution) GetMaskingJobId() string`

GetMaskingJobId returns the MaskingJobId field if non-nil, zero value otherwise.

### GetMaskingJobIdOk

`func (o *Execution) GetMaskingJobIdOk() (*string, bool)`

GetMaskingJobIdOk returns a tuple with the MaskingJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingJobId

`func (o *Execution) SetMaskingJobId(v string)`

SetMaskingJobId sets MaskingJobId field to given value.

### HasMaskingJobId

`func (o *Execution) HasMaskingJobId() bool`

HasMaskingJobId returns a boolean if a field has been set.

### GetSourceConnectorId

`func (o *Execution) GetSourceConnectorId() string`

GetSourceConnectorId returns the SourceConnectorId field if non-nil, zero value otherwise.

### GetSourceConnectorIdOk

`func (o *Execution) GetSourceConnectorIdOk() (*string, bool)`

GetSourceConnectorIdOk returns a tuple with the SourceConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnectorId

`func (o *Execution) SetSourceConnectorId(v string)`

SetSourceConnectorId sets SourceConnectorId field to given value.

### HasSourceConnectorId

`func (o *Execution) HasSourceConnectorId() bool`

HasSourceConnectorId returns a boolean if a field has been set.

### GetTargetConnectorId

`func (o *Execution) GetTargetConnectorId() string`

GetTargetConnectorId returns the TargetConnectorId field if non-nil, zero value otherwise.

### GetTargetConnectorIdOk

`func (o *Execution) GetTargetConnectorIdOk() (*string, bool)`

GetTargetConnectorIdOk returns a tuple with the TargetConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConnectorId

`func (o *Execution) SetTargetConnectorId(v string)`

SetTargetConnectorId sets TargetConnectorId field to given value.

### HasTargetConnectorId

`func (o *Execution) HasTargetConnectorId() bool`

HasTargetConnectorId returns a boolean if a field has been set.

### GetStatus

`func (o *Execution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Execution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Execution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Execution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRowsMasked

`func (o *Execution) GetRowsMasked() int64`

GetRowsMasked returns the RowsMasked field if non-nil, zero value otherwise.

### GetRowsMaskedOk

`func (o *Execution) GetRowsMaskedOk() (*int64, bool)`

GetRowsMaskedOk returns a tuple with the RowsMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsMasked

`func (o *Execution) SetRowsMasked(v int64)`

SetRowsMasked sets RowsMasked field to given value.

### HasRowsMasked

`func (o *Execution) HasRowsMasked() bool`

HasRowsMasked returns a boolean if a field has been set.

### GetRowsTotal

`func (o *Execution) GetRowsTotal() int64`

GetRowsTotal returns the RowsTotal field if non-nil, zero value otherwise.

### GetRowsTotalOk

`func (o *Execution) GetRowsTotalOk() (*int64, bool)`

GetRowsTotalOk returns a tuple with the RowsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsTotal

`func (o *Execution) SetRowsTotal(v int64)`

SetRowsTotal sets RowsTotal field to given value.

### HasRowsTotal

`func (o *Execution) HasRowsTotal() bool`

HasRowsTotal returns a boolean if a field has been set.

### GetBytesProcessed

`func (o *Execution) GetBytesProcessed() int64`

GetBytesProcessed returns the BytesProcessed field if non-nil, zero value otherwise.

### GetBytesProcessedOk

`func (o *Execution) GetBytesProcessedOk() (*int64, bool)`

GetBytesProcessedOk returns a tuple with the BytesProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesProcessed

`func (o *Execution) SetBytesProcessed(v int64)`

SetBytesProcessed sets BytesProcessed field to given value.

### HasBytesProcessed

`func (o *Execution) HasBytesProcessed() bool`

HasBytesProcessed returns a boolean if a field has been set.

### GetBytesTotal

`func (o *Execution) GetBytesTotal() int64`

GetBytesTotal returns the BytesTotal field if non-nil, zero value otherwise.

### GetBytesTotalOk

`func (o *Execution) GetBytesTotalOk() (*int64, bool)`

GetBytesTotalOk returns a tuple with the BytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTotal

`func (o *Execution) SetBytesTotal(v int64)`

SetBytesTotal sets BytesTotal field to given value.

### HasBytesTotal

`func (o *Execution) HasBytesTotal() bool`

HasBytesTotal returns a boolean if a field has been set.

### GetStartTime

`func (o *Execution) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Execution) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Execution) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Execution) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubmitTime

`func (o *Execution) GetSubmitTime() time.Time`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *Execution) GetSubmitTimeOk() (*time.Time, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *Execution) SetSubmitTime(v time.Time)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *Execution) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Execution) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Execution) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Execution) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Execution) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetTaskEvents

`func (o *Execution) GetTaskEvents() []TaskEvent`

GetTaskEvents returns the TaskEvents field if non-nil, zero value otherwise.

### GetTaskEventsOk

`func (o *Execution) GetTaskEventsOk() (*[]TaskEvent, bool)`

GetTaskEventsOk returns a tuple with the TaskEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskEvents

`func (o *Execution) SetTaskEvents(v []TaskEvent)`

SetTaskEvents sets TaskEvents field to given value.

### HasTaskEvents

`func (o *Execution) HasTaskEvents() bool`

HasTaskEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


