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

// checks if the HyperscaleConnectorCreateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleConnectorCreateParameters{}

// HyperscaleConnectorCreateParameters Parameters to create a Hyperscale Connector.
type HyperscaleConnectorCreateParameters struct {
	// The ID of the Hyperscale instance of this Connector.
	HyperscaleInstanceId *string `json:"hyperscale_instance_id,omitempty"`
	DataType *HyperscaleDataTypeEnum `json:"data_type,omitempty"`
	Name *string `json:"name,omitempty"`
	// The username this Connector will use to connect to the source database.
	SourceUsername *string `json:"source_username,omitempty"`
	// The password this Connector will use to connect to the source database.
	SourcePassword *string `json:"source_password,omitempty"`
	// The JDBC URL used to connect to the source database.
	SourceJdbcUrl *string `json:"source_jdbc_url,omitempty"`
	// The MongoDB URL used to connect to the source database.
	SourceMongoUrl *string `json:"source_mongo_url,omitempty"`
	// The path on the filesystem where source files must be read (Delimited files Only).
	SourceFilesystemPath *string `json:"source_filesystem_path,omitempty"`
	SourceConnectionProperties *map[string]string `json:"source_connection_properties,omitempty"`
	// The username this Connector will use to connect to the target database.
	TargetUsername *string `json:"target_username,omitempty"`
	// The username this Connector will use to connect to the target database.
	TargetPassword *string `json:"target_password,omitempty"`
	// The JDBC URL used to connect to the target database.
	TargetJdbcUrl *string `json:"target_jdbc_url,omitempty"`
	// The MongoDB URL used to connect to the target database.
	TargetMongoUrl *string `json:"target_mongo_url,omitempty"`
	// The path on the filesystem where target files must be written (Delimited files Only).
	TargetFilesystemPath *string `json:"target_filesystem_path,omitempty"`
	TargetConnectionProperties *map[string]string `json:"target_connection_properties,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
	// Whether the account creating this Hyperscale Connector must be configured as owner of it.
	MakeCurrentAccountOwner *bool `json:"make_current_account_owner,omitempty"`
}

// NewHyperscaleConnectorCreateParameters instantiates a new HyperscaleConnectorCreateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleConnectorCreateParameters() *HyperscaleConnectorCreateParameters {
	this := HyperscaleConnectorCreateParameters{}
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// NewHyperscaleConnectorCreateParametersWithDefaults instantiates a new HyperscaleConnectorCreateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleConnectorCreateParametersWithDefaults() *HyperscaleConnectorCreateParameters {
	this := HyperscaleConnectorCreateParameters{}
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// GetHyperscaleInstanceId returns the HyperscaleInstanceId field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetHyperscaleInstanceId() string {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		var ret string
		return ret
	}
	return *o.HyperscaleInstanceId
}

// GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetHyperscaleInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		return nil, false
	}
	return o.HyperscaleInstanceId, true
}

// HasHyperscaleInstanceId returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasHyperscaleInstanceId() bool {
	if o != nil && !IsNil(o.HyperscaleInstanceId) {
		return true
	}

	return false
}

// SetHyperscaleInstanceId gets a reference to the given string and assigns it to the HyperscaleInstanceId field.
func (o *HyperscaleConnectorCreateParameters) SetHyperscaleInstanceId(v string) {
	o.HyperscaleInstanceId = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetDataType() HyperscaleDataTypeEnum {
	if o == nil || IsNil(o.DataType) {
		var ret HyperscaleDataTypeEnum
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetDataTypeOk() (*HyperscaleDataTypeEnum, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given HyperscaleDataTypeEnum and assigns it to the DataType field.
func (o *HyperscaleConnectorCreateParameters) SetDataType(v HyperscaleDataTypeEnum) {
	o.DataType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperscaleConnectorCreateParameters) SetName(v string) {
	o.Name = &v
}

// GetSourceUsername returns the SourceUsername field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetSourceUsername() string {
	if o == nil || IsNil(o.SourceUsername) {
		var ret string
		return ret
	}
	return *o.SourceUsername
}

// GetSourceUsernameOk returns a tuple with the SourceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetSourceUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceUsername) {
		return nil, false
	}
	return o.SourceUsername, true
}

// HasSourceUsername returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasSourceUsername() bool {
	if o != nil && !IsNil(o.SourceUsername) {
		return true
	}

	return false
}

// SetSourceUsername gets a reference to the given string and assigns it to the SourceUsername field.
func (o *HyperscaleConnectorCreateParameters) SetSourceUsername(v string) {
	o.SourceUsername = &v
}

// GetSourcePassword returns the SourcePassword field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetSourcePassword() string {
	if o == nil || IsNil(o.SourcePassword) {
		var ret string
		return ret
	}
	return *o.SourcePassword
}

// GetSourcePasswordOk returns a tuple with the SourcePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetSourcePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SourcePassword) {
		return nil, false
	}
	return o.SourcePassword, true
}

// HasSourcePassword returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasSourcePassword() bool {
	if o != nil && !IsNil(o.SourcePassword) {
		return true
	}

	return false
}

// SetSourcePassword gets a reference to the given string and assigns it to the SourcePassword field.
func (o *HyperscaleConnectorCreateParameters) SetSourcePassword(v string) {
	o.SourcePassword = &v
}

// GetSourceJdbcUrl returns the SourceJdbcUrl field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetSourceJdbcUrl() string {
	if o == nil || IsNil(o.SourceJdbcUrl) {
		var ret string
		return ret
	}
	return *o.SourceJdbcUrl
}

// GetSourceJdbcUrlOk returns a tuple with the SourceJdbcUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetSourceJdbcUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceJdbcUrl) {
		return nil, false
	}
	return o.SourceJdbcUrl, true
}

// HasSourceJdbcUrl returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasSourceJdbcUrl() bool {
	if o != nil && !IsNil(o.SourceJdbcUrl) {
		return true
	}

	return false
}

// SetSourceJdbcUrl gets a reference to the given string and assigns it to the SourceJdbcUrl field.
func (o *HyperscaleConnectorCreateParameters) SetSourceJdbcUrl(v string) {
	o.SourceJdbcUrl = &v
}

// GetSourceMongoUrl returns the SourceMongoUrl field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetSourceMongoUrl() string {
	if o == nil || IsNil(o.SourceMongoUrl) {
		var ret string
		return ret
	}
	return *o.SourceMongoUrl
}

// GetSourceMongoUrlOk returns a tuple with the SourceMongoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetSourceMongoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceMongoUrl) {
		return nil, false
	}
	return o.SourceMongoUrl, true
}

// HasSourceMongoUrl returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasSourceMongoUrl() bool {
	if o != nil && !IsNil(o.SourceMongoUrl) {
		return true
	}

	return false
}

// SetSourceMongoUrl gets a reference to the given string and assigns it to the SourceMongoUrl field.
func (o *HyperscaleConnectorCreateParameters) SetSourceMongoUrl(v string) {
	o.SourceMongoUrl = &v
}

// GetSourceFilesystemPath returns the SourceFilesystemPath field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetSourceFilesystemPath() string {
	if o == nil || IsNil(o.SourceFilesystemPath) {
		var ret string
		return ret
	}
	return *o.SourceFilesystemPath
}

// GetSourceFilesystemPathOk returns a tuple with the SourceFilesystemPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetSourceFilesystemPathOk() (*string, bool) {
	if o == nil || IsNil(o.SourceFilesystemPath) {
		return nil, false
	}
	return o.SourceFilesystemPath, true
}

// HasSourceFilesystemPath returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasSourceFilesystemPath() bool {
	if o != nil && !IsNil(o.SourceFilesystemPath) {
		return true
	}

	return false
}

// SetSourceFilesystemPath gets a reference to the given string and assigns it to the SourceFilesystemPath field.
func (o *HyperscaleConnectorCreateParameters) SetSourceFilesystemPath(v string) {
	o.SourceFilesystemPath = &v
}

// GetSourceConnectionProperties returns the SourceConnectionProperties field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetSourceConnectionProperties() map[string]string {
	if o == nil || IsNil(o.SourceConnectionProperties) {
		var ret map[string]string
		return ret
	}
	return *o.SourceConnectionProperties
}

// GetSourceConnectionPropertiesOk returns a tuple with the SourceConnectionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetSourceConnectionPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.SourceConnectionProperties) {
		return nil, false
	}
	return o.SourceConnectionProperties, true
}

// HasSourceConnectionProperties returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasSourceConnectionProperties() bool {
	if o != nil && !IsNil(o.SourceConnectionProperties) {
		return true
	}

	return false
}

// SetSourceConnectionProperties gets a reference to the given map[string]string and assigns it to the SourceConnectionProperties field.
func (o *HyperscaleConnectorCreateParameters) SetSourceConnectionProperties(v map[string]string) {
	o.SourceConnectionProperties = &v
}

// GetTargetUsername returns the TargetUsername field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetTargetUsername() string {
	if o == nil || IsNil(o.TargetUsername) {
		var ret string
		return ret
	}
	return *o.TargetUsername
}

// GetTargetUsernameOk returns a tuple with the TargetUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetTargetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUsername) {
		return nil, false
	}
	return o.TargetUsername, true
}

// HasTargetUsername returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasTargetUsername() bool {
	if o != nil && !IsNil(o.TargetUsername) {
		return true
	}

	return false
}

// SetTargetUsername gets a reference to the given string and assigns it to the TargetUsername field.
func (o *HyperscaleConnectorCreateParameters) SetTargetUsername(v string) {
	o.TargetUsername = &v
}

// GetTargetPassword returns the TargetPassword field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetTargetPassword() string {
	if o == nil || IsNil(o.TargetPassword) {
		var ret string
		return ret
	}
	return *o.TargetPassword
}

// GetTargetPasswordOk returns a tuple with the TargetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetTargetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPassword) {
		return nil, false
	}
	return o.TargetPassword, true
}

// HasTargetPassword returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasTargetPassword() bool {
	if o != nil && !IsNil(o.TargetPassword) {
		return true
	}

	return false
}

// SetTargetPassword gets a reference to the given string and assigns it to the TargetPassword field.
func (o *HyperscaleConnectorCreateParameters) SetTargetPassword(v string) {
	o.TargetPassword = &v
}

// GetTargetJdbcUrl returns the TargetJdbcUrl field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetTargetJdbcUrl() string {
	if o == nil || IsNil(o.TargetJdbcUrl) {
		var ret string
		return ret
	}
	return *o.TargetJdbcUrl
}

// GetTargetJdbcUrlOk returns a tuple with the TargetJdbcUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetTargetJdbcUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetJdbcUrl) {
		return nil, false
	}
	return o.TargetJdbcUrl, true
}

// HasTargetJdbcUrl returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasTargetJdbcUrl() bool {
	if o != nil && !IsNil(o.TargetJdbcUrl) {
		return true
	}

	return false
}

// SetTargetJdbcUrl gets a reference to the given string and assigns it to the TargetJdbcUrl field.
func (o *HyperscaleConnectorCreateParameters) SetTargetJdbcUrl(v string) {
	o.TargetJdbcUrl = &v
}

// GetTargetMongoUrl returns the TargetMongoUrl field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetTargetMongoUrl() string {
	if o == nil || IsNil(o.TargetMongoUrl) {
		var ret string
		return ret
	}
	return *o.TargetMongoUrl
}

// GetTargetMongoUrlOk returns a tuple with the TargetMongoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetTargetMongoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMongoUrl) {
		return nil, false
	}
	return o.TargetMongoUrl, true
}

// HasTargetMongoUrl returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasTargetMongoUrl() bool {
	if o != nil && !IsNil(o.TargetMongoUrl) {
		return true
	}

	return false
}

// SetTargetMongoUrl gets a reference to the given string and assigns it to the TargetMongoUrl field.
func (o *HyperscaleConnectorCreateParameters) SetTargetMongoUrl(v string) {
	o.TargetMongoUrl = &v
}

// GetTargetFilesystemPath returns the TargetFilesystemPath field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetTargetFilesystemPath() string {
	if o == nil || IsNil(o.TargetFilesystemPath) {
		var ret string
		return ret
	}
	return *o.TargetFilesystemPath
}

// GetTargetFilesystemPathOk returns a tuple with the TargetFilesystemPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetTargetFilesystemPathOk() (*string, bool) {
	if o == nil || IsNil(o.TargetFilesystemPath) {
		return nil, false
	}
	return o.TargetFilesystemPath, true
}

// HasTargetFilesystemPath returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasTargetFilesystemPath() bool {
	if o != nil && !IsNil(o.TargetFilesystemPath) {
		return true
	}

	return false
}

// SetTargetFilesystemPath gets a reference to the given string and assigns it to the TargetFilesystemPath field.
func (o *HyperscaleConnectorCreateParameters) SetTargetFilesystemPath(v string) {
	o.TargetFilesystemPath = &v
}

// GetTargetConnectionProperties returns the TargetConnectionProperties field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetTargetConnectionProperties() map[string]string {
	if o == nil || IsNil(o.TargetConnectionProperties) {
		var ret map[string]string
		return ret
	}
	return *o.TargetConnectionProperties
}

// GetTargetConnectionPropertiesOk returns a tuple with the TargetConnectionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetTargetConnectionPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TargetConnectionProperties) {
		return nil, false
	}
	return o.TargetConnectionProperties, true
}

// HasTargetConnectionProperties returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasTargetConnectionProperties() bool {
	if o != nil && !IsNil(o.TargetConnectionProperties) {
		return true
	}

	return false
}

// SetTargetConnectionProperties gets a reference to the given map[string]string and assigns it to the TargetConnectionProperties field.
func (o *HyperscaleConnectorCreateParameters) SetTargetConnectionProperties(v map[string]string) {
	o.TargetConnectionProperties = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *HyperscaleConnectorCreateParameters) SetTags(v []Tag) {
	o.Tags = v
}

// GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field value if set, zero value otherwise.
func (o *HyperscaleConnectorCreateParameters) GetMakeCurrentAccountOwner() bool {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		var ret bool
		return ret
	}
	return *o.MakeCurrentAccountOwner
}

// GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleConnectorCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		return nil, false
	}
	return o.MakeCurrentAccountOwner, true
}

// HasMakeCurrentAccountOwner returns a boolean if a field has been set.
func (o *HyperscaleConnectorCreateParameters) HasMakeCurrentAccountOwner() bool {
	if o != nil && !IsNil(o.MakeCurrentAccountOwner) {
		return true
	}

	return false
}

// SetMakeCurrentAccountOwner gets a reference to the given bool and assigns it to the MakeCurrentAccountOwner field.
func (o *HyperscaleConnectorCreateParameters) SetMakeCurrentAccountOwner(v bool) {
	o.MakeCurrentAccountOwner = &v
}

func (o HyperscaleConnectorCreateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleConnectorCreateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HyperscaleInstanceId) {
		toSerialize["hyperscale_instance_id"] = o.HyperscaleInstanceId
	}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SourceUsername) {
		toSerialize["source_username"] = o.SourceUsername
	}
	if !IsNil(o.SourcePassword) {
		toSerialize["source_password"] = o.SourcePassword
	}
	if !IsNil(o.SourceJdbcUrl) {
		toSerialize["source_jdbc_url"] = o.SourceJdbcUrl
	}
	if !IsNil(o.SourceMongoUrl) {
		toSerialize["source_mongo_url"] = o.SourceMongoUrl
	}
	if !IsNil(o.SourceFilesystemPath) {
		toSerialize["source_filesystem_path"] = o.SourceFilesystemPath
	}
	if !IsNil(o.SourceConnectionProperties) {
		toSerialize["source_connection_properties"] = o.SourceConnectionProperties
	}
	if !IsNil(o.TargetUsername) {
		toSerialize["target_username"] = o.TargetUsername
	}
	if !IsNil(o.TargetPassword) {
		toSerialize["target_password"] = o.TargetPassword
	}
	if !IsNil(o.TargetJdbcUrl) {
		toSerialize["target_jdbc_url"] = o.TargetJdbcUrl
	}
	if !IsNil(o.TargetMongoUrl) {
		toSerialize["target_mongo_url"] = o.TargetMongoUrl
	}
	if !IsNil(o.TargetFilesystemPath) {
		toSerialize["target_filesystem_path"] = o.TargetFilesystemPath
	}
	if !IsNil(o.TargetConnectionProperties) {
		toSerialize["target_connection_properties"] = o.TargetConnectionProperties
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.MakeCurrentAccountOwner) {
		toSerialize["make_current_account_owner"] = o.MakeCurrentAccountOwner
	}
	return toSerialize, nil
}

type NullableHyperscaleConnectorCreateParameters struct {
	value *HyperscaleConnectorCreateParameters
	isSet bool
}

func (v NullableHyperscaleConnectorCreateParameters) Get() *HyperscaleConnectorCreateParameters {
	return v.value
}

func (v *NullableHyperscaleConnectorCreateParameters) Set(val *HyperscaleConnectorCreateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleConnectorCreateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleConnectorCreateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleConnectorCreateParameters(val *HyperscaleConnectorCreateParameters) *NullableHyperscaleConnectorCreateParameters {
	return &NullableHyperscaleConnectorCreateParameters{value: val, isSet: true}
}

func (v NullableHyperscaleConnectorCreateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleConnectorCreateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


