# ReportingRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportingMode** | [**NotificationMethod**](NotificationMethod.md) |  | 
**ReportingPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**ReportingThrs** | Pointer to [**[]ReportingThreshold**](ReportingThreshold.md) |  | [optional] 
**ImmRep** | Pointer to **bool** |  | [optional] 
**RepTerminMode** | Pointer to [**TerminationMode**](TerminationMode.md) |  | [optional] 
**ExpirationTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**MaxNumRep** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TermThr** | Pointer to [**MeasurementData**](MeasurementData.md) |  | [optional] 
**TermThrMode** | Pointer to [**ThresholdHandlingMode**](ThresholdHandlingMode.md) |  | [optional] 

## Methods

### NewReportingRequirements

`func NewReportingRequirements(reportingMode NotificationMethod, ) *ReportingRequirements`

NewReportingRequirements instantiates a new ReportingRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingRequirementsWithDefaults

`func NewReportingRequirementsWithDefaults() *ReportingRequirements`

NewReportingRequirementsWithDefaults instantiates a new ReportingRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportingMode

`func (o *ReportingRequirements) GetReportingMode() NotificationMethod`

GetReportingMode returns the ReportingMode field if non-nil, zero value otherwise.

### GetReportingModeOk

`func (o *ReportingRequirements) GetReportingModeOk() (*NotificationMethod, bool)`

GetReportingModeOk returns a tuple with the ReportingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingMode

`func (o *ReportingRequirements) SetReportingMode(v NotificationMethod)`

SetReportingMode sets ReportingMode field to given value.


### GetReportingPeriod

`func (o *ReportingRequirements) GetReportingPeriod() int32`

GetReportingPeriod returns the ReportingPeriod field if non-nil, zero value otherwise.

### GetReportingPeriodOk

`func (o *ReportingRequirements) GetReportingPeriodOk() (*int32, bool)`

GetReportingPeriodOk returns a tuple with the ReportingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingPeriod

`func (o *ReportingRequirements) SetReportingPeriod(v int32)`

SetReportingPeriod sets ReportingPeriod field to given value.

### HasReportingPeriod

`func (o *ReportingRequirements) HasReportingPeriod() bool`

HasReportingPeriod returns a boolean if a field has been set.

### GetReportingThrs

`func (o *ReportingRequirements) GetReportingThrs() []ReportingThreshold`

GetReportingThrs returns the ReportingThrs field if non-nil, zero value otherwise.

### GetReportingThrsOk

`func (o *ReportingRequirements) GetReportingThrsOk() (*[]ReportingThreshold, bool)`

GetReportingThrsOk returns a tuple with the ReportingThrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingThrs

`func (o *ReportingRequirements) SetReportingThrs(v []ReportingThreshold)`

SetReportingThrs sets ReportingThrs field to given value.

### HasReportingThrs

`func (o *ReportingRequirements) HasReportingThrs() bool`

HasReportingThrs returns a boolean if a field has been set.

### GetImmRep

`func (o *ReportingRequirements) GetImmRep() bool`

GetImmRep returns the ImmRep field if non-nil, zero value otherwise.

### GetImmRepOk

`func (o *ReportingRequirements) GetImmRepOk() (*bool, bool)`

GetImmRepOk returns a tuple with the ImmRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmRep

`func (o *ReportingRequirements) SetImmRep(v bool)`

SetImmRep sets ImmRep field to given value.

### HasImmRep

`func (o *ReportingRequirements) HasImmRep() bool`

HasImmRep returns a boolean if a field has been set.

### GetRepTerminMode

`func (o *ReportingRequirements) GetRepTerminMode() TerminationMode`

GetRepTerminMode returns the RepTerminMode field if non-nil, zero value otherwise.

### GetRepTerminModeOk

`func (o *ReportingRequirements) GetRepTerminModeOk() (*TerminationMode, bool)`

GetRepTerminModeOk returns a tuple with the RepTerminMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepTerminMode

`func (o *ReportingRequirements) SetRepTerminMode(v TerminationMode)`

SetRepTerminMode sets RepTerminMode field to given value.

### HasRepTerminMode

`func (o *ReportingRequirements) HasRepTerminMode() bool`

HasRepTerminMode returns a boolean if a field has been set.

### GetExpirationTimer

`func (o *ReportingRequirements) GetExpirationTimer() int32`

GetExpirationTimer returns the ExpirationTimer field if non-nil, zero value otherwise.

### GetExpirationTimerOk

`func (o *ReportingRequirements) GetExpirationTimerOk() (*int32, bool)`

GetExpirationTimerOk returns a tuple with the ExpirationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimer

`func (o *ReportingRequirements) SetExpirationTimer(v int32)`

SetExpirationTimer sets ExpirationTimer field to given value.

### HasExpirationTimer

`func (o *ReportingRequirements) HasExpirationTimer() bool`

HasExpirationTimer returns a boolean if a field has been set.

### GetMaxNumRep

`func (o *ReportingRequirements) GetMaxNumRep() int32`

GetMaxNumRep returns the MaxNumRep field if non-nil, zero value otherwise.

### GetMaxNumRepOk

`func (o *ReportingRequirements) GetMaxNumRepOk() (*int32, bool)`

GetMaxNumRepOk returns a tuple with the MaxNumRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumRep

`func (o *ReportingRequirements) SetMaxNumRep(v int32)`

SetMaxNumRep sets MaxNumRep field to given value.

### HasMaxNumRep

`func (o *ReportingRequirements) HasMaxNumRep() bool`

HasMaxNumRep returns a boolean if a field has been set.

### GetTermThr

`func (o *ReportingRequirements) GetTermThr() MeasurementData`

GetTermThr returns the TermThr field if non-nil, zero value otherwise.

### GetTermThrOk

`func (o *ReportingRequirements) GetTermThrOk() (*MeasurementData, bool)`

GetTermThrOk returns a tuple with the TermThr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermThr

`func (o *ReportingRequirements) SetTermThr(v MeasurementData)`

SetTermThr sets TermThr field to given value.

### HasTermThr

`func (o *ReportingRequirements) HasTermThr() bool`

HasTermThr returns a boolean if a field has been set.

### GetTermThrMode

`func (o *ReportingRequirements) GetTermThrMode() ThresholdHandlingMode`

GetTermThrMode returns the TermThrMode field if non-nil, zero value otherwise.

### GetTermThrModeOk

`func (o *ReportingRequirements) GetTermThrModeOk() (*ThresholdHandlingMode, bool)`

GetTermThrModeOk returns a tuple with the TermThrMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermThrMode

`func (o *ReportingRequirements) SetTermThrMode(v ThresholdHandlingMode)`

SetTermThrMode sets TermThrMode field to given value.

### HasTermThrMode

`func (o *ReportingRequirements) HasTermThrMode() bool`

HasTermThrMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


