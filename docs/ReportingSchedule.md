# ReportingSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **int32** |  | [optional] [readonly] 
**ReportType** | **string** |  | 
**CronExpression** | **string** | Standard cron expressions are supported e.g. 0 15 10 L * ?  - Schedule at 10:15 AM on the last day of every month, 0 0 2 ? * Mon-Fri - Schedule at 2:00 AM every Monday, Tuesday, Wednesday, Thursday and Friday. For more details kindly refer- \&quot;http://www.quartz-scheduler.org/documentation/\&quot; | 
**TimeZone** | Pointer to **string** | Timezones are specified according to the Olson tzinfo database - \&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot;. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**FileFormat** | **string** |  | 
**Enabled** | **bool** |  | [default to true]
**Recipients** | **[]string** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**SortColumn** | Pointer to **string** |  | [optional] 
**RowCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewReportingSchedule

`func NewReportingSchedule(reportType string, cronExpression string, fileFormat string, enabled bool, recipients []string, ) *ReportingSchedule`

NewReportingSchedule instantiates a new ReportingSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingScheduleWithDefaults

`func NewReportingScheduleWithDefaults() *ReportingSchedule`

NewReportingScheduleWithDefaults instantiates a new ReportingSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *ReportingSchedule) GetReportId() int32`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ReportingSchedule) GetReportIdOk() (*int32, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ReportingSchedule) SetReportId(v int32)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *ReportingSchedule) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetReportType

`func (o *ReportingSchedule) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportingSchedule) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportingSchedule) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetCronExpression

`func (o *ReportingSchedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ReportingSchedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ReportingSchedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetTimeZone

`func (o *ReportingSchedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ReportingSchedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ReportingSchedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ReportingSchedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetMessage

`func (o *ReportingSchedule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReportingSchedule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReportingSchedule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ReportingSchedule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFileFormat

`func (o *ReportingSchedule) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *ReportingSchedule) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *ReportingSchedule) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.


### GetEnabled

`func (o *ReportingSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReportingSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReportingSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRecipients

`func (o *ReportingSchedule) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ReportingSchedule) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ReportingSchedule) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.


### GetTags

`func (o *ReportingSchedule) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReportingSchedule) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReportingSchedule) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReportingSchedule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSortColumn

`func (o *ReportingSchedule) GetSortColumn() string`

GetSortColumn returns the SortColumn field if non-nil, zero value otherwise.

### GetSortColumnOk

`func (o *ReportingSchedule) GetSortColumnOk() (*string, bool)`

GetSortColumnOk returns a tuple with the SortColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortColumn

`func (o *ReportingSchedule) SetSortColumn(v string)`

SetSortColumn sets SortColumn field to given value.

### HasSortColumn

`func (o *ReportingSchedule) HasSortColumn() bool`

HasSortColumn returns a boolean if a field has been set.

### GetRowCount

`func (o *ReportingSchedule) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *ReportingSchedule) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *ReportingSchedule) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *ReportingSchedule) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


