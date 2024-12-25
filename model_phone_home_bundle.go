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
	"time"
	"bytes"
	"fmt"
)

// checks if the PhoneHomeBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneHomeBundle{}

// PhoneHomeBundle A bundle containing product identification, configuration information and API telemetry records.
type PhoneHomeBundle struct {
	ProductInfo PhoneHomeBundleProductInfo `json:"product_info"`
	// A list of Delphix Engines registered with this instance of Data Control Tower.
	RegisteredEngines []PhoneHomeBundleRegisteredEngine `json:"registered_engines"`
	// A list of API telemetry records.\"
	ApiTelemetry []PhoneHomeBundleApiTelemetry `json:"api_telemetry"`
	// A list of dates for which telemetry data is included in this bundle.
	Dates []string `json:"dates"`
	// The UTC time at bundle generation (ISO 8601 format).
	BundleGenerationTime time.Time `json:"bundle_generation_time"`
}

type _PhoneHomeBundle PhoneHomeBundle

// NewPhoneHomeBundle instantiates a new PhoneHomeBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneHomeBundle(productInfo PhoneHomeBundleProductInfo, registeredEngines []PhoneHomeBundleRegisteredEngine, apiTelemetry []PhoneHomeBundleApiTelemetry, dates []string, bundleGenerationTime time.Time) *PhoneHomeBundle {
	this := PhoneHomeBundle{}
	this.ProductInfo = productInfo
	this.RegisteredEngines = registeredEngines
	this.ApiTelemetry = apiTelemetry
	this.Dates = dates
	this.BundleGenerationTime = bundleGenerationTime
	return &this
}

// NewPhoneHomeBundleWithDefaults instantiates a new PhoneHomeBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneHomeBundleWithDefaults() *PhoneHomeBundle {
	this := PhoneHomeBundle{}
	return &this
}

// GetProductInfo returns the ProductInfo field value
func (o *PhoneHomeBundle) GetProductInfo() PhoneHomeBundleProductInfo {
	if o == nil {
		var ret PhoneHomeBundleProductInfo
		return ret
	}

	return o.ProductInfo
}

// GetProductInfoOk returns a tuple with the ProductInfo field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundle) GetProductInfoOk() (*PhoneHomeBundleProductInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductInfo, true
}

// SetProductInfo sets field value
func (o *PhoneHomeBundle) SetProductInfo(v PhoneHomeBundleProductInfo) {
	o.ProductInfo = v
}

// GetRegisteredEngines returns the RegisteredEngines field value
func (o *PhoneHomeBundle) GetRegisteredEngines() []PhoneHomeBundleRegisteredEngine {
	if o == nil {
		var ret []PhoneHomeBundleRegisteredEngine
		return ret
	}

	return o.RegisteredEngines
}

// GetRegisteredEnginesOk returns a tuple with the RegisteredEngines field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundle) GetRegisteredEnginesOk() ([]PhoneHomeBundleRegisteredEngine, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredEngines, true
}

// SetRegisteredEngines sets field value
func (o *PhoneHomeBundle) SetRegisteredEngines(v []PhoneHomeBundleRegisteredEngine) {
	o.RegisteredEngines = v
}

// GetApiTelemetry returns the ApiTelemetry field value
func (o *PhoneHomeBundle) GetApiTelemetry() []PhoneHomeBundleApiTelemetry {
	if o == nil {
		var ret []PhoneHomeBundleApiTelemetry
		return ret
	}

	return o.ApiTelemetry
}

// GetApiTelemetryOk returns a tuple with the ApiTelemetry field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundle) GetApiTelemetryOk() ([]PhoneHomeBundleApiTelemetry, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiTelemetry, true
}

// SetApiTelemetry sets field value
func (o *PhoneHomeBundle) SetApiTelemetry(v []PhoneHomeBundleApiTelemetry) {
	o.ApiTelemetry = v
}

// GetDates returns the Dates field value
func (o *PhoneHomeBundle) GetDates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundle) GetDatesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dates, true
}

// SetDates sets field value
func (o *PhoneHomeBundle) SetDates(v []string) {
	o.Dates = v
}

// GetBundleGenerationTime returns the BundleGenerationTime field value
func (o *PhoneHomeBundle) GetBundleGenerationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BundleGenerationTime
}

// GetBundleGenerationTimeOk returns a tuple with the BundleGenerationTime field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundle) GetBundleGenerationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleGenerationTime, true
}

// SetBundleGenerationTime sets field value
func (o *PhoneHomeBundle) SetBundleGenerationTime(v time.Time) {
	o.BundleGenerationTime = v
}

func (o PhoneHomeBundle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneHomeBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product_info"] = o.ProductInfo
	toSerialize["registered_engines"] = o.RegisteredEngines
	toSerialize["api_telemetry"] = o.ApiTelemetry
	toSerialize["dates"] = o.Dates
	toSerialize["bundle_generation_time"] = o.BundleGenerationTime
	return toSerialize, nil
}

func (o *PhoneHomeBundle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_info",
		"registered_engines",
		"api_telemetry",
		"dates",
		"bundle_generation_time",
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

	varPhoneHomeBundle := _PhoneHomeBundle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneHomeBundle)

	if err != nil {
		return err
	}

	*o = PhoneHomeBundle(varPhoneHomeBundle)

	return err
}

type NullablePhoneHomeBundle struct {
	value *PhoneHomeBundle
	isSet bool
}

func (v NullablePhoneHomeBundle) Get() *PhoneHomeBundle {
	return v.value
}

func (v *NullablePhoneHomeBundle) Set(val *PhoneHomeBundle) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneHomeBundle) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneHomeBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneHomeBundle(val *PhoneHomeBundle) *NullablePhoneHomeBundle {
	return &NullablePhoneHomeBundle{value: val, isSet: true}
}

func (v NullablePhoneHomeBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneHomeBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


