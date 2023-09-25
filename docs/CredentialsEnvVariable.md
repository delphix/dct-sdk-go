# CredentialsEnvVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseVarName** | **string** | Base name of the environment variables. Variables are named by appending &#39;_USER&#39;, &#39;_PASSWORD&#39;, &#39;_PUBKEY&#39; and &#39;_PRIVKEY&#39; to this base name, respectively. Variables whose values are not entered or are not present in the type of credential or vault selected, will not be set. | 
**Password** | Pointer to **string** | Password to assign to the environment variables. | [optional] 
**Vault** | Pointer to **string** | The name or reference of the vault to assign to the environment variables. | [optional] 
**HashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored. | [optional] 
**HashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored. | [optional] 
**HashicorpVaultUsernameKey** | Pointer to **string** | Hashicorp vault key for the username in the key-value store. | [optional] 
**HashicorpVaultSecretKey** | Pointer to **string** | Hashicorp vault key for the password in the key-value store. | [optional] 
**AzureVaultName** | Pointer to **string** | Azure key vault name. | [optional] 
**AzureVaultUsernameKey** | Pointer to **string** | Azure vault key in the key-value store. | [optional] 
**AzureVaultSecretKey** | Pointer to **string** | Azure vault key in the key-value store. | [optional] 
**CyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault. | [optional] 

## Methods

### NewCredentialsEnvVariable

`func NewCredentialsEnvVariable(baseVarName string, ) *CredentialsEnvVariable`

NewCredentialsEnvVariable instantiates a new CredentialsEnvVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsEnvVariableWithDefaults

`func NewCredentialsEnvVariableWithDefaults() *CredentialsEnvVariable`

NewCredentialsEnvVariableWithDefaults instantiates a new CredentialsEnvVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseVarName

`func (o *CredentialsEnvVariable) GetBaseVarName() string`

GetBaseVarName returns the BaseVarName field if non-nil, zero value otherwise.

### GetBaseVarNameOk

`func (o *CredentialsEnvVariable) GetBaseVarNameOk() (*string, bool)`

GetBaseVarNameOk returns a tuple with the BaseVarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVarName

`func (o *CredentialsEnvVariable) SetBaseVarName(v string)`

SetBaseVarName sets BaseVarName field to given value.


### GetPassword

`func (o *CredentialsEnvVariable) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialsEnvVariable) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialsEnvVariable) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CredentialsEnvVariable) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVault

`func (o *CredentialsEnvVariable) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *CredentialsEnvVariable) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *CredentialsEnvVariable) SetVault(v string)`

SetVault sets Vault field to given value.

### HasVault

`func (o *CredentialsEnvVariable) HasVault() bool`

HasVault returns a boolean if a field has been set.

### GetHashicorpVaultEngine

`func (o *CredentialsEnvVariable) GetHashicorpVaultEngine() string`

GetHashicorpVaultEngine returns the HashicorpVaultEngine field if non-nil, zero value otherwise.

### GetHashicorpVaultEngineOk

`func (o *CredentialsEnvVariable) GetHashicorpVaultEngineOk() (*string, bool)`

GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultEngine

`func (o *CredentialsEnvVariable) SetHashicorpVaultEngine(v string)`

SetHashicorpVaultEngine sets HashicorpVaultEngine field to given value.

### HasHashicorpVaultEngine

`func (o *CredentialsEnvVariable) HasHashicorpVaultEngine() bool`

HasHashicorpVaultEngine returns a boolean if a field has been set.

### GetHashicorpVaultSecretPath

`func (o *CredentialsEnvVariable) GetHashicorpVaultSecretPath() string`

GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretPathOk

`func (o *CredentialsEnvVariable) GetHashicorpVaultSecretPathOk() (*string, bool)`

GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretPath

`func (o *CredentialsEnvVariable) SetHashicorpVaultSecretPath(v string)`

SetHashicorpVaultSecretPath sets HashicorpVaultSecretPath field to given value.

### HasHashicorpVaultSecretPath

`func (o *CredentialsEnvVariable) HasHashicorpVaultSecretPath() bool`

HasHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetHashicorpVaultUsernameKey

`func (o *CredentialsEnvVariable) GetHashicorpVaultUsernameKey() string`

GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameKeyOk

`func (o *CredentialsEnvVariable) GetHashicorpVaultUsernameKeyOk() (*string, bool)`

GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameKey

`func (o *CredentialsEnvVariable) SetHashicorpVaultUsernameKey(v string)`

SetHashicorpVaultUsernameKey sets HashicorpVaultUsernameKey field to given value.

### HasHashicorpVaultUsernameKey

`func (o *CredentialsEnvVariable) HasHashicorpVaultUsernameKey() bool`

HasHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetHashicorpVaultSecretKey

`func (o *CredentialsEnvVariable) GetHashicorpVaultSecretKey() string`

GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretKeyOk

`func (o *CredentialsEnvVariable) GetHashicorpVaultSecretKeyOk() (*string, bool)`

GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretKey

`func (o *CredentialsEnvVariable) SetHashicorpVaultSecretKey(v string)`

SetHashicorpVaultSecretKey sets HashicorpVaultSecretKey field to given value.

### HasHashicorpVaultSecretKey

`func (o *CredentialsEnvVariable) HasHashicorpVaultSecretKey() bool`

HasHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetAzureVaultName

`func (o *CredentialsEnvVariable) GetAzureVaultName() string`

GetAzureVaultName returns the AzureVaultName field if non-nil, zero value otherwise.

### GetAzureVaultNameOk

`func (o *CredentialsEnvVariable) GetAzureVaultNameOk() (*string, bool)`

GetAzureVaultNameOk returns a tuple with the AzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultName

`func (o *CredentialsEnvVariable) SetAzureVaultName(v string)`

SetAzureVaultName sets AzureVaultName field to given value.

### HasAzureVaultName

`func (o *CredentialsEnvVariable) HasAzureVaultName() bool`

HasAzureVaultName returns a boolean if a field has been set.

### GetAzureVaultUsernameKey

`func (o *CredentialsEnvVariable) GetAzureVaultUsernameKey() string`

GetAzureVaultUsernameKey returns the AzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetAzureVaultUsernameKeyOk

`func (o *CredentialsEnvVariable) GetAzureVaultUsernameKeyOk() (*string, bool)`

GetAzureVaultUsernameKeyOk returns a tuple with the AzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultUsernameKey

`func (o *CredentialsEnvVariable) SetAzureVaultUsernameKey(v string)`

SetAzureVaultUsernameKey sets AzureVaultUsernameKey field to given value.

### HasAzureVaultUsernameKey

`func (o *CredentialsEnvVariable) HasAzureVaultUsernameKey() bool`

HasAzureVaultUsernameKey returns a boolean if a field has been set.

### GetAzureVaultSecretKey

`func (o *CredentialsEnvVariable) GetAzureVaultSecretKey() string`

GetAzureVaultSecretKey returns the AzureVaultSecretKey field if non-nil, zero value otherwise.

### GetAzureVaultSecretKeyOk

`func (o *CredentialsEnvVariable) GetAzureVaultSecretKeyOk() (*string, bool)`

GetAzureVaultSecretKeyOk returns a tuple with the AzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultSecretKey

`func (o *CredentialsEnvVariable) SetAzureVaultSecretKey(v string)`

SetAzureVaultSecretKey sets AzureVaultSecretKey field to given value.

### HasAzureVaultSecretKey

`func (o *CredentialsEnvVariable) HasAzureVaultSecretKey() bool`

HasAzureVaultSecretKey returns a boolean if a field has been set.

### GetCyberarkVaultQueryString

`func (o *CredentialsEnvVariable) GetCyberarkVaultQueryString() string`

GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetCyberarkVaultQueryStringOk

`func (o *CredentialsEnvVariable) GetCyberarkVaultQueryStringOk() (*string, bool)`

GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyberarkVaultQueryString

`func (o *CredentialsEnvVariable) SetCyberarkVaultQueryString(v string)`

SetCyberarkVaultQueryString sets CyberarkVaultQueryString field to given value.

### HasCyberarkVaultQueryString

`func (o *CredentialsEnvVariable) HasCyberarkVaultQueryString() bool`

HasCyberarkVaultQueryString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


