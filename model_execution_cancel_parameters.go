/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ExecutionCancelParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionCancelParameters{}

// ExecutionCancelParameters Parameters to cancel an execution.
type ExecutionCancelParameters struct {
	// The expected status of the execution to cancel to prevent cancelling a queued job that has transitioned to a running state since the request was issued.
	ExpectedStatus *string `json:"expected_status,omitempty"`
}

// NewExecutionCancelParameters instantiates a new ExecutionCancelParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionCancelParameters() *ExecutionCancelParameters {
	this := ExecutionCancelParameters{}
	return &this
}

// NewExecutionCancelParametersWithDefaults instantiates a new ExecutionCancelParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionCancelParametersWithDefaults() *ExecutionCancelParameters {
	this := ExecutionCancelParameters{}
	return &this
}

// GetExpectedStatus returns the ExpectedStatus field value if set, zero value otherwise.
func (o *ExecutionCancelParameters) GetExpectedStatus() string {
	if o == nil || IsNil(o.ExpectedStatus) {
		var ret string
		return ret
	}
	return *o.ExpectedStatus
}

// GetExpectedStatusOk returns a tuple with the ExpectedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionCancelParameters) GetExpectedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedStatus) {
		return nil, false
	}
	return o.ExpectedStatus, true
}

// HasExpectedStatus returns a boolean if a field has been set.
func (o *ExecutionCancelParameters) HasExpectedStatus() bool {
	if o != nil && !IsNil(o.ExpectedStatus) {
		return true
	}

	return false
}

// SetExpectedStatus gets a reference to the given string and assigns it to the ExpectedStatus field.
func (o *ExecutionCancelParameters) SetExpectedStatus(v string) {
	o.ExpectedStatus = &v
}

func (o ExecutionCancelParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionCancelParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpectedStatus) {
		toSerialize["expected_status"] = o.ExpectedStatus
	}
	return toSerialize, nil
}

type NullableExecutionCancelParameters struct {
	value *ExecutionCancelParameters
	isSet bool
}

func (v NullableExecutionCancelParameters) Get() *ExecutionCancelParameters {
	return v.value
}

func (v *NullableExecutionCancelParameters) Set(val *ExecutionCancelParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionCancelParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionCancelParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionCancelParameters(val *ExecutionCancelParameters) *NullableExecutionCancelParameters {
	return &NullableExecutionCancelParameters{value: val, isSet: true}
}

func (v NullableExecutionCancelParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionCancelParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


