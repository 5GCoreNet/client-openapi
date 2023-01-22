# AuthenticationInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationScheme** | Pointer to [**AuthenticationScheme**](AuthenticationScheme.md) |  | [optional] 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAuthenticationInfoRequest

`func NewAuthenticationInfoRequest() *AuthenticationInfoRequest`

NewAuthenticationInfoRequest instantiates a new AuthenticationInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationInfoRequestWithDefaults

`func NewAuthenticationInfoRequestWithDefaults() *AuthenticationInfoRequest`

NewAuthenticationInfoRequestWithDefaults instantiates a new AuthenticationInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationScheme

`func (o *AuthenticationInfoRequest) GetAuthenticationScheme() AuthenticationScheme`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *AuthenticationInfoRequest) GetAuthenticationSchemeOk() (*AuthenticationScheme, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *AuthenticationInfoRequest) SetAuthenticationScheme(v AuthenticationScheme)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *AuthenticationInfoRequest) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### GetResynchronizationInfo

`func (o *AuthenticationInfoRequest) GetResynchronizationInfo() ResynchronizationInfo`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *AuthenticationInfoRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *AuthenticationInfoRequest) SetResynchronizationInfo(v ResynchronizationInfo)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *AuthenticationInfoRequest) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AuthenticationInfoRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthenticationInfoRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthenticationInfoRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthenticationInfoRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


