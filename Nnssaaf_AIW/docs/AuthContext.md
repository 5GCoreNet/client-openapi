# AuthContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**AuthCtxId** | **string** | contains the resource ID of authentication context | 
**EapMessage** | Pointer to **NullableString** | contains an EAP packet | [optional] 
**TtlsInnerMethodContainer** | Pointer to **NullableString** | contains an EAP packet | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAuthContext

`func NewAuthContext(supi string, authCtxId string, ) *AuthContext`

NewAuthContext instantiates a new AuthContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthContextWithDefaults

`func NewAuthContextWithDefaults() *AuthContext`

NewAuthContextWithDefaults instantiates a new AuthContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *AuthContext) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AuthContext) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AuthContext) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetAuthCtxId

`func (o *AuthContext) GetAuthCtxId() string`

GetAuthCtxId returns the AuthCtxId field if non-nil, zero value otherwise.

### GetAuthCtxIdOk

`func (o *AuthContext) GetAuthCtxIdOk() (*string, bool)`

GetAuthCtxIdOk returns a tuple with the AuthCtxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCtxId

`func (o *AuthContext) SetAuthCtxId(v string)`

SetAuthCtxId sets AuthCtxId field to given value.


### GetEapMessage

`func (o *AuthContext) GetEapMessage() string`

GetEapMessage returns the EapMessage field if non-nil, zero value otherwise.

### GetEapMessageOk

`func (o *AuthContext) GetEapMessageOk() (*string, bool)`

GetEapMessageOk returns a tuple with the EapMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapMessage

`func (o *AuthContext) SetEapMessage(v string)`

SetEapMessage sets EapMessage field to given value.

### HasEapMessage

`func (o *AuthContext) HasEapMessage() bool`

HasEapMessage returns a boolean if a field has been set.

### SetEapMessageNil

`func (o *AuthContext) SetEapMessageNil(b bool)`

 SetEapMessageNil sets the value for EapMessage to be an explicit nil

### UnsetEapMessage
`func (o *AuthContext) UnsetEapMessage()`

UnsetEapMessage ensures that no value is present for EapMessage, not even an explicit nil
### GetTtlsInnerMethodContainer

`func (o *AuthContext) GetTtlsInnerMethodContainer() string`

GetTtlsInnerMethodContainer returns the TtlsInnerMethodContainer field if non-nil, zero value otherwise.

### GetTtlsInnerMethodContainerOk

`func (o *AuthContext) GetTtlsInnerMethodContainerOk() (*string, bool)`

GetTtlsInnerMethodContainerOk returns a tuple with the TtlsInnerMethodContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlsInnerMethodContainer

`func (o *AuthContext) SetTtlsInnerMethodContainer(v string)`

SetTtlsInnerMethodContainer sets TtlsInnerMethodContainer field to given value.

### HasTtlsInnerMethodContainer

`func (o *AuthContext) HasTtlsInnerMethodContainer() bool`

HasTtlsInnerMethodContainer returns a boolean if a field has been set.

### SetTtlsInnerMethodContainerNil

`func (o *AuthContext) SetTtlsInnerMethodContainerNil(b bool)`

 SetTtlsInnerMethodContainerNil sets the value for TtlsInnerMethodContainer to be an explicit nil

### UnsetTtlsInnerMethodContainer
`func (o *AuthContext) UnsetTtlsInnerMethodContainer()`

UnsetTtlsInnerMethodContainer ensures that no value is present for TtlsInnerMethodContainer, not even an explicit nil
### GetSupportedFeatures

`func (o *AuthContext) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthContext) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthContext) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthContext) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


