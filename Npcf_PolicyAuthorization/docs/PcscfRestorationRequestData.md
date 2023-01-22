# PcscfRestorationRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**SliceInfo** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UeIpv4** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**UeIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 

## Methods

### NewPcscfRestorationRequestData

`func NewPcscfRestorationRequestData() *PcscfRestorationRequestData`

NewPcscfRestorationRequestData instantiates a new PcscfRestorationRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcscfRestorationRequestDataWithDefaults

`func NewPcscfRestorationRequestDataWithDefaults() *PcscfRestorationRequestData`

NewPcscfRestorationRequestDataWithDefaults instantiates a new PcscfRestorationRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *PcscfRestorationRequestData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PcscfRestorationRequestData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PcscfRestorationRequestData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PcscfRestorationRequestData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetIpDomain

`func (o *PcscfRestorationRequestData) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *PcscfRestorationRequestData) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *PcscfRestorationRequestData) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *PcscfRestorationRequestData) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetSliceInfo

`func (o *PcscfRestorationRequestData) GetSliceInfo() Snssai`

GetSliceInfo returns the SliceInfo field if non-nil, zero value otherwise.

### GetSliceInfoOk

`func (o *PcscfRestorationRequestData) GetSliceInfoOk() (*Snssai, bool)`

GetSliceInfoOk returns a tuple with the SliceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceInfo

`func (o *PcscfRestorationRequestData) SetSliceInfo(v Snssai)`

SetSliceInfo sets SliceInfo field to given value.

### HasSliceInfo

`func (o *PcscfRestorationRequestData) HasSliceInfo() bool`

HasSliceInfo returns a boolean if a field has been set.

### GetSupi

`func (o *PcscfRestorationRequestData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PcscfRestorationRequestData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PcscfRestorationRequestData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PcscfRestorationRequestData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUeIpv4

`func (o *PcscfRestorationRequestData) GetUeIpv4() string`

GetUeIpv4 returns the UeIpv4 field if non-nil, zero value otherwise.

### GetUeIpv4Ok

`func (o *PcscfRestorationRequestData) GetUeIpv4Ok() (*string, bool)`

GetUeIpv4Ok returns a tuple with the UeIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4

`func (o *PcscfRestorationRequestData) SetUeIpv4(v string)`

SetUeIpv4 sets UeIpv4 field to given value.

### HasUeIpv4

`func (o *PcscfRestorationRequestData) HasUeIpv4() bool`

HasUeIpv4 returns a boolean if a field has been set.

### GetUeIpv6

`func (o *PcscfRestorationRequestData) GetUeIpv6() Ipv6Addr`

GetUeIpv6 returns the UeIpv6 field if non-nil, zero value otherwise.

### GetUeIpv6Ok

`func (o *PcscfRestorationRequestData) GetUeIpv6Ok() (*Ipv6Addr, bool)`

GetUeIpv6Ok returns a tuple with the UeIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6

`func (o *PcscfRestorationRequestData) SetUeIpv6(v Ipv6Addr)`

SetUeIpv6 sets UeIpv6 field to given value.

### HasUeIpv6

`func (o *PcscfRestorationRequestData) HasUeIpv6() bool`

HasUeIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


