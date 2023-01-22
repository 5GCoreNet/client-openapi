# AlarmsGet200ResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastNotificationHeader** | Pointer to [**NotificationHeader**](NotificationHeader.md) |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**NotificationId** | Pointer to **int32** |  | [optional] 
**AlarmRaisedTime** | Pointer to **time.Time** |  | [optional] 
**AlarmChangedTime** | Pointer to **time.Time** |  | [optional] 
**AlarmClearedTime** | Pointer to **time.Time** |  | [optional] 
**AlarmType** | Pointer to [**AlarmType**](AlarmType.md) |  | [optional] 
**ProbableCause** | Pointer to [**ProbableCause**](ProbableCause.md) |  | [optional] 
**SpecificProblem** | Pointer to [**SpecificProblem**](SpecificProblem.md) |  | [optional] 
**PerceivedSeverity** | Pointer to [**PerceivedSeverity**](PerceivedSeverity.md) |  | [optional] 
**BackedUpStatus** | Pointer to **bool** |  | [optional] 
**BackUpObject** | Pointer to **string** |  | [optional] 
**TrendIndication** | Pointer to [**TrendIndication**](TrendIndication.md) |  | [optional] 
**Thresholdinfo** | Pointer to [**ThresholdInfo**](ThresholdInfo.md) |  | [optional] 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**StateChangeDefinition** | Pointer to **[]map[string]interface{}** | The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values. | [optional] 
**MonitoredAttributes** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**ProposedRepairActions** | Pointer to **string** |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**RootCauseIndicator** | Pointer to **bool** |  | [optional] 
**AckTime** | Pointer to **time.Time** |  | [optional] 
**AckUserId** | Pointer to **string** |  | [optional] 
**AckSystemId** | Pointer to **string** |  | [optional] 
**AckState** | Pointer to [**AckState**](AckState.md) |  | [optional] 
**ClearUserId** | Pointer to **string** |  | [optional] 
**ClearSystemId** | Pointer to **string** |  | [optional] 
**ServiceUser** | Pointer to **string** |  | [optional] 
**ServiceProvider** | Pointer to **string** |  | [optional] 
**SecurityAlarmDetector** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to [**map[string]Comment**](Comment.md) | Collection of comments. The comment identifiers are allocated by the MnS producer and used as key in the map. | [optional] 

## Methods

### NewAlarmsGet200ResponseValue

`func NewAlarmsGet200ResponseValue() *AlarmsGet200ResponseValue`

NewAlarmsGet200ResponseValue instantiates a new AlarmsGet200ResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmsGet200ResponseValueWithDefaults

`func NewAlarmsGet200ResponseValueWithDefaults() *AlarmsGet200ResponseValue`

NewAlarmsGet200ResponseValueWithDefaults instantiates a new AlarmsGet200ResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastNotificationHeader

`func (o *AlarmsGet200ResponseValue) GetLastNotificationHeader() NotificationHeader`

GetLastNotificationHeader returns the LastNotificationHeader field if non-nil, zero value otherwise.

### GetLastNotificationHeaderOk

`func (o *AlarmsGet200ResponseValue) GetLastNotificationHeaderOk() (*NotificationHeader, bool)`

GetLastNotificationHeaderOk returns a tuple with the LastNotificationHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotificationHeader

`func (o *AlarmsGet200ResponseValue) SetLastNotificationHeader(v NotificationHeader)`

SetLastNotificationHeader sets LastNotificationHeader field to given value.

### HasLastNotificationHeader

`func (o *AlarmsGet200ResponseValue) HasLastNotificationHeader() bool`

HasLastNotificationHeader returns a boolean if a field has been set.

### GetObjectInstance

`func (o *AlarmsGet200ResponseValue) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *AlarmsGet200ResponseValue) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *AlarmsGet200ResponseValue) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *AlarmsGet200ResponseValue) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetNotificationId

`func (o *AlarmsGet200ResponseValue) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *AlarmsGet200ResponseValue) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *AlarmsGet200ResponseValue) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *AlarmsGet200ResponseValue) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetAlarmRaisedTime

`func (o *AlarmsGet200ResponseValue) GetAlarmRaisedTime() time.Time`

GetAlarmRaisedTime returns the AlarmRaisedTime field if non-nil, zero value otherwise.

### GetAlarmRaisedTimeOk

`func (o *AlarmsGet200ResponseValue) GetAlarmRaisedTimeOk() (*time.Time, bool)`

GetAlarmRaisedTimeOk returns a tuple with the AlarmRaisedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmRaisedTime

`func (o *AlarmsGet200ResponseValue) SetAlarmRaisedTime(v time.Time)`

SetAlarmRaisedTime sets AlarmRaisedTime field to given value.

### HasAlarmRaisedTime

`func (o *AlarmsGet200ResponseValue) HasAlarmRaisedTime() bool`

HasAlarmRaisedTime returns a boolean if a field has been set.

### GetAlarmChangedTime

`func (o *AlarmsGet200ResponseValue) GetAlarmChangedTime() time.Time`

GetAlarmChangedTime returns the AlarmChangedTime field if non-nil, zero value otherwise.

### GetAlarmChangedTimeOk

`func (o *AlarmsGet200ResponseValue) GetAlarmChangedTimeOk() (*time.Time, bool)`

GetAlarmChangedTimeOk returns a tuple with the AlarmChangedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmChangedTime

`func (o *AlarmsGet200ResponseValue) SetAlarmChangedTime(v time.Time)`

SetAlarmChangedTime sets AlarmChangedTime field to given value.

### HasAlarmChangedTime

`func (o *AlarmsGet200ResponseValue) HasAlarmChangedTime() bool`

HasAlarmChangedTime returns a boolean if a field has been set.

### GetAlarmClearedTime

`func (o *AlarmsGet200ResponseValue) GetAlarmClearedTime() time.Time`

GetAlarmClearedTime returns the AlarmClearedTime field if non-nil, zero value otherwise.

### GetAlarmClearedTimeOk

`func (o *AlarmsGet200ResponseValue) GetAlarmClearedTimeOk() (*time.Time, bool)`

GetAlarmClearedTimeOk returns a tuple with the AlarmClearedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmClearedTime

`func (o *AlarmsGet200ResponseValue) SetAlarmClearedTime(v time.Time)`

SetAlarmClearedTime sets AlarmClearedTime field to given value.

### HasAlarmClearedTime

`func (o *AlarmsGet200ResponseValue) HasAlarmClearedTime() bool`

HasAlarmClearedTime returns a boolean if a field has been set.

### GetAlarmType

`func (o *AlarmsGet200ResponseValue) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *AlarmsGet200ResponseValue) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *AlarmsGet200ResponseValue) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.

### HasAlarmType

`func (o *AlarmsGet200ResponseValue) HasAlarmType() bool`

HasAlarmType returns a boolean if a field has been set.

### GetProbableCause

`func (o *AlarmsGet200ResponseValue) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *AlarmsGet200ResponseValue) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *AlarmsGet200ResponseValue) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *AlarmsGet200ResponseValue) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetSpecificProblem

`func (o *AlarmsGet200ResponseValue) GetSpecificProblem() SpecificProblem`

GetSpecificProblem returns the SpecificProblem field if non-nil, zero value otherwise.

### GetSpecificProblemOk

`func (o *AlarmsGet200ResponseValue) GetSpecificProblemOk() (*SpecificProblem, bool)`

GetSpecificProblemOk returns a tuple with the SpecificProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificProblem

`func (o *AlarmsGet200ResponseValue) SetSpecificProblem(v SpecificProblem)`

SetSpecificProblem sets SpecificProblem field to given value.

### HasSpecificProblem

`func (o *AlarmsGet200ResponseValue) HasSpecificProblem() bool`

HasSpecificProblem returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *AlarmsGet200ResponseValue) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *AlarmsGet200ResponseValue) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *AlarmsGet200ResponseValue) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.

### HasPerceivedSeverity

`func (o *AlarmsGet200ResponseValue) HasPerceivedSeverity() bool`

HasPerceivedSeverity returns a boolean if a field has been set.

### GetBackedUpStatus

`func (o *AlarmsGet200ResponseValue) GetBackedUpStatus() bool`

GetBackedUpStatus returns the BackedUpStatus field if non-nil, zero value otherwise.

### GetBackedUpStatusOk

`func (o *AlarmsGet200ResponseValue) GetBackedUpStatusOk() (*bool, bool)`

GetBackedUpStatusOk returns a tuple with the BackedUpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpStatus

`func (o *AlarmsGet200ResponseValue) SetBackedUpStatus(v bool)`

SetBackedUpStatus sets BackedUpStatus field to given value.

### HasBackedUpStatus

`func (o *AlarmsGet200ResponseValue) HasBackedUpStatus() bool`

HasBackedUpStatus returns a boolean if a field has been set.

### GetBackUpObject

`func (o *AlarmsGet200ResponseValue) GetBackUpObject() string`

GetBackUpObject returns the BackUpObject field if non-nil, zero value otherwise.

### GetBackUpObjectOk

`func (o *AlarmsGet200ResponseValue) GetBackUpObjectOk() (*string, bool)`

GetBackUpObjectOk returns a tuple with the BackUpObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpObject

`func (o *AlarmsGet200ResponseValue) SetBackUpObject(v string)`

SetBackUpObject sets BackUpObject field to given value.

### HasBackUpObject

`func (o *AlarmsGet200ResponseValue) HasBackUpObject() bool`

HasBackUpObject returns a boolean if a field has been set.

### GetTrendIndication

`func (o *AlarmsGet200ResponseValue) GetTrendIndication() TrendIndication`

GetTrendIndication returns the TrendIndication field if non-nil, zero value otherwise.

### GetTrendIndicationOk

`func (o *AlarmsGet200ResponseValue) GetTrendIndicationOk() (*TrendIndication, bool)`

GetTrendIndicationOk returns a tuple with the TrendIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendIndication

`func (o *AlarmsGet200ResponseValue) SetTrendIndication(v TrendIndication)`

SetTrendIndication sets TrendIndication field to given value.

### HasTrendIndication

`func (o *AlarmsGet200ResponseValue) HasTrendIndication() bool`

HasTrendIndication returns a boolean if a field has been set.

### GetThresholdinfo

`func (o *AlarmsGet200ResponseValue) GetThresholdinfo() ThresholdInfo`

GetThresholdinfo returns the Thresholdinfo field if non-nil, zero value otherwise.

### GetThresholdinfoOk

`func (o *AlarmsGet200ResponseValue) GetThresholdinfoOk() (*ThresholdInfo, bool)`

GetThresholdinfoOk returns a tuple with the Thresholdinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdinfo

`func (o *AlarmsGet200ResponseValue) SetThresholdinfo(v ThresholdInfo)`

SetThresholdinfo sets Thresholdinfo field to given value.

### HasThresholdinfo

`func (o *AlarmsGet200ResponseValue) HasThresholdinfo() bool`

HasThresholdinfo returns a boolean if a field has been set.

### GetCorrelatedNotifications

`func (o *AlarmsGet200ResponseValue) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *AlarmsGet200ResponseValue) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *AlarmsGet200ResponseValue) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *AlarmsGet200ResponseValue) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetStateChangeDefinition

`func (o *AlarmsGet200ResponseValue) GetStateChangeDefinition() []map[string]interface{}`

GetStateChangeDefinition returns the StateChangeDefinition field if non-nil, zero value otherwise.

### GetStateChangeDefinitionOk

`func (o *AlarmsGet200ResponseValue) GetStateChangeDefinitionOk() (*[]map[string]interface{}, bool)`

GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeDefinition

`func (o *AlarmsGet200ResponseValue) SetStateChangeDefinition(v []map[string]interface{})`

SetStateChangeDefinition sets StateChangeDefinition field to given value.

### HasStateChangeDefinition

`func (o *AlarmsGet200ResponseValue) HasStateChangeDefinition() bool`

HasStateChangeDefinition returns a boolean if a field has been set.

### GetMonitoredAttributes

`func (o *AlarmsGet200ResponseValue) GetMonitoredAttributes() map[string]interface{}`

GetMonitoredAttributes returns the MonitoredAttributes field if non-nil, zero value otherwise.

### GetMonitoredAttributesOk

`func (o *AlarmsGet200ResponseValue) GetMonitoredAttributesOk() (*map[string]interface{}, bool)`

GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAttributes

`func (o *AlarmsGet200ResponseValue) SetMonitoredAttributes(v map[string]interface{})`

SetMonitoredAttributes sets MonitoredAttributes field to given value.

### HasMonitoredAttributes

`func (o *AlarmsGet200ResponseValue) HasMonitoredAttributes() bool`

HasMonitoredAttributes returns a boolean if a field has been set.

### GetProposedRepairActions

`func (o *AlarmsGet200ResponseValue) GetProposedRepairActions() string`

GetProposedRepairActions returns the ProposedRepairActions field if non-nil, zero value otherwise.

### GetProposedRepairActionsOk

`func (o *AlarmsGet200ResponseValue) GetProposedRepairActionsOk() (*string, bool)`

GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedRepairActions

`func (o *AlarmsGet200ResponseValue) SetProposedRepairActions(v string)`

SetProposedRepairActions sets ProposedRepairActions field to given value.

### HasProposedRepairActions

`func (o *AlarmsGet200ResponseValue) HasProposedRepairActions() bool`

HasProposedRepairActions returns a boolean if a field has been set.

### GetAdditionalText

`func (o *AlarmsGet200ResponseValue) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *AlarmsGet200ResponseValue) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *AlarmsGet200ResponseValue) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *AlarmsGet200ResponseValue) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *AlarmsGet200ResponseValue) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *AlarmsGet200ResponseValue) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *AlarmsGet200ResponseValue) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *AlarmsGet200ResponseValue) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *AlarmsGet200ResponseValue) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *AlarmsGet200ResponseValue) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *AlarmsGet200ResponseValue) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *AlarmsGet200ResponseValue) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.

### GetAckTime

`func (o *AlarmsGet200ResponseValue) GetAckTime() time.Time`

GetAckTime returns the AckTime field if non-nil, zero value otherwise.

### GetAckTimeOk

`func (o *AlarmsGet200ResponseValue) GetAckTimeOk() (*time.Time, bool)`

GetAckTimeOk returns a tuple with the AckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckTime

`func (o *AlarmsGet200ResponseValue) SetAckTime(v time.Time)`

SetAckTime sets AckTime field to given value.

### HasAckTime

`func (o *AlarmsGet200ResponseValue) HasAckTime() bool`

HasAckTime returns a boolean if a field has been set.

### GetAckUserId

`func (o *AlarmsGet200ResponseValue) GetAckUserId() string`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *AlarmsGet200ResponseValue) GetAckUserIdOk() (*string, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *AlarmsGet200ResponseValue) SetAckUserId(v string)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *AlarmsGet200ResponseValue) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetAckSystemId

`func (o *AlarmsGet200ResponseValue) GetAckSystemId() string`

GetAckSystemId returns the AckSystemId field if non-nil, zero value otherwise.

### GetAckSystemIdOk

`func (o *AlarmsGet200ResponseValue) GetAckSystemIdOk() (*string, bool)`

GetAckSystemIdOk returns a tuple with the AckSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSystemId

`func (o *AlarmsGet200ResponseValue) SetAckSystemId(v string)`

SetAckSystemId sets AckSystemId field to given value.

### HasAckSystemId

`func (o *AlarmsGet200ResponseValue) HasAckSystemId() bool`

HasAckSystemId returns a boolean if a field has been set.

### GetAckState

`func (o *AlarmsGet200ResponseValue) GetAckState() AckState`

GetAckState returns the AckState field if non-nil, zero value otherwise.

### GetAckStateOk

`func (o *AlarmsGet200ResponseValue) GetAckStateOk() (*AckState, bool)`

GetAckStateOk returns a tuple with the AckState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckState

`func (o *AlarmsGet200ResponseValue) SetAckState(v AckState)`

SetAckState sets AckState field to given value.

### HasAckState

`func (o *AlarmsGet200ResponseValue) HasAckState() bool`

HasAckState returns a boolean if a field has been set.

### GetClearUserId

`func (o *AlarmsGet200ResponseValue) GetClearUserId() string`

GetClearUserId returns the ClearUserId field if non-nil, zero value otherwise.

### GetClearUserIdOk

`func (o *AlarmsGet200ResponseValue) GetClearUserIdOk() (*string, bool)`

GetClearUserIdOk returns a tuple with the ClearUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearUserId

`func (o *AlarmsGet200ResponseValue) SetClearUserId(v string)`

SetClearUserId sets ClearUserId field to given value.

### HasClearUserId

`func (o *AlarmsGet200ResponseValue) HasClearUserId() bool`

HasClearUserId returns a boolean if a field has been set.

### GetClearSystemId

`func (o *AlarmsGet200ResponseValue) GetClearSystemId() string`

GetClearSystemId returns the ClearSystemId field if non-nil, zero value otherwise.

### GetClearSystemIdOk

`func (o *AlarmsGet200ResponseValue) GetClearSystemIdOk() (*string, bool)`

GetClearSystemIdOk returns a tuple with the ClearSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSystemId

`func (o *AlarmsGet200ResponseValue) SetClearSystemId(v string)`

SetClearSystemId sets ClearSystemId field to given value.

### HasClearSystemId

`func (o *AlarmsGet200ResponseValue) HasClearSystemId() bool`

HasClearSystemId returns a boolean if a field has been set.

### GetServiceUser

`func (o *AlarmsGet200ResponseValue) GetServiceUser() string`

GetServiceUser returns the ServiceUser field if non-nil, zero value otherwise.

### GetServiceUserOk

`func (o *AlarmsGet200ResponseValue) GetServiceUserOk() (*string, bool)`

GetServiceUserOk returns a tuple with the ServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUser

`func (o *AlarmsGet200ResponseValue) SetServiceUser(v string)`

SetServiceUser sets ServiceUser field to given value.

### HasServiceUser

`func (o *AlarmsGet200ResponseValue) HasServiceUser() bool`

HasServiceUser returns a boolean if a field has been set.

### GetServiceProvider

`func (o *AlarmsGet200ResponseValue) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *AlarmsGet200ResponseValue) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *AlarmsGet200ResponseValue) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *AlarmsGet200ResponseValue) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.

### GetSecurityAlarmDetector

`func (o *AlarmsGet200ResponseValue) GetSecurityAlarmDetector() string`

GetSecurityAlarmDetector returns the SecurityAlarmDetector field if non-nil, zero value otherwise.

### GetSecurityAlarmDetectorOk

`func (o *AlarmsGet200ResponseValue) GetSecurityAlarmDetectorOk() (*string, bool)`

GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAlarmDetector

`func (o *AlarmsGet200ResponseValue) SetSecurityAlarmDetector(v string)`

SetSecurityAlarmDetector sets SecurityAlarmDetector field to given value.

### HasSecurityAlarmDetector

`func (o *AlarmsGet200ResponseValue) HasSecurityAlarmDetector() bool`

HasSecurityAlarmDetector returns a boolean if a field has been set.

### GetComments

`func (o *AlarmsGet200ResponseValue) GetComments() map[string]Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AlarmsGet200ResponseValue) GetCommentsOk() (*map[string]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AlarmsGet200ResponseValue) SetComments(v map[string]Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AlarmsGet200ResponseValue) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


