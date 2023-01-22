# ExpectationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**ObjectContexts** | Pointer to [**[]ObjectContext**](ObjectContext.md) |  | [optional] 

## Methods

### NewExpectationObject

`func NewExpectationObject() *ExpectationObject`

NewExpectationObject instantiates a new ExpectationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectationObjectWithDefaults

`func NewExpectationObjectWithDefaults() *ExpectationObject`

NewExpectationObjectWithDefaults instantiates a new ExpectationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *ExpectationObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExpectationObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExpectationObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ExpectationObject) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ExpectationObject) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ExpectationObject) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ExpectationObject) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ExpectationObject) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetObjectContexts

`func (o *ExpectationObject) GetObjectContexts() []ObjectContext`

GetObjectContexts returns the ObjectContexts field if non-nil, zero value otherwise.

### GetObjectContextsOk

`func (o *ExpectationObject) GetObjectContextsOk() (*[]ObjectContext, bool)`

GetObjectContextsOk returns a tuple with the ObjectContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectContexts

`func (o *ExpectationObject) SetObjectContexts(v []ObjectContext)`

SetObjectContexts sets ObjectContexts field to given value.

### HasObjectContexts

`func (o *ExpectationObject) HasObjectContexts() bool`

HasObjectContexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


