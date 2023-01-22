# AmfCreatedEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**AmfEventSubscription**](AmfEventSubscription.md) |  | 
**SubscriptionId** | **string** | String providing an URI formatted according to RFC 3986. | 
**ReportList** | Pointer to [**[]AmfEventReport**](AmfEventReport.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAmfCreatedEventSubscription

`func NewAmfCreatedEventSubscription(subscription AmfEventSubscription, subscriptionId string, ) *AmfCreatedEventSubscription`

NewAmfCreatedEventSubscription instantiates a new AmfCreatedEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfCreatedEventSubscriptionWithDefaults

`func NewAmfCreatedEventSubscriptionWithDefaults() *AmfCreatedEventSubscription`

NewAmfCreatedEventSubscriptionWithDefaults instantiates a new AmfCreatedEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *AmfCreatedEventSubscription) GetSubscription() AmfEventSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AmfCreatedEventSubscription) GetSubscriptionOk() (*AmfEventSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AmfCreatedEventSubscription) SetSubscription(v AmfEventSubscription)`

SetSubscription sets Subscription field to given value.


### GetSubscriptionId

`func (o *AmfCreatedEventSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AmfCreatedEventSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AmfCreatedEventSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetReportList

`func (o *AmfCreatedEventSubscription) GetReportList() []AmfEventReport`

GetReportList returns the ReportList field if non-nil, zero value otherwise.

### GetReportListOk

`func (o *AmfCreatedEventSubscription) GetReportListOk() (*[]AmfEventReport, bool)`

GetReportListOk returns a tuple with the ReportList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportList

`func (o *AmfCreatedEventSubscription) SetReportList(v []AmfEventReport)`

SetReportList sets ReportList field to given value.

### HasReportList

`func (o *AmfCreatedEventSubscription) HasReportList() bool`

HasReportList returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AmfCreatedEventSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AmfCreatedEventSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AmfCreatedEventSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AmfCreatedEventSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


