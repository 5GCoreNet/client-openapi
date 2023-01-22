# PduSessionEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvNotif** | [**AfEventNotification**](AfEventNotification.md) |  | 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UeIpv4** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**UeIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**UeMac** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**Status** | Pointer to [**PduSessionStatus**](PduSessionStatus.md) |  | [optional] 
**PcfInfo** | Pointer to [**PcfAddressingInfo**](PcfAddressingInfo.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 

## Methods

### NewPduSessionEventNotification

`func NewPduSessionEventNotification(evNotif AfEventNotification, ) *PduSessionEventNotification`

NewPduSessionEventNotification instantiates a new PduSessionEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionEventNotificationWithDefaults

`func NewPduSessionEventNotificationWithDefaults() *PduSessionEventNotification`

NewPduSessionEventNotificationWithDefaults instantiates a new PduSessionEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvNotif

`func (o *PduSessionEventNotification) GetEvNotif() AfEventNotification`

GetEvNotif returns the EvNotif field if non-nil, zero value otherwise.

### GetEvNotifOk

`func (o *PduSessionEventNotification) GetEvNotifOk() (*AfEventNotification, bool)`

GetEvNotifOk returns a tuple with the EvNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvNotif

`func (o *PduSessionEventNotification) SetEvNotif(v AfEventNotification)`

SetEvNotif sets EvNotif field to given value.


### GetSupi

`func (o *PduSessionEventNotification) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PduSessionEventNotification) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PduSessionEventNotification) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PduSessionEventNotification) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUeIpv4

`func (o *PduSessionEventNotification) GetUeIpv4() string`

GetUeIpv4 returns the UeIpv4 field if non-nil, zero value otherwise.

### GetUeIpv4Ok

`func (o *PduSessionEventNotification) GetUeIpv4Ok() (*string, bool)`

GetUeIpv4Ok returns a tuple with the UeIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4

`func (o *PduSessionEventNotification) SetUeIpv4(v string)`

SetUeIpv4 sets UeIpv4 field to given value.

### HasUeIpv4

`func (o *PduSessionEventNotification) HasUeIpv4() bool`

HasUeIpv4 returns a boolean if a field has been set.

### GetUeIpv6

`func (o *PduSessionEventNotification) GetUeIpv6() Ipv6Addr`

GetUeIpv6 returns the UeIpv6 field if non-nil, zero value otherwise.

### GetUeIpv6Ok

`func (o *PduSessionEventNotification) GetUeIpv6Ok() (*Ipv6Addr, bool)`

GetUeIpv6Ok returns a tuple with the UeIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6

`func (o *PduSessionEventNotification) SetUeIpv6(v Ipv6Addr)`

SetUeIpv6 sets UeIpv6 field to given value.

### HasUeIpv6

`func (o *PduSessionEventNotification) HasUeIpv6() bool`

HasUeIpv6 returns a boolean if a field has been set.

### GetUeMac

`func (o *PduSessionEventNotification) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *PduSessionEventNotification) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *PduSessionEventNotification) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *PduSessionEventNotification) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetStatus

`func (o *PduSessionEventNotification) GetStatus() PduSessionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PduSessionEventNotification) GetStatusOk() (*PduSessionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PduSessionEventNotification) SetStatus(v PduSessionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PduSessionEventNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPcfInfo

`func (o *PduSessionEventNotification) GetPcfInfo() PcfAddressingInfo`

GetPcfInfo returns the PcfInfo field if non-nil, zero value otherwise.

### GetPcfInfoOk

`func (o *PduSessionEventNotification) GetPcfInfoOk() (*PcfAddressingInfo, bool)`

GetPcfInfoOk returns a tuple with the PcfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfInfo

`func (o *PduSessionEventNotification) SetPcfInfo(v PcfAddressingInfo)`

SetPcfInfo sets PcfInfo field to given value.

### HasPcfInfo

`func (o *PduSessionEventNotification) HasPcfInfo() bool`

HasPcfInfo returns a boolean if a field has been set.

### GetDnn

`func (o *PduSessionEventNotification) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionEventNotification) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionEventNotification) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PduSessionEventNotification) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *PduSessionEventNotification) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *PduSessionEventNotification) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *PduSessionEventNotification) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *PduSessionEventNotification) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetGpsi

`func (o *PduSessionEventNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PduSessionEventNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PduSessionEventNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PduSessionEventNotification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


