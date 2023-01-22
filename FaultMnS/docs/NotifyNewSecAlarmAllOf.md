# NotifyNewSecAlarmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**AlarmType** | [**AlarmType**](AlarmType.md) |  | 
**ProbableCause** | [**ProbableCause**](ProbableCause.md) |  | 
**PerceivedSeverity** | [**PerceivedSeverity**](PerceivedSeverity.md) |  | 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**RootCauseIndicator** | Pointer to **bool** |  | [optional] 
**ServiceUser** | **string** |  | 
**ServiceProvider** | **string** |  | 
**SecurityAlarmDetector** | **string** |  | 

## Methods

### NewNotifyNewSecAlarmAllOf

`func NewNotifyNewSecAlarmAllOf(alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, serviceUser string, serviceProvider string, securityAlarmDetector string, ) *NotifyNewSecAlarmAllOf`

NewNotifyNewSecAlarmAllOf instantiates a new NotifyNewSecAlarmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyNewSecAlarmAllOfWithDefaults

`func NewNotifyNewSecAlarmAllOfWithDefaults() *NotifyNewSecAlarmAllOf`

NewNotifyNewSecAlarmAllOfWithDefaults instantiates a new NotifyNewSecAlarmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *NotifyNewSecAlarmAllOf) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyNewSecAlarmAllOf) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyNewSecAlarmAllOf) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyNewSecAlarmAllOf) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyNewSecAlarmAllOf) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyNewSecAlarmAllOf) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyNewSecAlarmAllOf) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyNewSecAlarmAllOf) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyNewSecAlarmAllOf) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.


### GetPerceivedSeverity

`func (o *NotifyNewSecAlarmAllOf) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyNewSecAlarmAllOf) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyNewSecAlarmAllOf) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.


### GetCorrelatedNotifications

`func (o *NotifyNewSecAlarmAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyNewSecAlarmAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyNewSecAlarmAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyNewSecAlarmAllOf) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyNewSecAlarmAllOf) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyNewSecAlarmAllOf) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyNewSecAlarmAllOf) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyNewSecAlarmAllOf) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *NotifyNewSecAlarmAllOf) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *NotifyNewSecAlarmAllOf) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *NotifyNewSecAlarmAllOf) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *NotifyNewSecAlarmAllOf) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *NotifyNewSecAlarmAllOf) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *NotifyNewSecAlarmAllOf) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *NotifyNewSecAlarmAllOf) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *NotifyNewSecAlarmAllOf) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.

### GetServiceUser

`func (o *NotifyNewSecAlarmAllOf) GetServiceUser() string`

GetServiceUser returns the ServiceUser field if non-nil, zero value otherwise.

### GetServiceUserOk

`func (o *NotifyNewSecAlarmAllOf) GetServiceUserOk() (*string, bool)`

GetServiceUserOk returns a tuple with the ServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUser

`func (o *NotifyNewSecAlarmAllOf) SetServiceUser(v string)`

SetServiceUser sets ServiceUser field to given value.


### GetServiceProvider

`func (o *NotifyNewSecAlarmAllOf) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *NotifyNewSecAlarmAllOf) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *NotifyNewSecAlarmAllOf) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.


### GetSecurityAlarmDetector

`func (o *NotifyNewSecAlarmAllOf) GetSecurityAlarmDetector() string`

GetSecurityAlarmDetector returns the SecurityAlarmDetector field if non-nil, zero value otherwise.

### GetSecurityAlarmDetectorOk

`func (o *NotifyNewSecAlarmAllOf) GetSecurityAlarmDetectorOk() (*string, bool)`

GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAlarmDetector

`func (o *NotifyNewSecAlarmAllOf) SetSecurityAlarmDetector(v string)`

SetSecurityAlarmDetector sets SecurityAlarmDetector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


