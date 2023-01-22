# UAVAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AuthMsg** | Pointer to **string** |  | [optional] 
**AuthResult** | Pointer to [**AuthResult**](AuthResult.md) |  | [optional] 
**ServiceLevelId** | Pointer to **string** |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewUAVAuthResponse

`func NewUAVAuthResponse() *UAVAuthResponse`

NewUAVAuthResponse instantiates a new UAVAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUAVAuthResponseWithDefaults

`func NewUAVAuthResponseWithDefaults() *UAVAuthResponse`

NewUAVAuthResponseWithDefaults instantiates a new UAVAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *UAVAuthResponse) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UAVAuthResponse) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UAVAuthResponse) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *UAVAuthResponse) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetAuthMsg

`func (o *UAVAuthResponse) GetAuthMsg() string`

GetAuthMsg returns the AuthMsg field if non-nil, zero value otherwise.

### GetAuthMsgOk

`func (o *UAVAuthResponse) GetAuthMsgOk() (*string, bool)`

GetAuthMsgOk returns a tuple with the AuthMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMsg

`func (o *UAVAuthResponse) SetAuthMsg(v string)`

SetAuthMsg sets AuthMsg field to given value.

### HasAuthMsg

`func (o *UAVAuthResponse) HasAuthMsg() bool`

HasAuthMsg returns a boolean if a field has been set.

### GetAuthResult

`func (o *UAVAuthResponse) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *UAVAuthResponse) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *UAVAuthResponse) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *UAVAuthResponse) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.

### GetServiceLevelId

`func (o *UAVAuthResponse) GetServiceLevelId() string`

GetServiceLevelId returns the ServiceLevelId field if non-nil, zero value otherwise.

### GetServiceLevelIdOk

`func (o *UAVAuthResponse) GetServiceLevelIdOk() (*string, bool)`

GetServiceLevelIdOk returns a tuple with the ServiceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelId

`func (o *UAVAuthResponse) SetServiceLevelId(v string)`

SetServiceLevelId sets ServiceLevelId field to given value.

### HasServiceLevelId

`func (o *UAVAuthResponse) HasServiceLevelId() bool`

HasServiceLevelId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *UAVAuthResponse) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *UAVAuthResponse) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *UAVAuthResponse) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *UAVAuthResponse) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


