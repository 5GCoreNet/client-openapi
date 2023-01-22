# IwmscInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsisdnRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**ScNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewIwmscInfo

`func NewIwmscInfo() *IwmscInfo`

NewIwmscInfo instantiates a new IwmscInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIwmscInfoWithDefaults

`func NewIwmscInfoWithDefaults() *IwmscInfo`

NewIwmscInfoWithDefaults instantiates a new IwmscInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsisdnRanges

`func (o *IwmscInfo) GetMsisdnRanges() []IdentityRange`

GetMsisdnRanges returns the MsisdnRanges field if non-nil, zero value otherwise.

### GetMsisdnRangesOk

`func (o *IwmscInfo) GetMsisdnRangesOk() (*[]IdentityRange, bool)`

GetMsisdnRangesOk returns a tuple with the MsisdnRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdnRanges

`func (o *IwmscInfo) SetMsisdnRanges(v []IdentityRange)`

SetMsisdnRanges sets MsisdnRanges field to given value.

### HasMsisdnRanges

`func (o *IwmscInfo) HasMsisdnRanges() bool`

HasMsisdnRanges returns a boolean if a field has been set.

### GetSupiRanges

`func (o *IwmscInfo) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *IwmscInfo) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *IwmscInfo) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *IwmscInfo) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *IwmscInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *IwmscInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *IwmscInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *IwmscInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetScNumber

`func (o *IwmscInfo) GetScNumber() string`

GetScNumber returns the ScNumber field if non-nil, zero value otherwise.

### GetScNumberOk

`func (o *IwmscInfo) GetScNumberOk() (*string, bool)`

GetScNumberOk returns a tuple with the ScNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScNumber

`func (o *IwmscInfo) SetScNumber(v string)`

SetScNumber sets ScNumber field to given value.

### HasScNumber

`func (o *IwmscInfo) HasScNumber() bool`

HasScNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


