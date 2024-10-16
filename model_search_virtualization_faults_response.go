/*
Delphix DCT API

Delphix DCT API

API version: 3.17.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the SearchVirtualizationFaultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchVirtualizationFaultsResponse{}

// SearchVirtualizationFaultsResponse struct for SearchVirtualizationFaultsResponse
type SearchVirtualizationFaultsResponse struct {
	Items []VirtualizationFault `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchVirtualizationFaultsResponse instantiates a new SearchVirtualizationFaultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchVirtualizationFaultsResponse() *SearchVirtualizationFaultsResponse {
	this := SearchVirtualizationFaultsResponse{}
	return &this
}

// NewSearchVirtualizationFaultsResponseWithDefaults instantiates a new SearchVirtualizationFaultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchVirtualizationFaultsResponseWithDefaults() *SearchVirtualizationFaultsResponse {
	this := SearchVirtualizationFaultsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchVirtualizationFaultsResponse) GetItems() []VirtualizationFault {
	if o == nil || IsNil(o.Items) {
		var ret []VirtualizationFault
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVirtualizationFaultsResponse) GetItemsOk() ([]VirtualizationFault, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchVirtualizationFaultsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VirtualizationFault and assigns it to the Items field.
func (o *SearchVirtualizationFaultsResponse) SetItems(v []VirtualizationFault) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchVirtualizationFaultsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVirtualizationFaultsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchVirtualizationFaultsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchVirtualizationFaultsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchVirtualizationFaultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchVirtualizationFaultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchVirtualizationFaultsResponse struct {
	value *SearchVirtualizationFaultsResponse
	isSet bool
}

func (v NullableSearchVirtualizationFaultsResponse) Get() *SearchVirtualizationFaultsResponse {
	return v.value
}

func (v *NullableSearchVirtualizationFaultsResponse) Set(val *SearchVirtualizationFaultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchVirtualizationFaultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchVirtualizationFaultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchVirtualizationFaultsResponse(val *SearchVirtualizationFaultsResponse) *NullableSearchVirtualizationFaultsResponse {
	return &NullableSearchVirtualizationFaultsResponse{value: val, isSet: true}
}

func (v NullableSearchVirtualizationFaultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchVirtualizationFaultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


