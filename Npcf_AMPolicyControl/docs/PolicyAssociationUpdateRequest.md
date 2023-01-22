# PolicyAssociationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AltNotifIpv4Addrs** | Pointer to **[]string** | Alternate or backup IPv4 Address(es) where to send Notifications. | [optional] 
**AltNotifIpv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) | Alternate or backup IPv6 Address(es) where to send Notifications. | [optional] 
**AltNotifFqdns** | Pointer to **[]string** | Alternate or backup FQDN(s) where to send Notifications. | [optional] 
**Triggers** | Pointer to [**[]RequestTrigger**](RequestTrigger.md) | Request Triggers that the NF service consumer observes. | [optional] 
**ServAreaRes** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**WlServAreaRes** | Pointer to [**WirelineServiceAreaRestriction**](WirelineServiceAreaRestriction.md) |  | [optional] 
**Rfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**SmfSelInfo** | Pointer to [**NullableSmfSelectionData**](SmfSelectionData.md) |  | [optional] 
**UeAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**UeSliceMbrs** | Pointer to [**[]UeSliceMbr**](UeSliceMbr.md) | The subscribed UE-Slice-MBR for each subscribed S-NSSAI of the home PLMN mapping to a S-NSSAI of the serving PLMN Shall be provided for the \&quot;UE_SLICE_MBR_CH\&quot; policy control request trigger.  | [optional] 
**PraStatuses** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | Contains the UE presence status for tracking area for which changes of the UE presence occurred. The praId attribute within the PresenceInfo data type is the key of the map.  | [optional] 
**UserLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**AllowedSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of allowed S-NSSAIs for the 3GPP access. | [optional] 
**TargetSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of target S-NSSAIs. | [optional] 
**MappingSnssais** | Pointer to [**[]MappingOfSnssai**](MappingOfSnssai.md) | mapping of each S-NSSAI of the Allowed NSSAI to the corresponding S-NSSAI of the HPLMN.   | [optional] 
**AccessTypes** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**RatTypes** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**N3gAllowedSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of allowed S-NSSAIs for the Non-3GPP access. | [optional] 
**TraceReq** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**NwdafDatas** | Pointer to [**[]NwdafData**](NwdafData.md) |  | [optional] 

## Methods

### NewPolicyAssociationUpdateRequest

`func NewPolicyAssociationUpdateRequest() *PolicyAssociationUpdateRequest`

NewPolicyAssociationUpdateRequest instantiates a new PolicyAssociationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAssociationUpdateRequestWithDefaults

`func NewPolicyAssociationUpdateRequestWithDefaults() *PolicyAssociationUpdateRequest`

NewPolicyAssociationUpdateRequestWithDefaults instantiates a new PolicyAssociationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *PolicyAssociationUpdateRequest) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *PolicyAssociationUpdateRequest) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *PolicyAssociationUpdateRequest) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.

### HasNotificationUri

`func (o *PolicyAssociationUpdateRequest) HasNotificationUri() bool`

HasNotificationUri returns a boolean if a field has been set.

### GetAltNotifIpv4Addrs

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4Addrs() []string`

GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv4AddrsOk

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4AddrsOk() (*[]string, bool)`

GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv4Addrs

`func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv4Addrs(v []string)`

SetAltNotifIpv4Addrs sets AltNotifIpv4Addrs field to given value.

### HasAltNotifIpv4Addrs

`func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv4Addrs() bool`

HasAltNotifIpv4Addrs returns a boolean if a field has been set.

### GetAltNotifIpv6Addrs

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6Addrs() []Ipv6Addr`

GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv6AddrsOk

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv6Addrs

`func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv6Addrs(v []Ipv6Addr)`

SetAltNotifIpv6Addrs sets AltNotifIpv6Addrs field to given value.

### HasAltNotifIpv6Addrs

`func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv6Addrs() bool`

HasAltNotifIpv6Addrs returns a boolean if a field has been set.

### GetAltNotifFqdns

`func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdns() []string`

GetAltNotifFqdns returns the AltNotifFqdns field if non-nil, zero value otherwise.

### GetAltNotifFqdnsOk

`func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdnsOk() (*[]string, bool)`

GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifFqdns

`func (o *PolicyAssociationUpdateRequest) SetAltNotifFqdns(v []string)`

SetAltNotifFqdns sets AltNotifFqdns field to given value.

### HasAltNotifFqdns

`func (o *PolicyAssociationUpdateRequest) HasAltNotifFqdns() bool`

HasAltNotifFqdns returns a boolean if a field has been set.

### GetTriggers

`func (o *PolicyAssociationUpdateRequest) GetTriggers() []RequestTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PolicyAssociationUpdateRequest) GetTriggersOk() (*[]RequestTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PolicyAssociationUpdateRequest) SetTriggers(v []RequestTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PolicyAssociationUpdateRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetServAreaRes

`func (o *PolicyAssociationUpdateRequest) GetServAreaRes() ServiceAreaRestriction`

GetServAreaRes returns the ServAreaRes field if non-nil, zero value otherwise.

### GetServAreaResOk

`func (o *PolicyAssociationUpdateRequest) GetServAreaResOk() (*ServiceAreaRestriction, bool)`

GetServAreaResOk returns a tuple with the ServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAreaRes

`func (o *PolicyAssociationUpdateRequest) SetServAreaRes(v ServiceAreaRestriction)`

SetServAreaRes sets ServAreaRes field to given value.

### HasServAreaRes

`func (o *PolicyAssociationUpdateRequest) HasServAreaRes() bool`

HasServAreaRes returns a boolean if a field has been set.

### GetWlServAreaRes

`func (o *PolicyAssociationUpdateRequest) GetWlServAreaRes() WirelineServiceAreaRestriction`

GetWlServAreaRes returns the WlServAreaRes field if non-nil, zero value otherwise.

### GetWlServAreaResOk

`func (o *PolicyAssociationUpdateRequest) GetWlServAreaResOk() (*WirelineServiceAreaRestriction, bool)`

GetWlServAreaResOk returns a tuple with the WlServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlServAreaRes

`func (o *PolicyAssociationUpdateRequest) SetWlServAreaRes(v WirelineServiceAreaRestriction)`

SetWlServAreaRes sets WlServAreaRes field to given value.

### HasWlServAreaRes

`func (o *PolicyAssociationUpdateRequest) HasWlServAreaRes() bool`

HasWlServAreaRes returns a boolean if a field has been set.

### GetRfsp

`func (o *PolicyAssociationUpdateRequest) GetRfsp() int32`

GetRfsp returns the Rfsp field if non-nil, zero value otherwise.

### GetRfspOk

`func (o *PolicyAssociationUpdateRequest) GetRfspOk() (*int32, bool)`

GetRfspOk returns a tuple with the Rfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfsp

`func (o *PolicyAssociationUpdateRequest) SetRfsp(v int32)`

SetRfsp sets Rfsp field to given value.

### HasRfsp

`func (o *PolicyAssociationUpdateRequest) HasRfsp() bool`

HasRfsp returns a boolean if a field has been set.

### GetSmfSelInfo

`func (o *PolicyAssociationUpdateRequest) GetSmfSelInfo() SmfSelectionData`

GetSmfSelInfo returns the SmfSelInfo field if non-nil, zero value otherwise.

### GetSmfSelInfoOk

`func (o *PolicyAssociationUpdateRequest) GetSmfSelInfoOk() (*SmfSelectionData, bool)`

GetSmfSelInfoOk returns a tuple with the SmfSelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelInfo

`func (o *PolicyAssociationUpdateRequest) SetSmfSelInfo(v SmfSelectionData)`

SetSmfSelInfo sets SmfSelInfo field to given value.

### HasSmfSelInfo

`func (o *PolicyAssociationUpdateRequest) HasSmfSelInfo() bool`

HasSmfSelInfo returns a boolean if a field has been set.

### SetSmfSelInfoNil

`func (o *PolicyAssociationUpdateRequest) SetSmfSelInfoNil(b bool)`

 SetSmfSelInfoNil sets the value for SmfSelInfo to be an explicit nil

### UnsetSmfSelInfo
`func (o *PolicyAssociationUpdateRequest) UnsetSmfSelInfo()`

UnsetSmfSelInfo ensures that no value is present for SmfSelInfo, not even an explicit nil
### GetUeAmbr

`func (o *PolicyAssociationUpdateRequest) GetUeAmbr() Ambr`

GetUeAmbr returns the UeAmbr field if non-nil, zero value otherwise.

### GetUeAmbrOk

`func (o *PolicyAssociationUpdateRequest) GetUeAmbrOk() (*Ambr, bool)`

GetUeAmbrOk returns a tuple with the UeAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAmbr

`func (o *PolicyAssociationUpdateRequest) SetUeAmbr(v Ambr)`

SetUeAmbr sets UeAmbr field to given value.

### HasUeAmbr

`func (o *PolicyAssociationUpdateRequest) HasUeAmbr() bool`

HasUeAmbr returns a boolean if a field has been set.

### GetUeSliceMbrs

`func (o *PolicyAssociationUpdateRequest) GetUeSliceMbrs() []UeSliceMbr`

GetUeSliceMbrs returns the UeSliceMbrs field if non-nil, zero value otherwise.

### GetUeSliceMbrsOk

`func (o *PolicyAssociationUpdateRequest) GetUeSliceMbrsOk() (*[]UeSliceMbr, bool)`

GetUeSliceMbrsOk returns a tuple with the UeSliceMbrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSliceMbrs

`func (o *PolicyAssociationUpdateRequest) SetUeSliceMbrs(v []UeSliceMbr)`

SetUeSliceMbrs sets UeSliceMbrs field to given value.

### HasUeSliceMbrs

`func (o *PolicyAssociationUpdateRequest) HasUeSliceMbrs() bool`

HasUeSliceMbrs returns a boolean if a field has been set.

### GetPraStatuses

`func (o *PolicyAssociationUpdateRequest) GetPraStatuses() map[string]PresenceInfo`

GetPraStatuses returns the PraStatuses field if non-nil, zero value otherwise.

### GetPraStatusesOk

`func (o *PolicyAssociationUpdateRequest) GetPraStatusesOk() (*map[string]PresenceInfo, bool)`

GetPraStatusesOk returns a tuple with the PraStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraStatuses

`func (o *PolicyAssociationUpdateRequest) SetPraStatuses(v map[string]PresenceInfo)`

SetPraStatuses sets PraStatuses field to given value.

### HasPraStatuses

`func (o *PolicyAssociationUpdateRequest) HasPraStatuses() bool`

HasPraStatuses returns a boolean if a field has been set.

### GetUserLoc

`func (o *PolicyAssociationUpdateRequest) GetUserLoc() UserLocation`

GetUserLoc returns the UserLoc field if non-nil, zero value otherwise.

### GetUserLocOk

`func (o *PolicyAssociationUpdateRequest) GetUserLocOk() (*UserLocation, bool)`

GetUserLocOk returns a tuple with the UserLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoc

`func (o *PolicyAssociationUpdateRequest) SetUserLoc(v UserLocation)`

SetUserLoc sets UserLoc field to given value.

### HasUserLoc

`func (o *PolicyAssociationUpdateRequest) HasUserLoc() bool`

HasUserLoc returns a boolean if a field has been set.

### GetAllowedSnssais

`func (o *PolicyAssociationUpdateRequest) GetAllowedSnssais() []Snssai`

GetAllowedSnssais returns the AllowedSnssais field if non-nil, zero value otherwise.

### GetAllowedSnssaisOk

`func (o *PolicyAssociationUpdateRequest) GetAllowedSnssaisOk() (*[]Snssai, bool)`

GetAllowedSnssaisOk returns a tuple with the AllowedSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnssais

`func (o *PolicyAssociationUpdateRequest) SetAllowedSnssais(v []Snssai)`

SetAllowedSnssais sets AllowedSnssais field to given value.

### HasAllowedSnssais

`func (o *PolicyAssociationUpdateRequest) HasAllowedSnssais() bool`

HasAllowedSnssais returns a boolean if a field has been set.

### GetTargetSnssais

`func (o *PolicyAssociationUpdateRequest) GetTargetSnssais() []Snssai`

GetTargetSnssais returns the TargetSnssais field if non-nil, zero value otherwise.

### GetTargetSnssaisOk

`func (o *PolicyAssociationUpdateRequest) GetTargetSnssaisOk() (*[]Snssai, bool)`

GetTargetSnssaisOk returns a tuple with the TargetSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSnssais

`func (o *PolicyAssociationUpdateRequest) SetTargetSnssais(v []Snssai)`

SetTargetSnssais sets TargetSnssais field to given value.

### HasTargetSnssais

`func (o *PolicyAssociationUpdateRequest) HasTargetSnssais() bool`

HasTargetSnssais returns a boolean if a field has been set.

### GetMappingSnssais

`func (o *PolicyAssociationUpdateRequest) GetMappingSnssais() []MappingOfSnssai`

GetMappingSnssais returns the MappingSnssais field if non-nil, zero value otherwise.

### GetMappingSnssaisOk

`func (o *PolicyAssociationUpdateRequest) GetMappingSnssaisOk() (*[]MappingOfSnssai, bool)`

GetMappingSnssaisOk returns a tuple with the MappingSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingSnssais

`func (o *PolicyAssociationUpdateRequest) SetMappingSnssais(v []MappingOfSnssai)`

SetMappingSnssais sets MappingSnssais field to given value.

### HasMappingSnssais

`func (o *PolicyAssociationUpdateRequest) HasMappingSnssais() bool`

HasMappingSnssais returns a boolean if a field has been set.

### GetAccessTypes

`func (o *PolicyAssociationUpdateRequest) GetAccessTypes() []AccessType`

GetAccessTypes returns the AccessTypes field if non-nil, zero value otherwise.

### GetAccessTypesOk

`func (o *PolicyAssociationUpdateRequest) GetAccessTypesOk() (*[]AccessType, bool)`

GetAccessTypesOk returns a tuple with the AccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypes

`func (o *PolicyAssociationUpdateRequest) SetAccessTypes(v []AccessType)`

SetAccessTypes sets AccessTypes field to given value.

### HasAccessTypes

`func (o *PolicyAssociationUpdateRequest) HasAccessTypes() bool`

HasAccessTypes returns a boolean if a field has been set.

### GetRatTypes

`func (o *PolicyAssociationUpdateRequest) GetRatTypes() []RatType`

GetRatTypes returns the RatTypes field if non-nil, zero value otherwise.

### GetRatTypesOk

`func (o *PolicyAssociationUpdateRequest) GetRatTypesOk() (*[]RatType, bool)`

GetRatTypesOk returns a tuple with the RatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatTypes

`func (o *PolicyAssociationUpdateRequest) SetRatTypes(v []RatType)`

SetRatTypes sets RatTypes field to given value.

### HasRatTypes

`func (o *PolicyAssociationUpdateRequest) HasRatTypes() bool`

HasRatTypes returns a boolean if a field has been set.

### GetN3gAllowedSnssais

`func (o *PolicyAssociationUpdateRequest) GetN3gAllowedSnssais() []Snssai`

GetN3gAllowedSnssais returns the N3gAllowedSnssais field if non-nil, zero value otherwise.

### GetN3gAllowedSnssaisOk

`func (o *PolicyAssociationUpdateRequest) GetN3gAllowedSnssaisOk() (*[]Snssai, bool)`

GetN3gAllowedSnssaisOk returns a tuple with the N3gAllowedSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gAllowedSnssais

`func (o *PolicyAssociationUpdateRequest) SetN3gAllowedSnssais(v []Snssai)`

SetN3gAllowedSnssais sets N3gAllowedSnssais field to given value.

### HasN3gAllowedSnssais

`func (o *PolicyAssociationUpdateRequest) HasN3gAllowedSnssais() bool`

HasN3gAllowedSnssais returns a boolean if a field has been set.

### GetTraceReq

`func (o *PolicyAssociationUpdateRequest) GetTraceReq() TraceData`

GetTraceReq returns the TraceReq field if non-nil, zero value otherwise.

### GetTraceReqOk

`func (o *PolicyAssociationUpdateRequest) GetTraceReqOk() (*TraceData, bool)`

GetTraceReqOk returns a tuple with the TraceReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceReq

`func (o *PolicyAssociationUpdateRequest) SetTraceReq(v TraceData)`

SetTraceReq sets TraceReq field to given value.

### HasTraceReq

`func (o *PolicyAssociationUpdateRequest) HasTraceReq() bool`

HasTraceReq returns a boolean if a field has been set.

### SetTraceReqNil

`func (o *PolicyAssociationUpdateRequest) SetTraceReqNil(b bool)`

 SetTraceReqNil sets the value for TraceReq to be an explicit nil

### UnsetTraceReq
`func (o *PolicyAssociationUpdateRequest) UnsetTraceReq()`

UnsetTraceReq ensures that no value is present for TraceReq, not even an explicit nil
### GetGuami

`func (o *PolicyAssociationUpdateRequest) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *PolicyAssociationUpdateRequest) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *PolicyAssociationUpdateRequest) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *PolicyAssociationUpdateRequest) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetNwdafDatas

`func (o *PolicyAssociationUpdateRequest) GetNwdafDatas() []NwdafData`

GetNwdafDatas returns the NwdafDatas field if non-nil, zero value otherwise.

### GetNwdafDatasOk

`func (o *PolicyAssociationUpdateRequest) GetNwdafDatasOk() (*[]NwdafData, bool)`

GetNwdafDatasOk returns a tuple with the NwdafDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafDatas

`func (o *PolicyAssociationUpdateRequest) SetNwdafDatas(v []NwdafData)`

SetNwdafDatas sets NwdafDatas field to given value.

### HasNwdafDatas

`func (o *PolicyAssociationUpdateRequest) HasNwdafDatas() bool`

HasNwdafDatas returns a boolean if a field has been set.

### SetNwdafDatasNil

`func (o *PolicyAssociationUpdateRequest) SetNwdafDatasNil(b bool)`

 SetNwdafDatasNil sets the value for NwdafDatas to be an explicit nil

### UnsetNwdafDatas
`func (o *PolicyAssociationUpdateRequest) UnsetNwdafDatas()`

UnsetNwdafDatas ensures that no value is present for NwdafDatas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


