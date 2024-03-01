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

// checks if the SourceOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceOperation{}

// SourceOperation struct for SourceOperation
type SourceOperation struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Shell *string `json:"shell,omitempty"`
	// List of environment variables that will contain credentials for this operation.
	CredentialsEnvVars []CredentialsEnvVariable `json:"credentials_env_vars,omitempty"`
}

// NewSourceOperation instantiates a new SourceOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceOperation(name string, command string) *SourceOperation {
	this := SourceOperation{}
	this.Name = name
	this.Command = command
	var shell string = "bash"
	this.Shell = &shell
	return &this
}

// NewSourceOperationWithDefaults instantiates a new SourceOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceOperationWithDefaults() *SourceOperation {
	this := SourceOperation{}
	var shell string = "bash"
	this.Shell = &shell
	return &this
}

// GetName returns the Name field value
func (o *SourceOperation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceOperation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceOperation) SetName(v string) {
	o.Name = v
}

// GetCommand returns the Command field value
func (o *SourceOperation) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *SourceOperation) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *SourceOperation) SetCommand(v string) {
	o.Command = v
}

// GetShell returns the Shell field value if set, zero value otherwise.
func (o *SourceOperation) GetShell() string {
	if o == nil || IsNil(o.Shell) {
		var ret string
		return ret
	}
	return *o.Shell
}

// GetShellOk returns a tuple with the Shell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOperation) GetShellOk() (*string, bool) {
	if o == nil || IsNil(o.Shell) {
		return nil, false
	}
	return o.Shell, true
}

// HasShell returns a boolean if a field has been set.
func (o *SourceOperation) HasShell() bool {
	if o != nil && !IsNil(o.Shell) {
		return true
	}

	return false
}

// SetShell gets a reference to the given string and assigns it to the Shell field.
func (o *SourceOperation) SetShell(v string) {
	o.Shell = &v
}

// GetCredentialsEnvVars returns the CredentialsEnvVars field value if set, zero value otherwise.
func (o *SourceOperation) GetCredentialsEnvVars() []CredentialsEnvVariable {
	if o == nil || IsNil(o.CredentialsEnvVars) {
		var ret []CredentialsEnvVariable
		return ret
	}
	return o.CredentialsEnvVars
}

// GetCredentialsEnvVarsOk returns a tuple with the CredentialsEnvVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOperation) GetCredentialsEnvVarsOk() ([]CredentialsEnvVariable, bool) {
	if o == nil || IsNil(o.CredentialsEnvVars) {
		return nil, false
	}
	return o.CredentialsEnvVars, true
}

// HasCredentialsEnvVars returns a boolean if a field has been set.
func (o *SourceOperation) HasCredentialsEnvVars() bool {
	if o != nil && !IsNil(o.CredentialsEnvVars) {
		return true
	}

	return false
}

// SetCredentialsEnvVars gets a reference to the given []CredentialsEnvVariable and assigns it to the CredentialsEnvVars field.
func (o *SourceOperation) SetCredentialsEnvVars(v []CredentialsEnvVariable) {
	o.CredentialsEnvVars = v
}

func (o SourceOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["command"] = o.Command
	if !IsNil(o.Shell) {
		toSerialize["shell"] = o.Shell
	}
	if !IsNil(o.CredentialsEnvVars) {
		toSerialize["credentials_env_vars"] = o.CredentialsEnvVars
	}
	return toSerialize, nil
}

type NullableSourceOperation struct {
	value *SourceOperation
	isSet bool
}

func (v NullableSourceOperation) Get() *SourceOperation {
	return v.value
}

func (v *NullableSourceOperation) Set(val *SourceOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceOperation(val *SourceOperation) *NullableSourceOperation {
	return &NullableSourceOperation{value: val, isSet: true}
}

func (v NullableSourceOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


