# TestNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 

## Methods

### NewTestNotification

`func NewTestNotification(subscription string, ) *TestNotification`

NewTestNotification instantiates a new TestNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestNotificationWithDefaults

`func NewTestNotificationWithDefaults() *TestNotification`

NewTestNotificationWithDefaults instantiates a new TestNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *TestNotification) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *TestNotification) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *TestNotification) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


