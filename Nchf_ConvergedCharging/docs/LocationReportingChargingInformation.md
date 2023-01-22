# LocationReportingChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationReportingMessageType** | **int32** |  | 
**UserInformation** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**UserLocationinfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**PSCellInformation** | Pointer to [**PSCellInformation**](PSCellInformation.md) |  | [optional] 
**UetimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**RATType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**PresenceReportingAreaInformation** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 

## Methods

### NewLocationReportingChargingInformation

`func NewLocationReportingChargingInformation(locationReportingMessageType int32, ) *LocationReportingChargingInformation`

NewLocationReportingChargingInformation instantiates a new LocationReportingChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationReportingChargingInformationWithDefaults

`func NewLocationReportingChargingInformationWithDefaults() *LocationReportingChargingInformation`

NewLocationReportingChargingInformationWithDefaults instantiates a new LocationReportingChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationReportingMessageType

`func (o *LocationReportingChargingInformation) GetLocationReportingMessageType() int32`

GetLocationReportingMessageType returns the LocationReportingMessageType field if non-nil, zero value otherwise.

### GetLocationReportingMessageTypeOk

`func (o *LocationReportingChargingInformation) GetLocationReportingMessageTypeOk() (*int32, bool)`

GetLocationReportingMessageTypeOk returns a tuple with the LocationReportingMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReportingMessageType

`func (o *LocationReportingChargingInformation) SetLocationReportingMessageType(v int32)`

SetLocationReportingMessageType sets LocationReportingMessageType field to given value.


### GetUserInformation

`func (o *LocationReportingChargingInformation) GetUserInformation() UserInformation`

GetUserInformation returns the UserInformation field if non-nil, zero value otherwise.

### GetUserInformationOk

`func (o *LocationReportingChargingInformation) GetUserInformationOk() (*UserInformation, bool)`

GetUserInformationOk returns a tuple with the UserInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInformation

`func (o *LocationReportingChargingInformation) SetUserInformation(v UserInformation)`

SetUserInformation sets UserInformation field to given value.

### HasUserInformation

`func (o *LocationReportingChargingInformation) HasUserInformation() bool`

HasUserInformation returns a boolean if a field has been set.

### GetUserLocationinfo

`func (o *LocationReportingChargingInformation) GetUserLocationinfo() UserLocation`

GetUserLocationinfo returns the UserLocationinfo field if non-nil, zero value otherwise.

### GetUserLocationinfoOk

`func (o *LocationReportingChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool)`

GetUserLocationinfoOk returns a tuple with the UserLocationinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationinfo

`func (o *LocationReportingChargingInformation) SetUserLocationinfo(v UserLocation)`

SetUserLocationinfo sets UserLocationinfo field to given value.

### HasUserLocationinfo

`func (o *LocationReportingChargingInformation) HasUserLocationinfo() bool`

HasUserLocationinfo returns a boolean if a field has been set.

### GetPSCellInformation

`func (o *LocationReportingChargingInformation) GetPSCellInformation() PSCellInformation`

GetPSCellInformation returns the PSCellInformation field if non-nil, zero value otherwise.

### GetPSCellInformationOk

`func (o *LocationReportingChargingInformation) GetPSCellInformationOk() (*PSCellInformation, bool)`

GetPSCellInformationOk returns a tuple with the PSCellInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSCellInformation

`func (o *LocationReportingChargingInformation) SetPSCellInformation(v PSCellInformation)`

SetPSCellInformation sets PSCellInformation field to given value.

### HasPSCellInformation

`func (o *LocationReportingChargingInformation) HasPSCellInformation() bool`

HasPSCellInformation returns a boolean if a field has been set.

### GetUetimeZone

`func (o *LocationReportingChargingInformation) GetUetimeZone() string`

GetUetimeZone returns the UetimeZone field if non-nil, zero value otherwise.

### GetUetimeZoneOk

`func (o *LocationReportingChargingInformation) GetUetimeZoneOk() (*string, bool)`

GetUetimeZoneOk returns a tuple with the UetimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetimeZone

`func (o *LocationReportingChargingInformation) SetUetimeZone(v string)`

SetUetimeZone sets UetimeZone field to given value.

### HasUetimeZone

`func (o *LocationReportingChargingInformation) HasUetimeZone() bool`

HasUetimeZone returns a boolean if a field has been set.

### GetRATType

`func (o *LocationReportingChargingInformation) GetRATType() RatType`

GetRATType returns the RATType field if non-nil, zero value otherwise.

### GetRATTypeOk

`func (o *LocationReportingChargingInformation) GetRATTypeOk() (*RatType, bool)`

GetRATTypeOk returns a tuple with the RATType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRATType

`func (o *LocationReportingChargingInformation) SetRATType(v RatType)`

SetRATType sets RATType field to given value.

### HasRATType

`func (o *LocationReportingChargingInformation) HasRATType() bool`

HasRATType returns a boolean if a field has been set.

### GetPresenceReportingAreaInformation

`func (o *LocationReportingChargingInformation) GetPresenceReportingAreaInformation() map[string]PresenceInfo`

GetPresenceReportingAreaInformation returns the PresenceReportingAreaInformation field if non-nil, zero value otherwise.

### GetPresenceReportingAreaInformationOk

`func (o *LocationReportingChargingInformation) GetPresenceReportingAreaInformationOk() (*map[string]PresenceInfo, bool)`

GetPresenceReportingAreaInformationOk returns a tuple with the PresenceReportingAreaInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceReportingAreaInformation

`func (o *LocationReportingChargingInformation) SetPresenceReportingAreaInformation(v map[string]PresenceInfo)`

SetPresenceReportingAreaInformation sets PresenceReportingAreaInformation field to given value.

### HasPresenceReportingAreaInformation

`func (o *LocationReportingChargingInformation) HasPresenceReportingAreaInformation() bool`

HasPresenceReportingAreaInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


