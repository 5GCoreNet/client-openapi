# PdnConnectionInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**PdnConnectionStatus**](PdnConnectionStatus.md) |  | 
**Apn** | Pointer to **string** | Identify the APN, it is depending on the SCEF local configuration whether or not this attribute is sent to the SCS/AS. | [optional] 
**PdnType** | [**PdnType**](PdnType.md) |  | 
**InterfaceInd** | Pointer to [**InterfaceIndication**](InterfaceIndication.md) |  | [optional] 
**Ipv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**Ipv6Addrs** | Pointer to **[]string** |  | [optional] 
**MacAddrs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPdnConnectionInformation

`func NewPdnConnectionInformation(status PdnConnectionStatus, pdnType PdnType, ) *PdnConnectionInformation`

NewPdnConnectionInformation instantiates a new PdnConnectionInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdnConnectionInformationWithDefaults

`func NewPdnConnectionInformationWithDefaults() *PdnConnectionInformation`

NewPdnConnectionInformationWithDefaults instantiates a new PdnConnectionInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PdnConnectionInformation) GetStatus() PdnConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PdnConnectionInformation) GetStatusOk() (*PdnConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PdnConnectionInformation) SetStatus(v PdnConnectionStatus)`

SetStatus sets Status field to given value.


### GetApn

`func (o *PdnConnectionInformation) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *PdnConnectionInformation) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *PdnConnectionInformation) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *PdnConnectionInformation) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetPdnType

`func (o *PdnConnectionInformation) GetPdnType() PdnType`

GetPdnType returns the PdnType field if non-nil, zero value otherwise.

### GetPdnTypeOk

`func (o *PdnConnectionInformation) GetPdnTypeOk() (*PdnType, bool)`

GetPdnTypeOk returns a tuple with the PdnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnType

`func (o *PdnConnectionInformation) SetPdnType(v PdnType)`

SetPdnType sets PdnType field to given value.


### GetInterfaceInd

`func (o *PdnConnectionInformation) GetInterfaceInd() InterfaceIndication`

GetInterfaceInd returns the InterfaceInd field if non-nil, zero value otherwise.

### GetInterfaceIndOk

`func (o *PdnConnectionInformation) GetInterfaceIndOk() (*InterfaceIndication, bool)`

GetInterfaceIndOk returns a tuple with the InterfaceInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceInd

`func (o *PdnConnectionInformation) SetInterfaceInd(v InterfaceIndication)`

SetInterfaceInd sets InterfaceInd field to given value.

### HasInterfaceInd

`func (o *PdnConnectionInformation) HasInterfaceInd() bool`

HasInterfaceInd returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *PdnConnectionInformation) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *PdnConnectionInformation) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *PdnConnectionInformation) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *PdnConnectionInformation) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *PdnConnectionInformation) GetIpv6Addrs() []string`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *PdnConnectionInformation) GetIpv6AddrsOk() (*[]string, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *PdnConnectionInformation) SetIpv6Addrs(v []string)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *PdnConnectionInformation) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.

### GetMacAddrs

`func (o *PdnConnectionInformation) GetMacAddrs() []string`

GetMacAddrs returns the MacAddrs field if non-nil, zero value otherwise.

### GetMacAddrsOk

`func (o *PdnConnectionInformation) GetMacAddrsOk() (*[]string, bool)`

GetMacAddrsOk returns a tuple with the MacAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddrs

`func (o *PdnConnectionInformation) SetMacAddrs(v []string)`

SetMacAddrs sets MacAddrs field to given value.

### HasMacAddrs

`func (o *PdnConnectionInformation) HasMacAddrs() bool`

HasMacAddrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


