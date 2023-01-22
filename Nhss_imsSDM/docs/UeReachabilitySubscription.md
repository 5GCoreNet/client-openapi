# UeReachabilitySubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 

## Methods

### NewUeReachabilitySubscription

`func NewUeReachabilitySubscription(expiry time.Time, callbackReference string, ) *UeReachabilitySubscription`

NewUeReachabilitySubscription instantiates a new UeReachabilitySubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeReachabilitySubscriptionWithDefaults

`func NewUeReachabilitySubscriptionWithDefaults() *UeReachabilitySubscription`

NewUeReachabilitySubscriptionWithDefaults instantiates a new UeReachabilitySubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *UeReachabilitySubscription) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *UeReachabilitySubscription) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *UeReachabilitySubscription) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.


### GetCallbackReference

`func (o *UeReachabilitySubscription) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *UeReachabilitySubscription) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *UeReachabilitySubscription) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


