# GbaAuthenticationInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | [**GbaAuthType**](GbaAuthType.md) |  | 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo1**](ResynchronizationInfo1.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewGbaAuthenticationInfoRequest

`func NewGbaAuthenticationInfoRequest(authType GbaAuthType, ) *GbaAuthenticationInfoRequest`

NewGbaAuthenticationInfoRequest instantiates a new GbaAuthenticationInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGbaAuthenticationInfoRequestWithDefaults

`func NewGbaAuthenticationInfoRequestWithDefaults() *GbaAuthenticationInfoRequest`

NewGbaAuthenticationInfoRequestWithDefaults instantiates a new GbaAuthenticationInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *GbaAuthenticationInfoRequest) GetAuthType() GbaAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *GbaAuthenticationInfoRequest) GetAuthTypeOk() (*GbaAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *GbaAuthenticationInfoRequest) SetAuthType(v GbaAuthType)`

SetAuthType sets AuthType field to given value.


### GetResynchronizationInfo

`func (o *GbaAuthenticationInfoRequest) GetResynchronizationInfo() ResynchronizationInfo1`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *GbaAuthenticationInfoRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo1, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *GbaAuthenticationInfoRequest) SetResynchronizationInfo(v ResynchronizationInfo1)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *GbaAuthenticationInfoRequest) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *GbaAuthenticationInfoRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *GbaAuthenticationInfoRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *GbaAuthenticationInfoRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *GbaAuthenticationInfoRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


