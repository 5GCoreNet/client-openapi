# AuthenticationInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impi** | Pointer to **string** | IMS Private Identity of the UE | [optional] 
**Var3gAkaAv** | Pointer to [**Model3GAkaAv**](Model3GAkaAv.md) |  | [optional] 
**DigestAuth** | Pointer to [**DigestAuthentication**](DigestAuthentication.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAuthenticationInfoResult

`func NewAuthenticationInfoResult() *AuthenticationInfoResult`

NewAuthenticationInfoResult instantiates a new AuthenticationInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationInfoResultWithDefaults

`func NewAuthenticationInfoResultWithDefaults() *AuthenticationInfoResult`

NewAuthenticationInfoResultWithDefaults instantiates a new AuthenticationInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpi

`func (o *AuthenticationInfoResult) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *AuthenticationInfoResult) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *AuthenticationInfoResult) SetImpi(v string)`

SetImpi sets Impi field to given value.

### HasImpi

`func (o *AuthenticationInfoResult) HasImpi() bool`

HasImpi returns a boolean if a field has been set.

### GetVar3gAkaAv

`func (o *AuthenticationInfoResult) GetVar3gAkaAv() Model3GAkaAv`

GetVar3gAkaAv returns the Var3gAkaAv field if non-nil, zero value otherwise.

### GetVar3gAkaAvOk

`func (o *AuthenticationInfoResult) GetVar3gAkaAvOk() (*Model3GAkaAv, bool)`

GetVar3gAkaAvOk returns a tuple with the Var3gAkaAv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gAkaAv

`func (o *AuthenticationInfoResult) SetVar3gAkaAv(v Model3GAkaAv)`

SetVar3gAkaAv sets Var3gAkaAv field to given value.

### HasVar3gAkaAv

`func (o *AuthenticationInfoResult) HasVar3gAkaAv() bool`

HasVar3gAkaAv returns a boolean if a field has been set.

### GetDigestAuth

`func (o *AuthenticationInfoResult) GetDigestAuth() DigestAuthentication`

GetDigestAuth returns the DigestAuth field if non-nil, zero value otherwise.

### GetDigestAuthOk

`func (o *AuthenticationInfoResult) GetDigestAuthOk() (*DigestAuthentication, bool)`

GetDigestAuthOk returns a tuple with the DigestAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAuth

`func (o *AuthenticationInfoResult) SetDigestAuth(v DigestAuthentication)`

SetDigestAuth sets DigestAuth field to given value.

### HasDigestAuth

`func (o *AuthenticationInfoResult) HasDigestAuth() bool`

HasDigestAuth returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AuthenticationInfoResult) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthenticationInfoResult) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthenticationInfoResult) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthenticationInfoResult) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


