# MaskingJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The MaskingJob entity ID. | [optional] 
**Name** | Pointer to **string** | The name of this MaskingJob. | [optional] 
**Ruleset** | Pointer to [**MaskingRuleset**](MaskingRuleset.md) |  | [optional] 
**IsOnTheFlyMasking** | Pointer to **bool** | Whether this is an on-the-fly masking job. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date this MaskingJob was created. | [optional] 
**LastCompletedExecutionDate** | Pointer to **time.Time** | The date this MaskingJob was last executed to completion. | [optional] 
**LastExecutionStatus** | Pointer to **string** | The status of this MaskingJob&#39;s last execution. | [optional] 
**ConnectorUsername** | Pointer to **NullableString** | The username of the Connector used by the MaskingJob. | [optional] 
**ConnectorPassword** | Pointer to **NullableString** | The password of the Connector used by the MaskingJob. | [optional] 
**OnTheFlySourceConnectorUsername** | Pointer to **NullableString** | The username of the source Connector used by the on-the-fly MaskingJob. | [optional] 
**OnTheFlySourceConnectorPassword** | Pointer to **NullableString** | The password of the source Connector used by the on-the-fly MaskingJob. | [optional] 
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


