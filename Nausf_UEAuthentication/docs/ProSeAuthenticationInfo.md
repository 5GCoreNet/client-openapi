# ProSeAuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupiOrSuci** | Pointer to **string** | String identifying a SUPI or a SUCI. | [optional] 
**Var5gPrukId** | Pointer to **string** | The 5GPRUK ID is string in NAI format as specified in clause 28.7.19 of 3GPP TS 23.003.  | [optional] 
**RelayServiceCode** | **int32** | Relay Service Code to identify a connectivity service provided by the UE-to-Network relay.  | 
**Nonce1** | **NullableString** | contains an Nonce1 | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewProSeAuthenticationInfo

`func NewProSeAuthenticationInfo(relayServiceCode int32, nonce1 NullableString, ) *ProSeAuthenticationInfo`

NewProSeAuthenticationInfo instantiates a new ProSeAuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProSeAuthenticationInfoWithDefaults

`func NewProSeAuthenticationInfoWithDefaults() *ProSeAuthenticationInfo`

NewProSeAuthenticationInfoWithDefaults instantiates a new ProSeAuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupiOrSuci

`func (o *ProSeAuthenticationInfo) GetSupiOrSuci() string`

GetSupiOrSuci returns the SupiOrSuci field if non-nil, zero value otherwise.

### GetSupiOrSuciOk

`func (o *ProSeAuthenticationInfo) GetSupiOrSuciOk() (*string, bool)`

GetSupiOrSuciOk returns a tuple with the SupiOrSuci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiOrSuci

`func (o *ProSeAuthenticationInfo) SetSupiOrSuci(v string)`

SetSupiOrSuci sets SupiOrSuci field to given value.

### HasSupiOrSuci

`func (o *ProSeAuthenticationInfo) HasSupiOrSuci() bool`

HasSupiOrSuci returns a boolean if a field has been set.

### GetVar5gPrukId

`func (o *ProSeAuthenticationInfo) GetVar5gPrukId() string`

GetVar5gPrukId returns the Var5gPrukId field if non-nil, zero value otherwise.

### GetVar5gPrukIdOk

`func (o *ProSeAuthenticationInfo) GetVar5gPrukIdOk() (*string, bool)`

GetVar5gPrukIdOk returns a tuple with the Var5gPrukId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gPrukId

`func (o *ProSeAuthenticationInfo) SetVar5gPrukId(v string)`

SetVar5gPrukId sets Var5gPrukId field to given value.

### HasVar5gPrukId

`func (o *ProSeAuthenticationInfo) HasVar5gPrukId() bool`

HasVar5gPrukId returns a boolean if a field has been set.

### GetRelayServiceCode

`func (o *ProSeAuthenticationInfo) GetRelayServiceCode() int32`

GetRelayServiceCode returns the RelayServiceCode field if non-nil, zero value otherwise.

### GetRelayServiceCodeOk

`func (o *ProSeAuthenticationInfo) GetRelayServiceCodeOk() (*int32, bool)`

GetRelayServiceCodeOk returns a tuple with the RelayServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServiceCode

`func (o *ProSeAuthenticationInfo) SetRelayServiceCode(v int32)`

SetRelayServiceCode sets RelayServiceCode field to given value.


### GetNonce1

`func (o *ProSeAuthenticationInfo) GetNonce1() string`

GetNonce1 returns the Nonce1 field if non-nil, zero value otherwise.

### GetNonce1Ok

`func (o *ProSeAuthenticationInfo) GetNonce1Ok() (*string, bool)`

GetNonce1Ok returns a tuple with the Nonce1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce1

`func (o *ProSeAuthenticationInfo) SetNonce1(v string)`

SetNonce1 sets Nonce1 field to given value.


### SetNonce1Nil

`func (o *ProSeAuthenticationInfo) SetNonce1Nil(b bool)`

 SetNonce1Nil sets the value for Nonce1 to be an explicit nil

### UnsetNonce1
`func (o *ProSeAuthenticationInfo) UnsetNonce1()`

UnsetNonce1 ensures that no value is present for Nonce1, not even an explicit nil
### GetSupportedFeatures

`func (o *ProSeAuthenticationInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ProSeAuthenticationInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ProSeAuthenticationInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ProSeAuthenticationInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


