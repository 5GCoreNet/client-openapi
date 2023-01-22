# NiddDownlinkDataDeliveryStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NiddDownlinkDataTransfer** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**DeliveryStatus** | [**DeliveryStatus**](DeliveryStatus.md) |  | 
**RequestedRetransmissionTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewNiddDownlinkDataDeliveryStatusNotification

`func NewNiddDownlinkDataDeliveryStatusNotification(niddDownlinkDataTransfer string, deliveryStatus DeliveryStatus, ) *NiddDownlinkDataDeliveryStatusNotification`

NewNiddDownlinkDataDeliveryStatusNotification instantiates a new NiddDownlinkDataDeliveryStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddDownlinkDataDeliveryStatusNotificationWithDefaults

`func NewNiddDownlinkDataDeliveryStatusNotificationWithDefaults() *NiddDownlinkDataDeliveryStatusNotification`

NewNiddDownlinkDataDeliveryStatusNotificationWithDefaults instantiates a new NiddDownlinkDataDeliveryStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNiddDownlinkDataTransfer

`func (o *NiddDownlinkDataDeliveryStatusNotification) GetNiddDownlinkDataTransfer() string`

GetNiddDownlinkDataTransfer returns the NiddDownlinkDataTransfer field if non-nil, zero value otherwise.

### GetNiddDownlinkDataTransferOk

`func (o *NiddDownlinkDataDeliveryStatusNotification) GetNiddDownlinkDataTransferOk() (*string, bool)`

GetNiddDownlinkDataTransferOk returns a tuple with the NiddDownlinkDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddDownlinkDataTransfer

`func (o *NiddDownlinkDataDeliveryStatusNotification) SetNiddDownlinkDataTransfer(v string)`

SetNiddDownlinkDataTransfer sets NiddDownlinkDataTransfer field to given value.


### GetDeliveryStatus

`func (o *NiddDownlinkDataDeliveryStatusNotification) GetDeliveryStatus() DeliveryStatus`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *NiddDownlinkDataDeliveryStatusNotification) GetDeliveryStatusOk() (*DeliveryStatus, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *NiddDownlinkDataDeliveryStatusNotification) SetDeliveryStatus(v DeliveryStatus)`

SetDeliveryStatus sets DeliveryStatus field to given value.


### GetRequestedRetransmissionTime

`func (o *NiddDownlinkDataDeliveryStatusNotification) GetRequestedRetransmissionTime() time.Time`

GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field if non-nil, zero value otherwise.

### GetRequestedRetransmissionTimeOk

`func (o *NiddDownlinkDataDeliveryStatusNotification) GetRequestedRetransmissionTimeOk() (*time.Time, bool)`

GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRetransmissionTime

`func (o *NiddDownlinkDataDeliveryStatusNotification) SetRequestedRetransmissionTime(v time.Time)`

SetRequestedRetransmissionTime sets RequestedRetransmissionTime field to given value.

### HasRequestedRetransmissionTime

`func (o *NiddDownlinkDataDeliveryStatusNotification) HasRequestedRetransmissionTime() bool`

HasRequestedRetransmissionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


