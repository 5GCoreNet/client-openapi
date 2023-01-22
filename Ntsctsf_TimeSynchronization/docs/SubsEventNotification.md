# SubsEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**SubscribedEvent**](SubscribedEvent.md) |  | [optional] 
**TimeSyncCapas** | Pointer to [**[]TimeSyncCapability**](TimeSyncCapability.md) |  | [optional] 

## Methods

### NewSubsEventNotification

`func NewSubsEventNotification() *SubsEventNotification`

NewSubsEventNotification instantiates a new SubsEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubsEventNotificationWithDefaults

`func NewSubsEventNotificationWithDefaults() *SubsEventNotification`

NewSubsEventNotificationWithDefaults instantiates a new SubsEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *SubsEventNotification) GetEvent() SubscribedEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SubsEventNotification) GetEventOk() (*SubscribedEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SubsEventNotification) SetEvent(v SubscribedEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *SubsEventNotification) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTimeSyncCapas

`func (o *SubsEventNotification) GetTimeSyncCapas() []TimeSyncCapability`

GetTimeSyncCapas returns the TimeSyncCapas field if non-nil, zero value otherwise.

### GetTimeSyncCapasOk

`func (o *SubsEventNotification) GetTimeSyncCapasOk() (*[]TimeSyncCapability, bool)`

GetTimeSyncCapasOk returns a tuple with the TimeSyncCapas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSyncCapas

`func (o *SubsEventNotification) SetTimeSyncCapas(v []TimeSyncCapability)`

SetTimeSyncCapas sets TimeSyncCapas field to given value.

### HasTimeSyncCapas

`func (o *SubsEventNotification) HasTimeSyncCapas() bool`

HasTimeSyncCapas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


