# HostUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | host address of the machine. | [optional] 
**OracleClusterNodeName** | Pointer to **string** | The name of the associated OracleClusterNode. | [optional] 
**OracleClusterNodeEnabled** | Pointer to **bool** | Whether the associated OracleClusterNode is enabled. | [optional] 
**OracleClusterNodeVirtualIps** | Pointer to [**[]OracleVirtualIP**](OracleVirtualIP.md) | The Virtual IP addresses associated with the OracleClusterNode. | [optional] 
**NfsAddresses** | Pointer to **[]string** | array of ip addresses or hostnames | [optional] 
**SshPort** | Pointer to **int64** | ssh port of the host. | [optional] 
**ToolkitPath** | Pointer to **string** | The path for the toolkit that resides on the host. | [optional] 
**JavaHome** | Pointer to **string** | The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used. | [optional] 
**DspKeystorePath** | Pointer to **string** | DSP keystore path. | [optional] 
**DspKeystorePassword** | Pointer to **string** | DSP keystore password. | [optional] 
**DspKeystoreAlias** | Pointer to **string** | DSP keystore alias. | [optional] 
**DspTruststorePath** | Pointer to **string** | DSP truststore path. | [optional] 
**DspTruststorePassword** | Pointer to **string** | DSP truststore password. | [optional] 
**ConnectorPort** | Pointer to **int32** | Specify port on which Delphix connector will run. | [optional] 
**OracleJdbcKeystorePassword** | Pointer to **string** | The password for the user managed Oracle JDBC keystore. | [optional] 
**OracleTdeKeystoresRootPath** | Pointer to **string** | The path to the root of the Oracle TDE keystores artifact directories. | [optional] 
**SshVerificationStrategy** | Pointer to [**SSHVerificationStrategy**](SSHVerificationStrategy.md) |  | [optional] 
**ConnectorAuthenticationKey** | Pointer to **string** | Unique per Delphix key used to authenticate with the remote Delphix Connector. | [optional] 

## Methods

### NewHostUpdateParameters

`func NewHostUpdateParameters() *HostUpdateParameters`

NewHostUpdateParameters instantiates a new HostUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostUpdateParametersWithDefaults

`func NewHostUpdateParametersWithDefaults() *HostUpdateParameters`

NewHostUpdateParametersWithDefaults instantiates a new HostUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *HostUpdateParameters) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HostUpdateParameters) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HostUpdateParameters) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HostUpdateParameters) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOracleClusterNodeName

`func (o *HostUpdateParameters) GetOracleClusterNodeName() string`

GetOracleClusterNodeName returns the OracleClusterNodeName field if non-nil, zero value otherwise.

### GetOracleClusterNodeNameOk

`func (o *HostUpdateParameters) GetOracleClusterNodeNameOk() (*string, bool)`

GetOracleClusterNodeNameOk returns a tuple with the OracleClusterNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeName

`func (o *HostUpdateParameters) SetOracleClusterNodeName(v string)`

SetOracleClusterNodeName sets OracleClusterNodeName field to given value.

### HasOracleClusterNodeName

`func (o *HostUpdateParameters) HasOracleClusterNodeName() bool`

HasOracleClusterNodeName returns a boolean if a field has been set.

### GetOracleClusterNodeEnabled

`func (o *HostUpdateParameters) GetOracleClusterNodeEnabled() bool`

GetOracleClusterNodeEnabled returns the OracleClusterNodeEnabled field if non-nil, zero value otherwise.

### GetOracleClusterNodeEnabledOk

`func (o *HostUpdateParameters) GetOracleClusterNodeEnabledOk() (*bool, bool)`

GetOracleClusterNodeEnabledOk returns a tuple with the OracleClusterNodeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeEnabled

`func (o *HostUpdateParameters) SetOracleClusterNodeEnabled(v bool)`

SetOracleClusterNodeEnabled sets OracleClusterNodeEnabled field to given value.

### HasOracleClusterNodeEnabled

`func (o *HostUpdateParameters) HasOracleClusterNodeEnabled() bool`

HasOracleClusterNodeEnabled returns a boolean if a field has been set.

### GetOracleClusterNodeVirtualIps

`func (o *HostUpdateParameters) GetOracleClusterNodeVirtualIps() []OracleVirtualIP`

GetOracleClusterNodeVirtualIps returns the OracleClusterNodeVirtualIps field if non-nil, zero value otherwise.

### GetOracleClusterNodeVirtualIpsOk

`func (o *HostUpdateParameters) GetOracleClusterNodeVirtualIpsOk() (*[]OracleVirtualIP, bool)`

GetOracleClusterNodeVirtualIpsOk returns a tuple with the OracleClusterNodeVirtualIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeVirtualIps

`func (o *HostUpdateParameters) SetOracleClusterNodeVirtualIps(v []OracleVirtualIP)`

SetOracleClusterNodeVirtualIps sets OracleClusterNodeVirtualIps field to given value.

### HasOracleClusterNodeVirtualIps

`func (o *HostUpdateParameters) HasOracleClusterNodeVirtualIps() bool`

HasOracleClusterNodeVirtualIps returns a boolean if a field has been set.

### GetNfsAddresses

`func (o *HostUpdateParameters) GetNfsAddresses() []string`

GetNfsAddresses returns the NfsAddresses field if non-nil, zero value otherwise.

### GetNfsAddressesOk

`func (o *HostUpdateParameters) GetNfsAddressesOk() (*[]string, bool)`

GetNfsAddressesOk returns a tuple with the NfsAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAddresses

`func (o *HostUpdateParameters) SetNfsAddresses(v []string)`

SetNfsAddresses sets NfsAddresses field to given value.

### HasNfsAddresses

`func (o *HostUpdateParameters) HasNfsAddresses() bool`

HasNfsAddresses returns a boolean if a field has been set.

### GetSshPort

`func (o *HostUpdateParameters) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *HostUpdateParameters) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *HostUpdateParameters) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *HostUpdateParameters) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetToolkitPath

`func (o *HostUpdateParameters) GetToolkitPath() string`

GetToolkitPath returns the ToolkitPath field if non-nil, zero value otherwise.

### GetToolkitPathOk

`func (o *HostUpdateParameters) GetToolkitPathOk() (*string, bool)`

GetToolkitPathOk returns a tuple with the ToolkitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolkitPath

`func (o *HostUpdateParameters) SetToolkitPath(v string)`

SetToolkitPath sets ToolkitPath field to given value.

### HasToolkitPath

`func (o *HostUpdateParameters) HasToolkitPath() bool`

HasToolkitPath returns a boolean if a field has been set.

### GetJavaHome

`func (o *HostUpdateParameters) GetJavaHome() string`

GetJavaHome returns the JavaHome field if non-nil, zero value otherwise.

### GetJavaHomeOk

`func (o *HostUpdateParameters) GetJavaHomeOk() (*string, bool)`

GetJavaHomeOk returns a tuple with the JavaHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaHome

`func (o *HostUpdateParameters) SetJavaHome(v string)`

SetJavaHome sets JavaHome field to given value.

### HasJavaHome

`func (o *HostUpdateParameters) HasJavaHome() bool`

HasJavaHome returns a boolean if a field has been set.

### GetDspKeystorePath

`func (o *HostUpdateParameters) GetDspKeystorePath() string`

GetDspKeystorePath returns the DspKeystorePath field if non-nil, zero value otherwise.

### GetDspKeystorePathOk

`func (o *HostUpdateParameters) GetDspKeystorePathOk() (*string, bool)`

GetDspKeystorePathOk returns a tuple with the DspKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePath

`func (o *HostUpdateParameters) SetDspKeystorePath(v string)`

SetDspKeystorePath sets DspKeystorePath field to given value.

### HasDspKeystorePath

`func (o *HostUpdateParameters) HasDspKeystorePath() bool`

HasDspKeystorePath returns a boolean if a field has been set.

### GetDspKeystorePassword

`func (o *HostUpdateParameters) GetDspKeystorePassword() string`

GetDspKeystorePassword returns the DspKeystorePassword field if non-nil, zero value otherwise.

### GetDspKeystorePasswordOk

`func (o *HostUpdateParameters) GetDspKeystorePasswordOk() (*string, bool)`

GetDspKeystorePasswordOk returns a tuple with the DspKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePassword

`func (o *HostUpdateParameters) SetDspKeystorePassword(v string)`

SetDspKeystorePassword sets DspKeystorePassword field to given value.

### HasDspKeystorePassword

`func (o *HostUpdateParameters) HasDspKeystorePassword() bool`

HasDspKeystorePassword returns a boolean if a field has been set.

### GetDspKeystoreAlias

`func (o *HostUpdateParameters) GetDspKeystoreAlias() string`

GetDspKeystoreAlias returns the DspKeystoreAlias field if non-nil, zero value otherwise.

### GetDspKeystoreAliasOk

`func (o *HostUpdateParameters) GetDspKeystoreAliasOk() (*string, bool)`

GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystoreAlias

`func (o *HostUpdateParameters) SetDspKeystoreAlias(v string)`

SetDspKeystoreAlias sets DspKeystoreAlias field to given value.

### HasDspKeystoreAlias

`func (o *HostUpdateParameters) HasDspKeystoreAlias() bool`

HasDspKeystoreAlias returns a boolean if a field has been set.

### GetDspTruststorePath

`func (o *HostUpdateParameters) GetDspTruststorePath() string`

GetDspTruststorePath returns the DspTruststorePath field if non-nil, zero value otherwise.

### GetDspTruststorePathOk

`func (o *HostUpdateParameters) GetDspTruststorePathOk() (*string, bool)`

GetDspTruststorePathOk returns a tuple with the DspTruststorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePath

`func (o *HostUpdateParameters) SetDspTruststorePath(v string)`

SetDspTruststorePath sets DspTruststorePath field to given value.

### HasDspTruststorePath

`func (o *HostUpdateParameters) HasDspTruststorePath() bool`

HasDspTruststorePath returns a boolean if a field has been set.

### GetDspTruststorePassword

`func (o *HostUpdateParameters) GetDspTruststorePassword() string`

GetDspTruststorePassword returns the DspTruststorePassword field if non-nil, zero value otherwise.

### GetDspTruststorePasswordOk

`func (o *HostUpdateParameters) GetDspTruststorePasswordOk() (*string, bool)`

GetDspTruststorePasswordOk returns a tuple with the DspTruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePassword

`func (o *HostUpdateParameters) SetDspTruststorePassword(v string)`

SetDspTruststorePassword sets DspTruststorePassword field to given value.

### HasDspTruststorePassword

`func (o *HostUpdateParameters) HasDspTruststorePassword() bool`

HasDspTruststorePassword returns a boolean if a field has been set.

### GetConnectorPort

`func (o *HostUpdateParameters) GetConnectorPort() int32`

GetConnectorPort returns the ConnectorPort field if non-nil, zero value otherwise.

### GetConnectorPortOk

`func (o *HostUpdateParameters) GetConnectorPortOk() (*int32, bool)`

GetConnectorPortOk returns a tuple with the ConnectorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPort

`func (o *HostUpdateParameters) SetConnectorPort(v int32)`

SetConnectorPort sets ConnectorPort field to given value.

### HasConnectorPort

`func (o *HostUpdateParameters) HasConnectorPort() bool`

HasConnectorPort returns a boolean if a field has been set.

### GetOracleJdbcKeystorePassword

`func (o *HostUpdateParameters) GetOracleJdbcKeystorePassword() string`

GetOracleJdbcKeystorePassword returns the OracleJdbcKeystorePassword field if non-nil, zero value otherwise.

### GetOracleJdbcKeystorePasswordOk

`func (o *HostUpdateParameters) GetOracleJdbcKeystorePasswordOk() (*string, bool)`

GetOracleJdbcKeystorePasswordOk returns a tuple with the OracleJdbcKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleJdbcKeystorePassword

`func (o *HostUpdateParameters) SetOracleJdbcKeystorePassword(v string)`

SetOracleJdbcKeystorePassword sets OracleJdbcKeystorePassword field to given value.

### HasOracleJdbcKeystorePassword

`func (o *HostUpdateParameters) HasOracleJdbcKeystorePassword() bool`

HasOracleJdbcKeystorePassword returns a boolean if a field has been set.

### GetOracleTdeKeystoresRootPath

`func (o *HostUpdateParameters) GetOracleTdeKeystoresRootPath() string`

GetOracleTdeKeystoresRootPath returns the OracleTdeKeystoresRootPath field if non-nil, zero value otherwise.

### GetOracleTdeKeystoresRootPathOk

`func (o *HostUpdateParameters) GetOracleTdeKeystoresRootPathOk() (*string, bool)`

GetOracleTdeKeystoresRootPathOk returns a tuple with the OracleTdeKeystoresRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTdeKeystoresRootPath

`func (o *HostUpdateParameters) SetOracleTdeKeystoresRootPath(v string)`

SetOracleTdeKeystoresRootPath sets OracleTdeKeystoresRootPath field to given value.

### HasOracleTdeKeystoresRootPath

`func (o *HostUpdateParameters) HasOracleTdeKeystoresRootPath() bool`

HasOracleTdeKeystoresRootPath returns a boolean if a field has been set.

### GetSshVerificationStrategy

`func (o *HostUpdateParameters) GetSshVerificationStrategy() SSHVerificationStrategy`

GetSshVerificationStrategy returns the SshVerificationStrategy field if non-nil, zero value otherwise.

### GetSshVerificationStrategyOk

`func (o *HostUpdateParameters) GetSshVerificationStrategyOk() (*SSHVerificationStrategy, bool)`

GetSshVerificationStrategyOk returns a tuple with the SshVerificationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshVerificationStrategy

`func (o *HostUpdateParameters) SetSshVerificationStrategy(v SSHVerificationStrategy)`

SetSshVerificationStrategy sets SshVerificationStrategy field to given value.

### HasSshVerificationStrategy

`func (o *HostUpdateParameters) HasSshVerificationStrategy() bool`

HasSshVerificationStrategy returns a boolean if a field has been set.

### GetConnectorAuthenticationKey

`func (o *HostUpdateParameters) GetConnectorAuthenticationKey() string`

GetConnectorAuthenticationKey returns the ConnectorAuthenticationKey field if non-nil, zero value otherwise.

### GetConnectorAuthenticationKeyOk

`func (o *HostUpdateParameters) GetConnectorAuthenticationKeyOk() (*string, bool)`

GetConnectorAuthenticationKeyOk returns a tuple with the ConnectorAuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorAuthenticationKey

`func (o *HostUpdateParameters) SetConnectorAuthenticationKey(v string)`

SetConnectorAuthenticationKey sets ConnectorAuthenticationKey field to given value.

### HasConnectorAuthenticationKey

`func (o *HostUpdateParameters) HasConnectorAuthenticationKey() bool`

HasConnectorAuthenticationKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


