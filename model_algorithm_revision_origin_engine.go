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

// checks if the AlgorithmRevisionOriginEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlgorithmRevisionOriginEngine{}

// AlgorithmRevisionOriginEngine A record of engine that an algorithm revision originates from.
type AlgorithmRevisionOriginEngine struct {
	// The id of the engine that this algorithm revision is originated from.
	EngineId *string `json:"engine_id,omitempty"`
	// The name of the engine that this algorithm revision is originated from.
	EngineName *string `json:"engine_name,omitempty"`
}

// NewAlgorithmRevisionOriginEngine instantiates a new AlgorithmRevisionOriginEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgorithmRevisionOriginEngine() *AlgorithmRevisionOriginEngine {
	this := AlgorithmRevisionOriginEngine{}
	return &this
}

// NewAlgorithmRevisionOriginEngineWithDefaults instantiates a new AlgorithmRevisionOriginEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgorithmRevisionOriginEngineWithDefaults() *AlgorithmRevisionOriginEngine {
	this := AlgorithmRevisionOriginEngine{}
	return &this
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *AlgorithmRevisionOriginEngine) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgorithmRevisionOriginEngine) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *AlgorithmRevisionOriginEngine) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *AlgorithmRevisionOriginEngine) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *AlgorithmRevisionOriginEngine) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgorithmRevisionOriginEngine) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *AlgorithmRevisionOriginEngine) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *AlgorithmRevisionOriginEngine) SetEngineName(v string) {
	o.EngineName = &v
}

func (o AlgorithmRevisionOriginEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlgorithmRevisionOriginEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	return toSerialize, nil
}

type NullableAlgorithmRevisionOriginEngine struct {
	value *AlgorithmRevisionOriginEngine
	isSet bool
}

func (v NullableAlgorithmRevisionOriginEngine) Get() *AlgorithmRevisionOriginEngine {
	return v.value
}

func (v *NullableAlgorithmRevisionOriginEngine) Set(val *AlgorithmRevisionOriginEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgorithmRevisionOriginEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgorithmRevisionOriginEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgorithmRevisionOriginEngine(val *AlgorithmRevisionOriginEngine) *NullableAlgorithmRevisionOriginEngine {
	return &NullableAlgorithmRevisionOriginEngine{value: val, isSet: true}
}

func (v NullableAlgorithmRevisionOriginEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgorithmRevisionOriginEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


