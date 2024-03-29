/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the HyperscaleMountPointUpdateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleMountPointUpdateParameters{}

// HyperscaleMountPointUpdateParameters The updatable properties of a Hyperscale Mount Point.
type HyperscaleMountPointUpdateParameters struct {
	// Name of the mount, unique for a hyperscale instance.
	Name *string `json:"name,omitempty"`
	// The host name of the server.
	Hostname *string `json:"hostname,omitempty"`
	// The path to the directory on the filesystem to mount.
	MountPath *string `json:"mount_path,omitempty"`
	// The type of filesystem.
	MountType *string `json:"mount_type,omitempty"`
	// The mount options.
	Options *string `json:"options,omitempty"`
}

// NewHyperscaleMountPointUpdateParameters instantiates a new HyperscaleMountPointUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleMountPointUpdateParameters() *HyperscaleMountPointUpdateParameters {
	this := HyperscaleMountPointUpdateParameters{}
	return &this
}

// NewHyperscaleMountPointUpdateParametersWithDefaults instantiates a new HyperscaleMountPointUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleMountPointUpdateParametersWithDefaults() *HyperscaleMountPointUpdateParameters {
	this := HyperscaleMountPointUpdateParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperscaleMountPointUpdateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPointUpdateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperscaleMountPointUpdateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperscaleMountPointUpdateParameters) SetName(v string) {
	o.Name = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HyperscaleMountPointUpdateParameters) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPointUpdateParameters) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HyperscaleMountPointUpdateParameters) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HyperscaleMountPointUpdateParameters) SetHostname(v string) {
	o.Hostname = &v
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *HyperscaleMountPointUpdateParameters) GetMountPath() string {
	if o == nil || IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPointUpdateParameters) GetMountPathOk() (*string, bool) {
	if o == nil || IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *HyperscaleMountPointUpdateParameters) HasMountPath() bool {
	if o != nil && !IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the MountPath field.
func (o *HyperscaleMountPointUpdateParameters) SetMountPath(v string) {
	o.MountPath = &v
}

// GetMountType returns the MountType field value if set, zero value otherwise.
func (o *HyperscaleMountPointUpdateParameters) GetMountType() string {
	if o == nil || IsNil(o.MountType) {
		var ret string
		return ret
	}
	return *o.MountType
}

// GetMountTypeOk returns a tuple with the MountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPointUpdateParameters) GetMountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MountType) {
		return nil, false
	}
	return o.MountType, true
}

// HasMountType returns a boolean if a field has been set.
func (o *HyperscaleMountPointUpdateParameters) HasMountType() bool {
	if o != nil && !IsNil(o.MountType) {
		return true
	}

	return false
}

// SetMountType gets a reference to the given string and assigns it to the MountType field.
func (o *HyperscaleMountPointUpdateParameters) SetMountType(v string) {
	o.MountType = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *HyperscaleMountPointUpdateParameters) GetOptions() string {
	if o == nil || IsNil(o.Options) {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPointUpdateParameters) GetOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *HyperscaleMountPointUpdateParameters) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *HyperscaleMountPointUpdateParameters) SetOptions(v string) {
	o.Options = &v
}

func (o HyperscaleMountPointUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleMountPointUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.MountPath) {
		toSerialize["mount_path"] = o.MountPath
	}
	if !IsNil(o.MountType) {
		toSerialize["mount_type"] = o.MountType
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableHyperscaleMountPointUpdateParameters struct {
	value *HyperscaleMountPointUpdateParameters
	isSet bool
}

func (v NullableHyperscaleMountPointUpdateParameters) Get() *HyperscaleMountPointUpdateParameters {
	return v.value
}

func (v *NullableHyperscaleMountPointUpdateParameters) Set(val *HyperscaleMountPointUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleMountPointUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleMountPointUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleMountPointUpdateParameters(val *HyperscaleMountPointUpdateParameters) *NullableHyperscaleMountPointUpdateParameters {
	return &NullableHyperscaleMountPointUpdateParameters{value: val, isSet: true}
}

func (v NullableHyperscaleMountPointUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleMountPointUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


