# ReportingCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ReportingConditionType**](ReportingConditionType.md) |  | 
**Period** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**Parameter** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to [**ReportingConditionThreshold**](ReportingConditionThreshold.md) |  | [optional] 
**ReportWhenBelow** | Pointer to **bool** |  | [optional] 
**EventTrigger** | Pointer to [**ReportingEventTrigger**](ReportingEventTrigger.md) |  | [optional] 

## Methods

### NewReportingCondition

`func NewReportingCondition(type_ ReportingConditionType, ) *ReportingCondition`

NewReportingCondition instantiates a new ReportingCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingConditionWithDefaults

`func NewReportingConditionWithDefaults() *ReportingCondition`

NewReportingConditionWithDefaults instantiates a new ReportingCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportingCondition) GetType() ReportingConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportingCondition) GetTypeOk() (*ReportingConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportingCondition) SetType(v ReportingConditionType)`

SetType sets Type field to given value.


### GetPeriod

`func (o *ReportingCondition) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ReportingCondition) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ReportingCondition) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ReportingCondition) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetParameter

`func (o *ReportingCondition) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ReportingCondition) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ReportingCondition) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ReportingCondition) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetThreshold

`func (o *ReportingCondition) GetThreshold() ReportingConditionThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ReportingCondition) GetThresholdOk() (*ReportingConditionThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ReportingCondition) SetThreshold(v ReportingConditionThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ReportingCondition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetReportWhenBelow

`func (o *ReportingCondition) GetReportWhenBelow() bool`

GetReportWhenBelow returns the ReportWhenBelow field if non-nil, zero value otherwise.

### GetReportWhenBelowOk

`func (o *ReportingCondition) GetReportWhenBelowOk() (*bool, bool)`

GetReportWhenBelowOk returns a tuple with the ReportWhenBelow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportWhenBelow

`func (o *ReportingCondition) SetReportWhenBelow(v bool)`

SetReportWhenBelow sets ReportWhenBelow field to given value.

### HasReportWhenBelow

`func (o *ReportingCondition) HasReportWhenBelow() bool`

HasReportWhenBelow returns a boolean if a field has been set.

### GetEventTrigger

`func (o *ReportingCondition) GetEventTrigger() ReportingEventTrigger`

GetEventTrigger returns the EventTrigger field if non-nil, zero value otherwise.

### GetEventTriggerOk

`func (o *ReportingCondition) GetEventTriggerOk() (*ReportingEventTrigger, bool)`

GetEventTriggerOk returns a tuple with the EventTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTrigger

`func (o *ReportingCondition) SetEventTrigger(v ReportingEventTrigger)`

SetEventTrigger sets EventTrigger field to given value.

### HasEventTrigger

`func (o *ReportingCondition) HasEventTrigger() bool`

HasEventTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


