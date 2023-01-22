# AddressWithVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Address** | Pointer to **string** |  | [optional] 
**Ipv6Address** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 

## Methods

### NewAddressWithVlan

`func NewAddressWithVlan() *AddressWithVlan`

NewAddressWithVlan instantiates a new AddressWithVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithVlanWithDefaults

`func NewAddressWithVlanWithDefaults() *AddressWithVlan`

NewAddressWithVlanWithDefaults instantiates a new AddressWithVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Address

`func (o *AddressWithVlan) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *AddressWithVlan) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *AddressWithVlan) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *AddressWithVlan) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *AddressWithVlan) GetIpv6Address() Ipv6Addr`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *AddressWithVlan) GetIpv6AddressOk() (*Ipv6Addr, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *AddressWithVlan) SetIpv6Address(v Ipv6Addr)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *AddressWithVlan) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetVlanId

`func (o *AddressWithVlan) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *AddressWithVlan) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *AddressWithVlan) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *AddressWithVlan) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


