# ProSeAuthenticationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnrProSe** | Pointer to **string** | Contains the KNR_ProSe. | [optional] 
**Nonce2** | Pointer to **NullableString** | contains an Nonce2 | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewProSeAuthenticationResult

`func NewProSeAuthenticationResult() *ProSeAuthenticationResult`

NewProSeAuthenticationResult instantiates a new ProSeAuthenticationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProSeAuthenticationResultWithDefaults

`func NewProSeAuthenticationResultWithDefaults() *ProSeAuthenticationResult`

NewProSeAuthenticationResultWithDefaults instantiates a new ProSeAuthenticationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnrProSe

`func (o *ProSeAuthenticationResult) GetKnrProSe() string`

GetKnrProSe returns the KnrProSe field if non-nil, zero value otherwise.

### GetKnrProSeOk

`func (o *ProSeAuthenticationResult) GetKnrProSeOk() (*string, bool)`

GetKnrProSeOk returns a tuple with the KnrProSe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnrProSe

`func (o *ProSeAuthenticationResult) SetKnrProSe(v string)`

SetKnrProSe sets KnrProSe field to given value.

### HasKnrProSe

`func (o *ProSeAuthenticationResult) HasKnrProSe() bool`

HasKnrProSe returns a boolean if a field has been set.

### GetNonce2

`func (o *ProSeAuthenticationResult) GetNonce2() string`

GetNonce2 returns the Nonce2 field if non-nil, zero value otherwise.

### GetNonce2Ok

`func (o *ProSeAuthenticationResult) GetNonce2Ok() (*string, bool)`

GetNonce2Ok returns a tuple with the Nonce2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce2

`func (o *ProSeAuthenticationResult) SetNonce2(v string)`

SetNonce2 sets Nonce2 field to given value.

### HasNonce2

`func (o *ProSeAuthenticationResult) HasNonce2() bool`

HasNonce2 returns a boolean if a field has been set.

### SetNonce2Nil

`func (o *ProSeAuthenticationResult) SetNonce2Nil(b bool)`

 SetNonce2Nil sets the value for Nonce2 to be an explicit nil

### UnsetNonce2
`func (o *ProSeAuthenticationResult) UnsetNonce2()`

UnsetNonce2 ensures that no value is present for Nonce2, not even an explicit nil
### GetSupportedFeatures

`func (o *ProSeAuthenticationResult) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ProSeAuthenticationResult) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ProSeAuthenticationResult) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ProSeAuthenticationResult) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


