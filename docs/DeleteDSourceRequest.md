# DeleteDSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DsourceId** | **string** | Id of the dSource to delete. | 
**Force** | Pointer to **bool** | Flag indicating whether to continue the operation upon failures. | [optional] [default to false]
**OracleUsername** | Pointer to **string** | The name of the privileged user to run the delete operation as (Oracle only). | [optional] 
**OraclePassword** | Pointer to **string** | Password for privileged user (Oracle only). | [optional] 

## Methods

### NewDeleteDSourceRequest

`func NewDeleteDSourceRequest(dsourceId string, ) *DeleteDSourceRequest`

NewDeleteDSourceRequest instantiates a new DeleteDSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDSourceRequestWithDefaults

`func NewDeleteDSourceRequestWithDefaults() *DeleteDSourceRequest`

NewDeleteDSourceRequestWithDefaults instantiates a new DeleteDSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDsourceId

`func (o *DeleteDSourceRequest) GetDsourceId() string`

GetDsourceId returns the DsourceId field if non-nil, zero value otherwise.

### GetDsourceIdOk

`func (o *DeleteDSourceRequest) GetDsourceIdOk() (*string, bool)`

GetDsourceIdOk returns a tuple with the DsourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsourceId

`func (o *DeleteDSourceRequest) SetDsourceId(v string)`

SetDsourceId sets DsourceId field to given value.


### GetForce

`func (o *DeleteDSourceRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *DeleteDSourceRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *DeleteDSourceRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *DeleteDSourceRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetOracleUsername

`func (o *DeleteDSourceRequest) GetOracleUsername() string`

GetOracleUsername returns the OracleUsername field if non-nil, zero value otherwise.

### GetOracleUsernameOk

`func (o *DeleteDSourceRequest) GetOracleUsernameOk() (*string, bool)`

GetOracleUsernameOk returns a tuple with the OracleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleUsername

`func (o *DeleteDSourceRequest) SetOracleUsername(v string)`

SetOracleUsername sets OracleUsername field to given value.

### HasOracleUsername

`func (o *DeleteDSourceRequest) HasOracleUsername() bool`

HasOracleUsername returns a boolean if a field has been set.

### GetOraclePassword

`func (o *DeleteDSourceRequest) GetOraclePassword() string`

GetOraclePassword returns the OraclePassword field if non-nil, zero value otherwise.

### GetOraclePasswordOk

`func (o *DeleteDSourceRequest) GetOraclePasswordOk() (*string, bool)`

GetOraclePasswordOk returns a tuple with the OraclePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOraclePassword

`func (o *DeleteDSourceRequest) SetOraclePassword(v string)`

SetOraclePassword sets OraclePassword field to given value.

### HasOraclePassword

`func (o *DeleteDSourceRequest) HasOraclePassword() bool`

HasOraclePassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


