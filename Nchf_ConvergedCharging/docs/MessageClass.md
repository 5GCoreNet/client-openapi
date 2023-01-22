# MessageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassIdentifier** | Pointer to [**ClassIdentifier**](ClassIdentifier.md) |  | [optional] 
**TokenText** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageClass

`func NewMessageClass() *MessageClass`

NewMessageClass instantiates a new MessageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageClassWithDefaults

`func NewMessageClassWithDefaults() *MessageClass`

NewMessageClassWithDefaults instantiates a new MessageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassIdentifier

`func (o *MessageClass) GetClassIdentifier() ClassIdentifier`

GetClassIdentifier returns the ClassIdentifier field if non-nil, zero value otherwise.

### GetClassIdentifierOk

`func (o *MessageClass) GetClassIdentifierOk() (*ClassIdentifier, bool)`

GetClassIdentifierOk returns a tuple with the ClassIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassIdentifier

`func (o *MessageClass) SetClassIdentifier(v ClassIdentifier)`

SetClassIdentifier sets ClassIdentifier field to given value.

### HasClassIdentifier

`func (o *MessageClass) HasClassIdentifier() bool`

HasClassIdentifier returns a boolean if a field has been set.

### GetTokenText

`func (o *MessageClass) GetTokenText() string`

GetTokenText returns the TokenText field if non-nil, zero value otherwise.

### GetTokenTextOk

`func (o *MessageClass) GetTokenTextOk() (*string, bool)`

GetTokenTextOk returns a tuple with the TokenText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenText

`func (o *MessageClass) SetTokenText(v string)`

SetTokenText sets TokenText field to given value.

### HasTokenText

`func (o *MessageClass) HasTokenText() bool`

HasTokenText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


