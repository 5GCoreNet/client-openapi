# NwdafMLModelProvNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotifs** | [**[]MLEventNotif**](MLEventNotif.md) | Notifications about Individual Events. | 
**SubscriptionId** | **string** | String identifying a subscription to the Nnwdaf_MLModelProvision Service. | 

## Methods

### NewNwdafMLModelProvNotif

`func NewNwdafMLModelProvNotif(eventNotifs []MLEventNotif, subscriptionId string, ) *NwdafMLModelProvNotif`

NewNwdafMLModelProvNotif instantiates a new NwdafMLModelProvNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwdafMLModelProvNotifWithDefaults

`func NewNwdafMLModelProvNotifWithDefaults() *NwdafMLModelProvNotif`

NewNwdafMLModelProvNotifWithDefaults instantiates a new NwdafMLModelProvNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotifs

`func (o *NwdafMLModelProvNotif) GetEventNotifs() []MLEventNotif`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *NwdafMLModelProvNotif) GetEventNotifsOk() (*[]MLEventNotif, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *NwdafMLModelProvNotif) SetEventNotifs(v []MLEventNotif)`

SetEventNotifs sets EventNotifs field to given value.


### GetSubscriptionId

`func (o *NwdafMLModelProvNotif) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *NwdafMLModelProvNotif) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *NwdafMLModelProvNotif) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


