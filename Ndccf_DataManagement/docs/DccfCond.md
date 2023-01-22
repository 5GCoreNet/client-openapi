# DccfCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | **string** |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**ServingNfTypeList** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**ServingNfSetIdList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDccfCond

`func NewDccfCond(conditionType string, ) *DccfCond`

NewDccfCond instantiates a new DccfCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDccfCondWithDefaults

`func NewDccfCondWithDefaults() *DccfCond`

NewDccfCondWithDefaults instantiates a new DccfCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *DccfCond) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *DccfCond) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *DccfCond) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetTaiList

`func (o *DccfCond) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *DccfCond) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *DccfCond) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *DccfCond) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *DccfCond) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *DccfCond) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *DccfCond) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *DccfCond) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetServingNfTypeList

`func (o *DccfCond) GetServingNfTypeList() []NFType`

GetServingNfTypeList returns the ServingNfTypeList field if non-nil, zero value otherwise.

### GetServingNfTypeListOk

`func (o *DccfCond) GetServingNfTypeListOk() (*[]NFType, bool)`

GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfTypeList

`func (o *DccfCond) SetServingNfTypeList(v []NFType)`

SetServingNfTypeList sets ServingNfTypeList field to given value.

### HasServingNfTypeList

`func (o *DccfCond) HasServingNfTypeList() bool`

HasServingNfTypeList returns a boolean if a field has been set.

### GetServingNfSetIdList

`func (o *DccfCond) GetServingNfSetIdList() []string`

GetServingNfSetIdList returns the ServingNfSetIdList field if non-nil, zero value otherwise.

### GetServingNfSetIdListOk

`func (o *DccfCond) GetServingNfSetIdListOk() (*[]string, bool)`

GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfSetIdList

`func (o *DccfCond) SetServingNfSetIdList(v []string)`

SetServingNfSetIdList sets ServingNfSetIdList field to given value.

### HasServingNfSetIdList

`func (o *DccfCond) HasServingNfSetIdList() bool`

HasServingNfSetIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


