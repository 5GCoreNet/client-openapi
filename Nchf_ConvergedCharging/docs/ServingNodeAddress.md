# ServingNodeAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**Ipv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 

## Methods

### NewServingNodeAddress

`func NewServingNodeAddress() *ServingNodeAddress`

NewServingNodeAddress instantiates a new ServingNodeAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServingNodeAddressWithDefaults

`func NewServingNodeAddressWithDefaults() *ServingNodeAddress`

NewServingNodeAddressWithDefaults instantiates a new ServingNodeAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addr

`func (o *ServingNodeAddress) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *ServingNodeAddress) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *ServingNodeAddress) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *ServingNodeAddress) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *ServingNodeAddress) GetIpv6Addr() Ipv6Addr`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *ServingNodeAddress) GetIpv6AddrOk() (*Ipv6Addr, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *ServingNodeAddress) SetIpv6Addr(v Ipv6Addr)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *ServingNodeAddress) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


