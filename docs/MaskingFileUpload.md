# MaskingFileUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | Pointer to **string** | The id of the engine onto which the file was uploaded. | [optional] 
**Filename** | Pointer to **string** | Name of this file. | [optional] 
**FileReferenceId** | Pointer to **string** | An reference to this file. | [optional] 
**FileSize** | Pointer to **int64** | Size of this file in bytes. | [optional] 

## Methods

### NewMaskingFileUpload

`func NewMaskingFileUpload() *MaskingFileUpload`

NewMaskingFileUpload instantiates a new MaskingFileUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskingFileUploadWithDefaults

`func NewMaskingFileUploadWithDefaults() *MaskingFileUpload`

NewMaskingFileUploadWithDefaults instantiates a new MaskingFileUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *MaskingFileUpload) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *MaskingFileUpload) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *MaskingFileUpload) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *MaskingFileUpload) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetFilename

`func (o *MaskingFileUpload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MaskingFileUpload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MaskingFileUpload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MaskingFileUpload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFileReferenceId

`func (o *MaskingFileUpload) GetFileReferenceId() string`

GetFileReferenceId returns the FileReferenceId field if non-nil, zero value otherwise.

### GetFileReferenceIdOk

`func (o *MaskingFileUpload) GetFileReferenceIdOk() (*string, bool)`

GetFileReferenceIdOk returns a tuple with the FileReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileReferenceId

`func (o *MaskingFileUpload) SetFileReferenceId(v string)`

SetFileReferenceId sets FileReferenceId field to given value.

### HasFileReferenceId

`func (o *MaskingFileUpload) HasFileReferenceId() bool`

HasFileReferenceId returns a boolean if a field has been set.

### GetFileSize

`func (o *MaskingFileUpload) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *MaskingFileUpload) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *MaskingFileUpload) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *MaskingFileUpload) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


