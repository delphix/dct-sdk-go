# UpdateBookmarkParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The user-defined name of this bookmark. | [optional] 
**Expiration** | Pointer to **string** | The expiration for this Bookmark. Mutually exclusive with retain_forever. | [optional] 
**RetainForever** | Pointer to **bool** | Indicates that the Bookmark should be retained forever. | [optional] 

## Methods

### NewUpdateBookmarkParameters

`func NewUpdateBookmarkParameters() *UpdateBookmarkParameters`

NewUpdateBookmarkParameters instantiates a new UpdateBookmarkParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBookmarkParametersWithDefaults

`func NewUpdateBookmarkParametersWithDefaults() *UpdateBookmarkParameters`

NewUpdateBookmarkParametersWithDefaults instantiates a new UpdateBookmarkParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBookmarkParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBookmarkParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBookmarkParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBookmarkParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiration

`func (o *UpdateBookmarkParameters) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *UpdateBookmarkParameters) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *UpdateBookmarkParameters) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *UpdateBookmarkParameters) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRetainForever

`func (o *UpdateBookmarkParameters) GetRetainForever() bool`

GetRetainForever returns the RetainForever field if non-nil, zero value otherwise.

### GetRetainForeverOk

`func (o *UpdateBookmarkParameters) GetRetainForeverOk() (*bool, bool)`

GetRetainForeverOk returns a tuple with the RetainForever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainForever

`func (o *UpdateBookmarkParameters) SetRetainForever(v bool)`

SetRetainForever sets RetainForever field to given value.

### HasRetainForever

`func (o *UpdateBookmarkParameters) HasRetainForever() bool`

HasRetainForever returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


