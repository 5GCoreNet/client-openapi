# SubscriptionDataSubscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | Pointer to **string** | String represents the SUPI or GPSI | [optional] 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**OriginalCallbackReference** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**MonitoredResourceUris** | **[]string** |  | 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SdmSubscription** | Pointer to [**SdmSubscription**](SdmSubscription.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**UniqueSubscription** | Pointer to **bool** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSubscriptionDataSubscriptions

`func NewSubscriptionDataSubscriptions(callbackReference string, monitoredResourceUris []string, ) *SubscriptionDataSubscriptions`

NewSubscriptionDataSubscriptions instantiates a new SubscriptionDataSubscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataSubscriptionsWithDefaults

`func NewSubscriptionDataSubscriptionsWithDefaults() *SubscriptionDataSubscriptions`

NewSubscriptionDataSubscriptionsWithDefaults instantiates a new SubscriptionDataSubscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *SubscriptionDataSubscriptions) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *SubscriptionDataSubscriptions) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *SubscriptionDataSubscriptions) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *SubscriptionDataSubscriptions) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetCallbackReference

`func (o *SubscriptionDataSubscriptions) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *SubscriptionDataSubscriptions) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *SubscriptionDataSubscriptions) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetOriginalCallbackReference

`func (o *SubscriptionDataSubscriptions) GetOriginalCallbackReference() string`

GetOriginalCallbackReference returns the OriginalCallbackReference field if non-nil, zero value otherwise.

### GetOriginalCallbackReferenceOk

`func (o *SubscriptionDataSubscriptions) GetOriginalCallbackReferenceOk() (*string, bool)`

GetOriginalCallbackReferenceOk returns a tuple with the OriginalCallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCallbackReference

`func (o *SubscriptionDataSubscriptions) SetOriginalCallbackReference(v string)`

SetOriginalCallbackReference sets OriginalCallbackReference field to given value.

### HasOriginalCallbackReference

`func (o *SubscriptionDataSubscriptions) HasOriginalCallbackReference() bool`

HasOriginalCallbackReference returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *SubscriptionDataSubscriptions) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *SubscriptionDataSubscriptions) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *SubscriptionDataSubscriptions) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetExpiry

`func (o *SubscriptionDataSubscriptions) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SubscriptionDataSubscriptions) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SubscriptionDataSubscriptions) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SubscriptionDataSubscriptions) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSdmSubscription

`func (o *SubscriptionDataSubscriptions) GetSdmSubscription() SdmSubscription`

GetSdmSubscription returns the SdmSubscription field if non-nil, zero value otherwise.

### GetSdmSubscriptionOk

`func (o *SubscriptionDataSubscriptions) GetSdmSubscriptionOk() (*SdmSubscription, bool)`

GetSdmSubscriptionOk returns a tuple with the SdmSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdmSubscription

`func (o *SubscriptionDataSubscriptions) SetSdmSubscription(v SdmSubscription)`

SetSdmSubscription sets SdmSubscription field to given value.

### HasSdmSubscription

`func (o *SubscriptionDataSubscriptions) HasSdmSubscription() bool`

HasSdmSubscription returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SubscriptionDataSubscriptions) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionDataSubscriptions) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionDataSubscriptions) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SubscriptionDataSubscriptions) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUniqueSubscription

`func (o *SubscriptionDataSubscriptions) GetUniqueSubscription() bool`

GetUniqueSubscription returns the UniqueSubscription field if non-nil, zero value otherwise.

### GetUniqueSubscriptionOk

`func (o *SubscriptionDataSubscriptions) GetUniqueSubscriptionOk() (*bool, bool)`

GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSubscription

`func (o *SubscriptionDataSubscriptions) SetUniqueSubscription(v bool)`

SetUniqueSubscription sets UniqueSubscription field to given value.

### HasUniqueSubscription

`func (o *SubscriptionDataSubscriptions) HasUniqueSubscription() bool`

HasUniqueSubscription returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SubscriptionDataSubscriptions) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SubscriptionDataSubscriptions) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SubscriptionDataSubscriptions) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SubscriptionDataSubscriptions) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


