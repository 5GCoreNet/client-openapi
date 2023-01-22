# QosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosFlowsAddModRequestList** | Pointer to [**[]QosFlowAddModifyRequestItem**](QosFlowAddModifyRequestItem.md) |  | [optional] 
**QosFlowsRelRequestList** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewQosInfo

`func NewQosInfo() *QosInfo`

NewQosInfo instantiates a new QosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosInfoWithDefaults

`func NewQosInfoWithDefaults() *QosInfo`

NewQosInfoWithDefaults instantiates a new QosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosFlowsAddModRequestList

`func (o *QosInfo) GetQosFlowsAddModRequestList() []QosFlowAddModifyRequestItem`

GetQosFlowsAddModRequestList returns the QosFlowsAddModRequestList field if non-nil, zero value otherwise.

### GetQosFlowsAddModRequestListOk

`func (o *QosInfo) GetQosFlowsAddModRequestListOk() (*[]QosFlowAddModifyRequestItem, bool)`

GetQosFlowsAddModRequestListOk returns a tuple with the QosFlowsAddModRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsAddModRequestList

`func (o *QosInfo) SetQosFlowsAddModRequestList(v []QosFlowAddModifyRequestItem)`

SetQosFlowsAddModRequestList sets QosFlowsAddModRequestList field to given value.

### HasQosFlowsAddModRequestList

`func (o *QosInfo) HasQosFlowsAddModRequestList() bool`

HasQosFlowsAddModRequestList returns a boolean if a field has been set.

### GetQosFlowsRelRequestList

`func (o *QosInfo) GetQosFlowsRelRequestList() []int32`

GetQosFlowsRelRequestList returns the QosFlowsRelRequestList field if non-nil, zero value otherwise.

### GetQosFlowsRelRequestListOk

`func (o *QosInfo) GetQosFlowsRelRequestListOk() (*[]int32, bool)`

GetQosFlowsRelRequestListOk returns a tuple with the QosFlowsRelRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsRelRequestList

`func (o *QosInfo) SetQosFlowsRelRequestList(v []int32)`

SetQosFlowsRelRequestList sets QosFlowsRelRequestList field to given value.

### HasQosFlowsRelRequestList

`func (o *QosInfo) HasQosFlowsRelRequestList() bool`

HasQosFlowsRelRequestList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


