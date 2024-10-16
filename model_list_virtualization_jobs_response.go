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

// checks if the ListVirtualizationJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVirtualizationJobsResponse{}

// ListVirtualizationJobsResponse struct for ListVirtualizationJobsResponse
type ListVirtualizationJobsResponse struct {
	Items []VirtualizationJob `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListVirtualizationJobsResponse instantiates a new ListVirtualizationJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVirtualizationJobsResponse() *ListVirtualizationJobsResponse {
	this := ListVirtualizationJobsResponse{}
	return &this
}

// NewListVirtualizationJobsResponseWithDefaults instantiates a new ListVirtualizationJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVirtualizationJobsResponseWithDefaults() *ListVirtualizationJobsResponse {
	this := ListVirtualizationJobsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListVirtualizationJobsResponse) GetItems() []VirtualizationJob {
	if o == nil || IsNil(o.Items) {
		var ret []VirtualizationJob
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVirtualizationJobsResponse) GetItemsOk() ([]VirtualizationJob, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListVirtualizationJobsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VirtualizationJob and assigns it to the Items field.
func (o *ListVirtualizationJobsResponse) SetItems(v []VirtualizationJob) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListVirtualizationJobsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVirtualizationJobsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListVirtualizationJobsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListVirtualizationJobsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListVirtualizationJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVirtualizationJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListVirtualizationJobsResponse struct {
	value *ListVirtualizationJobsResponse
	isSet bool
}

func (v NullableListVirtualizationJobsResponse) Get() *ListVirtualizationJobsResponse {
	return v.value
}

func (v *NullableListVirtualizationJobsResponse) Set(val *ListVirtualizationJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVirtualizationJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVirtualizationJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVirtualizationJobsResponse(val *ListVirtualizationJobsResponse) *NullableListVirtualizationJobsResponse {
	return &NullableListVirtualizationJobsResponse{value: val, isSet: true}
}

func (v NullableListVirtualizationJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVirtualizationJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


