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

// ListApiClientsResponse struct for ListApiClientsResponse
type ListApiClientsResponse struct {
	Items []ApiClient `json:"items,omitempty"`
}

// NewListApiClientsResponse instantiates a new ListApiClientsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApiClientsResponse() *ListApiClientsResponse {
	this := ListApiClientsResponse{}
	return &this
}

// NewListApiClientsResponseWithDefaults instantiates a new ListApiClientsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApiClientsResponseWithDefaults() *ListApiClientsResponse {
	this := ListApiClientsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListApiClientsResponse) GetItems() []ApiClient {
	if o == nil || o.Items == nil {
		var ret []ApiClient
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApiClientsResponse) GetItemsOk() ([]ApiClient, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListApiClientsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ApiClient and assigns it to the Items field.
func (o *ListApiClientsResponse) SetItems(v []ApiClient) {
	o.Items = v
}

func (o ListApiClientsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableListApiClientsResponse struct {
	value *ListApiClientsResponse
	isSet bool
}

func (v NullableListApiClientsResponse) Get() *ListApiClientsResponse {
	return v.value
}

func (v *NullableListApiClientsResponse) Set(val *ListApiClientsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListApiClientsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListApiClientsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApiClientsResponse(val *ListApiClientsResponse) *NullableListApiClientsResponse {
	return &NullableListApiClientsResponse{value: val, isSet: true}
}

func (v NullableListApiClientsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApiClientsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


