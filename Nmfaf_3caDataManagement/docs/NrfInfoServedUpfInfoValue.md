# NrfInfoServedUpfInfoValue

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

### NewNrfInfoServedUpfInfoValue

`func NewNrfInfoServedUpfInfoValue(sNssaiUpfInfoList []SnssaiUpfInfoItem, ) *NrfInfoServedUpfInfoValue`

NewNrfInfoServedUpfInfoValue instantiates a new NrfInfoServedUpfInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedUpfInfoValueWithDefaults

`func NewNrfInfoServedUpfInfoValueWithDefaults() *NrfInfoServedUpfInfoValue`

NewNrfInfoServedUpfInfoValueWithDefaults instantiates a new NrfInfoServedUpfInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) GetSNssaiUpfInfoList() []SnssaiUpfInfoItem`

GetSNssaiUpfInfoList returns the SNssaiUpfInfoList field if non-nil, zero value otherwise.

### GetSNssaiUpfInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetSNssaiUpfInfoListOk() (*[]SnssaiUpfInfoItem, bool)`

GetSNssaiUpfInfoListOk returns a tuple with the SNssaiUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) SetSNssaiUpfInfoList(v []SnssaiUpfInfoItem)`

SetSNssaiUpfInfoList sets SNssaiUpfInfoList field to given value.


### GetSmfServingArea

`func (o *NrfInfoServedUpfInfoValue) GetSmfServingArea() []string`

GetSmfServingArea returns the SmfServingArea field if non-nil, zero value otherwise.

### GetSmfServingAreaOk

`func (o *NrfInfoServedUpfInfoValue) GetSmfServingAreaOk() (*[]string, bool)`

GetSmfServingAreaOk returns a tuple with the SmfServingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServingArea

`func (o *NrfInfoServedUpfInfoValue) SetSmfServingArea(v []string)`

SetSmfServingArea sets SmfServingArea field to given value.

### HasSmfServingArea

`func (o *NrfInfoServedUpfInfoValue) HasSmfServingArea() bool`

HasSmfServingArea returns a boolean if a field has been set.

### GetInterfaceUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem`

GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field if non-nil, zero value otherwise.

### GetInterfaceUpfInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetInterfaceUpfInfoListOk() (*[]InterfaceUpfInfoItem, bool)`

GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem)`

SetInterfaceUpfInfoList sets InterfaceUpfInfoList field to given value.

### HasInterfaceUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) HasInterfaceUpfInfoList() bool`

HasInterfaceUpfInfoList returns a boolean if a field has been set.

### GetIwkEpsInd

`func (o *NrfInfoServedUpfInfoValue) GetIwkEpsInd() bool`

GetIwkEpsInd returns the IwkEpsInd field if non-nil, zero value otherwise.

### GetIwkEpsIndOk

`func (o *NrfInfoServedUpfInfoValue) GetIwkEpsIndOk() (*bool, bool)`

GetIwkEpsIndOk returns a tuple with the IwkEpsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpsInd

`func (o *NrfInfoServedUpfInfoValue) SetIwkEpsInd(v bool)`

SetIwkEpsInd sets IwkEpsInd field to given value.

### HasIwkEpsInd

`func (o *NrfInfoServedUpfInfoValue) HasIwkEpsInd() bool`

HasIwkEpsInd returns a boolean if a field has been set.

### GetPduSessionTypes

`func (o *NrfInfoServedUpfInfoValue) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *NrfInfoServedUpfInfoValue) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *NrfInfoServedUpfInfoValue) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *NrfInfoServedUpfInfoValue) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetAtsssCapability

`func (o *NrfInfoServedUpfInfoValue) GetAtsssCapability() AtsssCapability`

GetAtsssCapability returns the AtsssCapability field if non-nil, zero value otherwise.

### GetAtsssCapabilityOk

`func (o *NrfInfoServedUpfInfoValue) GetAtsssCapabilityOk() (*AtsssCapability, bool)`

GetAtsssCapabilityOk returns a tuple with the AtsssCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssCapability

`func (o *NrfInfoServedUpfInfoValue) SetAtsssCapability(v AtsssCapability)`

SetAtsssCapability sets AtsssCapability field to given value.

### HasAtsssCapability

`func (o *NrfInfoServedUpfInfoValue) HasAtsssCapability() bool`

HasAtsssCapability returns a boolean if a field has been set.

### GetUeIpAddrInd

`func (o *NrfInfoServedUpfInfoValue) GetUeIpAddrInd() bool`

GetUeIpAddrInd returns the UeIpAddrInd field if non-nil, zero value otherwise.

### GetUeIpAddrIndOk

`func (o *NrfInfoServedUpfInfoValue) GetUeIpAddrIndOk() (*bool, bool)`

GetUeIpAddrIndOk returns a tuple with the UeIpAddrInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddrInd

`func (o *NrfInfoServedUpfInfoValue) SetUeIpAddrInd(v bool)`

SetUeIpAddrInd sets UeIpAddrInd field to given value.

### HasUeIpAddrInd

`func (o *NrfInfoServedUpfInfoValue) HasUeIpAddrInd() bool`

HasUeIpAddrInd returns a boolean if a field has been set.

### GetTaiList

`func (o *NrfInfoServedUpfInfoValue) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NrfInfoServedUpfInfoValue) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NrfInfoServedUpfInfoValue) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NrfInfoServedUpfInfoValue) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NrfInfoServedUpfInfoValue) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NrfInfoServedUpfInfoValue) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NrfInfoServedUpfInfoValue) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NrfInfoServedUpfInfoValue) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetWAgfInfo

`func (o *NrfInfoServedUpfInfoValue) GetWAgfInfo() WAgfInfo`

GetWAgfInfo returns the WAgfInfo field if non-nil, zero value otherwise.

### GetWAgfInfoOk

`func (o *NrfInfoServedUpfInfoValue) GetWAgfInfoOk() (*WAgfInfo, bool)`

GetWAgfInfoOk returns a tuple with the WAgfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWAgfInfo

`func (o *NrfInfoServedUpfInfoValue) SetWAgfInfo(v WAgfInfo)`

SetWAgfInfo sets WAgfInfo field to given value.

### HasWAgfInfo

`func (o *NrfInfoServedUpfInfoValue) HasWAgfInfo() bool`

HasWAgfInfo returns a boolean if a field has been set.

### GetTngfInfo

`func (o *NrfInfoServedUpfInfoValue) GetTngfInfo() TngfInfo`

GetTngfInfo returns the TngfInfo field if non-nil, zero value otherwise.

### GetTngfInfoOk

`func (o *NrfInfoServedUpfInfoValue) GetTngfInfoOk() (*TngfInfo, bool)`

GetTngfInfoOk returns a tuple with the TngfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfInfo

`func (o *NrfInfoServedUpfInfoValue) SetTngfInfo(v TngfInfo)`

SetTngfInfo sets TngfInfo field to given value.

### HasTngfInfo

`func (o *NrfInfoServedUpfInfoValue) HasTngfInfo() bool`

HasTngfInfo returns a boolean if a field has been set.

### GetTwifInfo

`func (o *NrfInfoServedUpfInfoValue) GetTwifInfo() TwifInfo`

GetTwifInfo returns the TwifInfo field if non-nil, zero value otherwise.

### GetTwifInfoOk

`func (o *NrfInfoServedUpfInfoValue) GetTwifInfoOk() (*TwifInfo, bool)`

GetTwifInfoOk returns a tuple with the TwifInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwifInfo

`func (o *NrfInfoServedUpfInfoValue) SetTwifInfo(v TwifInfo)`

SetTwifInfo sets TwifInfo field to given value.

### HasTwifInfo

`func (o *NrfInfoServedUpfInfoValue) HasTwifInfo() bool`

HasTwifInfo returns a boolean if a field has been set.

### GetPriority

`func (o *NrfInfoServedUpfInfoValue) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NrfInfoServedUpfInfoValue) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NrfInfoServedUpfInfoValue) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NrfInfoServedUpfInfoValue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRedundantGtpu

`func (o *NrfInfoServedUpfInfoValue) GetRedundantGtpu() bool`

GetRedundantGtpu returns the RedundantGtpu field if non-nil, zero value otherwise.

### GetRedundantGtpuOk

`func (o *NrfInfoServedUpfInfoValue) GetRedundantGtpuOk() (*bool, bool)`

GetRedundantGtpuOk returns a tuple with the RedundantGtpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantGtpu

`func (o *NrfInfoServedUpfInfoValue) SetRedundantGtpu(v bool)`

SetRedundantGtpu sets RedundantGtpu field to given value.

### HasRedundantGtpu

`func (o *NrfInfoServedUpfInfoValue) HasRedundantGtpu() bool`

HasRedundantGtpu returns a boolean if a field has been set.

### GetIpups

`func (o *NrfInfoServedUpfInfoValue) GetIpups() bool`

GetIpups returns the Ipups field if non-nil, zero value otherwise.

### GetIpupsOk

`func (o *NrfInfoServedUpfInfoValue) GetIpupsOk() (*bool, bool)`

GetIpupsOk returns a tuple with the Ipups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpups

`func (o *NrfInfoServedUpfInfoValue) SetIpups(v bool)`

SetIpups sets Ipups field to given value.

### HasIpups

`func (o *NrfInfoServedUpfInfoValue) HasIpups() bool`

HasIpups returns a boolean if a field has been set.

### GetDataForwarding

`func (o *NrfInfoServedUpfInfoValue) GetDataForwarding() bool`

GetDataForwarding returns the DataForwarding field if non-nil, zero value otherwise.

### GetDataForwardingOk

`func (o *NrfInfoServedUpfInfoValue) GetDataForwardingOk() (*bool, bool)`

GetDataForwardingOk returns a tuple with the DataForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwarding

`func (o *NrfInfoServedUpfInfoValue) SetDataForwarding(v bool)`

SetDataForwarding sets DataForwarding field to given value.

### HasDataForwarding

`func (o *NrfInfoServedUpfInfoValue) HasDataForwarding() bool`

HasDataForwarding returns a boolean if a field has been set.

### GetSupportedPfcpFeatures

`func (o *NrfInfoServedUpfInfoValue) GetSupportedPfcpFeatures() string`

GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field if non-nil, zero value otherwise.

### GetSupportedPfcpFeaturesOk

`func (o *NrfInfoServedUpfInfoValue) GetSupportedPfcpFeaturesOk() (*string, bool)`

GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPfcpFeatures

`func (o *NrfInfoServedUpfInfoValue) SetSupportedPfcpFeatures(v string)`

SetSupportedPfcpFeatures sets SupportedPfcpFeatures field to given value.

### HasSupportedPfcpFeatures

`func (o *NrfInfoServedUpfInfoValue) HasSupportedPfcpFeatures() bool`

HasSupportedPfcpFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


