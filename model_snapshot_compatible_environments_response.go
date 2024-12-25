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

// checks if the SnapshotCompatibleEnvironmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotCompatibleEnvironmentsResponse{}

// SnapshotCompatibleEnvironmentsResponse struct for SnapshotCompatibleEnvironmentsResponse
type SnapshotCompatibleEnvironmentsResponse struct {
	Items []Environment `json:"items,omitempty"`
}

// NewSnapshotCompatibleEnvironmentsResponse instantiates a new SnapshotCompatibleEnvironmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotCompatibleEnvironmentsResponse() *SnapshotCompatibleEnvironmentsResponse {
	this := SnapshotCompatibleEnvironmentsResponse{}
	return &this
}

// NewSnapshotCompatibleEnvironmentsResponseWithDefaults instantiates a new SnapshotCompatibleEnvironmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotCompatibleEnvironmentsResponseWithDefaults() *SnapshotCompatibleEnvironmentsResponse {
	this := SnapshotCompatibleEnvironmentsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SnapshotCompatibleEnvironmentsResponse) GetItems() []Environment {
	if o == nil || IsNil(o.Items) {
		var ret []Environment
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCompatibleEnvironmentsResponse) GetItemsOk() ([]Environment, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SnapshotCompatibleEnvironmentsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Environment and assigns it to the Items field.
func (o *SnapshotCompatibleEnvironmentsResponse) SetItems(v []Environment) {
	o.Items = v
}

func (o SnapshotCompatibleEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotCompatibleEnvironmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSnapshotCompatibleEnvironmentsResponse struct {
	value *SnapshotCompatibleEnvironmentsResponse
	isSet bool
}

func (v NullableSnapshotCompatibleEnvironmentsResponse) Get() *SnapshotCompatibleEnvironmentsResponse {
	return v.value
}

func (v *NullableSnapshotCompatibleEnvironmentsResponse) Set(val *SnapshotCompatibleEnvironmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotCompatibleEnvironmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotCompatibleEnvironmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotCompatibleEnvironmentsResponse(val *SnapshotCompatibleEnvironmentsResponse) *NullableSnapshotCompatibleEnvironmentsResponse {
	return &NullableSnapshotCompatibleEnvironmentsResponse{value: val, isSet: true}
}

func (v NullableSnapshotCompatibleEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotCompatibleEnvironmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


