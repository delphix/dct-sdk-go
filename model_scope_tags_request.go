/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ScopeTagsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeTagsRequest{}

// ScopeTagsRequest struct for ScopeTagsRequest
type ScopeTagsRequest struct {
	// Array of tags with key value pairs along with optional object_type and permissions
	Tags []ScopeTag `json:"tags"`
}

// NewScopeTagsRequest instantiates a new ScopeTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeTagsRequest(tags []ScopeTag) *ScopeTagsRequest {
	this := ScopeTagsRequest{}
	this.Tags = tags
	return &this
}

// NewScopeTagsRequestWithDefaults instantiates a new ScopeTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeTagsRequestWithDefaults() *ScopeTagsRequest {
	this := ScopeTagsRequest{}
	return &this
}

// GetTags returns the Tags field value
func (o *ScopeTagsRequest) GetTags() []ScopeTag {
	if o == nil {
		var ret []ScopeTag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ScopeTagsRequest) GetTagsOk() ([]ScopeTag, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ScopeTagsRequest) SetTags(v []ScopeTag) {
	o.Tags = v
}

func (o ScopeTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeTagsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

type NullableScopeTagsRequest struct {
	value *ScopeTagsRequest
	isSet bool
}

func (v NullableScopeTagsRequest) Get() *ScopeTagsRequest {
	return v.value
}

func (v *NullableScopeTagsRequest) Set(val *ScopeTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeTagsRequest(val *ScopeTagsRequest) *NullableScopeTagsRequest {
	return &NullableScopeTagsRequest{value: val, isSet: true}
}

func (v NullableScopeTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


