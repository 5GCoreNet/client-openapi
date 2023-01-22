# EapSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EapPayload** | **NullableString** | contains an EAP packet | 
**KSeaf** | Pointer to **string** | Contains the Kseaf. | [optional] 
**Links** | Pointer to [**map[string]LinksValueSchema**](LinksValueSchema.md) | A map(list of key-value pairs) where member serves as key | [optional] 
**AuthResult** | Pointer to [**AuthResult**](AuthResult.md) |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PvsInfo** | Pointer to [**[]ServerAddressingInfo**](ServerAddressingInfo.md) |  | [optional] 
**Msk** | Pointer to **string** | Contains the Master Session Key. | [optional] 

## Methods

### NewEapSession

`func NewEapSession(eapPayload NullableString, ) *EapSession`

NewEapSession instantiates a new EapSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEapSessionWithDefaults

`func NewEapSessionWithDefaults() *EapSession`

NewEapSessionWithDefaults instantiates a new EapSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEapPayload

`func (o *EapSession) GetEapPayload() string`

GetEapPayload returns the EapPayload field if non-nil, zero value otherwise.

### GetEapPayloadOk

`func (o *EapSession) GetEapPayloadOk() (*string, bool)`

GetEapPayloadOk returns a tuple with the EapPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapPayload

`func (o *EapSession) SetEapPayload(v string)`

SetEapPayload sets EapPayload field to given value.


### SetEapPayloadNil

`func (o *EapSession) SetEapPayloadNil(b bool)`

 SetEapPayloadNil sets the value for EapPayload to be an explicit nil

### UnsetEapPayload
`func (o *EapSession) UnsetEapPayload()`

UnsetEapPayload ensures that no value is present for EapPayload, not even an explicit nil
### GetKSeaf

`func (o *EapSession) GetKSeaf() string`

GetKSeaf returns the KSeaf field if non-nil, zero value otherwise.

### GetKSeafOk

`func (o *EapSession) GetKSeafOk() (*string, bool)`

GetKSeafOk returns a tuple with the KSeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKSeaf

`func (o *EapSession) SetKSeaf(v string)`

SetKSeaf sets KSeaf field to given value.

### HasKSeaf

`func (o *EapSession) HasKSeaf() bool`

HasKSeaf returns a boolean if a field has been set.

### GetLinks

`func (o *EapSession) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EapSession) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EapSession) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EapSession) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAuthResult

`func (o *EapSession) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *EapSession) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *EapSession) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *EapSession) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.

### GetSupi

`func (o *EapSession) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *EapSession) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *EapSession) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *EapSession) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EapSession) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EapSession) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EapSession) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EapSession) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPvsInfo

`func (o *EapSession) GetPvsInfo() []ServerAddressingInfo`

GetPvsInfo returns the PvsInfo field if non-nil, zero value otherwise.

### GetPvsInfoOk

`func (o *EapSession) GetPvsInfoOk() (*[]ServerAddressingInfo, bool)`

GetPvsInfoOk returns a tuple with the PvsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsInfo

`func (o *EapSession) SetPvsInfo(v []ServerAddressingInfo)`

SetPvsInfo sets PvsInfo field to given value.

### HasPvsInfo

`func (o *EapSession) HasPvsInfo() bool`

HasPvsInfo returns a boolean if a field has been set.

### GetMsk

`func (o *EapSession) GetMsk() string`

GetMsk returns the Msk field if non-nil, zero value otherwise.

### GetMskOk

`func (o *EapSession) GetMskOk() (*string, bool)`

GetMskOk returns a tuple with the Msk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsk

`func (o *EapSession) SetMsk(v string)`

SetMsk sets Msk field to given value.

### HasMsk

`func (o *EapSession) HasMsk() bool`

HasMsk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


