/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ListReportingScheduleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReportingScheduleResponse{}

// ListReportingScheduleResponse struct for ListReportingScheduleResponse
type ListReportingScheduleResponse struct {
	Items []ReportingSchedule `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListReportingScheduleResponse instantiates a new ListReportingScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReportingScheduleResponse() *ListReportingScheduleResponse {
	this := ListReportingScheduleResponse{}
	return &this
}

// NewListReportingScheduleResponseWithDefaults instantiates a new ListReportingScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReportingScheduleResponseWithDefaults() *ListReportingScheduleResponse {
	this := ListReportingScheduleResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListReportingScheduleResponse) GetItems() []ReportingSchedule {
	if o == nil || IsNil(o.Items) {
		var ret []ReportingSchedule
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReportingScheduleResponse) GetItemsOk() ([]ReportingSchedule, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListReportingScheduleResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ReportingSchedule and assigns it to the Items field.
func (o *ListReportingScheduleResponse) SetItems(v []ReportingSchedule) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListReportingScheduleResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReportingScheduleResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListReportingScheduleResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListReportingScheduleResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListReportingScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListReportingScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListReportingScheduleResponse struct {
	value *ListReportingScheduleResponse
	isSet bool
}

func (v NullableListReportingScheduleResponse) Get() *ListReportingScheduleResponse {
	return v.value
}

func (v *NullableListReportingScheduleResponse) Set(val *ListReportingScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListReportingScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListReportingScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReportingScheduleResponse(val *ListReportingScheduleResponse) *NullableListReportingScheduleResponse {
	return &NullableListReportingScheduleResponse{value: val, isSet: true}
}

func (v NullableListReportingScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReportingScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


