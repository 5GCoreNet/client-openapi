# CreatedSACEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**SACEventSubscription**](SACEventSubscription.md) |  | 
**SubscriptionId** | **string** |  | 
**Report** | Pointer to [**SACEventReportItem**](SACEventReportItem.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewCreatedSACEventSubscription

`func NewCreatedSACEventSubscription(subscription SACEventSubscription, subscriptionId string, ) *CreatedSACEventSubscription`

NewCreatedSACEventSubscription instantiates a new CreatedSACEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedSACEventSubscriptionWithDefaults

`func NewCreatedSACEventSubscriptionWithDefaults() *CreatedSACEventSubscription`

NewCreatedSACEventSubscriptionWithDefaults instantiates a new CreatedSACEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *CreatedSACEventSubscription) GetSubscription() SACEventSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CreatedSACEventSubscription) GetSubscriptionOk() (*SACEventSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CreatedSACEventSubscription) SetSubscription(v SACEventSubscription)`

SetSubscription sets Subscription field to given value.


### GetSubscriptionId

`func (o *CreatedSACEventSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreatedSACEventSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreatedSACEventSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetReport

`func (o *CreatedSACEventSubscription) GetReport() SACEventReportItem`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *CreatedSACEventSubscription) GetReportOk() (*SACEventReportItem, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *CreatedSACEventSubscription) SetReport(v SACEventReportItem)`

SetReport sets Report field to given value.

### HasReport

`func (o *CreatedSACEventSubscription) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *CreatedSACEventSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *CreatedSACEventSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *CreatedSACEventSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *CreatedSACEventSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


