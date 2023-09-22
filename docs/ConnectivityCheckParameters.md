# ConnectivityCheckParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | **string** | The ID of the engine to check. | 
**UseEnginePublicKey** | Pointer to **bool** | Whether to use public key authentication. | [optional] 
**OsName** | Pointer to **string** | Operating system type of the environment. | [optional] 
**StagingEnvironment** | Pointer to **string** | Id of the connector environment which is used to connect to this source environment. | [optional] 
**Host** | **string** | The hostname of the remote host machine to check. | 
**Port** | **NullableInt32** | The port of the remote host machine to check. For Windows, port on which Delphix connector is running. | 
**Username** | Pointer to **string** | The username of the remote host machine to check. Username is mandatory input with password/use_engine_public_key/kerberos_authentication. | [optional] 
**Password** | Pointer to **string** | The password of the remote host machine to check. | [optional] 
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

### NewConnectivityCheckParameters

`func NewConnectivityCheckParameters(engineId string, host string, port NullableInt32, ) *ConnectivityCheckParameters`

NewConnectivityCheckParameters instantiates a new ConnectivityCheckParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityCheckParametersWithDefaults

`func NewConnectivityCheckParametersWithDefaults() *ConnectivityCheckParameters`

NewConnectivityCheckParametersWithDefaults instantiates a new ConnectivityCheckParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *ConnectivityCheckParameters) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ConnectivityCheckParameters) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ConnectivityCheckParameters) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetUseEnginePublicKey

`func (o *ConnectivityCheckParameters) GetUseEnginePublicKey() bool`

GetUseEnginePublicKey returns the UseEnginePublicKey field if non-nil, zero value otherwise.

### GetUseEnginePublicKeyOk

`func (o *ConnectivityCheckParameters) GetUseEnginePublicKeyOk() (*bool, bool)`

GetUseEnginePublicKeyOk returns a tuple with the UseEnginePublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnginePublicKey

`func (o *ConnectivityCheckParameters) SetUseEnginePublicKey(v bool)`

SetUseEnginePublicKey sets UseEnginePublicKey field to given value.

### HasUseEnginePublicKey

`func (o *ConnectivityCheckParameters) HasUseEnginePublicKey() bool`

HasUseEnginePublicKey returns a boolean if a field has been set.

### GetOsName

`func (o *ConnectivityCheckParameters) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *ConnectivityCheckParameters) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *ConnectivityCheckParameters) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *ConnectivityCheckParameters) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetStagingEnvironment

`func (o *ConnectivityCheckParameters) GetStagingEnvironment() string`

GetStagingEnvironment returns the StagingEnvironment field if non-nil, zero value otherwise.

### GetStagingEnvironmentOk

`func (o *ConnectivityCheckParameters) GetStagingEnvironmentOk() (*string, bool)`

GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironment

`func (o *ConnectivityCheckParameters) SetStagingEnvironment(v string)`

SetStagingEnvironment sets StagingEnvironment field to given value.

### HasStagingEnvironment

`func (o *ConnectivityCheckParameters) HasStagingEnvironment() bool`

HasStagingEnvironment returns a boolean if a field has been set.

### GetHost

`func (o *ConnectivityCheckParameters) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ConnectivityCheckParameters) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ConnectivityCheckParameters) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ConnectivityCheckParameters) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnectivityCheckParameters) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnectivityCheckParameters) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *ConnectivityCheckParameters) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ConnectivityCheckParameters) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *ConnectivityCheckParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConnectivityCheckParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConnectivityCheckParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConnectivityCheckParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ConnectivityCheckParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConnectivityCheckParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConnectivityCheckParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConnectivityCheckParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVaultId

`func (o *ConnectivityCheckParameters) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *ConnectivityCheckParameters) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *ConnectivityCheckParameters) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *ConnectivityCheckParameters) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### GetHashicorpVaultEngine

`func (o *ConnectivityCheckParameters) GetHashicorpVaultEngine() string`

GetHashicorpVaultEngine returns the HashicorpVaultEngine field if non-nil, zero value otherwise.

### GetHashicorpVaultEngineOk

`func (o *ConnectivityCheckParameters) GetHashicorpVaultEngineOk() (*string, bool)`

GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultEngine

`func (o *ConnectivityCheckParameters) SetHashicorpVaultEngine(v string)`

SetHashicorpVaultEngine sets HashicorpVaultEngine field to given value.

### HasHashicorpVaultEngine

`func (o *ConnectivityCheckParameters) HasHashicorpVaultEngine() bool`

HasHashicorpVaultEngine returns a boolean if a field has been set.

### GetHashicorpVaultSecretPath

`func (o *ConnectivityCheckParameters) GetHashicorpVaultSecretPath() string`

GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretPathOk

`func (o *ConnectivityCheckParameters) GetHashicorpVaultSecretPathOk() (*string, bool)`

GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretPath

`func (o *ConnectivityCheckParameters) SetHashicorpVaultSecretPath(v string)`

SetHashicorpVaultSecretPath sets HashicorpVaultSecretPath field to given value.

### HasHashicorpVaultSecretPath

`func (o *ConnectivityCheckParameters) HasHashicorpVaultSecretPath() bool`

HasHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetHashicorpVaultUsernameKey

`func (o *ConnectivityCheckParameters) GetHashicorpVaultUsernameKey() string`

GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameKeyOk

`func (o *ConnectivityCheckParameters) GetHashicorpVaultUsernameKeyOk() (*string, bool)`

GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameKey

`func (o *ConnectivityCheckParameters) SetHashicorpVaultUsernameKey(v string)`

SetHashicorpVaultUsernameKey sets HashicorpVaultUsernameKey field to given value.

### HasHashicorpVaultUsernameKey

`func (o *ConnectivityCheckParameters) HasHashicorpVaultUsernameKey() bool`

HasHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetHashicorpVaultSecretKey

`func (o *ConnectivityCheckParameters) GetHashicorpVaultSecretKey() string`

GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretKeyOk

`func (o *ConnectivityCheckParameters) GetHashicorpVaultSecretKeyOk() (*string, bool)`

GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretKey

`func (o *ConnectivityCheckParameters) SetHashicorpVaultSecretKey(v string)`

SetHashicorpVaultSecretKey sets HashicorpVaultSecretKey field to given value.

### HasHashicorpVaultSecretKey

`func (o *ConnectivityCheckParameters) HasHashicorpVaultSecretKey() bool`

HasHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetAzureVaultName

`func (o *ConnectivityCheckParameters) GetAzureVaultName() string`

GetAzureVaultName returns the AzureVaultName field if non-nil, zero value otherwise.

### GetAzureVaultNameOk

`func (o *ConnectivityCheckParameters) GetAzureVaultNameOk() (*string, bool)`

GetAzureVaultNameOk returns a tuple with the AzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultName

`func (o *ConnectivityCheckParameters) SetAzureVaultName(v string)`

SetAzureVaultName sets AzureVaultName field to given value.

### HasAzureVaultName

`func (o *ConnectivityCheckParameters) HasAzureVaultName() bool`

HasAzureVaultName returns a boolean if a field has been set.

### GetAzureVaultUsernameKey

`func (o *ConnectivityCheckParameters) GetAzureVaultUsernameKey() string`

GetAzureVaultUsernameKey returns the AzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetAzureVaultUsernameKeyOk

`func (o *ConnectivityCheckParameters) GetAzureVaultUsernameKeyOk() (*string, bool)`

GetAzureVaultUsernameKeyOk returns a tuple with the AzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultUsernameKey

`func (o *ConnectivityCheckParameters) SetAzureVaultUsernameKey(v string)`

SetAzureVaultUsernameKey sets AzureVaultUsernameKey field to given value.

### HasAzureVaultUsernameKey

`func (o *ConnectivityCheckParameters) HasAzureVaultUsernameKey() bool`

HasAzureVaultUsernameKey returns a boolean if a field has been set.

### GetAzureVaultSecretKey

`func (o *ConnectivityCheckParameters) GetAzureVaultSecretKey() string`

GetAzureVaultSecretKey returns the AzureVaultSecretKey field if non-nil, zero value otherwise.

### GetAzureVaultSecretKeyOk

`func (o *ConnectivityCheckParameters) GetAzureVaultSecretKeyOk() (*string, bool)`

GetAzureVaultSecretKeyOk returns a tuple with the AzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultSecretKey

`func (o *ConnectivityCheckParameters) SetAzureVaultSecretKey(v string)`

SetAzureVaultSecretKey sets AzureVaultSecretKey field to given value.

### HasAzureVaultSecretKey

`func (o *ConnectivityCheckParameters) HasAzureVaultSecretKey() bool`

HasAzureVaultSecretKey returns a boolean if a field has been set.

### GetCyberarkVaultQueryString

`func (o *ConnectivityCheckParameters) GetCyberarkVaultQueryString() string`

GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetCyberarkVaultQueryStringOk

`func (o *ConnectivityCheckParameters) GetCyberarkVaultQueryStringOk() (*string, bool)`

GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyberarkVaultQueryString

`func (o *ConnectivityCheckParameters) SetCyberarkVaultQueryString(v string)`

SetCyberarkVaultQueryString sets CyberarkVaultQueryString field to given value.

### HasCyberarkVaultQueryString

`func (o *ConnectivityCheckParameters) HasCyberarkVaultQueryString() bool`

HasCyberarkVaultQueryString returns a boolean if a field has been set.

### GetUseKerberosAuthentication

`func (o *ConnectivityCheckParameters) GetUseKerberosAuthentication() bool`

GetUseKerberosAuthentication returns the UseKerberosAuthentication field if non-nil, zero value otherwise.

### GetUseKerberosAuthenticationOk

`func (o *ConnectivityCheckParameters) GetUseKerberosAuthenticationOk() (*bool, bool)`

GetUseKerberosAuthenticationOk returns a tuple with the UseKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKerberosAuthentication

`func (o *ConnectivityCheckParameters) SetUseKerberosAuthentication(v bool)`

SetUseKerberosAuthentication sets UseKerberosAuthentication field to given value.

### HasUseKerberosAuthentication

`func (o *ConnectivityCheckParameters) HasUseKerberosAuthentication() bool`

HasUseKerberosAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


