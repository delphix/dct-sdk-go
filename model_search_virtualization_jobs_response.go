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

// checks if the SearchVirtualizationJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchVirtualizationJobsResponse{}

// SearchVirtualizationJobsResponse struct for SearchVirtualizationJobsResponse
type SearchVirtualizationJobsResponse struct {
	Items []VirtualizationJob `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchVirtualizationJobsResponse instantiates a new SearchVirtualizationJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchVirtualizationJobsResponse() *SearchVirtualizationJobsResponse {
	this := SearchVirtualizationJobsResponse{}
	return &this
}

// NewSearchVirtualizationJobsResponseWithDefaults instantiates a new SearchVirtualizationJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchVirtualizationJobsResponseWithDefaults() *SearchVirtualizationJobsResponse {
	this := SearchVirtualizationJobsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchVirtualizationJobsResponse) GetItems() []VirtualizationJob {
	if o == nil || IsNil(o.Items) {
		var ret []VirtualizationJob
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVirtualizationJobsResponse) GetItemsOk() ([]VirtualizationJob, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchVirtualizationJobsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VirtualizationJob and assigns it to the Items field.
func (o *SearchVirtualizationJobsResponse) SetItems(v []VirtualizationJob) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchVirtualizationJobsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVirtualizationJobsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchVirtualizationJobsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchVirtualizationJobsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchVirtualizationJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchVirtualizationJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchVirtualizationJobsResponse struct {
	value *SearchVirtualizationJobsResponse
	isSet bool
}

func (v NullableSearchVirtualizationJobsResponse) Get() *SearchVirtualizationJobsResponse {
	return v.value
}

func (v *NullableSearchVirtualizationJobsResponse) Set(val *SearchVirtualizationJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchVirtualizationJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchVirtualizationJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchVirtualizationJobsResponse(val *SearchVirtualizationJobsResponse) *NullableSearchVirtualizationJobsResponse {
	return &NullableSearchVirtualizationJobsResponse{value: val, isSet: true}
}

func (v NullableSearchVirtualizationJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchVirtualizationJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


