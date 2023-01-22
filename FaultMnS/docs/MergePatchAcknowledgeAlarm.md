# MergePatchAcknowledgeAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckUserId** | **string** |  | 
**AckSystemId** | Pointer to **string** |  | [optional] 
**AckState** | [**AckState**](AckState.md) |  | 

## Methods

### NewMergePatchAcknowledgeAlarm

`func NewMergePatchAcknowledgeAlarm(ackUserId string, ackState AckState, ) *MergePatchAcknowledgeAlarm`

NewMergePatchAcknowledgeAlarm instantiates a new MergePatchAcknowledgeAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergePatchAcknowledgeAlarmWithDefaults

`func NewMergePatchAcknowledgeAlarmWithDefaults() *MergePatchAcknowledgeAlarm`

NewMergePatchAcknowledgeAlarmWithDefaults instantiates a new MergePatchAcknowledgeAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckUserId

`func (o *MergePatchAcknowledgeAlarm) GetAckUserId() string`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *MergePatchAcknowledgeAlarm) GetAckUserIdOk() (*string, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *MergePatchAcknowledgeAlarm) SetAckUserId(v string)`

SetAckUserId sets AckUserId field to given value.


### GetAckSystemId

`func (o *MergePatchAcknowledgeAlarm) GetAckSystemId() string`

GetAckSystemId returns the AckSystemId field if non-nil, zero value otherwise.

### GetAckSystemIdOk

`func (o *MergePatchAcknowledgeAlarm) GetAckSystemIdOk() (*string, bool)`

GetAckSystemIdOk returns a tuple with the AckSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSystemId

`func (o *MergePatchAcknowledgeAlarm) SetAckSystemId(v string)`

SetAckSystemId sets AckSystemId field to given value.

### HasAckSystemId

`func (o *MergePatchAcknowledgeAlarm) HasAckSystemId() bool`

HasAckSystemId returns a boolean if a field has been set.

### GetAckState

`func (o *MergePatchAcknowledgeAlarm) GetAckState() AckState`

GetAckState returns the AckState field if non-nil, zero value otherwise.

### GetAckStateOk

`func (o *MergePatchAcknowledgeAlarm) GetAckStateOk() (*AckState, bool)`

GetAckStateOk returns a tuple with the AckState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckState

`func (o *MergePatchAcknowledgeAlarm) SetAckState(v AckState)`

SetAckState sets AckState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


