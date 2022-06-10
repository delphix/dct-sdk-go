/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// ListRegisteredEnginesResponse struct for ListRegisteredEnginesResponse
type ListRegisteredEnginesResponse struct {
	Items []RegisteredEngine `json:"items,omitempty"`
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
	if o == nil || o.Items == nil {
		var ret []RegisteredEngine
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegisteredEnginesResponse) GetItemsOk() ([]RegisteredEngine, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListRegisteredEnginesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RegisteredEngine and assigns it to the Items field.
func (o *ListRegisteredEnginesResponse) SetItems(v []RegisteredEngine) {
	o.Items = v
}

func (o ListRegisteredEnginesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
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


