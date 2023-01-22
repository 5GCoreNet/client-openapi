# MbsPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsPolicyCtxtData** | [**MbsPolicyCtxtData**](MbsPolicyCtxtData.md) |  | 
**MbsPolicies** | Pointer to [**MbsPolicyDecision**](MbsPolicyDecision.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMbsPolicyData

`func NewMbsPolicyData(mbsPolicyCtxtData MbsPolicyCtxtData, ) *MbsPolicyData`

NewMbsPolicyData instantiates a new MbsPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsPolicyDataWithDefaults

`func NewMbsPolicyDataWithDefaults() *MbsPolicyData`

NewMbsPolicyDataWithDefaults instantiates a new MbsPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsPolicyCtxtData

`func (o *MbsPolicyData) GetMbsPolicyCtxtData() MbsPolicyCtxtData`

GetMbsPolicyCtxtData returns the MbsPolicyCtxtData field if non-nil, zero value otherwise.

### GetMbsPolicyCtxtDataOk

`func (o *MbsPolicyData) GetMbsPolicyCtxtDataOk() (*MbsPolicyCtxtData, bool)`

GetMbsPolicyCtxtDataOk returns a tuple with the MbsPolicyCtxtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsPolicyCtxtData

`func (o *MbsPolicyData) SetMbsPolicyCtxtData(v MbsPolicyCtxtData)`

SetMbsPolicyCtxtData sets MbsPolicyCtxtData field to given value.


### GetMbsPolicies

`func (o *MbsPolicyData) GetMbsPolicies() MbsPolicyDecision`

GetMbsPolicies returns the MbsPolicies field if non-nil, zero value otherwise.

### GetMbsPoliciesOk

`func (o *MbsPolicyData) GetMbsPoliciesOk() (*MbsPolicyDecision, bool)`

GetMbsPoliciesOk returns a tuple with the MbsPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsPolicies

`func (o *MbsPolicyData) SetMbsPolicies(v MbsPolicyDecision)`

SetMbsPolicies sets MbsPolicies field to given value.

### HasMbsPolicies

`func (o *MbsPolicyData) HasMbsPolicies() bool`

HasMbsPolicies returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MbsPolicyData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MbsPolicyData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MbsPolicyData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MbsPolicyData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


