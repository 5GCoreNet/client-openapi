# NotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NotificationEventType**](NotificationEventType.md) |  | 
**NfInstanceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NfProfile** | Pointer to [**NFProfile**](NFProfile.md) |  | [optional] 
**ProfileChanges** | Pointer to [**[]ChangeItem**](ChangeItem.md) |  | [optional] 
**ConditionEvent** | Pointer to [**ConditionEventType**](ConditionEventType.md) |  | [optional] 
**SubscriptionContext** | Pointer to [**SubscriptionContext**](SubscriptionContext.md) |  | [optional] 

## Methods

### NewNotificationData

`func NewNotificationData(event NotificationEventType, nfInstanceUri string, ) *NotificationData`

NewNotificationData instantiates a new NotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDataWithDefaults

`func NewNotificationDataWithDefaults() *NotificationData`

NewNotificationDataWithDefaults instantiates a new NotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NotificationData) GetEvent() NotificationEventType`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotificationData) GetEventOk() (*NotificationEventType, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotificationData) SetEvent(v NotificationEventType)`

SetEvent sets Event field to given value.


### GetNfInstanceUri

`func (o *NotificationData) GetNfInstanceUri() string`

GetNfInstanceUri returns the NfInstanceUri field if non-nil, zero value otherwise.

### GetNfInstanceUriOk

`func (o *NotificationData) GetNfInstanceUriOk() (*string, bool)`

GetNfInstanceUriOk returns a tuple with the NfInstanceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceUri

`func (o *NotificationData) SetNfInstanceUri(v string)`

SetNfInstanceUri sets NfInstanceUri field to given value.


### GetNfProfile

`func (o *NotificationData) GetNfProfile() NFProfile`

GetNfProfile returns the NfProfile field if non-nil, zero value otherwise.

### GetNfProfileOk

`func (o *NotificationData) GetNfProfileOk() (*NFProfile, bool)`

GetNfProfileOk returns a tuple with the NfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfile

`func (o *NotificationData) SetNfProfile(v NFProfile)`

SetNfProfile sets NfProfile field to given value.

### HasNfProfile

`func (o *NotificationData) HasNfProfile() bool`

HasNfProfile returns a boolean if a field has been set.

### GetProfileChanges

`func (o *NotificationData) GetProfileChanges() []ChangeItem`

GetProfileChanges returns the ProfileChanges field if non-nil, zero value otherwise.

### GetProfileChangesOk

`func (o *NotificationData) GetProfileChangesOk() (*[]ChangeItem, bool)`

GetProfileChangesOk returns a tuple with the ProfileChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileChanges

`func (o *NotificationData) SetProfileChanges(v []ChangeItem)`

SetProfileChanges sets ProfileChanges field to given value.

### HasProfileChanges

`func (o *NotificationData) HasProfileChanges() bool`

HasProfileChanges returns a boolean if a field has been set.

### GetConditionEvent

`func (o *NotificationData) GetConditionEvent() ConditionEventType`

GetConditionEvent returns the ConditionEvent field if non-nil, zero value otherwise.

### GetConditionEventOk

`func (o *NotificationData) GetConditionEventOk() (*ConditionEventType, bool)`

GetConditionEventOk returns a tuple with the ConditionEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionEvent

`func (o *NotificationData) SetConditionEvent(v ConditionEventType)`

SetConditionEvent sets ConditionEvent field to given value.

### HasConditionEvent

`func (o *NotificationData) HasConditionEvent() bool`

HasConditionEvent returns a boolean if a field has been set.

### GetSubscriptionContext

`func (o *NotificationData) GetSubscriptionContext() SubscriptionContext`

GetSubscriptionContext returns the SubscriptionContext field if non-nil, zero value otherwise.

### GetSubscriptionContextOk

`func (o *NotificationData) GetSubscriptionContextOk() (*SubscriptionContext, bool)`

GetSubscriptionContextOk returns a tuple with the SubscriptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionContext

`func (o *NotificationData) SetSubscriptionContext(v SubscriptionContext)`

SetSubscriptionContext sets SubscriptionContext field to given value.

### HasSubscriptionContext

`func (o *NotificationData) HasSubscriptionContext() bool`

HasSubscriptionContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


