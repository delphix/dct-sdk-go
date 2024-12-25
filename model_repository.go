/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the Repository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Repository{}

// Repository struct for Repository
type Repository struct {
	// Entity id of the repository.
	Id *string `json:"id,omitempty"`
	// Name of the repository.
	Name *string `json:"name,omitempty"`
	// The database type of this repository.
	DatabaseType *string `json:"database_type,omitempty"`
	// Flag indicating whether the repository should be used for provisioning.
	AllowProvisioning *bool `json:"allow_provisioning,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix Engine for internal processing.
	IsStaging *bool `json:"is_staging,omitempty"`
	// The Oracle base where database binaries are located.
	OracleBase *string `json:"oracle_base,omitempty"`
	// Version of the repository.
	Version *string `json:"version,omitempty"`
	// 32 or 64 bits.
	Bits *int32 `json:"bits,omitempty"`
	// Group name of the user that owns the install.
	InstallGroup *string `json:"install_group,omitempty"`
	// User name of the user that owns the install.
	InstallUser *string `json:"install_user,omitempty"`
	// Flag indicating whether the install supports Oracle RAC.
	Rac *bool `json:"rac,omitempty"`
	// The network ports for connecting to the database instance.
	Ports []int64 `json:"ports,omitempty"`
	// The network port for connecting to the SQL Server instance.
	Port *int64 `json:"port,omitempty"`
	// Fully qualified name of the dump history file.
	DumpHistoryFile *string `json:"dump_history_file,omitempty"`
	// Database page size for the SAP ASE instance.
	PageSize *int64 `json:"page_size,omitempty"`
	// Account the database server instance is running as.
	Owner *string `json:"owner,omitempty"`
	// Directory path where the installation is located.
	InstallationPath *string `json:"installation_path,omitempty"`
	// This property determines if the full-text search and semantic search is installed or not.
	FulltextInstalled *bool `json:"fulltext_installed,omitempty"`
	// The internal version is tied to the data format of a database and is used to detect compatibility.
	InternalVersion *int64 `json:"internal_version,omitempty"`
	// MSSQL cluster instances name.
	MssqlClusterInstancesName []string `json:"mssql_cluster_instances_name,omitempty"`
	// MSSQL cluster instances version.
	MssqlClusterInstancesVersion []string `json:"mssql_cluster_instances_version,omitempty"`
	// Directory where the installation home is located.
	InstallationHome *string `json:"installation_home,omitempty"`
	// MSSQL failover cluster drive letter.
	DriveLetter []string `json:"drive_letter,omitempty"`
	// Flag indicating whether the repository was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
}

// NewRepository instantiates a new Repository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepository() *Repository {
	this := Repository{}
	return &this
}

// NewRepositoryWithDefaults instantiates a new Repository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryWithDefaults() *Repository {
	this := Repository{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Repository) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Repository) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Repository) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Repository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Repository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Repository) SetName(v string) {
	o.Name = &v
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise.
func (o *Repository) GetDatabaseType() string {
	if o == nil || IsNil(o.DatabaseType) {
		var ret string
		return ret
	}
	return *o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetDatabaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseType) {
		return nil, false
	}
	return o.DatabaseType, true
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *Repository) HasDatabaseType() bool {
	if o != nil && !IsNil(o.DatabaseType) {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given string and assigns it to the DatabaseType field.
func (o *Repository) SetDatabaseType(v string) {
	o.DatabaseType = &v
}

// GetAllowProvisioning returns the AllowProvisioning field value if set, zero value otherwise.
func (o *Repository) GetAllowProvisioning() bool {
	if o == nil || IsNil(o.AllowProvisioning) {
		var ret bool
		return ret
	}
	return *o.AllowProvisioning
}

// GetAllowProvisioningOk returns a tuple with the AllowProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetAllowProvisioningOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowProvisioning) {
		return nil, false
	}
	return o.AllowProvisioning, true
}

// HasAllowProvisioning returns a boolean if a field has been set.
func (o *Repository) HasAllowProvisioning() bool {
	if o != nil && !IsNil(o.AllowProvisioning) {
		return true
	}

	return false
}

// SetAllowProvisioning gets a reference to the given bool and assigns it to the AllowProvisioning field.
func (o *Repository) SetAllowProvisioning(v bool) {
	o.AllowProvisioning = &v
}

// GetIsStaging returns the IsStaging field value if set, zero value otherwise.
func (o *Repository) GetIsStaging() bool {
	if o == nil || IsNil(o.IsStaging) {
		var ret bool
		return ret
	}
	return *o.IsStaging
}

// GetIsStagingOk returns a tuple with the IsStaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetIsStagingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStaging) {
		return nil, false
	}
	return o.IsStaging, true
}

// HasIsStaging returns a boolean if a field has been set.
func (o *Repository) HasIsStaging() bool {
	if o != nil && !IsNil(o.IsStaging) {
		return true
	}

	return false
}

// SetIsStaging gets a reference to the given bool and assigns it to the IsStaging field.
func (o *Repository) SetIsStaging(v bool) {
	o.IsStaging = &v
}

// GetOracleBase returns the OracleBase field value if set, zero value otherwise.
func (o *Repository) GetOracleBase() string {
	if o == nil || IsNil(o.OracleBase) {
		var ret string
		return ret
	}
	return *o.OracleBase
}

// GetOracleBaseOk returns a tuple with the OracleBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetOracleBaseOk() (*string, bool) {
	if o == nil || IsNil(o.OracleBase) {
		return nil, false
	}
	return o.OracleBase, true
}

// HasOracleBase returns a boolean if a field has been set.
func (o *Repository) HasOracleBase() bool {
	if o != nil && !IsNil(o.OracleBase) {
		return true
	}

	return false
}

// SetOracleBase gets a reference to the given string and assigns it to the OracleBase field.
func (o *Repository) SetOracleBase(v string) {
	o.OracleBase = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Repository) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Repository) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Repository) SetVersion(v string) {
	o.Version = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *Repository) GetBits() int32 {
	if o == nil || IsNil(o.Bits) {
		var ret int32
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetBitsOk() (*int32, bool) {
	if o == nil || IsNil(o.Bits) {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *Repository) HasBits() bool {
	if o != nil && !IsNil(o.Bits) {
		return true
	}

	return false
}

// SetBits gets a reference to the given int32 and assigns it to the Bits field.
func (o *Repository) SetBits(v int32) {
	o.Bits = &v
}

// GetInstallGroup returns the InstallGroup field value if set, zero value otherwise.
func (o *Repository) GetInstallGroup() string {
	if o == nil || IsNil(o.InstallGroup) {
		var ret string
		return ret
	}
	return *o.InstallGroup
}

// GetInstallGroupOk returns a tuple with the InstallGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetInstallGroupOk() (*string, bool) {
	if o == nil || IsNil(o.InstallGroup) {
		return nil, false
	}
	return o.InstallGroup, true
}

// HasInstallGroup returns a boolean if a field has been set.
func (o *Repository) HasInstallGroup() bool {
	if o != nil && !IsNil(o.InstallGroup) {
		return true
	}

	return false
}

// SetInstallGroup gets a reference to the given string and assigns it to the InstallGroup field.
func (o *Repository) SetInstallGroup(v string) {
	o.InstallGroup = &v
}

// GetInstallUser returns the InstallUser field value if set, zero value otherwise.
func (o *Repository) GetInstallUser() string {
	if o == nil || IsNil(o.InstallUser) {
		var ret string
		return ret
	}
	return *o.InstallUser
}

// GetInstallUserOk returns a tuple with the InstallUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetInstallUserOk() (*string, bool) {
	if o == nil || IsNil(o.InstallUser) {
		return nil, false
	}
	return o.InstallUser, true
}

// HasInstallUser returns a boolean if a field has been set.
func (o *Repository) HasInstallUser() bool {
	if o != nil && !IsNil(o.InstallUser) {
		return true
	}

	return false
}

// SetInstallUser gets a reference to the given string and assigns it to the InstallUser field.
func (o *Repository) SetInstallUser(v string) {
	o.InstallUser = &v
}

// GetRac returns the Rac field value if set, zero value otherwise.
func (o *Repository) GetRac() bool {
	if o == nil || IsNil(o.Rac) {
		var ret bool
		return ret
	}
	return *o.Rac
}

// GetRacOk returns a tuple with the Rac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetRacOk() (*bool, bool) {
	if o == nil || IsNil(o.Rac) {
		return nil, false
	}
	return o.Rac, true
}

// HasRac returns a boolean if a field has been set.
func (o *Repository) HasRac() bool {
	if o != nil && !IsNil(o.Rac) {
		return true
	}

	return false
}

// SetRac gets a reference to the given bool and assigns it to the Rac field.
func (o *Repository) SetRac(v bool) {
	o.Rac = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *Repository) GetPorts() []int64 {
	if o == nil || IsNil(o.Ports) {
		var ret []int64
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetPortsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *Repository) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []int64 and assigns it to the Ports field.
func (o *Repository) SetPorts(v []int64) {
	o.Ports = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Repository) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Repository) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *Repository) SetPort(v int64) {
	o.Port = &v
}

// GetDumpHistoryFile returns the DumpHistoryFile field value if set, zero value otherwise.
func (o *Repository) GetDumpHistoryFile() string {
	if o == nil || IsNil(o.DumpHistoryFile) {
		var ret string
		return ret
	}
	return *o.DumpHistoryFile
}

// GetDumpHistoryFileOk returns a tuple with the DumpHistoryFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetDumpHistoryFileOk() (*string, bool) {
	if o == nil || IsNil(o.DumpHistoryFile) {
		return nil, false
	}
	return o.DumpHistoryFile, true
}

// HasDumpHistoryFile returns a boolean if a field has been set.
func (o *Repository) HasDumpHistoryFile() bool {
	if o != nil && !IsNil(o.DumpHistoryFile) {
		return true
	}

	return false
}

// SetDumpHistoryFile gets a reference to the given string and assigns it to the DumpHistoryFile field.
func (o *Repository) SetDumpHistoryFile(v string) {
	o.DumpHistoryFile = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *Repository) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Repository) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *Repository) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Repository) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Repository) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Repository) SetOwner(v string) {
	o.Owner = &v
}

// GetInstallationPath returns the InstallationPath field value if set, zero value otherwise.
func (o *Repository) GetInstallationPath() string {
	if o == nil || IsNil(o.InstallationPath) {
		var ret string
		return ret
	}
	return *o.InstallationPath
}

// GetInstallationPathOk returns a tuple with the InstallationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetInstallationPathOk() (*string, bool) {
	if o == nil || IsNil(o.InstallationPath) {
		return nil, false
	}
	return o.InstallationPath, true
}

// HasInstallationPath returns a boolean if a field has been set.
func (o *Repository) HasInstallationPath() bool {
	if o != nil && !IsNil(o.InstallationPath) {
		return true
	}

	return false
}

// SetInstallationPath gets a reference to the given string and assigns it to the InstallationPath field.
func (o *Repository) SetInstallationPath(v string) {
	o.InstallationPath = &v
}

// GetFulltextInstalled returns the FulltextInstalled field value if set, zero value otherwise.
func (o *Repository) GetFulltextInstalled() bool {
	if o == nil || IsNil(o.FulltextInstalled) {
		var ret bool
		return ret
	}
	return *o.FulltextInstalled
}

// GetFulltextInstalledOk returns a tuple with the FulltextInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetFulltextInstalledOk() (*bool, bool) {
	if o == nil || IsNil(o.FulltextInstalled) {
		return nil, false
	}
	return o.FulltextInstalled, true
}

// HasFulltextInstalled returns a boolean if a field has been set.
func (o *Repository) HasFulltextInstalled() bool {
	if o != nil && !IsNil(o.FulltextInstalled) {
		return true
	}

	return false
}

// SetFulltextInstalled gets a reference to the given bool and assigns it to the FulltextInstalled field.
func (o *Repository) SetFulltextInstalled(v bool) {
	o.FulltextInstalled = &v
}

// GetInternalVersion returns the InternalVersion field value if set, zero value otherwise.
func (o *Repository) GetInternalVersion() int64 {
	if o == nil || IsNil(o.InternalVersion) {
		var ret int64
		return ret
	}
	return *o.InternalVersion
}

// GetInternalVersionOk returns a tuple with the InternalVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetInternalVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.InternalVersion) {
		return nil, false
	}
	return o.InternalVersion, true
}

// HasInternalVersion returns a boolean if a field has been set.
func (o *Repository) HasInternalVersion() bool {
	if o != nil && !IsNil(o.InternalVersion) {
		return true
	}

	return false
}

// SetInternalVersion gets a reference to the given int64 and assigns it to the InternalVersion field.
func (o *Repository) SetInternalVersion(v int64) {
	o.InternalVersion = &v
}

// GetMssqlClusterInstancesName returns the MssqlClusterInstancesName field value if set, zero value otherwise.
func (o *Repository) GetMssqlClusterInstancesName() []string {
	if o == nil || IsNil(o.MssqlClusterInstancesName) {
		var ret []string
		return ret
	}
	return o.MssqlClusterInstancesName
}

// GetMssqlClusterInstancesNameOk returns a tuple with the MssqlClusterInstancesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetMssqlClusterInstancesNameOk() ([]string, bool) {
	if o == nil || IsNil(o.MssqlClusterInstancesName) {
		return nil, false
	}
	return o.MssqlClusterInstancesName, true
}

// HasMssqlClusterInstancesName returns a boolean if a field has been set.
func (o *Repository) HasMssqlClusterInstancesName() bool {
	if o != nil && !IsNil(o.MssqlClusterInstancesName) {
		return true
	}

	return false
}

// SetMssqlClusterInstancesName gets a reference to the given []string and assigns it to the MssqlClusterInstancesName field.
func (o *Repository) SetMssqlClusterInstancesName(v []string) {
	o.MssqlClusterInstancesName = v
}

// GetMssqlClusterInstancesVersion returns the MssqlClusterInstancesVersion field value if set, zero value otherwise.
func (o *Repository) GetMssqlClusterInstancesVersion() []string {
	if o == nil || IsNil(o.MssqlClusterInstancesVersion) {
		var ret []string
		return ret
	}
	return o.MssqlClusterInstancesVersion
}

// GetMssqlClusterInstancesVersionOk returns a tuple with the MssqlClusterInstancesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetMssqlClusterInstancesVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.MssqlClusterInstancesVersion) {
		return nil, false
	}
	return o.MssqlClusterInstancesVersion, true
}

// HasMssqlClusterInstancesVersion returns a boolean if a field has been set.
func (o *Repository) HasMssqlClusterInstancesVersion() bool {
	if o != nil && !IsNil(o.MssqlClusterInstancesVersion) {
		return true
	}

	return false
}

// SetMssqlClusterInstancesVersion gets a reference to the given []string and assigns it to the MssqlClusterInstancesVersion field.
func (o *Repository) SetMssqlClusterInstancesVersion(v []string) {
	o.MssqlClusterInstancesVersion = v
}

// GetInstallationHome returns the InstallationHome field value if set, zero value otherwise.
func (o *Repository) GetInstallationHome() string {
	if o == nil || IsNil(o.InstallationHome) {
		var ret string
		return ret
	}
	return *o.InstallationHome
}

// GetInstallationHomeOk returns a tuple with the InstallationHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetInstallationHomeOk() (*string, bool) {
	if o == nil || IsNil(o.InstallationHome) {
		return nil, false
	}
	return o.InstallationHome, true
}

// HasInstallationHome returns a boolean if a field has been set.
func (o *Repository) HasInstallationHome() bool {
	if o != nil && !IsNil(o.InstallationHome) {
		return true
	}

	return false
}

// SetInstallationHome gets a reference to the given string and assigns it to the InstallationHome field.
func (o *Repository) SetInstallationHome(v string) {
	o.InstallationHome = &v
}

// GetDriveLetter returns the DriveLetter field value if set, zero value otherwise.
func (o *Repository) GetDriveLetter() []string {
	if o == nil || IsNil(o.DriveLetter) {
		var ret []string
		return ret
	}
	return o.DriveLetter
}

// GetDriveLetterOk returns a tuple with the DriveLetter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetDriveLetterOk() ([]string, bool) {
	if o == nil || IsNil(o.DriveLetter) {
		return nil, false
	}
	return o.DriveLetter, true
}

// HasDriveLetter returns a boolean if a field has been set.
func (o *Repository) HasDriveLetter() bool {
	if o != nil && !IsNil(o.DriveLetter) {
		return true
	}

	return false
}

// SetDriveLetter gets a reference to the given []string and assigns it to the DriveLetter field.
func (o *Repository) SetDriveLetter(v []string) {
	o.DriveLetter = v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *Repository) GetDiscovered() bool {
	if o == nil || IsNil(o.Discovered) {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetDiscoveredOk() (*bool, bool) {
	if o == nil || IsNil(o.Discovered) {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *Repository) HasDiscovered() bool {
	if o != nil && !IsNil(o.Discovered) {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *Repository) SetDiscovered(v bool) {
	o.Discovered = &v
}

func (o Repository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Repository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DatabaseType) {
		toSerialize["database_type"] = o.DatabaseType
	}
	if !IsNil(o.AllowProvisioning) {
		toSerialize["allow_provisioning"] = o.AllowProvisioning
	}
	if !IsNil(o.IsStaging) {
		toSerialize["is_staging"] = o.IsStaging
	}
	if !IsNil(o.OracleBase) {
		toSerialize["oracle_base"] = o.OracleBase
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Bits) {
		toSerialize["bits"] = o.Bits
	}
	if !IsNil(o.InstallGroup) {
		toSerialize["install_group"] = o.InstallGroup
	}
	if !IsNil(o.InstallUser) {
		toSerialize["install_user"] = o.InstallUser
	}
	if !IsNil(o.Rac) {
		toSerialize["rac"] = o.Rac
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.DumpHistoryFile) {
		toSerialize["dump_history_file"] = o.DumpHistoryFile
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.InstallationPath) {
		toSerialize["installation_path"] = o.InstallationPath
	}
	if !IsNil(o.FulltextInstalled) {
		toSerialize["fulltext_installed"] = o.FulltextInstalled
	}
	if !IsNil(o.InternalVersion) {
		toSerialize["internal_version"] = o.InternalVersion
	}
	if !IsNil(o.MssqlClusterInstancesName) {
		toSerialize["mssql_cluster_instances_name"] = o.MssqlClusterInstancesName
	}
	if !IsNil(o.MssqlClusterInstancesVersion) {
		toSerialize["mssql_cluster_instances_version"] = o.MssqlClusterInstancesVersion
	}
	if !IsNil(o.InstallationHome) {
		toSerialize["installation_home"] = o.InstallationHome
	}
	if !IsNil(o.DriveLetter) {
		toSerialize["drive_letter"] = o.DriveLetter
	}
	if !IsNil(o.Discovered) {
		toSerialize["discovered"] = o.Discovered
	}
	return toSerialize, nil
}

type NullableRepository struct {
	value *Repository
	isSet bool
}

func (v NullableRepository) Get() *Repository {
	return v.value
}

func (v *NullableRepository) Set(val *Repository) {
	v.value = val
	v.isSet = true
}

func (v NullableRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepository(val *Repository) *NullableRepository {
	return &NullableRepository{value: val, isSet: true}
}

func (v NullableRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


