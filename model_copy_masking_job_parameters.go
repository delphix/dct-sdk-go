/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the CopyMaskingJobParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyMaskingJobParameters{}

// CopyMaskingJobParameters Parameters to copy a masking job.
type CopyMaskingJobParameters struct {
	// The ID of the engine to copy the job to.
	TargetEngineId string `json:"target_engine_id"`
	// The ID or name of the source environment on the target engine. This only applies to On-The-Fly jobs.
	SourceEnvironmentId *string `json:"source_environment_id,omitempty"`
	// The ID or name of the target environment to copy the job into.
	TargetEnvironmentId *string `json:"target_environment_id,omitempty"`
	// Whether to overwrite objects that already exist on the target engine.
	ForceOverwrite *bool `json:"force_overwrite,omitempty"`
}

// NewCopyMaskingJobParameters instantiates a new CopyMaskingJobParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyMaskingJobParameters(targetEngineId string) *CopyMaskingJobParameters {
	this := CopyMaskingJobParameters{}
	this.TargetEngineId = targetEngineId
	var forceOverwrite bool = false
	this.ForceOverwrite = &forceOverwrite
	return &this
}

// NewCopyMaskingJobParametersWithDefaults instantiates a new CopyMaskingJobParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyMaskingJobParametersWithDefaults() *CopyMaskingJobParameters {
	this := CopyMaskingJobParameters{}
	var forceOverwrite bool = false
	this.ForceOverwrite = &forceOverwrite
	return &this
}

// GetTargetEngineId returns the TargetEngineId field value
func (o *CopyMaskingJobParameters) GetTargetEngineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEngineId
}

// GetTargetEngineIdOk returns a tuple with the TargetEngineId field value
// and a boolean to check if the value has been set.
func (o *CopyMaskingJobParameters) GetTargetEngineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEngineId, true
}

// SetTargetEngineId sets field value
func (o *CopyMaskingJobParameters) SetTargetEngineId(v string) {
	o.TargetEngineId = v
}

// GetSourceEnvironmentId returns the SourceEnvironmentId field value if set, zero value otherwise.
func (o *CopyMaskingJobParameters) GetSourceEnvironmentId() string {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		var ret string
		return ret
	}
	return *o.SourceEnvironmentId
}

// GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMaskingJobParameters) GetSourceEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		return nil, false
	}
	return o.SourceEnvironmentId, true
}

// HasSourceEnvironmentId returns a boolean if a field has been set.
func (o *CopyMaskingJobParameters) HasSourceEnvironmentId() bool {
	if o != nil && !IsNil(o.SourceEnvironmentId) {
		return true
	}

	return false
}

// SetSourceEnvironmentId gets a reference to the given string and assigns it to the SourceEnvironmentId field.
func (o *CopyMaskingJobParameters) SetSourceEnvironmentId(v string) {
	o.SourceEnvironmentId = &v
}

// GetTargetEnvironmentId returns the TargetEnvironmentId field value if set, zero value otherwise.
func (o *CopyMaskingJobParameters) GetTargetEnvironmentId() string {
	if o == nil || IsNil(o.TargetEnvironmentId) {
		var ret string
		return ret
	}
	return *o.TargetEnvironmentId
}

// GetTargetEnvironmentIdOk returns a tuple with the TargetEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMaskingJobParameters) GetTargetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetEnvironmentId) {
		return nil, false
	}
	return o.TargetEnvironmentId, true
}

// HasTargetEnvironmentId returns a boolean if a field has been set.
func (o *CopyMaskingJobParameters) HasTargetEnvironmentId() bool {
	if o != nil && !IsNil(o.TargetEnvironmentId) {
		return true
	}

	return false
}

// SetTargetEnvironmentId gets a reference to the given string and assigns it to the TargetEnvironmentId field.
func (o *CopyMaskingJobParameters) SetTargetEnvironmentId(v string) {
	o.TargetEnvironmentId = &v
}

// GetForceOverwrite returns the ForceOverwrite field value if set, zero value otherwise.
func (o *CopyMaskingJobParameters) GetForceOverwrite() bool {
	if o == nil || IsNil(o.ForceOverwrite) {
		var ret bool
		return ret
	}
	return *o.ForceOverwrite
}

// GetForceOverwriteOk returns a tuple with the ForceOverwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMaskingJobParameters) GetForceOverwriteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceOverwrite) {
		return nil, false
	}
	return o.ForceOverwrite, true
}

// HasForceOverwrite returns a boolean if a field has been set.
func (o *CopyMaskingJobParameters) HasForceOverwrite() bool {
	if o != nil && !IsNil(o.ForceOverwrite) {
		return true
	}

	return false
}

// SetForceOverwrite gets a reference to the given bool and assigns it to the ForceOverwrite field.
func (o *CopyMaskingJobParameters) SetForceOverwrite(v bool) {
	o.ForceOverwrite = &v
}

func (o CopyMaskingJobParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyMaskingJobParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target_engine_id"] = o.TargetEngineId
	if !IsNil(o.SourceEnvironmentId) {
		toSerialize["source_environment_id"] = o.SourceEnvironmentId
	}
	if !IsNil(o.TargetEnvironmentId) {
		toSerialize["target_environment_id"] = o.TargetEnvironmentId
	}
	if !IsNil(o.ForceOverwrite) {
		toSerialize["force_overwrite"] = o.ForceOverwrite
	}
	return toSerialize, nil
}

type NullableCopyMaskingJobParameters struct {
	value *CopyMaskingJobParameters
	isSet bool
}

func (v NullableCopyMaskingJobParameters) Get() *CopyMaskingJobParameters {
	return v.value
}

func (v *NullableCopyMaskingJobParameters) Set(val *CopyMaskingJobParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyMaskingJobParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyMaskingJobParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyMaskingJobParameters(val *CopyMaskingJobParameters) *NullableCopyMaskingJobParameters {
	return &NullableCopyMaskingJobParameters{value: val, isSet: true}
}

func (v NullableCopyMaskingJobParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyMaskingJobParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


