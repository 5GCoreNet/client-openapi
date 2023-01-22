# DeviceTriggeringDeliveryReportNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**Result** | [**DeliveryResult**](DeliveryResult.md) |  | 

## Methods

### NewDeviceTriggeringDeliveryReportNotification

`func NewDeviceTriggeringDeliveryReportNotification(transaction string, result DeliveryResult, ) *DeviceTriggeringDeliveryReportNotification`

NewDeviceTriggeringDeliveryReportNotification instantiates a new DeviceTriggeringDeliveryReportNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTriggeringDeliveryReportNotificationWithDefaults

`func NewDeviceTriggeringDeliveryReportNotificationWithDefaults() *DeviceTriggeringDeliveryReportNotification`

NewDeviceTriggeringDeliveryReportNotificationWithDefaults instantiates a new DeviceTriggeringDeliveryReportNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *DeviceTriggeringDeliveryReportNotification) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *DeviceTriggeringDeliveryReportNotification) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *DeviceTriggeringDeliveryReportNotification) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetResult

`func (o *DeviceTriggeringDeliveryReportNotification) GetResult() DeliveryResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeviceTriggeringDeliveryReportNotification) GetResultOk() (*DeliveryResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeviceTriggeringDeliveryReportNotification) SetResult(v DeliveryResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


