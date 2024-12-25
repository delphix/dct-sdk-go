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

// checks if the ProxyConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyConfiguration{}

// ProxyConfiguration Web proxy configuration that is used to communicate with Delphix Corp. for support, troubleshooting, upgrades, updates, and patches.
type ProxyConfiguration struct {
	// The host name or IP address of the proxy server.
	Host string `json:"host"`
	// The port number of the proxy server.
	Port int32 `json:"port"`
	// The username to use when authenticating with the proxy server.
	Username *string `json:"username,omitempty"`
	// The password to use when authenticating with the proxy server.
	Password *string `json:"password,omitempty"`
	// When set, these settings are enabled. True by default.
	Enabled bool `json:"enabled"`
	// File name of a truststore which can be used to validate the TLS certificate of the proxy server. The truststore must be available at /etc/config/certs/<truststore_filename>
	TruststoreFilename *string `json:"truststore_filename,omitempty" validate:"regexp=^$|^[a-zA-Z0-9_\\\\.\\\\-]+$"`
	// Password for reading trustStore file provided in 'truststore_filename' property
	TruststorePassword *string `json:"truststore_password,omitempty"`
}

type _ProxyConfiguration ProxyConfiguration

// NewProxyConfiguration instantiates a new ProxyConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyConfiguration(host string, port int32, enabled bool) *ProxyConfiguration {
	this := ProxyConfiguration{}
	this.Host = host
	this.Port = port
	this.Enabled = enabled
	return &this
}

// NewProxyConfigurationWithDefaults instantiates a new ProxyConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyConfigurationWithDefaults() *ProxyConfiguration {
	this := ProxyConfiguration{}
	return &this
}

// GetHost returns the Host field value
func (o *ProxyConfiguration) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ProxyConfiguration) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ProxyConfiguration) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *ProxyConfiguration) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ProxyConfiguration) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ProxyConfiguration) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ProxyConfiguration) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ProxyConfiguration) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ProxyConfiguration) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ProxyConfiguration) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ProxyConfiguration) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ProxyConfiguration) SetPassword(v string) {
	o.Password = &v
}

// GetEnabled returns the Enabled field value
func (o *ProxyConfiguration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProxyConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ProxyConfiguration) SetEnabled(v bool) {
	o.Enabled = v
}

// GetTruststoreFilename returns the TruststoreFilename field value if set, zero value otherwise.
func (o *ProxyConfiguration) GetTruststoreFilename() string {
	if o == nil || IsNil(o.TruststoreFilename) {
		var ret string
		return ret
	}
	return *o.TruststoreFilename
}

// GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfiguration) GetTruststoreFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.TruststoreFilename) {
		return nil, false
	}
	return o.TruststoreFilename, true
}

// HasTruststoreFilename returns a boolean if a field has been set.
func (o *ProxyConfiguration) HasTruststoreFilename() bool {
	if o != nil && !IsNil(o.TruststoreFilename) {
		return true
	}

	return false
}

// SetTruststoreFilename gets a reference to the given string and assigns it to the TruststoreFilename field.
func (o *ProxyConfiguration) SetTruststoreFilename(v string) {
	o.TruststoreFilename = &v
}

// GetTruststorePassword returns the TruststorePassword field value if set, zero value otherwise.
func (o *ProxyConfiguration) GetTruststorePassword() string {
	if o == nil || IsNil(o.TruststorePassword) {
		var ret string
		return ret
	}
	return *o.TruststorePassword
}

// GetTruststorePasswordOk returns a tuple with the TruststorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfiguration) GetTruststorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.TruststorePassword) {
		return nil, false
	}
	return o.TruststorePassword, true
}

// HasTruststorePassword returns a boolean if a field has been set.
func (o *ProxyConfiguration) HasTruststorePassword() bool {
	if o != nil && !IsNil(o.TruststorePassword) {
		return true
	}

	return false
}

// SetTruststorePassword gets a reference to the given string and assigns it to the TruststorePassword field.
func (o *ProxyConfiguration) SetTruststorePassword(v string) {
	o.TruststorePassword = &v
}

func (o ProxyConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.TruststoreFilename) {
		toSerialize["truststore_filename"] = o.TruststoreFilename
	}
	if !IsNil(o.TruststorePassword) {
		toSerialize["truststore_password"] = o.TruststorePassword
	}
	return toSerialize, nil
}

func (o *ProxyConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"port",
		"enabled",
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

	varProxyConfiguration := _ProxyConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProxyConfiguration)

	if err != nil {
		return err
	}

	*o = ProxyConfiguration(varProxyConfiguration)

	return err
}

type NullableProxyConfiguration struct {
	value *ProxyConfiguration
	isSet bool
}

func (v NullableProxyConfiguration) Get() *ProxyConfiguration {
	return v.value
}

func (v *NullableProxyConfiguration) Set(val *ProxyConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyConfiguration(val *ProxyConfiguration) *NullableProxyConfiguration {
	return &NullableProxyConfiguration{value: val, isSet: true}
}

func (v NullableProxyConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


