# NRFreqRelationSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OffsetMO** | Pointer to [**QOffsetRangeList**](QOffsetRangeList.md) |  | [optional] 
**BlockListEntry** | Pointer to **[]int32** |  | [optional] 
**BlockListEntryIdleMode** | Pointer to **int32** |  | [optional] 
**CellReselectionPriority** | Pointer to **int32** |  | [optional] 
**CellReselectionSubPriority** | Pointer to **float32** |  | [optional] 
**PMax** | Pointer to **int32** |  | [optional] 
**QOffsetFreq** | Pointer to **float32** |  | [optional] 
**QQualMin** | Pointer to **float32** |  | [optional] 
**QRxLevMin** | Pointer to **int32** |  | [optional] 
**ThreshXHighP** | Pointer to **int32** |  | [optional] 
**ThreshXHighQ** | Pointer to **int32** |  | [optional] 
**ThreshXLowP** | Pointer to **int32** |  | [optional] 
**ThreshXLowQ** | Pointer to **int32** |  | [optional] 
**TReselectionNr** | Pointer to **int32** |  | [optional] 
**TReselectionNRSfHigh** | Pointer to [**TReselectionNRSf**](TReselectionNRSf.md) |  | [optional] 
**TReselectionNRSfMedium** | Pointer to [**TReselectionNRSf**](TReselectionNRSf.md) |  | [optional] 
**NRFrequencyRef** | Pointer to **string** |  | [optional] 

## Methods

### NewNRFreqRelationSingleAllOfAttributes

`func NewNRFreqRelationSingleAllOfAttributes() *NRFreqRelationSingleAllOfAttributes`

NewNRFreqRelationSingleAllOfAttributes instantiates a new NRFreqRelationSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNRFreqRelationSingleAllOfAttributesWithDefaults

`func NewNRFreqRelationSingleAllOfAttributesWithDefaults() *NRFreqRelationSingleAllOfAttributes`

NewNRFreqRelationSingleAllOfAttributesWithDefaults instantiates a new NRFreqRelationSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffsetMO

`func (o *NRFreqRelationSingleAllOfAttributes) GetOffsetMO() QOffsetRangeList`

GetOffsetMO returns the OffsetMO field if non-nil, zero value otherwise.

### GetOffsetMOOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetOffsetMOOk() (*QOffsetRangeList, bool)`

GetOffsetMOOk returns a tuple with the OffsetMO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetMO

`func (o *NRFreqRelationSingleAllOfAttributes) SetOffsetMO(v QOffsetRangeList)`

SetOffsetMO sets OffsetMO field to given value.

### HasOffsetMO

`func (o *NRFreqRelationSingleAllOfAttributes) HasOffsetMO() bool`

HasOffsetMO returns a boolean if a field has been set.

### GetBlockListEntry

`func (o *NRFreqRelationSingleAllOfAttributes) GetBlockListEntry() []int32`

GetBlockListEntry returns the BlockListEntry field if non-nil, zero value otherwise.

### GetBlockListEntryOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetBlockListEntryOk() (*[]int32, bool)`

GetBlockListEntryOk returns a tuple with the BlockListEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockListEntry

`func (o *NRFreqRelationSingleAllOfAttributes) SetBlockListEntry(v []int32)`

SetBlockListEntry sets BlockListEntry field to given value.

### HasBlockListEntry

`func (o *NRFreqRelationSingleAllOfAttributes) HasBlockListEntry() bool`

HasBlockListEntry returns a boolean if a field has been set.

### GetBlockListEntryIdleMode

`func (o *NRFreqRelationSingleAllOfAttributes) GetBlockListEntryIdleMode() int32`

GetBlockListEntryIdleMode returns the BlockListEntryIdleMode field if non-nil, zero value otherwise.

### GetBlockListEntryIdleModeOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetBlockListEntryIdleModeOk() (*int32, bool)`

GetBlockListEntryIdleModeOk returns a tuple with the BlockListEntryIdleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockListEntryIdleMode

`func (o *NRFreqRelationSingleAllOfAttributes) SetBlockListEntryIdleMode(v int32)`

SetBlockListEntryIdleMode sets BlockListEntryIdleMode field to given value.

### HasBlockListEntryIdleMode

`func (o *NRFreqRelationSingleAllOfAttributes) HasBlockListEntryIdleMode() bool`

HasBlockListEntryIdleMode returns a boolean if a field has been set.

### GetCellReselectionPriority

`func (o *NRFreqRelationSingleAllOfAttributes) GetCellReselectionPriority() int32`

GetCellReselectionPriority returns the CellReselectionPriority field if non-nil, zero value otherwise.

### GetCellReselectionPriorityOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetCellReselectionPriorityOk() (*int32, bool)`

GetCellReselectionPriorityOk returns a tuple with the CellReselectionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellReselectionPriority

`func (o *NRFreqRelationSingleAllOfAttributes) SetCellReselectionPriority(v int32)`

SetCellReselectionPriority sets CellReselectionPriority field to given value.

### HasCellReselectionPriority

`func (o *NRFreqRelationSingleAllOfAttributes) HasCellReselectionPriority() bool`

HasCellReselectionPriority returns a boolean if a field has been set.

### GetCellReselectionSubPriority

`func (o *NRFreqRelationSingleAllOfAttributes) GetCellReselectionSubPriority() float32`

GetCellReselectionSubPriority returns the CellReselectionSubPriority field if non-nil, zero value otherwise.

### GetCellReselectionSubPriorityOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetCellReselectionSubPriorityOk() (*float32, bool)`

GetCellReselectionSubPriorityOk returns a tuple with the CellReselectionSubPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellReselectionSubPriority

`func (o *NRFreqRelationSingleAllOfAttributes) SetCellReselectionSubPriority(v float32)`

SetCellReselectionSubPriority sets CellReselectionSubPriority field to given value.

### HasCellReselectionSubPriority

`func (o *NRFreqRelationSingleAllOfAttributes) HasCellReselectionSubPriority() bool`

HasCellReselectionSubPriority returns a boolean if a field has been set.

### GetPMax

`func (o *NRFreqRelationSingleAllOfAttributes) GetPMax() int32`

GetPMax returns the PMax field if non-nil, zero value otherwise.

### GetPMaxOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetPMaxOk() (*int32, bool)`

GetPMaxOk returns a tuple with the PMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPMax

`func (o *NRFreqRelationSingleAllOfAttributes) SetPMax(v int32)`

SetPMax sets PMax field to given value.

### HasPMax

`func (o *NRFreqRelationSingleAllOfAttributes) HasPMax() bool`

HasPMax returns a boolean if a field has been set.

### GetQOffsetFreq

`func (o *NRFreqRelationSingleAllOfAttributes) GetQOffsetFreq() float32`

GetQOffsetFreq returns the QOffsetFreq field if non-nil, zero value otherwise.

### GetQOffsetFreqOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetQOffsetFreqOk() (*float32, bool)`

GetQOffsetFreqOk returns a tuple with the QOffsetFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQOffsetFreq

`func (o *NRFreqRelationSingleAllOfAttributes) SetQOffsetFreq(v float32)`

SetQOffsetFreq sets QOffsetFreq field to given value.

### HasQOffsetFreq

`func (o *NRFreqRelationSingleAllOfAttributes) HasQOffsetFreq() bool`

HasQOffsetFreq returns a boolean if a field has been set.

### GetQQualMin

`func (o *NRFreqRelationSingleAllOfAttributes) GetQQualMin() float32`

GetQQualMin returns the QQualMin field if non-nil, zero value otherwise.

### GetQQualMinOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetQQualMinOk() (*float32, bool)`

GetQQualMinOk returns a tuple with the QQualMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQQualMin

`func (o *NRFreqRelationSingleAllOfAttributes) SetQQualMin(v float32)`

SetQQualMin sets QQualMin field to given value.

### HasQQualMin

`func (o *NRFreqRelationSingleAllOfAttributes) HasQQualMin() bool`

HasQQualMin returns a boolean if a field has been set.

### GetQRxLevMin

`func (o *NRFreqRelationSingleAllOfAttributes) GetQRxLevMin() int32`

GetQRxLevMin returns the QRxLevMin field if non-nil, zero value otherwise.

### GetQRxLevMinOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetQRxLevMinOk() (*int32, bool)`

GetQRxLevMinOk returns a tuple with the QRxLevMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQRxLevMin

`func (o *NRFreqRelationSingleAllOfAttributes) SetQRxLevMin(v int32)`

SetQRxLevMin sets QRxLevMin field to given value.

### HasQRxLevMin

`func (o *NRFreqRelationSingleAllOfAttributes) HasQRxLevMin() bool`

HasQRxLevMin returns a boolean if a field has been set.

### GetThreshXHighP

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXHighP() int32`

GetThreshXHighP returns the ThreshXHighP field if non-nil, zero value otherwise.

### GetThreshXHighPOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXHighPOk() (*int32, bool)`

GetThreshXHighPOk returns a tuple with the ThreshXHighP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXHighP

`func (o *NRFreqRelationSingleAllOfAttributes) SetThreshXHighP(v int32)`

SetThreshXHighP sets ThreshXHighP field to given value.

### HasThreshXHighP

`func (o *NRFreqRelationSingleAllOfAttributes) HasThreshXHighP() bool`

HasThreshXHighP returns a boolean if a field has been set.

### GetThreshXHighQ

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXHighQ() int32`

GetThreshXHighQ returns the ThreshXHighQ field if non-nil, zero value otherwise.

### GetThreshXHighQOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXHighQOk() (*int32, bool)`

GetThreshXHighQOk returns a tuple with the ThreshXHighQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXHighQ

`func (o *NRFreqRelationSingleAllOfAttributes) SetThreshXHighQ(v int32)`

SetThreshXHighQ sets ThreshXHighQ field to given value.

### HasThreshXHighQ

`func (o *NRFreqRelationSingleAllOfAttributes) HasThreshXHighQ() bool`

HasThreshXHighQ returns a boolean if a field has been set.

### GetThreshXLowP

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXLowP() int32`

GetThreshXLowP returns the ThreshXLowP field if non-nil, zero value otherwise.

### GetThreshXLowPOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXLowPOk() (*int32, bool)`

GetThreshXLowPOk returns a tuple with the ThreshXLowP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXLowP

`func (o *NRFreqRelationSingleAllOfAttributes) SetThreshXLowP(v int32)`

SetThreshXLowP sets ThreshXLowP field to given value.

### HasThreshXLowP

`func (o *NRFreqRelationSingleAllOfAttributes) HasThreshXLowP() bool`

HasThreshXLowP returns a boolean if a field has been set.

### GetThreshXLowQ

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXLowQ() int32`

GetThreshXLowQ returns the ThreshXLowQ field if non-nil, zero value otherwise.

### GetThreshXLowQOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetThreshXLowQOk() (*int32, bool)`

GetThreshXLowQOk returns a tuple with the ThreshXLowQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXLowQ

`func (o *NRFreqRelationSingleAllOfAttributes) SetThreshXLowQ(v int32)`

SetThreshXLowQ sets ThreshXLowQ field to given value.

### HasThreshXLowQ

`func (o *NRFreqRelationSingleAllOfAttributes) HasThreshXLowQ() bool`

HasThreshXLowQ returns a boolean if a field has been set.

### GetTReselectionNr

`func (o *NRFreqRelationSingleAllOfAttributes) GetTReselectionNr() int32`

GetTReselectionNr returns the TReselectionNr field if non-nil, zero value otherwise.

### GetTReselectionNrOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetTReselectionNrOk() (*int32, bool)`

GetTReselectionNrOk returns a tuple with the TReselectionNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTReselectionNr

`func (o *NRFreqRelationSingleAllOfAttributes) SetTReselectionNr(v int32)`

SetTReselectionNr sets TReselectionNr field to given value.

### HasTReselectionNr

`func (o *NRFreqRelationSingleAllOfAttributes) HasTReselectionNr() bool`

HasTReselectionNr returns a boolean if a field has been set.

### GetTReselectionNRSfHigh

`func (o *NRFreqRelationSingleAllOfAttributes) GetTReselectionNRSfHigh() TReselectionNRSf`

GetTReselectionNRSfHigh returns the TReselectionNRSfHigh field if non-nil, zero value otherwise.

### GetTReselectionNRSfHighOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetTReselectionNRSfHighOk() (*TReselectionNRSf, bool)`

GetTReselectionNRSfHighOk returns a tuple with the TReselectionNRSfHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTReselectionNRSfHigh

`func (o *NRFreqRelationSingleAllOfAttributes) SetTReselectionNRSfHigh(v TReselectionNRSf)`

SetTReselectionNRSfHigh sets TReselectionNRSfHigh field to given value.

### HasTReselectionNRSfHigh

`func (o *NRFreqRelationSingleAllOfAttributes) HasTReselectionNRSfHigh() bool`

HasTReselectionNRSfHigh returns a boolean if a field has been set.

### GetTReselectionNRSfMedium

`func (o *NRFreqRelationSingleAllOfAttributes) GetTReselectionNRSfMedium() TReselectionNRSf`

GetTReselectionNRSfMedium returns the TReselectionNRSfMedium field if non-nil, zero value otherwise.

### GetTReselectionNRSfMediumOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetTReselectionNRSfMediumOk() (*TReselectionNRSf, bool)`

GetTReselectionNRSfMediumOk returns a tuple with the TReselectionNRSfMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTReselectionNRSfMedium

`func (o *NRFreqRelationSingleAllOfAttributes) SetTReselectionNRSfMedium(v TReselectionNRSf)`

SetTReselectionNRSfMedium sets TReselectionNRSfMedium field to given value.

### HasTReselectionNRSfMedium

`func (o *NRFreqRelationSingleAllOfAttributes) HasTReselectionNRSfMedium() bool`

HasTReselectionNRSfMedium returns a boolean if a field has been set.

### GetNRFrequencyRef

`func (o *NRFreqRelationSingleAllOfAttributes) GetNRFrequencyRef() string`

GetNRFrequencyRef returns the NRFrequencyRef field if non-nil, zero value otherwise.

### GetNRFrequencyRefOk

`func (o *NRFreqRelationSingleAllOfAttributes) GetNRFrequencyRefOk() (*string, bool)`

GetNRFrequencyRefOk returns a tuple with the NRFrequencyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRFrequencyRef

`func (o *NRFreqRelationSingleAllOfAttributes) SetNRFrequencyRef(v string)`

SetNRFrequencyRef sets NRFrequencyRef field to given value.

### HasNRFrequencyRef

`func (o *NRFreqRelationSingleAllOfAttributes) HasNRFrequencyRef() bool`

HasNRFrequencyRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


