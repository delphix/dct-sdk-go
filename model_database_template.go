/*
Delphix DCT API

Delphix DCT API

API version: 3.17.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DatabaseTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseTemplate{}

// DatabaseTemplate A database template to use for provisioning and refresh. If set, configParams will be ignored on provision or refresh.
type DatabaseTemplate struct {
	// The DatabaseTemplate entity ID.
	Id *string `json:"id,omitempty"`
	// The DatabaseTemplate name.
	Name string `json:"name"`
	// User provided description for this template.
	Description *string `json:"description,omitempty"`
	// The type of the source associated with the template.
	SourceType string `json:"source_type"`
	// A name/value map of string configuration parameters.
	Parameters *map[string]string `json:"parameters,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

type _DatabaseTemplate DatabaseTemplate

// NewDatabaseTemplate instantiates a new DatabaseTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseTemplate(name string, sourceType string) *DatabaseTemplate {
	this := DatabaseTemplate{}
	this.Name = name
	this.SourceType = sourceType
	return &this
}

// NewDatabaseTemplateWithDefaults instantiates a new DatabaseTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseTemplateWithDefaults() *DatabaseTemplate {
	this := DatabaseTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DatabaseTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DatabaseTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DatabaseTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DatabaseTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatabaseTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatabaseTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatabaseTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetSourceType returns the SourceType field value
func (o *DatabaseTemplate) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *DatabaseTemplate) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *DatabaseTemplate) SetSourceType(v string) {
	o.SourceType = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *DatabaseTemplate) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplate) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DatabaseTemplate) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *DatabaseTemplate) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DatabaseTemplate) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplate) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DatabaseTemplate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *DatabaseTemplate) SetTags(v []Tag) {
	o.Tags = v
}

func (o DatabaseTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["source_type"] = o.SourceType
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *DatabaseTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"source_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDatabaseTemplate := _DatabaseTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDatabaseTemplate)

	if err != nil {
		return err
	}

	*o = DatabaseTemplate(varDatabaseTemplate)

	return err
}

type NullableDatabaseTemplate struct {
	value *DatabaseTemplate
	isSet bool
}

func (v NullableDatabaseTemplate) Get() *DatabaseTemplate {
	return v.value
}

func (v *NullableDatabaseTemplate) Set(val *DatabaseTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseTemplate(val *DatabaseTemplate) *NullableDatabaseTemplate {
	return &NullableDatabaseTemplate{value: val, isSet: true}
}

func (v NullableDatabaseTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


