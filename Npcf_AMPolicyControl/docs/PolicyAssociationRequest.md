# PolicyAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**AltNotifIpv4Addrs** | Pointer to **[]string** | Alternate or backup IPv4 Address(es) where to send Notifications. | [optional] 
**AltNotifIpv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) | Alternate or backup IPv6 Address(es) where to send Notifications. | [optional] 
**AltNotifFqdns** | Pointer to **[]string** | Alternate or backup FQDN(s) where to send Notifications. | [optional] 
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AccessTypes** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**UserLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**TimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**ServingPlmn** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**RatTypes** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**ServAreaRes** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**WlServAreaRes** | Pointer to [**WirelineServiceAreaRestriction**](WirelineServiceAreaRestriction.md) |  | [optional] 
**Rfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**UeAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**UeSliceMbrs** | Pointer to [**[]UeSliceMbr**](UeSliceMbr.md) | The subscribed UE Slice-MBR for each subscribed S-NSSAI of the home PLMN mapping  to a S-NSSAI of the serving PLMN Shall be provided when available.  | [optional] 
**AllowedSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of allowed S-NSSAIs for the 3GPP access. | [optional] 
**TargetSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of target S-NSSAIs. | [optional] 
**MappingSnssais** | Pointer to [**[]MappingOfSnssai**](MappingOfSnssai.md) | mapping of each S-NSSAI of the Allowed NSSAI to the corresponding S-NSSAI of the HPLMN.   | [optional] 
**N3gAllowedSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of allowed S-NSSAIs for the Non-3GPP access. | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ServiveName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**TraceReq** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**NwdafDatas** | Pointer to [**[]NwdafData**](NwdafData.md) |  | [optional] 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 

## Methods

### NewPolicyAssociationRequest

`func NewPolicyAssociationRequest(notificationUri string, supi string, suppFeat string, ) *PolicyAssociationRequest`

NewPolicyAssociationRequest instantiates a new PolicyAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAssociationRequestWithDefaults

`func NewPolicyAssociationRequestWithDefaults() *PolicyAssociationRequest`

NewPolicyAssociationRequestWithDefaults instantiates a new PolicyAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *PolicyAssociationRequest) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *PolicyAssociationRequest) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *PolicyAssociationRequest) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetAltNotifIpv4Addrs

`func (o *PolicyAssociationRequest) GetAltNotifIpv4Addrs() []string`

GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv4AddrsOk

`func (o *PolicyAssociationRequest) GetAltNotifIpv4AddrsOk() (*[]string, bool)`

GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv4Addrs

`func (o *PolicyAssociationRequest) SetAltNotifIpv4Addrs(v []string)`

SetAltNotifIpv4Addrs sets AltNotifIpv4Addrs field to given value.

### HasAltNotifIpv4Addrs

`func (o *PolicyAssociationRequest) HasAltNotifIpv4Addrs() bool`

HasAltNotifIpv4Addrs returns a boolean if a field has been set.

### GetAltNotifIpv6Addrs

`func (o *PolicyAssociationRequest) GetAltNotifIpv6Addrs() []Ipv6Addr`

GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv6AddrsOk

`func (o *PolicyAssociationRequest) GetAltNotifIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv6Addrs

`func (o *PolicyAssociationRequest) SetAltNotifIpv6Addrs(v []Ipv6Addr)`

SetAltNotifIpv6Addrs sets AltNotifIpv6Addrs field to given value.

### HasAltNotifIpv6Addrs

`func (o *PolicyAssociationRequest) HasAltNotifIpv6Addrs() bool`

HasAltNotifIpv6Addrs returns a boolean if a field has been set.

### GetAltNotifFqdns

`func (o *PolicyAssociationRequest) GetAltNotifFqdns() []string`

GetAltNotifFqdns returns the AltNotifFqdns field if non-nil, zero value otherwise.

### GetAltNotifFqdnsOk

`func (o *PolicyAssociationRequest) GetAltNotifFqdnsOk() (*[]string, bool)`

GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifFqdns

`func (o *PolicyAssociationRequest) SetAltNotifFqdns(v []string)`

SetAltNotifFqdns sets AltNotifFqdns field to given value.

### HasAltNotifFqdns

`func (o *PolicyAssociationRequest) HasAltNotifFqdns() bool`

HasAltNotifFqdns returns a boolean if a field has been set.

### GetSupi

`func (o *PolicyAssociationRequest) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PolicyAssociationRequest) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PolicyAssociationRequest) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetGpsi

`func (o *PolicyAssociationRequest) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PolicyAssociationRequest) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PolicyAssociationRequest) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PolicyAssociationRequest) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetAccessType

`func (o *PolicyAssociationRequest) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PolicyAssociationRequest) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PolicyAssociationRequest) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *PolicyAssociationRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAccessTypes

`func (o *PolicyAssociationRequest) GetAccessTypes() []AccessType`

GetAccessTypes returns the AccessTypes field if non-nil, zero value otherwise.

### GetAccessTypesOk

`func (o *PolicyAssociationRequest) GetAccessTypesOk() (*[]AccessType, bool)`

GetAccessTypesOk returns a tuple with the AccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypes

`func (o *PolicyAssociationRequest) SetAccessTypes(v []AccessType)`

SetAccessTypes sets AccessTypes field to given value.

### HasAccessTypes

`func (o *PolicyAssociationRequest) HasAccessTypes() bool`

HasAccessTypes returns a boolean if a field has been set.

### GetPei

`func (o *PolicyAssociationRequest) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *PolicyAssociationRequest) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *PolicyAssociationRequest) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *PolicyAssociationRequest) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetUserLoc

`func (o *PolicyAssociationRequest) GetUserLoc() UserLocation`

GetUserLoc returns the UserLoc field if non-nil, zero value otherwise.

### GetUserLocOk

`func (o *PolicyAssociationRequest) GetUserLocOk() (*UserLocation, bool)`

GetUserLocOk returns a tuple with the UserLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoc

`func (o *PolicyAssociationRequest) SetUserLoc(v UserLocation)`

SetUserLoc sets UserLoc field to given value.

### HasUserLoc

`func (o *PolicyAssociationRequest) HasUserLoc() bool`

HasUserLoc returns a boolean if a field has been set.

### GetTimeZone

`func (o *PolicyAssociationRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PolicyAssociationRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PolicyAssociationRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PolicyAssociationRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetServingPlmn

`func (o *PolicyAssociationRequest) GetServingPlmn() PlmnIdNid`

GetServingPlmn returns the ServingPlmn field if non-nil, zero value otherwise.

### GetServingPlmnOk

`func (o *PolicyAssociationRequest) GetServingPlmnOk() (*PlmnIdNid, bool)`

GetServingPlmnOk returns a tuple with the ServingPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingPlmn

`func (o *PolicyAssociationRequest) SetServingPlmn(v PlmnIdNid)`

SetServingPlmn sets ServingPlmn field to given value.

### HasServingPlmn

`func (o *PolicyAssociationRequest) HasServingPlmn() bool`

HasServingPlmn returns a boolean if a field has been set.

### GetRatType

`func (o *PolicyAssociationRequest) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *PolicyAssociationRequest) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *PolicyAssociationRequest) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *PolicyAssociationRequest) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetRatTypes

`func (o *PolicyAssociationRequest) GetRatTypes() []RatType`

GetRatTypes returns the RatTypes field if non-nil, zero value otherwise.

### GetRatTypesOk

`func (o *PolicyAssociationRequest) GetRatTypesOk() (*[]RatType, bool)`

GetRatTypesOk returns a tuple with the RatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatTypes

`func (o *PolicyAssociationRequest) SetRatTypes(v []RatType)`

SetRatTypes sets RatTypes field to given value.

### HasRatTypes

`func (o *PolicyAssociationRequest) HasRatTypes() bool`

HasRatTypes returns a boolean if a field has been set.

### GetGroupIds

`func (o *PolicyAssociationRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *PolicyAssociationRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *PolicyAssociationRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *PolicyAssociationRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetServAreaRes

`func (o *PolicyAssociationRequest) GetServAreaRes() ServiceAreaRestriction`

GetServAreaRes returns the ServAreaRes field if non-nil, zero value otherwise.

### GetServAreaResOk

`func (o *PolicyAssociationRequest) GetServAreaResOk() (*ServiceAreaRestriction, bool)`

GetServAreaResOk returns a tuple with the ServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAreaRes

`func (o *PolicyAssociationRequest) SetServAreaRes(v ServiceAreaRestriction)`

SetServAreaRes sets ServAreaRes field to given value.

### HasServAreaRes

`func (o *PolicyAssociationRequest) HasServAreaRes() bool`

HasServAreaRes returns a boolean if a field has been set.

### GetWlServAreaRes

`func (o *PolicyAssociationRequest) GetWlServAreaRes() WirelineServiceAreaRestriction`

GetWlServAreaRes returns the WlServAreaRes field if non-nil, zero value otherwise.

### GetWlServAreaResOk

`func (o *PolicyAssociationRequest) GetWlServAreaResOk() (*WirelineServiceAreaRestriction, bool)`

GetWlServAreaResOk returns a tuple with the WlServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlServAreaRes

`func (o *PolicyAssociationRequest) SetWlServAreaRes(v WirelineServiceAreaRestriction)`

SetWlServAreaRes sets WlServAreaRes field to given value.

### HasWlServAreaRes

`func (o *PolicyAssociationRequest) HasWlServAreaRes() bool`

HasWlServAreaRes returns a boolean if a field has been set.

### GetRfsp

`func (o *PolicyAssociationRequest) GetRfsp() int32`

GetRfsp returns the Rfsp field if non-nil, zero value otherwise.

### GetRfspOk

`func (o *PolicyAssociationRequest) GetRfspOk() (*int32, bool)`

GetRfspOk returns a tuple with the Rfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfsp

`func (o *PolicyAssociationRequest) SetRfsp(v int32)`

SetRfsp sets Rfsp field to given value.

### HasRfsp

`func (o *PolicyAssociationRequest) HasRfsp() bool`

HasRfsp returns a boolean if a field has been set.

### GetUeAmbr

`func (o *PolicyAssociationRequest) GetUeAmbr() Ambr`

GetUeAmbr returns the UeAmbr field if non-nil, zero value otherwise.

### GetUeAmbrOk

`func (o *PolicyAssociationRequest) GetUeAmbrOk() (*Ambr, bool)`

GetUeAmbrOk returns a tuple with the UeAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAmbr

`func (o *PolicyAssociationRequest) SetUeAmbr(v Ambr)`

SetUeAmbr sets UeAmbr field to given value.

### HasUeAmbr

`func (o *PolicyAssociationRequest) HasUeAmbr() bool`

HasUeAmbr returns a boolean if a field has been set.

### GetUeSliceMbrs

`func (o *PolicyAssociationRequest) GetUeSliceMbrs() []UeSliceMbr`

GetUeSliceMbrs returns the UeSliceMbrs field if non-nil, zero value otherwise.

### GetUeSliceMbrsOk

`func (o *PolicyAssociationRequest) GetUeSliceMbrsOk() (*[]UeSliceMbr, bool)`

GetUeSliceMbrsOk returns a tuple with the UeSliceMbrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSliceMbrs

`func (o *PolicyAssociationRequest) SetUeSliceMbrs(v []UeSliceMbr)`

SetUeSliceMbrs sets UeSliceMbrs field to given value.

### HasUeSliceMbrs

`func (o *PolicyAssociationRequest) HasUeSliceMbrs() bool`

HasUeSliceMbrs returns a boolean if a field has been set.

### GetAllowedSnssais

`func (o *PolicyAssociationRequest) GetAllowedSnssais() []Snssai`

GetAllowedSnssais returns the AllowedSnssais field if non-nil, zero value otherwise.

### GetAllowedSnssaisOk

`func (o *PolicyAssociationRequest) GetAllowedSnssaisOk() (*[]Snssai, bool)`

GetAllowedSnssaisOk returns a tuple with the AllowedSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnssais

`func (o *PolicyAssociationRequest) SetAllowedSnssais(v []Snssai)`

SetAllowedSnssais sets AllowedSnssais field to given value.

### HasAllowedSnssais

`func (o *PolicyAssociationRequest) HasAllowedSnssais() bool`

HasAllowedSnssais returns a boolean if a field has been set.

### GetTargetSnssais

`func (o *PolicyAssociationRequest) GetTargetSnssais() []Snssai`

GetTargetSnssais returns the TargetSnssais field if non-nil, zero value otherwise.

### GetTargetSnssaisOk

`func (o *PolicyAssociationRequest) GetTargetSnssaisOk() (*[]Snssai, bool)`

GetTargetSnssaisOk returns a tuple with the TargetSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSnssais

`func (o *PolicyAssociationRequest) SetTargetSnssais(v []Snssai)`

SetTargetSnssais sets TargetSnssais field to given value.

### HasTargetSnssais

`func (o *PolicyAssociationRequest) HasTargetSnssais() bool`

HasTargetSnssais returns a boolean if a field has been set.

### GetMappingSnssais

`func (o *PolicyAssociationRequest) GetMappingSnssais() []MappingOfSnssai`

GetMappingSnssais returns the MappingSnssais field if non-nil, zero value otherwise.

### GetMappingSnssaisOk

`func (o *PolicyAssociationRequest) GetMappingSnssaisOk() (*[]MappingOfSnssai, bool)`

GetMappingSnssaisOk returns a tuple with the MappingSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingSnssais

`func (o *PolicyAssociationRequest) SetMappingSnssais(v []MappingOfSnssai)`

SetMappingSnssais sets MappingSnssais field to given value.

### HasMappingSnssais

`func (o *PolicyAssociationRequest) HasMappingSnssais() bool`

HasMappingSnssais returns a boolean if a field has been set.

### GetN3gAllowedSnssais

`func (o *PolicyAssociationRequest) GetN3gAllowedSnssais() []Snssai`

GetN3gAllowedSnssais returns the N3gAllowedSnssais field if non-nil, zero value otherwise.

### GetN3gAllowedSnssaisOk

`func (o *PolicyAssociationRequest) GetN3gAllowedSnssaisOk() (*[]Snssai, bool)`

GetN3gAllowedSnssaisOk returns a tuple with the N3gAllowedSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gAllowedSnssais

`func (o *PolicyAssociationRequest) SetN3gAllowedSnssais(v []Snssai)`

SetN3gAllowedSnssais sets N3gAllowedSnssais field to given value.

### HasN3gAllowedSnssais

`func (o *PolicyAssociationRequest) HasN3gAllowedSnssais() bool`

HasN3gAllowedSnssais returns a boolean if a field has been set.

### GetGuami

`func (o *PolicyAssociationRequest) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *PolicyAssociationRequest) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *PolicyAssociationRequest) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *PolicyAssociationRequest) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetServiveName

`func (o *PolicyAssociationRequest) GetServiveName() ServiceName`

GetServiveName returns the ServiveName field if non-nil, zero value otherwise.

### GetServiveNameOk

`func (o *PolicyAssociationRequest) GetServiveNameOk() (*ServiceName, bool)`

GetServiveNameOk returns a tuple with the ServiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiveName

`func (o *PolicyAssociationRequest) SetServiveName(v ServiceName)`

SetServiveName sets ServiveName field to given value.

### HasServiveName

`func (o *PolicyAssociationRequest) HasServiveName() bool`

HasServiveName returns a boolean if a field has been set.

### GetTraceReq

`func (o *PolicyAssociationRequest) GetTraceReq() TraceData`

GetTraceReq returns the TraceReq field if non-nil, zero value otherwise.

### GetTraceReqOk

`func (o *PolicyAssociationRequest) GetTraceReqOk() (*TraceData, bool)`

GetTraceReqOk returns a tuple with the TraceReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceReq

`func (o *PolicyAssociationRequest) SetTraceReq(v TraceData)`

SetTraceReq sets TraceReq field to given value.

### HasTraceReq

`func (o *PolicyAssociationRequest) HasTraceReq() bool`

HasTraceReq returns a boolean if a field has been set.

### SetTraceReqNil

`func (o *PolicyAssociationRequest) SetTraceReqNil(b bool)`

 SetTraceReqNil sets the value for TraceReq to be an explicit nil

### UnsetTraceReq
`func (o *PolicyAssociationRequest) UnsetTraceReq()`

UnsetTraceReq ensures that no value is present for TraceReq, not even an explicit nil
### GetNwdafDatas

`func (o *PolicyAssociationRequest) GetNwdafDatas() []NwdafData`

GetNwdafDatas returns the NwdafDatas field if non-nil, zero value otherwise.

### GetNwdafDatasOk

`func (o *PolicyAssociationRequest) GetNwdafDatasOk() (*[]NwdafData, bool)`

GetNwdafDatasOk returns a tuple with the NwdafDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafDatas

`func (o *PolicyAssociationRequest) SetNwdafDatas(v []NwdafData)`

SetNwdafDatas sets NwdafDatas field to given value.

### HasNwdafDatas

`func (o *PolicyAssociationRequest) HasNwdafDatas() bool`

HasNwdafDatas returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PolicyAssociationRequest) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PolicyAssociationRequest) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PolicyAssociationRequest) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


