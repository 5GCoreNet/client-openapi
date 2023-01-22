# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AllowedNFTypes** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**OperationSemantics** | Pointer to [**OperationSemantics**](OperationSemantics.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Operation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Operation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Operation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Operation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowedNFTypes

`func (o *Operation) GetAllowedNFTypes() NFType`

GetAllowedNFTypes returns the AllowedNFTypes field if non-nil, zero value otherwise.

### GetAllowedNFTypesOk

`func (o *Operation) GetAllowedNFTypesOk() (*NFType, bool)`

GetAllowedNFTypesOk returns a tuple with the AllowedNFTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNFTypes

`func (o *Operation) SetAllowedNFTypes(v NFType)`

SetAllowedNFTypes sets AllowedNFTypes field to given value.

### HasAllowedNFTypes

`func (o *Operation) HasAllowedNFTypes() bool`

HasAllowedNFTypes returns a boolean if a field has been set.

### GetOperationSemantics

`func (o *Operation) GetOperationSemantics() OperationSemantics`

GetOperationSemantics returns the OperationSemantics field if non-nil, zero value otherwise.

### GetOperationSemanticsOk

`func (o *Operation) GetOperationSemanticsOk() (*OperationSemantics, bool)`

GetOperationSemanticsOk returns a tuple with the OperationSemantics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationSemantics

`func (o *Operation) SetOperationSemantics(v OperationSemantics)`

SetOperationSemantics sets OperationSemantics field to given value.

### HasOperationSemantics

`func (o *Operation) HasOperationSemantics() bool`

HasOperationSemantics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


