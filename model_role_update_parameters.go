/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the RoleUpdateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleUpdateParameters{}

// RoleUpdateParameters struct for RoleUpdateParameters
type RoleUpdateParameters struct {
	// The role name
	Name *string `json:"name,omitempty"`
	// The role description
	Description *string `json:"description,omitempty"`
}

// NewRoleUpdateParameters instantiates a new RoleUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleUpdateParameters() *RoleUpdateParameters {
	this := RoleUpdateParameters{}
	return &this
}

// NewRoleUpdateParametersWithDefaults instantiates a new RoleUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleUpdateParametersWithDefaults() *RoleUpdateParameters {
	this := RoleUpdateParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleUpdateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleUpdateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleUpdateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleUpdateParameters) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleUpdateParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleUpdateParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleUpdateParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleUpdateParameters) SetDescription(v string) {
	o.Description = &v
}

func (o RoleUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableRoleUpdateParameters struct {
	value *RoleUpdateParameters
	isSet bool
}

func (v NullableRoleUpdateParameters) Get() *RoleUpdateParameters {
	return v.value
}

func (v *NullableRoleUpdateParameters) Set(val *RoleUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleUpdateParameters(val *RoleUpdateParameters) *NullableRoleUpdateParameters {
	return &NullableRoleUpdateParameters{value: val, isSet: true}
}

func (v NullableRoleUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


