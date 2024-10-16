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

// checks if the ListMaskingFileConsumersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMaskingFileConsumersResponse{}

// ListMaskingFileConsumersResponse struct for ListMaskingFileConsumersResponse
type ListMaskingFileConsumersResponse struct {
	Items []Consumer `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListMaskingFileConsumersResponse instantiates a new ListMaskingFileConsumersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMaskingFileConsumersResponse() *ListMaskingFileConsumersResponse {
	this := ListMaskingFileConsumersResponse{}
	return &this
}

// NewListMaskingFileConsumersResponseWithDefaults instantiates a new ListMaskingFileConsumersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMaskingFileConsumersResponseWithDefaults() *ListMaskingFileConsumersResponse {
	this := ListMaskingFileConsumersResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListMaskingFileConsumersResponse) GetItems() []Consumer {
	if o == nil || IsNil(o.Items) {
		var ret []Consumer
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMaskingFileConsumersResponse) GetItemsOk() ([]Consumer, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListMaskingFileConsumersResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Consumer and assigns it to the Items field.
func (o *ListMaskingFileConsumersResponse) SetItems(v []Consumer) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListMaskingFileConsumersResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMaskingFileConsumersResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListMaskingFileConsumersResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListMaskingFileConsumersResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListMaskingFileConsumersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMaskingFileConsumersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListMaskingFileConsumersResponse struct {
	value *ListMaskingFileConsumersResponse
	isSet bool
}

func (v NullableListMaskingFileConsumersResponse) Get() *ListMaskingFileConsumersResponse {
	return v.value
}

func (v *NullableListMaskingFileConsumersResponse) Set(val *ListMaskingFileConsumersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMaskingFileConsumersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMaskingFileConsumersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMaskingFileConsumersResponse(val *ListMaskingFileConsumersResponse) *NullableListMaskingFileConsumersResponse {
	return &NullableListMaskingFileConsumersResponse{value: val, isSet: true}
}

func (v NullableListMaskingFileConsumersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMaskingFileConsumersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


