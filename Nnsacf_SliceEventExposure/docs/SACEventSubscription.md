# SACEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**SACEvent**](SACEvent.md) |  | 
**EventNotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NfId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 
**MaxReports** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSACEventSubscription

`func NewSACEventSubscription(event SACEvent, eventNotifyUri string, nfId string, ) *SACEventSubscription`

NewSACEventSubscription instantiates a new SACEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSACEventSubscriptionWithDefaults

`func NewSACEventSubscriptionWithDefaults() *SACEventSubscription`

NewSACEventSubscriptionWithDefaults instantiates a new SACEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *SACEventSubscription) GetEvent() SACEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SACEventSubscription) GetEventOk() (*SACEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SACEventSubscription) SetEvent(v SACEvent)`

SetEvent sets Event field to given value.


### GetEventNotifyUri

`func (o *SACEventSubscription) GetEventNotifyUri() string`

GetEventNotifyUri returns the EventNotifyUri field if non-nil, zero value otherwise.

### GetEventNotifyUriOk

`func (o *SACEventSubscription) GetEventNotifyUriOk() (*string, bool)`

GetEventNotifyUriOk returns a tuple with the EventNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyUri

`func (o *SACEventSubscription) SetEventNotifyUri(v string)`

SetEventNotifyUri sets EventNotifyUri field to given value.


### GetNfId

`func (o *SACEventSubscription) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *SACEventSubscription) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *SACEventSubscription) SetNfId(v string)`

SetNfId sets NfId field to given value.


### GetNotifyCorrelationId

`func (o *SACEventSubscription) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *SACEventSubscription) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *SACEventSubscription) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *SACEventSubscription) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetMaxReports

`func (o *SACEventSubscription) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *SACEventSubscription) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *SACEventSubscription) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *SACEventSubscription) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetExpiry

`func (o *SACEventSubscription) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SACEventSubscription) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SACEventSubscription) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SACEventSubscription) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SACEventSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SACEventSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SACEventSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SACEventSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


