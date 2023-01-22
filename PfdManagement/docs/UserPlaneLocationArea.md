# UserPlaneLocationArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationArea** | Pointer to [**LocationArea**](LocationArea.md) |  | [optional] 
**LocationArea5G** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**Dnais** | Pointer to **[]string** | Identifies a list of DNAI which the user plane functions support. | [optional] 

## Methods

### NewUserPlaneLocationArea

`func NewUserPlaneLocationArea() *UserPlaneLocationArea`

NewUserPlaneLocationArea instantiates a new UserPlaneLocationArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPlaneLocationAreaWithDefaults

`func NewUserPlaneLocationAreaWithDefaults() *UserPlaneLocationArea`

NewUserPlaneLocationAreaWithDefaults instantiates a new UserPlaneLocationArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationArea

`func (o *UserPlaneLocationArea) GetLocationArea() LocationArea`

GetLocationArea returns the LocationArea field if non-nil, zero value otherwise.

### GetLocationAreaOk

`func (o *UserPlaneLocationArea) GetLocationAreaOk() (*LocationArea, bool)`

GetLocationAreaOk returns a tuple with the LocationArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea

`func (o *UserPlaneLocationArea) SetLocationArea(v LocationArea)`

SetLocationArea sets LocationArea field to given value.

### HasLocationArea

`func (o *UserPlaneLocationArea) HasLocationArea() bool`

HasLocationArea returns a boolean if a field has been set.

### GetLocationArea5G

`func (o *UserPlaneLocationArea) GetLocationArea5G() LocationArea5G`

GetLocationArea5G returns the LocationArea5G field if non-nil, zero value otherwise.

### GetLocationArea5GOk

`func (o *UserPlaneLocationArea) GetLocationArea5GOk() (*LocationArea5G, bool)`

GetLocationArea5GOk returns a tuple with the LocationArea5G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea5G

`func (o *UserPlaneLocationArea) SetLocationArea5G(v LocationArea5G)`

SetLocationArea5G sets LocationArea5G field to given value.

### HasLocationArea5G

`func (o *UserPlaneLocationArea) HasLocationArea5G() bool`

HasLocationArea5G returns a boolean if a field has been set.

### GetDnais

`func (o *UserPlaneLocationArea) GetDnais() []string`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *UserPlaneLocationArea) GetDnaisOk() (*[]string, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *UserPlaneLocationArea) SetDnais(v []string)`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *UserPlaneLocationArea) HasDnais() bool`

HasDnais returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


