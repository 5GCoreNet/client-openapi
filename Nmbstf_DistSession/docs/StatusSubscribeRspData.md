# StatusSubscribeRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**DistSessionSubscription**](DistSessionSubscription.md) |  | 
**ReportList** | Pointer to [**DistSessionEventReportList**](DistSessionEventReportList.md) |  | [optional] 

## Methods

### NewStatusSubscribeRspData

`func NewStatusSubscribeRspData(subscription DistSessionSubscription, ) *StatusSubscribeRspData`

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

`func (o *StatusSubscribeRspData) GetSubscription() DistSessionSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *StatusSubscribeRspData) GetSubscriptionOk() (*DistSessionSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *StatusSubscribeRspData) SetSubscription(v DistSessionSubscription)`

SetSubscription sets Subscription field to given value.


### GetReportList

`func (o *StatusSubscribeRspData) GetReportList() DistSessionEventReportList`

GetReportList returns the ReportList field if non-nil, zero value otherwise.

### GetReportListOk

`func (o *StatusSubscribeRspData) GetReportListOk() (*DistSessionEventReportList, bool)`

GetReportListOk returns a tuple with the ReportList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportList

`func (o *StatusSubscribeRspData) SetReportList(v DistSessionEventReportList)`

SetReportList sets ReportList field to given value.

### HasReportList

`func (o *StatusSubscribeRspData) HasReportList() bool`

HasReportList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


