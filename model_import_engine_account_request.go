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

// checks if the ImportEngineAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportEngineAccountRequest{}

// ImportEngineAccountRequest struct for ImportEngineAccountRequest
type ImportEngineAccountRequest struct {
	// List of data-layout ids for which accounts should be imported to DCT. This is mutually exclusive with `engine_ids` and `import_all`.
	DataLayoutIds []string `json:"data_layout_ids,omitempty"`
	// List of engine ids for which accounts should be imported to DCT. This is mutually exclusive with `data_layout_ids` and `import_all`.
	EngineIds []string `json:"engine_ids,omitempty"`
	// All self-service accounts across engines should imported to DCT. This is mutually exclusive with `data_layout_ids` and `engine_ids`.
	ImportAll *bool `json:"import_all,omitempty"`
}

// NewImportEngineAccountRequest instantiates a new ImportEngineAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportEngineAccountRequest() *ImportEngineAccountRequest {
	this := ImportEngineAccountRequest{}
	var importAll bool = false
	this.ImportAll = &importAll
	return &this
}

// NewImportEngineAccountRequestWithDefaults instantiates a new ImportEngineAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportEngineAccountRequestWithDefaults() *ImportEngineAccountRequest {
	this := ImportEngineAccountRequest{}
	var importAll bool = false
	this.ImportAll = &importAll
	return &this
}

// GetDataLayoutIds returns the DataLayoutIds field value if set, zero value otherwise.
func (o *ImportEngineAccountRequest) GetDataLayoutIds() []string {
	if o == nil || IsNil(o.DataLayoutIds) {
		var ret []string
		return ret
	}
	return o.DataLayoutIds
}

// GetDataLayoutIdsOk returns a tuple with the DataLayoutIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEngineAccountRequest) GetDataLayoutIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DataLayoutIds) {
		return nil, false
	}
	return o.DataLayoutIds, true
}

// HasDataLayoutIds returns a boolean if a field has been set.
func (o *ImportEngineAccountRequest) HasDataLayoutIds() bool {
	if o != nil && !IsNil(o.DataLayoutIds) {
		return true
	}

	return false
}

// SetDataLayoutIds gets a reference to the given []string and assigns it to the DataLayoutIds field.
func (o *ImportEngineAccountRequest) SetDataLayoutIds(v []string) {
	o.DataLayoutIds = v
}

// GetEngineIds returns the EngineIds field value if set, zero value otherwise.
func (o *ImportEngineAccountRequest) GetEngineIds() []string {
	if o == nil || IsNil(o.EngineIds) {
		var ret []string
		return ret
	}
	return o.EngineIds
}

// GetEngineIdsOk returns a tuple with the EngineIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEngineAccountRequest) GetEngineIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EngineIds) {
		return nil, false
	}
	return o.EngineIds, true
}

// HasEngineIds returns a boolean if a field has been set.
func (o *ImportEngineAccountRequest) HasEngineIds() bool {
	if o != nil && !IsNil(o.EngineIds) {
		return true
	}

	return false
}

// SetEngineIds gets a reference to the given []string and assigns it to the EngineIds field.
func (o *ImportEngineAccountRequest) SetEngineIds(v []string) {
	o.EngineIds = v
}

// GetImportAll returns the ImportAll field value if set, zero value otherwise.
func (o *ImportEngineAccountRequest) GetImportAll() bool {
	if o == nil || IsNil(o.ImportAll) {
		var ret bool
		return ret
	}
	return *o.ImportAll
}

// GetImportAllOk returns a tuple with the ImportAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEngineAccountRequest) GetImportAllOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportAll) {
		return nil, false
	}
	return o.ImportAll, true
}

// HasImportAll returns a boolean if a field has been set.
func (o *ImportEngineAccountRequest) HasImportAll() bool {
	if o != nil && !IsNil(o.ImportAll) {
		return true
	}

	return false
}

// SetImportAll gets a reference to the given bool and assigns it to the ImportAll field.
func (o *ImportEngineAccountRequest) SetImportAll(v bool) {
	o.ImportAll = &v
}

func (o ImportEngineAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportEngineAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataLayoutIds) {
		toSerialize["data_layout_ids"] = o.DataLayoutIds
	}
	if !IsNil(o.EngineIds) {
		toSerialize["engine_ids"] = o.EngineIds
	}
	if !IsNil(o.ImportAll) {
		toSerialize["import_all"] = o.ImportAll
	}
	return toSerialize, nil
}

type NullableImportEngineAccountRequest struct {
	value *ImportEngineAccountRequest
	isSet bool
}

func (v NullableImportEngineAccountRequest) Get() *ImportEngineAccountRequest {
	return v.value
}

func (v *NullableImportEngineAccountRequest) Set(val *ImportEngineAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportEngineAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportEngineAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportEngineAccountRequest(val *ImportEngineAccountRequest) *NullableImportEngineAccountRequest {
	return &NullableImportEngineAccountRequest{value: val, isSet: true}
}

func (v NullableImportEngineAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportEngineAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


