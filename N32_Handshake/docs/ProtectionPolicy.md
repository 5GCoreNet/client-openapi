# ProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiIeMappingList** | [**[]ApiIeMapping**](ApiIeMapping.md) |  | 
**DataTypeEncPolicy** | Pointer to [**[]IeType**](IeType.md) |  | [optional] 

## Methods

### NewProtectionPolicy

`func NewProtectionPolicy(apiIeMappingList []ApiIeMapping, ) *ProtectionPolicy`

NewProtectionPolicy instantiates a new ProtectionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionPolicyWithDefaults

`func NewProtectionPolicyWithDefaults() *ProtectionPolicy`

NewProtectionPolicyWithDefaults instantiates a new ProtectionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiIeMappingList

`func (o *ProtectionPolicy) GetApiIeMappingList() []ApiIeMapping`

GetApiIeMappingList returns the ApiIeMappingList field if non-nil, zero value otherwise.

### GetApiIeMappingListOk

`func (o *ProtectionPolicy) GetApiIeMappingListOk() (*[]ApiIeMapping, bool)`

GetApiIeMappingListOk returns a tuple with the ApiIeMappingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiIeMappingList

`func (o *ProtectionPolicy) SetApiIeMappingList(v []ApiIeMapping)`

SetApiIeMappingList sets ApiIeMappingList field to given value.


### GetDataTypeEncPolicy

`func (o *ProtectionPolicy) GetDataTypeEncPolicy() []IeType`

GetDataTypeEncPolicy returns the DataTypeEncPolicy field if non-nil, zero value otherwise.

### GetDataTypeEncPolicyOk

`func (o *ProtectionPolicy) GetDataTypeEncPolicyOk() (*[]IeType, bool)`

GetDataTypeEncPolicyOk returns a tuple with the DataTypeEncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeEncPolicy

`func (o *ProtectionPolicy) SetDataTypeEncPolicy(v []IeType)`

SetDataTypeEncPolicy sets DataTypeEncPolicy field to given value.

### HasDataTypeEncPolicy

`func (o *ProtectionPolicy) HasDataTypeEncPolicy() bool`

HasDataTypeEncPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


