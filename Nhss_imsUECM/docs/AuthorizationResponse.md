# AuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationResult** | [**AuthorizationResult**](AuthorizationResult.md) |  | 
**CscfServerName** | Pointer to **string** |  | [optional] 
**ScscfSelectionAssistanceInfo** | Pointer to [**ScscfSelectionAssistanceInformation**](ScscfSelectionAssistanceInformation.md) |  | [optional] 

## Methods

### NewAuthorizationResponse

`func NewAuthorizationResponse(authorizationResult AuthorizationResult, ) *AuthorizationResponse`

NewAuthorizationResponse instantiates a new AuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationResponseWithDefaults

`func NewAuthorizationResponseWithDefaults() *AuthorizationResponse`

NewAuthorizationResponseWithDefaults instantiates a new AuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationResult

`func (o *AuthorizationResponse) GetAuthorizationResult() AuthorizationResult`

GetAuthorizationResult returns the AuthorizationResult field if non-nil, zero value otherwise.

### GetAuthorizationResultOk

`func (o *AuthorizationResponse) GetAuthorizationResultOk() (*AuthorizationResult, bool)`

GetAuthorizationResultOk returns a tuple with the AuthorizationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationResult

`func (o *AuthorizationResponse) SetAuthorizationResult(v AuthorizationResult)`

SetAuthorizationResult sets AuthorizationResult field to given value.


### GetCscfServerName

`func (o *AuthorizationResponse) GetCscfServerName() string`

GetCscfServerName returns the CscfServerName field if non-nil, zero value otherwise.

### GetCscfServerNameOk

`func (o *AuthorizationResponse) GetCscfServerNameOk() (*string, bool)`

GetCscfServerNameOk returns a tuple with the CscfServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscfServerName

`func (o *AuthorizationResponse) SetCscfServerName(v string)`

SetCscfServerName sets CscfServerName field to given value.

### HasCscfServerName

`func (o *AuthorizationResponse) HasCscfServerName() bool`

HasCscfServerName returns a boolean if a field has been set.

### GetScscfSelectionAssistanceInfo

`func (o *AuthorizationResponse) GetScscfSelectionAssistanceInfo() ScscfSelectionAssistanceInformation`

GetScscfSelectionAssistanceInfo returns the ScscfSelectionAssistanceInfo field if non-nil, zero value otherwise.

### GetScscfSelectionAssistanceInfoOk

`func (o *AuthorizationResponse) GetScscfSelectionAssistanceInfoOk() (*ScscfSelectionAssistanceInformation, bool)`

GetScscfSelectionAssistanceInfoOk returns a tuple with the ScscfSelectionAssistanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScscfSelectionAssistanceInfo

`func (o *AuthorizationResponse) SetScscfSelectionAssistanceInfo(v ScscfSelectionAssistanceInformation)`

SetScscfSelectionAssistanceInfo sets ScscfSelectionAssistanceInfo field to given value.

### HasScscfSelectionAssistanceInfo

`func (o *AuthorizationResponse) HasScscfSelectionAssistanceInfo() bool`

HasScscfSelectionAssistanceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


