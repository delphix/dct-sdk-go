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

// checks if the GetEngineGlobalObjectStateReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEngineGlobalObjectStateReportResponse{}

// GetEngineGlobalObjectStateReportResponse struct for GetEngineGlobalObjectStateReportResponse
type GetEngineGlobalObjectStateReportResponse struct {
	Items []EngineGlobalObjectStateData `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewGetEngineGlobalObjectStateReportResponse instantiates a new GetEngineGlobalObjectStateReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEngineGlobalObjectStateReportResponse() *GetEngineGlobalObjectStateReportResponse {
	this := GetEngineGlobalObjectStateReportResponse{}
	return &this
}

// NewGetEngineGlobalObjectStateReportResponseWithDefaults instantiates a new GetEngineGlobalObjectStateReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEngineGlobalObjectStateReportResponseWithDefaults() *GetEngineGlobalObjectStateReportResponse {
	this := GetEngineGlobalObjectStateReportResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetEngineGlobalObjectStateReportResponse) GetItems() []EngineGlobalObjectStateData {
	if o == nil || IsNil(o.Items) {
		var ret []EngineGlobalObjectStateData
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEngineGlobalObjectStateReportResponse) GetItemsOk() ([]EngineGlobalObjectStateData, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetEngineGlobalObjectStateReportResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []EngineGlobalObjectStateData and assigns it to the Items field.
func (o *GetEngineGlobalObjectStateReportResponse) SetItems(v []EngineGlobalObjectStateData) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *GetEngineGlobalObjectStateReportResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEngineGlobalObjectStateReportResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *GetEngineGlobalObjectStateReportResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *GetEngineGlobalObjectStateReportResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o GetEngineGlobalObjectStateReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEngineGlobalObjectStateReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableGetEngineGlobalObjectStateReportResponse struct {
	value *GetEngineGlobalObjectStateReportResponse
	isSet bool
}

func (v NullableGetEngineGlobalObjectStateReportResponse) Get() *GetEngineGlobalObjectStateReportResponse {
	return v.value
}

func (v *NullableGetEngineGlobalObjectStateReportResponse) Set(val *GetEngineGlobalObjectStateReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEngineGlobalObjectStateReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEngineGlobalObjectStateReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEngineGlobalObjectStateReportResponse(val *GetEngineGlobalObjectStateReportResponse) *NullableGetEngineGlobalObjectStateReportResponse {
	return &NullableGetEngineGlobalObjectStateReportResponse{value: val, isSet: true}
}

func (v NullableGetEngineGlobalObjectStateReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEngineGlobalObjectStateReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


