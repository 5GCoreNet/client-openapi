# NRCellRelationSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NRTCI** | Pointer to **int32** |  | [optional] 
**CellIndividualOffset** | Pointer to [**CellIndividualOffset**](CellIndividualOffset.md) |  | [optional] 
**AdjacentNRCellRef** | Pointer to **string** |  | [optional] 
**NRFreqRelationRef** | Pointer to **string** |  | [optional] 
**IsRemoveAllowed** | Pointer to **bool** |  | [optional] 
**IsHOAllowed** | Pointer to **bool** |  | [optional] 
**IsESCoveredBy** | Pointer to [**IsESCoveredBy**](IsESCoveredBy.md) |  | [optional] 
**IsENDCAllowed** | Pointer to **bool** |  | [optional] 
**IsMLBAllowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewNRCellRelationSingleAllOfAttributes

`func NewNRCellRelationSingleAllOfAttributes() *NRCellRelationSingleAllOfAttributes`

NewNRCellRelationSingleAllOfAttributes instantiates a new NRCellRelationSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNRCellRelationSingleAllOfAttributesWithDefaults

`func NewNRCellRelationSingleAllOfAttributesWithDefaults() *NRCellRelationSingleAllOfAttributes`

NewNRCellRelationSingleAllOfAttributesWithDefaults instantiates a new NRCellRelationSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNRTCI

`func (o *NRCellRelationSingleAllOfAttributes) GetNRTCI() int32`

GetNRTCI returns the NRTCI field if non-nil, zero value otherwise.

### GetNRTCIOk

`func (o *NRCellRelationSingleAllOfAttributes) GetNRTCIOk() (*int32, bool)`

GetNRTCIOk returns a tuple with the NRTCI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRTCI

`func (o *NRCellRelationSingleAllOfAttributes) SetNRTCI(v int32)`

SetNRTCI sets NRTCI field to given value.

### HasNRTCI

`func (o *NRCellRelationSingleAllOfAttributes) HasNRTCI() bool`

HasNRTCI returns a boolean if a field has been set.

### GetCellIndividualOffset

`func (o *NRCellRelationSingleAllOfAttributes) GetCellIndividualOffset() CellIndividualOffset`

GetCellIndividualOffset returns the CellIndividualOffset field if non-nil, zero value otherwise.

### GetCellIndividualOffsetOk

`func (o *NRCellRelationSingleAllOfAttributes) GetCellIndividualOffsetOk() (*CellIndividualOffset, bool)`

GetCellIndividualOffsetOk returns a tuple with the CellIndividualOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellIndividualOffset

`func (o *NRCellRelationSingleAllOfAttributes) SetCellIndividualOffset(v CellIndividualOffset)`

SetCellIndividualOffset sets CellIndividualOffset field to given value.

### HasCellIndividualOffset

`func (o *NRCellRelationSingleAllOfAttributes) HasCellIndividualOffset() bool`

HasCellIndividualOffset returns a boolean if a field has been set.

### GetAdjacentNRCellRef

`func (o *NRCellRelationSingleAllOfAttributes) GetAdjacentNRCellRef() string`

GetAdjacentNRCellRef returns the AdjacentNRCellRef field if non-nil, zero value otherwise.

### GetAdjacentNRCellRefOk

`func (o *NRCellRelationSingleAllOfAttributes) GetAdjacentNRCellRefOk() (*string, bool)`

GetAdjacentNRCellRefOk returns a tuple with the AdjacentNRCellRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentNRCellRef

`func (o *NRCellRelationSingleAllOfAttributes) SetAdjacentNRCellRef(v string)`

SetAdjacentNRCellRef sets AdjacentNRCellRef field to given value.

### HasAdjacentNRCellRef

`func (o *NRCellRelationSingleAllOfAttributes) HasAdjacentNRCellRef() bool`

HasAdjacentNRCellRef returns a boolean if a field has been set.

### GetNRFreqRelationRef

`func (o *NRCellRelationSingleAllOfAttributes) GetNRFreqRelationRef() string`

GetNRFreqRelationRef returns the NRFreqRelationRef field if non-nil, zero value otherwise.

### GetNRFreqRelationRefOk

`func (o *NRCellRelationSingleAllOfAttributes) GetNRFreqRelationRefOk() (*string, bool)`

GetNRFreqRelationRefOk returns a tuple with the NRFreqRelationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRFreqRelationRef

`func (o *NRCellRelationSingleAllOfAttributes) SetNRFreqRelationRef(v string)`

SetNRFreqRelationRef sets NRFreqRelationRef field to given value.

### HasNRFreqRelationRef

`func (o *NRCellRelationSingleAllOfAttributes) HasNRFreqRelationRef() bool`

HasNRFreqRelationRef returns a boolean if a field has been set.

### GetIsRemoveAllowed

`func (o *NRCellRelationSingleAllOfAttributes) GetIsRemoveAllowed() bool`

GetIsRemoveAllowed returns the IsRemoveAllowed field if non-nil, zero value otherwise.

### GetIsRemoveAllowedOk

`func (o *NRCellRelationSingleAllOfAttributes) GetIsRemoveAllowedOk() (*bool, bool)`

GetIsRemoveAllowedOk returns a tuple with the IsRemoveAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoveAllowed

`func (o *NRCellRelationSingleAllOfAttributes) SetIsRemoveAllowed(v bool)`

SetIsRemoveAllowed sets IsRemoveAllowed field to given value.

### HasIsRemoveAllowed

`func (o *NRCellRelationSingleAllOfAttributes) HasIsRemoveAllowed() bool`

HasIsRemoveAllowed returns a boolean if a field has been set.

### GetIsHOAllowed

`func (o *NRCellRelationSingleAllOfAttributes) GetIsHOAllowed() bool`

GetIsHOAllowed returns the IsHOAllowed field if non-nil, zero value otherwise.

### GetIsHOAllowedOk

`func (o *NRCellRelationSingleAllOfAttributes) GetIsHOAllowedOk() (*bool, bool)`

GetIsHOAllowedOk returns a tuple with the IsHOAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHOAllowed

`func (o *NRCellRelationSingleAllOfAttributes) SetIsHOAllowed(v bool)`

SetIsHOAllowed sets IsHOAllowed field to given value.

### HasIsHOAllowed

`func (o *NRCellRelationSingleAllOfAttributes) HasIsHOAllowed() bool`

HasIsHOAllowed returns a boolean if a field has been set.

### GetIsESCoveredBy

`func (o *NRCellRelationSingleAllOfAttributes) GetIsESCoveredBy() IsESCoveredBy`

GetIsESCoveredBy returns the IsESCoveredBy field if non-nil, zero value otherwise.

### GetIsESCoveredByOk

`func (o *NRCellRelationSingleAllOfAttributes) GetIsESCoveredByOk() (*IsESCoveredBy, bool)`

GetIsESCoveredByOk returns a tuple with the IsESCoveredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsESCoveredBy

`func (o *NRCellRelationSingleAllOfAttributes) SetIsESCoveredBy(v IsESCoveredBy)`

SetIsESCoveredBy sets IsESCoveredBy field to given value.

### HasIsESCoveredBy

`func (o *NRCellRelationSingleAllOfAttributes) HasIsESCoveredBy() bool`

HasIsESCoveredBy returns a boolean if a field has been set.

### GetIsENDCAllowed

`func (o *NRCellRelationSingleAllOfAttributes) GetIsENDCAllowed() bool`

GetIsENDCAllowed returns the IsENDCAllowed field if non-nil, zero value otherwise.

### GetIsENDCAllowedOk

`func (o *NRCellRelationSingleAllOfAttributes) GetIsENDCAllowedOk() (*bool, bool)`

GetIsENDCAllowedOk returns a tuple with the IsENDCAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsENDCAllowed

`func (o *NRCellRelationSingleAllOfAttributes) SetIsENDCAllowed(v bool)`

SetIsENDCAllowed sets IsENDCAllowed field to given value.

### HasIsENDCAllowed

`func (o *NRCellRelationSingleAllOfAttributes) HasIsENDCAllowed() bool`

HasIsENDCAllowed returns a boolean if a field has been set.

### GetIsMLBAllowed

`func (o *NRCellRelationSingleAllOfAttributes) GetIsMLBAllowed() bool`

GetIsMLBAllowed returns the IsMLBAllowed field if non-nil, zero value otherwise.

### GetIsMLBAllowedOk

`func (o *NRCellRelationSingleAllOfAttributes) GetIsMLBAllowedOk() (*bool, bool)`

GetIsMLBAllowedOk returns a tuple with the IsMLBAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMLBAllowed

`func (o *NRCellRelationSingleAllOfAttributes) SetIsMLBAllowed(v bool)`

SetIsMLBAllowed sets IsMLBAllowed field to given value.

### HasIsMLBAllowed

`func (o *NRCellRelationSingleAllOfAttributes) HasIsMLBAllowed() bool`

HasIsMLBAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


