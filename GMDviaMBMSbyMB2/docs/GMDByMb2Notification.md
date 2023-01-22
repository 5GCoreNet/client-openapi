# GMDByMb2Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**DeliveryTriggerStatus** | **bool** | Indicates whether delivery of group message payload corresponding to the TMGI was successful (TRUE) or not (FALSE) | 

## Methods

### NewGMDByMb2Notification

`func NewGMDByMb2Notification(transaction string, deliveryTriggerStatus bool, ) *GMDByMb2Notification`

NewGMDByMb2Notification instantiates a new GMDByMb2Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGMDByMb2NotificationWithDefaults

`func NewGMDByMb2NotificationWithDefaults() *GMDByMb2Notification`

NewGMDByMb2NotificationWithDefaults instantiates a new GMDByMb2Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *GMDByMb2Notification) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *GMDByMb2Notification) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *GMDByMb2Notification) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetDeliveryTriggerStatus

`func (o *GMDByMb2Notification) GetDeliveryTriggerStatus() bool`

GetDeliveryTriggerStatus returns the DeliveryTriggerStatus field if non-nil, zero value otherwise.

### GetDeliveryTriggerStatusOk

`func (o *GMDByMb2Notification) GetDeliveryTriggerStatusOk() (*bool, bool)`

GetDeliveryTriggerStatusOk returns a tuple with the DeliveryTriggerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTriggerStatus

`func (o *GMDByMb2Notification) SetDeliveryTriggerStatus(v bool)`

SetDeliveryTriggerStatus sets DeliveryTriggerStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


