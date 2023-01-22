# NrfInfoServedNefInfoValue

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

### NewNrfInfoServedNefInfoValue

`func NewNrfInfoServedNefInfoValue() *NrfInfoServedNefInfoValue`

NewNrfInfoServedNefInfoValue instantiates a new NrfInfoServedNefInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedNefInfoValueWithDefaults

`func NewNrfInfoServedNefInfoValueWithDefaults() *NrfInfoServedNefInfoValue`

NewNrfInfoServedNefInfoValueWithDefaults instantiates a new NrfInfoServedNefInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNefId

`func (o *NrfInfoServedNefInfoValue) GetNefId() string`

GetNefId returns the NefId field if non-nil, zero value otherwise.

### GetNefIdOk

`func (o *NrfInfoServedNefInfoValue) GetNefIdOk() (*string, bool)`

GetNefIdOk returns a tuple with the NefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefId

`func (o *NrfInfoServedNefInfoValue) SetNefId(v string)`

SetNefId sets NefId field to given value.

### HasNefId

`func (o *NrfInfoServedNefInfoValue) HasNefId() bool`

HasNefId returns a boolean if a field has been set.

### GetPfdData

`func (o *NrfInfoServedNefInfoValue) GetPfdData() PfdData`

GetPfdData returns the PfdData field if non-nil, zero value otherwise.

### GetPfdDataOk

`func (o *NrfInfoServedNefInfoValue) GetPfdDataOk() (*PfdData, bool)`

GetPfdDataOk returns a tuple with the PfdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdData

`func (o *NrfInfoServedNefInfoValue) SetPfdData(v PfdData)`

SetPfdData sets PfdData field to given value.

### HasPfdData

`func (o *NrfInfoServedNefInfoValue) HasPfdData() bool`

HasPfdData returns a boolean if a field has been set.

### GetAfEeData

`func (o *NrfInfoServedNefInfoValue) GetAfEeData() AfEventExposureData`

GetAfEeData returns the AfEeData field if non-nil, zero value otherwise.

### GetAfEeDataOk

`func (o *NrfInfoServedNefInfoValue) GetAfEeDataOk() (*AfEventExposureData, bool)`

GetAfEeDataOk returns a tuple with the AfEeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEeData

`func (o *NrfInfoServedNefInfoValue) SetAfEeData(v AfEventExposureData)`

SetAfEeData sets AfEeData field to given value.

### HasAfEeData

`func (o *NrfInfoServedNefInfoValue) HasAfEeData() bool`

HasAfEeData returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *NrfInfoServedNefInfoValue) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *NrfInfoServedNefInfoValue) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *NrfInfoServedNefInfoValue) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *NrfInfoServedNefInfoValue) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedNefInfoValue) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *NrfInfoServedNefInfoValue) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedNefInfoValue) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *NrfInfoServedNefInfoValue) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetServedFqdnList

`func (o *NrfInfoServedNefInfoValue) GetServedFqdnList() []string`

GetServedFqdnList returns the ServedFqdnList field if non-nil, zero value otherwise.

### GetServedFqdnListOk

`func (o *NrfInfoServedNefInfoValue) GetServedFqdnListOk() (*[]string, bool)`

GetServedFqdnListOk returns a tuple with the ServedFqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedFqdnList

`func (o *NrfInfoServedNefInfoValue) SetServedFqdnList(v []string)`

SetServedFqdnList sets ServedFqdnList field to given value.

### HasServedFqdnList

`func (o *NrfInfoServedNefInfoValue) HasServedFqdnList() bool`

HasServedFqdnList returns a boolean if a field has been set.

### GetTaiList

`func (o *NrfInfoServedNefInfoValue) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NrfInfoServedNefInfoValue) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NrfInfoServedNefInfoValue) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NrfInfoServedNefInfoValue) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NrfInfoServedNefInfoValue) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NrfInfoServedNefInfoValue) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NrfInfoServedNefInfoValue) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NrfInfoServedNefInfoValue) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetDnaiList

`func (o *NrfInfoServedNefInfoValue) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *NrfInfoServedNefInfoValue) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *NrfInfoServedNefInfoValue) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *NrfInfoServedNefInfoValue) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetUnTrustAfInfoList

`func (o *NrfInfoServedNefInfoValue) GetUnTrustAfInfoList() []UnTrustAfInfo`

GetUnTrustAfInfoList returns the UnTrustAfInfoList field if non-nil, zero value otherwise.

### GetUnTrustAfInfoListOk

`func (o *NrfInfoServedNefInfoValue) GetUnTrustAfInfoListOk() (*[]UnTrustAfInfo, bool)`

GetUnTrustAfInfoListOk returns a tuple with the UnTrustAfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnTrustAfInfoList

`func (o *NrfInfoServedNefInfoValue) SetUnTrustAfInfoList(v []UnTrustAfInfo)`

SetUnTrustAfInfoList sets UnTrustAfInfoList field to given value.

### HasUnTrustAfInfoList

`func (o *NrfInfoServedNefInfoValue) HasUnTrustAfInfoList() bool`

HasUnTrustAfInfoList returns a boolean if a field has been set.

### GetUasNfFunctionalityInd

`func (o *NrfInfoServedNefInfoValue) GetUasNfFunctionalityInd() bool`

GetUasNfFunctionalityInd returns the UasNfFunctionalityInd field if non-nil, zero value otherwise.

### GetUasNfFunctionalityIndOk

`func (o *NrfInfoServedNefInfoValue) GetUasNfFunctionalityIndOk() (*bool, bool)`

GetUasNfFunctionalityIndOk returns a tuple with the UasNfFunctionalityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUasNfFunctionalityInd

`func (o *NrfInfoServedNefInfoValue) SetUasNfFunctionalityInd(v bool)`

SetUasNfFunctionalityInd sets UasNfFunctionalityInd field to given value.

### HasUasNfFunctionalityInd

`func (o *NrfInfoServedNefInfoValue) HasUasNfFunctionalityInd() bool`

HasUasNfFunctionalityInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


