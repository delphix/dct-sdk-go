# DatabaseConnectivityCheckParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsType** | **string** | The type of credentials. | 
**SourceId** | **string** | Source database config Id. | 
**Username** | Pointer to **string** | Database username (Not applicable for MSSQL_ENVIRONMENT_USER). | [optional] 
**Password** | Pointer to **string** | Database password (Not applicable for MSSQL_ENVIRONMENT_USER and mutually exclusive with vault attributes). | [optional] 
**Vault** | Pointer to **string** | The name or reference of the vault from which to read the database credentials (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**HashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**HashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**HashicorpVaultUsernameKey** | Pointer to **string** | Hashicorp vault key for the username in the key-value store (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**HashicorpVaultSecretKey** | Pointer to **string** | Hashicorp vault key for the password in the key-value store (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**AzureVaultName** | Pointer to **string** | Azure key vault name (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**AzureVaultUsernameKey** | Pointer to **string** | Azure vault key for the username in the key-value store (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**AzureVaultSecretKey** | Pointer to **string** | Azure vault key for the password in the key-value store (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**CyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault (ORACLE, ASE and MSSQL_DOMAIN_USER only). | [optional] 
**EnvironmentId** | Pointer to **string** | Id of the environment to which environment user belongs (MSSQL_ENVIRONMENT_USER only). | [optional] 
**EnvironmentUser** | Pointer to **string** | Reference to the environment user (MSSQL_ENVIRONMENT_USER only). | [optional] 

## Methods

### NewDatabaseConnectivityCheckParameters

`func NewDatabaseConnectivityCheckParameters(credentialsType string, sourceId string, ) *DatabaseConnectivityCheckParameters`

NewDatabaseConnectivityCheckParameters instantiates a new DatabaseConnectivityCheckParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConnectivityCheckParametersWithDefaults

`func NewDatabaseConnectivityCheckParametersWithDefaults() *DatabaseConnectivityCheckParameters`

NewDatabaseConnectivityCheckParametersWithDefaults instantiates a new DatabaseConnectivityCheckParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsType

`func (o *DatabaseConnectivityCheckParameters) GetCredentialsType() string`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *DatabaseConnectivityCheckParameters) GetCredentialsTypeOk() (*string, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *DatabaseConnectivityCheckParameters) SetCredentialsType(v string)`

SetCredentialsType sets CredentialsType field to given value.


### GetSourceId

`func (o *DatabaseConnectivityCheckParameters) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DatabaseConnectivityCheckParameters) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DatabaseConnectivityCheckParameters) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetUsername

`func (o *DatabaseConnectivityCheckParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DatabaseConnectivityCheckParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DatabaseConnectivityCheckParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DatabaseConnectivityCheckParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *DatabaseConnectivityCheckParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DatabaseConnectivityCheckParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DatabaseConnectivityCheckParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DatabaseConnectivityCheckParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVault

`func (o *DatabaseConnectivityCheckParameters) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *DatabaseConnectivityCheckParameters) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *DatabaseConnectivityCheckParameters) SetVault(v string)`

SetVault sets Vault field to given value.

### HasVault

`func (o *DatabaseConnectivityCheckParameters) HasVault() bool`

HasVault returns a boolean if a field has been set.

### GetHashicorpVaultEngine

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultEngine() string`

GetHashicorpVaultEngine returns the HashicorpVaultEngine field if non-nil, zero value otherwise.

### GetHashicorpVaultEngineOk

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultEngineOk() (*string, bool)`

GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultEngine

`func (o *DatabaseConnectivityCheckParameters) SetHashicorpVaultEngine(v string)`

SetHashicorpVaultEngine sets HashicorpVaultEngine field to given value.

### HasHashicorpVaultEngine

`func (o *DatabaseConnectivityCheckParameters) HasHashicorpVaultEngine() bool`

HasHashicorpVaultEngine returns a boolean if a field has been set.

### GetHashicorpVaultSecretPath

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultSecretPath() string`

GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretPathOk

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultSecretPathOk() (*string, bool)`

GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretPath

`func (o *DatabaseConnectivityCheckParameters) SetHashicorpVaultSecretPath(v string)`

SetHashicorpVaultSecretPath sets HashicorpVaultSecretPath field to given value.

### HasHashicorpVaultSecretPath

`func (o *DatabaseConnectivityCheckParameters) HasHashicorpVaultSecretPath() bool`

HasHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetHashicorpVaultUsernameKey

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultUsernameKey() string`

GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameKeyOk

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultUsernameKeyOk() (*string, bool)`

GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameKey

`func (o *DatabaseConnectivityCheckParameters) SetHashicorpVaultUsernameKey(v string)`

SetHashicorpVaultUsernameKey sets HashicorpVaultUsernameKey field to given value.

### HasHashicorpVaultUsernameKey

`func (o *DatabaseConnectivityCheckParameters) HasHashicorpVaultUsernameKey() bool`

HasHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetHashicorpVaultSecretKey

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultSecretKey() string`

GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretKeyOk

`func (o *DatabaseConnectivityCheckParameters) GetHashicorpVaultSecretKeyOk() (*string, bool)`

GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretKey

`func (o *DatabaseConnectivityCheckParameters) SetHashicorpVaultSecretKey(v string)`

SetHashicorpVaultSecretKey sets HashicorpVaultSecretKey field to given value.

### HasHashicorpVaultSecretKey

`func (o *DatabaseConnectivityCheckParameters) HasHashicorpVaultSecretKey() bool`

HasHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetAzureVaultName

`func (o *DatabaseConnectivityCheckParameters) GetAzureVaultName() string`

GetAzureVaultName returns the AzureVaultName field if non-nil, zero value otherwise.

### GetAzureVaultNameOk

`func (o *DatabaseConnectivityCheckParameters) GetAzureVaultNameOk() (*string, bool)`

GetAzureVaultNameOk returns a tuple with the AzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultName

`func (o *DatabaseConnectivityCheckParameters) SetAzureVaultName(v string)`

SetAzureVaultName sets AzureVaultName field to given value.

### HasAzureVaultName

`func (o *DatabaseConnectivityCheckParameters) HasAzureVaultName() bool`

HasAzureVaultName returns a boolean if a field has been set.

### GetAzureVaultUsernameKey

`func (o *DatabaseConnectivityCheckParameters) GetAzureVaultUsernameKey() string`

GetAzureVaultUsernameKey returns the AzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetAzureVaultUsernameKeyOk

`func (o *DatabaseConnectivityCheckParameters) GetAzureVaultUsernameKeyOk() (*string, bool)`

GetAzureVaultUsernameKeyOk returns a tuple with the AzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultUsernameKey

`func (o *DatabaseConnectivityCheckParameters) SetAzureVaultUsernameKey(v string)`

SetAzureVaultUsernameKey sets AzureVaultUsernameKey field to given value.

### HasAzureVaultUsernameKey

`func (o *DatabaseConnectivityCheckParameters) HasAzureVaultUsernameKey() bool`

HasAzureVaultUsernameKey returns a boolean if a field has been set.

### GetAzureVaultSecretKey

`func (o *DatabaseConnectivityCheckParameters) GetAzureVaultSecretKey() string`

GetAzureVaultSecretKey returns the AzureVaultSecretKey field if non-nil, zero value otherwise.

### GetAzureVaultSecretKeyOk

`func (o *DatabaseConnectivityCheckParameters) GetAzureVaultSecretKeyOk() (*string, bool)`

GetAzureVaultSecretKeyOk returns a tuple with the AzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVaultSecretKey

`func (o *DatabaseConnectivityCheckParameters) SetAzureVaultSecretKey(v string)`

SetAzureVaultSecretKey sets AzureVaultSecretKey field to given value.

### HasAzureVaultSecretKey

`func (o *DatabaseConnectivityCheckParameters) HasAzureVaultSecretKey() bool`

HasAzureVaultSecretKey returns a boolean if a field has been set.

### GetCyberarkVaultQueryString

`func (o *DatabaseConnectivityCheckParameters) GetCyberarkVaultQueryString() string`

GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetCyberarkVaultQueryStringOk

`func (o *DatabaseConnectivityCheckParameters) GetCyberarkVaultQueryStringOk() (*string, bool)`

GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyberarkVaultQueryString

`func (o *DatabaseConnectivityCheckParameters) SetCyberarkVaultQueryString(v string)`

SetCyberarkVaultQueryString sets CyberarkVaultQueryString field to given value.

### HasCyberarkVaultQueryString

`func (o *DatabaseConnectivityCheckParameters) HasCyberarkVaultQueryString() bool`

HasCyberarkVaultQueryString returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DatabaseConnectivityCheckParameters) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DatabaseConnectivityCheckParameters) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DatabaseConnectivityCheckParameters) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DatabaseConnectivityCheckParameters) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentUser

`func (o *DatabaseConnectivityCheckParameters) GetEnvironmentUser() string`

GetEnvironmentUser returns the EnvironmentUser field if non-nil, zero value otherwise.

### GetEnvironmentUserOk

`func (o *DatabaseConnectivityCheckParameters) GetEnvironmentUserOk() (*string, bool)`

GetEnvironmentUserOk returns a tuple with the EnvironmentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUser

`func (o *DatabaseConnectivityCheckParameters) SetEnvironmentUser(v string)`

SetEnvironmentUser sets EnvironmentUser field to given value.

### HasEnvironmentUser

`func (o *DatabaseConnectivityCheckParameters) HasEnvironmentUser() bool`

HasEnvironmentUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


