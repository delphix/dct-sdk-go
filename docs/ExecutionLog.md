# ExecutionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ExecutionLog entity ID. | [optional] 
**ExecutionId** | Pointer to **string** | The ID of the execution. | [optional] 
**MaskingJobId** | Pointer to **string** | The ID of the masking job that is being executed. | [optional] 
**Status** | Pointer to **string** | The status of the execution regarding its completion. | [optional] 
**Log** | Pointer to **string** | The log file contents. | [optional] 

## Methods

### NewExecutionLog

`func NewExecutionLog() *ExecutionLog`

NewExecutionLog instantiates a new ExecutionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionLogWithDefaults

`func NewExecutionLogWithDefaults() *ExecutionLog`

NewExecutionLogWithDefaults instantiates a new ExecutionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExecutionId

`func (o *ExecutionLog) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExecutionLog) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExecutionLog) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ExecutionLog) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetMaskingJobId

`func (o *ExecutionLog) GetMaskingJobId() string`

GetMaskingJobId returns the MaskingJobId field if non-nil, zero value otherwise.

### GetMaskingJobIdOk

`func (o *ExecutionLog) GetMaskingJobIdOk() (*string, bool)`

GetMaskingJobIdOk returns a tuple with the MaskingJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingJobId

`func (o *ExecutionLog) SetMaskingJobId(v string)`

SetMaskingJobId sets MaskingJobId field to given value.

### HasMaskingJobId

`func (o *ExecutionLog) HasMaskingJobId() bool`

HasMaskingJobId returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLog

`func (o *ExecutionLog) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ExecutionLog) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ExecutionLog) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *ExecutionLog) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


