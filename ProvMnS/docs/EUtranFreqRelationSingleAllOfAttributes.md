# EUtranFreqRelationSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellIndividualOffset** | Pointer to [**CellIndividualOffset**](CellIndividualOffset.md) |  | [optional] 
**BlackListEntry** | Pointer to **[]int32** |  | [optional] 
**BlackListEntryIdleMode** | Pointer to **int32** |  | [optional] 
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
**TReselectionEutran** | Pointer to **int32** |  | [optional] 
**TReselectionNRSfHigh** | Pointer to [**TReselectionNRSf**](TReselectionNRSf.md) |  | [optional] 
**TReselectionNRSfMedium** | Pointer to [**TReselectionNRSf**](TReselectionNRSf.md) |  | [optional] 
**EUTranFrequencyRef** | Pointer to **string** |  | [optional] 

## Methods

### NewEUtranFreqRelationSingleAllOfAttributes

`func NewEUtranFreqRelationSingleAllOfAttributes() *EUtranFreqRelationSingleAllOfAttributes`

NewEUtranFreqRelationSingleAllOfAttributes instantiates a new EUtranFreqRelationSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEUtranFreqRelationSingleAllOfAttributesWithDefaults

`func NewEUtranFreqRelationSingleAllOfAttributesWithDefaults() *EUtranFreqRelationSingleAllOfAttributes`

NewEUtranFreqRelationSingleAllOfAttributesWithDefaults instantiates a new EUtranFreqRelationSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellIndividualOffset

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellIndividualOffset() CellIndividualOffset`

GetCellIndividualOffset returns the CellIndividualOffset field if non-nil, zero value otherwise.

### GetCellIndividualOffsetOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellIndividualOffsetOk() (*CellIndividualOffset, bool)`

GetCellIndividualOffsetOk returns a tuple with the CellIndividualOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellIndividualOffset

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetCellIndividualOffset(v CellIndividualOffset)`

SetCellIndividualOffset sets CellIndividualOffset field to given value.

### HasCellIndividualOffset

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasCellIndividualOffset() bool`

HasCellIndividualOffset returns a boolean if a field has been set.

### GetBlackListEntry

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntry() []int32`

GetBlackListEntry returns the BlackListEntry field if non-nil, zero value otherwise.

### GetBlackListEntryOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntryOk() (*[]int32, bool)`

GetBlackListEntryOk returns a tuple with the BlackListEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackListEntry

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetBlackListEntry(v []int32)`

SetBlackListEntry sets BlackListEntry field to given value.

### HasBlackListEntry

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasBlackListEntry() bool`

HasBlackListEntry returns a boolean if a field has been set.

### GetBlackListEntryIdleMode

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntryIdleMode() int32`

GetBlackListEntryIdleMode returns the BlackListEntryIdleMode field if non-nil, zero value otherwise.

### GetBlackListEntryIdleModeOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntryIdleModeOk() (*int32, bool)`

GetBlackListEntryIdleModeOk returns a tuple with the BlackListEntryIdleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackListEntryIdleMode

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetBlackListEntryIdleMode(v int32)`

SetBlackListEntryIdleMode sets BlackListEntryIdleMode field to given value.

### HasBlackListEntryIdleMode

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasBlackListEntryIdleMode() bool`

HasBlackListEntryIdleMode returns a boolean if a field has been set.

### GetCellReselectionPriority

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionPriority() int32`

GetCellReselectionPriority returns the CellReselectionPriority field if non-nil, zero value otherwise.

### GetCellReselectionPriorityOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionPriorityOk() (*int32, bool)`

GetCellReselectionPriorityOk returns a tuple with the CellReselectionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellReselectionPriority

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetCellReselectionPriority(v int32)`

SetCellReselectionPriority sets CellReselectionPriority field to given value.

### HasCellReselectionPriority

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasCellReselectionPriority() bool`

HasCellReselectionPriority returns a boolean if a field has been set.

### GetCellReselectionSubPriority

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionSubPriority() float32`

GetCellReselectionSubPriority returns the CellReselectionSubPriority field if non-nil, zero value otherwise.

### GetCellReselectionSubPriorityOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionSubPriorityOk() (*float32, bool)`

GetCellReselectionSubPriorityOk returns a tuple with the CellReselectionSubPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellReselectionSubPriority

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetCellReselectionSubPriority(v float32)`

SetCellReselectionSubPriority sets CellReselectionSubPriority field to given value.

### HasCellReselectionSubPriority

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasCellReselectionSubPriority() bool`

HasCellReselectionSubPriority returns a boolean if a field has been set.

### GetPMax

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetPMax() int32`

GetPMax returns the PMax field if non-nil, zero value otherwise.

### GetPMaxOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetPMaxOk() (*int32, bool)`

GetPMaxOk returns a tuple with the PMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPMax

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetPMax(v int32)`

SetPMax sets PMax field to given value.

### HasPMax

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasPMax() bool`

HasPMax returns a boolean if a field has been set.

### GetQOffsetFreq

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetQOffsetFreq() float32`

GetQOffsetFreq returns the QOffsetFreq field if non-nil, zero value otherwise.

### GetQOffsetFreqOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetQOffsetFreqOk() (*float32, bool)`

GetQOffsetFreqOk returns a tuple with the QOffsetFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQOffsetFreq

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetQOffsetFreq(v float32)`

SetQOffsetFreq sets QOffsetFreq field to given value.

### HasQOffsetFreq

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasQOffsetFreq() bool`

HasQOffsetFreq returns a boolean if a field has been set.

### GetQQualMin

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetQQualMin() float32`

GetQQualMin returns the QQualMin field if non-nil, zero value otherwise.

### GetQQualMinOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetQQualMinOk() (*float32, bool)`

GetQQualMinOk returns a tuple with the QQualMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQQualMin

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetQQualMin(v float32)`

SetQQualMin sets QQualMin field to given value.

### HasQQualMin

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasQQualMin() bool`

HasQQualMin returns a boolean if a field has been set.

### GetQRxLevMin

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetQRxLevMin() int32`

GetQRxLevMin returns the QRxLevMin field if non-nil, zero value otherwise.

### GetQRxLevMinOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetQRxLevMinOk() (*int32, bool)`

GetQRxLevMinOk returns a tuple with the QRxLevMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQRxLevMin

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetQRxLevMin(v int32)`

SetQRxLevMin sets QRxLevMin field to given value.

### HasQRxLevMin

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasQRxLevMin() bool`

HasQRxLevMin returns a boolean if a field has been set.

### GetThreshXHighP

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighP() int32`

GetThreshXHighP returns the ThreshXHighP field if non-nil, zero value otherwise.

### GetThreshXHighPOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighPOk() (*int32, bool)`

GetThreshXHighPOk returns a tuple with the ThreshXHighP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXHighP

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXHighP(v int32)`

SetThreshXHighP sets ThreshXHighP field to given value.

### HasThreshXHighP

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXHighP() bool`

HasThreshXHighP returns a boolean if a field has been set.

### GetThreshXHighQ

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighQ() int32`

GetThreshXHighQ returns the ThreshXHighQ field if non-nil, zero value otherwise.

### GetThreshXHighQOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighQOk() (*int32, bool)`

GetThreshXHighQOk returns a tuple with the ThreshXHighQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXHighQ

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXHighQ(v int32)`

SetThreshXHighQ sets ThreshXHighQ field to given value.

### HasThreshXHighQ

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXHighQ() bool`

HasThreshXHighQ returns a boolean if a field has been set.

### GetThreshXLowP

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowP() int32`

GetThreshXLowP returns the ThreshXLowP field if non-nil, zero value otherwise.

### GetThreshXLowPOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowPOk() (*int32, bool)`

GetThreshXLowPOk returns a tuple with the ThreshXLowP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXLowP

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXLowP(v int32)`

SetThreshXLowP sets ThreshXLowP field to given value.

### HasThreshXLowP

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXLowP() bool`

HasThreshXLowP returns a boolean if a field has been set.

### GetThreshXLowQ

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowQ() int32`

GetThreshXLowQ returns the ThreshXLowQ field if non-nil, zero value otherwise.

### GetThreshXLowQOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowQOk() (*int32, bool)`

GetThreshXLowQOk returns a tuple with the ThreshXLowQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshXLowQ

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXLowQ(v int32)`

SetThreshXLowQ sets ThreshXLowQ field to given value.

### HasThreshXLowQ

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXLowQ() bool`

HasThreshXLowQ returns a boolean if a field has been set.

### GetTReselectionEutran

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionEutran() int32`

GetTReselectionEutran returns the TReselectionEutran field if non-nil, zero value otherwise.

### GetTReselectionEutranOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionEutranOk() (*int32, bool)`

GetTReselectionEutranOk returns a tuple with the TReselectionEutran field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTReselectionEutran

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetTReselectionEutran(v int32)`

SetTReselectionEutran sets TReselectionEutran field to given value.

### HasTReselectionEutran

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasTReselectionEutran() bool`

HasTReselectionEutran returns a boolean if a field has been set.

### GetTReselectionNRSfHigh

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfHigh() TReselectionNRSf`

GetTReselectionNRSfHigh returns the TReselectionNRSfHigh field if non-nil, zero value otherwise.

### GetTReselectionNRSfHighOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfHighOk() (*TReselectionNRSf, bool)`

GetTReselectionNRSfHighOk returns a tuple with the TReselectionNRSfHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTReselectionNRSfHigh

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetTReselectionNRSfHigh(v TReselectionNRSf)`

SetTReselectionNRSfHigh sets TReselectionNRSfHigh field to given value.

### HasTReselectionNRSfHigh

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasTReselectionNRSfHigh() bool`

HasTReselectionNRSfHigh returns a boolean if a field has been set.

### GetTReselectionNRSfMedium

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfMedium() TReselectionNRSf`

GetTReselectionNRSfMedium returns the TReselectionNRSfMedium field if non-nil, zero value otherwise.

### GetTReselectionNRSfMediumOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfMediumOk() (*TReselectionNRSf, bool)`

GetTReselectionNRSfMediumOk returns a tuple with the TReselectionNRSfMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTReselectionNRSfMedium

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetTReselectionNRSfMedium(v TReselectionNRSf)`

SetTReselectionNRSfMedium sets TReselectionNRSfMedium field to given value.

### HasTReselectionNRSfMedium

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasTReselectionNRSfMedium() bool`

HasTReselectionNRSfMedium returns a boolean if a field has been set.

### GetEUTranFrequencyRef

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetEUTranFrequencyRef() string`

GetEUTranFrequencyRef returns the EUTranFrequencyRef field if non-nil, zero value otherwise.

### GetEUTranFrequencyRefOk

`func (o *EUtranFreqRelationSingleAllOfAttributes) GetEUTranFrequencyRefOk() (*string, bool)`

GetEUTranFrequencyRefOk returns a tuple with the EUTranFrequencyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUTranFrequencyRef

`func (o *EUtranFreqRelationSingleAllOfAttributes) SetEUTranFrequencyRef(v string)`

SetEUTranFrequencyRef sets EUTranFrequencyRef field to given value.

### HasEUTranFrequencyRef

`func (o *EUtranFreqRelationSingleAllOfAttributes) HasEUTranFrequencyRef() bool`

HasEUTranFrequencyRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


