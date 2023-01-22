# AccessTokenRsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | JWS Compact Serialized representation of JWS signed JSON object (AccessTokenClaims)  | 
**TokenType** | **string** |  | 
**ExpiresIn** | **int32** | Unsigned integer identifying a period of time in units of seconds. | 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessTokenRsp

`func NewAccessTokenRsp(accessToken string, tokenType string, expiresIn int32, ) *AccessTokenRsp`

NewAccessTokenRsp instantiates a new AccessTokenRsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenRspWithDefaults

`func NewAccessTokenRspWithDefaults() *AccessTokenRsp`

NewAccessTokenRspWithDefaults instantiates a new AccessTokenRsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AccessTokenRsp) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccessTokenRsp) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccessTokenRsp) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *AccessTokenRsp) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AccessTokenRsp) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AccessTokenRsp) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *AccessTokenRsp) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AccessTokenRsp) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AccessTokenRsp) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetScope

`func (o *AccessTokenRsp) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessTokenRsp) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessTokenRsp) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccessTokenRsp) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


