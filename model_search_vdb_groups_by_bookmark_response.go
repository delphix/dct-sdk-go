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

// checks if the SearchVDBGroupsByBookmarkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchVDBGroupsByBookmarkResponse{}

// SearchVDBGroupsByBookmarkResponse struct for SearchVDBGroupsByBookmarkResponse
type SearchVDBGroupsByBookmarkResponse struct {
	Items []VDBGroup `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchVDBGroupsByBookmarkResponse instantiates a new SearchVDBGroupsByBookmarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchVDBGroupsByBookmarkResponse() *SearchVDBGroupsByBookmarkResponse {
	this := SearchVDBGroupsByBookmarkResponse{}
	return &this
}

// NewSearchVDBGroupsByBookmarkResponseWithDefaults instantiates a new SearchVDBGroupsByBookmarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchVDBGroupsByBookmarkResponseWithDefaults() *SearchVDBGroupsByBookmarkResponse {
	this := SearchVDBGroupsByBookmarkResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchVDBGroupsByBookmarkResponse) GetItems() []VDBGroup {
	if o == nil || IsNil(o.Items) {
		var ret []VDBGroup
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVDBGroupsByBookmarkResponse) GetItemsOk() ([]VDBGroup, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchVDBGroupsByBookmarkResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VDBGroup and assigns it to the Items field.
func (o *SearchVDBGroupsByBookmarkResponse) SetItems(v []VDBGroup) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchVDBGroupsByBookmarkResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVDBGroupsByBookmarkResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchVDBGroupsByBookmarkResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchVDBGroupsByBookmarkResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchVDBGroupsByBookmarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchVDBGroupsByBookmarkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchVDBGroupsByBookmarkResponse struct {
	value *SearchVDBGroupsByBookmarkResponse
	isSet bool
}

func (v NullableSearchVDBGroupsByBookmarkResponse) Get() *SearchVDBGroupsByBookmarkResponse {
	return v.value
}

func (v *NullableSearchVDBGroupsByBookmarkResponse) Set(val *SearchVDBGroupsByBookmarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchVDBGroupsByBookmarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchVDBGroupsByBookmarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchVDBGroupsByBookmarkResponse(val *SearchVDBGroupsByBookmarkResponse) *NullableSearchVDBGroupsByBookmarkResponse {
	return &NullableSearchVDBGroupsByBookmarkResponse{value: val, isSet: true}
}

func (v NullableSearchVDBGroupsByBookmarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchVDBGroupsByBookmarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


