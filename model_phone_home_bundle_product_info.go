/*
Delphix DCT API

Delphix DCT API

API version: 3.16.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PhoneHomeBundleProductInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneHomeBundleProductInfo{}

// PhoneHomeBundleProductInfo Information to identify this Data Control Tower instance.
type PhoneHomeBundleProductInfo struct {
	// A unique identifier for this Data Control Tower instance.
	SystemUuid string `json:"system_uuid"`
	// The currently running version of this instance of Data Control Tower.
	ProductVersion string `json:"product_version"`
	// The API version in use for this instance of Data Control Tower.
	ApiVersion string `json:"api_version"`
	// The upgrade history of this instance of Data Control Tower.
	ProductUpgradeHistory []PhoneHomeBundleProductHistory `json:"product_upgrade_history"`
}

type _PhoneHomeBundleProductInfo PhoneHomeBundleProductInfo

// NewPhoneHomeBundleProductInfo instantiates a new PhoneHomeBundleProductInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneHomeBundleProductInfo(systemUuid string, productVersion string, apiVersion string, productUpgradeHistory []PhoneHomeBundleProductHistory) *PhoneHomeBundleProductInfo {
	this := PhoneHomeBundleProductInfo{}
	this.SystemUuid = systemUuid
	this.ProductVersion = productVersion
	this.ApiVersion = apiVersion
	this.ProductUpgradeHistory = productUpgradeHistory
	return &this
}

// NewPhoneHomeBundleProductInfoWithDefaults instantiates a new PhoneHomeBundleProductInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneHomeBundleProductInfoWithDefaults() *PhoneHomeBundleProductInfo {
	this := PhoneHomeBundleProductInfo{}
	return &this
}

// GetSystemUuid returns the SystemUuid field value
func (o *PhoneHomeBundleProductInfo) GetSystemUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemUuid
}

// GetSystemUuidOk returns a tuple with the SystemUuid field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundleProductInfo) GetSystemUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemUuid, true
}

// SetSystemUuid sets field value
func (o *PhoneHomeBundleProductInfo) SetSystemUuid(v string) {
	o.SystemUuid = v
}

// GetProductVersion returns the ProductVersion field value
func (o *PhoneHomeBundleProductInfo) GetProductVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundleProductInfo) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVersion, true
}

// SetProductVersion sets field value
func (o *PhoneHomeBundleProductInfo) SetProductVersion(v string) {
	o.ProductVersion = v
}

// GetApiVersion returns the ApiVersion field value
func (o *PhoneHomeBundleProductInfo) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundleProductInfo) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *PhoneHomeBundleProductInfo) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetProductUpgradeHistory returns the ProductUpgradeHistory field value
func (o *PhoneHomeBundleProductInfo) GetProductUpgradeHistory() []PhoneHomeBundleProductHistory {
	if o == nil {
		var ret []PhoneHomeBundleProductHistory
		return ret
	}

	return o.ProductUpgradeHistory
}

// GetProductUpgradeHistoryOk returns a tuple with the ProductUpgradeHistory field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundleProductInfo) GetProductUpgradeHistoryOk() ([]PhoneHomeBundleProductHistory, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductUpgradeHistory, true
}

// SetProductUpgradeHistory sets field value
func (o *PhoneHomeBundleProductInfo) SetProductUpgradeHistory(v []PhoneHomeBundleProductHistory) {
	o.ProductUpgradeHistory = v
}

func (o PhoneHomeBundleProductInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneHomeBundleProductInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["system_uuid"] = o.SystemUuid
	toSerialize["product_version"] = o.ProductVersion
	toSerialize["api_version"] = o.ApiVersion
	toSerialize["product_upgrade_history"] = o.ProductUpgradeHistory
	return toSerialize, nil
}

func (o *PhoneHomeBundleProductInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"system_uuid",
		"product_version",
		"api_version",
		"product_upgrade_history",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPhoneHomeBundleProductInfo := _PhoneHomeBundleProductInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneHomeBundleProductInfo)

	if err != nil {
		return err
	}

	*o = PhoneHomeBundleProductInfo(varPhoneHomeBundleProductInfo)

	return err
}

type NullablePhoneHomeBundleProductInfo struct {
	value *PhoneHomeBundleProductInfo
	isSet bool
}

func (v NullablePhoneHomeBundleProductInfo) Get() *PhoneHomeBundleProductInfo {
	return v.value
}

func (v *NullablePhoneHomeBundleProductInfo) Set(val *PhoneHomeBundleProductInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneHomeBundleProductInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneHomeBundleProductInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneHomeBundleProductInfo(val *PhoneHomeBundleProductInfo) *NullablePhoneHomeBundleProductInfo {
	return &NullablePhoneHomeBundleProductInfo{value: val, isSet: true}
}

func (v NullablePhoneHomeBundleProductInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneHomeBundleProductInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


