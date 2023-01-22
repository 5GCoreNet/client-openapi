# GMDByxMBNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**DeliveryTriggerStatus** | **bool** | Indicates whether delivery of group message payload was successful(TRUE) or not (FALSE) | 

## Methods

### NewGMDByxMBNotification

`func NewGMDByxMBNotification(transaction string, deliveryTriggerStatus bool, ) *GMDByxMBNotification`

NewGMDByxMBNotification instantiates a new GMDByxMBNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGMDByxMBNotificationWithDefaults

`func NewGMDByxMBNotificationWithDefaults() *GMDByxMBNotification`

NewGMDByxMBNotificationWithDefaults instantiates a new GMDByxMBNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *GMDByxMBNotification) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *GMDByxMBNotification) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *GMDByxMBNotification) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetDeliveryTriggerStatus

`func (o *GMDByxMBNotification) GetDeliveryTriggerStatus() bool`

GetDeliveryTriggerStatus returns the DeliveryTriggerStatus field if non-nil, zero value otherwise.

### GetDeliveryTriggerStatusOk

`func (o *GMDByxMBNotification) GetDeliveryTriggerStatusOk() (*bool, bool)`

GetDeliveryTriggerStatusOk returns a tuple with the DeliveryTriggerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTriggerStatus

`func (o *GMDByxMBNotification) SetDeliveryTriggerStatus(v bool)`

SetDeliveryTriggerStatus sets DeliveryTriggerStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


