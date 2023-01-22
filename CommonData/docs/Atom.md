# Atom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | **string** | contains the name of a defined query parameter. | 
**Value** | **interface{}** |  | 
**Negative** | Pointer to **bool** | indicates whether the negative condition applies for the query condition. | [optional] 

## Methods

### NewAtom

`func NewAtom(attr string, value interface{}, ) *Atom`

NewAtom instantiates a new Atom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtomWithDefaults

`func NewAtomWithDefaults() *Atom`

NewAtomWithDefaults instantiates a new Atom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *Atom) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *Atom) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *Atom) SetAttr(v string)`

SetAttr sets Attr field to given value.


### GetValue

`func (o *Atom) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Atom) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Atom) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Atom) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Atom) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetNegative

`func (o *Atom) GetNegative() bool`

GetNegative returns the Negative field if non-nil, zero value otherwise.

### GetNegativeOk

`func (o *Atom) GetNegativeOk() (*bool, bool)`

GetNegativeOk returns a tuple with the Negative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegative

`func (o *Atom) SetNegative(v bool)`

SetNegative sets Negative field to given value.

### HasNegative

`func (o *Atom) HasNegative() bool`

HasNegative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


