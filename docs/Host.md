# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The entity ID of this Host. | [optional] 
**Hostname** | Pointer to **string** | The hostname or IP address of this host. | [optional] 
**OsName** | Pointer to **string** | The name of the OS on this host. | [optional] 
**OsVersion** | Pointer to **string** | The version of the OS on this host. | [optional] 
**MemorySize** | Pointer to **int64** | The total amount of memory on this host in bytes. | [optional] 
**Available** | Pointer to **bool** | True if the host is up and a connection can be established from the engine. | [optional] 
**AvailableTimestamp** | Pointer to **time.Time** | The last time the available property was updated. | [optional] 
**NotAvailableReason** | Pointer to **string** | The reason why the host is not available. | [optional] 
**OracleClusterNodeReference** | Pointer to **string** | The reference to the associated OracleClusterNode. | [optional] 
**OracleClusterNodeName** | Pointer to **string** | The name of the associated OracleClusterNode. | [optional] 
**OracleClusterNodeEnabled** | Pointer to **bool** | Whether the associated OracleClusterNode is enabled. | [optional] 
**OracleClusterNodeDiscovered** | Pointer to **bool** | Whether the associated OracleClusterNode was discovered. | [optional] 
**OracleClusterNodeVirtualIps** | Pointer to [**[]OracleVirtualIP**](OracleVirtualIP.md) | The Virtual IP addresses associated with the OracleClusterNode. | [optional] 
**OracleClusterNodeInstances** | Pointer to [**[]OracleClusterNodeInstance**](OracleClusterNodeInstance.md) | The instances associated with the OracleClusterNode. | [optional] 
**WindowsClusterNodeReference** | Pointer to **string** | The reference to the associated WindowsClusterNode. | [optional] 
**WindowsClusterNodeName** | Pointer to **string** | The name of the associated WindowsClusterNode. | [optional] 
**WindowsClusterNodeDiscovered** | Pointer to **bool** | Whether the associated Windows cluster node was discovered. | [optional] 
**NfsAddresses** | Pointer to **[]string** | The list of host/IP addresses to use for NFS export. | [optional] 
**DspKeystoreAlias** | Pointer to **string** | The lowercase alias to use inside the user managed DSP keystore. | [optional] 
**DspKeystorePath** | Pointer to **string** | The path to the user managed DSP keystore. | [optional] 
**DspTruststorePath** | Pointer to **string** | The path to the user managed DSP truststore. | [optional] 
**JavaHome** | Pointer to **string** | The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used. | [optional] 
**SshPort** | Pointer to **int32** | The port number used to connect to the host via SSH. | [optional] 
**ToolkitPath** | Pointer to **string** | The path for the toolkit that resides on the host. | [optional] 
**OracleTdeKeystoresRootPath** | Pointer to **string** | The path to the root of the Oracle TDE keystores artifact directories. | [optional] 
**ProcessorType** | Pointer to **string** | The platform for the host machine. | [optional] 
**Timezone** | Pointer to **string** | The OS timezone. | [optional] 

## Methods

### NewHost

`func NewHost() *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Host) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Host) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHostname

`func (o *Host) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Host) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Host) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Host) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOsName

`func (o *Host) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *Host) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *Host) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *Host) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *Host) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Host) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Host) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *Host) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetMemorySize

`func (o *Host) GetMemorySize() int64`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *Host) GetMemorySizeOk() (*int64, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *Host) SetMemorySize(v int64)`

SetMemorySize sets MemorySize field to given value.

### HasMemorySize

`func (o *Host) HasMemorySize() bool`

HasMemorySize returns a boolean if a field has been set.

### GetAvailable

`func (o *Host) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Host) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Host) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Host) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetAvailableTimestamp

`func (o *Host) GetAvailableTimestamp() time.Time`

GetAvailableTimestamp returns the AvailableTimestamp field if non-nil, zero value otherwise.

### GetAvailableTimestampOk

`func (o *Host) GetAvailableTimestampOk() (*time.Time, bool)`

GetAvailableTimestampOk returns a tuple with the AvailableTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTimestamp

`func (o *Host) SetAvailableTimestamp(v time.Time)`

SetAvailableTimestamp sets AvailableTimestamp field to given value.

### HasAvailableTimestamp

`func (o *Host) HasAvailableTimestamp() bool`

HasAvailableTimestamp returns a boolean if a field has been set.

### GetNotAvailableReason

`func (o *Host) GetNotAvailableReason() string`

GetNotAvailableReason returns the NotAvailableReason field if non-nil, zero value otherwise.

### GetNotAvailableReasonOk

`func (o *Host) GetNotAvailableReasonOk() (*string, bool)`

GetNotAvailableReasonOk returns a tuple with the NotAvailableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAvailableReason

`func (o *Host) SetNotAvailableReason(v string)`

SetNotAvailableReason sets NotAvailableReason field to given value.

### HasNotAvailableReason

`func (o *Host) HasNotAvailableReason() bool`

HasNotAvailableReason returns a boolean if a field has been set.

### GetOracleClusterNodeReference

`func (o *Host) GetOracleClusterNodeReference() string`

GetOracleClusterNodeReference returns the OracleClusterNodeReference field if non-nil, zero value otherwise.

### GetOracleClusterNodeReferenceOk

`func (o *Host) GetOracleClusterNodeReferenceOk() (*string, bool)`

GetOracleClusterNodeReferenceOk returns a tuple with the OracleClusterNodeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeReference

`func (o *Host) SetOracleClusterNodeReference(v string)`

SetOracleClusterNodeReference sets OracleClusterNodeReference field to given value.

### HasOracleClusterNodeReference

`func (o *Host) HasOracleClusterNodeReference() bool`

HasOracleClusterNodeReference returns a boolean if a field has been set.

### GetOracleClusterNodeName

`func (o *Host) GetOracleClusterNodeName() string`

GetOracleClusterNodeName returns the OracleClusterNodeName field if non-nil, zero value otherwise.

### GetOracleClusterNodeNameOk

`func (o *Host) GetOracleClusterNodeNameOk() (*string, bool)`

GetOracleClusterNodeNameOk returns a tuple with the OracleClusterNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeName

`func (o *Host) SetOracleClusterNodeName(v string)`

SetOracleClusterNodeName sets OracleClusterNodeName field to given value.

### HasOracleClusterNodeName

`func (o *Host) HasOracleClusterNodeName() bool`

HasOracleClusterNodeName returns a boolean if a field has been set.

### GetOracleClusterNodeEnabled

`func (o *Host) GetOracleClusterNodeEnabled() bool`

GetOracleClusterNodeEnabled returns the OracleClusterNodeEnabled field if non-nil, zero value otherwise.

### GetOracleClusterNodeEnabledOk

`func (o *Host) GetOracleClusterNodeEnabledOk() (*bool, bool)`

GetOracleClusterNodeEnabledOk returns a tuple with the OracleClusterNodeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeEnabled

`func (o *Host) SetOracleClusterNodeEnabled(v bool)`

SetOracleClusterNodeEnabled sets OracleClusterNodeEnabled field to given value.

### HasOracleClusterNodeEnabled

`func (o *Host) HasOracleClusterNodeEnabled() bool`

HasOracleClusterNodeEnabled returns a boolean if a field has been set.

### GetOracleClusterNodeDiscovered

`func (o *Host) GetOracleClusterNodeDiscovered() bool`

GetOracleClusterNodeDiscovered returns the OracleClusterNodeDiscovered field if non-nil, zero value otherwise.

### GetOracleClusterNodeDiscoveredOk

`func (o *Host) GetOracleClusterNodeDiscoveredOk() (*bool, bool)`

GetOracleClusterNodeDiscoveredOk returns a tuple with the OracleClusterNodeDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeDiscovered

`func (o *Host) SetOracleClusterNodeDiscovered(v bool)`

SetOracleClusterNodeDiscovered sets OracleClusterNodeDiscovered field to given value.

### HasOracleClusterNodeDiscovered

`func (o *Host) HasOracleClusterNodeDiscovered() bool`

HasOracleClusterNodeDiscovered returns a boolean if a field has been set.

### GetOracleClusterNodeVirtualIps

`func (o *Host) GetOracleClusterNodeVirtualIps() []OracleVirtualIP`

GetOracleClusterNodeVirtualIps returns the OracleClusterNodeVirtualIps field if non-nil, zero value otherwise.

### GetOracleClusterNodeVirtualIpsOk

`func (o *Host) GetOracleClusterNodeVirtualIpsOk() (*[]OracleVirtualIP, bool)`

GetOracleClusterNodeVirtualIpsOk returns a tuple with the OracleClusterNodeVirtualIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeVirtualIps

`func (o *Host) SetOracleClusterNodeVirtualIps(v []OracleVirtualIP)`

SetOracleClusterNodeVirtualIps sets OracleClusterNodeVirtualIps field to given value.

### HasOracleClusterNodeVirtualIps

`func (o *Host) HasOracleClusterNodeVirtualIps() bool`

HasOracleClusterNodeVirtualIps returns a boolean if a field has been set.

### GetOracleClusterNodeInstances

`func (o *Host) GetOracleClusterNodeInstances() []OracleClusterNodeInstance`

GetOracleClusterNodeInstances returns the OracleClusterNodeInstances field if non-nil, zero value otherwise.

### GetOracleClusterNodeInstancesOk

`func (o *Host) GetOracleClusterNodeInstancesOk() (*[]OracleClusterNodeInstance, bool)`

GetOracleClusterNodeInstancesOk returns a tuple with the OracleClusterNodeInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleClusterNodeInstances

`func (o *Host) SetOracleClusterNodeInstances(v []OracleClusterNodeInstance)`

SetOracleClusterNodeInstances sets OracleClusterNodeInstances field to given value.

### HasOracleClusterNodeInstances

`func (o *Host) HasOracleClusterNodeInstances() bool`

HasOracleClusterNodeInstances returns a boolean if a field has been set.

### GetWindowsClusterNodeReference

`func (o *Host) GetWindowsClusterNodeReference() string`

GetWindowsClusterNodeReference returns the WindowsClusterNodeReference field if non-nil, zero value otherwise.

### GetWindowsClusterNodeReferenceOk

`func (o *Host) GetWindowsClusterNodeReferenceOk() (*string, bool)`

GetWindowsClusterNodeReferenceOk returns a tuple with the WindowsClusterNodeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsClusterNodeReference

`func (o *Host) SetWindowsClusterNodeReference(v string)`

SetWindowsClusterNodeReference sets WindowsClusterNodeReference field to given value.

### HasWindowsClusterNodeReference

`func (o *Host) HasWindowsClusterNodeReference() bool`

HasWindowsClusterNodeReference returns a boolean if a field has been set.

### GetWindowsClusterNodeName

`func (o *Host) GetWindowsClusterNodeName() string`

GetWindowsClusterNodeName returns the WindowsClusterNodeName field if non-nil, zero value otherwise.

### GetWindowsClusterNodeNameOk

`func (o *Host) GetWindowsClusterNodeNameOk() (*string, bool)`

GetWindowsClusterNodeNameOk returns a tuple with the WindowsClusterNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsClusterNodeName

`func (o *Host) SetWindowsClusterNodeName(v string)`

SetWindowsClusterNodeName sets WindowsClusterNodeName field to given value.

### HasWindowsClusterNodeName

`func (o *Host) HasWindowsClusterNodeName() bool`

HasWindowsClusterNodeName returns a boolean if a field has been set.

### GetWindowsClusterNodeDiscovered

`func (o *Host) GetWindowsClusterNodeDiscovered() bool`

GetWindowsClusterNodeDiscovered returns the WindowsClusterNodeDiscovered field if non-nil, zero value otherwise.

### GetWindowsClusterNodeDiscoveredOk

`func (o *Host) GetWindowsClusterNodeDiscoveredOk() (*bool, bool)`

GetWindowsClusterNodeDiscoveredOk returns a tuple with the WindowsClusterNodeDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsClusterNodeDiscovered

`func (o *Host) SetWindowsClusterNodeDiscovered(v bool)`

SetWindowsClusterNodeDiscovered sets WindowsClusterNodeDiscovered field to given value.

### HasWindowsClusterNodeDiscovered

`func (o *Host) HasWindowsClusterNodeDiscovered() bool`

HasWindowsClusterNodeDiscovered returns a boolean if a field has been set.

### GetNfsAddresses

`func (o *Host) GetNfsAddresses() []string`

GetNfsAddresses returns the NfsAddresses field if non-nil, zero value otherwise.

### GetNfsAddressesOk

`func (o *Host) GetNfsAddressesOk() (*[]string, bool)`

GetNfsAddressesOk returns a tuple with the NfsAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAddresses

`func (o *Host) SetNfsAddresses(v []string)`

SetNfsAddresses sets NfsAddresses field to given value.

### HasNfsAddresses

`func (o *Host) HasNfsAddresses() bool`

HasNfsAddresses returns a boolean if a field has been set.

### GetDspKeystoreAlias

`func (o *Host) GetDspKeystoreAlias() string`

GetDspKeystoreAlias returns the DspKeystoreAlias field if non-nil, zero value otherwise.

### GetDspKeystoreAliasOk

`func (o *Host) GetDspKeystoreAliasOk() (*string, bool)`

GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystoreAlias

`func (o *Host) SetDspKeystoreAlias(v string)`

SetDspKeystoreAlias sets DspKeystoreAlias field to given value.

### HasDspKeystoreAlias

`func (o *Host) HasDspKeystoreAlias() bool`

HasDspKeystoreAlias returns a boolean if a field has been set.

### GetDspKeystorePath

`func (o *Host) GetDspKeystorePath() string`

GetDspKeystorePath returns the DspKeystorePath field if non-nil, zero value otherwise.

### GetDspKeystorePathOk

`func (o *Host) GetDspKeystorePathOk() (*string, bool)`

GetDspKeystorePathOk returns a tuple with the DspKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePath

`func (o *Host) SetDspKeystorePath(v string)`

SetDspKeystorePath sets DspKeystorePath field to given value.

### HasDspKeystorePath

`func (o *Host) HasDspKeystorePath() bool`

HasDspKeystorePath returns a boolean if a field has been set.

### GetDspTruststorePath

`func (o *Host) GetDspTruststorePath() string`

GetDspTruststorePath returns the DspTruststorePath field if non-nil, zero value otherwise.

### GetDspTruststorePathOk

`func (o *Host) GetDspTruststorePathOk() (*string, bool)`

GetDspTruststorePathOk returns a tuple with the DspTruststorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePath

`func (o *Host) SetDspTruststorePath(v string)`

SetDspTruststorePath sets DspTruststorePath field to given value.

### HasDspTruststorePath

`func (o *Host) HasDspTruststorePath() bool`

HasDspTruststorePath returns a boolean if a field has been set.

### GetJavaHome

`func (o *Host) GetJavaHome() string`

GetJavaHome returns the JavaHome field if non-nil, zero value otherwise.

### GetJavaHomeOk

`func (o *Host) GetJavaHomeOk() (*string, bool)`

GetJavaHomeOk returns a tuple with the JavaHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaHome

`func (o *Host) SetJavaHome(v string)`

SetJavaHome sets JavaHome field to given value.

### HasJavaHome

`func (o *Host) HasJavaHome() bool`

HasJavaHome returns a boolean if a field has been set.

### GetSshPort

`func (o *Host) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *Host) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *Host) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *Host) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetToolkitPath

`func (o *Host) GetToolkitPath() string`

GetToolkitPath returns the ToolkitPath field if non-nil, zero value otherwise.

### GetToolkitPathOk

`func (o *Host) GetToolkitPathOk() (*string, bool)`

GetToolkitPathOk returns a tuple with the ToolkitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolkitPath

`func (o *Host) SetToolkitPath(v string)`

SetToolkitPath sets ToolkitPath field to given value.

### HasToolkitPath

`func (o *Host) HasToolkitPath() bool`

HasToolkitPath returns a boolean if a field has been set.

### GetOracleTdeKeystoresRootPath

`func (o *Host) GetOracleTdeKeystoresRootPath() string`

GetOracleTdeKeystoresRootPath returns the OracleTdeKeystoresRootPath field if non-nil, zero value otherwise.

### GetOracleTdeKeystoresRootPathOk

`func (o *Host) GetOracleTdeKeystoresRootPathOk() (*string, bool)`

GetOracleTdeKeystoresRootPathOk returns a tuple with the OracleTdeKeystoresRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTdeKeystoresRootPath

`func (o *Host) SetOracleTdeKeystoresRootPath(v string)`

SetOracleTdeKeystoresRootPath sets OracleTdeKeystoresRootPath field to given value.

### HasOracleTdeKeystoresRootPath

`func (o *Host) HasOracleTdeKeystoresRootPath() bool`

HasOracleTdeKeystoresRootPath returns a boolean if a field has been set.

### GetProcessorType

`func (o *Host) GetProcessorType() string`

GetProcessorType returns the ProcessorType field if non-nil, zero value otherwise.

### GetProcessorTypeOk

`func (o *Host) GetProcessorTypeOk() (*string, bool)`

GetProcessorTypeOk returns a tuple with the ProcessorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorType

`func (o *Host) SetProcessorType(v string)`

SetProcessorType sets ProcessorType field to given value.

### HasProcessorType

`func (o *Host) HasProcessorType() bool`

HasProcessorType returns a boolean if a field has been set.

### GetTimezone

`func (o *Host) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Host) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Host) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Host) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


