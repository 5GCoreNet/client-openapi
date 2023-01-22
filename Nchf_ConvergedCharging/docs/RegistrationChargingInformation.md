# RegistrationChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationMessagetype** | [**RegistrationMessageType**](RegistrationMessageType.md) |  | 
**UserInformation** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**UserLocationinfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**PSCellInformation** | Pointer to [**PSCellInformation**](PSCellInformation.md) |  | [optional] 
**UetimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**RATType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**Var5GMMCapability** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**MICOModeIndication** | Pointer to [**MICOModeIndication**](MICOModeIndication.md) |  | [optional] 
**SmsIndication** | Pointer to [**SmsIndication**](SmsIndication.md) |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**ServiceAreaRestriction** | Pointer to [**[]ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**RequestedNSSAI** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**AllowedNSSAI** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RejectedNSSAI** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**NSSAIMapList** | Pointer to [**[]NSSAIMap**](NSSAIMap.md) |  | [optional] 
**AmfUeNgapId** | Pointer to **int32** |  | [optional] 
**RanUeNgapId** | Pointer to **int32** |  | [optional] 
**RanNodeId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 

## Methods

### NewRegistrationChargingInformation

`func NewRegistrationChargingInformation(registrationMessagetype RegistrationMessageType, ) *RegistrationChargingInformation`

NewRegistrationChargingInformation instantiates a new RegistrationChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationChargingInformationWithDefaults

`func NewRegistrationChargingInformationWithDefaults() *RegistrationChargingInformation`

NewRegistrationChargingInformationWithDefaults instantiates a new RegistrationChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationMessagetype

`func (o *RegistrationChargingInformation) GetRegistrationMessagetype() RegistrationMessageType`

GetRegistrationMessagetype returns the RegistrationMessagetype field if non-nil, zero value otherwise.

### GetRegistrationMessagetypeOk

`func (o *RegistrationChargingInformation) GetRegistrationMessagetypeOk() (*RegistrationMessageType, bool)`

GetRegistrationMessagetypeOk returns a tuple with the RegistrationMessagetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationMessagetype

`func (o *RegistrationChargingInformation) SetRegistrationMessagetype(v RegistrationMessageType)`

SetRegistrationMessagetype sets RegistrationMessagetype field to given value.


### GetUserInformation

`func (o *RegistrationChargingInformation) GetUserInformation() UserInformation`

GetUserInformation returns the UserInformation field if non-nil, zero value otherwise.

### GetUserInformationOk

`func (o *RegistrationChargingInformation) GetUserInformationOk() (*UserInformation, bool)`

GetUserInformationOk returns a tuple with the UserInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInformation

`func (o *RegistrationChargingInformation) SetUserInformation(v UserInformation)`

SetUserInformation sets UserInformation field to given value.

### HasUserInformation

`func (o *RegistrationChargingInformation) HasUserInformation() bool`

HasUserInformation returns a boolean if a field has been set.

### GetUserLocationinfo

`func (o *RegistrationChargingInformation) GetUserLocationinfo() UserLocation`

GetUserLocationinfo returns the UserLocationinfo field if non-nil, zero value otherwise.

### GetUserLocationinfoOk

`func (o *RegistrationChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool)`

GetUserLocationinfoOk returns a tuple with the UserLocationinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationinfo

`func (o *RegistrationChargingInformation) SetUserLocationinfo(v UserLocation)`

SetUserLocationinfo sets UserLocationinfo field to given value.

### HasUserLocationinfo

`func (o *RegistrationChargingInformation) HasUserLocationinfo() bool`

HasUserLocationinfo returns a boolean if a field has been set.

### GetPSCellInformation

`func (o *RegistrationChargingInformation) GetPSCellInformation() PSCellInformation`

GetPSCellInformation returns the PSCellInformation field if non-nil, zero value otherwise.

### GetPSCellInformationOk

`func (o *RegistrationChargingInformation) GetPSCellInformationOk() (*PSCellInformation, bool)`

GetPSCellInformationOk returns a tuple with the PSCellInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSCellInformation

`func (o *RegistrationChargingInformation) SetPSCellInformation(v PSCellInformation)`

SetPSCellInformation sets PSCellInformation field to given value.

### HasPSCellInformation

`func (o *RegistrationChargingInformation) HasPSCellInformation() bool`

HasPSCellInformation returns a boolean if a field has been set.

### GetUetimeZone

`func (o *RegistrationChargingInformation) GetUetimeZone() string`

GetUetimeZone returns the UetimeZone field if non-nil, zero value otherwise.

### GetUetimeZoneOk

`func (o *RegistrationChargingInformation) GetUetimeZoneOk() (*string, bool)`

GetUetimeZoneOk returns a tuple with the UetimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetimeZone

`func (o *RegistrationChargingInformation) SetUetimeZone(v string)`

SetUetimeZone sets UetimeZone field to given value.

### HasUetimeZone

`func (o *RegistrationChargingInformation) HasUetimeZone() bool`

HasUetimeZone returns a boolean if a field has been set.

### GetRATType

`func (o *RegistrationChargingInformation) GetRATType() RatType`

GetRATType returns the RATType field if non-nil, zero value otherwise.

### GetRATTypeOk

`func (o *RegistrationChargingInformation) GetRATTypeOk() (*RatType, bool)`

GetRATTypeOk returns a tuple with the RATType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRATType

`func (o *RegistrationChargingInformation) SetRATType(v RatType)`

SetRATType sets RATType field to given value.

### HasRATType

`func (o *RegistrationChargingInformation) HasRATType() bool`

HasRATType returns a boolean if a field has been set.

### GetVar5GMMCapability

`func (o *RegistrationChargingInformation) GetVar5GMMCapability() string`

GetVar5GMMCapability returns the Var5GMMCapability field if non-nil, zero value otherwise.

### GetVar5GMMCapabilityOk

`func (o *RegistrationChargingInformation) GetVar5GMMCapabilityOk() (*string, bool)`

GetVar5GMMCapabilityOk returns a tuple with the Var5GMMCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5GMMCapability

`func (o *RegistrationChargingInformation) SetVar5GMMCapability(v string)`

SetVar5GMMCapability sets Var5GMMCapability field to given value.

### HasVar5GMMCapability

`func (o *RegistrationChargingInformation) HasVar5GMMCapability() bool`

HasVar5GMMCapability returns a boolean if a field has been set.

### GetMICOModeIndication

`func (o *RegistrationChargingInformation) GetMICOModeIndication() MICOModeIndication`

GetMICOModeIndication returns the MICOModeIndication field if non-nil, zero value otherwise.

### GetMICOModeIndicationOk

`func (o *RegistrationChargingInformation) GetMICOModeIndicationOk() (*MICOModeIndication, bool)`

GetMICOModeIndicationOk returns a tuple with the MICOModeIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMICOModeIndication

`func (o *RegistrationChargingInformation) SetMICOModeIndication(v MICOModeIndication)`

SetMICOModeIndication sets MICOModeIndication field to given value.

### HasMICOModeIndication

`func (o *RegistrationChargingInformation) HasMICOModeIndication() bool`

HasMICOModeIndication returns a boolean if a field has been set.

### GetSmsIndication

`func (o *RegistrationChargingInformation) GetSmsIndication() SmsIndication`

GetSmsIndication returns the SmsIndication field if non-nil, zero value otherwise.

### GetSmsIndicationOk

`func (o *RegistrationChargingInformation) GetSmsIndicationOk() (*SmsIndication, bool)`

GetSmsIndicationOk returns a tuple with the SmsIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsIndication

`func (o *RegistrationChargingInformation) SetSmsIndication(v SmsIndication)`

SetSmsIndication sets SmsIndication field to given value.

### HasSmsIndication

`func (o *RegistrationChargingInformation) HasSmsIndication() bool`

HasSmsIndication returns a boolean if a field has been set.

### GetTaiList

`func (o *RegistrationChargingInformation) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *RegistrationChargingInformation) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *RegistrationChargingInformation) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *RegistrationChargingInformation) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetServiceAreaRestriction

`func (o *RegistrationChargingInformation) GetServiceAreaRestriction() []ServiceAreaRestriction`

GetServiceAreaRestriction returns the ServiceAreaRestriction field if non-nil, zero value otherwise.

### GetServiceAreaRestrictionOk

`func (o *RegistrationChargingInformation) GetServiceAreaRestrictionOk() (*[]ServiceAreaRestriction, bool)`

GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreaRestriction

`func (o *RegistrationChargingInformation) SetServiceAreaRestriction(v []ServiceAreaRestriction)`

SetServiceAreaRestriction sets ServiceAreaRestriction field to given value.

### HasServiceAreaRestriction

`func (o *RegistrationChargingInformation) HasServiceAreaRestriction() bool`

HasServiceAreaRestriction returns a boolean if a field has been set.

### GetRequestedNSSAI

`func (o *RegistrationChargingInformation) GetRequestedNSSAI() []Snssai`

GetRequestedNSSAI returns the RequestedNSSAI field if non-nil, zero value otherwise.

### GetRequestedNSSAIOk

`func (o *RegistrationChargingInformation) GetRequestedNSSAIOk() (*[]Snssai, bool)`

GetRequestedNSSAIOk returns a tuple with the RequestedNSSAI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedNSSAI

`func (o *RegistrationChargingInformation) SetRequestedNSSAI(v []Snssai)`

SetRequestedNSSAI sets RequestedNSSAI field to given value.

### HasRequestedNSSAI

`func (o *RegistrationChargingInformation) HasRequestedNSSAI() bool`

HasRequestedNSSAI returns a boolean if a field has been set.

### GetAllowedNSSAI

`func (o *RegistrationChargingInformation) GetAllowedNSSAI() []Snssai`

GetAllowedNSSAI returns the AllowedNSSAI field if non-nil, zero value otherwise.

### GetAllowedNSSAIOk

`func (o *RegistrationChargingInformation) GetAllowedNSSAIOk() (*[]Snssai, bool)`

GetAllowedNSSAIOk returns a tuple with the AllowedNSSAI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNSSAI

`func (o *RegistrationChargingInformation) SetAllowedNSSAI(v []Snssai)`

SetAllowedNSSAI sets AllowedNSSAI field to given value.

### HasAllowedNSSAI

`func (o *RegistrationChargingInformation) HasAllowedNSSAI() bool`

HasAllowedNSSAI returns a boolean if a field has been set.

### GetRejectedNSSAI

`func (o *RegistrationChargingInformation) GetRejectedNSSAI() []Snssai`

GetRejectedNSSAI returns the RejectedNSSAI field if non-nil, zero value otherwise.

### GetRejectedNSSAIOk

`func (o *RegistrationChargingInformation) GetRejectedNSSAIOk() (*[]Snssai, bool)`

GetRejectedNSSAIOk returns a tuple with the RejectedNSSAI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNSSAI

`func (o *RegistrationChargingInformation) SetRejectedNSSAI(v []Snssai)`

SetRejectedNSSAI sets RejectedNSSAI field to given value.

### HasRejectedNSSAI

`func (o *RegistrationChargingInformation) HasRejectedNSSAI() bool`

HasRejectedNSSAI returns a boolean if a field has been set.

### GetNSSAIMapList

`func (o *RegistrationChargingInformation) GetNSSAIMapList() []NSSAIMap`

GetNSSAIMapList returns the NSSAIMapList field if non-nil, zero value otherwise.

### GetNSSAIMapListOk

`func (o *RegistrationChargingInformation) GetNSSAIMapListOk() (*[]NSSAIMap, bool)`

GetNSSAIMapListOk returns a tuple with the NSSAIMapList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSSAIMapList

`func (o *RegistrationChargingInformation) SetNSSAIMapList(v []NSSAIMap)`

SetNSSAIMapList sets NSSAIMapList field to given value.

### HasNSSAIMapList

`func (o *RegistrationChargingInformation) HasNSSAIMapList() bool`

HasNSSAIMapList returns a boolean if a field has been set.

### GetAmfUeNgapId

`func (o *RegistrationChargingInformation) GetAmfUeNgapId() int32`

GetAmfUeNgapId returns the AmfUeNgapId field if non-nil, zero value otherwise.

### GetAmfUeNgapIdOk

`func (o *RegistrationChargingInformation) GetAmfUeNgapIdOk() (*int32, bool)`

GetAmfUeNgapIdOk returns a tuple with the AmfUeNgapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfUeNgapId

`func (o *RegistrationChargingInformation) SetAmfUeNgapId(v int32)`

SetAmfUeNgapId sets AmfUeNgapId field to given value.

### HasAmfUeNgapId

`func (o *RegistrationChargingInformation) HasAmfUeNgapId() bool`

HasAmfUeNgapId returns a boolean if a field has been set.

### GetRanUeNgapId

`func (o *RegistrationChargingInformation) GetRanUeNgapId() int32`

GetRanUeNgapId returns the RanUeNgapId field if non-nil, zero value otherwise.

### GetRanUeNgapIdOk

`func (o *RegistrationChargingInformation) GetRanUeNgapIdOk() (*int32, bool)`

GetRanUeNgapIdOk returns a tuple with the RanUeNgapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUeNgapId

`func (o *RegistrationChargingInformation) SetRanUeNgapId(v int32)`

SetRanUeNgapId sets RanUeNgapId field to given value.

### HasRanUeNgapId

`func (o *RegistrationChargingInformation) HasRanUeNgapId() bool`

HasRanUeNgapId returns a boolean if a field has been set.

### GetRanNodeId

`func (o *RegistrationChargingInformation) GetRanNodeId() GlobalRanNodeId`

GetRanNodeId returns the RanNodeId field if non-nil, zero value otherwise.

### GetRanNodeIdOk

`func (o *RegistrationChargingInformation) GetRanNodeIdOk() (*GlobalRanNodeId, bool)`

GetRanNodeIdOk returns a tuple with the RanNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNodeId

`func (o *RegistrationChargingInformation) SetRanNodeId(v GlobalRanNodeId)`

SetRanNodeId sets RanNodeId field to given value.

### HasRanNodeId

`func (o *RegistrationChargingInformation) HasRanNodeId() bool`

HasRanNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


