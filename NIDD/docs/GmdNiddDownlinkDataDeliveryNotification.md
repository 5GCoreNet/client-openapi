# GmdNiddDownlinkDataDeliveryNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NiddDownlinkDataTransfer** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**GmdResults** | [**[]GmdResult**](GmdResult.md) | Indicates the group message delivery result. | 

## Methods

### NewGmdNiddDownlinkDataDeliveryNotification

`func NewGmdNiddDownlinkDataDeliveryNotification(niddDownlinkDataTransfer string, gmdResults []GmdResult, ) *GmdNiddDownlinkDataDeliveryNotification`

NewGmdNiddDownlinkDataDeliveryNotification instantiates a new GmdNiddDownlinkDataDeliveryNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGmdNiddDownlinkDataDeliveryNotificationWithDefaults

`func NewGmdNiddDownlinkDataDeliveryNotificationWithDefaults() *GmdNiddDownlinkDataDeliveryNotification`

NewGmdNiddDownlinkDataDeliveryNotificationWithDefaults instantiates a new GmdNiddDownlinkDataDeliveryNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNiddDownlinkDataTransfer

`func (o *GmdNiddDownlinkDataDeliveryNotification) GetNiddDownlinkDataTransfer() string`

GetNiddDownlinkDataTransfer returns the NiddDownlinkDataTransfer field if non-nil, zero value otherwise.

### GetNiddDownlinkDataTransferOk

`func (o *GmdNiddDownlinkDataDeliveryNotification) GetNiddDownlinkDataTransferOk() (*string, bool)`

GetNiddDownlinkDataTransferOk returns a tuple with the NiddDownlinkDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddDownlinkDataTransfer

`func (o *GmdNiddDownlinkDataDeliveryNotification) SetNiddDownlinkDataTransfer(v string)`

SetNiddDownlinkDataTransfer sets NiddDownlinkDataTransfer field to given value.


### GetGmdResults

`func (o *GmdNiddDownlinkDataDeliveryNotification) GetGmdResults() []GmdResult`

GetGmdResults returns the GmdResults field if non-nil, zero value otherwise.

### GetGmdResultsOk

`func (o *GmdNiddDownlinkDataDeliveryNotification) GetGmdResultsOk() (*[]GmdResult, bool)`

GetGmdResultsOk returns a tuple with the GmdResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmdResults

`func (o *GmdNiddDownlinkDataDeliveryNotification) SetGmdResults(v []GmdResult)`

SetGmdResults sets GmdResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


