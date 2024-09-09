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

// checks if the UpdateRepositoryParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRepositoryParameters{}

// UpdateRepositoryParameters struct for UpdateRepositoryParameters
type UpdateRepositoryParameters struct {
	// The database type of this repository.
	DatabaseType string `json:"database_type"`
	// Flag indicating whether the repository should be used for provisioning.
	AllowProvisioning *bool `json:"allow_provisioning,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix Engine for internal processing.
	IsStaging *bool `json:"is_staging,omitempty"`
	// Version of the repository.
	Version *string `json:"version,omitempty"`
	// The Oracle base where database binaries are located.
	OracleBase *string `json:"oracle_base,omitempty"`
	// 32 or 64 bits.
	Bits *int32 `json:"bits,omitempty"`
}

type _UpdateRepositoryParameters UpdateRepositoryParameters

// NewUpdateRepositoryParameters instantiates a new UpdateRepositoryParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRepositoryParameters(databaseType string) *UpdateRepositoryParameters {
	this := UpdateRepositoryParameters{}
	this.DatabaseType = databaseType
	return &this
}

// NewUpdateRepositoryParametersWithDefaults instantiates a new UpdateRepositoryParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRepositoryParametersWithDefaults() *UpdateRepositoryParameters {
	this := UpdateRepositoryParameters{}
	return &this
}

// GetDatabaseType returns the DatabaseType field value
func (o *UpdateRepositoryParameters) GetDatabaseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value
// and a boolean to check if the value has been set.
func (o *UpdateRepositoryParameters) GetDatabaseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseType, true
}

// SetDatabaseType sets field value
func (o *UpdateRepositoryParameters) SetDatabaseType(v string) {
	o.DatabaseType = v
}

// GetAllowProvisioning returns the AllowProvisioning field value if set, zero value otherwise.
func (o *UpdateRepositoryParameters) GetAllowProvisioning() bool {
	if o == nil || IsNil(o.AllowProvisioning) {
		var ret bool
		return ret
	}
	return *o.AllowProvisioning
}

// GetAllowProvisioningOk returns a tuple with the AllowProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRepositoryParameters) GetAllowProvisioningOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowProvisioning) {
		return nil, false
	}
	return o.AllowProvisioning, true
}

// HasAllowProvisioning returns a boolean if a field has been set.
func (o *UpdateRepositoryParameters) HasAllowProvisioning() bool {
	if o != nil && !IsNil(o.AllowProvisioning) {
		return true
	}

	return false
}

// SetAllowProvisioning gets a reference to the given bool and assigns it to the AllowProvisioning field.
func (o *UpdateRepositoryParameters) SetAllowProvisioning(v bool) {
	o.AllowProvisioning = &v
}

// GetIsStaging returns the IsStaging field value if set, zero value otherwise.
func (o *UpdateRepositoryParameters) GetIsStaging() bool {
	if o == nil || IsNil(o.IsStaging) {
		var ret bool
		return ret
	}
	return *o.IsStaging
}

// GetIsStagingOk returns a tuple with the IsStaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRepositoryParameters) GetIsStagingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStaging) {
		return nil, false
	}
	return o.IsStaging, true
}

// HasIsStaging returns a boolean if a field has been set.
func (o *UpdateRepositoryParameters) HasIsStaging() bool {
	if o != nil && !IsNil(o.IsStaging) {
		return true
	}

	return false
}

// SetIsStaging gets a reference to the given bool and assigns it to the IsStaging field.
func (o *UpdateRepositoryParameters) SetIsStaging(v bool) {
	o.IsStaging = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpdateRepositoryParameters) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRepositoryParameters) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpdateRepositoryParameters) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UpdateRepositoryParameters) SetVersion(v string) {
	o.Version = &v
}

// GetOracleBase returns the OracleBase field value if set, zero value otherwise.
func (o *UpdateRepositoryParameters) GetOracleBase() string {
	if o == nil || IsNil(o.OracleBase) {
		var ret string
		return ret
	}
	return *o.OracleBase
}

// GetOracleBaseOk returns a tuple with the OracleBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRepositoryParameters) GetOracleBaseOk() (*string, bool) {
	if o == nil || IsNil(o.OracleBase) {
		return nil, false
	}
	return o.OracleBase, true
}

// HasOracleBase returns a boolean if a field has been set.
func (o *UpdateRepositoryParameters) HasOracleBase() bool {
	if o != nil && !IsNil(o.OracleBase) {
		return true
	}

	return false
}

// SetOracleBase gets a reference to the given string and assigns it to the OracleBase field.
func (o *UpdateRepositoryParameters) SetOracleBase(v string) {
	o.OracleBase = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *UpdateRepositoryParameters) GetBits() int32 {
	if o == nil || IsNil(o.Bits) {
		var ret int32
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRepositoryParameters) GetBitsOk() (*int32, bool) {
	if o == nil || IsNil(o.Bits) {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *UpdateRepositoryParameters) HasBits() bool {
	if o != nil && !IsNil(o.Bits) {
		return true
	}

	return false
}

// SetBits gets a reference to the given int32 and assigns it to the Bits field.
func (o *UpdateRepositoryParameters) SetBits(v int32) {
	o.Bits = &v
}

func (o UpdateRepositoryParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRepositoryParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["database_type"] = o.DatabaseType
	if !IsNil(o.AllowProvisioning) {
		toSerialize["allow_provisioning"] = o.AllowProvisioning
	}
	if !IsNil(o.IsStaging) {
		toSerialize["is_staging"] = o.IsStaging
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.OracleBase) {
		toSerialize["oracle_base"] = o.OracleBase
	}
	if !IsNil(o.Bits) {
		toSerialize["bits"] = o.Bits
	}
	return toSerialize, nil
}

func (o *UpdateRepositoryParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"database_type",
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

	varUpdateRepositoryParameters := _UpdateRepositoryParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateRepositoryParameters)

	if err != nil {
		return err
	}

	*o = UpdateRepositoryParameters(varUpdateRepositoryParameters)

	return err
}

type NullableUpdateRepositoryParameters struct {
	value *UpdateRepositoryParameters
	isSet bool
}

func (v NullableUpdateRepositoryParameters) Get() *UpdateRepositoryParameters {
	return v.value
}

func (v *NullableUpdateRepositoryParameters) Set(val *UpdateRepositoryParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRepositoryParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRepositoryParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRepositoryParameters(val *UpdateRepositoryParameters) *NullableUpdateRepositoryParameters {
	return &NullableUpdateRepositoryParameters{value: val, isSet: true}
}

func (v NullableUpdateRepositoryParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRepositoryParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


