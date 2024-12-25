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

// checks if the SnapshotsDayRangesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotsDayRangesResponse{}

// SnapshotsDayRangesResponse struct for SnapshotsDayRangesResponse
type SnapshotsDayRangesResponse struct {
	Items []SnapshotDayRange `json:"items,omitempty"`
}

// NewSnapshotsDayRangesResponse instantiates a new SnapshotsDayRangesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotsDayRangesResponse() *SnapshotsDayRangesResponse {
	this := SnapshotsDayRangesResponse{}
	return &this
}

// NewSnapshotsDayRangesResponseWithDefaults instantiates a new SnapshotsDayRangesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotsDayRangesResponseWithDefaults() *SnapshotsDayRangesResponse {
	this := SnapshotsDayRangesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SnapshotsDayRangesResponse) GetItems() []SnapshotDayRange {
	if o == nil || IsNil(o.Items) {
		var ret []SnapshotDayRange
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsDayRangesResponse) GetItemsOk() ([]SnapshotDayRange, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SnapshotsDayRangesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SnapshotDayRange and assigns it to the Items field.
func (o *SnapshotsDayRangesResponse) SetItems(v []SnapshotDayRange) {
	o.Items = v
}

func (o SnapshotsDayRangesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotsDayRangesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSnapshotsDayRangesResponse struct {
	value *SnapshotsDayRangesResponse
	isSet bool
}

func (v NullableSnapshotsDayRangesResponse) Get() *SnapshotsDayRangesResponse {
	return v.value
}

func (v *NullableSnapshotsDayRangesResponse) Set(val *SnapshotsDayRangesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotsDayRangesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotsDayRangesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotsDayRangesResponse(val *SnapshotsDayRangesResponse) *NullableSnapshotsDayRangesResponse {
	return &NullableSnapshotsDayRangesResponse{value: val, isSet: true}
}

func (v NullableSnapshotsDayRangesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotsDayRangesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


