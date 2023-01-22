# ScsAsIdConfigurationsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NiddDownlinkDataTransfer** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**DeliveryStatus** | [**DeliveryStatus**](DeliveryStatus.md) |  | 
**RequestedRetransmissionTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**GmdResults** | [**[]GmdResult**](GmdResult.md) | Indicates the group message delivery result. | 

## Methods

### NewScsAsIdConfigurationsGetRequest

`func NewScsAsIdConfigurationsGetRequest(niddDownlinkDataTransfer string, deliveryStatus DeliveryStatus, gmdResults []GmdResult, ) *ScsAsIdConfigurationsGetRequest`

NewScsAsIdConfigurationsGetRequest instantiates a new ScsAsIdConfigurationsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScsAsIdConfigurationsGetRequestWithDefaults

`func NewScsAsIdConfigurationsGetRequestWithDefaults() *ScsAsIdConfigurationsGetRequest`

NewScsAsIdConfigurationsGetRequestWithDefaults instantiates a new ScsAsIdConfigurationsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNiddDownlinkDataTransfer

`func (o *ScsAsIdConfigurationsGetRequest) GetNiddDownlinkDataTransfer() string`

GetNiddDownlinkDataTransfer returns the NiddDownlinkDataTransfer field if non-nil, zero value otherwise.

### GetNiddDownlinkDataTransferOk

`func (o *ScsAsIdConfigurationsGetRequest) GetNiddDownlinkDataTransferOk() (*string, bool)`

GetNiddDownlinkDataTransferOk returns a tuple with the NiddDownlinkDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddDownlinkDataTransfer

`func (o *ScsAsIdConfigurationsGetRequest) SetNiddDownlinkDataTransfer(v string)`

SetNiddDownlinkDataTransfer sets NiddDownlinkDataTransfer field to given value.


### GetDeliveryStatus

`func (o *ScsAsIdConfigurationsGetRequest) GetDeliveryStatus() DeliveryStatus`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *ScsAsIdConfigurationsGetRequest) GetDeliveryStatusOk() (*DeliveryStatus, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *ScsAsIdConfigurationsGetRequest) SetDeliveryStatus(v DeliveryStatus)`

SetDeliveryStatus sets DeliveryStatus field to given value.


### GetRequestedRetransmissionTime

`func (o *ScsAsIdConfigurationsGetRequest) GetRequestedRetransmissionTime() time.Time`

GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field if non-nil, zero value otherwise.

### GetRequestedRetransmissionTimeOk

`func (o *ScsAsIdConfigurationsGetRequest) GetRequestedRetransmissionTimeOk() (*time.Time, bool)`

GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRetransmissionTime

`func (o *ScsAsIdConfigurationsGetRequest) SetRequestedRetransmissionTime(v time.Time)`

SetRequestedRetransmissionTime sets RequestedRetransmissionTime field to given value.

### HasRequestedRetransmissionTime

`func (o *ScsAsIdConfigurationsGetRequest) HasRequestedRetransmissionTime() bool`

HasRequestedRetransmissionTime returns a boolean if a field has been set.

### GetGmdResults

`func (o *ScsAsIdConfigurationsGetRequest) GetGmdResults() []GmdResult`

GetGmdResults returns the GmdResults field if non-nil, zero value otherwise.

### GetGmdResultsOk

`func (o *ScsAsIdConfigurationsGetRequest) GetGmdResultsOk() (*[]GmdResult, bool)`

GetGmdResultsOk returns a tuple with the GmdResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmdResults

`func (o *ScsAsIdConfigurationsGetRequest) SetGmdResults(v []GmdResult)`

SetGmdResults sets GmdResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


