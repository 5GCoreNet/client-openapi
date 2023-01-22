# CancelPosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**HgmlcCallBackURI** | **string** | String providing an URI formatted according to RFC 3986. | 
**LdrReference** | **string** | LDR Reference. | 
**ServingLMFIdentification** | Pointer to **string** | LMF identification. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewCancelPosInfo

`func NewCancelPosInfo(supi string, hgmlcCallBackURI string, ldrReference string, ) *CancelPosInfo`

NewCancelPosInfo instantiates a new CancelPosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelPosInfoWithDefaults

`func NewCancelPosInfoWithDefaults() *CancelPosInfo`

NewCancelPosInfoWithDefaults instantiates a new CancelPosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *CancelPosInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *CancelPosInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *CancelPosInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetHgmlcCallBackURI

`func (o *CancelPosInfo) GetHgmlcCallBackURI() string`

GetHgmlcCallBackURI returns the HgmlcCallBackURI field if non-nil, zero value otherwise.

### GetHgmlcCallBackURIOk

`func (o *CancelPosInfo) GetHgmlcCallBackURIOk() (*string, bool)`

GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgmlcCallBackURI

`func (o *CancelPosInfo) SetHgmlcCallBackURI(v string)`

SetHgmlcCallBackURI sets HgmlcCallBackURI field to given value.


### GetLdrReference

`func (o *CancelPosInfo) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *CancelPosInfo) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *CancelPosInfo) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.


### GetServingLMFIdentification

`func (o *CancelPosInfo) GetServingLMFIdentification() string`

GetServingLMFIdentification returns the ServingLMFIdentification field if non-nil, zero value otherwise.

### GetServingLMFIdentificationOk

`func (o *CancelPosInfo) GetServingLMFIdentificationOk() (*string, bool)`

GetServingLMFIdentificationOk returns a tuple with the ServingLMFIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingLMFIdentification

`func (o *CancelPosInfo) SetServingLMFIdentification(v string)`

SetServingLMFIdentification sets ServingLMFIdentification field to given value.

### HasServingLMFIdentification

`func (o *CancelPosInfo) HasServingLMFIdentification() bool`

HasServingLMFIdentification returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *CancelPosInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *CancelPosInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *CancelPosInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *CancelPosInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


