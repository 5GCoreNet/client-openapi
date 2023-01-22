# NefEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NefEvent**](NefEvent.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SvcExprcInfos** | Pointer to [**[]ServiceExperienceInfo**](ServiceExperienceInfo.md) |  | [optional] 
**UeMobilityInfos** | Pointer to [**[]UeMobilityInfo**](UeMobilityInfo.md) |  | [optional] 
**UeCommInfos** | Pointer to [**[]UeCommunicationInfo**](UeCommunicationInfo.md) |  | [optional] 
**ExcepInfos** | Pointer to [**[]ExceptionInfo**](ExceptionInfo.md) |  | [optional] 
**CongestionInfos** | Pointer to [**[]UserDataCongestionCollection**](UserDataCongestionCollection.md) |  | [optional] 
**PerfDataInfos** | Pointer to [**[]PerformanceDataInfo**](PerformanceDataInfo.md) |  | [optional] 
**DispersionInfos** | Pointer to [**[]DispersionCollection**](DispersionCollection.md) |  | [optional] 
**CollBhvrInfs** | Pointer to [**[]CollectiveBehaviourInfo**](CollectiveBehaviourInfo.md) |  | [optional] 
**MsQoeMetrInfos** | Pointer to [**[]MsQoeMetricsCollection**](MsQoeMetricsCollection.md) |  | [optional] 
**MsConsumpInfos** | Pointer to [**[]MsConsumptionCollection**](MsConsumptionCollection.md) |  | [optional] 
**MsNetAssInvInfos** | Pointer to [**[]MsNetAssInvocationCollection**](MsNetAssInvocationCollection.md) |  | [optional] 
**MsDynPlyInvInfos** | Pointer to [**[]MsDynPolicyInvocationCollection**](MsDynPolicyInvocationCollection.md) |  | [optional] 
**MsAccActInfos** | Pointer to [**[]MSAccessActivityCollection**](MSAccessActivityCollection.md) |  | [optional] 

## Methods

### NewNefEventNotification

`func NewNefEventNotification(event NefEvent, timeStamp time.Time, ) *NefEventNotification`

NewNefEventNotification instantiates a new NefEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefEventNotificationWithDefaults

`func NewNefEventNotificationWithDefaults() *NefEventNotification`

NewNefEventNotificationWithDefaults instantiates a new NefEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NefEventNotification) GetEvent() NefEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NefEventNotification) GetEventOk() (*NefEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NefEventNotification) SetEvent(v NefEvent)`

SetEvent sets Event field to given value.


### GetTimeStamp

`func (o *NefEventNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *NefEventNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *NefEventNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSvcExprcInfos

`func (o *NefEventNotification) GetSvcExprcInfos() []ServiceExperienceInfo`

GetSvcExprcInfos returns the SvcExprcInfos field if non-nil, zero value otherwise.

### GetSvcExprcInfosOk

`func (o *NefEventNotification) GetSvcExprcInfosOk() (*[]ServiceExperienceInfo, bool)`

GetSvcExprcInfosOk returns a tuple with the SvcExprcInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExprcInfos

`func (o *NefEventNotification) SetSvcExprcInfos(v []ServiceExperienceInfo)`

SetSvcExprcInfos sets SvcExprcInfos field to given value.

### HasSvcExprcInfos

`func (o *NefEventNotification) HasSvcExprcInfos() bool`

HasSvcExprcInfos returns a boolean if a field has been set.

### GetUeMobilityInfos

`func (o *NefEventNotification) GetUeMobilityInfos() []UeMobilityInfo`

GetUeMobilityInfos returns the UeMobilityInfos field if non-nil, zero value otherwise.

### GetUeMobilityInfosOk

`func (o *NefEventNotification) GetUeMobilityInfosOk() (*[]UeMobilityInfo, bool)`

GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobilityInfos

`func (o *NefEventNotification) SetUeMobilityInfos(v []UeMobilityInfo)`

SetUeMobilityInfos sets UeMobilityInfos field to given value.

### HasUeMobilityInfos

`func (o *NefEventNotification) HasUeMobilityInfos() bool`

HasUeMobilityInfos returns a boolean if a field has been set.

### GetUeCommInfos

`func (o *NefEventNotification) GetUeCommInfos() []UeCommunicationInfo`

GetUeCommInfos returns the UeCommInfos field if non-nil, zero value otherwise.

### GetUeCommInfosOk

`func (o *NefEventNotification) GetUeCommInfosOk() (*[]UeCommunicationInfo, bool)`

GetUeCommInfosOk returns a tuple with the UeCommInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCommInfos

`func (o *NefEventNotification) SetUeCommInfos(v []UeCommunicationInfo)`

SetUeCommInfos sets UeCommInfos field to given value.

### HasUeCommInfos

`func (o *NefEventNotification) HasUeCommInfos() bool`

HasUeCommInfos returns a boolean if a field has been set.

### GetExcepInfos

`func (o *NefEventNotification) GetExcepInfos() []ExceptionInfo`

GetExcepInfos returns the ExcepInfos field if non-nil, zero value otherwise.

### GetExcepInfosOk

`func (o *NefEventNotification) GetExcepInfosOk() (*[]ExceptionInfo, bool)`

GetExcepInfosOk returns a tuple with the ExcepInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepInfos

`func (o *NefEventNotification) SetExcepInfos(v []ExceptionInfo)`

SetExcepInfos sets ExcepInfos field to given value.

### HasExcepInfos

`func (o *NefEventNotification) HasExcepInfos() bool`

HasExcepInfos returns a boolean if a field has been set.

### GetCongestionInfos

`func (o *NefEventNotification) GetCongestionInfos() []UserDataCongestionCollection`

GetCongestionInfos returns the CongestionInfos field if non-nil, zero value otherwise.

### GetCongestionInfosOk

`func (o *NefEventNotification) GetCongestionInfosOk() (*[]UserDataCongestionCollection, bool)`

GetCongestionInfosOk returns a tuple with the CongestionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestionInfos

`func (o *NefEventNotification) SetCongestionInfos(v []UserDataCongestionCollection)`

SetCongestionInfos sets CongestionInfos field to given value.

### HasCongestionInfos

`func (o *NefEventNotification) HasCongestionInfos() bool`

HasCongestionInfos returns a boolean if a field has been set.

### GetPerfDataInfos

`func (o *NefEventNotification) GetPerfDataInfos() []PerformanceDataInfo`

GetPerfDataInfos returns the PerfDataInfos field if non-nil, zero value otherwise.

### GetPerfDataInfosOk

`func (o *NefEventNotification) GetPerfDataInfosOk() (*[]PerformanceDataInfo, bool)`

GetPerfDataInfosOk returns a tuple with the PerfDataInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfDataInfos

`func (o *NefEventNotification) SetPerfDataInfos(v []PerformanceDataInfo)`

SetPerfDataInfos sets PerfDataInfos field to given value.

### HasPerfDataInfos

`func (o *NefEventNotification) HasPerfDataInfos() bool`

HasPerfDataInfos returns a boolean if a field has been set.

### GetDispersionInfos

`func (o *NefEventNotification) GetDispersionInfos() []DispersionCollection`

GetDispersionInfos returns the DispersionInfos field if non-nil, zero value otherwise.

### GetDispersionInfosOk

`func (o *NefEventNotification) GetDispersionInfosOk() (*[]DispersionCollection, bool)`

GetDispersionInfosOk returns a tuple with the DispersionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispersionInfos

`func (o *NefEventNotification) SetDispersionInfos(v []DispersionCollection)`

SetDispersionInfos sets DispersionInfos field to given value.

### HasDispersionInfos

`func (o *NefEventNotification) HasDispersionInfos() bool`

HasDispersionInfos returns a boolean if a field has been set.

### GetCollBhvrInfs

`func (o *NefEventNotification) GetCollBhvrInfs() []CollectiveBehaviourInfo`

GetCollBhvrInfs returns the CollBhvrInfs field if non-nil, zero value otherwise.

### GetCollBhvrInfsOk

`func (o *NefEventNotification) GetCollBhvrInfsOk() (*[]CollectiveBehaviourInfo, bool)`

GetCollBhvrInfsOk returns a tuple with the CollBhvrInfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollBhvrInfs

`func (o *NefEventNotification) SetCollBhvrInfs(v []CollectiveBehaviourInfo)`

SetCollBhvrInfs sets CollBhvrInfs field to given value.

### HasCollBhvrInfs

`func (o *NefEventNotification) HasCollBhvrInfs() bool`

HasCollBhvrInfs returns a boolean if a field has been set.

### GetMsQoeMetrInfos

`func (o *NefEventNotification) GetMsQoeMetrInfos() []MsQoeMetricsCollection`

GetMsQoeMetrInfos returns the MsQoeMetrInfos field if non-nil, zero value otherwise.

### GetMsQoeMetrInfosOk

`func (o *NefEventNotification) GetMsQoeMetrInfosOk() (*[]MsQoeMetricsCollection, bool)`

GetMsQoeMetrInfosOk returns a tuple with the MsQoeMetrInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsQoeMetrInfos

`func (o *NefEventNotification) SetMsQoeMetrInfos(v []MsQoeMetricsCollection)`

SetMsQoeMetrInfos sets MsQoeMetrInfos field to given value.

### HasMsQoeMetrInfos

`func (o *NefEventNotification) HasMsQoeMetrInfos() bool`

HasMsQoeMetrInfos returns a boolean if a field has been set.

### GetMsConsumpInfos

`func (o *NefEventNotification) GetMsConsumpInfos() []MsConsumptionCollection`

GetMsConsumpInfos returns the MsConsumpInfos field if non-nil, zero value otherwise.

### GetMsConsumpInfosOk

`func (o *NefEventNotification) GetMsConsumpInfosOk() (*[]MsConsumptionCollection, bool)`

GetMsConsumpInfosOk returns a tuple with the MsConsumpInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsConsumpInfos

`func (o *NefEventNotification) SetMsConsumpInfos(v []MsConsumptionCollection)`

SetMsConsumpInfos sets MsConsumpInfos field to given value.

### HasMsConsumpInfos

`func (o *NefEventNotification) HasMsConsumpInfos() bool`

HasMsConsumpInfos returns a boolean if a field has been set.

### GetMsNetAssInvInfos

`func (o *NefEventNotification) GetMsNetAssInvInfos() []MsNetAssInvocationCollection`

GetMsNetAssInvInfos returns the MsNetAssInvInfos field if non-nil, zero value otherwise.

### GetMsNetAssInvInfosOk

`func (o *NefEventNotification) GetMsNetAssInvInfosOk() (*[]MsNetAssInvocationCollection, bool)`

GetMsNetAssInvInfosOk returns a tuple with the MsNetAssInvInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsNetAssInvInfos

`func (o *NefEventNotification) SetMsNetAssInvInfos(v []MsNetAssInvocationCollection)`

SetMsNetAssInvInfos sets MsNetAssInvInfos field to given value.

### HasMsNetAssInvInfos

`func (o *NefEventNotification) HasMsNetAssInvInfos() bool`

HasMsNetAssInvInfos returns a boolean if a field has been set.

### GetMsDynPlyInvInfos

`func (o *NefEventNotification) GetMsDynPlyInvInfos() []MsDynPolicyInvocationCollection`

GetMsDynPlyInvInfos returns the MsDynPlyInvInfos field if non-nil, zero value otherwise.

### GetMsDynPlyInvInfosOk

`func (o *NefEventNotification) GetMsDynPlyInvInfosOk() (*[]MsDynPolicyInvocationCollection, bool)`

GetMsDynPlyInvInfosOk returns a tuple with the MsDynPlyInvInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDynPlyInvInfos

`func (o *NefEventNotification) SetMsDynPlyInvInfos(v []MsDynPolicyInvocationCollection)`

SetMsDynPlyInvInfos sets MsDynPlyInvInfos field to given value.

### HasMsDynPlyInvInfos

`func (o *NefEventNotification) HasMsDynPlyInvInfos() bool`

HasMsDynPlyInvInfos returns a boolean if a field has been set.

### GetMsAccActInfos

`func (o *NefEventNotification) GetMsAccActInfos() []MSAccessActivityCollection`

GetMsAccActInfos returns the MsAccActInfos field if non-nil, zero value otherwise.

### GetMsAccActInfosOk

`func (o *NefEventNotification) GetMsAccActInfosOk() (*[]MSAccessActivityCollection, bool)`

GetMsAccActInfosOk returns a tuple with the MsAccActInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAccActInfos

`func (o *NefEventNotification) SetMsAccActInfos(v []MSAccessActivityCollection)`

SetMsAccActInfos sets MsAccActInfos field to given value.

### HasMsAccActInfos

`func (o *NefEventNotification) HasMsAccActInfos() bool`

HasMsAccActInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


