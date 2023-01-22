# UpPathChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | [**IndUeIdentification**](IndUeIdentification.md) |  | 
**DnaiChgType** | [**DnaiChangeType**](DnaiChangeType.md) |  | 
**SourceTrafficRoute** | Pointer to [**NullableRouteToLocation**](RouteToLocation.md) |  | [optional] 
**TargetTrafficRoute** | Pointer to [**NullableRouteToLocation**](RouteToLocation.md) |  | [optional] 
**SourceDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**TargetDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**SrcUeIpv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**SrcUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**TgtUeIpv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**TgtUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 

## Methods

### NewUpPathChangeInfo

`func NewUpPathChangeInfo(ueId IndUeIdentification, dnaiChgType DnaiChangeType, ) *UpPathChangeInfo`

NewUpPathChangeInfo instantiates a new UpPathChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpPathChangeInfoWithDefaults

`func NewUpPathChangeInfoWithDefaults() *UpPathChangeInfo`

NewUpPathChangeInfoWithDefaults instantiates a new UpPathChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *UpPathChangeInfo) GetUeId() IndUeIdentification`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *UpPathChangeInfo) GetUeIdOk() (*IndUeIdentification, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *UpPathChangeInfo) SetUeId(v IndUeIdentification)`

SetUeId sets UeId field to given value.


### GetDnaiChgType

`func (o *UpPathChangeInfo) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *UpPathChangeInfo) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *UpPathChangeInfo) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.


### GetSourceTrafficRoute

`func (o *UpPathChangeInfo) GetSourceTrafficRoute() RouteToLocation`

GetSourceTrafficRoute returns the SourceTrafficRoute field if non-nil, zero value otherwise.

### GetSourceTrafficRouteOk

`func (o *UpPathChangeInfo) GetSourceTrafficRouteOk() (*RouteToLocation, bool)`

GetSourceTrafficRouteOk returns a tuple with the SourceTrafficRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTrafficRoute

`func (o *UpPathChangeInfo) SetSourceTrafficRoute(v RouteToLocation)`

SetSourceTrafficRoute sets SourceTrafficRoute field to given value.

### HasSourceTrafficRoute

`func (o *UpPathChangeInfo) HasSourceTrafficRoute() bool`

HasSourceTrafficRoute returns a boolean if a field has been set.

### SetSourceTrafficRouteNil

`func (o *UpPathChangeInfo) SetSourceTrafficRouteNil(b bool)`

 SetSourceTrafficRouteNil sets the value for SourceTrafficRoute to be an explicit nil

### UnsetSourceTrafficRoute
`func (o *UpPathChangeInfo) UnsetSourceTrafficRoute()`

UnsetSourceTrafficRoute ensures that no value is present for SourceTrafficRoute, not even an explicit nil
### GetTargetTrafficRoute

`func (o *UpPathChangeInfo) GetTargetTrafficRoute() RouteToLocation`

GetTargetTrafficRoute returns the TargetTrafficRoute field if non-nil, zero value otherwise.

### GetTargetTrafficRouteOk

`func (o *UpPathChangeInfo) GetTargetTrafficRouteOk() (*RouteToLocation, bool)`

GetTargetTrafficRouteOk returns a tuple with the TargetTrafficRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTrafficRoute

`func (o *UpPathChangeInfo) SetTargetTrafficRoute(v RouteToLocation)`

SetTargetTrafficRoute sets TargetTrafficRoute field to given value.

### HasTargetTrafficRoute

`func (o *UpPathChangeInfo) HasTargetTrafficRoute() bool`

HasTargetTrafficRoute returns a boolean if a field has been set.

### SetTargetTrafficRouteNil

`func (o *UpPathChangeInfo) SetTargetTrafficRouteNil(b bool)`

 SetTargetTrafficRouteNil sets the value for TargetTrafficRoute to be an explicit nil

### UnsetTargetTrafficRoute
`func (o *UpPathChangeInfo) UnsetTargetTrafficRoute()`

UnsetTargetTrafficRoute ensures that no value is present for TargetTrafficRoute, not even an explicit nil
### GetSourceDnai

`func (o *UpPathChangeInfo) GetSourceDnai() string`

GetSourceDnai returns the SourceDnai field if non-nil, zero value otherwise.

### GetSourceDnaiOk

`func (o *UpPathChangeInfo) GetSourceDnaiOk() (*string, bool)`

GetSourceDnaiOk returns a tuple with the SourceDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDnai

`func (o *UpPathChangeInfo) SetSourceDnai(v string)`

SetSourceDnai sets SourceDnai field to given value.

### HasSourceDnai

`func (o *UpPathChangeInfo) HasSourceDnai() bool`

HasSourceDnai returns a boolean if a field has been set.

### GetTargetDnai

`func (o *UpPathChangeInfo) GetTargetDnai() string`

GetTargetDnai returns the TargetDnai field if non-nil, zero value otherwise.

### GetTargetDnaiOk

`func (o *UpPathChangeInfo) GetTargetDnaiOk() (*string, bool)`

GetTargetDnaiOk returns a tuple with the TargetDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnai

`func (o *UpPathChangeInfo) SetTargetDnai(v string)`

SetTargetDnai sets TargetDnai field to given value.

### HasTargetDnai

`func (o *UpPathChangeInfo) HasTargetDnai() bool`

HasTargetDnai returns a boolean if a field has been set.

### GetSrcUeIpv4Addr

`func (o *UpPathChangeInfo) GetSrcUeIpv4Addr() string`

GetSrcUeIpv4Addr returns the SrcUeIpv4Addr field if non-nil, zero value otherwise.

### GetSrcUeIpv4AddrOk

`func (o *UpPathChangeInfo) GetSrcUeIpv4AddrOk() (*string, bool)`

GetSrcUeIpv4AddrOk returns a tuple with the SrcUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUeIpv4Addr

`func (o *UpPathChangeInfo) SetSrcUeIpv4Addr(v string)`

SetSrcUeIpv4Addr sets SrcUeIpv4Addr field to given value.

### HasSrcUeIpv4Addr

`func (o *UpPathChangeInfo) HasSrcUeIpv4Addr() bool`

HasSrcUeIpv4Addr returns a boolean if a field has been set.

### GetSrcUeIpv6Prefix

`func (o *UpPathChangeInfo) GetSrcUeIpv6Prefix() Ipv6Prefix`

GetSrcUeIpv6Prefix returns the SrcUeIpv6Prefix field if non-nil, zero value otherwise.

### GetSrcUeIpv6PrefixOk

`func (o *UpPathChangeInfo) GetSrcUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetSrcUeIpv6PrefixOk returns a tuple with the SrcUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUeIpv6Prefix

`func (o *UpPathChangeInfo) SetSrcUeIpv6Prefix(v Ipv6Prefix)`

SetSrcUeIpv6Prefix sets SrcUeIpv6Prefix field to given value.

### HasSrcUeIpv6Prefix

`func (o *UpPathChangeInfo) HasSrcUeIpv6Prefix() bool`

HasSrcUeIpv6Prefix returns a boolean if a field has been set.

### GetTgtUeIpv4Addr

`func (o *UpPathChangeInfo) GetTgtUeIpv4Addr() string`

GetTgtUeIpv4Addr returns the TgtUeIpv4Addr field if non-nil, zero value otherwise.

### GetTgtUeIpv4AddrOk

`func (o *UpPathChangeInfo) GetTgtUeIpv4AddrOk() (*string, bool)`

GetTgtUeIpv4AddrOk returns a tuple with the TgtUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUeIpv4Addr

`func (o *UpPathChangeInfo) SetTgtUeIpv4Addr(v string)`

SetTgtUeIpv4Addr sets TgtUeIpv4Addr field to given value.

### HasTgtUeIpv4Addr

`func (o *UpPathChangeInfo) HasTgtUeIpv4Addr() bool`

HasTgtUeIpv4Addr returns a boolean if a field has been set.

### GetTgtUeIpv6Prefix

`func (o *UpPathChangeInfo) GetTgtUeIpv6Prefix() Ipv6Prefix`

GetTgtUeIpv6Prefix returns the TgtUeIpv6Prefix field if non-nil, zero value otherwise.

### GetTgtUeIpv6PrefixOk

`func (o *UpPathChangeInfo) GetTgtUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetTgtUeIpv6PrefixOk returns a tuple with the TgtUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUeIpv6Prefix

`func (o *UpPathChangeInfo) SetTgtUeIpv6Prefix(v Ipv6Prefix)`

SetTgtUeIpv6Prefix sets TgtUeIpv6Prefix field to given value.

### HasTgtUeIpv6Prefix

`func (o *UpPathChangeInfo) HasTgtUeIpv6Prefix() bool`

HasTgtUeIpv6Prefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


