# CopyMaskingJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 
**MaskingJobId** | Pointer to **string** |  | [optional] 

## Methods

### NewCopyMaskingJobResponse

`func NewCopyMaskingJobResponse() *CopyMaskingJobResponse`

NewCopyMaskingJobResponse instantiates a new CopyMaskingJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyMaskingJobResponseWithDefaults

`func NewCopyMaskingJobResponseWithDefaults() *CopyMaskingJobResponse`

NewCopyMaskingJobResponseWithDefaults instantiates a new CopyMaskingJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *CopyMaskingJobResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CopyMaskingJobResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CopyMaskingJobResponse) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *CopyMaskingJobResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetMaskingJobId

`func (o *CopyMaskingJobResponse) GetMaskingJobId() string`

GetMaskingJobId returns the MaskingJobId field if non-nil, zero value otherwise.

### GetMaskingJobIdOk

`func (o *CopyMaskingJobResponse) GetMaskingJobIdOk() (*string, bool)`

GetMaskingJobIdOk returns a tuple with the MaskingJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingJobId

`func (o *CopyMaskingJobResponse) SetMaskingJobId(v string)`

SetMaskingJobId sets MaskingJobId field to given value.

### HasMaskingJobId

`func (o *CopyMaskingJobResponse) HasMaskingJobId() bool`

HasMaskingJobId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


