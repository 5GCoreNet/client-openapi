# SubscriptionTransferInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransReqType** | [**TransferRequestType**](TransferRequestType.md) |  | 
**NwdafEvSub** | [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | 
**ConsumerId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**ContextId** | Pointer to [**AnalyticsContextIdentifier**](AnalyticsContextIdentifier.md) |  | [optional] 
**SourceNfIds** | Pointer to **[]string** |  | [optional] 
**SourceSetIds** | Pointer to **[]string** |  | [optional] 
**ModelInfo** | Pointer to [**[]ModelInfo**](ModelInfo.md) |  | [optional] 

## Methods

### NewSubscriptionTransferInfo

`func NewSubscriptionTransferInfo(transReqType TransferRequestType, nwdafEvSub NnwdafEventsSubscription, consumerId string, ) *SubscriptionTransferInfo`

NewSubscriptionTransferInfo instantiates a new SubscriptionTransferInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTransferInfoWithDefaults

`func NewSubscriptionTransferInfoWithDefaults() *SubscriptionTransferInfo`

NewSubscriptionTransferInfoWithDefaults instantiates a new SubscriptionTransferInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransReqType

`func (o *SubscriptionTransferInfo) GetTransReqType() TransferRequestType`

GetTransReqType returns the TransReqType field if non-nil, zero value otherwise.

### GetTransReqTypeOk

`func (o *SubscriptionTransferInfo) GetTransReqTypeOk() (*TransferRequestType, bool)`

GetTransReqTypeOk returns a tuple with the TransReqType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransReqType

`func (o *SubscriptionTransferInfo) SetTransReqType(v TransferRequestType)`

SetTransReqType sets TransReqType field to given value.


### GetNwdafEvSub

`func (o *SubscriptionTransferInfo) GetNwdafEvSub() NnwdafEventsSubscription`

GetNwdafEvSub returns the NwdafEvSub field if non-nil, zero value otherwise.

### GetNwdafEvSubOk

`func (o *SubscriptionTransferInfo) GetNwdafEvSubOk() (*NnwdafEventsSubscription, bool)`

GetNwdafEvSubOk returns a tuple with the NwdafEvSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEvSub

`func (o *SubscriptionTransferInfo) SetNwdafEvSub(v NnwdafEventsSubscription)`

SetNwdafEvSub sets NwdafEvSub field to given value.


### GetConsumerId

`func (o *SubscriptionTransferInfo) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *SubscriptionTransferInfo) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *SubscriptionTransferInfo) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.


### GetContextId

`func (o *SubscriptionTransferInfo) GetContextId() AnalyticsContextIdentifier`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *SubscriptionTransferInfo) GetContextIdOk() (*AnalyticsContextIdentifier, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *SubscriptionTransferInfo) SetContextId(v AnalyticsContextIdentifier)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *SubscriptionTransferInfo) HasContextId() bool`

HasContextId returns a boolean if a field has been set.

### GetSourceNfIds

`func (o *SubscriptionTransferInfo) GetSourceNfIds() []string`

GetSourceNfIds returns the SourceNfIds field if non-nil, zero value otherwise.

### GetSourceNfIdsOk

`func (o *SubscriptionTransferInfo) GetSourceNfIdsOk() (*[]string, bool)`

GetSourceNfIdsOk returns a tuple with the SourceNfIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNfIds

`func (o *SubscriptionTransferInfo) SetSourceNfIds(v []string)`

SetSourceNfIds sets SourceNfIds field to given value.

### HasSourceNfIds

`func (o *SubscriptionTransferInfo) HasSourceNfIds() bool`

HasSourceNfIds returns a boolean if a field has been set.

### GetSourceSetIds

`func (o *SubscriptionTransferInfo) GetSourceSetIds() []string`

GetSourceSetIds returns the SourceSetIds field if non-nil, zero value otherwise.

### GetSourceSetIdsOk

`func (o *SubscriptionTransferInfo) GetSourceSetIdsOk() (*[]string, bool)`

GetSourceSetIdsOk returns a tuple with the SourceSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSetIds

`func (o *SubscriptionTransferInfo) SetSourceSetIds(v []string)`

SetSourceSetIds sets SourceSetIds field to given value.

### HasSourceSetIds

`func (o *SubscriptionTransferInfo) HasSourceSetIds() bool`

HasSourceSetIds returns a boolean if a field has been set.

### GetModelInfo

`func (o *SubscriptionTransferInfo) GetModelInfo() []ModelInfo`

GetModelInfo returns the ModelInfo field if non-nil, zero value otherwise.

### GetModelInfoOk

`func (o *SubscriptionTransferInfo) GetModelInfoOk() (*[]ModelInfo, bool)`

GetModelInfoOk returns a tuple with the ModelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelInfo

`func (o *SubscriptionTransferInfo) SetModelInfo(v []ModelInfo)`

SetModelInfo sets ModelInfo field to given value.

### HasModelInfo

`func (o *SubscriptionTransferInfo) HasModelInfo() bool`

HasModelInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


