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
)

// checks if the EngineAutoTaggingConfigUpdateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EngineAutoTaggingConfigUpdateParameters{}

// EngineAutoTaggingConfigUpdateParameters Configuration settings for auto tagging.
type EngineAutoTaggingConfigUpdateParameters struct {
	// Include dataset group names as tags for dSources, VDBs, and vCDBs (key 'dlpx-dataset-group').
	EnableVirtualizationDatasetGroups *bool `json:"enable_virtualization_dataset_groups,omitempty"`
	// Include masking environment names as tags for masking Jobs and Connectors (key 'dlpx-environment').
	EnableMaskingEnvironments *bool `json:"enable_masking_environments,omitempty"`
	// Include masking app names as tags for masking Jobs and Connectors (key 'dlpx-application').
	EnableMaskingApplications *bool `json:"enable_masking_applications,omitempty"`
	// Include the engine name as a tag on discovered objects (key 'dlpx-engine').
	EnableEngineName *bool `json:"enable_engine_name,omitempty"`
	// List of new custom tags to be added to discovered objects. These are appended to the AutoTaggingConfig's custom_tags list.
	CustomTagsToAdd []Tag `json:"custom_tags_to_add,omitempty"`
	// List of tags to remove from the AutoTaggingConfig's custom_tags list (applied AFTER custom_tags_to_add).
	CustomTagsToRemove []Tag `json:"custom_tags_to_remove,omitempty"`
}

// NewEngineAutoTaggingConfigUpdateParameters instantiates a new EngineAutoTaggingConfigUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineAutoTaggingConfigUpdateParameters() *EngineAutoTaggingConfigUpdateParameters {
	this := EngineAutoTaggingConfigUpdateParameters{}
	return &this
}

// NewEngineAutoTaggingConfigUpdateParametersWithDefaults instantiates a new EngineAutoTaggingConfigUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineAutoTaggingConfigUpdateParametersWithDefaults() *EngineAutoTaggingConfigUpdateParameters {
	this := EngineAutoTaggingConfigUpdateParameters{}
	return &this
}

// GetEnableVirtualizationDatasetGroups returns the EnableVirtualizationDatasetGroups field value if set, zero value otherwise.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableVirtualizationDatasetGroups() bool {
	if o == nil || IsNil(o.EnableVirtualizationDatasetGroups) {
		var ret bool
		return ret
	}
	return *o.EnableVirtualizationDatasetGroups
}

// GetEnableVirtualizationDatasetGroupsOk returns a tuple with the EnableVirtualizationDatasetGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableVirtualizationDatasetGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableVirtualizationDatasetGroups) {
		return nil, false
	}
	return o.EnableVirtualizationDatasetGroups, true
}

// HasEnableVirtualizationDatasetGroups returns a boolean if a field has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) HasEnableVirtualizationDatasetGroups() bool {
	if o != nil && !IsNil(o.EnableVirtualizationDatasetGroups) {
		return true
	}

	return false
}

// SetEnableVirtualizationDatasetGroups gets a reference to the given bool and assigns it to the EnableVirtualizationDatasetGroups field.
func (o *EngineAutoTaggingConfigUpdateParameters) SetEnableVirtualizationDatasetGroups(v bool) {
	o.EnableVirtualizationDatasetGroups = &v
}

// GetEnableMaskingEnvironments returns the EnableMaskingEnvironments field value if set, zero value otherwise.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableMaskingEnvironments() bool {
	if o == nil || IsNil(o.EnableMaskingEnvironments) {
		var ret bool
		return ret
	}
	return *o.EnableMaskingEnvironments
}

// GetEnableMaskingEnvironmentsOk returns a tuple with the EnableMaskingEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableMaskingEnvironmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMaskingEnvironments) {
		return nil, false
	}
	return o.EnableMaskingEnvironments, true
}

// HasEnableMaskingEnvironments returns a boolean if a field has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) HasEnableMaskingEnvironments() bool {
	if o != nil && !IsNil(o.EnableMaskingEnvironments) {
		return true
	}

	return false
}

// SetEnableMaskingEnvironments gets a reference to the given bool and assigns it to the EnableMaskingEnvironments field.
func (o *EngineAutoTaggingConfigUpdateParameters) SetEnableMaskingEnvironments(v bool) {
	o.EnableMaskingEnvironments = &v
}

// GetEnableMaskingApplications returns the EnableMaskingApplications field value if set, zero value otherwise.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableMaskingApplications() bool {
	if o == nil || IsNil(o.EnableMaskingApplications) {
		var ret bool
		return ret
	}
	return *o.EnableMaskingApplications
}

// GetEnableMaskingApplicationsOk returns a tuple with the EnableMaskingApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableMaskingApplicationsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMaskingApplications) {
		return nil, false
	}
	return o.EnableMaskingApplications, true
}

// HasEnableMaskingApplications returns a boolean if a field has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) HasEnableMaskingApplications() bool {
	if o != nil && !IsNil(o.EnableMaskingApplications) {
		return true
	}

	return false
}

// SetEnableMaskingApplications gets a reference to the given bool and assigns it to the EnableMaskingApplications field.
func (o *EngineAutoTaggingConfigUpdateParameters) SetEnableMaskingApplications(v bool) {
	o.EnableMaskingApplications = &v
}

// GetEnableEngineName returns the EnableEngineName field value if set, zero value otherwise.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableEngineName() bool {
	if o == nil || IsNil(o.EnableEngineName) {
		var ret bool
		return ret
	}
	return *o.EnableEngineName
}

// GetEnableEngineNameOk returns a tuple with the EnableEngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) GetEnableEngineNameOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEngineName) {
		return nil, false
	}
	return o.EnableEngineName, true
}

// HasEnableEngineName returns a boolean if a field has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) HasEnableEngineName() bool {
	if o != nil && !IsNil(o.EnableEngineName) {
		return true
	}

	return false
}

// SetEnableEngineName gets a reference to the given bool and assigns it to the EnableEngineName field.
func (o *EngineAutoTaggingConfigUpdateParameters) SetEnableEngineName(v bool) {
	o.EnableEngineName = &v
}

// GetCustomTagsToAdd returns the CustomTagsToAdd field value if set, zero value otherwise.
func (o *EngineAutoTaggingConfigUpdateParameters) GetCustomTagsToAdd() []Tag {
	if o == nil || IsNil(o.CustomTagsToAdd) {
		var ret []Tag
		return ret
	}
	return o.CustomTagsToAdd
}

// GetCustomTagsToAddOk returns a tuple with the CustomTagsToAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) GetCustomTagsToAddOk() ([]Tag, bool) {
	if o == nil || IsNil(o.CustomTagsToAdd) {
		return nil, false
	}
	return o.CustomTagsToAdd, true
}

// HasCustomTagsToAdd returns a boolean if a field has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) HasCustomTagsToAdd() bool {
	if o != nil && !IsNil(o.CustomTagsToAdd) {
		return true
	}

	return false
}

// SetCustomTagsToAdd gets a reference to the given []Tag and assigns it to the CustomTagsToAdd field.
func (o *EngineAutoTaggingConfigUpdateParameters) SetCustomTagsToAdd(v []Tag) {
	o.CustomTagsToAdd = v
}

// GetCustomTagsToRemove returns the CustomTagsToRemove field value if set, zero value otherwise.
func (o *EngineAutoTaggingConfigUpdateParameters) GetCustomTagsToRemove() []Tag {
	if o == nil || IsNil(o.CustomTagsToRemove) {
		var ret []Tag
		return ret
	}
	return o.CustomTagsToRemove
}

// GetCustomTagsToRemoveOk returns a tuple with the CustomTagsToRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) GetCustomTagsToRemoveOk() ([]Tag, bool) {
	if o == nil || IsNil(o.CustomTagsToRemove) {
		return nil, false
	}
	return o.CustomTagsToRemove, true
}

// HasCustomTagsToRemove returns a boolean if a field has been set.
func (o *EngineAutoTaggingConfigUpdateParameters) HasCustomTagsToRemove() bool {
	if o != nil && !IsNil(o.CustomTagsToRemove) {
		return true
	}

	return false
}

// SetCustomTagsToRemove gets a reference to the given []Tag and assigns it to the CustomTagsToRemove field.
func (o *EngineAutoTaggingConfigUpdateParameters) SetCustomTagsToRemove(v []Tag) {
	o.CustomTagsToRemove = v
}

func (o EngineAutoTaggingConfigUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EngineAutoTaggingConfigUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableVirtualizationDatasetGroups) {
		toSerialize["enable_virtualization_dataset_groups"] = o.EnableVirtualizationDatasetGroups
	}
	if !IsNil(o.EnableMaskingEnvironments) {
		toSerialize["enable_masking_environments"] = o.EnableMaskingEnvironments
	}
	if !IsNil(o.EnableMaskingApplications) {
		toSerialize["enable_masking_applications"] = o.EnableMaskingApplications
	}
	if !IsNil(o.EnableEngineName) {
		toSerialize["enable_engine_name"] = o.EnableEngineName
	}
	if !IsNil(o.CustomTagsToAdd) {
		toSerialize["custom_tags_to_add"] = o.CustomTagsToAdd
	}
	if !IsNil(o.CustomTagsToRemove) {
		toSerialize["custom_tags_to_remove"] = o.CustomTagsToRemove
	}
	return toSerialize, nil
}

type NullableEngineAutoTaggingConfigUpdateParameters struct {
	value *EngineAutoTaggingConfigUpdateParameters
	isSet bool
}

func (v NullableEngineAutoTaggingConfigUpdateParameters) Get() *EngineAutoTaggingConfigUpdateParameters {
	return v.value
}

func (v *NullableEngineAutoTaggingConfigUpdateParameters) Set(val *EngineAutoTaggingConfigUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineAutoTaggingConfigUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineAutoTaggingConfigUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineAutoTaggingConfigUpdateParameters(val *EngineAutoTaggingConfigUpdateParameters) *NullableEngineAutoTaggingConfigUpdateParameters {
	return &NullableEngineAutoTaggingConfigUpdateParameters{value: val, isSet: true}
}

func (v NullableEngineAutoTaggingConfigUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineAutoTaggingConfigUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


