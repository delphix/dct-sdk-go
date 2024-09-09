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

// checks if the SyncEnginesHyperscaleParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncEnginesHyperscaleParameters{}

// SyncEnginesHyperscaleParameters Parameters to sync the global object from a source engine to a list of target engines registered with a Hyperscale Instance.
type SyncEnginesHyperscaleParameters struct {
	// The ID of the engine to copy the data from.
	SourceEngineId string `json:"source_engine_id"`
	// The IDs of the target engines to copy the data into.
	TargetEngineIds []string `json:"target_engine_ids"`
}

type _SyncEnginesHyperscaleParameters SyncEnginesHyperscaleParameters

// NewSyncEnginesHyperscaleParameters instantiates a new SyncEnginesHyperscaleParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncEnginesHyperscaleParameters(sourceEngineId string, targetEngineIds []string) *SyncEnginesHyperscaleParameters {
	this := SyncEnginesHyperscaleParameters{}
	this.SourceEngineId = sourceEngineId
	this.TargetEngineIds = targetEngineIds
	return &this
}

// NewSyncEnginesHyperscaleParametersWithDefaults instantiates a new SyncEnginesHyperscaleParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncEnginesHyperscaleParametersWithDefaults() *SyncEnginesHyperscaleParameters {
	this := SyncEnginesHyperscaleParameters{}
	return &this
}

// GetSourceEngineId returns the SourceEngineId field value
func (o *SyncEnginesHyperscaleParameters) GetSourceEngineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceEngineId
}

// GetSourceEngineIdOk returns a tuple with the SourceEngineId field value
// and a boolean to check if the value has been set.
func (o *SyncEnginesHyperscaleParameters) GetSourceEngineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceEngineId, true
}

// SetSourceEngineId sets field value
func (o *SyncEnginesHyperscaleParameters) SetSourceEngineId(v string) {
	o.SourceEngineId = v
}

// GetTargetEngineIds returns the TargetEngineIds field value
func (o *SyncEnginesHyperscaleParameters) GetTargetEngineIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetEngineIds
}

// GetTargetEngineIdsOk returns a tuple with the TargetEngineIds field value
// and a boolean to check if the value has been set.
func (o *SyncEnginesHyperscaleParameters) GetTargetEngineIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetEngineIds, true
}

// SetTargetEngineIds sets field value
func (o *SyncEnginesHyperscaleParameters) SetTargetEngineIds(v []string) {
	o.TargetEngineIds = v
}

func (o SyncEnginesHyperscaleParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncEnginesHyperscaleParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_engine_id"] = o.SourceEngineId
	toSerialize["target_engine_ids"] = o.TargetEngineIds
	return toSerialize, nil
}

func (o *SyncEnginesHyperscaleParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_engine_id",
		"target_engine_ids",
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

	varSyncEnginesHyperscaleParameters := _SyncEnginesHyperscaleParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSyncEnginesHyperscaleParameters)

	if err != nil {
		return err
	}

	*o = SyncEnginesHyperscaleParameters(varSyncEnginesHyperscaleParameters)

	return err
}

type NullableSyncEnginesHyperscaleParameters struct {
	value *SyncEnginesHyperscaleParameters
	isSet bool
}

func (v NullableSyncEnginesHyperscaleParameters) Get() *SyncEnginesHyperscaleParameters {
	return v.value
}

func (v *NullableSyncEnginesHyperscaleParameters) Set(val *SyncEnginesHyperscaleParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncEnginesHyperscaleParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncEnginesHyperscaleParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncEnginesHyperscaleParameters(val *SyncEnginesHyperscaleParameters) *NullableSyncEnginesHyperscaleParameters {
	return &NullableSyncEnginesHyperscaleParameters{value: val, isSet: true}
}

func (v NullableSyncEnginesHyperscaleParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncEnginesHyperscaleParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


