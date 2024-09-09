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
	"fmt"
)

// HyperscaleDataTypeEnum The data type designation for the hyperscale deployment.
type HyperscaleDataTypeEnum string

// List of HyperscaleDataTypeEnum
const (
	HYPERSCALEDATATYPEENUM_ORACLE HyperscaleDataTypeEnum = "ORACLE"
	HYPERSCALEDATATYPEENUM_MSSQL HyperscaleDataTypeEnum = "MSSQL"
	HYPERSCALEDATATYPEENUM_DELIMITED_FILES HyperscaleDataTypeEnum = "DELIMITED_FILES"
	HYPERSCALEDATATYPEENUM_MONGO_DB HyperscaleDataTypeEnum = "MONGO_DB"
)

// All allowed values of HyperscaleDataTypeEnum enum
var AllowedHyperscaleDataTypeEnumEnumValues = []HyperscaleDataTypeEnum{
	"ORACLE",
	"MSSQL",
	"DELIMITED_FILES",
	"MONGO_DB",
}

func (v *HyperscaleDataTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HyperscaleDataTypeEnum(value)
	for _, existing := range AllowedHyperscaleDataTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HyperscaleDataTypeEnum", value)
}

// NewHyperscaleDataTypeEnumFromValue returns a pointer to a valid HyperscaleDataTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHyperscaleDataTypeEnumFromValue(v string) (*HyperscaleDataTypeEnum, error) {
	ev := HyperscaleDataTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HyperscaleDataTypeEnum: valid values are %v", v, AllowedHyperscaleDataTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HyperscaleDataTypeEnum) IsValid() bool {
	for _, existing := range AllowedHyperscaleDataTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HyperscaleDataTypeEnum value
func (v HyperscaleDataTypeEnum) Ptr() *HyperscaleDataTypeEnum {
	return &v
}

type NullableHyperscaleDataTypeEnum struct {
	value *HyperscaleDataTypeEnum
	isSet bool
}

func (v NullableHyperscaleDataTypeEnum) Get() *HyperscaleDataTypeEnum {
	return v.value
}

func (v *NullableHyperscaleDataTypeEnum) Set(val *HyperscaleDataTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleDataTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleDataTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleDataTypeEnum(val *HyperscaleDataTypeEnum) *NullableHyperscaleDataTypeEnum {
	return &NullableHyperscaleDataTypeEnum{value: val, isSet: true}
}

func (v NullableHyperscaleDataTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleDataTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

