# N3gMessageDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriAddr** | [**Address**](Address.md) |  | 
**DestAddr** | [**Address**](Address.md) |  | 
**AppId** | Pointer to **string** |  | [optional] 
**MsgId** | **string** |  | 
**DelivStReqInd** | Pointer to **bool** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**SegInd** | Pointer to **bool** |  | [optional] 
**SegParams** | Pointer to [**MessageSegmentParameters**](MessageSegmentParameters.md) |  | [optional] 

## Methods

### NewN3gMessageDelivery

`func NewN3gMessageDelivery(oriAddr Address, destAddr Address, msgId string, ) *N3gMessageDelivery`

NewN3gMessageDelivery instantiates a new N3gMessageDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN3gMessageDeliveryWithDefaults

`func NewN3gMessageDeliveryWithDefaults() *N3gMessageDelivery`

NewN3gMessageDeliveryWithDefaults instantiates a new N3gMessageDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriAddr

`func (o *N3gMessageDelivery) GetOriAddr() Address`

GetOriAddr returns the OriAddr field if non-nil, zero value otherwise.

### GetOriAddrOk

`func (o *N3gMessageDelivery) GetOriAddrOk() (*Address, bool)`

GetOriAddrOk returns a tuple with the OriAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriAddr

`func (o *N3gMessageDelivery) SetOriAddr(v Address)`

SetOriAddr sets OriAddr field to given value.


### GetDestAddr

`func (o *N3gMessageDelivery) GetDestAddr() Address`

GetDestAddr returns the DestAddr field if non-nil, zero value otherwise.

### GetDestAddrOk

`func (o *N3gMessageDelivery) GetDestAddrOk() (*Address, bool)`

GetDestAddrOk returns a tuple with the DestAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAddr

`func (o *N3gMessageDelivery) SetDestAddr(v Address)`

SetDestAddr sets DestAddr field to given value.


### GetAppId

`func (o *N3gMessageDelivery) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *N3gMessageDelivery) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *N3gMessageDelivery) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *N3gMessageDelivery) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetMsgId

`func (o *N3gMessageDelivery) GetMsgId() string`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *N3gMessageDelivery) GetMsgIdOk() (*string, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *N3gMessageDelivery) SetMsgId(v string)`

SetMsgId sets MsgId field to given value.


### GetDelivStReqInd

`func (o *N3gMessageDelivery) GetDelivStReqInd() bool`

GetDelivStReqInd returns the DelivStReqInd field if non-nil, zero value otherwise.

### GetDelivStReqIndOk

`func (o *N3gMessageDelivery) GetDelivStReqIndOk() (*bool, bool)`

GetDelivStReqIndOk returns a tuple with the DelivStReqInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivStReqInd

`func (o *N3gMessageDelivery) SetDelivStReqInd(v bool)`

SetDelivStReqInd sets DelivStReqInd field to given value.

### HasDelivStReqInd

`func (o *N3gMessageDelivery) HasDelivStReqInd() bool`

HasDelivStReqInd returns a boolean if a field has been set.

### GetPayload

`func (o *N3gMessageDelivery) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *N3gMessageDelivery) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *N3gMessageDelivery) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *N3gMessageDelivery) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSegInd

`func (o *N3gMessageDelivery) GetSegInd() bool`

GetSegInd returns the SegInd field if non-nil, zero value otherwise.

### GetSegIndOk

`func (o *N3gMessageDelivery) GetSegIndOk() (*bool, bool)`

GetSegIndOk returns a tuple with the SegInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegInd

`func (o *N3gMessageDelivery) SetSegInd(v bool)`

SetSegInd sets SegInd field to given value.

### HasSegInd

`func (o *N3gMessageDelivery) HasSegInd() bool`

HasSegInd returns a boolean if a field has been set.

### GetSegParams

`func (o *N3gMessageDelivery) GetSegParams() MessageSegmentParameters`

GetSegParams returns the SegParams field if non-nil, zero value otherwise.

### GetSegParamsOk

`func (o *N3gMessageDelivery) GetSegParamsOk() (*MessageSegmentParameters, bool)`

GetSegParamsOk returns a tuple with the SegParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegParams

`func (o *N3gMessageDelivery) SetSegParams(v MessageSegmentParameters)`

SetSegParams sets SegParams field to given value.

### HasSegParams

`func (o *N3gMessageDelivery) HasSegParams() bool`

HasSegParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


