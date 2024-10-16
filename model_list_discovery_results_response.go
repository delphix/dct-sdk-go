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

// checks if the ListDiscoveryResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDiscoveryResultsResponse{}

// ListDiscoveryResultsResponse struct for ListDiscoveryResultsResponse
type ListDiscoveryResultsResponse struct {
	Items []DiscoveryResult `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListDiscoveryResultsResponse instantiates a new ListDiscoveryResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDiscoveryResultsResponse() *ListDiscoveryResultsResponse {
	this := ListDiscoveryResultsResponse{}
	return &this
}

// NewListDiscoveryResultsResponseWithDefaults instantiates a new ListDiscoveryResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDiscoveryResultsResponseWithDefaults() *ListDiscoveryResultsResponse {
	this := ListDiscoveryResultsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListDiscoveryResultsResponse) GetItems() []DiscoveryResult {
	if o == nil || IsNil(o.Items) {
		var ret []DiscoveryResult
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiscoveryResultsResponse) GetItemsOk() ([]DiscoveryResult, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListDiscoveryResultsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DiscoveryResult and assigns it to the Items field.
func (o *ListDiscoveryResultsResponse) SetItems(v []DiscoveryResult) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListDiscoveryResultsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiscoveryResultsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListDiscoveryResultsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListDiscoveryResultsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListDiscoveryResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDiscoveryResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListDiscoveryResultsResponse struct {
	value *ListDiscoveryResultsResponse
	isSet bool
}

func (v NullableListDiscoveryResultsResponse) Get() *ListDiscoveryResultsResponse {
	return v.value
}

func (v *NullableListDiscoveryResultsResponse) Set(val *ListDiscoveryResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDiscoveryResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDiscoveryResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDiscoveryResultsResponse(val *ListDiscoveryResultsResponse) *NullableListDiscoveryResultsResponse {
	return &NullableListDiscoveryResultsResponse{value: val, isSet: true}
}

func (v NullableListDiscoveryResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDiscoveryResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


