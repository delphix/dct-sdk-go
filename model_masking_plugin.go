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

// checks if the MaskingPlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskingPlugin{}

// MaskingPlugin A masking plugin.
type MaskingPlugin struct {
	// The masking plugin entity ID.
	Id *string `json:"id,omitempty"`
	// The name of this plugin.
	Name *string `json:"name,omitempty"`
	// The plugin type.
	PluginType *string `json:"plugin_type,omitempty"`
	// A description of this plugin.
	Description NullableString `json:"description,omitempty"`
	// The date and time when this plugin is installed.
	InstallDate *time.Time `json:"install_date,omitempty"`
	// Whether this plugin is a built-in plugin.
	BuiltIn *bool `json:"built_in,omitempty"`
	// The masking SDK version that this plugin is built from.
	SdkVersion *string `json:"sdk_version,omitempty"`
	// The version of this plugin.
	Version *string `json:"version,omitempty"`
	// The author of this plugin.
	Author *string `json:"author,omitempty"`
	// The list of frameworks for this plugin.
	Frameworks []Framework `json:"frameworks,omitempty"`
	// The tags of this plugin.
	Tags []Tag `json:"tags,omitempty"`
}

// NewMaskingPlugin instantiates a new MaskingPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskingPlugin() *MaskingPlugin {
	this := MaskingPlugin{}
	return &this
}

// NewMaskingPluginWithDefaults instantiates a new MaskingPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskingPluginWithDefaults() *MaskingPlugin {
	this := MaskingPlugin{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MaskingPlugin) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MaskingPlugin) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MaskingPlugin) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MaskingPlugin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MaskingPlugin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MaskingPlugin) SetName(v string) {
	o.Name = &v
}

// GetPluginType returns the PluginType field value if set, zero value otherwise.
func (o *MaskingPlugin) GetPluginType() string {
	if o == nil || IsNil(o.PluginType) {
		var ret string
		return ret
	}
	return *o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetPluginTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PluginType) {
		return nil, false
	}
	return o.PluginType, true
}

// HasPluginType returns a boolean if a field has been set.
func (o *MaskingPlugin) HasPluginType() bool {
	if o != nil && !IsNil(o.PluginType) {
		return true
	}

	return false
}

// SetPluginType gets a reference to the given string and assigns it to the PluginType field.
func (o *MaskingPlugin) SetPluginType(v string) {
	o.PluginType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingPlugin) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingPlugin) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MaskingPlugin) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MaskingPlugin) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MaskingPlugin) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MaskingPlugin) UnsetDescription() {
	o.Description.Unset()
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise.
func (o *MaskingPlugin) GetInstallDate() time.Time {
	if o == nil || IsNil(o.InstallDate) {
		var ret time.Time
		return ret
	}
	return *o.InstallDate
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetInstallDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InstallDate) {
		return nil, false
	}
	return o.InstallDate, true
}

// HasInstallDate returns a boolean if a field has been set.
func (o *MaskingPlugin) HasInstallDate() bool {
	if o != nil && !IsNil(o.InstallDate) {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given time.Time and assigns it to the InstallDate field.
func (o *MaskingPlugin) SetInstallDate(v time.Time) {
	o.InstallDate = &v
}

// GetBuiltIn returns the BuiltIn field value if set, zero value otherwise.
func (o *MaskingPlugin) GetBuiltIn() bool {
	if o == nil || IsNil(o.BuiltIn) {
		var ret bool
		return ret
	}
	return *o.BuiltIn
}

// GetBuiltInOk returns a tuple with the BuiltIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetBuiltInOk() (*bool, bool) {
	if o == nil || IsNil(o.BuiltIn) {
		return nil, false
	}
	return o.BuiltIn, true
}

// HasBuiltIn returns a boolean if a field has been set.
func (o *MaskingPlugin) HasBuiltIn() bool {
	if o != nil && !IsNil(o.BuiltIn) {
		return true
	}

	return false
}

// SetBuiltIn gets a reference to the given bool and assigns it to the BuiltIn field.
func (o *MaskingPlugin) SetBuiltIn(v bool) {
	o.BuiltIn = &v
}

// GetSdkVersion returns the SdkVersion field value if set, zero value otherwise.
func (o *MaskingPlugin) GetSdkVersion() string {
	if o == nil || IsNil(o.SdkVersion) {
		var ret string
		return ret
	}
	return *o.SdkVersion
}

// GetSdkVersionOk returns a tuple with the SdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetSdkVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SdkVersion) {
		return nil, false
	}
	return o.SdkVersion, true
}

// HasSdkVersion returns a boolean if a field has been set.
func (o *MaskingPlugin) HasSdkVersion() bool {
	if o != nil && !IsNil(o.SdkVersion) {
		return true
	}

	return false
}

// SetSdkVersion gets a reference to the given string and assigns it to the SdkVersion field.
func (o *MaskingPlugin) SetSdkVersion(v string) {
	o.SdkVersion = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MaskingPlugin) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MaskingPlugin) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MaskingPlugin) SetVersion(v string) {
	o.Version = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *MaskingPlugin) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MaskingPlugin) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *MaskingPlugin) SetAuthor(v string) {
	o.Author = &v
}

// GetFrameworks returns the Frameworks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingPlugin) GetFrameworks() []Framework {
	if o == nil {
		var ret []Framework
		return ret
	}
	return o.Frameworks
}

// GetFrameworksOk returns a tuple with the Frameworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingPlugin) GetFrameworksOk() ([]Framework, bool) {
	if o == nil || IsNil(o.Frameworks) {
		return nil, false
	}
	return o.Frameworks, true
}

// HasFrameworks returns a boolean if a field has been set.
func (o *MaskingPlugin) HasFrameworks() bool {
	if o != nil && IsNil(o.Frameworks) {
		return true
	}

	return false
}

// SetFrameworks gets a reference to the given []Framework and assigns it to the Frameworks field.
func (o *MaskingPlugin) SetFrameworks(v []Framework) {
	o.Frameworks = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MaskingPlugin) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingPlugin) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MaskingPlugin) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *MaskingPlugin) SetTags(v []Tag) {
	o.Tags = v
}

func (o MaskingPlugin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskingPlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PluginType) {
		toSerialize["plugin_type"] = o.PluginType
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.InstallDate) {
		toSerialize["install_date"] = o.InstallDate
	}
	if !IsNil(o.BuiltIn) {
		toSerialize["built_in"] = o.BuiltIn
	}
	if !IsNil(o.SdkVersion) {
		toSerialize["sdk_version"] = o.SdkVersion
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if o.Frameworks != nil {
		toSerialize["frameworks"] = o.Frameworks
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableMaskingPlugin struct {
	value *MaskingPlugin
	isSet bool
}

func (v NullableMaskingPlugin) Get() *MaskingPlugin {
	return v.value
}

func (v *NullableMaskingPlugin) Set(val *MaskingPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskingPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskingPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskingPlugin(val *MaskingPlugin) *NullableMaskingPlugin {
	return &NullableMaskingPlugin{value: val, isSet: true}
}

func (v NullableMaskingPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskingPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


