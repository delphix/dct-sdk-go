# EnvironmentUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username of the OS. | [optional] 
**Password** | Pointer to **string** | Password of the OS. | [optional] 
**Vault** | Pointer to **string** | The name or reference of the vault from which to read the host credentials. | [optional] 
**VaultUsername** | Pointer to **string** | Delphix display name for the vault user | [optional] 
**HashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored. | [optional] 
**HashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored. | [optional] 
**HashicorpVaultUsernameKey** | Pointer to **string** | Key for the username in the key-value store. | [optional] 
**HashicorpVaultSecretKey** | Pointer to **string** | Key for the password in the key-value store. | [optional] 
**CyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault. | [optional] 
**UseKerberosAuthentication** | Pointer to **bool** | Whether to use kerberos authentication. | [optional] 
**UseEnginePublicKey** | Pointer to **bool** | Whether to use public key authentication. | [optional] 

## Methods

### NewEnvironmentUserParams

`func NewEnvironmentUserParams() *EnvironmentUserParams`

NewEnvironmentUserParams instantiates a new EnvironmentUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentUserParamsWithDefaults

`func NewEnvironmentUserParamsWithDefaults() *EnvironmentUserParams`

NewEnvironmentUserParamsWithDefaults instantiates a new EnvironmentUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *EnvironmentUserParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnvironmentUserParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnvironmentUserParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EnvironmentUserParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *EnvironmentUserParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnvironmentUserParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnvironmentUserParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnvironmentUserParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVault

`func (o *EnvironmentUserParams) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *EnvironmentUserParams) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *EnvironmentUserParams) SetVault(v string)`

SetVault sets Vault field to given value.

### HasVault

`func (o *EnvironmentUserParams) HasVault() bool`

HasVault returns a boolean if a field has been set.

### GetVaultUsername

`func (o *EnvironmentUserParams) GetVaultUsername() string`

GetVaultUsername returns the VaultUsername field if non-nil, zero value otherwise.

### GetVaultUsernameOk

`func (o *EnvironmentUserParams) GetVaultUsernameOk() (*string, bool)`

GetVaultUsernameOk returns a tuple with the VaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultUsername

`func (o *EnvironmentUserParams) SetVaultUsername(v string)`

SetVaultUsername sets VaultUsername field to given value.

### HasVaultUsername

`func (o *EnvironmentUserParams) HasVaultUsername() bool`

HasVaultUsername returns a boolean if a field has been set.

### GetHashicorpVaultEngine

`func (o *EnvironmentUserParams) GetHashicorpVaultEngine() string`

GetHashicorpVaultEngine returns the HashicorpVaultEngine field if non-nil, zero value otherwise.

### GetHashicorpVaultEngineOk

`func (o *EnvironmentUserParams) GetHashicorpVaultEngineOk() (*string, bool)`

GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultEngine

`func (o *EnvironmentUserParams) SetHashicorpVaultEngine(v string)`

SetHashicorpVaultEngine sets HashicorpVaultEngine field to given value.

### HasHashicorpVaultEngine

`func (o *EnvironmentUserParams) HasHashicorpVaultEngine() bool`

HasHashicorpVaultEngine returns a boolean if a field has been set.

### GetHashicorpVaultSecretPath

`func (o *EnvironmentUserParams) GetHashicorpVaultSecretPath() string`

GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretPathOk

`func (o *EnvironmentUserParams) GetHashicorpVaultSecretPathOk() (*string, bool)`

GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretPath

`func (o *EnvironmentUserParams) SetHashicorpVaultSecretPath(v string)`

SetHashicorpVaultSecretPath sets HashicorpVaultSecretPath field to given value.

### HasHashicorpVaultSecretPath

`func (o *EnvironmentUserParams) HasHashicorpVaultSecretPath() bool`

HasHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetHashicorpVaultUsernameKey

`func (o *EnvironmentUserParams) GetHashicorpVaultUsernameKey() string`

GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameKeyOk

`func (o *EnvironmentUserParams) GetHashicorpVaultUsernameKeyOk() (*string, bool)`

GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameKey

`func (o *EnvironmentUserParams) SetHashicorpVaultUsernameKey(v string)`

SetHashicorpVaultUsernameKey sets HashicorpVaultUsernameKey field to given value.

### HasHashicorpVaultUsernameKey

`func (o *EnvironmentUserParams) HasHashicorpVaultUsernameKey() bool`

HasHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetHashicorpVaultSecretKey

`func (o *EnvironmentUserParams) GetHashicorpVaultSecretKey() string`

GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretKeyOk

`func (o *EnvironmentUserParams) GetHashicorpVaultSecretKeyOk() (*string, bool)`

GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretKey

`func (o *EnvironmentUserParams) SetHashicorpVaultSecretKey(v string)`

SetHashicorpVaultSecretKey sets HashicorpVaultSecretKey field to given value.

### HasHashicorpVaultSecretKey

`func (o *EnvironmentUserParams) HasHashicorpVaultSecretKey() bool`

HasHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetCyberarkVaultQueryString

`func (o *EnvironmentUserParams) GetCyberarkVaultQueryString() string`

GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetCyberarkVaultQueryStringOk

`func (o *EnvironmentUserParams) GetCyberarkVaultQueryStringOk() (*string, bool)`

GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyberarkVaultQueryString

`func (o *EnvironmentUserParams) SetCyberarkVaultQueryString(v string)`

SetCyberarkVaultQueryString sets CyberarkVaultQueryString field to given value.

### HasCyberarkVaultQueryString

`func (o *EnvironmentUserParams) HasCyberarkVaultQueryString() bool`

HasCyberarkVaultQueryString returns a boolean if a field has been set.

### GetUseKerberosAuthentication

`func (o *EnvironmentUserParams) GetUseKerberosAuthentication() bool`

GetUseKerberosAuthentication returns the UseKerberosAuthentication field if non-nil, zero value otherwise.

### GetUseKerberosAuthenticationOk

`func (o *EnvironmentUserParams) GetUseKerberosAuthenticationOk() (*bool, bool)`

GetUseKerberosAuthenticationOk returns a tuple with the UseKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKerberosAuthentication

`func (o *EnvironmentUserParams) SetUseKerberosAuthentication(v bool)`

SetUseKerberosAuthentication sets UseKerberosAuthentication field to given value.

### HasUseKerberosAuthentication

`func (o *EnvironmentUserParams) HasUseKerberosAuthentication() bool`

HasUseKerberosAuthentication returns a boolean if a field has been set.

### GetUseEnginePublicKey

`func (o *EnvironmentUserParams) GetUseEnginePublicKey() bool`

GetUseEnginePublicKey returns the UseEnginePublicKey field if non-nil, zero value otherwise.

### GetUseEnginePublicKeyOk

`func (o *EnvironmentUserParams) GetUseEnginePublicKeyOk() (*bool, bool)`

GetUseEnginePublicKeyOk returns a tuple with the UseEnginePublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnginePublicKey

`func (o *EnvironmentUserParams) SetUseEnginePublicKey(v bool)`

SetUseEnginePublicKey sets UseEnginePublicKey field to given value.

### HasUseEnginePublicKey

`func (o *EnvironmentUserParams) HasUseEnginePublicKey() bool`

HasUseEnginePublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


