/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ListBookmarksByVDBGroupsResponse struct for ListBookmarksByVDBGroupsResponse
type ListBookmarksByVDBGroupsResponse struct {
	Items []Bookmark `json:"items,omitempty"`
}

// NewListBookmarksByVDBGroupsResponse instantiates a new ListBookmarksByVDBGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBookmarksByVDBGroupsResponse() *ListBookmarksByVDBGroupsResponse {
	this := ListBookmarksByVDBGroupsResponse{}
	return &this
}

// NewListBookmarksByVDBGroupsResponseWithDefaults instantiates a new ListBookmarksByVDBGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBookmarksByVDBGroupsResponseWithDefaults() *ListBookmarksByVDBGroupsResponse {
	this := ListBookmarksByVDBGroupsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListBookmarksByVDBGroupsResponse) GetItems() []Bookmark {
	if o == nil || o.Items == nil {
		var ret []Bookmark
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBookmarksByVDBGroupsResponse) GetItemsOk() ([]Bookmark, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListBookmarksByVDBGroupsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Bookmark and assigns it to the Items field.
func (o *ListBookmarksByVDBGroupsResponse) SetItems(v []Bookmark) {
	o.Items = v
}

func (o ListBookmarksByVDBGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableListBookmarksByVDBGroupsResponse struct {
	value *ListBookmarksByVDBGroupsResponse
	isSet bool
}

func (v NullableListBookmarksByVDBGroupsResponse) Get() *ListBookmarksByVDBGroupsResponse {
	return v.value
}

func (v *NullableListBookmarksByVDBGroupsResponse) Set(val *ListBookmarksByVDBGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBookmarksByVDBGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBookmarksByVDBGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBookmarksByVDBGroupsResponse(val *ListBookmarksByVDBGroupsResponse) *NullableListBookmarksByVDBGroupsResponse {
	return &NullableListBookmarksByVDBGroupsResponse{value: val, isSet: true}
}

func (v NullableListBookmarksByVDBGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBookmarksByVDBGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

