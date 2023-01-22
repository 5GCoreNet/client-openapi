# CreatedSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DicEntryId** | **int32** |  | 
**ConfirmedExpires** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewCreatedSubscription

`func NewCreatedSubscription(dicEntryId int32, ) *CreatedSubscription`

NewCreatedSubscription instantiates a new CreatedSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedSubscriptionWithDefaults

`func NewCreatedSubscriptionWithDefaults() *CreatedSubscription`

NewCreatedSubscriptionWithDefaults instantiates a new CreatedSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDicEntryId

`func (o *CreatedSubscription) GetDicEntryId() int32`

GetDicEntryId returns the DicEntryId field if non-nil, zero value otherwise.

### GetDicEntryIdOk

`func (o *CreatedSubscription) GetDicEntryIdOk() (*int32, bool)`

GetDicEntryIdOk returns a tuple with the DicEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDicEntryId

`func (o *CreatedSubscription) SetDicEntryId(v int32)`

SetDicEntryId sets DicEntryId field to given value.


### GetConfirmedExpires

`func (o *CreatedSubscription) GetConfirmedExpires() time.Time`

GetConfirmedExpires returns the ConfirmedExpires field if non-nil, zero value otherwise.

### GetConfirmedExpiresOk

`func (o *CreatedSubscription) GetConfirmedExpiresOk() (*time.Time, bool)`

GetConfirmedExpiresOk returns a tuple with the ConfirmedExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedExpires

`func (o *CreatedSubscription) SetConfirmedExpires(v time.Time)`

SetConfirmedExpires sets ConfirmedExpires field to given value.

### HasConfirmedExpires

`func (o *CreatedSubscription) HasConfirmedExpires() bool`

HasConfirmedExpires returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *CreatedSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *CreatedSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *CreatedSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *CreatedSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


