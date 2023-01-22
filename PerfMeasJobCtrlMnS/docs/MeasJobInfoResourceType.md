# MeasJobInfoResourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**IOCName** | Pointer to **string** |  | [optional] 
**IOCInstanceList** | Pointer to **[]string** |  | [optional] 
**MeasurementCategoryList** | Pointer to **[]string** |  | [optional] 
**ReportingMethod** | Pointer to [**ReportingMethodType**](ReportingMethodType.md) |  | [optional] 
**GranularityPeriod** | Pointer to **int32** |  | [optional] 
**ReportingPeriod** | Pointer to **int32** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**StopTime** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**ScheduleType**](ScheduleType.md) |  | [optional] 
**StreamTarget** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**PriorityType**](PriorityType.md) |  | [optional] 
**Reliability** | Pointer to **string** |  | [optional] 

## Methods

### NewMeasJobInfoResourceType

`func NewMeasJobInfoResourceType() *MeasJobInfoResourceType`

NewMeasJobInfoResourceType instantiates a new MeasJobInfoResourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasJobInfoResourceTypeWithDefaults

`func NewMeasJobInfoResourceTypeWithDefaults() *MeasJobInfoResourceType`

NewMeasJobInfoResourceTypeWithDefaults instantiates a new MeasJobInfoResourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *MeasJobInfoResourceType) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MeasJobInfoResourceType) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MeasJobInfoResourceType) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MeasJobInfoResourceType) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetIOCName

`func (o *MeasJobInfoResourceType) GetIOCName() string`

GetIOCName returns the IOCName field if non-nil, zero value otherwise.

### GetIOCNameOk

`func (o *MeasJobInfoResourceType) GetIOCNameOk() (*string, bool)`

GetIOCNameOk returns a tuple with the IOCName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOCName

`func (o *MeasJobInfoResourceType) SetIOCName(v string)`

SetIOCName sets IOCName field to given value.

### HasIOCName

`func (o *MeasJobInfoResourceType) HasIOCName() bool`

HasIOCName returns a boolean if a field has been set.

### GetIOCInstanceList

`func (o *MeasJobInfoResourceType) GetIOCInstanceList() []string`

GetIOCInstanceList returns the IOCInstanceList field if non-nil, zero value otherwise.

### GetIOCInstanceListOk

`func (o *MeasJobInfoResourceType) GetIOCInstanceListOk() (*[]string, bool)`

GetIOCInstanceListOk returns a tuple with the IOCInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOCInstanceList

`func (o *MeasJobInfoResourceType) SetIOCInstanceList(v []string)`

SetIOCInstanceList sets IOCInstanceList field to given value.

### HasIOCInstanceList

`func (o *MeasJobInfoResourceType) HasIOCInstanceList() bool`

HasIOCInstanceList returns a boolean if a field has been set.

### GetMeasurementCategoryList

`func (o *MeasJobInfoResourceType) GetMeasurementCategoryList() []string`

GetMeasurementCategoryList returns the MeasurementCategoryList field if non-nil, zero value otherwise.

### GetMeasurementCategoryListOk

`func (o *MeasJobInfoResourceType) GetMeasurementCategoryListOk() (*[]string, bool)`

GetMeasurementCategoryListOk returns a tuple with the MeasurementCategoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementCategoryList

`func (o *MeasJobInfoResourceType) SetMeasurementCategoryList(v []string)`

SetMeasurementCategoryList sets MeasurementCategoryList field to given value.

### HasMeasurementCategoryList

`func (o *MeasJobInfoResourceType) HasMeasurementCategoryList() bool`

HasMeasurementCategoryList returns a boolean if a field has been set.

### GetReportingMethod

`func (o *MeasJobInfoResourceType) GetReportingMethod() ReportingMethodType`

GetReportingMethod returns the ReportingMethod field if non-nil, zero value otherwise.

### GetReportingMethodOk

`func (o *MeasJobInfoResourceType) GetReportingMethodOk() (*ReportingMethodType, bool)`

GetReportingMethodOk returns a tuple with the ReportingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingMethod

`func (o *MeasJobInfoResourceType) SetReportingMethod(v ReportingMethodType)`

SetReportingMethod sets ReportingMethod field to given value.

### HasReportingMethod

`func (o *MeasJobInfoResourceType) HasReportingMethod() bool`

HasReportingMethod returns a boolean if a field has been set.

### GetGranularityPeriod

`func (o *MeasJobInfoResourceType) GetGranularityPeriod() int32`

GetGranularityPeriod returns the GranularityPeriod field if non-nil, zero value otherwise.

### GetGranularityPeriodOk

`func (o *MeasJobInfoResourceType) GetGranularityPeriodOk() (*int32, bool)`

GetGranularityPeriodOk returns a tuple with the GranularityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityPeriod

`func (o *MeasJobInfoResourceType) SetGranularityPeriod(v int32)`

SetGranularityPeriod sets GranularityPeriod field to given value.

### HasGranularityPeriod

`func (o *MeasJobInfoResourceType) HasGranularityPeriod() bool`

HasGranularityPeriod returns a boolean if a field has been set.

### GetReportingPeriod

`func (o *MeasJobInfoResourceType) GetReportingPeriod() int32`

GetReportingPeriod returns the ReportingPeriod field if non-nil, zero value otherwise.

### GetReportingPeriodOk

`func (o *MeasJobInfoResourceType) GetReportingPeriodOk() (*int32, bool)`

GetReportingPeriodOk returns a tuple with the ReportingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingPeriod

`func (o *MeasJobInfoResourceType) SetReportingPeriod(v int32)`

SetReportingPeriod sets ReportingPeriod field to given value.

### HasReportingPeriod

`func (o *MeasJobInfoResourceType) HasReportingPeriod() bool`

HasReportingPeriod returns a boolean if a field has been set.

### GetStartTime

`func (o *MeasJobInfoResourceType) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MeasJobInfoResourceType) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MeasJobInfoResourceType) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MeasJobInfoResourceType) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStopTime

`func (o *MeasJobInfoResourceType) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *MeasJobInfoResourceType) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *MeasJobInfoResourceType) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *MeasJobInfoResourceType) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### GetSchedule

`func (o *MeasJobInfoResourceType) GetSchedule() ScheduleType`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MeasJobInfoResourceType) GetScheduleOk() (*ScheduleType, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MeasJobInfoResourceType) SetSchedule(v ScheduleType)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MeasJobInfoResourceType) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStreamTarget

`func (o *MeasJobInfoResourceType) GetStreamTarget() string`

GetStreamTarget returns the StreamTarget field if non-nil, zero value otherwise.

### GetStreamTargetOk

`func (o *MeasJobInfoResourceType) GetStreamTargetOk() (*string, bool)`

GetStreamTargetOk returns a tuple with the StreamTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamTarget

`func (o *MeasJobInfoResourceType) SetStreamTarget(v string)`

SetStreamTarget sets StreamTarget field to given value.

### HasStreamTarget

`func (o *MeasJobInfoResourceType) HasStreamTarget() bool`

HasStreamTarget returns a boolean if a field has been set.

### GetPriority

`func (o *MeasJobInfoResourceType) GetPriority() PriorityType`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MeasJobInfoResourceType) GetPriorityOk() (*PriorityType, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MeasJobInfoResourceType) SetPriority(v PriorityType)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MeasJobInfoResourceType) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReliability

`func (o *MeasJobInfoResourceType) GetReliability() string`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *MeasJobInfoResourceType) GetReliabilityOk() (*string, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *MeasJobInfoResourceType) SetReliability(v string)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *MeasJobInfoResourceType) HasReliability() bool`

HasReliability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


