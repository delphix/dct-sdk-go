/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// RefreshVDBGroupParameters Parameters to refresh a VDB Group.
type RefreshVDBGroupParameters struct {
	// ID of a bookmark to refresh this VDB Group to.
	BookmarkId string `json:"bookmark_id"`
}

// NewRefreshVDBGroupParameters instantiates a new RefreshVDBGroupParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBGroupParameters(bookmarkId string) *RefreshVDBGroupParameters {
	this := RefreshVDBGroupParameters{}
	this.BookmarkId = bookmarkId
	return &this
}

// NewRefreshVDBGroupParametersWithDefaults instantiates a new RefreshVDBGroupParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBGroupParametersWithDefaults() *RefreshVDBGroupParameters {
	this := RefreshVDBGroupParameters{}
	return &this
}

// GetBookmarkId returns the BookmarkId field value
func (o *RefreshVDBGroupParameters) GetBookmarkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BookmarkId
}

// GetBookmarkIdOk returns a tuple with the BookmarkId field value
// and a boolean to check if the value has been set.
func (o *RefreshVDBGroupParameters) GetBookmarkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BookmarkId, true
}

// SetBookmarkId sets field value
func (o *RefreshVDBGroupParameters) SetBookmarkId(v string) {
	o.BookmarkId = v
}

func (o RefreshVDBGroupParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bookmark_id"] = o.BookmarkId
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshVDBGroupParameters struct {
	value *RefreshVDBGroupParameters
	isSet bool
}

func (v NullableRefreshVDBGroupParameters) Get() *RefreshVDBGroupParameters {
	return v.value
}

func (v *NullableRefreshVDBGroupParameters) Set(val *RefreshVDBGroupParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBGroupParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBGroupParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBGroupParameters(val *RefreshVDBGroupParameters) *NullableRefreshVDBGroupParameters {
	return &NullableRefreshVDBGroupParameters{value: val, isSet: true}
}

func (v NullableRefreshVDBGroupParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBGroupParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


