# UeSliceMbr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SliceMbr** | [**map[string]SliceMbr**](SliceMbr.md) | Contains the MBR for uplink and the MBR for downlink. | 
**ServingSnssai** | [**Snssai**](Snssai.md) |  | 
**MappedHomeSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewUeSliceMbr

`func NewUeSliceMbr(sliceMbr map[string]SliceMbr, servingSnssai Snssai, ) *UeSliceMbr`

NewUeSliceMbr instantiates a new UeSliceMbr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeSliceMbrWithDefaults

`func NewUeSliceMbrWithDefaults() *UeSliceMbr`

NewUeSliceMbrWithDefaults instantiates a new UeSliceMbr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSliceMbr

`func (o *UeSliceMbr) GetSliceMbr() map[string]SliceMbr`

GetSliceMbr returns the SliceMbr field if non-nil, zero value otherwise.

### GetSliceMbrOk

`func (o *UeSliceMbr) GetSliceMbrOk() (*map[string]SliceMbr, bool)`

GetSliceMbrOk returns a tuple with the SliceMbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMbr

`func (o *UeSliceMbr) SetSliceMbr(v map[string]SliceMbr)`

SetSliceMbr sets SliceMbr field to given value.


### GetServingSnssai

`func (o *UeSliceMbr) GetServingSnssai() Snssai`

GetServingSnssai returns the ServingSnssai field if non-nil, zero value otherwise.

### GetServingSnssaiOk

`func (o *UeSliceMbr) GetServingSnssaiOk() (*Snssai, bool)`

GetServingSnssaiOk returns a tuple with the ServingSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingSnssai

`func (o *UeSliceMbr) SetServingSnssai(v Snssai)`

SetServingSnssai sets ServingSnssai field to given value.


### GetMappedHomeSnssai

`func (o *UeSliceMbr) GetMappedHomeSnssai() Snssai`

GetMappedHomeSnssai returns the MappedHomeSnssai field if non-nil, zero value otherwise.

### GetMappedHomeSnssaiOk

`func (o *UeSliceMbr) GetMappedHomeSnssaiOk() (*Snssai, bool)`

GetMappedHomeSnssaiOk returns a tuple with the MappedHomeSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedHomeSnssai

`func (o *UeSliceMbr) SetMappedHomeSnssai(v Snssai)`

SetMappedHomeSnssai sets MappedHomeSnssai field to given value.

### HasMappedHomeSnssai

`func (o *UeSliceMbr) HasMappedHomeSnssai() bool`

HasMappedHomeSnssai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


