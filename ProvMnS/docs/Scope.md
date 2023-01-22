# Scope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeType** | Pointer to [**ScopeType**](ScopeType.md) |  | [optional] 
**ScopeLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewScope

`func NewScope() *Scope`

NewScope instantiates a new Scope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeWithDefaults

`func NewScopeWithDefaults() *Scope`

NewScopeWithDefaults instantiates a new Scope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeType

`func (o *Scope) GetScopeType() ScopeType`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *Scope) GetScopeTypeOk() (*ScopeType, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *Scope) SetScopeType(v ScopeType)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *Scope) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetScopeLevel

`func (o *Scope) GetScopeLevel() int32`

GetScopeLevel returns the ScopeLevel field if non-nil, zero value otherwise.

### GetScopeLevelOk

`func (o *Scope) GetScopeLevelOk() (*int32, bool)`

GetScopeLevelOk returns a tuple with the ScopeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeLevel

`func (o *Scope) SetScopeLevel(v int32)`

SetScopeLevel sets ScopeLevel field to given value.

### HasScopeLevel

`func (o *Scope) HasScopeLevel() bool`

HasScopeLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


