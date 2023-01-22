# StatusSubscribeReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**MbsSessionSubscription**](MbsSessionSubscription.md) |  | 

## Methods

### NewStatusSubscribeReqData

`func NewStatusSubscribeReqData(subscription MbsSessionSubscription, ) *StatusSubscribeReqData`

NewStatusSubscribeReqData instantiates a new StatusSubscribeReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSubscribeReqDataWithDefaults

`func NewStatusSubscribeReqDataWithDefaults() *StatusSubscribeReqData`

NewStatusSubscribeReqDataWithDefaults instantiates a new StatusSubscribeReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *StatusSubscribeReqData) GetSubscription() MbsSessionSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *StatusSubscribeReqData) GetSubscriptionOk() (*MbsSessionSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *StatusSubscribeReqData) SetSubscription(v MbsSessionSubscription)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


