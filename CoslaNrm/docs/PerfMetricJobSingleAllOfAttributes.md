# PerfMetricJobSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativeState** | Pointer to [**AdministrativeState**](AdministrativeState.md) |  | [optional] 
**OperationalState** | Pointer to [**OperationalState**](OperationalState.md) |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**PerformanceMetrics** | Pointer to **[]string** |  | [optional] 
**GranularityPeriod** | Pointer to **int32** |  | [optional] 
**ObjectInstances** | Pointer to **[]string** |  | [optional] 
**RootObjectInstances** | Pointer to **[]string** |  | [optional] 
**ReportingCtrl** | Pointer to [**ReportingCtrl**](ReportingCtrl.md) |  | [optional] 

## Methods

### NewPerfMetricJobSingleAllOfAttributes

`func NewPerfMetricJobSingleAllOfAttributes() *PerfMetricJobSingleAllOfAttributes`

NewPerfMetricJobSingleAllOfAttributes instantiates a new PerfMetricJobSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfMetricJobSingleAllOfAttributesWithDefaults

`func NewPerfMetricJobSingleAllOfAttributesWithDefaults() *PerfMetricJobSingleAllOfAttributes`

NewPerfMetricJobSingleAllOfAttributesWithDefaults instantiates a new PerfMetricJobSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativeState

`func (o *PerfMetricJobSingleAllOfAttributes) GetAdministrativeState() AdministrativeState`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *PerfMetricJobSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState)`

SetAdministrativeState sets AdministrativeState field to given value.

### HasAdministrativeState

`func (o *PerfMetricJobSingleAllOfAttributes) HasAdministrativeState() bool`

HasAdministrativeState returns a boolean if a field has been set.

### GetOperationalState

`func (o *PerfMetricJobSingleAllOfAttributes) GetOperationalState() OperationalState`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *PerfMetricJobSingleAllOfAttributes) SetOperationalState(v OperationalState)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *PerfMetricJobSingleAllOfAttributes) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetJobId

`func (o *PerfMetricJobSingleAllOfAttributes) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *PerfMetricJobSingleAllOfAttributes) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *PerfMetricJobSingleAllOfAttributes) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetPerformanceMetrics

`func (o *PerfMetricJobSingleAllOfAttributes) GetPerformanceMetrics() []string`

GetPerformanceMetrics returns the PerformanceMetrics field if non-nil, zero value otherwise.

### GetPerformanceMetricsOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetPerformanceMetricsOk() (*[]string, bool)`

GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceMetrics

`func (o *PerfMetricJobSingleAllOfAttributes) SetPerformanceMetrics(v []string)`

SetPerformanceMetrics sets PerformanceMetrics field to given value.

### HasPerformanceMetrics

`func (o *PerfMetricJobSingleAllOfAttributes) HasPerformanceMetrics() bool`

HasPerformanceMetrics returns a boolean if a field has been set.

### GetGranularityPeriod

`func (o *PerfMetricJobSingleAllOfAttributes) GetGranularityPeriod() int32`

GetGranularityPeriod returns the GranularityPeriod field if non-nil, zero value otherwise.

### GetGranularityPeriodOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetGranularityPeriodOk() (*int32, bool)`

GetGranularityPeriodOk returns a tuple with the GranularityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityPeriod

`func (o *PerfMetricJobSingleAllOfAttributes) SetGranularityPeriod(v int32)`

SetGranularityPeriod sets GranularityPeriod field to given value.

### HasGranularityPeriod

`func (o *PerfMetricJobSingleAllOfAttributes) HasGranularityPeriod() bool`

HasGranularityPeriod returns a boolean if a field has been set.

### GetObjectInstances

`func (o *PerfMetricJobSingleAllOfAttributes) GetObjectInstances() []string`

GetObjectInstances returns the ObjectInstances field if non-nil, zero value otherwise.

### GetObjectInstancesOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetObjectInstancesOk() (*[]string, bool)`

GetObjectInstancesOk returns a tuple with the ObjectInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstances

`func (o *PerfMetricJobSingleAllOfAttributes) SetObjectInstances(v []string)`

SetObjectInstances sets ObjectInstances field to given value.

### HasObjectInstances

`func (o *PerfMetricJobSingleAllOfAttributes) HasObjectInstances() bool`

HasObjectInstances returns a boolean if a field has been set.

### GetRootObjectInstances

`func (o *PerfMetricJobSingleAllOfAttributes) GetRootObjectInstances() []string`

GetRootObjectInstances returns the RootObjectInstances field if non-nil, zero value otherwise.

### GetRootObjectInstancesOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetRootObjectInstancesOk() (*[]string, bool)`

GetRootObjectInstancesOk returns a tuple with the RootObjectInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootObjectInstances

`func (o *PerfMetricJobSingleAllOfAttributes) SetRootObjectInstances(v []string)`

SetRootObjectInstances sets RootObjectInstances field to given value.

### HasRootObjectInstances

`func (o *PerfMetricJobSingleAllOfAttributes) HasRootObjectInstances() bool`

HasRootObjectInstances returns a boolean if a field has been set.

### GetReportingCtrl

`func (o *PerfMetricJobSingleAllOfAttributes) GetReportingCtrl() ReportingCtrl`

GetReportingCtrl returns the ReportingCtrl field if non-nil, zero value otherwise.

### GetReportingCtrlOk

`func (o *PerfMetricJobSingleAllOfAttributes) GetReportingCtrlOk() (*ReportingCtrl, bool)`

GetReportingCtrlOk returns a tuple with the ReportingCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCtrl

`func (o *PerfMetricJobSingleAllOfAttributes) SetReportingCtrl(v ReportingCtrl)`

SetReportingCtrl sets ReportingCtrl field to given value.

### HasReportingCtrl

`func (o *PerfMetricJobSingleAllOfAttributes) HasReportingCtrl() bool`

HasReportingCtrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


