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

// checks if the EnvironmentUserParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentUserParams{}

// EnvironmentUserParams struct for EnvironmentUserParams
type EnvironmentUserParams struct {
	// Username of the OS.
	Username *string `json:"username,omitempty"`
	// Password of the OS.
	Password *string `json:"password,omitempty"`
	// The name or reference of the vault from which to read the host credentials.
	Vault *string `json:"vault,omitempty"`
	// Delphix display name for the vault user
	VaultUsername *string `json:"vault_username,omitempty"`
	// Vault engine name where the credential is stored.
	HashicorpVaultEngine *string `json:"hashicorp_vault_engine,omitempty"`
	// Path in the vault engine where the credential is stored.
	HashicorpVaultSecretPath *string `json:"hashicorp_vault_secret_path,omitempty"`
	// Key for the username in the key-value store.
	HashicorpVaultUsernameKey *string `json:"hashicorp_vault_username_key,omitempty"`
	// Key for the password in the key-value store.
	HashicorpVaultSecretKey *string `json:"hashicorp_vault_secret_key,omitempty"`
	// Query to find a credential in the CyberArk vault.
	CyberarkVaultQueryString *string `json:"cyberark_vault_query_string,omitempty"`
	// Whether to use kerberos authentication.
	UseKerberosAuthentication *bool `json:"use_kerberos_authentication,omitempty"`
	// Whether to use public key authentication.
	UseEnginePublicKey *bool `json:"use_engine_public_key,omitempty"`
}

// NewEnvironmentUserParams instantiates a new EnvironmentUserParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentUserParams() *EnvironmentUserParams {
	this := EnvironmentUserParams{}
	return &this
}

// NewEnvironmentUserParamsWithDefaults instantiates a new EnvironmentUserParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentUserParamsWithDefaults() *EnvironmentUserParams {
	this := EnvironmentUserParams{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EnvironmentUserParams) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EnvironmentUserParams) SetPassword(v string) {
	o.Password = &v
}

// GetVault returns the Vault field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetVault() string {
	if o == nil || IsNil(o.Vault) {
		var ret string
		return ret
	}
	return *o.Vault
}

// GetVaultOk returns a tuple with the Vault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetVaultOk() (*string, bool) {
	if o == nil || IsNil(o.Vault) {
		return nil, false
	}
	return o.Vault, true
}

// HasVault returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasVault() bool {
	if o != nil && !IsNil(o.Vault) {
		return true
	}

	return false
}

// SetVault gets a reference to the given string and assigns it to the Vault field.
func (o *EnvironmentUserParams) SetVault(v string) {
	o.Vault = &v
}

// GetVaultUsername returns the VaultUsername field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetVaultUsername() string {
	if o == nil || IsNil(o.VaultUsername) {
		var ret string
		return ret
	}
	return *o.VaultUsername
}

// GetVaultUsernameOk returns a tuple with the VaultUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetVaultUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.VaultUsername) {
		return nil, false
	}
	return o.VaultUsername, true
}

// HasVaultUsername returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasVaultUsername() bool {
	if o != nil && !IsNil(o.VaultUsername) {
		return true
	}

	return false
}

// SetVaultUsername gets a reference to the given string and assigns it to the VaultUsername field.
func (o *EnvironmentUserParams) SetVaultUsername(v string) {
	o.VaultUsername = &v
}

// GetHashicorpVaultEngine returns the HashicorpVaultEngine field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetHashicorpVaultEngine() string {
	if o == nil || IsNil(o.HashicorpVaultEngine) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultEngine
}

// GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetHashicorpVaultEngineOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultEngine) {
		return nil, false
	}
	return o.HashicorpVaultEngine, true
}

// HasHashicorpVaultEngine returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasHashicorpVaultEngine() bool {
	if o != nil && !IsNil(o.HashicorpVaultEngine) {
		return true
	}

	return false
}

// SetHashicorpVaultEngine gets a reference to the given string and assigns it to the HashicorpVaultEngine field.
func (o *EnvironmentUserParams) SetHashicorpVaultEngine(v string) {
	o.HashicorpVaultEngine = &v
}

// GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetHashicorpVaultSecretPath() string {
	if o == nil || IsNil(o.HashicorpVaultSecretPath) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultSecretPath
}

// GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetHashicorpVaultSecretPathOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultSecretPath) {
		return nil, false
	}
	return o.HashicorpVaultSecretPath, true
}

// HasHashicorpVaultSecretPath returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasHashicorpVaultSecretPath() bool {
	if o != nil && !IsNil(o.HashicorpVaultSecretPath) {
		return true
	}

	return false
}

// SetHashicorpVaultSecretPath gets a reference to the given string and assigns it to the HashicorpVaultSecretPath field.
func (o *EnvironmentUserParams) SetHashicorpVaultSecretPath(v string) {
	o.HashicorpVaultSecretPath = &v
}

// GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetHashicorpVaultUsernameKey() string {
	if o == nil || IsNil(o.HashicorpVaultUsernameKey) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultUsernameKey
}

// GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetHashicorpVaultUsernameKeyOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultUsernameKey) {
		return nil, false
	}
	return o.HashicorpVaultUsernameKey, true
}

// HasHashicorpVaultUsernameKey returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasHashicorpVaultUsernameKey() bool {
	if o != nil && !IsNil(o.HashicorpVaultUsernameKey) {
		return true
	}

	return false
}

// SetHashicorpVaultUsernameKey gets a reference to the given string and assigns it to the HashicorpVaultUsernameKey field.
func (o *EnvironmentUserParams) SetHashicorpVaultUsernameKey(v string) {
	o.HashicorpVaultUsernameKey = &v
}

// GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetHashicorpVaultSecretKey() string {
	if o == nil || IsNil(o.HashicorpVaultSecretKey) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultSecretKey
}

// GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetHashicorpVaultSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultSecretKey) {
		return nil, false
	}
	return o.HashicorpVaultSecretKey, true
}

// HasHashicorpVaultSecretKey returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasHashicorpVaultSecretKey() bool {
	if o != nil && !IsNil(o.HashicorpVaultSecretKey) {
		return true
	}

	return false
}

// SetHashicorpVaultSecretKey gets a reference to the given string and assigns it to the HashicorpVaultSecretKey field.
func (o *EnvironmentUserParams) SetHashicorpVaultSecretKey(v string) {
	o.HashicorpVaultSecretKey = &v
}

// GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetCyberarkVaultQueryString() string {
	if o == nil || IsNil(o.CyberarkVaultQueryString) {
		var ret string
		return ret
	}
	return *o.CyberarkVaultQueryString
}

// GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetCyberarkVaultQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.CyberarkVaultQueryString) {
		return nil, false
	}
	return o.CyberarkVaultQueryString, true
}

// HasCyberarkVaultQueryString returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasCyberarkVaultQueryString() bool {
	if o != nil && !IsNil(o.CyberarkVaultQueryString) {
		return true
	}

	return false
}

// SetCyberarkVaultQueryString gets a reference to the given string and assigns it to the CyberarkVaultQueryString field.
func (o *EnvironmentUserParams) SetCyberarkVaultQueryString(v string) {
	o.CyberarkVaultQueryString = &v
}

// GetUseKerberosAuthentication returns the UseKerberosAuthentication field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetUseKerberosAuthentication() bool {
	if o == nil || IsNil(o.UseKerberosAuthentication) {
		var ret bool
		return ret
	}
	return *o.UseKerberosAuthentication
}

// GetUseKerberosAuthenticationOk returns a tuple with the UseKerberosAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetUseKerberosAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.UseKerberosAuthentication) {
		return nil, false
	}
	return o.UseKerberosAuthentication, true
}

// HasUseKerberosAuthentication returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasUseKerberosAuthentication() bool {
	if o != nil && !IsNil(o.UseKerberosAuthentication) {
		return true
	}

	return false
}

// SetUseKerberosAuthentication gets a reference to the given bool and assigns it to the UseKerberosAuthentication field.
func (o *EnvironmentUserParams) SetUseKerberosAuthentication(v bool) {
	o.UseKerberosAuthentication = &v
}

// GetUseEnginePublicKey returns the UseEnginePublicKey field value if set, zero value otherwise.
func (o *EnvironmentUserParams) GetUseEnginePublicKey() bool {
	if o == nil || IsNil(o.UseEnginePublicKey) {
		var ret bool
		return ret
	}
	return *o.UseEnginePublicKey
}

// GetUseEnginePublicKeyOk returns a tuple with the UseEnginePublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUserParams) GetUseEnginePublicKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseEnginePublicKey) {
		return nil, false
	}
	return o.UseEnginePublicKey, true
}

// HasUseEnginePublicKey returns a boolean if a field has been set.
func (o *EnvironmentUserParams) HasUseEnginePublicKey() bool {
	if o != nil && !IsNil(o.UseEnginePublicKey) {
		return true
	}

	return false
}

// SetUseEnginePublicKey gets a reference to the given bool and assigns it to the UseEnginePublicKey field.
func (o *EnvironmentUserParams) SetUseEnginePublicKey(v bool) {
	o.UseEnginePublicKey = &v
}

func (o EnvironmentUserParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentUserParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Vault) {
		toSerialize["vault"] = o.Vault
	}
	if !IsNil(o.VaultUsername) {
		toSerialize["vault_username"] = o.VaultUsername
	}
	if !IsNil(o.HashicorpVaultEngine) {
		toSerialize["hashicorp_vault_engine"] = o.HashicorpVaultEngine
	}
	if !IsNil(o.HashicorpVaultSecretPath) {
		toSerialize["hashicorp_vault_secret_path"] = o.HashicorpVaultSecretPath
	}
	if !IsNil(o.HashicorpVaultUsernameKey) {
		toSerialize["hashicorp_vault_username_key"] = o.HashicorpVaultUsernameKey
	}
	if !IsNil(o.HashicorpVaultSecretKey) {
		toSerialize["hashicorp_vault_secret_key"] = o.HashicorpVaultSecretKey
	}
	if !IsNil(o.CyberarkVaultQueryString) {
		toSerialize["cyberark_vault_query_string"] = o.CyberarkVaultQueryString
	}
	if !IsNil(o.UseKerberosAuthentication) {
		toSerialize["use_kerberos_authentication"] = o.UseKerberosAuthentication
	}
	if !IsNil(o.UseEnginePublicKey) {
		toSerialize["use_engine_public_key"] = o.UseEnginePublicKey
	}
	return toSerialize, nil
}

type NullableEnvironmentUserParams struct {
	value *EnvironmentUserParams
	isSet bool
}

func (v NullableEnvironmentUserParams) Get() *EnvironmentUserParams {
	return v.value
}

func (v *NullableEnvironmentUserParams) Set(val *EnvironmentUserParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentUserParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentUserParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentUserParams(val *EnvironmentUserParams) *NullableEnvironmentUserParams {
	return &NullableEnvironmentUserParams{value: val, isSet: true}
}

func (v NullableEnvironmentUserParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentUserParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


