# ThresholdMonitorSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativeState** | Pointer to [**AdministrativeState**](AdministrativeState.md) |  | [optional] 
**OperationalState** | Pointer to [**OperationalState**](OperationalState.md) |  | [optional] 
**PerformanceMetrics** | Pointer to **[]string** |  | [optional] 
**ThresholdInfoList** | Pointer to [**[]ThresholdInfo**](ThresholdInfo.md) |  | [optional] 
**MonitorGranularityPeriod** | Pointer to **int32** |  | [optional] 
**ObjectInstances** | Pointer to **[]string** |  | [optional] 
**RootObjectInstances** | Pointer to **[]string** |  | [optional] 

## Methods

### NewThresholdMonitorSingleAllOfAttributes

`func NewThresholdMonitorSingleAllOfAttributes() *ThresholdMonitorSingleAllOfAttributes`

NewThresholdMonitorSingleAllOfAttributes instantiates a new ThresholdMonitorSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdMonitorSingleAllOfAttributesWithDefaults

`func NewThresholdMonitorSingleAllOfAttributesWithDefaults() *ThresholdMonitorSingleAllOfAttributes`

NewThresholdMonitorSingleAllOfAttributesWithDefaults instantiates a new ThresholdMonitorSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativeState

`func (o *ThresholdMonitorSingleAllOfAttributes) GetAdministrativeState() AdministrativeState`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *ThresholdMonitorSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *ThresholdMonitorSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState)`

SetAdministrativeState sets AdministrativeState field to given value.

### HasAdministrativeState

`func (o *ThresholdMonitorSingleAllOfAttributes) HasAdministrativeState() bool`

HasAdministrativeState returns a boolean if a field has been set.

### GetOperationalState

`func (o *ThresholdMonitorSingleAllOfAttributes) GetOperationalState() OperationalState`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *ThresholdMonitorSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *ThresholdMonitorSingleAllOfAttributes) SetOperationalState(v OperationalState)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *ThresholdMonitorSingleAllOfAttributes) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetPerformanceMetrics

`func (o *ThresholdMonitorSingleAllOfAttributes) GetPerformanceMetrics() []string`

GetPerformanceMetrics returns the PerformanceMetrics field if non-nil, zero value otherwise.

### GetPerformanceMetricsOk

`func (o *ThresholdMonitorSingleAllOfAttributes) GetPerformanceMetricsOk() (*[]string, bool)`

GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceMetrics

`func (o *ThresholdMonitorSingleAllOfAttributes) SetPerformanceMetrics(v []string)`

SetPerformanceMetrics sets PerformanceMetrics field to given value.

### HasPerformanceMetrics

`func (o *ThresholdMonitorSingleAllOfAttributes) HasPerformanceMetrics() bool`

HasPerformanceMetrics returns a boolean if a field has been set.

### GetThresholdInfoList

`func (o *ThresholdMonitorSingleAllOfAttributes) GetThresholdInfoList() []ThresholdInfo`

GetThresholdInfoList returns the ThresholdInfoList field if non-nil, zero value otherwise.

### GetThresholdInfoListOk

`func (o *ThresholdMonitorSingleAllOfAttributes) GetThresholdInfoListOk() (*[]ThresholdInfo, bool)`

GetThresholdInfoListOk returns a tuple with the ThresholdInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdInfoList

`func (o *ThresholdMonitorSingleAllOfAttributes) SetThresholdInfoList(v []ThresholdInfo)`

SetThresholdInfoList sets ThresholdInfoList field to given value.

### HasThresholdInfoList

`func (o *ThresholdMonitorSingleAllOfAttributes) HasThresholdInfoList() bool`

HasThresholdInfoList returns a boolean if a field has been set.

### GetMonitorGranularityPeriod

`func (o *ThresholdMonitorSingleAllOfAttributes) GetMonitorGranularityPeriod() int32`

GetMonitorGranularityPeriod returns the MonitorGranularityPeriod field if non-nil, zero value otherwise.

### GetMonitorGranularityPeriodOk

`func (o *ThresholdMonitorSingleAllOfAttributes) GetMonitorGranularityPeriodOk() (*int32, bool)`

GetMonitorGranularityPeriodOk returns a tuple with the MonitorGranularityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGranularityPeriod

`func (o *ThresholdMonitorSingleAllOfAttributes) SetMonitorGranularityPeriod(v int32)`

SetMonitorGranularityPeriod sets MonitorGranularityPeriod field to given value.

### HasMonitorGranularityPeriod

`func (o *ThresholdMonitorSingleAllOfAttributes) HasMonitorGranularityPeriod() bool`

HasMonitorGranularityPeriod returns a boolean if a field has been set.

### GetObjectInstances

`func (o *ThresholdMonitorSingleAllOfAttributes) GetObjectInstances() []string`

GetObjectInstances returns the ObjectInstances field if non-nil, zero value otherwise.

### GetObjectInstancesOk

`func (o *ThresholdMonitorSingleAllOfAttributes) GetObjectInstancesOk() (*[]string, bool)`

GetObjectInstancesOk returns a tuple with the ObjectInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstances

`func (o *ThresholdMonitorSingleAllOfAttributes) SetObjectInstances(v []string)`

SetObjectInstances sets ObjectInstances field to given value.

### HasObjectInstances

`func (o *ThresholdMonitorSingleAllOfAttributes) HasObjectInstances() bool`

HasObjectInstances returns a boolean if a field has been set.

### GetRootObjectInstances

`func (o *ThresholdMonitorSingleAllOfAttributes) GetRootObjectInstances() []string`

GetRootObjectInstances returns the RootObjectInstances field if non-nil, zero value otherwise.

### GetRootObjectInstancesOk

`func (o *ThresholdMonitorSingleAllOfAttributes) GetRootObjectInstancesOk() (*[]string, bool)`

GetRootObjectInstancesOk returns a tuple with the RootObjectInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootObjectInstances

`func (o *ThresholdMonitorSingleAllOfAttributes) SetRootObjectInstances(v []string)`

SetRootObjectInstances sets RootObjectInstances field to given value.

### HasRootObjectInstances

`func (o *ThresholdMonitorSingleAllOfAttributes) HasRootObjectInstances() bool`

HasRootObjectInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


