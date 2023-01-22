# MulticastAccessControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrcIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**SrcIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**MulticastV4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**MulticastV6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**AccStatus** | [**AccessRightStatus**](AccessRightStatus.md) |  | 

## Methods

### NewMulticastAccessControl

`func NewMulticastAccessControl(accStatus AccessRightStatus, ) *MulticastAccessControl`

NewMulticastAccessControl instantiates a new MulticastAccessControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticastAccessControlWithDefaults

`func NewMulticastAccessControlWithDefaults() *MulticastAccessControl`

NewMulticastAccessControlWithDefaults instantiates a new MulticastAccessControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrcIpv4Addr

`func (o *MulticastAccessControl) GetSrcIpv4Addr() string`

GetSrcIpv4Addr returns the SrcIpv4Addr field if non-nil, zero value otherwise.

### GetSrcIpv4AddrOk

`func (o *MulticastAccessControl) GetSrcIpv4AddrOk() (*string, bool)`

GetSrcIpv4AddrOk returns a tuple with the SrcIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIpv4Addr

`func (o *MulticastAccessControl) SetSrcIpv4Addr(v string)`

SetSrcIpv4Addr sets SrcIpv4Addr field to given value.

### HasSrcIpv4Addr

`func (o *MulticastAccessControl) HasSrcIpv4Addr() bool`

HasSrcIpv4Addr returns a boolean if a field has been set.

### GetSrcIpv6Addr

`func (o *MulticastAccessControl) GetSrcIpv6Addr() Ipv6Addr`

GetSrcIpv6Addr returns the SrcIpv6Addr field if non-nil, zero value otherwise.

### GetSrcIpv6AddrOk

`func (o *MulticastAccessControl) GetSrcIpv6AddrOk() (*Ipv6Addr, bool)`

GetSrcIpv6AddrOk returns a tuple with the SrcIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIpv6Addr

`func (o *MulticastAccessControl) SetSrcIpv6Addr(v Ipv6Addr)`

SetSrcIpv6Addr sets SrcIpv6Addr field to given value.

### HasSrcIpv6Addr

`func (o *MulticastAccessControl) HasSrcIpv6Addr() bool`

HasSrcIpv6Addr returns a boolean if a field has been set.

### GetMulticastV4Addr

`func (o *MulticastAccessControl) GetMulticastV4Addr() string`

GetMulticastV4Addr returns the MulticastV4Addr field if non-nil, zero value otherwise.

### GetMulticastV4AddrOk

`func (o *MulticastAccessControl) GetMulticastV4AddrOk() (*string, bool)`

GetMulticastV4AddrOk returns a tuple with the MulticastV4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastV4Addr

`func (o *MulticastAccessControl) SetMulticastV4Addr(v string)`

SetMulticastV4Addr sets MulticastV4Addr field to given value.

### HasMulticastV4Addr

`func (o *MulticastAccessControl) HasMulticastV4Addr() bool`

HasMulticastV4Addr returns a boolean if a field has been set.

### GetMulticastV6Addr

`func (o *MulticastAccessControl) GetMulticastV6Addr() Ipv6Addr`

GetMulticastV6Addr returns the MulticastV6Addr field if non-nil, zero value otherwise.

### GetMulticastV6AddrOk

`func (o *MulticastAccessControl) GetMulticastV6AddrOk() (*Ipv6Addr, bool)`

GetMulticastV6AddrOk returns a tuple with the MulticastV6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastV6Addr

`func (o *MulticastAccessControl) SetMulticastV6Addr(v Ipv6Addr)`

SetMulticastV6Addr sets MulticastV6Addr field to given value.

### HasMulticastV6Addr

`func (o *MulticastAccessControl) HasMulticastV6Addr() bool`

HasMulticastV6Addr returns a boolean if a field has been set.

### GetAccStatus

`func (o *MulticastAccessControl) GetAccStatus() AccessRightStatus`

GetAccStatus returns the AccStatus field if non-nil, zero value otherwise.

### GetAccStatusOk

`func (o *MulticastAccessControl) GetAccStatusOk() (*AccessRightStatus, bool)`

GetAccStatusOk returns a tuple with the AccStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccStatus

`func (o *MulticastAccessControl) SetAccStatus(v AccessRightStatus)`

SetAccStatus sets AccStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


