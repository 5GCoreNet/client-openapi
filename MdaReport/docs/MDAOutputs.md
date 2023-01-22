# MDAOutputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MDAType** | Pointer to **string** |  | [optional] 
**MdaOutputList** | Pointer to [**[]MDAOutputEntry**](MDAOutputEntry.md) |  | [optional] 
**MDARequestRef** | Pointer to **string** |  | [optional] 

## Methods

### NewMDAOutputs

`func NewMDAOutputs() *MDAOutputs`

NewMDAOutputs instantiates a new MDAOutputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMDAOutputsWithDefaults

`func NewMDAOutputsWithDefaults() *MDAOutputs`

NewMDAOutputsWithDefaults instantiates a new MDAOutputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMDAType

`func (o *MDAOutputs) GetMDAType() string`

GetMDAType returns the MDAType field if non-nil, zero value otherwise.

### GetMDATypeOk

`func (o *MDAOutputs) GetMDATypeOk() (*string, bool)`

GetMDATypeOk returns a tuple with the MDAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDAType

`func (o *MDAOutputs) SetMDAType(v string)`

SetMDAType sets MDAType field to given value.

### HasMDAType

`func (o *MDAOutputs) HasMDAType() bool`

HasMDAType returns a boolean if a field has been set.

### GetMdaOutputList

`func (o *MDAOutputs) GetMdaOutputList() []MDAOutputEntry`

GetMdaOutputList returns the MdaOutputList field if non-nil, zero value otherwise.

### GetMdaOutputListOk

`func (o *MDAOutputs) GetMdaOutputListOk() (*[]MDAOutputEntry, bool)`

GetMdaOutputListOk returns a tuple with the MdaOutputList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdaOutputList

`func (o *MDAOutputs) SetMdaOutputList(v []MDAOutputEntry)`

SetMdaOutputList sets MdaOutputList field to given value.

### HasMdaOutputList

`func (o *MDAOutputs) HasMdaOutputList() bool`

HasMdaOutputList returns a boolean if a field has been set.

### GetMDARequestRef

`func (o *MDAOutputs) GetMDARequestRef() string`

GetMDARequestRef returns the MDARequestRef field if non-nil, zero value otherwise.

### GetMDARequestRefOk

`func (o *MDAOutputs) GetMDARequestRefOk() (*string, bool)`

GetMDARequestRefOk returns a tuple with the MDARequestRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDARequestRef

`func (o *MDAOutputs) SetMDARequestRef(v string)`

SetMDARequestRef sets MDARequestRef field to given value.

### HasMDARequestRef

`func (o *MDAOutputs) HasMDARequestRef() bool`

HasMDARequestRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


