# EnvironmentRepository

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
**EnvironmentId** | Pointer to **string** | The environment ID. | [optional] 

## Methods

### NewEnvironmentRepository

`func NewEnvironmentRepository() *EnvironmentRepository`

NewEnvironmentRepository instantiates a new EnvironmentRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentRepositoryWithDefaults

`func NewEnvironmentRepositoryWithDefaults() *EnvironmentRepository`

NewEnvironmentRepositoryWithDefaults instantiates a new EnvironmentRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentRepository) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatabaseType

`func (o *EnvironmentRepository) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *EnvironmentRepository) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *EnvironmentRepository) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *EnvironmentRepository) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetAllowProvisioning

`func (o *EnvironmentRepository) GetAllowProvisioning() bool`

GetAllowProvisioning returns the AllowProvisioning field if non-nil, zero value otherwise.

### GetAllowProvisioningOk

`func (o *EnvironmentRepository) GetAllowProvisioningOk() (*bool, bool)`

GetAllowProvisioningOk returns a tuple with the AllowProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvisioning

`func (o *EnvironmentRepository) SetAllowProvisioning(v bool)`

SetAllowProvisioning sets AllowProvisioning field to given value.

### HasAllowProvisioning

`func (o *EnvironmentRepository) HasAllowProvisioning() bool`

HasAllowProvisioning returns a boolean if a field has been set.

### GetIsStaging

`func (o *EnvironmentRepository) GetIsStaging() bool`

GetIsStaging returns the IsStaging field if non-nil, zero value otherwise.

### GetIsStagingOk

`func (o *EnvironmentRepository) GetIsStagingOk() (*bool, bool)`

GetIsStagingOk returns a tuple with the IsStaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaging

`func (o *EnvironmentRepository) SetIsStaging(v bool)`

SetIsStaging sets IsStaging field to given value.

### HasIsStaging

`func (o *EnvironmentRepository) HasIsStaging() bool`

HasIsStaging returns a boolean if a field has been set.

### GetOracleBase

`func (o *EnvironmentRepository) GetOracleBase() string`

GetOracleBase returns the OracleBase field if non-nil, zero value otherwise.

### GetOracleBaseOk

`func (o *EnvironmentRepository) GetOracleBaseOk() (*string, bool)`

GetOracleBaseOk returns a tuple with the OracleBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleBase

`func (o *EnvironmentRepository) SetOracleBase(v string)`

SetOracleBase sets OracleBase field to given value.

### HasOracleBase

`func (o *EnvironmentRepository) HasOracleBase() bool`

HasOracleBase returns a boolean if a field has been set.

### GetVersion

`func (o *EnvironmentRepository) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnvironmentRepository) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnvironmentRepository) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EnvironmentRepository) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBits

`func (o *EnvironmentRepository) GetBits() int32`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *EnvironmentRepository) GetBitsOk() (*int32, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *EnvironmentRepository) SetBits(v int32)`

SetBits sets Bits field to given value.

### HasBits

`func (o *EnvironmentRepository) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetInstallGroup

`func (o *EnvironmentRepository) GetInstallGroup() string`

GetInstallGroup returns the InstallGroup field if non-nil, zero value otherwise.

### GetInstallGroupOk

`func (o *EnvironmentRepository) GetInstallGroupOk() (*string, bool)`

GetInstallGroupOk returns a tuple with the InstallGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGroup

`func (o *EnvironmentRepository) SetInstallGroup(v string)`

SetInstallGroup sets InstallGroup field to given value.

### HasInstallGroup

`func (o *EnvironmentRepository) HasInstallGroup() bool`

HasInstallGroup returns a boolean if a field has been set.

### GetInstallUser

`func (o *EnvironmentRepository) GetInstallUser() string`

GetInstallUser returns the InstallUser field if non-nil, zero value otherwise.

### GetInstallUserOk

`func (o *EnvironmentRepository) GetInstallUserOk() (*string, bool)`

GetInstallUserOk returns a tuple with the InstallUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUser

`func (o *EnvironmentRepository) SetInstallUser(v string)`

SetInstallUser sets InstallUser field to given value.

### HasInstallUser

`func (o *EnvironmentRepository) HasInstallUser() bool`

HasInstallUser returns a boolean if a field has been set.

### GetRac

`func (o *EnvironmentRepository) GetRac() bool`

GetRac returns the Rac field if non-nil, zero value otherwise.

### GetRacOk

`func (o *EnvironmentRepository) GetRacOk() (*bool, bool)`

GetRacOk returns a tuple with the Rac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRac

`func (o *EnvironmentRepository) SetRac(v bool)`

SetRac sets Rac field to given value.

### HasRac

`func (o *EnvironmentRepository) HasRac() bool`

HasRac returns a boolean if a field has been set.

### GetPorts

`func (o *EnvironmentRepository) GetPorts() []int64`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *EnvironmentRepository) GetPortsOk() (*[]int64, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *EnvironmentRepository) SetPorts(v []int64)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *EnvironmentRepository) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetDumpHistoryFile

`func (o *EnvironmentRepository) GetDumpHistoryFile() string`

GetDumpHistoryFile returns the DumpHistoryFile field if non-nil, zero value otherwise.

### GetDumpHistoryFileOk

`func (o *EnvironmentRepository) GetDumpHistoryFileOk() (*string, bool)`

GetDumpHistoryFileOk returns a tuple with the DumpHistoryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpHistoryFile

`func (o *EnvironmentRepository) SetDumpHistoryFile(v string)`

SetDumpHistoryFile sets DumpHistoryFile field to given value.

### HasDumpHistoryFile

`func (o *EnvironmentRepository) HasDumpHistoryFile() bool`

HasDumpHistoryFile returns a boolean if a field has been set.

### GetPageSize

`func (o *EnvironmentRepository) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EnvironmentRepository) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EnvironmentRepository) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EnvironmentRepository) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetOwner

`func (o *EnvironmentRepository) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EnvironmentRepository) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EnvironmentRepository) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EnvironmentRepository) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetInstallationPath

`func (o *EnvironmentRepository) GetInstallationPath() string`

GetInstallationPath returns the InstallationPath field if non-nil, zero value otherwise.

### GetInstallationPathOk

`func (o *EnvironmentRepository) GetInstallationPathOk() (*string, bool)`

GetInstallationPathOk returns a tuple with the InstallationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationPath

`func (o *EnvironmentRepository) SetInstallationPath(v string)`

SetInstallationPath sets InstallationPath field to given value.

### HasInstallationPath

`func (o *EnvironmentRepository) HasInstallationPath() bool`

HasInstallationPath returns a boolean if a field has been set.

### GetFulltextInstalled

`func (o *EnvironmentRepository) GetFulltextInstalled() bool`

GetFulltextInstalled returns the FulltextInstalled field if non-nil, zero value otherwise.

### GetFulltextInstalledOk

`func (o *EnvironmentRepository) GetFulltextInstalledOk() (*bool, bool)`

GetFulltextInstalledOk returns a tuple with the FulltextInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltextInstalled

`func (o *EnvironmentRepository) SetFulltextInstalled(v bool)`

SetFulltextInstalled sets FulltextInstalled field to given value.

### HasFulltextInstalled

`func (o *EnvironmentRepository) HasFulltextInstalled() bool`

HasFulltextInstalled returns a boolean if a field has been set.

### GetInternalVersion

`func (o *EnvironmentRepository) GetInternalVersion() int64`

GetInternalVersion returns the InternalVersion field if non-nil, zero value otherwise.

### GetInternalVersionOk

`func (o *EnvironmentRepository) GetInternalVersionOk() (*int64, bool)`

GetInternalVersionOk returns a tuple with the InternalVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalVersion

`func (o *EnvironmentRepository) SetInternalVersion(v int64)`

SetInternalVersion sets InternalVersion field to given value.

### HasInternalVersion

`func (o *EnvironmentRepository) HasInternalVersion() bool`

HasInternalVersion returns a boolean if a field has been set.

### GetMssqlClusterInstancesName

`func (o *EnvironmentRepository) GetMssqlClusterInstancesName() []string`

GetMssqlClusterInstancesName returns the MssqlClusterInstancesName field if non-nil, zero value otherwise.

### GetMssqlClusterInstancesNameOk

`func (o *EnvironmentRepository) GetMssqlClusterInstancesNameOk() (*[]string, bool)`

GetMssqlClusterInstancesNameOk returns a tuple with the MssqlClusterInstancesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlClusterInstancesName

`func (o *EnvironmentRepository) SetMssqlClusterInstancesName(v []string)`

SetMssqlClusterInstancesName sets MssqlClusterInstancesName field to given value.

### HasMssqlClusterInstancesName

`func (o *EnvironmentRepository) HasMssqlClusterInstancesName() bool`

HasMssqlClusterInstancesName returns a boolean if a field has been set.

### GetMssqlClusterInstancesVersion

`func (o *EnvironmentRepository) GetMssqlClusterInstancesVersion() []string`

GetMssqlClusterInstancesVersion returns the MssqlClusterInstancesVersion field if non-nil, zero value otherwise.

### GetMssqlClusterInstancesVersionOk

`func (o *EnvironmentRepository) GetMssqlClusterInstancesVersionOk() (*[]string, bool)`

GetMssqlClusterInstancesVersionOk returns a tuple with the MssqlClusterInstancesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlClusterInstancesVersion

`func (o *EnvironmentRepository) SetMssqlClusterInstancesVersion(v []string)`

SetMssqlClusterInstancesVersion sets MssqlClusterInstancesVersion field to given value.

### HasMssqlClusterInstancesVersion

`func (o *EnvironmentRepository) HasMssqlClusterInstancesVersion() bool`

HasMssqlClusterInstancesVersion returns a boolean if a field has been set.

### GetInstallationHome

`func (o *EnvironmentRepository) GetInstallationHome() string`

GetInstallationHome returns the InstallationHome field if non-nil, zero value otherwise.

### GetInstallationHomeOk

`func (o *EnvironmentRepository) GetInstallationHomeOk() (*string, bool)`

GetInstallationHomeOk returns a tuple with the InstallationHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationHome

`func (o *EnvironmentRepository) SetInstallationHome(v string)`

SetInstallationHome sets InstallationHome field to given value.

### HasInstallationHome

`func (o *EnvironmentRepository) HasInstallationHome() bool`

HasInstallationHome returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *EnvironmentRepository) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EnvironmentRepository) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EnvironmentRepository) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *EnvironmentRepository) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


