# NotifyNewAlarmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**AlarmType** | [**AlarmType**](AlarmType.md) |  | 
**ProbableCause** | [**ProbableCause**](ProbableCause.md) |  | 
**SpecificProblem** | Pointer to [**SpecificProblem**](SpecificProblem.md) |  | [optional] 
**PerceivedSeverity** | [**PerceivedSeverity**](PerceivedSeverity.md) |  | 
**BackedUpStatus** | Pointer to **bool** |  | [optional] 
**BackUpObject** | Pointer to **string** |  | [optional] 
**TrendIndication** | Pointer to [**TrendIndication**](TrendIndication.md) |  | [optional] 
**ThresholdInfo** | Pointer to [**ThresholdInfo**](ThresholdInfo.md) |  | [optional] 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**StateChangeDefinition** | Pointer to **[]map[string]interface{}** | The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values. | [optional] 
**MonitoredAttributes** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**ProposedRepairActions** | Pointer to **string** |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**RootCauseIndicator** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotifyNewAlarmAllOf

`func NewNotifyNewAlarmAllOf(alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, ) *NotifyNewAlarmAllOf`

NewNotifyNewAlarmAllOf instantiates a new NotifyNewAlarmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyNewAlarmAllOfWithDefaults

`func NewNotifyNewAlarmAllOfWithDefaults() *NotifyNewAlarmAllOf`

NewNotifyNewAlarmAllOfWithDefaults instantiates a new NotifyNewAlarmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *NotifyNewAlarmAllOf) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyNewAlarmAllOf) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyNewAlarmAllOf) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyNewAlarmAllOf) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyNewAlarmAllOf) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyNewAlarmAllOf) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyNewAlarmAllOf) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyNewAlarmAllOf) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyNewAlarmAllOf) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.


### GetSpecificProblem

`func (o *NotifyNewAlarmAllOf) GetSpecificProblem() SpecificProblem`

GetSpecificProblem returns the SpecificProblem field if non-nil, zero value otherwise.

### GetSpecificProblemOk

`func (o *NotifyNewAlarmAllOf) GetSpecificProblemOk() (*SpecificProblem, bool)`

GetSpecificProblemOk returns a tuple with the SpecificProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificProblem

`func (o *NotifyNewAlarmAllOf) SetSpecificProblem(v SpecificProblem)`

SetSpecificProblem sets SpecificProblem field to given value.

### HasSpecificProblem

`func (o *NotifyNewAlarmAllOf) HasSpecificProblem() bool`

HasSpecificProblem returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *NotifyNewAlarmAllOf) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyNewAlarmAllOf) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyNewAlarmAllOf) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.


### GetBackedUpStatus

`func (o *NotifyNewAlarmAllOf) GetBackedUpStatus() bool`

GetBackedUpStatus returns the BackedUpStatus field if non-nil, zero value otherwise.

### GetBackedUpStatusOk

`func (o *NotifyNewAlarmAllOf) GetBackedUpStatusOk() (*bool, bool)`

GetBackedUpStatusOk returns a tuple with the BackedUpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpStatus

`func (o *NotifyNewAlarmAllOf) SetBackedUpStatus(v bool)`

SetBackedUpStatus sets BackedUpStatus field to given value.

### HasBackedUpStatus

`func (o *NotifyNewAlarmAllOf) HasBackedUpStatus() bool`

HasBackedUpStatus returns a boolean if a field has been set.

### GetBackUpObject

`func (o *NotifyNewAlarmAllOf) GetBackUpObject() string`

GetBackUpObject returns the BackUpObject field if non-nil, zero value otherwise.

### GetBackUpObjectOk

`func (o *NotifyNewAlarmAllOf) GetBackUpObjectOk() (*string, bool)`

GetBackUpObjectOk returns a tuple with the BackUpObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpObject

`func (o *NotifyNewAlarmAllOf) SetBackUpObject(v string)`

SetBackUpObject sets BackUpObject field to given value.

### HasBackUpObject

`func (o *NotifyNewAlarmAllOf) HasBackUpObject() bool`

HasBackUpObject returns a boolean if a field has been set.

### GetTrendIndication

`func (o *NotifyNewAlarmAllOf) GetTrendIndication() TrendIndication`

GetTrendIndication returns the TrendIndication field if non-nil, zero value otherwise.

### GetTrendIndicationOk

`func (o *NotifyNewAlarmAllOf) GetTrendIndicationOk() (*TrendIndication, bool)`

GetTrendIndicationOk returns a tuple with the TrendIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendIndication

`func (o *NotifyNewAlarmAllOf) SetTrendIndication(v TrendIndication)`

SetTrendIndication sets TrendIndication field to given value.

### HasTrendIndication

`func (o *NotifyNewAlarmAllOf) HasTrendIndication() bool`

HasTrendIndication returns a boolean if a field has been set.

### GetThresholdInfo

`func (o *NotifyNewAlarmAllOf) GetThresholdInfo() ThresholdInfo`

GetThresholdInfo returns the ThresholdInfo field if non-nil, zero value otherwise.

### GetThresholdInfoOk

`func (o *NotifyNewAlarmAllOf) GetThresholdInfoOk() (*ThresholdInfo, bool)`

GetThresholdInfoOk returns a tuple with the ThresholdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdInfo

`func (o *NotifyNewAlarmAllOf) SetThresholdInfo(v ThresholdInfo)`

SetThresholdInfo sets ThresholdInfo field to given value.

### HasThresholdInfo

`func (o *NotifyNewAlarmAllOf) HasThresholdInfo() bool`

HasThresholdInfo returns a boolean if a field has been set.

### GetCorrelatedNotifications

`func (o *NotifyNewAlarmAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyNewAlarmAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyNewAlarmAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyNewAlarmAllOf) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetStateChangeDefinition

`func (o *NotifyNewAlarmAllOf) GetStateChangeDefinition() []map[string]interface{}`

GetStateChangeDefinition returns the StateChangeDefinition field if non-nil, zero value otherwise.

### GetStateChangeDefinitionOk

`func (o *NotifyNewAlarmAllOf) GetStateChangeDefinitionOk() (*[]map[string]interface{}, bool)`

GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeDefinition

`func (o *NotifyNewAlarmAllOf) SetStateChangeDefinition(v []map[string]interface{})`

SetStateChangeDefinition sets StateChangeDefinition field to given value.

### HasStateChangeDefinition

`func (o *NotifyNewAlarmAllOf) HasStateChangeDefinition() bool`

HasStateChangeDefinition returns a boolean if a field has been set.

### GetMonitoredAttributes

`func (o *NotifyNewAlarmAllOf) GetMonitoredAttributes() map[string]interface{}`

GetMonitoredAttributes returns the MonitoredAttributes field if non-nil, zero value otherwise.

### GetMonitoredAttributesOk

`func (o *NotifyNewAlarmAllOf) GetMonitoredAttributesOk() (*map[string]interface{}, bool)`

GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAttributes

`func (o *NotifyNewAlarmAllOf) SetMonitoredAttributes(v map[string]interface{})`

SetMonitoredAttributes sets MonitoredAttributes field to given value.

### HasMonitoredAttributes

`func (o *NotifyNewAlarmAllOf) HasMonitoredAttributes() bool`

HasMonitoredAttributes returns a boolean if a field has been set.

### GetProposedRepairActions

`func (o *NotifyNewAlarmAllOf) GetProposedRepairActions() string`

GetProposedRepairActions returns the ProposedRepairActions field if non-nil, zero value otherwise.

### GetProposedRepairActionsOk

`func (o *NotifyNewAlarmAllOf) GetProposedRepairActionsOk() (*string, bool)`

GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedRepairActions

`func (o *NotifyNewAlarmAllOf) SetProposedRepairActions(v string)`

SetProposedRepairActions sets ProposedRepairActions field to given value.

### HasProposedRepairActions

`func (o *NotifyNewAlarmAllOf) HasProposedRepairActions() bool`

HasProposedRepairActions returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyNewAlarmAllOf) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyNewAlarmAllOf) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyNewAlarmAllOf) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyNewAlarmAllOf) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *NotifyNewAlarmAllOf) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *NotifyNewAlarmAllOf) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *NotifyNewAlarmAllOf) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *NotifyNewAlarmAllOf) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *NotifyNewAlarmAllOf) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *NotifyNewAlarmAllOf) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *NotifyNewAlarmAllOf) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *NotifyNewAlarmAllOf) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


