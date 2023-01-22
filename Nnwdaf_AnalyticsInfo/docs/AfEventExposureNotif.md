# AfEventExposureNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** |  | 
**EventNotifs** | [**[]AfEventNotification**](AfEventNotification.md) |  | 

## Methods

### NewAfEventExposureNotif

`func NewAfEventExposureNotif(notifId string, eventNotifs []AfEventNotification, ) *AfEventExposureNotif`

NewAfEventExposureNotif instantiates a new AfEventExposureNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfEventExposureNotifWithDefaults

`func NewAfEventExposureNotifWithDefaults() *AfEventExposureNotif`

NewAfEventExposureNotifWithDefaults instantiates a new AfEventExposureNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *AfEventExposureNotif) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *AfEventExposureNotif) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *AfEventExposureNotif) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *AfEventExposureNotif) GetEventNotifs() []AfEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *AfEventExposureNotif) GetEventNotifsOk() (*[]AfEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *AfEventExposureNotif) SetEventNotifs(v []AfEventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


