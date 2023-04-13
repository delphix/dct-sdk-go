/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ReportingScheduleCreateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingScheduleCreateParameters{}

// ReportingScheduleCreateParameters struct for ReportingScheduleCreateParameters
type ReportingScheduleCreateParameters struct {
	ReportType string `json:"report_type"`
	// Standard cron expressions are supported e.g. 0 15 10 L * ?  - Schedule at 10:15 AM on the last day of every month, 0 0 2 ? * Mon-Fri - Schedule at 2:00 AM every Monday, Tuesday, Wednesday, Thursday and Friday. For more details kindly refer- \"http://www.quartz-scheduler.org/documentation/\"
	CronExpression string `json:"cron_expression"`
	// Timezones are specified according to the Olson tzinfo database - \"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\".
	TimeZone *string `json:"time_zone,omitempty"`
	Message *string `json:"message,omitempty"`
	FileFormat string `json:"file_format"`
	Enabled bool `json:"enabled"`
	Recipients []string `json:"recipients"`
	SortColumn *string `json:"sort_column,omitempty"`
	RowCount *int32 `json:"row_count,omitempty"`
	// Whether the account creating this reporting schedule must be configured as owner of the reporting schedule.
	MakeCurrentAccountOwner *bool `json:"make_current_account_owner,omitempty"`
}

// NewReportingScheduleCreateParameters instantiates a new ReportingScheduleCreateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingScheduleCreateParameters(reportType string, cronExpression string, fileFormat string, enabled bool, recipients []string) *ReportingScheduleCreateParameters {
	this := ReportingScheduleCreateParameters{}
	this.ReportType = reportType
	this.CronExpression = cronExpression
	this.FileFormat = fileFormat
	this.Enabled = enabled
	this.Recipients = recipients
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// NewReportingScheduleCreateParametersWithDefaults instantiates a new ReportingScheduleCreateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingScheduleCreateParametersWithDefaults() *ReportingScheduleCreateParameters {
	this := ReportingScheduleCreateParameters{}
	var enabled bool = true
	this.Enabled = enabled
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// GetReportType returns the ReportType field value
func (o *ReportingScheduleCreateParameters) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *ReportingScheduleCreateParameters) SetReportType(v string) {
	o.ReportType = v
}

// GetCronExpression returns the CronExpression field value
func (o *ReportingScheduleCreateParameters) GetCronExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetCronExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronExpression, true
}

// SetCronExpression sets field value
func (o *ReportingScheduleCreateParameters) SetCronExpression(v string) {
	o.CronExpression = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ReportingScheduleCreateParameters) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ReportingScheduleCreateParameters) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ReportingScheduleCreateParameters) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReportingScheduleCreateParameters) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReportingScheduleCreateParameters) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReportingScheduleCreateParameters) SetMessage(v string) {
	o.Message = &v
}

// GetFileFormat returns the FileFormat field value
func (o *ReportingScheduleCreateParameters) GetFileFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetFileFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileFormat, true
}

// SetFileFormat sets field value
func (o *ReportingScheduleCreateParameters) SetFileFormat(v string) {
	o.FileFormat = v
}

// GetEnabled returns the Enabled field value
func (o *ReportingScheduleCreateParameters) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReportingScheduleCreateParameters) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRecipients returns the Recipients field value
func (o *ReportingScheduleCreateParameters) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetRecipientsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ReportingScheduleCreateParameters) SetRecipients(v []string) {
	o.Recipients = v
}

// GetSortColumn returns the SortColumn field value if set, zero value otherwise.
func (o *ReportingScheduleCreateParameters) GetSortColumn() string {
	if o == nil || IsNil(o.SortColumn) {
		var ret string
		return ret
	}
	return *o.SortColumn
}

// GetSortColumnOk returns a tuple with the SortColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetSortColumnOk() (*string, bool) {
	if o == nil || IsNil(o.SortColumn) {
		return nil, false
	}
	return o.SortColumn, true
}

// HasSortColumn returns a boolean if a field has been set.
func (o *ReportingScheduleCreateParameters) HasSortColumn() bool {
	if o != nil && !IsNil(o.SortColumn) {
		return true
	}

	return false
}

// SetSortColumn gets a reference to the given string and assigns it to the SortColumn field.
func (o *ReportingScheduleCreateParameters) SetSortColumn(v string) {
	o.SortColumn = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *ReportingScheduleCreateParameters) GetRowCount() int32 {
	if o == nil || IsNil(o.RowCount) {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetRowCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RowCount) {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *ReportingScheduleCreateParameters) HasRowCount() bool {
	if o != nil && !IsNil(o.RowCount) {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *ReportingScheduleCreateParameters) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field value if set, zero value otherwise.
func (o *ReportingScheduleCreateParameters) GetMakeCurrentAccountOwner() bool {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		var ret bool
		return ret
	}
	return *o.MakeCurrentAccountOwner
}

// GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingScheduleCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		return nil, false
	}
	return o.MakeCurrentAccountOwner, true
}

// HasMakeCurrentAccountOwner returns a boolean if a field has been set.
func (o *ReportingScheduleCreateParameters) HasMakeCurrentAccountOwner() bool {
	if o != nil && !IsNil(o.MakeCurrentAccountOwner) {
		return true
	}

	return false
}

// SetMakeCurrentAccountOwner gets a reference to the given bool and assigns it to the MakeCurrentAccountOwner field.
func (o *ReportingScheduleCreateParameters) SetMakeCurrentAccountOwner(v bool) {
	o.MakeCurrentAccountOwner = &v
}

func (o ReportingScheduleCreateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingScheduleCreateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["report_type"] = o.ReportType
	toSerialize["cron_expression"] = o.CronExpression
	if !IsNil(o.TimeZone) {
		toSerialize["time_zone"] = o.TimeZone
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["file_format"] = o.FileFormat
	toSerialize["enabled"] = o.Enabled
	toSerialize["recipients"] = o.Recipients
	if !IsNil(o.SortColumn) {
		toSerialize["sort_column"] = o.SortColumn
	}
	if !IsNil(o.RowCount) {
		toSerialize["row_count"] = o.RowCount
	}
	if !IsNil(o.MakeCurrentAccountOwner) {
		toSerialize["make_current_account_owner"] = o.MakeCurrentAccountOwner
	}
	return toSerialize, nil
}

type NullableReportingScheduleCreateParameters struct {
	value *ReportingScheduleCreateParameters
	isSet bool
}

func (v NullableReportingScheduleCreateParameters) Get() *ReportingScheduleCreateParameters {
	return v.value
}

func (v *NullableReportingScheduleCreateParameters) Set(val *ReportingScheduleCreateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingScheduleCreateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingScheduleCreateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingScheduleCreateParameters(val *ReportingScheduleCreateParameters) *NullableReportingScheduleCreateParameters {
	return &NullableReportingScheduleCreateParameters{value: val, isSet: true}
}

func (v NullableReportingScheduleCreateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingScheduleCreateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


