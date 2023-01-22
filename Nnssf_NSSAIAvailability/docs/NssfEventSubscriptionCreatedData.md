# NssfEventSubscriptionCreatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AuthorizedNssaiAvailabilityData** | Pointer to [**[]AuthorizedNssaiAvailabilityData**](AuthorizedNssaiAvailabilityData.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewNssfEventSubscriptionCreatedData

`func NewNssfEventSubscriptionCreatedData(subscriptionId string, ) *NssfEventSubscriptionCreatedData`

NewNssfEventSubscriptionCreatedData instantiates a new NssfEventSubscriptionCreatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNssfEventSubscriptionCreatedDataWithDefaults

`func NewNssfEventSubscriptionCreatedDataWithDefaults() *NssfEventSubscriptionCreatedData`

NewNssfEventSubscriptionCreatedDataWithDefaults instantiates a new NssfEventSubscriptionCreatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *NssfEventSubscriptionCreatedData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *NssfEventSubscriptionCreatedData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *NssfEventSubscriptionCreatedData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetExpiry

`func (o *NssfEventSubscriptionCreatedData) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *NssfEventSubscriptionCreatedData) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *NssfEventSubscriptionCreatedData) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *NssfEventSubscriptionCreatedData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetAuthorizedNssaiAvailabilityData

`func (o *NssfEventSubscriptionCreatedData) GetAuthorizedNssaiAvailabilityData() []AuthorizedNssaiAvailabilityData`

GetAuthorizedNssaiAvailabilityData returns the AuthorizedNssaiAvailabilityData field if non-nil, zero value otherwise.

### GetAuthorizedNssaiAvailabilityDataOk

`func (o *NssfEventSubscriptionCreatedData) GetAuthorizedNssaiAvailabilityDataOk() (*[]AuthorizedNssaiAvailabilityData, bool)`

GetAuthorizedNssaiAvailabilityDataOk returns a tuple with the AuthorizedNssaiAvailabilityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedNssaiAvailabilityData

`func (o *NssfEventSubscriptionCreatedData) SetAuthorizedNssaiAvailabilityData(v []AuthorizedNssaiAvailabilityData)`

SetAuthorizedNssaiAvailabilityData sets AuthorizedNssaiAvailabilityData field to given value.

### HasAuthorizedNssaiAvailabilityData

`func (o *NssfEventSubscriptionCreatedData) HasAuthorizedNssaiAvailabilityData() bool`

HasAuthorizedNssaiAvailabilityData returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NssfEventSubscriptionCreatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NssfEventSubscriptionCreatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NssfEventSubscriptionCreatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NssfEventSubscriptionCreatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


