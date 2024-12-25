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

// checks if the LicenseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseInfo{}

// LicenseInfo Information about the current DCT license.
type LicenseInfo struct {
	Tier *LicenseTier `json:"tier,omitempty"`
	// The number of virtualization engines counting against the limit.
	VirtualizationEngineCount NullableInt32 `json:"virtualization_engine_count,omitempty"`
	// The number of masking engines counting against the limit. Masking engines added to a Hyperscale Instance's pool do not count against the limit.
	MaskingEngineCount NullableInt32 `json:"masking_engine_count,omitempty"`
	// The maximum number of registered virtualization engines allowed for the current license tier.
	VirtualizationEngineCountLimit NullableInt32 `json:"virtualization_engine_count_limit,omitempty"`
	// The maximum number of registered masking engines allowed for the current license tier. Masking engines added to a Hyperscale Instance's pool do not count against the limit.
	MaskingEngineCountLimit NullableInt32 `json:"masking_engine_count_limit,omitempty"`
}

// NewLicenseInfo instantiates a new LicenseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseInfo() *LicenseInfo {
	this := LicenseInfo{}
	return &this
}

// NewLicenseInfoWithDefaults instantiates a new LicenseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseInfoWithDefaults() *LicenseInfo {
	this := LicenseInfo{}
	return &this
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *LicenseInfo) GetTier() LicenseTier {
	if o == nil || IsNil(o.Tier) {
		var ret LicenseTier
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseInfo) GetTierOk() (*LicenseTier, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *LicenseInfo) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given LicenseTier and assigns it to the Tier field.
func (o *LicenseInfo) SetTier(v LicenseTier) {
	o.Tier = &v
}

// GetVirtualizationEngineCount returns the VirtualizationEngineCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfo) GetVirtualizationEngineCount() int32 {
	if o == nil || IsNil(o.VirtualizationEngineCount.Get()) {
		var ret int32
		return ret
	}
	return *o.VirtualizationEngineCount.Get()
}

// GetVirtualizationEngineCountOk returns a tuple with the VirtualizationEngineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseInfo) GetVirtualizationEngineCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualizationEngineCount.Get(), o.VirtualizationEngineCount.IsSet()
}

// HasVirtualizationEngineCount returns a boolean if a field has been set.
func (o *LicenseInfo) HasVirtualizationEngineCount() bool {
	if o != nil && o.VirtualizationEngineCount.IsSet() {
		return true
	}

	return false
}

// SetVirtualizationEngineCount gets a reference to the given NullableInt32 and assigns it to the VirtualizationEngineCount field.
func (o *LicenseInfo) SetVirtualizationEngineCount(v int32) {
	o.VirtualizationEngineCount.Set(&v)
}
// SetVirtualizationEngineCountNil sets the value for VirtualizationEngineCount to be an explicit nil
func (o *LicenseInfo) SetVirtualizationEngineCountNil() {
	o.VirtualizationEngineCount.Set(nil)
}

// UnsetVirtualizationEngineCount ensures that no value is present for VirtualizationEngineCount, not even an explicit nil
func (o *LicenseInfo) UnsetVirtualizationEngineCount() {
	o.VirtualizationEngineCount.Unset()
}

// GetMaskingEngineCount returns the MaskingEngineCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfo) GetMaskingEngineCount() int32 {
	if o == nil || IsNil(o.MaskingEngineCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaskingEngineCount.Get()
}

// GetMaskingEngineCountOk returns a tuple with the MaskingEngineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseInfo) GetMaskingEngineCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskingEngineCount.Get(), o.MaskingEngineCount.IsSet()
}

// HasMaskingEngineCount returns a boolean if a field has been set.
func (o *LicenseInfo) HasMaskingEngineCount() bool {
	if o != nil && o.MaskingEngineCount.IsSet() {
		return true
	}

	return false
}

// SetMaskingEngineCount gets a reference to the given NullableInt32 and assigns it to the MaskingEngineCount field.
func (o *LicenseInfo) SetMaskingEngineCount(v int32) {
	o.MaskingEngineCount.Set(&v)
}
// SetMaskingEngineCountNil sets the value for MaskingEngineCount to be an explicit nil
func (o *LicenseInfo) SetMaskingEngineCountNil() {
	o.MaskingEngineCount.Set(nil)
}

// UnsetMaskingEngineCount ensures that no value is present for MaskingEngineCount, not even an explicit nil
func (o *LicenseInfo) UnsetMaskingEngineCount() {
	o.MaskingEngineCount.Unset()
}

// GetVirtualizationEngineCountLimit returns the VirtualizationEngineCountLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfo) GetVirtualizationEngineCountLimit() int32 {
	if o == nil || IsNil(o.VirtualizationEngineCountLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.VirtualizationEngineCountLimit.Get()
}

// GetVirtualizationEngineCountLimitOk returns a tuple with the VirtualizationEngineCountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseInfo) GetVirtualizationEngineCountLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualizationEngineCountLimit.Get(), o.VirtualizationEngineCountLimit.IsSet()
}

// HasVirtualizationEngineCountLimit returns a boolean if a field has been set.
func (o *LicenseInfo) HasVirtualizationEngineCountLimit() bool {
	if o != nil && o.VirtualizationEngineCountLimit.IsSet() {
		return true
	}

	return false
}

// SetVirtualizationEngineCountLimit gets a reference to the given NullableInt32 and assigns it to the VirtualizationEngineCountLimit field.
func (o *LicenseInfo) SetVirtualizationEngineCountLimit(v int32) {
	o.VirtualizationEngineCountLimit.Set(&v)
}
// SetVirtualizationEngineCountLimitNil sets the value for VirtualizationEngineCountLimit to be an explicit nil
func (o *LicenseInfo) SetVirtualizationEngineCountLimitNil() {
	o.VirtualizationEngineCountLimit.Set(nil)
}

// UnsetVirtualizationEngineCountLimit ensures that no value is present for VirtualizationEngineCountLimit, not even an explicit nil
func (o *LicenseInfo) UnsetVirtualizationEngineCountLimit() {
	o.VirtualizationEngineCountLimit.Unset()
}

// GetMaskingEngineCountLimit returns the MaskingEngineCountLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfo) GetMaskingEngineCountLimit() int32 {
	if o == nil || IsNil(o.MaskingEngineCountLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.MaskingEngineCountLimit.Get()
}

// GetMaskingEngineCountLimitOk returns a tuple with the MaskingEngineCountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseInfo) GetMaskingEngineCountLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskingEngineCountLimit.Get(), o.MaskingEngineCountLimit.IsSet()
}

// HasMaskingEngineCountLimit returns a boolean if a field has been set.
func (o *LicenseInfo) HasMaskingEngineCountLimit() bool {
	if o != nil && o.MaskingEngineCountLimit.IsSet() {
		return true
	}

	return false
}

// SetMaskingEngineCountLimit gets a reference to the given NullableInt32 and assigns it to the MaskingEngineCountLimit field.
func (o *LicenseInfo) SetMaskingEngineCountLimit(v int32) {
	o.MaskingEngineCountLimit.Set(&v)
}
// SetMaskingEngineCountLimitNil sets the value for MaskingEngineCountLimit to be an explicit nil
func (o *LicenseInfo) SetMaskingEngineCountLimitNil() {
	o.MaskingEngineCountLimit.Set(nil)
}

// UnsetMaskingEngineCountLimit ensures that no value is present for MaskingEngineCountLimit, not even an explicit nil
func (o *LicenseInfo) UnsetMaskingEngineCountLimit() {
	o.MaskingEngineCountLimit.Unset()
}

func (o LicenseInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if o.VirtualizationEngineCount.IsSet() {
		toSerialize["virtualization_engine_count"] = o.VirtualizationEngineCount.Get()
	}
	if o.MaskingEngineCount.IsSet() {
		toSerialize["masking_engine_count"] = o.MaskingEngineCount.Get()
	}
	if o.VirtualizationEngineCountLimit.IsSet() {
		toSerialize["virtualization_engine_count_limit"] = o.VirtualizationEngineCountLimit.Get()
	}
	if o.MaskingEngineCountLimit.IsSet() {
		toSerialize["masking_engine_count_limit"] = o.MaskingEngineCountLimit.Get()
	}
	return toSerialize, nil
}

type NullableLicenseInfo struct {
	value *LicenseInfo
	isSet bool
}

func (v NullableLicenseInfo) Get() *LicenseInfo {
	return v.value
}

func (v *NullableLicenseInfo) Set(val *LicenseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseInfo(val *LicenseInfo) *NullableLicenseInfo {
	return &NullableLicenseInfo{value: val, isSet: true}
}

func (v NullableLicenseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


