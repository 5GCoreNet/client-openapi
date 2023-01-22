# EventNotification1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**SmfEvent**](SmfEvent.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**TransacInfos** | Pointer to [**[]TransactionInfo**](TransactionInfo.md) | Transaction Information. | [optional] 
**SourceDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**TargetDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**DnaiChgType** | Pointer to [**DnaiChangeType**](DnaiChangeType.md) |  | [optional] 
**SourceUeIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**SourceUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**TargetUeIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**TargetUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**SourceTraRouting** | Pointer to [**NullableRouteToLocation**](RouteToLocation.md) |  | [optional] 
**TargetTraRouting** | Pointer to [**NullableRouteToLocation**](RouteToLocation.md) |  | [optional] 
**UeMac** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**AdIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**AdIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**ReIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**ReIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AccType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**PduSeId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**DddStatus** | Pointer to [**DlDataDeliveryStatus**](DlDataDeliveryStatus.md) |  | [optional] 
**DddTraDescriptor** | Pointer to [**DddTrafficDescriptor**](DddTrafficDescriptor.md) |  | [optional] 
**MaxWaitTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**CommFailure** | Pointer to [**CommunicationFailure**](CommunicationFailure.md) |  | [optional] 
**Ipv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**Ipv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**Ipv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PduSessType** | Pointer to [**PduSessionType**](PduSessionType.md) |  | [optional] 
**Qfi** | Pointer to **int32** | Unsigned integer identifying a QoS flow, within the range 0 to 63. | [optional] 
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**EthFlowDescs** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Descriptor(s) for non-IP traffic. It allows the encoding of multiple UL and/or DL flows. Each entry of the array describes a single Ethernet flow.  | [optional] 
**EthfDescs** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Contains the UL and/or DL Ethernet flows. Each entry of the array describes a single Ethernet flow.  | [optional] 
**FlowDescs** | Pointer to **[]string** | Descriptor(s) for IP traffic. It allows the encoding of multiple UL and/or DL flows. Each entry of the array describes a single IP flow.  | [optional] 
**FDescs** | Pointer to **[]string** | Contains the UL and/or DL IP flows. Each entry of the array describes a single IP flow.  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**UlDelays** | Pointer to **[]int32** |  | [optional] 
**DlDelays** | Pointer to **[]int32** |  | [optional] 
**RtDelays** | Pointer to **[]int32** |  | [optional] 
**TimeWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**SmNasFromUe** | Pointer to [**SmNasFromUe**](SmNasFromUe.md) |  | [optional] 
**SmNasFromSmf** | Pointer to [**SmNasFromSmf**](SmNasFromSmf.md) |  | [optional] 
**UpRedTrans** | Pointer to **bool** | Indicates whether the redundant transmission is setup or terminated. Set to \&quot;true\&quot; if  the redundant transmission is setup, otherwise set to \&quot;false\&quot; if the redundant  transmission is terminated. Default value is set to \&quot;false\&quot;.  | [optional] 
**SsId** | Pointer to **string** |  | [optional] 
**BssId** | Pointer to **string** |  | [optional] 
**StartWlan** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**EndWlan** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PduSessInfos** | Pointer to [**[]PduSessionInformation**](PduSessionInformation.md) |  | [optional] 
**UpfInfo** | Pointer to [**UpfInformation**](UpfInformation.md) |  | [optional] 

## Methods

### NewEventNotification1

`func NewEventNotification1(event SmfEvent, timeStamp time.Time, ) *EventNotification1`

NewEventNotification1 instantiates a new EventNotification1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotification1WithDefaults

`func NewEventNotification1WithDefaults() *EventNotification1`

NewEventNotification1WithDefaults instantiates a new EventNotification1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventNotification1) GetEvent() SmfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventNotification1) GetEventOk() (*SmfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventNotification1) SetEvent(v SmfEvent)`

SetEvent sets Event field to given value.


### GetTimeStamp

`func (o *EventNotification1) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *EventNotification1) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *EventNotification1) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSupi

`func (o *EventNotification1) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *EventNotification1) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *EventNotification1) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *EventNotification1) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *EventNotification1) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *EventNotification1) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *EventNotification1) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *EventNotification1) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *EventNotification1) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *EventNotification1) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *EventNotification1) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *EventNotification1) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.

### GetTransacInfos

`func (o *EventNotification1) GetTransacInfos() []TransactionInfo`

GetTransacInfos returns the TransacInfos field if non-nil, zero value otherwise.

### GetTransacInfosOk

`func (o *EventNotification1) GetTransacInfosOk() (*[]TransactionInfo, bool)`

GetTransacInfosOk returns a tuple with the TransacInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransacInfos

`func (o *EventNotification1) SetTransacInfos(v []TransactionInfo)`

SetTransacInfos sets TransacInfos field to given value.

### HasTransacInfos

`func (o *EventNotification1) HasTransacInfos() bool`

HasTransacInfos returns a boolean if a field has been set.

### GetSourceDnai

`func (o *EventNotification1) GetSourceDnai() string`

GetSourceDnai returns the SourceDnai field if non-nil, zero value otherwise.

### GetSourceDnaiOk

`func (o *EventNotification1) GetSourceDnaiOk() (*string, bool)`

GetSourceDnaiOk returns a tuple with the SourceDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDnai

`func (o *EventNotification1) SetSourceDnai(v string)`

SetSourceDnai sets SourceDnai field to given value.

### HasSourceDnai

`func (o *EventNotification1) HasSourceDnai() bool`

HasSourceDnai returns a boolean if a field has been set.

### GetTargetDnai

`func (o *EventNotification1) GetTargetDnai() string`

GetTargetDnai returns the TargetDnai field if non-nil, zero value otherwise.

### GetTargetDnaiOk

`func (o *EventNotification1) GetTargetDnaiOk() (*string, bool)`

GetTargetDnaiOk returns a tuple with the TargetDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnai

`func (o *EventNotification1) SetTargetDnai(v string)`

SetTargetDnai sets TargetDnai field to given value.

### HasTargetDnai

`func (o *EventNotification1) HasTargetDnai() bool`

HasTargetDnai returns a boolean if a field has been set.

### GetDnaiChgType

`func (o *EventNotification1) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *EventNotification1) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *EventNotification1) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.

### HasDnaiChgType

`func (o *EventNotification1) HasDnaiChgType() bool`

HasDnaiChgType returns a boolean if a field has been set.

### GetSourceUeIpv4Addr

`func (o *EventNotification1) GetSourceUeIpv4Addr() string`

GetSourceUeIpv4Addr returns the SourceUeIpv4Addr field if non-nil, zero value otherwise.

### GetSourceUeIpv4AddrOk

`func (o *EventNotification1) GetSourceUeIpv4AddrOk() (*string, bool)`

GetSourceUeIpv4AddrOk returns a tuple with the SourceUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUeIpv4Addr

`func (o *EventNotification1) SetSourceUeIpv4Addr(v string)`

SetSourceUeIpv4Addr sets SourceUeIpv4Addr field to given value.

### HasSourceUeIpv4Addr

`func (o *EventNotification1) HasSourceUeIpv4Addr() bool`

HasSourceUeIpv4Addr returns a boolean if a field has been set.

### GetSourceUeIpv6Prefix

`func (o *EventNotification1) GetSourceUeIpv6Prefix() Ipv6Prefix`

GetSourceUeIpv6Prefix returns the SourceUeIpv6Prefix field if non-nil, zero value otherwise.

### GetSourceUeIpv6PrefixOk

`func (o *EventNotification1) GetSourceUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetSourceUeIpv6PrefixOk returns a tuple with the SourceUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUeIpv6Prefix

`func (o *EventNotification1) SetSourceUeIpv6Prefix(v Ipv6Prefix)`

SetSourceUeIpv6Prefix sets SourceUeIpv6Prefix field to given value.

### HasSourceUeIpv6Prefix

`func (o *EventNotification1) HasSourceUeIpv6Prefix() bool`

HasSourceUeIpv6Prefix returns a boolean if a field has been set.

### GetTargetUeIpv4Addr

`func (o *EventNotification1) GetTargetUeIpv4Addr() string`

GetTargetUeIpv4Addr returns the TargetUeIpv4Addr field if non-nil, zero value otherwise.

### GetTargetUeIpv4AddrOk

`func (o *EventNotification1) GetTargetUeIpv4AddrOk() (*string, bool)`

GetTargetUeIpv4AddrOk returns a tuple with the TargetUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUeIpv4Addr

`func (o *EventNotification1) SetTargetUeIpv4Addr(v string)`

SetTargetUeIpv4Addr sets TargetUeIpv4Addr field to given value.

### HasTargetUeIpv4Addr

`func (o *EventNotification1) HasTargetUeIpv4Addr() bool`

HasTargetUeIpv4Addr returns a boolean if a field has been set.

### GetTargetUeIpv6Prefix

`func (o *EventNotification1) GetTargetUeIpv6Prefix() Ipv6Prefix`

GetTargetUeIpv6Prefix returns the TargetUeIpv6Prefix field if non-nil, zero value otherwise.

### GetTargetUeIpv6PrefixOk

`func (o *EventNotification1) GetTargetUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetTargetUeIpv6PrefixOk returns a tuple with the TargetUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUeIpv6Prefix

`func (o *EventNotification1) SetTargetUeIpv6Prefix(v Ipv6Prefix)`

SetTargetUeIpv6Prefix sets TargetUeIpv6Prefix field to given value.

### HasTargetUeIpv6Prefix

`func (o *EventNotification1) HasTargetUeIpv6Prefix() bool`

HasTargetUeIpv6Prefix returns a boolean if a field has been set.

### GetSourceTraRouting

`func (o *EventNotification1) GetSourceTraRouting() RouteToLocation`

GetSourceTraRouting returns the SourceTraRouting field if non-nil, zero value otherwise.

### GetSourceTraRoutingOk

`func (o *EventNotification1) GetSourceTraRoutingOk() (*RouteToLocation, bool)`

GetSourceTraRoutingOk returns a tuple with the SourceTraRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTraRouting

`func (o *EventNotification1) SetSourceTraRouting(v RouteToLocation)`

SetSourceTraRouting sets SourceTraRouting field to given value.

### HasSourceTraRouting

`func (o *EventNotification1) HasSourceTraRouting() bool`

HasSourceTraRouting returns a boolean if a field has been set.

### SetSourceTraRoutingNil

`func (o *EventNotification1) SetSourceTraRoutingNil(b bool)`

 SetSourceTraRoutingNil sets the value for SourceTraRouting to be an explicit nil

### UnsetSourceTraRouting
`func (o *EventNotification1) UnsetSourceTraRouting()`

UnsetSourceTraRouting ensures that no value is present for SourceTraRouting, not even an explicit nil
### GetTargetTraRouting

`func (o *EventNotification1) GetTargetTraRouting() RouteToLocation`

GetTargetTraRouting returns the TargetTraRouting field if non-nil, zero value otherwise.

### GetTargetTraRoutingOk

`func (o *EventNotification1) GetTargetTraRoutingOk() (*RouteToLocation, bool)`

GetTargetTraRoutingOk returns a tuple with the TargetTraRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTraRouting

`func (o *EventNotification1) SetTargetTraRouting(v RouteToLocation)`

SetTargetTraRouting sets TargetTraRouting field to given value.

### HasTargetTraRouting

`func (o *EventNotification1) HasTargetTraRouting() bool`

HasTargetTraRouting returns a boolean if a field has been set.

### SetTargetTraRoutingNil

`func (o *EventNotification1) SetTargetTraRoutingNil(b bool)`

 SetTargetTraRoutingNil sets the value for TargetTraRouting to be an explicit nil

### UnsetTargetTraRouting
`func (o *EventNotification1) UnsetTargetTraRouting()`

UnsetTargetTraRouting ensures that no value is present for TargetTraRouting, not even an explicit nil
### GetUeMac

`func (o *EventNotification1) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *EventNotification1) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *EventNotification1) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *EventNotification1) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetAdIpv4Addr

`func (o *EventNotification1) GetAdIpv4Addr() string`

GetAdIpv4Addr returns the AdIpv4Addr field if non-nil, zero value otherwise.

### GetAdIpv4AddrOk

`func (o *EventNotification1) GetAdIpv4AddrOk() (*string, bool)`

GetAdIpv4AddrOk returns a tuple with the AdIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdIpv4Addr

`func (o *EventNotification1) SetAdIpv4Addr(v string)`

SetAdIpv4Addr sets AdIpv4Addr field to given value.

### HasAdIpv4Addr

`func (o *EventNotification1) HasAdIpv4Addr() bool`

HasAdIpv4Addr returns a boolean if a field has been set.

### GetAdIpv6Prefix

`func (o *EventNotification1) GetAdIpv6Prefix() Ipv6Prefix`

GetAdIpv6Prefix returns the AdIpv6Prefix field if non-nil, zero value otherwise.

### GetAdIpv6PrefixOk

`func (o *EventNotification1) GetAdIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetAdIpv6PrefixOk returns a tuple with the AdIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdIpv6Prefix

`func (o *EventNotification1) SetAdIpv6Prefix(v Ipv6Prefix)`

SetAdIpv6Prefix sets AdIpv6Prefix field to given value.

### HasAdIpv6Prefix

`func (o *EventNotification1) HasAdIpv6Prefix() bool`

HasAdIpv6Prefix returns a boolean if a field has been set.

### GetReIpv4Addr

`func (o *EventNotification1) GetReIpv4Addr() string`

GetReIpv4Addr returns the ReIpv4Addr field if non-nil, zero value otherwise.

### GetReIpv4AddrOk

`func (o *EventNotification1) GetReIpv4AddrOk() (*string, bool)`

GetReIpv4AddrOk returns a tuple with the ReIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReIpv4Addr

`func (o *EventNotification1) SetReIpv4Addr(v string)`

SetReIpv4Addr sets ReIpv4Addr field to given value.

### HasReIpv4Addr

`func (o *EventNotification1) HasReIpv4Addr() bool`

HasReIpv4Addr returns a boolean if a field has been set.

### GetReIpv6Prefix

`func (o *EventNotification1) GetReIpv6Prefix() Ipv6Prefix`

GetReIpv6Prefix returns the ReIpv6Prefix field if non-nil, zero value otherwise.

### GetReIpv6PrefixOk

`func (o *EventNotification1) GetReIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetReIpv6PrefixOk returns a tuple with the ReIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReIpv6Prefix

`func (o *EventNotification1) SetReIpv6Prefix(v Ipv6Prefix)`

SetReIpv6Prefix sets ReIpv6Prefix field to given value.

### HasReIpv6Prefix

`func (o *EventNotification1) HasReIpv6Prefix() bool`

HasReIpv6Prefix returns a boolean if a field has been set.

### GetPlmnId

`func (o *EventNotification1) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *EventNotification1) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *EventNotification1) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *EventNotification1) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetAccType

`func (o *EventNotification1) GetAccType() AccessType`

GetAccType returns the AccType field if non-nil, zero value otherwise.

### GetAccTypeOk

`func (o *EventNotification1) GetAccTypeOk() (*AccessType, bool)`

GetAccTypeOk returns a tuple with the AccType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccType

`func (o *EventNotification1) SetAccType(v AccessType)`

SetAccType sets AccType field to given value.

### HasAccType

`func (o *EventNotification1) HasAccType() bool`

HasAccType returns a boolean if a field has been set.

### GetPduSeId

`func (o *EventNotification1) GetPduSeId() int32`

GetPduSeId returns the PduSeId field if non-nil, zero value otherwise.

### GetPduSeIdOk

`func (o *EventNotification1) GetPduSeIdOk() (*int32, bool)`

GetPduSeIdOk returns a tuple with the PduSeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSeId

`func (o *EventNotification1) SetPduSeId(v int32)`

SetPduSeId sets PduSeId field to given value.

### HasPduSeId

`func (o *EventNotification1) HasPduSeId() bool`

HasPduSeId returns a boolean if a field has been set.

### GetRatType

`func (o *EventNotification1) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *EventNotification1) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *EventNotification1) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *EventNotification1) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetDddStatus

`func (o *EventNotification1) GetDddStatus() DlDataDeliveryStatus`

GetDddStatus returns the DddStatus field if non-nil, zero value otherwise.

### GetDddStatusOk

`func (o *EventNotification1) GetDddStatusOk() (*DlDataDeliveryStatus, bool)`

GetDddStatusOk returns a tuple with the DddStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddStatus

`func (o *EventNotification1) SetDddStatus(v DlDataDeliveryStatus)`

SetDddStatus sets DddStatus field to given value.

### HasDddStatus

`func (o *EventNotification1) HasDddStatus() bool`

HasDddStatus returns a boolean if a field has been set.

### GetDddTraDescriptor

`func (o *EventNotification1) GetDddTraDescriptor() DddTrafficDescriptor`

GetDddTraDescriptor returns the DddTraDescriptor field if non-nil, zero value otherwise.

### GetDddTraDescriptorOk

`func (o *EventNotification1) GetDddTraDescriptorOk() (*DddTrafficDescriptor, bool)`

GetDddTraDescriptorOk returns a tuple with the DddTraDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddTraDescriptor

`func (o *EventNotification1) SetDddTraDescriptor(v DddTrafficDescriptor)`

SetDddTraDescriptor sets DddTraDescriptor field to given value.

### HasDddTraDescriptor

`func (o *EventNotification1) HasDddTraDescriptor() bool`

HasDddTraDescriptor returns a boolean if a field has been set.

### GetMaxWaitTime

`func (o *EventNotification1) GetMaxWaitTime() time.Time`

GetMaxWaitTime returns the MaxWaitTime field if non-nil, zero value otherwise.

### GetMaxWaitTimeOk

`func (o *EventNotification1) GetMaxWaitTimeOk() (*time.Time, bool)`

GetMaxWaitTimeOk returns a tuple with the MaxWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitTime

`func (o *EventNotification1) SetMaxWaitTime(v time.Time)`

SetMaxWaitTime sets MaxWaitTime field to given value.

### HasMaxWaitTime

`func (o *EventNotification1) HasMaxWaitTime() bool`

HasMaxWaitTime returns a boolean if a field has been set.

### GetCommFailure

`func (o *EventNotification1) GetCommFailure() CommunicationFailure`

GetCommFailure returns the CommFailure field if non-nil, zero value otherwise.

### GetCommFailureOk

`func (o *EventNotification1) GetCommFailureOk() (*CommunicationFailure, bool)`

GetCommFailureOk returns a tuple with the CommFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommFailure

`func (o *EventNotification1) SetCommFailure(v CommunicationFailure)`

SetCommFailure sets CommFailure field to given value.

### HasCommFailure

`func (o *EventNotification1) HasCommFailure() bool`

HasCommFailure returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *EventNotification1) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *EventNotification1) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *EventNotification1) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *EventNotification1) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Prefixes

`func (o *EventNotification1) GetIpv6Prefixes() []Ipv6Prefix`

GetIpv6Prefixes returns the Ipv6Prefixes field if non-nil, zero value otherwise.

### GetIpv6PrefixesOk

`func (o *EventNotification1) GetIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefixes

`func (o *EventNotification1) SetIpv6Prefixes(v []Ipv6Prefix)`

SetIpv6Prefixes sets Ipv6Prefixes field to given value.

### HasIpv6Prefixes

`func (o *EventNotification1) HasIpv6Prefixes() bool`

HasIpv6Prefixes returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *EventNotification1) GetIpv6Addrs() []Ipv6Addr`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *EventNotification1) GetIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *EventNotification1) SetIpv6Addrs(v []Ipv6Addr)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *EventNotification1) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.

### GetPduSessType

`func (o *EventNotification1) GetPduSessType() PduSessionType`

GetPduSessType returns the PduSessType field if non-nil, zero value otherwise.

### GetPduSessTypeOk

`func (o *EventNotification1) GetPduSessTypeOk() (*PduSessionType, bool)`

GetPduSessTypeOk returns a tuple with the PduSessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessType

`func (o *EventNotification1) SetPduSessType(v PduSessionType)`

SetPduSessType sets PduSessType field to given value.

### HasPduSessType

`func (o *EventNotification1) HasPduSessType() bool`

HasPduSessType returns a boolean if a field has been set.

### GetQfi

`func (o *EventNotification1) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *EventNotification1) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *EventNotification1) SetQfi(v int32)`

SetQfi sets Qfi field to given value.

### HasQfi

`func (o *EventNotification1) HasQfi() bool`

HasQfi returns a boolean if a field has been set.

### GetAppId

`func (o *EventNotification1) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EventNotification1) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EventNotification1) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *EventNotification1) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetEthFlowDescs

`func (o *EventNotification1) GetEthFlowDescs() []EthFlowDescription`

GetEthFlowDescs returns the EthFlowDescs field if non-nil, zero value otherwise.

### GetEthFlowDescsOk

`func (o *EventNotification1) GetEthFlowDescsOk() (*[]EthFlowDescription, bool)`

GetEthFlowDescsOk returns a tuple with the EthFlowDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowDescs

`func (o *EventNotification1) SetEthFlowDescs(v []EthFlowDescription)`

SetEthFlowDescs sets EthFlowDescs field to given value.

### HasEthFlowDescs

`func (o *EventNotification1) HasEthFlowDescs() bool`

HasEthFlowDescs returns a boolean if a field has been set.

### GetEthfDescs

`func (o *EventNotification1) GetEthfDescs() []EthFlowDescription`

GetEthfDescs returns the EthfDescs field if non-nil, zero value otherwise.

### GetEthfDescsOk

`func (o *EventNotification1) GetEthfDescsOk() (*[]EthFlowDescription, bool)`

GetEthfDescsOk returns a tuple with the EthfDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthfDescs

`func (o *EventNotification1) SetEthfDescs(v []EthFlowDescription)`

SetEthfDescs sets EthfDescs field to given value.

### HasEthfDescs

`func (o *EventNotification1) HasEthfDescs() bool`

HasEthfDescs returns a boolean if a field has been set.

### GetFlowDescs

`func (o *EventNotification1) GetFlowDescs() []string`

GetFlowDescs returns the FlowDescs field if non-nil, zero value otherwise.

### GetFlowDescsOk

`func (o *EventNotification1) GetFlowDescsOk() (*[]string, bool)`

GetFlowDescsOk returns a tuple with the FlowDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescs

`func (o *EventNotification1) SetFlowDescs(v []string)`

SetFlowDescs sets FlowDescs field to given value.

### HasFlowDescs

`func (o *EventNotification1) HasFlowDescs() bool`

HasFlowDescs returns a boolean if a field has been set.

### GetFDescs

`func (o *EventNotification1) GetFDescs() []string`

GetFDescs returns the FDescs field if non-nil, zero value otherwise.

### GetFDescsOk

`func (o *EventNotification1) GetFDescsOk() (*[]string, bool)`

GetFDescsOk returns a tuple with the FDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDescs

`func (o *EventNotification1) SetFDescs(v []string)`

SetFDescs sets FDescs field to given value.

### HasFDescs

`func (o *EventNotification1) HasFDescs() bool`

HasFDescs returns a boolean if a field has been set.

### GetDnn

`func (o *EventNotification1) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *EventNotification1) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *EventNotification1) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *EventNotification1) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *EventNotification1) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *EventNotification1) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *EventNotification1) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *EventNotification1) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetUlDelays

`func (o *EventNotification1) GetUlDelays() []int32`

GetUlDelays returns the UlDelays field if non-nil, zero value otherwise.

### GetUlDelaysOk

`func (o *EventNotification1) GetUlDelaysOk() (*[]int32, bool)`

GetUlDelaysOk returns a tuple with the UlDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlDelays

`func (o *EventNotification1) SetUlDelays(v []int32)`

SetUlDelays sets UlDelays field to given value.

### HasUlDelays

`func (o *EventNotification1) HasUlDelays() bool`

HasUlDelays returns a boolean if a field has been set.

### GetDlDelays

`func (o *EventNotification1) GetDlDelays() []int32`

GetDlDelays returns the DlDelays field if non-nil, zero value otherwise.

### GetDlDelaysOk

`func (o *EventNotification1) GetDlDelaysOk() (*[]int32, bool)`

GetDlDelaysOk returns a tuple with the DlDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlDelays

`func (o *EventNotification1) SetDlDelays(v []int32)`

SetDlDelays sets DlDelays field to given value.

### HasDlDelays

`func (o *EventNotification1) HasDlDelays() bool`

HasDlDelays returns a boolean if a field has been set.

### GetRtDelays

`func (o *EventNotification1) GetRtDelays() []int32`

GetRtDelays returns the RtDelays field if non-nil, zero value otherwise.

### GetRtDelaysOk

`func (o *EventNotification1) GetRtDelaysOk() (*[]int32, bool)`

GetRtDelaysOk returns a tuple with the RtDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtDelays

`func (o *EventNotification1) SetRtDelays(v []int32)`

SetRtDelays sets RtDelays field to given value.

### HasRtDelays

`func (o *EventNotification1) HasRtDelays() bool`

HasRtDelays returns a boolean if a field has been set.

### GetTimeWindow

`func (o *EventNotification1) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *EventNotification1) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *EventNotification1) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *EventNotification1) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetSmNasFromUe

`func (o *EventNotification1) GetSmNasFromUe() SmNasFromUe`

GetSmNasFromUe returns the SmNasFromUe field if non-nil, zero value otherwise.

### GetSmNasFromUeOk

`func (o *EventNotification1) GetSmNasFromUeOk() (*SmNasFromUe, bool)`

GetSmNasFromUeOk returns a tuple with the SmNasFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmNasFromUe

`func (o *EventNotification1) SetSmNasFromUe(v SmNasFromUe)`

SetSmNasFromUe sets SmNasFromUe field to given value.

### HasSmNasFromUe

`func (o *EventNotification1) HasSmNasFromUe() bool`

HasSmNasFromUe returns a boolean if a field has been set.

### GetSmNasFromSmf

`func (o *EventNotification1) GetSmNasFromSmf() SmNasFromSmf`

GetSmNasFromSmf returns the SmNasFromSmf field if non-nil, zero value otherwise.

### GetSmNasFromSmfOk

`func (o *EventNotification1) GetSmNasFromSmfOk() (*SmNasFromSmf, bool)`

GetSmNasFromSmfOk returns a tuple with the SmNasFromSmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmNasFromSmf

`func (o *EventNotification1) SetSmNasFromSmf(v SmNasFromSmf)`

SetSmNasFromSmf sets SmNasFromSmf field to given value.

### HasSmNasFromSmf

`func (o *EventNotification1) HasSmNasFromSmf() bool`

HasSmNasFromSmf returns a boolean if a field has been set.

### GetUpRedTrans

`func (o *EventNotification1) GetUpRedTrans() bool`

GetUpRedTrans returns the UpRedTrans field if non-nil, zero value otherwise.

### GetUpRedTransOk

`func (o *EventNotification1) GetUpRedTransOk() (*bool, bool)`

GetUpRedTransOk returns a tuple with the UpRedTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRedTrans

`func (o *EventNotification1) SetUpRedTrans(v bool)`

SetUpRedTrans sets UpRedTrans field to given value.

### HasUpRedTrans

`func (o *EventNotification1) HasUpRedTrans() bool`

HasUpRedTrans returns a boolean if a field has been set.

### GetSsId

`func (o *EventNotification1) GetSsId() string`

GetSsId returns the SsId field if non-nil, zero value otherwise.

### GetSsIdOk

`func (o *EventNotification1) GetSsIdOk() (*string, bool)`

GetSsIdOk returns a tuple with the SsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsId

`func (o *EventNotification1) SetSsId(v string)`

SetSsId sets SsId field to given value.

### HasSsId

`func (o *EventNotification1) HasSsId() bool`

HasSsId returns a boolean if a field has been set.

### GetBssId

`func (o *EventNotification1) GetBssId() string`

GetBssId returns the BssId field if non-nil, zero value otherwise.

### GetBssIdOk

`func (o *EventNotification1) GetBssIdOk() (*string, bool)`

GetBssIdOk returns a tuple with the BssId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssId

`func (o *EventNotification1) SetBssId(v string)`

SetBssId sets BssId field to given value.

### HasBssId

`func (o *EventNotification1) HasBssId() bool`

HasBssId returns a boolean if a field has been set.

### GetStartWlan

`func (o *EventNotification1) GetStartWlan() time.Time`

GetStartWlan returns the StartWlan field if non-nil, zero value otherwise.

### GetStartWlanOk

`func (o *EventNotification1) GetStartWlanOk() (*time.Time, bool)`

GetStartWlanOk returns a tuple with the StartWlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWlan

`func (o *EventNotification1) SetStartWlan(v time.Time)`

SetStartWlan sets StartWlan field to given value.

### HasStartWlan

`func (o *EventNotification1) HasStartWlan() bool`

HasStartWlan returns a boolean if a field has been set.

### GetEndWlan

`func (o *EventNotification1) GetEndWlan() time.Time`

GetEndWlan returns the EndWlan field if non-nil, zero value otherwise.

### GetEndWlanOk

`func (o *EventNotification1) GetEndWlanOk() (*time.Time, bool)`

GetEndWlanOk returns a tuple with the EndWlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndWlan

`func (o *EventNotification1) SetEndWlan(v time.Time)`

SetEndWlan sets EndWlan field to given value.

### HasEndWlan

`func (o *EventNotification1) HasEndWlan() bool`

HasEndWlan returns a boolean if a field has been set.

### GetPduSessInfos

`func (o *EventNotification1) GetPduSessInfos() []PduSessionInformation`

GetPduSessInfos returns the PduSessInfos field if non-nil, zero value otherwise.

### GetPduSessInfosOk

`func (o *EventNotification1) GetPduSessInfosOk() (*[]PduSessionInformation, bool)`

GetPduSessInfosOk returns a tuple with the PduSessInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessInfos

`func (o *EventNotification1) SetPduSessInfos(v []PduSessionInformation)`

SetPduSessInfos sets PduSessInfos field to given value.

### HasPduSessInfos

`func (o *EventNotification1) HasPduSessInfos() bool`

HasPduSessInfos returns a boolean if a field has been set.

### GetUpfInfo

`func (o *EventNotification1) GetUpfInfo() UpfInformation`

GetUpfInfo returns the UpfInfo field if non-nil, zero value otherwise.

### GetUpfInfoOk

`func (o *EventNotification1) GetUpfInfoOk() (*UpfInformation, bool)`

GetUpfInfoOk returns a tuple with the UpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfo

`func (o *EventNotification1) SetUpfInfo(v UpfInformation)`

SetUpfInfo sets UpfInfo field to given value.

### HasUpfInfo

`func (o *EventNotification1) HasUpfInfo() bool`

HasUpfInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


