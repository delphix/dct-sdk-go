/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the SearchReplicationProfilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchReplicationProfilesResponse{}

// SearchReplicationProfilesResponse struct for SearchReplicationProfilesResponse
type SearchReplicationProfilesResponse struct {
	Items []ReplicationProfile `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchReplicationProfilesResponse instantiates a new SearchReplicationProfilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchReplicationProfilesResponse() *SearchReplicationProfilesResponse {
	this := SearchReplicationProfilesResponse{}
	return &this
}

// NewSearchReplicationProfilesResponseWithDefaults instantiates a new SearchReplicationProfilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchReplicationProfilesResponseWithDefaults() *SearchReplicationProfilesResponse {
	this := SearchReplicationProfilesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchReplicationProfilesResponse) GetItems() []ReplicationProfile {
	if o == nil || IsNil(o.Items) {
		var ret []ReplicationProfile
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchReplicationProfilesResponse) GetItemsOk() ([]ReplicationProfile, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchReplicationProfilesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ReplicationProfile and assigns it to the Items field.
func (o *SearchReplicationProfilesResponse) SetItems(v []ReplicationProfile) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchReplicationProfilesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchReplicationProfilesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchReplicationProfilesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchReplicationProfilesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchReplicationProfilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchReplicationProfilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchReplicationProfilesResponse struct {
	value *SearchReplicationProfilesResponse
	isSet bool
}

func (v NullableSearchReplicationProfilesResponse) Get() *SearchReplicationProfilesResponse {
	return v.value
}

func (v *NullableSearchReplicationProfilesResponse) Set(val *SearchReplicationProfilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchReplicationProfilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchReplicationProfilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchReplicationProfilesResponse(val *SearchReplicationProfilesResponse) *NullableSearchReplicationProfilesResponse {
	return &NullableSearchReplicationProfilesResponse{value: val, isSet: true}
}

func (v NullableSearchReplicationProfilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchReplicationProfilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


