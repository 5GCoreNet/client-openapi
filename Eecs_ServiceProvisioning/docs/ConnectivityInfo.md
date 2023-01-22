# ConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**SsId** | Pointer to **string** | Identifies the SSID of the access point to which the UE is attached. | [optional] 

## Methods

### NewConnectivityInfo

`func NewConnectivityInfo() *ConnectivityInfo`

NewConnectivityInfo instantiates a new ConnectivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityInfoWithDefaults

`func NewConnectivityInfoWithDefaults() *ConnectivityInfo`

NewConnectivityInfoWithDefaults instantiates a new ConnectivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *ConnectivityInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *ConnectivityInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *ConnectivityInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *ConnectivityInfo) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetSsId

`func (o *ConnectivityInfo) GetSsId() string`

GetSsId returns the SsId field if non-nil, zero value otherwise.

### GetSsIdOk

`func (o *ConnectivityInfo) GetSsIdOk() (*string, bool)`

GetSsIdOk returns a tuple with the SsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsId

`func (o *ConnectivityInfo) SetSsId(v string)`

SetSsId sets SsId field to given value.

### HasSsId

`func (o *ConnectivityInfo) HasSsId() bool`

HasSsId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


