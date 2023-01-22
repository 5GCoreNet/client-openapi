# NefInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NefId** | Pointer to **string** | Identity of the NEF | [optional] 
**PfdData** | Pointer to [**PfdData**](PfdData.md) |  | [optional] 
**AfEeData** | Pointer to [**AfEventExposureData**](AfEventExposureData.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ServedFqdnList** | Pointer to **[]string** |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**UnTrustAfInfoList** | Pointer to [**[]UnTrustAfInfo**](UnTrustAfInfo.md) |  | [optional] 
**UasNfFunctionalityInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNefInfo

`func NewNefInfo() *NefInfo`

NewNefInfo instantiates a new NefInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefInfoWithDefaults

`func NewNefInfoWithDefaults() *NefInfo`

NewNefInfoWithDefaults instantiates a new NefInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNefId

`func (o *NefInfo) GetNefId() string`

GetNefId returns the NefId field if non-nil, zero value otherwise.

### GetNefIdOk

`func (o *NefInfo) GetNefIdOk() (*string, bool)`

GetNefIdOk returns a tuple with the NefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefId

`func (o *NefInfo) SetNefId(v string)`

SetNefId sets NefId field to given value.

### HasNefId

`func (o *NefInfo) HasNefId() bool`

HasNefId returns a boolean if a field has been set.

### GetPfdData

`func (o *NefInfo) GetPfdData() PfdData`

GetPfdData returns the PfdData field if non-nil, zero value otherwise.

### GetPfdDataOk

`func (o *NefInfo) GetPfdDataOk() (*PfdData, bool)`

GetPfdDataOk returns a tuple with the PfdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdData

`func (o *NefInfo) SetPfdData(v PfdData)`

SetPfdData sets PfdData field to given value.

### HasPfdData

`func (o *NefInfo) HasPfdData() bool`

HasPfdData returns a boolean if a field has been set.

### GetAfEeData

`func (o *NefInfo) GetAfEeData() AfEventExposureData`

GetAfEeData returns the AfEeData field if non-nil, zero value otherwise.

### GetAfEeDataOk

`func (o *NefInfo) GetAfEeDataOk() (*AfEventExposureData, bool)`

GetAfEeDataOk returns a tuple with the AfEeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEeData

`func (o *NefInfo) SetAfEeData(v AfEventExposureData)`

SetAfEeData sets AfEeData field to given value.

### HasAfEeData

`func (o *NefInfo) HasAfEeData() bool`

HasAfEeData returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *NefInfo) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *NefInfo) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *NefInfo) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *NefInfo) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *NefInfo) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *NefInfo) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *NefInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *NefInfo) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetServedFqdnList

`func (o *NefInfo) GetServedFqdnList() []string`

GetServedFqdnList returns the ServedFqdnList field if non-nil, zero value otherwise.

### GetServedFqdnListOk

`func (o *NefInfo) GetServedFqdnListOk() (*[]string, bool)`

GetServedFqdnListOk returns a tuple with the ServedFqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedFqdnList

`func (o *NefInfo) SetServedFqdnList(v []string)`

SetServedFqdnList sets ServedFqdnList field to given value.

### HasServedFqdnList

`func (o *NefInfo) HasServedFqdnList() bool`

HasServedFqdnList returns a boolean if a field has been set.

### GetTaiList

`func (o *NefInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NefInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NefInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NefInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NefInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NefInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NefInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NefInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetDnaiList

`func (o *NefInfo) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *NefInfo) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *NefInfo) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *NefInfo) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetUnTrustAfInfoList

`func (o *NefInfo) GetUnTrustAfInfoList() []UnTrustAfInfo`

GetUnTrustAfInfoList returns the UnTrustAfInfoList field if non-nil, zero value otherwise.

### GetUnTrustAfInfoListOk

`func (o *NefInfo) GetUnTrustAfInfoListOk() (*[]UnTrustAfInfo, bool)`

GetUnTrustAfInfoListOk returns a tuple with the UnTrustAfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnTrustAfInfoList

`func (o *NefInfo) SetUnTrustAfInfoList(v []UnTrustAfInfo)`

SetUnTrustAfInfoList sets UnTrustAfInfoList field to given value.

### HasUnTrustAfInfoList

`func (o *NefInfo) HasUnTrustAfInfoList() bool`

HasUnTrustAfInfoList returns a boolean if a field has been set.

### GetUasNfFunctionalityInd

`func (o *NefInfo) GetUasNfFunctionalityInd() bool`

GetUasNfFunctionalityInd returns the UasNfFunctionalityInd field if non-nil, zero value otherwise.

### GetUasNfFunctionalityIndOk

`func (o *NefInfo) GetUasNfFunctionalityIndOk() (*bool, bool)`

GetUasNfFunctionalityIndOk returns a tuple with the UasNfFunctionalityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUasNfFunctionalityInd

`func (o *NefInfo) SetUasNfFunctionalityInd(v bool)`

SetUasNfFunctionalityInd sets UasNfFunctionalityInd field to given value.

### HasUasNfFunctionalityInd

`func (o *NefInfo) HasUasNfFunctionalityInd() bool`

HasUasNfFunctionalityInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


