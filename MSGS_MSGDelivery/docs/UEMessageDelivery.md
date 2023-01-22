# UEMessageDelivery

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
**SegInd** | Pointer to **bool** |  | [optional] 
**SegParams** | Pointer to [**MessageSegmentParameters**](MessageSegmentParameters.md) |  | [optional] 
**StoAndFwInd** | **bool** |  | 
**StoAndFwParams** | Pointer to [**StoreAndForwardParameters**](StoreAndForwardParameters.md) |  | [optional] 

## Methods

### NewUEMessageDelivery

`func NewUEMessageDelivery(oriAddr Address, destAddr Address, msgId string, stoAndFwInd bool, ) *UEMessageDelivery`

NewUEMessageDelivery instantiates a new UEMessageDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUEMessageDeliveryWithDefaults

`func NewUEMessageDeliveryWithDefaults() *UEMessageDelivery`

NewUEMessageDeliveryWithDefaults instantiates a new UEMessageDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriAddr

`func (o *UEMessageDelivery) GetOriAddr() Address`

GetOriAddr returns the OriAddr field if non-nil, zero value otherwise.

### GetOriAddrOk

`func (o *UEMessageDelivery) GetOriAddrOk() (*Address, bool)`

GetOriAddrOk returns a tuple with the OriAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriAddr

`func (o *UEMessageDelivery) SetOriAddr(v Address)`

SetOriAddr sets OriAddr field to given value.


### GetDestAddr

`func (o *UEMessageDelivery) GetDestAddr() Address`

GetDestAddr returns the DestAddr field if non-nil, zero value otherwise.

### GetDestAddrOk

`func (o *UEMessageDelivery) GetDestAddrOk() (*Address, bool)`

GetDestAddrOk returns a tuple with the DestAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAddr

`func (o *UEMessageDelivery) SetDestAddr(v Address)`

SetDestAddr sets DestAddr field to given value.


### GetAppId

`func (o *UEMessageDelivery) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UEMessageDelivery) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UEMessageDelivery) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UEMessageDelivery) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetMsgId

`func (o *UEMessageDelivery) GetMsgId() string`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *UEMessageDelivery) GetMsgIdOk() (*string, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *UEMessageDelivery) SetMsgId(v string)`

SetMsgId sets MsgId field to given value.


### GetSecCred

`func (o *UEMessageDelivery) GetSecCred() string`

GetSecCred returns the SecCred field if non-nil, zero value otherwise.

### GetSecCredOk

`func (o *UEMessageDelivery) GetSecCredOk() (*string, bool)`

GetSecCredOk returns a tuple with the SecCred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecCred

`func (o *UEMessageDelivery) SetSecCred(v string)`

SetSecCred sets SecCred field to given value.

### HasSecCred

`func (o *UEMessageDelivery) HasSecCred() bool`

HasSecCred returns a boolean if a field has been set.

### GetDelivStReqInd

`func (o *UEMessageDelivery) GetDelivStReqInd() bool`

GetDelivStReqInd returns the DelivStReqInd field if non-nil, zero value otherwise.

### GetDelivStReqIndOk

`func (o *UEMessageDelivery) GetDelivStReqIndOk() (*bool, bool)`

GetDelivStReqIndOk returns a tuple with the DelivStReqInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivStReqInd

`func (o *UEMessageDelivery) SetDelivStReqInd(v bool)`

SetDelivStReqInd sets DelivStReqInd field to given value.

### HasDelivStReqInd

`func (o *UEMessageDelivery) HasDelivStReqInd() bool`

HasDelivStReqInd returns a boolean if a field has been set.

### GetPayload

`func (o *UEMessageDelivery) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UEMessageDelivery) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UEMessageDelivery) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UEMessageDelivery) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSegInd

`func (o *UEMessageDelivery) GetSegInd() bool`

GetSegInd returns the SegInd field if non-nil, zero value otherwise.

### GetSegIndOk

`func (o *UEMessageDelivery) GetSegIndOk() (*bool, bool)`

GetSegIndOk returns a tuple with the SegInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegInd

`func (o *UEMessageDelivery) SetSegInd(v bool)`

SetSegInd sets SegInd field to given value.

### HasSegInd

`func (o *UEMessageDelivery) HasSegInd() bool`

HasSegInd returns a boolean if a field has been set.

### GetSegParams

`func (o *UEMessageDelivery) GetSegParams() MessageSegmentParameters`

GetSegParams returns the SegParams field if non-nil, zero value otherwise.

### GetSegParamsOk

`func (o *UEMessageDelivery) GetSegParamsOk() (*MessageSegmentParameters, bool)`

GetSegParamsOk returns a tuple with the SegParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegParams

`func (o *UEMessageDelivery) SetSegParams(v MessageSegmentParameters)`

SetSegParams sets SegParams field to given value.

### HasSegParams

`func (o *UEMessageDelivery) HasSegParams() bool`

HasSegParams returns a boolean if a field has been set.

### GetStoAndFwInd

`func (o *UEMessageDelivery) GetStoAndFwInd() bool`

GetStoAndFwInd returns the StoAndFwInd field if non-nil, zero value otherwise.

### GetStoAndFwIndOk

`func (o *UEMessageDelivery) GetStoAndFwIndOk() (*bool, bool)`

GetStoAndFwIndOk returns a tuple with the StoAndFwInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoAndFwInd

`func (o *UEMessageDelivery) SetStoAndFwInd(v bool)`

SetStoAndFwInd sets StoAndFwInd field to given value.


### GetStoAndFwParams

`func (o *UEMessageDelivery) GetStoAndFwParams() StoreAndForwardParameters`

GetStoAndFwParams returns the StoAndFwParams field if non-nil, zero value otherwise.

### GetStoAndFwParamsOk

`func (o *UEMessageDelivery) GetStoAndFwParamsOk() (*StoreAndForwardParameters, bool)`

GetStoAndFwParamsOk returns a tuple with the StoAndFwParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoAndFwParams

`func (o *UEMessageDelivery) SetStoAndFwParams(v StoreAndForwardParameters)`

SetStoAndFwParams sets StoAndFwParams field to given value.

### HasStoAndFwParams

`func (o *UEMessageDelivery) HasStoAndFwParams() bool`

HasStoAndFwParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


