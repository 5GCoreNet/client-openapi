# SubscriptionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**SubscrCond** | Pointer to [**SubscrCond**](SubscrCond.md) |  | [optional] 

## Methods

### NewSubscriptionContext

`func NewSubscriptionContext(subscriptionId string, ) *SubscriptionContext`

NewSubscriptionContext instantiates a new SubscriptionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionContextWithDefaults

`func NewSubscriptionContextWithDefaults() *SubscriptionContext`

NewSubscriptionContextWithDefaults instantiates a new SubscriptionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *SubscriptionContext) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionContext) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionContext) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSubscrCond

`func (o *SubscriptionContext) GetSubscrCond() SubscrCond`

GetSubscrCond returns the SubscrCond field if non-nil, zero value otherwise.

### GetSubscrCondOk

`func (o *SubscriptionContext) GetSubscrCondOk() (*SubscrCond, bool)`

GetSubscrCondOk returns a tuple with the SubscrCond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscrCond

`func (o *SubscriptionContext) SetSubscrCond(v SubscrCond)`

SetSubscrCond sets SubscrCond field to given value.

### HasSubscrCond

`func (o *SubscriptionContext) HasSubscrCond() bool`

HasSubscrCond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


