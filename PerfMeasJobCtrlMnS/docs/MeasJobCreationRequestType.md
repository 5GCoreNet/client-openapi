# MeasJobCreationRequestType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewMeasJobCreationRequestType

`func NewMeasJobCreationRequestType() *MeasJobCreationRequestType`

NewMeasJobCreationRequestType instantiates a new MeasJobCreationRequestType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasJobCreationRequestTypeWithDefaults

`func NewMeasJobCreationRequestTypeWithDefaults() *MeasJobCreationRequestType`

NewMeasJobCreationRequestTypeWithDefaults instantiates a new MeasJobCreationRequestType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIOCName

`func (o *MeasJobCreationRequestType) GetIOCName() string`

GetIOCName returns the IOCName field if non-nil, zero value otherwise.

### GetIOCNameOk

`func (o *MeasJobCreationRequestType) GetIOCNameOk() (*string, bool)`

GetIOCNameOk returns a tuple with the IOCName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOCName

`func (o *MeasJobCreationRequestType) SetIOCName(v string)`

SetIOCName sets IOCName field to given value.

### HasIOCName

`func (o *MeasJobCreationRequestType) HasIOCName() bool`

HasIOCName returns a boolean if a field has been set.

### GetIOCInstanceList

`func (o *MeasJobCreationRequestType) GetIOCInstanceList() []string`

GetIOCInstanceList returns the IOCInstanceList field if non-nil, zero value otherwise.

### GetIOCInstanceListOk

`func (o *MeasJobCreationRequestType) GetIOCInstanceListOk() (*[]string, bool)`

GetIOCInstanceListOk returns a tuple with the IOCInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOCInstanceList

`func (o *MeasJobCreationRequestType) SetIOCInstanceList(v []string)`

SetIOCInstanceList sets IOCInstanceList field to given value.

### HasIOCInstanceList

`func (o *MeasJobCreationRequestType) HasIOCInstanceList() bool`

HasIOCInstanceList returns a boolean if a field has been set.

### GetMeasurementCategoryList

`func (o *MeasJobCreationRequestType) GetMeasurementCategoryList() []string`

GetMeasurementCategoryList returns the MeasurementCategoryList field if non-nil, zero value otherwise.

### GetMeasurementCategoryListOk

`func (o *MeasJobCreationRequestType) GetMeasurementCategoryListOk() (*[]string, bool)`

GetMeasurementCategoryListOk returns a tuple with the MeasurementCategoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementCategoryList

`func (o *MeasJobCreationRequestType) SetMeasurementCategoryList(v []string)`

SetMeasurementCategoryList sets MeasurementCategoryList field to given value.

### HasMeasurementCategoryList

`func (o *MeasJobCreationRequestType) HasMeasurementCategoryList() bool`

HasMeasurementCategoryList returns a boolean if a field has been set.

### GetReportingMethod

`func (o *MeasJobCreationRequestType) GetReportingMethod() ReportingMethodType`

GetReportingMethod returns the ReportingMethod field if non-nil, zero value otherwise.

### GetReportingMethodOk

`func (o *MeasJobCreationRequestType) GetReportingMethodOk() (*ReportingMethodType, bool)`

GetReportingMethodOk returns a tuple with the ReportingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingMethod

`func (o *MeasJobCreationRequestType) SetReportingMethod(v ReportingMethodType)`

SetReportingMethod sets ReportingMethod field to given value.

### HasReportingMethod

`func (o *MeasJobCreationRequestType) HasReportingMethod() bool`

HasReportingMethod returns a boolean if a field has been set.

### GetGranularityPeriod

`func (o *MeasJobCreationRequestType) GetGranularityPeriod() int32`

GetGranularityPeriod returns the GranularityPeriod field if non-nil, zero value otherwise.

### GetGranularityPeriodOk

`func (o *MeasJobCreationRequestType) GetGranularityPeriodOk() (*int32, bool)`

GetGranularityPeriodOk returns a tuple with the GranularityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityPeriod

`func (o *MeasJobCreationRequestType) SetGranularityPeriod(v int32)`

SetGranularityPeriod sets GranularityPeriod field to given value.

### HasGranularityPeriod

`func (o *MeasJobCreationRequestType) HasGranularityPeriod() bool`

HasGranularityPeriod returns a boolean if a field has been set.

### GetReportingPeriod

`func (o *MeasJobCreationRequestType) GetReportingPeriod() int32`

GetReportingPeriod returns the ReportingPeriod field if non-nil, zero value otherwise.

### GetReportingPeriodOk

`func (o *MeasJobCreationRequestType) GetReportingPeriodOk() (*int32, bool)`

GetReportingPeriodOk returns a tuple with the ReportingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingPeriod

`func (o *MeasJobCreationRequestType) SetReportingPeriod(v int32)`

SetReportingPeriod sets ReportingPeriod field to given value.

### HasReportingPeriod

`func (o *MeasJobCreationRequestType) HasReportingPeriod() bool`

HasReportingPeriod returns a boolean if a field has been set.

### GetStartTime

`func (o *MeasJobCreationRequestType) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MeasJobCreationRequestType) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MeasJobCreationRequestType) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MeasJobCreationRequestType) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStopTime

`func (o *MeasJobCreationRequestType) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *MeasJobCreationRequestType) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *MeasJobCreationRequestType) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *MeasJobCreationRequestType) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### GetSchedule

`func (o *MeasJobCreationRequestType) GetSchedule() ScheduleType`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MeasJobCreationRequestType) GetScheduleOk() (*ScheduleType, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MeasJobCreationRequestType) SetSchedule(v ScheduleType)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MeasJobCreationRequestType) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStreamTarget

`func (o *MeasJobCreationRequestType) GetStreamTarget() string`

GetStreamTarget returns the StreamTarget field if non-nil, zero value otherwise.

### GetStreamTargetOk

`func (o *MeasJobCreationRequestType) GetStreamTargetOk() (*string, bool)`

GetStreamTargetOk returns a tuple with the StreamTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamTarget

`func (o *MeasJobCreationRequestType) SetStreamTarget(v string)`

SetStreamTarget sets StreamTarget field to given value.

### HasStreamTarget

`func (o *MeasJobCreationRequestType) HasStreamTarget() bool`

HasStreamTarget returns a boolean if a field has been set.

### GetPriority

`func (o *MeasJobCreationRequestType) GetPriority() PriorityType`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MeasJobCreationRequestType) GetPriorityOk() (*PriorityType, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MeasJobCreationRequestType) SetPriority(v PriorityType)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MeasJobCreationRequestType) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReliability

`func (o *MeasJobCreationRequestType) GetReliability() string`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *MeasJobCreationRequestType) GetReliabilityOk() (*string, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *MeasJobCreationRequestType) SetReliability(v string)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *MeasJobCreationRequestType) HasReliability() bool`

HasReliability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


