/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ConnectorUpdateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorUpdateParameters{}

// ConnectorUpdateParameters Parameters used to update a Masking Connector.
type ConnectorUpdateParameters struct {
	// The Connector name.
	Name *string `json:"name,omitempty"`
	// The network hostname or IP address of the database server.
	Hostname *string `json:"hostname,omitempty"`
	// The TCP port of the server.
	Port *int32 `json:"port,omitempty"`
	// The username this Connector will use to connect to the database.
	Username *string `json:"username,omitempty"`
	// The password this Connector will use to connect to the database.
	Password *string `json:"password,omitempty"`
}

// NewConnectorUpdateParameters instantiates a new ConnectorUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorUpdateParameters() *ConnectorUpdateParameters {
	this := ConnectorUpdateParameters{}
	return &this
}

// NewConnectorUpdateParametersWithDefaults instantiates a new ConnectorUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorUpdateParametersWithDefaults() *ConnectorUpdateParameters {
	this := ConnectorUpdateParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorUpdateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUpdateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorUpdateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorUpdateParameters) SetName(v string) {
	o.Name = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ConnectorUpdateParameters) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUpdateParameters) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ConnectorUpdateParameters) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ConnectorUpdateParameters) SetHostname(v string) {
	o.Hostname = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ConnectorUpdateParameters) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUpdateParameters) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ConnectorUpdateParameters) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ConnectorUpdateParameters) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ConnectorUpdateParameters) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUpdateParameters) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ConnectorUpdateParameters) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ConnectorUpdateParameters) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ConnectorUpdateParameters) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUpdateParameters) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ConnectorUpdateParameters) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ConnectorUpdateParameters) SetPassword(v string) {
	o.Password = &v
}

func (o ConnectorUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableConnectorUpdateParameters struct {
	value *ConnectorUpdateParameters
	isSet bool
}

func (v NullableConnectorUpdateParameters) Get() *ConnectorUpdateParameters {
	return v.value
}

func (v *NullableConnectorUpdateParameters) Set(val *ConnectorUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorUpdateParameters(val *ConnectorUpdateParameters) *NullableConnectorUpdateParameters {
	return &NullableConnectorUpdateParameters{value: val, isSet: true}
}

func (v NullableConnectorUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


