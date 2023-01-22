# N2ConnectionChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N2ConnectionMessageType** | **int32** |  | 
**UserInformation** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**UserLocationinfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**PSCellInformation** | Pointer to [**PSCellInformation**](PSCellInformation.md) |  | [optional] 
**UetimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**RATType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**AmfUeNgapId** | Pointer to **int32** |  | [optional] 
**RanUeNgapId** | Pointer to **int32** |  | [optional] 
**RanNodeId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**RestrictedRatList** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**ForbiddenAreaList** | Pointer to [**[]Area**](Area.md) |  | [optional] 
**ServiceAreaRestriction** | Pointer to [**[]ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**RestrictedCnList** | Pointer to [**[]CoreNetworkType**](CoreNetworkType.md) |  | [optional] 
**AllowedNSSAI** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RrcEstCause** | Pointer to **string** |  | [optional] 

## Methods

### NewN2ConnectionChargingInformation

`func NewN2ConnectionChargingInformation(n2ConnectionMessageType int32, ) *N2ConnectionChargingInformation`

NewN2ConnectionChargingInformation instantiates a new N2ConnectionChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2ConnectionChargingInformationWithDefaults

`func NewN2ConnectionChargingInformationWithDefaults() *N2ConnectionChargingInformation`

NewN2ConnectionChargingInformationWithDefaults instantiates a new N2ConnectionChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN2ConnectionMessageType

`func (o *N2ConnectionChargingInformation) GetN2ConnectionMessageType() int32`

GetN2ConnectionMessageType returns the N2ConnectionMessageType field if non-nil, zero value otherwise.

### GetN2ConnectionMessageTypeOk

`func (o *N2ConnectionChargingInformation) GetN2ConnectionMessageTypeOk() (*int32, bool)`

GetN2ConnectionMessageTypeOk returns a tuple with the N2ConnectionMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2ConnectionMessageType

`func (o *N2ConnectionChargingInformation) SetN2ConnectionMessageType(v int32)`

SetN2ConnectionMessageType sets N2ConnectionMessageType field to given value.


### GetUserInformation

`func (o *N2ConnectionChargingInformation) GetUserInformation() UserInformation`

GetUserInformation returns the UserInformation field if non-nil, zero value otherwise.

### GetUserInformationOk

`func (o *N2ConnectionChargingInformation) GetUserInformationOk() (*UserInformation, bool)`

GetUserInformationOk returns a tuple with the UserInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInformation

`func (o *N2ConnectionChargingInformation) SetUserInformation(v UserInformation)`

SetUserInformation sets UserInformation field to given value.

### HasUserInformation

`func (o *N2ConnectionChargingInformation) HasUserInformation() bool`

HasUserInformation returns a boolean if a field has been set.

### GetUserLocationinfo

`func (o *N2ConnectionChargingInformation) GetUserLocationinfo() UserLocation`

GetUserLocationinfo returns the UserLocationinfo field if non-nil, zero value otherwise.

### GetUserLocationinfoOk

`func (o *N2ConnectionChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool)`

GetUserLocationinfoOk returns a tuple with the UserLocationinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationinfo

`func (o *N2ConnectionChargingInformation) SetUserLocationinfo(v UserLocation)`

SetUserLocationinfo sets UserLocationinfo field to given value.

### HasUserLocationinfo

`func (o *N2ConnectionChargingInformation) HasUserLocationinfo() bool`

HasUserLocationinfo returns a boolean if a field has been set.

### GetPSCellInformation

`func (o *N2ConnectionChargingInformation) GetPSCellInformation() PSCellInformation`

GetPSCellInformation returns the PSCellInformation field if non-nil, zero value otherwise.

### GetPSCellInformationOk

`func (o *N2ConnectionChargingInformation) GetPSCellInformationOk() (*PSCellInformation, bool)`

GetPSCellInformationOk returns a tuple with the PSCellInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSCellInformation

`func (o *N2ConnectionChargingInformation) SetPSCellInformation(v PSCellInformation)`

SetPSCellInformation sets PSCellInformation field to given value.

### HasPSCellInformation

`func (o *N2ConnectionChargingInformation) HasPSCellInformation() bool`

HasPSCellInformation returns a boolean if a field has been set.

### GetUetimeZone

`func (o *N2ConnectionChargingInformation) GetUetimeZone() string`

GetUetimeZone returns the UetimeZone field if non-nil, zero value otherwise.

### GetUetimeZoneOk

`func (o *N2ConnectionChargingInformation) GetUetimeZoneOk() (*string, bool)`

GetUetimeZoneOk returns a tuple with the UetimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetimeZone

`func (o *N2ConnectionChargingInformation) SetUetimeZone(v string)`

SetUetimeZone sets UetimeZone field to given value.

### HasUetimeZone

`func (o *N2ConnectionChargingInformation) HasUetimeZone() bool`

HasUetimeZone returns a boolean if a field has been set.

### GetRATType

`func (o *N2ConnectionChargingInformation) GetRATType() RatType`

GetRATType returns the RATType field if non-nil, zero value otherwise.

### GetRATTypeOk

`func (o *N2ConnectionChargingInformation) GetRATTypeOk() (*RatType, bool)`

GetRATTypeOk returns a tuple with the RATType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRATType

`func (o *N2ConnectionChargingInformation) SetRATType(v RatType)`

SetRATType sets RATType field to given value.

### HasRATType

`func (o *N2ConnectionChargingInformation) HasRATType() bool`

HasRATType returns a boolean if a field has been set.

### GetAmfUeNgapId

`func (o *N2ConnectionChargingInformation) GetAmfUeNgapId() int32`

GetAmfUeNgapId returns the AmfUeNgapId field if non-nil, zero value otherwise.

### GetAmfUeNgapIdOk

`func (o *N2ConnectionChargingInformation) GetAmfUeNgapIdOk() (*int32, bool)`

GetAmfUeNgapIdOk returns a tuple with the AmfUeNgapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfUeNgapId

`func (o *N2ConnectionChargingInformation) SetAmfUeNgapId(v int32)`

SetAmfUeNgapId sets AmfUeNgapId field to given value.

### HasAmfUeNgapId

`func (o *N2ConnectionChargingInformation) HasAmfUeNgapId() bool`

HasAmfUeNgapId returns a boolean if a field has been set.

### GetRanUeNgapId

`func (o *N2ConnectionChargingInformation) GetRanUeNgapId() int32`

GetRanUeNgapId returns the RanUeNgapId field if non-nil, zero value otherwise.

### GetRanUeNgapIdOk

`func (o *N2ConnectionChargingInformation) GetRanUeNgapIdOk() (*int32, bool)`

GetRanUeNgapIdOk returns a tuple with the RanUeNgapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUeNgapId

`func (o *N2ConnectionChargingInformation) SetRanUeNgapId(v int32)`

SetRanUeNgapId sets RanUeNgapId field to given value.

### HasRanUeNgapId

`func (o *N2ConnectionChargingInformation) HasRanUeNgapId() bool`

HasRanUeNgapId returns a boolean if a field has been set.

### GetRanNodeId

`func (o *N2ConnectionChargingInformation) GetRanNodeId() GlobalRanNodeId`

GetRanNodeId returns the RanNodeId field if non-nil, zero value otherwise.

### GetRanNodeIdOk

`func (o *N2ConnectionChargingInformation) GetRanNodeIdOk() (*GlobalRanNodeId, bool)`

GetRanNodeIdOk returns a tuple with the RanNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNodeId

`func (o *N2ConnectionChargingInformation) SetRanNodeId(v GlobalRanNodeId)`

SetRanNodeId sets RanNodeId field to given value.

### HasRanNodeId

`func (o *N2ConnectionChargingInformation) HasRanNodeId() bool`

HasRanNodeId returns a boolean if a field has been set.

### GetRestrictedRatList

`func (o *N2ConnectionChargingInformation) GetRestrictedRatList() []RatType`

GetRestrictedRatList returns the RestrictedRatList field if non-nil, zero value otherwise.

### GetRestrictedRatListOk

`func (o *N2ConnectionChargingInformation) GetRestrictedRatListOk() (*[]RatType, bool)`

GetRestrictedRatListOk returns a tuple with the RestrictedRatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedRatList

`func (o *N2ConnectionChargingInformation) SetRestrictedRatList(v []RatType)`

SetRestrictedRatList sets RestrictedRatList field to given value.

### HasRestrictedRatList

`func (o *N2ConnectionChargingInformation) HasRestrictedRatList() bool`

HasRestrictedRatList returns a boolean if a field has been set.

### GetForbiddenAreaList

`func (o *N2ConnectionChargingInformation) GetForbiddenAreaList() []Area`

GetForbiddenAreaList returns the ForbiddenAreaList field if non-nil, zero value otherwise.

### GetForbiddenAreaListOk

`func (o *N2ConnectionChargingInformation) GetForbiddenAreaListOk() (*[]Area, bool)`

GetForbiddenAreaListOk returns a tuple with the ForbiddenAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAreaList

`func (o *N2ConnectionChargingInformation) SetForbiddenAreaList(v []Area)`

SetForbiddenAreaList sets ForbiddenAreaList field to given value.

### HasForbiddenAreaList

`func (o *N2ConnectionChargingInformation) HasForbiddenAreaList() bool`

HasForbiddenAreaList returns a boolean if a field has been set.

### GetServiceAreaRestriction

`func (o *N2ConnectionChargingInformation) GetServiceAreaRestriction() []ServiceAreaRestriction`

GetServiceAreaRestriction returns the ServiceAreaRestriction field if non-nil, zero value otherwise.

### GetServiceAreaRestrictionOk

`func (o *N2ConnectionChargingInformation) GetServiceAreaRestrictionOk() (*[]ServiceAreaRestriction, bool)`

GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreaRestriction

`func (o *N2ConnectionChargingInformation) SetServiceAreaRestriction(v []ServiceAreaRestriction)`

SetServiceAreaRestriction sets ServiceAreaRestriction field to given value.

### HasServiceAreaRestriction

`func (o *N2ConnectionChargingInformation) HasServiceAreaRestriction() bool`

HasServiceAreaRestriction returns a boolean if a field has been set.

### GetRestrictedCnList

`func (o *N2ConnectionChargingInformation) GetRestrictedCnList() []CoreNetworkType`

GetRestrictedCnList returns the RestrictedCnList field if non-nil, zero value otherwise.

### GetRestrictedCnListOk

`func (o *N2ConnectionChargingInformation) GetRestrictedCnListOk() (*[]CoreNetworkType, bool)`

GetRestrictedCnListOk returns a tuple with the RestrictedCnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCnList

`func (o *N2ConnectionChargingInformation) SetRestrictedCnList(v []CoreNetworkType)`

SetRestrictedCnList sets RestrictedCnList field to given value.

### HasRestrictedCnList

`func (o *N2ConnectionChargingInformation) HasRestrictedCnList() bool`

HasRestrictedCnList returns a boolean if a field has been set.

### GetAllowedNSSAI

`func (o *N2ConnectionChargingInformation) GetAllowedNSSAI() []Snssai`

GetAllowedNSSAI returns the AllowedNSSAI field if non-nil, zero value otherwise.

### GetAllowedNSSAIOk

`func (o *N2ConnectionChargingInformation) GetAllowedNSSAIOk() (*[]Snssai, bool)`

GetAllowedNSSAIOk returns a tuple with the AllowedNSSAI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNSSAI

`func (o *N2ConnectionChargingInformation) SetAllowedNSSAI(v []Snssai)`

SetAllowedNSSAI sets AllowedNSSAI field to given value.

### HasAllowedNSSAI

`func (o *N2ConnectionChargingInformation) HasAllowedNSSAI() bool`

HasAllowedNSSAI returns a boolean if a field has been set.

### GetRrcEstCause

`func (o *N2ConnectionChargingInformation) GetRrcEstCause() string`

GetRrcEstCause returns the RrcEstCause field if non-nil, zero value otherwise.

### GetRrcEstCauseOk

`func (o *N2ConnectionChargingInformation) GetRrcEstCauseOk() (*string, bool)`

GetRrcEstCauseOk returns a tuple with the RrcEstCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrcEstCause

`func (o *N2ConnectionChargingInformation) SetRrcEstCause(v string)`

SetRrcEstCause sets RrcEstCause field to given value.

### HasRrcEstCause

`func (o *N2ConnectionChargingInformation) HasRrcEstCause() bool`

HasRrcEstCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


