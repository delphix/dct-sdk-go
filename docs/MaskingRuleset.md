# MaskingRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | The Ruleset type. | [optional] 
**Name** | Pointer to **string** | The name of this Ruleset. | [optional] 
**RefreshDropsTables** | Pointer to **NullableBool** | Whether refresh drops tables. Only applicable to database ruleset type. | [optional] 
**Algorithms** | Pointer to **[]string** | The list of algorithms for this Ruleset. | [optional] 

## Methods

### NewMaskingRuleset

`func NewMaskingRuleset() *MaskingRuleset`

NewMaskingRuleset instantiates a new MaskingRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskingRulesetWithDefaults

`func NewMaskingRulesetWithDefaults() *MaskingRuleset`

NewMaskingRulesetWithDefaults instantiates a new MaskingRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MaskingRuleset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaskingRuleset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaskingRuleset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MaskingRuleset) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MaskingRuleset) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MaskingRuleset) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *MaskingRuleset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaskingRuleset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaskingRuleset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MaskingRuleset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefreshDropsTables

`func (o *MaskingRuleset) GetRefreshDropsTables() bool`

GetRefreshDropsTables returns the RefreshDropsTables field if non-nil, zero value otherwise.

### GetRefreshDropsTablesOk

`func (o *MaskingRuleset) GetRefreshDropsTablesOk() (*bool, bool)`

GetRefreshDropsTablesOk returns a tuple with the RefreshDropsTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshDropsTables

`func (o *MaskingRuleset) SetRefreshDropsTables(v bool)`

SetRefreshDropsTables sets RefreshDropsTables field to given value.

### HasRefreshDropsTables

`func (o *MaskingRuleset) HasRefreshDropsTables() bool`

HasRefreshDropsTables returns a boolean if a field has been set.

### SetRefreshDropsTablesNil

`func (o *MaskingRuleset) SetRefreshDropsTablesNil(b bool)`

 SetRefreshDropsTablesNil sets the value for RefreshDropsTables to be an explicit nil

### UnsetRefreshDropsTables
`func (o *MaskingRuleset) UnsetRefreshDropsTables()`

UnsetRefreshDropsTables ensures that no value is present for RefreshDropsTables, not even an explicit nil
### GetAlgorithms

`func (o *MaskingRuleset) GetAlgorithms() []string`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *MaskingRuleset) GetAlgorithmsOk() (*[]string, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *MaskingRuleset) SetAlgorithms(v []string)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *MaskingRuleset) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


