# NadrfDataRetrievalSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnaSub** | Pointer to [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | [optional] 
**DataSub** | Pointer to [**DataSubscription**](DataSubscription.md) |  | [optional] 
**NotificationURI** | **string** | String providing an URI formatted according to RFC 3986. | 
**TimePeriod** | [**TimeWindow**](TimeWindow.md) |  | 
**NotifCorrId** | **string** | Notification correlation identifier. | 

## Methods

### NewNadrfDataRetrievalSubscription

`func NewNadrfDataRetrievalSubscription(notificationURI string, timePeriod TimeWindow, notifCorrId string, ) *NadrfDataRetrievalSubscription`

NewNadrfDataRetrievalSubscription instantiates a new NadrfDataRetrievalSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNadrfDataRetrievalSubscriptionWithDefaults

`func NewNadrfDataRetrievalSubscriptionWithDefaults() *NadrfDataRetrievalSubscription`

NewNadrfDataRetrievalSubscriptionWithDefaults instantiates a new NadrfDataRetrievalSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnaSub

`func (o *NadrfDataRetrievalSubscription) GetAnaSub() NnwdafEventsSubscription`

GetAnaSub returns the AnaSub field if non-nil, zero value otherwise.

### GetAnaSubOk

`func (o *NadrfDataRetrievalSubscription) GetAnaSubOk() (*NnwdafEventsSubscription, bool)`

GetAnaSubOk returns a tuple with the AnaSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaSub

`func (o *NadrfDataRetrievalSubscription) SetAnaSub(v NnwdafEventsSubscription)`

SetAnaSub sets AnaSub field to given value.

### HasAnaSub

`func (o *NadrfDataRetrievalSubscription) HasAnaSub() bool`

HasAnaSub returns a boolean if a field has been set.

### GetDataSub

`func (o *NadrfDataRetrievalSubscription) GetDataSub() DataSubscription`

GetDataSub returns the DataSub field if non-nil, zero value otherwise.

### GetDataSubOk

`func (o *NadrfDataRetrievalSubscription) GetDataSubOk() (*DataSubscription, bool)`

GetDataSubOk returns a tuple with the DataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSub

`func (o *NadrfDataRetrievalSubscription) SetDataSub(v DataSubscription)`

SetDataSub sets DataSub field to given value.

### HasDataSub

`func (o *NadrfDataRetrievalSubscription) HasDataSub() bool`

HasDataSub returns a boolean if a field has been set.

### GetNotificationURI

`func (o *NadrfDataRetrievalSubscription) GetNotificationURI() string`

GetNotificationURI returns the NotificationURI field if non-nil, zero value otherwise.

### GetNotificationURIOk

`func (o *NadrfDataRetrievalSubscription) GetNotificationURIOk() (*string, bool)`

GetNotificationURIOk returns a tuple with the NotificationURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationURI

`func (o *NadrfDataRetrievalSubscription) SetNotificationURI(v string)`

SetNotificationURI sets NotificationURI field to given value.


### GetTimePeriod

`func (o *NadrfDataRetrievalSubscription) GetTimePeriod() TimeWindow`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *NadrfDataRetrievalSubscription) GetTimePeriodOk() (*TimeWindow, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *NadrfDataRetrievalSubscription) SetTimePeriod(v TimeWindow)`

SetTimePeriod sets TimePeriod field to given value.


### GetNotifCorrId

`func (o *NadrfDataRetrievalSubscription) GetNotifCorrId() string`

GetNotifCorrId returns the NotifCorrId field if non-nil, zero value otherwise.

### GetNotifCorrIdOk

`func (o *NadrfDataRetrievalSubscription) GetNotifCorrIdOk() (*string, bool)`

GetNotifCorrIdOk returns a tuple with the NotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorrId

`func (o *NadrfDataRetrievalSubscription) SetNotifCorrId(v string)`

SetNotifCorrId sets NotifCorrId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


