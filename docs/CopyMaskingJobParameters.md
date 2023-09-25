# CopyMaskingJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEngineId** | **string** | The ID of the engine to copy the job to. | 
**SourceEnvironmentId** | Pointer to **string** | The ID or name of the source environment on the target engine. This only applies to On-The-Fly jobs. | [optional] 
**TargetEnvironmentId** | Pointer to **string** | The ID or name of the target environment to copy the job into. | [optional] 

## Methods

### NewCopyMaskingJobParameters

`func NewCopyMaskingJobParameters(targetEngineId string, ) *CopyMaskingJobParameters`

NewCopyMaskingJobParameters instantiates a new CopyMaskingJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyMaskingJobParametersWithDefaults

`func NewCopyMaskingJobParametersWithDefaults() *CopyMaskingJobParameters`

NewCopyMaskingJobParametersWithDefaults instantiates a new CopyMaskingJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEngineId

`func (o *CopyMaskingJobParameters) GetTargetEngineId() string`

GetTargetEngineId returns the TargetEngineId field if non-nil, zero value otherwise.

### GetTargetEngineIdOk

`func (o *CopyMaskingJobParameters) GetTargetEngineIdOk() (*string, bool)`

GetTargetEngineIdOk returns a tuple with the TargetEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEngineId

`func (o *CopyMaskingJobParameters) SetTargetEngineId(v string)`

SetTargetEngineId sets TargetEngineId field to given value.


### GetSourceEnvironmentId

`func (o *CopyMaskingJobParameters) GetSourceEnvironmentId() string`

GetSourceEnvironmentId returns the SourceEnvironmentId field if non-nil, zero value otherwise.

### GetSourceEnvironmentIdOk

`func (o *CopyMaskingJobParameters) GetSourceEnvironmentIdOk() (*string, bool)`

GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentId

`func (o *CopyMaskingJobParameters) SetSourceEnvironmentId(v string)`

SetSourceEnvironmentId sets SourceEnvironmentId field to given value.

### HasSourceEnvironmentId

`func (o *CopyMaskingJobParameters) HasSourceEnvironmentId() bool`

HasSourceEnvironmentId returns a boolean if a field has been set.

### GetTargetEnvironmentId

`func (o *CopyMaskingJobParameters) GetTargetEnvironmentId() string`

GetTargetEnvironmentId returns the TargetEnvironmentId field if non-nil, zero value otherwise.

### GetTargetEnvironmentIdOk

`func (o *CopyMaskingJobParameters) GetTargetEnvironmentIdOk() (*string, bool)`

GetTargetEnvironmentIdOk returns a tuple with the TargetEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironmentId

`func (o *CopyMaskingJobParameters) SetTargetEnvironmentId(v string)`

SetTargetEnvironmentId sets TargetEnvironmentId field to given value.

### HasTargetEnvironmentId

`func (o *CopyMaskingJobParameters) HasTargetEnvironmentId() bool`

HasTargetEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


