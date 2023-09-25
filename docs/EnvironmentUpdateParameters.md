# EnvironmentUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the environment. | [optional] 
**StagingEnvironment** | Pointer to **string** | Id of the connector environment which is used to connect to this source environment. | [optional] 
**ClusterAddress** | Pointer to **string** | Address of the cluster. This property can be modified for Windows cluster only. | [optional] 
**ClusterHome** | Pointer to **string** | Absolute path to cluster home directory. This parameter is for UNIX cluster environments. | [optional] 
**AseDbUsername** | Pointer to **string** | username of the SAP ASE database. | [optional] 
**AseDbPassword** | Pointer to **string** | password of the SAP ASE database. | [optional] 
**AseDbVault** | Pointer to **string** | The name or reference of the vault from which to read the ASE database credentials. | [optional] 
**AseDbVaultUsername** | Pointer to **string** | Delphix display name for the vault user | [optional] 
**AseDbHashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored. | [optional] 
**AseDbHashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored. | [optional] 
**AseDbHashicorpVaultUsernameKey** | Pointer to **string** | Key for the username in the key-value store. | [optional] 
**AseDbHashicorpVaultSecretKey** | Pointer to **string** | Key for the password in the key-value store. | [optional] 
**AseDbCyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault. | [optional] 
**AseDbUseKerberosAuthentication** | Pointer to **bool** | Whether to use kerberos authentication for ASE DB discovery. | [optional] 
**Description** | Pointer to **string** | The environment description. | [optional] 

## Methods

### NewEnvironmentUpdateParameters

`func NewEnvironmentUpdateParameters() *EnvironmentUpdateParameters`

NewEnvironmentUpdateParameters instantiates a new EnvironmentUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentUpdateParametersWithDefaults

`func NewEnvironmentUpdateParametersWithDefaults() *EnvironmentUpdateParameters`

NewEnvironmentUpdateParametersWithDefaults instantiates a new EnvironmentUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentUpdateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentUpdateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentUpdateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentUpdateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStagingEnvironment

`func (o *EnvironmentUpdateParameters) GetStagingEnvironment() string`

GetStagingEnvironment returns the StagingEnvironment field if non-nil, zero value otherwise.

### GetStagingEnvironmentOk

`func (o *EnvironmentUpdateParameters) GetStagingEnvironmentOk() (*string, bool)`

GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironment

`func (o *EnvironmentUpdateParameters) SetStagingEnvironment(v string)`

SetStagingEnvironment sets StagingEnvironment field to given value.

### HasStagingEnvironment

`func (o *EnvironmentUpdateParameters) HasStagingEnvironment() bool`

HasStagingEnvironment returns a boolean if a field has been set.

### GetClusterAddress

`func (o *EnvironmentUpdateParameters) GetClusterAddress() string`

GetClusterAddress returns the ClusterAddress field if non-nil, zero value otherwise.

### GetClusterAddressOk

`func (o *EnvironmentUpdateParameters) GetClusterAddressOk() (*string, bool)`

GetClusterAddressOk returns a tuple with the ClusterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAddress

`func (o *EnvironmentUpdateParameters) SetClusterAddress(v string)`

SetClusterAddress sets ClusterAddress field to given value.

### HasClusterAddress

`func (o *EnvironmentUpdateParameters) HasClusterAddress() bool`

HasClusterAddress returns a boolean if a field has been set.

### GetClusterHome

`func (o *EnvironmentUpdateParameters) GetClusterHome() string`

GetClusterHome returns the ClusterHome field if non-nil, zero value otherwise.

### GetClusterHomeOk

`func (o *EnvironmentUpdateParameters) GetClusterHomeOk() (*string, bool)`

GetClusterHomeOk returns a tuple with the ClusterHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHome

`func (o *EnvironmentUpdateParameters) SetClusterHome(v string)`

SetClusterHome sets ClusterHome field to given value.

### HasClusterHome

`func (o *EnvironmentUpdateParameters) HasClusterHome() bool`

HasClusterHome returns a boolean if a field has been set.

### GetAseDbUsername

`func (o *EnvironmentUpdateParameters) GetAseDbUsername() string`

GetAseDbUsername returns the AseDbUsername field if non-nil, zero value otherwise.

### GetAseDbUsernameOk

`func (o *EnvironmentUpdateParameters) GetAseDbUsernameOk() (*string, bool)`

GetAseDbUsernameOk returns a tuple with the AseDbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbUsername

`func (o *EnvironmentUpdateParameters) SetAseDbUsername(v string)`

SetAseDbUsername sets AseDbUsername field to given value.

### HasAseDbUsername

`func (o *EnvironmentUpdateParameters) HasAseDbUsername() bool`

HasAseDbUsername returns a boolean if a field has been set.

### GetAseDbPassword

`func (o *EnvironmentUpdateParameters) GetAseDbPassword() string`

GetAseDbPassword returns the AseDbPassword field if non-nil, zero value otherwise.

### GetAseDbPasswordOk

`func (o *EnvironmentUpdateParameters) GetAseDbPasswordOk() (*string, bool)`

GetAseDbPasswordOk returns a tuple with the AseDbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbPassword

`func (o *EnvironmentUpdateParameters) SetAseDbPassword(v string)`

SetAseDbPassword sets AseDbPassword field to given value.

### HasAseDbPassword

`func (o *EnvironmentUpdateParameters) HasAseDbPassword() bool`

HasAseDbPassword returns a boolean if a field has been set.

### GetAseDbVault

`func (o *EnvironmentUpdateParameters) GetAseDbVault() string`

GetAseDbVault returns the AseDbVault field if non-nil, zero value otherwise.

### GetAseDbVaultOk

`func (o *EnvironmentUpdateParameters) GetAseDbVaultOk() (*string, bool)`

GetAseDbVaultOk returns a tuple with the AseDbVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbVault

`func (o *EnvironmentUpdateParameters) SetAseDbVault(v string)`

SetAseDbVault sets AseDbVault field to given value.

### HasAseDbVault

`func (o *EnvironmentUpdateParameters) HasAseDbVault() bool`

HasAseDbVault returns a boolean if a field has been set.

### GetAseDbVaultUsername

`func (o *EnvironmentUpdateParameters) GetAseDbVaultUsername() string`

GetAseDbVaultUsername returns the AseDbVaultUsername field if non-nil, zero value otherwise.

### GetAseDbVaultUsernameOk

`func (o *EnvironmentUpdateParameters) GetAseDbVaultUsernameOk() (*string, bool)`

GetAseDbVaultUsernameOk returns a tuple with the AseDbVaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbVaultUsername

`func (o *EnvironmentUpdateParameters) SetAseDbVaultUsername(v string)`

SetAseDbVaultUsername sets AseDbVaultUsername field to given value.

### HasAseDbVaultUsername

`func (o *EnvironmentUpdateParameters) HasAseDbVaultUsername() bool`

HasAseDbVaultUsername returns a boolean if a field has been set.

### GetAseDbHashicorpVaultEngine

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultEngine() string`

GetAseDbHashicorpVaultEngine returns the AseDbHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultEngineOk

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultEngineOk() (*string, bool)`

GetAseDbHashicorpVaultEngineOk returns a tuple with the AseDbHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultEngine

`func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultEngine(v string)`

SetAseDbHashicorpVaultEngine sets AseDbHashicorpVaultEngine field to given value.

### HasAseDbHashicorpVaultEngine

`func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultEngine() bool`

HasAseDbHashicorpVaultEngine returns a boolean if a field has been set.

### GetAseDbHashicorpVaultSecretPath

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretPath() string`

GetAseDbHashicorpVaultSecretPath returns the AseDbHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultSecretPathOk

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretPathOk() (*string, bool)`

GetAseDbHashicorpVaultSecretPathOk returns a tuple with the AseDbHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultSecretPath

`func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultSecretPath(v string)`

SetAseDbHashicorpVaultSecretPath sets AseDbHashicorpVaultSecretPath field to given value.

### HasAseDbHashicorpVaultSecretPath

`func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultSecretPath() bool`

HasAseDbHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetAseDbHashicorpVaultUsernameKey

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultUsernameKey() string`

GetAseDbHashicorpVaultUsernameKey returns the AseDbHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultUsernameKeyOk

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultUsernameKeyOk() (*string, bool)`

GetAseDbHashicorpVaultUsernameKeyOk returns a tuple with the AseDbHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultUsernameKey

`func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultUsernameKey(v string)`

SetAseDbHashicorpVaultUsernameKey sets AseDbHashicorpVaultUsernameKey field to given value.

### HasAseDbHashicorpVaultUsernameKey

`func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultUsernameKey() bool`

HasAseDbHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetAseDbHashicorpVaultSecretKey

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretKey() string`

GetAseDbHashicorpVaultSecretKey returns the AseDbHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultSecretKeyOk

`func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretKeyOk() (*string, bool)`

GetAseDbHashicorpVaultSecretKeyOk returns a tuple with the AseDbHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultSecretKey

`func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultSecretKey(v string)`

SetAseDbHashicorpVaultSecretKey sets AseDbHashicorpVaultSecretKey field to given value.

### HasAseDbHashicorpVaultSecretKey

`func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultSecretKey() bool`

HasAseDbHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetAseDbCyberarkVaultQueryString

`func (o *EnvironmentUpdateParameters) GetAseDbCyberarkVaultQueryString() string`

GetAseDbCyberarkVaultQueryString returns the AseDbCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetAseDbCyberarkVaultQueryStringOk

`func (o *EnvironmentUpdateParameters) GetAseDbCyberarkVaultQueryStringOk() (*string, bool)`

GetAseDbCyberarkVaultQueryStringOk returns a tuple with the AseDbCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbCyberarkVaultQueryString

`func (o *EnvironmentUpdateParameters) SetAseDbCyberarkVaultQueryString(v string)`

SetAseDbCyberarkVaultQueryString sets AseDbCyberarkVaultQueryString field to given value.

### HasAseDbCyberarkVaultQueryString

`func (o *EnvironmentUpdateParameters) HasAseDbCyberarkVaultQueryString() bool`

HasAseDbCyberarkVaultQueryString returns a boolean if a field has been set.

### GetAseDbUseKerberosAuthentication

`func (o *EnvironmentUpdateParameters) GetAseDbUseKerberosAuthentication() bool`

GetAseDbUseKerberosAuthentication returns the AseDbUseKerberosAuthentication field if non-nil, zero value otherwise.

### GetAseDbUseKerberosAuthenticationOk

`func (o *EnvironmentUpdateParameters) GetAseDbUseKerberosAuthenticationOk() (*bool, bool)`

GetAseDbUseKerberosAuthenticationOk returns a tuple with the AseDbUseKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbUseKerberosAuthentication

`func (o *EnvironmentUpdateParameters) SetAseDbUseKerberosAuthentication(v bool)`

SetAseDbUseKerberosAuthentication sets AseDbUseKerberosAuthentication field to given value.

### HasAseDbUseKerberosAuthentication

`func (o *EnvironmentUpdateParameters) HasAseDbUseKerberosAuthentication() bool`

HasAseDbUseKerberosAuthentication returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentUpdateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentUpdateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentUpdateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentUpdateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


