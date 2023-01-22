# PcfForPduSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**PcfFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) | IP end points of the PCF hosting the Npcf_AmPolicyAuthorization service. | [optional] 
**Ipv4Addr** | Pointer to **NullableString** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166 with the OpenAPI defined &#39;nullable: true&#39; property.  | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**Ipv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) | The IPv6 Address Prefixes of the served UE. | [optional] 
**MacAddrs** | Pointer to **[]string** | The MAC Addresses of the served UE. | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**BindLevel** | Pointer to [**BindingLevel**](BindingLevel.md) |  | [optional] 

## Methods

### NewPcfForPduSessionInfo

`func NewPcfForPduSessionInfo(dnn string, snssai Snssai, ) *PcfForPduSessionInfo`

NewPcfForPduSessionInfo instantiates a new PcfForPduSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfForPduSessionInfoWithDefaults

`func NewPcfForPduSessionInfoWithDefaults() *PcfForPduSessionInfo`

NewPcfForPduSessionInfoWithDefaults instantiates a new PcfForPduSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *PcfForPduSessionInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PcfForPduSessionInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PcfForPduSessionInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSnssai

`func (o *PcfForPduSessionInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *PcfForPduSessionInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *PcfForPduSessionInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetPcfFqdn

`func (o *PcfForPduSessionInfo) GetPcfFqdn() string`

GetPcfFqdn returns the PcfFqdn field if non-nil, zero value otherwise.

### GetPcfFqdnOk

`func (o *PcfForPduSessionInfo) GetPcfFqdnOk() (*string, bool)`

GetPcfFqdnOk returns a tuple with the PcfFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFqdn

`func (o *PcfForPduSessionInfo) SetPcfFqdn(v string)`

SetPcfFqdn sets PcfFqdn field to given value.

### HasPcfFqdn

`func (o *PcfForPduSessionInfo) HasPcfFqdn() bool`

HasPcfFqdn returns a boolean if a field has been set.

### GetPcfIpEndPoints

`func (o *PcfForPduSessionInfo) GetPcfIpEndPoints() []IpEndPoint`

GetPcfIpEndPoints returns the PcfIpEndPoints field if non-nil, zero value otherwise.

### GetPcfIpEndPointsOk

`func (o *PcfForPduSessionInfo) GetPcfIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfIpEndPoints

`func (o *PcfForPduSessionInfo) SetPcfIpEndPoints(v []IpEndPoint)`

SetPcfIpEndPoints sets PcfIpEndPoints field to given value.

### HasPcfIpEndPoints

`func (o *PcfForPduSessionInfo) HasPcfIpEndPoints() bool`

HasPcfIpEndPoints returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *PcfForPduSessionInfo) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *PcfForPduSessionInfo) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *PcfForPduSessionInfo) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *PcfForPduSessionInfo) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### SetIpv4AddrNil

`func (o *PcfForPduSessionInfo) SetIpv4AddrNil(b bool)`

 SetIpv4AddrNil sets the value for Ipv4Addr to be an explicit nil

### UnsetIpv4Addr
`func (o *PcfForPduSessionInfo) UnsetIpv4Addr()`

UnsetIpv4Addr ensures that no value is present for Ipv4Addr, not even an explicit nil
### GetIpDomain

`func (o *PcfForPduSessionInfo) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *PcfForPduSessionInfo) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *PcfForPduSessionInfo) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *PcfForPduSessionInfo) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetIpv6Prefixes

`func (o *PcfForPduSessionInfo) GetIpv6Prefixes() []Ipv6Prefix`

GetIpv6Prefixes returns the Ipv6Prefixes field if non-nil, zero value otherwise.

### GetIpv6PrefixesOk

`func (o *PcfForPduSessionInfo) GetIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefixes

`func (o *PcfForPduSessionInfo) SetIpv6Prefixes(v []Ipv6Prefix)`

SetIpv6Prefixes sets Ipv6Prefixes field to given value.

### HasIpv6Prefixes

`func (o *PcfForPduSessionInfo) HasIpv6Prefixes() bool`

HasIpv6Prefixes returns a boolean if a field has been set.

### GetMacAddrs

`func (o *PcfForPduSessionInfo) GetMacAddrs() []string`

GetMacAddrs returns the MacAddrs field if non-nil, zero value otherwise.

### GetMacAddrsOk

`func (o *PcfForPduSessionInfo) GetMacAddrsOk() (*[]string, bool)`

GetMacAddrsOk returns a tuple with the MacAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddrs

`func (o *PcfForPduSessionInfo) SetMacAddrs(v []string)`

SetMacAddrs sets MacAddrs field to given value.

### HasMacAddrs

`func (o *PcfForPduSessionInfo) HasMacAddrs() bool`

HasMacAddrs returns a boolean if a field has been set.

### GetPcfId

`func (o *PcfForPduSessionInfo) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *PcfForPduSessionInfo) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *PcfForPduSessionInfo) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *PcfForPduSessionInfo) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *PcfForPduSessionInfo) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *PcfForPduSessionInfo) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *PcfForPduSessionInfo) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *PcfForPduSessionInfo) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetBindLevel

`func (o *PcfForPduSessionInfo) GetBindLevel() BindingLevel`

GetBindLevel returns the BindLevel field if non-nil, zero value otherwise.

### GetBindLevelOk

`func (o *PcfForPduSessionInfo) GetBindLevelOk() (*BindingLevel, bool)`

GetBindLevelOk returns a tuple with the BindLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindLevel

`func (o *PcfForPduSessionInfo) SetBindLevel(v BindingLevel)`

SetBindLevel sets BindLevel field to given value.

### HasBindLevel

`func (o *PcfForPduSessionInfo) HasBindLevel() bool`

HasBindLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


