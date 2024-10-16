/*
Delphix DCT API

Delphix DCT API

API version: 3.17.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ExecutionLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionLog{}

// ExecutionLog A log containing warnings or errors associated with a job execution.
type ExecutionLog struct {
	// The ExecutionLog entity ID.
	Id *string `json:"id,omitempty"`
	// The ID of the execution.
	ExecutionId *string `json:"execution_id,omitempty"`
	// The ID of the masking job that is being executed.
	MaskingJobId *string `json:"masking_job_id,omitempty"`
	Status *ExecutionStatus `json:"status,omitempty"`
	// The log file contents.
	Log *string `json:"log,omitempty"`
}

// NewExecutionLog instantiates a new ExecutionLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionLog() *ExecutionLog {
	this := ExecutionLog{}
	return &this
}

// NewExecutionLogWithDefaults instantiates a new ExecutionLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionLogWithDefaults() *ExecutionLog {
	this := ExecutionLog{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExecutionLog) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExecutionLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExecutionLog) SetId(v string) {
	o.Id = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *ExecutionLog) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *ExecutionLog) HasExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *ExecutionLog) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetMaskingJobId returns the MaskingJobId field value if set, zero value otherwise.
func (o *ExecutionLog) GetMaskingJobId() string {
	if o == nil || IsNil(o.MaskingJobId) {
		var ret string
		return ret
	}
	return *o.MaskingJobId
}

// GetMaskingJobIdOk returns a tuple with the MaskingJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetMaskingJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.MaskingJobId) {
		return nil, false
	}
	return o.MaskingJobId, true
}

// HasMaskingJobId returns a boolean if a field has been set.
func (o *ExecutionLog) HasMaskingJobId() bool {
	if o != nil && !IsNil(o.MaskingJobId) {
		return true
	}

	return false
}

// SetMaskingJobId gets a reference to the given string and assigns it to the MaskingJobId field.
func (o *ExecutionLog) SetMaskingJobId(v string) {
	o.MaskingJobId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExecutionLog) GetStatus() ExecutionStatus {
	if o == nil || IsNil(o.Status) {
		var ret ExecutionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetStatusOk() (*ExecutionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExecutionLog) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ExecutionStatus and assigns it to the Status field.
func (o *ExecutionLog) SetStatus(v ExecutionStatus) {
	o.Status = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *ExecutionLog) GetLog() string {
	if o == nil || IsNil(o.Log) {
		var ret string
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetLogOk() (*string, bool) {
	if o == nil || IsNil(o.Log) {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *ExecutionLog) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given string and assigns it to the Log field.
func (o *ExecutionLog) SetLog(v string) {
	o.Log = &v
}

func (o ExecutionLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ExecutionId) {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if !IsNil(o.MaskingJobId) {
		toSerialize["masking_job_id"] = o.MaskingJobId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	return toSerialize, nil
}

type NullableExecutionLog struct {
	value *ExecutionLog
	isSet bool
}

func (v NullableExecutionLog) Get() *ExecutionLog {
	return v.value
}

func (v *NullableExecutionLog) Set(val *ExecutionLog) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionLog) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionLog(val *ExecutionLog) *NullableExecutionLog {
	return &NullableExecutionLog{value: val, isSet: true}
}

func (v NullableExecutionLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


