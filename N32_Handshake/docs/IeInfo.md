# IeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IeLoc** | [**IeLocation**](IeLocation.md) |  | 
**IeType** | [**IeType**](IeType.md) |  | 
**ReqIe** | Pointer to **string** |  | [optional] 
**RspIe** | Pointer to **string** |  | [optional] 
**IsModifiable** | Pointer to **bool** |  | [optional] 
**IsModifiableByIpx** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewIeInfo

`func NewIeInfo(ieLoc IeLocation, ieType IeType, ) *IeInfo`

NewIeInfo instantiates a new IeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIeInfoWithDefaults

`func NewIeInfoWithDefaults() *IeInfo`

NewIeInfoWithDefaults instantiates a new IeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIeLoc

`func (o *IeInfo) GetIeLoc() IeLocation`

GetIeLoc returns the IeLoc field if non-nil, zero value otherwise.

### GetIeLocOk

`func (o *IeInfo) GetIeLocOk() (*IeLocation, bool)`

GetIeLocOk returns a tuple with the IeLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIeLoc

`func (o *IeInfo) SetIeLoc(v IeLocation)`

SetIeLoc sets IeLoc field to given value.


### GetIeType

`func (o *IeInfo) GetIeType() IeType`

GetIeType returns the IeType field if non-nil, zero value otherwise.

### GetIeTypeOk

`func (o *IeInfo) GetIeTypeOk() (*IeType, bool)`

GetIeTypeOk returns a tuple with the IeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIeType

`func (o *IeInfo) SetIeType(v IeType)`

SetIeType sets IeType field to given value.


### GetReqIe

`func (o *IeInfo) GetReqIe() string`

GetReqIe returns the ReqIe field if non-nil, zero value otherwise.

### GetReqIeOk

`func (o *IeInfo) GetReqIeOk() (*string, bool)`

GetReqIeOk returns a tuple with the ReqIe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqIe

`func (o *IeInfo) SetReqIe(v string)`

SetReqIe sets ReqIe field to given value.

### HasReqIe

`func (o *IeInfo) HasReqIe() bool`

HasReqIe returns a boolean if a field has been set.

### GetRspIe

`func (o *IeInfo) GetRspIe() string`

GetRspIe returns the RspIe field if non-nil, zero value otherwise.

### GetRspIeOk

`func (o *IeInfo) GetRspIeOk() (*string, bool)`

GetRspIeOk returns a tuple with the RspIe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRspIe

`func (o *IeInfo) SetRspIe(v string)`

SetRspIe sets RspIe field to given value.

### HasRspIe

`func (o *IeInfo) HasRspIe() bool`

HasRspIe returns a boolean if a field has been set.

### GetIsModifiable

`func (o *IeInfo) GetIsModifiable() bool`

GetIsModifiable returns the IsModifiable field if non-nil, zero value otherwise.

### GetIsModifiableOk

`func (o *IeInfo) GetIsModifiableOk() (*bool, bool)`

GetIsModifiableOk returns a tuple with the IsModifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModifiable

`func (o *IeInfo) SetIsModifiable(v bool)`

SetIsModifiable sets IsModifiable field to given value.

### HasIsModifiable

`func (o *IeInfo) HasIsModifiable() bool`

HasIsModifiable returns a boolean if a field has been set.

### GetIsModifiableByIpx

`func (o *IeInfo) GetIsModifiableByIpx() map[string]bool`

GetIsModifiableByIpx returns the IsModifiableByIpx field if non-nil, zero value otherwise.

### GetIsModifiableByIpxOk

`func (o *IeInfo) GetIsModifiableByIpxOk() (*map[string]bool, bool)`

GetIsModifiableByIpxOk returns a tuple with the IsModifiableByIpx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModifiableByIpx

`func (o *IeInfo) SetIsModifiableByIpx(v map[string]bool)`

SetIsModifiableByIpx sets IsModifiableByIpx field to given value.

### HasIsModifiableByIpx

`func (o *IeInfo) HasIsModifiableByIpx() bool`

HasIsModifiableByIpx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


