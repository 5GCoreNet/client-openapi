# MbSmfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiInfoList** | Pointer to [**map[string]SnssaiMbSmfInfoItem**](SnssaiMbSmfInfoItem.md) | A map (list of key-value pairs) where a valid JSON string serves as key | [optional] 
**TmgiRangeList** | Pointer to [**map[string]TmgiRange**](TmgiRange.md) | A map (list of key-value pairs) where a valid JSON string serves as key | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**MbsSessionList** | Pointer to [**map[string]MbsSession**](MbsSession.md) | A map (list of key-value pairs) where a valid JSON string serves as key | [optional] 

## Methods

### NewMbSmfInfo

`func NewMbSmfInfo() *MbSmfInfo`

NewMbSmfInfo instantiates a new MbSmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbSmfInfoWithDefaults

`func NewMbSmfInfoWithDefaults() *MbSmfInfo`

NewMbSmfInfoWithDefaults instantiates a new MbSmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiInfoList

`func (o *MbSmfInfo) GetSNssaiInfoList() map[string]SnssaiMbSmfInfoItem`

GetSNssaiInfoList returns the SNssaiInfoList field if non-nil, zero value otherwise.

### GetSNssaiInfoListOk

`func (o *MbSmfInfo) GetSNssaiInfoListOk() (*map[string]SnssaiMbSmfInfoItem, bool)`

GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiInfoList

`func (o *MbSmfInfo) SetSNssaiInfoList(v map[string]SnssaiMbSmfInfoItem)`

SetSNssaiInfoList sets SNssaiInfoList field to given value.

### HasSNssaiInfoList

`func (o *MbSmfInfo) HasSNssaiInfoList() bool`

HasSNssaiInfoList returns a boolean if a field has been set.

### GetTmgiRangeList

`func (o *MbSmfInfo) GetTmgiRangeList() map[string]TmgiRange`

GetTmgiRangeList returns the TmgiRangeList field if non-nil, zero value otherwise.

### GetTmgiRangeListOk

`func (o *MbSmfInfo) GetTmgiRangeListOk() (*map[string]TmgiRange, bool)`

GetTmgiRangeListOk returns a tuple with the TmgiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgiRangeList

`func (o *MbSmfInfo) SetTmgiRangeList(v map[string]TmgiRange)`

SetTmgiRangeList sets TmgiRangeList field to given value.

### HasTmgiRangeList

`func (o *MbSmfInfo) HasTmgiRangeList() bool`

HasTmgiRangeList returns a boolean if a field has been set.

### GetTaiList

`func (o *MbSmfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *MbSmfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *MbSmfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *MbSmfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *MbSmfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *MbSmfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *MbSmfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *MbSmfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetMbsSessionList

`func (o *MbSmfInfo) GetMbsSessionList() map[string]MbsSession`

GetMbsSessionList returns the MbsSessionList field if non-nil, zero value otherwise.

### GetMbsSessionListOk

`func (o *MbSmfInfo) GetMbsSessionListOk() (*map[string]MbsSession, bool)`

GetMbsSessionListOk returns a tuple with the MbsSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionList

`func (o *MbSmfInfo) SetMbsSessionList(v map[string]MbsSession)`

SetMbsSessionList sets MbsSessionList field to given value.

### HasMbsSessionList

`func (o *MbSmfInfo) HasMbsSessionList() bool`

HasMbsSessionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


