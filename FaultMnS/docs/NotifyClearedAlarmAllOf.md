# NotifyClearedAlarmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**AlarmType** | [**AlarmType**](AlarmType.md) |  | 
**ProbableCause** | [**ProbableCause**](ProbableCause.md) |  | 
**PerceivedSeverity** | [**PerceivedSeverity**](PerceivedSeverity.md) |  | 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**ClearUserId** | Pointer to **string** |  | [optional] 
**ClearSystemId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyClearedAlarmAllOf

`func NewNotifyClearedAlarmAllOf(alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, ) *NotifyClearedAlarmAllOf`

NewNotifyClearedAlarmAllOf instantiates a new NotifyClearedAlarmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyClearedAlarmAllOfWithDefaults

`func NewNotifyClearedAlarmAllOfWithDefaults() *NotifyClearedAlarmAllOf`

NewNotifyClearedAlarmAllOfWithDefaults instantiates a new NotifyClearedAlarmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *NotifyClearedAlarmAllOf) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyClearedAlarmAllOf) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyClearedAlarmAllOf) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyClearedAlarmAllOf) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyClearedAlarmAllOf) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyClearedAlarmAllOf) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyClearedAlarmAllOf) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyClearedAlarmAllOf) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyClearedAlarmAllOf) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.


### GetPerceivedSeverity

`func (o *NotifyClearedAlarmAllOf) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyClearedAlarmAllOf) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyClearedAlarmAllOf) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.


### GetCorrelatedNotifications

`func (o *NotifyClearedAlarmAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyClearedAlarmAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyClearedAlarmAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyClearedAlarmAllOf) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetClearUserId

`func (o *NotifyClearedAlarmAllOf) GetClearUserId() string`

GetClearUserId returns the ClearUserId field if non-nil, zero value otherwise.

### GetClearUserIdOk

`func (o *NotifyClearedAlarmAllOf) GetClearUserIdOk() (*string, bool)`

GetClearUserIdOk returns a tuple with the ClearUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearUserId

`func (o *NotifyClearedAlarmAllOf) SetClearUserId(v string)`

SetClearUserId sets ClearUserId field to given value.

### HasClearUserId

`func (o *NotifyClearedAlarmAllOf) HasClearUserId() bool`

HasClearUserId returns a boolean if a field has been set.

### GetClearSystemId

`func (o *NotifyClearedAlarmAllOf) GetClearSystemId() string`

GetClearSystemId returns the ClearSystemId field if non-nil, zero value otherwise.

### GetClearSystemIdOk

`func (o *NotifyClearedAlarmAllOf) GetClearSystemIdOk() (*string, bool)`

GetClearSystemIdOk returns a tuple with the ClearSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSystemId

`func (o *NotifyClearedAlarmAllOf) SetClearSystemId(v string)`

SetClearSystemId sets ClearSystemId field to given value.

### HasClearSystemId

`func (o *NotifyClearedAlarmAllOf) HasClearSystemId() bool`

HasClearSystemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


