# ReportingScheduleCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** |  | 
**CronExpression** | **string** | Standard cron expressions are supported e.g. 0 15 10 L * ?  - Schedule at 10:15 AM on the last day of every month, 0 0 2 ? * Mon-Fri - Schedule at 2:00 AM every Monday, Tuesday, Wednesday, Thursday and Friday. For more details kindly refer- \&quot;http://www.quartz-scheduler.org/documentation/\&quot; | 
**TimeZone** | Pointer to **string** | Timezones are specified according to the Olson tzinfo database - \&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot;. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**FileFormat** | **string** |  | 
**Enabled** | **bool** |  | [default to true]
**Recipients** | **[]string** |  | 
**SortColumn** | Pointer to **string** |  | [optional] 
**RowCount** | Pointer to **int32** |  | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this reporting schedule must be configured as owner of the reporting schedule. | [optional] [default to true]

## Methods

### NewReportingScheduleCreateParameters

`func NewReportingScheduleCreateParameters(reportType string, cronExpression string, fileFormat string, enabled bool, recipients []string, ) *ReportingScheduleCreateParameters`

NewReportingScheduleCreateParameters instantiates a new ReportingScheduleCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingScheduleCreateParametersWithDefaults

`func NewReportingScheduleCreateParametersWithDefaults() *ReportingScheduleCreateParameters`

NewReportingScheduleCreateParametersWithDefaults instantiates a new ReportingScheduleCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *ReportingScheduleCreateParameters) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportingScheduleCreateParameters) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportingScheduleCreateParameters) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetCronExpression

`func (o *ReportingScheduleCreateParameters) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ReportingScheduleCreateParameters) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ReportingScheduleCreateParameters) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetTimeZone

`func (o *ReportingScheduleCreateParameters) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ReportingScheduleCreateParameters) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ReportingScheduleCreateParameters) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ReportingScheduleCreateParameters) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetMessage

`func (o *ReportingScheduleCreateParameters) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReportingScheduleCreateParameters) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReportingScheduleCreateParameters) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ReportingScheduleCreateParameters) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFileFormat

`func (o *ReportingScheduleCreateParameters) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *ReportingScheduleCreateParameters) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *ReportingScheduleCreateParameters) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.


### GetEnabled

`func (o *ReportingScheduleCreateParameters) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReportingScheduleCreateParameters) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReportingScheduleCreateParameters) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRecipients

`func (o *ReportingScheduleCreateParameters) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ReportingScheduleCreateParameters) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ReportingScheduleCreateParameters) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.


### GetSortColumn

`func (o *ReportingScheduleCreateParameters) GetSortColumn() string`

GetSortColumn returns the SortColumn field if non-nil, zero value otherwise.

### GetSortColumnOk

`func (o *ReportingScheduleCreateParameters) GetSortColumnOk() (*string, bool)`

GetSortColumnOk returns a tuple with the SortColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortColumn

`func (o *ReportingScheduleCreateParameters) SetSortColumn(v string)`

SetSortColumn sets SortColumn field to given value.

### HasSortColumn

`func (o *ReportingScheduleCreateParameters) HasSortColumn() bool`

HasSortColumn returns a boolean if a field has been set.

### GetRowCount

`func (o *ReportingScheduleCreateParameters) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *ReportingScheduleCreateParameters) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *ReportingScheduleCreateParameters) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *ReportingScheduleCreateParameters) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *ReportingScheduleCreateParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ReportingScheduleCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ReportingScheduleCreateParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ReportingScheduleCreateParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


