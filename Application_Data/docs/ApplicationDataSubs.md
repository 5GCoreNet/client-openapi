# ApplicationDataSubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**DataFilters** | Pointer to [**[]DataFilter**](DataFilter.md) |  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ImmRep** | Pointer to **bool** | Immediate reporting indication. | [optional] 
**AmInfluEntries** | Pointer to [**[]AmInfluData**](AmInfluData.md) | The AM Influence Data entries stored in the UDR that match a subscription. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApplicationDataSubs

`func NewApplicationDataSubs(notificationUri string, ) *ApplicationDataSubs`

NewApplicationDataSubs instantiates a new ApplicationDataSubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDataSubsWithDefaults

`func NewApplicationDataSubsWithDefaults() *ApplicationDataSubs`

NewApplicationDataSubsWithDefaults instantiates a new ApplicationDataSubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *ApplicationDataSubs) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *ApplicationDataSubs) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *ApplicationDataSubs) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetDataFilters

`func (o *ApplicationDataSubs) GetDataFilters() []DataFilter`

GetDataFilters returns the DataFilters field if non-nil, zero value otherwise.

### GetDataFiltersOk

`func (o *ApplicationDataSubs) GetDataFiltersOk() (*[]DataFilter, bool)`

GetDataFiltersOk returns a tuple with the DataFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFilters

`func (o *ApplicationDataSubs) SetDataFilters(v []DataFilter)`

SetDataFilters sets DataFilters field to given value.

### HasDataFilters

`func (o *ApplicationDataSubs) HasDataFilters() bool`

HasDataFilters returns a boolean if a field has been set.

### GetExpiry

`func (o *ApplicationDataSubs) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ApplicationDataSubs) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ApplicationDataSubs) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ApplicationDataSubs) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetImmRep

`func (o *ApplicationDataSubs) GetImmRep() bool`

GetImmRep returns the ImmRep field if non-nil, zero value otherwise.

### GetImmRepOk

`func (o *ApplicationDataSubs) GetImmRepOk() (*bool, bool)`

GetImmRepOk returns a tuple with the ImmRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmRep

`func (o *ApplicationDataSubs) SetImmRep(v bool)`

SetImmRep sets ImmRep field to given value.

### HasImmRep

`func (o *ApplicationDataSubs) HasImmRep() bool`

HasImmRep returns a boolean if a field has been set.

### GetAmInfluEntries

`func (o *ApplicationDataSubs) GetAmInfluEntries() []AmInfluData`

GetAmInfluEntries returns the AmInfluEntries field if non-nil, zero value otherwise.

### GetAmInfluEntriesOk

`func (o *ApplicationDataSubs) GetAmInfluEntriesOk() (*[]AmInfluData, bool)`

GetAmInfluEntriesOk returns a tuple with the AmInfluEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInfluEntries

`func (o *ApplicationDataSubs) SetAmInfluEntries(v []AmInfluData)`

SetAmInfluEntries sets AmInfluEntries field to given value.

### HasAmInfluEntries

`func (o *ApplicationDataSubs) HasAmInfluEntries() bool`

HasAmInfluEntries returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ApplicationDataSubs) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ApplicationDataSubs) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ApplicationDataSubs) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ApplicationDataSubs) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetResetIds

`func (o *ApplicationDataSubs) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *ApplicationDataSubs) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *ApplicationDataSubs) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *ApplicationDataSubs) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


