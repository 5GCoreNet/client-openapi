# DnaiInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnai** | **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | 
**DnsServIds** | Pointer to [**[]DnsServerIdentifier**](DnsServerIdentifier.md) |  | [optional] 
**EasIpAddrs** | Pointer to [**[]IpAddr**](IpAddr.md) |  | [optional] 

## Methods

### NewDnaiInformation

`func NewDnaiInformation(dnai string, ) *DnaiInformation`

NewDnaiInformation instantiates a new DnaiInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaiInformationWithDefaults

`func NewDnaiInformationWithDefaults() *DnaiInformation`

NewDnaiInformationWithDefaults instantiates a new DnaiInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnai

`func (o *DnaiInformation) GetDnai() string`

GetDnai returns the Dnai field if non-nil, zero value otherwise.

### GetDnaiOk

`func (o *DnaiInformation) GetDnaiOk() (*string, bool)`

GetDnaiOk returns a tuple with the Dnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnai

`func (o *DnaiInformation) SetDnai(v string)`

SetDnai sets Dnai field to given value.


### GetDnsServIds

`func (o *DnaiInformation) GetDnsServIds() []DnsServerIdentifier`

GetDnsServIds returns the DnsServIds field if non-nil, zero value otherwise.

### GetDnsServIdsOk

`func (o *DnaiInformation) GetDnsServIdsOk() (*[]DnsServerIdentifier, bool)`

GetDnsServIdsOk returns a tuple with the DnsServIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServIds

`func (o *DnaiInformation) SetDnsServIds(v []DnsServerIdentifier)`

SetDnsServIds sets DnsServIds field to given value.

### HasDnsServIds

`func (o *DnaiInformation) HasDnsServIds() bool`

HasDnsServIds returns a boolean if a field has been set.

### GetEasIpAddrs

`func (o *DnaiInformation) GetEasIpAddrs() []IpAddr`

GetEasIpAddrs returns the EasIpAddrs field if non-nil, zero value otherwise.

### GetEasIpAddrsOk

`func (o *DnaiInformation) GetEasIpAddrsOk() (*[]IpAddr, bool)`

GetEasIpAddrsOk returns a tuple with the EasIpAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpAddrs

`func (o *DnaiInformation) SetEasIpAddrs(v []IpAddr)`

SetEasIpAddrs sets EasIpAddrs field to given value.

### HasEasIpAddrs

`func (o *DnaiInformation) HasEasIpAddrs() bool`

HasEasIpAddrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


