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

// checks if the SearchDiscoveryResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchDiscoveryResultsResponse{}

// SearchDiscoveryResultsResponse struct for SearchDiscoveryResultsResponse
type SearchDiscoveryResultsResponse struct {
	Items []DiscoveryResult `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchDiscoveryResultsResponse instantiates a new SearchDiscoveryResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchDiscoveryResultsResponse() *SearchDiscoveryResultsResponse {
	this := SearchDiscoveryResultsResponse{}
	return &this
}

// NewSearchDiscoveryResultsResponseWithDefaults instantiates a new SearchDiscoveryResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchDiscoveryResultsResponseWithDefaults() *SearchDiscoveryResultsResponse {
	this := SearchDiscoveryResultsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchDiscoveryResultsResponse) GetItems() []DiscoveryResult {
	if o == nil || IsNil(o.Items) {
		var ret []DiscoveryResult
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDiscoveryResultsResponse) GetItemsOk() ([]DiscoveryResult, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchDiscoveryResultsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DiscoveryResult and assigns it to the Items field.
func (o *SearchDiscoveryResultsResponse) SetItems(v []DiscoveryResult) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchDiscoveryResultsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDiscoveryResultsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchDiscoveryResultsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchDiscoveryResultsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchDiscoveryResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchDiscoveryResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchDiscoveryResultsResponse struct {
	value *SearchDiscoveryResultsResponse
	isSet bool
}

func (v NullableSearchDiscoveryResultsResponse) Get() *SearchDiscoveryResultsResponse {
	return v.value
}

func (v *NullableSearchDiscoveryResultsResponse) Set(val *SearchDiscoveryResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchDiscoveryResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchDiscoveryResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchDiscoveryResultsResponse(val *SearchDiscoveryResultsResponse) *NullableSearchDiscoveryResultsResponse {
	return &NullableSearchDiscoveryResultsResponse{value: val, isSet: true}
}

func (v NullableSearchDiscoveryResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchDiscoveryResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


