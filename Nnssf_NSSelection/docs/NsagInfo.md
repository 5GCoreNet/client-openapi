# NsagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsagIds** | **[]int32** |  | 
**SnssaiList** | [**[]Snssai**](Snssai.md) |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 

## Methods

### NewNsagInfo

`func NewNsagInfo(nsagIds []int32, snssaiList []Snssai, ) *NsagInfo`

NewNsagInfo instantiates a new NsagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsagInfoWithDefaults

`func NewNsagInfoWithDefaults() *NsagInfo`

NewNsagInfoWithDefaults instantiates a new NsagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNsagIds

`func (o *NsagInfo) GetNsagIds() []int32`

GetNsagIds returns the NsagIds field if non-nil, zero value otherwise.

### GetNsagIdsOk

`func (o *NsagInfo) GetNsagIdsOk() (*[]int32, bool)`

GetNsagIdsOk returns a tuple with the NsagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsagIds

`func (o *NsagInfo) SetNsagIds(v []int32)`

SetNsagIds sets NsagIds field to given value.


### GetSnssaiList

`func (o *NsagInfo) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *NsagInfo) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *NsagInfo) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.


### GetTaiList

`func (o *NsagInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NsagInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NsagInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NsagInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NsagInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NsagInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NsagInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NsagInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


