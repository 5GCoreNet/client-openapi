# ProSeEapSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EapPayload** | **NullableString** | contains an EAP packet | 
**KnrProSe** | Pointer to **string** | Contains the KNR_ProSe. | [optional] 
**Links** | Pointer to [**map[string]LinksValueSchema**](LinksValueSchema.md) | A map(list of key-value pairs) where member serves as key | [optional] 
**AuthResult** | Pointer to [**AuthResult**](AuthResult.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**Nonce2** | Pointer to **NullableString** | contains an Nonce2 | [optional] 
**Var5gPrukId** | Pointer to **string** | The 5GPRUK ID is string in NAI format as specified in clause 28.7.19 of 3GPP TS 23.003.  | [optional] 

## Methods

### NewProSeEapSession

`func NewProSeEapSession(eapPayload NullableString, ) *ProSeEapSession`

NewProSeEapSession instantiates a new ProSeEapSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProSeEapSessionWithDefaults

`func NewProSeEapSessionWithDefaults() *ProSeEapSession`

NewProSeEapSessionWithDefaults instantiates a new ProSeEapSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEapPayload

`func (o *ProSeEapSession) GetEapPayload() string`

GetEapPayload returns the EapPayload field if non-nil, zero value otherwise.

### GetEapPayloadOk

`func (o *ProSeEapSession) GetEapPayloadOk() (*string, bool)`

GetEapPayloadOk returns a tuple with the EapPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapPayload

`func (o *ProSeEapSession) SetEapPayload(v string)`

SetEapPayload sets EapPayload field to given value.


### SetEapPayloadNil

`func (o *ProSeEapSession) SetEapPayloadNil(b bool)`

 SetEapPayloadNil sets the value for EapPayload to be an explicit nil

### UnsetEapPayload
`func (o *ProSeEapSession) UnsetEapPayload()`

UnsetEapPayload ensures that no value is present for EapPayload, not even an explicit nil
### GetKnrProSe

`func (o *ProSeEapSession) GetKnrProSe() string`

GetKnrProSe returns the KnrProSe field if non-nil, zero value otherwise.

### GetKnrProSeOk

`func (o *ProSeEapSession) GetKnrProSeOk() (*string, bool)`

GetKnrProSeOk returns a tuple with the KnrProSe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnrProSe

`func (o *ProSeEapSession) SetKnrProSe(v string)`

SetKnrProSe sets KnrProSe field to given value.

### HasKnrProSe

`func (o *ProSeEapSession) HasKnrProSe() bool`

HasKnrProSe returns a boolean if a field has been set.

### GetLinks

`func (o *ProSeEapSession) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProSeEapSession) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProSeEapSession) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProSeEapSession) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAuthResult

`func (o *ProSeEapSession) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *ProSeEapSession) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *ProSeEapSession) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *ProSeEapSession) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ProSeEapSession) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ProSeEapSession) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ProSeEapSession) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ProSeEapSession) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetNonce2

`func (o *ProSeEapSession) GetNonce2() string`

GetNonce2 returns the Nonce2 field if non-nil, zero value otherwise.

### GetNonce2Ok

`func (o *ProSeEapSession) GetNonce2Ok() (*string, bool)`

GetNonce2Ok returns a tuple with the Nonce2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce2

`func (o *ProSeEapSession) SetNonce2(v string)`

SetNonce2 sets Nonce2 field to given value.

### HasNonce2

`func (o *ProSeEapSession) HasNonce2() bool`

HasNonce2 returns a boolean if a field has been set.

### SetNonce2Nil

`func (o *ProSeEapSession) SetNonce2Nil(b bool)`

 SetNonce2Nil sets the value for Nonce2 to be an explicit nil

### UnsetNonce2
`func (o *ProSeEapSession) UnsetNonce2()`

UnsetNonce2 ensures that no value is present for Nonce2, not even an explicit nil
### GetVar5gPrukId

`func (o *ProSeEapSession) GetVar5gPrukId() string`

GetVar5gPrukId returns the Var5gPrukId field if non-nil, zero value otherwise.

### GetVar5gPrukIdOk

`func (o *ProSeEapSession) GetVar5gPrukIdOk() (*string, bool)`

GetVar5gPrukIdOk returns a tuple with the Var5gPrukId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gPrukId

`func (o *ProSeEapSession) SetVar5gPrukId(v string)`

SetVar5gPrukId sets Var5gPrukId field to given value.

### HasVar5gPrukId

`func (o *ProSeEapSession) HasVar5gPrukId() bool`

HasVar5gPrukId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


