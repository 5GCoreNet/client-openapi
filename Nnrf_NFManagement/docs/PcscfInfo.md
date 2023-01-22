# PcscfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**DnnList** | Pointer to **[]string** |  | [optional] 
**GmFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**GmIpv4Addresses** | Pointer to **[]string** |  | [optional] 
**GmIpv6Addresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**MwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**MwIpv4Addresses** | Pointer to **[]string** |  | [optional] 
**MwIpv6Addresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**ServedIpv4AddressRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**ServedIpv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 

## Methods

### NewPcscfInfo

`func NewPcscfInfo() *PcscfInfo`

NewPcscfInfo instantiates a new PcscfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcscfInfoWithDefaults

`func NewPcscfInfoWithDefaults() *PcscfInfo`

NewPcscfInfoWithDefaults instantiates a new PcscfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *PcscfInfo) GetAccessType() []AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PcscfInfo) GetAccessTypeOk() (*[]AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PcscfInfo) SetAccessType(v []AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *PcscfInfo) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetDnnList

`func (o *PcscfInfo) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *PcscfInfo) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *PcscfInfo) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *PcscfInfo) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetGmFqdn

`func (o *PcscfInfo) GetGmFqdn() string`

GetGmFqdn returns the GmFqdn field if non-nil, zero value otherwise.

### GetGmFqdnOk

`func (o *PcscfInfo) GetGmFqdnOk() (*string, bool)`

GetGmFqdnOk returns a tuple with the GmFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmFqdn

`func (o *PcscfInfo) SetGmFqdn(v string)`

SetGmFqdn sets GmFqdn field to given value.

### HasGmFqdn

`func (o *PcscfInfo) HasGmFqdn() bool`

HasGmFqdn returns a boolean if a field has been set.

### GetGmIpv4Addresses

`func (o *PcscfInfo) GetGmIpv4Addresses() []string`

GetGmIpv4Addresses returns the GmIpv4Addresses field if non-nil, zero value otherwise.

### GetGmIpv4AddressesOk

`func (o *PcscfInfo) GetGmIpv4AddressesOk() (*[]string, bool)`

GetGmIpv4AddressesOk returns a tuple with the GmIpv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmIpv4Addresses

`func (o *PcscfInfo) SetGmIpv4Addresses(v []string)`

SetGmIpv4Addresses sets GmIpv4Addresses field to given value.

### HasGmIpv4Addresses

`func (o *PcscfInfo) HasGmIpv4Addresses() bool`

HasGmIpv4Addresses returns a boolean if a field has been set.

### GetGmIpv6Addresses

`func (o *PcscfInfo) GetGmIpv6Addresses() []Ipv6Addr`

GetGmIpv6Addresses returns the GmIpv6Addresses field if non-nil, zero value otherwise.

### GetGmIpv6AddressesOk

`func (o *PcscfInfo) GetGmIpv6AddressesOk() (*[]Ipv6Addr, bool)`

GetGmIpv6AddressesOk returns a tuple with the GmIpv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmIpv6Addresses

`func (o *PcscfInfo) SetGmIpv6Addresses(v []Ipv6Addr)`

SetGmIpv6Addresses sets GmIpv6Addresses field to given value.

### HasGmIpv6Addresses

`func (o *PcscfInfo) HasGmIpv6Addresses() bool`

HasGmIpv6Addresses returns a boolean if a field has been set.

### GetMwFqdn

`func (o *PcscfInfo) GetMwFqdn() string`

GetMwFqdn returns the MwFqdn field if non-nil, zero value otherwise.

### GetMwFqdnOk

`func (o *PcscfInfo) GetMwFqdnOk() (*string, bool)`

GetMwFqdnOk returns a tuple with the MwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMwFqdn

`func (o *PcscfInfo) SetMwFqdn(v string)`

SetMwFqdn sets MwFqdn field to given value.

### HasMwFqdn

`func (o *PcscfInfo) HasMwFqdn() bool`

HasMwFqdn returns a boolean if a field has been set.

### GetMwIpv4Addresses

`func (o *PcscfInfo) GetMwIpv4Addresses() []string`

GetMwIpv4Addresses returns the MwIpv4Addresses field if non-nil, zero value otherwise.

### GetMwIpv4AddressesOk

`func (o *PcscfInfo) GetMwIpv4AddressesOk() (*[]string, bool)`

GetMwIpv4AddressesOk returns a tuple with the MwIpv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMwIpv4Addresses

`func (o *PcscfInfo) SetMwIpv4Addresses(v []string)`

SetMwIpv4Addresses sets MwIpv4Addresses field to given value.

### HasMwIpv4Addresses

`func (o *PcscfInfo) HasMwIpv4Addresses() bool`

HasMwIpv4Addresses returns a boolean if a field has been set.

### GetMwIpv6Addresses

`func (o *PcscfInfo) GetMwIpv6Addresses() []Ipv6Addr`

GetMwIpv6Addresses returns the MwIpv6Addresses field if non-nil, zero value otherwise.

### GetMwIpv6AddressesOk

`func (o *PcscfInfo) GetMwIpv6AddressesOk() (*[]Ipv6Addr, bool)`

GetMwIpv6AddressesOk returns a tuple with the MwIpv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMwIpv6Addresses

`func (o *PcscfInfo) SetMwIpv6Addresses(v []Ipv6Addr)`

SetMwIpv6Addresses sets MwIpv6Addresses field to given value.

### HasMwIpv6Addresses

`func (o *PcscfInfo) HasMwIpv6Addresses() bool`

HasMwIpv6Addresses returns a boolean if a field has been set.

### GetServedIpv4AddressRanges

`func (o *PcscfInfo) GetServedIpv4AddressRanges() []Ipv4AddressRange`

GetServedIpv4AddressRanges returns the ServedIpv4AddressRanges field if non-nil, zero value otherwise.

### GetServedIpv4AddressRangesOk

`func (o *PcscfInfo) GetServedIpv4AddressRangesOk() (*[]Ipv4AddressRange, bool)`

GetServedIpv4AddressRangesOk returns a tuple with the ServedIpv4AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedIpv4AddressRanges

`func (o *PcscfInfo) SetServedIpv4AddressRanges(v []Ipv4AddressRange)`

SetServedIpv4AddressRanges sets ServedIpv4AddressRanges field to given value.

### HasServedIpv4AddressRanges

`func (o *PcscfInfo) HasServedIpv4AddressRanges() bool`

HasServedIpv4AddressRanges returns a boolean if a field has been set.

### GetServedIpv6PrefixRanges

`func (o *PcscfInfo) GetServedIpv6PrefixRanges() []Ipv6PrefixRange`

GetServedIpv6PrefixRanges returns the ServedIpv6PrefixRanges field if non-nil, zero value otherwise.

### GetServedIpv6PrefixRangesOk

`func (o *PcscfInfo) GetServedIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetServedIpv6PrefixRangesOk returns a tuple with the ServedIpv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedIpv6PrefixRanges

`func (o *PcscfInfo) SetServedIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetServedIpv6PrefixRanges sets ServedIpv6PrefixRanges field to given value.

### HasServedIpv6PrefixRanges

`func (o *PcscfInfo) HasServedIpv6PrefixRanges() bool`

HasServedIpv6PrefixRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


