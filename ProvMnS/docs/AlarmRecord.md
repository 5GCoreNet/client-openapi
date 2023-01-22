# AlarmRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Thresholdinfo** | Pointer to [**ThresholdInfo1**](ThresholdInfo1.md) |  | [optional] 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification1**](CorrelatedNotification1.md) |  | [optional] 
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

## Methods

### NewAlarmRecord

`func NewAlarmRecord() *AlarmRecord`

NewAlarmRecord instantiates a new AlarmRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmRecordWithDefaults

`func NewAlarmRecordWithDefaults() *AlarmRecord`

NewAlarmRecordWithDefaults instantiates a new AlarmRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectInstance

`func (o *AlarmRecord) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *AlarmRecord) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *AlarmRecord) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *AlarmRecord) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetNotificationId

`func (o *AlarmRecord) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *AlarmRecord) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *AlarmRecord) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *AlarmRecord) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetAlarmRaisedTime

`func (o *AlarmRecord) GetAlarmRaisedTime() time.Time`

GetAlarmRaisedTime returns the AlarmRaisedTime field if non-nil, zero value otherwise.

### GetAlarmRaisedTimeOk

`func (o *AlarmRecord) GetAlarmRaisedTimeOk() (*time.Time, bool)`

GetAlarmRaisedTimeOk returns a tuple with the AlarmRaisedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmRaisedTime

`func (o *AlarmRecord) SetAlarmRaisedTime(v time.Time)`

SetAlarmRaisedTime sets AlarmRaisedTime field to given value.

### HasAlarmRaisedTime

`func (o *AlarmRecord) HasAlarmRaisedTime() bool`

HasAlarmRaisedTime returns a boolean if a field has been set.

### GetAlarmChangedTime

`func (o *AlarmRecord) GetAlarmChangedTime() time.Time`

GetAlarmChangedTime returns the AlarmChangedTime field if non-nil, zero value otherwise.

### GetAlarmChangedTimeOk

`func (o *AlarmRecord) GetAlarmChangedTimeOk() (*time.Time, bool)`

GetAlarmChangedTimeOk returns a tuple with the AlarmChangedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmChangedTime

`func (o *AlarmRecord) SetAlarmChangedTime(v time.Time)`

SetAlarmChangedTime sets AlarmChangedTime field to given value.

### HasAlarmChangedTime

`func (o *AlarmRecord) HasAlarmChangedTime() bool`

HasAlarmChangedTime returns a boolean if a field has been set.

### GetAlarmClearedTime

`func (o *AlarmRecord) GetAlarmClearedTime() time.Time`

GetAlarmClearedTime returns the AlarmClearedTime field if non-nil, zero value otherwise.

### GetAlarmClearedTimeOk

`func (o *AlarmRecord) GetAlarmClearedTimeOk() (*time.Time, bool)`

GetAlarmClearedTimeOk returns a tuple with the AlarmClearedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmClearedTime

`func (o *AlarmRecord) SetAlarmClearedTime(v time.Time)`

SetAlarmClearedTime sets AlarmClearedTime field to given value.

### HasAlarmClearedTime

`func (o *AlarmRecord) HasAlarmClearedTime() bool`

HasAlarmClearedTime returns a boolean if a field has been set.

### GetAlarmType

`func (o *AlarmRecord) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *AlarmRecord) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *AlarmRecord) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.

### HasAlarmType

`func (o *AlarmRecord) HasAlarmType() bool`

HasAlarmType returns a boolean if a field has been set.

### GetProbableCause

`func (o *AlarmRecord) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *AlarmRecord) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *AlarmRecord) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *AlarmRecord) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetSpecificProblem

`func (o *AlarmRecord) GetSpecificProblem() SpecificProblem`

GetSpecificProblem returns the SpecificProblem field if non-nil, zero value otherwise.

### GetSpecificProblemOk

`func (o *AlarmRecord) GetSpecificProblemOk() (*SpecificProblem, bool)`

GetSpecificProblemOk returns a tuple with the SpecificProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificProblem

`func (o *AlarmRecord) SetSpecificProblem(v SpecificProblem)`

SetSpecificProblem sets SpecificProblem field to given value.

### HasSpecificProblem

`func (o *AlarmRecord) HasSpecificProblem() bool`

HasSpecificProblem returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *AlarmRecord) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *AlarmRecord) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *AlarmRecord) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.

### HasPerceivedSeverity

`func (o *AlarmRecord) HasPerceivedSeverity() bool`

HasPerceivedSeverity returns a boolean if a field has been set.

### GetBackedUpStatus

`func (o *AlarmRecord) GetBackedUpStatus() bool`

GetBackedUpStatus returns the BackedUpStatus field if non-nil, zero value otherwise.

### GetBackedUpStatusOk

`func (o *AlarmRecord) GetBackedUpStatusOk() (*bool, bool)`

GetBackedUpStatusOk returns a tuple with the BackedUpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpStatus

`func (o *AlarmRecord) SetBackedUpStatus(v bool)`

SetBackedUpStatus sets BackedUpStatus field to given value.

### HasBackedUpStatus

`func (o *AlarmRecord) HasBackedUpStatus() bool`

HasBackedUpStatus returns a boolean if a field has been set.

### GetBackUpObject

`func (o *AlarmRecord) GetBackUpObject() string`

GetBackUpObject returns the BackUpObject field if non-nil, zero value otherwise.

### GetBackUpObjectOk

`func (o *AlarmRecord) GetBackUpObjectOk() (*string, bool)`

GetBackUpObjectOk returns a tuple with the BackUpObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpObject

`func (o *AlarmRecord) SetBackUpObject(v string)`

SetBackUpObject sets BackUpObject field to given value.

### HasBackUpObject

`func (o *AlarmRecord) HasBackUpObject() bool`

HasBackUpObject returns a boolean if a field has been set.

### GetTrendIndication

`func (o *AlarmRecord) GetTrendIndication() TrendIndication`

GetTrendIndication returns the TrendIndication field if non-nil, zero value otherwise.

### GetTrendIndicationOk

`func (o *AlarmRecord) GetTrendIndicationOk() (*TrendIndication, bool)`

GetTrendIndicationOk returns a tuple with the TrendIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendIndication

`func (o *AlarmRecord) SetTrendIndication(v TrendIndication)`

SetTrendIndication sets TrendIndication field to given value.

### HasTrendIndication

`func (o *AlarmRecord) HasTrendIndication() bool`

HasTrendIndication returns a boolean if a field has been set.

### GetThresholdinfo

`func (o *AlarmRecord) GetThresholdinfo() ThresholdInfo1`

GetThresholdinfo returns the Thresholdinfo field if non-nil, zero value otherwise.

### GetThresholdinfoOk

`func (o *AlarmRecord) GetThresholdinfoOk() (*ThresholdInfo1, bool)`

GetThresholdinfoOk returns a tuple with the Thresholdinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdinfo

`func (o *AlarmRecord) SetThresholdinfo(v ThresholdInfo1)`

SetThresholdinfo sets Thresholdinfo field to given value.

### HasThresholdinfo

`func (o *AlarmRecord) HasThresholdinfo() bool`

HasThresholdinfo returns a boolean if a field has been set.

### GetCorrelatedNotifications

`func (o *AlarmRecord) GetCorrelatedNotifications() []CorrelatedNotification1`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *AlarmRecord) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification1, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *AlarmRecord) SetCorrelatedNotifications(v []CorrelatedNotification1)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *AlarmRecord) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetStateChangeDefinition

`func (o *AlarmRecord) GetStateChangeDefinition() []map[string]interface{}`

GetStateChangeDefinition returns the StateChangeDefinition field if non-nil, zero value otherwise.

### GetStateChangeDefinitionOk

`func (o *AlarmRecord) GetStateChangeDefinitionOk() (*[]map[string]interface{}, bool)`

GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeDefinition

`func (o *AlarmRecord) SetStateChangeDefinition(v []map[string]interface{})`

SetStateChangeDefinition sets StateChangeDefinition field to given value.

### HasStateChangeDefinition

`func (o *AlarmRecord) HasStateChangeDefinition() bool`

HasStateChangeDefinition returns a boolean if a field has been set.

### GetMonitoredAttributes

`func (o *AlarmRecord) GetMonitoredAttributes() map[string]interface{}`

GetMonitoredAttributes returns the MonitoredAttributes field if non-nil, zero value otherwise.

### GetMonitoredAttributesOk

`func (o *AlarmRecord) GetMonitoredAttributesOk() (*map[string]interface{}, bool)`

GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAttributes

`func (o *AlarmRecord) SetMonitoredAttributes(v map[string]interface{})`

SetMonitoredAttributes sets MonitoredAttributes field to given value.

### HasMonitoredAttributes

`func (o *AlarmRecord) HasMonitoredAttributes() bool`

HasMonitoredAttributes returns a boolean if a field has been set.

### GetProposedRepairActions

`func (o *AlarmRecord) GetProposedRepairActions() string`

GetProposedRepairActions returns the ProposedRepairActions field if non-nil, zero value otherwise.

### GetProposedRepairActionsOk

`func (o *AlarmRecord) GetProposedRepairActionsOk() (*string, bool)`

GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedRepairActions

`func (o *AlarmRecord) SetProposedRepairActions(v string)`

SetProposedRepairActions sets ProposedRepairActions field to given value.

### HasProposedRepairActions

`func (o *AlarmRecord) HasProposedRepairActions() bool`

HasProposedRepairActions returns a boolean if a field has been set.

### GetAdditionalText

`func (o *AlarmRecord) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *AlarmRecord) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *AlarmRecord) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *AlarmRecord) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *AlarmRecord) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *AlarmRecord) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *AlarmRecord) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *AlarmRecord) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *AlarmRecord) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *AlarmRecord) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *AlarmRecord) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *AlarmRecord) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.

### GetAckTime

`func (o *AlarmRecord) GetAckTime() time.Time`

GetAckTime returns the AckTime field if non-nil, zero value otherwise.

### GetAckTimeOk

`func (o *AlarmRecord) GetAckTimeOk() (*time.Time, bool)`

GetAckTimeOk returns a tuple with the AckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckTime

`func (o *AlarmRecord) SetAckTime(v time.Time)`

SetAckTime sets AckTime field to given value.

### HasAckTime

`func (o *AlarmRecord) HasAckTime() bool`

HasAckTime returns a boolean if a field has been set.

### GetAckUserId

`func (o *AlarmRecord) GetAckUserId() string`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *AlarmRecord) GetAckUserIdOk() (*string, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *AlarmRecord) SetAckUserId(v string)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *AlarmRecord) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetAckSystemId

`func (o *AlarmRecord) GetAckSystemId() string`

GetAckSystemId returns the AckSystemId field if non-nil, zero value otherwise.

### GetAckSystemIdOk

`func (o *AlarmRecord) GetAckSystemIdOk() (*string, bool)`

GetAckSystemIdOk returns a tuple with the AckSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSystemId

`func (o *AlarmRecord) SetAckSystemId(v string)`

SetAckSystemId sets AckSystemId field to given value.

### HasAckSystemId

`func (o *AlarmRecord) HasAckSystemId() bool`

HasAckSystemId returns a boolean if a field has been set.

### GetAckState

`func (o *AlarmRecord) GetAckState() AckState`

GetAckState returns the AckState field if non-nil, zero value otherwise.

### GetAckStateOk

`func (o *AlarmRecord) GetAckStateOk() (*AckState, bool)`

GetAckStateOk returns a tuple with the AckState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckState

`func (o *AlarmRecord) SetAckState(v AckState)`

SetAckState sets AckState field to given value.

### HasAckState

`func (o *AlarmRecord) HasAckState() bool`

HasAckState returns a boolean if a field has been set.

### GetClearUserId

`func (o *AlarmRecord) GetClearUserId() string`

GetClearUserId returns the ClearUserId field if non-nil, zero value otherwise.

### GetClearUserIdOk

`func (o *AlarmRecord) GetClearUserIdOk() (*string, bool)`

GetClearUserIdOk returns a tuple with the ClearUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearUserId

`func (o *AlarmRecord) SetClearUserId(v string)`

SetClearUserId sets ClearUserId field to given value.

### HasClearUserId

`func (o *AlarmRecord) HasClearUserId() bool`

HasClearUserId returns a boolean if a field has been set.

### GetClearSystemId

`func (o *AlarmRecord) GetClearSystemId() string`

GetClearSystemId returns the ClearSystemId field if non-nil, zero value otherwise.

### GetClearSystemIdOk

`func (o *AlarmRecord) GetClearSystemIdOk() (*string, bool)`

GetClearSystemIdOk returns a tuple with the ClearSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSystemId

`func (o *AlarmRecord) SetClearSystemId(v string)`

SetClearSystemId sets ClearSystemId field to given value.

### HasClearSystemId

`func (o *AlarmRecord) HasClearSystemId() bool`

HasClearSystemId returns a boolean if a field has been set.

### GetServiceUser

`func (o *AlarmRecord) GetServiceUser() string`

GetServiceUser returns the ServiceUser field if non-nil, zero value otherwise.

### GetServiceUserOk

`func (o *AlarmRecord) GetServiceUserOk() (*string, bool)`

GetServiceUserOk returns a tuple with the ServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUser

`func (o *AlarmRecord) SetServiceUser(v string)`

SetServiceUser sets ServiceUser field to given value.

### HasServiceUser

`func (o *AlarmRecord) HasServiceUser() bool`

HasServiceUser returns a boolean if a field has been set.

### GetServiceProvider

`func (o *AlarmRecord) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *AlarmRecord) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *AlarmRecord) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *AlarmRecord) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.

### GetSecurityAlarmDetector

`func (o *AlarmRecord) GetSecurityAlarmDetector() string`

GetSecurityAlarmDetector returns the SecurityAlarmDetector field if non-nil, zero value otherwise.

### GetSecurityAlarmDetectorOk

`func (o *AlarmRecord) GetSecurityAlarmDetectorOk() (*string, bool)`

GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAlarmDetector

`func (o *AlarmRecord) SetSecurityAlarmDetector(v string)`

SetSecurityAlarmDetector sets SecurityAlarmDetector field to given value.

### HasSecurityAlarmDetector

`func (o *AlarmRecord) HasSecurityAlarmDetector() bool`

HasSecurityAlarmDetector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


