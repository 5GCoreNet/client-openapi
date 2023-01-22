# NotifyChangedAlarmGeneralAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**AlarmType** | [**AlarmType**](AlarmType.md) |  | 
**ProbableCause** | Pointer to [**ProbableCause**](ProbableCause.md) |  | [optional] 
**SpecificProblem** | Pointer to [**SpecificProblem**](SpecificProblem.md) |  | [optional] 
**PerceivedSeverity** | Pointer to [**PerceivedSeverity**](PerceivedSeverity.md) |  | [optional] 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**BackedUpStatus** | Pointer to **bool** |  | [optional] 
**BackUpObject** | Pointer to **string** |  | [optional] 
**TrendIndication** | Pointer to [**TrendIndication**](TrendIndication.md) |  | [optional] 
**ThresholdInfo** | Pointer to [**ThresholdInfo**](ThresholdInfo.md) |  | [optional] 
**StateChangeDefinition** | Pointer to **[]map[string]interface{}** | The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values. | [optional] 
**MonitoredAttributes** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**ProposedRepairActions** | Pointer to **string** |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**RootCauseIndicator** | Pointer to **bool** |  | [optional] 
**ChangedAlarmAttributes** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 

## Methods

### NewNotifyChangedAlarmGeneralAllOf

`func NewNotifyChangedAlarmGeneralAllOf(alarmId string, alarmType AlarmType, ) *NotifyChangedAlarmGeneralAllOf`

NewNotifyChangedAlarmGeneralAllOf instantiates a new NotifyChangedAlarmGeneralAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyChangedAlarmGeneralAllOfWithDefaults

`func NewNotifyChangedAlarmGeneralAllOfWithDefaults() *NotifyChangedAlarmGeneralAllOf`

NewNotifyChangedAlarmGeneralAllOfWithDefaults instantiates a new NotifyChangedAlarmGeneralAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *NotifyChangedAlarmGeneralAllOf) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyChangedAlarmGeneralAllOf) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyChangedAlarmGeneralAllOf) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyChangedAlarmGeneralAllOf) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyChangedAlarmGeneralAllOf) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyChangedAlarmGeneralAllOf) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *NotifyChangedAlarmGeneralAllOf) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetSpecificProblem

`func (o *NotifyChangedAlarmGeneralAllOf) GetSpecificProblem() SpecificProblem`

GetSpecificProblem returns the SpecificProblem field if non-nil, zero value otherwise.

### GetSpecificProblemOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetSpecificProblemOk() (*SpecificProblem, bool)`

GetSpecificProblemOk returns a tuple with the SpecificProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificProblem

`func (o *NotifyChangedAlarmGeneralAllOf) SetSpecificProblem(v SpecificProblem)`

SetSpecificProblem sets SpecificProblem field to given value.

### HasSpecificProblem

`func (o *NotifyChangedAlarmGeneralAllOf) HasSpecificProblem() bool`

HasSpecificProblem returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *NotifyChangedAlarmGeneralAllOf) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyChangedAlarmGeneralAllOf) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.

### HasPerceivedSeverity

`func (o *NotifyChangedAlarmGeneralAllOf) HasPerceivedSeverity() bool`

HasPerceivedSeverity returns a boolean if a field has been set.

### GetCorrelatedNotifications

`func (o *NotifyChangedAlarmGeneralAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyChangedAlarmGeneralAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyChangedAlarmGeneralAllOf) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetBackedUpStatus

`func (o *NotifyChangedAlarmGeneralAllOf) GetBackedUpStatus() bool`

GetBackedUpStatus returns the BackedUpStatus field if non-nil, zero value otherwise.

### GetBackedUpStatusOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetBackedUpStatusOk() (*bool, bool)`

GetBackedUpStatusOk returns a tuple with the BackedUpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpStatus

`func (o *NotifyChangedAlarmGeneralAllOf) SetBackedUpStatus(v bool)`

SetBackedUpStatus sets BackedUpStatus field to given value.

### HasBackedUpStatus

`func (o *NotifyChangedAlarmGeneralAllOf) HasBackedUpStatus() bool`

HasBackedUpStatus returns a boolean if a field has been set.

### GetBackUpObject

`func (o *NotifyChangedAlarmGeneralAllOf) GetBackUpObject() string`

GetBackUpObject returns the BackUpObject field if non-nil, zero value otherwise.

### GetBackUpObjectOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetBackUpObjectOk() (*string, bool)`

GetBackUpObjectOk returns a tuple with the BackUpObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpObject

`func (o *NotifyChangedAlarmGeneralAllOf) SetBackUpObject(v string)`

SetBackUpObject sets BackUpObject field to given value.

### HasBackUpObject

`func (o *NotifyChangedAlarmGeneralAllOf) HasBackUpObject() bool`

HasBackUpObject returns a boolean if a field has been set.

### GetTrendIndication

`func (o *NotifyChangedAlarmGeneralAllOf) GetTrendIndication() TrendIndication`

GetTrendIndication returns the TrendIndication field if non-nil, zero value otherwise.

### GetTrendIndicationOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetTrendIndicationOk() (*TrendIndication, bool)`

GetTrendIndicationOk returns a tuple with the TrendIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendIndication

`func (o *NotifyChangedAlarmGeneralAllOf) SetTrendIndication(v TrendIndication)`

SetTrendIndication sets TrendIndication field to given value.

### HasTrendIndication

`func (o *NotifyChangedAlarmGeneralAllOf) HasTrendIndication() bool`

HasTrendIndication returns a boolean if a field has been set.

### GetThresholdInfo

`func (o *NotifyChangedAlarmGeneralAllOf) GetThresholdInfo() ThresholdInfo`

GetThresholdInfo returns the ThresholdInfo field if non-nil, zero value otherwise.

### GetThresholdInfoOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetThresholdInfoOk() (*ThresholdInfo, bool)`

GetThresholdInfoOk returns a tuple with the ThresholdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdInfo

`func (o *NotifyChangedAlarmGeneralAllOf) SetThresholdInfo(v ThresholdInfo)`

SetThresholdInfo sets ThresholdInfo field to given value.

### HasThresholdInfo

`func (o *NotifyChangedAlarmGeneralAllOf) HasThresholdInfo() bool`

HasThresholdInfo returns a boolean if a field has been set.

### GetStateChangeDefinition

`func (o *NotifyChangedAlarmGeneralAllOf) GetStateChangeDefinition() []map[string]interface{}`

GetStateChangeDefinition returns the StateChangeDefinition field if non-nil, zero value otherwise.

### GetStateChangeDefinitionOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetStateChangeDefinitionOk() (*[]map[string]interface{}, bool)`

GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeDefinition

`func (o *NotifyChangedAlarmGeneralAllOf) SetStateChangeDefinition(v []map[string]interface{})`

SetStateChangeDefinition sets StateChangeDefinition field to given value.

### HasStateChangeDefinition

`func (o *NotifyChangedAlarmGeneralAllOf) HasStateChangeDefinition() bool`

HasStateChangeDefinition returns a boolean if a field has been set.

### GetMonitoredAttributes

`func (o *NotifyChangedAlarmGeneralAllOf) GetMonitoredAttributes() map[string]interface{}`

GetMonitoredAttributes returns the MonitoredAttributes field if non-nil, zero value otherwise.

### GetMonitoredAttributesOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetMonitoredAttributesOk() (*map[string]interface{}, bool)`

GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAttributes

`func (o *NotifyChangedAlarmGeneralAllOf) SetMonitoredAttributes(v map[string]interface{})`

SetMonitoredAttributes sets MonitoredAttributes field to given value.

### HasMonitoredAttributes

`func (o *NotifyChangedAlarmGeneralAllOf) HasMonitoredAttributes() bool`

HasMonitoredAttributes returns a boolean if a field has been set.

### GetProposedRepairActions

`func (o *NotifyChangedAlarmGeneralAllOf) GetProposedRepairActions() string`

GetProposedRepairActions returns the ProposedRepairActions field if non-nil, zero value otherwise.

### GetProposedRepairActionsOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetProposedRepairActionsOk() (*string, bool)`

GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedRepairActions

`func (o *NotifyChangedAlarmGeneralAllOf) SetProposedRepairActions(v string)`

SetProposedRepairActions sets ProposedRepairActions field to given value.

### HasProposedRepairActions

`func (o *NotifyChangedAlarmGeneralAllOf) HasProposedRepairActions() bool`

HasProposedRepairActions returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyChangedAlarmGeneralAllOf) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyChangedAlarmGeneralAllOf) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyChangedAlarmGeneralAllOf) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *NotifyChangedAlarmGeneralAllOf) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *NotifyChangedAlarmGeneralAllOf) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *NotifyChangedAlarmGeneralAllOf) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *NotifyChangedAlarmGeneralAllOf) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *NotifyChangedAlarmGeneralAllOf) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *NotifyChangedAlarmGeneralAllOf) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.

### GetChangedAlarmAttributes

`func (o *NotifyChangedAlarmGeneralAllOf) GetChangedAlarmAttributes() map[string]interface{}`

GetChangedAlarmAttributes returns the ChangedAlarmAttributes field if non-nil, zero value otherwise.

### GetChangedAlarmAttributesOk

`func (o *NotifyChangedAlarmGeneralAllOf) GetChangedAlarmAttributesOk() (*map[string]interface{}, bool)`

GetChangedAlarmAttributesOk returns a tuple with the ChangedAlarmAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAlarmAttributes

`func (o *NotifyChangedAlarmGeneralAllOf) SetChangedAlarmAttributes(v map[string]interface{})`

SetChangedAlarmAttributes sets ChangedAlarmAttributes field to given value.

### HasChangedAlarmAttributes

`func (o *NotifyChangedAlarmGeneralAllOf) HasChangedAlarmAttributes() bool`

HasChangedAlarmAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


