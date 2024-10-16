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
	"fmt"
)

// OracleConfigTypeEnum the model 'OracleConfigTypeEnum'
type OracleConfigTypeEnum string

// List of OracleConfigTypeEnum
const (
	ORACLECONFIGTYPEENUM_ORACLE_RAC_CONFIG OracleConfigTypeEnum = "OracleRACConfig"
	ORACLECONFIGTYPEENUM_ORACLE_SI_CONFIG OracleConfigTypeEnum = "OracleSIConfig"
	ORACLECONFIGTYPEENUM_ORACLE_PDB_CONFIG OracleConfigTypeEnum = "OraclePDBConfig"
)

// All allowed values of OracleConfigTypeEnum enum
var AllowedOracleConfigTypeEnumEnumValues = []OracleConfigTypeEnum{
	"OracleRACConfig",
	"OracleSIConfig",
	"OraclePDBConfig",
}

func (v *OracleConfigTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OracleConfigTypeEnum(value)
	for _, existing := range AllowedOracleConfigTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OracleConfigTypeEnum", value)
}

// NewOracleConfigTypeEnumFromValue returns a pointer to a valid OracleConfigTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOracleConfigTypeEnumFromValue(v string) (*OracleConfigTypeEnum, error) {
	ev := OracleConfigTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OracleConfigTypeEnum: valid values are %v", v, AllowedOracleConfigTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OracleConfigTypeEnum) IsValid() bool {
	for _, existing := range AllowedOracleConfigTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OracleConfigTypeEnum value
func (v OracleConfigTypeEnum) Ptr() *OracleConfigTypeEnum {
	return &v
}

type NullableOracleConfigTypeEnum struct {
	value *OracleConfigTypeEnum
	isSet bool
}

func (v NullableOracleConfigTypeEnum) Get() *OracleConfigTypeEnum {
	return v.value
}

func (v *NullableOracleConfigTypeEnum) Set(val *OracleConfigTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleConfigTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleConfigTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleConfigTypeEnum(val *OracleConfigTypeEnum) *NullableOracleConfigTypeEnum {
	return &NullableOracleConfigTypeEnum{value: val, isSet: true}
}

func (v NullableOracleConfigTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleConfigTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

