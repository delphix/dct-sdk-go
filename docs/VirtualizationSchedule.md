# VirtualizationSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronString** | **string** |  | 
**CutoffTime** | **int64** |  | 

## Methods

### NewVirtualizationSchedule

`func NewVirtualizationSchedule(cronString string, cutoffTime int64, ) *VirtualizationSchedule`

NewVirtualizationSchedule instantiates a new VirtualizationSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationScheduleWithDefaults

`func NewVirtualizationScheduleWithDefaults() *VirtualizationSchedule`

NewVirtualizationScheduleWithDefaults instantiates a new VirtualizationSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronString

`func (o *VirtualizationSchedule) GetCronString() string`

GetCronString returns the CronString field if non-nil, zero value otherwise.

### GetCronStringOk

`func (o *VirtualizationSchedule) GetCronStringOk() (*string, bool)`

GetCronStringOk returns a tuple with the CronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronString

`func (o *VirtualizationSchedule) SetCronString(v string)`

SetCronString sets CronString field to given value.


### GetCutoffTime

`func (o *VirtualizationSchedule) GetCutoffTime() int64`

GetCutoffTime returns the CutoffTime field if non-nil, zero value otherwise.

### GetCutoffTimeOk

`func (o *VirtualizationSchedule) GetCutoffTimeOk() (*int64, bool)`

GetCutoffTimeOk returns a tuple with the CutoffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffTime

`func (o *VirtualizationSchedule) SetCutoffTime(v int64)`

SetCutoffTime sets CutoffTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


