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

// checks if the DatabaseTemplateCreateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseTemplateCreateParameters{}

// DatabaseTemplateCreateParameters Parameters to create a database template.
type DatabaseTemplateCreateParameters struct {
	// The DatabaseTemplate name.
	Name string `json:"name"`
	// User provided description for this template.
	Description *string `json:"description,omitempty"`
	// The type of the source associated with the template.
	SourceType string `json:"source_type"`
	// A name/value map of string configuration parameters.
	Parameters *map[string]string `json:"parameters,omitempty"`
	// Whether the account creating this database template must be configured as owner of the database template.
	MakeCurrentAccountOwner *bool `json:"make_current_account_owner,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

// NewDatabaseTemplateCreateParameters instantiates a new DatabaseTemplateCreateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseTemplateCreateParameters(name string, sourceType string) *DatabaseTemplateCreateParameters {
	this := DatabaseTemplateCreateParameters{}
	this.Name = name
	this.SourceType = sourceType
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// NewDatabaseTemplateCreateParametersWithDefaults instantiates a new DatabaseTemplateCreateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseTemplateCreateParametersWithDefaults() *DatabaseTemplateCreateParameters {
	this := DatabaseTemplateCreateParameters{}
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// GetName returns the Name field value
func (o *DatabaseTemplateCreateParameters) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseTemplateCreateParameters) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseTemplateCreateParameters) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatabaseTemplateCreateParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplateCreateParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatabaseTemplateCreateParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatabaseTemplateCreateParameters) SetDescription(v string) {
	o.Description = &v
}

// GetSourceType returns the SourceType field value
func (o *DatabaseTemplateCreateParameters) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *DatabaseTemplateCreateParameters) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *DatabaseTemplateCreateParameters) SetSourceType(v string) {
	o.SourceType = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *DatabaseTemplateCreateParameters) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplateCreateParameters) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DatabaseTemplateCreateParameters) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *DatabaseTemplateCreateParameters) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field value if set, zero value otherwise.
func (o *DatabaseTemplateCreateParameters) GetMakeCurrentAccountOwner() bool {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		var ret bool
		return ret
	}
	return *o.MakeCurrentAccountOwner
}

// GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplateCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		return nil, false
	}
	return o.MakeCurrentAccountOwner, true
}

// HasMakeCurrentAccountOwner returns a boolean if a field has been set.
func (o *DatabaseTemplateCreateParameters) HasMakeCurrentAccountOwner() bool {
	if o != nil && !IsNil(o.MakeCurrentAccountOwner) {
		return true
	}

	return false
}

// SetMakeCurrentAccountOwner gets a reference to the given bool and assigns it to the MakeCurrentAccountOwner field.
func (o *DatabaseTemplateCreateParameters) SetMakeCurrentAccountOwner(v bool) {
	o.MakeCurrentAccountOwner = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DatabaseTemplateCreateParameters) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTemplateCreateParameters) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DatabaseTemplateCreateParameters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *DatabaseTemplateCreateParameters) SetTags(v []Tag) {
	o.Tags = v
}

func (o DatabaseTemplateCreateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseTemplateCreateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["source_type"] = o.SourceType
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.MakeCurrentAccountOwner) {
		toSerialize["make_current_account_owner"] = o.MakeCurrentAccountOwner
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableDatabaseTemplateCreateParameters struct {
	value *DatabaseTemplateCreateParameters
	isSet bool
}

func (v NullableDatabaseTemplateCreateParameters) Get() *DatabaseTemplateCreateParameters {
	return v.value
}

func (v *NullableDatabaseTemplateCreateParameters) Set(val *DatabaseTemplateCreateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseTemplateCreateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseTemplateCreateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseTemplateCreateParameters(val *DatabaseTemplateCreateParameters) *NullableDatabaseTemplateCreateParameters {
	return &NullableDatabaseTemplateCreateParameters{value: val, isSet: true}
}

func (v NullableDatabaseTemplateCreateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseTemplateCreateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


