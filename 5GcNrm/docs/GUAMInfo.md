# GUAMInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PLMNId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AMFIdentifier** | Pointer to **int32** |  | [optional] 

## Methods

### NewGUAMInfo

`func NewGUAMInfo() *GUAMInfo`

NewGUAMInfo instantiates a new GUAMInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGUAMInfoWithDefaults

`func NewGUAMInfoWithDefaults() *GUAMInfo`

NewGUAMInfoWithDefaults instantiates a new GUAMInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPLMNId

`func (o *GUAMInfo) GetPLMNId() PlmnId`

GetPLMNId returns the PLMNId field if non-nil, zero value otherwise.

### GetPLMNIdOk

`func (o *GUAMInfo) GetPLMNIdOk() (*PlmnId, bool)`

GetPLMNIdOk returns a tuple with the PLMNId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPLMNId

`func (o *GUAMInfo) SetPLMNId(v PlmnId)`

SetPLMNId sets PLMNId field to given value.

### HasPLMNId

`func (o *GUAMInfo) HasPLMNId() bool`

HasPLMNId returns a boolean if a field has been set.

### GetAMFIdentifier

`func (o *GUAMInfo) GetAMFIdentifier() int32`

GetAMFIdentifier returns the AMFIdentifier field if non-nil, zero value otherwise.

### GetAMFIdentifierOk

`func (o *GUAMInfo) GetAMFIdentifierOk() (*int32, bool)`

GetAMFIdentifierOk returns a tuple with the AMFIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAMFIdentifier

`func (o *GUAMInfo) SetAMFIdentifier(v int32)`

SetAMFIdentifier sets AMFIdentifier field to given value.

### HasAMFIdentifier

`func (o *GUAMInfo) HasAMFIdentifier() bool`

HasAMFIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


