# NdccfAnalyticsSubscriptionNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnaNotifCorrId** | **string** | Notification correlation identifier. | 
**AnaNotifications** | Pointer to [**[]NnwdafEventsSubscriptionNotification**](NnwdafEventsSubscriptionNotification.md) | List of analytics subscription notifications. | [optional] 
**AnaReports** | Pointer to [**[]NotifSummaryReport**](NotifSummaryReport.md) | List of reports with summarized data from multiple analytics notifications that the DCCF has received from NWDAF.  | [optional] 
**FetchInstruct** | Pointer to [**FetchInstruction**](FetchInstruction.md) |  | [optional] 
**TerminationReq** | Pointer to **bool** | It indicates the termination of the data management subscription that requested by the DCCF.  | [optional] 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewNdccfAnalyticsSubscriptionNotification

`func NewNdccfAnalyticsSubscriptionNotification(anaNotifCorrId string, timeStamp time.Time, ) *NdccfAnalyticsSubscriptionNotification`

NewNdccfAnalyticsSubscriptionNotification instantiates a new NdccfAnalyticsSubscriptionNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNdccfAnalyticsSubscriptionNotificationWithDefaults

`func NewNdccfAnalyticsSubscriptionNotificationWithDefaults() *NdccfAnalyticsSubscriptionNotification`

NewNdccfAnalyticsSubscriptionNotificationWithDefaults instantiates a new NdccfAnalyticsSubscriptionNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnaNotifCorrId

`func (o *NdccfAnalyticsSubscriptionNotification) GetAnaNotifCorrId() string`

GetAnaNotifCorrId returns the AnaNotifCorrId field if non-nil, zero value otherwise.

### GetAnaNotifCorrIdOk

`func (o *NdccfAnalyticsSubscriptionNotification) GetAnaNotifCorrIdOk() (*string, bool)`

GetAnaNotifCorrIdOk returns a tuple with the AnaNotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaNotifCorrId

`func (o *NdccfAnalyticsSubscriptionNotification) SetAnaNotifCorrId(v string)`

SetAnaNotifCorrId sets AnaNotifCorrId field to given value.


### GetAnaNotifications

`func (o *NdccfAnalyticsSubscriptionNotification) GetAnaNotifications() []NnwdafEventsSubscriptionNotification`

GetAnaNotifications returns the AnaNotifications field if non-nil, zero value otherwise.

### GetAnaNotificationsOk

`func (o *NdccfAnalyticsSubscriptionNotification) GetAnaNotificationsOk() (*[]NnwdafEventsSubscriptionNotification, bool)`

GetAnaNotificationsOk returns a tuple with the AnaNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaNotifications

`func (o *NdccfAnalyticsSubscriptionNotification) SetAnaNotifications(v []NnwdafEventsSubscriptionNotification)`

SetAnaNotifications sets AnaNotifications field to given value.

### HasAnaNotifications

`func (o *NdccfAnalyticsSubscriptionNotification) HasAnaNotifications() bool`

HasAnaNotifications returns a boolean if a field has been set.

### GetAnaReports

`func (o *NdccfAnalyticsSubscriptionNotification) GetAnaReports() []NotifSummaryReport`

GetAnaReports returns the AnaReports field if non-nil, zero value otherwise.

### GetAnaReportsOk

`func (o *NdccfAnalyticsSubscriptionNotification) GetAnaReportsOk() (*[]NotifSummaryReport, bool)`

GetAnaReportsOk returns a tuple with the AnaReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaReports

`func (o *NdccfAnalyticsSubscriptionNotification) SetAnaReports(v []NotifSummaryReport)`

SetAnaReports sets AnaReports field to given value.

### HasAnaReports

`func (o *NdccfAnalyticsSubscriptionNotification) HasAnaReports() bool`

HasAnaReports returns a boolean if a field has been set.

### GetFetchInstruct

`func (o *NdccfAnalyticsSubscriptionNotification) GetFetchInstruct() FetchInstruction`

GetFetchInstruct returns the FetchInstruct field if non-nil, zero value otherwise.

### GetFetchInstructOk

`func (o *NdccfAnalyticsSubscriptionNotification) GetFetchInstructOk() (*FetchInstruction, bool)`

GetFetchInstructOk returns a tuple with the FetchInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchInstruct

`func (o *NdccfAnalyticsSubscriptionNotification) SetFetchInstruct(v FetchInstruction)`

SetFetchInstruct sets FetchInstruct field to given value.

### HasFetchInstruct

`func (o *NdccfAnalyticsSubscriptionNotification) HasFetchInstruct() bool`

HasFetchInstruct returns a boolean if a field has been set.

### GetTerminationReq

`func (o *NdccfAnalyticsSubscriptionNotification) GetTerminationReq() bool`

GetTerminationReq returns the TerminationReq field if non-nil, zero value otherwise.

### GetTerminationReqOk

`func (o *NdccfAnalyticsSubscriptionNotification) GetTerminationReqOk() (*bool, bool)`

GetTerminationReqOk returns a tuple with the TerminationReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReq

`func (o *NdccfAnalyticsSubscriptionNotification) SetTerminationReq(v bool)`

SetTerminationReq sets TerminationReq field to given value.

### HasTerminationReq

`func (o *NdccfAnalyticsSubscriptionNotification) HasTerminationReq() bool`

HasTerminationReq returns a boolean if a field has been set.

### GetTimeStamp

`func (o *NdccfAnalyticsSubscriptionNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *NdccfAnalyticsSubscriptionNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *NdccfAnalyticsSubscriptionNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


