# EnvironmentCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the environment. | [optional] 
**EngineId** | **string** | The ID of the Engine onto which to create the environment. | 
**OsName** | **string** | Operating system type of the environment. | 
**IsCluster** | Pointer to **bool** | Whether the environment to be created is a cluster. | [optional] [default to false]
**ClusterHome** | Pointer to **string** | Absolute path to cluster home drectory. This parameter is mandatory for UNIX cluster environments. | [optional] 
**Hostname** | **string** | host address of the machine. | 
**StagingEnvironment** | Pointer to **string** | Id of the connector environment which is used to connect to this source environment. This is mandatory parameter when creating Windows source environments. | [optional] 
**ConnectorPort** | Pointer to **int32** | Specify port on which Delphix connector will run. This is mandatory parameter when creating Windows target environments. | [optional] 
**ConnectorAuthenticationKey** | Pointer to **string** | Unique per Delphix key used to authenticate with the remote Delphix Connector. | [optional] 
**IsTarget** | Pointer to **bool** | Whether the environment to be created is a target cluster environment. This property is used only when creating Windows cluster environments. | [optional] 
**SshPort** | Pointer to **int64** | ssh port of the host. | [optional] [default to 22]
**ToolkitPath** | Pointer to **string** | The path for the toolkit that resides on the host. | [optional] 
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
**NfsAddresses** | Pointer to **[]string** | array of ip address or hostnames | [optional] 
**AseDbVaultUsername** | Pointer to **string** | Delphix display name for the vault user | [optional] 
**AseDbUsername** | Pointer to **string** | username of the SAP ASE database. | [optional] 
**AseDbPassword** | Pointer to **string** | password of the SAP ASE database. | [optional] 
**AseDbVault** | Pointer to **string** | The name or reference of the vault from which to read the ASE database credentials. | [optional] 
**AseDbHashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored. | [optional] 
**AseDbHashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored. | [optional] 
**AseDbHashicorpVaultUsernameKey** | Pointer to **string** | Key for the username in the key-value store. | [optional] 
**AseDbHashicorpVaultSecretKey** | Pointer to **string** | Key for the password in the key-value store. | [optional] 
**AseDbCyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault. | [optional] 
**AseDbUseKerberosAuthentication** | Pointer to **bool** | Whether to use kerberos authentication for ASE DB discovery. | [optional] 
**JavaHome** | Pointer to **string** | The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used. | [optional] 
**DspKeystorePath** | Pointer to **string** | DSP keystore path. | [optional] 
**DspKeystorePassword** | Pointer to **string** | DSP keystore password. | [optional] 
**DspKeystoreAlias** | Pointer to **string** | DSP keystore alias. | [optional] 
**DspTruststorePath** | Pointer to **string** | DSP truststore path. | [optional] 
**DspTruststorePassword** | Pointer to **string** | DSP truststore password. | [optional] 
**Description** | Pointer to **string** | The environment description. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this environment. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this environment must be configured as owner of the environment. | [optional] [default to true]

## Methods

### NewEnvironmentCreateParameters

`func NewEnvironmentCreateParameters(engineId string, osName string, hostname string, ) *EnvironmentCreateParameters`

NewEnvironmentCreateParameters instantiates a new EnvironmentCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCreateParametersWithDefaults

`func NewEnvironmentCreateParametersWithDefaults() *EnvironmentCreateParameters`

NewEnvironmentCreateParametersWithDefaults instantiates a new EnvironmentCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentCreateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCreateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCreateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentCreateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEngineId

`func (o *EnvironmentCreateParameters) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *EnvironmentCreateParameters) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *EnvironmentCreateParameters) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetOsName

`func (o *EnvironmentCreateParameters) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *EnvironmentCreateParameters) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *EnvironmentCreateParameters) SetOsName(v string)`

SetOsName sets OsName field to given value.


### GetIsCluster

`func (o *EnvironmentCreateParameters) GetIsCluster() bool`

GetIsCluster returns the IsCluster field if non-nil, zero value otherwise.

### GetIsClusterOk

`func (o *EnvironmentCreateParameters) GetIsClusterOk() (*bool, bool)`

GetIsClusterOk returns a tuple with the IsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCluster

`func (o *EnvironmentCreateParameters) SetIsCluster(v bool)`

SetIsCluster sets IsCluster field to given value.

### HasIsCluster

`func (o *EnvironmentCreateParameters) HasIsCluster() bool`

HasIsCluster returns a boolean if a field has been set.

### GetClusterHome

`func (o *EnvironmentCreateParameters) GetClusterHome() string`

GetClusterHome returns the ClusterHome field if non-nil, zero value otherwise.

### GetClusterHomeOk

`func (o *EnvironmentCreateParameters) GetClusterHomeOk() (*string, bool)`

GetClusterHomeOk returns a tuple with the ClusterHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHome

`func (o *EnvironmentCreateParameters) SetClusterHome(v string)`

SetClusterHome sets ClusterHome field to given value.

### HasClusterHome

`func (o *EnvironmentCreateParameters) HasClusterHome() bool`

HasClusterHome returns a boolean if a field has been set.

### GetHostname

`func (o *EnvironmentCreateParameters) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EnvironmentCreateParameters) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EnvironmentCreateParameters) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStagingEnvironment

`func (o *EnvironmentCreateParameters) GetStagingEnvironment() string`

GetStagingEnvironment returns the StagingEnvironment field if non-nil, zero value otherwise.

### GetStagingEnvironmentOk

`func (o *EnvironmentCreateParameters) GetStagingEnvironmentOk() (*string, bool)`

GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironment

`func (o *EnvironmentCreateParameters) SetStagingEnvironment(v string)`

SetStagingEnvironment sets StagingEnvironment field to given value.

### HasStagingEnvironment

`func (o *EnvironmentCreateParameters) HasStagingEnvironment() bool`

HasStagingEnvironment returns a boolean if a field has been set.

### GetConnectorPort

`func (o *EnvironmentCreateParameters) GetConnectorPort() int32`

GetConnectorPort returns the ConnectorPort field if non-nil, zero value otherwise.

### GetConnectorPortOk

`func (o *EnvironmentCreateParameters) GetConnectorPortOk() (*int32, bool)`

GetConnectorPortOk returns a tuple with the ConnectorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPort

`func (o *EnvironmentCreateParameters) SetConnectorPort(v int32)`

SetConnectorPort sets ConnectorPort field to given value.

### HasConnectorPort

`func (o *EnvironmentCreateParameters) HasConnectorPort() bool`

HasConnectorPort returns a boolean if a field has been set.

### GetConnectorAuthenticationKey

`func (o *EnvironmentCreateParameters) GetConnectorAuthenticationKey() string`

GetConnectorAuthenticationKey returns the ConnectorAuthenticationKey field if non-nil, zero value otherwise.

### GetConnectorAuthenticationKeyOk

`func (o *EnvironmentCreateParameters) GetConnectorAuthenticationKeyOk() (*string, bool)`

GetConnectorAuthenticationKeyOk returns a tuple with the ConnectorAuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorAuthenticationKey

`func (o *EnvironmentCreateParameters) SetConnectorAuthenticationKey(v string)`

SetConnectorAuthenticationKey sets ConnectorAuthenticationKey field to given value.

### HasConnectorAuthenticationKey

`func (o *EnvironmentCreateParameters) HasConnectorAuthenticationKey() bool`

HasConnectorAuthenticationKey returns a boolean if a field has been set.

### GetIsTarget

`func (o *EnvironmentCreateParameters) GetIsTarget() bool`

GetIsTarget returns the IsTarget field if non-nil, zero value otherwise.

### GetIsTargetOk

`func (o *EnvironmentCreateParameters) GetIsTargetOk() (*bool, bool)`

GetIsTargetOk returns a tuple with the IsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTarget

`func (o *EnvironmentCreateParameters) SetIsTarget(v bool)`

SetIsTarget sets IsTarget field to given value.

### HasIsTarget

`func (o *EnvironmentCreateParameters) HasIsTarget() bool`

HasIsTarget returns a boolean if a field has been set.

### GetSshPort

`func (o *EnvironmentCreateParameters) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *EnvironmentCreateParameters) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *EnvironmentCreateParameters) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *EnvironmentCreateParameters) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetToolkitPath

`func (o *EnvironmentCreateParameters) GetToolkitPath() string`

GetToolkitPath returns the ToolkitPath field if non-nil, zero value otherwise.

### GetToolkitPathOk

`func (o *EnvironmentCreateParameters) GetToolkitPathOk() (*string, bool)`

GetToolkitPathOk returns a tuple with the ToolkitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolkitPath

`func (o *EnvironmentCreateParameters) SetToolkitPath(v string)`

SetToolkitPath sets ToolkitPath field to given value.

### HasToolkitPath

`func (o *EnvironmentCreateParameters) HasToolkitPath() bool`

HasToolkitPath returns a boolean if a field has been set.

### GetUsername

`func (o *EnvironmentCreateParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnvironmentCreateParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnvironmentCreateParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EnvironmentCreateParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *EnvironmentCreateParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnvironmentCreateParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnvironmentCreateParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnvironmentCreateParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVault

`func (o *EnvironmentCreateParameters) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *EnvironmentCreateParameters) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *EnvironmentCreateParameters) SetVault(v string)`

SetVault sets Vault field to given value.

### HasVault

`func (o *EnvironmentCreateParameters) HasVault() bool`

HasVault returns a boolean if a field has been set.

### GetVaultUsername

`func (o *EnvironmentCreateParameters) GetVaultUsername() string`

GetVaultUsername returns the VaultUsername field if non-nil, zero value otherwise.

### GetVaultUsernameOk

`func (o *EnvironmentCreateParameters) GetVaultUsernameOk() (*string, bool)`

GetVaultUsernameOk returns a tuple with the VaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultUsername

`func (o *EnvironmentCreateParameters) SetVaultUsername(v string)`

SetVaultUsername sets VaultUsername field to given value.

### HasVaultUsername

`func (o *EnvironmentCreateParameters) HasVaultUsername() bool`

HasVaultUsername returns a boolean if a field has been set.

### GetHashicorpVaultEngine

`func (o *EnvironmentCreateParameters) GetHashicorpVaultEngine() string`

GetHashicorpVaultEngine returns the HashicorpVaultEngine field if non-nil, zero value otherwise.

### GetHashicorpVaultEngineOk

`func (o *EnvironmentCreateParameters) GetHashicorpVaultEngineOk() (*string, bool)`

GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultEngine

`func (o *EnvironmentCreateParameters) SetHashicorpVaultEngine(v string)`

SetHashicorpVaultEngine sets HashicorpVaultEngine field to given value.

### HasHashicorpVaultEngine

`func (o *EnvironmentCreateParameters) HasHashicorpVaultEngine() bool`

HasHashicorpVaultEngine returns a boolean if a field has been set.

### GetHashicorpVaultSecretPath

`func (o *EnvironmentCreateParameters) GetHashicorpVaultSecretPath() string`

GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretPathOk

`func (o *EnvironmentCreateParameters) GetHashicorpVaultSecretPathOk() (*string, bool)`

GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretPath

`func (o *EnvironmentCreateParameters) SetHashicorpVaultSecretPath(v string)`

SetHashicorpVaultSecretPath sets HashicorpVaultSecretPath field to given value.

### HasHashicorpVaultSecretPath

`func (o *EnvironmentCreateParameters) HasHashicorpVaultSecretPath() bool`

HasHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetHashicorpVaultUsernameKey

`func (o *EnvironmentCreateParameters) GetHashicorpVaultUsernameKey() string`

GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameKeyOk

`func (o *EnvironmentCreateParameters) GetHashicorpVaultUsernameKeyOk() (*string, bool)`

GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameKey

`func (o *EnvironmentCreateParameters) SetHashicorpVaultUsernameKey(v string)`

SetHashicorpVaultUsernameKey sets HashicorpVaultUsernameKey field to given value.

### HasHashicorpVaultUsernameKey

`func (o *EnvironmentCreateParameters) HasHashicorpVaultUsernameKey() bool`

HasHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetHashicorpVaultSecretKey

`func (o *EnvironmentCreateParameters) GetHashicorpVaultSecretKey() string`

GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetHashicorpVaultSecretKeyOk

`func (o *EnvironmentCreateParameters) GetHashicorpVaultSecretKeyOk() (*string, bool)`

GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultSecretKey

`func (o *EnvironmentCreateParameters) SetHashicorpVaultSecretKey(v string)`

SetHashicorpVaultSecretKey sets HashicorpVaultSecretKey field to given value.

### HasHashicorpVaultSecretKey

`func (o *EnvironmentCreateParameters) HasHashicorpVaultSecretKey() bool`

HasHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetCyberarkVaultQueryString

`func (o *EnvironmentCreateParameters) GetCyberarkVaultQueryString() string`

GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetCyberarkVaultQueryStringOk

`func (o *EnvironmentCreateParameters) GetCyberarkVaultQueryStringOk() (*string, bool)`

GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyberarkVaultQueryString

`func (o *EnvironmentCreateParameters) SetCyberarkVaultQueryString(v string)`

SetCyberarkVaultQueryString sets CyberarkVaultQueryString field to given value.

### HasCyberarkVaultQueryString

`func (o *EnvironmentCreateParameters) HasCyberarkVaultQueryString() bool`

HasCyberarkVaultQueryString returns a boolean if a field has been set.

### GetUseKerberosAuthentication

`func (o *EnvironmentCreateParameters) GetUseKerberosAuthentication() bool`

GetUseKerberosAuthentication returns the UseKerberosAuthentication field if non-nil, zero value otherwise.

### GetUseKerberosAuthenticationOk

`func (o *EnvironmentCreateParameters) GetUseKerberosAuthenticationOk() (*bool, bool)`

GetUseKerberosAuthenticationOk returns a tuple with the UseKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKerberosAuthentication

`func (o *EnvironmentCreateParameters) SetUseKerberosAuthentication(v bool)`

SetUseKerberosAuthentication sets UseKerberosAuthentication field to given value.

### HasUseKerberosAuthentication

`func (o *EnvironmentCreateParameters) HasUseKerberosAuthentication() bool`

HasUseKerberosAuthentication returns a boolean if a field has been set.

### GetUseEnginePublicKey

`func (o *EnvironmentCreateParameters) GetUseEnginePublicKey() bool`

GetUseEnginePublicKey returns the UseEnginePublicKey field if non-nil, zero value otherwise.

### GetUseEnginePublicKeyOk

`func (o *EnvironmentCreateParameters) GetUseEnginePublicKeyOk() (*bool, bool)`

GetUseEnginePublicKeyOk returns a tuple with the UseEnginePublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnginePublicKey

`func (o *EnvironmentCreateParameters) SetUseEnginePublicKey(v bool)`

SetUseEnginePublicKey sets UseEnginePublicKey field to given value.

### HasUseEnginePublicKey

`func (o *EnvironmentCreateParameters) HasUseEnginePublicKey() bool`

HasUseEnginePublicKey returns a boolean if a field has been set.

### GetNfsAddresses

`func (o *EnvironmentCreateParameters) GetNfsAddresses() []string`

GetNfsAddresses returns the NfsAddresses field if non-nil, zero value otherwise.

### GetNfsAddressesOk

`func (o *EnvironmentCreateParameters) GetNfsAddressesOk() (*[]string, bool)`

GetNfsAddressesOk returns a tuple with the NfsAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAddresses

`func (o *EnvironmentCreateParameters) SetNfsAddresses(v []string)`

SetNfsAddresses sets NfsAddresses field to given value.

### HasNfsAddresses

`func (o *EnvironmentCreateParameters) HasNfsAddresses() bool`

HasNfsAddresses returns a boolean if a field has been set.

### GetAseDbVaultUsername

`func (o *EnvironmentCreateParameters) GetAseDbVaultUsername() string`

GetAseDbVaultUsername returns the AseDbVaultUsername field if non-nil, zero value otherwise.

### GetAseDbVaultUsernameOk

`func (o *EnvironmentCreateParameters) GetAseDbVaultUsernameOk() (*string, bool)`

GetAseDbVaultUsernameOk returns a tuple with the AseDbVaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbVaultUsername

`func (o *EnvironmentCreateParameters) SetAseDbVaultUsername(v string)`

SetAseDbVaultUsername sets AseDbVaultUsername field to given value.

### HasAseDbVaultUsername

`func (o *EnvironmentCreateParameters) HasAseDbVaultUsername() bool`

HasAseDbVaultUsername returns a boolean if a field has been set.

### GetAseDbUsername

`func (o *EnvironmentCreateParameters) GetAseDbUsername() string`

GetAseDbUsername returns the AseDbUsername field if non-nil, zero value otherwise.

### GetAseDbUsernameOk

`func (o *EnvironmentCreateParameters) GetAseDbUsernameOk() (*string, bool)`

GetAseDbUsernameOk returns a tuple with the AseDbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbUsername

`func (o *EnvironmentCreateParameters) SetAseDbUsername(v string)`

SetAseDbUsername sets AseDbUsername field to given value.

### HasAseDbUsername

`func (o *EnvironmentCreateParameters) HasAseDbUsername() bool`

HasAseDbUsername returns a boolean if a field has been set.

### GetAseDbPassword

`func (o *EnvironmentCreateParameters) GetAseDbPassword() string`

GetAseDbPassword returns the AseDbPassword field if non-nil, zero value otherwise.

### GetAseDbPasswordOk

`func (o *EnvironmentCreateParameters) GetAseDbPasswordOk() (*string, bool)`

GetAseDbPasswordOk returns a tuple with the AseDbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbPassword

`func (o *EnvironmentCreateParameters) SetAseDbPassword(v string)`

SetAseDbPassword sets AseDbPassword field to given value.

### HasAseDbPassword

`func (o *EnvironmentCreateParameters) HasAseDbPassword() bool`

HasAseDbPassword returns a boolean if a field has been set.

### GetAseDbVault

`func (o *EnvironmentCreateParameters) GetAseDbVault() string`

GetAseDbVault returns the AseDbVault field if non-nil, zero value otherwise.

### GetAseDbVaultOk

`func (o *EnvironmentCreateParameters) GetAseDbVaultOk() (*string, bool)`

GetAseDbVaultOk returns a tuple with the AseDbVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbVault

`func (o *EnvironmentCreateParameters) SetAseDbVault(v string)`

SetAseDbVault sets AseDbVault field to given value.

### HasAseDbVault

`func (o *EnvironmentCreateParameters) HasAseDbVault() bool`

HasAseDbVault returns a boolean if a field has been set.

### GetAseDbHashicorpVaultEngine

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultEngine() string`

GetAseDbHashicorpVaultEngine returns the AseDbHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultEngineOk

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultEngineOk() (*string, bool)`

GetAseDbHashicorpVaultEngineOk returns a tuple with the AseDbHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultEngine

`func (o *EnvironmentCreateParameters) SetAseDbHashicorpVaultEngine(v string)`

SetAseDbHashicorpVaultEngine sets AseDbHashicorpVaultEngine field to given value.

### HasAseDbHashicorpVaultEngine

`func (o *EnvironmentCreateParameters) HasAseDbHashicorpVaultEngine() bool`

HasAseDbHashicorpVaultEngine returns a boolean if a field has been set.

### GetAseDbHashicorpVaultSecretPath

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultSecretPath() string`

GetAseDbHashicorpVaultSecretPath returns the AseDbHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultSecretPathOk

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultSecretPathOk() (*string, bool)`

GetAseDbHashicorpVaultSecretPathOk returns a tuple with the AseDbHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultSecretPath

`func (o *EnvironmentCreateParameters) SetAseDbHashicorpVaultSecretPath(v string)`

SetAseDbHashicorpVaultSecretPath sets AseDbHashicorpVaultSecretPath field to given value.

### HasAseDbHashicorpVaultSecretPath

`func (o *EnvironmentCreateParameters) HasAseDbHashicorpVaultSecretPath() bool`

HasAseDbHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetAseDbHashicorpVaultUsernameKey

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultUsernameKey() string`

GetAseDbHashicorpVaultUsernameKey returns the AseDbHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultUsernameKeyOk

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultUsernameKeyOk() (*string, bool)`

GetAseDbHashicorpVaultUsernameKeyOk returns a tuple with the AseDbHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultUsernameKey

`func (o *EnvironmentCreateParameters) SetAseDbHashicorpVaultUsernameKey(v string)`

SetAseDbHashicorpVaultUsernameKey sets AseDbHashicorpVaultUsernameKey field to given value.

### HasAseDbHashicorpVaultUsernameKey

`func (o *EnvironmentCreateParameters) HasAseDbHashicorpVaultUsernameKey() bool`

HasAseDbHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetAseDbHashicorpVaultSecretKey

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultSecretKey() string`

GetAseDbHashicorpVaultSecretKey returns the AseDbHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetAseDbHashicorpVaultSecretKeyOk

`func (o *EnvironmentCreateParameters) GetAseDbHashicorpVaultSecretKeyOk() (*string, bool)`

GetAseDbHashicorpVaultSecretKeyOk returns a tuple with the AseDbHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbHashicorpVaultSecretKey

`func (o *EnvironmentCreateParameters) SetAseDbHashicorpVaultSecretKey(v string)`

SetAseDbHashicorpVaultSecretKey sets AseDbHashicorpVaultSecretKey field to given value.

### HasAseDbHashicorpVaultSecretKey

`func (o *EnvironmentCreateParameters) HasAseDbHashicorpVaultSecretKey() bool`

HasAseDbHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetAseDbCyberarkVaultQueryString

`func (o *EnvironmentCreateParameters) GetAseDbCyberarkVaultQueryString() string`

GetAseDbCyberarkVaultQueryString returns the AseDbCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetAseDbCyberarkVaultQueryStringOk

`func (o *EnvironmentCreateParameters) GetAseDbCyberarkVaultQueryStringOk() (*string, bool)`

GetAseDbCyberarkVaultQueryStringOk returns a tuple with the AseDbCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbCyberarkVaultQueryString

`func (o *EnvironmentCreateParameters) SetAseDbCyberarkVaultQueryString(v string)`

SetAseDbCyberarkVaultQueryString sets AseDbCyberarkVaultQueryString field to given value.

### HasAseDbCyberarkVaultQueryString

`func (o *EnvironmentCreateParameters) HasAseDbCyberarkVaultQueryString() bool`

HasAseDbCyberarkVaultQueryString returns a boolean if a field has been set.

### GetAseDbUseKerberosAuthentication

`func (o *EnvironmentCreateParameters) GetAseDbUseKerberosAuthentication() bool`

GetAseDbUseKerberosAuthentication returns the AseDbUseKerberosAuthentication field if non-nil, zero value otherwise.

### GetAseDbUseKerberosAuthenticationOk

`func (o *EnvironmentCreateParameters) GetAseDbUseKerberosAuthenticationOk() (*bool, bool)`

GetAseDbUseKerberosAuthenticationOk returns a tuple with the AseDbUseKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbUseKerberosAuthentication

`func (o *EnvironmentCreateParameters) SetAseDbUseKerberosAuthentication(v bool)`

SetAseDbUseKerberosAuthentication sets AseDbUseKerberosAuthentication field to given value.

### HasAseDbUseKerberosAuthentication

`func (o *EnvironmentCreateParameters) HasAseDbUseKerberosAuthentication() bool`

HasAseDbUseKerberosAuthentication returns a boolean if a field has been set.

### GetJavaHome

`func (o *EnvironmentCreateParameters) GetJavaHome() string`

GetJavaHome returns the JavaHome field if non-nil, zero value otherwise.

### GetJavaHomeOk

`func (o *EnvironmentCreateParameters) GetJavaHomeOk() (*string, bool)`

GetJavaHomeOk returns a tuple with the JavaHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaHome

`func (o *EnvironmentCreateParameters) SetJavaHome(v string)`

SetJavaHome sets JavaHome field to given value.

### HasJavaHome

`func (o *EnvironmentCreateParameters) HasJavaHome() bool`

HasJavaHome returns a boolean if a field has been set.

### GetDspKeystorePath

`func (o *EnvironmentCreateParameters) GetDspKeystorePath() string`

GetDspKeystorePath returns the DspKeystorePath field if non-nil, zero value otherwise.

### GetDspKeystorePathOk

`func (o *EnvironmentCreateParameters) GetDspKeystorePathOk() (*string, bool)`

GetDspKeystorePathOk returns a tuple with the DspKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePath

`func (o *EnvironmentCreateParameters) SetDspKeystorePath(v string)`

SetDspKeystorePath sets DspKeystorePath field to given value.

### HasDspKeystorePath

`func (o *EnvironmentCreateParameters) HasDspKeystorePath() bool`

HasDspKeystorePath returns a boolean if a field has been set.

### GetDspKeystorePassword

`func (o *EnvironmentCreateParameters) GetDspKeystorePassword() string`

GetDspKeystorePassword returns the DspKeystorePassword field if non-nil, zero value otherwise.

### GetDspKeystorePasswordOk

`func (o *EnvironmentCreateParameters) GetDspKeystorePasswordOk() (*string, bool)`

GetDspKeystorePasswordOk returns a tuple with the DspKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePassword

`func (o *EnvironmentCreateParameters) SetDspKeystorePassword(v string)`

SetDspKeystorePassword sets DspKeystorePassword field to given value.

### HasDspKeystorePassword

`func (o *EnvironmentCreateParameters) HasDspKeystorePassword() bool`

HasDspKeystorePassword returns a boolean if a field has been set.

### GetDspKeystoreAlias

`func (o *EnvironmentCreateParameters) GetDspKeystoreAlias() string`

GetDspKeystoreAlias returns the DspKeystoreAlias field if non-nil, zero value otherwise.

### GetDspKeystoreAliasOk

`func (o *EnvironmentCreateParameters) GetDspKeystoreAliasOk() (*string, bool)`

GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystoreAlias

`func (o *EnvironmentCreateParameters) SetDspKeystoreAlias(v string)`

SetDspKeystoreAlias sets DspKeystoreAlias field to given value.

### HasDspKeystoreAlias

`func (o *EnvironmentCreateParameters) HasDspKeystoreAlias() bool`

HasDspKeystoreAlias returns a boolean if a field has been set.

### GetDspTruststorePath

`func (o *EnvironmentCreateParameters) GetDspTruststorePath() string`

GetDspTruststorePath returns the DspTruststorePath field if non-nil, zero value otherwise.

### GetDspTruststorePathOk

`func (o *EnvironmentCreateParameters) GetDspTruststorePathOk() (*string, bool)`

GetDspTruststorePathOk returns a tuple with the DspTruststorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePath

`func (o *EnvironmentCreateParameters) SetDspTruststorePath(v string)`

SetDspTruststorePath sets DspTruststorePath field to given value.

### HasDspTruststorePath

`func (o *EnvironmentCreateParameters) HasDspTruststorePath() bool`

HasDspTruststorePath returns a boolean if a field has been set.

### GetDspTruststorePassword

`func (o *EnvironmentCreateParameters) GetDspTruststorePassword() string`

GetDspTruststorePassword returns the DspTruststorePassword field if non-nil, zero value otherwise.

### GetDspTruststorePasswordOk

`func (o *EnvironmentCreateParameters) GetDspTruststorePasswordOk() (*string, bool)`

GetDspTruststorePasswordOk returns a tuple with the DspTruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePassword

`func (o *EnvironmentCreateParameters) SetDspTruststorePassword(v string)`

SetDspTruststorePassword sets DspTruststorePassword field to given value.

### HasDspTruststorePassword

`func (o *EnvironmentCreateParameters) HasDspTruststorePassword() bool`

HasDspTruststorePassword returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentCreateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentCreateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentCreateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentCreateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *EnvironmentCreateParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EnvironmentCreateParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EnvironmentCreateParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EnvironmentCreateParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *EnvironmentCreateParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *EnvironmentCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *EnvironmentCreateParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *EnvironmentCreateParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


