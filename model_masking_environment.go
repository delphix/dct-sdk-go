/*
Delphix DCT API

Delphix DCT API

API version: 3.5.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the MaskingEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskingEnvironment{}

// MaskingEnvironment struct for MaskingEnvironment
type MaskingEnvironment struct {
	// The MaskingEnvironment entity ID.
	Id *string `json:"id,omitempty"`
	// The ID of the Engine where this MaskingEnvironment resides.
	EngineId *string `json:"engine_id,omitempty"`
	// The name of the Engine where this MaskingEnvironment resides.
	EngineName *string `json:"engine_name,omitempty"`
	// The name of this MaskingEnvironment.
	Name *string `json:"name,omitempty"`
	// The purpose of this MaskingEnvironment. MaskingEnvironments with a 'MASK' purpose will have access to Masking and Profiling jobs, whereas Environments with a 'TOKENIZE' purpose will have access to Tokenization and Re-Identification jobs. Note that any custom purposes created through the UI will be represented as 'MASK' purposes, due to the jobs that they have access to.
	Purpose *string `json:"purpose,omitempty"`
	IsWorkflowEnabled *bool `json:"is_workflow_enabled,omitempty"`
}

// NewMaskingEnvironment instantiates a new MaskingEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskingEnvironment() *MaskingEnvironment {
	this := MaskingEnvironment{}
	return &this
}

// NewMaskingEnvironmentWithDefaults instantiates a new MaskingEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskingEnvironmentWithDefaults() *MaskingEnvironment {
	this := MaskingEnvironment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MaskingEnvironment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingEnvironment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MaskingEnvironment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MaskingEnvironment) SetId(v string) {
	o.Id = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *MaskingEnvironment) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingEnvironment) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *MaskingEnvironment) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *MaskingEnvironment) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *MaskingEnvironment) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingEnvironment) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *MaskingEnvironment) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *MaskingEnvironment) SetEngineName(v string) {
	o.EngineName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MaskingEnvironment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingEnvironment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MaskingEnvironment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MaskingEnvironment) SetName(v string) {
	o.Name = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *MaskingEnvironment) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingEnvironment) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *MaskingEnvironment) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *MaskingEnvironment) SetPurpose(v string) {
	o.Purpose = &v
}

// GetIsWorkflowEnabled returns the IsWorkflowEnabled field value if set, zero value otherwise.
func (o *MaskingEnvironment) GetIsWorkflowEnabled() bool {
	if o == nil || IsNil(o.IsWorkflowEnabled) {
		var ret bool
		return ret
	}
	return *o.IsWorkflowEnabled
}

// GetIsWorkflowEnabledOk returns a tuple with the IsWorkflowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingEnvironment) GetIsWorkflowEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWorkflowEnabled) {
		return nil, false
	}
	return o.IsWorkflowEnabled, true
}

// HasIsWorkflowEnabled returns a boolean if a field has been set.
func (o *MaskingEnvironment) HasIsWorkflowEnabled() bool {
	if o != nil && !IsNil(o.IsWorkflowEnabled) {
		return true
	}

	return false
}

// SetIsWorkflowEnabled gets a reference to the given bool and assigns it to the IsWorkflowEnabled field.
func (o *MaskingEnvironment) SetIsWorkflowEnabled(v bool) {
	o.IsWorkflowEnabled = &v
}

func (o MaskingEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskingEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.IsWorkflowEnabled) {
		toSerialize["is_workflow_enabled"] = o.IsWorkflowEnabled
	}
	return toSerialize, nil
}

type NullableMaskingEnvironment struct {
	value *MaskingEnvironment
	isSet bool
}

func (v NullableMaskingEnvironment) Get() *MaskingEnvironment {
	return v.value
}

func (v *NullableMaskingEnvironment) Set(val *MaskingEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskingEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskingEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskingEnvironment(val *MaskingEnvironment) *NullableMaskingEnvironment {
	return &NullableMaskingEnvironment{value: val, isSet: true}
}

func (v NullableMaskingEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskingEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

