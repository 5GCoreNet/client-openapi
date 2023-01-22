# ContextUpdateRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N2MbsSmInfoList** | Pointer to [**[]N2MbsSmInfo**](N2MbsSmInfo.md) |  | [optional] 
**OperationStatus** | Pointer to [**OperationStatus**](OperationStatus.md) |  | [optional] 

## Methods

### NewContextUpdateRspData

`func NewContextUpdateRspData() *ContextUpdateRspData`

NewContextUpdateRspData instantiates a new ContextUpdateRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextUpdateRspDataWithDefaults

`func NewContextUpdateRspDataWithDefaults() *ContextUpdateRspData`

NewContextUpdateRspDataWithDefaults instantiates a new ContextUpdateRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN2MbsSmInfoList

`func (o *ContextUpdateRspData) GetN2MbsSmInfoList() []N2MbsSmInfo`

GetN2MbsSmInfoList returns the N2MbsSmInfoList field if non-nil, zero value otherwise.

### GetN2MbsSmInfoListOk

`func (o *ContextUpdateRspData) GetN2MbsSmInfoListOk() (*[]N2MbsSmInfo, bool)`

GetN2MbsSmInfoListOk returns a tuple with the N2MbsSmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsSmInfoList

`func (o *ContextUpdateRspData) SetN2MbsSmInfoList(v []N2MbsSmInfo)`

SetN2MbsSmInfoList sets N2MbsSmInfoList field to given value.

### HasN2MbsSmInfoList

`func (o *ContextUpdateRspData) HasN2MbsSmInfoList() bool`

HasN2MbsSmInfoList returns a boolean if a field has been set.

### GetOperationStatus

`func (o *ContextUpdateRspData) GetOperationStatus() OperationStatus`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *ContextUpdateRspData) GetOperationStatusOk() (*OperationStatus, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *ContextUpdateRspData) SetOperationStatus(v OperationStatus)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *ContextUpdateRspData) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


