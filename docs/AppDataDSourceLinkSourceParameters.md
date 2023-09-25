# AppDataDSourceLinkSourceParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the dSource to be created. | [optional] 
**SourceId** | Pointer to **string** | Id of the source to link. | [optional] 
**GroupId** | Pointer to **string** | Id of the dataset group where this dSource should belong to. | [optional] 
**Description** | Pointer to **string** | The notes/description for the dSource. | [optional] 
**LogSyncEnabled** | Pointer to **bool** | True if LogSync should run for this database. | [optional] [default to false]
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this reporting schedule must be configured as owner of the reporting schedule. | [optional] [default to true]
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for dSource. | [optional] 
**OpsPreSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing. | [optional] 
**OpsPostSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform after syncing a created dSource. | [optional] 
**LinkType** | Pointer to **string** | The type of link to create. Default is AppDataDirect.  * &#x60;AppDataDirect&#x60; - Represents the AppData specific parameters of a link request for a source directly replicated into the Delphix Engine. * &#x60;AppDataStaged&#x60; - Represents the AppData specific parameters of a link request for a source with a staging source.  | [optional] [default to "AppDataDirect"]
**StagingMountBase** | Pointer to **string** | The base mount point for the NFS mount on the staging environment [AppDataStaged only]. | [optional] 
**StagingEnvironment** | Pointer to **string** | The environment used as an intermediate stage to pull data into Delphix [AppDataStaged only]. | [optional] 
**StagingEnvironmentUser** | Pointer to **string** | The environment user used to access the staging environment [AppDataStaged only]. | [optional] 
**EnvironmentUser** | **string** | The OS user to use for linking. | 
**Excludes** | Pointer to **[]string** | List of subdirectories in the source to exclude when syncing data. These paths are relative to the root of the source directory. [AppDataDirect only] | [optional] 
**FollowSymlinks** | Pointer to **[]string** | List of symlinks in the source to follow when syncing data. These paths are relative to the root of the source directory. All other symlinks are preserved. [AppDataDirect only] | [optional] 
**Parameters** | **map[string]interface{}** | The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. | 
**SyncParameters** | **map[string]interface{}** | The JSON payload conforming to the snapshot parameters definition in a LUA toolkit or platform plugin. | 

## Methods

### NewAppDataDSourceLinkSourceParameters

`func NewAppDataDSourceLinkSourceParameters(environmentUser string, parameters map[string]interface{}, syncParameters map[string]interface{}, ) *AppDataDSourceLinkSourceParameters`

NewAppDataDSourceLinkSourceParameters instantiates a new AppDataDSourceLinkSourceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataDSourceLinkSourceParametersWithDefaults

`func NewAppDataDSourceLinkSourceParametersWithDefaults() *AppDataDSourceLinkSourceParameters`

NewAppDataDSourceLinkSourceParametersWithDefaults instantiates a new AppDataDSourceLinkSourceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppDataDSourceLinkSourceParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDataDSourceLinkSourceParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDataDSourceLinkSourceParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppDataDSourceLinkSourceParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *AppDataDSourceLinkSourceParameters) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AppDataDSourceLinkSourceParameters) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AppDataDSourceLinkSourceParameters) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AppDataDSourceLinkSourceParameters) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetGroupId

`func (o *AppDataDSourceLinkSourceParameters) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AppDataDSourceLinkSourceParameters) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AppDataDSourceLinkSourceParameters) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AppDataDSourceLinkSourceParameters) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDescription

`func (o *AppDataDSourceLinkSourceParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDataDSourceLinkSourceParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDataDSourceLinkSourceParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDataDSourceLinkSourceParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogSyncEnabled

`func (o *AppDataDSourceLinkSourceParameters) GetLogSyncEnabled() bool`

GetLogSyncEnabled returns the LogSyncEnabled field if non-nil, zero value otherwise.

### GetLogSyncEnabledOk

`func (o *AppDataDSourceLinkSourceParameters) GetLogSyncEnabledOk() (*bool, bool)`

GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncEnabled

`func (o *AppDataDSourceLinkSourceParameters) SetLogSyncEnabled(v bool)`

SetLogSyncEnabled sets LogSyncEnabled field to given value.

### HasLogSyncEnabled

`func (o *AppDataDSourceLinkSourceParameters) HasLogSyncEnabled() bool`

HasLogSyncEnabled returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *AppDataDSourceLinkSourceParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *AppDataDSourceLinkSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *AppDataDSourceLinkSourceParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *AppDataDSourceLinkSourceParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.

### GetTags

`func (o *AppDataDSourceLinkSourceParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AppDataDSourceLinkSourceParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AppDataDSourceLinkSourceParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AppDataDSourceLinkSourceParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOpsPreSync

`func (o *AppDataDSourceLinkSourceParameters) GetOpsPreSync() []SourceOperation`

GetOpsPreSync returns the OpsPreSync field if non-nil, zero value otherwise.

### GetOpsPreSyncOk

`func (o *AppDataDSourceLinkSourceParameters) GetOpsPreSyncOk() (*[]SourceOperation, bool)`

GetOpsPreSyncOk returns a tuple with the OpsPreSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreSync

`func (o *AppDataDSourceLinkSourceParameters) SetOpsPreSync(v []SourceOperation)`

SetOpsPreSync sets OpsPreSync field to given value.

### HasOpsPreSync

`func (o *AppDataDSourceLinkSourceParameters) HasOpsPreSync() bool`

HasOpsPreSync returns a boolean if a field has been set.

### GetOpsPostSync

`func (o *AppDataDSourceLinkSourceParameters) GetOpsPostSync() []SourceOperation`

GetOpsPostSync returns the OpsPostSync field if non-nil, zero value otherwise.

### GetOpsPostSyncOk

`func (o *AppDataDSourceLinkSourceParameters) GetOpsPostSyncOk() (*[]SourceOperation, bool)`

GetOpsPostSyncOk returns a tuple with the OpsPostSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPostSync

`func (o *AppDataDSourceLinkSourceParameters) SetOpsPostSync(v []SourceOperation)`

SetOpsPostSync sets OpsPostSync field to given value.

### HasOpsPostSync

`func (o *AppDataDSourceLinkSourceParameters) HasOpsPostSync() bool`

HasOpsPostSync returns a boolean if a field has been set.

### GetLinkType

`func (o *AppDataDSourceLinkSourceParameters) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *AppDataDSourceLinkSourceParameters) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *AppDataDSourceLinkSourceParameters) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *AppDataDSourceLinkSourceParameters) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetStagingMountBase

`func (o *AppDataDSourceLinkSourceParameters) GetStagingMountBase() string`

GetStagingMountBase returns the StagingMountBase field if non-nil, zero value otherwise.

### GetStagingMountBaseOk

`func (o *AppDataDSourceLinkSourceParameters) GetStagingMountBaseOk() (*string, bool)`

GetStagingMountBaseOk returns a tuple with the StagingMountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingMountBase

`func (o *AppDataDSourceLinkSourceParameters) SetStagingMountBase(v string)`

SetStagingMountBase sets StagingMountBase field to given value.

### HasStagingMountBase

`func (o *AppDataDSourceLinkSourceParameters) HasStagingMountBase() bool`

HasStagingMountBase returns a boolean if a field has been set.

### GetStagingEnvironment

`func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironment() string`

GetStagingEnvironment returns the StagingEnvironment field if non-nil, zero value otherwise.

### GetStagingEnvironmentOk

`func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironmentOk() (*string, bool)`

GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironment

`func (o *AppDataDSourceLinkSourceParameters) SetStagingEnvironment(v string)`

SetStagingEnvironment sets StagingEnvironment field to given value.

### HasStagingEnvironment

`func (o *AppDataDSourceLinkSourceParameters) HasStagingEnvironment() bool`

HasStagingEnvironment returns a boolean if a field has been set.

### GetStagingEnvironmentUser

`func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironmentUser() string`

GetStagingEnvironmentUser returns the StagingEnvironmentUser field if non-nil, zero value otherwise.

### GetStagingEnvironmentUserOk

`func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironmentUserOk() (*string, bool)`

GetStagingEnvironmentUserOk returns a tuple with the StagingEnvironmentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironmentUser

`func (o *AppDataDSourceLinkSourceParameters) SetStagingEnvironmentUser(v string)`

SetStagingEnvironmentUser sets StagingEnvironmentUser field to given value.

### HasStagingEnvironmentUser

`func (o *AppDataDSourceLinkSourceParameters) HasStagingEnvironmentUser() bool`

HasStagingEnvironmentUser returns a boolean if a field has been set.

### GetEnvironmentUser

`func (o *AppDataDSourceLinkSourceParameters) GetEnvironmentUser() string`

GetEnvironmentUser returns the EnvironmentUser field if non-nil, zero value otherwise.

### GetEnvironmentUserOk

`func (o *AppDataDSourceLinkSourceParameters) GetEnvironmentUserOk() (*string, bool)`

GetEnvironmentUserOk returns a tuple with the EnvironmentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUser

`func (o *AppDataDSourceLinkSourceParameters) SetEnvironmentUser(v string)`

SetEnvironmentUser sets EnvironmentUser field to given value.


### GetExcludes

`func (o *AppDataDSourceLinkSourceParameters) GetExcludes() []string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *AppDataDSourceLinkSourceParameters) GetExcludesOk() (*[]string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *AppDataDSourceLinkSourceParameters) SetExcludes(v []string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *AppDataDSourceLinkSourceParameters) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### GetFollowSymlinks

`func (o *AppDataDSourceLinkSourceParameters) GetFollowSymlinks() []string`

GetFollowSymlinks returns the FollowSymlinks field if non-nil, zero value otherwise.

### GetFollowSymlinksOk

`func (o *AppDataDSourceLinkSourceParameters) GetFollowSymlinksOk() (*[]string, bool)`

GetFollowSymlinksOk returns a tuple with the FollowSymlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowSymlinks

`func (o *AppDataDSourceLinkSourceParameters) SetFollowSymlinks(v []string)`

SetFollowSymlinks sets FollowSymlinks field to given value.

### HasFollowSymlinks

`func (o *AppDataDSourceLinkSourceParameters) HasFollowSymlinks() bool`

HasFollowSymlinks returns a boolean if a field has been set.

### GetParameters

`func (o *AppDataDSourceLinkSourceParameters) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AppDataDSourceLinkSourceParameters) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AppDataDSourceLinkSourceParameters) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetSyncParameters

`func (o *AppDataDSourceLinkSourceParameters) GetSyncParameters() map[string]interface{}`

GetSyncParameters returns the SyncParameters field if non-nil, zero value otherwise.

### GetSyncParametersOk

`func (o *AppDataDSourceLinkSourceParameters) GetSyncParametersOk() (*map[string]interface{}, bool)`

GetSyncParametersOk returns a tuple with the SyncParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncParameters

`func (o *AppDataDSourceLinkSourceParameters) SetSyncParameters(v map[string]interface{})`

SetSyncParameters sets SyncParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


