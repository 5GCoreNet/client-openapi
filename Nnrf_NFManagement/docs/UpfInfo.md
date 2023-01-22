# UpfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiUpfInfoList** | [**[]SnssaiUpfInfoItem**](SnssaiUpfInfoItem.md) |  | 
**SmfServingArea** | Pointer to **[]string** |  | [optional] 
**InterfaceUpfInfoList** | Pointer to [**[]InterfaceUpfInfoItem**](InterfaceUpfInfoItem.md) |  | [optional] 
**IwkEpsInd** | Pointer to **bool** |  | [optional] [default to false]
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**AtsssCapability** | Pointer to [**AtsssCapability**](AtsssCapability.md) |  | [optional] 
**UeIpAddrInd** | Pointer to **bool** |  | [optional] [default to false]
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**WAgfInfo** | Pointer to [**WAgfInfo**](WAgfInfo.md) |  | [optional] 
**TngfInfo** | Pointer to [**TngfInfo**](TngfInfo.md) |  | [optional] 
**TwifInfo** | Pointer to [**TwifInfo**](TwifInfo.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RedundantGtpu** | Pointer to **bool** |  | [optional] [default to false]
**Ipups** | Pointer to **bool** |  | [optional] [default to false]
**DataForwarding** | Pointer to **bool** |  | [optional] [default to false]
**SupportedPfcpFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewUpfInfo

`func NewUpfInfo(sNssaiUpfInfoList []SnssaiUpfInfoItem, ) *UpfInfo`

NewUpfInfo instantiates a new UpfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpfInfoWithDefaults

`func NewUpfInfoWithDefaults() *UpfInfo`

NewUpfInfoWithDefaults instantiates a new UpfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiUpfInfoList

`func (o *UpfInfo) GetSNssaiUpfInfoList() []SnssaiUpfInfoItem`

GetSNssaiUpfInfoList returns the SNssaiUpfInfoList field if non-nil, zero value otherwise.

### GetSNssaiUpfInfoListOk

`func (o *UpfInfo) GetSNssaiUpfInfoListOk() (*[]SnssaiUpfInfoItem, bool)`

GetSNssaiUpfInfoListOk returns a tuple with the SNssaiUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiUpfInfoList

`func (o *UpfInfo) SetSNssaiUpfInfoList(v []SnssaiUpfInfoItem)`

SetSNssaiUpfInfoList sets SNssaiUpfInfoList field to given value.


### GetSmfServingArea

`func (o *UpfInfo) GetSmfServingArea() []string`

GetSmfServingArea returns the SmfServingArea field if non-nil, zero value otherwise.

### GetSmfServingAreaOk

`func (o *UpfInfo) GetSmfServingAreaOk() (*[]string, bool)`

GetSmfServingAreaOk returns a tuple with the SmfServingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServingArea

`func (o *UpfInfo) SetSmfServingArea(v []string)`

SetSmfServingArea sets SmfServingArea field to given value.

### HasSmfServingArea

`func (o *UpfInfo) HasSmfServingArea() bool`

HasSmfServingArea returns a boolean if a field has been set.

### GetInterfaceUpfInfoList

`func (o *UpfInfo) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem`

GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field if non-nil, zero value otherwise.

### GetInterfaceUpfInfoListOk

`func (o *UpfInfo) GetInterfaceUpfInfoListOk() (*[]InterfaceUpfInfoItem, bool)`

GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUpfInfoList

`func (o *UpfInfo) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem)`

SetInterfaceUpfInfoList sets InterfaceUpfInfoList field to given value.

### HasInterfaceUpfInfoList

`func (o *UpfInfo) HasInterfaceUpfInfoList() bool`

HasInterfaceUpfInfoList returns a boolean if a field has been set.

### GetIwkEpsInd

`func (o *UpfInfo) GetIwkEpsInd() bool`

GetIwkEpsInd returns the IwkEpsInd field if non-nil, zero value otherwise.

### GetIwkEpsIndOk

`func (o *UpfInfo) GetIwkEpsIndOk() (*bool, bool)`

GetIwkEpsIndOk returns a tuple with the IwkEpsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpsInd

`func (o *UpfInfo) SetIwkEpsInd(v bool)`

SetIwkEpsInd sets IwkEpsInd field to given value.

### HasIwkEpsInd

`func (o *UpfInfo) HasIwkEpsInd() bool`

HasIwkEpsInd returns a boolean if a field has been set.

### GetPduSessionTypes

`func (o *UpfInfo) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *UpfInfo) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *UpfInfo) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *UpfInfo) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetAtsssCapability

`func (o *UpfInfo) GetAtsssCapability() AtsssCapability`

GetAtsssCapability returns the AtsssCapability field if non-nil, zero value otherwise.

### GetAtsssCapabilityOk

`func (o *UpfInfo) GetAtsssCapabilityOk() (*AtsssCapability, bool)`

GetAtsssCapabilityOk returns a tuple with the AtsssCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssCapability

`func (o *UpfInfo) SetAtsssCapability(v AtsssCapability)`

SetAtsssCapability sets AtsssCapability field to given value.

### HasAtsssCapability

`func (o *UpfInfo) HasAtsssCapability() bool`

HasAtsssCapability returns a boolean if a field has been set.

### GetUeIpAddrInd

`func (o *UpfInfo) GetUeIpAddrInd() bool`

GetUeIpAddrInd returns the UeIpAddrInd field if non-nil, zero value otherwise.

### GetUeIpAddrIndOk

`func (o *UpfInfo) GetUeIpAddrIndOk() (*bool, bool)`

GetUeIpAddrIndOk returns a tuple with the UeIpAddrInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddrInd

`func (o *UpfInfo) SetUeIpAddrInd(v bool)`

SetUeIpAddrInd sets UeIpAddrInd field to given value.

### HasUeIpAddrInd

`func (o *UpfInfo) HasUeIpAddrInd() bool`

HasUeIpAddrInd returns a boolean if a field has been set.

### GetTaiList

`func (o *UpfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *UpfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *UpfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *UpfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *UpfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *UpfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *UpfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *UpfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetWAgfInfo

`func (o *UpfInfo) GetWAgfInfo() WAgfInfo`

GetWAgfInfo returns the WAgfInfo field if non-nil, zero value otherwise.

### GetWAgfInfoOk

`func (o *UpfInfo) GetWAgfInfoOk() (*WAgfInfo, bool)`

GetWAgfInfoOk returns a tuple with the WAgfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWAgfInfo

`func (o *UpfInfo) SetWAgfInfo(v WAgfInfo)`

SetWAgfInfo sets WAgfInfo field to given value.

### HasWAgfInfo

`func (o *UpfInfo) HasWAgfInfo() bool`

HasWAgfInfo returns a boolean if a field has been set.

### GetTngfInfo

`func (o *UpfInfo) GetTngfInfo() TngfInfo`

GetTngfInfo returns the TngfInfo field if non-nil, zero value otherwise.

### GetTngfInfoOk

`func (o *UpfInfo) GetTngfInfoOk() (*TngfInfo, bool)`

GetTngfInfoOk returns a tuple with the TngfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfInfo

`func (o *UpfInfo) SetTngfInfo(v TngfInfo)`

SetTngfInfo sets TngfInfo field to given value.

### HasTngfInfo

`func (o *UpfInfo) HasTngfInfo() bool`

HasTngfInfo returns a boolean if a field has been set.

### GetTwifInfo

`func (o *UpfInfo) GetTwifInfo() TwifInfo`

GetTwifInfo returns the TwifInfo field if non-nil, zero value otherwise.

### GetTwifInfoOk

`func (o *UpfInfo) GetTwifInfoOk() (*TwifInfo, bool)`

GetTwifInfoOk returns a tuple with the TwifInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwifInfo

`func (o *UpfInfo) SetTwifInfo(v TwifInfo)`

SetTwifInfo sets TwifInfo field to given value.

### HasTwifInfo

`func (o *UpfInfo) HasTwifInfo() bool`

HasTwifInfo returns a boolean if a field has been set.

### GetPriority

`func (o *UpfInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpfInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpfInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpfInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRedundantGtpu

`func (o *UpfInfo) GetRedundantGtpu() bool`

GetRedundantGtpu returns the RedundantGtpu field if non-nil, zero value otherwise.

### GetRedundantGtpuOk

`func (o *UpfInfo) GetRedundantGtpuOk() (*bool, bool)`

GetRedundantGtpuOk returns a tuple with the RedundantGtpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantGtpu

`func (o *UpfInfo) SetRedundantGtpu(v bool)`

SetRedundantGtpu sets RedundantGtpu field to given value.

### HasRedundantGtpu

`func (o *UpfInfo) HasRedundantGtpu() bool`

HasRedundantGtpu returns a boolean if a field has been set.

### GetIpups

`func (o *UpfInfo) GetIpups() bool`

GetIpups returns the Ipups field if non-nil, zero value otherwise.

### GetIpupsOk

`func (o *UpfInfo) GetIpupsOk() (*bool, bool)`

GetIpupsOk returns a tuple with the Ipups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpups

`func (o *UpfInfo) SetIpups(v bool)`

SetIpups sets Ipups field to given value.

### HasIpups

`func (o *UpfInfo) HasIpups() bool`

HasIpups returns a boolean if a field has been set.

### GetDataForwarding

`func (o *UpfInfo) GetDataForwarding() bool`

GetDataForwarding returns the DataForwarding field if non-nil, zero value otherwise.

### GetDataForwardingOk

`func (o *UpfInfo) GetDataForwardingOk() (*bool, bool)`

GetDataForwardingOk returns a tuple with the DataForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwarding

`func (o *UpfInfo) SetDataForwarding(v bool)`

SetDataForwarding sets DataForwarding field to given value.

### HasDataForwarding

`func (o *UpfInfo) HasDataForwarding() bool`

HasDataForwarding returns a boolean if a field has been set.

### GetSupportedPfcpFeatures

`func (o *UpfInfo) GetSupportedPfcpFeatures() string`

GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field if non-nil, zero value otherwise.

### GetSupportedPfcpFeaturesOk

`func (o *UpfInfo) GetSupportedPfcpFeaturesOk() (*string, bool)`

GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPfcpFeatures

`func (o *UpfInfo) SetSupportedPfcpFeatures(v string)`

SetSupportedPfcpFeatures sets SupportedPfcpFeatures field to given value.

### HasSupportedPfcpFeatures

`func (o *UpfInfo) HasSupportedPfcpFeatures() bool`

HasSupportedPfcpFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


