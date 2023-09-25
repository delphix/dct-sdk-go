# MaskingJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Masking Job entity ID. | [optional] 
**Name** | Pointer to **string** | The name of this Masking Job. | [optional] 
**Ruleset** | Pointer to [**MaskingRuleset**](MaskingRuleset.md) |  | [optional] 
**ConnectorType** | Pointer to **string** | The type of data being masked by this Job. If the Masking Job is masking a database this is the type of the database (Standard Job only). | [optional] 
**IsOnTheFlyMasking** | Pointer to **bool** | Whether this is an on-the-fly masking job (Standard Job only). | [optional] 
**CreationDate** | Pointer to **time.Time** | The date this MaskingJob was created (Standard Job only). | [optional] 
**LastCompletedExecutionDate** | Pointer to **time.Time** | The date this MaskingJob was last executed to completion. | [optional] 
**LastExecutionStatus** | Pointer to **string** | The status of this MaskingJob&#39;s last execution. | [optional] 
**LastExecutionRowsMasked** | Pointer to **int64** | The number of rows masked by the last successful execution. This is not applicable for JSON file type. | [optional] 
**LastExecutionRowsTotal** | Pointer to **int64** | The total number of rows masked by the last successful execution. This is not applicable for JSON file type. | [optional] 
**LastExecutionBytesMasked** | Pointer to **int64** | The number of bytes masked by the last successful execution. This is only applicable for JSON file type. | [optional] 
**LastExecutionBytesTotal** | Pointer to **int64** | The total number of bytes masked by the last successful execution. This is only applicable for JSON file type. | [optional] 
**ConnectorUsername** | Pointer to **NullableString** | The username of the Connector used by the MaskingJob (Standard Job only). For hyperscale jobs, see the connector of the dataset. | [optional] 
**ConnectorPassword** | Pointer to **NullableString** | The password of the Connector used by the MaskingJob (Standard Job only). For hyperscale jobs, see the connector of the dataset. | [optional] 
**OnTheFlySourceConnectorUsername** | Pointer to **NullableString** | The username of the source Connector used by the on-the-fly MaskingJob (Standard Job only). | [optional] 
**OnTheFlySourceConnectorPassword** | Pointer to **NullableString** | The password of the source Connector used by the on-the-fly MaskingJob (Standard Job only). | [optional] 
**JobType** | Pointer to **string** | The type of this Job. | [optional] 
**HyperscaleInstanceId** | Pointer to **string** | The ID of the Hyperscale instance of this Job (Hyperscale Job only). | [optional] 
**Description** | Pointer to **string** | Description of the Job (Hyperscale Job only). | [optional] 
**DatasetId** | Pointer to **string** | Dataset of the Hyperscale Job (Hyperscale Job only). | [optional] 
**RetainExecutionData** | Pointer to **string** | Defines whether execution data will be stored after execution is complete (Hyperscale Job only). | [optional] 
**MaxMemory** | Pointer to **int32** | Maximum memory to be allocated for each Masking job (Hyperscale Job only). | [optional] 
**MinMemory** | Pointer to **int32** | Minimum memory to be allocated for each Masking job (Hyperscale Job only). | [optional] 
**FeedbackSize** | Pointer to **int32** | Feedback Size for each Masking job (Hyperscale Job only). | [optional] 
**StreamRowLimit** | Pointer to **int32** | Stream Row Limit for each Masking job (Hyperscale Job only). | [optional] 
**NumInputStreams** | Pointer to **int32** | Number of input streams to be configured for Masking Job (Hyperscale Job only). | [optional] 
**MaxConcurrentSourceConnections** | Pointer to **int32** | Maximum number of parallel connection that the Hyperscale instance can have with the source datasource. | [optional] 
**MaxConcurrentTargetConnections** | Pointer to **int32** | Maximum number of parallel connection that the Hyperscale instance can have with the target datasource. | [optional] 
**ParallelismDegree** | Pointer to **int32** | The degree of parallelism (DOP) per Oracle job to recreate the index in the post-load process. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewMaskingJob

`func NewMaskingJob() *MaskingJob`

NewMaskingJob instantiates a new MaskingJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskingJobWithDefaults

`func NewMaskingJobWithDefaults() *MaskingJob`

NewMaskingJobWithDefaults instantiates a new MaskingJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaskingJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaskingJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaskingJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MaskingJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MaskingJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaskingJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaskingJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MaskingJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleset

`func (o *MaskingJob) GetRuleset() MaskingRuleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *MaskingJob) GetRulesetOk() (*MaskingRuleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *MaskingJob) SetRuleset(v MaskingRuleset)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *MaskingJob) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetConnectorType

`func (o *MaskingJob) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *MaskingJob) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *MaskingJob) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *MaskingJob) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetIsOnTheFlyMasking

`func (o *MaskingJob) GetIsOnTheFlyMasking() bool`

GetIsOnTheFlyMasking returns the IsOnTheFlyMasking field if non-nil, zero value otherwise.

### GetIsOnTheFlyMaskingOk

`func (o *MaskingJob) GetIsOnTheFlyMaskingOk() (*bool, bool)`

GetIsOnTheFlyMaskingOk returns a tuple with the IsOnTheFlyMasking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnTheFlyMasking

`func (o *MaskingJob) SetIsOnTheFlyMasking(v bool)`

SetIsOnTheFlyMasking sets IsOnTheFlyMasking field to given value.

### HasIsOnTheFlyMasking

`func (o *MaskingJob) HasIsOnTheFlyMasking() bool`

HasIsOnTheFlyMasking returns a boolean if a field has been set.

### GetCreationDate

`func (o *MaskingJob) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *MaskingJob) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *MaskingJob) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *MaskingJob) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastCompletedExecutionDate

`func (o *MaskingJob) GetLastCompletedExecutionDate() time.Time`

GetLastCompletedExecutionDate returns the LastCompletedExecutionDate field if non-nil, zero value otherwise.

### GetLastCompletedExecutionDateOk

`func (o *MaskingJob) GetLastCompletedExecutionDateOk() (*time.Time, bool)`

GetLastCompletedExecutionDateOk returns a tuple with the LastCompletedExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletedExecutionDate

`func (o *MaskingJob) SetLastCompletedExecutionDate(v time.Time)`

SetLastCompletedExecutionDate sets LastCompletedExecutionDate field to given value.

### HasLastCompletedExecutionDate

`func (o *MaskingJob) HasLastCompletedExecutionDate() bool`

HasLastCompletedExecutionDate returns a boolean if a field has been set.

### GetLastExecutionStatus

`func (o *MaskingJob) GetLastExecutionStatus() string`

GetLastExecutionStatus returns the LastExecutionStatus field if non-nil, zero value otherwise.

### GetLastExecutionStatusOk

`func (o *MaskingJob) GetLastExecutionStatusOk() (*string, bool)`

GetLastExecutionStatusOk returns a tuple with the LastExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionStatus

`func (o *MaskingJob) SetLastExecutionStatus(v string)`

SetLastExecutionStatus sets LastExecutionStatus field to given value.

### HasLastExecutionStatus

`func (o *MaskingJob) HasLastExecutionStatus() bool`

HasLastExecutionStatus returns a boolean if a field has been set.

### GetLastExecutionRowsMasked

`func (o *MaskingJob) GetLastExecutionRowsMasked() int64`

GetLastExecutionRowsMasked returns the LastExecutionRowsMasked field if non-nil, zero value otherwise.

### GetLastExecutionRowsMaskedOk

`func (o *MaskingJob) GetLastExecutionRowsMaskedOk() (*int64, bool)`

GetLastExecutionRowsMaskedOk returns a tuple with the LastExecutionRowsMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionRowsMasked

`func (o *MaskingJob) SetLastExecutionRowsMasked(v int64)`

SetLastExecutionRowsMasked sets LastExecutionRowsMasked field to given value.

### HasLastExecutionRowsMasked

`func (o *MaskingJob) HasLastExecutionRowsMasked() bool`

HasLastExecutionRowsMasked returns a boolean if a field has been set.

### GetLastExecutionRowsTotal

`func (o *MaskingJob) GetLastExecutionRowsTotal() int64`

GetLastExecutionRowsTotal returns the LastExecutionRowsTotal field if non-nil, zero value otherwise.

### GetLastExecutionRowsTotalOk

`func (o *MaskingJob) GetLastExecutionRowsTotalOk() (*int64, bool)`

GetLastExecutionRowsTotalOk returns a tuple with the LastExecutionRowsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionRowsTotal

`func (o *MaskingJob) SetLastExecutionRowsTotal(v int64)`

SetLastExecutionRowsTotal sets LastExecutionRowsTotal field to given value.

### HasLastExecutionRowsTotal

`func (o *MaskingJob) HasLastExecutionRowsTotal() bool`

HasLastExecutionRowsTotal returns a boolean if a field has been set.

### GetLastExecutionBytesMasked

`func (o *MaskingJob) GetLastExecutionBytesMasked() int64`

GetLastExecutionBytesMasked returns the LastExecutionBytesMasked field if non-nil, zero value otherwise.

### GetLastExecutionBytesMaskedOk

`func (o *MaskingJob) GetLastExecutionBytesMaskedOk() (*int64, bool)`

GetLastExecutionBytesMaskedOk returns a tuple with the LastExecutionBytesMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionBytesMasked

`func (o *MaskingJob) SetLastExecutionBytesMasked(v int64)`

SetLastExecutionBytesMasked sets LastExecutionBytesMasked field to given value.

### HasLastExecutionBytesMasked

`func (o *MaskingJob) HasLastExecutionBytesMasked() bool`

HasLastExecutionBytesMasked returns a boolean if a field has been set.

### GetLastExecutionBytesTotal

`func (o *MaskingJob) GetLastExecutionBytesTotal() int64`

GetLastExecutionBytesTotal returns the LastExecutionBytesTotal field if non-nil, zero value otherwise.

### GetLastExecutionBytesTotalOk

`func (o *MaskingJob) GetLastExecutionBytesTotalOk() (*int64, bool)`

GetLastExecutionBytesTotalOk returns a tuple with the LastExecutionBytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionBytesTotal

`func (o *MaskingJob) SetLastExecutionBytesTotal(v int64)`

SetLastExecutionBytesTotal sets LastExecutionBytesTotal field to given value.

### HasLastExecutionBytesTotal

`func (o *MaskingJob) HasLastExecutionBytesTotal() bool`

HasLastExecutionBytesTotal returns a boolean if a field has been set.

### GetConnectorUsername

`func (o *MaskingJob) GetConnectorUsername() string`

GetConnectorUsername returns the ConnectorUsername field if non-nil, zero value otherwise.

### GetConnectorUsernameOk

`func (o *MaskingJob) GetConnectorUsernameOk() (*string, bool)`

GetConnectorUsernameOk returns a tuple with the ConnectorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUsername

`func (o *MaskingJob) SetConnectorUsername(v string)`

SetConnectorUsername sets ConnectorUsername field to given value.

### HasConnectorUsername

`func (o *MaskingJob) HasConnectorUsername() bool`

HasConnectorUsername returns a boolean if a field has been set.

### SetConnectorUsernameNil

`func (o *MaskingJob) SetConnectorUsernameNil(b bool)`

 SetConnectorUsernameNil sets the value for ConnectorUsername to be an explicit nil

### UnsetConnectorUsername
`func (o *MaskingJob) UnsetConnectorUsername()`

UnsetConnectorUsername ensures that no value is present for ConnectorUsername, not even an explicit nil
### GetConnectorPassword

`func (o *MaskingJob) GetConnectorPassword() string`

GetConnectorPassword returns the ConnectorPassword field if non-nil, zero value otherwise.

### GetConnectorPasswordOk

`func (o *MaskingJob) GetConnectorPasswordOk() (*string, bool)`

GetConnectorPasswordOk returns a tuple with the ConnectorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPassword

`func (o *MaskingJob) SetConnectorPassword(v string)`

SetConnectorPassword sets ConnectorPassword field to given value.

### HasConnectorPassword

`func (o *MaskingJob) HasConnectorPassword() bool`

HasConnectorPassword returns a boolean if a field has been set.

### SetConnectorPasswordNil

`func (o *MaskingJob) SetConnectorPasswordNil(b bool)`

 SetConnectorPasswordNil sets the value for ConnectorPassword to be an explicit nil

### UnsetConnectorPassword
`func (o *MaskingJob) UnsetConnectorPassword()`

UnsetConnectorPassword ensures that no value is present for ConnectorPassword, not even an explicit nil
### GetOnTheFlySourceConnectorUsername

`func (o *MaskingJob) GetOnTheFlySourceConnectorUsername() string`

GetOnTheFlySourceConnectorUsername returns the OnTheFlySourceConnectorUsername field if non-nil, zero value otherwise.

### GetOnTheFlySourceConnectorUsernameOk

`func (o *MaskingJob) GetOnTheFlySourceConnectorUsernameOk() (*string, bool)`

GetOnTheFlySourceConnectorUsernameOk returns a tuple with the OnTheFlySourceConnectorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTheFlySourceConnectorUsername

`func (o *MaskingJob) SetOnTheFlySourceConnectorUsername(v string)`

SetOnTheFlySourceConnectorUsername sets OnTheFlySourceConnectorUsername field to given value.

### HasOnTheFlySourceConnectorUsername

`func (o *MaskingJob) HasOnTheFlySourceConnectorUsername() bool`

HasOnTheFlySourceConnectorUsername returns a boolean if a field has been set.

### SetOnTheFlySourceConnectorUsernameNil

`func (o *MaskingJob) SetOnTheFlySourceConnectorUsernameNil(b bool)`

 SetOnTheFlySourceConnectorUsernameNil sets the value for OnTheFlySourceConnectorUsername to be an explicit nil

### UnsetOnTheFlySourceConnectorUsername
`func (o *MaskingJob) UnsetOnTheFlySourceConnectorUsername()`

UnsetOnTheFlySourceConnectorUsername ensures that no value is present for OnTheFlySourceConnectorUsername, not even an explicit nil
### GetOnTheFlySourceConnectorPassword

`func (o *MaskingJob) GetOnTheFlySourceConnectorPassword() string`

GetOnTheFlySourceConnectorPassword returns the OnTheFlySourceConnectorPassword field if non-nil, zero value otherwise.

### GetOnTheFlySourceConnectorPasswordOk

`func (o *MaskingJob) GetOnTheFlySourceConnectorPasswordOk() (*string, bool)`

GetOnTheFlySourceConnectorPasswordOk returns a tuple with the OnTheFlySourceConnectorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTheFlySourceConnectorPassword

`func (o *MaskingJob) SetOnTheFlySourceConnectorPassword(v string)`

SetOnTheFlySourceConnectorPassword sets OnTheFlySourceConnectorPassword field to given value.

### HasOnTheFlySourceConnectorPassword

`func (o *MaskingJob) HasOnTheFlySourceConnectorPassword() bool`

HasOnTheFlySourceConnectorPassword returns a boolean if a field has been set.

### SetOnTheFlySourceConnectorPasswordNil

`func (o *MaskingJob) SetOnTheFlySourceConnectorPasswordNil(b bool)`

 SetOnTheFlySourceConnectorPasswordNil sets the value for OnTheFlySourceConnectorPassword to be an explicit nil

### UnsetOnTheFlySourceConnectorPassword
`func (o *MaskingJob) UnsetOnTheFlySourceConnectorPassword()`

UnsetOnTheFlySourceConnectorPassword ensures that no value is present for OnTheFlySourceConnectorPassword, not even an explicit nil
### GetJobType

`func (o *MaskingJob) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *MaskingJob) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *MaskingJob) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *MaskingJob) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetHyperscaleInstanceId

`func (o *MaskingJob) GetHyperscaleInstanceId() string`

GetHyperscaleInstanceId returns the HyperscaleInstanceId field if non-nil, zero value otherwise.

### GetHyperscaleInstanceIdOk

`func (o *MaskingJob) GetHyperscaleInstanceIdOk() (*string, bool)`

GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperscaleInstanceId

`func (o *MaskingJob) SetHyperscaleInstanceId(v string)`

SetHyperscaleInstanceId sets HyperscaleInstanceId field to given value.

### HasHyperscaleInstanceId

`func (o *MaskingJob) HasHyperscaleInstanceId() bool`

HasHyperscaleInstanceId returns a boolean if a field has been set.

### GetDescription

`func (o *MaskingJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaskingJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaskingJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MaskingJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDatasetId

`func (o *MaskingJob) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *MaskingJob) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *MaskingJob) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *MaskingJob) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetRetainExecutionData

`func (o *MaskingJob) GetRetainExecutionData() string`

GetRetainExecutionData returns the RetainExecutionData field if non-nil, zero value otherwise.

### GetRetainExecutionDataOk

`func (o *MaskingJob) GetRetainExecutionDataOk() (*string, bool)`

GetRetainExecutionDataOk returns a tuple with the RetainExecutionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainExecutionData

`func (o *MaskingJob) SetRetainExecutionData(v string)`

SetRetainExecutionData sets RetainExecutionData field to given value.

### HasRetainExecutionData

`func (o *MaskingJob) HasRetainExecutionData() bool`

HasRetainExecutionData returns a boolean if a field has been set.

### GetMaxMemory

`func (o *MaskingJob) GetMaxMemory() int32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *MaskingJob) GetMaxMemoryOk() (*int32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *MaskingJob) SetMaxMemory(v int32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *MaskingJob) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMinMemory

`func (o *MaskingJob) GetMinMemory() int32`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *MaskingJob) GetMinMemoryOk() (*int32, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *MaskingJob) SetMinMemory(v int32)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *MaskingJob) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetFeedbackSize

`func (o *MaskingJob) GetFeedbackSize() int32`

GetFeedbackSize returns the FeedbackSize field if non-nil, zero value otherwise.

### GetFeedbackSizeOk

`func (o *MaskingJob) GetFeedbackSizeOk() (*int32, bool)`

GetFeedbackSizeOk returns a tuple with the FeedbackSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackSize

`func (o *MaskingJob) SetFeedbackSize(v int32)`

SetFeedbackSize sets FeedbackSize field to given value.

### HasFeedbackSize

`func (o *MaskingJob) HasFeedbackSize() bool`

HasFeedbackSize returns a boolean if a field has been set.

### GetStreamRowLimit

`func (o *MaskingJob) GetStreamRowLimit() int32`

GetStreamRowLimit returns the StreamRowLimit field if non-nil, zero value otherwise.

### GetStreamRowLimitOk

`func (o *MaskingJob) GetStreamRowLimitOk() (*int32, bool)`

GetStreamRowLimitOk returns a tuple with the StreamRowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamRowLimit

`func (o *MaskingJob) SetStreamRowLimit(v int32)`

SetStreamRowLimit sets StreamRowLimit field to given value.

### HasStreamRowLimit

`func (o *MaskingJob) HasStreamRowLimit() bool`

HasStreamRowLimit returns a boolean if a field has been set.

### GetNumInputStreams

`func (o *MaskingJob) GetNumInputStreams() int32`

GetNumInputStreams returns the NumInputStreams field if non-nil, zero value otherwise.

### GetNumInputStreamsOk

`func (o *MaskingJob) GetNumInputStreamsOk() (*int32, bool)`

GetNumInputStreamsOk returns a tuple with the NumInputStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInputStreams

`func (o *MaskingJob) SetNumInputStreams(v int32)`

SetNumInputStreams sets NumInputStreams field to given value.

### HasNumInputStreams

`func (o *MaskingJob) HasNumInputStreams() bool`

HasNumInputStreams returns a boolean if a field has been set.

### GetMaxConcurrentSourceConnections

`func (o *MaskingJob) GetMaxConcurrentSourceConnections() int32`

GetMaxConcurrentSourceConnections returns the MaxConcurrentSourceConnections field if non-nil, zero value otherwise.

### GetMaxConcurrentSourceConnectionsOk

`func (o *MaskingJob) GetMaxConcurrentSourceConnectionsOk() (*int32, bool)`

GetMaxConcurrentSourceConnectionsOk returns a tuple with the MaxConcurrentSourceConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentSourceConnections

`func (o *MaskingJob) SetMaxConcurrentSourceConnections(v int32)`

SetMaxConcurrentSourceConnections sets MaxConcurrentSourceConnections field to given value.

### HasMaxConcurrentSourceConnections

`func (o *MaskingJob) HasMaxConcurrentSourceConnections() bool`

HasMaxConcurrentSourceConnections returns a boolean if a field has been set.

### GetMaxConcurrentTargetConnections

`func (o *MaskingJob) GetMaxConcurrentTargetConnections() int32`

GetMaxConcurrentTargetConnections returns the MaxConcurrentTargetConnections field if non-nil, zero value otherwise.

### GetMaxConcurrentTargetConnectionsOk

`func (o *MaskingJob) GetMaxConcurrentTargetConnectionsOk() (*int32, bool)`

GetMaxConcurrentTargetConnectionsOk returns a tuple with the MaxConcurrentTargetConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentTargetConnections

`func (o *MaskingJob) SetMaxConcurrentTargetConnections(v int32)`

SetMaxConcurrentTargetConnections sets MaxConcurrentTargetConnections field to given value.

### HasMaxConcurrentTargetConnections

`func (o *MaskingJob) HasMaxConcurrentTargetConnections() bool`

HasMaxConcurrentTargetConnections returns a boolean if a field has been set.

### GetParallelismDegree

`func (o *MaskingJob) GetParallelismDegree() int32`

GetParallelismDegree returns the ParallelismDegree field if non-nil, zero value otherwise.

### GetParallelismDegreeOk

`func (o *MaskingJob) GetParallelismDegreeOk() (*int32, bool)`

GetParallelismDegreeOk returns a tuple with the ParallelismDegree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismDegree

`func (o *MaskingJob) SetParallelismDegree(v int32)`

SetParallelismDegree sets ParallelismDegree field to given value.

### HasParallelismDegree

`func (o *MaskingJob) HasParallelismDegree() bool`

HasParallelismDegree returns a boolean if a field has been set.

### GetTags

`func (o *MaskingJob) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MaskingJob) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MaskingJob) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MaskingJob) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


