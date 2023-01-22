# FailedAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**FailureReason** | **string** |  | 

## Methods

### NewFailedAlarm

`func NewFailedAlarm(alarmId string, failureReason string, ) *FailedAlarm`

NewFailedAlarm instantiates a new FailedAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedAlarmWithDefaults

`func NewFailedAlarmWithDefaults() *FailedAlarm`

NewFailedAlarmWithDefaults instantiates a new FailedAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *FailedAlarm) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *FailedAlarm) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *FailedAlarm) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetFailureReason

`func (o *FailedAlarm) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *FailedAlarm) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *FailedAlarm) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


