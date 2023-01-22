# RoutingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4AddrRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**Ipv6AddrRanges** | Pointer to [**[]Ipv6AddressRange**](Ipv6AddressRange.md) |  | [optional] 
**AefProfile** | [**AefProfile**](AefProfile.md) |  | 

## Methods

### NewRoutingRule

`func NewRoutingRule(aefProfile AefProfile, ) *RoutingRule`

NewRoutingRule instantiates a new RoutingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRuleWithDefaults

`func NewRoutingRuleWithDefaults() *RoutingRule`

NewRoutingRuleWithDefaults instantiates a new RoutingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4AddrRanges

`func (o *RoutingRule) GetIpv4AddrRanges() []Ipv4AddressRange`

GetIpv4AddrRanges returns the Ipv4AddrRanges field if non-nil, zero value otherwise.

### GetIpv4AddrRangesOk

`func (o *RoutingRule) GetIpv4AddrRangesOk() (*[]Ipv4AddressRange, bool)`

GetIpv4AddrRangesOk returns a tuple with the Ipv4AddrRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddrRanges

`func (o *RoutingRule) SetIpv4AddrRanges(v []Ipv4AddressRange)`

SetIpv4AddrRanges sets Ipv4AddrRanges field to given value.

### HasIpv4AddrRanges

`func (o *RoutingRule) HasIpv4AddrRanges() bool`

HasIpv4AddrRanges returns a boolean if a field has been set.

### GetIpv6AddrRanges

`func (o *RoutingRule) GetIpv6AddrRanges() []Ipv6AddressRange`

GetIpv6AddrRanges returns the Ipv6AddrRanges field if non-nil, zero value otherwise.

### GetIpv6AddrRangesOk

`func (o *RoutingRule) GetIpv6AddrRangesOk() (*[]Ipv6AddressRange, bool)`

GetIpv6AddrRangesOk returns a tuple with the Ipv6AddrRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddrRanges

`func (o *RoutingRule) SetIpv6AddrRanges(v []Ipv6AddressRange)`

SetIpv6AddrRanges sets Ipv6AddrRanges field to given value.

### HasIpv6AddrRanges

`func (o *RoutingRule) HasIpv6AddrRanges() bool`

HasIpv6AddrRanges returns a boolean if a field has been set.

### GetAefProfile

`func (o *RoutingRule) GetAefProfile() AefProfile`

GetAefProfile returns the AefProfile field if non-nil, zero value otherwise.

### GetAefProfileOk

`func (o *RoutingRule) GetAefProfileOk() (*AefProfile, bool)`

GetAefProfileOk returns a tuple with the AefProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefProfile

`func (o *RoutingRule) SetAefProfile(v AefProfile)`

SetAefProfile sets AefProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


