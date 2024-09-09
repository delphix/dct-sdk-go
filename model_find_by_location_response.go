/*
Delphix DCT API

Delphix DCT API

API version: 3.16.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the FindByLocationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindByLocationResponse{}

// FindByLocationResponse struct for FindByLocationResponse
type FindByLocationResponse struct {
	Items []Snapshot `json:"items,omitempty"`
}

// NewFindByLocationResponse instantiates a new FindByLocationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindByLocationResponse() *FindByLocationResponse {
	this := FindByLocationResponse{}
	return &this
}

// NewFindByLocationResponseWithDefaults instantiates a new FindByLocationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindByLocationResponseWithDefaults() *FindByLocationResponse {
	this := FindByLocationResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *FindByLocationResponse) GetItems() []Snapshot {
	if o == nil || IsNil(o.Items) {
		var ret []Snapshot
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindByLocationResponse) GetItemsOk() ([]Snapshot, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *FindByLocationResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Snapshot and assigns it to the Items field.
func (o *FindByLocationResponse) SetItems(v []Snapshot) {
	o.Items = v
}

func (o FindByLocationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindByLocationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableFindByLocationResponse struct {
	value *FindByLocationResponse
	isSet bool
}

func (v NullableFindByLocationResponse) Get() *FindByLocationResponse {
	return v.value
}

func (v *NullableFindByLocationResponse) Set(val *FindByLocationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFindByLocationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFindByLocationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindByLocationResponse(val *FindByLocationResponse) *NullableFindByLocationResponse {
	return &NullableFindByLocationResponse{value: val, isSet: true}
}

func (v NullableFindByLocationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindByLocationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


