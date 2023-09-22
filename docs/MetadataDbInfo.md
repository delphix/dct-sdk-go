# MetadataDbInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | Pointer to **bool** | True if an external database, i.e a database running outside of the application cluster, is in use. | [optional] 
**Version** | Pointer to **string** | The full database version in String format | [optional] 
**DatabaseProductName** | Pointer to **string** | The database product name as reported by the database itself. | [optional] 
**MajorVersion** | Pointer to **int32** | The database major version. | [optional] 
**MinorVersion** | Pointer to **int32** | The database minor version | [optional] 
**MinSupportedMajorVersion** | Pointer to **int32** | The minimum supported major version of PostgreSQL. | [optional] 
**MinSupportedMinorVersion** | Pointer to **int32** | The minimum supported minor version of PostgreSQL. | [optional] 
**MaxSupportedMajorVersion** | Pointer to **int32** | The maximum supported major version of PostgreSQL. | [optional] 
**MaxSupportedMinorVersion** | Pointer to **int32** | The maximum supported minor version of PostgreSQL. | [optional] 
**Compatible** | Pointer to **bool** | Whether the database is recognized as valid for this product. In order to be compatible, the database product name must be a recognized PostgreSQL and the database version must be greater than or equal to the minimum minor version and smaller than or equal to the maximum support version. | [optional] 

## Methods

### NewMetadataDbInfo

`func NewMetadataDbInfo() *MetadataDbInfo`

NewMetadataDbInfo instantiates a new MetadataDbInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataDbInfoWithDefaults

`func NewMetadataDbInfoWithDefaults() *MetadataDbInfo`

NewMetadataDbInfoWithDefaults instantiates a new MetadataDbInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *MetadataDbInfo) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *MetadataDbInfo) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *MetadataDbInfo) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *MetadataDbInfo) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetVersion

`func (o *MetadataDbInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetadataDbInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetadataDbInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MetadataDbInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatabaseProductName

`func (o *MetadataDbInfo) GetDatabaseProductName() string`

GetDatabaseProductName returns the DatabaseProductName field if non-nil, zero value otherwise.

### GetDatabaseProductNameOk

`func (o *MetadataDbInfo) GetDatabaseProductNameOk() (*string, bool)`

GetDatabaseProductNameOk returns a tuple with the DatabaseProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseProductName

`func (o *MetadataDbInfo) SetDatabaseProductName(v string)`

SetDatabaseProductName sets DatabaseProductName field to given value.

### HasDatabaseProductName

`func (o *MetadataDbInfo) HasDatabaseProductName() bool`

HasDatabaseProductName returns a boolean if a field has been set.

### GetMajorVersion

`func (o *MetadataDbInfo) GetMajorVersion() int32`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *MetadataDbInfo) GetMajorVersionOk() (*int32, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *MetadataDbInfo) SetMajorVersion(v int32)`

SetMajorVersion sets MajorVersion field to given value.

### HasMajorVersion

`func (o *MetadataDbInfo) HasMajorVersion() bool`

HasMajorVersion returns a boolean if a field has been set.

### GetMinorVersion

`func (o *MetadataDbInfo) GetMinorVersion() int32`

GetMinorVersion returns the MinorVersion field if non-nil, zero value otherwise.

### GetMinorVersionOk

`func (o *MetadataDbInfo) GetMinorVersionOk() (*int32, bool)`

GetMinorVersionOk returns a tuple with the MinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorVersion

`func (o *MetadataDbInfo) SetMinorVersion(v int32)`

SetMinorVersion sets MinorVersion field to given value.

### HasMinorVersion

`func (o *MetadataDbInfo) HasMinorVersion() bool`

HasMinorVersion returns a boolean if a field has been set.

### GetMinSupportedMajorVersion

`func (o *MetadataDbInfo) GetMinSupportedMajorVersion() int32`

GetMinSupportedMajorVersion returns the MinSupportedMajorVersion field if non-nil, zero value otherwise.

### GetMinSupportedMajorVersionOk

`func (o *MetadataDbInfo) GetMinSupportedMajorVersionOk() (*int32, bool)`

GetMinSupportedMajorVersionOk returns a tuple with the MinSupportedMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSupportedMajorVersion

`func (o *MetadataDbInfo) SetMinSupportedMajorVersion(v int32)`

SetMinSupportedMajorVersion sets MinSupportedMajorVersion field to given value.

### HasMinSupportedMajorVersion

`func (o *MetadataDbInfo) HasMinSupportedMajorVersion() bool`

HasMinSupportedMajorVersion returns a boolean if a field has been set.

### GetMinSupportedMinorVersion

`func (o *MetadataDbInfo) GetMinSupportedMinorVersion() int32`

GetMinSupportedMinorVersion returns the MinSupportedMinorVersion field if non-nil, zero value otherwise.

### GetMinSupportedMinorVersionOk

`func (o *MetadataDbInfo) GetMinSupportedMinorVersionOk() (*int32, bool)`

GetMinSupportedMinorVersionOk returns a tuple with the MinSupportedMinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSupportedMinorVersion

`func (o *MetadataDbInfo) SetMinSupportedMinorVersion(v int32)`

SetMinSupportedMinorVersion sets MinSupportedMinorVersion field to given value.

### HasMinSupportedMinorVersion

`func (o *MetadataDbInfo) HasMinSupportedMinorVersion() bool`

HasMinSupportedMinorVersion returns a boolean if a field has been set.

### GetMaxSupportedMajorVersion

`func (o *MetadataDbInfo) GetMaxSupportedMajorVersion() int32`

GetMaxSupportedMajorVersion returns the MaxSupportedMajorVersion field if non-nil, zero value otherwise.

### GetMaxSupportedMajorVersionOk

`func (o *MetadataDbInfo) GetMaxSupportedMajorVersionOk() (*int32, bool)`

GetMaxSupportedMajorVersionOk returns a tuple with the MaxSupportedMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupportedMajorVersion

`func (o *MetadataDbInfo) SetMaxSupportedMajorVersion(v int32)`

SetMaxSupportedMajorVersion sets MaxSupportedMajorVersion field to given value.

### HasMaxSupportedMajorVersion

`func (o *MetadataDbInfo) HasMaxSupportedMajorVersion() bool`

HasMaxSupportedMajorVersion returns a boolean if a field has been set.

### GetMaxSupportedMinorVersion

`func (o *MetadataDbInfo) GetMaxSupportedMinorVersion() int32`

GetMaxSupportedMinorVersion returns the MaxSupportedMinorVersion field if non-nil, zero value otherwise.

### GetMaxSupportedMinorVersionOk

`func (o *MetadataDbInfo) GetMaxSupportedMinorVersionOk() (*int32, bool)`

GetMaxSupportedMinorVersionOk returns a tuple with the MaxSupportedMinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupportedMinorVersion

`func (o *MetadataDbInfo) SetMaxSupportedMinorVersion(v int32)`

SetMaxSupportedMinorVersion sets MaxSupportedMinorVersion field to given value.

### HasMaxSupportedMinorVersion

`func (o *MetadataDbInfo) HasMaxSupportedMinorVersion() bool`

HasMaxSupportedMinorVersion returns a boolean if a field has been set.

### GetCompatible

`func (o *MetadataDbInfo) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *MetadataDbInfo) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *MetadataDbInfo) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *MetadataDbInfo) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


