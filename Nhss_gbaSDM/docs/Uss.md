# Uss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GsId** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**GsType** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**UeIds** | [**[]UeIdsItem**](UeIdsItem.md) |  | 
**NafGroup** | Pointer to **string** | Character string representing a NAF Group | [optional] 
**Flags** | Pointer to [**[]FlagsItem**](FlagsItem.md) |  | [optional] 
**KeyChoice** | Pointer to [**KeyChoice**](KeyChoice.md) |  | [optional] 

## Methods

### NewUss

`func NewUss(gsId int32, gsType int32, ueIds []UeIdsItem, ) *Uss`

NewUss instantiates a new Uss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUssWithDefaults

`func NewUssWithDefaults() *Uss`

NewUssWithDefaults instantiates a new Uss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGsId

`func (o *Uss) GetGsId() int32`

GetGsId returns the GsId field if non-nil, zero value otherwise.

### GetGsIdOk

`func (o *Uss) GetGsIdOk() (*int32, bool)`

GetGsIdOk returns a tuple with the GsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGsId

`func (o *Uss) SetGsId(v int32)`

SetGsId sets GsId field to given value.


### GetGsType

`func (o *Uss) GetGsType() int32`

GetGsType returns the GsType field if non-nil, zero value otherwise.

### GetGsTypeOk

`func (o *Uss) GetGsTypeOk() (*int32, bool)`

GetGsTypeOk returns a tuple with the GsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGsType

`func (o *Uss) SetGsType(v int32)`

SetGsType sets GsType field to given value.


### GetUeIds

`func (o *Uss) GetUeIds() []UeIdsItem`

GetUeIds returns the UeIds field if non-nil, zero value otherwise.

### GetUeIdsOk

`func (o *Uss) GetUeIdsOk() (*[]UeIdsItem, bool)`

GetUeIdsOk returns a tuple with the UeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIds

`func (o *Uss) SetUeIds(v []UeIdsItem)`

SetUeIds sets UeIds field to given value.


### GetNafGroup

`func (o *Uss) GetNafGroup() string`

GetNafGroup returns the NafGroup field if non-nil, zero value otherwise.

### GetNafGroupOk

`func (o *Uss) GetNafGroupOk() (*string, bool)`

GetNafGroupOk returns a tuple with the NafGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNafGroup

`func (o *Uss) SetNafGroup(v string)`

SetNafGroup sets NafGroup field to given value.

### HasNafGroup

`func (o *Uss) HasNafGroup() bool`

HasNafGroup returns a boolean if a field has been set.

### GetFlags

`func (o *Uss) GetFlags() []FlagsItem`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Uss) GetFlagsOk() (*[]FlagsItem, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Uss) SetFlags(v []FlagsItem)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Uss) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetKeyChoice

`func (o *Uss) GetKeyChoice() KeyChoice`

GetKeyChoice returns the KeyChoice field if non-nil, zero value otherwise.

### GetKeyChoiceOk

`func (o *Uss) GetKeyChoiceOk() (*KeyChoice, bool)`

GetKeyChoiceOk returns a tuple with the KeyChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyChoice

`func (o *Uss) SetKeyChoice(v KeyChoice)`

SetKeyChoice sets KeyChoice field to given value.

### HasKeyChoice

`func (o *Uss) HasKeyChoice() bool`

HasKeyChoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


