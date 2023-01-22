# SmfInfo

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

### NewSmfInfo

`func NewSmfInfo(sNssaiSmfInfoList []SnssaiSmfInfoItem, ) *SmfInfo`

NewSmfInfo instantiates a new SmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfInfoWithDefaults

`func NewSmfInfoWithDefaults() *SmfInfo`

NewSmfInfoWithDefaults instantiates a new SmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiSmfInfoList

`func (o *SmfInfo) GetSNssaiSmfInfoList() []SnssaiSmfInfoItem`

GetSNssaiSmfInfoList returns the SNssaiSmfInfoList field if non-nil, zero value otherwise.

### GetSNssaiSmfInfoListOk

`func (o *SmfInfo) GetSNssaiSmfInfoListOk() (*[]SnssaiSmfInfoItem, bool)`

GetSNssaiSmfInfoListOk returns a tuple with the SNssaiSmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiSmfInfoList

`func (o *SmfInfo) SetSNssaiSmfInfoList(v []SnssaiSmfInfoItem)`

SetSNssaiSmfInfoList sets SNssaiSmfInfoList field to given value.


### GetTaiList

`func (o *SmfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *SmfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *SmfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *SmfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *SmfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *SmfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *SmfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *SmfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetPgwFqdn

`func (o *SmfInfo) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *SmfInfo) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *SmfInfo) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *SmfInfo) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.

### GetPgwIpAddrList

`func (o *SmfInfo) GetPgwIpAddrList() []IpAddr`

GetPgwIpAddrList returns the PgwIpAddrList field if non-nil, zero value otherwise.

### GetPgwIpAddrListOk

`func (o *SmfInfo) GetPgwIpAddrListOk() (*[]IpAddr, bool)`

GetPgwIpAddrListOk returns a tuple with the PgwIpAddrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwIpAddrList

`func (o *SmfInfo) SetPgwIpAddrList(v []IpAddr)`

SetPgwIpAddrList sets PgwIpAddrList field to given value.

### HasPgwIpAddrList

`func (o *SmfInfo) HasPgwIpAddrList() bool`

HasPgwIpAddrList returns a boolean if a field has been set.

### GetAccessType

`func (o *SmfInfo) GetAccessType() []AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SmfInfo) GetAccessTypeOk() (*[]AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SmfInfo) SetAccessType(v []AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SmfInfo) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetPriority

`func (o *SmfInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SmfInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SmfInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SmfInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetVsmfSupportInd

`func (o *SmfInfo) GetVsmfSupportInd() bool`

GetVsmfSupportInd returns the VsmfSupportInd field if non-nil, zero value otherwise.

### GetVsmfSupportIndOk

`func (o *SmfInfo) GetVsmfSupportIndOk() (*bool, bool)`

GetVsmfSupportIndOk returns a tuple with the VsmfSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfSupportInd

`func (o *SmfInfo) SetVsmfSupportInd(v bool)`

SetVsmfSupportInd sets VsmfSupportInd field to given value.

### HasVsmfSupportInd

`func (o *SmfInfo) HasVsmfSupportInd() bool`

HasVsmfSupportInd returns a boolean if a field has been set.

### GetPgwFqdnList

`func (o *SmfInfo) GetPgwFqdnList() []string`

GetPgwFqdnList returns the PgwFqdnList field if non-nil, zero value otherwise.

### GetPgwFqdnListOk

`func (o *SmfInfo) GetPgwFqdnListOk() (*[]string, bool)`

GetPgwFqdnListOk returns a tuple with the PgwFqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdnList

`func (o *SmfInfo) SetPgwFqdnList(v []string)`

SetPgwFqdnList sets PgwFqdnList field to given value.

### HasPgwFqdnList

`func (o *SmfInfo) HasPgwFqdnList() bool`

HasPgwFqdnList returns a boolean if a field has been set.

### GetSmfOnboardingCapability

`func (o *SmfInfo) GetSmfOnboardingCapability() bool`

GetSmfOnboardingCapability returns the SmfOnboardingCapability field if non-nil, zero value otherwise.

### GetSmfOnboardingCapabilityOk

`func (o *SmfInfo) GetSmfOnboardingCapabilityOk() (*bool, bool)`

GetSmfOnboardingCapabilityOk returns a tuple with the SmfOnboardingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfOnboardingCapability

`func (o *SmfInfo) SetSmfOnboardingCapability(v bool)`

SetSmfOnboardingCapability sets SmfOnboardingCapability field to given value.

### HasSmfOnboardingCapability

`func (o *SmfInfo) HasSmfOnboardingCapability() bool`

HasSmfOnboardingCapability returns a boolean if a field has been set.

### GetIsmfSupportInd

`func (o *SmfInfo) GetIsmfSupportInd() bool`

GetIsmfSupportInd returns the IsmfSupportInd field if non-nil, zero value otherwise.

### GetIsmfSupportIndOk

`func (o *SmfInfo) GetIsmfSupportIndOk() (*bool, bool)`

GetIsmfSupportIndOk returns a tuple with the IsmfSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfSupportInd

`func (o *SmfInfo) SetIsmfSupportInd(v bool)`

SetIsmfSupportInd sets IsmfSupportInd field to given value.

### HasIsmfSupportInd

`func (o *SmfInfo) HasIsmfSupportInd() bool`

HasIsmfSupportInd returns a boolean if a field has been set.

### GetSmfUPRPCapability

`func (o *SmfInfo) GetSmfUPRPCapability() bool`

GetSmfUPRPCapability returns the SmfUPRPCapability field if non-nil, zero value otherwise.

### GetSmfUPRPCapabilityOk

`func (o *SmfInfo) GetSmfUPRPCapabilityOk() (*bool, bool)`

GetSmfUPRPCapabilityOk returns a tuple with the SmfUPRPCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfUPRPCapability

`func (o *SmfInfo) SetSmfUPRPCapability(v bool)`

SetSmfUPRPCapability sets SmfUPRPCapability field to given value.

### HasSmfUPRPCapability

`func (o *SmfInfo) HasSmfUPRPCapability() bool`

HasSmfUPRPCapability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


