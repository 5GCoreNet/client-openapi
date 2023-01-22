# AreaEventInfoExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaDefinition** | [**[]ReportingArea**](ReportingArea.md) |  | 
**OccurrenceInfo** | Pointer to [**OccurrenceInfo**](OccurrenceInfo.md) |  | [optional] 
**MinimumInterval** | Pointer to **int32** | Minimum interval between event reports. | [optional] 
**MaximumInterval** | Pointer to **int32** | Maximum interval between event reports. | [optional] 
**SamplingInterval** | Pointer to **int32** | Maximum time interval between consecutive evaluations by a UE of a trigger event. | [optional] 
**ReportingDuration** | Pointer to **int32** | Maximum duration of event reporting. | [optional] 
**ReportingLocationReq** | Pointer to **bool** |  | [optional] [default to true]
**GeoAreaList** | Pointer to [**[]GeographicArea**](GeographicArea.md) |  | [optional] 
**IgnoreAreaDefInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAreaEventInfoExt

`func NewAreaEventInfoExt(areaDefinition []ReportingArea, ) *AreaEventInfoExt`

NewAreaEventInfoExt instantiates a new AreaEventInfoExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAreaEventInfoExtWithDefaults

`func NewAreaEventInfoExtWithDefaults() *AreaEventInfoExt`

NewAreaEventInfoExtWithDefaults instantiates a new AreaEventInfoExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaDefinition

`func (o *AreaEventInfoExt) GetAreaDefinition() []ReportingArea`

GetAreaDefinition returns the AreaDefinition field if non-nil, zero value otherwise.

### GetAreaDefinitionOk

`func (o *AreaEventInfoExt) GetAreaDefinitionOk() (*[]ReportingArea, bool)`

GetAreaDefinitionOk returns a tuple with the AreaDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaDefinition

`func (o *AreaEventInfoExt) SetAreaDefinition(v []ReportingArea)`

SetAreaDefinition sets AreaDefinition field to given value.


### GetOccurrenceInfo

`func (o *AreaEventInfoExt) GetOccurrenceInfo() OccurrenceInfo`

GetOccurrenceInfo returns the OccurrenceInfo field if non-nil, zero value otherwise.

### GetOccurrenceInfoOk

`func (o *AreaEventInfoExt) GetOccurrenceInfoOk() (*OccurrenceInfo, bool)`

GetOccurrenceInfoOk returns a tuple with the OccurrenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceInfo

`func (o *AreaEventInfoExt) SetOccurrenceInfo(v OccurrenceInfo)`

SetOccurrenceInfo sets OccurrenceInfo field to given value.

### HasOccurrenceInfo

`func (o *AreaEventInfoExt) HasOccurrenceInfo() bool`

HasOccurrenceInfo returns a boolean if a field has been set.

### GetMinimumInterval

`func (o *AreaEventInfoExt) GetMinimumInterval() int32`

GetMinimumInterval returns the MinimumInterval field if non-nil, zero value otherwise.

### GetMinimumIntervalOk

`func (o *AreaEventInfoExt) GetMinimumIntervalOk() (*int32, bool)`

GetMinimumIntervalOk returns a tuple with the MinimumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumInterval

`func (o *AreaEventInfoExt) SetMinimumInterval(v int32)`

SetMinimumInterval sets MinimumInterval field to given value.

### HasMinimumInterval

`func (o *AreaEventInfoExt) HasMinimumInterval() bool`

HasMinimumInterval returns a boolean if a field has been set.

### GetMaximumInterval

`func (o *AreaEventInfoExt) GetMaximumInterval() int32`

GetMaximumInterval returns the MaximumInterval field if non-nil, zero value otherwise.

### GetMaximumIntervalOk

`func (o *AreaEventInfoExt) GetMaximumIntervalOk() (*int32, bool)`

GetMaximumIntervalOk returns a tuple with the MaximumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInterval

`func (o *AreaEventInfoExt) SetMaximumInterval(v int32)`

SetMaximumInterval sets MaximumInterval field to given value.

### HasMaximumInterval

`func (o *AreaEventInfoExt) HasMaximumInterval() bool`

HasMaximumInterval returns a boolean if a field has been set.

### GetSamplingInterval

`func (o *AreaEventInfoExt) GetSamplingInterval() int32`

GetSamplingInterval returns the SamplingInterval field if non-nil, zero value otherwise.

### GetSamplingIntervalOk

`func (o *AreaEventInfoExt) GetSamplingIntervalOk() (*int32, bool)`

GetSamplingIntervalOk returns a tuple with the SamplingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingInterval

`func (o *AreaEventInfoExt) SetSamplingInterval(v int32)`

SetSamplingInterval sets SamplingInterval field to given value.

### HasSamplingInterval

`func (o *AreaEventInfoExt) HasSamplingInterval() bool`

HasSamplingInterval returns a boolean if a field has been set.

### GetReportingDuration

`func (o *AreaEventInfoExt) GetReportingDuration() int32`

GetReportingDuration returns the ReportingDuration field if non-nil, zero value otherwise.

### GetReportingDurationOk

`func (o *AreaEventInfoExt) GetReportingDurationOk() (*int32, bool)`

GetReportingDurationOk returns a tuple with the ReportingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingDuration

`func (o *AreaEventInfoExt) SetReportingDuration(v int32)`

SetReportingDuration sets ReportingDuration field to given value.

### HasReportingDuration

`func (o *AreaEventInfoExt) HasReportingDuration() bool`

HasReportingDuration returns a boolean if a field has been set.

### GetReportingLocationReq

`func (o *AreaEventInfoExt) GetReportingLocationReq() bool`

GetReportingLocationReq returns the ReportingLocationReq field if non-nil, zero value otherwise.

### GetReportingLocationReqOk

`func (o *AreaEventInfoExt) GetReportingLocationReqOk() (*bool, bool)`

GetReportingLocationReqOk returns a tuple with the ReportingLocationReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingLocationReq

`func (o *AreaEventInfoExt) SetReportingLocationReq(v bool)`

SetReportingLocationReq sets ReportingLocationReq field to given value.

### HasReportingLocationReq

`func (o *AreaEventInfoExt) HasReportingLocationReq() bool`

HasReportingLocationReq returns a boolean if a field has been set.

### GetGeoAreaList

`func (o *AreaEventInfoExt) GetGeoAreaList() []GeographicArea`

GetGeoAreaList returns the GeoAreaList field if non-nil, zero value otherwise.

### GetGeoAreaListOk

`func (o *AreaEventInfoExt) GetGeoAreaListOk() (*[]GeographicArea, bool)`

GetGeoAreaListOk returns a tuple with the GeoAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoAreaList

`func (o *AreaEventInfoExt) SetGeoAreaList(v []GeographicArea)`

SetGeoAreaList sets GeoAreaList field to given value.

### HasGeoAreaList

`func (o *AreaEventInfoExt) HasGeoAreaList() bool`

HasGeoAreaList returns a boolean if a field has been set.

### GetIgnoreAreaDefInd

`func (o *AreaEventInfoExt) GetIgnoreAreaDefInd() bool`

GetIgnoreAreaDefInd returns the IgnoreAreaDefInd field if non-nil, zero value otherwise.

### GetIgnoreAreaDefIndOk

`func (o *AreaEventInfoExt) GetIgnoreAreaDefIndOk() (*bool, bool)`

GetIgnoreAreaDefIndOk returns a tuple with the IgnoreAreaDefInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAreaDefInd

`func (o *AreaEventInfoExt) SetIgnoreAreaDefInd(v bool)`

SetIgnoreAreaDefInd sets IgnoreAreaDefInd field to given value.

### HasIgnoreAreaDefInd

`func (o *AreaEventInfoExt) HasIgnoreAreaDefInd() bool`

HasIgnoreAreaDefInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


