# Ifcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IfcList** | Pointer to [**[]Ifc**](Ifc.md) |  | [optional] 
**CscfFilterSetIdList** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewIfcs

`func NewIfcs() *Ifcs`

NewIfcs instantiates a new Ifcs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIfcsWithDefaults

`func NewIfcsWithDefaults() *Ifcs`

NewIfcsWithDefaults instantiates a new Ifcs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIfcList

`func (o *Ifcs) GetIfcList() []Ifc`

GetIfcList returns the IfcList field if non-nil, zero value otherwise.

### GetIfcListOk

`func (o *Ifcs) GetIfcListOk() (*[]Ifc, bool)`

GetIfcListOk returns a tuple with the IfcList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfcList

`func (o *Ifcs) SetIfcList(v []Ifc)`

SetIfcList sets IfcList field to given value.

### HasIfcList

`func (o *Ifcs) HasIfcList() bool`

HasIfcList returns a boolean if a field has been set.

### GetCscfFilterSetIdList

`func (o *Ifcs) GetCscfFilterSetIdList() []int32`

GetCscfFilterSetIdList returns the CscfFilterSetIdList field if non-nil, zero value otherwise.

### GetCscfFilterSetIdListOk

`func (o *Ifcs) GetCscfFilterSetIdListOk() (*[]int32, bool)`

GetCscfFilterSetIdListOk returns a tuple with the CscfFilterSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscfFilterSetIdList

`func (o *Ifcs) SetCscfFilterSetIdList(v []int32)`

SetCscfFilterSetIdList sets CscfFilterSetIdList field to given value.

### HasCscfFilterSetIdList

`func (o *Ifcs) HasCscfFilterSetIdList() bool`

HasCscfFilterSetIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


