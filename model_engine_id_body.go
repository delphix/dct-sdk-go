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

// checks if the EngineIdBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EngineIdBody{}

// EngineIdBody The ID of the registered engine.
type EngineIdBody struct {
	EngineId *string `json:"engine_id,omitempty"`
}

// NewEngineIdBody instantiates a new EngineIdBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineIdBody() *EngineIdBody {
	this := EngineIdBody{}
	return &this
}

// NewEngineIdBodyWithDefaults instantiates a new EngineIdBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineIdBodyWithDefaults() *EngineIdBody {
	this := EngineIdBody{}
	return &this
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *EngineIdBody) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineIdBody) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *EngineIdBody) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *EngineIdBody) SetEngineId(v string) {
	o.EngineId = &v
}

func (o EngineIdBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EngineIdBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	return toSerialize, nil
}

type NullableEngineIdBody struct {
	value *EngineIdBody
	isSet bool
}

func (v NullableEngineIdBody) Get() *EngineIdBody {
	return v.value
}

func (v *NullableEngineIdBody) Set(val *EngineIdBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineIdBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineIdBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineIdBody(val *EngineIdBody) *NullableEngineIdBody {
	return &NullableEngineIdBody{value: val, isSet: true}
}

func (v NullableEngineIdBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineIdBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


