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

// checks if the ClusterNodeInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterNodeInstance{}

// ClusterNodeInstance struct for ClusterNodeInstance
type ClusterNodeInstance struct {
	// The cluster node id, name or addresses for this provision operation
	NodeReference string `json:"node_reference"`
	// The instance number for this provision operation
	InstanceNumber int32 `json:"instance_number"`
	// The instance name for this provision operation
	InstanceName string `json:"instance_name"`
}

type _ClusterNodeInstance ClusterNodeInstance

// NewClusterNodeInstance instantiates a new ClusterNodeInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNodeInstance(nodeReference string, instanceNumber int32, instanceName string) *ClusterNodeInstance {
	this := ClusterNodeInstance{}
	this.NodeReference = nodeReference
	this.InstanceNumber = instanceNumber
	this.InstanceName = instanceName
	return &this
}

// NewClusterNodeInstanceWithDefaults instantiates a new ClusterNodeInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodeInstanceWithDefaults() *ClusterNodeInstance {
	this := ClusterNodeInstance{}
	return &this
}

// GetNodeReference returns the NodeReference field value
func (o *ClusterNodeInstance) GetNodeReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeReference
}

// GetNodeReferenceOk returns a tuple with the NodeReference field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeInstance) GetNodeReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeReference, true
}

// SetNodeReference sets field value
func (o *ClusterNodeInstance) SetNodeReference(v string) {
	o.NodeReference = v
}

// GetInstanceNumber returns the InstanceNumber field value
func (o *ClusterNodeInstance) GetInstanceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InstanceNumber
}

// GetInstanceNumberOk returns a tuple with the InstanceNumber field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeInstance) GetInstanceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceNumber, true
}

// SetInstanceNumber sets field value
func (o *ClusterNodeInstance) SetInstanceNumber(v int32) {
	o.InstanceNumber = v
}

// GetInstanceName returns the InstanceName field value
func (o *ClusterNodeInstance) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value
func (o *ClusterNodeInstance) SetInstanceName(v string) {
	o.InstanceName = v
}

func (o ClusterNodeInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterNodeInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_reference"] = o.NodeReference
	toSerialize["instance_number"] = o.InstanceNumber
	toSerialize["instance_name"] = o.InstanceName
	return toSerialize, nil
}

func (o *ClusterNodeInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_reference",
		"instance_number",
		"instance_name",
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

	varClusterNodeInstance := _ClusterNodeInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterNodeInstance)

	if err != nil {
		return err
	}

	*o = ClusterNodeInstance(varClusterNodeInstance)

	return err
}

type NullableClusterNodeInstance struct {
	value *ClusterNodeInstance
	isSet bool
}

func (v NullableClusterNodeInstance) Get() *ClusterNodeInstance {
	return v.value
}

func (v *NullableClusterNodeInstance) Set(val *ClusterNodeInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNodeInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNodeInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNodeInstance(val *ClusterNodeInstance) *NullableClusterNodeInstance {
	return &NullableClusterNodeInstance{value: val, isSet: true}
}

func (v NullableClusterNodeInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNodeInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


