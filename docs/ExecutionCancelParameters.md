# ExecutionCancelParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedStatus** | Pointer to **string** | The expected status of the execution to cancel to prevent cancelling a queued job that has transitioned to a running state since the request was issued. | [optional] 

## Methods

### NewExecutionCancelParameters

`func NewExecutionCancelParameters() *ExecutionCancelParameters`

NewExecutionCancelParameters instantiates a new ExecutionCancelParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionCancelParametersWithDefaults

`func NewExecutionCancelParametersWithDefaults() *ExecutionCancelParameters`

NewExecutionCancelParametersWithDefaults instantiates a new ExecutionCancelParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedStatus

`func (o *ExecutionCancelParameters) GetExpectedStatus() string`

GetExpectedStatus returns the ExpectedStatus field if non-nil, zero value otherwise.

### GetExpectedStatusOk

`func (o *ExecutionCancelParameters) GetExpectedStatusOk() (*string, bool)`

GetExpectedStatusOk returns a tuple with the ExpectedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStatus

`func (o *ExecutionCancelParameters) SetExpectedStatus(v string)`

SetExpectedStatus sets ExpectedStatus field to given value.

### HasExpectedStatus

`func (o *ExecutionCancelParameters) HasExpectedStatus() bool`

HasExpectedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


