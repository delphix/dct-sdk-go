/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the HyperscaleDatasetTableOrFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleDatasetTableOrFile{}

// HyperscaleDatasetTableOrFile A table or file which is part of a hyperscale dataset.
type HyperscaleDatasetTableOrFile struct {
	// The ID of the Hyperscale Dataset table or file.
	Id *string `json:"id,omitempty"`
	// Name of the table schema (Oracle/MSSql only).
	SchemaName *string `json:"schema_name,omitempty"`
	// Name of the table (Oracle/MSSql only).
	TableName *string `json:"table_name,omitempty"`
	// Name of the collection (MongoDB only).
	CollectionName *string `json:"collection_name,omitempty"`
	// Name of the database (MongoDB only).
	DatabaseName *string `json:"database_name,omitempty"`
	// The unique database column field to filter the source data.
	FilterKey *string `json:"filter_key,omitempty"`
	// The number of column array rows to be used by the sqlldr oracle utility which determines the number of rows loaded before the stream buffer is built.
	ColumnArrayRows *int64 `json:"column_array_rows,omitempty"`
	// The number of unloaded files to be generated from the source database.
	UnloadSplit *int64 `json:"unload_split,omitempty"`
	// Long The stream size to be used by the sqlldr oracle utility which specifies the size (in bytes) of the data stream sent from the client to the server.
	StreamSize *int64 `json:"stream_size,omitempty"`
	// The end of line character. Support values are \\n, \\r  and \\r\\n (Delimited files only).
	EndOfRecord *string `json:"end_of_record,omitempty"`
	// The single character length delimiter used in source files (Delimited files only).
	Delimiter *string `json:"delimiter,omitempty"`
	// The single character length quote character used in the source files (Delimited files only).
	Enclosure *string `json:"enclosure,omitempty"`
	// The escape character used to escape quote characters (Delimited files only).
	EnclosureEscapeCharacter *string `json:"enclosure_escape_character,omitempty"`
	// Whether to escape the enclosure escape character (Delimited files only).
	EscapeEnclosureEscapeCharacter *bool `json:"escape_enclosure_escape_character,omitempty"`
	// Whether source files have header column names or not (Delimited files only). If set to true, format files with the same column names are created and the same can be used for the masking inventory. If set to false, the column names of pattern f0, f1, f2, and so on are used to create the format files for delimited file masking.
	HasHeaders *bool `json:"has_headers,omitempty"`
	// This is the source key that maps the load-service and masking-service data sets with the unload-service data set (Delimited files only). Please ensure that this value is different for each HyperscaleDatasetTableOrFile.
	UniqueSourceFilesIdentifier *string `json:"unique_source_files_identifier,omitempty"`
	// List of all source files that need to be masked (Delimited files only). All files should have the same delimiter character and other helper characters. All files should have the same number of columns and same column names if it has a header line.
	SourceFiles []string `json:"source_files,omitempty"`
	// Whether the split files must be joined (Delimited files only).
	PerformJoin *bool `json:"perform_join,omitempty"`
	// DataSet information for masking inventory.
	MaskingInventory []HyperscaleColumnOrField `json:"masking_inventory,omitempty"`
}

// NewHyperscaleDatasetTableOrFile instantiates a new HyperscaleDatasetTableOrFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleDatasetTableOrFile() *HyperscaleDatasetTableOrFile {
	this := HyperscaleDatasetTableOrFile{}
	return &this
}

// NewHyperscaleDatasetTableOrFileWithDefaults instantiates a new HyperscaleDatasetTableOrFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleDatasetTableOrFileWithDefaults() *HyperscaleDatasetTableOrFile {
	this := HyperscaleDatasetTableOrFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HyperscaleDatasetTableOrFile) SetId(v string) {
	o.Id = &v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetSchemaName() string {
	if o == nil || IsNil(o.SchemaName) {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetSchemaNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaName) {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasSchemaName() bool {
	if o != nil && !IsNil(o.SchemaName) {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *HyperscaleDatasetTableOrFile) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *HyperscaleDatasetTableOrFile) SetTableName(v string) {
	o.TableName = &v
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetCollectionName() string {
	if o == nil || IsNil(o.CollectionName) {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetCollectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionName) {
		return nil, false
	}
	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasCollectionName() bool {
	if o != nil && !IsNil(o.CollectionName) {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *HyperscaleDatasetTableOrFile) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *HyperscaleDatasetTableOrFile) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetFilterKey returns the FilterKey field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetFilterKey() string {
	if o == nil || IsNil(o.FilterKey) {
		var ret string
		return ret
	}
	return *o.FilterKey
}

// GetFilterKeyOk returns a tuple with the FilterKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetFilterKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FilterKey) {
		return nil, false
	}
	return o.FilterKey, true
}

// HasFilterKey returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasFilterKey() bool {
	if o != nil && !IsNil(o.FilterKey) {
		return true
	}

	return false
}

// SetFilterKey gets a reference to the given string and assigns it to the FilterKey field.
func (o *HyperscaleDatasetTableOrFile) SetFilterKey(v string) {
	o.FilterKey = &v
}

// GetColumnArrayRows returns the ColumnArrayRows field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetColumnArrayRows() int64 {
	if o == nil || IsNil(o.ColumnArrayRows) {
		var ret int64
		return ret
	}
	return *o.ColumnArrayRows
}

// GetColumnArrayRowsOk returns a tuple with the ColumnArrayRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetColumnArrayRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.ColumnArrayRows) {
		return nil, false
	}
	return o.ColumnArrayRows, true
}

// HasColumnArrayRows returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasColumnArrayRows() bool {
	if o != nil && !IsNil(o.ColumnArrayRows) {
		return true
	}

	return false
}

// SetColumnArrayRows gets a reference to the given int64 and assigns it to the ColumnArrayRows field.
func (o *HyperscaleDatasetTableOrFile) SetColumnArrayRows(v int64) {
	o.ColumnArrayRows = &v
}

// GetUnloadSplit returns the UnloadSplit field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetUnloadSplit() int64 {
	if o == nil || IsNil(o.UnloadSplit) {
		var ret int64
		return ret
	}
	return *o.UnloadSplit
}

// GetUnloadSplitOk returns a tuple with the UnloadSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetUnloadSplitOk() (*int64, bool) {
	if o == nil || IsNil(o.UnloadSplit) {
		return nil, false
	}
	return o.UnloadSplit, true
}

// HasUnloadSplit returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasUnloadSplit() bool {
	if o != nil && !IsNil(o.UnloadSplit) {
		return true
	}

	return false
}

// SetUnloadSplit gets a reference to the given int64 and assigns it to the UnloadSplit field.
func (o *HyperscaleDatasetTableOrFile) SetUnloadSplit(v int64) {
	o.UnloadSplit = &v
}

// GetStreamSize returns the StreamSize field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetStreamSize() int64 {
	if o == nil || IsNil(o.StreamSize) {
		var ret int64
		return ret
	}
	return *o.StreamSize
}

// GetStreamSizeOk returns a tuple with the StreamSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetStreamSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.StreamSize) {
		return nil, false
	}
	return o.StreamSize, true
}

// HasStreamSize returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasStreamSize() bool {
	if o != nil && !IsNil(o.StreamSize) {
		return true
	}

	return false
}

// SetStreamSize gets a reference to the given int64 and assigns it to the StreamSize field.
func (o *HyperscaleDatasetTableOrFile) SetStreamSize(v int64) {
	o.StreamSize = &v
}

// GetEndOfRecord returns the EndOfRecord field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetEndOfRecord() string {
	if o == nil || IsNil(o.EndOfRecord) {
		var ret string
		return ret
	}
	return *o.EndOfRecord
}

// GetEndOfRecordOk returns a tuple with the EndOfRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetEndOfRecordOk() (*string, bool) {
	if o == nil || IsNil(o.EndOfRecord) {
		return nil, false
	}
	return o.EndOfRecord, true
}

// HasEndOfRecord returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasEndOfRecord() bool {
	if o != nil && !IsNil(o.EndOfRecord) {
		return true
	}

	return false
}

// SetEndOfRecord gets a reference to the given string and assigns it to the EndOfRecord field.
func (o *HyperscaleDatasetTableOrFile) SetEndOfRecord(v string) {
	o.EndOfRecord = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetDelimiter() string {
	if o == nil || IsNil(o.Delimiter) {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetDelimiterOk() (*string, bool) {
	if o == nil || IsNil(o.Delimiter) {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasDelimiter() bool {
	if o != nil && !IsNil(o.Delimiter) {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *HyperscaleDatasetTableOrFile) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetEnclosure returns the Enclosure field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetEnclosure() string {
	if o == nil || IsNil(o.Enclosure) {
		var ret string
		return ret
	}
	return *o.Enclosure
}

// GetEnclosureOk returns a tuple with the Enclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetEnclosureOk() (*string, bool) {
	if o == nil || IsNil(o.Enclosure) {
		return nil, false
	}
	return o.Enclosure, true
}

// HasEnclosure returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasEnclosure() bool {
	if o != nil && !IsNil(o.Enclosure) {
		return true
	}

	return false
}

// SetEnclosure gets a reference to the given string and assigns it to the Enclosure field.
func (o *HyperscaleDatasetTableOrFile) SetEnclosure(v string) {
	o.Enclosure = &v
}

// GetEnclosureEscapeCharacter returns the EnclosureEscapeCharacter field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetEnclosureEscapeCharacter() string {
	if o == nil || IsNil(o.EnclosureEscapeCharacter) {
		var ret string
		return ret
	}
	return *o.EnclosureEscapeCharacter
}

// GetEnclosureEscapeCharacterOk returns a tuple with the EnclosureEscapeCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetEnclosureEscapeCharacterOk() (*string, bool) {
	if o == nil || IsNil(o.EnclosureEscapeCharacter) {
		return nil, false
	}
	return o.EnclosureEscapeCharacter, true
}

// HasEnclosureEscapeCharacter returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasEnclosureEscapeCharacter() bool {
	if o != nil && !IsNil(o.EnclosureEscapeCharacter) {
		return true
	}

	return false
}

// SetEnclosureEscapeCharacter gets a reference to the given string and assigns it to the EnclosureEscapeCharacter field.
func (o *HyperscaleDatasetTableOrFile) SetEnclosureEscapeCharacter(v string) {
	o.EnclosureEscapeCharacter = &v
}

// GetEscapeEnclosureEscapeCharacter returns the EscapeEnclosureEscapeCharacter field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetEscapeEnclosureEscapeCharacter() bool {
	if o == nil || IsNil(o.EscapeEnclosureEscapeCharacter) {
		var ret bool
		return ret
	}
	return *o.EscapeEnclosureEscapeCharacter
}

// GetEscapeEnclosureEscapeCharacterOk returns a tuple with the EscapeEnclosureEscapeCharacter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetEscapeEnclosureEscapeCharacterOk() (*bool, bool) {
	if o == nil || IsNil(o.EscapeEnclosureEscapeCharacter) {
		return nil, false
	}
	return o.EscapeEnclosureEscapeCharacter, true
}

// HasEscapeEnclosureEscapeCharacter returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasEscapeEnclosureEscapeCharacter() bool {
	if o != nil && !IsNil(o.EscapeEnclosureEscapeCharacter) {
		return true
	}

	return false
}

// SetEscapeEnclosureEscapeCharacter gets a reference to the given bool and assigns it to the EscapeEnclosureEscapeCharacter field.
func (o *HyperscaleDatasetTableOrFile) SetEscapeEnclosureEscapeCharacter(v bool) {
	o.EscapeEnclosureEscapeCharacter = &v
}

// GetHasHeaders returns the HasHeaders field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetHasHeaders() bool {
	if o == nil || IsNil(o.HasHeaders) {
		var ret bool
		return ret
	}
	return *o.HasHeaders
}

// GetHasHeadersOk returns a tuple with the HasHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetHasHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.HasHeaders) {
		return nil, false
	}
	return o.HasHeaders, true
}

// HasHasHeaders returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasHasHeaders() bool {
	if o != nil && !IsNil(o.HasHeaders) {
		return true
	}

	return false
}

// SetHasHeaders gets a reference to the given bool and assigns it to the HasHeaders field.
func (o *HyperscaleDatasetTableOrFile) SetHasHeaders(v bool) {
	o.HasHeaders = &v
}

// GetUniqueSourceFilesIdentifier returns the UniqueSourceFilesIdentifier field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetUniqueSourceFilesIdentifier() string {
	if o == nil || IsNil(o.UniqueSourceFilesIdentifier) {
		var ret string
		return ret
	}
	return *o.UniqueSourceFilesIdentifier
}

// GetUniqueSourceFilesIdentifierOk returns a tuple with the UniqueSourceFilesIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetUniqueSourceFilesIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueSourceFilesIdentifier) {
		return nil, false
	}
	return o.UniqueSourceFilesIdentifier, true
}

// HasUniqueSourceFilesIdentifier returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasUniqueSourceFilesIdentifier() bool {
	if o != nil && !IsNil(o.UniqueSourceFilesIdentifier) {
		return true
	}

	return false
}

// SetUniqueSourceFilesIdentifier gets a reference to the given string and assigns it to the UniqueSourceFilesIdentifier field.
func (o *HyperscaleDatasetTableOrFile) SetUniqueSourceFilesIdentifier(v string) {
	o.UniqueSourceFilesIdentifier = &v
}

// GetSourceFiles returns the SourceFiles field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetSourceFiles() []string {
	if o == nil || IsNil(o.SourceFiles) {
		var ret []string
		return ret
	}
	return o.SourceFiles
}

// GetSourceFilesOk returns a tuple with the SourceFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetSourceFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceFiles) {
		return nil, false
	}
	return o.SourceFiles, true
}

// HasSourceFiles returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasSourceFiles() bool {
	if o != nil && !IsNil(o.SourceFiles) {
		return true
	}

	return false
}

// SetSourceFiles gets a reference to the given []string and assigns it to the SourceFiles field.
func (o *HyperscaleDatasetTableOrFile) SetSourceFiles(v []string) {
	o.SourceFiles = v
}

// GetPerformJoin returns the PerformJoin field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetPerformJoin() bool {
	if o == nil || IsNil(o.PerformJoin) {
		var ret bool
		return ret
	}
	return *o.PerformJoin
}

// GetPerformJoinOk returns a tuple with the PerformJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetPerformJoinOk() (*bool, bool) {
	if o == nil || IsNil(o.PerformJoin) {
		return nil, false
	}
	return o.PerformJoin, true
}

// HasPerformJoin returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasPerformJoin() bool {
	if o != nil && !IsNil(o.PerformJoin) {
		return true
	}

	return false
}

// SetPerformJoin gets a reference to the given bool and assigns it to the PerformJoin field.
func (o *HyperscaleDatasetTableOrFile) SetPerformJoin(v bool) {
	o.PerformJoin = &v
}

// GetMaskingInventory returns the MaskingInventory field value if set, zero value otherwise.
func (o *HyperscaleDatasetTableOrFile) GetMaskingInventory() []HyperscaleColumnOrField {
	if o == nil || IsNil(o.MaskingInventory) {
		var ret []HyperscaleColumnOrField
		return ret
	}
	return o.MaskingInventory
}

// GetMaskingInventoryOk returns a tuple with the MaskingInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleDatasetTableOrFile) GetMaskingInventoryOk() ([]HyperscaleColumnOrField, bool) {
	if o == nil || IsNil(o.MaskingInventory) {
		return nil, false
	}
	return o.MaskingInventory, true
}

// HasMaskingInventory returns a boolean if a field has been set.
func (o *HyperscaleDatasetTableOrFile) HasMaskingInventory() bool {
	if o != nil && !IsNil(o.MaskingInventory) {
		return true
	}

	return false
}

// SetMaskingInventory gets a reference to the given []HyperscaleColumnOrField and assigns it to the MaskingInventory field.
func (o *HyperscaleDatasetTableOrFile) SetMaskingInventory(v []HyperscaleColumnOrField) {
	o.MaskingInventory = v
}

func (o HyperscaleDatasetTableOrFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleDatasetTableOrFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SchemaName) {
		toSerialize["schema_name"] = o.SchemaName
	}
	if !IsNil(o.TableName) {
		toSerialize["table_name"] = o.TableName
	}
	if !IsNil(o.CollectionName) {
		toSerialize["collection_name"] = o.CollectionName
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["database_name"] = o.DatabaseName
	}
	if !IsNil(o.FilterKey) {
		toSerialize["filter_key"] = o.FilterKey
	}
	if !IsNil(o.ColumnArrayRows) {
		toSerialize["column_array_rows"] = o.ColumnArrayRows
	}
	if !IsNil(o.UnloadSplit) {
		toSerialize["unload_split"] = o.UnloadSplit
	}
	if !IsNil(o.StreamSize) {
		toSerialize["stream_size"] = o.StreamSize
	}
	if !IsNil(o.EndOfRecord) {
		toSerialize["end_of_record"] = o.EndOfRecord
	}
	if !IsNil(o.Delimiter) {
		toSerialize["delimiter"] = o.Delimiter
	}
	if !IsNil(o.Enclosure) {
		toSerialize["enclosure"] = o.Enclosure
	}
	if !IsNil(o.EnclosureEscapeCharacter) {
		toSerialize["enclosure_escape_character"] = o.EnclosureEscapeCharacter
	}
	if !IsNil(o.EscapeEnclosureEscapeCharacter) {
		toSerialize["escape_enclosure_escape_character"] = o.EscapeEnclosureEscapeCharacter
	}
	if !IsNil(o.HasHeaders) {
		toSerialize["has_headers"] = o.HasHeaders
	}
	if !IsNil(o.UniqueSourceFilesIdentifier) {
		toSerialize["unique_source_files_identifier"] = o.UniqueSourceFilesIdentifier
	}
	if !IsNil(o.SourceFiles) {
		toSerialize["source_files"] = o.SourceFiles
	}
	if !IsNil(o.PerformJoin) {
		toSerialize["perform_join"] = o.PerformJoin
	}
	if !IsNil(o.MaskingInventory) {
		toSerialize["masking_inventory"] = o.MaskingInventory
	}
	return toSerialize, nil
}

type NullableHyperscaleDatasetTableOrFile struct {
	value *HyperscaleDatasetTableOrFile
	isSet bool
}

func (v NullableHyperscaleDatasetTableOrFile) Get() *HyperscaleDatasetTableOrFile {
	return v.value
}

func (v *NullableHyperscaleDatasetTableOrFile) Set(val *HyperscaleDatasetTableOrFile) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleDatasetTableOrFile) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleDatasetTableOrFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleDatasetTableOrFile(val *HyperscaleDatasetTableOrFile) *NullableHyperscaleDatasetTableOrFile {
	return &NullableHyperscaleDatasetTableOrFile{value: val, isSet: true}
}

func (v NullableHyperscaleDatasetTableOrFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleDatasetTableOrFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


