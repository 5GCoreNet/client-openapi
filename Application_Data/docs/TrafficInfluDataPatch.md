# TrafficInfluDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpPathChgNotifCorreId** | Pointer to **string** | Contains the Notification Correlation Id allocated by the NEF for the UP path change notification.  | [optional] 
**AppReloInd** | Pointer to **bool** | Identifies whether an application can be relocated once a location of the application has been selected.  | [optional] 
**EthTrafficFilters** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet filters. Either \&quot;trafficFilters\&quot; or \&quot;ethTrafficFilters\&quot; shall be included if applicable.  | [optional] 
**TrafficFilters** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Identifies IP packet filters. Either \&quot;trafficFilters\&quot; or \&quot;ethTrafficFilters\&quot; shall be included if applicable.  | [optional] 
**TrafficRoutes** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) | Identifies the N6 traffic routing requirement. | [optional] 
**TraffCorreInd** | Pointer to **bool** |  | [optional] 
**ValidStartTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ValidEndTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TempValidities** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) | Identifies the temporal validities for the N6 traffic routing requirement. | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**UpPathChgNotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Headers** | Pointer to **[]string** |  | [optional] 
**AfAckInd** | Pointer to **bool** |  | [optional] 
**AddrPreserInd** | Pointer to **bool** |  | [optional] 
**MaxAllowedUpLat** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**SimConnInd** | Pointer to **bool** | Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.  | [optional] 
**SimConnTerm** | Pointer to **NullableInt32** | indicating a time in seconds with OpenAPI defined &#39;nullable: true&#39; property. | [optional] 

## Methods

### NewTrafficInfluDataPatch

`func NewTrafficInfluDataPatch() *TrafficInfluDataPatch`

NewTrafficInfluDataPatch instantiates a new TrafficInfluDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficInfluDataPatchWithDefaults

`func NewTrafficInfluDataPatchWithDefaults() *TrafficInfluDataPatch`

NewTrafficInfluDataPatchWithDefaults instantiates a new TrafficInfluDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpPathChgNotifCorreId

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifCorreId() string`

GetUpPathChgNotifCorreId returns the UpPathChgNotifCorreId field if non-nil, zero value otherwise.

### GetUpPathChgNotifCorreIdOk

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifCorreIdOk() (*string, bool)`

GetUpPathChgNotifCorreIdOk returns a tuple with the UpPathChgNotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgNotifCorreId

`func (o *TrafficInfluDataPatch) SetUpPathChgNotifCorreId(v string)`

SetUpPathChgNotifCorreId sets UpPathChgNotifCorreId field to given value.

### HasUpPathChgNotifCorreId

`func (o *TrafficInfluDataPatch) HasUpPathChgNotifCorreId() bool`

HasUpPathChgNotifCorreId returns a boolean if a field has been set.

### GetAppReloInd

`func (o *TrafficInfluDataPatch) GetAppReloInd() bool`

GetAppReloInd returns the AppReloInd field if non-nil, zero value otherwise.

### GetAppReloIndOk

`func (o *TrafficInfluDataPatch) GetAppReloIndOk() (*bool, bool)`

GetAppReloIndOk returns a tuple with the AppReloInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloInd

`func (o *TrafficInfluDataPatch) SetAppReloInd(v bool)`

SetAppReloInd sets AppReloInd field to given value.

### HasAppReloInd

`func (o *TrafficInfluDataPatch) HasAppReloInd() bool`

HasAppReloInd returns a boolean if a field has been set.

### GetEthTrafficFilters

`func (o *TrafficInfluDataPatch) GetEthTrafficFilters() []EthFlowDescription`

GetEthTrafficFilters returns the EthTrafficFilters field if non-nil, zero value otherwise.

### GetEthTrafficFiltersOk

`func (o *TrafficInfluDataPatch) GetEthTrafficFiltersOk() (*[]EthFlowDescription, bool)`

GetEthTrafficFiltersOk returns a tuple with the EthTrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthTrafficFilters

`func (o *TrafficInfluDataPatch) SetEthTrafficFilters(v []EthFlowDescription)`

SetEthTrafficFilters sets EthTrafficFilters field to given value.

### HasEthTrafficFilters

`func (o *TrafficInfluDataPatch) HasEthTrafficFilters() bool`

HasEthTrafficFilters returns a boolean if a field has been set.

### GetTrafficFilters

`func (o *TrafficInfluDataPatch) GetTrafficFilters() []FlowInfo`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *TrafficInfluDataPatch) GetTrafficFiltersOk() (*[]FlowInfo, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *TrafficInfluDataPatch) SetTrafficFilters(v []FlowInfo)`

SetTrafficFilters sets TrafficFilters field to given value.

### HasTrafficFilters

`func (o *TrafficInfluDataPatch) HasTrafficFilters() bool`

HasTrafficFilters returns a boolean if a field has been set.

### GetTrafficRoutes

`func (o *TrafficInfluDataPatch) GetTrafficRoutes() []RouteToLocation`

GetTrafficRoutes returns the TrafficRoutes field if non-nil, zero value otherwise.

### GetTrafficRoutesOk

`func (o *TrafficInfluDataPatch) GetTrafficRoutesOk() (*[]RouteToLocation, bool)`

GetTrafficRoutesOk returns a tuple with the TrafficRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRoutes

`func (o *TrafficInfluDataPatch) SetTrafficRoutes(v []RouteToLocation)`

SetTrafficRoutes sets TrafficRoutes field to given value.

### HasTrafficRoutes

`func (o *TrafficInfluDataPatch) HasTrafficRoutes() bool`

HasTrafficRoutes returns a boolean if a field has been set.

### GetTraffCorreInd

`func (o *TrafficInfluDataPatch) GetTraffCorreInd() bool`

GetTraffCorreInd returns the TraffCorreInd field if non-nil, zero value otherwise.

### GetTraffCorreIndOk

`func (o *TrafficInfluDataPatch) GetTraffCorreIndOk() (*bool, bool)`

GetTraffCorreIndOk returns a tuple with the TraffCorreInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffCorreInd

`func (o *TrafficInfluDataPatch) SetTraffCorreInd(v bool)`

SetTraffCorreInd sets TraffCorreInd field to given value.

### HasTraffCorreInd

`func (o *TrafficInfluDataPatch) HasTraffCorreInd() bool`

HasTraffCorreInd returns a boolean if a field has been set.

### GetValidStartTime

`func (o *TrafficInfluDataPatch) GetValidStartTime() time.Time`

GetValidStartTime returns the ValidStartTime field if non-nil, zero value otherwise.

### GetValidStartTimeOk

`func (o *TrafficInfluDataPatch) GetValidStartTimeOk() (*time.Time, bool)`

GetValidStartTimeOk returns a tuple with the ValidStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidStartTime

`func (o *TrafficInfluDataPatch) SetValidStartTime(v time.Time)`

SetValidStartTime sets ValidStartTime field to given value.

### HasValidStartTime

`func (o *TrafficInfluDataPatch) HasValidStartTime() bool`

HasValidStartTime returns a boolean if a field has been set.

### GetValidEndTime

`func (o *TrafficInfluDataPatch) GetValidEndTime() time.Time`

GetValidEndTime returns the ValidEndTime field if non-nil, zero value otherwise.

### GetValidEndTimeOk

`func (o *TrafficInfluDataPatch) GetValidEndTimeOk() (*time.Time, bool)`

GetValidEndTimeOk returns a tuple with the ValidEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidEndTime

`func (o *TrafficInfluDataPatch) SetValidEndTime(v time.Time)`

SetValidEndTime sets ValidEndTime field to given value.

### HasValidEndTime

`func (o *TrafficInfluDataPatch) HasValidEndTime() bool`

HasValidEndTime returns a boolean if a field has been set.

### GetTempValidities

`func (o *TrafficInfluDataPatch) GetTempValidities() []TemporalValidity`

GetTempValidities returns the TempValidities field if non-nil, zero value otherwise.

### GetTempValiditiesOk

`func (o *TrafficInfluDataPatch) GetTempValiditiesOk() (*[]TemporalValidity, bool)`

GetTempValiditiesOk returns a tuple with the TempValidities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempValidities

`func (o *TrafficInfluDataPatch) SetTempValidities(v []TemporalValidity)`

SetTempValidities sets TempValidities field to given value.

### HasTempValidities

`func (o *TrafficInfluDataPatch) HasTempValidities() bool`

HasTempValidities returns a boolean if a field has been set.

### SetTempValiditiesNil

`func (o *TrafficInfluDataPatch) SetTempValiditiesNil(b bool)`

 SetTempValiditiesNil sets the value for TempValidities to be an explicit nil

### UnsetTempValidities
`func (o *TrafficInfluDataPatch) UnsetTempValidities()`

UnsetTempValidities ensures that no value is present for TempValidities, not even an explicit nil
### GetNwAreaInfo

`func (o *TrafficInfluDataPatch) GetNwAreaInfo() NetworkAreaInfo`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *TrafficInfluDataPatch) GetNwAreaInfoOk() (*NetworkAreaInfo, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *TrafficInfluDataPatch) SetNwAreaInfo(v NetworkAreaInfo)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *TrafficInfluDataPatch) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetUpPathChgNotifUri

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifUri() string`

GetUpPathChgNotifUri returns the UpPathChgNotifUri field if non-nil, zero value otherwise.

### GetUpPathChgNotifUriOk

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifUriOk() (*string, bool)`

GetUpPathChgNotifUriOk returns a tuple with the UpPathChgNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgNotifUri

`func (o *TrafficInfluDataPatch) SetUpPathChgNotifUri(v string)`

SetUpPathChgNotifUri sets UpPathChgNotifUri field to given value.

### HasUpPathChgNotifUri

`func (o *TrafficInfluDataPatch) HasUpPathChgNotifUri() bool`

HasUpPathChgNotifUri returns a boolean if a field has been set.

### GetHeaders

`func (o *TrafficInfluDataPatch) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TrafficInfluDataPatch) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TrafficInfluDataPatch) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TrafficInfluDataPatch) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetAfAckInd

`func (o *TrafficInfluDataPatch) GetAfAckInd() bool`

GetAfAckInd returns the AfAckInd field if non-nil, zero value otherwise.

### GetAfAckIndOk

`func (o *TrafficInfluDataPatch) GetAfAckIndOk() (*bool, bool)`

GetAfAckIndOk returns a tuple with the AfAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAckInd

`func (o *TrafficInfluDataPatch) SetAfAckInd(v bool)`

SetAfAckInd sets AfAckInd field to given value.

### HasAfAckInd

`func (o *TrafficInfluDataPatch) HasAfAckInd() bool`

HasAfAckInd returns a boolean if a field has been set.

### GetAddrPreserInd

`func (o *TrafficInfluDataPatch) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *TrafficInfluDataPatch) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *TrafficInfluDataPatch) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *TrafficInfluDataPatch) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.

### GetMaxAllowedUpLat

`func (o *TrafficInfluDataPatch) GetMaxAllowedUpLat() int32`

GetMaxAllowedUpLat returns the MaxAllowedUpLat field if non-nil, zero value otherwise.

### GetMaxAllowedUpLatOk

`func (o *TrafficInfluDataPatch) GetMaxAllowedUpLatOk() (*int32, bool)`

GetMaxAllowedUpLatOk returns a tuple with the MaxAllowedUpLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedUpLat

`func (o *TrafficInfluDataPatch) SetMaxAllowedUpLat(v int32)`

SetMaxAllowedUpLat sets MaxAllowedUpLat field to given value.

### HasMaxAllowedUpLat

`func (o *TrafficInfluDataPatch) HasMaxAllowedUpLat() bool`

HasMaxAllowedUpLat returns a boolean if a field has been set.

### SetMaxAllowedUpLatNil

`func (o *TrafficInfluDataPatch) SetMaxAllowedUpLatNil(b bool)`

 SetMaxAllowedUpLatNil sets the value for MaxAllowedUpLat to be an explicit nil

### UnsetMaxAllowedUpLat
`func (o *TrafficInfluDataPatch) UnsetMaxAllowedUpLat()`

UnsetMaxAllowedUpLat ensures that no value is present for MaxAllowedUpLat, not even an explicit nil
### GetSimConnInd

`func (o *TrafficInfluDataPatch) GetSimConnInd() bool`

GetSimConnInd returns the SimConnInd field if non-nil, zero value otherwise.

### GetSimConnIndOk

`func (o *TrafficInfluDataPatch) GetSimConnIndOk() (*bool, bool)`

GetSimConnIndOk returns a tuple with the SimConnInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnInd

`func (o *TrafficInfluDataPatch) SetSimConnInd(v bool)`

SetSimConnInd sets SimConnInd field to given value.

### HasSimConnInd

`func (o *TrafficInfluDataPatch) HasSimConnInd() bool`

HasSimConnInd returns a boolean if a field has been set.

### GetSimConnTerm

`func (o *TrafficInfluDataPatch) GetSimConnTerm() int32`

GetSimConnTerm returns the SimConnTerm field if non-nil, zero value otherwise.

### GetSimConnTermOk

`func (o *TrafficInfluDataPatch) GetSimConnTermOk() (*int32, bool)`

GetSimConnTermOk returns a tuple with the SimConnTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnTerm

`func (o *TrafficInfluDataPatch) SetSimConnTerm(v int32)`

SetSimConnTerm sets SimConnTerm field to given value.

### HasSimConnTerm

`func (o *TrafficInfluDataPatch) HasSimConnTerm() bool`

HasSimConnTerm returns a boolean if a field has been set.

### SetSimConnTermNil

`func (o *TrafficInfluDataPatch) SetSimConnTermNil(b bool)`

 SetSimConnTermNil sets the value for SimConnTerm to be an explicit nil

### UnsetSimConnTerm
`func (o *TrafficInfluDataPatch) UnsetSimConnTerm()`

UnsetSimConnTerm ensures that no value is present for SimConnTerm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


