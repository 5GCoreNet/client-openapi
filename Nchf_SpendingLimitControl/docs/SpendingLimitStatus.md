# SpendingLimitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**NotifId** | Pointer to **string** |  | [optional] 
**StatusInfos** | Pointer to [**map[string]PolicyCounterInfo**](PolicyCounterInfo.md) | Status of the requested policy counters. The key of the map is the attribute \&quot;policyCounterId\&quot;.  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSpendingLimitStatus

`func NewSpendingLimitStatus() *SpendingLimitStatus`

NewSpendingLimitStatus instantiates a new SpendingLimitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingLimitStatusWithDefaults

`func NewSpendingLimitStatusWithDefaults() *SpendingLimitStatus`

NewSpendingLimitStatusWithDefaults instantiates a new SpendingLimitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *SpendingLimitStatus) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *SpendingLimitStatus) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *SpendingLimitStatus) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *SpendingLimitStatus) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetNotifId

`func (o *SpendingLimitStatus) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *SpendingLimitStatus) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *SpendingLimitStatus) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.

### HasNotifId

`func (o *SpendingLimitStatus) HasNotifId() bool`

HasNotifId returns a boolean if a field has been set.

### GetStatusInfos

`func (o *SpendingLimitStatus) GetStatusInfos() map[string]PolicyCounterInfo`

GetStatusInfos returns the StatusInfos field if non-nil, zero value otherwise.

### GetStatusInfosOk

`func (o *SpendingLimitStatus) GetStatusInfosOk() (*map[string]PolicyCounterInfo, bool)`

GetStatusInfosOk returns a tuple with the StatusInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfos

`func (o *SpendingLimitStatus) SetStatusInfos(v map[string]PolicyCounterInfo)`

SetStatusInfos sets StatusInfos field to given value.

### HasStatusInfos

`func (o *SpendingLimitStatus) HasStatusInfos() bool`

HasStatusInfos returns a boolean if a field has been set.

### GetExpiry

`func (o *SpendingLimitStatus) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SpendingLimitStatus) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SpendingLimitStatus) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SpendingLimitStatus) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SpendingLimitStatus) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SpendingLimitStatus) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SpendingLimitStatus) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SpendingLimitStatus) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


