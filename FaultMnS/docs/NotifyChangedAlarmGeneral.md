# NotifyChangedAlarmGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
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

### NewNotifyChangedAlarmGeneral

`func NewNotifyChangedAlarmGeneral(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, ) *NotifyChangedAlarmGeneral`

NewNotifyChangedAlarmGeneral instantiates a new NotifyChangedAlarmGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyChangedAlarmGeneralWithDefaults

`func NewNotifyChangedAlarmGeneralWithDefaults() *NotifyChangedAlarmGeneral`

NewNotifyChangedAlarmGeneralWithDefaults instantiates a new NotifyChangedAlarmGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyChangedAlarmGeneral) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyChangedAlarmGeneral) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyChangedAlarmGeneral) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyChangedAlarmGeneral) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyChangedAlarmGeneral) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyChangedAlarmGeneral) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyChangedAlarmGeneral) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyChangedAlarmGeneral) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyChangedAlarmGeneral) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyChangedAlarmGeneral) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyChangedAlarmGeneral) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyChangedAlarmGeneral) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyChangedAlarmGeneral) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyChangedAlarmGeneral) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyChangedAlarmGeneral) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetAlarmId

`func (o *NotifyChangedAlarmGeneral) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyChangedAlarmGeneral) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyChangedAlarmGeneral) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyChangedAlarmGeneral) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyChangedAlarmGeneral) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyChangedAlarmGeneral) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyChangedAlarmGeneral) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyChangedAlarmGeneral) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyChangedAlarmGeneral) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *NotifyChangedAlarmGeneral) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetSpecificProblem

`func (o *NotifyChangedAlarmGeneral) GetSpecificProblem() SpecificProblem`

GetSpecificProblem returns the SpecificProblem field if non-nil, zero value otherwise.

### GetSpecificProblemOk

`func (o *NotifyChangedAlarmGeneral) GetSpecificProblemOk() (*SpecificProblem, bool)`

GetSpecificProblemOk returns a tuple with the SpecificProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificProblem

`func (o *NotifyChangedAlarmGeneral) SetSpecificProblem(v SpecificProblem)`

SetSpecificProblem sets SpecificProblem field to given value.

### HasSpecificProblem

`func (o *NotifyChangedAlarmGeneral) HasSpecificProblem() bool`

HasSpecificProblem returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *NotifyChangedAlarmGeneral) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyChangedAlarmGeneral) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyChangedAlarmGeneral) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.

### HasPerceivedSeverity

`func (o *NotifyChangedAlarmGeneral) HasPerceivedSeverity() bool`

HasPerceivedSeverity returns a boolean if a field has been set.

### GetCorrelatedNotifications

`func (o *NotifyChangedAlarmGeneral) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyChangedAlarmGeneral) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyChangedAlarmGeneral) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyChangedAlarmGeneral) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetBackedUpStatus

`func (o *NotifyChangedAlarmGeneral) GetBackedUpStatus() bool`

GetBackedUpStatus returns the BackedUpStatus field if non-nil, zero value otherwise.

### GetBackedUpStatusOk

`func (o *NotifyChangedAlarmGeneral) GetBackedUpStatusOk() (*bool, bool)`

GetBackedUpStatusOk returns a tuple with the BackedUpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpStatus

`func (o *NotifyChangedAlarmGeneral) SetBackedUpStatus(v bool)`

SetBackedUpStatus sets BackedUpStatus field to given value.

### HasBackedUpStatus

`func (o *NotifyChangedAlarmGeneral) HasBackedUpStatus() bool`

HasBackedUpStatus returns a boolean if a field has been set.

### GetBackUpObject

`func (o *NotifyChangedAlarmGeneral) GetBackUpObject() string`

GetBackUpObject returns the BackUpObject field if non-nil, zero value otherwise.

### GetBackUpObjectOk

`func (o *NotifyChangedAlarmGeneral) GetBackUpObjectOk() (*string, bool)`

GetBackUpObjectOk returns a tuple with the BackUpObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpObject

`func (o *NotifyChangedAlarmGeneral) SetBackUpObject(v string)`

SetBackUpObject sets BackUpObject field to given value.

### HasBackUpObject

`func (o *NotifyChangedAlarmGeneral) HasBackUpObject() bool`

HasBackUpObject returns a boolean if a field has been set.

### GetTrendIndication

`func (o *NotifyChangedAlarmGeneral) GetTrendIndication() TrendIndication`

GetTrendIndication returns the TrendIndication field if non-nil, zero value otherwise.

### GetTrendIndicationOk

`func (o *NotifyChangedAlarmGeneral) GetTrendIndicationOk() (*TrendIndication, bool)`

GetTrendIndicationOk returns a tuple with the TrendIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendIndication

`func (o *NotifyChangedAlarmGeneral) SetTrendIndication(v TrendIndication)`

SetTrendIndication sets TrendIndication field to given value.

### HasTrendIndication

`func (o *NotifyChangedAlarmGeneral) HasTrendIndication() bool`

HasTrendIndication returns a boolean if a field has been set.

### GetThresholdInfo

`func (o *NotifyChangedAlarmGeneral) GetThresholdInfo() ThresholdInfo`

GetThresholdInfo returns the ThresholdInfo field if non-nil, zero value otherwise.

### GetThresholdInfoOk

`func (o *NotifyChangedAlarmGeneral) GetThresholdInfoOk() (*ThresholdInfo, bool)`

GetThresholdInfoOk returns a tuple with the ThresholdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdInfo

`func (o *NotifyChangedAlarmGeneral) SetThresholdInfo(v ThresholdInfo)`

SetThresholdInfo sets ThresholdInfo field to given value.

### HasThresholdInfo

`func (o *NotifyChangedAlarmGeneral) HasThresholdInfo() bool`

HasThresholdInfo returns a boolean if a field has been set.

### GetStateChangeDefinition

`func (o *NotifyChangedAlarmGeneral) GetStateChangeDefinition() []map[string]interface{}`

GetStateChangeDefinition returns the StateChangeDefinition field if non-nil, zero value otherwise.

### GetStateChangeDefinitionOk

`func (o *NotifyChangedAlarmGeneral) GetStateChangeDefinitionOk() (*[]map[string]interface{}, bool)`

GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeDefinition

`func (o *NotifyChangedAlarmGeneral) SetStateChangeDefinition(v []map[string]interface{})`

SetStateChangeDefinition sets StateChangeDefinition field to given value.

### HasStateChangeDefinition

`func (o *NotifyChangedAlarmGeneral) HasStateChangeDefinition() bool`

HasStateChangeDefinition returns a boolean if a field has been set.

### GetMonitoredAttributes

`func (o *NotifyChangedAlarmGeneral) GetMonitoredAttributes() map[string]interface{}`

GetMonitoredAttributes returns the MonitoredAttributes field if non-nil, zero value otherwise.

### GetMonitoredAttributesOk

`func (o *NotifyChangedAlarmGeneral) GetMonitoredAttributesOk() (*map[string]interface{}, bool)`

GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAttributes

`func (o *NotifyChangedAlarmGeneral) SetMonitoredAttributes(v map[string]interface{})`

SetMonitoredAttributes sets MonitoredAttributes field to given value.

### HasMonitoredAttributes

`func (o *NotifyChangedAlarmGeneral) HasMonitoredAttributes() bool`

HasMonitoredAttributes returns a boolean if a field has been set.

### GetProposedRepairActions

`func (o *NotifyChangedAlarmGeneral) GetProposedRepairActions() string`

GetProposedRepairActions returns the ProposedRepairActions field if non-nil, zero value otherwise.

### GetProposedRepairActionsOk

`func (o *NotifyChangedAlarmGeneral) GetProposedRepairActionsOk() (*string, bool)`

GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedRepairActions

`func (o *NotifyChangedAlarmGeneral) SetProposedRepairActions(v string)`

SetProposedRepairActions sets ProposedRepairActions field to given value.

### HasProposedRepairActions

`func (o *NotifyChangedAlarmGeneral) HasProposedRepairActions() bool`

HasProposedRepairActions returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyChangedAlarmGeneral) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyChangedAlarmGeneral) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyChangedAlarmGeneral) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyChangedAlarmGeneral) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *NotifyChangedAlarmGeneral) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *NotifyChangedAlarmGeneral) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *NotifyChangedAlarmGeneral) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *NotifyChangedAlarmGeneral) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *NotifyChangedAlarmGeneral) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *NotifyChangedAlarmGeneral) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *NotifyChangedAlarmGeneral) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *NotifyChangedAlarmGeneral) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.

### GetChangedAlarmAttributes

`func (o *NotifyChangedAlarmGeneral) GetChangedAlarmAttributes() map[string]interface{}`

GetChangedAlarmAttributes returns the ChangedAlarmAttributes field if non-nil, zero value otherwise.

### GetChangedAlarmAttributesOk

`func (o *NotifyChangedAlarmGeneral) GetChangedAlarmAttributesOk() (*map[string]interface{}, bool)`

GetChangedAlarmAttributesOk returns a tuple with the ChangedAlarmAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAlarmAttributes

`func (o *NotifyChangedAlarmGeneral) SetChangedAlarmAttributes(v map[string]interface{})`

SetChangedAlarmAttributes sets ChangedAlarmAttributes field to given value.

### HasChangedAlarmAttributes

`func (o *NotifyChangedAlarmGeneral) HasChangedAlarmAttributes() bool`

HasChangedAlarmAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


