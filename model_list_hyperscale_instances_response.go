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

// checks if the ListHyperscaleInstancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHyperscaleInstancesResponse{}

// ListHyperscaleInstancesResponse struct for ListHyperscaleInstancesResponse
type ListHyperscaleInstancesResponse struct {
	Items []HyperscaleInstance `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListHyperscaleInstancesResponse instantiates a new ListHyperscaleInstancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHyperscaleInstancesResponse() *ListHyperscaleInstancesResponse {
	this := ListHyperscaleInstancesResponse{}
	return &this
}

// NewListHyperscaleInstancesResponseWithDefaults instantiates a new ListHyperscaleInstancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHyperscaleInstancesResponseWithDefaults() *ListHyperscaleInstancesResponse {
	this := ListHyperscaleInstancesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListHyperscaleInstancesResponse) GetItems() []HyperscaleInstance {
	if o == nil || IsNil(o.Items) {
		var ret []HyperscaleInstance
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHyperscaleInstancesResponse) GetItemsOk() ([]HyperscaleInstance, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListHyperscaleInstancesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []HyperscaleInstance and assigns it to the Items field.
func (o *ListHyperscaleInstancesResponse) SetItems(v []HyperscaleInstance) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListHyperscaleInstancesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHyperscaleInstancesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListHyperscaleInstancesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListHyperscaleInstancesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListHyperscaleInstancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHyperscaleInstancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListHyperscaleInstancesResponse struct {
	value *ListHyperscaleInstancesResponse
	isSet bool
}

func (v NullableListHyperscaleInstancesResponse) Get() *ListHyperscaleInstancesResponse {
	return v.value
}

func (v *NullableListHyperscaleInstancesResponse) Set(val *ListHyperscaleInstancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListHyperscaleInstancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListHyperscaleInstancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHyperscaleInstancesResponse(val *ListHyperscaleInstancesResponse) *NullableListHyperscaleInstancesResponse {
	return &NullableListHyperscaleInstancesResponse{value: val, isSet: true}
}

func (v NullableListHyperscaleInstancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHyperscaleInstancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


