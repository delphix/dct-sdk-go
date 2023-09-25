# Algorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Algorithm entity ID. | [optional] 
**Name** | Pointer to **string** | The name of this Algorithm. | [optional] 
**Type** | Pointer to **string** | The algorithm type. | [optional] 
**Description** | Pointer to **NullableString** | A description of this algorithm. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this algorithm belongs to. | [optional] 
**IsTokenizationSupported** | Pointer to **bool** | Whether tokenization is supported on this algorithm. | [optional] 
**Config** | Pointer to **map[string]interface{}** | The configuration of this algorithm. | [optional] 
**PluginId** | Pointer to **NullableInt32** | Entity ID for the plugin that provides this algorithm. | [optional] 
**FrameworkId** | Pointer to **NullableInt32** | Entity ID for the framework of this algorithm. | [optional] 

## Methods

### NewAlgorithm

`func NewAlgorithm() *Algorithm`

NewAlgorithm instantiates a new Algorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgorithmWithDefaults

`func NewAlgorithmWithDefaults() *Algorithm`

NewAlgorithmWithDefaults instantiates a new Algorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Algorithm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Algorithm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Algorithm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Algorithm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Algorithm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Algorithm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Algorithm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Algorithm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Algorithm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Algorithm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Algorithm) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Algorithm) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Algorithm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Algorithm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Algorithm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Algorithm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Algorithm) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Algorithm) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEngineId

`func (o *Algorithm) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Algorithm) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Algorithm) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *Algorithm) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetIsTokenizationSupported

`func (o *Algorithm) GetIsTokenizationSupported() bool`

GetIsTokenizationSupported returns the IsTokenizationSupported field if non-nil, zero value otherwise.

### GetIsTokenizationSupportedOk

`func (o *Algorithm) GetIsTokenizationSupportedOk() (*bool, bool)`

GetIsTokenizationSupportedOk returns a tuple with the IsTokenizationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenizationSupported

`func (o *Algorithm) SetIsTokenizationSupported(v bool)`

SetIsTokenizationSupported sets IsTokenizationSupported field to given value.

### HasIsTokenizationSupported

`func (o *Algorithm) HasIsTokenizationSupported() bool`

HasIsTokenizationSupported returns a boolean if a field has been set.

### GetConfig

`func (o *Algorithm) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Algorithm) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Algorithm) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Algorithm) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *Algorithm) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *Algorithm) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetPluginId

`func (o *Algorithm) GetPluginId() int32`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *Algorithm) GetPluginIdOk() (*int32, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *Algorithm) SetPluginId(v int32)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *Algorithm) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### SetPluginIdNil

`func (o *Algorithm) SetPluginIdNil(b bool)`

 SetPluginIdNil sets the value for PluginId to be an explicit nil

### UnsetPluginId
`func (o *Algorithm) UnsetPluginId()`

UnsetPluginId ensures that no value is present for PluginId, not even an explicit nil
### GetFrameworkId

`func (o *Algorithm) GetFrameworkId() int32`

GetFrameworkId returns the FrameworkId field if non-nil, zero value otherwise.

### GetFrameworkIdOk

`func (o *Algorithm) GetFrameworkIdOk() (*int32, bool)`

GetFrameworkIdOk returns a tuple with the FrameworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkId

`func (o *Algorithm) SetFrameworkId(v int32)`

SetFrameworkId sets FrameworkId field to given value.

### HasFrameworkId

`func (o *Algorithm) HasFrameworkId() bool`

HasFrameworkId returns a boolean if a field has been set.

### SetFrameworkIdNil

`func (o *Algorithm) SetFrameworkIdNil(b bool)`

 SetFrameworkIdNil sets the value for FrameworkId to be an explicit nil

### UnsetFrameworkId
`func (o *Algorithm) UnsetFrameworkId()`

UnsetFrameworkId ensures that no value is present for FrameworkId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


