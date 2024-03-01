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

// checks if the HyperscaleDataset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleDataset{}

// HyperscaleDataset A Hyperscale Dataset.
type HyperscaleDataset struct {
	// The ID of the Hyperscale Dataset.
	Id *string `json:"id,omitempty"`
	// The ID of the Hyperscale instance of this Dataset.
	HyperscaleInstanceId *string `json:"hyperscale_instance_id,omitempty"`
	DataType *HyperscaleDataTypeEnum `json:"data_type,omitempty"`
	// The Id of the Hyperscale Mount Point used for this Dataset.
	MountPointId *string `json:"mount_point_id,omitempty"`
	// Id the Hyperscale Connector used to read sensitive data and write masked data.
	ConnectorId *string `json:"connector_id,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

// NewHyperscaleDataset instantiates a new HyperscaleDataset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleDataset() *HyperscaleDataset {
	this := HyperscaleDataset{}
	return &this
}

// NewHyperscaleDatasetWithDefaults instantiates a new HyperscaleDataset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleDatasetWithDefaults() *HyperscaleDataset {
	this := HyperscaleDataset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HyperscaleDataset) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDataset) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HyperscaleDataset) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HyperscaleDataset) SetId(v string) {
	o.Id = &v
}

// GetHyperscaleInstanceId returns the HyperscaleInstanceId field value if set, zero value otherwise.
func (o *HyperscaleDataset) GetHyperscaleInstanceId() string {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		var ret string
		return ret
	}
	return *o.HyperscaleInstanceId
}

// GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDataset) GetHyperscaleInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		return nil, false
	}
	return o.HyperscaleInstanceId, true
}

// HasHyperscaleInstanceId returns a boolean if a field has been set.
func (o *HyperscaleDataset) HasHyperscaleInstanceId() bool {
	if o != nil && !IsNil(o.HyperscaleInstanceId) {
		return true
	}

	return false
}

// SetHyperscaleInstanceId gets a reference to the given string and assigns it to the HyperscaleInstanceId field.
func (o *HyperscaleDataset) SetHyperscaleInstanceId(v string) {
	o.HyperscaleInstanceId = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *HyperscaleDataset) GetDataType() HyperscaleDataTypeEnum {
	if o == nil || IsNil(o.DataType) {
		var ret HyperscaleDataTypeEnum
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDataset) GetDataTypeOk() (*HyperscaleDataTypeEnum, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *HyperscaleDataset) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given HyperscaleDataTypeEnum and assigns it to the DataType field.
func (o *HyperscaleDataset) SetDataType(v HyperscaleDataTypeEnum) {
	o.DataType = &v
}

// GetMountPointId returns the MountPointId field value if set, zero value otherwise.
func (o *HyperscaleDataset) GetMountPointId() string {
	if o == nil || IsNil(o.MountPointId) {
		var ret string
		return ret
	}
	return *o.MountPointId
}

// GetMountPointIdOk returns a tuple with the MountPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDataset) GetMountPointIdOk() (*string, bool) {
	if o == nil || IsNil(o.MountPointId) {
		return nil, false
	}
	return o.MountPointId, true
}

// HasMountPointId returns a boolean if a field has been set.
func (o *HyperscaleDataset) HasMountPointId() bool {
	if o != nil && !IsNil(o.MountPointId) {
		return true
	}

	return false
}

// SetMountPointId gets a reference to the given string and assigns it to the MountPointId field.
func (o *HyperscaleDataset) SetMountPointId(v string) {
	o.MountPointId = &v
}

// GetConnectorId returns the ConnectorId field value if set, zero value otherwise.
func (o *HyperscaleDataset) GetConnectorId() string {
	if o == nil || IsNil(o.ConnectorId) {
		var ret string
		return ret
	}
	return *o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDataset) GetConnectorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorId) {
		return nil, false
	}
	return o.ConnectorId, true
}

// HasConnectorId returns a boolean if a field has been set.
func (o *HyperscaleDataset) HasConnectorId() bool {
	if o != nil && !IsNil(o.ConnectorId) {
		return true
	}

	return false
}

// SetConnectorId gets a reference to the given string and assigns it to the ConnectorId field.
func (o *HyperscaleDataset) SetConnectorId(v string) {
	o.ConnectorId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *HyperscaleDataset) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDataset) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *HyperscaleDataset) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *HyperscaleDataset) SetTags(v []Tag) {
	o.Tags = v
}

func (o HyperscaleDataset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleDataset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.HyperscaleInstanceId) {
		toSerialize["hyperscale_instance_id"] = o.HyperscaleInstanceId
	}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !IsNil(o.MountPointId) {
		toSerialize["mount_point_id"] = o.MountPointId
	}
	if !IsNil(o.ConnectorId) {
		toSerialize["connector_id"] = o.ConnectorId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableHyperscaleDataset struct {
	value *HyperscaleDataset
	isSet bool
}

func (v NullableHyperscaleDataset) Get() *HyperscaleDataset {
	return v.value
}

func (v *NullableHyperscaleDataset) Set(val *HyperscaleDataset) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleDataset(val *HyperscaleDataset) *NullableHyperscaleDataset {
	return &NullableHyperscaleDataset{value: val, isSet: true}
}

func (v NullableHyperscaleDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


