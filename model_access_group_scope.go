/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccessGroupScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessGroupScope{}

// AccessGroupScope An Access group scope allows a role to be granted to accounts in the Access group.
type AccessGroupScope struct {
	// The Access group scope ID.
	Id *string `json:"id,omitempty"`
	// The Access group scope name.
	Name *string `json:"name,omitempty"`
	// The Access group role id.
	RoleId string `json:"role_id"`
	// Specifies the type of the scope. Scope of type SIMPLE would grant access to all DCT objects. Scope of type SCOPED would grant access to all objects based on objects and object-tags and permissions defined in linked role. Scope of type ADVANCED would grant access to DCT objects based on objects and object-tags and the individual permissions.
	ScopeType *string `json:"scope_type,omitempty"`
	// The permissions in this access group scope will be granted to all DCT objects tagged with tags matching this property. This is cumulative with objects defined in the 'objects' property, and mutually exclusive with scope_type 'SIMPLE'.
	ObjectTags []ScopeTag `json:"object_tags,omitempty"`
	// The permissions in this access group scope will be granted to all DCT objects with matching object id and object type, mutually exclusive with  scope_type 'SIMPLE'.
	Objects []ScopedObjectItem `json:"objects,omitempty"`
	// An array of always allowed permissions which can be used to specify object type and permission. An object with same object type and permission can not be added in 'objects' array.
	AlwaysAllowedPermissions []AlwaysAllowedPermission `json:"always_allowed_permissions,omitempty"`
}

type _AccessGroupScope AccessGroupScope

// NewAccessGroupScope instantiates a new AccessGroupScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessGroupScope(roleId string) *AccessGroupScope {
	this := AccessGroupScope{}
	this.RoleId = roleId
	return &this
}

// NewAccessGroupScopeWithDefaults instantiates a new AccessGroupScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessGroupScopeWithDefaults() *AccessGroupScope {
	this := AccessGroupScope{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessGroupScope) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupScope) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessGroupScope) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessGroupScope) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessGroupScope) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupScope) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessGroupScope) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessGroupScope) SetName(v string) {
	o.Name = &v
}

// GetRoleId returns the RoleId field value
func (o *AccessGroupScope) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *AccessGroupScope) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *AccessGroupScope) SetRoleId(v string) {
	o.RoleId = v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise.
func (o *AccessGroupScope) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType) {
		var ret string
		return ret
	}
	return *o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupScope) GetScopeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeType) {
		return nil, false
	}
	return o.ScopeType, true
}

// HasScopeType returns a boolean if a field has been set.
func (o *AccessGroupScope) HasScopeType() bool {
	if o != nil && !IsNil(o.ScopeType) {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given string and assigns it to the ScopeType field.
func (o *AccessGroupScope) SetScopeType(v string) {
	o.ScopeType = &v
}

// GetObjectTags returns the ObjectTags field value if set, zero value otherwise.
func (o *AccessGroupScope) GetObjectTags() []ScopeTag {
	if o == nil || IsNil(o.ObjectTags) {
		var ret []ScopeTag
		return ret
	}
	return o.ObjectTags
}

// GetObjectTagsOk returns a tuple with the ObjectTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupScope) GetObjectTagsOk() ([]ScopeTag, bool) {
	if o == nil || IsNil(o.ObjectTags) {
		return nil, false
	}
	return o.ObjectTags, true
}

// HasObjectTags returns a boolean if a field has been set.
func (o *AccessGroupScope) HasObjectTags() bool {
	if o != nil && !IsNil(o.ObjectTags) {
		return true
	}

	return false
}

// SetObjectTags gets a reference to the given []ScopeTag and assigns it to the ObjectTags field.
func (o *AccessGroupScope) SetObjectTags(v []ScopeTag) {
	o.ObjectTags = v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *AccessGroupScope) GetObjects() []ScopedObjectItem {
	if o == nil || IsNil(o.Objects) {
		var ret []ScopedObjectItem
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupScope) GetObjectsOk() ([]ScopedObjectItem, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *AccessGroupScope) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []ScopedObjectItem and assigns it to the Objects field.
func (o *AccessGroupScope) SetObjects(v []ScopedObjectItem) {
	o.Objects = v
}

// GetAlwaysAllowedPermissions returns the AlwaysAllowedPermissions field value if set, zero value otherwise.
func (o *AccessGroupScope) GetAlwaysAllowedPermissions() []AlwaysAllowedPermission {
	if o == nil || IsNil(o.AlwaysAllowedPermissions) {
		var ret []AlwaysAllowedPermission
		return ret
	}
	return o.AlwaysAllowedPermissions
}

// GetAlwaysAllowedPermissionsOk returns a tuple with the AlwaysAllowedPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessGroupScope) GetAlwaysAllowedPermissionsOk() ([]AlwaysAllowedPermission, bool) {
	if o == nil || IsNil(o.AlwaysAllowedPermissions) {
		return nil, false
	}
	return o.AlwaysAllowedPermissions, true
}

// HasAlwaysAllowedPermissions returns a boolean if a field has been set.
func (o *AccessGroupScope) HasAlwaysAllowedPermissions() bool {
	if o != nil && !IsNil(o.AlwaysAllowedPermissions) {
		return true
	}

	return false
}

// SetAlwaysAllowedPermissions gets a reference to the given []AlwaysAllowedPermission and assigns it to the AlwaysAllowedPermissions field.
func (o *AccessGroupScope) SetAlwaysAllowedPermissions(v []AlwaysAllowedPermission) {
	o.AlwaysAllowedPermissions = v
}

func (o AccessGroupScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessGroupScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["role_id"] = o.RoleId
	if !IsNil(o.ScopeType) {
		toSerialize["scope_type"] = o.ScopeType
	}
	if !IsNil(o.ObjectTags) {
		toSerialize["object_tags"] = o.ObjectTags
	}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.AlwaysAllowedPermissions) {
		toSerialize["always_allowed_permissions"] = o.AlwaysAllowedPermissions
	}
	return toSerialize, nil
}

func (o *AccessGroupScope) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role_id",
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

	varAccessGroupScope := _AccessGroupScope{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessGroupScope)

	if err != nil {
		return err
	}

	*o = AccessGroupScope(varAccessGroupScope)

	return err
}

type NullableAccessGroupScope struct {
	value *AccessGroupScope
	isSet bool
}

func (v NullableAccessGroupScope) Get() *AccessGroupScope {
	return v.value
}

func (v *NullableAccessGroupScope) Set(val *AccessGroupScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessGroupScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessGroupScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessGroupScope(val *AccessGroupScope) *NullableAccessGroupScope {
	return &NullableAccessGroupScope{value: val, isSet: true}
}

func (v NullableAccessGroupScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessGroupScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


