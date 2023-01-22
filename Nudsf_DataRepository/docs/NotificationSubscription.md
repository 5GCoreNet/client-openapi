# NotificationSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | [**ClientId**](ClientId.md) |  | 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**ExpiryCallbackReference** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ExpiryNotification** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**SubFilter** | Pointer to [**SubscriptionFilter**](SubscriptionFilter.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewNotificationSubscription

`func NewNotificationSubscription(clientId ClientId, callbackReference string, ) *NotificationSubscription`

NewNotificationSubscription instantiates a new NotificationSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSubscriptionWithDefaults

`func NewNotificationSubscriptionWithDefaults() *NotificationSubscription`

NewNotificationSubscriptionWithDefaults instantiates a new NotificationSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *NotificationSubscription) GetClientId() ClientId`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NotificationSubscription) GetClientIdOk() (*ClientId, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NotificationSubscription) SetClientId(v ClientId)`

SetClientId sets ClientId field to given value.


### GetCallbackReference

`func (o *NotificationSubscription) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *NotificationSubscription) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *NotificationSubscription) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetExpiryCallbackReference

`func (o *NotificationSubscription) GetExpiryCallbackReference() string`

GetExpiryCallbackReference returns the ExpiryCallbackReference field if non-nil, zero value otherwise.

### GetExpiryCallbackReferenceOk

`func (o *NotificationSubscription) GetExpiryCallbackReferenceOk() (*string, bool)`

GetExpiryCallbackReferenceOk returns a tuple with the ExpiryCallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryCallbackReference

`func (o *NotificationSubscription) SetExpiryCallbackReference(v string)`

SetExpiryCallbackReference sets ExpiryCallbackReference field to given value.

### HasExpiryCallbackReference

`func (o *NotificationSubscription) HasExpiryCallbackReference() bool`

HasExpiryCallbackReference returns a boolean if a field has been set.

### GetExpiry

`func (o *NotificationSubscription) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *NotificationSubscription) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *NotificationSubscription) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *NotificationSubscription) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetExpiryNotification

`func (o *NotificationSubscription) GetExpiryNotification() int32`

GetExpiryNotification returns the ExpiryNotification field if non-nil, zero value otherwise.

### GetExpiryNotificationOk

`func (o *NotificationSubscription) GetExpiryNotificationOk() (*int32, bool)`

GetExpiryNotificationOk returns a tuple with the ExpiryNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryNotification

`func (o *NotificationSubscription) SetExpiryNotification(v int32)`

SetExpiryNotification sets ExpiryNotification field to given value.

### HasExpiryNotification

`func (o *NotificationSubscription) HasExpiryNotification() bool`

HasExpiryNotification returns a boolean if a field has been set.

### GetSubFilter

`func (o *NotificationSubscription) GetSubFilter() SubscriptionFilter`

GetSubFilter returns the SubFilter field if non-nil, zero value otherwise.

### GetSubFilterOk

`func (o *NotificationSubscription) GetSubFilterOk() (*SubscriptionFilter, bool)`

GetSubFilterOk returns a tuple with the SubFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFilter

`func (o *NotificationSubscription) SetSubFilter(v SubscriptionFilter)`

SetSubFilter sets SubFilter field to given value.

### HasSubFilter

`func (o *NotificationSubscription) HasSubFilter() bool`

HasSubFilter returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NotificationSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NotificationSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NotificationSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NotificationSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


