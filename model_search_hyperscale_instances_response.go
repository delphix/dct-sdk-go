/*
Delphix DCT API

Delphix DCT API

API version: 3.5.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the SearchHyperscaleInstancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchHyperscaleInstancesResponse{}

// SearchHyperscaleInstancesResponse struct for SearchHyperscaleInstancesResponse
type SearchHyperscaleInstancesResponse struct {
	Items []HyperscaleInstance `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchHyperscaleInstancesResponse instantiates a new SearchHyperscaleInstancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHyperscaleInstancesResponse() *SearchHyperscaleInstancesResponse {
	this := SearchHyperscaleInstancesResponse{}
	return &this
}

// NewSearchHyperscaleInstancesResponseWithDefaults instantiates a new SearchHyperscaleInstancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchHyperscaleInstancesResponseWithDefaults() *SearchHyperscaleInstancesResponse {
	this := SearchHyperscaleInstancesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchHyperscaleInstancesResponse) GetItems() []HyperscaleInstance {
	if o == nil || IsNil(o.Items) {
		var ret []HyperscaleInstance
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHyperscaleInstancesResponse) GetItemsOk() ([]HyperscaleInstance, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchHyperscaleInstancesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []HyperscaleInstance and assigns it to the Items field.
func (o *SearchHyperscaleInstancesResponse) SetItems(v []HyperscaleInstance) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchHyperscaleInstancesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHyperscaleInstancesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchHyperscaleInstancesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchHyperscaleInstancesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchHyperscaleInstancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchHyperscaleInstancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchHyperscaleInstancesResponse struct {
	value *SearchHyperscaleInstancesResponse
	isSet bool
}

func (v NullableSearchHyperscaleInstancesResponse) Get() *SearchHyperscaleInstancesResponse {
	return v.value
}

func (v *NullableSearchHyperscaleInstancesResponse) Set(val *SearchHyperscaleInstancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHyperscaleInstancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHyperscaleInstancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHyperscaleInstancesResponse(val *SearchHyperscaleInstancesResponse) *NullableSearchHyperscaleInstancesResponse {
	return &NullableSearchHyperscaleInstancesResponse{value: val, isSet: true}
}

func (v NullableSearchHyperscaleInstancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHyperscaleInstancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

