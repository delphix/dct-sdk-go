/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// ProvisionVDBBySnapshotParametersAllOf struct for ProvisionVDBBySnapshotParametersAllOf
type ProvisionVDBBySnapshotParametersAllOf struct {
	// The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.
	EngineId *string `json:"engine_id,omitempty"`
	// The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. If this property is not set, the data_source of the snapshot_id will be used.
	SourceDataId *string `json:"source_data_id,omitempty"`
}

// NewProvisionVDBBySnapshotParametersAllOf instantiates a new ProvisionVDBBySnapshotParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionVDBBySnapshotParametersAllOf() *ProvisionVDBBySnapshotParametersAllOf {
	this := ProvisionVDBBySnapshotParametersAllOf{}
	return &this
}

// NewProvisionVDBBySnapshotParametersAllOfWithDefaults instantiates a new ProvisionVDBBySnapshotParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionVDBBySnapshotParametersAllOfWithDefaults() *ProvisionVDBBySnapshotParametersAllOf {
	this := ProvisionVDBBySnapshotParametersAllOf{}
	return &this
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *ProvisionVDBBySnapshotParametersAllOf) GetEngineId() string {
	if o == nil || o.EngineId == nil {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBBySnapshotParametersAllOf) GetEngineIdOk() (*string, bool) {
	if o == nil || o.EngineId == nil {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *ProvisionVDBBySnapshotParametersAllOf) HasEngineId() bool {
	if o != nil && o.EngineId != nil {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *ProvisionVDBBySnapshotParametersAllOf) SetEngineId(v string) {
	o.EngineId = &v
}

// GetSourceDataId returns the SourceDataId field value if set, zero value otherwise.
func (o *ProvisionVDBBySnapshotParametersAllOf) GetSourceDataId() string {
	if o == nil || o.SourceDataId == nil {
		var ret string
		return ret
	}
	return *o.SourceDataId
}

// GetSourceDataIdOk returns a tuple with the SourceDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBBySnapshotParametersAllOf) GetSourceDataIdOk() (*string, bool) {
	if o == nil || o.SourceDataId == nil {
		return nil, false
	}
	return o.SourceDataId, true
}

// HasSourceDataId returns a boolean if a field has been set.
func (o *ProvisionVDBBySnapshotParametersAllOf) HasSourceDataId() bool {
	if o != nil && o.SourceDataId != nil {
		return true
	}

	return false
}

// SetSourceDataId gets a reference to the given string and assigns it to the SourceDataId field.
func (o *ProvisionVDBBySnapshotParametersAllOf) SetSourceDataId(v string) {
	o.SourceDataId = &v
}

func (o ProvisionVDBBySnapshotParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EngineId != nil {
		toSerialize["engine_id"] = o.EngineId
	}
	if o.SourceDataId != nil {
		toSerialize["source_data_id"] = o.SourceDataId
	}
	return json.Marshal(toSerialize)
}

type NullableProvisionVDBBySnapshotParametersAllOf struct {
	value *ProvisionVDBBySnapshotParametersAllOf
	isSet bool
}

func (v NullableProvisionVDBBySnapshotParametersAllOf) Get() *ProvisionVDBBySnapshotParametersAllOf {
	return v.value
}

func (v *NullableProvisionVDBBySnapshotParametersAllOf) Set(val *ProvisionVDBBySnapshotParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionVDBBySnapshotParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionVDBBySnapshotParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionVDBBySnapshotParametersAllOf(val *ProvisionVDBBySnapshotParametersAllOf) *NullableProvisionVDBBySnapshotParametersAllOf {
	return &NullableProvisionVDBBySnapshotParametersAllOf{value: val, isSet: true}
}

func (v NullableProvisionVDBBySnapshotParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionVDBBySnapshotParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


