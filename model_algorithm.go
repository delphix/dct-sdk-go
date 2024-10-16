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
	"time"
)

// checks if the Algorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Algorithm{}

// Algorithm A masking algorithm.
type Algorithm struct {
	// The algorithm entity ID.
	Id *string `json:"id,omitempty"`
	// The name of this algorithm.
	Name *string `json:"name,omitempty"`
	// A description of this algorithm.
	Description NullableString `json:"description,omitempty"`
	// The name of the plugin that this algorithm belongs to.
	PluginName *string `json:"plugin_name,omitempty"`
	// The id of the plugin that this algorithm belongs to.
	PluginId *string `json:"plugin_id,omitempty"`
	// The name of the framework that this algorithm belongs to.
	FrameworkName *string `json:"framework_name,omitempty"`
	// The id of the framework that this algorithm belongs to.
	FrameworkId *string `json:"framework_id,omitempty"`
	// The name of the origin engine that this algorithm belongs to.
	EngineName *string `json:"engine_name,omitempty"`
	// The id of the origin engine that this algorithm belongs to.
	EngineId *string `json:"engine_id,omitempty"`
	// The export revision hash of this algorithm from the source engine.
	RevisionHash *string `json:"revision_hash,omitempty"`
	// The configuration of this algorithm.
	Config map[string]interface{} `json:"config,omitempty"`
	// The date and time this algorithm was created.
	CreateDate *time.Time `json:"create_date,omitempty"`
	// The tags of this algorithm.
	Tags []Tag `json:"tags,omitempty"`
}

// NewAlgorithm instantiates a new Algorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgorithm() *Algorithm {
	this := Algorithm{}
	return &this
}

// NewAlgorithmWithDefaults instantiates a new Algorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgorithmWithDefaults() *Algorithm {
	this := Algorithm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Algorithm) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Algorithm) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Algorithm) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Algorithm) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Algorithm) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Algorithm) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Algorithm) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Algorithm) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Algorithm) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Algorithm) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Algorithm) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Algorithm) UnsetDescription() {
	o.Description.Unset()
}

// GetPluginName returns the PluginName field value if set, zero value otherwise.
func (o *Algorithm) GetPluginName() string {
	if o == nil || IsNil(o.PluginName) {
		var ret string
		return ret
	}
	return *o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetPluginNameOk() (*string, bool) {
	if o == nil || IsNil(o.PluginName) {
		return nil, false
	}
	return o.PluginName, true
}

// HasPluginName returns a boolean if a field has been set.
func (o *Algorithm) HasPluginName() bool {
	if o != nil && !IsNil(o.PluginName) {
		return true
	}

	return false
}

// SetPluginName gets a reference to the given string and assigns it to the PluginName field.
func (o *Algorithm) SetPluginName(v string) {
	o.PluginName = &v
}

// GetPluginId returns the PluginId field value if set, zero value otherwise.
func (o *Algorithm) GetPluginId() string {
	if o == nil || IsNil(o.PluginId) {
		var ret string
		return ret
	}
	return *o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetPluginIdOk() (*string, bool) {
	if o == nil || IsNil(o.PluginId) {
		return nil, false
	}
	return o.PluginId, true
}

// HasPluginId returns a boolean if a field has been set.
func (o *Algorithm) HasPluginId() bool {
	if o != nil && !IsNil(o.PluginId) {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given string and assigns it to the PluginId field.
func (o *Algorithm) SetPluginId(v string) {
	o.PluginId = &v
}

// GetFrameworkName returns the FrameworkName field value if set, zero value otherwise.
func (o *Algorithm) GetFrameworkName() string {
	if o == nil || IsNil(o.FrameworkName) {
		var ret string
		return ret
	}
	return *o.FrameworkName
}

// GetFrameworkNameOk returns a tuple with the FrameworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetFrameworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.FrameworkName) {
		return nil, false
	}
	return o.FrameworkName, true
}

// HasFrameworkName returns a boolean if a field has been set.
func (o *Algorithm) HasFrameworkName() bool {
	if o != nil && !IsNil(o.FrameworkName) {
		return true
	}

	return false
}

// SetFrameworkName gets a reference to the given string and assigns it to the FrameworkName field.
func (o *Algorithm) SetFrameworkName(v string) {
	o.FrameworkName = &v
}

// GetFrameworkId returns the FrameworkId field value if set, zero value otherwise.
func (o *Algorithm) GetFrameworkId() string {
	if o == nil || IsNil(o.FrameworkId) {
		var ret string
		return ret
	}
	return *o.FrameworkId
}

// GetFrameworkIdOk returns a tuple with the FrameworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetFrameworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.FrameworkId) {
		return nil, false
	}
	return o.FrameworkId, true
}

// HasFrameworkId returns a boolean if a field has been set.
func (o *Algorithm) HasFrameworkId() bool {
	if o != nil && !IsNil(o.FrameworkId) {
		return true
	}

	return false
}

// SetFrameworkId gets a reference to the given string and assigns it to the FrameworkId field.
func (o *Algorithm) SetFrameworkId(v string) {
	o.FrameworkId = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *Algorithm) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *Algorithm) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *Algorithm) SetEngineName(v string) {
	o.EngineName = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *Algorithm) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *Algorithm) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *Algorithm) SetEngineId(v string) {
	o.EngineId = &v
}

// GetRevisionHash returns the RevisionHash field value if set, zero value otherwise.
func (o *Algorithm) GetRevisionHash() string {
	if o == nil || IsNil(o.RevisionHash) {
		var ret string
		return ret
	}
	return *o.RevisionHash
}

// GetRevisionHashOk returns a tuple with the RevisionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetRevisionHashOk() (*string, bool) {
	if o == nil || IsNil(o.RevisionHash) {
		return nil, false
	}
	return o.RevisionHash, true
}

// HasRevisionHash returns a boolean if a field has been set.
func (o *Algorithm) HasRevisionHash() bool {
	if o != nil && !IsNil(o.RevisionHash) {
		return true
	}

	return false
}

// SetRevisionHash gets a reference to the given string and assigns it to the RevisionHash field.
func (o *Algorithm) SetRevisionHash(v string) {
	o.RevisionHash = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Algorithm) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Algorithm) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Algorithm) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *Algorithm) GetCreateDate() time.Time {
	if o == nil || IsNil(o.CreateDate) {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *Algorithm) HasCreateDate() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *Algorithm) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Algorithm) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Algorithm) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Algorithm) SetTags(v []Tag) {
	o.Tags = v
}

func (o Algorithm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Algorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.PluginName) {
		toSerialize["plugin_name"] = o.PluginName
	}
	if !IsNil(o.PluginId) {
		toSerialize["plugin_id"] = o.PluginId
	}
	if !IsNil(o.FrameworkName) {
		toSerialize["framework_name"] = o.FrameworkName
	}
	if !IsNil(o.FrameworkId) {
		toSerialize["framework_id"] = o.FrameworkId
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.RevisionHash) {
		toSerialize["revision_hash"] = o.RevisionHash
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.CreateDate) {
		toSerialize["create_date"] = o.CreateDate
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAlgorithm struct {
	value *Algorithm
	isSet bool
}

func (v NullableAlgorithm) Get() *Algorithm {
	return v.value
}

func (v *NullableAlgorithm) Set(val *Algorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgorithm(val *Algorithm) *NullableAlgorithm {
	return &NullableAlgorithm{value: val, isSet: true}
}

func (v NullableAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


