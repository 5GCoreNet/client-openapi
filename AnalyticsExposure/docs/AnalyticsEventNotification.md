# AnalyticsEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** |  | 
**AnalyEventNotifs** | [**[]AnalyticsEventNotif**](AnalyticsEventNotif.md) |  | 

## Methods

### NewAnalyticsEventNotification

`func NewAnalyticsEventNotification(notifId string, analyEventNotifs []AnalyticsEventNotif, ) *AnalyticsEventNotification`

NewAnalyticsEventNotification instantiates a new AnalyticsEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsEventNotificationWithDefaults

`func NewAnalyticsEventNotificationWithDefaults() *AnalyticsEventNotification`

NewAnalyticsEventNotificationWithDefaults instantiates a new AnalyticsEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *AnalyticsEventNotification) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *AnalyticsEventNotification) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *AnalyticsEventNotification) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetAnalyEventNotifs

`func (o *AnalyticsEventNotification) GetAnalyEventNotifs() []AnalyticsEventNotif`

GetAnalyEventNotifs returns the AnalyEventNotifs field if non-nil, zero value otherwise.

### GetAnalyEventNotifsOk

`func (o *AnalyticsEventNotification) GetAnalyEventNotifsOk() (*[]AnalyticsEventNotif, bool)`

GetAnalyEventNotifsOk returns a tuple with the AnalyEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyEventNotifs

`func (o *AnalyticsEventNotification) SetAnalyEventNotifs(v []AnalyticsEventNotif)`

SetAnalyEventNotifs sets AnalyEventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


