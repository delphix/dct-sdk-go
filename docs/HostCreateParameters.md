# HostCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name to associate with the host. | [optional] 
**Hostname** | Pointer to **string** | The hostname or IP address of this host. | [optional] 
**NfsAddresses** | Pointer to **[]string** | The list of host/IP addresses to use for NFS export. | [optional] 
**SshPort** | Pointer to **int32** | The port number used to connect to the host via SSH. | [optional] [default to 22]
**PrivilegeElevationProfileReference** | Pointer to **string** | Reference to a profile for escalating user privileges. | [optional] 
**DspKeystoreAlias** | Pointer to **string** | The lowercase alias to use inside the user managed DSP keystore. | [optional] 
**DspKeystorePassword** | Pointer to **string** | The password for the user managed DSP keystore. | [optional] 
**DspKeystorePath** | Pointer to **string** | The path to the user managed DSP keystore. | [optional] 
**DspTruststorePassword** | Pointer to **string** | The password for the user managed DSP truststore. | [optional] 
**DspTruststorePath** | Pointer to **string** | The path to the user managed DSP truststore. | [optional] 
**JavaHome** | Pointer to **string** | The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used. | [optional] 
**ToolkitPath** | Pointer to **string** | The path for the toolkit that resides on the host. | [optional] 
**OracleJdbcKeystorePassword** | Pointer to **string** | The password for the user managed Oracle JDBC keystore. | [optional] 
**OracleTdeKeystoresRootPath** | Pointer to **string** | The path to the root of the Oracle TDE keystores artifact directories. | [optional] 
**SshVerificationStrategy** | Pointer to [**SSHVerificationStrategy**](SSHVerificationStrategy.md) |  | [optional] 
**OracleClusterNodeVirtualIps** | Pointer to [**[]OracleVirtualIP**](OracleVirtualIP.md) | The Virtual IP addresses associated with the OracleClusterNode. | [optional] 

## Methods

### NewHostCreateParameters

`func NewHostCreateParameters() *HostCreateParameters`

NewHostCreateParameters instantiates a new HostCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostCreateParametersWithDefaults

`func NewHostCreateParametersWithDefaults() *HostCreateParameters`

NewHostCreateParametersWithDefaults instantiates a new HostCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HostCreateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostCreateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostCreateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostCreateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *HostCreateParameters) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HostCreateParameters) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HostCreateParameters) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HostCreateParameters) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNfsAddresses

`func (o *HostCreateParameters) GetNfsAddresses() []string`

GetNfsAddresses returns the NfsAddresses field if non-nil, zero value otherwise.

### GetNfsAddressesOk

`func (o *HostCreateParameters) GetNfsAddressesOk() (*[]string, bool)`

GetNfsAddressesOk returns a tuple with the NfsAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAddresses

`func (o *HostCreateParameters) SetNfsAddresses(v []string)`

SetNfsAddresses sets NfsAddresses field to given value.

### HasNfsAddresses

`func (o *HostCreateParameters) HasNfsAddresses() bool`

HasNfsAddresses returns a boolean if a field has been set.

### GetSshPort

`func (o *HostCreateParameters) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *HostCreateParameters) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *HostCreateParameters) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *HostCreateParameters) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetPrivilegeElevationProfileReference

`func (o *HostCreateParameters) GetPrivilegeElevationProfileReference() string`

GetPrivilegeElevationProfileReference returns the PrivilegeElevationProfileReference field if non-nil, zero value otherwise.

### GetPrivilegeElevationProfileReferenceOk

`func (o *HostCreateParameters) GetPrivilegeElevationProfileReferenceOk() (*string, bool)`

GetPrivilegeElevationProfileReferenceOk returns a tuple with the PrivilegeElevationProfileReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeElevationProfileReference

`func (o *HostCreateParameters) SetPrivilegeElevationProfileReference(v string)`

SetPrivilegeElevationProfileReference sets PrivilegeElevationProfileReference field to given value.

### HasPrivilegeElevationProfileReference

`func (o *HostCreateParameters) HasPrivilegeElevationProfileReference() bool`

HasPrivilegeElevationProfileReference returns a boolean if a field has been set.

### GetDspKeystoreAlias

`func (o *HostCreateParameters) GetDspKeystoreAlias() string`

GetDspKeystoreAlias returns the DspKeystoreAlias field if non-nil, zero value otherwise.

### GetDspKeystoreAliasOk

`func (o *HostCreateParameters) GetDspKeystoreAliasOk() (*string, bool)`

GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystoreAlias

`func (o *HostCreateParameters) SetDspKeystoreAlias(v string)`

SetDspKeystoreAlias sets DspKeystoreAlias field to given value.

### HasDspKeystoreAlias

`func (o *HostCreateParameters) HasDspKeystoreAlias() bool`

HasDspKeystoreAlias returns a boolean if a field has been set.

### GetDspKeystorePassword

`func (o *HostCreateParameters) GetDspKeystorePassword() string`

GetDspKeystorePassword returns the DspKeystorePassword field if non-nil, zero value otherwise.

### GetDspKeystorePasswordOk

`func (o *HostCreateParameters) GetDspKeystorePasswordOk() (*string, bool)`

GetDspKeystorePasswordOk returns a tuple with the DspKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePassword

`func (o *HostCreateParameters) SetDspKeystorePassword(v string)`

SetDspKeystorePassword sets DspKeystorePassword field to given value.

### HasDspKeystorePassword

`func (o *HostCreateParameters) HasDspKeystorePassword() bool`

HasDspKeystorePassword returns a boolean if a field has been set.

### GetDspKeystorePath

`func (o *HostCreateParameters) GetDspKeystorePath() string`

GetDspKeystorePath returns the DspKeystorePath field if non-nil, zero value otherwise.

### GetDspKeystorePathOk

`func (o *HostCreateParameters) GetDspKeystorePathOk() (*string, bool)`

GetDspKeystorePathOk returns a tuple with the DspKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePath

`func (o *HostCreateParameters) SetDspKeystorePath(v string)`

SetDspKeystorePath sets DspKeystorePath field to given value.

### HasDspKeystorePath

`func (o *HostCreateParameters) HasDspKeystorePath() bool`

HasDspKeystorePath returns a boolean if a field has been set.

### GetDspTruststorePassword

`func (o *HostCreateParameters) GetDspTruststorePassword() string`

GetDspTruststorePassword returns the DspTruststorePassword field if non-nil, zero value otherwise.

### GetDspTruststorePasswordOk

`func (o *HostCreateParameters) GetDspTruststorePasswordOk() (*string, bool)`

GetDspTruststorePasswordOk returns a tuple with the DspTruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePassword

`func (o *HostCreateParameters) SetDspTruststorePassword(v string)`

SetDspTruststorePassword sets DspTruststorePassword field to given value.

### HasDspTruststorePassword

`func (o *HostCreateParameters) HasDspTruststorePassword() bool`

HasDspTruststorePassword returns a boolean if a field has been set.

### GetDspTruststorePath

`func (o *HostCreateParameters) GetDspTruststorePath() string`

GetDspTruststorePath returns the DspTruststorePath field if non-nil, zero value otherwise.

### GetDspTruststorePathOk

`func (o *HostCreateParameters) GetDspTruststorePathOk() (*string, bool)`

GetDspTruststorePathOk returns a tuple with the DspTruststorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePath

`func (o *HostCreateParameters) SetDspTruststorePath(v string)`

SetDspTruststorePath sets DspTruststorePath field to given value.

### HasDspTruststorePath

`func (o *HostCreateParameters) HasDspTruststorePath() bool`

HasDspTruststorePath returns a boolean if a field has been set.

### GetJavaHome

`func (o *HostCreateParameters) GetJavaHome() string`

GetJavaHome returns the JavaHome field if non-nil, zero value otherwise.

### GetJavaHomeOk

`func (o *HostCreateParameters) GetJavaHomeOk() (*string, bool)`

GetJavaHomeOk returns a tuple with the JavaHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaHome

`func (o *HostCreateParameters) SetJavaHome(v string)`

SetJavaHome sets JavaHome field to given value.

### HasJavaHome

`func (o *HostCreateParameters) HasJavaHome() bool`

HasJavaHome returns a boolean if a field has been set.

### GetToolkitPath

`func (o *HostCreateParameters) GetToolkitPath() string`

GetToolkitPath returns the ToolkitPath field if non-nil, zero value otherwise.

### GetToolkitPathOk

`func (o *HostCreateParameters) GetToolkitPathOk() (*string, bool)`

GetToolkitPathOk returns a tuple with the ToolkitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolkitPath

`func (o *HostCreateParameters) SetToolkitPath(v string)`

SetToolkitPath sets ToolkitPath field to given value.

### HasToolkitPath

`func (o *HostCreateParameters) HasToolkitPath() bool`

HasToolkitPath returns a boolean if a field has been set.

### GetOracleJdbcKeystorePassword

`func (o *HostCreateParameters) GetOracleJdbcKeystorePassword() string`

GetOracleJdbcKeystorePassword returns the OracleJdbcKeystorePassword field if non-nil, zero value otherwise.

### GetOracleJdbcKeystorePasswordOk

`func (o *HostCreateParameters) GetOracleJdbcKeystorePasswordOk() (*string, bool)`

GetOracleJdbcKeystorePasswordOk returns a tuple with the OracleJdbcKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleJdbcKeystorePassword

`func (o *HostCreateParameters) SetOracleJdbcKeystorePassword(v string)`

SetOracleJdbcKeystorePassword sets OracleJdbcKeystorePassword field to given value.

### HasOracleJdbcKeystorePassword

`func (o *HostCreateParameters) HasOracleJdbcKeystorePassword() bool`

HasOracleJdbcKeystorePassword returns a boolean if a field has been set.

### GetOracleTdeKeystoresRootPath

`func (o *HostCreateParameters) GetOracleTdeKeystoresRootPath() string`

GetOracleTdeKeystoresRootPath returns the OracleTdeKeystoresRootPath field if non-nil, zero value otherwise.

### GetOracleTdeKeystoresRootPathOk

`func (o *HostCreateParameters) GetOracleTdeKeystoresRootPathOk() (*string, bool)`

GetOracleTdeKeystoresRootPathOk returns a tuple with the OracleTdeKeystoresRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTdeKeystoresRootPath

`func (o *HostCreateParameters) SetOracleTdeKeystoresRootPath(v string)`

SetOracleTdeKeystoresRootPath sets OracleTdeKeystoresRootPath field to given value.

### HasOracleTdeKeystoresRootPath

`func (o *HostCreateParameters) HasOracleTdeKeystoresRootPath() bool`

HasOracleTdeKeystoresRootPath returns a boolean if a field has been set.

### GetSshVerificationStrategy

`func (o *HostCreateParameters) GetSshVerificationStrategy() SSHVerificationStrategy`

GetSshVerificationStrategy returns the SshVerificationStrategy field if non-nil, zero value otherwise.

### GetSshVerificationStrategyOk

`func (o *HostCreateParameters) GetSshVerificationStrategyOk() (*SSHVerificationStrategy, bool)`

GetSshVerificationStrategyOk returns a tuple with the SshVerificationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshVerificationStrategy

`func (o *HostCreateParameters) SetSshVerificationStrategy(v SSHVerificationStrategy)`

SetSshVerificationStrategy sets SshVerificationStrategy field to given value.

### HasSshVerificationStrategy

`func (o *HostCreateParameters) HasSshVerificationStrategy() bool`

HasSshVerificationStrategy returns a boolean if a field has been set.

### GetOracleClusterNodeVirtualIps

`func (o *HostCreateParameters) GetOracleClusterNodeVirtualIps() []OracleVirtualIP`

GetOracleClusterNodeVirtualIps returns the OracleClusterNodeVirtualIps field if non-nil, zero value otherwise.

### GetOracleClusterNodeVirtualIpsOk

`func (o *HostCreateParameters) GetOracleClusterNodeVirtualIpsOk() (*[]OracleVirtualIP, bool)`

GetOracleClusterNodeVirtualIpsOk returns a tuple with the OracleClusterNodeVirtualIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeVirtualIps

`func (o *HostCreateParameters) SetOracleClusterNodeVirtualIps(v []OracleVirtualIP)`

SetOracleClusterNodeVirtualIps sets OracleClusterNodeVirtualIps field to given value.

### HasOracleClusterNodeVirtualIps

`func (o *HostCreateParameters) HasOracleClusterNodeVirtualIps() bool`

HasOracleClusterNodeVirtualIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


