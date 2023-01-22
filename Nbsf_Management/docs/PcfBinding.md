# PcfBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Ipv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**Ipv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**AddIpv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) | The additional IPv6 Address Prefixes of the served UE. | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**MacAddr48** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**AddMacAddrs** | Pointer to **[]string** | The additional MAC Addresses of the served UE. | [optional] 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**PcfFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) | IP end points of the PCF hosting the Npcf_PolicyAuthorization service | [optional] 
**PcfDiamHost** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfDiamRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfSmFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfSmIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) | IP end points of the PCF hosting the Npcf_SMPolicyControl service. | [optional] 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ParaCom** | Pointer to [**ParameterCombination**](ParameterCombination.md) |  | [optional] 
**BindLevel** | Pointer to [**BindingLevel**](BindingLevel.md) |  | [optional] 
**Ipv4FrameRouteList** | Pointer to **[]string** |  | [optional] 
**Ipv6FrameRouteList** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 

## Methods

### NewPcfBinding

`func NewPcfBinding(dnn string, snssai Snssai, ) *PcfBinding`

NewPcfBinding instantiates a new PcfBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfBindingWithDefaults

`func NewPcfBindingWithDefaults() *PcfBinding`

NewPcfBindingWithDefaults instantiates a new PcfBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *PcfBinding) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PcfBinding) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PcfBinding) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PcfBinding) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *PcfBinding) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PcfBinding) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PcfBinding) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PcfBinding) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *PcfBinding) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *PcfBinding) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *PcfBinding) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *PcfBinding) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Prefix

`func (o *PcfBinding) GetIpv6Prefix() Ipv6Prefix`

GetIpv6Prefix returns the Ipv6Prefix field if non-nil, zero value otherwise.

### GetIpv6PrefixOk

`func (o *PcfBinding) GetIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefix

`func (o *PcfBinding) SetIpv6Prefix(v Ipv6Prefix)`

SetIpv6Prefix sets Ipv6Prefix field to given value.

### HasIpv6Prefix

`func (o *PcfBinding) HasIpv6Prefix() bool`

HasIpv6Prefix returns a boolean if a field has been set.

### GetAddIpv6Prefixes

`func (o *PcfBinding) GetAddIpv6Prefixes() []Ipv6Prefix`

GetAddIpv6Prefixes returns the AddIpv6Prefixes field if non-nil, zero value otherwise.

### GetAddIpv6PrefixesOk

`func (o *PcfBinding) GetAddIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetAddIpv6PrefixesOk returns a tuple with the AddIpv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddIpv6Prefixes

`func (o *PcfBinding) SetAddIpv6Prefixes(v []Ipv6Prefix)`

SetAddIpv6Prefixes sets AddIpv6Prefixes field to given value.

### HasAddIpv6Prefixes

`func (o *PcfBinding) HasAddIpv6Prefixes() bool`

HasAddIpv6Prefixes returns a boolean if a field has been set.

### GetIpDomain

`func (o *PcfBinding) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *PcfBinding) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *PcfBinding) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *PcfBinding) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetMacAddr48

`func (o *PcfBinding) GetMacAddr48() string`

GetMacAddr48 returns the MacAddr48 field if non-nil, zero value otherwise.

### GetMacAddr48Ok

`func (o *PcfBinding) GetMacAddr48Ok() (*string, bool)`

GetMacAddr48Ok returns a tuple with the MacAddr48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddr48

`func (o *PcfBinding) SetMacAddr48(v string)`

SetMacAddr48 sets MacAddr48 field to given value.

### HasMacAddr48

`func (o *PcfBinding) HasMacAddr48() bool`

HasMacAddr48 returns a boolean if a field has been set.

### GetAddMacAddrs

`func (o *PcfBinding) GetAddMacAddrs() []string`

GetAddMacAddrs returns the AddMacAddrs field if non-nil, zero value otherwise.

### GetAddMacAddrsOk

`func (o *PcfBinding) GetAddMacAddrsOk() (*[]string, bool)`

GetAddMacAddrsOk returns a tuple with the AddMacAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMacAddrs

`func (o *PcfBinding) SetAddMacAddrs(v []string)`

SetAddMacAddrs sets AddMacAddrs field to given value.

### HasAddMacAddrs

`func (o *PcfBinding) HasAddMacAddrs() bool`

HasAddMacAddrs returns a boolean if a field has been set.

### GetDnn

`func (o *PcfBinding) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PcfBinding) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PcfBinding) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetPcfFqdn

`func (o *PcfBinding) GetPcfFqdn() string`

GetPcfFqdn returns the PcfFqdn field if non-nil, zero value otherwise.

### GetPcfFqdnOk

`func (o *PcfBinding) GetPcfFqdnOk() (*string, bool)`

GetPcfFqdnOk returns a tuple with the PcfFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFqdn

`func (o *PcfBinding) SetPcfFqdn(v string)`

SetPcfFqdn sets PcfFqdn field to given value.

### HasPcfFqdn

`func (o *PcfBinding) HasPcfFqdn() bool`

HasPcfFqdn returns a boolean if a field has been set.

### GetPcfIpEndPoints

`func (o *PcfBinding) GetPcfIpEndPoints() []IpEndPoint`

GetPcfIpEndPoints returns the PcfIpEndPoints field if non-nil, zero value otherwise.

### GetPcfIpEndPointsOk

`func (o *PcfBinding) GetPcfIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfIpEndPoints

`func (o *PcfBinding) SetPcfIpEndPoints(v []IpEndPoint)`

SetPcfIpEndPoints sets PcfIpEndPoints field to given value.

### HasPcfIpEndPoints

`func (o *PcfBinding) HasPcfIpEndPoints() bool`

HasPcfIpEndPoints returns a boolean if a field has been set.

### GetPcfDiamHost

`func (o *PcfBinding) GetPcfDiamHost() string`

GetPcfDiamHost returns the PcfDiamHost field if non-nil, zero value otherwise.

### GetPcfDiamHostOk

`func (o *PcfBinding) GetPcfDiamHostOk() (*string, bool)`

GetPcfDiamHostOk returns a tuple with the PcfDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfDiamHost

`func (o *PcfBinding) SetPcfDiamHost(v string)`

SetPcfDiamHost sets PcfDiamHost field to given value.

### HasPcfDiamHost

`func (o *PcfBinding) HasPcfDiamHost() bool`

HasPcfDiamHost returns a boolean if a field has been set.

### GetPcfDiamRealm

`func (o *PcfBinding) GetPcfDiamRealm() string`

GetPcfDiamRealm returns the PcfDiamRealm field if non-nil, zero value otherwise.

### GetPcfDiamRealmOk

`func (o *PcfBinding) GetPcfDiamRealmOk() (*string, bool)`

GetPcfDiamRealmOk returns a tuple with the PcfDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfDiamRealm

`func (o *PcfBinding) SetPcfDiamRealm(v string)`

SetPcfDiamRealm sets PcfDiamRealm field to given value.

### HasPcfDiamRealm

`func (o *PcfBinding) HasPcfDiamRealm() bool`

HasPcfDiamRealm returns a boolean if a field has been set.

### GetPcfSmFqdn

`func (o *PcfBinding) GetPcfSmFqdn() string`

GetPcfSmFqdn returns the PcfSmFqdn field if non-nil, zero value otherwise.

### GetPcfSmFqdnOk

`func (o *PcfBinding) GetPcfSmFqdnOk() (*string, bool)`

GetPcfSmFqdnOk returns a tuple with the PcfSmFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSmFqdn

`func (o *PcfBinding) SetPcfSmFqdn(v string)`

SetPcfSmFqdn sets PcfSmFqdn field to given value.

### HasPcfSmFqdn

`func (o *PcfBinding) HasPcfSmFqdn() bool`

HasPcfSmFqdn returns a boolean if a field has been set.

### GetPcfSmIpEndPoints

`func (o *PcfBinding) GetPcfSmIpEndPoints() []IpEndPoint`

GetPcfSmIpEndPoints returns the PcfSmIpEndPoints field if non-nil, zero value otherwise.

### GetPcfSmIpEndPointsOk

`func (o *PcfBinding) GetPcfSmIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfSmIpEndPointsOk returns a tuple with the PcfSmIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSmIpEndPoints

`func (o *PcfBinding) SetPcfSmIpEndPoints(v []IpEndPoint)`

SetPcfSmIpEndPoints sets PcfSmIpEndPoints field to given value.

### HasPcfSmIpEndPoints

`func (o *PcfBinding) HasPcfSmIpEndPoints() bool`

HasPcfSmIpEndPoints returns a boolean if a field has been set.

### GetSnssai

`func (o *PcfBinding) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *PcfBinding) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *PcfBinding) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetSuppFeat

`func (o *PcfBinding) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PcfBinding) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PcfBinding) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *PcfBinding) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetPcfId

`func (o *PcfBinding) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *PcfBinding) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *PcfBinding) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *PcfBinding) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *PcfBinding) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *PcfBinding) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *PcfBinding) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *PcfBinding) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *PcfBinding) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *PcfBinding) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *PcfBinding) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *PcfBinding) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetParaCom

`func (o *PcfBinding) GetParaCom() ParameterCombination`

GetParaCom returns the ParaCom field if non-nil, zero value otherwise.

### GetParaComOk

`func (o *PcfBinding) GetParaComOk() (*ParameterCombination, bool)`

GetParaComOk returns a tuple with the ParaCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParaCom

`func (o *PcfBinding) SetParaCom(v ParameterCombination)`

SetParaCom sets ParaCom field to given value.

### HasParaCom

`func (o *PcfBinding) HasParaCom() bool`

HasParaCom returns a boolean if a field has been set.

### GetBindLevel

`func (o *PcfBinding) GetBindLevel() BindingLevel`

GetBindLevel returns the BindLevel field if non-nil, zero value otherwise.

### GetBindLevelOk

`func (o *PcfBinding) GetBindLevelOk() (*BindingLevel, bool)`

GetBindLevelOk returns a tuple with the BindLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindLevel

`func (o *PcfBinding) SetBindLevel(v BindingLevel)`

SetBindLevel sets BindLevel field to given value.

### HasBindLevel

`func (o *PcfBinding) HasBindLevel() bool`

HasBindLevel returns a boolean if a field has been set.

### GetIpv4FrameRouteList

`func (o *PcfBinding) GetIpv4FrameRouteList() []string`

GetIpv4FrameRouteList returns the Ipv4FrameRouteList field if non-nil, zero value otherwise.

### GetIpv4FrameRouteListOk

`func (o *PcfBinding) GetIpv4FrameRouteListOk() (*[]string, bool)`

GetIpv4FrameRouteListOk returns a tuple with the Ipv4FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4FrameRouteList

`func (o *PcfBinding) SetIpv4FrameRouteList(v []string)`

SetIpv4FrameRouteList sets Ipv4FrameRouteList field to given value.

### HasIpv4FrameRouteList

`func (o *PcfBinding) HasIpv4FrameRouteList() bool`

HasIpv4FrameRouteList returns a boolean if a field has been set.

### GetIpv6FrameRouteList

`func (o *PcfBinding) GetIpv6FrameRouteList() []Ipv6Prefix`

GetIpv6FrameRouteList returns the Ipv6FrameRouteList field if non-nil, zero value otherwise.

### GetIpv6FrameRouteListOk

`func (o *PcfBinding) GetIpv6FrameRouteListOk() (*[]Ipv6Prefix, bool)`

GetIpv6FrameRouteListOk returns a tuple with the Ipv6FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6FrameRouteList

`func (o *PcfBinding) SetIpv6FrameRouteList(v []Ipv6Prefix)`

SetIpv6FrameRouteList sets Ipv6FrameRouteList field to given value.

### HasIpv6FrameRouteList

`func (o *PcfBinding) HasIpv6FrameRouteList() bool`

HasIpv6FrameRouteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


