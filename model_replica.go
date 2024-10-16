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

// checks if the Replica type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Replica{}

// Replica struct for Replica
type Replica struct {
	// The ID of the replicated object.
	ReplicaId *string `json:"replica_id,omitempty"`
	// The ID of the replicated object's engine.
	ReplicaEngineId *string `json:"replica_engine_id,omitempty"`
	// The name of the replicated object's engine.
	ReplicaEngineName *string `json:"replica_engine_name,omitempty"`
	// The namespace id of the replicated object.
	ReplicaNamespaceId *string `json:"replica_namespace_id,omitempty"`
}

// NewReplica instantiates a new Replica object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplica() *Replica {
	this := Replica{}
	return &this
}

// NewReplicaWithDefaults instantiates a new Replica object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaWithDefaults() *Replica {
	this := Replica{}
	return &this
}

// GetReplicaId returns the ReplicaId field value if set, zero value otherwise.
func (o *Replica) GetReplicaId() string {
	if o == nil || IsNil(o.ReplicaId) {
		var ret string
		return ret
	}
	return *o.ReplicaId
}

// GetReplicaIdOk returns a tuple with the ReplicaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Replica) GetReplicaIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaId) {
		return nil, false
	}
	return o.ReplicaId, true
}

// HasReplicaId returns a boolean if a field has been set.
func (o *Replica) HasReplicaId() bool {
	if o != nil && !IsNil(o.ReplicaId) {
		return true
	}

	return false
}

// SetReplicaId gets a reference to the given string and assigns it to the ReplicaId field.
func (o *Replica) SetReplicaId(v string) {
	o.ReplicaId = &v
}

// GetReplicaEngineId returns the ReplicaEngineId field value if set, zero value otherwise.
func (o *Replica) GetReplicaEngineId() string {
	if o == nil || IsNil(o.ReplicaEngineId) {
		var ret string
		return ret
	}
	return *o.ReplicaEngineId
}

// GetReplicaEngineIdOk returns a tuple with the ReplicaEngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Replica) GetReplicaEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaEngineId) {
		return nil, false
	}
	return o.ReplicaEngineId, true
}

// HasReplicaEngineId returns a boolean if a field has been set.
func (o *Replica) HasReplicaEngineId() bool {
	if o != nil && !IsNil(o.ReplicaEngineId) {
		return true
	}

	return false
}

// SetReplicaEngineId gets a reference to the given string and assigns it to the ReplicaEngineId field.
func (o *Replica) SetReplicaEngineId(v string) {
	o.ReplicaEngineId = &v
}

// GetReplicaEngineName returns the ReplicaEngineName field value if set, zero value otherwise.
func (o *Replica) GetReplicaEngineName() string {
	if o == nil || IsNil(o.ReplicaEngineName) {
		var ret string
		return ret
	}
	return *o.ReplicaEngineName
}

// GetReplicaEngineNameOk returns a tuple with the ReplicaEngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Replica) GetReplicaEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaEngineName) {
		return nil, false
	}
	return o.ReplicaEngineName, true
}

// HasReplicaEngineName returns a boolean if a field has been set.
func (o *Replica) HasReplicaEngineName() bool {
	if o != nil && !IsNil(o.ReplicaEngineName) {
		return true
	}

	return false
}

// SetReplicaEngineName gets a reference to the given string and assigns it to the ReplicaEngineName field.
func (o *Replica) SetReplicaEngineName(v string) {
	o.ReplicaEngineName = &v
}

// GetReplicaNamespaceId returns the ReplicaNamespaceId field value if set, zero value otherwise.
func (o *Replica) GetReplicaNamespaceId() string {
	if o == nil || IsNil(o.ReplicaNamespaceId) {
		var ret string
		return ret
	}
	return *o.ReplicaNamespaceId
}

// GetReplicaNamespaceIdOk returns a tuple with the ReplicaNamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Replica) GetReplicaNamespaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaNamespaceId) {
		return nil, false
	}
	return o.ReplicaNamespaceId, true
}

// HasReplicaNamespaceId returns a boolean if a field has been set.
func (o *Replica) HasReplicaNamespaceId() bool {
	if o != nil && !IsNil(o.ReplicaNamespaceId) {
		return true
	}

	return false
}

// SetReplicaNamespaceId gets a reference to the given string and assigns it to the ReplicaNamespaceId field.
func (o *Replica) SetReplicaNamespaceId(v string) {
	o.ReplicaNamespaceId = &v
}

func (o Replica) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Replica) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReplicaId) {
		toSerialize["replica_id"] = o.ReplicaId
	}
	if !IsNil(o.ReplicaEngineId) {
		toSerialize["replica_engine_id"] = o.ReplicaEngineId
	}
	if !IsNil(o.ReplicaEngineName) {
		toSerialize["replica_engine_name"] = o.ReplicaEngineName
	}
	if !IsNil(o.ReplicaNamespaceId) {
		toSerialize["replica_namespace_id"] = o.ReplicaNamespaceId
	}
	return toSerialize, nil
}

type NullableReplica struct {
	value *Replica
	isSet bool
}

func (v NullableReplica) Get() *Replica {
	return v.value
}

func (v *NullableReplica) Set(val *Replica) {
	v.value = val
	v.isSet = true
}

func (v NullableReplica) IsSet() bool {
	return v.isSet
}

func (v *NullableReplica) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplica(val *Replica) *NullableReplica {
	return &NullableReplica{value: val, isSet: true}
}

func (v NullableReplica) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplica) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


