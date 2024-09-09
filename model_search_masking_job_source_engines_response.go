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

// checks if the SearchMaskingJobSourceEnginesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchMaskingJobSourceEnginesResponse{}

// SearchMaskingJobSourceEnginesResponse struct for SearchMaskingJobSourceEnginesResponse
type SearchMaskingJobSourceEnginesResponse struct {
	Items []MaskingJobSourceEngine `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchMaskingJobSourceEnginesResponse instantiates a new SearchMaskingJobSourceEnginesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchMaskingJobSourceEnginesResponse() *SearchMaskingJobSourceEnginesResponse {
	this := SearchMaskingJobSourceEnginesResponse{}
	return &this
}

// NewSearchMaskingJobSourceEnginesResponseWithDefaults instantiates a new SearchMaskingJobSourceEnginesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchMaskingJobSourceEnginesResponseWithDefaults() *SearchMaskingJobSourceEnginesResponse {
	this := SearchMaskingJobSourceEnginesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchMaskingJobSourceEnginesResponse) GetItems() []MaskingJobSourceEngine {
	if o == nil || IsNil(o.Items) {
		var ret []MaskingJobSourceEngine
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMaskingJobSourceEnginesResponse) GetItemsOk() ([]MaskingJobSourceEngine, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchMaskingJobSourceEnginesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MaskingJobSourceEngine and assigns it to the Items field.
func (o *SearchMaskingJobSourceEnginesResponse) SetItems(v []MaskingJobSourceEngine) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchMaskingJobSourceEnginesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMaskingJobSourceEnginesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchMaskingJobSourceEnginesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchMaskingJobSourceEnginesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchMaskingJobSourceEnginesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchMaskingJobSourceEnginesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchMaskingJobSourceEnginesResponse struct {
	value *SearchMaskingJobSourceEnginesResponse
	isSet bool
}

func (v NullableSearchMaskingJobSourceEnginesResponse) Get() *SearchMaskingJobSourceEnginesResponse {
	return v.value
}

func (v *NullableSearchMaskingJobSourceEnginesResponse) Set(val *SearchMaskingJobSourceEnginesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMaskingJobSourceEnginesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMaskingJobSourceEnginesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMaskingJobSourceEnginesResponse(val *SearchMaskingJobSourceEnginesResponse) *NullableSearchMaskingJobSourceEnginesResponse {
	return &NullableSearchMaskingJobSourceEnginesResponse{value: val, isSet: true}
}

func (v NullableSearchMaskingJobSourceEnginesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMaskingJobSourceEnginesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


