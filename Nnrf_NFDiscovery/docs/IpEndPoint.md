# IpEndPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Address** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**Ipv6Address** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**Transport** | Pointer to [**TransportProtocol**](TransportProtocol.md) |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 

## Methods

### NewIpEndPoint

`func NewIpEndPoint() *IpEndPoint`

NewIpEndPoint instantiates a new IpEndPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpEndPointWithDefaults

`func NewIpEndPointWithDefaults() *IpEndPoint`

NewIpEndPointWithDefaults instantiates a new IpEndPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Address

`func (o *IpEndPoint) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *IpEndPoint) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *IpEndPoint) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *IpEndPoint) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *IpEndPoint) GetIpv6Address() Ipv6Addr`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *IpEndPoint) GetIpv6AddressOk() (*Ipv6Addr, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *IpEndPoint) SetIpv6Address(v Ipv6Addr)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *IpEndPoint) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetTransport

`func (o *IpEndPoint) GetTransport() TransportProtocol`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *IpEndPoint) GetTransportOk() (*TransportProtocol, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *IpEndPoint) SetTransport(v TransportProtocol)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *IpEndPoint) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetPort

`func (o *IpEndPoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IpEndPoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IpEndPoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IpEndPoint) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


