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

// checks if the SearchAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAccountsResponse{}

// SearchAccountsResponse struct for SearchAccountsResponse
type SearchAccountsResponse struct {
	Items []Account `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchAccountsResponse instantiates a new SearchAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAccountsResponse() *SearchAccountsResponse {
	this := SearchAccountsResponse{}
	return &this
}

// NewSearchAccountsResponseWithDefaults instantiates a new SearchAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAccountsResponseWithDefaults() *SearchAccountsResponse {
	this := SearchAccountsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchAccountsResponse) GetItems() []Account {
	if o == nil || IsNil(o.Items) {
		var ret []Account
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAccountsResponse) GetItemsOk() ([]Account, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchAccountsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Account and assigns it to the Items field.
func (o *SearchAccountsResponse) SetItems(v []Account) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchAccountsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAccountsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchAccountsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchAccountsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchAccountsResponse struct {
	value *SearchAccountsResponse
	isSet bool
}

func (v NullableSearchAccountsResponse) Get() *SearchAccountsResponse {
	return v.value
}

func (v *NullableSearchAccountsResponse) Set(val *SearchAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAccountsResponse(val *SearchAccountsResponse) *NullableSearchAccountsResponse {
	return &NullableSearchAccountsResponse{value: val, isSet: true}
}

func (v NullableSearchAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


