# PduSessionTsnBridge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsnBridgeInfo** | [**TsnBridgeInfo**](TsnBridgeInfo.md) |  | 
**TsnBridgeManCont** | Pointer to [**BridgeManagementContainer**](BridgeManagementContainer.md) |  | [optional] 
**TsnPortManContDstt** | Pointer to [**PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**TsnPortManContNwtts** | Pointer to [**[]PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**UeIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**IpDomain** | Pointer to **string** | IPv4 address domain identifier. | [optional] 
**UeIpv6AddrPrefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 

## Methods

### NewPduSessionTsnBridge

`func NewPduSessionTsnBridge(tsnBridgeInfo TsnBridgeInfo, ) *PduSessionTsnBridge`

NewPduSessionTsnBridge instantiates a new PduSessionTsnBridge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionTsnBridgeWithDefaults

`func NewPduSessionTsnBridgeWithDefaults() *PduSessionTsnBridge`

NewPduSessionTsnBridgeWithDefaults instantiates a new PduSessionTsnBridge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsnBridgeInfo

`func (o *PduSessionTsnBridge) GetTsnBridgeInfo() TsnBridgeInfo`

GetTsnBridgeInfo returns the TsnBridgeInfo field if non-nil, zero value otherwise.

### GetTsnBridgeInfoOk

`func (o *PduSessionTsnBridge) GetTsnBridgeInfoOk() (*TsnBridgeInfo, bool)`

GetTsnBridgeInfoOk returns a tuple with the TsnBridgeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeInfo

`func (o *PduSessionTsnBridge) SetTsnBridgeInfo(v TsnBridgeInfo)`

SetTsnBridgeInfo sets TsnBridgeInfo field to given value.


### GetTsnBridgeManCont

`func (o *PduSessionTsnBridge) GetTsnBridgeManCont() BridgeManagementContainer`

GetTsnBridgeManCont returns the TsnBridgeManCont field if non-nil, zero value otherwise.

### GetTsnBridgeManContOk

`func (o *PduSessionTsnBridge) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool)`

GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeManCont

`func (o *PduSessionTsnBridge) SetTsnBridgeManCont(v BridgeManagementContainer)`

SetTsnBridgeManCont sets TsnBridgeManCont field to given value.

### HasTsnBridgeManCont

`func (o *PduSessionTsnBridge) HasTsnBridgeManCont() bool`

HasTsnBridgeManCont returns a boolean if a field has been set.

### GetTsnPortManContDstt

`func (o *PduSessionTsnBridge) GetTsnPortManContDstt() PortManagementContainer`

GetTsnPortManContDstt returns the TsnPortManContDstt field if non-nil, zero value otherwise.

### GetTsnPortManContDsttOk

`func (o *PduSessionTsnBridge) GetTsnPortManContDsttOk() (*PortManagementContainer, bool)`

GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContDstt

`func (o *PduSessionTsnBridge) SetTsnPortManContDstt(v PortManagementContainer)`

SetTsnPortManContDstt sets TsnPortManContDstt field to given value.

### HasTsnPortManContDstt

`func (o *PduSessionTsnBridge) HasTsnPortManContDstt() bool`

HasTsnPortManContDstt returns a boolean if a field has been set.

### GetTsnPortManContNwtts

`func (o *PduSessionTsnBridge) GetTsnPortManContNwtts() []PortManagementContainer`

GetTsnPortManContNwtts returns the TsnPortManContNwtts field if non-nil, zero value otherwise.

### GetTsnPortManContNwttsOk

`func (o *PduSessionTsnBridge) GetTsnPortManContNwttsOk() (*[]PortManagementContainer, bool)`

GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContNwtts

`func (o *PduSessionTsnBridge) SetTsnPortManContNwtts(v []PortManagementContainer)`

SetTsnPortManContNwtts sets TsnPortManContNwtts field to given value.

### HasTsnPortManContNwtts

`func (o *PduSessionTsnBridge) HasTsnPortManContNwtts() bool`

HasTsnPortManContNwtts returns a boolean if a field has been set.

### GetUeIpv4Addr

`func (o *PduSessionTsnBridge) GetUeIpv4Addr() string`

GetUeIpv4Addr returns the UeIpv4Addr field if non-nil, zero value otherwise.

### GetUeIpv4AddrOk

`func (o *PduSessionTsnBridge) GetUeIpv4AddrOk() (*string, bool)`

GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Addr

`func (o *PduSessionTsnBridge) SetUeIpv4Addr(v string)`

SetUeIpv4Addr sets UeIpv4Addr field to given value.

### HasUeIpv4Addr

`func (o *PduSessionTsnBridge) HasUeIpv4Addr() bool`

HasUeIpv4Addr returns a boolean if a field has been set.

### GetDnn

`func (o *PduSessionTsnBridge) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionTsnBridge) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionTsnBridge) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PduSessionTsnBridge) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *PduSessionTsnBridge) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *PduSessionTsnBridge) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *PduSessionTsnBridge) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *PduSessionTsnBridge) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetIpDomain

`func (o *PduSessionTsnBridge) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *PduSessionTsnBridge) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *PduSessionTsnBridge) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *PduSessionTsnBridge) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetUeIpv6AddrPrefix

`func (o *PduSessionTsnBridge) GetUeIpv6AddrPrefix() Ipv6Prefix`

GetUeIpv6AddrPrefix returns the UeIpv6AddrPrefix field if non-nil, zero value otherwise.

### GetUeIpv6AddrPrefixOk

`func (o *PduSessionTsnBridge) GetUeIpv6AddrPrefixOk() (*Ipv6Prefix, bool)`

GetUeIpv6AddrPrefixOk returns a tuple with the UeIpv6AddrPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6AddrPrefix

`func (o *PduSessionTsnBridge) SetUeIpv6AddrPrefix(v Ipv6Prefix)`

SetUeIpv6AddrPrefix sets UeIpv6AddrPrefix field to given value.

### HasUeIpv6AddrPrefix

`func (o *PduSessionTsnBridge) HasUeIpv6AddrPrefix() bool`

HasUeIpv6AddrPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


