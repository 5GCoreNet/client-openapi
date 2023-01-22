# CreateSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**UcmfNotificationUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**SuggestedExpires** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewCreateSubscription

`func NewCreateSubscription(ucmfNotificationUri string, ) *CreateSubscription`

NewCreateSubscription instantiates a new CreateSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionWithDefaults

`func NewCreateSubscriptionWithDefaults() *CreateSubscription`

NewCreateSubscriptionWithDefaults instantiates a new CreateSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfId

`func (o *CreateSubscription) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *CreateSubscription) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *CreateSubscription) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *CreateSubscription) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetUcmfNotificationUri

`func (o *CreateSubscription) GetUcmfNotificationUri() string`

GetUcmfNotificationUri returns the UcmfNotificationUri field if non-nil, zero value otherwise.

### GetUcmfNotificationUriOk

`func (o *CreateSubscription) GetUcmfNotificationUriOk() (*string, bool)`

GetUcmfNotificationUriOk returns a tuple with the UcmfNotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcmfNotificationUri

`func (o *CreateSubscription) SetUcmfNotificationUri(v string)`

SetUcmfNotificationUri sets UcmfNotificationUri field to given value.


### GetSuggestedExpires

`func (o *CreateSubscription) GetSuggestedExpires() time.Time`

GetSuggestedExpires returns the SuggestedExpires field if non-nil, zero value otherwise.

### GetSuggestedExpiresOk

`func (o *CreateSubscription) GetSuggestedExpiresOk() (*time.Time, bool)`

GetSuggestedExpiresOk returns a tuple with the SuggestedExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedExpires

`func (o *CreateSubscription) SetSuggestedExpires(v time.Time)`

SetSuggestedExpires sets SuggestedExpires field to given value.

### HasSuggestedExpires

`func (o *CreateSubscription) HasSuggestedExpires() bool`

HasSuggestedExpires returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *CreateSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *CreateSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *CreateSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *CreateSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


