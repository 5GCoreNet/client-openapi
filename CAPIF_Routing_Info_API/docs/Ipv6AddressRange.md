# Ipv6AddressRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | 
**End** | **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | 

## Methods

### NewIpv6AddressRange

`func NewIpv6AddressRange(start string, end string, ) *Ipv6AddressRange`

NewIpv6AddressRange instantiates a new Ipv6AddressRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6AddressRangeWithDefaults

`func NewIpv6AddressRangeWithDefaults() *Ipv6AddressRange`

NewIpv6AddressRangeWithDefaults instantiates a new Ipv6AddressRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *Ipv6AddressRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ipv6AddressRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Ipv6AddressRange) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Ipv6AddressRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ipv6AddressRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Ipv6AddressRange) SetEnd(v string)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


