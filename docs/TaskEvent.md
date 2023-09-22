# TaskEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The steps or events a task will perform. | [optional] 
**Status** | Pointer to **string** | The state of result of the task event. | [optional] 

## Methods

### NewTaskEvent

`func NewTaskEvent() *TaskEvent`

NewTaskEvent instantiates a new TaskEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskEventWithDefaults

`func NewTaskEventWithDefaults() *TaskEvent`

NewTaskEventWithDefaults instantiates a new TaskEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TaskEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TaskEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TaskEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *TaskEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetStatus

`func (o *TaskEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


