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

// checks if the ListRolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRolesResponse{}

// ListRolesResponse struct for ListRolesResponse
type ListRolesResponse struct {
	Items []Role `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListRolesResponse instantiates a new ListRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRolesResponse() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// NewListRolesResponseWithDefaults instantiates a new ListRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRolesResponseWithDefaults() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListRolesResponse) GetItems() []Role {
	if o == nil || IsNil(o.Items) {
		var ret []Role
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetItemsOk() ([]Role, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListRolesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Role and assigns it to the Items field.
func (o *ListRolesResponse) SetItems(v []Role) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListRolesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListRolesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListRolesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListRolesResponse struct {
	value *ListRolesResponse
	isSet bool
}

func (v NullableListRolesResponse) Get() *ListRolesResponse {
	return v.value
}

func (v *NullableListRolesResponse) Set(val *ListRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRolesResponse(val *ListRolesResponse) *NullableListRolesResponse {
	return &NullableListRolesResponse{value: val, isSet: true}
}

func (v NullableListRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


