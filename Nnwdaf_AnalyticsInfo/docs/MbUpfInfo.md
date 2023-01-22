# MbUpfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiMbUpfInfoList** | [**[]SnssaiUpfInfoItem**](SnssaiUpfInfoItem.md) |  | 
**MbSmfServingArea** | Pointer to **[]string** |  | [optional] 
**InterfaceMbUpfInfoList** | Pointer to [**[]InterfaceUpfInfoItem**](InterfaceUpfInfoItem.md) |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**SupportedPfcpFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewMbUpfInfo

`func NewMbUpfInfo(sNssaiMbUpfInfoList []SnssaiUpfInfoItem, ) *MbUpfInfo`

NewMbUpfInfo instantiates a new MbUpfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbUpfInfoWithDefaults

`func NewMbUpfInfoWithDefaults() *MbUpfInfo`

NewMbUpfInfoWithDefaults instantiates a new MbUpfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiMbUpfInfoList

`func (o *MbUpfInfo) GetSNssaiMbUpfInfoList() []SnssaiUpfInfoItem`

GetSNssaiMbUpfInfoList returns the SNssaiMbUpfInfoList field if non-nil, zero value otherwise.

### GetSNssaiMbUpfInfoListOk

`func (o *MbUpfInfo) GetSNssaiMbUpfInfoListOk() (*[]SnssaiUpfInfoItem, bool)`

GetSNssaiMbUpfInfoListOk returns a tuple with the SNssaiMbUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiMbUpfInfoList

`func (o *MbUpfInfo) SetSNssaiMbUpfInfoList(v []SnssaiUpfInfoItem)`

SetSNssaiMbUpfInfoList sets SNssaiMbUpfInfoList field to given value.


### GetMbSmfServingArea

`func (o *MbUpfInfo) GetMbSmfServingArea() []string`

GetMbSmfServingArea returns the MbSmfServingArea field if non-nil, zero value otherwise.

### GetMbSmfServingAreaOk

`func (o *MbUpfInfo) GetMbSmfServingAreaOk() (*[]string, bool)`

GetMbSmfServingAreaOk returns a tuple with the MbSmfServingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbSmfServingArea

`func (o *MbUpfInfo) SetMbSmfServingArea(v []string)`

SetMbSmfServingArea sets MbSmfServingArea field to given value.

### HasMbSmfServingArea

`func (o *MbUpfInfo) HasMbSmfServingArea() bool`

HasMbSmfServingArea returns a boolean if a field has been set.

### GetInterfaceMbUpfInfoList

`func (o *MbUpfInfo) GetInterfaceMbUpfInfoList() []InterfaceUpfInfoItem`

GetInterfaceMbUpfInfoList returns the InterfaceMbUpfInfoList field if non-nil, zero value otherwise.

### GetInterfaceMbUpfInfoListOk

`func (o *MbUpfInfo) GetInterfaceMbUpfInfoListOk() (*[]InterfaceUpfInfoItem, bool)`

GetInterfaceMbUpfInfoListOk returns a tuple with the InterfaceMbUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceMbUpfInfoList

`func (o *MbUpfInfo) SetInterfaceMbUpfInfoList(v []InterfaceUpfInfoItem)`

SetInterfaceMbUpfInfoList sets InterfaceMbUpfInfoList field to given value.

### HasInterfaceMbUpfInfoList

`func (o *MbUpfInfo) HasInterfaceMbUpfInfoList() bool`

HasInterfaceMbUpfInfoList returns a boolean if a field has been set.

### GetTaiList

`func (o *MbUpfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *MbUpfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *MbUpfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *MbUpfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *MbUpfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *MbUpfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *MbUpfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *MbUpfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetPriority

`func (o *MbUpfInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MbUpfInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MbUpfInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MbUpfInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSupportedPfcpFeatures

`func (o *MbUpfInfo) GetSupportedPfcpFeatures() string`

GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field if non-nil, zero value otherwise.

### GetSupportedPfcpFeaturesOk

`func (o *MbUpfInfo) GetSupportedPfcpFeaturesOk() (*string, bool)`

GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPfcpFeatures

`func (o *MbUpfInfo) SetSupportedPfcpFeatures(v string)`

SetSupportedPfcpFeatures sets SupportedPfcpFeatures field to given value.

### HasSupportedPfcpFeatures

`func (o *MbUpfInfo) HasSupportedPfcpFeatures() bool`

HasSupportedPfcpFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


