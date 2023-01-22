# AreaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreqInfo** | Pointer to [**FreqInfo**](FreqInfo.md) |  | [optional] 
**PciList** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewAreaConfig

`func NewAreaConfig() *AreaConfig`

NewAreaConfig instantiates a new AreaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAreaConfigWithDefaults

`func NewAreaConfigWithDefaults() *AreaConfig`

NewAreaConfigWithDefaults instantiates a new AreaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreqInfo

`func (o *AreaConfig) GetFreqInfo() FreqInfo`

GetFreqInfo returns the FreqInfo field if non-nil, zero value otherwise.

### GetFreqInfoOk

`func (o *AreaConfig) GetFreqInfoOk() (*FreqInfo, bool)`

GetFreqInfoOk returns a tuple with the FreqInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreqInfo

`func (o *AreaConfig) SetFreqInfo(v FreqInfo)`

SetFreqInfo sets FreqInfo field to given value.

### HasFreqInfo

`func (o *AreaConfig) HasFreqInfo() bool`

HasFreqInfo returns a boolean if a field has been set.

### GetPciList

`func (o *AreaConfig) GetPciList() []int32`

GetPciList returns the PciList field if non-nil, zero value otherwise.

### GetPciListOk

`func (o *AreaConfig) GetPciListOk() (*[]int32, bool)`

GetPciListOk returns a tuple with the PciList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciList

`func (o *AreaConfig) SetPciList(v []int32)`

SetPciList sets PciList field to given value.

### HasPciList

`func (o *AreaConfig) HasPciList() bool`

HasPciList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


