/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"time"
)

// checks if the Host type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Host{}

// Host A physical/virtual server.
type Host struct {
	// The entity ID of this Host.
	Id *string `json:"id,omitempty"`
	// The hostname or IP address of this host.
	Hostname *string `json:"hostname,omitempty"`
	// The name of the OS on this host.
	OsName *string `json:"os_name,omitempty"`
	// The version of the OS on this host.
	OsVersion *string `json:"os_version,omitempty"`
	// The total amount of memory on this host in bytes.
	MemorySize *int64 `json:"memory_size,omitempty"`
	// True if the host is up and a connection can be established from the engine.
	Available *bool `json:"available,omitempty"`
	// The last time the available property was updated.
	AvailableTimestamp *time.Time `json:"available_timestamp,omitempty"`
	// The reason why the host is not available.
	NotAvailableReason *string `json:"not_available_reason,omitempty"`
	// The reference to the associated OracleClusterNode.
	OracleClusterNodeReference *string `json:"oracle_cluster_node_reference,omitempty"`
	// The name of the associated OracleClusterNode.
	OracleClusterNodeName *string `json:"oracle_cluster_node_name,omitempty"`
	// Whether the associated OracleClusterNode is enabled.
	OracleClusterNodeEnabled *bool `json:"oracle_cluster_node_enabled,omitempty"`
	// Whether the associated OracleClusterNode was discovered.
	OracleClusterNodeDiscovered *bool `json:"oracle_cluster_node_discovered,omitempty"`
	// The Virtual IP addresses associated with the OracleClusterNode.
	OracleClusterNodeVirtualIps []OracleVirtualIP `json:"oracle_cluster_node_virtual_ips,omitempty"`
	// The reference to the associated WindowsClusterNode.
	WindowsClusterNodeReference *string `json:"windows_cluster_node_reference,omitempty"`
	// The name of the associated WindowsClusterNode.
	WindowsClusterNodeName *string `json:"windows_cluster_node_name,omitempty"`
	// Whether the associated Windows cluster node was discovered.
	WindowsClusterNodeDiscovered *bool `json:"windows_cluster_node_discovered,omitempty"`
	// The list of host/IP addresses to use for NFS export.
	NfsAddresses []string `json:"nfs_addresses,omitempty"`
	// The lowercase alias to use inside the user managed DSP keystore.
	DspKeystoreAlias *string `json:"dsp_keystore_alias,omitempty"`
	// The path to the user managed DSP keystore.
	DspKeystorePath *string `json:"dsp_keystore_path,omitempty"`
	// The path to the user managed DSP truststore.
	DspTruststorePath *string `json:"dsp_truststore_path,omitempty"`
	// The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used.
	JavaHome *string `json:"java_home,omitempty"`
	// The port number used to connect to the host via SSH.
	SshPort *int32 `json:"ssh_port,omitempty"`
	// The path for the toolkit that resides on the host.
	ToolkitPath *string `json:"toolkit_path,omitempty"`
	// The path to the root of the Oracle TDE keystores artifact directories.
	OracleTdeKeystoresRootPath *string `json:"oracle_tde_keystores_root_path,omitempty"`
	// The platform for the host machine.
	ProcessorType *string `json:"processor_type,omitempty"`
	// The OS timezone.
	Timezone *string `json:"timezone,omitempty"`
}

// NewHost instantiates a new Host object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHost() *Host {
	this := Host{}
	return &this
}

// NewHostWithDefaults instantiates a new Host object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostWithDefaults() *Host {
	this := Host{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Host) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Host) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Host) SetId(v string) {
	o.Id = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *Host) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *Host) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *Host) SetHostname(v string) {
	o.Hostname = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *Host) GetOsName() string {
	if o == nil || IsNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOsNameOk() (*string, bool) {
	if o == nil || IsNil(o.OsName) {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *Host) HasOsName() bool {
	if o != nil && !IsNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *Host) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *Host) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *Host) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *Host) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *Host) GetMemorySize() int64 {
	if o == nil || IsNil(o.MemorySize) {
		var ret int64
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetMemorySizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MemorySize) {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *Host) HasMemorySize() bool {
	if o != nil && !IsNil(o.MemorySize) {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given int64 and assigns it to the MemorySize field.
func (o *Host) SetMemorySize(v int64) {
	o.MemorySize = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *Host) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *Host) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *Host) SetAvailable(v bool) {
	o.Available = &v
}

// GetAvailableTimestamp returns the AvailableTimestamp field value if set, zero value otherwise.
func (o *Host) GetAvailableTimestamp() time.Time {
	if o == nil || IsNil(o.AvailableTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.AvailableTimestamp
}

// GetAvailableTimestampOk returns a tuple with the AvailableTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetAvailableTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AvailableTimestamp) {
		return nil, false
	}
	return o.AvailableTimestamp, true
}

// HasAvailableTimestamp returns a boolean if a field has been set.
func (o *Host) HasAvailableTimestamp() bool {
	if o != nil && !IsNil(o.AvailableTimestamp) {
		return true
	}

	return false
}

// SetAvailableTimestamp gets a reference to the given time.Time and assigns it to the AvailableTimestamp field.
func (o *Host) SetAvailableTimestamp(v time.Time) {
	o.AvailableTimestamp = &v
}

// GetNotAvailableReason returns the NotAvailableReason field value if set, zero value otherwise.
func (o *Host) GetNotAvailableReason() string {
	if o == nil || IsNil(o.NotAvailableReason) {
		var ret string
		return ret
	}
	return *o.NotAvailableReason
}

// GetNotAvailableReasonOk returns a tuple with the NotAvailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetNotAvailableReasonOk() (*string, bool) {
	if o == nil || IsNil(o.NotAvailableReason) {
		return nil, false
	}
	return o.NotAvailableReason, true
}

// HasNotAvailableReason returns a boolean if a field has been set.
func (o *Host) HasNotAvailableReason() bool {
	if o != nil && !IsNil(o.NotAvailableReason) {
		return true
	}

	return false
}

// SetNotAvailableReason gets a reference to the given string and assigns it to the NotAvailableReason field.
func (o *Host) SetNotAvailableReason(v string) {
	o.NotAvailableReason = &v
}

// GetOracleClusterNodeReference returns the OracleClusterNodeReference field value if set, zero value otherwise.
func (o *Host) GetOracleClusterNodeReference() string {
	if o == nil || IsNil(o.OracleClusterNodeReference) {
		var ret string
		return ret
	}
	return *o.OracleClusterNodeReference
}

// GetOracleClusterNodeReferenceOk returns a tuple with the OracleClusterNodeReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOracleClusterNodeReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.OracleClusterNodeReference) {
		return nil, false
	}
	return o.OracleClusterNodeReference, true
}

// HasOracleClusterNodeReference returns a boolean if a field has been set.
func (o *Host) HasOracleClusterNodeReference() bool {
	if o != nil && !IsNil(o.OracleClusterNodeReference) {
		return true
	}

	return false
}

// SetOracleClusterNodeReference gets a reference to the given string and assigns it to the OracleClusterNodeReference field.
func (o *Host) SetOracleClusterNodeReference(v string) {
	o.OracleClusterNodeReference = &v
}

// GetOracleClusterNodeName returns the OracleClusterNodeName field value if set, zero value otherwise.
func (o *Host) GetOracleClusterNodeName() string {
	if o == nil || IsNil(o.OracleClusterNodeName) {
		var ret string
		return ret
	}
	return *o.OracleClusterNodeName
}

// GetOracleClusterNodeNameOk returns a tuple with the OracleClusterNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOracleClusterNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.OracleClusterNodeName) {
		return nil, false
	}
	return o.OracleClusterNodeName, true
}

// HasOracleClusterNodeName returns a boolean if a field has been set.
func (o *Host) HasOracleClusterNodeName() bool {
	if o != nil && !IsNil(o.OracleClusterNodeName) {
		return true
	}

	return false
}

// SetOracleClusterNodeName gets a reference to the given string and assigns it to the OracleClusterNodeName field.
func (o *Host) SetOracleClusterNodeName(v string) {
	o.OracleClusterNodeName = &v
}

// GetOracleClusterNodeEnabled returns the OracleClusterNodeEnabled field value if set, zero value otherwise.
func (o *Host) GetOracleClusterNodeEnabled() bool {
	if o == nil || IsNil(o.OracleClusterNodeEnabled) {
		var ret bool
		return ret
	}
	return *o.OracleClusterNodeEnabled
}

// GetOracleClusterNodeEnabledOk returns a tuple with the OracleClusterNodeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOracleClusterNodeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OracleClusterNodeEnabled) {
		return nil, false
	}
	return o.OracleClusterNodeEnabled, true
}

// HasOracleClusterNodeEnabled returns a boolean if a field has been set.
func (o *Host) HasOracleClusterNodeEnabled() bool {
	if o != nil && !IsNil(o.OracleClusterNodeEnabled) {
		return true
	}

	return false
}

// SetOracleClusterNodeEnabled gets a reference to the given bool and assigns it to the OracleClusterNodeEnabled field.
func (o *Host) SetOracleClusterNodeEnabled(v bool) {
	o.OracleClusterNodeEnabled = &v
}

// GetOracleClusterNodeDiscovered returns the OracleClusterNodeDiscovered field value if set, zero value otherwise.
func (o *Host) GetOracleClusterNodeDiscovered() bool {
	if o == nil || IsNil(o.OracleClusterNodeDiscovered) {
		var ret bool
		return ret
	}
	return *o.OracleClusterNodeDiscovered
}

// GetOracleClusterNodeDiscoveredOk returns a tuple with the OracleClusterNodeDiscovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOracleClusterNodeDiscoveredOk() (*bool, bool) {
	if o == nil || IsNil(o.OracleClusterNodeDiscovered) {
		return nil, false
	}
	return o.OracleClusterNodeDiscovered, true
}

// HasOracleClusterNodeDiscovered returns a boolean if a field has been set.
func (o *Host) HasOracleClusterNodeDiscovered() bool {
	if o != nil && !IsNil(o.OracleClusterNodeDiscovered) {
		return true
	}

	return false
}

// SetOracleClusterNodeDiscovered gets a reference to the given bool and assigns it to the OracleClusterNodeDiscovered field.
func (o *Host) SetOracleClusterNodeDiscovered(v bool) {
	o.OracleClusterNodeDiscovered = &v
}

// GetOracleClusterNodeVirtualIps returns the OracleClusterNodeVirtualIps field value if set, zero value otherwise.
func (o *Host) GetOracleClusterNodeVirtualIps() []OracleVirtualIP {
	if o == nil || IsNil(o.OracleClusterNodeVirtualIps) {
		var ret []OracleVirtualIP
		return ret
	}
	return o.OracleClusterNodeVirtualIps
}

// GetOracleClusterNodeVirtualIpsOk returns a tuple with the OracleClusterNodeVirtualIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOracleClusterNodeVirtualIpsOk() ([]OracleVirtualIP, bool) {
	if o == nil || IsNil(o.OracleClusterNodeVirtualIps) {
		return nil, false
	}
	return o.OracleClusterNodeVirtualIps, true
}

// HasOracleClusterNodeVirtualIps returns a boolean if a field has been set.
func (o *Host) HasOracleClusterNodeVirtualIps() bool {
	if o != nil && !IsNil(o.OracleClusterNodeVirtualIps) {
		return true
	}

	return false
}

// SetOracleClusterNodeVirtualIps gets a reference to the given []OracleVirtualIP and assigns it to the OracleClusterNodeVirtualIps field.
func (o *Host) SetOracleClusterNodeVirtualIps(v []OracleVirtualIP) {
	o.OracleClusterNodeVirtualIps = v
}

// GetWindowsClusterNodeReference returns the WindowsClusterNodeReference field value if set, zero value otherwise.
func (o *Host) GetWindowsClusterNodeReference() string {
	if o == nil || IsNil(o.WindowsClusterNodeReference) {
		var ret string
		return ret
	}
	return *o.WindowsClusterNodeReference
}

// GetWindowsClusterNodeReferenceOk returns a tuple with the WindowsClusterNodeReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetWindowsClusterNodeReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsClusterNodeReference) {
		return nil, false
	}
	return o.WindowsClusterNodeReference, true
}

// HasWindowsClusterNodeReference returns a boolean if a field has been set.
func (o *Host) HasWindowsClusterNodeReference() bool {
	if o != nil && !IsNil(o.WindowsClusterNodeReference) {
		return true
	}

	return false
}

// SetWindowsClusterNodeReference gets a reference to the given string and assigns it to the WindowsClusterNodeReference field.
func (o *Host) SetWindowsClusterNodeReference(v string) {
	o.WindowsClusterNodeReference = &v
}

// GetWindowsClusterNodeName returns the WindowsClusterNodeName field value if set, zero value otherwise.
func (o *Host) GetWindowsClusterNodeName() string {
	if o == nil || IsNil(o.WindowsClusterNodeName) {
		var ret string
		return ret
	}
	return *o.WindowsClusterNodeName
}

// GetWindowsClusterNodeNameOk returns a tuple with the WindowsClusterNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetWindowsClusterNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsClusterNodeName) {
		return nil, false
	}
	return o.WindowsClusterNodeName, true
}

// HasWindowsClusterNodeName returns a boolean if a field has been set.
func (o *Host) HasWindowsClusterNodeName() bool {
	if o != nil && !IsNil(o.WindowsClusterNodeName) {
		return true
	}

	return false
}

// SetWindowsClusterNodeName gets a reference to the given string and assigns it to the WindowsClusterNodeName field.
func (o *Host) SetWindowsClusterNodeName(v string) {
	o.WindowsClusterNodeName = &v
}

// GetWindowsClusterNodeDiscovered returns the WindowsClusterNodeDiscovered field value if set, zero value otherwise.
func (o *Host) GetWindowsClusterNodeDiscovered() bool {
	if o == nil || IsNil(o.WindowsClusterNodeDiscovered) {
		var ret bool
		return ret
	}
	return *o.WindowsClusterNodeDiscovered
}

// GetWindowsClusterNodeDiscoveredOk returns a tuple with the WindowsClusterNodeDiscovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetWindowsClusterNodeDiscoveredOk() (*bool, bool) {
	if o == nil || IsNil(o.WindowsClusterNodeDiscovered) {
		return nil, false
	}
	return o.WindowsClusterNodeDiscovered, true
}

// HasWindowsClusterNodeDiscovered returns a boolean if a field has been set.
func (o *Host) HasWindowsClusterNodeDiscovered() bool {
	if o != nil && !IsNil(o.WindowsClusterNodeDiscovered) {
		return true
	}

	return false
}

// SetWindowsClusterNodeDiscovered gets a reference to the given bool and assigns it to the WindowsClusterNodeDiscovered field.
func (o *Host) SetWindowsClusterNodeDiscovered(v bool) {
	o.WindowsClusterNodeDiscovered = &v
}

// GetNfsAddresses returns the NfsAddresses field value if set, zero value otherwise.
func (o *Host) GetNfsAddresses() []string {
	if o == nil || IsNil(o.NfsAddresses) {
		var ret []string
		return ret
	}
	return o.NfsAddresses
}

// GetNfsAddressesOk returns a tuple with the NfsAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetNfsAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.NfsAddresses) {
		return nil, false
	}
	return o.NfsAddresses, true
}

// HasNfsAddresses returns a boolean if a field has been set.
func (o *Host) HasNfsAddresses() bool {
	if o != nil && !IsNil(o.NfsAddresses) {
		return true
	}

	return false
}

// SetNfsAddresses gets a reference to the given []string and assigns it to the NfsAddresses field.
func (o *Host) SetNfsAddresses(v []string) {
	o.NfsAddresses = v
}

// GetDspKeystoreAlias returns the DspKeystoreAlias field value if set, zero value otherwise.
func (o *Host) GetDspKeystoreAlias() string {
	if o == nil || IsNil(o.DspKeystoreAlias) {
		var ret string
		return ret
	}
	return *o.DspKeystoreAlias
}

// GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetDspKeystoreAliasOk() (*string, bool) {
	if o == nil || IsNil(o.DspKeystoreAlias) {
		return nil, false
	}
	return o.DspKeystoreAlias, true
}

// HasDspKeystoreAlias returns a boolean if a field has been set.
func (o *Host) HasDspKeystoreAlias() bool {
	if o != nil && !IsNil(o.DspKeystoreAlias) {
		return true
	}

	return false
}

// SetDspKeystoreAlias gets a reference to the given string and assigns it to the DspKeystoreAlias field.
func (o *Host) SetDspKeystoreAlias(v string) {
	o.DspKeystoreAlias = &v
}

// GetDspKeystorePath returns the DspKeystorePath field value if set, zero value otherwise.
func (o *Host) GetDspKeystorePath() string {
	if o == nil || IsNil(o.DspKeystorePath) {
		var ret string
		return ret
	}
	return *o.DspKeystorePath
}

// GetDspKeystorePathOk returns a tuple with the DspKeystorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetDspKeystorePathOk() (*string, bool) {
	if o == nil || IsNil(o.DspKeystorePath) {
		return nil, false
	}
	return o.DspKeystorePath, true
}

// HasDspKeystorePath returns a boolean if a field has been set.
func (o *Host) HasDspKeystorePath() bool {
	if o != nil && !IsNil(o.DspKeystorePath) {
		return true
	}

	return false
}

// SetDspKeystorePath gets a reference to the given string and assigns it to the DspKeystorePath field.
func (o *Host) SetDspKeystorePath(v string) {
	o.DspKeystorePath = &v
}

// GetDspTruststorePath returns the DspTruststorePath field value if set, zero value otherwise.
func (o *Host) GetDspTruststorePath() string {
	if o == nil || IsNil(o.DspTruststorePath) {
		var ret string
		return ret
	}
	return *o.DspTruststorePath
}

// GetDspTruststorePathOk returns a tuple with the DspTruststorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetDspTruststorePathOk() (*string, bool) {
	if o == nil || IsNil(o.DspTruststorePath) {
		return nil, false
	}
	return o.DspTruststorePath, true
}

// HasDspTruststorePath returns a boolean if a field has been set.
func (o *Host) HasDspTruststorePath() bool {
	if o != nil && !IsNil(o.DspTruststorePath) {
		return true
	}

	return false
}

// SetDspTruststorePath gets a reference to the given string and assigns it to the DspTruststorePath field.
func (o *Host) SetDspTruststorePath(v string) {
	o.DspTruststorePath = &v
}

// GetJavaHome returns the JavaHome field value if set, zero value otherwise.
func (o *Host) GetJavaHome() string {
	if o == nil || IsNil(o.JavaHome) {
		var ret string
		return ret
	}
	return *o.JavaHome
}

// GetJavaHomeOk returns a tuple with the JavaHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetJavaHomeOk() (*string, bool) {
	if o == nil || IsNil(o.JavaHome) {
		return nil, false
	}
	return o.JavaHome, true
}

// HasJavaHome returns a boolean if a field has been set.
func (o *Host) HasJavaHome() bool {
	if o != nil && !IsNil(o.JavaHome) {
		return true
	}

	return false
}

// SetJavaHome gets a reference to the given string and assigns it to the JavaHome field.
func (o *Host) SetJavaHome(v string) {
	o.JavaHome = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *Host) GetSshPort() int32 {
	if o == nil || IsNil(o.SshPort) {
		var ret int32
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetSshPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SshPort) {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *Host) HasSshPort() bool {
	if o != nil && !IsNil(o.SshPort) {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int32 and assigns it to the SshPort field.
func (o *Host) SetSshPort(v int32) {
	o.SshPort = &v
}

// GetToolkitPath returns the ToolkitPath field value if set, zero value otherwise.
func (o *Host) GetToolkitPath() string {
	if o == nil || IsNil(o.ToolkitPath) {
		var ret string
		return ret
	}
	return *o.ToolkitPath
}

// GetToolkitPathOk returns a tuple with the ToolkitPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetToolkitPathOk() (*string, bool) {
	if o == nil || IsNil(o.ToolkitPath) {
		return nil, false
	}
	return o.ToolkitPath, true
}

// HasToolkitPath returns a boolean if a field has been set.
func (o *Host) HasToolkitPath() bool {
	if o != nil && !IsNil(o.ToolkitPath) {
		return true
	}

	return false
}

// SetToolkitPath gets a reference to the given string and assigns it to the ToolkitPath field.
func (o *Host) SetToolkitPath(v string) {
	o.ToolkitPath = &v
}

// GetOracleTdeKeystoresRootPath returns the OracleTdeKeystoresRootPath field value if set, zero value otherwise.
func (o *Host) GetOracleTdeKeystoresRootPath() string {
	if o == nil || IsNil(o.OracleTdeKeystoresRootPath) {
		var ret string
		return ret
	}
	return *o.OracleTdeKeystoresRootPath
}

// GetOracleTdeKeystoresRootPathOk returns a tuple with the OracleTdeKeystoresRootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOracleTdeKeystoresRootPathOk() (*string, bool) {
	if o == nil || IsNil(o.OracleTdeKeystoresRootPath) {
		return nil, false
	}
	return o.OracleTdeKeystoresRootPath, true
}

// HasOracleTdeKeystoresRootPath returns a boolean if a field has been set.
func (o *Host) HasOracleTdeKeystoresRootPath() bool {
	if o != nil && !IsNil(o.OracleTdeKeystoresRootPath) {
		return true
	}

	return false
}

// SetOracleTdeKeystoresRootPath gets a reference to the given string and assigns it to the OracleTdeKeystoresRootPath field.
func (o *Host) SetOracleTdeKeystoresRootPath(v string) {
	o.OracleTdeKeystoresRootPath = &v
}

// GetProcessorType returns the ProcessorType field value if set, zero value otherwise.
func (o *Host) GetProcessorType() string {
	if o == nil || IsNil(o.ProcessorType) {
		var ret string
		return ret
	}
	return *o.ProcessorType
}

// GetProcessorTypeOk returns a tuple with the ProcessorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetProcessorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorType) {
		return nil, false
	}
	return o.ProcessorType, true
}

// HasProcessorType returns a boolean if a field has been set.
func (o *Host) HasProcessorType() bool {
	if o != nil && !IsNil(o.ProcessorType) {
		return true
	}

	return false
}

// SetProcessorType gets a reference to the given string and assigns it to the ProcessorType field.
func (o *Host) SetProcessorType(v string) {
	o.ProcessorType = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Host) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Host) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Host) SetTimezone(v string) {
	o.Timezone = &v
}

func (o Host) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Host) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.OsName) {
		toSerialize["os_name"] = o.OsName
	}
	if !IsNil(o.OsVersion) {
		toSerialize["os_version"] = o.OsVersion
	}
	if !IsNil(o.MemorySize) {
		toSerialize["memory_size"] = o.MemorySize
	}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.AvailableTimestamp) {
		toSerialize["available_timestamp"] = o.AvailableTimestamp
	}
	if !IsNil(o.NotAvailableReason) {
		toSerialize["not_available_reason"] = o.NotAvailableReason
	}
	if !IsNil(o.OracleClusterNodeReference) {
		toSerialize["oracle_cluster_node_reference"] = o.OracleClusterNodeReference
	}
	if !IsNil(o.OracleClusterNodeName) {
		toSerialize["oracle_cluster_node_name"] = o.OracleClusterNodeName
	}
	if !IsNil(o.OracleClusterNodeEnabled) {
		toSerialize["oracle_cluster_node_enabled"] = o.OracleClusterNodeEnabled
	}
	if !IsNil(o.OracleClusterNodeDiscovered) {
		toSerialize["oracle_cluster_node_discovered"] = o.OracleClusterNodeDiscovered
	}
	if !IsNil(o.OracleClusterNodeVirtualIps) {
		toSerialize["oracle_cluster_node_virtual_ips"] = o.OracleClusterNodeVirtualIps
	}
	if !IsNil(o.WindowsClusterNodeReference) {
		toSerialize["windows_cluster_node_reference"] = o.WindowsClusterNodeReference
	}
	if !IsNil(o.WindowsClusterNodeName) {
		toSerialize["windows_cluster_node_name"] = o.WindowsClusterNodeName
	}
	if !IsNil(o.WindowsClusterNodeDiscovered) {
		toSerialize["windows_cluster_node_discovered"] = o.WindowsClusterNodeDiscovered
	}
	if !IsNil(o.NfsAddresses) {
		toSerialize["nfs_addresses"] = o.NfsAddresses
	}
	if !IsNil(o.DspKeystoreAlias) {
		toSerialize["dsp_keystore_alias"] = o.DspKeystoreAlias
	}
	if !IsNil(o.DspKeystorePath) {
		toSerialize["dsp_keystore_path"] = o.DspKeystorePath
	}
	if !IsNil(o.DspTruststorePath) {
		toSerialize["dsp_truststore_path"] = o.DspTruststorePath
	}
	if !IsNil(o.JavaHome) {
		toSerialize["java_home"] = o.JavaHome
	}
	if !IsNil(o.SshPort) {
		toSerialize["ssh_port"] = o.SshPort
	}
	if !IsNil(o.ToolkitPath) {
		toSerialize["toolkit_path"] = o.ToolkitPath
	}
	if !IsNil(o.OracleTdeKeystoresRootPath) {
		toSerialize["oracle_tde_keystores_root_path"] = o.OracleTdeKeystoresRootPath
	}
	if !IsNil(o.ProcessorType) {
		toSerialize["processor_type"] = o.ProcessorType
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

type NullableHost struct {
	value *Host
	isSet bool
}

func (v NullableHost) Get() *Host {
	return v.value
}

func (v *NullableHost) Set(val *Host) {
	v.value = val
	v.isSet = true
}

func (v NullableHost) IsSet() bool {
	return v.isSet
}

func (v *NullableHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHost(val *Host) *NullableHost {
	return &NullableHost{value: val, isSet: true}
}

func (v NullableHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


