# StatusSubscribeRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**MbsSessionSubscription**](MbsSessionSubscription.md) |  | 
**EventList** | Pointer to [**MbsSessionEventReportList**](MbsSessionEventReportList.md) |  | [optional] 

## Methods

### NewStatusSubscribeRspData

`func NewStatusSubscribeRspData(subscription MbsSessionSubscription, ) *StatusSubscribeRspData`

NewStatusSubscribeRspData instantiates a new StatusSubscribeRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSubscribeRspDataWithDefaults

`func NewStatusSubscribeRspDataWithDefaults() *StatusSubscribeRspData`

NewStatusSubscribeRspDataWithDefaults instantiates a new StatusSubscribeRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *StatusSubscribeRspData) GetSubscription() MbsSessionSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *StatusSubscribeRspData) GetSubscriptionOk() (*MbsSessionSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *StatusSubscribeRspData) SetSubscription(v MbsSessionSubscription)`

SetSubscription sets Subscription field to given value.


### GetEventList

`func (o *StatusSubscribeRspData) GetEventList() MbsSessionEventReportList`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *StatusSubscribeRspData) GetEventListOk() (*MbsSessionEventReportList, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *StatusSubscribeRspData) SetEventList(v MbsSessionEventReportList)`

SetEventList sets EventList field to given value.

### HasEventList

`func (o *StatusSubscribeRspData) HasEventList() bool`

HasEventList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


