# DatabaseTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The DatabaseTemplate entity ID. | [optional] [readonly] 
**Name** | **string** | The DatabaseTemplate name. | 
**Description** | Pointer to **string** | User provided description for this template. | [optional] 
**SourceType** | **string** | The type of the source associated with the template. | 
**Parameters** | Pointer to **map[string]string** | A name/value map of string configuration parameters. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewDatabaseTemplate

`func NewDatabaseTemplate(name string, sourceType string, ) *DatabaseTemplate`

NewDatabaseTemplate instantiates a new DatabaseTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseTemplateWithDefaults

`func NewDatabaseTemplateWithDefaults() *DatabaseTemplate`

NewDatabaseTemplateWithDefaults instantiates a new DatabaseTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatabaseTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DatabaseTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceType

`func (o *DatabaseTemplate) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DatabaseTemplate) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DatabaseTemplate) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetParameters

`func (o *DatabaseTemplate) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DatabaseTemplate) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DatabaseTemplate) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DatabaseTemplate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTags

`func (o *DatabaseTemplate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DatabaseTemplate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DatabaseTemplate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DatabaseTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


