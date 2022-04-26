/*
Delphix DCT API

Delphix DCT API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// SearchSourcesResponse struct for SearchSourcesResponse
type SearchSourcesResponse struct {
	Items []Source `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchSourcesResponse instantiates a new SearchSourcesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSourcesResponse() *SearchSourcesResponse {
	this := SearchSourcesResponse{}
	return &this
}

// NewSearchSourcesResponseWithDefaults instantiates a new SearchSourcesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSourcesResponseWithDefaults() *SearchSourcesResponse {
	this := SearchSourcesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchSourcesResponse) GetItems() []Source {
	if o == nil || o.Items == nil {
		var ret []Source
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSourcesResponse) GetItemsOk() ([]Source, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchSourcesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Source and assigns it to the Items field.
func (o *SearchSourcesResponse) SetItems(v []Source) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchSourcesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || o.ResponseMetadata == nil {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSourcesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || o.ResponseMetadata == nil {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchSourcesResponse) HasResponseMetadata() bool {
	if o != nil && o.ResponseMetadata != nil {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchSourcesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchSourcesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ResponseMetadata != nil {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableSearchSourcesResponse struct {
	value *SearchSourcesResponse
	isSet bool
}

func (v NullableSearchSourcesResponse) Get() *SearchSourcesResponse {
	return v.value
}

func (v *NullableSearchSourcesResponse) Set(val *SearchSourcesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSourcesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSourcesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSourcesResponse(val *SearchSourcesResponse) *NullableSearchSourcesResponse {
	return &NullableSearchSourcesResponse{value: val, isSet: true}
}

func (v NullableSearchSourcesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSourcesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


