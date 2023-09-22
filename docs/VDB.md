# VDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The VDB object entity ID. | [optional] 
**DatabaseType** | Pointer to **NullableString** | The database type of this VDB. | [optional] 
**Name** | Pointer to **NullableString** | The container name of this VDB. | [optional] 
**DatabaseVersion** | Pointer to **NullableString** | The database version of this VDB. | [optional] 
**Size** | Pointer to **NullableInt64** | The total size of this VDB, in bytes. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this VDB belongs to. | [optional] 
**Status** | Pointer to **NullableString** | The runtime status of the VDB. &#39;Unknown&#39; if all attempts to connect to the dataset failed. | [optional] 
**Masked** | Pointer to **NullableBool** | The VDB is masked or not. | [optional] 
**ParentTimeflowTimestamp** | Pointer to **NullableTime** | The timestamp for parent timeflow. | [optional] 
**ParentTimeflowTimezone** | Pointer to **NullableString** | The timezone for parent timeflow. | [optional] 
**EnvironmentId** | Pointer to **NullableString** | A reference to the Environment that hosts this VDB. | [optional] 
**IpAddress** | Pointer to **NullableString** | The IP address of the VDB&#39;s host. | [optional] 
**Fqdn** | Pointer to **NullableString** | The FQDN of the VDB&#39;s host. | [optional] 
**ParentId** | Pointer to **NullableString** | A reference to the parent dataset of this VDB. | [optional] 
**ParentDsourceId** | Pointer to **NullableString** | A reference to the parent dSource of this VDB. | [optional] 
**GroupName** | Pointer to **NullableString** | The name of the group containing this VDB. | [optional] 
**EngineName** | Pointer to **NullableString** | Name of the Engine where this VDB is hosted | [optional] 
**CdbId** | Pointer to **NullableString** | A reference to the CDB or VCDB associated with this VDB. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**CreationDate** | Pointer to **NullableTime** | The date this VDB was created. | [optional] 
**Hooks** | Pointer to [**VirtualDatasetHooks**](VirtualDatasetHooks.md) |  | [optional] 
**AppdataSourceParams** | Pointer to **map[string]interface{}** | The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. | [optional] 
**TemplateId** | Pointer to **NullableString** | A reference to the Database Template. | [optional] 
**ConfigParams** | Pointer to **map[string]interface{}** | Database configuration parameter overrides. | [optional] 
**AdditionalMountPoints** | Pointer to [**[]AdditionalMountPoint**](AdditionalMountPoint.md) | Specifies additional locations on which to mount a subdirectory of an AppData container. Can only be updated while the VDB is disabled. | [optional] 
**AppdataConfigParams** | Pointer to **map[string]interface{}** | The parameters specified by the source config schema in the toolkit | [optional] 
**MountPoint** | Pointer to **string** | Mount point for the VDB (Oracle, ASE, AppData). | [optional] 
**CurrentTimeflowId** | Pointer to **string** | A reference to the currently active timeflow for this VDB. | [optional] 
**PreviousTimeflowId** | Pointer to **string** | A reference to the previous timeflow for this VDB. | [optional] 

## Methods

### NewVDB

`func NewVDB() *VDB`

NewVDB instantiates a new VDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDBWithDefaults

`func NewVDBWithDefaults() *VDB`

NewVDBWithDefaults instantiates a new VDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VDB) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VDB) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VDB) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VDB) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDatabaseType

`func (o *VDB) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *VDB) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *VDB) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *VDB) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### SetDatabaseTypeNil

`func (o *VDB) SetDatabaseTypeNil(b bool)`

 SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil

### UnsetDatabaseType
`func (o *VDB) UnsetDatabaseType()`

UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
### GetName

`func (o *VDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VDB) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VDB) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VDB) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VDB) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDatabaseVersion

`func (o *VDB) GetDatabaseVersion() string`

GetDatabaseVersion returns the DatabaseVersion field if non-nil, zero value otherwise.

### GetDatabaseVersionOk

`func (o *VDB) GetDatabaseVersionOk() (*string, bool)`

GetDatabaseVersionOk returns a tuple with the DatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVersion

`func (o *VDB) SetDatabaseVersion(v string)`

SetDatabaseVersion sets DatabaseVersion field to given value.

### HasDatabaseVersion

`func (o *VDB) HasDatabaseVersion() bool`

HasDatabaseVersion returns a boolean if a field has been set.

### SetDatabaseVersionNil

`func (o *VDB) SetDatabaseVersionNil(b bool)`

 SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil

### UnsetDatabaseVersion
`func (o *VDB) UnsetDatabaseVersion()`

UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
### GetSize

`func (o *VDB) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VDB) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VDB) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VDB) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VDB) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VDB) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetEngineId

`func (o *VDB) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *VDB) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *VDB) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *VDB) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetStatus

`func (o *VDB) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VDB) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VDB) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VDB) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VDB) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VDB) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMasked

`func (o *VDB) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *VDB) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *VDB) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *VDB) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### SetMaskedNil

`func (o *VDB) SetMaskedNil(b bool)`

 SetMaskedNil sets the value for Masked to be an explicit nil

### UnsetMasked
`func (o *VDB) UnsetMasked()`

UnsetMasked ensures that no value is present for Masked, not even an explicit nil
### GetParentTimeflowTimestamp

`func (o *VDB) GetParentTimeflowTimestamp() time.Time`

GetParentTimeflowTimestamp returns the ParentTimeflowTimestamp field if non-nil, zero value otherwise.

### GetParentTimeflowTimestampOk

`func (o *VDB) GetParentTimeflowTimestampOk() (*time.Time, bool)`

GetParentTimeflowTimestampOk returns a tuple with the ParentTimeflowTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeflowTimestamp

`func (o *VDB) SetParentTimeflowTimestamp(v time.Time)`

SetParentTimeflowTimestamp sets ParentTimeflowTimestamp field to given value.

### HasParentTimeflowTimestamp

`func (o *VDB) HasParentTimeflowTimestamp() bool`

HasParentTimeflowTimestamp returns a boolean if a field has been set.

### SetParentTimeflowTimestampNil

`func (o *VDB) SetParentTimeflowTimestampNil(b bool)`

 SetParentTimeflowTimestampNil sets the value for ParentTimeflowTimestamp to be an explicit nil

### UnsetParentTimeflowTimestamp
`func (o *VDB) UnsetParentTimeflowTimestamp()`

UnsetParentTimeflowTimestamp ensures that no value is present for ParentTimeflowTimestamp, not even an explicit nil
### GetParentTimeflowTimezone

`func (o *VDB) GetParentTimeflowTimezone() string`

GetParentTimeflowTimezone returns the ParentTimeflowTimezone field if non-nil, zero value otherwise.

### GetParentTimeflowTimezoneOk

`func (o *VDB) GetParentTimeflowTimezoneOk() (*string, bool)`

GetParentTimeflowTimezoneOk returns a tuple with the ParentTimeflowTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeflowTimezone

`func (o *VDB) SetParentTimeflowTimezone(v string)`

SetParentTimeflowTimezone sets ParentTimeflowTimezone field to given value.

### HasParentTimeflowTimezone

`func (o *VDB) HasParentTimeflowTimezone() bool`

HasParentTimeflowTimezone returns a boolean if a field has been set.

### SetParentTimeflowTimezoneNil

`func (o *VDB) SetParentTimeflowTimezoneNil(b bool)`

 SetParentTimeflowTimezoneNil sets the value for ParentTimeflowTimezone to be an explicit nil

### UnsetParentTimeflowTimezone
`func (o *VDB) UnsetParentTimeflowTimezone()`

UnsetParentTimeflowTimezone ensures that no value is present for ParentTimeflowTimezone, not even an explicit nil
### GetEnvironmentId

`func (o *VDB) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *VDB) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *VDB) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *VDB) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *VDB) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *VDB) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetIpAddress

`func (o *VDB) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VDB) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VDB) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VDB) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VDB) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VDB) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetFqdn

`func (o *VDB) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *VDB) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *VDB) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *VDB) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *VDB) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *VDB) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetParentId

`func (o *VDB) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *VDB) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *VDB) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *VDB) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *VDB) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *VDB) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentDsourceId

`func (o *VDB) GetParentDsourceId() string`

GetParentDsourceId returns the ParentDsourceId field if non-nil, zero value otherwise.

### GetParentDsourceIdOk

`func (o *VDB) GetParentDsourceIdOk() (*string, bool)`

GetParentDsourceIdOk returns a tuple with the ParentDsourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDsourceId

`func (o *VDB) SetParentDsourceId(v string)`

SetParentDsourceId sets ParentDsourceId field to given value.

### HasParentDsourceId

`func (o *VDB) HasParentDsourceId() bool`

HasParentDsourceId returns a boolean if a field has been set.

### SetParentDsourceIdNil

`func (o *VDB) SetParentDsourceIdNil(b bool)`

 SetParentDsourceIdNil sets the value for ParentDsourceId to be an explicit nil

### UnsetParentDsourceId
`func (o *VDB) UnsetParentDsourceId()`

UnsetParentDsourceId ensures that no value is present for ParentDsourceId, not even an explicit nil
### GetGroupName

`func (o *VDB) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *VDB) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *VDB) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *VDB) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *VDB) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *VDB) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetEngineName

`func (o *VDB) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *VDB) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *VDB) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *VDB) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### SetEngineNameNil

`func (o *VDB) SetEngineNameNil(b bool)`

 SetEngineNameNil sets the value for EngineName to be an explicit nil

### UnsetEngineName
`func (o *VDB) UnsetEngineName()`

UnsetEngineName ensures that no value is present for EngineName, not even an explicit nil
### GetCdbId

`func (o *VDB) GetCdbId() string`

GetCdbId returns the CdbId field if non-nil, zero value otherwise.

### GetCdbIdOk

`func (o *VDB) GetCdbIdOk() (*string, bool)`

GetCdbIdOk returns a tuple with the CdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbId

`func (o *VDB) SetCdbId(v string)`

SetCdbId sets CdbId field to given value.

### HasCdbId

`func (o *VDB) HasCdbId() bool`

HasCdbId returns a boolean if a field has been set.

### SetCdbIdNil

`func (o *VDB) SetCdbIdNil(b bool)`

 SetCdbIdNil sets the value for CdbId to be an explicit nil

### UnsetCdbId
`func (o *VDB) UnsetCdbId()`

UnsetCdbId ensures that no value is present for CdbId, not even an explicit nil
### GetTags

`func (o *VDB) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VDB) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VDB) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VDB) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreationDate

`func (o *VDB) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *VDB) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *VDB) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *VDB) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *VDB) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *VDB) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
### GetHooks

`func (o *VDB) GetHooks() VirtualDatasetHooks`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *VDB) GetHooksOk() (*VirtualDatasetHooks, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *VDB) SetHooks(v VirtualDatasetHooks)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *VDB) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetAppdataSourceParams

`func (o *VDB) GetAppdataSourceParams() map[string]interface{}`

GetAppdataSourceParams returns the AppdataSourceParams field if non-nil, zero value otherwise.

### GetAppdataSourceParamsOk

`func (o *VDB) GetAppdataSourceParamsOk() (*map[string]interface{}, bool)`

GetAppdataSourceParamsOk returns a tuple with the AppdataSourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataSourceParams

`func (o *VDB) SetAppdataSourceParams(v map[string]interface{})`

SetAppdataSourceParams sets AppdataSourceParams field to given value.

### HasAppdataSourceParams

`func (o *VDB) HasAppdataSourceParams() bool`

HasAppdataSourceParams returns a boolean if a field has been set.

### SetAppdataSourceParamsNil

`func (o *VDB) SetAppdataSourceParamsNil(b bool)`

 SetAppdataSourceParamsNil sets the value for AppdataSourceParams to be an explicit nil

### UnsetAppdataSourceParams
`func (o *VDB) UnsetAppdataSourceParams()`

UnsetAppdataSourceParams ensures that no value is present for AppdataSourceParams, not even an explicit nil
### GetTemplateId

`func (o *VDB) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *VDB) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *VDB) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *VDB) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *VDB) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *VDB) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetConfigParams

`func (o *VDB) GetConfigParams() map[string]interface{}`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *VDB) GetConfigParamsOk() (*map[string]interface{}, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *VDB) SetConfigParams(v map[string]interface{})`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *VDB) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### SetConfigParamsNil

`func (o *VDB) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *VDB) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
### GetAdditionalMountPoints

`func (o *VDB) GetAdditionalMountPoints() []AdditionalMountPoint`

GetAdditionalMountPoints returns the AdditionalMountPoints field if non-nil, zero value otherwise.

### GetAdditionalMountPointsOk

`func (o *VDB) GetAdditionalMountPointsOk() (*[]AdditionalMountPoint, bool)`

GetAdditionalMountPointsOk returns a tuple with the AdditionalMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMountPoints

`func (o *VDB) SetAdditionalMountPoints(v []AdditionalMountPoint)`

SetAdditionalMountPoints sets AdditionalMountPoints field to given value.

### HasAdditionalMountPoints

`func (o *VDB) HasAdditionalMountPoints() bool`

HasAdditionalMountPoints returns a boolean if a field has been set.

### SetAdditionalMountPointsNil

`func (o *VDB) SetAdditionalMountPointsNil(b bool)`

 SetAdditionalMountPointsNil sets the value for AdditionalMountPoints to be an explicit nil

### UnsetAdditionalMountPoints
`func (o *VDB) UnsetAdditionalMountPoints()`

UnsetAdditionalMountPoints ensures that no value is present for AdditionalMountPoints, not even an explicit nil
### GetAppdataConfigParams

`func (o *VDB) GetAppdataConfigParams() map[string]interface{}`

GetAppdataConfigParams returns the AppdataConfigParams field if non-nil, zero value otherwise.

### GetAppdataConfigParamsOk

`func (o *VDB) GetAppdataConfigParamsOk() (*map[string]interface{}, bool)`

GetAppdataConfigParamsOk returns a tuple with the AppdataConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataConfigParams

`func (o *VDB) SetAppdataConfigParams(v map[string]interface{})`

SetAppdataConfigParams sets AppdataConfigParams field to given value.

### HasAppdataConfigParams

`func (o *VDB) HasAppdataConfigParams() bool`

HasAppdataConfigParams returns a boolean if a field has been set.

### SetAppdataConfigParamsNil

`func (o *VDB) SetAppdataConfigParamsNil(b bool)`

 SetAppdataConfigParamsNil sets the value for AppdataConfigParams to be an explicit nil

### UnsetAppdataConfigParams
`func (o *VDB) UnsetAppdataConfigParams()`

UnsetAppdataConfigParams ensures that no value is present for AppdataConfigParams, not even an explicit nil
### GetMountPoint

`func (o *VDB) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *VDB) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *VDB) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *VDB) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetCurrentTimeflowId

`func (o *VDB) GetCurrentTimeflowId() string`

GetCurrentTimeflowId returns the CurrentTimeflowId field if non-nil, zero value otherwise.

### GetCurrentTimeflowIdOk

`func (o *VDB) GetCurrentTimeflowIdOk() (*string, bool)`

GetCurrentTimeflowIdOk returns a tuple with the CurrentTimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTimeflowId

`func (o *VDB) SetCurrentTimeflowId(v string)`

SetCurrentTimeflowId sets CurrentTimeflowId field to given value.

### HasCurrentTimeflowId

`func (o *VDB) HasCurrentTimeflowId() bool`

HasCurrentTimeflowId returns a boolean if a field has been set.

### GetPreviousTimeflowId

`func (o *VDB) GetPreviousTimeflowId() string`

GetPreviousTimeflowId returns the PreviousTimeflowId field if non-nil, zero value otherwise.

### GetPreviousTimeflowIdOk

`func (o *VDB) GetPreviousTimeflowIdOk() (*string, bool)`

GetPreviousTimeflowIdOk returns a tuple with the PreviousTimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTimeflowId

`func (o *VDB) SetPreviousTimeflowId(v string)`

SetPreviousTimeflowId sets PreviousTimeflowId field to given value.

### HasPreviousTimeflowId

`func (o *VDB) HasPreviousTimeflowId() bool`

HasPreviousTimeflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


