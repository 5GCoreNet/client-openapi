# FreqInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arfcn** | Pointer to **int32** |  | [optional] 
**FreqBands** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewFreqInfo

`func NewFreqInfo() *FreqInfo`

NewFreqInfo instantiates a new FreqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreqInfoWithDefaults

`func NewFreqInfoWithDefaults() *FreqInfo`

NewFreqInfoWithDefaults instantiates a new FreqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArfcn

`func (o *FreqInfo) GetArfcn() int32`

GetArfcn returns the Arfcn field if non-nil, zero value otherwise.

### GetArfcnOk

`func (o *FreqInfo) GetArfcnOk() (*int32, bool)`

GetArfcnOk returns a tuple with the Arfcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArfcn

`func (o *FreqInfo) SetArfcn(v int32)`

SetArfcn sets Arfcn field to given value.

### HasArfcn

`func (o *FreqInfo) HasArfcn() bool`

HasArfcn returns a boolean if a field has been set.

### GetFreqBands

`func (o *FreqInfo) GetFreqBands() []int32`

GetFreqBands returns the FreqBands field if non-nil, zero value otherwise.

### GetFreqBandsOk

`func (o *FreqInfo) GetFreqBandsOk() (*[]int32, bool)`

GetFreqBandsOk returns a tuple with the FreqBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreqBands

`func (o *FreqInfo) SetFreqBands(v []int32)`

SetFreqBands sets FreqBands field to given value.

### HasFreqBands

`func (o *FreqInfo) HasFreqBands() bool`

HasFreqBands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


