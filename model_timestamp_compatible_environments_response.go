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

// checks if the TimestampCompatibleEnvironmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimestampCompatibleEnvironmentsResponse{}

// TimestampCompatibleEnvironmentsResponse struct for TimestampCompatibleEnvironmentsResponse
type TimestampCompatibleEnvironmentsResponse struct {
	Items []Environment `json:"items,omitempty"`
}

// NewTimestampCompatibleEnvironmentsResponse instantiates a new TimestampCompatibleEnvironmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimestampCompatibleEnvironmentsResponse() *TimestampCompatibleEnvironmentsResponse {
	this := TimestampCompatibleEnvironmentsResponse{}
	return &this
}

// NewTimestampCompatibleEnvironmentsResponseWithDefaults instantiates a new TimestampCompatibleEnvironmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimestampCompatibleEnvironmentsResponseWithDefaults() *TimestampCompatibleEnvironmentsResponse {
	this := TimestampCompatibleEnvironmentsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TimestampCompatibleEnvironmentsResponse) GetItems() []Environment {
	if o == nil || IsNil(o.Items) {
		var ret []Environment
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampCompatibleEnvironmentsResponse) GetItemsOk() ([]Environment, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TimestampCompatibleEnvironmentsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Environment and assigns it to the Items field.
func (o *TimestampCompatibleEnvironmentsResponse) SetItems(v []Environment) {
	o.Items = v
}

func (o TimestampCompatibleEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimestampCompatibleEnvironmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableTimestampCompatibleEnvironmentsResponse struct {
	value *TimestampCompatibleEnvironmentsResponse
	isSet bool
}

func (v NullableTimestampCompatibleEnvironmentsResponse) Get() *TimestampCompatibleEnvironmentsResponse {
	return v.value
}

func (v *NullableTimestampCompatibleEnvironmentsResponse) Set(val *TimestampCompatibleEnvironmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimestampCompatibleEnvironmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimestampCompatibleEnvironmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimestampCompatibleEnvironmentsResponse(val *TimestampCompatibleEnvironmentsResponse) *NullableTimestampCompatibleEnvironmentsResponse {
	return &NullableTimestampCompatibleEnvironmentsResponse{value: val, isSet: true}
}

func (v NullableTimestampCompatibleEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimestampCompatibleEnvironmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

