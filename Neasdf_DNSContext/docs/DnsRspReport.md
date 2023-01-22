# DnsRspReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**EasIpv4Addresses** | Pointer to **[]string** |  | [optional] 
**EasIpv6Addresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**EcsOption** | Pointer to [**EcsOption**](EcsOption.md) |  | [optional] 

## Methods

### NewDnsRspReport

`func NewDnsRspReport() *DnsRspReport`

NewDnsRspReport instantiates a new DnsRspReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRspReportWithDefaults

`func NewDnsRspReportWithDefaults() *DnsRspReport`

NewDnsRspReportWithDefaults instantiates a new DnsRspReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *DnsRspReport) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DnsRspReport) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DnsRspReport) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DnsRspReport) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetEasIpv4Addresses

`func (o *DnsRspReport) GetEasIpv4Addresses() []string`

GetEasIpv4Addresses returns the EasIpv4Addresses field if non-nil, zero value otherwise.

### GetEasIpv4AddressesOk

`func (o *DnsRspReport) GetEasIpv4AddressesOk() (*[]string, bool)`

GetEasIpv4AddressesOk returns a tuple with the EasIpv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpv4Addresses

`func (o *DnsRspReport) SetEasIpv4Addresses(v []string)`

SetEasIpv4Addresses sets EasIpv4Addresses field to given value.

### HasEasIpv4Addresses

`func (o *DnsRspReport) HasEasIpv4Addresses() bool`

HasEasIpv4Addresses returns a boolean if a field has been set.

### GetEasIpv6Addresses

`func (o *DnsRspReport) GetEasIpv6Addresses() []Ipv6Addr`

GetEasIpv6Addresses returns the EasIpv6Addresses field if non-nil, zero value otherwise.

### GetEasIpv6AddressesOk

`func (o *DnsRspReport) GetEasIpv6AddressesOk() (*[]Ipv6Addr, bool)`

GetEasIpv6AddressesOk returns a tuple with the EasIpv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpv6Addresses

`func (o *DnsRspReport) SetEasIpv6Addresses(v []Ipv6Addr)`

SetEasIpv6Addresses sets EasIpv6Addresses field to given value.

### HasEasIpv6Addresses

`func (o *DnsRspReport) HasEasIpv6Addresses() bool`

HasEasIpv6Addresses returns a boolean if a field has been set.

### GetEcsOption

`func (o *DnsRspReport) GetEcsOption() EcsOption`

GetEcsOption returns the EcsOption field if non-nil, zero value otherwise.

### GetEcsOptionOk

`func (o *DnsRspReport) GetEcsOptionOk() (*EcsOption, bool)`

GetEcsOptionOk returns a tuple with the EcsOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsOption

`func (o *DnsRspReport) SetEcsOption(v EcsOption)`

SetEcsOption sets EcsOption field to given value.

### HasEcsOption

`func (o *DnsRspReport) HasEcsOption() bool`

HasEcsOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


