# SubscriptionsPostRequest1

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
**ServiceUser** | **string** |  | 
**ServiceProvider** | **string** |  | 
**SecurityAlarmDetector** | **string** |  | 

## Methods

### NewSubscriptionsPostRequest1

`func NewSubscriptionsPostRequest1(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, serviceUser string, serviceProvider string, securityAlarmDetector string, ) *SubscriptionsPostRequest1`

NewSubscriptionsPostRequest1 instantiates a new SubscriptionsPostRequest1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsPostRequest1WithDefaults

`func NewSubscriptionsPostRequest1WithDefaults() *SubscriptionsPostRequest1`

NewSubscriptionsPostRequest1WithDefaults instantiates a new SubscriptionsPostRequest1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SubscriptionsPostRequest1) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SubscriptionsPostRequest1) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SubscriptionsPostRequest1) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *SubscriptionsPostRequest1) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *SubscriptionsPostRequest1) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *SubscriptionsPostRequest1) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *SubscriptionsPostRequest1) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *SubscriptionsPostRequest1) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *SubscriptionsPostRequest1) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *SubscriptionsPostRequest1) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *SubscriptionsPostRequest1) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *SubscriptionsPostRequest1) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *SubscriptionsPostRequest1) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *SubscriptionsPostRequest1) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *SubscriptionsPostRequest1) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetAlarmId

`func (o *SubscriptionsPostRequest1) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *SubscriptionsPostRequest1) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *SubscriptionsPostRequest1) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *SubscriptionsPostRequest1) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *SubscriptionsPostRequest1) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *SubscriptionsPostRequest1) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *SubscriptionsPostRequest1) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *SubscriptionsPostRequest1) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *SubscriptionsPostRequest1) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *SubscriptionsPostRequest1) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetSpecificProblem

`func (o *SubscriptionsPostRequest1) GetSpecificProblem() SpecificProblem`

GetSpecificProblem returns the SpecificProblem field if non-nil, zero value otherwise.

### GetSpecificProblemOk

`func (o *SubscriptionsPostRequest1) GetSpecificProblemOk() (*SpecificProblem, bool)`

GetSpecificProblemOk returns a tuple with the SpecificProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificProblem

`func (o *SubscriptionsPostRequest1) SetSpecificProblem(v SpecificProblem)`

SetSpecificProblem sets SpecificProblem field to given value.

### HasSpecificProblem

`func (o *SubscriptionsPostRequest1) HasSpecificProblem() bool`

HasSpecificProblem returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *SubscriptionsPostRequest1) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *SubscriptionsPostRequest1) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *SubscriptionsPostRequest1) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.

### HasPerceivedSeverity

`func (o *SubscriptionsPostRequest1) HasPerceivedSeverity() bool`

HasPerceivedSeverity returns a boolean if a field has been set.

### GetCorrelatedNotifications

`func (o *SubscriptionsPostRequest1) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *SubscriptionsPostRequest1) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *SubscriptionsPostRequest1) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *SubscriptionsPostRequest1) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetBackedUpStatus

`func (o *SubscriptionsPostRequest1) GetBackedUpStatus() bool`

GetBackedUpStatus returns the BackedUpStatus field if non-nil, zero value otherwise.

### GetBackedUpStatusOk

`func (o *SubscriptionsPostRequest1) GetBackedUpStatusOk() (*bool, bool)`

GetBackedUpStatusOk returns a tuple with the BackedUpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpStatus

`func (o *SubscriptionsPostRequest1) SetBackedUpStatus(v bool)`

SetBackedUpStatus sets BackedUpStatus field to given value.

### HasBackedUpStatus

`func (o *SubscriptionsPostRequest1) HasBackedUpStatus() bool`

HasBackedUpStatus returns a boolean if a field has been set.

### GetBackUpObject

`func (o *SubscriptionsPostRequest1) GetBackUpObject() string`

GetBackUpObject returns the BackUpObject field if non-nil, zero value otherwise.

### GetBackUpObjectOk

`func (o *SubscriptionsPostRequest1) GetBackUpObjectOk() (*string, bool)`

GetBackUpObjectOk returns a tuple with the BackUpObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpObject

`func (o *SubscriptionsPostRequest1) SetBackUpObject(v string)`

SetBackUpObject sets BackUpObject field to given value.

### HasBackUpObject

`func (o *SubscriptionsPostRequest1) HasBackUpObject() bool`

HasBackUpObject returns a boolean if a field has been set.

### GetTrendIndication

`func (o *SubscriptionsPostRequest1) GetTrendIndication() TrendIndication`

GetTrendIndication returns the TrendIndication field if non-nil, zero value otherwise.

### GetTrendIndicationOk

`func (o *SubscriptionsPostRequest1) GetTrendIndicationOk() (*TrendIndication, bool)`

GetTrendIndicationOk returns a tuple with the TrendIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendIndication

`func (o *SubscriptionsPostRequest1) SetTrendIndication(v TrendIndication)`

SetTrendIndication sets TrendIndication field to given value.

### HasTrendIndication

`func (o *SubscriptionsPostRequest1) HasTrendIndication() bool`

HasTrendIndication returns a boolean if a field has been set.

### GetThresholdInfo

`func (o *SubscriptionsPostRequest1) GetThresholdInfo() ThresholdInfo`

GetThresholdInfo returns the ThresholdInfo field if non-nil, zero value otherwise.

### GetThresholdInfoOk

`func (o *SubscriptionsPostRequest1) GetThresholdInfoOk() (*ThresholdInfo, bool)`

GetThresholdInfoOk returns a tuple with the ThresholdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdInfo

`func (o *SubscriptionsPostRequest1) SetThresholdInfo(v ThresholdInfo)`

SetThresholdInfo sets ThresholdInfo field to given value.

### HasThresholdInfo

`func (o *SubscriptionsPostRequest1) HasThresholdInfo() bool`

HasThresholdInfo returns a boolean if a field has been set.

### GetStateChangeDefinition

`func (o *SubscriptionsPostRequest1) GetStateChangeDefinition() []map[string]interface{}`

GetStateChangeDefinition returns the StateChangeDefinition field if non-nil, zero value otherwise.

### GetStateChangeDefinitionOk

`func (o *SubscriptionsPostRequest1) GetStateChangeDefinitionOk() (*[]map[string]interface{}, bool)`

GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeDefinition

`func (o *SubscriptionsPostRequest1) SetStateChangeDefinition(v []map[string]interface{})`

SetStateChangeDefinition sets StateChangeDefinition field to given value.

### HasStateChangeDefinition

`func (o *SubscriptionsPostRequest1) HasStateChangeDefinition() bool`

HasStateChangeDefinition returns a boolean if a field has been set.

### GetMonitoredAttributes

`func (o *SubscriptionsPostRequest1) GetMonitoredAttributes() map[string]interface{}`

GetMonitoredAttributes returns the MonitoredAttributes field if non-nil, zero value otherwise.

### GetMonitoredAttributesOk

`func (o *SubscriptionsPostRequest1) GetMonitoredAttributesOk() (*map[string]interface{}, bool)`

GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAttributes

`func (o *SubscriptionsPostRequest1) SetMonitoredAttributes(v map[string]interface{})`

SetMonitoredAttributes sets MonitoredAttributes field to given value.

### HasMonitoredAttributes

`func (o *SubscriptionsPostRequest1) HasMonitoredAttributes() bool`

HasMonitoredAttributes returns a boolean if a field has been set.

### GetProposedRepairActions

`func (o *SubscriptionsPostRequest1) GetProposedRepairActions() string`

GetProposedRepairActions returns the ProposedRepairActions field if non-nil, zero value otherwise.

### GetProposedRepairActionsOk

`func (o *SubscriptionsPostRequest1) GetProposedRepairActionsOk() (*string, bool)`

GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedRepairActions

`func (o *SubscriptionsPostRequest1) SetProposedRepairActions(v string)`

SetProposedRepairActions sets ProposedRepairActions field to given value.

### HasProposedRepairActions

`func (o *SubscriptionsPostRequest1) HasProposedRepairActions() bool`

HasProposedRepairActions returns a boolean if a field has been set.

### GetAdditionalText

`func (o *SubscriptionsPostRequest1) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *SubscriptionsPostRequest1) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *SubscriptionsPostRequest1) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *SubscriptionsPostRequest1) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *SubscriptionsPostRequest1) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *SubscriptionsPostRequest1) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *SubscriptionsPostRequest1) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *SubscriptionsPostRequest1) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *SubscriptionsPostRequest1) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *SubscriptionsPostRequest1) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *SubscriptionsPostRequest1) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *SubscriptionsPostRequest1) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.

### GetChangedAlarmAttributes

`func (o *SubscriptionsPostRequest1) GetChangedAlarmAttributes() map[string]interface{}`

GetChangedAlarmAttributes returns the ChangedAlarmAttributes field if non-nil, zero value otherwise.

### GetChangedAlarmAttributesOk

`func (o *SubscriptionsPostRequest1) GetChangedAlarmAttributesOk() (*map[string]interface{}, bool)`

GetChangedAlarmAttributesOk returns a tuple with the ChangedAlarmAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAlarmAttributes

`func (o *SubscriptionsPostRequest1) SetChangedAlarmAttributes(v map[string]interface{})`

SetChangedAlarmAttributes sets ChangedAlarmAttributes field to given value.

### HasChangedAlarmAttributes

`func (o *SubscriptionsPostRequest1) HasChangedAlarmAttributes() bool`

HasChangedAlarmAttributes returns a boolean if a field has been set.

### GetServiceUser

`func (o *SubscriptionsPostRequest1) GetServiceUser() string`

GetServiceUser returns the ServiceUser field if non-nil, zero value otherwise.

### GetServiceUserOk

`func (o *SubscriptionsPostRequest1) GetServiceUserOk() (*string, bool)`

GetServiceUserOk returns a tuple with the ServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUser

`func (o *SubscriptionsPostRequest1) SetServiceUser(v string)`

SetServiceUser sets ServiceUser field to given value.


### GetServiceProvider

`func (o *SubscriptionsPostRequest1) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *SubscriptionsPostRequest1) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *SubscriptionsPostRequest1) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.


### GetSecurityAlarmDetector

`func (o *SubscriptionsPostRequest1) GetSecurityAlarmDetector() string`

GetSecurityAlarmDetector returns the SecurityAlarmDetector field if non-nil, zero value otherwise.

### GetSecurityAlarmDetectorOk

`func (o *SubscriptionsPostRequest1) GetSecurityAlarmDetectorOk() (*string, bool)`

GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAlarmDetector

`func (o *SubscriptionsPostRequest1) SetSecurityAlarmDetector(v string)`

SetSecurityAlarmDetector sets SecurityAlarmDetector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


