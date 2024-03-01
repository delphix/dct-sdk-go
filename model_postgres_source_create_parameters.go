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

// checks if the PostgresSourceCreateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostgresSourceCreateParameters{}

// PostgresSourceCreateParameters struct for PostgresSourceCreateParameters
type PostgresSourceCreateParameters struct {
	// The name of the source.
	Name string `json:"name"`
	// The ID of the Repository onto which the source will be created.
	RepositoryId *string `json:"repository_id,omitempty"`
	// The ID of the engine to create the source on.
	EngineId *string `json:"engine_id,omitempty"`
	// The ID of the environment to create the source on.
	EnvironmentId *string `json:"environment_id,omitempty"`
}

// NewPostgresSourceCreateParameters instantiates a new PostgresSourceCreateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresSourceCreateParameters(name string) *PostgresSourceCreateParameters {
	this := PostgresSourceCreateParameters{}
	this.Name = name
	return &this
}

// NewPostgresSourceCreateParametersWithDefaults instantiates a new PostgresSourceCreateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresSourceCreateParametersWithDefaults() *PostgresSourceCreateParameters {
	this := PostgresSourceCreateParameters{}
	return &this
}

// GetName returns the Name field value
func (o *PostgresSourceCreateParameters) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostgresSourceCreateParameters) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostgresSourceCreateParameters) SetName(v string) {
	o.Name = v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *PostgresSourceCreateParameters) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId) {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresSourceCreateParameters) GetRepositoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryId) {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *PostgresSourceCreateParameters) HasRepositoryId() bool {
	if o != nil && !IsNil(o.RepositoryId) {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *PostgresSourceCreateParameters) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *PostgresSourceCreateParameters) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresSourceCreateParameters) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *PostgresSourceCreateParameters) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *PostgresSourceCreateParameters) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *PostgresSourceCreateParameters) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresSourceCreateParameters) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *PostgresSourceCreateParameters) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *PostgresSourceCreateParameters) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o PostgresSourceCreateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostgresSourceCreateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.RepositoryId) {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	return toSerialize, nil
}

type NullablePostgresSourceCreateParameters struct {
	value *PostgresSourceCreateParameters
	isSet bool
}

func (v NullablePostgresSourceCreateParameters) Get() *PostgresSourceCreateParameters {
	return v.value
}

func (v *NullablePostgresSourceCreateParameters) Set(val *PostgresSourceCreateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresSourceCreateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresSourceCreateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresSourceCreateParameters(val *PostgresSourceCreateParameters) *NullablePostgresSourceCreateParameters {
	return &NullablePostgresSourceCreateParameters{value: val, isSet: true}
}

func (v NullablePostgresSourceCreateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresSourceCreateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


