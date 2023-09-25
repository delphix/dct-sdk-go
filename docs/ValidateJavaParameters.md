# ValidateJavaParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JavaHome** | **string** | Path pointing to java home on the remote machine. | 
**Port** | **NullableInt32** | SSH port of the remote host machine that will be used to establish SSH connection. | 
**Username** | Pointer to **string** | The username of the user that will be used to connect to the remote host machine. | [optional] 
**Password** | Pointer to **string** | The password of the user that will be used to connect to the remote host machine. | [optional] 
**HostName** | **string** | Hostname of the remote host machine that will be used to establish connection. | 
**UseEnginePublicKey** | Pointer to **bool** | Whether to use public key authentication. | [optional] 
**VaultId** | Pointer to **string** | The DCT id or name of the vault from which to read the host credentials. | [optional] 
**HashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored. | [optional] 
**HashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored. | [optional] 
**HashicorpVaultUsernameKey** | Pointer to **string** | Key for the username in the key-value store. | [optional] 
**HashicorpVaultSecretKey** | Pointer to **string** | Key for the password in the key-value store. | [optional] 
**AzureVaultName** | Pointer to **string** | Azure key vault name (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**AzureVaultUsernameKey** | Pointer to **string** | Azure vault key for the username in the key-value store (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**AzureVaultSecretKey** | Pointer to **string** | Azure vault key for the password in the key-value store (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**CyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault. | [optional] 
**UseKerberosAuthentication** | Pointer to **bool** | Whether to use kerberos authentication. | [optional] 

## Methods

### NewValidateJavaParameters

`func NewValidateJavaParameters(javaHome string, port NullableInt32, hostName string, ) *ValidateJavaParameters`

NewValidateJavaParameters instantiates a new ValidateJavaParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateJavaParametersWithDefaults

`func NewValidateJavaParametersWithDefaults() *ValidateJavaParameters`

NewValidateJavaParametersWithDefaults instantiates a new ValidateJavaParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJavaHome

`func (o *ValidateJavaParameters) GetJavaHome() string`

GetJavaHome returns the JavaHome field if non-nil, zero value otherwise.

### GetJavaHomeOk

`func (o *ValidateJavaParameters) GetJavaHomeOk() (*string, bool)`

GetJavaHomeOk returns a tuple with the JavaHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaHome

`func (o *ValidateJavaParameters) SetJavaHome(v string)`

SetJavaHome sets JavaHome field to given value.


### GetPort

`func (o *ValidateJavaParameters) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ValidateJavaParameters) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ValidateJavaParameters) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *ValidateJavaParameters) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ValidateJavaParameters) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *ValidateJavaParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ValidateJavaParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ValidateJavaParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ValidateJavaParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ValidateJavaParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ValidateJavaParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ValidateJavaParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ValidateJavaParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHostName

`func (o *ValidateJavaParameters) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ValidateJavaParameters) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ValidateJavaParameters) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetUseEnginePublicKey

`func (o *ValidateJavaParameters) GetUseEnginePublicKey() bool`

GetUseEnginePublicKey returns the UseEnginePublicKey field if non-nil, zero value otherwise.

### GetUseEnginePublicKeyOk

`func (o *ValidateJavaParameters) GetUseEnginePublicKeyOk() (*bool, bool)`

GetUseEnginePublicKeyOk returns a tuple with the UseEnginePublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnginePublicKey

`func (o *ValidateJavaParameters) SetUseEnginePublicKey(v bool)`

SetUseEnginePublicKey sets UseEnginePublicKey field to given value.

### HasUseEnginePublicKey

`func (o *ValidateJavaParameters) HasUseEnginePublicKey() bool`

HasUseEnginePublicKey returns a boolean if a field has been set.

### GetVaultId

`func (o *ValidateJavaParameters) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *ValidateJavaParameters) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *ValidateJavaParameters) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *ValidateJavaParameters) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### GetHashicorpVaultEngine

`func (o *ValidateJavaParameters) GetHashicorpVaultEngine() string`

GetHashicorpVaultEngine returns the HashicorpVaultEngine field if non-nil, zero value otherwise.

### GetHashicorpVaultEngineOk

`func (o *ValidateJavaParameters) GetHashicorpVaultEngineOk() (*string, bool)`

GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultEngine

`func (o *ValidateJavaParameters) SetHashicorpVaultEngine(v string)`

SetHashicorpVaultEngine sets HashicorpVaultEngine field to given value.

### HasHashicorpVaultEngine

`func (o *ValidateJavaParameters) HasHashicorpVaultEngine() bool`

HasHashicorpVaultEngine returns a boolean if a field has been set.

### GetHashicorpVaultSecretPath

`func (o *ValidateJavaParameters) GetHashicorpVaultSecretPath() string`

GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretPathOk

`func (o *ValidateJavaParameters) GetHashicorpVaultSecretPathOk() (*string, bool)`

GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretPath

`func (o *ValidateJavaParameters) SetHashicorpVaultSecretPath(v string)`

SetHashicorpVaultSecretPath sets HashicorpVaultSecretPath field to given value.

### HasHashicorpVaultSecretPath

`func (o *ValidateJavaParameters) HasHashicorpVaultSecretPath() bool`

HasHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetHashicorpVaultUsernameKey

`func (o *ValidateJavaParameters) GetHashicorpVaultUsernameKey() string`

GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameKeyOk

`func (o *ValidateJavaParameters) GetHashicorpVaultUsernameKeyOk() (*string, bool)`

GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameKey

`func (o *ValidateJavaParameters) SetHashicorpVaultUsernameKey(v string)`

SetHashicorpVaultUsernameKey sets HashicorpVaultUsernameKey field to given value.

### HasHashicorpVaultUsernameKey

`func (o *ValidateJavaParameters) HasHashicorpVaultUsernameKey() bool`

HasHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetHashicorpVaultSecretKey

`func (o *ValidateJavaParameters) GetHashicorpVaultSecretKey() string`

GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretKeyOk

`func (o *ValidateJavaParameters) GetHashicorpVaultSecretKeyOk() (*string, bool)`

GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretKey

`func (o *ValidateJavaParameters) SetHashicorpVaultSecretKey(v string)`

SetHashicorpVaultSecretKey sets HashicorpVaultSecretKey field to given value.

### HasHashicorpVaultSecretKey

`func (o *ValidateJavaParameters) HasHashicorpVaultSecretKey() bool`

HasHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetAzureVaultName

`func (o *ValidateJavaParameters) GetAzureVaultName() string`

GetAzureVaultName returns the AzureVaultName field if non-nil, zero value otherwise.

### GetAzureVaultNameOk

`func (o *ValidateJavaParameters) GetAzureVaultNameOk() (*string, bool)`

GetAzureVaultNameOk returns a tuple with the AzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultName

`func (o *ValidateJavaParameters) SetAzureVaultName(v string)`

SetAzureVaultName sets AzureVaultName field to given value.

### HasAzureVaultName

`func (o *ValidateJavaParameters) HasAzureVaultName() bool`

HasAzureVaultName returns a boolean if a field has been set.

### GetAzureVaultUsernameKey

`func (o *ValidateJavaParameters) GetAzureVaultUsernameKey() string`

GetAzureVaultUsernameKey returns the AzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetAzureVaultUsernameKeyOk

`func (o *ValidateJavaParameters) GetAzureVaultUsernameKeyOk() (*string, bool)`

GetAzureVaultUsernameKeyOk returns a tuple with the AzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultUsernameKey

`func (o *ValidateJavaParameters) SetAzureVaultUsernameKey(v string)`

SetAzureVaultUsernameKey sets AzureVaultUsernameKey field to given value.

### HasAzureVaultUsernameKey

`func (o *ValidateJavaParameters) HasAzureVaultUsernameKey() bool`

HasAzureVaultUsernameKey returns a boolean if a field has been set.

### GetAzureVaultSecretKey

`func (o *ValidateJavaParameters) GetAzureVaultSecretKey() string`

GetAzureVaultSecretKey returns the AzureVaultSecretKey field if non-nil, zero value otherwise.

### GetAzureVaultSecretKeyOk

`func (o *ValidateJavaParameters) GetAzureVaultSecretKeyOk() (*string, bool)`

GetAzureVaultSecretKeyOk returns a tuple with the AzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultSecretKey

`func (o *ValidateJavaParameters) SetAzureVaultSecretKey(v string)`

SetAzureVaultSecretKey sets AzureVaultSecretKey field to given value.

### HasAzureVaultSecretKey

`func (o *ValidateJavaParameters) HasAzureVaultSecretKey() bool`

HasAzureVaultSecretKey returns a boolean if a field has been set.

### GetCyberarkVaultQueryString

`func (o *ValidateJavaParameters) GetCyberarkVaultQueryString() string`

GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetCyberarkVaultQueryStringOk

`func (o *ValidateJavaParameters) GetCyberarkVaultQueryStringOk() (*string, bool)`

GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyberarkVaultQueryString

`func (o *ValidateJavaParameters) SetCyberarkVaultQueryString(v string)`

SetCyberarkVaultQueryString sets CyberarkVaultQueryString field to given value.

### HasCyberarkVaultQueryString

`func (o *ValidateJavaParameters) HasCyberarkVaultQueryString() bool`

HasCyberarkVaultQueryString returns a boolean if a field has been set.

### GetUseKerberosAuthentication

`func (o *ValidateJavaParameters) GetUseKerberosAuthentication() bool`

GetUseKerberosAuthentication returns the UseKerberosAuthentication field if non-nil, zero value otherwise.

### GetUseKerberosAuthenticationOk

`func (o *ValidateJavaParameters) GetUseKerberosAuthenticationOk() (*bool, bool)`

GetUseKerberosAuthenticationOk returns a tuple with the UseKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKerberosAuthentication

`func (o *ValidateJavaParameters) SetUseKerberosAuthentication(v bool)`

SetUseKerberosAuthentication sets UseKerberosAuthentication field to given value.

### HasUseKerberosAuthentication

`func (o *ValidateJavaParameters) HasUseKerberosAuthentication() bool`

HasUseKerberosAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


