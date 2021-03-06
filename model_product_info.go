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

// ProductInfo Product Information Response
type ProductInfo struct {
	// Current API version.
	ApiVersion *string `json:"api_version,omitempty"`
	// Current installed product version.
	ProductVersion *string `json:"product_version,omitempty"`
	// Product upgrade history.
	ProductUpgradeHistory []ProductHistory `json:"product_upgrade_history,omitempty"`
	// All the supported API versions.
	SupportedApiVersions []string `json:"supported_api_versions,omitempty"`
}

// NewProductInfo instantiates a new ProductInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductInfo() *ProductInfo {
	this := ProductInfo{}
	return &this
}

// NewProductInfoWithDefaults instantiates a new ProductInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductInfoWithDefaults() *ProductInfo {
	this := ProductInfo{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ProductInfo) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfo) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ProductInfo) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ProductInfo) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ProductInfo) GetProductVersion() string {
	if o == nil || o.ProductVersion == nil {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfo) GetProductVersionOk() (*string, bool) {
	if o == nil || o.ProductVersion == nil {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ProductInfo) HasProductVersion() bool {
	if o != nil && o.ProductVersion != nil {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *ProductInfo) SetProductVersion(v string) {
	o.ProductVersion = &v
}

// GetProductUpgradeHistory returns the ProductUpgradeHistory field value if set, zero value otherwise.
func (o *ProductInfo) GetProductUpgradeHistory() []ProductHistory {
	if o == nil || o.ProductUpgradeHistory == nil {
		var ret []ProductHistory
		return ret
	}
	return o.ProductUpgradeHistory
}

// GetProductUpgradeHistoryOk returns a tuple with the ProductUpgradeHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfo) GetProductUpgradeHistoryOk() ([]ProductHistory, bool) {
	if o == nil || o.ProductUpgradeHistory == nil {
		return nil, false
	}
	return o.ProductUpgradeHistory, true
}

// HasProductUpgradeHistory returns a boolean if a field has been set.
func (o *ProductInfo) HasProductUpgradeHistory() bool {
	if o != nil && o.ProductUpgradeHistory != nil {
		return true
	}

	return false
}

// SetProductUpgradeHistory gets a reference to the given []ProductHistory and assigns it to the ProductUpgradeHistory field.
func (o *ProductInfo) SetProductUpgradeHistory(v []ProductHistory) {
	o.ProductUpgradeHistory = v
}

// GetSupportedApiVersions returns the SupportedApiVersions field value if set, zero value otherwise.
func (o *ProductInfo) GetSupportedApiVersions() []string {
	if o == nil || o.SupportedApiVersions == nil {
		var ret []string
		return ret
	}
	return o.SupportedApiVersions
}

// GetSupportedApiVersionsOk returns a tuple with the SupportedApiVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfo) GetSupportedApiVersionsOk() ([]string, bool) {
	if o == nil || o.SupportedApiVersions == nil {
		return nil, false
	}
	return o.SupportedApiVersions, true
}

// HasSupportedApiVersions returns a boolean if a field has been set.
func (o *ProductInfo) HasSupportedApiVersions() bool {
	if o != nil && o.SupportedApiVersions != nil {
		return true
	}

	return false
}

// SetSupportedApiVersions gets a reference to the given []string and assigns it to the SupportedApiVersions field.
func (o *ProductInfo) SetSupportedApiVersions(v []string) {
	o.SupportedApiVersions = v
}

func (o ProductInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.ProductVersion != nil {
		toSerialize["product_version"] = o.ProductVersion
	}
	if o.ProductUpgradeHistory != nil {
		toSerialize["product_upgrade_history"] = o.ProductUpgradeHistory
	}
	if o.SupportedApiVersions != nil {
		toSerialize["supported_api_versions"] = o.SupportedApiVersions
	}
	return json.Marshal(toSerialize)
}

type NullableProductInfo struct {
	value *ProductInfo
	isSet bool
}

func (v NullableProductInfo) Get() *ProductInfo {
	return v.value
}

func (v *NullableProductInfo) Set(val *ProductInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProductInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProductInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductInfo(val *ProductInfo) *NullableProductInfo {
	return &NullableProductInfo{value: val, isSet: true}
}

func (v NullableProductInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


