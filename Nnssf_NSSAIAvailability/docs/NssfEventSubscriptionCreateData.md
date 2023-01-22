# NssfEventSubscriptionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfNssaiAvailabilityUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**TaiList** | [**[]Tai**](Tai.md) |  | 
**Event** | [**NssfEventType**](NssfEventType.md) |  | 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AmfSetId** | Pointer to **string** |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**AmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AllAmfSetTaiInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNssfEventSubscriptionCreateData

`func NewNssfEventSubscriptionCreateData(nfNssaiAvailabilityUri string, taiList []Tai, event NssfEventType, ) *NssfEventSubscriptionCreateData`

NewNssfEventSubscriptionCreateData instantiates a new NssfEventSubscriptionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNssfEventSubscriptionCreateDataWithDefaults

`func NewNssfEventSubscriptionCreateDataWithDefaults() *NssfEventSubscriptionCreateData`

NewNssfEventSubscriptionCreateDataWithDefaults instantiates a new NssfEventSubscriptionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfNssaiAvailabilityUri

`func (o *NssfEventSubscriptionCreateData) GetNfNssaiAvailabilityUri() string`

GetNfNssaiAvailabilityUri returns the NfNssaiAvailabilityUri field if non-nil, zero value otherwise.

### GetNfNssaiAvailabilityUriOk

`func (o *NssfEventSubscriptionCreateData) GetNfNssaiAvailabilityUriOk() (*string, bool)`

GetNfNssaiAvailabilityUriOk returns a tuple with the NfNssaiAvailabilityUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfNssaiAvailabilityUri

`func (o *NssfEventSubscriptionCreateData) SetNfNssaiAvailabilityUri(v string)`

SetNfNssaiAvailabilityUri sets NfNssaiAvailabilityUri field to given value.


### GetTaiList

`func (o *NssfEventSubscriptionCreateData) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NssfEventSubscriptionCreateData) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NssfEventSubscriptionCreateData) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.


### GetEvent

`func (o *NssfEventSubscriptionCreateData) GetEvent() NssfEventType`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NssfEventSubscriptionCreateData) GetEventOk() (*NssfEventType, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NssfEventSubscriptionCreateData) SetEvent(v NssfEventType)`

SetEvent sets Event field to given value.


### GetExpiry

`func (o *NssfEventSubscriptionCreateData) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *NssfEventSubscriptionCreateData) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *NssfEventSubscriptionCreateData) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *NssfEventSubscriptionCreateData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetAmfSetId

`func (o *NssfEventSubscriptionCreateData) GetAmfSetId() string`

GetAmfSetId returns the AmfSetId field if non-nil, zero value otherwise.

### GetAmfSetIdOk

`func (o *NssfEventSubscriptionCreateData) GetAmfSetIdOk() (*string, bool)`

GetAmfSetIdOk returns a tuple with the AmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSetId

`func (o *NssfEventSubscriptionCreateData) SetAmfSetId(v string)`

SetAmfSetId sets AmfSetId field to given value.

### HasAmfSetId

`func (o *NssfEventSubscriptionCreateData) HasAmfSetId() bool`

HasAmfSetId returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NssfEventSubscriptionCreateData) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NssfEventSubscriptionCreateData) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NssfEventSubscriptionCreateData) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NssfEventSubscriptionCreateData) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetAmfId

`func (o *NssfEventSubscriptionCreateData) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *NssfEventSubscriptionCreateData) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *NssfEventSubscriptionCreateData) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.

### HasAmfId

`func (o *NssfEventSubscriptionCreateData) HasAmfId() bool`

HasAmfId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NssfEventSubscriptionCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NssfEventSubscriptionCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NssfEventSubscriptionCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NssfEventSubscriptionCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAllAmfSetTaiInd

`func (o *NssfEventSubscriptionCreateData) GetAllAmfSetTaiInd() bool`

GetAllAmfSetTaiInd returns the AllAmfSetTaiInd field if non-nil, zero value otherwise.

### GetAllAmfSetTaiIndOk

`func (o *NssfEventSubscriptionCreateData) GetAllAmfSetTaiIndOk() (*bool, bool)`

GetAllAmfSetTaiIndOk returns a tuple with the AllAmfSetTaiInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAmfSetTaiInd

`func (o *NssfEventSubscriptionCreateData) SetAllAmfSetTaiInd(v bool)`

SetAllAmfSetTaiInd sets AllAmfSetTaiInd field to given value.

### HasAllAmfSetTaiInd

`func (o *NssfEventSubscriptionCreateData) HasAllAmfSetTaiInd() bool`

HasAllAmfSetTaiInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


