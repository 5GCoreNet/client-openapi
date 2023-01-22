# AccessTokenClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iss** | **string** |  | 
**Scope** | **string** |  | 
**Exp** | **int32** | Unsigned integer identifying a period of time in units of seconds. | 

## Methods

### NewAccessTokenClaims

`func NewAccessTokenClaims(iss string, scope string, exp int32, ) *AccessTokenClaims`

NewAccessTokenClaims instantiates a new AccessTokenClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenClaimsWithDefaults

`func NewAccessTokenClaimsWithDefaults() *AccessTokenClaims`

NewAccessTokenClaimsWithDefaults instantiates a new AccessTokenClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIss

`func (o *AccessTokenClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *AccessTokenClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *AccessTokenClaims) SetIss(v string)`

SetIss sets Iss field to given value.


### GetScope

`func (o *AccessTokenClaims) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessTokenClaims) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessTokenClaims) SetScope(v string)`

SetScope sets Scope field to given value.


### GetExp

`func (o *AccessTokenClaims) GetExp() int32`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *AccessTokenClaims) GetExpOk() (*int32, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *AccessTokenClaims) SetExp(v int32)`

SetExp sets Exp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


