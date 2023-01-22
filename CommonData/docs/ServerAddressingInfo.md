# ServerAddressingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Ipv6Addresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**FqdnList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServerAddressingInfo

`func NewServerAddressingInfo() *ServerAddressingInfo`

NewServerAddressingInfo instantiates a new ServerAddressingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerAddressingInfoWithDefaults

`func NewServerAddressingInfoWithDefaults() *ServerAddressingInfo`

NewServerAddressingInfoWithDefaults instantiates a new ServerAddressingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addresses

`func (o *ServerAddressingInfo) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *ServerAddressingInfo) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *ServerAddressingInfo) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *ServerAddressingInfo) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *ServerAddressingInfo) GetIpv6Addresses() []Ipv6Addr`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *ServerAddressingInfo) GetIpv6AddressesOk() (*[]Ipv6Addr, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *ServerAddressingInfo) SetIpv6Addresses(v []Ipv6Addr)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *ServerAddressingInfo) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetFqdnList

`func (o *ServerAddressingInfo) GetFqdnList() []string`

GetFqdnList returns the FqdnList field if non-nil, zero value otherwise.

### GetFqdnListOk

`func (o *ServerAddressingInfo) GetFqdnListOk() (*[]string, bool)`

GetFqdnListOk returns a tuple with the FqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnList

`func (o *ServerAddressingInfo) SetFqdnList(v []string)`

SetFqdnList sets FqdnList field to given value.

### HasFqdnList

`func (o *ServerAddressingInfo) HasFqdnList() bool`

HasFqdnList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


