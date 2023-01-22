# NotifyAckStateChangedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**AlarmType** | [**AlarmType**](AlarmType.md) |  | 
**ProbableCause** | [**ProbableCause**](ProbableCause.md) |  | 
**PerceivedSeverity** | [**PerceivedSeverity**](PerceivedSeverity.md) |  | 
**AckState** | [**AckState**](AckState.md) |  | 
**AckUserId** | **string** |  | 
**AckSystemId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyAckStateChangedAllOf

`func NewNotifyAckStateChangedAllOf(alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, ackState AckState, ackUserId string, ) *NotifyAckStateChangedAllOf`

NewNotifyAckStateChangedAllOf instantiates a new NotifyAckStateChangedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyAckStateChangedAllOfWithDefaults

`func NewNotifyAckStateChangedAllOfWithDefaults() *NotifyAckStateChangedAllOf`

NewNotifyAckStateChangedAllOfWithDefaults instantiates a new NotifyAckStateChangedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *NotifyAckStateChangedAllOf) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyAckStateChangedAllOf) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyAckStateChangedAllOf) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyAckStateChangedAllOf) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyAckStateChangedAllOf) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyAckStateChangedAllOf) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyAckStateChangedAllOf) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyAckStateChangedAllOf) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyAckStateChangedAllOf) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.


### GetPerceivedSeverity

`func (o *NotifyAckStateChangedAllOf) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyAckStateChangedAllOf) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyAckStateChangedAllOf) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.


### GetAckState

`func (o *NotifyAckStateChangedAllOf) GetAckState() AckState`

GetAckState returns the AckState field if non-nil, zero value otherwise.

### GetAckStateOk

`func (o *NotifyAckStateChangedAllOf) GetAckStateOk() (*AckState, bool)`

GetAckStateOk returns a tuple with the AckState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckState

`func (o *NotifyAckStateChangedAllOf) SetAckState(v AckState)`

SetAckState sets AckState field to given value.


### GetAckUserId

`func (o *NotifyAckStateChangedAllOf) GetAckUserId() string`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *NotifyAckStateChangedAllOf) GetAckUserIdOk() (*string, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *NotifyAckStateChangedAllOf) SetAckUserId(v string)`

SetAckUserId sets AckUserId field to given value.


### GetAckSystemId

`func (o *NotifyAckStateChangedAllOf) GetAckSystemId() string`

GetAckSystemId returns the AckSystemId field if non-nil, zero value otherwise.

### GetAckSystemIdOk

`func (o *NotifyAckStateChangedAllOf) GetAckSystemIdOk() (*string, bool)`

GetAckSystemIdOk returns a tuple with the AckSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSystemId

`func (o *NotifyAckStateChangedAllOf) SetAckSystemId(v string)`

SetAckSystemId sets AckSystemId field to given value.

### HasAckSystemId

`func (o *NotifyAckStateChangedAllOf) HasAckSystemId() bool`

HasAckSystemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


