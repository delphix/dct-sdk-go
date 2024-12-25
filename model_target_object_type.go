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
	"fmt"
)

// TargetObjectType A DCT specific object type.
type TargetObjectType string

// List of TargetObjectType
const (
	TARGETOBJECTTYPE_REPLICATION_PROFILE TargetObjectType = "REPLICATION_PROFILE"
)

// All allowed values of TargetObjectType enum
var AllowedTargetObjectTypeEnumValues = []TargetObjectType{
	"REPLICATION_PROFILE",
}

func (v *TargetObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TargetObjectType(value)
	for _, existing := range AllowedTargetObjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TargetObjectType", value)
}

// NewTargetObjectTypeFromValue returns a pointer to a valid TargetObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTargetObjectTypeFromValue(v string) (*TargetObjectType, error) {
	ev := TargetObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TargetObjectType: valid values are %v", v, AllowedTargetObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TargetObjectType) IsValid() bool {
	for _, existing := range AllowedTargetObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TargetObjectType value
func (v TargetObjectType) Ptr() *TargetObjectType {
	return &v
}

type NullableTargetObjectType struct {
	value *TargetObjectType
	isSet bool
}

func (v NullableTargetObjectType) Get() *TargetObjectType {
	return v.value
}

func (v *NullableTargetObjectType) Set(val *TargetObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetObjectType(val *TargetObjectType) *NullableTargetObjectType {
	return &NullableTargetObjectType{value: val, isSet: true}
}

func (v NullableTargetObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

