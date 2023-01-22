# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusEvent** | [**Event**](Event.md) |  | 
**MbsDisSessionId** | Pointer to **string** |  | [optional] 
**StatusAddInfo** | Pointer to **string** |  | [optional] 
**TimeStamp** | **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 

## Methods

### NewEventNotification

`func NewEventNotification(statusEvent Event, timeStamp time.Time, ) *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusEvent

`func (o *EventNotification) GetStatusEvent() Event`

GetStatusEvent returns the StatusEvent field if non-nil, zero value otherwise.

### GetStatusEventOk

`func (o *EventNotification) GetStatusEventOk() (*Event, bool)`

GetStatusEventOk returns a tuple with the StatusEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEvent

`func (o *EventNotification) SetStatusEvent(v Event)`

SetStatusEvent sets StatusEvent field to given value.


### GetMbsDisSessionId

`func (o *EventNotification) GetMbsDisSessionId() string`

GetMbsDisSessionId returns the MbsDisSessionId field if non-nil, zero value otherwise.

### GetMbsDisSessionIdOk

`func (o *EventNotification) GetMbsDisSessionIdOk() (*string, bool)`

GetMbsDisSessionIdOk returns a tuple with the MbsDisSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDisSessionId

`func (o *EventNotification) SetMbsDisSessionId(v string)`

SetMbsDisSessionId sets MbsDisSessionId field to given value.

### HasMbsDisSessionId

`func (o *EventNotification) HasMbsDisSessionId() bool`

HasMbsDisSessionId returns a boolean if a field has been set.

### GetStatusAddInfo

`func (o *EventNotification) GetStatusAddInfo() string`

GetStatusAddInfo returns the StatusAddInfo field if non-nil, zero value otherwise.

### GetStatusAddInfoOk

`func (o *EventNotification) GetStatusAddInfoOk() (*string, bool)`

GetStatusAddInfoOk returns a tuple with the StatusAddInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAddInfo

`func (o *EventNotification) SetStatusAddInfo(v string)`

SetStatusAddInfo sets StatusAddInfo field to given value.

### HasStatusAddInfo

`func (o *EventNotification) HasStatusAddInfo() bool`

HasStatusAddInfo returns a boolean if a field has been set.

### GetTimeStamp

`func (o *EventNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *EventNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *EventNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


