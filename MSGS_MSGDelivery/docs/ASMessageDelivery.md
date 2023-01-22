# ASMessageDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriAddr** | [**Address**](Address.md) |  | 
**DestAddr** | [**Address**](Address.md) |  | 
**AppId** | Pointer to **string** |  | [optional] 
**MsgId** | **string** |  | 
**SecCred** | Pointer to **string** |  | [optional] 
**DelivStReqInd** | Pointer to **bool** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**Priority**](Priority.md) |  | [optional] 
**SegInd** | Pointer to **bool** |  | [optional] 
**SegParams** | Pointer to [**MessageSegmentParameters**](MessageSegmentParameters.md) |  | [optional] 
**StoAndFwInd** | **bool** |  | 
**StoAndFwParams** | Pointer to [**StoreAndForwardParameters**](StoreAndForwardParameters.md) |  | [optional] 
**Latency** | Pointer to **int32** |  | [optional] 

## Methods

### NewASMessageDelivery

`func NewASMessageDelivery(oriAddr Address, destAddr Address, msgId string, stoAndFwInd bool, ) *ASMessageDelivery`

NewASMessageDelivery instantiates a new ASMessageDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASMessageDeliveryWithDefaults

`func NewASMessageDeliveryWithDefaults() *ASMessageDelivery`

NewASMessageDeliveryWithDefaults instantiates a new ASMessageDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriAddr

`func (o *ASMessageDelivery) GetOriAddr() Address`

GetOriAddr returns the OriAddr field if non-nil, zero value otherwise.

### GetOriAddrOk

`func (o *ASMessageDelivery) GetOriAddrOk() (*Address, bool)`

GetOriAddrOk returns a tuple with the OriAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriAddr

`func (o *ASMessageDelivery) SetOriAddr(v Address)`

SetOriAddr sets OriAddr field to given value.


### GetDestAddr

`func (o *ASMessageDelivery) GetDestAddr() Address`

GetDestAddr returns the DestAddr field if non-nil, zero value otherwise.

### GetDestAddrOk

`func (o *ASMessageDelivery) GetDestAddrOk() (*Address, bool)`

GetDestAddrOk returns a tuple with the DestAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAddr

`func (o *ASMessageDelivery) SetDestAddr(v Address)`

SetDestAddr sets DestAddr field to given value.


### GetAppId

`func (o *ASMessageDelivery) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ASMessageDelivery) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ASMessageDelivery) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ASMessageDelivery) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetMsgId

`func (o *ASMessageDelivery) GetMsgId() string`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *ASMessageDelivery) GetMsgIdOk() (*string, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *ASMessageDelivery) SetMsgId(v string)`

SetMsgId sets MsgId field to given value.


### GetSecCred

`func (o *ASMessageDelivery) GetSecCred() string`

GetSecCred returns the SecCred field if non-nil, zero value otherwise.

### GetSecCredOk

`func (o *ASMessageDelivery) GetSecCredOk() (*string, bool)`

GetSecCredOk returns a tuple with the SecCred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecCred

`func (o *ASMessageDelivery) SetSecCred(v string)`

SetSecCred sets SecCred field to given value.

### HasSecCred

`func (o *ASMessageDelivery) HasSecCred() bool`

HasSecCred returns a boolean if a field has been set.

### GetDelivStReqInd

`func (o *ASMessageDelivery) GetDelivStReqInd() bool`

GetDelivStReqInd returns the DelivStReqInd field if non-nil, zero value otherwise.

### GetDelivStReqIndOk

`func (o *ASMessageDelivery) GetDelivStReqIndOk() (*bool, bool)`

GetDelivStReqIndOk returns a tuple with the DelivStReqInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivStReqInd

`func (o *ASMessageDelivery) SetDelivStReqInd(v bool)`

SetDelivStReqInd sets DelivStReqInd field to given value.

### HasDelivStReqInd

`func (o *ASMessageDelivery) HasDelivStReqInd() bool`

HasDelivStReqInd returns a boolean if a field has been set.

### GetPayload

`func (o *ASMessageDelivery) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ASMessageDelivery) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ASMessageDelivery) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ASMessageDelivery) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPriority

`func (o *ASMessageDelivery) GetPriority() Priority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ASMessageDelivery) GetPriorityOk() (*Priority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ASMessageDelivery) SetPriority(v Priority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ASMessageDelivery) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSegInd

`func (o *ASMessageDelivery) GetSegInd() bool`

GetSegInd returns the SegInd field if non-nil, zero value otherwise.

### GetSegIndOk

`func (o *ASMessageDelivery) GetSegIndOk() (*bool, bool)`

GetSegIndOk returns a tuple with the SegInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegInd

`func (o *ASMessageDelivery) SetSegInd(v bool)`

SetSegInd sets SegInd field to given value.

### HasSegInd

`func (o *ASMessageDelivery) HasSegInd() bool`

HasSegInd returns a boolean if a field has been set.

### GetSegParams

`func (o *ASMessageDelivery) GetSegParams() MessageSegmentParameters`

GetSegParams returns the SegParams field if non-nil, zero value otherwise.

### GetSegParamsOk

`func (o *ASMessageDelivery) GetSegParamsOk() (*MessageSegmentParameters, bool)`

GetSegParamsOk returns a tuple with the SegParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegParams

`func (o *ASMessageDelivery) SetSegParams(v MessageSegmentParameters)`

SetSegParams sets SegParams field to given value.

### HasSegParams

`func (o *ASMessageDelivery) HasSegParams() bool`

HasSegParams returns a boolean if a field has been set.

### GetStoAndFwInd

`func (o *ASMessageDelivery) GetStoAndFwInd() bool`

GetStoAndFwInd returns the StoAndFwInd field if non-nil, zero value otherwise.

### GetStoAndFwIndOk

`func (o *ASMessageDelivery) GetStoAndFwIndOk() (*bool, bool)`

GetStoAndFwIndOk returns a tuple with the StoAndFwInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoAndFwInd

`func (o *ASMessageDelivery) SetStoAndFwInd(v bool)`

SetStoAndFwInd sets StoAndFwInd field to given value.


### GetStoAndFwParams

`func (o *ASMessageDelivery) GetStoAndFwParams() StoreAndForwardParameters`

GetStoAndFwParams returns the StoAndFwParams field if non-nil, zero value otherwise.

### GetStoAndFwParamsOk

`func (o *ASMessageDelivery) GetStoAndFwParamsOk() (*StoreAndForwardParameters, bool)`

GetStoAndFwParamsOk returns a tuple with the StoAndFwParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoAndFwParams

`func (o *ASMessageDelivery) SetStoAndFwParams(v StoreAndForwardParameters)`

SetStoAndFwParams sets StoAndFwParams field to given value.

### HasStoAndFwParams

`func (o *ASMessageDelivery) HasStoAndFwParams() bool`

HasStoAndFwParams returns a boolean if a field has been set.

### GetLatency

`func (o *ASMessageDelivery) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *ASMessageDelivery) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *ASMessageDelivery) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *ASMessageDelivery) HasLatency() bool`

HasLatency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


