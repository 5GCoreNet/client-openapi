# MbsPolicyCtxtData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**MbsServInfo** | Pointer to [**MbsServiceInfo**](MbsServiceInfo.md) |  | [optional] 
**NotificationUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMbsPolicyCtxtData

`func NewMbsPolicyCtxtData(mbsSessionId MbsSessionId, ) *MbsPolicyCtxtData`

NewMbsPolicyCtxtData instantiates a new MbsPolicyCtxtData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsPolicyCtxtDataWithDefaults

`func NewMbsPolicyCtxtDataWithDefaults() *MbsPolicyCtxtData`

NewMbsPolicyCtxtDataWithDefaults instantiates a new MbsPolicyCtxtData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *MbsPolicyCtxtData) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MbsPolicyCtxtData) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MbsPolicyCtxtData) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetDnn

`func (o *MbsPolicyCtxtData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *MbsPolicyCtxtData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *MbsPolicyCtxtData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *MbsPolicyCtxtData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *MbsPolicyCtxtData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *MbsPolicyCtxtData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *MbsPolicyCtxtData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *MbsPolicyCtxtData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetMbsServInfo

`func (o *MbsPolicyCtxtData) GetMbsServInfo() MbsServiceInfo`

GetMbsServInfo returns the MbsServInfo field if non-nil, zero value otherwise.

### GetMbsServInfoOk

`func (o *MbsPolicyCtxtData) GetMbsServInfoOk() (*MbsServiceInfo, bool)`

GetMbsServInfoOk returns a tuple with the MbsServInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServInfo

`func (o *MbsPolicyCtxtData) SetMbsServInfo(v MbsServiceInfo)`

SetMbsServInfo sets MbsServInfo field to given value.

### HasMbsServInfo

`func (o *MbsPolicyCtxtData) HasMbsServInfo() bool`

HasMbsServInfo returns a boolean if a field has been set.

### GetNotificationUri

`func (o *MbsPolicyCtxtData) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *MbsPolicyCtxtData) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *MbsPolicyCtxtData) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.

### HasNotificationUri

`func (o *MbsPolicyCtxtData) HasNotificationUri() bool`

HasNotificationUri returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MbsPolicyCtxtData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MbsPolicyCtxtData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MbsPolicyCtxtData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MbsPolicyCtxtData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


