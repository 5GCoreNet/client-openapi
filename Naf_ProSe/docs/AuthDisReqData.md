# AuthDisReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthRequestType** | [**AuthRequestType**](AuthRequestType.md) |  | 
**ProseAppId** | Pointer to **[]string** |  | [optional] 
**AllowedSuffixNum** | Pointer to **int32** | contains the allowed number of suffixes. | [optional] 
**AppLevelContainer** | Pointer to **string** | Contains the Application Level Container. | [optional] 
**Rpauid** | Pointer to **string** | Contains the RPAUID. | [optional] 
**TargetRpauid** | Pointer to **string** | Contains the RPAUID. | [optional] 
**AuthUpdateCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewAuthDisReqData

`func NewAuthDisReqData(authRequestType AuthRequestType, ) *AuthDisReqData`

NewAuthDisReqData instantiates a new AuthDisReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthDisReqDataWithDefaults

`func NewAuthDisReqDataWithDefaults() *AuthDisReqData`

NewAuthDisReqDataWithDefaults instantiates a new AuthDisReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthRequestType

`func (o *AuthDisReqData) GetAuthRequestType() AuthRequestType`

GetAuthRequestType returns the AuthRequestType field if non-nil, zero value otherwise.

### GetAuthRequestTypeOk

`func (o *AuthDisReqData) GetAuthRequestTypeOk() (*AuthRequestType, bool)`

GetAuthRequestTypeOk returns a tuple with the AuthRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRequestType

`func (o *AuthDisReqData) SetAuthRequestType(v AuthRequestType)`

SetAuthRequestType sets AuthRequestType field to given value.


### GetProseAppId

`func (o *AuthDisReqData) GetProseAppId() []string`

GetProseAppId returns the ProseAppId field if non-nil, zero value otherwise.

### GetProseAppIdOk

`func (o *AuthDisReqData) GetProseAppIdOk() (*[]string, bool)`

GetProseAppIdOk returns a tuple with the ProseAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppId

`func (o *AuthDisReqData) SetProseAppId(v []string)`

SetProseAppId sets ProseAppId field to given value.

### HasProseAppId

`func (o *AuthDisReqData) HasProseAppId() bool`

HasProseAppId returns a boolean if a field has been set.

### GetAllowedSuffixNum

`func (o *AuthDisReqData) GetAllowedSuffixNum() int32`

GetAllowedSuffixNum returns the AllowedSuffixNum field if non-nil, zero value otherwise.

### GetAllowedSuffixNumOk

`func (o *AuthDisReqData) GetAllowedSuffixNumOk() (*int32, bool)`

GetAllowedSuffixNumOk returns a tuple with the AllowedSuffixNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSuffixNum

`func (o *AuthDisReqData) SetAllowedSuffixNum(v int32)`

SetAllowedSuffixNum sets AllowedSuffixNum field to given value.

### HasAllowedSuffixNum

`func (o *AuthDisReqData) HasAllowedSuffixNum() bool`

HasAllowedSuffixNum returns a boolean if a field has been set.

### GetAppLevelContainer

`func (o *AuthDisReqData) GetAppLevelContainer() string`

GetAppLevelContainer returns the AppLevelContainer field if non-nil, zero value otherwise.

### GetAppLevelContainerOk

`func (o *AuthDisReqData) GetAppLevelContainerOk() (*string, bool)`

GetAppLevelContainerOk returns a tuple with the AppLevelContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLevelContainer

`func (o *AuthDisReqData) SetAppLevelContainer(v string)`

SetAppLevelContainer sets AppLevelContainer field to given value.

### HasAppLevelContainer

`func (o *AuthDisReqData) HasAppLevelContainer() bool`

HasAppLevelContainer returns a boolean if a field has been set.

### GetRpauid

`func (o *AuthDisReqData) GetRpauid() string`

GetRpauid returns the Rpauid field if non-nil, zero value otherwise.

### GetRpauidOk

`func (o *AuthDisReqData) GetRpauidOk() (*string, bool)`

GetRpauidOk returns a tuple with the Rpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpauid

`func (o *AuthDisReqData) SetRpauid(v string)`

SetRpauid sets Rpauid field to given value.

### HasRpauid

`func (o *AuthDisReqData) HasRpauid() bool`

HasRpauid returns a boolean if a field has been set.

### GetTargetRpauid

`func (o *AuthDisReqData) GetTargetRpauid() string`

GetTargetRpauid returns the TargetRpauid field if non-nil, zero value otherwise.

### GetTargetRpauidOk

`func (o *AuthDisReqData) GetTargetRpauidOk() (*string, bool)`

GetTargetRpauidOk returns a tuple with the TargetRpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRpauid

`func (o *AuthDisReqData) SetTargetRpauid(v string)`

SetTargetRpauid sets TargetRpauid field to given value.

### HasTargetRpauid

`func (o *AuthDisReqData) HasTargetRpauid() bool`

HasTargetRpauid returns a boolean if a field has been set.

### GetAuthUpdateCallbackUri

`func (o *AuthDisReqData) GetAuthUpdateCallbackUri() string`

GetAuthUpdateCallbackUri returns the AuthUpdateCallbackUri field if non-nil, zero value otherwise.

### GetAuthUpdateCallbackUriOk

`func (o *AuthDisReqData) GetAuthUpdateCallbackUriOk() (*string, bool)`

GetAuthUpdateCallbackUriOk returns a tuple with the AuthUpdateCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUpdateCallbackUri

`func (o *AuthDisReqData) SetAuthUpdateCallbackUri(v string)`

SetAuthUpdateCallbackUri sets AuthUpdateCallbackUri field to given value.

### HasAuthUpdateCallbackUri

`func (o *AuthDisReqData) HasAuthUpdateCallbackUri() bool`

HasAuthUpdateCallbackUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


