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

// checks if the Toolkit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Toolkit{}

// Toolkit A toolkit or plugin.
type Toolkit struct {
	// Id of the toolkit.
	Id *string `json:"id,omitempty"`
	// Specifies whether this object is toolkit or plugin
	Type *string `json:"type,omitempty"`
	// The object reference.
	Reference *string `json:"reference,omitempty"`
	// Name of the engine.
	EngineName *string `json:"engine_name,omitempty"`
	// Id of the engine.
	EngineId *string `json:"engine_id,omitempty"`
	// Definition of how to provision virtual sources of this type
	VirtualSourceDefinition map[string]interface{} `json:"virtual_source_definition,omitempty"`
	// Definition of how to link sources of this type.
	LinkedSourceDefinition map[string]interface{} `json:"linked_source_definition,omitempty"`
	// Definition of how to discover sources of this type.
	DiscoveryDefinition map[string]interface{} `json:"discovery_definition,omitempty"`
	// Definition of how to upgrade sources of this type.
	UpgradeDefinition map[string]interface{} `json:"upgrade_definition,omitempty"`
	// The schema that defines the structure of the fields in AppDataSyncParameters.
	SnapshotParametersDefinition map[string]interface{} `json:"snapshot_parameters_definition,omitempty"`
	// Tags associated to this toolkit.
	Tags []Tag `json:"tags,omitempty"`
}

// NewToolkit instantiates a new Toolkit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolkit() *Toolkit {
	this := Toolkit{}
	return &this
}

// NewToolkitWithDefaults instantiates a new Toolkit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolkitWithDefaults() *Toolkit {
	this := Toolkit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Toolkit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Toolkit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Toolkit) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Toolkit) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Toolkit) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Toolkit) SetType(v string) {
	o.Type = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Toolkit) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Toolkit) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Toolkit) SetReference(v string) {
	o.Reference = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *Toolkit) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *Toolkit) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *Toolkit) SetEngineName(v string) {
	o.EngineName = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *Toolkit) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *Toolkit) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *Toolkit) SetEngineId(v string) {
	o.EngineId = &v
}

// GetVirtualSourceDefinition returns the VirtualSourceDefinition field value if set, zero value otherwise.
func (o *Toolkit) GetVirtualSourceDefinition() map[string]interface{} {
	if o == nil || IsNil(o.VirtualSourceDefinition) {
		var ret map[string]interface{}
		return ret
	}
	return o.VirtualSourceDefinition
}

// GetVirtualSourceDefinitionOk returns a tuple with the VirtualSourceDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetVirtualSourceDefinitionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.VirtualSourceDefinition) {
		return map[string]interface{}{}, false
	}
	return o.VirtualSourceDefinition, true
}

// HasVirtualSourceDefinition returns a boolean if a field has been set.
func (o *Toolkit) HasVirtualSourceDefinition() bool {
	if o != nil && !IsNil(o.VirtualSourceDefinition) {
		return true
	}

	return false
}

// SetVirtualSourceDefinition gets a reference to the given map[string]interface{} and assigns it to the VirtualSourceDefinition field.
func (o *Toolkit) SetVirtualSourceDefinition(v map[string]interface{}) {
	o.VirtualSourceDefinition = v
}

// GetLinkedSourceDefinition returns the LinkedSourceDefinition field value if set, zero value otherwise.
func (o *Toolkit) GetLinkedSourceDefinition() map[string]interface{} {
	if o == nil || IsNil(o.LinkedSourceDefinition) {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedSourceDefinition
}

// GetLinkedSourceDefinitionOk returns a tuple with the LinkedSourceDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetLinkedSourceDefinitionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LinkedSourceDefinition) {
		return map[string]interface{}{}, false
	}
	return o.LinkedSourceDefinition, true
}

// HasLinkedSourceDefinition returns a boolean if a field has been set.
func (o *Toolkit) HasLinkedSourceDefinition() bool {
	if o != nil && !IsNil(o.LinkedSourceDefinition) {
		return true
	}

	return false
}

// SetLinkedSourceDefinition gets a reference to the given map[string]interface{} and assigns it to the LinkedSourceDefinition field.
func (o *Toolkit) SetLinkedSourceDefinition(v map[string]interface{}) {
	o.LinkedSourceDefinition = v
}

// GetDiscoveryDefinition returns the DiscoveryDefinition field value if set, zero value otherwise.
func (o *Toolkit) GetDiscoveryDefinition() map[string]interface{} {
	if o == nil || IsNil(o.DiscoveryDefinition) {
		var ret map[string]interface{}
		return ret
	}
	return o.DiscoveryDefinition
}

// GetDiscoveryDefinitionOk returns a tuple with the DiscoveryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetDiscoveryDefinitionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DiscoveryDefinition) {
		return map[string]interface{}{}, false
	}
	return o.DiscoveryDefinition, true
}

// HasDiscoveryDefinition returns a boolean if a field has been set.
func (o *Toolkit) HasDiscoveryDefinition() bool {
	if o != nil && !IsNil(o.DiscoveryDefinition) {
		return true
	}

	return false
}

// SetDiscoveryDefinition gets a reference to the given map[string]interface{} and assigns it to the DiscoveryDefinition field.
func (o *Toolkit) SetDiscoveryDefinition(v map[string]interface{}) {
	o.DiscoveryDefinition = v
}

// GetUpgradeDefinition returns the UpgradeDefinition field value if set, zero value otherwise.
func (o *Toolkit) GetUpgradeDefinition() map[string]interface{} {
	if o == nil || IsNil(o.UpgradeDefinition) {
		var ret map[string]interface{}
		return ret
	}
	return o.UpgradeDefinition
}

// GetUpgradeDefinitionOk returns a tuple with the UpgradeDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetUpgradeDefinitionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UpgradeDefinition) {
		return map[string]interface{}{}, false
	}
	return o.UpgradeDefinition, true
}

// HasUpgradeDefinition returns a boolean if a field has been set.
func (o *Toolkit) HasUpgradeDefinition() bool {
	if o != nil && !IsNil(o.UpgradeDefinition) {
		return true
	}

	return false
}

// SetUpgradeDefinition gets a reference to the given map[string]interface{} and assigns it to the UpgradeDefinition field.
func (o *Toolkit) SetUpgradeDefinition(v map[string]interface{}) {
	o.UpgradeDefinition = v
}

// GetSnapshotParametersDefinition returns the SnapshotParametersDefinition field value if set, zero value otherwise.
func (o *Toolkit) GetSnapshotParametersDefinition() map[string]interface{} {
	if o == nil || IsNil(o.SnapshotParametersDefinition) {
		var ret map[string]interface{}
		return ret
	}
	return o.SnapshotParametersDefinition
}

// GetSnapshotParametersDefinitionOk returns a tuple with the SnapshotParametersDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetSnapshotParametersDefinitionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SnapshotParametersDefinition) {
		return map[string]interface{}{}, false
	}
	return o.SnapshotParametersDefinition, true
}

// HasSnapshotParametersDefinition returns a boolean if a field has been set.
func (o *Toolkit) HasSnapshotParametersDefinition() bool {
	if o != nil && !IsNil(o.SnapshotParametersDefinition) {
		return true
	}

	return false
}

// SetSnapshotParametersDefinition gets a reference to the given map[string]interface{} and assigns it to the SnapshotParametersDefinition field.
func (o *Toolkit) SetSnapshotParametersDefinition(v map[string]interface{}) {
	o.SnapshotParametersDefinition = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Toolkit) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toolkit) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Toolkit) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Toolkit) SetTags(v []Tag) {
	o.Tags = v
}

func (o Toolkit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Toolkit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.VirtualSourceDefinition) {
		toSerialize["virtual_source_definition"] = o.VirtualSourceDefinition
	}
	if !IsNil(o.LinkedSourceDefinition) {
		toSerialize["linked_source_definition"] = o.LinkedSourceDefinition
	}
	if !IsNil(o.DiscoveryDefinition) {
		toSerialize["discovery_definition"] = o.DiscoveryDefinition
	}
	if !IsNil(o.UpgradeDefinition) {
		toSerialize["upgrade_definition"] = o.UpgradeDefinition
	}
	if !IsNil(o.SnapshotParametersDefinition) {
		toSerialize["snapshot_parameters_definition"] = o.SnapshotParametersDefinition
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableToolkit struct {
	value *Toolkit
	isSet bool
}

func (v NullableToolkit) Get() *Toolkit {
	return v.value
}

func (v *NullableToolkit) Set(val *Toolkit) {
	v.value = val
	v.isSet = true
}

func (v NullableToolkit) IsSet() bool {
	return v.isSet
}

func (v *NullableToolkit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolkit(val *Toolkit) *NullableToolkit {
	return &NullableToolkit{value: val, isSet: true}
}

func (v NullableToolkit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolkit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


