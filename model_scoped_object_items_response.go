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

// checks if the ScopedObjectItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopedObjectItemsResponse{}

// ScopedObjectItemsResponse struct for ScopedObjectItemsResponse
type ScopedObjectItemsResponse struct {
	// Array of access group scope objects with object id and object type
	Objects []ScopedObjectItem `json:"objects,omitempty"`
}

// NewScopedObjectItemsResponse instantiates a new ScopedObjectItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopedObjectItemsResponse() *ScopedObjectItemsResponse {
	this := ScopedObjectItemsResponse{}
	return &this
}

// NewScopedObjectItemsResponseWithDefaults instantiates a new ScopedObjectItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopedObjectItemsResponseWithDefaults() *ScopedObjectItemsResponse {
	this := ScopedObjectItemsResponse{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *ScopedObjectItemsResponse) GetObjects() []ScopedObjectItem {
	if o == nil || IsNil(o.Objects) {
		var ret []ScopedObjectItem
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopedObjectItemsResponse) GetObjectsOk() ([]ScopedObjectItem, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *ScopedObjectItemsResponse) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []ScopedObjectItem and assigns it to the Objects field.
func (o *ScopedObjectItemsResponse) SetObjects(v []ScopedObjectItem) {
	o.Objects = v
}

func (o ScopedObjectItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopedObjectItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableScopedObjectItemsResponse struct {
	value *ScopedObjectItemsResponse
	isSet bool
}

func (v NullableScopedObjectItemsResponse) Get() *ScopedObjectItemsResponse {
	return v.value
}

func (v *NullableScopedObjectItemsResponse) Set(val *ScopedObjectItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScopedObjectItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScopedObjectItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopedObjectItemsResponse(val *ScopedObjectItemsResponse) *NullableScopedObjectItemsResponse {
	return &NullableScopedObjectItemsResponse{value: val, isSet: true}
}

func (v NullableScopedObjectItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopedObjectItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


