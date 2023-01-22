# EndPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**Ipv4Addrs** | Pointer to **[]string** | IPv4 addresses of the edge server. | [optional] 
**Ipv6Addrs** | Pointer to **[]string** | IPv6 addresses of the edge server. | [optional] 
**Uri** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 

## Methods

### NewEndPoint

`func NewEndPoint() *EndPoint`

NewEndPoint instantiates a new EndPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndPointWithDefaults

`func NewEndPointWithDefaults() *EndPoint`

NewEndPointWithDefaults instantiates a new EndPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *EndPoint) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *EndPoint) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *EndPoint) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *EndPoint) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetIpv4Addrs

`func (o *EndPoint) GetIpv4Addrs() []string`

GetIpv4Addrs returns the Ipv4Addrs field if non-nil, zero value otherwise.

### GetIpv4AddrsOk

`func (o *EndPoint) GetIpv4AddrsOk() (*[]string, bool)`

GetIpv4AddrsOk returns a tuple with the Ipv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addrs

`func (o *EndPoint) SetIpv4Addrs(v []string)`

SetIpv4Addrs sets Ipv4Addrs field to given value.

### HasIpv4Addrs

`func (o *EndPoint) HasIpv4Addrs() bool`

HasIpv4Addrs returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *EndPoint) GetIpv6Addrs() []string`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *EndPoint) GetIpv6AddrsOk() (*[]string, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *EndPoint) SetIpv6Addrs(v []string)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *EndPoint) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.

### GetUri

`func (o *EndPoint) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *EndPoint) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *EndPoint) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *EndPoint) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


