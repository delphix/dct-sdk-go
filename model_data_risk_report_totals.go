/*
Delphix DCT API

Delphix DCT API

API version: 3.16.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the DataRiskReportTotals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataRiskReportTotals{}

// DataRiskReportTotals The global totals of masking data risk metrics across all connectors.
type DataRiskReportTotals struct {
	// The total number of connectors.
	ConnectorTotal *int64 `json:"connector_total,omitempty"`
	// The number of connectors that have been profiled.
	ProfiledConnectorCount *int64 `json:"profiled_connector_count,omitempty"`
	// The number of connectors found to have sensitive data.
	SensitiveConnectorCount *int64 `json:"sensitive_connector_count,omitempty"`
	// The number of connectors having had a successfully run masking job.
	MaskedConnectorCount *int64 `json:"masked_connector_count,omitempty"`
	// The number of connectors with sensitive data that have not been masked.
	AtRiskConnectorCount *int64 `json:"at_risk_connector_count,omitempty"`
	// The total number of data elements across connectors.
	DataElementsTotal *int64 `json:"data_elements_total,omitempty"`
	// The number of data elements categorized as not sensitive.
	DataElementsNotSensitive *int64 `json:"data_elements_not_sensitive,omitempty"`
	// The number of sensitive data elements that have been masked.
	DataElementsSensitiveMasked *int64 `json:"data_elements_sensitive_masked,omitempty"`
	// The number of sensitive data elements that have not been masked.
	DataElementsSensitiveUnmasked *int64 `json:"data_elements_sensitive_unmasked,omitempty"`
	// The total number of records across connectors.
	RecordsTotal *int64 `json:"records_total,omitempty"`
	// The number of records found to be not sensitive.
	RecordsNotSensitive *int64 `json:"records_not_sensitive,omitempty"`
	// The number of sensitive records that have been masked.
	RecordsSensitiveMasked *int64 `json:"records_sensitive_masked,omitempty"`
	// The number of sensitive records that have not been masked.
	RecordsSensitiveUnmasked *int64 `json:"records_sensitive_unmasked,omitempty"`
	// An explanation if the records coverage is not provided.
	RecordsCoverageMissingReason *string `json:"records_coverage_missing_reason,omitempty"`
}

// NewDataRiskReportTotals instantiates a new DataRiskReportTotals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataRiskReportTotals() *DataRiskReportTotals {
	this := DataRiskReportTotals{}
	return &this
}

// NewDataRiskReportTotalsWithDefaults instantiates a new DataRiskReportTotals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataRiskReportTotalsWithDefaults() *DataRiskReportTotals {
	this := DataRiskReportTotals{}
	return &this
}

// GetConnectorTotal returns the ConnectorTotal field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetConnectorTotal() int64 {
	if o == nil || IsNil(o.ConnectorTotal) {
		var ret int64
		return ret
	}
	return *o.ConnectorTotal
}

// GetConnectorTotalOk returns a tuple with the ConnectorTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetConnectorTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectorTotal) {
		return nil, false
	}
	return o.ConnectorTotal, true
}

// HasConnectorTotal returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasConnectorTotal() bool {
	if o != nil && !IsNil(o.ConnectorTotal) {
		return true
	}

	return false
}

// SetConnectorTotal gets a reference to the given int64 and assigns it to the ConnectorTotal field.
func (o *DataRiskReportTotals) SetConnectorTotal(v int64) {
	o.ConnectorTotal = &v
}

// GetProfiledConnectorCount returns the ProfiledConnectorCount field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetProfiledConnectorCount() int64 {
	if o == nil || IsNil(o.ProfiledConnectorCount) {
		var ret int64
		return ret
	}
	return *o.ProfiledConnectorCount
}

// GetProfiledConnectorCountOk returns a tuple with the ProfiledConnectorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetProfiledConnectorCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ProfiledConnectorCount) {
		return nil, false
	}
	return o.ProfiledConnectorCount, true
}

// HasProfiledConnectorCount returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasProfiledConnectorCount() bool {
	if o != nil && !IsNil(o.ProfiledConnectorCount) {
		return true
	}

	return false
}

// SetProfiledConnectorCount gets a reference to the given int64 and assigns it to the ProfiledConnectorCount field.
func (o *DataRiskReportTotals) SetProfiledConnectorCount(v int64) {
	o.ProfiledConnectorCount = &v
}

// GetSensitiveConnectorCount returns the SensitiveConnectorCount field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetSensitiveConnectorCount() int64 {
	if o == nil || IsNil(o.SensitiveConnectorCount) {
		var ret int64
		return ret
	}
	return *o.SensitiveConnectorCount
}

// GetSensitiveConnectorCountOk returns a tuple with the SensitiveConnectorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetSensitiveConnectorCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SensitiveConnectorCount) {
		return nil, false
	}
	return o.SensitiveConnectorCount, true
}

// HasSensitiveConnectorCount returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasSensitiveConnectorCount() bool {
	if o != nil && !IsNil(o.SensitiveConnectorCount) {
		return true
	}

	return false
}

// SetSensitiveConnectorCount gets a reference to the given int64 and assigns it to the SensitiveConnectorCount field.
func (o *DataRiskReportTotals) SetSensitiveConnectorCount(v int64) {
	o.SensitiveConnectorCount = &v
}

// GetMaskedConnectorCount returns the MaskedConnectorCount field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetMaskedConnectorCount() int64 {
	if o == nil || IsNil(o.MaskedConnectorCount) {
		var ret int64
		return ret
	}
	return *o.MaskedConnectorCount
}

// GetMaskedConnectorCountOk returns a tuple with the MaskedConnectorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetMaskedConnectorCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MaskedConnectorCount) {
		return nil, false
	}
	return o.MaskedConnectorCount, true
}

// HasMaskedConnectorCount returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasMaskedConnectorCount() bool {
	if o != nil && !IsNil(o.MaskedConnectorCount) {
		return true
	}

	return false
}

// SetMaskedConnectorCount gets a reference to the given int64 and assigns it to the MaskedConnectorCount field.
func (o *DataRiskReportTotals) SetMaskedConnectorCount(v int64) {
	o.MaskedConnectorCount = &v
}

// GetAtRiskConnectorCount returns the AtRiskConnectorCount field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetAtRiskConnectorCount() int64 {
	if o == nil || IsNil(o.AtRiskConnectorCount) {
		var ret int64
		return ret
	}
	return *o.AtRiskConnectorCount
}

// GetAtRiskConnectorCountOk returns a tuple with the AtRiskConnectorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetAtRiskConnectorCountOk() (*int64, bool) {
	if o == nil || IsNil(o.AtRiskConnectorCount) {
		return nil, false
	}
	return o.AtRiskConnectorCount, true
}

// HasAtRiskConnectorCount returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasAtRiskConnectorCount() bool {
	if o != nil && !IsNil(o.AtRiskConnectorCount) {
		return true
	}

	return false
}

// SetAtRiskConnectorCount gets a reference to the given int64 and assigns it to the AtRiskConnectorCount field.
func (o *DataRiskReportTotals) SetAtRiskConnectorCount(v int64) {
	o.AtRiskConnectorCount = &v
}

// GetDataElementsTotal returns the DataElementsTotal field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetDataElementsTotal() int64 {
	if o == nil || IsNil(o.DataElementsTotal) {
		var ret int64
		return ret
	}
	return *o.DataElementsTotal
}

// GetDataElementsTotalOk returns a tuple with the DataElementsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetDataElementsTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.DataElementsTotal) {
		return nil, false
	}
	return o.DataElementsTotal, true
}

// HasDataElementsTotal returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasDataElementsTotal() bool {
	if o != nil && !IsNil(o.DataElementsTotal) {
		return true
	}

	return false
}

// SetDataElementsTotal gets a reference to the given int64 and assigns it to the DataElementsTotal field.
func (o *DataRiskReportTotals) SetDataElementsTotal(v int64) {
	o.DataElementsTotal = &v
}

// GetDataElementsNotSensitive returns the DataElementsNotSensitive field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetDataElementsNotSensitive() int64 {
	if o == nil || IsNil(o.DataElementsNotSensitive) {
		var ret int64
		return ret
	}
	return *o.DataElementsNotSensitive
}

// GetDataElementsNotSensitiveOk returns a tuple with the DataElementsNotSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetDataElementsNotSensitiveOk() (*int64, bool) {
	if o == nil || IsNil(o.DataElementsNotSensitive) {
		return nil, false
	}
	return o.DataElementsNotSensitive, true
}

// HasDataElementsNotSensitive returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasDataElementsNotSensitive() bool {
	if o != nil && !IsNil(o.DataElementsNotSensitive) {
		return true
	}

	return false
}

// SetDataElementsNotSensitive gets a reference to the given int64 and assigns it to the DataElementsNotSensitive field.
func (o *DataRiskReportTotals) SetDataElementsNotSensitive(v int64) {
	o.DataElementsNotSensitive = &v
}

// GetDataElementsSensitiveMasked returns the DataElementsSensitiveMasked field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetDataElementsSensitiveMasked() int64 {
	if o == nil || IsNil(o.DataElementsSensitiveMasked) {
		var ret int64
		return ret
	}
	return *o.DataElementsSensitiveMasked
}

// GetDataElementsSensitiveMaskedOk returns a tuple with the DataElementsSensitiveMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetDataElementsSensitiveMaskedOk() (*int64, bool) {
	if o == nil || IsNil(o.DataElementsSensitiveMasked) {
		return nil, false
	}
	return o.DataElementsSensitiveMasked, true
}

// HasDataElementsSensitiveMasked returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasDataElementsSensitiveMasked() bool {
	if o != nil && !IsNil(o.DataElementsSensitiveMasked) {
		return true
	}

	return false
}

// SetDataElementsSensitiveMasked gets a reference to the given int64 and assigns it to the DataElementsSensitiveMasked field.
func (o *DataRiskReportTotals) SetDataElementsSensitiveMasked(v int64) {
	o.DataElementsSensitiveMasked = &v
}

// GetDataElementsSensitiveUnmasked returns the DataElementsSensitiveUnmasked field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetDataElementsSensitiveUnmasked() int64 {
	if o == nil || IsNil(o.DataElementsSensitiveUnmasked) {
		var ret int64
		return ret
	}
	return *o.DataElementsSensitiveUnmasked
}

// GetDataElementsSensitiveUnmaskedOk returns a tuple with the DataElementsSensitiveUnmasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetDataElementsSensitiveUnmaskedOk() (*int64, bool) {
	if o == nil || IsNil(o.DataElementsSensitiveUnmasked) {
		return nil, false
	}
	return o.DataElementsSensitiveUnmasked, true
}

// HasDataElementsSensitiveUnmasked returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasDataElementsSensitiveUnmasked() bool {
	if o != nil && !IsNil(o.DataElementsSensitiveUnmasked) {
		return true
	}

	return false
}

// SetDataElementsSensitiveUnmasked gets a reference to the given int64 and assigns it to the DataElementsSensitiveUnmasked field.
func (o *DataRiskReportTotals) SetDataElementsSensitiveUnmasked(v int64) {
	o.DataElementsSensitiveUnmasked = &v
}

// GetRecordsTotal returns the RecordsTotal field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetRecordsTotal() int64 {
	if o == nil || IsNil(o.RecordsTotal) {
		var ret int64
		return ret
	}
	return *o.RecordsTotal
}

// GetRecordsTotalOk returns a tuple with the RecordsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetRecordsTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.RecordsTotal) {
		return nil, false
	}
	return o.RecordsTotal, true
}

// HasRecordsTotal returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasRecordsTotal() bool {
	if o != nil && !IsNil(o.RecordsTotal) {
		return true
	}

	return false
}

// SetRecordsTotal gets a reference to the given int64 and assigns it to the RecordsTotal field.
func (o *DataRiskReportTotals) SetRecordsTotal(v int64) {
	o.RecordsTotal = &v
}

// GetRecordsNotSensitive returns the RecordsNotSensitive field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetRecordsNotSensitive() int64 {
	if o == nil || IsNil(o.RecordsNotSensitive) {
		var ret int64
		return ret
	}
	return *o.RecordsNotSensitive
}

// GetRecordsNotSensitiveOk returns a tuple with the RecordsNotSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetRecordsNotSensitiveOk() (*int64, bool) {
	if o == nil || IsNil(o.RecordsNotSensitive) {
		return nil, false
	}
	return o.RecordsNotSensitive, true
}

// HasRecordsNotSensitive returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasRecordsNotSensitive() bool {
	if o != nil && !IsNil(o.RecordsNotSensitive) {
		return true
	}

	return false
}

// SetRecordsNotSensitive gets a reference to the given int64 and assigns it to the RecordsNotSensitive field.
func (o *DataRiskReportTotals) SetRecordsNotSensitive(v int64) {
	o.RecordsNotSensitive = &v
}

// GetRecordsSensitiveMasked returns the RecordsSensitiveMasked field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetRecordsSensitiveMasked() int64 {
	if o == nil || IsNil(o.RecordsSensitiveMasked) {
		var ret int64
		return ret
	}
	return *o.RecordsSensitiveMasked
}

// GetRecordsSensitiveMaskedOk returns a tuple with the RecordsSensitiveMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetRecordsSensitiveMaskedOk() (*int64, bool) {
	if o == nil || IsNil(o.RecordsSensitiveMasked) {
		return nil, false
	}
	return o.RecordsSensitiveMasked, true
}

// HasRecordsSensitiveMasked returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasRecordsSensitiveMasked() bool {
	if o != nil && !IsNil(o.RecordsSensitiveMasked) {
		return true
	}

	return false
}

// SetRecordsSensitiveMasked gets a reference to the given int64 and assigns it to the RecordsSensitiveMasked field.
func (o *DataRiskReportTotals) SetRecordsSensitiveMasked(v int64) {
	o.RecordsSensitiveMasked = &v
}

// GetRecordsSensitiveUnmasked returns the RecordsSensitiveUnmasked field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetRecordsSensitiveUnmasked() int64 {
	if o == nil || IsNil(o.RecordsSensitiveUnmasked) {
		var ret int64
		return ret
	}
	return *o.RecordsSensitiveUnmasked
}

// GetRecordsSensitiveUnmaskedOk returns a tuple with the RecordsSensitiveUnmasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetRecordsSensitiveUnmaskedOk() (*int64, bool) {
	if o == nil || IsNil(o.RecordsSensitiveUnmasked) {
		return nil, false
	}
	return o.RecordsSensitiveUnmasked, true
}

// HasRecordsSensitiveUnmasked returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasRecordsSensitiveUnmasked() bool {
	if o != nil && !IsNil(o.RecordsSensitiveUnmasked) {
		return true
	}

	return false
}

// SetRecordsSensitiveUnmasked gets a reference to the given int64 and assigns it to the RecordsSensitiveUnmasked field.
func (o *DataRiskReportTotals) SetRecordsSensitiveUnmasked(v int64) {
	o.RecordsSensitiveUnmasked = &v
}

// GetRecordsCoverageMissingReason returns the RecordsCoverageMissingReason field value if set, zero value otherwise.
func (o *DataRiskReportTotals) GetRecordsCoverageMissingReason() string {
	if o == nil || IsNil(o.RecordsCoverageMissingReason) {
		var ret string
		return ret
	}
	return *o.RecordsCoverageMissingReason
}

// GetRecordsCoverageMissingReasonOk returns a tuple with the RecordsCoverageMissingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRiskReportTotals) GetRecordsCoverageMissingReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RecordsCoverageMissingReason) {
		return nil, false
	}
	return o.RecordsCoverageMissingReason, true
}

// HasRecordsCoverageMissingReason returns a boolean if a field has been set.
func (o *DataRiskReportTotals) HasRecordsCoverageMissingReason() bool {
	if o != nil && !IsNil(o.RecordsCoverageMissingReason) {
		return true
	}

	return false
}

// SetRecordsCoverageMissingReason gets a reference to the given string and assigns it to the RecordsCoverageMissingReason field.
func (o *DataRiskReportTotals) SetRecordsCoverageMissingReason(v string) {
	o.RecordsCoverageMissingReason = &v
}

func (o DataRiskReportTotals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataRiskReportTotals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectorTotal) {
		toSerialize["connector_total"] = o.ConnectorTotal
	}
	if !IsNil(o.ProfiledConnectorCount) {
		toSerialize["profiled_connector_count"] = o.ProfiledConnectorCount
	}
	if !IsNil(o.SensitiveConnectorCount) {
		toSerialize["sensitive_connector_count"] = o.SensitiveConnectorCount
	}
	if !IsNil(o.MaskedConnectorCount) {
		toSerialize["masked_connector_count"] = o.MaskedConnectorCount
	}
	if !IsNil(o.AtRiskConnectorCount) {
		toSerialize["at_risk_connector_count"] = o.AtRiskConnectorCount
	}
	if !IsNil(o.DataElementsTotal) {
		toSerialize["data_elements_total"] = o.DataElementsTotal
	}
	if !IsNil(o.DataElementsNotSensitive) {
		toSerialize["data_elements_not_sensitive"] = o.DataElementsNotSensitive
	}
	if !IsNil(o.DataElementsSensitiveMasked) {
		toSerialize["data_elements_sensitive_masked"] = o.DataElementsSensitiveMasked
	}
	if !IsNil(o.DataElementsSensitiveUnmasked) {
		toSerialize["data_elements_sensitive_unmasked"] = o.DataElementsSensitiveUnmasked
	}
	if !IsNil(o.RecordsTotal) {
		toSerialize["records_total"] = o.RecordsTotal
	}
	if !IsNil(o.RecordsNotSensitive) {
		toSerialize["records_not_sensitive"] = o.RecordsNotSensitive
	}
	if !IsNil(o.RecordsSensitiveMasked) {
		toSerialize["records_sensitive_masked"] = o.RecordsSensitiveMasked
	}
	if !IsNil(o.RecordsSensitiveUnmasked) {
		toSerialize["records_sensitive_unmasked"] = o.RecordsSensitiveUnmasked
	}
	if !IsNil(o.RecordsCoverageMissingReason) {
		toSerialize["records_coverage_missing_reason"] = o.RecordsCoverageMissingReason
	}
	return toSerialize, nil
}

type NullableDataRiskReportTotals struct {
	value *DataRiskReportTotals
	isSet bool
}

func (v NullableDataRiskReportTotals) Get() *DataRiskReportTotals {
	return v.value
}

func (v *NullableDataRiskReportTotals) Set(val *DataRiskReportTotals) {
	v.value = val
	v.isSet = true
}

func (v NullableDataRiskReportTotals) IsSet() bool {
	return v.isSet
}

func (v *NullableDataRiskReportTotals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataRiskReportTotals(val *DataRiskReportTotals) *NullableDataRiskReportTotals {
	return &NullableDataRiskReportTotals{value: val, isSet: true}
}

func (v NullableDataRiskReportTotals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataRiskReportTotals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


