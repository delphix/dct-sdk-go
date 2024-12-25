/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NetbackupConnectivityCheckParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetbackupConnectivityCheckParameters{}

// NetbackupConnectivityCheckParameters Parameters to test NetBackup master server and client connectivity on an environment.
type NetbackupConnectivityCheckParameters struct {
	// Id of the target environment to test NetBackup connectivity from.
	EnvironmentId string `json:"environment_id"`
	// Id of the environment user.
	EnvironmentUserId string `json:"environment_user_id"`
	// The name of the NetBackup master server to attempt to connect to.
	MasterServerName string `json:"master_server_name"`
	// The name of the NetBackup client to attempt to connect with.
	SourceClientName string `json:"source_client_name"`
}

type _NetbackupConnectivityCheckParameters NetbackupConnectivityCheckParameters

// NewNetbackupConnectivityCheckParameters instantiates a new NetbackupConnectivityCheckParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetbackupConnectivityCheckParameters(environmentId string, environmentUserId string, masterServerName string, sourceClientName string) *NetbackupConnectivityCheckParameters {
	this := NetbackupConnectivityCheckParameters{}
	this.EnvironmentId = environmentId
	this.EnvironmentUserId = environmentUserId
	this.MasterServerName = masterServerName
	this.SourceClientName = sourceClientName
	return &this
}

// NewNetbackupConnectivityCheckParametersWithDefaults instantiates a new NetbackupConnectivityCheckParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetbackupConnectivityCheckParametersWithDefaults() *NetbackupConnectivityCheckParameters {
	this := NetbackupConnectivityCheckParameters{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *NetbackupConnectivityCheckParameters) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *NetbackupConnectivityCheckParameters) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *NetbackupConnectivityCheckParameters) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetEnvironmentUserId returns the EnvironmentUserId field value
func (o *NetbackupConnectivityCheckParameters) GetEnvironmentUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentUserId
}

// GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field value
// and a boolean to check if the value has been set.
func (o *NetbackupConnectivityCheckParameters) GetEnvironmentUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentUserId, true
}

// SetEnvironmentUserId sets field value
func (o *NetbackupConnectivityCheckParameters) SetEnvironmentUserId(v string) {
	o.EnvironmentUserId = v
}

// GetMasterServerName returns the MasterServerName field value
func (o *NetbackupConnectivityCheckParameters) GetMasterServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MasterServerName
}

// GetMasterServerNameOk returns a tuple with the MasterServerName field value
// and a boolean to check if the value has been set.
func (o *NetbackupConnectivityCheckParameters) GetMasterServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MasterServerName, true
}

// SetMasterServerName sets field value
func (o *NetbackupConnectivityCheckParameters) SetMasterServerName(v string) {
	o.MasterServerName = v
}

// GetSourceClientName returns the SourceClientName field value
func (o *NetbackupConnectivityCheckParameters) GetSourceClientName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceClientName
}

// GetSourceClientNameOk returns a tuple with the SourceClientName field value
// and a boolean to check if the value has been set.
func (o *NetbackupConnectivityCheckParameters) GetSourceClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceClientName, true
}

// SetSourceClientName sets field value
func (o *NetbackupConnectivityCheckParameters) SetSourceClientName(v string) {
	o.SourceClientName = v
}

func (o NetbackupConnectivityCheckParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetbackupConnectivityCheckParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["environment_user_id"] = o.EnvironmentUserId
	toSerialize["master_server_name"] = o.MasterServerName
	toSerialize["source_client_name"] = o.SourceClientName
	return toSerialize, nil
}

func (o *NetbackupConnectivityCheckParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment_id",
		"environment_user_id",
		"master_server_name",
		"source_client_name",
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

	varNetbackupConnectivityCheckParameters := _NetbackupConnectivityCheckParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetbackupConnectivityCheckParameters)

	if err != nil {
		return err
	}

	*o = NetbackupConnectivityCheckParameters(varNetbackupConnectivityCheckParameters)

	return err
}

type NullableNetbackupConnectivityCheckParameters struct {
	value *NetbackupConnectivityCheckParameters
	isSet bool
}

func (v NullableNetbackupConnectivityCheckParameters) Get() *NetbackupConnectivityCheckParameters {
	return v.value
}

func (v *NullableNetbackupConnectivityCheckParameters) Set(val *NetbackupConnectivityCheckParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableNetbackupConnectivityCheckParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableNetbackupConnectivityCheckParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetbackupConnectivityCheckParameters(val *NetbackupConnectivityCheckParameters) *NullableNetbackupConnectivityCheckParameters {
	return &NullableNetbackupConnectivityCheckParameters{value: val, isSet: true}
}

func (v NullableNetbackupConnectivityCheckParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetbackupConnectivityCheckParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


