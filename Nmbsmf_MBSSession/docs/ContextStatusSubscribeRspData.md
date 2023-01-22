# ContextStatusSubscribeRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**ContextStatusSubscription**](ContextStatusSubscription.md) |  | 
**ReportList** | Pointer to [**[]ContextStatusEventReport**](ContextStatusEventReport.md) |  | [optional] 
**MbsContextInfo** | Pointer to [**MbsContextInfo**](MbsContextInfo.md) |  | [optional] 

## Methods

### NewContextStatusSubscribeRspData

`func NewContextStatusSubscribeRspData(subscription ContextStatusSubscription, ) *ContextStatusSubscribeRspData`

NewContextStatusSubscribeRspData instantiates a new ContextStatusSubscribeRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextStatusSubscribeRspDataWithDefaults

`func NewContextStatusSubscribeRspDataWithDefaults() *ContextStatusSubscribeRspData`

NewContextStatusSubscribeRspDataWithDefaults instantiates a new ContextStatusSubscribeRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *ContextStatusSubscribeRspData) GetSubscription() ContextStatusSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ContextStatusSubscribeRspData) GetSubscriptionOk() (*ContextStatusSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ContextStatusSubscribeRspData) SetSubscription(v ContextStatusSubscription)`

SetSubscription sets Subscription field to given value.


### GetReportList

`func (o *ContextStatusSubscribeRspData) GetReportList() []ContextStatusEventReport`

GetReportList returns the ReportList field if non-nil, zero value otherwise.

### GetReportListOk

`func (o *ContextStatusSubscribeRspData) GetReportListOk() (*[]ContextStatusEventReport, bool)`

GetReportListOk returns a tuple with the ReportList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportList

`func (o *ContextStatusSubscribeRspData) SetReportList(v []ContextStatusEventReport)`

SetReportList sets ReportList field to given value.

### HasReportList

`func (o *ContextStatusSubscribeRspData) HasReportList() bool`

HasReportList returns a boolean if a field has been set.

### GetMbsContextInfo

`func (o *ContextStatusSubscribeRspData) GetMbsContextInfo() MbsContextInfo`

GetMbsContextInfo returns the MbsContextInfo field if non-nil, zero value otherwise.

### GetMbsContextInfoOk

`func (o *ContextStatusSubscribeRspData) GetMbsContextInfoOk() (*MbsContextInfo, bool)`

GetMbsContextInfoOk returns a tuple with the MbsContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsContextInfo

`func (o *ContextStatusSubscribeRspData) SetMbsContextInfo(v MbsContextInfo)`

SetMbsContextInfo sets MbsContextInfo field to given value.

### HasMbsContextInfo

`func (o *ContextStatusSubscribeRspData) HasMbsContextInfo() bool`

HasMbsContextInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


