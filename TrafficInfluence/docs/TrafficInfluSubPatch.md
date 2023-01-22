# TrafficInfluSubPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppReloInd** | Pointer to **NullableBool** | Identifies whether an application can be relocated once a location of the application has been selected.  | [optional] 
**TrafficFilters** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Identifies IP packet filters. | [optional] 
**EthTrafficFilters** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet filters. | [optional] 
**TrafficRoutes** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) | Identifies the N6 traffic routing requirement. | [optional] 
**TfcCorrInd** | Pointer to **NullableBool** |  | [optional] 
**TempValidities** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 
**ValidGeoZoneIds** | Pointer to **[]string** | Identifies a geographic zone that the AF request applies only to the traffic of UE(s) located in this specific zone.  | [optional] 
**GeoAreas** | Pointer to [**[]GeographicalArea**](GeographicalArea.md) | Identifies geographical areas within which the AF request applies. | [optional] 
**AfAckInd** | Pointer to **NullableBool** |  | [optional] 
**AddrPreserInd** | Pointer to **NullableBool** |  | [optional] 
**SimConnInd** | Pointer to **bool** | Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.  | [optional] 
**SimConnTerm** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**MaxAllowedUpLat** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**EasIpReplaceInfos** | Pointer to [**[]EasIpReplacementInfo**](EasIpReplacementInfo.md) | Contains EAS IP replacement information. | [optional] 
**EasRedisInd** | Pointer to **bool** | Indicates the EAS rediscovery is required for the application if it is included and set to \&quot;true\&quot;. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 

## Methods

### NewTrafficInfluSubPatch

`func NewTrafficInfluSubPatch() *TrafficInfluSubPatch`

NewTrafficInfluSubPatch instantiates a new TrafficInfluSubPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficInfluSubPatchWithDefaults

`func NewTrafficInfluSubPatchWithDefaults() *TrafficInfluSubPatch`

NewTrafficInfluSubPatchWithDefaults instantiates a new TrafficInfluSubPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppReloInd

`func (o *TrafficInfluSubPatch) GetAppReloInd() bool`

GetAppReloInd returns the AppReloInd field if non-nil, zero value otherwise.

### GetAppReloIndOk

`func (o *TrafficInfluSubPatch) GetAppReloIndOk() (*bool, bool)`

GetAppReloIndOk returns a tuple with the AppReloInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloInd

`func (o *TrafficInfluSubPatch) SetAppReloInd(v bool)`

SetAppReloInd sets AppReloInd field to given value.

### HasAppReloInd

`func (o *TrafficInfluSubPatch) HasAppReloInd() bool`

HasAppReloInd returns a boolean if a field has been set.

### SetAppReloIndNil

`func (o *TrafficInfluSubPatch) SetAppReloIndNil(b bool)`

 SetAppReloIndNil sets the value for AppReloInd to be an explicit nil

### UnsetAppReloInd
`func (o *TrafficInfluSubPatch) UnsetAppReloInd()`

UnsetAppReloInd ensures that no value is present for AppReloInd, not even an explicit nil
### GetTrafficFilters

`func (o *TrafficInfluSubPatch) GetTrafficFilters() []FlowInfo`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *TrafficInfluSubPatch) GetTrafficFiltersOk() (*[]FlowInfo, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *TrafficInfluSubPatch) SetTrafficFilters(v []FlowInfo)`

SetTrafficFilters sets TrafficFilters field to given value.

### HasTrafficFilters

`func (o *TrafficInfluSubPatch) HasTrafficFilters() bool`

HasTrafficFilters returns a boolean if a field has been set.

### GetEthTrafficFilters

`func (o *TrafficInfluSubPatch) GetEthTrafficFilters() []EthFlowDescription`

GetEthTrafficFilters returns the EthTrafficFilters field if non-nil, zero value otherwise.

### GetEthTrafficFiltersOk

`func (o *TrafficInfluSubPatch) GetEthTrafficFiltersOk() (*[]EthFlowDescription, bool)`

GetEthTrafficFiltersOk returns a tuple with the EthTrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthTrafficFilters

`func (o *TrafficInfluSubPatch) SetEthTrafficFilters(v []EthFlowDescription)`

SetEthTrafficFilters sets EthTrafficFilters field to given value.

### HasEthTrafficFilters

`func (o *TrafficInfluSubPatch) HasEthTrafficFilters() bool`

HasEthTrafficFilters returns a boolean if a field has been set.

### GetTrafficRoutes

`func (o *TrafficInfluSubPatch) GetTrafficRoutes() []RouteToLocation`

GetTrafficRoutes returns the TrafficRoutes field if non-nil, zero value otherwise.

### GetTrafficRoutesOk

`func (o *TrafficInfluSubPatch) GetTrafficRoutesOk() (*[]RouteToLocation, bool)`

GetTrafficRoutesOk returns a tuple with the TrafficRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRoutes

`func (o *TrafficInfluSubPatch) SetTrafficRoutes(v []RouteToLocation)`

SetTrafficRoutes sets TrafficRoutes field to given value.

### HasTrafficRoutes

`func (o *TrafficInfluSubPatch) HasTrafficRoutes() bool`

HasTrafficRoutes returns a boolean if a field has been set.

### GetTfcCorrInd

`func (o *TrafficInfluSubPatch) GetTfcCorrInd() bool`

GetTfcCorrInd returns the TfcCorrInd field if non-nil, zero value otherwise.

### GetTfcCorrIndOk

`func (o *TrafficInfluSubPatch) GetTfcCorrIndOk() (*bool, bool)`

GetTfcCorrIndOk returns a tuple with the TfcCorrInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfcCorrInd

`func (o *TrafficInfluSubPatch) SetTfcCorrInd(v bool)`

SetTfcCorrInd sets TfcCorrInd field to given value.

### HasTfcCorrInd

`func (o *TrafficInfluSubPatch) HasTfcCorrInd() bool`

HasTfcCorrInd returns a boolean if a field has been set.

### SetTfcCorrIndNil

`func (o *TrafficInfluSubPatch) SetTfcCorrIndNil(b bool)`

 SetTfcCorrIndNil sets the value for TfcCorrInd to be an explicit nil

### UnsetTfcCorrInd
`func (o *TrafficInfluSubPatch) UnsetTfcCorrInd()`

UnsetTfcCorrInd ensures that no value is present for TfcCorrInd, not even an explicit nil
### GetTempValidities

`func (o *TrafficInfluSubPatch) GetTempValidities() []TemporalValidity`

GetTempValidities returns the TempValidities field if non-nil, zero value otherwise.

### GetTempValiditiesOk

`func (o *TrafficInfluSubPatch) GetTempValiditiesOk() (*[]TemporalValidity, bool)`

GetTempValiditiesOk returns a tuple with the TempValidities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempValidities

`func (o *TrafficInfluSubPatch) SetTempValidities(v []TemporalValidity)`

SetTempValidities sets TempValidities field to given value.

### HasTempValidities

`func (o *TrafficInfluSubPatch) HasTempValidities() bool`

HasTempValidities returns a boolean if a field has been set.

### SetTempValiditiesNil

`func (o *TrafficInfluSubPatch) SetTempValiditiesNil(b bool)`

 SetTempValiditiesNil sets the value for TempValidities to be an explicit nil

### UnsetTempValidities
`func (o *TrafficInfluSubPatch) UnsetTempValidities()`

UnsetTempValidities ensures that no value is present for TempValidities, not even an explicit nil
### GetValidGeoZoneIds

`func (o *TrafficInfluSubPatch) GetValidGeoZoneIds() []string`

GetValidGeoZoneIds returns the ValidGeoZoneIds field if non-nil, zero value otherwise.

### GetValidGeoZoneIdsOk

`func (o *TrafficInfluSubPatch) GetValidGeoZoneIdsOk() (*[]string, bool)`

GetValidGeoZoneIdsOk returns a tuple with the ValidGeoZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidGeoZoneIds

`func (o *TrafficInfluSubPatch) SetValidGeoZoneIds(v []string)`

SetValidGeoZoneIds sets ValidGeoZoneIds field to given value.

### HasValidGeoZoneIds

`func (o *TrafficInfluSubPatch) HasValidGeoZoneIds() bool`

HasValidGeoZoneIds returns a boolean if a field has been set.

### SetValidGeoZoneIdsNil

`func (o *TrafficInfluSubPatch) SetValidGeoZoneIdsNil(b bool)`

 SetValidGeoZoneIdsNil sets the value for ValidGeoZoneIds to be an explicit nil

### UnsetValidGeoZoneIds
`func (o *TrafficInfluSubPatch) UnsetValidGeoZoneIds()`

UnsetValidGeoZoneIds ensures that no value is present for ValidGeoZoneIds, not even an explicit nil
### GetGeoAreas

`func (o *TrafficInfluSubPatch) GetGeoAreas() []GeographicalArea`

GetGeoAreas returns the GeoAreas field if non-nil, zero value otherwise.

### GetGeoAreasOk

`func (o *TrafficInfluSubPatch) GetGeoAreasOk() (*[]GeographicalArea, bool)`

GetGeoAreasOk returns a tuple with the GeoAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoAreas

`func (o *TrafficInfluSubPatch) SetGeoAreas(v []GeographicalArea)`

SetGeoAreas sets GeoAreas field to given value.

### HasGeoAreas

`func (o *TrafficInfluSubPatch) HasGeoAreas() bool`

HasGeoAreas returns a boolean if a field has been set.

### SetGeoAreasNil

`func (o *TrafficInfluSubPatch) SetGeoAreasNil(b bool)`

 SetGeoAreasNil sets the value for GeoAreas to be an explicit nil

### UnsetGeoAreas
`func (o *TrafficInfluSubPatch) UnsetGeoAreas()`

UnsetGeoAreas ensures that no value is present for GeoAreas, not even an explicit nil
### GetAfAckInd

`func (o *TrafficInfluSubPatch) GetAfAckInd() bool`

GetAfAckInd returns the AfAckInd field if non-nil, zero value otherwise.

### GetAfAckIndOk

`func (o *TrafficInfluSubPatch) GetAfAckIndOk() (*bool, bool)`

GetAfAckIndOk returns a tuple with the AfAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAckInd

`func (o *TrafficInfluSubPatch) SetAfAckInd(v bool)`

SetAfAckInd sets AfAckInd field to given value.

### HasAfAckInd

`func (o *TrafficInfluSubPatch) HasAfAckInd() bool`

HasAfAckInd returns a boolean if a field has been set.

### SetAfAckIndNil

`func (o *TrafficInfluSubPatch) SetAfAckIndNil(b bool)`

 SetAfAckIndNil sets the value for AfAckInd to be an explicit nil

### UnsetAfAckInd
`func (o *TrafficInfluSubPatch) UnsetAfAckInd()`

UnsetAfAckInd ensures that no value is present for AfAckInd, not even an explicit nil
### GetAddrPreserInd

`func (o *TrafficInfluSubPatch) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *TrafficInfluSubPatch) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *TrafficInfluSubPatch) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *TrafficInfluSubPatch) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.

### SetAddrPreserIndNil

`func (o *TrafficInfluSubPatch) SetAddrPreserIndNil(b bool)`

 SetAddrPreserIndNil sets the value for AddrPreserInd to be an explicit nil

### UnsetAddrPreserInd
`func (o *TrafficInfluSubPatch) UnsetAddrPreserInd()`

UnsetAddrPreserInd ensures that no value is present for AddrPreserInd, not even an explicit nil
### GetSimConnInd

`func (o *TrafficInfluSubPatch) GetSimConnInd() bool`

GetSimConnInd returns the SimConnInd field if non-nil, zero value otherwise.

### GetSimConnIndOk

`func (o *TrafficInfluSubPatch) GetSimConnIndOk() (*bool, bool)`

GetSimConnIndOk returns a tuple with the SimConnInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnInd

`func (o *TrafficInfluSubPatch) SetSimConnInd(v bool)`

SetSimConnInd sets SimConnInd field to given value.

### HasSimConnInd

`func (o *TrafficInfluSubPatch) HasSimConnInd() bool`

HasSimConnInd returns a boolean if a field has been set.

### GetSimConnTerm

`func (o *TrafficInfluSubPatch) GetSimConnTerm() int32`

GetSimConnTerm returns the SimConnTerm field if non-nil, zero value otherwise.

### GetSimConnTermOk

`func (o *TrafficInfluSubPatch) GetSimConnTermOk() (*int32, bool)`

GetSimConnTermOk returns a tuple with the SimConnTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnTerm

`func (o *TrafficInfluSubPatch) SetSimConnTerm(v int32)`

SetSimConnTerm sets SimConnTerm field to given value.

### HasSimConnTerm

`func (o *TrafficInfluSubPatch) HasSimConnTerm() bool`

HasSimConnTerm returns a boolean if a field has been set.

### GetMaxAllowedUpLat

`func (o *TrafficInfluSubPatch) GetMaxAllowedUpLat() int32`

GetMaxAllowedUpLat returns the MaxAllowedUpLat field if non-nil, zero value otherwise.

### GetMaxAllowedUpLatOk

`func (o *TrafficInfluSubPatch) GetMaxAllowedUpLatOk() (*int32, bool)`

GetMaxAllowedUpLatOk returns a tuple with the MaxAllowedUpLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedUpLat

`func (o *TrafficInfluSubPatch) SetMaxAllowedUpLat(v int32)`

SetMaxAllowedUpLat sets MaxAllowedUpLat field to given value.

### HasMaxAllowedUpLat

`func (o *TrafficInfluSubPatch) HasMaxAllowedUpLat() bool`

HasMaxAllowedUpLat returns a boolean if a field has been set.

### SetMaxAllowedUpLatNil

`func (o *TrafficInfluSubPatch) SetMaxAllowedUpLatNil(b bool)`

 SetMaxAllowedUpLatNil sets the value for MaxAllowedUpLat to be an explicit nil

### UnsetMaxAllowedUpLat
`func (o *TrafficInfluSubPatch) UnsetMaxAllowedUpLat()`

UnsetMaxAllowedUpLat ensures that no value is present for MaxAllowedUpLat, not even an explicit nil
### GetEasIpReplaceInfos

`func (o *TrafficInfluSubPatch) GetEasIpReplaceInfos() []EasIpReplacementInfo`

GetEasIpReplaceInfos returns the EasIpReplaceInfos field if non-nil, zero value otherwise.

### GetEasIpReplaceInfosOk

`func (o *TrafficInfluSubPatch) GetEasIpReplaceInfosOk() (*[]EasIpReplacementInfo, bool)`

GetEasIpReplaceInfosOk returns a tuple with the EasIpReplaceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpReplaceInfos

`func (o *TrafficInfluSubPatch) SetEasIpReplaceInfos(v []EasIpReplacementInfo)`

SetEasIpReplaceInfos sets EasIpReplaceInfos field to given value.

### HasEasIpReplaceInfos

`func (o *TrafficInfluSubPatch) HasEasIpReplaceInfos() bool`

HasEasIpReplaceInfos returns a boolean if a field has been set.

### SetEasIpReplaceInfosNil

`func (o *TrafficInfluSubPatch) SetEasIpReplaceInfosNil(b bool)`

 SetEasIpReplaceInfosNil sets the value for EasIpReplaceInfos to be an explicit nil

### UnsetEasIpReplaceInfos
`func (o *TrafficInfluSubPatch) UnsetEasIpReplaceInfos()`

UnsetEasIpReplaceInfos ensures that no value is present for EasIpReplaceInfos, not even an explicit nil
### GetEasRedisInd

`func (o *TrafficInfluSubPatch) GetEasRedisInd() bool`

GetEasRedisInd returns the EasRedisInd field if non-nil, zero value otherwise.

### GetEasRedisIndOk

`func (o *TrafficInfluSubPatch) GetEasRedisIndOk() (*bool, bool)`

GetEasRedisIndOk returns a tuple with the EasRedisInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasRedisInd

`func (o *TrafficInfluSubPatch) SetEasRedisInd(v bool)`

SetEasRedisInd sets EasRedisInd field to given value.

### HasEasRedisInd

`func (o *TrafficInfluSubPatch) HasEasRedisInd() bool`

HasEasRedisInd returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *TrafficInfluSubPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *TrafficInfluSubPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *TrafficInfluSubPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *TrafficInfluSubPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetEventReq

`func (o *TrafficInfluSubPatch) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *TrafficInfluSubPatch) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *TrafficInfluSubPatch) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *TrafficInfluSubPatch) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


