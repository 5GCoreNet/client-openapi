# MessageFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReqUe** | [**ValTargetUe**](ValTargetUe.md) |  | 
**TgtUe** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of VAL User or UE IDs whose message to be sent. | [optional] 
**MaxMsgs** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**Scheds** | Pointer to [**[]ScheduledCommunicationTime**](ScheduledCommunicationTime.md) | Time frame associated with total number of messages. | [optional] 
**MsgTypes** | Pointer to **[]string** | List of message types to be sent to VAL UE. | [optional] 

## Methods

### NewMessageFilter

`func NewMessageFilter(reqUe ValTargetUe, ) *MessageFilter`

NewMessageFilter instantiates a new MessageFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageFilterWithDefaults

`func NewMessageFilterWithDefaults() *MessageFilter`

NewMessageFilterWithDefaults instantiates a new MessageFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReqUe

`func (o *MessageFilter) GetReqUe() ValTargetUe`

GetReqUe returns the ReqUe field if non-nil, zero value otherwise.

### GetReqUeOk

`func (o *MessageFilter) GetReqUeOk() (*ValTargetUe, bool)`

GetReqUeOk returns a tuple with the ReqUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUe

`func (o *MessageFilter) SetReqUe(v ValTargetUe)`

SetReqUe sets ReqUe field to given value.


### GetTgtUe

`func (o *MessageFilter) GetTgtUe() []ValTargetUe`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *MessageFilter) GetTgtUeOk() (*[]ValTargetUe, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *MessageFilter) SetTgtUe(v []ValTargetUe)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *MessageFilter) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *MessageFilter) GetMaxMsgs() int32`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *MessageFilter) GetMaxMsgsOk() (*int32, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *MessageFilter) SetMaxMsgs(v int32)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *MessageFilter) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetScheds

`func (o *MessageFilter) GetScheds() []ScheduledCommunicationTime`

GetScheds returns the Scheds field if non-nil, zero value otherwise.

### GetSchedsOk

`func (o *MessageFilter) GetSchedsOk() (*[]ScheduledCommunicationTime, bool)`

GetSchedsOk returns a tuple with the Scheds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheds

`func (o *MessageFilter) SetScheds(v []ScheduledCommunicationTime)`

SetScheds sets Scheds field to given value.

### HasScheds

`func (o *MessageFilter) HasScheds() bool`

HasScheds returns a boolean if a field has been set.

### GetMsgTypes

`func (o *MessageFilter) GetMsgTypes() []string`

GetMsgTypes returns the MsgTypes field if non-nil, zero value otherwise.

### GetMsgTypesOk

`func (o *MessageFilter) GetMsgTypesOk() (*[]string, bool)`

GetMsgTypesOk returns a tuple with the MsgTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgTypes

`func (o *MessageFilter) SetMsgTypes(v []string)`

SetMsgTypes sets MsgTypes field to given value.

### HasMsgTypes

`func (o *MessageFilter) HasMsgTypes() bool`

HasMsgTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


