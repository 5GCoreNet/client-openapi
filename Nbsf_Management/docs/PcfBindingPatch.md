# PcfBindingPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addr** | Pointer to **NullableString** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166 with the OpenAPI defined &#39;nullable: true&#39; property.  | [optional] 
**IpDomain** | Pointer to **NullableString** |  | [optional] 
**Ipv6Prefix** | Pointer to [**NullableIpv6PrefixRm**](Ipv6PrefixRm.md) |  | [optional] 
**AddIpv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) | The additional IPv6 Address Prefixes of the served UE. | [optional] 
**MacAddr48** | Pointer to **NullableString** | \&quot;String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042 with the OpenAPI &#39;nullable: true&#39; property.\&quot;  | [optional] 
**AddMacAddrs** | Pointer to **[]string** | The additional MAC Addresses of the served UE. | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) | IP end points of the PCF hosting the Npcf_PolicyAuthorization service. | [optional] 
**PcfDiamHost** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfDiamRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewPcfBindingPatch

`func NewPcfBindingPatch() *PcfBindingPatch`

NewPcfBindingPatch instantiates a new PcfBindingPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfBindingPatchWithDefaults

`func NewPcfBindingPatchWithDefaults() *PcfBindingPatch`

NewPcfBindingPatchWithDefaults instantiates a new PcfBindingPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addr

`func (o *PcfBindingPatch) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *PcfBindingPatch) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *PcfBindingPatch) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *PcfBindingPatch) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### SetIpv4AddrNil

`func (o *PcfBindingPatch) SetIpv4AddrNil(b bool)`

 SetIpv4AddrNil sets the value for Ipv4Addr to be an explicit nil

### UnsetIpv4Addr
`func (o *PcfBindingPatch) UnsetIpv4Addr()`

UnsetIpv4Addr ensures that no value is present for Ipv4Addr, not even an explicit nil
### GetIpDomain

`func (o *PcfBindingPatch) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *PcfBindingPatch) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *PcfBindingPatch) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *PcfBindingPatch) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### SetIpDomainNil

`func (o *PcfBindingPatch) SetIpDomainNil(b bool)`

 SetIpDomainNil sets the value for IpDomain to be an explicit nil

### UnsetIpDomain
`func (o *PcfBindingPatch) UnsetIpDomain()`

UnsetIpDomain ensures that no value is present for IpDomain, not even an explicit nil
### GetIpv6Prefix

`func (o *PcfBindingPatch) GetIpv6Prefix() Ipv6PrefixRm`

GetIpv6Prefix returns the Ipv6Prefix field if non-nil, zero value otherwise.

### GetIpv6PrefixOk

`func (o *PcfBindingPatch) GetIpv6PrefixOk() (*Ipv6PrefixRm, bool)`

GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefix

`func (o *PcfBindingPatch) SetIpv6Prefix(v Ipv6PrefixRm)`

SetIpv6Prefix sets Ipv6Prefix field to given value.

### HasIpv6Prefix

`func (o *PcfBindingPatch) HasIpv6Prefix() bool`

HasIpv6Prefix returns a boolean if a field has been set.

### SetIpv6PrefixNil

`func (o *PcfBindingPatch) SetIpv6PrefixNil(b bool)`

 SetIpv6PrefixNil sets the value for Ipv6Prefix to be an explicit nil

### UnsetIpv6Prefix
`func (o *PcfBindingPatch) UnsetIpv6Prefix()`

UnsetIpv6Prefix ensures that no value is present for Ipv6Prefix, not even an explicit nil
### GetAddIpv6Prefixes

`func (o *PcfBindingPatch) GetAddIpv6Prefixes() []Ipv6Prefix`

GetAddIpv6Prefixes returns the AddIpv6Prefixes field if non-nil, zero value otherwise.

### GetAddIpv6PrefixesOk

`func (o *PcfBindingPatch) GetAddIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetAddIpv6PrefixesOk returns a tuple with the AddIpv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddIpv6Prefixes

`func (o *PcfBindingPatch) SetAddIpv6Prefixes(v []Ipv6Prefix)`

SetAddIpv6Prefixes sets AddIpv6Prefixes field to given value.

### HasAddIpv6Prefixes

`func (o *PcfBindingPatch) HasAddIpv6Prefixes() bool`

HasAddIpv6Prefixes returns a boolean if a field has been set.

### SetAddIpv6PrefixesNil

`func (o *PcfBindingPatch) SetAddIpv6PrefixesNil(b bool)`

 SetAddIpv6PrefixesNil sets the value for AddIpv6Prefixes to be an explicit nil

### UnsetAddIpv6Prefixes
`func (o *PcfBindingPatch) UnsetAddIpv6Prefixes()`

UnsetAddIpv6Prefixes ensures that no value is present for AddIpv6Prefixes, not even an explicit nil
### GetMacAddr48

`func (o *PcfBindingPatch) GetMacAddr48() string`

GetMacAddr48 returns the MacAddr48 field if non-nil, zero value otherwise.

### GetMacAddr48Ok

`func (o *PcfBindingPatch) GetMacAddr48Ok() (*string, bool)`

GetMacAddr48Ok returns a tuple with the MacAddr48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddr48

`func (o *PcfBindingPatch) SetMacAddr48(v string)`

SetMacAddr48 sets MacAddr48 field to given value.

### HasMacAddr48

`func (o *PcfBindingPatch) HasMacAddr48() bool`

HasMacAddr48 returns a boolean if a field has been set.

### SetMacAddr48Nil

`func (o *PcfBindingPatch) SetMacAddr48Nil(b bool)`

 SetMacAddr48Nil sets the value for MacAddr48 to be an explicit nil

### UnsetMacAddr48
`func (o *PcfBindingPatch) UnsetMacAddr48()`

UnsetMacAddr48 ensures that no value is present for MacAddr48, not even an explicit nil
### GetAddMacAddrs

`func (o *PcfBindingPatch) GetAddMacAddrs() []string`

GetAddMacAddrs returns the AddMacAddrs field if non-nil, zero value otherwise.

### GetAddMacAddrsOk

`func (o *PcfBindingPatch) GetAddMacAddrsOk() (*[]string, bool)`

GetAddMacAddrsOk returns a tuple with the AddMacAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMacAddrs

`func (o *PcfBindingPatch) SetAddMacAddrs(v []string)`

SetAddMacAddrs sets AddMacAddrs field to given value.

### HasAddMacAddrs

`func (o *PcfBindingPatch) HasAddMacAddrs() bool`

HasAddMacAddrs returns a boolean if a field has been set.

### SetAddMacAddrsNil

`func (o *PcfBindingPatch) SetAddMacAddrsNil(b bool)`

 SetAddMacAddrsNil sets the value for AddMacAddrs to be an explicit nil

### UnsetAddMacAddrs
`func (o *PcfBindingPatch) UnsetAddMacAddrs()`

UnsetAddMacAddrs ensures that no value is present for AddMacAddrs, not even an explicit nil
### GetPcfId

`func (o *PcfBindingPatch) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *PcfBindingPatch) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *PcfBindingPatch) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *PcfBindingPatch) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfFqdn

`func (o *PcfBindingPatch) GetPcfFqdn() string`

GetPcfFqdn returns the PcfFqdn field if non-nil, zero value otherwise.

### GetPcfFqdnOk

`func (o *PcfBindingPatch) GetPcfFqdnOk() (*string, bool)`

GetPcfFqdnOk returns a tuple with the PcfFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFqdn

`func (o *PcfBindingPatch) SetPcfFqdn(v string)`

SetPcfFqdn sets PcfFqdn field to given value.

### HasPcfFqdn

`func (o *PcfBindingPatch) HasPcfFqdn() bool`

HasPcfFqdn returns a boolean if a field has been set.

### GetPcfIpEndPoints

`func (o *PcfBindingPatch) GetPcfIpEndPoints() []IpEndPoint`

GetPcfIpEndPoints returns the PcfIpEndPoints field if non-nil, zero value otherwise.

### GetPcfIpEndPointsOk

`func (o *PcfBindingPatch) GetPcfIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfIpEndPoints

`func (o *PcfBindingPatch) SetPcfIpEndPoints(v []IpEndPoint)`

SetPcfIpEndPoints sets PcfIpEndPoints field to given value.

### HasPcfIpEndPoints

`func (o *PcfBindingPatch) HasPcfIpEndPoints() bool`

HasPcfIpEndPoints returns a boolean if a field has been set.

### GetPcfDiamHost

`func (o *PcfBindingPatch) GetPcfDiamHost() string`

GetPcfDiamHost returns the PcfDiamHost field if non-nil, zero value otherwise.

### GetPcfDiamHostOk

`func (o *PcfBindingPatch) GetPcfDiamHostOk() (*string, bool)`

GetPcfDiamHostOk returns a tuple with the PcfDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfDiamHost

`func (o *PcfBindingPatch) SetPcfDiamHost(v string)`

SetPcfDiamHost sets PcfDiamHost field to given value.

### HasPcfDiamHost

`func (o *PcfBindingPatch) HasPcfDiamHost() bool`

HasPcfDiamHost returns a boolean if a field has been set.

### GetPcfDiamRealm

`func (o *PcfBindingPatch) GetPcfDiamRealm() string`

GetPcfDiamRealm returns the PcfDiamRealm field if non-nil, zero value otherwise.

### GetPcfDiamRealmOk

`func (o *PcfBindingPatch) GetPcfDiamRealmOk() (*string, bool)`

GetPcfDiamRealmOk returns a tuple with the PcfDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfDiamRealm

`func (o *PcfBindingPatch) SetPcfDiamRealm(v string)`

SetPcfDiamRealm sets PcfDiamRealm field to given value.

### HasPcfDiamRealm

`func (o *PcfBindingPatch) HasPcfDiamRealm() bool`

HasPcfDiamRealm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


