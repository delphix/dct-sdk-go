# MaskingFileUploadParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | **string** | The id of the engine onto which the file will be uploaded. | 
**File** | ***os.File** | The file to upload. | 

## Methods

### NewMaskingFileUploadParameters

`func NewMaskingFileUploadParameters(engineId string, file *os.File, ) *MaskingFileUploadParameters`

NewMaskingFileUploadParameters instantiates a new MaskingFileUploadParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskingFileUploadParametersWithDefaults

`func NewMaskingFileUploadParametersWithDefaults() *MaskingFileUploadParameters`

NewMaskingFileUploadParametersWithDefaults instantiates a new MaskingFileUploadParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *MaskingFileUploadParameters) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *MaskingFileUploadParameters) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *MaskingFileUploadParameters) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetFile

`func (o *MaskingFileUploadParameters) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MaskingFileUploadParameters) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *MaskingFileUploadParameters) SetFile(v *os.File)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


