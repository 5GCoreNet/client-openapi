# EndpointAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**Ipv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PortNumber** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | 

## Methods

### NewEndpointAddress

`func NewEndpointAddress(portNumber int32, ) *EndpointAddress`

NewEndpointAddress instantiates a new EndpointAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointAddressWithDefaults

`func NewEndpointAddressWithDefaults() *EndpointAddress`

NewEndpointAddressWithDefaults instantiates a new EndpointAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addr

`func (o *EndpointAddress) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *EndpointAddress) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *EndpointAddress) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *EndpointAddress) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *EndpointAddress) GetIpv6Addr() Ipv6Addr`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *EndpointAddress) GetIpv6AddrOk() (*Ipv6Addr, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *EndpointAddress) SetIpv6Addr(v Ipv6Addr)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *EndpointAddress) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.

### GetPortNumber

`func (o *EndpointAddress) GetPortNumber() int32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *EndpointAddress) GetPortNumberOk() (*int32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *EndpointAddress) SetPortNumber(v int32)`

SetPortNumber sets PortNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


