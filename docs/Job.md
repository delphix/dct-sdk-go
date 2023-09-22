# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Job entity ID. | [optional] 
**Status** | Pointer to **string** | The status of the job. | [optional] 
**IsWaitingForTelemetry** | Pointer to **bool** | Indicates that the operations performed by this Job have completed successfully, but the object changes are not yet reflected. This is only set when when the JOB is in STARTED status, with the guarantee that the job will not transition to the FAILED status. Note that this flag will likely be replaced with a new status in future API versions and be deprecated. | [optional] 
**Type** | Pointer to **string** | The type of job being done. | [optional] 
**LocalizedType** | Pointer to **string** | The i18n translated type of job being done. | [optional] 
**ErrorDetails** | Pointer to **string** | Details about the failure for FAILED jobs. | [optional] 
**WarningMessage** | Pointer to **string** | Warnings for the job. | [optional] 
**TargetId** | Pointer to **string** | A reference to the job&#39;s target. | [optional] 
**StartTime** | Pointer to **time.Time** | The time the job started executing. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time the job was last updated. | [optional] 
**TraceId** | Pointer to **string** | traceId of the request which created this Job | [optional] 
**EngineIds** | Pointer to **[]string** | IDs of the engines this Job is executing on. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Engines** | Pointer to [**[]Engine**](Engine.md) |  | [optional] 
**AccountId** | Pointer to **int32** | The ID of the account who initiated this job. | [optional] 
**AccountName** | Pointer to **string** | The account name which initiated this job. It can be either firstname and lastname combination or firstname or lastname or username or email address or Account-&lt;id&gt;. | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Job) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsWaitingForTelemetry

`func (o *Job) GetIsWaitingForTelemetry() bool`

GetIsWaitingForTelemetry returns the IsWaitingForTelemetry field if non-nil, zero value otherwise.

### GetIsWaitingForTelemetryOk

`func (o *Job) GetIsWaitingForTelemetryOk() (*bool, bool)`

GetIsWaitingForTelemetryOk returns a tuple with the IsWaitingForTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWaitingForTelemetry

`func (o *Job) SetIsWaitingForTelemetry(v bool)`

SetIsWaitingForTelemetry sets IsWaitingForTelemetry field to given value.

### HasIsWaitingForTelemetry

`func (o *Job) HasIsWaitingForTelemetry() bool`

HasIsWaitingForTelemetry returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Job) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLocalizedType

`func (o *Job) GetLocalizedType() string`

GetLocalizedType returns the LocalizedType field if non-nil, zero value otherwise.

### GetLocalizedTypeOk

`func (o *Job) GetLocalizedTypeOk() (*string, bool)`

GetLocalizedTypeOk returns a tuple with the LocalizedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedType

`func (o *Job) SetLocalizedType(v string)`

SetLocalizedType sets LocalizedType field to given value.

### HasLocalizedType

`func (o *Job) HasLocalizedType() bool`

HasLocalizedType returns a boolean if a field has been set.

### GetErrorDetails

`func (o *Job) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *Job) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *Job) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *Job) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetWarningMessage

`func (o *Job) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *Job) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *Job) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *Job) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### GetTargetId

`func (o *Job) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Job) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Job) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Job) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetStartTime

`func (o *Job) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Job) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Job) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Job) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Job) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Job) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Job) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Job) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTraceId

`func (o *Job) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *Job) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *Job) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *Job) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetEngineIds

`func (o *Job) GetEngineIds() []string`

GetEngineIds returns the EngineIds field if non-nil, zero value otherwise.

### GetEngineIdsOk

`func (o *Job) GetEngineIdsOk() (*[]string, bool)`

GetEngineIdsOk returns a tuple with the EngineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIds

`func (o *Job) SetEngineIds(v []string)`

SetEngineIds sets EngineIds field to given value.

### HasEngineIds

`func (o *Job) HasEngineIds() bool`

HasEngineIds returns a boolean if a field has been set.

### GetTags

`func (o *Job) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Job) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Job) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Job) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEngines

`func (o *Job) GetEngines() []Engine`

GetEngines returns the Engines field if non-nil, zero value otherwise.

### GetEnginesOk

`func (o *Job) GetEnginesOk() (*[]Engine, bool)`

GetEnginesOk returns a tuple with the Engines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngines

`func (o *Job) SetEngines(v []Engine)`

SetEngines sets Engines field to given value.

### HasEngines

`func (o *Job) HasEngines() bool`

HasEngines returns a boolean if a field has been set.

### GetAccountId

`func (o *Job) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Job) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Job) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Job) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *Job) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Job) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Job) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Job) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


