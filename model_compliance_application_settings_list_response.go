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

// checks if the ComplianceApplicationSettingsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComplianceApplicationSettingsListResponse{}

// ComplianceApplicationSettingsListResponse struct for ComplianceApplicationSettingsListResponse
type ComplianceApplicationSettingsListResponse struct {
	Items []ComplianceApplicationSetting `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewComplianceApplicationSettingsListResponse instantiates a new ComplianceApplicationSettingsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplianceApplicationSettingsListResponse() *ComplianceApplicationSettingsListResponse {
	this := ComplianceApplicationSettingsListResponse{}
	return &this
}

// NewComplianceApplicationSettingsListResponseWithDefaults instantiates a new ComplianceApplicationSettingsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceApplicationSettingsListResponseWithDefaults() *ComplianceApplicationSettingsListResponse {
	this := ComplianceApplicationSettingsListResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ComplianceApplicationSettingsListResponse) GetItems() []ComplianceApplicationSetting {
	if o == nil || IsNil(o.Items) {
		var ret []ComplianceApplicationSetting
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceApplicationSettingsListResponse) GetItemsOk() ([]ComplianceApplicationSetting, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ComplianceApplicationSettingsListResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ComplianceApplicationSetting and assigns it to the Items field.
func (o *ComplianceApplicationSettingsListResponse) SetItems(v []ComplianceApplicationSetting) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ComplianceApplicationSettingsListResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceApplicationSettingsListResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ComplianceApplicationSettingsListResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ComplianceApplicationSettingsListResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ComplianceApplicationSettingsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComplianceApplicationSettingsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableComplianceApplicationSettingsListResponse struct {
	value *ComplianceApplicationSettingsListResponse
	isSet bool
}

func (v NullableComplianceApplicationSettingsListResponse) Get() *ComplianceApplicationSettingsListResponse {
	return v.value
}

func (v *NullableComplianceApplicationSettingsListResponse) Set(val *ComplianceApplicationSettingsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableComplianceApplicationSettingsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableComplianceApplicationSettingsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplianceApplicationSettingsListResponse(val *ComplianceApplicationSettingsListResponse) *NullableComplianceApplicationSettingsListResponse {
	return &NullableComplianceApplicationSettingsListResponse{value: val, isSet: true}
}

func (v NullableComplianceApplicationSettingsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplianceApplicationSettingsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


