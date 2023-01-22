# ProSeAuthenticationCtx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | [**AuthType**](AuthType.md) |  | 
**Links** | [**map[string]LinksValueSchema**](LinksValueSchema.md) | A map(list of key-value pairs) where member serves as key | 
**ProSeAuthData** | [**ProSeAuthData**](ProSeAuthData.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewProSeAuthenticationCtx

`func NewProSeAuthenticationCtx(authType AuthType, links map[string]LinksValueSchema, proSeAuthData ProSeAuthData, ) *ProSeAuthenticationCtx`

NewProSeAuthenticationCtx instantiates a new ProSeAuthenticationCtx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProSeAuthenticationCtxWithDefaults

`func NewProSeAuthenticationCtxWithDefaults() *ProSeAuthenticationCtx`

NewProSeAuthenticationCtxWithDefaults instantiates a new ProSeAuthenticationCtx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *ProSeAuthenticationCtx) GetAuthType() AuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ProSeAuthenticationCtx) GetAuthTypeOk() (*AuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ProSeAuthenticationCtx) SetAuthType(v AuthType)`

SetAuthType sets AuthType field to given value.


### GetLinks

`func (o *ProSeAuthenticationCtx) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProSeAuthenticationCtx) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProSeAuthenticationCtx) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.


### GetProSeAuthData

`func (o *ProSeAuthenticationCtx) GetProSeAuthData() ProSeAuthData`

GetProSeAuthData returns the ProSeAuthData field if non-nil, zero value otherwise.

### GetProSeAuthDataOk

`func (o *ProSeAuthenticationCtx) GetProSeAuthDataOk() (*ProSeAuthData, bool)`

GetProSeAuthDataOk returns a tuple with the ProSeAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProSeAuthData

`func (o *ProSeAuthenticationCtx) SetProSeAuthData(v ProSeAuthData)`

SetProSeAuthData sets ProSeAuthData field to given value.


### GetSupportedFeatures

`func (o *ProSeAuthenticationCtx) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ProSeAuthenticationCtx) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ProSeAuthenticationCtx) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ProSeAuthenticationCtx) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


