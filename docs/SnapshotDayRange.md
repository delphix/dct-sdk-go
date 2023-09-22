# SnapshotDayRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Number of TimeFlow snapshots on that day. | [optional] 
**Date** | Pointer to **string** | Date for which TimeFlow snapshots have been aggregated. | [optional] 
**StartOfDay** | Pointer to **time.Time** | Start of day of this range in the time zone used for computation. | [optional] 
**EndOfDay** | Pointer to **time.Time** | End of day of this range in the time zone used for computation. | [optional] 

## Methods

### NewSnapshotDayRange

`func NewSnapshotDayRange() *SnapshotDayRange`

NewSnapshotDayRange instantiates a new SnapshotDayRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotDayRangeWithDefaults

`func NewSnapshotDayRangeWithDefaults() *SnapshotDayRange`

NewSnapshotDayRangeWithDefaults instantiates a new SnapshotDayRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SnapshotDayRange) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SnapshotDayRange) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SnapshotDayRange) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SnapshotDayRange) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDate

`func (o *SnapshotDayRange) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SnapshotDayRange) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SnapshotDayRange) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *SnapshotDayRange) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStartOfDay

`func (o *SnapshotDayRange) GetStartOfDay() time.Time`

GetStartOfDay returns the StartOfDay field if non-nil, zero value otherwise.

### GetStartOfDayOk

`func (o *SnapshotDayRange) GetStartOfDayOk() (*time.Time, bool)`

GetStartOfDayOk returns a tuple with the StartOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOfDay

`func (o *SnapshotDayRange) SetStartOfDay(v time.Time)`

SetStartOfDay sets StartOfDay field to given value.

### HasStartOfDay

`func (o *SnapshotDayRange) HasStartOfDay() bool`

HasStartOfDay returns a boolean if a field has been set.

### GetEndOfDay

`func (o *SnapshotDayRange) GetEndOfDay() time.Time`

GetEndOfDay returns the EndOfDay field if non-nil, zero value otherwise.

### GetEndOfDayOk

`func (o *SnapshotDayRange) GetEndOfDayOk() (*time.Time, bool)`

GetEndOfDayOk returns a tuple with the EndOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfDay

`func (o *SnapshotDayRange) SetEndOfDay(v time.Time)`

SetEndOfDay sets EndOfDay field to given value.

### HasEndOfDay

`func (o *SnapshotDayRange) HasEndOfDay() bool`

HasEndOfDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


