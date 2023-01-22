# AlarmsAlarmIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckUserId** | **string** |  | 
**AckSystemId** | Pointer to **string** |  | [optional] 
**AckState** | [**AckState**](AckState.md) |  | 
**ClearUserId** | **string** |  | 
**ClearSystemId** | Pointer to **string** |  | [optional] 
**PerceivedSeverity** | **string** |  | 

## Methods

### NewAlarmsAlarmIdPatchRequest

`func NewAlarmsAlarmIdPatchRequest(ackUserId string, ackState AckState, clearUserId string, perceivedSeverity string, ) *AlarmsAlarmIdPatchRequest`

NewAlarmsAlarmIdPatchRequest instantiates a new AlarmsAlarmIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmsAlarmIdPatchRequestWithDefaults

`func NewAlarmsAlarmIdPatchRequestWithDefaults() *AlarmsAlarmIdPatchRequest`

NewAlarmsAlarmIdPatchRequestWithDefaults instantiates a new AlarmsAlarmIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckUserId

`func (o *AlarmsAlarmIdPatchRequest) GetAckUserId() string`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *AlarmsAlarmIdPatchRequest) GetAckUserIdOk() (*string, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *AlarmsAlarmIdPatchRequest) SetAckUserId(v string)`

SetAckUserId sets AckUserId field to given value.


### GetAckSystemId

`func (o *AlarmsAlarmIdPatchRequest) GetAckSystemId() string`

GetAckSystemId returns the AckSystemId field if non-nil, zero value otherwise.

### GetAckSystemIdOk

`func (o *AlarmsAlarmIdPatchRequest) GetAckSystemIdOk() (*string, bool)`

GetAckSystemIdOk returns a tuple with the AckSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSystemId

`func (o *AlarmsAlarmIdPatchRequest) SetAckSystemId(v string)`

SetAckSystemId sets AckSystemId field to given value.

### HasAckSystemId

`func (o *AlarmsAlarmIdPatchRequest) HasAckSystemId() bool`

HasAckSystemId returns a boolean if a field has been set.

### GetAckState

`func (o *AlarmsAlarmIdPatchRequest) GetAckState() AckState`

GetAckState returns the AckState field if non-nil, zero value otherwise.

### GetAckStateOk

`func (o *AlarmsAlarmIdPatchRequest) GetAckStateOk() (*AckState, bool)`

GetAckStateOk returns a tuple with the AckState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckState

`func (o *AlarmsAlarmIdPatchRequest) SetAckState(v AckState)`

SetAckState sets AckState field to given value.


### GetClearUserId

`func (o *AlarmsAlarmIdPatchRequest) GetClearUserId() string`

GetClearUserId returns the ClearUserId field if non-nil, zero value otherwise.

### GetClearUserIdOk

`func (o *AlarmsAlarmIdPatchRequest) GetClearUserIdOk() (*string, bool)`

GetClearUserIdOk returns a tuple with the ClearUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearUserId

`func (o *AlarmsAlarmIdPatchRequest) SetClearUserId(v string)`

SetClearUserId sets ClearUserId field to given value.


### GetClearSystemId

`func (o *AlarmsAlarmIdPatchRequest) GetClearSystemId() string`

GetClearSystemId returns the ClearSystemId field if non-nil, zero value otherwise.

### GetClearSystemIdOk

`func (o *AlarmsAlarmIdPatchRequest) GetClearSystemIdOk() (*string, bool)`

GetClearSystemIdOk returns a tuple with the ClearSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSystemId

`func (o *AlarmsAlarmIdPatchRequest) SetClearSystemId(v string)`

SetClearSystemId sets ClearSystemId field to given value.

### HasClearSystemId

`func (o *AlarmsAlarmIdPatchRequest) HasClearSystemId() bool`

HasClearSystemId returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *AlarmsAlarmIdPatchRequest) GetPerceivedSeverity() string`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *AlarmsAlarmIdPatchRequest) GetPerceivedSeverityOk() (*string, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *AlarmsAlarmIdPatchRequest) SetPerceivedSeverity(v string)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


