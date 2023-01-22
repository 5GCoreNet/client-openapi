# IpInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4EndpointAddresses** | Pointer to **string** |  | [optional] 
**Ipv6EndpointAddresses** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 

## Methods

### NewIpInterface

`func NewIpInterface() *IpInterface`

NewIpInterface instantiates a new IpInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpInterfaceWithDefaults

`func NewIpInterfaceWithDefaults() *IpInterface`

NewIpInterfaceWithDefaults instantiates a new IpInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4EndpointAddresses

`func (o *IpInterface) GetIpv4EndpointAddresses() string`

GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv4EndpointAddressesOk

`func (o *IpInterface) GetIpv4EndpointAddressesOk() (*string, bool)`

GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4EndpointAddresses

`func (o *IpInterface) SetIpv4EndpointAddresses(v string)`

SetIpv4EndpointAddresses sets Ipv4EndpointAddresses field to given value.

### HasIpv4EndpointAddresses

`func (o *IpInterface) HasIpv4EndpointAddresses() bool`

HasIpv4EndpointAddresses returns a boolean if a field has been set.

### GetIpv6EndpointAddresses

`func (o *IpInterface) GetIpv6EndpointAddresses() Ipv6Addr`

GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv6EndpointAddressesOk

`func (o *IpInterface) GetIpv6EndpointAddressesOk() (*Ipv6Addr, bool)`

GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EndpointAddresses

`func (o *IpInterface) SetIpv6EndpointAddresses(v Ipv6Addr)`

SetIpv6EndpointAddresses sets Ipv6EndpointAddresses field to given value.

### HasIpv6EndpointAddresses

`func (o *IpInterface) HasIpv6EndpointAddresses() bool`

HasIpv6EndpointAddresses returns a boolean if a field has been set.

### GetFqdn

`func (o *IpInterface) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *IpInterface) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *IpInterface) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *IpInterface) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


