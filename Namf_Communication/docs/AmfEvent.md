# AmfEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AmfEventType**](AmfEventType.md) |  | 
**ImmediateFlag** | Pointer to **bool** |  | [optional] [default to false]
**AreaList** | Pointer to [**[]AmfEventArea**](AmfEventArea.md) |  | [optional] 
**LocationFilterList** | Pointer to [**[]LocationFilter**](LocationFilter.md) |  | [optional] 
**RefId** | Pointer to **int32** |  | [optional] 
**TrafficDescriptorList** | Pointer to [**[]TrafficDescriptor**](TrafficDescriptor.md) |  | [optional] 
**ReportUeReachable** | Pointer to **bool** |  | [optional] [default to false]
**ReachabilityFilter** | Pointer to [**ReachabilityFilter**](ReachabilityFilter.md) |  | [optional] 
**UdmDetectInd** | Pointer to **bool** |  | [optional] [default to false]
**MaxReports** | Pointer to **int32** |  | [optional] 
**PresenceInfoList** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | A map(list of key-value pairs) where praId serves as key. | [optional] 
**MaxResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**TargetArea** | Pointer to [**TargetArea**](TargetArea.md) |  | [optional] 
**SnssaiFilter** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**UeInAreaFilter** | Pointer to [**UeInAreaFilter**](UeInAreaFilter.md) |  | [optional] 
**MinInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NextReport** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**IdleStatusInd** | Pointer to **bool** |  | [optional] [default to false]
**DispersionArea** | Pointer to [**DispersionArea**](DispersionArea.md) |  | [optional] 

## Methods

### NewAmfEvent

`func NewAmfEvent(type_ AmfEventType, ) *AmfEvent`

NewAmfEvent instantiates a new AmfEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventWithDefaults

`func NewAmfEventWithDefaults() *AmfEvent`

NewAmfEventWithDefaults instantiates a new AmfEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AmfEvent) GetType() AmfEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmfEvent) GetTypeOk() (*AmfEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmfEvent) SetType(v AmfEventType)`

SetType sets Type field to given value.


### GetImmediateFlag

`func (o *AmfEvent) GetImmediateFlag() bool`

GetImmediateFlag returns the ImmediateFlag field if non-nil, zero value otherwise.

### GetImmediateFlagOk

`func (o *AmfEvent) GetImmediateFlagOk() (*bool, bool)`

GetImmediateFlagOk returns a tuple with the ImmediateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFlag

`func (o *AmfEvent) SetImmediateFlag(v bool)`

SetImmediateFlag sets ImmediateFlag field to given value.

### HasImmediateFlag

`func (o *AmfEvent) HasImmediateFlag() bool`

HasImmediateFlag returns a boolean if a field has been set.

### GetAreaList

`func (o *AmfEvent) GetAreaList() []AmfEventArea`

GetAreaList returns the AreaList field if non-nil, zero value otherwise.

### GetAreaListOk

`func (o *AmfEvent) GetAreaListOk() (*[]AmfEventArea, bool)`

GetAreaListOk returns a tuple with the AreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaList

`func (o *AmfEvent) SetAreaList(v []AmfEventArea)`

SetAreaList sets AreaList field to given value.

### HasAreaList

`func (o *AmfEvent) HasAreaList() bool`

HasAreaList returns a boolean if a field has been set.

### GetLocationFilterList

`func (o *AmfEvent) GetLocationFilterList() []LocationFilter`

GetLocationFilterList returns the LocationFilterList field if non-nil, zero value otherwise.

### GetLocationFilterListOk

`func (o *AmfEvent) GetLocationFilterListOk() (*[]LocationFilter, bool)`

GetLocationFilterListOk returns a tuple with the LocationFilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationFilterList

`func (o *AmfEvent) SetLocationFilterList(v []LocationFilter)`

SetLocationFilterList sets LocationFilterList field to given value.

### HasLocationFilterList

`func (o *AmfEvent) HasLocationFilterList() bool`

HasLocationFilterList returns a boolean if a field has been set.

### GetRefId

`func (o *AmfEvent) GetRefId() int32`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AmfEvent) GetRefIdOk() (*int32, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AmfEvent) SetRefId(v int32)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AmfEvent) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetTrafficDescriptorList

`func (o *AmfEvent) GetTrafficDescriptorList() []TrafficDescriptor`

GetTrafficDescriptorList returns the TrafficDescriptorList field if non-nil, zero value otherwise.

### GetTrafficDescriptorListOk

`func (o *AmfEvent) GetTrafficDescriptorListOk() (*[]TrafficDescriptor, bool)`

GetTrafficDescriptorListOk returns a tuple with the TrafficDescriptorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDescriptorList

`func (o *AmfEvent) SetTrafficDescriptorList(v []TrafficDescriptor)`

SetTrafficDescriptorList sets TrafficDescriptorList field to given value.

### HasTrafficDescriptorList

`func (o *AmfEvent) HasTrafficDescriptorList() bool`

HasTrafficDescriptorList returns a boolean if a field has been set.

### GetReportUeReachable

`func (o *AmfEvent) GetReportUeReachable() bool`

GetReportUeReachable returns the ReportUeReachable field if non-nil, zero value otherwise.

### GetReportUeReachableOk

`func (o *AmfEvent) GetReportUeReachableOk() (*bool, bool)`

GetReportUeReachableOk returns a tuple with the ReportUeReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUeReachable

`func (o *AmfEvent) SetReportUeReachable(v bool)`

SetReportUeReachable sets ReportUeReachable field to given value.

### HasReportUeReachable

`func (o *AmfEvent) HasReportUeReachable() bool`

HasReportUeReachable returns a boolean if a field has been set.

### GetReachabilityFilter

`func (o *AmfEvent) GetReachabilityFilter() ReachabilityFilter`

GetReachabilityFilter returns the ReachabilityFilter field if non-nil, zero value otherwise.

### GetReachabilityFilterOk

`func (o *AmfEvent) GetReachabilityFilterOk() (*ReachabilityFilter, bool)`

GetReachabilityFilterOk returns a tuple with the ReachabilityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityFilter

`func (o *AmfEvent) SetReachabilityFilter(v ReachabilityFilter)`

SetReachabilityFilter sets ReachabilityFilter field to given value.

### HasReachabilityFilter

`func (o *AmfEvent) HasReachabilityFilter() bool`

HasReachabilityFilter returns a boolean if a field has been set.

### GetUdmDetectInd

`func (o *AmfEvent) GetUdmDetectInd() bool`

GetUdmDetectInd returns the UdmDetectInd field if non-nil, zero value otherwise.

### GetUdmDetectIndOk

`func (o *AmfEvent) GetUdmDetectIndOk() (*bool, bool)`

GetUdmDetectIndOk returns a tuple with the UdmDetectInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmDetectInd

`func (o *AmfEvent) SetUdmDetectInd(v bool)`

SetUdmDetectInd sets UdmDetectInd field to given value.

### HasUdmDetectInd

`func (o *AmfEvent) HasUdmDetectInd() bool`

HasUdmDetectInd returns a boolean if a field has been set.

### GetMaxReports

`func (o *AmfEvent) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *AmfEvent) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *AmfEvent) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *AmfEvent) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetPresenceInfoList

`func (o *AmfEvent) GetPresenceInfoList() map[string]PresenceInfo`

GetPresenceInfoList returns the PresenceInfoList field if non-nil, zero value otherwise.

### GetPresenceInfoListOk

`func (o *AmfEvent) GetPresenceInfoListOk() (*map[string]PresenceInfo, bool)`

GetPresenceInfoListOk returns a tuple with the PresenceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInfoList

`func (o *AmfEvent) SetPresenceInfoList(v map[string]PresenceInfo)`

SetPresenceInfoList sets PresenceInfoList field to given value.

### HasPresenceInfoList

`func (o *AmfEvent) HasPresenceInfoList() bool`

HasPresenceInfoList returns a boolean if a field has been set.

### GetMaxResponseTime

`func (o *AmfEvent) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *AmfEvent) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *AmfEvent) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *AmfEvent) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetTargetArea

`func (o *AmfEvent) GetTargetArea() TargetArea`

GetTargetArea returns the TargetArea field if non-nil, zero value otherwise.

### GetTargetAreaOk

`func (o *AmfEvent) GetTargetAreaOk() (*TargetArea, bool)`

GetTargetAreaOk returns a tuple with the TargetArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetArea

`func (o *AmfEvent) SetTargetArea(v TargetArea)`

SetTargetArea sets TargetArea field to given value.

### HasTargetArea

`func (o *AmfEvent) HasTargetArea() bool`

HasTargetArea returns a boolean if a field has been set.

### GetSnssaiFilter

`func (o *AmfEvent) GetSnssaiFilter() []ExtSnssai`

GetSnssaiFilter returns the SnssaiFilter field if non-nil, zero value otherwise.

### GetSnssaiFilterOk

`func (o *AmfEvent) GetSnssaiFilterOk() (*[]ExtSnssai, bool)`

GetSnssaiFilterOk returns a tuple with the SnssaiFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiFilter

`func (o *AmfEvent) SetSnssaiFilter(v []ExtSnssai)`

SetSnssaiFilter sets SnssaiFilter field to given value.

### HasSnssaiFilter

`func (o *AmfEvent) HasSnssaiFilter() bool`

HasSnssaiFilter returns a boolean if a field has been set.

### GetUeInAreaFilter

`func (o *AmfEvent) GetUeInAreaFilter() UeInAreaFilter`

GetUeInAreaFilter returns the UeInAreaFilter field if non-nil, zero value otherwise.

### GetUeInAreaFilterOk

`func (o *AmfEvent) GetUeInAreaFilterOk() (*UeInAreaFilter, bool)`

GetUeInAreaFilterOk returns a tuple with the UeInAreaFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeInAreaFilter

`func (o *AmfEvent) SetUeInAreaFilter(v UeInAreaFilter)`

SetUeInAreaFilter sets UeInAreaFilter field to given value.

### HasUeInAreaFilter

`func (o *AmfEvent) HasUeInAreaFilter() bool`

HasUeInAreaFilter returns a boolean if a field has been set.

### GetMinInterval

`func (o *AmfEvent) GetMinInterval() int32`

GetMinInterval returns the MinInterval field if non-nil, zero value otherwise.

### GetMinIntervalOk

`func (o *AmfEvent) GetMinIntervalOk() (*int32, bool)`

GetMinIntervalOk returns a tuple with the MinInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInterval

`func (o *AmfEvent) SetMinInterval(v int32)`

SetMinInterval sets MinInterval field to given value.

### HasMinInterval

`func (o *AmfEvent) HasMinInterval() bool`

HasMinInterval returns a boolean if a field has been set.

### GetNextReport

`func (o *AmfEvent) GetNextReport() time.Time`

GetNextReport returns the NextReport field if non-nil, zero value otherwise.

### GetNextReportOk

`func (o *AmfEvent) GetNextReportOk() (*time.Time, bool)`

GetNextReportOk returns a tuple with the NextReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReport

`func (o *AmfEvent) SetNextReport(v time.Time)`

SetNextReport sets NextReport field to given value.

### HasNextReport

`func (o *AmfEvent) HasNextReport() bool`

HasNextReport returns a boolean if a field has been set.

### GetIdleStatusInd

`func (o *AmfEvent) GetIdleStatusInd() bool`

GetIdleStatusInd returns the IdleStatusInd field if non-nil, zero value otherwise.

### GetIdleStatusIndOk

`func (o *AmfEvent) GetIdleStatusIndOk() (*bool, bool)`

GetIdleStatusIndOk returns a tuple with the IdleStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusInd

`func (o *AmfEvent) SetIdleStatusInd(v bool)`

SetIdleStatusInd sets IdleStatusInd field to given value.

### HasIdleStatusInd

`func (o *AmfEvent) HasIdleStatusInd() bool`

HasIdleStatusInd returns a boolean if a field has been set.

### GetDispersionArea

`func (o *AmfEvent) GetDispersionArea() DispersionArea`

GetDispersionArea returns the DispersionArea field if non-nil, zero value otherwise.

### GetDispersionAreaOk

`func (o *AmfEvent) GetDispersionAreaOk() (*DispersionArea, bool)`

GetDispersionAreaOk returns a tuple with the DispersionArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispersionArea

`func (o *AmfEvent) SetDispersionArea(v DispersionArea)`

SetDispersionArea sets DispersionArea field to given value.

### HasDispersionArea

`func (o *AmfEvent) HasDispersionArea() bool`

HasDispersionArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


