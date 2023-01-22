# TceMappingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TceIPAddress** | Pointer to [**TceMappingInfoTceIPAddress**](TceMappingInfoTceIPAddress.md) |  | [optional] 
**TceID** | Pointer to **int32** |  | [optional] 
**PlmnTarget** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 

## Methods

### NewTceMappingInfo

`func NewTceMappingInfo() *TceMappingInfo`

NewTceMappingInfo instantiates a new TceMappingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTceMappingInfoWithDefaults

`func NewTceMappingInfoWithDefaults() *TceMappingInfo`

NewTceMappingInfoWithDefaults instantiates a new TceMappingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTceIPAddress

`func (o *TceMappingInfo) GetTceIPAddress() TceMappingInfoTceIPAddress`

GetTceIPAddress returns the TceIPAddress field if non-nil, zero value otherwise.

### GetTceIPAddressOk

`func (o *TceMappingInfo) GetTceIPAddressOk() (*TceMappingInfoTceIPAddress, bool)`

GetTceIPAddressOk returns a tuple with the TceIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTceIPAddress

`func (o *TceMappingInfo) SetTceIPAddress(v TceMappingInfoTceIPAddress)`

SetTceIPAddress sets TceIPAddress field to given value.

### HasTceIPAddress

`func (o *TceMappingInfo) HasTceIPAddress() bool`

HasTceIPAddress returns a boolean if a field has been set.

### GetTceID

`func (o *TceMappingInfo) GetTceID() int32`

GetTceID returns the TceID field if non-nil, zero value otherwise.

### GetTceIDOk

`func (o *TceMappingInfo) GetTceIDOk() (*int32, bool)`

GetTceIDOk returns a tuple with the TceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTceID

`func (o *TceMappingInfo) SetTceID(v int32)`

SetTceID sets TceID field to given value.

### HasTceID

`func (o *TceMappingInfo) HasTceID() bool`

HasTceID returns a boolean if a field has been set.

### GetPlmnTarget

`func (o *TceMappingInfo) GetPlmnTarget() PlmnId`

GetPlmnTarget returns the PlmnTarget field if non-nil, zero value otherwise.

### GetPlmnTargetOk

`func (o *TceMappingInfo) GetPlmnTargetOk() (*PlmnId, bool)`

GetPlmnTargetOk returns a tuple with the PlmnTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnTarget

`func (o *TceMappingInfo) SetPlmnTarget(v PlmnId)`

SetPlmnTarget sets PlmnTarget field to given value.

### HasPlmnTarget

`func (o *TceMappingInfo) HasPlmnTarget() bool`

HasPlmnTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


