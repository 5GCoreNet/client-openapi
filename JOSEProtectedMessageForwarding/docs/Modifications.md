# Modifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** | Fully Qualified Domain Name | 
**Operations** | Pointer to [**[]PatchItem**](PatchItem.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewModifications

`func NewModifications(identity string, ) *Modifications`

NewModifications instantiates a new Modifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModificationsWithDefaults

`func NewModificationsWithDefaults() *Modifications`

NewModificationsWithDefaults instantiates a new Modifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *Modifications) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Modifications) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Modifications) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### GetOperations

`func (o *Modifications) GetOperations() []PatchItem`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Modifications) GetOperationsOk() (*[]PatchItem, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Modifications) SetOperations(v []PatchItem)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Modifications) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetTag

`func (o *Modifications) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Modifications) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Modifications) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Modifications) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


