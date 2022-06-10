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
	"time"
)

// ProductHistory struct for ProductHistory
type ProductHistory struct {
	// Product Version.
	Version *string `json:"version,omitempty"`
	// This version installed on date.
	InstalledOn *time.Time `json:"installed_on,omitempty"`
}

// NewProductHistory instantiates a new ProductHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductHistory() *ProductHistory {
	this := ProductHistory{}
	return &this
}

// NewProductHistoryWithDefaults instantiates a new ProductHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductHistoryWithDefaults() *ProductHistory {
	this := ProductHistory{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ProductHistory) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductHistory) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ProductHistory) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ProductHistory) SetVersion(v string) {
	o.Version = &v
}

// GetInstalledOn returns the InstalledOn field value if set, zero value otherwise.
func (o *ProductHistory) GetInstalledOn() time.Time {
	if o == nil || o.InstalledOn == nil {
		var ret time.Time
		return ret
	}
	return *o.InstalledOn
}

// GetInstalledOnOk returns a tuple with the InstalledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductHistory) GetInstalledOnOk() (*time.Time, bool) {
	if o == nil || o.InstalledOn == nil {
		return nil, false
	}
	return o.InstalledOn, true
}

// HasInstalledOn returns a boolean if a field has been set.
func (o *ProductHistory) HasInstalledOn() bool {
	if o != nil && o.InstalledOn != nil {
		return true
	}

	return false
}

// SetInstalledOn gets a reference to the given time.Time and assigns it to the InstalledOn field.
func (o *ProductHistory) SetInstalledOn(v time.Time) {
	o.InstalledOn = &v
}

func (o ProductHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.InstalledOn != nil {
		toSerialize["installed_on"] = o.InstalledOn
	}
	return json.Marshal(toSerialize)
}

type NullableProductHistory struct {
	value *ProductHistory
	isSet bool
}

func (v NullableProductHistory) Get() *ProductHistory {
	return v.value
}

func (v *NullableProductHistory) Set(val *ProductHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableProductHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableProductHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductHistory(val *ProductHistory) *NullableProductHistory {
	return &NullableProductHistory{value: val, isSet: true}
}

func (v NullableProductHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


