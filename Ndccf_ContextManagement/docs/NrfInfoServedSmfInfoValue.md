# NrfInfoServedSmfInfoValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiSmfInfoList** | [**[]SnssaiSmfInfoItem**](SnssaiSmfInfoItem.md) |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**PgwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PgwIpAddrList** | Pointer to [**[]IpAddr**](IpAddr.md) |  | [optional] 
**AccessType** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**VsmfSupportInd** | Pointer to **bool** |  | [optional] 
**PgwFqdnList** | Pointer to **[]string** |  | [optional] 
**SmfOnboardingCapability** | Pointer to **bool** |  | [optional] [default to false]
**IsmfSupportInd** | Pointer to **bool** |  | [optional] 
**SmfUPRPCapability** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNrfInfoServedSmfInfoValue

`func NewNrfInfoServedSmfInfoValue(sNssaiSmfInfoList []SnssaiSmfInfoItem, ) *NrfInfoServedSmfInfoValue`

NewNrfInfoServedSmfInfoValue instantiates a new NrfInfoServedSmfInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedSmfInfoValueWithDefaults

`func NewNrfInfoServedSmfInfoValueWithDefaults() *NrfInfoServedSmfInfoValue`

NewNrfInfoServedSmfInfoValueWithDefaults instantiates a new NrfInfoServedSmfInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiSmfInfoList

`func (o *NrfInfoServedSmfInfoValue) GetSNssaiSmfInfoList() []SnssaiSmfInfoItem`

GetSNssaiSmfInfoList returns the SNssaiSmfInfoList field if non-nil, zero value otherwise.

### GetSNssaiSmfInfoListOk

`func (o *NrfInfoServedSmfInfoValue) GetSNssaiSmfInfoListOk() (*[]SnssaiSmfInfoItem, bool)`

GetSNssaiSmfInfoListOk returns a tuple with the SNssaiSmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiSmfInfoList

`func (o *NrfInfoServedSmfInfoValue) SetSNssaiSmfInfoList(v []SnssaiSmfInfoItem)`

SetSNssaiSmfInfoList sets SNssaiSmfInfoList field to given value.


### GetTaiList

`func (o *NrfInfoServedSmfInfoValue) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NrfInfoServedSmfInfoValue) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NrfInfoServedSmfInfoValue) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NrfInfoServedSmfInfoValue) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NrfInfoServedSmfInfoValue) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NrfInfoServedSmfInfoValue) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NrfInfoServedSmfInfoValue) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NrfInfoServedSmfInfoValue) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetPgwFqdn

`func (o *NrfInfoServedSmfInfoValue) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *NrfInfoServedSmfInfoValue) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *NrfInfoServedSmfInfoValue) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *NrfInfoServedSmfInfoValue) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.

### GetPgwIpAddrList

`func (o *NrfInfoServedSmfInfoValue) GetPgwIpAddrList() []IpAddr`

GetPgwIpAddrList returns the PgwIpAddrList field if non-nil, zero value otherwise.

### GetPgwIpAddrListOk

`func (o *NrfInfoServedSmfInfoValue) GetPgwIpAddrListOk() (*[]IpAddr, bool)`

GetPgwIpAddrListOk returns a tuple with the PgwIpAddrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwIpAddrList

`func (o *NrfInfoServedSmfInfoValue) SetPgwIpAddrList(v []IpAddr)`

SetPgwIpAddrList sets PgwIpAddrList field to given value.

### HasPgwIpAddrList

`func (o *NrfInfoServedSmfInfoValue) HasPgwIpAddrList() bool`

HasPgwIpAddrList returns a boolean if a field has been set.

### GetAccessType

`func (o *NrfInfoServedSmfInfoValue) GetAccessType() []AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *NrfInfoServedSmfInfoValue) GetAccessTypeOk() (*[]AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *NrfInfoServedSmfInfoValue) SetAccessType(v []AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *NrfInfoServedSmfInfoValue) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetPriority

`func (o *NrfInfoServedSmfInfoValue) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NrfInfoServedSmfInfoValue) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NrfInfoServedSmfInfoValue) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NrfInfoServedSmfInfoValue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetVsmfSupportInd

`func (o *NrfInfoServedSmfInfoValue) GetVsmfSupportInd() bool`

GetVsmfSupportInd returns the VsmfSupportInd field if non-nil, zero value otherwise.

### GetVsmfSupportIndOk

`func (o *NrfInfoServedSmfInfoValue) GetVsmfSupportIndOk() (*bool, bool)`

GetVsmfSupportIndOk returns a tuple with the VsmfSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfSupportInd

`func (o *NrfInfoServedSmfInfoValue) SetVsmfSupportInd(v bool)`

SetVsmfSupportInd sets VsmfSupportInd field to given value.

### HasVsmfSupportInd

`func (o *NrfInfoServedSmfInfoValue) HasVsmfSupportInd() bool`

HasVsmfSupportInd returns a boolean if a field has been set.

### GetPgwFqdnList

`func (o *NrfInfoServedSmfInfoValue) GetPgwFqdnList() []string`

GetPgwFqdnList returns the PgwFqdnList field if non-nil, zero value otherwise.

### GetPgwFqdnListOk

`func (o *NrfInfoServedSmfInfoValue) GetPgwFqdnListOk() (*[]string, bool)`

GetPgwFqdnListOk returns a tuple with the PgwFqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdnList

`func (o *NrfInfoServedSmfInfoValue) SetPgwFqdnList(v []string)`

SetPgwFqdnList sets PgwFqdnList field to given value.

### HasPgwFqdnList

`func (o *NrfInfoServedSmfInfoValue) HasPgwFqdnList() bool`

HasPgwFqdnList returns a boolean if a field has been set.

### GetSmfOnboardingCapability

`func (o *NrfInfoServedSmfInfoValue) GetSmfOnboardingCapability() bool`

GetSmfOnboardingCapability returns the SmfOnboardingCapability field if non-nil, zero value otherwise.

### GetSmfOnboardingCapabilityOk

`func (o *NrfInfoServedSmfInfoValue) GetSmfOnboardingCapabilityOk() (*bool, bool)`

GetSmfOnboardingCapabilityOk returns a tuple with the SmfOnboardingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfOnboardingCapability

`func (o *NrfInfoServedSmfInfoValue) SetSmfOnboardingCapability(v bool)`

SetSmfOnboardingCapability sets SmfOnboardingCapability field to given value.

### HasSmfOnboardingCapability

`func (o *NrfInfoServedSmfInfoValue) HasSmfOnboardingCapability() bool`

HasSmfOnboardingCapability returns a boolean if a field has been set.

### GetIsmfSupportInd

`func (o *NrfInfoServedSmfInfoValue) GetIsmfSupportInd() bool`

GetIsmfSupportInd returns the IsmfSupportInd field if non-nil, zero value otherwise.

### GetIsmfSupportIndOk

`func (o *NrfInfoServedSmfInfoValue) GetIsmfSupportIndOk() (*bool, bool)`

GetIsmfSupportIndOk returns a tuple with the IsmfSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfSupportInd

`func (o *NrfInfoServedSmfInfoValue) SetIsmfSupportInd(v bool)`

SetIsmfSupportInd sets IsmfSupportInd field to given value.

### HasIsmfSupportInd

`func (o *NrfInfoServedSmfInfoValue) HasIsmfSupportInd() bool`

HasIsmfSupportInd returns a boolean if a field has been set.

### GetSmfUPRPCapability

`func (o *NrfInfoServedSmfInfoValue) GetSmfUPRPCapability() bool`

GetSmfUPRPCapability returns the SmfUPRPCapability field if non-nil, zero value otherwise.

### GetSmfUPRPCapabilityOk

`func (o *NrfInfoServedSmfInfoValue) GetSmfUPRPCapabilityOk() (*bool, bool)`

GetSmfUPRPCapabilityOk returns a tuple with the SmfUPRPCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfUPRPCapability

`func (o *NrfInfoServedSmfInfoValue) SetSmfUPRPCapability(v bool)`

SetSmfUPRPCapability sets SmfUPRPCapability field to given value.

### HasSmfUPRPCapability

`func (o *NrfInfoServedSmfInfoValue) HasSmfUPRPCapability() bool`

HasSmfUPRPCapability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


