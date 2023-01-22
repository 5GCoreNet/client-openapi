# TmgiAllocate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TmgiNumber** | Pointer to **int32** | The number of requested TMGIs | [optional] 
**TmgiList** | Pointer to [**[]Tmgi**](Tmgi.md) | The list of TMGIs to be refreshed | [optional] 

## Methods

### NewTmgiAllocate

`func NewTmgiAllocate() *TmgiAllocate`

NewTmgiAllocate instantiates a new TmgiAllocate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTmgiAllocateWithDefaults

`func NewTmgiAllocateWithDefaults() *TmgiAllocate`

NewTmgiAllocateWithDefaults instantiates a new TmgiAllocate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmgiNumber

`func (o *TmgiAllocate) GetTmgiNumber() int32`

GetTmgiNumber returns the TmgiNumber field if non-nil, zero value otherwise.

### GetTmgiNumberOk

`func (o *TmgiAllocate) GetTmgiNumberOk() (*int32, bool)`

GetTmgiNumberOk returns a tuple with the TmgiNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgiNumber

`func (o *TmgiAllocate) SetTmgiNumber(v int32)`

SetTmgiNumber sets TmgiNumber field to given value.

### HasTmgiNumber

`func (o *TmgiAllocate) HasTmgiNumber() bool`

HasTmgiNumber returns a boolean if a field has been set.

### GetTmgiList

`func (o *TmgiAllocate) GetTmgiList() []Tmgi`

GetTmgiList returns the TmgiList field if non-nil, zero value otherwise.

### GetTmgiListOk

`func (o *TmgiAllocate) GetTmgiListOk() (*[]Tmgi, bool)`

GetTmgiListOk returns a tuple with the TmgiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgiList

`func (o *TmgiAllocate) SetTmgiList(v []Tmgi)`

SetTmgiList sets TmgiList field to given value.

### HasTmgiList

`func (o *TmgiAllocate) HasTmgiList() bool`

HasTmgiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


