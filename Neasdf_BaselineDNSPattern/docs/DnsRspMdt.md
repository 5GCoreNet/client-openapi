# DnsRspMdt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MdtId** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**FqdnPatternList** | Pointer to [**[]FqdnPatternMatchingRule**](FqdnPatternMatchingRule.md) |  | [optional] 
**EasIpv4AddrRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**EasIpv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 

## Methods

### NewDnsRspMdt

`func NewDnsRspMdt(mdtId string, ) *DnsRspMdt`

NewDnsRspMdt instantiates a new DnsRspMdt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRspMdtWithDefaults

`func NewDnsRspMdtWithDefaults() *DnsRspMdt`

NewDnsRspMdtWithDefaults instantiates a new DnsRspMdt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMdtId

`func (o *DnsRspMdt) GetMdtId() string`

GetMdtId returns the MdtId field if non-nil, zero value otherwise.

### GetMdtIdOk

`func (o *DnsRspMdt) GetMdtIdOk() (*string, bool)`

GetMdtIdOk returns a tuple with the MdtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtId

`func (o *DnsRspMdt) SetMdtId(v string)`

SetMdtId sets MdtId field to given value.


### GetLabel

`func (o *DnsRspMdt) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DnsRspMdt) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DnsRspMdt) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DnsRspMdt) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetFqdnPatternList

`func (o *DnsRspMdt) GetFqdnPatternList() []FqdnPatternMatchingRule`

GetFqdnPatternList returns the FqdnPatternList field if non-nil, zero value otherwise.

### GetFqdnPatternListOk

`func (o *DnsRspMdt) GetFqdnPatternListOk() (*[]FqdnPatternMatchingRule, bool)`

GetFqdnPatternListOk returns a tuple with the FqdnPatternList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnPatternList

`func (o *DnsRspMdt) SetFqdnPatternList(v []FqdnPatternMatchingRule)`

SetFqdnPatternList sets FqdnPatternList field to given value.

### HasFqdnPatternList

`func (o *DnsRspMdt) HasFqdnPatternList() bool`

HasFqdnPatternList returns a boolean if a field has been set.

### GetEasIpv4AddrRanges

`func (o *DnsRspMdt) GetEasIpv4AddrRanges() []Ipv4AddressRange`

GetEasIpv4AddrRanges returns the EasIpv4AddrRanges field if non-nil, zero value otherwise.

### GetEasIpv4AddrRangesOk

`func (o *DnsRspMdt) GetEasIpv4AddrRangesOk() (*[]Ipv4AddressRange, bool)`

GetEasIpv4AddrRangesOk returns a tuple with the EasIpv4AddrRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpv4AddrRanges

`func (o *DnsRspMdt) SetEasIpv4AddrRanges(v []Ipv4AddressRange)`

SetEasIpv4AddrRanges sets EasIpv4AddrRanges field to given value.

### HasEasIpv4AddrRanges

`func (o *DnsRspMdt) HasEasIpv4AddrRanges() bool`

HasEasIpv4AddrRanges returns a boolean if a field has been set.

### GetEasIpv6PrefixRanges

`func (o *DnsRspMdt) GetEasIpv6PrefixRanges() []Ipv6PrefixRange`

GetEasIpv6PrefixRanges returns the EasIpv6PrefixRanges field if non-nil, zero value otherwise.

### GetEasIpv6PrefixRangesOk

`func (o *DnsRspMdt) GetEasIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetEasIpv6PrefixRangesOk returns a tuple with the EasIpv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpv6PrefixRanges

`func (o *DnsRspMdt) SetEasIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetEasIpv6PrefixRanges sets EasIpv6PrefixRanges field to given value.

### HasEasIpv6PrefixRanges

`func (o *DnsRspMdt) HasEasIpv6PrefixRanges() bool`

HasEasIpv6PrefixRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


