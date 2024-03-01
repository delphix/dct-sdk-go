/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ListRegisteredEnginesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRegisteredEnginesResponse{}

// ListRegisteredEnginesResponse struct for ListRegisteredEnginesResponse
type ListRegisteredEnginesResponse struct {
	Items []RegisteredEngine `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListRegisteredEnginesResponse instantiates a new ListRegisteredEnginesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRegisteredEnginesResponse() *ListRegisteredEnginesResponse {
	this := ListRegisteredEnginesResponse{}
	return &this
}

// NewListRegisteredEnginesResponseWithDefaults instantiates a new ListRegisteredEnginesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRegisteredEnginesResponseWithDefaults() *ListRegisteredEnginesResponse {
	this := ListRegisteredEnginesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListRegisteredEnginesResponse) GetItems() []RegisteredEngine {
	if o == nil || IsNil(o.Items) {
		var ret []RegisteredEngine
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegisteredEnginesResponse) GetItemsOk() ([]RegisteredEngine, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListRegisteredEnginesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RegisteredEngine and assigns it to the Items field.
func (o *ListRegisteredEnginesResponse) SetItems(v []RegisteredEngine) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListRegisteredEnginesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegisteredEnginesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListRegisteredEnginesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListRegisteredEnginesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListRegisteredEnginesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRegisteredEnginesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListRegisteredEnginesResponse struct {
	value *ListRegisteredEnginesResponse
	isSet bool
}

func (v NullableListRegisteredEnginesResponse) Get() *ListRegisteredEnginesResponse {
	return v.value
}

func (v *NullableListRegisteredEnginesResponse) Set(val *ListRegisteredEnginesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRegisteredEnginesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRegisteredEnginesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRegisteredEnginesResponse(val *ListRegisteredEnginesResponse) *NullableListRegisteredEnginesResponse {
	return &NullableListRegisteredEnginesResponse{value: val, isSet: true}
}

func (v NullableListRegisteredEnginesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRegisteredEnginesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


