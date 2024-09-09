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

// checks if the ListKerberosConfigsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListKerberosConfigsResponse{}

// ListKerberosConfigsResponse struct for ListKerberosConfigsResponse
type ListKerberosConfigsResponse struct {
	Items []KerberosConfig `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListKerberosConfigsResponse instantiates a new ListKerberosConfigsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListKerberosConfigsResponse() *ListKerberosConfigsResponse {
	this := ListKerberosConfigsResponse{}
	return &this
}

// NewListKerberosConfigsResponseWithDefaults instantiates a new ListKerberosConfigsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListKerberosConfigsResponseWithDefaults() *ListKerberosConfigsResponse {
	this := ListKerberosConfigsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListKerberosConfigsResponse) GetItems() []KerberosConfig {
	if o == nil || IsNil(o.Items) {
		var ret []KerberosConfig
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListKerberosConfigsResponse) GetItemsOk() ([]KerberosConfig, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListKerberosConfigsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []KerberosConfig and assigns it to the Items field.
func (o *ListKerberosConfigsResponse) SetItems(v []KerberosConfig) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListKerberosConfigsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListKerberosConfigsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListKerberosConfigsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListKerberosConfigsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListKerberosConfigsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListKerberosConfigsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListKerberosConfigsResponse struct {
	value *ListKerberosConfigsResponse
	isSet bool
}

func (v NullableListKerberosConfigsResponse) Get() *ListKerberosConfigsResponse {
	return v.value
}

func (v *NullableListKerberosConfigsResponse) Set(val *ListKerberosConfigsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListKerberosConfigsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListKerberosConfigsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListKerberosConfigsResponse(val *ListKerberosConfigsResponse) *NullableListKerberosConfigsResponse {
	return &NullableListKerberosConfigsResponse{value: val, isSet: true}
}

func (v NullableListKerberosConfigsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListKerberosConfigsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


