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

// checks if the HyperscaleColumnOrField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleColumnOrField{}

// HyperscaleColumnOrField struct for HyperscaleColumnOrField
type HyperscaleColumnOrField struct {
	// Name of the field.
	FieldName *string `json:"field_name,omitempty"`
	// The name of the domain assigned to this field.
	DomainName *string `json:"domain_name,omitempty"`
	// The name of the algorithm assigned to this field.
	AlgorithmName *string `json:"algorithm_name,omitempty"`
	// The format of the date assigned to this field.
	DateFormat *string `json:"date_format,omitempty"`
	// The name of the algorithm field that is associated with this algorithm.
	AlgorithmFieldName *string `json:"algorithm_field_name,omitempty"`
	// The group number of the algorithm to identify a set of columns associated with one instance of algorithm.
	AlgorithmGroupNo *int64 `json:"algorithm_group_no,omitempty"`
}

// NewHyperscaleColumnOrField instantiates a new HyperscaleColumnOrField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleColumnOrField() *HyperscaleColumnOrField {
	this := HyperscaleColumnOrField{}
	return &this
}

// NewHyperscaleColumnOrFieldWithDefaults instantiates a new HyperscaleColumnOrField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleColumnOrFieldWithDefaults() *HyperscaleColumnOrField {
	this := HyperscaleColumnOrField{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *HyperscaleColumnOrField) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleColumnOrField) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *HyperscaleColumnOrField) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *HyperscaleColumnOrField) SetFieldName(v string) {
	o.FieldName = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *HyperscaleColumnOrField) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleColumnOrField) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *HyperscaleColumnOrField) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *HyperscaleColumnOrField) SetDomainName(v string) {
	o.DomainName = &v
}

// GetAlgorithmName returns the AlgorithmName field value if set, zero value otherwise.
func (o *HyperscaleColumnOrField) GetAlgorithmName() string {
	if o == nil || IsNil(o.AlgorithmName) {
		var ret string
		return ret
	}
	return *o.AlgorithmName
}

// GetAlgorithmNameOk returns a tuple with the AlgorithmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleColumnOrField) GetAlgorithmNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlgorithmName) {
		return nil, false
	}
	return o.AlgorithmName, true
}

// HasAlgorithmName returns a boolean if a field has been set.
func (o *HyperscaleColumnOrField) HasAlgorithmName() bool {
	if o != nil && !IsNil(o.AlgorithmName) {
		return true
	}

	return false
}

// SetAlgorithmName gets a reference to the given string and assigns it to the AlgorithmName field.
func (o *HyperscaleColumnOrField) SetAlgorithmName(v string) {
	o.AlgorithmName = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *HyperscaleColumnOrField) GetDateFormat() string {
	if o == nil || IsNil(o.DateFormat) {
		var ret string
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleColumnOrField) GetDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DateFormat) {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *HyperscaleColumnOrField) HasDateFormat() bool {
	if o != nil && !IsNil(o.DateFormat) {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.
func (o *HyperscaleColumnOrField) SetDateFormat(v string) {
	o.DateFormat = &v
}

// GetAlgorithmFieldName returns the AlgorithmFieldName field value if set, zero value otherwise.
func (o *HyperscaleColumnOrField) GetAlgorithmFieldName() string {
	if o == nil || IsNil(o.AlgorithmFieldName) {
		var ret string
		return ret
	}
	return *o.AlgorithmFieldName
}

// GetAlgorithmFieldNameOk returns a tuple with the AlgorithmFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleColumnOrField) GetAlgorithmFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlgorithmFieldName) {
		return nil, false
	}
	return o.AlgorithmFieldName, true
}

// HasAlgorithmFieldName returns a boolean if a field has been set.
func (o *HyperscaleColumnOrField) HasAlgorithmFieldName() bool {
	if o != nil && !IsNil(o.AlgorithmFieldName) {
		return true
	}

	return false
}

// SetAlgorithmFieldName gets a reference to the given string and assigns it to the AlgorithmFieldName field.
func (o *HyperscaleColumnOrField) SetAlgorithmFieldName(v string) {
	o.AlgorithmFieldName = &v
}

// GetAlgorithmGroupNo returns the AlgorithmGroupNo field value if set, zero value otherwise.
func (o *HyperscaleColumnOrField) GetAlgorithmGroupNo() int64 {
	if o == nil || IsNil(o.AlgorithmGroupNo) {
		var ret int64
		return ret
	}
	return *o.AlgorithmGroupNo
}

// GetAlgorithmGroupNoOk returns a tuple with the AlgorithmGroupNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleColumnOrField) GetAlgorithmGroupNoOk() (*int64, bool) {
	if o == nil || IsNil(o.AlgorithmGroupNo) {
		return nil, false
	}
	return o.AlgorithmGroupNo, true
}

// HasAlgorithmGroupNo returns a boolean if a field has been set.
func (o *HyperscaleColumnOrField) HasAlgorithmGroupNo() bool {
	if o != nil && !IsNil(o.AlgorithmGroupNo) {
		return true
	}

	return false
}

// SetAlgorithmGroupNo gets a reference to the given int64 and assigns it to the AlgorithmGroupNo field.
func (o *HyperscaleColumnOrField) SetAlgorithmGroupNo(v int64) {
	o.AlgorithmGroupNo = &v
}

func (o HyperscaleColumnOrField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleColumnOrField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.AlgorithmName) {
		toSerialize["algorithm_name"] = o.AlgorithmName
	}
	if !IsNil(o.DateFormat) {
		toSerialize["date_format"] = o.DateFormat
	}
	if !IsNil(o.AlgorithmFieldName) {
		toSerialize["algorithm_field_name"] = o.AlgorithmFieldName
	}
	if !IsNil(o.AlgorithmGroupNo) {
		toSerialize["algorithm_group_no"] = o.AlgorithmGroupNo
	}
	return toSerialize, nil
}

type NullableHyperscaleColumnOrField struct {
	value *HyperscaleColumnOrField
	isSet bool
}

func (v NullableHyperscaleColumnOrField) Get() *HyperscaleColumnOrField {
	return v.value
}

func (v *NullableHyperscaleColumnOrField) Set(val *HyperscaleColumnOrField) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleColumnOrField) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleColumnOrField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleColumnOrField(val *HyperscaleColumnOrField) *NullableHyperscaleColumnOrField {
	return &NullableHyperscaleColumnOrField{value: val, isSet: true}
}

func (v NullableHyperscaleColumnOrField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleColumnOrField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


