# Ipv4AddressRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**End** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 

## Methods

### NewIpv4AddressRange

`func NewIpv4AddressRange() *Ipv4AddressRange`

NewIpv4AddressRange instantiates a new Ipv4AddressRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4AddressRangeWithDefaults

`func NewIpv4AddressRangeWithDefaults() *Ipv4AddressRange`

NewIpv4AddressRangeWithDefaults instantiates a new Ipv4AddressRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *Ipv4AddressRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ipv4AddressRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Ipv4AddressRange) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *Ipv4AddressRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Ipv4AddressRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ipv4AddressRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Ipv4AddressRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Ipv4AddressRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


