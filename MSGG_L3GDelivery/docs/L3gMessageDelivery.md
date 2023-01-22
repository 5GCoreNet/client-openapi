# L3gMessageDelivery

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

### NewL3gMessageDelivery

`func NewL3gMessageDelivery(oriAddr Address, destAddr Address, msgId string, ) *L3gMessageDelivery`

NewL3gMessageDelivery instantiates a new L3gMessageDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL3gMessageDeliveryWithDefaults

`func NewL3gMessageDeliveryWithDefaults() *L3gMessageDelivery`

NewL3gMessageDeliveryWithDefaults instantiates a new L3gMessageDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriAddr

`func (o *L3gMessageDelivery) GetOriAddr() Address`

GetOriAddr returns the OriAddr field if non-nil, zero value otherwise.

### GetOriAddrOk

`func (o *L3gMessageDelivery) GetOriAddrOk() (*Address, bool)`

GetOriAddrOk returns a tuple with the OriAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriAddr

`func (o *L3gMessageDelivery) SetOriAddr(v Address)`

SetOriAddr sets OriAddr field to given value.


### GetDestAddr

`func (o *L3gMessageDelivery) GetDestAddr() Address`

GetDestAddr returns the DestAddr field if non-nil, zero value otherwise.

### GetDestAddrOk

`func (o *L3gMessageDelivery) GetDestAddrOk() (*Address, bool)`

GetDestAddrOk returns a tuple with the DestAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAddr

`func (o *L3gMessageDelivery) SetDestAddr(v Address)`

SetDestAddr sets DestAddr field to given value.


### GetAppId

`func (o *L3gMessageDelivery) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *L3gMessageDelivery) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *L3gMessageDelivery) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *L3gMessageDelivery) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetMsgId

`func (o *L3gMessageDelivery) GetMsgId() string`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *L3gMessageDelivery) GetMsgIdOk() (*string, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *L3gMessageDelivery) SetMsgId(v string)`

SetMsgId sets MsgId field to given value.


### GetDelivStReqInd

`func (o *L3gMessageDelivery) GetDelivStReqInd() bool`

GetDelivStReqInd returns the DelivStReqInd field if non-nil, zero value otherwise.

### GetDelivStReqIndOk

`func (o *L3gMessageDelivery) GetDelivStReqIndOk() (*bool, bool)`

GetDelivStReqIndOk returns a tuple with the DelivStReqInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivStReqInd

`func (o *L3gMessageDelivery) SetDelivStReqInd(v bool)`

SetDelivStReqInd sets DelivStReqInd field to given value.

### HasDelivStReqInd

`func (o *L3gMessageDelivery) HasDelivStReqInd() bool`

HasDelivStReqInd returns a boolean if a field has been set.

### GetPayload

`func (o *L3gMessageDelivery) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *L3gMessageDelivery) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *L3gMessageDelivery) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *L3gMessageDelivery) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSegInd

`func (o *L3gMessageDelivery) GetSegInd() bool`

GetSegInd returns the SegInd field if non-nil, zero value otherwise.

### GetSegIndOk

`func (o *L3gMessageDelivery) GetSegIndOk() (*bool, bool)`

GetSegIndOk returns a tuple with the SegInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegInd

`func (o *L3gMessageDelivery) SetSegInd(v bool)`

SetSegInd sets SegInd field to given value.

### HasSegInd

`func (o *L3gMessageDelivery) HasSegInd() bool`

HasSegInd returns a boolean if a field has been set.

### GetSegParams

`func (o *L3gMessageDelivery) GetSegParams() MessageSegmentParameters`

GetSegParams returns the SegParams field if non-nil, zero value otherwise.

### GetSegParamsOk

`func (o *L3gMessageDelivery) GetSegParamsOk() (*MessageSegmentParameters, bool)`

GetSegParamsOk returns a tuple with the SegParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegParams

`func (o *L3gMessageDelivery) SetSegParams(v MessageSegmentParameters)`

SetSegParams sets SegParams field to given value.

### HasSegParams

`func (o *L3gMessageDelivery) HasSegParams() bool`

HasSegParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


