# AlgorithmCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this Algorithm. | 
**Description** | Pointer to **string** | A description of this algorithm. | [optional] 
**EngineId** | **string** | The id of the engine onto which this algorithm will be created. | 
**Config** | **map[string]interface{}** | The configuration of this algorithm. | 
**FrameworkId** | **int32** | Entity ID for the framework of this algorithm. | 

## Methods

### NewAlgorithmCreateParameters

`func NewAlgorithmCreateParameters(name string, engineId string, config map[string]interface{}, frameworkId int32, ) *AlgorithmCreateParameters`

NewAlgorithmCreateParameters instantiates a new AlgorithmCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgorithmCreateParametersWithDefaults

`func NewAlgorithmCreateParametersWithDefaults() *AlgorithmCreateParameters`

NewAlgorithmCreateParametersWithDefaults instantiates a new AlgorithmCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlgorithmCreateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlgorithmCreateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlgorithmCreateParameters) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlgorithmCreateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlgorithmCreateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlgorithmCreateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlgorithmCreateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEngineId

`func (o *AlgorithmCreateParameters) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *AlgorithmCreateParameters) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *AlgorithmCreateParameters) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetConfig

`func (o *AlgorithmCreateParameters) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AlgorithmCreateParameters) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AlgorithmCreateParameters) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetFrameworkId

`func (o *AlgorithmCreateParameters) GetFrameworkId() int32`

GetFrameworkId returns the FrameworkId field if non-nil, zero value otherwise.

### GetFrameworkIdOk

`func (o *AlgorithmCreateParameters) GetFrameworkIdOk() (*int32, bool)`

GetFrameworkIdOk returns a tuple with the FrameworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkId

`func (o *AlgorithmCreateParameters) SetFrameworkId(v int32)`

SetFrameworkId sets FrameworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


