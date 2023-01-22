# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfTransId** | Pointer to **string** | Identifies an NEF Northbound interface transaction, generated by the AF. | [optional] 
**DnaiChgType** | [**DnaiChangeType**](DnaiChangeType.md) |  | 
**SourceTrafficRoute** | Pointer to [**NullableRouteToLocation**](RouteToLocation.md) |  | [optional] 
**SubscribedEvent** | [**SubscribedEvent**](SubscribedEvent.md) |  | 
**TargetTrafficRoute** | Pointer to [**NullableRouteToLocation**](RouteToLocation.md) |  | [optional] 
**SourceDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**TargetDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**SrcUeIpv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**SrcUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**TgtUeIpv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**TgtUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**UeMac** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**AfAckUri** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewEventNotification

`func NewEventNotification(dnaiChgType DnaiChangeType, subscribedEvent SubscribedEvent, ) *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfTransId

`func (o *EventNotification) GetAfTransId() string`

GetAfTransId returns the AfTransId field if non-nil, zero value otherwise.

### GetAfTransIdOk

`func (o *EventNotification) GetAfTransIdOk() (*string, bool)`

GetAfTransIdOk returns a tuple with the AfTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfTransId

`func (o *EventNotification) SetAfTransId(v string)`

SetAfTransId sets AfTransId field to given value.

### HasAfTransId

`func (o *EventNotification) HasAfTransId() bool`

HasAfTransId returns a boolean if a field has been set.

### GetDnaiChgType

`func (o *EventNotification) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *EventNotification) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *EventNotification) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.


### GetSourceTrafficRoute

`func (o *EventNotification) GetSourceTrafficRoute() RouteToLocation`

GetSourceTrafficRoute returns the SourceTrafficRoute field if non-nil, zero value otherwise.

### GetSourceTrafficRouteOk

`func (o *EventNotification) GetSourceTrafficRouteOk() (*RouteToLocation, bool)`

GetSourceTrafficRouteOk returns a tuple with the SourceTrafficRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTrafficRoute

`func (o *EventNotification) SetSourceTrafficRoute(v RouteToLocation)`

SetSourceTrafficRoute sets SourceTrafficRoute field to given value.

### HasSourceTrafficRoute

`func (o *EventNotification) HasSourceTrafficRoute() bool`

HasSourceTrafficRoute returns a boolean if a field has been set.

### SetSourceTrafficRouteNil

`func (o *EventNotification) SetSourceTrafficRouteNil(b bool)`

 SetSourceTrafficRouteNil sets the value for SourceTrafficRoute to be an explicit nil

### UnsetSourceTrafficRoute
`func (o *EventNotification) UnsetSourceTrafficRoute()`

UnsetSourceTrafficRoute ensures that no value is present for SourceTrafficRoute, not even an explicit nil
### GetSubscribedEvent

`func (o *EventNotification) GetSubscribedEvent() SubscribedEvent`

GetSubscribedEvent returns the SubscribedEvent field if non-nil, zero value otherwise.

### GetSubscribedEventOk

`func (o *EventNotification) GetSubscribedEventOk() (*SubscribedEvent, bool)`

GetSubscribedEventOk returns a tuple with the SubscribedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvent

`func (o *EventNotification) SetSubscribedEvent(v SubscribedEvent)`

SetSubscribedEvent sets SubscribedEvent field to given value.


### GetTargetTrafficRoute

`func (o *EventNotification) GetTargetTrafficRoute() RouteToLocation`

GetTargetTrafficRoute returns the TargetTrafficRoute field if non-nil, zero value otherwise.

### GetTargetTrafficRouteOk

`func (o *EventNotification) GetTargetTrafficRouteOk() (*RouteToLocation, bool)`

GetTargetTrafficRouteOk returns a tuple with the TargetTrafficRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTrafficRoute

`func (o *EventNotification) SetTargetTrafficRoute(v RouteToLocation)`

SetTargetTrafficRoute sets TargetTrafficRoute field to given value.

### HasTargetTrafficRoute

`func (o *EventNotification) HasTargetTrafficRoute() bool`

HasTargetTrafficRoute returns a boolean if a field has been set.

### SetTargetTrafficRouteNil

`func (o *EventNotification) SetTargetTrafficRouteNil(b bool)`

 SetTargetTrafficRouteNil sets the value for TargetTrafficRoute to be an explicit nil

### UnsetTargetTrafficRoute
`func (o *EventNotification) UnsetTargetTrafficRoute()`

UnsetTargetTrafficRoute ensures that no value is present for TargetTrafficRoute, not even an explicit nil
### GetSourceDnai

`func (o *EventNotification) GetSourceDnai() string`

GetSourceDnai returns the SourceDnai field if non-nil, zero value otherwise.

### GetSourceDnaiOk

`func (o *EventNotification) GetSourceDnaiOk() (*string, bool)`

GetSourceDnaiOk returns a tuple with the SourceDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDnai

`func (o *EventNotification) SetSourceDnai(v string)`

SetSourceDnai sets SourceDnai field to given value.

### HasSourceDnai

`func (o *EventNotification) HasSourceDnai() bool`

HasSourceDnai returns a boolean if a field has been set.

### GetTargetDnai

`func (o *EventNotification) GetTargetDnai() string`

GetTargetDnai returns the TargetDnai field if non-nil, zero value otherwise.

### GetTargetDnaiOk

`func (o *EventNotification) GetTargetDnaiOk() (*string, bool)`

GetTargetDnaiOk returns a tuple with the TargetDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnai

`func (o *EventNotification) SetTargetDnai(v string)`

SetTargetDnai sets TargetDnai field to given value.

### HasTargetDnai

`func (o *EventNotification) HasTargetDnai() bool`

HasTargetDnai returns a boolean if a field has been set.

### GetGpsi

`func (o *EventNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *EventNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *EventNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *EventNotification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSrcUeIpv4Addr

`func (o *EventNotification) GetSrcUeIpv4Addr() string`

GetSrcUeIpv4Addr returns the SrcUeIpv4Addr field if non-nil, zero value otherwise.

### GetSrcUeIpv4AddrOk

`func (o *EventNotification) GetSrcUeIpv4AddrOk() (*string, bool)`

GetSrcUeIpv4AddrOk returns a tuple with the SrcUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUeIpv4Addr

`func (o *EventNotification) SetSrcUeIpv4Addr(v string)`

SetSrcUeIpv4Addr sets SrcUeIpv4Addr field to given value.

### HasSrcUeIpv4Addr

`func (o *EventNotification) HasSrcUeIpv4Addr() bool`

HasSrcUeIpv4Addr returns a boolean if a field has been set.

### GetSrcUeIpv6Prefix

`func (o *EventNotification) GetSrcUeIpv6Prefix() Ipv6Prefix`

GetSrcUeIpv6Prefix returns the SrcUeIpv6Prefix field if non-nil, zero value otherwise.

### GetSrcUeIpv6PrefixOk

`func (o *EventNotification) GetSrcUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetSrcUeIpv6PrefixOk returns a tuple with the SrcUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUeIpv6Prefix

`func (o *EventNotification) SetSrcUeIpv6Prefix(v Ipv6Prefix)`

SetSrcUeIpv6Prefix sets SrcUeIpv6Prefix field to given value.

### HasSrcUeIpv6Prefix

`func (o *EventNotification) HasSrcUeIpv6Prefix() bool`

HasSrcUeIpv6Prefix returns a boolean if a field has been set.

### GetTgtUeIpv4Addr

`func (o *EventNotification) GetTgtUeIpv4Addr() string`

GetTgtUeIpv4Addr returns the TgtUeIpv4Addr field if non-nil, zero value otherwise.

### GetTgtUeIpv4AddrOk

`func (o *EventNotification) GetTgtUeIpv4AddrOk() (*string, bool)`

GetTgtUeIpv4AddrOk returns a tuple with the TgtUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUeIpv4Addr

`func (o *EventNotification) SetTgtUeIpv4Addr(v string)`

SetTgtUeIpv4Addr sets TgtUeIpv4Addr field to given value.

### HasTgtUeIpv4Addr

`func (o *EventNotification) HasTgtUeIpv4Addr() bool`

HasTgtUeIpv4Addr returns a boolean if a field has been set.

### GetTgtUeIpv6Prefix

`func (o *EventNotification) GetTgtUeIpv6Prefix() Ipv6Prefix`

GetTgtUeIpv6Prefix returns the TgtUeIpv6Prefix field if non-nil, zero value otherwise.

### GetTgtUeIpv6PrefixOk

`func (o *EventNotification) GetTgtUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetTgtUeIpv6PrefixOk returns a tuple with the TgtUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUeIpv6Prefix

`func (o *EventNotification) SetTgtUeIpv6Prefix(v Ipv6Prefix)`

SetTgtUeIpv6Prefix sets TgtUeIpv6Prefix field to given value.

### HasTgtUeIpv6Prefix

`func (o *EventNotification) HasTgtUeIpv6Prefix() bool`

HasTgtUeIpv6Prefix returns a boolean if a field has been set.

### GetUeMac

`func (o *EventNotification) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *EventNotification) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *EventNotification) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *EventNotification) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetAfAckUri

`func (o *EventNotification) GetAfAckUri() string`

GetAfAckUri returns the AfAckUri field if non-nil, zero value otherwise.

### GetAfAckUriOk

`func (o *EventNotification) GetAfAckUriOk() (*string, bool)`

GetAfAckUriOk returns a tuple with the AfAckUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAckUri

`func (o *EventNotification) SetAfAckUri(v string)`

SetAfAckUri sets AfAckUri field to given value.

### HasAfAckUri

`func (o *EventNotification) HasAfAckUri() bool`

HasAfAckUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

