# NotifyChangedSecAlarmGeneralAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**AlarmType** | [**AlarmType**](AlarmType.md) |  | 
**ProbableCause** | Pointer to [**ProbableCause**](ProbableCause.md) |  | [optional] 
**PerceivedSeverity** | Pointer to [**PerceivedSeverity**](PerceivedSeverity.md) |  | [optional] 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 
**RootCauseIndicator** | Pointer to **bool** |  | [optional] 
**ServiceUser** | **string** |  | 
**ServiceProvider** | **string** |  | 
**SecurityAlarmDetector** | **string** |  | 
**ChangedAlarmAttributes** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 

## Methods

### NewNotifyChangedSecAlarmGeneralAllOf

`func NewNotifyChangedSecAlarmGeneralAllOf(alarmId string, alarmType AlarmType, serviceUser string, serviceProvider string, securityAlarmDetector string, ) *NotifyChangedSecAlarmGeneralAllOf`

NewNotifyChangedSecAlarmGeneralAllOf instantiates a new NotifyChangedSecAlarmGeneralAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyChangedSecAlarmGeneralAllOfWithDefaults

`func NewNotifyChangedSecAlarmGeneralAllOfWithDefaults() *NotifyChangedSecAlarmGeneralAllOf`

NewNotifyChangedSecAlarmGeneralAllOfWithDefaults instantiates a new NotifyChangedSecAlarmGeneralAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *NotifyChangedSecAlarmGeneralAllOf) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.

### HasPerceivedSeverity

`func (o *NotifyChangedSecAlarmGeneralAllOf) HasPerceivedSeverity() bool`

HasPerceivedSeverity returns a boolean if a field has been set.

### GetCorrelatedNotifications

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyChangedSecAlarmGeneralAllOf) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyChangedSecAlarmGeneralAllOf) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *NotifyChangedSecAlarmGeneralAllOf) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRootCauseIndicator

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *NotifyChangedSecAlarmGeneralAllOf) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.

### GetServiceUser

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetServiceUser() string`

GetServiceUser returns the ServiceUser field if non-nil, zero value otherwise.

### GetServiceUserOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetServiceUserOk() (*string, bool)`

GetServiceUserOk returns a tuple with the ServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUser

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetServiceUser(v string)`

SetServiceUser sets ServiceUser field to given value.


### GetServiceProvider

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.


### GetSecurityAlarmDetector

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetSecurityAlarmDetector() string`

GetSecurityAlarmDetector returns the SecurityAlarmDetector field if non-nil, zero value otherwise.

### GetSecurityAlarmDetectorOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetSecurityAlarmDetectorOk() (*string, bool)`

GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAlarmDetector

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetSecurityAlarmDetector(v string)`

SetSecurityAlarmDetector sets SecurityAlarmDetector field to given value.


### GetChangedAlarmAttributes

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetChangedAlarmAttributes() map[string]interface{}`

GetChangedAlarmAttributes returns the ChangedAlarmAttributes field if non-nil, zero value otherwise.

### GetChangedAlarmAttributesOk

`func (o *NotifyChangedSecAlarmGeneralAllOf) GetChangedAlarmAttributesOk() (*map[string]interface{}, bool)`

GetChangedAlarmAttributesOk returns a tuple with the ChangedAlarmAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAlarmAttributes

`func (o *NotifyChangedSecAlarmGeneralAllOf) SetChangedAlarmAttributes(v map[string]interface{})`

SetChangedAlarmAttributes sets ChangedAlarmAttributes field to given value.

### HasChangedAlarmAttributes

`func (o *NotifyChangedSecAlarmGeneralAllOf) HasChangedAlarmAttributes() bool`

HasChangedAlarmAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


