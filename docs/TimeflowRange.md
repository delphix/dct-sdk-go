# TimeflowRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartPointLocation** | Pointer to **string** | The starting location of this range. | [optional] 
**StartPointTimestamp** | Pointer to **time.Time** | The starting timestamp of this range. | [optional] 
**EndPointLocation** | Pointer to **string** | The ending location of this range. | [optional] 
**EndPointTimestamp** | Pointer to **time.Time** | The ending timestamp of this range. | [optional] 
**TimeflowId** | Pointer to **string** | A reference to the timeflow of this range. | [optional] 
**Provisionable** | Pointer to **bool** | Whether or not this range is provisionable. | [optional] 

## Methods

### NewTimeflowRange

`func NewTimeflowRange() *TimeflowRange`

NewTimeflowRange instantiates a new TimeflowRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeflowRangeWithDefaults

`func NewTimeflowRangeWithDefaults() *TimeflowRange`

NewTimeflowRangeWithDefaults instantiates a new TimeflowRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartPointLocation

`func (o *TimeflowRange) GetStartPointLocation() string`

GetStartPointLocation returns the StartPointLocation field if non-nil, zero value otherwise.

### GetStartPointLocationOk

`func (o *TimeflowRange) GetStartPointLocationOk() (*string, bool)`

GetStartPointLocationOk returns a tuple with the StartPointLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointLocation

`func (o *TimeflowRange) SetStartPointLocation(v string)`

SetStartPointLocation sets StartPointLocation field to given value.

### HasStartPointLocation

`func (o *TimeflowRange) HasStartPointLocation() bool`

HasStartPointLocation returns a boolean if a field has been set.

### GetStartPointTimestamp

`func (o *TimeflowRange) GetStartPointTimestamp() time.Time`

GetStartPointTimestamp returns the StartPointTimestamp field if non-nil, zero value otherwise.

### GetStartPointTimestampOk

`func (o *TimeflowRange) GetStartPointTimestampOk() (*time.Time, bool)`

GetStartPointTimestampOk returns a tuple with the StartPointTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointTimestamp

`func (o *TimeflowRange) SetStartPointTimestamp(v time.Time)`

SetStartPointTimestamp sets StartPointTimestamp field to given value.

### HasStartPointTimestamp

`func (o *TimeflowRange) HasStartPointTimestamp() bool`

HasStartPointTimestamp returns a boolean if a field has been set.

### GetEndPointLocation

`func (o *TimeflowRange) GetEndPointLocation() string`

GetEndPointLocation returns the EndPointLocation field if non-nil, zero value otherwise.

### GetEndPointLocationOk

`func (o *TimeflowRange) GetEndPointLocationOk() (*string, bool)`

GetEndPointLocationOk returns a tuple with the EndPointLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointLocation

`func (o *TimeflowRange) SetEndPointLocation(v string)`

SetEndPointLocation sets EndPointLocation field to given value.

### HasEndPointLocation

`func (o *TimeflowRange) HasEndPointLocation() bool`

HasEndPointLocation returns a boolean if a field has been set.

### GetEndPointTimestamp

`func (o *TimeflowRange) GetEndPointTimestamp() time.Time`

GetEndPointTimestamp returns the EndPointTimestamp field if non-nil, zero value otherwise.

### GetEndPointTimestampOk

`func (o *TimeflowRange) GetEndPointTimestampOk() (*time.Time, bool)`

GetEndPointTimestampOk returns a tuple with the EndPointTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointTimestamp

`func (o *TimeflowRange) SetEndPointTimestamp(v time.Time)`

SetEndPointTimestamp sets EndPointTimestamp field to given value.

### HasEndPointTimestamp

`func (o *TimeflowRange) HasEndPointTimestamp() bool`

HasEndPointTimestamp returns a boolean if a field has been set.

### GetTimeflowId

`func (o *TimeflowRange) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *TimeflowRange) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *TimeflowRange) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *TimeflowRange) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.

### GetProvisionable

`func (o *TimeflowRange) GetProvisionable() bool`

GetProvisionable returns the Provisionable field if non-nil, zero value otherwise.

### GetProvisionableOk

`func (o *TimeflowRange) GetProvisionableOk() (*bool, bool)`

GetProvisionableOk returns a tuple with the Provisionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionable

`func (o *TimeflowRange) SetProvisionable(v bool)`

SetProvisionable sets Provisionable field to given value.

### HasProvisionable

`func (o *TimeflowRange) HasProvisionable() bool`

HasProvisionable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


