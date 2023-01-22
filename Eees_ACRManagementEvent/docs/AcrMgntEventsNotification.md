# AcrMgntEventsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubpId** | **string** | String identifying the Individual ACR Management Events Subscription for which the notification is delivered.  | 
**EventReports** | [**[]AcrMgntEventReport**](AcrMgntEventReport.md) | A list of ACR management event reports. | 

## Methods

### NewAcrMgntEventsNotification

`func NewAcrMgntEventsNotification(subpId string, eventReports []AcrMgntEventReport, ) *AcrMgntEventsNotification`

NewAcrMgntEventsNotification instantiates a new AcrMgntEventsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrMgntEventsNotificationWithDefaults

`func NewAcrMgntEventsNotificationWithDefaults() *AcrMgntEventsNotification`

NewAcrMgntEventsNotificationWithDefaults instantiates a new AcrMgntEventsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubpId

`func (o *AcrMgntEventsNotification) GetSubpId() string`

GetSubpId returns the SubpId field if non-nil, zero value otherwise.

### GetSubpIdOk

`func (o *AcrMgntEventsNotification) GetSubpIdOk() (*string, bool)`

GetSubpIdOk returns a tuple with the SubpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubpId

`func (o *AcrMgntEventsNotification) SetSubpId(v string)`

SetSubpId sets SubpId field to given value.


### GetEventReports

`func (o *AcrMgntEventsNotification) GetEventReports() []AcrMgntEventReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *AcrMgntEventsNotification) GetEventReportsOk() (*[]AcrMgntEventReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *AcrMgntEventsNotification) SetEventReports(v []AcrMgntEventReport)`

SetEventReports sets EventReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


