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

// checks if the VDBInventoryReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VDBInventoryReportResponse{}

// VDBInventoryReportResponse struct for VDBInventoryReportResponse
type VDBInventoryReportResponse struct {
	Items []VDBInventoryData `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewVDBInventoryReportResponse instantiates a new VDBInventoryReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVDBInventoryReportResponse() *VDBInventoryReportResponse {
	this := VDBInventoryReportResponse{}
	return &this
}

// NewVDBInventoryReportResponseWithDefaults instantiates a new VDBInventoryReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVDBInventoryReportResponseWithDefaults() *VDBInventoryReportResponse {
	this := VDBInventoryReportResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *VDBInventoryReportResponse) GetItems() []VDBInventoryData {
	if o == nil || IsNil(o.Items) {
		var ret []VDBInventoryData
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryReportResponse) GetItemsOk() ([]VDBInventoryData, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *VDBInventoryReportResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VDBInventoryData and assigns it to the Items field.
func (o *VDBInventoryReportResponse) SetItems(v []VDBInventoryData) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *VDBInventoryReportResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryReportResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *VDBInventoryReportResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *VDBInventoryReportResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o VDBInventoryReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VDBInventoryReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableVDBInventoryReportResponse struct {
	value *VDBInventoryReportResponse
	isSet bool
}

func (v NullableVDBInventoryReportResponse) Get() *VDBInventoryReportResponse {
	return v.value
}

func (v *NullableVDBInventoryReportResponse) Set(val *VDBInventoryReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVDBInventoryReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVDBInventoryReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVDBInventoryReportResponse(val *VDBInventoryReportResponse) *NullableVDBInventoryReportResponse {
	return &NullableVDBInventoryReportResponse{value: val, isSet: true}
}

func (v NullableVDBInventoryReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVDBInventoryReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


