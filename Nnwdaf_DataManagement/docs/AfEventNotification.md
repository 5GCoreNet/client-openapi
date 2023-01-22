# AfEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AfEvent**](AfEvent.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SvcExprcInfos** | Pointer to [**[]ServiceExperienceInfoPerApp**](ServiceExperienceInfoPerApp.md) |  | [optional] 
**UeMobilityInfos** | Pointer to [**[]UeMobilityCollection**](UeMobilityCollection.md) |  | [optional] 
**UeCommInfos** | Pointer to [**[]UeCommunicationCollection**](UeCommunicationCollection.md) |  | [optional] 
**ExcepInfos** | Pointer to [**[]ExceptionInfo**](ExceptionInfo.md) |  | [optional] 
**CongestionInfos** | Pointer to [**[]UserDataCongestionCollection**](UserDataCongestionCollection.md) |  | [optional] 
**PerfDataInfos** | Pointer to [**[]PerformanceDataCollection**](PerformanceDataCollection.md) |  | [optional] 
**DispersionInfos** | Pointer to [**[]DispersionCollection1**](DispersionCollection1.md) |  | [optional] 
**CollBhvrInfs** | Pointer to [**[]CollectiveBehaviourInfo**](CollectiveBehaviourInfo.md) |  | [optional] 
**MsQoeMetrInfos** | Pointer to [**[]MsQoeMetricsCollection**](MsQoeMetricsCollection.md) |  | [optional] 
**MsConsumpInfos** | Pointer to [**[]MsConsumptionCollection**](MsConsumptionCollection.md) |  | [optional] 
**MsNetAssInvInfos** | Pointer to [**[]MsNetAssInvocationCollection**](MsNetAssInvocationCollection.md) |  | [optional] 
**MsDynPlyInvInfos** | Pointer to [**[]MsDynPolicyInvocationCollection**](MsDynPolicyInvocationCollection.md) |  | [optional] 
**MsAccActInfos** | Pointer to [**[]MSAccessActivityCollection**](MSAccessActivityCollection.md) |  | [optional] 

## Methods

### NewAfEventNotification

`func NewAfEventNotification(event AfEvent, timeStamp time.Time, ) *AfEventNotification`

NewAfEventNotification instantiates a new AfEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfEventNotificationWithDefaults

`func NewAfEventNotificationWithDefaults() *AfEventNotification`

NewAfEventNotificationWithDefaults instantiates a new AfEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AfEventNotification) GetEvent() AfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AfEventNotification) GetEventOk() (*AfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AfEventNotification) SetEvent(v AfEvent)`

SetEvent sets Event field to given value.


### GetTimeStamp

`func (o *AfEventNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AfEventNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AfEventNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSvcExprcInfos

`func (o *AfEventNotification) GetSvcExprcInfos() []ServiceExperienceInfoPerApp`

GetSvcExprcInfos returns the SvcExprcInfos field if non-nil, zero value otherwise.

### GetSvcExprcInfosOk

`func (o *AfEventNotification) GetSvcExprcInfosOk() (*[]ServiceExperienceInfoPerApp, bool)`

GetSvcExprcInfosOk returns a tuple with the SvcExprcInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExprcInfos

`func (o *AfEventNotification) SetSvcExprcInfos(v []ServiceExperienceInfoPerApp)`

SetSvcExprcInfos sets SvcExprcInfos field to given value.

### HasSvcExprcInfos

`func (o *AfEventNotification) HasSvcExprcInfos() bool`

HasSvcExprcInfos returns a boolean if a field has been set.

### GetUeMobilityInfos

`func (o *AfEventNotification) GetUeMobilityInfos() []UeMobilityCollection`

GetUeMobilityInfos returns the UeMobilityInfos field if non-nil, zero value otherwise.

### GetUeMobilityInfosOk

`func (o *AfEventNotification) GetUeMobilityInfosOk() (*[]UeMobilityCollection, bool)`

GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobilityInfos

`func (o *AfEventNotification) SetUeMobilityInfos(v []UeMobilityCollection)`

SetUeMobilityInfos sets UeMobilityInfos field to given value.

### HasUeMobilityInfos

`func (o *AfEventNotification) HasUeMobilityInfos() bool`

HasUeMobilityInfos returns a boolean if a field has been set.

### GetUeCommInfos

`func (o *AfEventNotification) GetUeCommInfos() []UeCommunicationCollection`

GetUeCommInfos returns the UeCommInfos field if non-nil, zero value otherwise.

### GetUeCommInfosOk

`func (o *AfEventNotification) GetUeCommInfosOk() (*[]UeCommunicationCollection, bool)`

GetUeCommInfosOk returns a tuple with the UeCommInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCommInfos

`func (o *AfEventNotification) SetUeCommInfos(v []UeCommunicationCollection)`

SetUeCommInfos sets UeCommInfos field to given value.

### HasUeCommInfos

`func (o *AfEventNotification) HasUeCommInfos() bool`

HasUeCommInfos returns a boolean if a field has been set.

### GetExcepInfos

`func (o *AfEventNotification) GetExcepInfos() []ExceptionInfo`

GetExcepInfos returns the ExcepInfos field if non-nil, zero value otherwise.

### GetExcepInfosOk

`func (o *AfEventNotification) GetExcepInfosOk() (*[]ExceptionInfo, bool)`

GetExcepInfosOk returns a tuple with the ExcepInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepInfos

`func (o *AfEventNotification) SetExcepInfos(v []ExceptionInfo)`

SetExcepInfos sets ExcepInfos field to given value.

### HasExcepInfos

`func (o *AfEventNotification) HasExcepInfos() bool`

HasExcepInfos returns a boolean if a field has been set.

### GetCongestionInfos

`func (o *AfEventNotification) GetCongestionInfos() []UserDataCongestionCollection`

GetCongestionInfos returns the CongestionInfos field if non-nil, zero value otherwise.

### GetCongestionInfosOk

`func (o *AfEventNotification) GetCongestionInfosOk() (*[]UserDataCongestionCollection, bool)`

GetCongestionInfosOk returns a tuple with the CongestionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestionInfos

`func (o *AfEventNotification) SetCongestionInfos(v []UserDataCongestionCollection)`

SetCongestionInfos sets CongestionInfos field to given value.

### HasCongestionInfos

`func (o *AfEventNotification) HasCongestionInfos() bool`

HasCongestionInfos returns a boolean if a field has been set.

### GetPerfDataInfos

`func (o *AfEventNotification) GetPerfDataInfos() []PerformanceDataCollection`

GetPerfDataInfos returns the PerfDataInfos field if non-nil, zero value otherwise.

### GetPerfDataInfosOk

`func (o *AfEventNotification) GetPerfDataInfosOk() (*[]PerformanceDataCollection, bool)`

GetPerfDataInfosOk returns a tuple with the PerfDataInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfDataInfos

`func (o *AfEventNotification) SetPerfDataInfos(v []PerformanceDataCollection)`

SetPerfDataInfos sets PerfDataInfos field to given value.

### HasPerfDataInfos

`func (o *AfEventNotification) HasPerfDataInfos() bool`

HasPerfDataInfos returns a boolean if a field has been set.

### GetDispersionInfos

`func (o *AfEventNotification) GetDispersionInfos() []DispersionCollection1`

GetDispersionInfos returns the DispersionInfos field if non-nil, zero value otherwise.

### GetDispersionInfosOk

`func (o *AfEventNotification) GetDispersionInfosOk() (*[]DispersionCollection1, bool)`

GetDispersionInfosOk returns a tuple with the DispersionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispersionInfos

`func (o *AfEventNotification) SetDispersionInfos(v []DispersionCollection1)`

SetDispersionInfos sets DispersionInfos field to given value.

### HasDispersionInfos

`func (o *AfEventNotification) HasDispersionInfos() bool`

HasDispersionInfos returns a boolean if a field has been set.

### GetCollBhvrInfs

`func (o *AfEventNotification) GetCollBhvrInfs() []CollectiveBehaviourInfo`

GetCollBhvrInfs returns the CollBhvrInfs field if non-nil, zero value otherwise.

### GetCollBhvrInfsOk

`func (o *AfEventNotification) GetCollBhvrInfsOk() (*[]CollectiveBehaviourInfo, bool)`

GetCollBhvrInfsOk returns a tuple with the CollBhvrInfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollBhvrInfs

`func (o *AfEventNotification) SetCollBhvrInfs(v []CollectiveBehaviourInfo)`

SetCollBhvrInfs sets CollBhvrInfs field to given value.

### HasCollBhvrInfs

`func (o *AfEventNotification) HasCollBhvrInfs() bool`

HasCollBhvrInfs returns a boolean if a field has been set.

### GetMsQoeMetrInfos

`func (o *AfEventNotification) GetMsQoeMetrInfos() []MsQoeMetricsCollection`

GetMsQoeMetrInfos returns the MsQoeMetrInfos field if non-nil, zero value otherwise.

### GetMsQoeMetrInfosOk

`func (o *AfEventNotification) GetMsQoeMetrInfosOk() (*[]MsQoeMetricsCollection, bool)`

GetMsQoeMetrInfosOk returns a tuple with the MsQoeMetrInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsQoeMetrInfos

`func (o *AfEventNotification) SetMsQoeMetrInfos(v []MsQoeMetricsCollection)`

SetMsQoeMetrInfos sets MsQoeMetrInfos field to given value.

### HasMsQoeMetrInfos

`func (o *AfEventNotification) HasMsQoeMetrInfos() bool`

HasMsQoeMetrInfos returns a boolean if a field has been set.

### GetMsConsumpInfos

`func (o *AfEventNotification) GetMsConsumpInfos() []MsConsumptionCollection`

GetMsConsumpInfos returns the MsConsumpInfos field if non-nil, zero value otherwise.

### GetMsConsumpInfosOk

`func (o *AfEventNotification) GetMsConsumpInfosOk() (*[]MsConsumptionCollection, bool)`

GetMsConsumpInfosOk returns a tuple with the MsConsumpInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsConsumpInfos

`func (o *AfEventNotification) SetMsConsumpInfos(v []MsConsumptionCollection)`

SetMsConsumpInfos sets MsConsumpInfos field to given value.

### HasMsConsumpInfos

`func (o *AfEventNotification) HasMsConsumpInfos() bool`

HasMsConsumpInfos returns a boolean if a field has been set.

### GetMsNetAssInvInfos

`func (o *AfEventNotification) GetMsNetAssInvInfos() []MsNetAssInvocationCollection`

GetMsNetAssInvInfos returns the MsNetAssInvInfos field if non-nil, zero value otherwise.

### GetMsNetAssInvInfosOk

`func (o *AfEventNotification) GetMsNetAssInvInfosOk() (*[]MsNetAssInvocationCollection, bool)`

GetMsNetAssInvInfosOk returns a tuple with the MsNetAssInvInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsNetAssInvInfos

`func (o *AfEventNotification) SetMsNetAssInvInfos(v []MsNetAssInvocationCollection)`

SetMsNetAssInvInfos sets MsNetAssInvInfos field to given value.

### HasMsNetAssInvInfos

`func (o *AfEventNotification) HasMsNetAssInvInfos() bool`

HasMsNetAssInvInfos returns a boolean if a field has been set.

### GetMsDynPlyInvInfos

`func (o *AfEventNotification) GetMsDynPlyInvInfos() []MsDynPolicyInvocationCollection`

GetMsDynPlyInvInfos returns the MsDynPlyInvInfos field if non-nil, zero value otherwise.

### GetMsDynPlyInvInfosOk

`func (o *AfEventNotification) GetMsDynPlyInvInfosOk() (*[]MsDynPolicyInvocationCollection, bool)`

GetMsDynPlyInvInfosOk returns a tuple with the MsDynPlyInvInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDynPlyInvInfos

`func (o *AfEventNotification) SetMsDynPlyInvInfos(v []MsDynPolicyInvocationCollection)`

SetMsDynPlyInvInfos sets MsDynPlyInvInfos field to given value.

### HasMsDynPlyInvInfos

`func (o *AfEventNotification) HasMsDynPlyInvInfos() bool`

HasMsDynPlyInvInfos returns a boolean if a field has been set.

### GetMsAccActInfos

`func (o *AfEventNotification) GetMsAccActInfos() []MSAccessActivityCollection`

GetMsAccActInfos returns the MsAccActInfos field if non-nil, zero value otherwise.

### GetMsAccActInfosOk

`func (o *AfEventNotification) GetMsAccActInfosOk() (*[]MSAccessActivityCollection, bool)`

GetMsAccActInfosOk returns a tuple with the MsAccActInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAccActInfos

`func (o *AfEventNotification) SetMsAccActInfos(v []MSAccessActivityCollection)`

SetMsAccActInfos sets MsAccActInfos field to given value.

### HasMsAccActInfos

`func (o *AfEventNotification) HasMsAccActInfos() bool`

HasMsAccActInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


