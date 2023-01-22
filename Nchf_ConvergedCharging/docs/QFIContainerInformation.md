# QFIContainerInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QFI** | Pointer to **int32** | Unsigned integer identifying a QoS flow, within the range 0 to 63. | [optional] 
**ReportTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**TimeofFirstUsage** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeofLastUsage** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**QoSInformation** | Pointer to [**NullableQosData**](QosData.md) |  | [optional] 
**QoSCharacteristics** | Pointer to [**QosCharacteristics**](QosCharacteristics.md) |  | [optional] 
**UserLocationInformation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UetimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**PresenceReportingAreaInformation** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 
**RATType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**ServingNetworkFunctionID** | Pointer to [**[]ServingNetworkFunctionID**](ServingNetworkFunctionID.md) |  | [optional] 
**Var3gppPSDataOffStatus** | Pointer to [**Model3GPPPSDataOffStatus**](Model3GPPPSDataOffStatus.md) |  | [optional] 
**Var3gppChargingId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**Diagnostics** | Pointer to **int32** |  | [optional] 
**EnhancedDiagnostics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewQFIContainerInformation

`func NewQFIContainerInformation(reportTime time.Time, ) *QFIContainerInformation`

NewQFIContainerInformation instantiates a new QFIContainerInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQFIContainerInformationWithDefaults

`func NewQFIContainerInformationWithDefaults() *QFIContainerInformation`

NewQFIContainerInformationWithDefaults instantiates a new QFIContainerInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQFI

`func (o *QFIContainerInformation) GetQFI() int32`

GetQFI returns the QFI field if non-nil, zero value otherwise.

### GetQFIOk

`func (o *QFIContainerInformation) GetQFIOk() (*int32, bool)`

GetQFIOk returns a tuple with the QFI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQFI

`func (o *QFIContainerInformation) SetQFI(v int32)`

SetQFI sets QFI field to given value.

### HasQFI

`func (o *QFIContainerInformation) HasQFI() bool`

HasQFI returns a boolean if a field has been set.

### GetReportTime

`func (o *QFIContainerInformation) GetReportTime() time.Time`

GetReportTime returns the ReportTime field if non-nil, zero value otherwise.

### GetReportTimeOk

`func (o *QFIContainerInformation) GetReportTimeOk() (*time.Time, bool)`

GetReportTimeOk returns a tuple with the ReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTime

`func (o *QFIContainerInformation) SetReportTime(v time.Time)`

SetReportTime sets ReportTime field to given value.


### GetTimeofFirstUsage

`func (o *QFIContainerInformation) GetTimeofFirstUsage() time.Time`

GetTimeofFirstUsage returns the TimeofFirstUsage field if non-nil, zero value otherwise.

### GetTimeofFirstUsageOk

`func (o *QFIContainerInformation) GetTimeofFirstUsageOk() (*time.Time, bool)`

GetTimeofFirstUsageOk returns a tuple with the TimeofFirstUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeofFirstUsage

`func (o *QFIContainerInformation) SetTimeofFirstUsage(v time.Time)`

SetTimeofFirstUsage sets TimeofFirstUsage field to given value.

### HasTimeofFirstUsage

`func (o *QFIContainerInformation) HasTimeofFirstUsage() bool`

HasTimeofFirstUsage returns a boolean if a field has been set.

### GetTimeofLastUsage

`func (o *QFIContainerInformation) GetTimeofLastUsage() time.Time`

GetTimeofLastUsage returns the TimeofLastUsage field if non-nil, zero value otherwise.

### GetTimeofLastUsageOk

`func (o *QFIContainerInformation) GetTimeofLastUsageOk() (*time.Time, bool)`

GetTimeofLastUsageOk returns a tuple with the TimeofLastUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeofLastUsage

`func (o *QFIContainerInformation) SetTimeofLastUsage(v time.Time)`

SetTimeofLastUsage sets TimeofLastUsage field to given value.

### HasTimeofLastUsage

`func (o *QFIContainerInformation) HasTimeofLastUsage() bool`

HasTimeofLastUsage returns a boolean if a field has been set.

### GetQoSInformation

`func (o *QFIContainerInformation) GetQoSInformation() QosData`

GetQoSInformation returns the QoSInformation field if non-nil, zero value otherwise.

### GetQoSInformationOk

`func (o *QFIContainerInformation) GetQoSInformationOk() (*QosData, bool)`

GetQoSInformationOk returns a tuple with the QoSInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoSInformation

`func (o *QFIContainerInformation) SetQoSInformation(v QosData)`

SetQoSInformation sets QoSInformation field to given value.

### HasQoSInformation

`func (o *QFIContainerInformation) HasQoSInformation() bool`

HasQoSInformation returns a boolean if a field has been set.

### SetQoSInformationNil

`func (o *QFIContainerInformation) SetQoSInformationNil(b bool)`

 SetQoSInformationNil sets the value for QoSInformation to be an explicit nil

### UnsetQoSInformation
`func (o *QFIContainerInformation) UnsetQoSInformation()`

UnsetQoSInformation ensures that no value is present for QoSInformation, not even an explicit nil
### GetQoSCharacteristics

`func (o *QFIContainerInformation) GetQoSCharacteristics() QosCharacteristics`

GetQoSCharacteristics returns the QoSCharacteristics field if non-nil, zero value otherwise.

### GetQoSCharacteristicsOk

`func (o *QFIContainerInformation) GetQoSCharacteristicsOk() (*QosCharacteristics, bool)`

GetQoSCharacteristicsOk returns a tuple with the QoSCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoSCharacteristics

`func (o *QFIContainerInformation) SetQoSCharacteristics(v QosCharacteristics)`

SetQoSCharacteristics sets QoSCharacteristics field to given value.

### HasQoSCharacteristics

`func (o *QFIContainerInformation) HasQoSCharacteristics() bool`

HasQoSCharacteristics returns a boolean if a field has been set.

### GetUserLocationInformation

`func (o *QFIContainerInformation) GetUserLocationInformation() UserLocation`

GetUserLocationInformation returns the UserLocationInformation field if non-nil, zero value otherwise.

### GetUserLocationInformationOk

`func (o *QFIContainerInformation) GetUserLocationInformationOk() (*UserLocation, bool)`

GetUserLocationInformationOk returns a tuple with the UserLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInformation

`func (o *QFIContainerInformation) SetUserLocationInformation(v UserLocation)`

SetUserLocationInformation sets UserLocationInformation field to given value.

### HasUserLocationInformation

`func (o *QFIContainerInformation) HasUserLocationInformation() bool`

HasUserLocationInformation returns a boolean if a field has been set.

### GetUetimeZone

`func (o *QFIContainerInformation) GetUetimeZone() string`

GetUetimeZone returns the UetimeZone field if non-nil, zero value otherwise.

### GetUetimeZoneOk

`func (o *QFIContainerInformation) GetUetimeZoneOk() (*string, bool)`

GetUetimeZoneOk returns a tuple with the UetimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetimeZone

`func (o *QFIContainerInformation) SetUetimeZone(v string)`

SetUetimeZone sets UetimeZone field to given value.

### HasUetimeZone

`func (o *QFIContainerInformation) HasUetimeZone() bool`

HasUetimeZone returns a boolean if a field has been set.

### GetPresenceReportingAreaInformation

`func (o *QFIContainerInformation) GetPresenceReportingAreaInformation() map[string]PresenceInfo`

GetPresenceReportingAreaInformation returns the PresenceReportingAreaInformation field if non-nil, zero value otherwise.

### GetPresenceReportingAreaInformationOk

`func (o *QFIContainerInformation) GetPresenceReportingAreaInformationOk() (*map[string]PresenceInfo, bool)`

GetPresenceReportingAreaInformationOk returns a tuple with the PresenceReportingAreaInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceReportingAreaInformation

`func (o *QFIContainerInformation) SetPresenceReportingAreaInformation(v map[string]PresenceInfo)`

SetPresenceReportingAreaInformation sets PresenceReportingAreaInformation field to given value.

### HasPresenceReportingAreaInformation

`func (o *QFIContainerInformation) HasPresenceReportingAreaInformation() bool`

HasPresenceReportingAreaInformation returns a boolean if a field has been set.

### GetRATType

`func (o *QFIContainerInformation) GetRATType() RatType`

GetRATType returns the RATType field if non-nil, zero value otherwise.

### GetRATTypeOk

`func (o *QFIContainerInformation) GetRATTypeOk() (*RatType, bool)`

GetRATTypeOk returns a tuple with the RATType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRATType

`func (o *QFIContainerInformation) SetRATType(v RatType)`

SetRATType sets RATType field to given value.

### HasRATType

`func (o *QFIContainerInformation) HasRATType() bool`

HasRATType returns a boolean if a field has been set.

### GetServingNetworkFunctionID

`func (o *QFIContainerInformation) GetServingNetworkFunctionID() []ServingNetworkFunctionID`

GetServingNetworkFunctionID returns the ServingNetworkFunctionID field if non-nil, zero value otherwise.

### GetServingNetworkFunctionIDOk

`func (o *QFIContainerInformation) GetServingNetworkFunctionIDOk() (*[]ServingNetworkFunctionID, bool)`

GetServingNetworkFunctionIDOk returns a tuple with the ServingNetworkFunctionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkFunctionID

`func (o *QFIContainerInformation) SetServingNetworkFunctionID(v []ServingNetworkFunctionID)`

SetServingNetworkFunctionID sets ServingNetworkFunctionID field to given value.

### HasServingNetworkFunctionID

`func (o *QFIContainerInformation) HasServingNetworkFunctionID() bool`

HasServingNetworkFunctionID returns a boolean if a field has been set.

### GetVar3gppPSDataOffStatus

`func (o *QFIContainerInformation) GetVar3gppPSDataOffStatus() Model3GPPPSDataOffStatus`

GetVar3gppPSDataOffStatus returns the Var3gppPSDataOffStatus field if non-nil, zero value otherwise.

### GetVar3gppPSDataOffStatusOk

`func (o *QFIContainerInformation) GetVar3gppPSDataOffStatusOk() (*Model3GPPPSDataOffStatus, bool)`

GetVar3gppPSDataOffStatusOk returns a tuple with the Var3gppPSDataOffStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppPSDataOffStatus

`func (o *QFIContainerInformation) SetVar3gppPSDataOffStatus(v Model3GPPPSDataOffStatus)`

SetVar3gppPSDataOffStatus sets Var3gppPSDataOffStatus field to given value.

### HasVar3gppPSDataOffStatus

`func (o *QFIContainerInformation) HasVar3gppPSDataOffStatus() bool`

HasVar3gppPSDataOffStatus returns a boolean if a field has been set.

### GetVar3gppChargingId

`func (o *QFIContainerInformation) GetVar3gppChargingId() int32`

GetVar3gppChargingId returns the Var3gppChargingId field if non-nil, zero value otherwise.

### GetVar3gppChargingIdOk

`func (o *QFIContainerInformation) GetVar3gppChargingIdOk() (*int32, bool)`

GetVar3gppChargingIdOk returns a tuple with the Var3gppChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppChargingId

`func (o *QFIContainerInformation) SetVar3gppChargingId(v int32)`

SetVar3gppChargingId sets Var3gppChargingId field to given value.

### HasVar3gppChargingId

`func (o *QFIContainerInformation) HasVar3gppChargingId() bool`

HasVar3gppChargingId returns a boolean if a field has been set.

### GetDiagnostics

`func (o *QFIContainerInformation) GetDiagnostics() int32`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *QFIContainerInformation) GetDiagnosticsOk() (*int32, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *QFIContainerInformation) SetDiagnostics(v int32)`

SetDiagnostics sets Diagnostics field to given value.

### HasDiagnostics

`func (o *QFIContainerInformation) HasDiagnostics() bool`

HasDiagnostics returns a boolean if a field has been set.

### GetEnhancedDiagnostics

`func (o *QFIContainerInformation) GetEnhancedDiagnostics() []string`

GetEnhancedDiagnostics returns the EnhancedDiagnostics field if non-nil, zero value otherwise.

### GetEnhancedDiagnosticsOk

`func (o *QFIContainerInformation) GetEnhancedDiagnosticsOk() (*[]string, bool)`

GetEnhancedDiagnosticsOk returns a tuple with the EnhancedDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedDiagnostics

`func (o *QFIContainerInformation) SetEnhancedDiagnostics(v []string)`

SetEnhancedDiagnostics sets EnhancedDiagnostics field to given value.

### HasEnhancedDiagnostics

`func (o *QFIContainerInformation) HasEnhancedDiagnostics() bool`

HasEnhancedDiagnostics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


