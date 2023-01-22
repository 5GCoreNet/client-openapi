# NdccfDataSubscriptionNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataNotifCorrId** | **string** | Notification correlation identifier. | 
**DataNotif** | Pointer to [**DataNotification**](DataNotification.md) |  | [optional] 
**DataReports** | Pointer to [**[]NotifSummaryReport**](NotifSummaryReport.md) | List of reports with summarized data from multiple notifications received from data producer.  | [optional] 
**FetchInstruct** | Pointer to [**FetchInstruction**](FetchInstruction.md) |  | [optional] 
**TerminationReq** | Pointer to **bool** | It indicates the termination of the data management subscription that requested by the DCCF.  | [optional] 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewNdccfDataSubscriptionNotification

`func NewNdccfDataSubscriptionNotification(dataNotifCorrId string, timeStamp time.Time, ) *NdccfDataSubscriptionNotification`

NewNdccfDataSubscriptionNotification instantiates a new NdccfDataSubscriptionNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNdccfDataSubscriptionNotificationWithDefaults

`func NewNdccfDataSubscriptionNotificationWithDefaults() *NdccfDataSubscriptionNotification`

NewNdccfDataSubscriptionNotificationWithDefaults instantiates a new NdccfDataSubscriptionNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataNotifCorrId

`func (o *NdccfDataSubscriptionNotification) GetDataNotifCorrId() string`

GetDataNotifCorrId returns the DataNotifCorrId field if non-nil, zero value otherwise.

### GetDataNotifCorrIdOk

`func (o *NdccfDataSubscriptionNotification) GetDataNotifCorrIdOk() (*string, bool)`

GetDataNotifCorrIdOk returns a tuple with the DataNotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNotifCorrId

`func (o *NdccfDataSubscriptionNotification) SetDataNotifCorrId(v string)`

SetDataNotifCorrId sets DataNotifCorrId field to given value.


### GetDataNotif

`func (o *NdccfDataSubscriptionNotification) GetDataNotif() DataNotification`

GetDataNotif returns the DataNotif field if non-nil, zero value otherwise.

### GetDataNotifOk

`func (o *NdccfDataSubscriptionNotification) GetDataNotifOk() (*DataNotification, bool)`

GetDataNotifOk returns a tuple with the DataNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNotif

`func (o *NdccfDataSubscriptionNotification) SetDataNotif(v DataNotification)`

SetDataNotif sets DataNotif field to given value.

### HasDataNotif

`func (o *NdccfDataSubscriptionNotification) HasDataNotif() bool`

HasDataNotif returns a boolean if a field has been set.

### GetDataReports

`func (o *NdccfDataSubscriptionNotification) GetDataReports() []NotifSummaryReport`

GetDataReports returns the DataReports field if non-nil, zero value otherwise.

### GetDataReportsOk

`func (o *NdccfDataSubscriptionNotification) GetDataReportsOk() (*[]NotifSummaryReport, bool)`

GetDataReportsOk returns a tuple with the DataReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReports

`func (o *NdccfDataSubscriptionNotification) SetDataReports(v []NotifSummaryReport)`

SetDataReports sets DataReports field to given value.

### HasDataReports

`func (o *NdccfDataSubscriptionNotification) HasDataReports() bool`

HasDataReports returns a boolean if a field has been set.

### GetFetchInstruct

`func (o *NdccfDataSubscriptionNotification) GetFetchInstruct() FetchInstruction`

GetFetchInstruct returns the FetchInstruct field if non-nil, zero value otherwise.

### GetFetchInstructOk

`func (o *NdccfDataSubscriptionNotification) GetFetchInstructOk() (*FetchInstruction, bool)`

GetFetchInstructOk returns a tuple with the FetchInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchInstruct

`func (o *NdccfDataSubscriptionNotification) SetFetchInstruct(v FetchInstruction)`

SetFetchInstruct sets FetchInstruct field to given value.

### HasFetchInstruct

`func (o *NdccfDataSubscriptionNotification) HasFetchInstruct() bool`

HasFetchInstruct returns a boolean if a field has been set.

### GetTerminationReq

`func (o *NdccfDataSubscriptionNotification) GetTerminationReq() bool`

GetTerminationReq returns the TerminationReq field if non-nil, zero value otherwise.

### GetTerminationReqOk

`func (o *NdccfDataSubscriptionNotification) GetTerminationReqOk() (*bool, bool)`

GetTerminationReqOk returns a tuple with the TerminationReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReq

`func (o *NdccfDataSubscriptionNotification) SetTerminationReq(v bool)`

SetTerminationReq sets TerminationReq field to given value.

### HasTerminationReq

`func (o *NdccfDataSubscriptionNotification) HasTerminationReq() bool`

HasTerminationReq returns a boolean if a field has been set.

### GetTimeStamp

`func (o *NdccfDataSubscriptionNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *NdccfDataSubscriptionNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *NdccfDataSubscriptionNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


