/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"time"
)

// checks if the DatabaseTableMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseTableMetadata{}

// DatabaseTableMetadata An object describe a database table of a rule set with some settings.
type DatabaseTableMetadata struct {
	// The id of this table metadata.
	Id *int64 `json:"id,omitempty"`
	// The name of the table.
	TableName *string `json:"table_name,omitempty"`
	// The id of the rule set that this table metadata belongs to.
	RuleSetId *string `json:"rule_set_id,omitempty"`
	// The name of the rule set that this table metadata belongs to.
	RuleSetName *string `json:"rule_set_name,omitempty"`
	// Custom SQL for the table.
	CustomSql *string `json:"custom_sql,omitempty"`
	// SQL where clause for the table.
	WhereClause *string `json:"where_clause,omitempty"`
	// SQL having clause for the table.
	HavingClause *string `json:"having_clause,omitempty"`
	// Key Column for the table.
	KeyColumn *string `json:"key_column,omitempty"`
	// This field indicates whether or not a table has sensitive data. This field is assigned by DCT to true or false based on whether the table is assigned an algorithm and domain.
	IsSensitive *bool `json:"is_sensitive,omitempty"`
	RowCount *int64 `json:"row_count,omitempty"`
	LastRefreshTime *time.Time `json:"last_refresh_time,omitempty"`
	LastRowCountTime *time.Time `json:"last_row_count_time,omitempty"`
	// The id of the engine associated with this column.
	EngineId *string `json:"engine_id,omitempty"`
	// The name of the engine associated with this column.
	EngineName *string `json:"engine_name,omitempty"`
}

// NewDatabaseTableMetadata instantiates a new DatabaseTableMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseTableMetadata() *DatabaseTableMetadata {
	this := DatabaseTableMetadata{}
	return &this
}

// NewDatabaseTableMetadataWithDefaults instantiates a new DatabaseTableMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseTableMetadataWithDefaults() *DatabaseTableMetadata {
	this := DatabaseTableMetadata{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DatabaseTableMetadata) SetId(v int64) {
	o.Id = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *DatabaseTableMetadata) SetTableName(v string) {
	o.TableName = &v
}

// GetRuleSetId returns the RuleSetId field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetRuleSetId() string {
	if o == nil || IsNil(o.RuleSetId) {
		var ret string
		return ret
	}
	return *o.RuleSetId
}

// GetRuleSetIdOk returns a tuple with the RuleSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetRuleSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleSetId) {
		return nil, false
	}
	return o.RuleSetId, true
}

// HasRuleSetId returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasRuleSetId() bool {
	if o != nil && !IsNil(o.RuleSetId) {
		return true
	}

	return false
}

// SetRuleSetId gets a reference to the given string and assigns it to the RuleSetId field.
func (o *DatabaseTableMetadata) SetRuleSetId(v string) {
	o.RuleSetId = &v
}

// GetRuleSetName returns the RuleSetName field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetRuleSetName() string {
	if o == nil || IsNil(o.RuleSetName) {
		var ret string
		return ret
	}
	return *o.RuleSetName
}

// GetRuleSetNameOk returns a tuple with the RuleSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetRuleSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleSetName) {
		return nil, false
	}
	return o.RuleSetName, true
}

// HasRuleSetName returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasRuleSetName() bool {
	if o != nil && !IsNil(o.RuleSetName) {
		return true
	}

	return false
}

// SetRuleSetName gets a reference to the given string and assigns it to the RuleSetName field.
func (o *DatabaseTableMetadata) SetRuleSetName(v string) {
	o.RuleSetName = &v
}

// GetCustomSql returns the CustomSql field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetCustomSql() string {
	if o == nil || IsNil(o.CustomSql) {
		var ret string
		return ret
	}
	return *o.CustomSql
}

// GetCustomSqlOk returns a tuple with the CustomSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetCustomSqlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomSql) {
		return nil, false
	}
	return o.CustomSql, true
}

// HasCustomSql returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasCustomSql() bool {
	if o != nil && !IsNil(o.CustomSql) {
		return true
	}

	return false
}

// SetCustomSql gets a reference to the given string and assigns it to the CustomSql field.
func (o *DatabaseTableMetadata) SetCustomSql(v string) {
	o.CustomSql = &v
}

// GetWhereClause returns the WhereClause field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetWhereClause() string {
	if o == nil || IsNil(o.WhereClause) {
		var ret string
		return ret
	}
	return *o.WhereClause
}

// GetWhereClauseOk returns a tuple with the WhereClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetWhereClauseOk() (*string, bool) {
	if o == nil || IsNil(o.WhereClause) {
		return nil, false
	}
	return o.WhereClause, true
}

// HasWhereClause returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasWhereClause() bool {
	if o != nil && !IsNil(o.WhereClause) {
		return true
	}

	return false
}

// SetWhereClause gets a reference to the given string and assigns it to the WhereClause field.
func (o *DatabaseTableMetadata) SetWhereClause(v string) {
	o.WhereClause = &v
}

// GetHavingClause returns the HavingClause field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetHavingClause() string {
	if o == nil || IsNil(o.HavingClause) {
		var ret string
		return ret
	}
	return *o.HavingClause
}

// GetHavingClauseOk returns a tuple with the HavingClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetHavingClauseOk() (*string, bool) {
	if o == nil || IsNil(o.HavingClause) {
		return nil, false
	}
	return o.HavingClause, true
}

// HasHavingClause returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasHavingClause() bool {
	if o != nil && !IsNil(o.HavingClause) {
		return true
	}

	return false
}

// SetHavingClause gets a reference to the given string and assigns it to the HavingClause field.
func (o *DatabaseTableMetadata) SetHavingClause(v string) {
	o.HavingClause = &v
}

// GetKeyColumn returns the KeyColumn field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetKeyColumn() string {
	if o == nil || IsNil(o.KeyColumn) {
		var ret string
		return ret
	}
	return *o.KeyColumn
}

// GetKeyColumnOk returns a tuple with the KeyColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetKeyColumnOk() (*string, bool) {
	if o == nil || IsNil(o.KeyColumn) {
		return nil, false
	}
	return o.KeyColumn, true
}

// HasKeyColumn returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasKeyColumn() bool {
	if o != nil && !IsNil(o.KeyColumn) {
		return true
	}

	return false
}

// SetKeyColumn gets a reference to the given string and assigns it to the KeyColumn field.
func (o *DatabaseTableMetadata) SetKeyColumn(v string) {
	o.KeyColumn = &v
}

// GetIsSensitive returns the IsSensitive field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetIsSensitive() bool {
	if o == nil || IsNil(o.IsSensitive) {
		var ret bool
		return ret
	}
	return *o.IsSensitive
}

// GetIsSensitiveOk returns a tuple with the IsSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetIsSensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSensitive) {
		return nil, false
	}
	return o.IsSensitive, true
}

// HasIsSensitive returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasIsSensitive() bool {
	if o != nil && !IsNil(o.IsSensitive) {
		return true
	}

	return false
}

// SetIsSensitive gets a reference to the given bool and assigns it to the IsSensitive field.
func (o *DatabaseTableMetadata) SetIsSensitive(v bool) {
	o.IsSensitive = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetRowCount() int64 {
	if o == nil || IsNil(o.RowCount) {
		var ret int64
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetRowCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RowCount) {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasRowCount() bool {
	if o != nil && !IsNil(o.RowCount) {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int64 and assigns it to the RowCount field.
func (o *DatabaseTableMetadata) SetRowCount(v int64) {
	o.RowCount = &v
}

// GetLastRefreshTime returns the LastRefreshTime field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetLastRefreshTime() time.Time {
	if o == nil || IsNil(o.LastRefreshTime) {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshTime
}

// GetLastRefreshTimeOk returns a tuple with the LastRefreshTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetLastRefreshTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRefreshTime) {
		return nil, false
	}
	return o.LastRefreshTime, true
}

// HasLastRefreshTime returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasLastRefreshTime() bool {
	if o != nil && !IsNil(o.LastRefreshTime) {
		return true
	}

	return false
}

// SetLastRefreshTime gets a reference to the given time.Time and assigns it to the LastRefreshTime field.
func (o *DatabaseTableMetadata) SetLastRefreshTime(v time.Time) {
	o.LastRefreshTime = &v
}

// GetLastRowCountTime returns the LastRowCountTime field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetLastRowCountTime() time.Time {
	if o == nil || IsNil(o.LastRowCountTime) {
		var ret time.Time
		return ret
	}
	return *o.LastRowCountTime
}

// GetLastRowCountTimeOk returns a tuple with the LastRowCountTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetLastRowCountTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRowCountTime) {
		return nil, false
	}
	return o.LastRowCountTime, true
}

// HasLastRowCountTime returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasLastRowCountTime() bool {
	if o != nil && !IsNil(o.LastRowCountTime) {
		return true
	}

	return false
}

// SetLastRowCountTime gets a reference to the given time.Time and assigns it to the LastRowCountTime field.
func (o *DatabaseTableMetadata) SetLastRowCountTime(v time.Time) {
	o.LastRowCountTime = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *DatabaseTableMetadata) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *DatabaseTableMetadata) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadata) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *DatabaseTableMetadata) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *DatabaseTableMetadata) SetEngineName(v string) {
	o.EngineName = &v
}

func (o DatabaseTableMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseTableMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TableName) {
		toSerialize["table_name"] = o.TableName
	}
	if !IsNil(o.RuleSetId) {
		toSerialize["rule_set_id"] = o.RuleSetId
	}
	if !IsNil(o.RuleSetName) {
		toSerialize["rule_set_name"] = o.RuleSetName
	}
	if !IsNil(o.CustomSql) {
		toSerialize["custom_sql"] = o.CustomSql
	}
	if !IsNil(o.WhereClause) {
		toSerialize["where_clause"] = o.WhereClause
	}
	if !IsNil(o.HavingClause) {
		toSerialize["having_clause"] = o.HavingClause
	}
	if !IsNil(o.KeyColumn) {
		toSerialize["key_column"] = o.KeyColumn
	}
	if !IsNil(o.IsSensitive) {
		toSerialize["is_sensitive"] = o.IsSensitive
	}
	if !IsNil(o.RowCount) {
		toSerialize["row_count"] = o.RowCount
	}
	if !IsNil(o.LastRefreshTime) {
		toSerialize["last_refresh_time"] = o.LastRefreshTime
	}
	if !IsNil(o.LastRowCountTime) {
		toSerialize["last_row_count_time"] = o.LastRowCountTime
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	return toSerialize, nil
}

type NullableDatabaseTableMetadata struct {
	value *DatabaseTableMetadata
	isSet bool
}

func (v NullableDatabaseTableMetadata) Get() *DatabaseTableMetadata {
	return v.value
}

func (v *NullableDatabaseTableMetadata) Set(val *DatabaseTableMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseTableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseTableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseTableMetadata(val *DatabaseTableMetadata) *NullableDatabaseTableMetadata {
	return &NullableDatabaseTableMetadata{value: val, isSet: true}
}

func (v NullableDatabaseTableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseTableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


