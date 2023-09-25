# UpdateDatabaseTemplateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The DatabaseTemplate name. | [optional] 
**Description** | Pointer to **string** | User provided description for this template. | [optional] 
**SourceType** | Pointer to **string** | The type of the source associated with the template. | [optional] 
**Parameters** | Pointer to **map[string]string** | A name/value map of string configuration parameters. | [optional] 

## Methods

### NewUpdateDatabaseTemplateParameters

`func NewUpdateDatabaseTemplateParameters() *UpdateDatabaseTemplateParameters`

NewUpdateDatabaseTemplateParameters instantiates a new UpdateDatabaseTemplateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatabaseTemplateParametersWithDefaults

`func NewUpdateDatabaseTemplateParametersWithDefaults() *UpdateDatabaseTemplateParameters`

NewUpdateDatabaseTemplateParametersWithDefaults instantiates a new UpdateDatabaseTemplateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDatabaseTemplateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDatabaseTemplateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDatabaseTemplateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDatabaseTemplateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDatabaseTemplateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDatabaseTemplateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDatabaseTemplateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDatabaseTemplateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceType

`func (o *UpdateDatabaseTemplateParameters) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UpdateDatabaseTemplateParameters) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UpdateDatabaseTemplateParameters) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *UpdateDatabaseTemplateParameters) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateDatabaseTemplateParameters) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateDatabaseTemplateParameters) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateDatabaseTemplateParameters) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateDatabaseTemplateParameters) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


