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

// DataPointFromBookmarkParameters struct for DataPointFromBookmarkParameters
type DataPointFromBookmarkParameters struct {
	// The ID of the bookmark from which to execute the operation. The boomkark must contain only one VDB.
	BookmarkId string `json:"bookmark_id"`
}

// NewDataPointFromBookmarkParameters instantiates a new DataPointFromBookmarkParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPointFromBookmarkParameters(bookmarkId string) *DataPointFromBookmarkParameters {
	this := DataPointFromBookmarkParameters{}
	this.BookmarkId = bookmarkId
	return &this
}

// NewDataPointFromBookmarkParametersWithDefaults instantiates a new DataPointFromBookmarkParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPointFromBookmarkParametersWithDefaults() *DataPointFromBookmarkParameters {
	this := DataPointFromBookmarkParameters{}
	return &this
}

// GetBookmarkId returns the BookmarkId field value
func (o *DataPointFromBookmarkParameters) GetBookmarkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BookmarkId
}

// GetBookmarkIdOk returns a tuple with the BookmarkId field value
// and a boolean to check if the value has been set.
func (o *DataPointFromBookmarkParameters) GetBookmarkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BookmarkId, true
}

// SetBookmarkId sets field value
func (o *DataPointFromBookmarkParameters) SetBookmarkId(v string) {
	o.BookmarkId = v
}

func (o DataPointFromBookmarkParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bookmark_id"] = o.BookmarkId
	}
	return json.Marshal(toSerialize)
}

type NullableDataPointFromBookmarkParameters struct {
	value *DataPointFromBookmarkParameters
	isSet bool
}

func (v NullableDataPointFromBookmarkParameters) Get() *DataPointFromBookmarkParameters {
	return v.value
}

func (v *NullableDataPointFromBookmarkParameters) Set(val *DataPointFromBookmarkParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPointFromBookmarkParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPointFromBookmarkParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPointFromBookmarkParameters(val *DataPointFromBookmarkParameters) *NullableDataPointFromBookmarkParameters {
	return &NullableDataPointFromBookmarkParameters{value: val, isSet: true}
}

func (v NullableDataPointFromBookmarkParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPointFromBookmarkParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


