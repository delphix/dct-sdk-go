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

// checks if the BookmarkDSources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkDSources{}

// BookmarkDSources dSource id and name associated with bookmark.
type BookmarkDSources struct {
	// The dSource id.
	DsourceId *string `json:"dsource_id,omitempty"`
	// The dSource name.
	DsourceName *string `json:"dsource_name,omitempty"`
}

// NewBookmarkDSources instantiates a new BookmarkDSources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkDSources() *BookmarkDSources {
	this := BookmarkDSources{}
	return &this
}

// NewBookmarkDSourcesWithDefaults instantiates a new BookmarkDSources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkDSourcesWithDefaults() *BookmarkDSources {
	this := BookmarkDSources{}
	return &this
}

// GetDsourceId returns the DsourceId field value if set, zero value otherwise.
func (o *BookmarkDSources) GetDsourceId() string {
	if o == nil || IsNil(o.DsourceId) {
		var ret string
		return ret
	}
	return *o.DsourceId
}

// GetDsourceIdOk returns a tuple with the DsourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkDSources) GetDsourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DsourceId) {
		return nil, false
	}
	return o.DsourceId, true
}

// HasDsourceId returns a boolean if a field has been set.
func (o *BookmarkDSources) HasDsourceId() bool {
	if o != nil && !IsNil(o.DsourceId) {
		return true
	}

	return false
}

// SetDsourceId gets a reference to the given string and assigns it to the DsourceId field.
func (o *BookmarkDSources) SetDsourceId(v string) {
	o.DsourceId = &v
}

// GetDsourceName returns the DsourceName field value if set, zero value otherwise.
func (o *BookmarkDSources) GetDsourceName() string {
	if o == nil || IsNil(o.DsourceName) {
		var ret string
		return ret
	}
	return *o.DsourceName
}

// GetDsourceNameOk returns a tuple with the DsourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkDSources) GetDsourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DsourceName) {
		return nil, false
	}
	return o.DsourceName, true
}

// HasDsourceName returns a boolean if a field has been set.
func (o *BookmarkDSources) HasDsourceName() bool {
	if o != nil && !IsNil(o.DsourceName) {
		return true
	}

	return false
}

// SetDsourceName gets a reference to the given string and assigns it to the DsourceName field.
func (o *BookmarkDSources) SetDsourceName(v string) {
	o.DsourceName = &v
}

func (o BookmarkDSources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkDSources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DsourceId) {
		toSerialize["dsource_id"] = o.DsourceId
	}
	if !IsNil(o.DsourceName) {
		toSerialize["dsource_name"] = o.DsourceName
	}
	return toSerialize, nil
}

type NullableBookmarkDSources struct {
	value *BookmarkDSources
	isSet bool
}

func (v NullableBookmarkDSources) Get() *BookmarkDSources {
	return v.value
}

func (v *NullableBookmarkDSources) Set(val *BookmarkDSources) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkDSources) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkDSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkDSources(val *BookmarkDSources) *NullableBookmarkDSources {
	return &NullableBookmarkDSources{value: val, isSet: true}
}

func (v NullableBookmarkDSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkDSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


