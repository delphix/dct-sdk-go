# DSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The dSource object entity ID. | [optional] 
**DatabaseType** | Pointer to **NullableString** | The database type of this dSource. | [optional] 
**Name** | Pointer to **NullableString** | The container name of this dSource. | [optional] 
**NamespaceId** | Pointer to **NullableString** | The namespace id of this dSource. | [optional] 
**NamespaceName** | Pointer to **NullableString** | The namespace name of this dSource. | [optional] 
**IsReplica** | Pointer to **NullableBool** | Is this a replicated object. | [optional] 
**DatabaseVersion** | Pointer to **NullableString** | The database version of this dSource. | [optional] 
**ContentType** | Pointer to **NullableString** | The content type of the dSource. | [optional] 
**DataUuid** | Pointer to **NullableString** | A universal ID that uniquely identifies the dSource database. | [optional] 
**StorageSize** | Pointer to **NullableInt64** | The actual space used by this dSource, in bytes. | [optional] 
**PluginVersion** | Pointer to **NullableString** | The version of the plugin associated with this source database. | [optional] 
**CreationDate** | Pointer to **NullableTime** | The date this dSource was created. | [optional] 
**GroupName** | Pointer to **NullableString** | The name of the group containing this dSource. | [optional] 
**Enabled** | Pointer to **NullableBool** | A value indicating whether this dSource is enabled. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this dSource belongs to. | [optional] 
**SourceId** | Pointer to **NullableString** | A reference to the Source associated with this dSource. | [optional] 
**Status** | Pointer to **NullableString** | The runtime status of the dSource. &#39;Unknown&#39; if all attempts to connect to the source failed. | [optional] 
**EngineName** | Pointer to **NullableString** | Name of the Engine where this DSource is hosted | [optional] 
**CdbId** | Pointer to **NullableString** | A reference to the CDB associated with this dSource. | [optional] 
**CurrentTimeflowId** | Pointer to **string** | A reference to the currently active timeflow for this dSource. | [optional] 
**PreviousTimeflowId** | Pointer to **string** | A reference to the previous timeflow for this dSource. | [optional] 
**IsAppdata** | Pointer to **bool** | Indicates whether this dSource has an AppData database. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewDSource

`func NewDSource() *DSource`

NewDSource instantiates a new DSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSourceWithDefaults

`func NewDSourceWithDefaults() *DSource`

NewDSourceWithDefaults instantiates a new DSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDatabaseType

`func (o *DSource) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *DSource) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *DSource) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *DSource) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### SetDatabaseTypeNil

`func (o *DSource) SetDatabaseTypeNil(b bool)`

 SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil

### UnsetDatabaseType
`func (o *DSource) UnsetDatabaseType()`

UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
### GetName

`func (o *DSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespaceId

`func (o *DSource) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *DSource) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *DSource) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *DSource) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### SetNamespaceIdNil

`func (o *DSource) SetNamespaceIdNil(b bool)`

 SetNamespaceIdNil sets the value for NamespaceId to be an explicit nil

### UnsetNamespaceId
`func (o *DSource) UnsetNamespaceId()`

UnsetNamespaceId ensures that no value is present for NamespaceId, not even an explicit nil
### GetNamespaceName

`func (o *DSource) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *DSource) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *DSource) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *DSource) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### SetNamespaceNameNil

`func (o *DSource) SetNamespaceNameNil(b bool)`

 SetNamespaceNameNil sets the value for NamespaceName to be an explicit nil

### UnsetNamespaceName
`func (o *DSource) UnsetNamespaceName()`

UnsetNamespaceName ensures that no value is present for NamespaceName, not even an explicit nil
### GetIsReplica

`func (o *DSource) GetIsReplica() bool`

GetIsReplica returns the IsReplica field if non-nil, zero value otherwise.

### GetIsReplicaOk

`func (o *DSource) GetIsReplicaOk() (*bool, bool)`

GetIsReplicaOk returns a tuple with the IsReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplica

`func (o *DSource) SetIsReplica(v bool)`

SetIsReplica sets IsReplica field to given value.

### HasIsReplica

`func (o *DSource) HasIsReplica() bool`

HasIsReplica returns a boolean if a field has been set.

### SetIsReplicaNil

`func (o *DSource) SetIsReplicaNil(b bool)`

 SetIsReplicaNil sets the value for IsReplica to be an explicit nil

### UnsetIsReplica
`func (o *DSource) UnsetIsReplica()`

UnsetIsReplica ensures that no value is present for IsReplica, not even an explicit nil
### GetDatabaseVersion

`func (o *DSource) GetDatabaseVersion() string`

GetDatabaseVersion returns the DatabaseVersion field if non-nil, zero value otherwise.

### GetDatabaseVersionOk

`func (o *DSource) GetDatabaseVersionOk() (*string, bool)`

GetDatabaseVersionOk returns a tuple with the DatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVersion

`func (o *DSource) SetDatabaseVersion(v string)`

SetDatabaseVersion sets DatabaseVersion field to given value.

### HasDatabaseVersion

`func (o *DSource) HasDatabaseVersion() bool`

HasDatabaseVersion returns a boolean if a field has been set.

### SetDatabaseVersionNil

`func (o *DSource) SetDatabaseVersionNil(b bool)`

 SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil

### UnsetDatabaseVersion
`func (o *DSource) UnsetDatabaseVersion()`

UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
### GetContentType

`func (o *DSource) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DSource) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DSource) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DSource) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *DSource) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *DSource) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDataUuid

`func (o *DSource) GetDataUuid() string`

GetDataUuid returns the DataUuid field if non-nil, zero value otherwise.

### GetDataUuidOk

`func (o *DSource) GetDataUuidOk() (*string, bool)`

GetDataUuidOk returns a tuple with the DataUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUuid

`func (o *DSource) SetDataUuid(v string)`

SetDataUuid sets DataUuid field to given value.

### HasDataUuid

`func (o *DSource) HasDataUuid() bool`

HasDataUuid returns a boolean if a field has been set.

### SetDataUuidNil

`func (o *DSource) SetDataUuidNil(b bool)`

 SetDataUuidNil sets the value for DataUuid to be an explicit nil

### UnsetDataUuid
`func (o *DSource) UnsetDataUuid()`

UnsetDataUuid ensures that no value is present for DataUuid, not even an explicit nil
### GetStorageSize

`func (o *DSource) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *DSource) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *DSource) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *DSource) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *DSource) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *DSource) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetPluginVersion

`func (o *DSource) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *DSource) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *DSource) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.

### HasPluginVersion

`func (o *DSource) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.

### SetPluginVersionNil

`func (o *DSource) SetPluginVersionNil(b bool)`

 SetPluginVersionNil sets the value for PluginVersion to be an explicit nil

### UnsetPluginVersion
`func (o *DSource) UnsetPluginVersion()`

UnsetPluginVersion ensures that no value is present for PluginVersion, not even an explicit nil
### GetCreationDate

`func (o *DSource) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *DSource) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *DSource) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *DSource) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *DSource) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *DSource) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
### GetGroupName

`func (o *DSource) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DSource) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DSource) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DSource) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *DSource) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *DSource) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetEnabled

`func (o *DSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *DSource) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *DSource) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetEngineId

`func (o *DSource) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *DSource) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *DSource) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *DSource) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSourceId

`func (o *DSource) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DSource) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DSource) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DSource) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *DSource) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *DSource) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetStatus

`func (o *DSource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DSource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DSource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DSource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DSource) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DSource) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEngineName

`func (o *DSource) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *DSource) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *DSource) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *DSource) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### SetEngineNameNil

`func (o *DSource) SetEngineNameNil(b bool)`

 SetEngineNameNil sets the value for EngineName to be an explicit nil

### UnsetEngineName
`func (o *DSource) UnsetEngineName()`

UnsetEngineName ensures that no value is present for EngineName, not even an explicit nil
### GetCdbId

`func (o *DSource) GetCdbId() string`

GetCdbId returns the CdbId field if non-nil, zero value otherwise.

### GetCdbIdOk

`func (o *DSource) GetCdbIdOk() (*string, bool)`

GetCdbIdOk returns a tuple with the CdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbId

`func (o *DSource) SetCdbId(v string)`

SetCdbId sets CdbId field to given value.

### HasCdbId

`func (o *DSource) HasCdbId() bool`

HasCdbId returns a boolean if a field has been set.

### SetCdbIdNil

`func (o *DSource) SetCdbIdNil(b bool)`

 SetCdbIdNil sets the value for CdbId to be an explicit nil

### UnsetCdbId
`func (o *DSource) UnsetCdbId()`

UnsetCdbId ensures that no value is present for CdbId, not even an explicit nil
### GetCurrentTimeflowId

`func (o *DSource) GetCurrentTimeflowId() string`

GetCurrentTimeflowId returns the CurrentTimeflowId field if non-nil, zero value otherwise.

### GetCurrentTimeflowIdOk

`func (o *DSource) GetCurrentTimeflowIdOk() (*string, bool)`

GetCurrentTimeflowIdOk returns a tuple with the CurrentTimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTimeflowId

`func (o *DSource) SetCurrentTimeflowId(v string)`

SetCurrentTimeflowId sets CurrentTimeflowId field to given value.

### HasCurrentTimeflowId

`func (o *DSource) HasCurrentTimeflowId() bool`

HasCurrentTimeflowId returns a boolean if a field has been set.

### GetPreviousTimeflowId

`func (o *DSource) GetPreviousTimeflowId() string`

GetPreviousTimeflowId returns the PreviousTimeflowId field if non-nil, zero value otherwise.

### GetPreviousTimeflowIdOk

`func (o *DSource) GetPreviousTimeflowIdOk() (*string, bool)`

GetPreviousTimeflowIdOk returns a tuple with the PreviousTimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTimeflowId

`func (o *DSource) SetPreviousTimeflowId(v string)`

SetPreviousTimeflowId sets PreviousTimeflowId field to given value.

### HasPreviousTimeflowId

`func (o *DSource) HasPreviousTimeflowId() bool`

HasPreviousTimeflowId returns a boolean if a field has been set.

### GetIsAppdata

`func (o *DSource) GetIsAppdata() bool`

GetIsAppdata returns the IsAppdata field if non-nil, zero value otherwise.

### GetIsAppdataOk

`func (o *DSource) GetIsAppdataOk() (*bool, bool)`

GetIsAppdataOk returns a tuple with the IsAppdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppdata

`func (o *DSource) SetIsAppdata(v bool)`

SetIsAppdata sets IsAppdata field to given value.

### HasIsAppdata

`func (o *DSource) HasIsAppdata() bool`

HasIsAppdata returns a boolean if a field has been set.

### GetTags

`func (o *DSource) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DSource) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DSource) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DSource) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


