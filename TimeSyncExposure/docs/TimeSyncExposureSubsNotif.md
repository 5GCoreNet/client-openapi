# TimeSyncExposureSubsNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubsNotifId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**EventNotifs** | [**[]SubsEventNotification**](SubsEventNotification.md) |  | 

## Methods

### NewTimeSyncExposureSubsNotif

`func NewTimeSyncExposureSubsNotif(subsNotifId string, eventNotifs []SubsEventNotification, ) *TimeSyncExposureSubsNotif`

NewTimeSyncExposureSubsNotif instantiates a new TimeSyncExposureSubsNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncExposureSubsNotifWithDefaults

`func NewTimeSyncExposureSubsNotifWithDefaults() *TimeSyncExposureSubsNotif`

NewTimeSyncExposureSubsNotifWithDefaults instantiates a new TimeSyncExposureSubsNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubsNotifId

`func (o *TimeSyncExposureSubsNotif) GetSubsNotifId() string`

GetSubsNotifId returns the SubsNotifId field if non-nil, zero value otherwise.

### GetSubsNotifIdOk

`func (o *TimeSyncExposureSubsNotif) GetSubsNotifIdOk() (*string, bool)`

GetSubsNotifIdOk returns a tuple with the SubsNotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsNotifId

`func (o *TimeSyncExposureSubsNotif) SetSubsNotifId(v string)`

SetSubsNotifId sets SubsNotifId field to given value.


### GetEventNotifs

`func (o *TimeSyncExposureSubsNotif) GetEventNotifs() []SubsEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *TimeSyncExposureSubsNotif) GetEventNotifsOk() (*[]SubsEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *TimeSyncExposureSubsNotif) SetEventNotifs(v []SubsEventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


