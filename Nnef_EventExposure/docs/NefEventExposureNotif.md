# NefEventExposureNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** |  | 
**EventNotifs** | [**[]NefEventNotification**](NefEventNotification.md) |  | 

## Methods

### NewNefEventExposureNotif

`func NewNefEventExposureNotif(notifId string, eventNotifs []NefEventNotification, ) *NefEventExposureNotif`

NewNefEventExposureNotif instantiates a new NefEventExposureNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefEventExposureNotifWithDefaults

`func NewNefEventExposureNotifWithDefaults() *NefEventExposureNotif`

NewNefEventExposureNotifWithDefaults instantiates a new NefEventExposureNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *NefEventExposureNotif) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *NefEventExposureNotif) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *NefEventExposureNotif) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *NefEventExposureNotif) GetEventNotifs() []NefEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *NefEventExposureNotif) GetEventNotifsOk() (*[]NefEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *NefEventExposureNotif) SetEventNotifs(v []NefEventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


