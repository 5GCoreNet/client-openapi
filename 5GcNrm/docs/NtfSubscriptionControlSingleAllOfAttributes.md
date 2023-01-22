# NtfSubscriptionControlSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationRecipientAddress** | Pointer to **string** |  | [optional] 
**NotificationTypes** | Pointer to [**[]NotificationType1**](NotificationType1.md) |  | [optional] 
**Scope** | Pointer to [**Scope**](Scope.md) |  | [optional] 
**NotificationFilter** | Pointer to **string** | The filter format shall be compliant to XPath 1.0. | [optional] 

## Methods

### NewNtfSubscriptionControlSingleAllOfAttributes

`func NewNtfSubscriptionControlSingleAllOfAttributes() *NtfSubscriptionControlSingleAllOfAttributes`

NewNtfSubscriptionControlSingleAllOfAttributes instantiates a new NtfSubscriptionControlSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtfSubscriptionControlSingleAllOfAttributesWithDefaults

`func NewNtfSubscriptionControlSingleAllOfAttributesWithDefaults() *NtfSubscriptionControlSingleAllOfAttributes`

NewNtfSubscriptionControlSingleAllOfAttributesWithDefaults instantiates a new NtfSubscriptionControlSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationRecipientAddress

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationRecipientAddress() string`

GetNotificationRecipientAddress returns the NotificationRecipientAddress field if non-nil, zero value otherwise.

### GetNotificationRecipientAddressOk

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationRecipientAddressOk() (*string, bool)`

GetNotificationRecipientAddressOk returns a tuple with the NotificationRecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRecipientAddress

`func (o *NtfSubscriptionControlSingleAllOfAttributes) SetNotificationRecipientAddress(v string)`

SetNotificationRecipientAddress sets NotificationRecipientAddress field to given value.

### HasNotificationRecipientAddress

`func (o *NtfSubscriptionControlSingleAllOfAttributes) HasNotificationRecipientAddress() bool`

HasNotificationRecipientAddress returns a boolean if a field has been set.

### GetNotificationTypes

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationTypes() []NotificationType1`

GetNotificationTypes returns the NotificationTypes field if non-nil, zero value otherwise.

### GetNotificationTypesOk

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationTypesOk() (*[]NotificationType1, bool)`

GetNotificationTypesOk returns a tuple with the NotificationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTypes

`func (o *NtfSubscriptionControlSingleAllOfAttributes) SetNotificationTypes(v []NotificationType1)`

SetNotificationTypes sets NotificationTypes field to given value.

### HasNotificationTypes

`func (o *NtfSubscriptionControlSingleAllOfAttributes) HasNotificationTypes() bool`

HasNotificationTypes returns a boolean if a field has been set.

### GetScope

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NtfSubscriptionControlSingleAllOfAttributes) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *NtfSubscriptionControlSingleAllOfAttributes) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetNotificationFilter

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationFilter() string`

GetNotificationFilter returns the NotificationFilter field if non-nil, zero value otherwise.

### GetNotificationFilterOk

`func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationFilterOk() (*string, bool)`

GetNotificationFilterOk returns a tuple with the NotificationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationFilter

`func (o *NtfSubscriptionControlSingleAllOfAttributes) SetNotificationFilter(v string)`

SetNotificationFilter sets NotificationFilter field to given value.

### HasNotificationFilter

`func (o *NtfSubscriptionControlSingleAllOfAttributes) HasNotificationFilter() bool`

HasNotificationFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


