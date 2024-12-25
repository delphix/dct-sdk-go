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

// checks if the AccessGroupUpdateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessGroupUpdateParameters{}

// AccessGroupUpdateParameters struct for AccessGroupUpdateParameters
type AccessGroupUpdateParameters struct {
	// The Access group name
	Name *string `json:"name,omitempty"`
}

// NewAccessGroupUpdateParameters instantiates a new AccessGroupUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessGroupUpdateParameters() *AccessGroupUpdateParameters {
	this := AccessGroupUpdateParameters{}
	return &this
}

// NewAccessGroupUpdateParametersWithDefaults instantiates a new AccessGroupUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessGroupUpdateParametersWithDefaults() *AccessGroupUpdateParameters {
	this := AccessGroupUpdateParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessGroupUpdateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupUpdateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessGroupUpdateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessGroupUpdateParameters) SetName(v string) {
	o.Name = &v
}

func (o AccessGroupUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessGroupUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAccessGroupUpdateParameters struct {
	value *AccessGroupUpdateParameters
	isSet bool
}

func (v NullableAccessGroupUpdateParameters) Get() *AccessGroupUpdateParameters {
	return v.value
}

func (v *NullableAccessGroupUpdateParameters) Set(val *AccessGroupUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessGroupUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessGroupUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessGroupUpdateParameters(val *AccessGroupUpdateParameters) *NullableAccessGroupUpdateParameters {
	return &NullableAccessGroupUpdateParameters{value: val, isSet: true}
}

func (v NullableAccessGroupUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessGroupUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


