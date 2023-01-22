# NrmEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NrmEvent**](NrmEvent.md) |  | 
**Ts** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**DeliveryMode** | Pointer to [**DeliveryMode**](DeliveryMode.md) |  | [optional] 
**StreamIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNrmEventNotification

`func NewNrmEventNotification(event NrmEvent, ts time.Time, ) *NrmEventNotification`

NewNrmEventNotification instantiates a new NrmEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrmEventNotificationWithDefaults

`func NewNrmEventNotificationWithDefaults() *NrmEventNotification`

NewNrmEventNotificationWithDefaults instantiates a new NrmEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NrmEventNotification) GetEvent() NrmEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NrmEventNotification) GetEventOk() (*NrmEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NrmEventNotification) SetEvent(v NrmEvent)`

SetEvent sets Event field to given value.


### GetTs

`func (o *NrmEventNotification) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *NrmEventNotification) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *NrmEventNotification) SetTs(v time.Time)`

SetTs sets Ts field to given value.


### GetDeliveryMode

`func (o *NrmEventNotification) GetDeliveryMode() DeliveryMode`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *NrmEventNotification) GetDeliveryModeOk() (*DeliveryMode, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *NrmEventNotification) SetDeliveryMode(v DeliveryMode)`

SetDeliveryMode sets DeliveryMode field to given value.

### HasDeliveryMode

`func (o *NrmEventNotification) HasDeliveryMode() bool`

HasDeliveryMode returns a boolean if a field has been set.

### GetStreamIds

`func (o *NrmEventNotification) GetStreamIds() []string`

GetStreamIds returns the StreamIds field if non-nil, zero value otherwise.

### GetStreamIdsOk

`func (o *NrmEventNotification) GetStreamIdsOk() (*[]string, bool)`

GetStreamIdsOk returns a tuple with the StreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamIds

`func (o *NrmEventNotification) SetStreamIds(v []string)`

SetStreamIds sets StreamIds field to given value.

### HasStreamIds

`func (o *NrmEventNotification) HasStreamIds() bool`

HasStreamIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


