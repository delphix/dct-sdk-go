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

// checks if the HyperscaleMountPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleMountPoint{}

// HyperscaleMountPoint A Mount Point for the Hyperscale instance to write to a staging area and engines to read from.
type HyperscaleMountPoint struct {
	// The ID of the Hyperscale Mount Point.
	Id *string `json:"id,omitempty"`
	// The ID of the Hyperscale instance of this Mount Point.
	HyperscaleInstanceId string `json:"hyperscale_instance_id"`
	// Name of the mount, unique for a hyperscale instance. This name will be used as a directory name by the Hyperscale instance.
	Name string `json:"name"`
	// The host name of the server.
	Hostname string `json:"hostname"`
	// The path to the directory on the filesystem to mount.
	MountPath string `json:"mount_path"`
	// The type of filesystem. Enum having values- CIFS, NFS3, NFS4.
	MountType string `json:"mount_type"`
	// The options for mount. The endpoint will return all default options and user specified options.
	Options *string `json:"options,omitempty"`
}

// NewHyperscaleMountPoint instantiates a new HyperscaleMountPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleMountPoint(hyperscaleInstanceId string, name string, hostname string, mountPath string, mountType string) *HyperscaleMountPoint {
	this := HyperscaleMountPoint{}
	this.HyperscaleInstanceId = hyperscaleInstanceId
	this.Name = name
	this.Hostname = hostname
	this.MountPath = mountPath
	this.MountType = mountType
	return &this
}

// NewHyperscaleMountPointWithDefaults instantiates a new HyperscaleMountPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleMountPointWithDefaults() *HyperscaleMountPoint {
	this := HyperscaleMountPoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HyperscaleMountPoint) SetId(v string) {
	o.Id = &v
}

// GetHyperscaleInstanceId returns the HyperscaleInstanceId field value
func (o *HyperscaleMountPoint) GetHyperscaleInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HyperscaleInstanceId
}

// GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field value
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetHyperscaleInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HyperscaleInstanceId, true
}

// SetHyperscaleInstanceId sets field value
func (o *HyperscaleMountPoint) SetHyperscaleInstanceId(v string) {
	o.HyperscaleInstanceId = v
}

// GetName returns the Name field value
func (o *HyperscaleMountPoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HyperscaleMountPoint) SetName(v string) {
	o.Name = v
}

// GetHostname returns the Hostname field value
func (o *HyperscaleMountPoint) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *HyperscaleMountPoint) SetHostname(v string) {
	o.Hostname = v
}

// GetMountPath returns the MountPath field value
func (o *HyperscaleMountPoint) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPath, true
}

// SetMountPath sets field value
func (o *HyperscaleMountPoint) SetMountPath(v string) {
	o.MountPath = v
}

// GetMountType returns the MountType field value
func (o *HyperscaleMountPoint) GetMountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountType
}

// GetMountTypeOk returns a tuple with the MountType field value
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetMountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountType, true
}

// SetMountType sets field value
func (o *HyperscaleMountPoint) SetMountType(v string) {
	o.MountType = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetOptions() string {
	if o == nil || IsNil(o.Options) {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *HyperscaleMountPoint) SetOptions(v string) {
	o.Options = &v
}

func (o HyperscaleMountPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleMountPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["hyperscale_instance_id"] = o.HyperscaleInstanceId
	toSerialize["name"] = o.Name
	toSerialize["hostname"] = o.Hostname
	toSerialize["mount_path"] = o.MountPath
	toSerialize["mount_type"] = o.MountType
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableHyperscaleMountPoint struct {
	value *HyperscaleMountPoint
	isSet bool
}

func (v NullableHyperscaleMountPoint) Get() *HyperscaleMountPoint {
	return v.value
}

func (v *NullableHyperscaleMountPoint) Set(val *HyperscaleMountPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleMountPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleMountPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleMountPoint(val *HyperscaleMountPoint) *NullableHyperscaleMountPoint {
	return &NullableHyperscaleMountPoint{value: val, isSet: true}
}

func (v NullableHyperscaleMountPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleMountPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


