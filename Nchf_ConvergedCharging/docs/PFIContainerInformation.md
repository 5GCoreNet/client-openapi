# PFIContainerInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PFI** | Pointer to **string** |  | [optional] 
**ReportTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeofFirstUsage** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeofLastUsage** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**QoSInformation** | Pointer to [**NullableQosData**](QosData.md) |  | [optional] 
**QoSCharacteristics** | Pointer to [**QosCharacteristics**](QosCharacteristics.md) |  | [optional] 
**UserLocationInformation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UetimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**PresenceReportingAreaInformation** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 

## Methods

### NewPFIContainerInformation

`func NewPFIContainerInformation() *PFIContainerInformation`

NewPFIContainerInformation instantiates a new PFIContainerInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPFIContainerInformationWithDefaults

`func NewPFIContainerInformationWithDefaults() *PFIContainerInformation`

NewPFIContainerInformationWithDefaults instantiates a new PFIContainerInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPFI

`func (o *PFIContainerInformation) GetPFI() string`

GetPFI returns the PFI field if non-nil, zero value otherwise.

### GetPFIOk

`func (o *PFIContainerInformation) GetPFIOk() (*string, bool)`

GetPFIOk returns a tuple with the PFI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPFI

`func (o *PFIContainerInformation) SetPFI(v string)`

SetPFI sets PFI field to given value.

### HasPFI

`func (o *PFIContainerInformation) HasPFI() bool`

HasPFI returns a boolean if a field has been set.

### GetReportTime

`func (o *PFIContainerInformation) GetReportTime() time.Time`

GetReportTime returns the ReportTime field if non-nil, zero value otherwise.

### GetReportTimeOk

`func (o *PFIContainerInformation) GetReportTimeOk() (*time.Time, bool)`

GetReportTimeOk returns a tuple with the ReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTime

`func (o *PFIContainerInformation) SetReportTime(v time.Time)`

SetReportTime sets ReportTime field to given value.

### HasReportTime

`func (o *PFIContainerInformation) HasReportTime() bool`

HasReportTime returns a boolean if a field has been set.

### GetTimeofFirstUsage

`func (o *PFIContainerInformation) GetTimeofFirstUsage() time.Time`

GetTimeofFirstUsage returns the TimeofFirstUsage field if non-nil, zero value otherwise.

### GetTimeofFirstUsageOk

`func (o *PFIContainerInformation) GetTimeofFirstUsageOk() (*time.Time, bool)`

GetTimeofFirstUsageOk returns a tuple with the TimeofFirstUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeofFirstUsage

`func (o *PFIContainerInformation) SetTimeofFirstUsage(v time.Time)`

SetTimeofFirstUsage sets TimeofFirstUsage field to given value.

### HasTimeofFirstUsage

`func (o *PFIContainerInformation) HasTimeofFirstUsage() bool`

HasTimeofFirstUsage returns a boolean if a field has been set.

### GetTimeofLastUsage

`func (o *PFIContainerInformation) GetTimeofLastUsage() time.Time`

GetTimeofLastUsage returns the TimeofLastUsage field if non-nil, zero value otherwise.

### GetTimeofLastUsageOk

`func (o *PFIContainerInformation) GetTimeofLastUsageOk() (*time.Time, bool)`

GetTimeofLastUsageOk returns a tuple with the TimeofLastUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeofLastUsage

`func (o *PFIContainerInformation) SetTimeofLastUsage(v time.Time)`

SetTimeofLastUsage sets TimeofLastUsage field to given value.

### HasTimeofLastUsage

`func (o *PFIContainerInformation) HasTimeofLastUsage() bool`

HasTimeofLastUsage returns a boolean if a field has been set.

### GetQoSInformation

`func (o *PFIContainerInformation) GetQoSInformation() QosData`

GetQoSInformation returns the QoSInformation field if non-nil, zero value otherwise.

### GetQoSInformationOk

`func (o *PFIContainerInformation) GetQoSInformationOk() (*QosData, bool)`

GetQoSInformationOk returns a tuple with the QoSInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoSInformation

`func (o *PFIContainerInformation) SetQoSInformation(v QosData)`

SetQoSInformation sets QoSInformation field to given value.

### HasQoSInformation

`func (o *PFIContainerInformation) HasQoSInformation() bool`

HasQoSInformation returns a boolean if a field has been set.

### SetQoSInformationNil

`func (o *PFIContainerInformation) SetQoSInformationNil(b bool)`

 SetQoSInformationNil sets the value for QoSInformation to be an explicit nil

### UnsetQoSInformation
`func (o *PFIContainerInformation) UnsetQoSInformation()`

UnsetQoSInformation ensures that no value is present for QoSInformation, not even an explicit nil
### GetQoSCharacteristics

`func (o *PFIContainerInformation) GetQoSCharacteristics() QosCharacteristics`

GetQoSCharacteristics returns the QoSCharacteristics field if non-nil, zero value otherwise.

### GetQoSCharacteristicsOk

`func (o *PFIContainerInformation) GetQoSCharacteristicsOk() (*QosCharacteristics, bool)`

GetQoSCharacteristicsOk returns a tuple with the QoSCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoSCharacteristics

`func (o *PFIContainerInformation) SetQoSCharacteristics(v QosCharacteristics)`

SetQoSCharacteristics sets QoSCharacteristics field to given value.

### HasQoSCharacteristics

`func (o *PFIContainerInformation) HasQoSCharacteristics() bool`

HasQoSCharacteristics returns a boolean if a field has been set.

### GetUserLocationInformation

`func (o *PFIContainerInformation) GetUserLocationInformation() UserLocation`

GetUserLocationInformation returns the UserLocationInformation field if non-nil, zero value otherwise.

### GetUserLocationInformationOk

`func (o *PFIContainerInformation) GetUserLocationInformationOk() (*UserLocation, bool)`

GetUserLocationInformationOk returns a tuple with the UserLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInformation

`func (o *PFIContainerInformation) SetUserLocationInformation(v UserLocation)`

SetUserLocationInformation sets UserLocationInformation field to given value.

### HasUserLocationInformation

`func (o *PFIContainerInformation) HasUserLocationInformation() bool`

HasUserLocationInformation returns a boolean if a field has been set.

### GetUetimeZone

`func (o *PFIContainerInformation) GetUetimeZone() string`

GetUetimeZone returns the UetimeZone field if non-nil, zero value otherwise.

### GetUetimeZoneOk

`func (o *PFIContainerInformation) GetUetimeZoneOk() (*string, bool)`

GetUetimeZoneOk returns a tuple with the UetimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetimeZone

`func (o *PFIContainerInformation) SetUetimeZone(v string)`

SetUetimeZone sets UetimeZone field to given value.

### HasUetimeZone

`func (o *PFIContainerInformation) HasUetimeZone() bool`

HasUetimeZone returns a boolean if a field has been set.

### GetPresenceReportingAreaInformation

`func (o *PFIContainerInformation) GetPresenceReportingAreaInformation() map[string]PresenceInfo`

GetPresenceReportingAreaInformation returns the PresenceReportingAreaInformation field if non-nil, zero value otherwise.

### GetPresenceReportingAreaInformationOk

`func (o *PFIContainerInformation) GetPresenceReportingAreaInformationOk() (*map[string]PresenceInfo, bool)`

GetPresenceReportingAreaInformationOk returns a tuple with the PresenceReportingAreaInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceReportingAreaInformation

`func (o *PFIContainerInformation) SetPresenceReportingAreaInformation(v map[string]PresenceInfo)`

SetPresenceReportingAreaInformation sets PresenceReportingAreaInformation field to given value.

### HasPresenceReportingAreaInformation

`func (o *PFIContainerInformation) HasPresenceReportingAreaInformation() bool`

HasPresenceReportingAreaInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


