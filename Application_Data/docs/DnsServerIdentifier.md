# DnsServerIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServIpAddr** | [**IpAddr**](IpAddr.md) |  | 
**PortNumber** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 

## Methods

### NewDnsServerIdentifier

`func NewDnsServerIdentifier(dnsServIpAddr IpAddr, portNumber int32, ) *DnsServerIdentifier`

NewDnsServerIdentifier instantiates a new DnsServerIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsServerIdentifierWithDefaults

`func NewDnsServerIdentifierWithDefaults() *DnsServerIdentifier`

NewDnsServerIdentifierWithDefaults instantiates a new DnsServerIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServIpAddr

`func (o *DnsServerIdentifier) GetDnsServIpAddr() IpAddr`

GetDnsServIpAddr returns the DnsServIpAddr field if non-nil, zero value otherwise.

### GetDnsServIpAddrOk

`func (o *DnsServerIdentifier) GetDnsServIpAddrOk() (*IpAddr, bool)`

GetDnsServIpAddrOk returns a tuple with the DnsServIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServIpAddr

`func (o *DnsServerIdentifier) SetDnsServIpAddr(v IpAddr)`

SetDnsServIpAddr sets DnsServIpAddr field to given value.


### GetPortNumber

`func (o *DnsServerIdentifier) GetPortNumber() int32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *DnsServerIdentifier) GetPortNumberOk() (*int32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *DnsServerIdentifier) SetPortNumber(v int32)`

SetPortNumber sets PortNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


