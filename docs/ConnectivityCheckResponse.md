# ConnectivityCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A message describing the result of the connectivity check. | 
**Status** | Pointer to **string** | A status describing the status of the connectivity check. | [optional] 

## Methods

### NewConnectivityCheckResponse

`func NewConnectivityCheckResponse(message string, ) *ConnectivityCheckResponse`

NewConnectivityCheckResponse instantiates a new ConnectivityCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityCheckResponseWithDefaults

`func NewConnectivityCheckResponseWithDefaults() *ConnectivityCheckResponse`

NewConnectivityCheckResponseWithDefaults instantiates a new ConnectivityCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConnectivityCheckResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectivityCheckResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectivityCheckResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *ConnectivityCheckResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectivityCheckResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectivityCheckResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectivityCheckResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


