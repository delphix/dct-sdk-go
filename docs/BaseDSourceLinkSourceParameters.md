# BaseDSourceLinkSourceParameters

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

## Methods

### NewBaseDSourceLinkSourceParameters

`func NewBaseDSourceLinkSourceParameters() *BaseDSourceLinkSourceParameters`

NewBaseDSourceLinkSourceParameters instantiates a new BaseDSourceLinkSourceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDSourceLinkSourceParametersWithDefaults

`func NewBaseDSourceLinkSourceParametersWithDefaults() *BaseDSourceLinkSourceParameters`

NewBaseDSourceLinkSourceParametersWithDefaults instantiates a new BaseDSourceLinkSourceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseDSourceLinkSourceParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseDSourceLinkSourceParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseDSourceLinkSourceParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseDSourceLinkSourceParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *BaseDSourceLinkSourceParameters) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseDSourceLinkSourceParameters) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseDSourceLinkSourceParameters) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BaseDSourceLinkSourceParameters) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetGroupId

`func (o *BaseDSourceLinkSourceParameters) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BaseDSourceLinkSourceParameters) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BaseDSourceLinkSourceParameters) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BaseDSourceLinkSourceParameters) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDescription

`func (o *BaseDSourceLinkSourceParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseDSourceLinkSourceParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseDSourceLinkSourceParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseDSourceLinkSourceParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogSyncEnabled

`func (o *BaseDSourceLinkSourceParameters) GetLogSyncEnabled() bool`

GetLogSyncEnabled returns the LogSyncEnabled field if non-nil, zero value otherwise.

### GetLogSyncEnabledOk

`func (o *BaseDSourceLinkSourceParameters) GetLogSyncEnabledOk() (*bool, bool)`

GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncEnabled

`func (o *BaseDSourceLinkSourceParameters) SetLogSyncEnabled(v bool)`

SetLogSyncEnabled sets LogSyncEnabled field to given value.

### HasLogSyncEnabled

`func (o *BaseDSourceLinkSourceParameters) HasLogSyncEnabled() bool`

HasLogSyncEnabled returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *BaseDSourceLinkSourceParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *BaseDSourceLinkSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *BaseDSourceLinkSourceParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *BaseDSourceLinkSourceParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.

### GetTags

`func (o *BaseDSourceLinkSourceParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BaseDSourceLinkSourceParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BaseDSourceLinkSourceParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BaseDSourceLinkSourceParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOpsPreSync

`func (o *BaseDSourceLinkSourceParameters) GetOpsPreSync() []SourceOperation`

GetOpsPreSync returns the OpsPreSync field if non-nil, zero value otherwise.

### GetOpsPreSyncOk

`func (o *BaseDSourceLinkSourceParameters) GetOpsPreSyncOk() (*[]SourceOperation, bool)`

GetOpsPreSyncOk returns a tuple with the OpsPreSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreSync

`func (o *BaseDSourceLinkSourceParameters) SetOpsPreSync(v []SourceOperation)`

SetOpsPreSync sets OpsPreSync field to given value.

### HasOpsPreSync

`func (o *BaseDSourceLinkSourceParameters) HasOpsPreSync() bool`

HasOpsPreSync returns a boolean if a field has been set.

### GetOpsPostSync

`func (o *BaseDSourceLinkSourceParameters) GetOpsPostSync() []SourceOperation`

GetOpsPostSync returns the OpsPostSync field if non-nil, zero value otherwise.

### GetOpsPostSyncOk

`func (o *BaseDSourceLinkSourceParameters) GetOpsPostSyncOk() (*[]SourceOperation, bool)`

GetOpsPostSyncOk returns a tuple with the OpsPostSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPostSync

`func (o *BaseDSourceLinkSourceParameters) SetOpsPostSync(v []SourceOperation)`

SetOpsPostSync sets OpsPostSync field to given value.

### HasOpsPostSync

`func (o *BaseDSourceLinkSourceParameters) HasOpsPostSync() bool`

HasOpsPostSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


