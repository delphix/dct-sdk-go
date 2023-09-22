# MaskingEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The MaskingEnvironment entity ID. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine where this MaskingEnvironment resides. | [optional] 
**EngineName** | Pointer to **string** | The name of the Engine where this MaskingEnvironment resides. | [optional] 
**Name** | Pointer to **string** | The name of this MaskingEnvironment. | [optional] 
**Purpose** | Pointer to **string** | The purpose of this MaskingEnvironment. MaskingEnvironments with a &#39;MASK&#39; purpose will have access to Masking and Profiling jobs, whereas Environments with a &#39;TOKENIZE&#39; purpose will have access to Tokenization and Re-Identification jobs. Note that any custom purposes created through the UI will be represented as &#39;MASK&#39; purposes, due to the jobs that they have access to. | [optional] 
**IsWorkflowEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewMaskingEnvironment

`func NewMaskingEnvironment() *MaskingEnvironment`

NewMaskingEnvironment instantiates a new MaskingEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskingEnvironmentWithDefaults

`func NewMaskingEnvironmentWithDefaults() *MaskingEnvironment`

NewMaskingEnvironmentWithDefaults instantiates a new MaskingEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaskingEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaskingEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaskingEnvironment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MaskingEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEngineId

`func (o *MaskingEnvironment) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *MaskingEnvironment) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *MaskingEnvironment) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *MaskingEnvironment) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEngineName

`func (o *MaskingEnvironment) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *MaskingEnvironment) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *MaskingEnvironment) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *MaskingEnvironment) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetName

`func (o *MaskingEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaskingEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaskingEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MaskingEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurpose

`func (o *MaskingEnvironment) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *MaskingEnvironment) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *MaskingEnvironment) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *MaskingEnvironment) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetIsWorkflowEnabled

`func (o *MaskingEnvironment) GetIsWorkflowEnabled() bool`

GetIsWorkflowEnabled returns the IsWorkflowEnabled field if non-nil, zero value otherwise.

### GetIsWorkflowEnabledOk

`func (o *MaskingEnvironment) GetIsWorkflowEnabledOk() (*bool, bool)`

GetIsWorkflowEnabledOk returns a tuple with the IsWorkflowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorkflowEnabled

`func (o *MaskingEnvironment) SetIsWorkflowEnabled(v bool)`

SetIsWorkflowEnabled sets IsWorkflowEnabled field to given value.

### HasIsWorkflowEnabled

`func (o *MaskingEnvironment) HasIsWorkflowEnabled() bool`

HasIsWorkflowEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


