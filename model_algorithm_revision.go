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
	"time"
)

// checks if the AlgorithmRevision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlgorithmRevision{}

// AlgorithmRevision A masking algorithm revision.
type AlgorithmRevision struct {
	// The algorithm revision ID.
	Id string `json:"id"`
	// The name of this algorithm revision.
	Name string `json:"name"`
	// A note of this algorithm.
	Note NullableString `json:"note,omitempty"`
	// The id of the algorithm that this revision belongs to.
	AlgorithmId string `json:"algorithm_id"`
	// The id of the plugin that this algorithm revision belongs to.
	PluginId string `json:"plugin_id"`
	// The id of the framework that this algorithm revision belongs to.
	FrameworkId string `json:"framework_id"`
	// Whether this algorithm revision is the primary revision of its algorithm.
	IsPrimary bool `json:"is_primary"`
	// Whether this algorithm revision is defined in a plugin as a default instance.
	IsDefaultInstance bool `json:"is_default_instance"`
	// The date and time this algorithm revision was created.
	CreateDate *time.Time `json:"create_date,omitempty"`
	// The engines that this algorithm revision is originated from.
	OriginEngines []AlgorithmRevisionOriginEngine `json:"origin_engines,omitempty"`
	// The configuration of this algorithm revision.
	Config map[string]interface{} `json:"config,omitempty"`
	// The tags of this algorithm revision.
	Tags []Tag `json:"tags,omitempty"`
}

// NewAlgorithmRevision instantiates a new AlgorithmRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgorithmRevision(id string, name string, algorithmId string, pluginId string, frameworkId string, isPrimary bool, isDefaultInstance bool) *AlgorithmRevision {
	this := AlgorithmRevision{}
	this.Id = id
	this.Name = name
	this.AlgorithmId = algorithmId
	this.PluginId = pluginId
	this.FrameworkId = frameworkId
	this.IsPrimary = isPrimary
	this.IsDefaultInstance = isDefaultInstance
	return &this
}

// NewAlgorithmRevisionWithDefaults instantiates a new AlgorithmRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgorithmRevisionWithDefaults() *AlgorithmRevision {
	this := AlgorithmRevision{}
	return &this
}

// GetId returns the Id field value
func (o *AlgorithmRevision) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AlgorithmRevision) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AlgorithmRevision) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlgorithmRevision) SetName(v string) {
	o.Name = v
}

// GetNote returns the Note field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlgorithmRevision) GetNote() string {
	if o == nil || IsNil(o.Note.Get()) {
		var ret string
		return ret
	}
	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlgorithmRevision) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// HasNote returns a boolean if a field has been set.
func (o *AlgorithmRevision) HasNote() bool {
	if o != nil && o.Note.IsSet() {
		return true
	}

	return false
}

// SetNote gets a reference to the given NullableString and assigns it to the Note field.
func (o *AlgorithmRevision) SetNote(v string) {
	o.Note.Set(&v)
}
// SetNoteNil sets the value for Note to be an explicit nil
func (o *AlgorithmRevision) SetNoteNil() {
	o.Note.Set(nil)
}

// UnsetNote ensures that no value is present for Note, not even an explicit nil
func (o *AlgorithmRevision) UnsetNote() {
	o.Note.Unset()
}

// GetAlgorithmId returns the AlgorithmId field value
func (o *AlgorithmRevision) GetAlgorithmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgorithmId
}

// GetAlgorithmIdOk returns a tuple with the AlgorithmId field value
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetAlgorithmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlgorithmId, true
}

// SetAlgorithmId sets field value
func (o *AlgorithmRevision) SetAlgorithmId(v string) {
	o.AlgorithmId = v
}

// GetPluginId returns the PluginId field value
func (o *AlgorithmRevision) GetPluginId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetPluginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginId, true
}

// SetPluginId sets field value
func (o *AlgorithmRevision) SetPluginId(v string) {
	o.PluginId = v
}

// GetFrameworkId returns the FrameworkId field value
func (o *AlgorithmRevision) GetFrameworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FrameworkId
}

// GetFrameworkIdOk returns a tuple with the FrameworkId field value
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetFrameworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrameworkId, true
}

// SetFrameworkId sets field value
func (o *AlgorithmRevision) SetFrameworkId(v string) {
	o.FrameworkId = v
}

// GetIsPrimary returns the IsPrimary field value
func (o *AlgorithmRevision) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value
func (o *AlgorithmRevision) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetIsDefaultInstance returns the IsDefaultInstance field value
func (o *AlgorithmRevision) GetIsDefaultInstance() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefaultInstance
}

// GetIsDefaultInstanceOk returns a tuple with the IsDefaultInstance field value
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetIsDefaultInstanceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefaultInstance, true
}

// SetIsDefaultInstance sets field value
func (o *AlgorithmRevision) SetIsDefaultInstance(v bool) {
	o.IsDefaultInstance = v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *AlgorithmRevision) GetCreateDate() time.Time {
	if o == nil || IsNil(o.CreateDate) {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *AlgorithmRevision) HasCreateDate() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *AlgorithmRevision) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetOriginEngines returns the OriginEngines field value if set, zero value otherwise.
func (o *AlgorithmRevision) GetOriginEngines() []AlgorithmRevisionOriginEngine {
	if o == nil || IsNil(o.OriginEngines) {
		var ret []AlgorithmRevisionOriginEngine
		return ret
	}
	return o.OriginEngines
}

// GetOriginEnginesOk returns a tuple with the OriginEngines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetOriginEnginesOk() ([]AlgorithmRevisionOriginEngine, bool) {
	if o == nil || IsNil(o.OriginEngines) {
		return nil, false
	}
	return o.OriginEngines, true
}

// HasOriginEngines returns a boolean if a field has been set.
func (o *AlgorithmRevision) HasOriginEngines() bool {
	if o != nil && !IsNil(o.OriginEngines) {
		return true
	}

	return false
}

// SetOriginEngines gets a reference to the given []AlgorithmRevisionOriginEngine and assigns it to the OriginEngines field.
func (o *AlgorithmRevision) SetOriginEngines(v []AlgorithmRevisionOriginEngine) {
	o.OriginEngines = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AlgorithmRevision) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AlgorithmRevision) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *AlgorithmRevision) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AlgorithmRevision) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgorithmRevision) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AlgorithmRevision) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *AlgorithmRevision) SetTags(v []Tag) {
	o.Tags = v
}

func (o AlgorithmRevision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlgorithmRevision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Note.IsSet() {
		toSerialize["note"] = o.Note.Get()
	}
	toSerialize["algorithm_id"] = o.AlgorithmId
	toSerialize["plugin_id"] = o.PluginId
	toSerialize["framework_id"] = o.FrameworkId
	toSerialize["is_primary"] = o.IsPrimary
	toSerialize["is_default_instance"] = o.IsDefaultInstance
	if !IsNil(o.CreateDate) {
		toSerialize["create_date"] = o.CreateDate
	}
	if !IsNil(o.OriginEngines) {
		toSerialize["origin_engines"] = o.OriginEngines
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAlgorithmRevision struct {
	value *AlgorithmRevision
	isSet bool
}

func (v NullableAlgorithmRevision) Get() *AlgorithmRevision {
	return v.value
}

func (v *NullableAlgorithmRevision) Set(val *AlgorithmRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgorithmRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgorithmRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgorithmRevision(val *AlgorithmRevision) *NullableAlgorithmRevision {
	return &NullableAlgorithmRevision{value: val, isSet: true}
}

func (v NullableAlgorithmRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgorithmRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


