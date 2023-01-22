# TargetArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**AnyTa** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewTargetArea

`func NewTargetArea() *TargetArea`

NewTargetArea instantiates a new TargetArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetAreaWithDefaults

`func NewTargetAreaWithDefaults() *TargetArea`

NewTargetAreaWithDefaults instantiates a new TargetArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaList

`func (o *TargetArea) GetTaList() []Tai`

GetTaList returns the TaList field if non-nil, zero value otherwise.

### GetTaListOk

`func (o *TargetArea) GetTaListOk() (*[]Tai, bool)`

GetTaListOk returns a tuple with the TaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaList

`func (o *TargetArea) SetTaList(v []Tai)`

SetTaList sets TaList field to given value.

### HasTaList

`func (o *TargetArea) HasTaList() bool`

HasTaList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *TargetArea) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *TargetArea) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *TargetArea) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *TargetArea) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetAnyTa

`func (o *TargetArea) GetAnyTa() bool`

GetAnyTa returns the AnyTa field if non-nil, zero value otherwise.

### GetAnyTaOk

`func (o *TargetArea) GetAnyTaOk() (*bool, bool)`

GetAnyTaOk returns a tuple with the AnyTa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyTa

`func (o *TargetArea) SetAnyTa(v bool)`

SetAnyTa sets AnyTa field to given value.

### HasAnyTa

`func (o *TargetArea) HasAnyTa() bool`

HasAnyTa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


