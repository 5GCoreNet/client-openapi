# NrfInfoServedScpInfoListValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScpDomainInfoList** | Pointer to [**map[string]ScpDomainInfo**](ScpDomainInfo.md) | A map (list of key-value pairs) where the key of the map shall be the string identifying an SCP domain  | [optional] 
**ScpPrefix** | Pointer to **string** |  | [optional] 
**ScpPorts** | Pointer to **map[string]int32** | Port numbers for HTTP and HTTPS. The key of the map shall be \&quot;http\&quot; or \&quot;https\&quot;.  | [optional] 
**AddressDomains** | Pointer to **[]string** |  | [optional] 
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Ipv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**Ipv4AddrRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**Ipv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 
**ServedNfSetIdList** | Pointer to **[]string** |  | [optional] 
**RemotePlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**RemoteSnpnList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**IpReachability** | Pointer to [**IpReachability**](IpReachability.md) |  | [optional] 
**ScpCapabilities** | Pointer to [**[]ScpCapability**](ScpCapability.md) |  | [optional] 

## Methods

### NewNrfInfoServedScpInfoListValue

`func NewNrfInfoServedScpInfoListValue() *NrfInfoServedScpInfoListValue`

NewNrfInfoServedScpInfoListValue instantiates a new NrfInfoServedScpInfoListValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedScpInfoListValueWithDefaults

`func NewNrfInfoServedScpInfoListValueWithDefaults() *NrfInfoServedScpInfoListValue`

NewNrfInfoServedScpInfoListValueWithDefaults instantiates a new NrfInfoServedScpInfoListValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScpDomainInfoList

`func (o *NrfInfoServedScpInfoListValue) GetScpDomainInfoList() map[string]ScpDomainInfo`

GetScpDomainInfoList returns the ScpDomainInfoList field if non-nil, zero value otherwise.

### GetScpDomainInfoListOk

`func (o *NrfInfoServedScpInfoListValue) GetScpDomainInfoListOk() (*map[string]ScpDomainInfo, bool)`

GetScpDomainInfoListOk returns a tuple with the ScpDomainInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpDomainInfoList

`func (o *NrfInfoServedScpInfoListValue) SetScpDomainInfoList(v map[string]ScpDomainInfo)`

SetScpDomainInfoList sets ScpDomainInfoList field to given value.

### HasScpDomainInfoList

`func (o *NrfInfoServedScpInfoListValue) HasScpDomainInfoList() bool`

HasScpDomainInfoList returns a boolean if a field has been set.

### GetScpPrefix

`func (o *NrfInfoServedScpInfoListValue) GetScpPrefix() string`

GetScpPrefix returns the ScpPrefix field if non-nil, zero value otherwise.

### GetScpPrefixOk

`func (o *NrfInfoServedScpInfoListValue) GetScpPrefixOk() (*string, bool)`

GetScpPrefixOk returns a tuple with the ScpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpPrefix

`func (o *NrfInfoServedScpInfoListValue) SetScpPrefix(v string)`

SetScpPrefix sets ScpPrefix field to given value.

### HasScpPrefix

`func (o *NrfInfoServedScpInfoListValue) HasScpPrefix() bool`

HasScpPrefix returns a boolean if a field has been set.

### GetScpPorts

`func (o *NrfInfoServedScpInfoListValue) GetScpPorts() map[string]int32`

GetScpPorts returns the ScpPorts field if non-nil, zero value otherwise.

### GetScpPortsOk

`func (o *NrfInfoServedScpInfoListValue) GetScpPortsOk() (*map[string]int32, bool)`

GetScpPortsOk returns a tuple with the ScpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpPorts

`func (o *NrfInfoServedScpInfoListValue) SetScpPorts(v map[string]int32)`

SetScpPorts sets ScpPorts field to given value.

### HasScpPorts

`func (o *NrfInfoServedScpInfoListValue) HasScpPorts() bool`

HasScpPorts returns a boolean if a field has been set.

### GetAddressDomains

`func (o *NrfInfoServedScpInfoListValue) GetAddressDomains() []string`

GetAddressDomains returns the AddressDomains field if non-nil, zero value otherwise.

### GetAddressDomainsOk

`func (o *NrfInfoServedScpInfoListValue) GetAddressDomainsOk() (*[]string, bool)`

GetAddressDomainsOk returns a tuple with the AddressDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressDomains

`func (o *NrfInfoServedScpInfoListValue) SetAddressDomains(v []string)`

SetAddressDomains sets AddressDomains field to given value.

### HasAddressDomains

`func (o *NrfInfoServedScpInfoListValue) HasAddressDomains() bool`

HasAddressDomains returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *NrfInfoServedScpInfoListValue) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *NrfInfoServedScpInfoListValue) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *NrfInfoServedScpInfoListValue) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *NrfInfoServedScpInfoListValue) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Prefixes

`func (o *NrfInfoServedScpInfoListValue) GetIpv6Prefixes() []Ipv6Prefix`

GetIpv6Prefixes returns the Ipv6Prefixes field if non-nil, zero value otherwise.

### GetIpv6PrefixesOk

`func (o *NrfInfoServedScpInfoListValue) GetIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefixes

`func (o *NrfInfoServedScpInfoListValue) SetIpv6Prefixes(v []Ipv6Prefix)`

SetIpv6Prefixes sets Ipv6Prefixes field to given value.

### HasIpv6Prefixes

`func (o *NrfInfoServedScpInfoListValue) HasIpv6Prefixes() bool`

HasIpv6Prefixes returns a boolean if a field has been set.

### GetIpv4AddrRanges

`func (o *NrfInfoServedScpInfoListValue) GetIpv4AddrRanges() []Ipv4AddressRange`

GetIpv4AddrRanges returns the Ipv4AddrRanges field if non-nil, zero value otherwise.

### GetIpv4AddrRangesOk

`func (o *NrfInfoServedScpInfoListValue) GetIpv4AddrRangesOk() (*[]Ipv4AddressRange, bool)`

GetIpv4AddrRangesOk returns a tuple with the Ipv4AddrRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddrRanges

`func (o *NrfInfoServedScpInfoListValue) SetIpv4AddrRanges(v []Ipv4AddressRange)`

SetIpv4AddrRanges sets Ipv4AddrRanges field to given value.

### HasIpv4AddrRanges

`func (o *NrfInfoServedScpInfoListValue) HasIpv4AddrRanges() bool`

HasIpv4AddrRanges returns a boolean if a field has been set.

### GetIpv6PrefixRanges

`func (o *NrfInfoServedScpInfoListValue) GetIpv6PrefixRanges() []Ipv6PrefixRange`

GetIpv6PrefixRanges returns the Ipv6PrefixRanges field if non-nil, zero value otherwise.

### GetIpv6PrefixRangesOk

`func (o *NrfInfoServedScpInfoListValue) GetIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixRanges

`func (o *NrfInfoServedScpInfoListValue) SetIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetIpv6PrefixRanges sets Ipv6PrefixRanges field to given value.

### HasIpv6PrefixRanges

`func (o *NrfInfoServedScpInfoListValue) HasIpv6PrefixRanges() bool`

HasIpv6PrefixRanges returns a boolean if a field has been set.

### GetServedNfSetIdList

`func (o *NrfInfoServedScpInfoListValue) GetServedNfSetIdList() []string`

GetServedNfSetIdList returns the ServedNfSetIdList field if non-nil, zero value otherwise.

### GetServedNfSetIdListOk

`func (o *NrfInfoServedScpInfoListValue) GetServedNfSetIdListOk() (*[]string, bool)`

GetServedNfSetIdListOk returns a tuple with the ServedNfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedNfSetIdList

`func (o *NrfInfoServedScpInfoListValue) SetServedNfSetIdList(v []string)`

SetServedNfSetIdList sets ServedNfSetIdList field to given value.

### HasServedNfSetIdList

`func (o *NrfInfoServedScpInfoListValue) HasServedNfSetIdList() bool`

HasServedNfSetIdList returns a boolean if a field has been set.

### GetRemotePlmnList

`func (o *NrfInfoServedScpInfoListValue) GetRemotePlmnList() []PlmnId`

GetRemotePlmnList returns the RemotePlmnList field if non-nil, zero value otherwise.

### GetRemotePlmnListOk

`func (o *NrfInfoServedScpInfoListValue) GetRemotePlmnListOk() (*[]PlmnId, bool)`

GetRemotePlmnListOk returns a tuple with the RemotePlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePlmnList

`func (o *NrfInfoServedScpInfoListValue) SetRemotePlmnList(v []PlmnId)`

SetRemotePlmnList sets RemotePlmnList field to given value.

### HasRemotePlmnList

`func (o *NrfInfoServedScpInfoListValue) HasRemotePlmnList() bool`

HasRemotePlmnList returns a boolean if a field has been set.

### GetRemoteSnpnList

`func (o *NrfInfoServedScpInfoListValue) GetRemoteSnpnList() []PlmnIdNid`

GetRemoteSnpnList returns the RemoteSnpnList field if non-nil, zero value otherwise.

### GetRemoteSnpnListOk

`func (o *NrfInfoServedScpInfoListValue) GetRemoteSnpnListOk() (*[]PlmnIdNid, bool)`

GetRemoteSnpnListOk returns a tuple with the RemoteSnpnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSnpnList

`func (o *NrfInfoServedScpInfoListValue) SetRemoteSnpnList(v []PlmnIdNid)`

SetRemoteSnpnList sets RemoteSnpnList field to given value.

### HasRemoteSnpnList

`func (o *NrfInfoServedScpInfoListValue) HasRemoteSnpnList() bool`

HasRemoteSnpnList returns a boolean if a field has been set.

### GetIpReachability

`func (o *NrfInfoServedScpInfoListValue) GetIpReachability() IpReachability`

GetIpReachability returns the IpReachability field if non-nil, zero value otherwise.

### GetIpReachabilityOk

`func (o *NrfInfoServedScpInfoListValue) GetIpReachabilityOk() (*IpReachability, bool)`

GetIpReachabilityOk returns a tuple with the IpReachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReachability

`func (o *NrfInfoServedScpInfoListValue) SetIpReachability(v IpReachability)`

SetIpReachability sets IpReachability field to given value.

### HasIpReachability

`func (o *NrfInfoServedScpInfoListValue) HasIpReachability() bool`

HasIpReachability returns a boolean if a field has been set.

### GetScpCapabilities

`func (o *NrfInfoServedScpInfoListValue) GetScpCapabilities() []ScpCapability`

GetScpCapabilities returns the ScpCapabilities field if non-nil, zero value otherwise.

### GetScpCapabilitiesOk

`func (o *NrfInfoServedScpInfoListValue) GetScpCapabilitiesOk() (*[]ScpCapability, bool)`

GetScpCapabilitiesOk returns a tuple with the ScpCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpCapabilities

`func (o *NrfInfoServedScpInfoListValue) SetScpCapabilities(v []ScpCapability)`

SetScpCapabilities sets ScpCapabilities field to given value.

### HasScpCapabilities

`func (o *NrfInfoServedScpInfoListValue) HasScpCapabilities() bool`

HasScpCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


