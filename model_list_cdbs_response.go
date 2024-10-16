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

// checks if the ListCDBsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCDBsResponse{}

// ListCDBsResponse struct for ListCDBsResponse
type ListCDBsResponse struct {
	Items []CDB `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListCDBsResponse instantiates a new ListCDBsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCDBsResponse() *ListCDBsResponse {
	this := ListCDBsResponse{}
	return &this
}

// NewListCDBsResponseWithDefaults instantiates a new ListCDBsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCDBsResponseWithDefaults() *ListCDBsResponse {
	this := ListCDBsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListCDBsResponse) GetItems() []CDB {
	if o == nil || IsNil(o.Items) {
		var ret []CDB
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCDBsResponse) GetItemsOk() ([]CDB, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListCDBsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CDB and assigns it to the Items field.
func (o *ListCDBsResponse) SetItems(v []CDB) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListCDBsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCDBsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListCDBsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListCDBsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListCDBsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCDBsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListCDBsResponse struct {
	value *ListCDBsResponse
	isSet bool
}

func (v NullableListCDBsResponse) Get() *ListCDBsResponse {
	return v.value
}

func (v *NullableListCDBsResponse) Set(val *ListCDBsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCDBsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCDBsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCDBsResponse(val *ListCDBsResponse) *NullableListCDBsResponse {
	return &NullableListCDBsResponse{value: val, isSet: true}
}

func (v NullableListCDBsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCDBsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


