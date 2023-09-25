# ExecutionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ExecutionEvent entity ID. | [optional] 
**ExecutionId** | Pointer to **string** | The ID of the execution. | [optional] 
**EventType** | Pointer to **string** | The type of execution event. | [optional] 
**Severity** | Pointer to **string** | The severity of the execution event. | [optional] 
**Cause** | Pointer to **string** | The cause of the execution event. | [optional] 
**Count** | Pointer to **int64** | The number of times the execution event occurred. | [optional] 
**Timestamp** | Pointer to **time.Time** | The date and time that this execution event first occurred. | [optional] 
**MaskedObjectName** | Pointer to **string** | The name of the column, field, or other object being masked when this event occurred, if applicable. | [optional] 
**AlgorithmName** | Pointer to **string** | The name of the masking algorithm running when this event occurred, if applicable. | [optional] 
**ExceptionType** | Pointer to **string** | The Java class of the exception that triggered this event, if applicable. | [optional] 
**ExceptionDetail** | Pointer to **string** | The details associated with the Java exception that triggered this event, if applicable. | [optional] 

## Methods

### NewExecutionEvent

`func NewExecutionEvent() *ExecutionEvent`

NewExecutionEvent instantiates a new ExecutionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionEventWithDefaults

`func NewExecutionEventWithDefaults() *ExecutionEvent`

NewExecutionEventWithDefaults instantiates a new ExecutionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExecutionId

`func (o *ExecutionEvent) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExecutionEvent) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExecutionEvent) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ExecutionEvent) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetEventType

`func (o *ExecutionEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ExecutionEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ExecutionEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ExecutionEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSeverity

`func (o *ExecutionEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ExecutionEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ExecutionEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ExecutionEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCause

`func (o *ExecutionEvent) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ExecutionEvent) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ExecutionEvent) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ExecutionEvent) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCount

`func (o *ExecutionEvent) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ExecutionEvent) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ExecutionEvent) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ExecutionEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTimestamp

`func (o *ExecutionEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExecutionEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExecutionEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExecutionEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMaskedObjectName

`func (o *ExecutionEvent) GetMaskedObjectName() string`

GetMaskedObjectName returns the MaskedObjectName field if non-nil, zero value otherwise.

### GetMaskedObjectNameOk

`func (o *ExecutionEvent) GetMaskedObjectNameOk() (*string, bool)`

GetMaskedObjectNameOk returns a tuple with the MaskedObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedObjectName

`func (o *ExecutionEvent) SetMaskedObjectName(v string)`

SetMaskedObjectName sets MaskedObjectName field to given value.

### HasMaskedObjectName

`func (o *ExecutionEvent) HasMaskedObjectName() bool`

HasMaskedObjectName returns a boolean if a field has been set.

### GetAlgorithmName

`func (o *ExecutionEvent) GetAlgorithmName() string`

GetAlgorithmName returns the AlgorithmName field if non-nil, zero value otherwise.

### GetAlgorithmNameOk

`func (o *ExecutionEvent) GetAlgorithmNameOk() (*string, bool)`

GetAlgorithmNameOk returns a tuple with the AlgorithmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmName

`func (o *ExecutionEvent) SetAlgorithmName(v string)`

SetAlgorithmName sets AlgorithmName field to given value.

### HasAlgorithmName

`func (o *ExecutionEvent) HasAlgorithmName() bool`

HasAlgorithmName returns a boolean if a field has been set.

### GetExceptionType

`func (o *ExecutionEvent) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *ExecutionEvent) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *ExecutionEvent) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *ExecutionEvent) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetExceptionDetail

`func (o *ExecutionEvent) GetExceptionDetail() string`

GetExceptionDetail returns the ExceptionDetail field if non-nil, zero value otherwise.

### GetExceptionDetailOk

`func (o *ExecutionEvent) GetExceptionDetailOk() (*string, bool)`

GetExceptionDetailOk returns a tuple with the ExceptionDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionDetail

`func (o *ExecutionEvent) SetExceptionDetail(v string)`

SetExceptionDetail sets ExceptionDetail field to given value.

### HasExceptionDetail

`func (o *ExecutionEvent) HasExceptionDetail() bool`

HasExceptionDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


