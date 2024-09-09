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

// checks if the AuditLogsSummaryReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogsSummaryReportResponse{}

// AuditLogsSummaryReportResponse struct for AuditLogsSummaryReportResponse
type AuditLogsSummaryReportResponse struct {
	Items []AuditLogsSummary `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
	Totals *AuditLogsSummaryTotals `json:"totals,omitempty"`
}

// NewAuditLogsSummaryReportResponse instantiates a new AuditLogsSummaryReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogsSummaryReportResponse() *AuditLogsSummaryReportResponse {
	this := AuditLogsSummaryReportResponse{}
	return &this
}

// NewAuditLogsSummaryReportResponseWithDefaults instantiates a new AuditLogsSummaryReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogsSummaryReportResponseWithDefaults() *AuditLogsSummaryReportResponse {
	this := AuditLogsSummaryReportResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AuditLogsSummaryReportResponse) GetItems() []AuditLogsSummary {
	if o == nil || IsNil(o.Items) {
		var ret []AuditLogsSummary
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogsSummaryReportResponse) GetItemsOk() ([]AuditLogsSummary, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AuditLogsSummaryReportResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AuditLogsSummary and assigns it to the Items field.
func (o *AuditLogsSummaryReportResponse) SetItems(v []AuditLogsSummary) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *AuditLogsSummaryReportResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogsSummaryReportResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *AuditLogsSummaryReportResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *AuditLogsSummaryReportResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

// GetTotals returns the Totals field value if set, zero value otherwise.
func (o *AuditLogsSummaryReportResponse) GetTotals() AuditLogsSummaryTotals {
	if o == nil || IsNil(o.Totals) {
		var ret AuditLogsSummaryTotals
		return ret
	}
	return *o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogsSummaryReportResponse) GetTotalsOk() (*AuditLogsSummaryTotals, bool) {
	if o == nil || IsNil(o.Totals) {
		return nil, false
	}
	return o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *AuditLogsSummaryReportResponse) HasTotals() bool {
	if o != nil && !IsNil(o.Totals) {
		return true
	}

	return false
}

// SetTotals gets a reference to the given AuditLogsSummaryTotals and assigns it to the Totals field.
func (o *AuditLogsSummaryReportResponse) SetTotals(v AuditLogsSummaryTotals) {
	o.Totals = &v
}

func (o AuditLogsSummaryReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogsSummaryReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	if !IsNil(o.Totals) {
		toSerialize["totals"] = o.Totals
	}
	return toSerialize, nil
}

type NullableAuditLogsSummaryReportResponse struct {
	value *AuditLogsSummaryReportResponse
	isSet bool
}

func (v NullableAuditLogsSummaryReportResponse) Get() *AuditLogsSummaryReportResponse {
	return v.value
}

func (v *NullableAuditLogsSummaryReportResponse) Set(val *AuditLogsSummaryReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogsSummaryReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogsSummaryReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogsSummaryReportResponse(val *AuditLogsSummaryReportResponse) *NullableAuditLogsSummaryReportResponse {
	return &NullableAuditLogsSummaryReportResponse{value: val, isSet: true}
}

func (v NullableAuditLogsSummaryReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogsSummaryReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


