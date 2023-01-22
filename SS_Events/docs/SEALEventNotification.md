# SEALEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Identifier of the subscription resource. | 
**EventDetails** | [**[]SEALEventDetail**](SEALEventDetail.md) | Detailed notifications of individual events. | 

## Methods

### NewSEALEventNotification

`func NewSEALEventNotification(subscriptionId string, eventDetails []SEALEventDetail, ) *SEALEventNotification`

NewSEALEventNotification instantiates a new SEALEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSEALEventNotificationWithDefaults

`func NewSEALEventNotificationWithDefaults() *SEALEventNotification`

NewSEALEventNotificationWithDefaults instantiates a new SEALEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *SEALEventNotification) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SEALEventNotification) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SEALEventNotification) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetEventDetails

`func (o *SEALEventNotification) GetEventDetails() []SEALEventDetail`

GetEventDetails returns the EventDetails field if non-nil, zero value otherwise.

### GetEventDetailsOk

`func (o *SEALEventNotification) GetEventDetailsOk() (*[]SEALEventDetail, bool)`

GetEventDetailsOk returns a tuple with the EventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetails

`func (o *SEALEventNotification) SetEventDetails(v []SEALEventDetail)`

SetEventDetails sets EventDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


