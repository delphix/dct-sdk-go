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

// PermissionEnum Type of the permission on DCT object.
type PermissionEnum string

// List of PermissionEnum
const (
	PERMISSIONENUM_READ PermissionEnum = "READ"
	PERMISSIONENUM_UPDATE PermissionEnum = "UPDATE"
	PERMISSIONENUM_DELETE PermissionEnum = "DELETE"
	PERMISSIONENUM_EXECUTE PermissionEnum = "EXECUTE"
	PERMISSIONENUM_CANCEL PermissionEnum = "CANCEL"
	PERMISSIONENUM_MIGRATE PermissionEnum = "MIGRATE"
	PERMISSIONENUM_REFRESH PermissionEnum = "REFRESH"
	PERMISSIONENUM_DISABLE PermissionEnum = "DISABLE"
	PERMISSIONENUM_ENABLE PermissionEnum = "ENABLE"
	PERMISSIONENUM_ABANDON PermissionEnum = "ABANDON"
	PERMISSIONENUM_VALIDATE PermissionEnum = "VALIDATE"
	PERMISSIONENUM_START PermissionEnum = "START"
	PERMISSIONENUM_STOP PermissionEnum = "STOP"
	PERMISSIONENUM_SNAPSHOT PermissionEnum = "SNAPSHOT"
	PERMISSIONENUM_COPY PermissionEnum = "COPY"
	PERMISSIONENUM_REMOVE_JOB PermissionEnum = "REMOVE_JOB"
	PERMISSIONENUM_PASSWORD_RESET PermissionEnum = "PASSWORD_RESET"
	PERMISSIONENUM_UNDO_IMPORT PermissionEnum = "UNDO_IMPORT"
	PERMISSIONENUM_IMPORT PermissionEnum = "IMPORT"
	PERMISSIONENUM_PROVISION_FROM_BOOKMARK PermissionEnum = "PROVISION_FROM_BOOKMARK"
	PERMISSIONENUM_PROVISION PermissionEnum = "PROVISION"
	PERMISSIONENUM_REFRESH_FROM_BOOKMARK PermissionEnum = "REFRESH_FROM_BOOKMARK"
	PERMISSIONENUM_REFRESH_FROM_SNAPSHOT PermissionEnum = "REFRESH_FROM_SNAPSHOT"
	PERMISSIONENUM_REFRESH_FROM_TIMESTAMP PermissionEnum = "REFRESH_FROM_TIMESTAMP"
	PERMISSIONENUM_REFRESH_FROM_LOCATION PermissionEnum = "REFRESH_FROM_LOCATION"
	PERMISSIONENUM_CREATE_ENVIRONMENT PermissionEnum = "CREATE_ENVIRONMENT"
	PERMISSIONENUM_CREATE_BOOKMARK PermissionEnum = "CREATE_BOOKMARK"
	PERMISSIONENUM_CREATE_VDBGROUP PermissionEnum = "CREATE_VDBGROUP"
	PERMISSIONENUM_MANAGE_TAGS PermissionEnum = "MANAGE_TAGS"
	PERMISSIONENUM_LINK PermissionEnum = "LINK"
	PERMISSIONENUM_REPLICATE PermissionEnum = "REPLICATE"
	PERMISSIONENUM_REPLICATE_TO PermissionEnum = "REPLICATE_TO"
	PERMISSIONENUM_CONVERT_AND_DROP PermissionEnum = "CONVERT_AND_DROP"
	PERMISSIONENUM_IMPORT_BOOKMARKS PermissionEnum = "IMPORT_BOOKMARKS"
	PERMISSIONENUM_FAILOVER PermissionEnum = "FAILOVER"
	PERMISSIONENUM_COMMIT_FAILOVER PermissionEnum = "COMMIT_FAILOVER"
	PERMISSIONENUM_FAILBACK PermissionEnum = "FAILBACK"
	PERMISSIONENUM_DISCARD PermissionEnum = "DISCARD"
	PERMISSIONENUM_LOCK PermissionEnum = "LOCK"
	PERMISSIONENUM_UNLOCK PermissionEnum = "UNLOCK"
	PERMISSIONENUM_FORCE_UNLOCK PermissionEnum = "FORCE_UNLOCK"
	PERMISSIONENUM_LOCK_FOR_OTHER_ACCOUNT PermissionEnum = "LOCK_FOR_OTHER_ACCOUNT"
	PERMISSIONENUM_UPDATE_TIMEFLOW PermissionEnum = "UPDATE_TIMEFLOW"
	PERMISSIONENUM_SNAPSHOT_DELETE PermissionEnum = "SNAPSHOT_DELETE"
	PERMISSIONENUM_SWITCH_TIMEFLOW PermissionEnum = "SWITCH_TIMEFLOW"
	PERMISSIONENUM_DELETE_TIMEFLOW PermissionEnum = "DELETE_TIMEFLOW"
	PERMISSIONENUM_SNAPSHOT_UPDATE PermissionEnum = "SNAPSHOT_UPDATE"
	PERMISSIONENUM_IMPORT_ACCOUNTS PermissionEnum = "IMPORT_ACCOUNTS"
	PERMISSIONENUM_DETACH_SOURCE PermissionEnum = "DETACH_SOURCE"
	PERMISSIONENUM_ATTACH_SOURCE PermissionEnum = "ATTACH_SOURCE"
	PERMISSIONENUM_RESOLVE PermissionEnum = "RESOLVE"
	PERMISSIONENUM_RESOLVE_ALL PermissionEnum = "RESOLVE_ALL"
	PERMISSIONENUM_RESOLVE_OR_IGNORE PermissionEnum = "RESOLVE_OR_IGNORE"
	PERMISSIONENUM_API_KEY_RESET PermissionEnum = "API_KEY_RESET"
	PERMISSIONENUM_API_KEY_DELETE PermissionEnum = "API_KEY_DELETE"
	PERMISSIONENUM_DELETE_BOOKMARK PermissionEnum = "DELETE_BOOKMARK"
	PERMISSIONENUM_UPDATE_BOOKMARK PermissionEnum = "UPDATE_BOOKMARK"
	PERMISSIONENUM_READ_BOOKMARK PermissionEnum = "READ_BOOKMARK"
	PERMISSIONENUM_UPLOAD_MASKING_FILE PermissionEnum = "UPLOAD_MASKING_FILE"
	PERMISSIONENUM_GLOBAL_SYNC_ENGINES PermissionEnum = "GLOBAL_SYNC_ENGINES"
	PERMISSIONENUM_ADD_ENGINE_TO_HYPERSCALE PermissionEnum = "ADD_ENGINE_TO_HYPERSCALE"
	PERMISSIONENUM_CONFIGURE_CUSTOM_AUTO_TAGGING PermissionEnum = "CONFIGURE_CUSTOM_AUTO_TAGGING"
	PERMISSIONENUM_CONFIGURE_PREDEFINED_AUTO_TAGGING PermissionEnum = "CONFIGURE_PREDEFINED_AUTO_TAGGING"
	PERMISSIONENUM_APPLY PermissionEnum = "APPLY"
	PERMISSIONENUM_UNAPPLY PermissionEnum = "UNAPPLY"
	PERMISSIONENUM_UNDO_REFRESH PermissionEnum = "UNDO_REFRESH"
	PERMISSIONENUM_CONVERT PermissionEnum = "CONVERT"
)

// All allowed values of PermissionEnum enum
var AllowedPermissionEnumEnumValues = []PermissionEnum{
	"READ",
	"UPDATE",
	"DELETE",
	"EXECUTE",
	"CANCEL",
	"MIGRATE",
	"REFRESH",
	"DISABLE",
	"ENABLE",
	"ABANDON",
	"VALIDATE",
	"START",
	"STOP",
	"SNAPSHOT",
	"COPY",
	"REMOVE_JOB",
	"PASSWORD_RESET",
	"UNDO_IMPORT",
	"IMPORT",
	"PROVISION_FROM_BOOKMARK",
	"PROVISION",
	"REFRESH_FROM_BOOKMARK",
	"REFRESH_FROM_SNAPSHOT",
	"REFRESH_FROM_TIMESTAMP",
	"REFRESH_FROM_LOCATION",
	"CREATE_ENVIRONMENT",
	"CREATE_BOOKMARK",
	"CREATE_VDBGROUP",
	"MANAGE_TAGS",
	"LINK",
	"REPLICATE",
	"REPLICATE_TO",
	"CONVERT_AND_DROP",
	"IMPORT_BOOKMARKS",
	"FAILOVER",
	"COMMIT_FAILOVER",
	"FAILBACK",
	"DISCARD",
	"LOCK",
	"UNLOCK",
	"FORCE_UNLOCK",
	"LOCK_FOR_OTHER_ACCOUNT",
	"UPDATE_TIMEFLOW",
	"SNAPSHOT_DELETE",
	"SWITCH_TIMEFLOW",
	"DELETE_TIMEFLOW",
	"SNAPSHOT_UPDATE",
	"IMPORT_ACCOUNTS",
	"DETACH_SOURCE",
	"ATTACH_SOURCE",
	"RESOLVE",
	"RESOLVE_ALL",
	"RESOLVE_OR_IGNORE",
	"API_KEY_RESET",
	"API_KEY_DELETE",
	"DELETE_BOOKMARK",
	"UPDATE_BOOKMARK",
	"READ_BOOKMARK",
	"UPLOAD_MASKING_FILE",
	"GLOBAL_SYNC_ENGINES",
	"ADD_ENGINE_TO_HYPERSCALE",
	"CONFIGURE_CUSTOM_AUTO_TAGGING",
	"CONFIGURE_PREDEFINED_AUTO_TAGGING",
	"APPLY",
	"UNAPPLY",
	"UNDO_REFRESH",
	"CONVERT",
}

func (v *PermissionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PermissionEnum(value)
	for _, existing := range AllowedPermissionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PermissionEnum", value)
}

// NewPermissionEnumFromValue returns a pointer to a valid PermissionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPermissionEnumFromValue(v string) (*PermissionEnum, error) {
	ev := PermissionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PermissionEnum: valid values are %v", v, AllowedPermissionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PermissionEnum) IsValid() bool {
	for _, existing := range AllowedPermissionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PermissionEnum value
func (v PermissionEnum) Ptr() *PermissionEnum {
	return &v
}

type NullablePermissionEnum struct {
	value *PermissionEnum
	isSet bool
}

func (v NullablePermissionEnum) Get() *PermissionEnum {
	return v.value
}

func (v *NullablePermissionEnum) Set(val *PermissionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionEnum(val *PermissionEnum) *NullablePermissionEnum {
	return &NullablePermissionEnum{value: val, isSet: true}
}

func (v NullablePermissionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

