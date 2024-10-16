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
)

// checks if the MetadataDbInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataDbInfo{}

// MetadataDbInfo Information about the product's metadata database.
type MetadataDbInfo struct {
	// True if an external database, i.e a database running outside of the application cluster, is in use.
	External *bool `json:"external,omitempty"`
	// The full database version in String format
	Version *string `json:"version,omitempty"`
	// The database product name as reported by the database itself.
	DatabaseProductName *string `json:"database_product_name,omitempty"`
	// The database major version.
	MajorVersion *int32 `json:"major_version,omitempty"`
	// The database minor version
	MinorVersion *int32 `json:"minor_version,omitempty"`
	// The minimum supported major version of PostgreSQL.
	MinSupportedMajorVersion *int32 `json:"min_supported_major_version,omitempty"`
	// The minimum supported minor version of PostgreSQL.
	MinSupportedMinorVersion *int32 `json:"min_supported_minor_version,omitempty"`
	// The maximum supported major version of PostgreSQL.
	MaxSupportedMajorVersion *int32 `json:"max_supported_major_version,omitempty"`
	// The maximum supported minor version of PostgreSQL.
	MaxSupportedMinorVersion *int32 `json:"max_supported_minor_version,omitempty"`
	// Whether the database is recognized as valid for this product. In order to be compatible, the database product name must be a recognized PostgreSQL and the database version must be greater than or equal to the minimum minor version and smaller than or equal to the maximum support version.
	Compatible *bool `json:"compatible,omitempty"`
}

// NewMetadataDbInfo instantiates a new MetadataDbInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataDbInfo() *MetadataDbInfo {
	this := MetadataDbInfo{}
	return &this
}

// NewMetadataDbInfoWithDefaults instantiates a new MetadataDbInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataDbInfoWithDefaults() *MetadataDbInfo {
	this := MetadataDbInfo{}
	return &this
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *MetadataDbInfo) SetExternal(v bool) {
	o.External = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MetadataDbInfo) SetVersion(v string) {
	o.Version = &v
}

// GetDatabaseProductName returns the DatabaseProductName field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetDatabaseProductName() string {
	if o == nil || IsNil(o.DatabaseProductName) {
		var ret string
		return ret
	}
	return *o.DatabaseProductName
}

// GetDatabaseProductNameOk returns a tuple with the DatabaseProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetDatabaseProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseProductName) {
		return nil, false
	}
	return o.DatabaseProductName, true
}

// HasDatabaseProductName returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasDatabaseProductName() bool {
	if o != nil && !IsNil(o.DatabaseProductName) {
		return true
	}

	return false
}

// SetDatabaseProductName gets a reference to the given string and assigns it to the DatabaseProductName field.
func (o *MetadataDbInfo) SetDatabaseProductName(v string) {
	o.DatabaseProductName = &v
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetMajorVersion() int32 {
	if o == nil || IsNil(o.MajorVersion) {
		var ret int32
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetMajorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MajorVersion) {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasMajorVersion() bool {
	if o != nil && !IsNil(o.MajorVersion) {
		return true
	}

	return false
}

// SetMajorVersion gets a reference to the given int32 and assigns it to the MajorVersion field.
func (o *MetadataDbInfo) SetMajorVersion(v int32) {
	o.MajorVersion = &v
}

// GetMinorVersion returns the MinorVersion field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetMinorVersion() int32 {
	if o == nil || IsNil(o.MinorVersion) {
		var ret int32
		return ret
	}
	return *o.MinorVersion
}

// GetMinorVersionOk returns a tuple with the MinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetMinorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MinorVersion) {
		return nil, false
	}
	return o.MinorVersion, true
}

// HasMinorVersion returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasMinorVersion() bool {
	if o != nil && !IsNil(o.MinorVersion) {
		return true
	}

	return false
}

// SetMinorVersion gets a reference to the given int32 and assigns it to the MinorVersion field.
func (o *MetadataDbInfo) SetMinorVersion(v int32) {
	o.MinorVersion = &v
}

// GetMinSupportedMajorVersion returns the MinSupportedMajorVersion field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetMinSupportedMajorVersion() int32 {
	if o == nil || IsNil(o.MinSupportedMajorVersion) {
		var ret int32
		return ret
	}
	return *o.MinSupportedMajorVersion
}

// GetMinSupportedMajorVersionOk returns a tuple with the MinSupportedMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetMinSupportedMajorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSupportedMajorVersion) {
		return nil, false
	}
	return o.MinSupportedMajorVersion, true
}

// HasMinSupportedMajorVersion returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasMinSupportedMajorVersion() bool {
	if o != nil && !IsNil(o.MinSupportedMajorVersion) {
		return true
	}

	return false
}

// SetMinSupportedMajorVersion gets a reference to the given int32 and assigns it to the MinSupportedMajorVersion field.
func (o *MetadataDbInfo) SetMinSupportedMajorVersion(v int32) {
	o.MinSupportedMajorVersion = &v
}

// GetMinSupportedMinorVersion returns the MinSupportedMinorVersion field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetMinSupportedMinorVersion() int32 {
	if o == nil || IsNil(o.MinSupportedMinorVersion) {
		var ret int32
		return ret
	}
	return *o.MinSupportedMinorVersion
}

// GetMinSupportedMinorVersionOk returns a tuple with the MinSupportedMinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetMinSupportedMinorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSupportedMinorVersion) {
		return nil, false
	}
	return o.MinSupportedMinorVersion, true
}

// HasMinSupportedMinorVersion returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasMinSupportedMinorVersion() bool {
	if o != nil && !IsNil(o.MinSupportedMinorVersion) {
		return true
	}

	return false
}

// SetMinSupportedMinorVersion gets a reference to the given int32 and assigns it to the MinSupportedMinorVersion field.
func (o *MetadataDbInfo) SetMinSupportedMinorVersion(v int32) {
	o.MinSupportedMinorVersion = &v
}

// GetMaxSupportedMajorVersion returns the MaxSupportedMajorVersion field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetMaxSupportedMajorVersion() int32 {
	if o == nil || IsNil(o.MaxSupportedMajorVersion) {
		var ret int32
		return ret
	}
	return *o.MaxSupportedMajorVersion
}

// GetMaxSupportedMajorVersionOk returns a tuple with the MaxSupportedMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetMaxSupportedMajorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSupportedMajorVersion) {
		return nil, false
	}
	return o.MaxSupportedMajorVersion, true
}

// HasMaxSupportedMajorVersion returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasMaxSupportedMajorVersion() bool {
	if o != nil && !IsNil(o.MaxSupportedMajorVersion) {
		return true
	}

	return false
}

// SetMaxSupportedMajorVersion gets a reference to the given int32 and assigns it to the MaxSupportedMajorVersion field.
func (o *MetadataDbInfo) SetMaxSupportedMajorVersion(v int32) {
	o.MaxSupportedMajorVersion = &v
}

// GetMaxSupportedMinorVersion returns the MaxSupportedMinorVersion field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetMaxSupportedMinorVersion() int32 {
	if o == nil || IsNil(o.MaxSupportedMinorVersion) {
		var ret int32
		return ret
	}
	return *o.MaxSupportedMinorVersion
}

// GetMaxSupportedMinorVersionOk returns a tuple with the MaxSupportedMinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetMaxSupportedMinorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSupportedMinorVersion) {
		return nil, false
	}
	return o.MaxSupportedMinorVersion, true
}

// HasMaxSupportedMinorVersion returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasMaxSupportedMinorVersion() bool {
	if o != nil && !IsNil(o.MaxSupportedMinorVersion) {
		return true
	}

	return false
}

// SetMaxSupportedMinorVersion gets a reference to the given int32 and assigns it to the MaxSupportedMinorVersion field.
func (o *MetadataDbInfo) SetMaxSupportedMinorVersion(v int32) {
	o.MaxSupportedMinorVersion = &v
}

// GetCompatible returns the Compatible field value if set, zero value otherwise.
func (o *MetadataDbInfo) GetCompatible() bool {
	if o == nil || IsNil(o.Compatible) {
		var ret bool
		return ret
	}
	return *o.Compatible
}

// GetCompatibleOk returns a tuple with the Compatible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDbInfo) GetCompatibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Compatible) {
		return nil, false
	}
	return o.Compatible, true
}

// HasCompatible returns a boolean if a field has been set.
func (o *MetadataDbInfo) HasCompatible() bool {
	if o != nil && !IsNil(o.Compatible) {
		return true
	}

	return false
}

// SetCompatible gets a reference to the given bool and assigns it to the Compatible field.
func (o *MetadataDbInfo) SetCompatible(v bool) {
	o.Compatible = &v
}

func (o MetadataDbInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataDbInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.DatabaseProductName) {
		toSerialize["database_product_name"] = o.DatabaseProductName
	}
	if !IsNil(o.MajorVersion) {
		toSerialize["major_version"] = o.MajorVersion
	}
	if !IsNil(o.MinorVersion) {
		toSerialize["minor_version"] = o.MinorVersion
	}
	if !IsNil(o.MinSupportedMajorVersion) {
		toSerialize["min_supported_major_version"] = o.MinSupportedMajorVersion
	}
	if !IsNil(o.MinSupportedMinorVersion) {
		toSerialize["min_supported_minor_version"] = o.MinSupportedMinorVersion
	}
	if !IsNil(o.MaxSupportedMajorVersion) {
		toSerialize["max_supported_major_version"] = o.MaxSupportedMajorVersion
	}
	if !IsNil(o.MaxSupportedMinorVersion) {
		toSerialize["max_supported_minor_version"] = o.MaxSupportedMinorVersion
	}
	if !IsNil(o.Compatible) {
		toSerialize["compatible"] = o.Compatible
	}
	return toSerialize, nil
}

type NullableMetadataDbInfo struct {
	value *MetadataDbInfo
	isSet bool
}

func (v NullableMetadataDbInfo) Get() *MetadataDbInfo {
	return v.value
}

func (v *NullableMetadataDbInfo) Set(val *MetadataDbInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataDbInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataDbInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataDbInfo(val *MetadataDbInfo) *NullableMetadataDbInfo {
	return &NullableMetadataDbInfo{value: val, isSet: true}
}

func (v NullableMetadataDbInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataDbInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


