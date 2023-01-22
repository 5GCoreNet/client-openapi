# NotificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiredSubscriptions** | [**[]NotificationSubscription**](NotificationSubscription.md) |  | 

## Methods

### NewNotificationInfo

`func NewNotificationInfo(expiredSubscriptions []NotificationSubscription, ) *NotificationInfo`

NewNotificationInfo instantiates a new NotificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationInfoWithDefaults

`func NewNotificationInfoWithDefaults() *NotificationInfo`

NewNotificationInfoWithDefaults instantiates a new NotificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiredSubscriptions

`func (o *NotificationInfo) GetExpiredSubscriptions() []NotificationSubscription`

GetExpiredSubscriptions returns the ExpiredSubscriptions field if non-nil, zero value otherwise.

### GetExpiredSubscriptionsOk

`func (o *NotificationInfo) GetExpiredSubscriptionsOk() (*[]NotificationSubscription, bool)`

GetExpiredSubscriptionsOk returns a tuple with the ExpiredSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredSubscriptions

`func (o *NotificationInfo) SetExpiredSubscriptions(v []NotificationSubscription)`

SetExpiredSubscriptions sets ExpiredSubscriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


