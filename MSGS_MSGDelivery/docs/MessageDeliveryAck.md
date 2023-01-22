# MessageDeliveryAck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriAddr** | [**Address**](Address.md) |  | 
**MsgId** | **string** |  | 
**Status** | Pointer to [**DeliveryStatus**](DeliveryStatus.md) |  | [optional] 
**FailureCause** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageDeliveryAck

`func NewMessageDeliveryAck(oriAddr Address, msgId string, ) *MessageDeliveryAck`

NewMessageDeliveryAck instantiates a new MessageDeliveryAck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDeliveryAckWithDefaults

`func NewMessageDeliveryAckWithDefaults() *MessageDeliveryAck`

NewMessageDeliveryAckWithDefaults instantiates a new MessageDeliveryAck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriAddr

`func (o *MessageDeliveryAck) GetOriAddr() Address`

GetOriAddr returns the OriAddr field if non-nil, zero value otherwise.

### GetOriAddrOk

`func (o *MessageDeliveryAck) GetOriAddrOk() (*Address, bool)`

GetOriAddrOk returns a tuple with the OriAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriAddr

`func (o *MessageDeliveryAck) SetOriAddr(v Address)`

SetOriAddr sets OriAddr field to given value.


### GetMsgId

`func (o *MessageDeliveryAck) GetMsgId() string`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MessageDeliveryAck) GetMsgIdOk() (*string, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MessageDeliveryAck) SetMsgId(v string)`

SetMsgId sets MsgId field to given value.


### GetStatus

`func (o *MessageDeliveryAck) GetStatus() DeliveryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageDeliveryAck) GetStatusOk() (*DeliveryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageDeliveryAck) SetStatus(v DeliveryStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageDeliveryAck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailureCause

`func (o *MessageDeliveryAck) GetFailureCause() string`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *MessageDeliveryAck) GetFailureCauseOk() (*string, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *MessageDeliveryAck) SetFailureCause(v string)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *MessageDeliveryAck) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


