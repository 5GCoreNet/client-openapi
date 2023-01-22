# NadrfDataRetrievalNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifCorrId** | **string** | Notification correlation identifier. | 
**AnaNotifications** | Pointer to [**[]NnwdafEventsSubscriptionNotification**](NnwdafEventsSubscriptionNotification.md) | List of analytics subscription notifications. | [optional] 
**DataNotif** | Pointer to [**DataNotification**](DataNotification.md) |  | [optional] 
**FetchInstruct** | Pointer to [**FetchInstruction**](FetchInstruction.md) |  | [optional] 
**TerminationReq** | Pointer to **bool** | It indicates the termination of the data management subscription that requested by the ADRF.  | [optional] 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewNadrfDataRetrievalNotification

`func NewNadrfDataRetrievalNotification(notifCorrId string, timeStamp time.Time, ) *NadrfDataRetrievalNotification`

NewNadrfDataRetrievalNotification instantiates a new NadrfDataRetrievalNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNadrfDataRetrievalNotificationWithDefaults

`func NewNadrfDataRetrievalNotificationWithDefaults() *NadrfDataRetrievalNotification`

NewNadrfDataRetrievalNotificationWithDefaults instantiates a new NadrfDataRetrievalNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifCorrId

`func (o *NadrfDataRetrievalNotification) GetNotifCorrId() string`

GetNotifCorrId returns the NotifCorrId field if non-nil, zero value otherwise.

### GetNotifCorrIdOk

`func (o *NadrfDataRetrievalNotification) GetNotifCorrIdOk() (*string, bool)`

GetNotifCorrIdOk returns a tuple with the NotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorrId

`func (o *NadrfDataRetrievalNotification) SetNotifCorrId(v string)`

SetNotifCorrId sets NotifCorrId field to given value.


### GetAnaNotifications

`func (o *NadrfDataRetrievalNotification) GetAnaNotifications() []NnwdafEventsSubscriptionNotification`

GetAnaNotifications returns the AnaNotifications field if non-nil, zero value otherwise.

### GetAnaNotificationsOk

`func (o *NadrfDataRetrievalNotification) GetAnaNotificationsOk() (*[]NnwdafEventsSubscriptionNotification, bool)`

GetAnaNotificationsOk returns a tuple with the AnaNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaNotifications

`func (o *NadrfDataRetrievalNotification) SetAnaNotifications(v []NnwdafEventsSubscriptionNotification)`

SetAnaNotifications sets AnaNotifications field to given value.

### HasAnaNotifications

`func (o *NadrfDataRetrievalNotification) HasAnaNotifications() bool`

HasAnaNotifications returns a boolean if a field has been set.

### GetDataNotif

`func (o *NadrfDataRetrievalNotification) GetDataNotif() DataNotification`

GetDataNotif returns the DataNotif field if non-nil, zero value otherwise.

### GetDataNotifOk

`func (o *NadrfDataRetrievalNotification) GetDataNotifOk() (*DataNotification, bool)`

GetDataNotifOk returns a tuple with the DataNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNotif

`func (o *NadrfDataRetrievalNotification) SetDataNotif(v DataNotification)`

SetDataNotif sets DataNotif field to given value.

### HasDataNotif

`func (o *NadrfDataRetrievalNotification) HasDataNotif() bool`

HasDataNotif returns a boolean if a field has been set.

### GetFetchInstruct

`func (o *NadrfDataRetrievalNotification) GetFetchInstruct() FetchInstruction`

GetFetchInstruct returns the FetchInstruct field if non-nil, zero value otherwise.

### GetFetchInstructOk

`func (o *NadrfDataRetrievalNotification) GetFetchInstructOk() (*FetchInstruction, bool)`

GetFetchInstructOk returns a tuple with the FetchInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchInstruct

`func (o *NadrfDataRetrievalNotification) SetFetchInstruct(v FetchInstruction)`

SetFetchInstruct sets FetchInstruct field to given value.

### HasFetchInstruct

`func (o *NadrfDataRetrievalNotification) HasFetchInstruct() bool`

HasFetchInstruct returns a boolean if a field has been set.

### GetTerminationReq

`func (o *NadrfDataRetrievalNotification) GetTerminationReq() bool`

GetTerminationReq returns the TerminationReq field if non-nil, zero value otherwise.

### GetTerminationReqOk

`func (o *NadrfDataRetrievalNotification) GetTerminationReqOk() (*bool, bool)`

GetTerminationReqOk returns a tuple with the TerminationReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReq

`func (o *NadrfDataRetrievalNotification) SetTerminationReq(v bool)`

SetTerminationReq sets TerminationReq field to given value.

### HasTerminationReq

`func (o *NadrfDataRetrievalNotification) HasTerminationReq() bool`

HasTerminationReq returns a boolean if a field has been set.

### GetTimeStamp

`func (o *NadrfDataRetrievalNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *NadrfDataRetrievalNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *NadrfDataRetrievalNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


