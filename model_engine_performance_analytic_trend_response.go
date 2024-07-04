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

// checks if the EnginePerformanceAnalyticTrendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnginePerformanceAnalyticTrendResponse{}

// EnginePerformanceAnalyticTrendResponse struct for EnginePerformanceAnalyticTrendResponse
type EnginePerformanceAnalyticTrendResponse struct {
	Items []EnginePerformanceAnalyticTrend `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewEnginePerformanceAnalyticTrendResponse instantiates a new EnginePerformanceAnalyticTrendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnginePerformanceAnalyticTrendResponse() *EnginePerformanceAnalyticTrendResponse {
	this := EnginePerformanceAnalyticTrendResponse{}
	return &this
}

// NewEnginePerformanceAnalyticTrendResponseWithDefaults instantiates a new EnginePerformanceAnalyticTrendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnginePerformanceAnalyticTrendResponseWithDefaults() *EnginePerformanceAnalyticTrendResponse {
	this := EnginePerformanceAnalyticTrendResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *EnginePerformanceAnalyticTrendResponse) GetItems() []EnginePerformanceAnalyticTrend {
	if o == nil || IsNil(o.Items) {
		var ret []EnginePerformanceAnalyticTrend
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnginePerformanceAnalyticTrendResponse) GetItemsOk() ([]EnginePerformanceAnalyticTrend, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *EnginePerformanceAnalyticTrendResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []EnginePerformanceAnalyticTrend and assigns it to the Items field.
func (o *EnginePerformanceAnalyticTrendResponse) SetItems(v []EnginePerformanceAnalyticTrend) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *EnginePerformanceAnalyticTrendResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnginePerformanceAnalyticTrendResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *EnginePerformanceAnalyticTrendResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *EnginePerformanceAnalyticTrendResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o EnginePerformanceAnalyticTrendResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnginePerformanceAnalyticTrendResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableEnginePerformanceAnalyticTrendResponse struct {
	value *EnginePerformanceAnalyticTrendResponse
	isSet bool
}

func (v NullableEnginePerformanceAnalyticTrendResponse) Get() *EnginePerformanceAnalyticTrendResponse {
	return v.value
}

func (v *NullableEnginePerformanceAnalyticTrendResponse) Set(val *EnginePerformanceAnalyticTrendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnginePerformanceAnalyticTrendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnginePerformanceAnalyticTrendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnginePerformanceAnalyticTrendResponse(val *EnginePerformanceAnalyticTrendResponse) *NullableEnginePerformanceAnalyticTrendResponse {
	return &NullableEnginePerformanceAnalyticTrendResponse{value: val, isSet: true}
}

func (v NullableEnginePerformanceAnalyticTrendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnginePerformanceAnalyticTrendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

