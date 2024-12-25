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

// checks if the SearchMaskingJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchMaskingJobsResponse{}

// SearchMaskingJobsResponse struct for SearchMaskingJobsResponse
type SearchMaskingJobsResponse struct {
	Items []MaskingJob `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchMaskingJobsResponse instantiates a new SearchMaskingJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchMaskingJobsResponse() *SearchMaskingJobsResponse {
	this := SearchMaskingJobsResponse{}
	return &this
}

// NewSearchMaskingJobsResponseWithDefaults instantiates a new SearchMaskingJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchMaskingJobsResponseWithDefaults() *SearchMaskingJobsResponse {
	this := SearchMaskingJobsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchMaskingJobsResponse) GetItems() []MaskingJob {
	if o == nil || IsNil(o.Items) {
		var ret []MaskingJob
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMaskingJobsResponse) GetItemsOk() ([]MaskingJob, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchMaskingJobsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MaskingJob and assigns it to the Items field.
func (o *SearchMaskingJobsResponse) SetItems(v []MaskingJob) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchMaskingJobsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMaskingJobsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchMaskingJobsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchMaskingJobsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchMaskingJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchMaskingJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchMaskingJobsResponse struct {
	value *SearchMaskingJobsResponse
	isSet bool
}

func (v NullableSearchMaskingJobsResponse) Get() *SearchMaskingJobsResponse {
	return v.value
}

func (v *NullableSearchMaskingJobsResponse) Set(val *SearchMaskingJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMaskingJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMaskingJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMaskingJobsResponse(val *SearchMaskingJobsResponse) *NullableSearchMaskingJobsResponse {
	return &NullableSearchMaskingJobsResponse{value: val, isSet: true}
}

func (v NullableSearchMaskingJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMaskingJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


