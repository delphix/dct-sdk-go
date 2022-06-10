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

// ListReportingScheduleResponse struct for ListReportingScheduleResponse
type ListReportingScheduleResponse struct {
	Items []ReportingSchedule `json:"items,omitempty"`
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
	if o == nil || o.Items == nil {
		var ret []ReportingSchedule
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReportingScheduleResponse) GetItemsOk() ([]ReportingSchedule, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListReportingScheduleResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ReportingSchedule and assigns it to the Items field.
func (o *ListReportingScheduleResponse) SetItems(v []ReportingSchedule) {
	o.Items = v
}

func (o ListReportingScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
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

