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

// checks if the SearchVirtualizationActionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchVirtualizationActionsResponse{}

// SearchVirtualizationActionsResponse struct for SearchVirtualizationActionsResponse
type SearchVirtualizationActionsResponse struct {
	Items []VirtualizationAction `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchVirtualizationActionsResponse instantiates a new SearchVirtualizationActionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchVirtualizationActionsResponse() *SearchVirtualizationActionsResponse {
	this := SearchVirtualizationActionsResponse{}
	return &this
}

// NewSearchVirtualizationActionsResponseWithDefaults instantiates a new SearchVirtualizationActionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchVirtualizationActionsResponseWithDefaults() *SearchVirtualizationActionsResponse {
	this := SearchVirtualizationActionsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchVirtualizationActionsResponse) GetItems() []VirtualizationAction {
	if o == nil || IsNil(o.Items) {
		var ret []VirtualizationAction
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVirtualizationActionsResponse) GetItemsOk() ([]VirtualizationAction, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchVirtualizationActionsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VirtualizationAction and assigns it to the Items field.
func (o *SearchVirtualizationActionsResponse) SetItems(v []VirtualizationAction) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchVirtualizationActionsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchVirtualizationActionsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchVirtualizationActionsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchVirtualizationActionsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchVirtualizationActionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchVirtualizationActionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchVirtualizationActionsResponse struct {
	value *SearchVirtualizationActionsResponse
	isSet bool
}

func (v NullableSearchVirtualizationActionsResponse) Get() *SearchVirtualizationActionsResponse {
	return v.value
}

func (v *NullableSearchVirtualizationActionsResponse) Set(val *SearchVirtualizationActionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchVirtualizationActionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchVirtualizationActionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchVirtualizationActionsResponse(val *SearchVirtualizationActionsResponse) *NullableSearchVirtualizationActionsResponse {
	return &NullableSearchVirtualizationActionsResponse{value: val, isSet: true}
}

func (v NullableSearchVirtualizationActionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchVirtualizationActionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


