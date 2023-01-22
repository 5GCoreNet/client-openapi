# SgsnAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SgsnIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**SgsnIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 

## Methods

### NewSgsnAddress

`func NewSgsnAddress() *SgsnAddress`

NewSgsnAddress instantiates a new SgsnAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSgsnAddressWithDefaults

`func NewSgsnAddressWithDefaults() *SgsnAddress`

NewSgsnAddressWithDefaults instantiates a new SgsnAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSgsnIpv4Addr

`func (o *SgsnAddress) GetSgsnIpv4Addr() string`

GetSgsnIpv4Addr returns the SgsnIpv4Addr field if non-nil, zero value otherwise.

### GetSgsnIpv4AddrOk

`func (o *SgsnAddress) GetSgsnIpv4AddrOk() (*string, bool)`

GetSgsnIpv4AddrOk returns a tuple with the SgsnIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgsnIpv4Addr

`func (o *SgsnAddress) SetSgsnIpv4Addr(v string)`

SetSgsnIpv4Addr sets SgsnIpv4Addr field to given value.

### HasSgsnIpv4Addr

`func (o *SgsnAddress) HasSgsnIpv4Addr() bool`

HasSgsnIpv4Addr returns a boolean if a field has been set.

### GetSgsnIpv6Addr

`func (o *SgsnAddress) GetSgsnIpv6Addr() Ipv6Addr`

GetSgsnIpv6Addr returns the SgsnIpv6Addr field if non-nil, zero value otherwise.

### GetSgsnIpv6AddrOk

`func (o *SgsnAddress) GetSgsnIpv6AddrOk() (*Ipv6Addr, bool)`

GetSgsnIpv6AddrOk returns a tuple with the SgsnIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgsnIpv6Addr

`func (o *SgsnAddress) SetSgsnIpv6Addr(v Ipv6Addr)`

SetSgsnIpv6Addr sets SgsnIpv6Addr field to given value.

### HasSgsnIpv6Addr

`func (o *SgsnAddress) HasSgsnIpv6Addr() bool`

HasSgsnIpv6Addr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


