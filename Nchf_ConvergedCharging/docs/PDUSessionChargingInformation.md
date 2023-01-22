# PDUSessionChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargingId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**SMFchargingId** | Pointer to **string** |  | [optional] 
**HomeProvidedChargingId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**SMFHomeProvidedChargingId** | Pointer to **string** |  | [optional] 
**UserInformation** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**UserLocationinfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**MAPDUNon3GPPUserLocationInfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**Non3GPPUserLocationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**MAPDUNon3GPPUserLocationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PresenceReportingAreaInformation** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 
**UetimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**PduSessionInformation** | Pointer to [**PDUSessionInformation**](PDUSessionInformation.md) |  | [optional] 
**UnitCountInactivityTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**RANSecondaryRATUsageReport** | Pointer to [**RANSecondaryRATUsageReport**](RANSecondaryRATUsageReport.md) |  | [optional] 

## Methods

### NewPDUSessionChargingInformation

`func NewPDUSessionChargingInformation() *PDUSessionChargingInformation`

NewPDUSessionChargingInformation instantiates a new PDUSessionChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPDUSessionChargingInformationWithDefaults

`func NewPDUSessionChargingInformationWithDefaults() *PDUSessionChargingInformation`

NewPDUSessionChargingInformationWithDefaults instantiates a new PDUSessionChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargingId

`func (o *PDUSessionChargingInformation) GetChargingId() int32`

GetChargingId returns the ChargingId field if non-nil, zero value otherwise.

### GetChargingIdOk

`func (o *PDUSessionChargingInformation) GetChargingIdOk() (*int32, bool)`

GetChargingIdOk returns a tuple with the ChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingId

`func (o *PDUSessionChargingInformation) SetChargingId(v int32)`

SetChargingId sets ChargingId field to given value.

### HasChargingId

`func (o *PDUSessionChargingInformation) HasChargingId() bool`

HasChargingId returns a boolean if a field has been set.

### GetSMFchargingId

`func (o *PDUSessionChargingInformation) GetSMFchargingId() string`

GetSMFchargingId returns the SMFchargingId field if non-nil, zero value otherwise.

### GetSMFchargingIdOk

`func (o *PDUSessionChargingInformation) GetSMFchargingIdOk() (*string, bool)`

GetSMFchargingIdOk returns a tuple with the SMFchargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMFchargingId

`func (o *PDUSessionChargingInformation) SetSMFchargingId(v string)`

SetSMFchargingId sets SMFchargingId field to given value.

### HasSMFchargingId

`func (o *PDUSessionChargingInformation) HasSMFchargingId() bool`

HasSMFchargingId returns a boolean if a field has been set.

### GetHomeProvidedChargingId

`func (o *PDUSessionChargingInformation) GetHomeProvidedChargingId() int32`

GetHomeProvidedChargingId returns the HomeProvidedChargingId field if non-nil, zero value otherwise.

### GetHomeProvidedChargingIdOk

`func (o *PDUSessionChargingInformation) GetHomeProvidedChargingIdOk() (*int32, bool)`

GetHomeProvidedChargingIdOk returns a tuple with the HomeProvidedChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeProvidedChargingId

`func (o *PDUSessionChargingInformation) SetHomeProvidedChargingId(v int32)`

SetHomeProvidedChargingId sets HomeProvidedChargingId field to given value.

### HasHomeProvidedChargingId

`func (o *PDUSessionChargingInformation) HasHomeProvidedChargingId() bool`

HasHomeProvidedChargingId returns a boolean if a field has been set.

### GetSMFHomeProvidedChargingId

`func (o *PDUSessionChargingInformation) GetSMFHomeProvidedChargingId() string`

GetSMFHomeProvidedChargingId returns the SMFHomeProvidedChargingId field if non-nil, zero value otherwise.

### GetSMFHomeProvidedChargingIdOk

`func (o *PDUSessionChargingInformation) GetSMFHomeProvidedChargingIdOk() (*string, bool)`

GetSMFHomeProvidedChargingIdOk returns a tuple with the SMFHomeProvidedChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMFHomeProvidedChargingId

`func (o *PDUSessionChargingInformation) SetSMFHomeProvidedChargingId(v string)`

SetSMFHomeProvidedChargingId sets SMFHomeProvidedChargingId field to given value.

### HasSMFHomeProvidedChargingId

`func (o *PDUSessionChargingInformation) HasSMFHomeProvidedChargingId() bool`

HasSMFHomeProvidedChargingId returns a boolean if a field has been set.

### GetUserInformation

`func (o *PDUSessionChargingInformation) GetUserInformation() UserInformation`

GetUserInformation returns the UserInformation field if non-nil, zero value otherwise.

### GetUserInformationOk

`func (o *PDUSessionChargingInformation) GetUserInformationOk() (*UserInformation, bool)`

GetUserInformationOk returns a tuple with the UserInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInformation

`func (o *PDUSessionChargingInformation) SetUserInformation(v UserInformation)`

SetUserInformation sets UserInformation field to given value.

### HasUserInformation

`func (o *PDUSessionChargingInformation) HasUserInformation() bool`

HasUserInformation returns a boolean if a field has been set.

### GetUserLocationinfo

`func (o *PDUSessionChargingInformation) GetUserLocationinfo() UserLocation`

GetUserLocationinfo returns the UserLocationinfo field if non-nil, zero value otherwise.

### GetUserLocationinfoOk

`func (o *PDUSessionChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool)`

GetUserLocationinfoOk returns a tuple with the UserLocationinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationinfo

`func (o *PDUSessionChargingInformation) SetUserLocationinfo(v UserLocation)`

SetUserLocationinfo sets UserLocationinfo field to given value.

### HasUserLocationinfo

`func (o *PDUSessionChargingInformation) HasUserLocationinfo() bool`

HasUserLocationinfo returns a boolean if a field has been set.

### GetMAPDUNon3GPPUserLocationInfo

`func (o *PDUSessionChargingInformation) GetMAPDUNon3GPPUserLocationInfo() UserLocation`

GetMAPDUNon3GPPUserLocationInfo returns the MAPDUNon3GPPUserLocationInfo field if non-nil, zero value otherwise.

### GetMAPDUNon3GPPUserLocationInfoOk

`func (o *PDUSessionChargingInformation) GetMAPDUNon3GPPUserLocationInfoOk() (*UserLocation, bool)`

GetMAPDUNon3GPPUserLocationInfoOk returns a tuple with the MAPDUNon3GPPUserLocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAPDUNon3GPPUserLocationInfo

`func (o *PDUSessionChargingInformation) SetMAPDUNon3GPPUserLocationInfo(v UserLocation)`

SetMAPDUNon3GPPUserLocationInfo sets MAPDUNon3GPPUserLocationInfo field to given value.

### HasMAPDUNon3GPPUserLocationInfo

`func (o *PDUSessionChargingInformation) HasMAPDUNon3GPPUserLocationInfo() bool`

HasMAPDUNon3GPPUserLocationInfo returns a boolean if a field has been set.

### GetNon3GPPUserLocationTime

`func (o *PDUSessionChargingInformation) GetNon3GPPUserLocationTime() time.Time`

GetNon3GPPUserLocationTime returns the Non3GPPUserLocationTime field if non-nil, zero value otherwise.

### GetNon3GPPUserLocationTimeOk

`func (o *PDUSessionChargingInformation) GetNon3GPPUserLocationTimeOk() (*time.Time, bool)`

GetNon3GPPUserLocationTimeOk returns a tuple with the Non3GPPUserLocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNon3GPPUserLocationTime

`func (o *PDUSessionChargingInformation) SetNon3GPPUserLocationTime(v time.Time)`

SetNon3GPPUserLocationTime sets Non3GPPUserLocationTime field to given value.

### HasNon3GPPUserLocationTime

`func (o *PDUSessionChargingInformation) HasNon3GPPUserLocationTime() bool`

HasNon3GPPUserLocationTime returns a boolean if a field has been set.

### GetMAPDUNon3GPPUserLocationTime

`func (o *PDUSessionChargingInformation) GetMAPDUNon3GPPUserLocationTime() time.Time`

GetMAPDUNon3GPPUserLocationTime returns the MAPDUNon3GPPUserLocationTime field if non-nil, zero value otherwise.

### GetMAPDUNon3GPPUserLocationTimeOk

`func (o *PDUSessionChargingInformation) GetMAPDUNon3GPPUserLocationTimeOk() (*time.Time, bool)`

GetMAPDUNon3GPPUserLocationTimeOk returns a tuple with the MAPDUNon3GPPUserLocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAPDUNon3GPPUserLocationTime

`func (o *PDUSessionChargingInformation) SetMAPDUNon3GPPUserLocationTime(v time.Time)`

SetMAPDUNon3GPPUserLocationTime sets MAPDUNon3GPPUserLocationTime field to given value.

### HasMAPDUNon3GPPUserLocationTime

`func (o *PDUSessionChargingInformation) HasMAPDUNon3GPPUserLocationTime() bool`

HasMAPDUNon3GPPUserLocationTime returns a boolean if a field has been set.

### GetPresenceReportingAreaInformation

`func (o *PDUSessionChargingInformation) GetPresenceReportingAreaInformation() map[string]PresenceInfo`

GetPresenceReportingAreaInformation returns the PresenceReportingAreaInformation field if non-nil, zero value otherwise.

### GetPresenceReportingAreaInformationOk

`func (o *PDUSessionChargingInformation) GetPresenceReportingAreaInformationOk() (*map[string]PresenceInfo, bool)`

GetPresenceReportingAreaInformationOk returns a tuple with the PresenceReportingAreaInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceReportingAreaInformation

`func (o *PDUSessionChargingInformation) SetPresenceReportingAreaInformation(v map[string]PresenceInfo)`

SetPresenceReportingAreaInformation sets PresenceReportingAreaInformation field to given value.

### HasPresenceReportingAreaInformation

`func (o *PDUSessionChargingInformation) HasPresenceReportingAreaInformation() bool`

HasPresenceReportingAreaInformation returns a boolean if a field has been set.

### GetUetimeZone

`func (o *PDUSessionChargingInformation) GetUetimeZone() string`

GetUetimeZone returns the UetimeZone field if non-nil, zero value otherwise.

### GetUetimeZoneOk

`func (o *PDUSessionChargingInformation) GetUetimeZoneOk() (*string, bool)`

GetUetimeZoneOk returns a tuple with the UetimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetimeZone

`func (o *PDUSessionChargingInformation) SetUetimeZone(v string)`

SetUetimeZone sets UetimeZone field to given value.

### HasUetimeZone

`func (o *PDUSessionChargingInformation) HasUetimeZone() bool`

HasUetimeZone returns a boolean if a field has been set.

### GetPduSessionInformation

`func (o *PDUSessionChargingInformation) GetPduSessionInformation() PDUSessionInformation`

GetPduSessionInformation returns the PduSessionInformation field if non-nil, zero value otherwise.

### GetPduSessionInformationOk

`func (o *PDUSessionChargingInformation) GetPduSessionInformationOk() (*PDUSessionInformation, bool)`

GetPduSessionInformationOk returns a tuple with the PduSessionInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionInformation

`func (o *PDUSessionChargingInformation) SetPduSessionInformation(v PDUSessionInformation)`

SetPduSessionInformation sets PduSessionInformation field to given value.

### HasPduSessionInformation

`func (o *PDUSessionChargingInformation) HasPduSessionInformation() bool`

HasPduSessionInformation returns a boolean if a field has been set.

### GetUnitCountInactivityTimer

`func (o *PDUSessionChargingInformation) GetUnitCountInactivityTimer() int32`

GetUnitCountInactivityTimer returns the UnitCountInactivityTimer field if non-nil, zero value otherwise.

### GetUnitCountInactivityTimerOk

`func (o *PDUSessionChargingInformation) GetUnitCountInactivityTimerOk() (*int32, bool)`

GetUnitCountInactivityTimerOk returns a tuple with the UnitCountInactivityTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCountInactivityTimer

`func (o *PDUSessionChargingInformation) SetUnitCountInactivityTimer(v int32)`

SetUnitCountInactivityTimer sets UnitCountInactivityTimer field to given value.

### HasUnitCountInactivityTimer

`func (o *PDUSessionChargingInformation) HasUnitCountInactivityTimer() bool`

HasUnitCountInactivityTimer returns a boolean if a field has been set.

### GetRANSecondaryRATUsageReport

`func (o *PDUSessionChargingInformation) GetRANSecondaryRATUsageReport() RANSecondaryRATUsageReport`

GetRANSecondaryRATUsageReport returns the RANSecondaryRATUsageReport field if non-nil, zero value otherwise.

### GetRANSecondaryRATUsageReportOk

`func (o *PDUSessionChargingInformation) GetRANSecondaryRATUsageReportOk() (*RANSecondaryRATUsageReport, bool)`

GetRANSecondaryRATUsageReportOk returns a tuple with the RANSecondaryRATUsageReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRANSecondaryRATUsageReport

`func (o *PDUSessionChargingInformation) SetRANSecondaryRATUsageReport(v RANSecondaryRATUsageReport)`

SetRANSecondaryRATUsageReport sets RANSecondaryRATUsageReport field to given value.

### HasRANSecondaryRATUsageReport

`func (o *PDUSessionChargingInformation) HasRANSecondaryRATUsageReport() bool`

HasRANSecondaryRATUsageReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


