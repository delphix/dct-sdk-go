# DatabaseTemplateCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The DatabaseTemplate name. | 
**Description** | Pointer to **string** | User provided description for this template. | [optional] 
**SourceType** | **string** | The type of the source associated with the template. | 
**Parameters** | Pointer to **map[string]string** | A name/value map of string configuration parameters. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this database template must be configured as owner of the database template. | [optional] [default to true]
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewDatabaseTemplateCreateParameters

`func NewDatabaseTemplateCreateParameters(name string, sourceType string, ) *DatabaseTemplateCreateParameters`

NewDatabaseTemplateCreateParameters instantiates a new DatabaseTemplateCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseTemplateCreateParametersWithDefaults

`func NewDatabaseTemplateCreateParametersWithDefaults() *DatabaseTemplateCreateParameters`

NewDatabaseTemplateCreateParametersWithDefaults instantiates a new DatabaseTemplateCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseTemplateCreateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseTemplateCreateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseTemplateCreateParameters) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseTemplateCreateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseTemplateCreateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseTemplateCreateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseTemplateCreateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceType

`func (o *DatabaseTemplateCreateParameters) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DatabaseTemplateCreateParameters) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DatabaseTemplateCreateParameters) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetParameters

`func (o *DatabaseTemplateCreateParameters) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DatabaseTemplateCreateParameters) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DatabaseTemplateCreateParameters) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DatabaseTemplateCreateParameters) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *DatabaseTemplateCreateParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *DatabaseTemplateCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *DatabaseTemplateCreateParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *DatabaseTemplateCreateParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.

### GetTags

`func (o *DatabaseTemplateCreateParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DatabaseTemplateCreateParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DatabaseTemplateCreateParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DatabaseTemplateCreateParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


