# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Entity id of the repository. | [optional] 
**Name** | Pointer to **string** | Name of the repository. | [optional] 
**DatabaseType** | Pointer to **string** | The database type of this repository. | [optional] 
**AllowProvisioning** | Pointer to **bool** | Flag indicating whether the repository should be used for provisioning. | [optional] 
**IsStaging** | Pointer to **bool** | Flag indicating whether this repository can be used by the Delphix Engine for internal processing. | [optional] 
**OracleBase** | Pointer to **string** | The Oracle base where database binaries are located. | [optional] 
**Version** | Pointer to **string** | Version of the repository. | [optional] 
**Bits** | Pointer to **int32** | 32 or 64 bits. | [optional] 
**InstallGroup** | Pointer to **string** | Group name of the user that owns the install. | [optional] 
**InstallUser** | Pointer to **string** | User name of the user that owns the install. | [optional] 
**Rac** | Pointer to **bool** | Flag indicating whether the install supports Oracle RAC. | [optional] 
**Ports** | Pointer to **[]int64** | The network ports for connecting to the database instance. | [optional] 
**DumpHistoryFile** | Pointer to **string** | Fully qualified name of the dump history file. | [optional] 
**PageSize** | Pointer to **int64** | Database page size for the SAP ASE instance. | [optional] 
**Owner** | Pointer to **string** | Account the database server instance is running as. | [optional] 
**InstallationPath** | Pointer to **string** | Directory path where the installation is located. | [optional] 
**FulltextInstalled** | Pointer to **bool** | This property determines if the full-text search and semantic search is installed or not. | [optional] 
**InternalVersion** | Pointer to **int64** | The internal version is tied to the data format of a database and is used to detect compatibility. | [optional] 
**MssqlClusterInstancesName** | Pointer to **[]string** | MSSQL cluster instances name. | [optional] 
**MssqlClusterInstancesVersion** | Pointer to **[]string** | MSSQL cluster instances version. | [optional] 
**InstallationHome** | Pointer to **string** | Directory where the installation home is located. | [optional] 

## Methods

### NewRepository

`func NewRepository() *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Repository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Repository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Repository) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Repository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Repository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatabaseType

`func (o *Repository) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *Repository) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *Repository) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *Repository) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetAllowProvisioning

`func (o *Repository) GetAllowProvisioning() bool`

GetAllowProvisioning returns the AllowProvisioning field if non-nil, zero value otherwise.

### GetAllowProvisioningOk

`func (o *Repository) GetAllowProvisioningOk() (*bool, bool)`

GetAllowProvisioningOk returns a tuple with the AllowProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvisioning

`func (o *Repository) SetAllowProvisioning(v bool)`

SetAllowProvisioning sets AllowProvisioning field to given value.

### HasAllowProvisioning

`func (o *Repository) HasAllowProvisioning() bool`

HasAllowProvisioning returns a boolean if a field has been set.

### GetIsStaging

`func (o *Repository) GetIsStaging() bool`

GetIsStaging returns the IsStaging field if non-nil, zero value otherwise.

### GetIsStagingOk

`func (o *Repository) GetIsStagingOk() (*bool, bool)`

GetIsStagingOk returns a tuple with the IsStaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaging

`func (o *Repository) SetIsStaging(v bool)`

SetIsStaging sets IsStaging field to given value.

### HasIsStaging

`func (o *Repository) HasIsStaging() bool`

HasIsStaging returns a boolean if a field has been set.

### GetOracleBase

`func (o *Repository) GetOracleBase() string`

GetOracleBase returns the OracleBase field if non-nil, zero value otherwise.

### GetOracleBaseOk

`func (o *Repository) GetOracleBaseOk() (*string, bool)`

GetOracleBaseOk returns a tuple with the OracleBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleBase

`func (o *Repository) SetOracleBase(v string)`

SetOracleBase sets OracleBase field to given value.

### HasOracleBase

`func (o *Repository) HasOracleBase() bool`

HasOracleBase returns a boolean if a field has been set.

### GetVersion

`func (o *Repository) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Repository) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Repository) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Repository) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBits

`func (o *Repository) GetBits() int32`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *Repository) GetBitsOk() (*int32, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *Repository) SetBits(v int32)`

SetBits sets Bits field to given value.

### HasBits

`func (o *Repository) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetInstallGroup

`func (o *Repository) GetInstallGroup() string`

GetInstallGroup returns the InstallGroup field if non-nil, zero value otherwise.

### GetInstallGroupOk

`func (o *Repository) GetInstallGroupOk() (*string, bool)`

GetInstallGroupOk returns a tuple with the InstallGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGroup

`func (o *Repository) SetInstallGroup(v string)`

SetInstallGroup sets InstallGroup field to given value.

### HasInstallGroup

`func (o *Repository) HasInstallGroup() bool`

HasInstallGroup returns a boolean if a field has been set.

### GetInstallUser

`func (o *Repository) GetInstallUser() string`

GetInstallUser returns the InstallUser field if non-nil, zero value otherwise.

### GetInstallUserOk

`func (o *Repository) GetInstallUserOk() (*string, bool)`

GetInstallUserOk returns a tuple with the InstallUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUser

`func (o *Repository) SetInstallUser(v string)`

SetInstallUser sets InstallUser field to given value.

### HasInstallUser

`func (o *Repository) HasInstallUser() bool`

HasInstallUser returns a boolean if a field has been set.

### GetRac

`func (o *Repository) GetRac() bool`

GetRac returns the Rac field if non-nil, zero value otherwise.

### GetRacOk

`func (o *Repository) GetRacOk() (*bool, bool)`

GetRacOk returns a tuple with the Rac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRac

`func (o *Repository) SetRac(v bool)`

SetRac sets Rac field to given value.

### HasRac

`func (o *Repository) HasRac() bool`

HasRac returns a boolean if a field has been set.

### GetPorts

`func (o *Repository) GetPorts() []int64`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Repository) GetPortsOk() (*[]int64, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Repository) SetPorts(v []int64)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Repository) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetDumpHistoryFile

`func (o *Repository) GetDumpHistoryFile() string`

GetDumpHistoryFile returns the DumpHistoryFile field if non-nil, zero value otherwise.

### GetDumpHistoryFileOk

`func (o *Repository) GetDumpHistoryFileOk() (*string, bool)`

GetDumpHistoryFileOk returns a tuple with the DumpHistoryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpHistoryFile

`func (o *Repository) SetDumpHistoryFile(v string)`

SetDumpHistoryFile sets DumpHistoryFile field to given value.

### HasDumpHistoryFile

`func (o *Repository) HasDumpHistoryFile() bool`

HasDumpHistoryFile returns a boolean if a field has been set.

### GetPageSize

`func (o *Repository) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Repository) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Repository) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Repository) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetOwner

`func (o *Repository) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Repository) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Repository) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Repository) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetInstallationPath

`func (o *Repository) GetInstallationPath() string`

GetInstallationPath returns the InstallationPath field if non-nil, zero value otherwise.

### GetInstallationPathOk

`func (o *Repository) GetInstallationPathOk() (*string, bool)`

GetInstallationPathOk returns a tuple with the InstallationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationPath

`func (o *Repository) SetInstallationPath(v string)`

SetInstallationPath sets InstallationPath field to given value.

### HasInstallationPath

`func (o *Repository) HasInstallationPath() bool`

HasInstallationPath returns a boolean if a field has been set.

### GetFulltextInstalled

`func (o *Repository) GetFulltextInstalled() bool`

GetFulltextInstalled returns the FulltextInstalled field if non-nil, zero value otherwise.

### GetFulltextInstalledOk

`func (o *Repository) GetFulltextInstalledOk() (*bool, bool)`

GetFulltextInstalledOk returns a tuple with the FulltextInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltextInstalled

`func (o *Repository) SetFulltextInstalled(v bool)`

SetFulltextInstalled sets FulltextInstalled field to given value.

### HasFulltextInstalled

`func (o *Repository) HasFulltextInstalled() bool`

HasFulltextInstalled returns a boolean if a field has been set.

### GetInternalVersion

`func (o *Repository) GetInternalVersion() int64`

GetInternalVersion returns the InternalVersion field if non-nil, zero value otherwise.

### GetInternalVersionOk

`func (o *Repository) GetInternalVersionOk() (*int64, bool)`

GetInternalVersionOk returns a tuple with the InternalVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalVersion

`func (o *Repository) SetInternalVersion(v int64)`

SetInternalVersion sets InternalVersion field to given value.

### HasInternalVersion

`func (o *Repository) HasInternalVersion() bool`

HasInternalVersion returns a boolean if a field has been set.

### GetMssqlClusterInstancesName

`func (o *Repository) GetMssqlClusterInstancesName() []string`

GetMssqlClusterInstancesName returns the MssqlClusterInstancesName field if non-nil, zero value otherwise.

### GetMssqlClusterInstancesNameOk

`func (o *Repository) GetMssqlClusterInstancesNameOk() (*[]string, bool)`

GetMssqlClusterInstancesNameOk returns a tuple with the MssqlClusterInstancesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlClusterInstancesName

`func (o *Repository) SetMssqlClusterInstancesName(v []string)`

SetMssqlClusterInstancesName sets MssqlClusterInstancesName field to given value.

### HasMssqlClusterInstancesName

`func (o *Repository) HasMssqlClusterInstancesName() bool`

HasMssqlClusterInstancesName returns a boolean if a field has been set.

### GetMssqlClusterInstancesVersion

`func (o *Repository) GetMssqlClusterInstancesVersion() []string`

GetMssqlClusterInstancesVersion returns the MssqlClusterInstancesVersion field if non-nil, zero value otherwise.

### GetMssqlClusterInstancesVersionOk

`func (o *Repository) GetMssqlClusterInstancesVersionOk() (*[]string, bool)`

GetMssqlClusterInstancesVersionOk returns a tuple with the MssqlClusterInstancesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlClusterInstancesVersion

`func (o *Repository) SetMssqlClusterInstancesVersion(v []string)`

SetMssqlClusterInstancesVersion sets MssqlClusterInstancesVersion field to given value.

### HasMssqlClusterInstancesVersion

`func (o *Repository) HasMssqlClusterInstancesVersion() bool`

HasMssqlClusterInstancesVersion returns a boolean if a field has been set.

### GetInstallationHome

`func (o *Repository) GetInstallationHome() string`

GetInstallationHome returns the InstallationHome field if non-nil, zero value otherwise.

### GetInstallationHomeOk

`func (o *Repository) GetInstallationHomeOk() (*string, bool)`

GetInstallationHomeOk returns a tuple with the InstallationHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationHome

`func (o *Repository) SetInstallationHome(v string)`

SetInstallationHome sets InstallationHome field to given value.

### HasInstallationHome

`func (o *Repository) HasInstallationHome() bool`

HasInstallationHome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


