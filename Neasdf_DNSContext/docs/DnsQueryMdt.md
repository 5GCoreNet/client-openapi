# DnsQueryMdt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MdtId** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**SourceIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**SourceIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**FqdnPatternList** | Pointer to [**[]FqdnPatternMatchingRule**](FqdnPatternMatchingRule.md) |  | [optional] 

## Methods

### NewDnsQueryMdt

`func NewDnsQueryMdt(mdtId string, ) *DnsQueryMdt`

NewDnsQueryMdt instantiates a new DnsQueryMdt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsQueryMdtWithDefaults

`func NewDnsQueryMdtWithDefaults() *DnsQueryMdt`

NewDnsQueryMdtWithDefaults instantiates a new DnsQueryMdt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMdtId

`func (o *DnsQueryMdt) GetMdtId() string`

GetMdtId returns the MdtId field if non-nil, zero value otherwise.

### GetMdtIdOk

`func (o *DnsQueryMdt) GetMdtIdOk() (*string, bool)`

GetMdtIdOk returns a tuple with the MdtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtId

`func (o *DnsQueryMdt) SetMdtId(v string)`

SetMdtId sets MdtId field to given value.


### GetLabel

`func (o *DnsQueryMdt) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DnsQueryMdt) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DnsQueryMdt) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DnsQueryMdt) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSourceIpv4Addr

`func (o *DnsQueryMdt) GetSourceIpv4Addr() string`

GetSourceIpv4Addr returns the SourceIpv4Addr field if non-nil, zero value otherwise.

### GetSourceIpv4AddrOk

`func (o *DnsQueryMdt) GetSourceIpv4AddrOk() (*string, bool)`

GetSourceIpv4AddrOk returns a tuple with the SourceIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpv4Addr

`func (o *DnsQueryMdt) SetSourceIpv4Addr(v string)`

SetSourceIpv4Addr sets SourceIpv4Addr field to given value.

### HasSourceIpv4Addr

`func (o *DnsQueryMdt) HasSourceIpv4Addr() bool`

HasSourceIpv4Addr returns a boolean if a field has been set.

### GetSourceIpv6Prefix

`func (o *DnsQueryMdt) GetSourceIpv6Prefix() Ipv6Prefix`

GetSourceIpv6Prefix returns the SourceIpv6Prefix field if non-nil, zero value otherwise.

### GetSourceIpv6PrefixOk

`func (o *DnsQueryMdt) GetSourceIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetSourceIpv6PrefixOk returns a tuple with the SourceIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpv6Prefix

`func (o *DnsQueryMdt) SetSourceIpv6Prefix(v Ipv6Prefix)`

SetSourceIpv6Prefix sets SourceIpv6Prefix field to given value.

### HasSourceIpv6Prefix

`func (o *DnsQueryMdt) HasSourceIpv6Prefix() bool`

HasSourceIpv6Prefix returns a boolean if a field has been set.

### GetFqdnPatternList

`func (o *DnsQueryMdt) GetFqdnPatternList() []FqdnPatternMatchingRule`

GetFqdnPatternList returns the FqdnPatternList field if non-nil, zero value otherwise.

### GetFqdnPatternListOk

`func (o *DnsQueryMdt) GetFqdnPatternListOk() (*[]FqdnPatternMatchingRule, bool)`

GetFqdnPatternListOk returns a tuple with the FqdnPatternList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnPatternList

`func (o *DnsQueryMdt) SetFqdnPatternList(v []FqdnPatternMatchingRule)`

SetFqdnPatternList sets FqdnPatternList field to given value.

### HasFqdnPatternList

`func (o *DnsQueryMdt) HasFqdnPatternList() bool`

HasFqdnPatternList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


