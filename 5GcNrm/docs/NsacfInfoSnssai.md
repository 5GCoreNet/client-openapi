# NsacfInfoSnssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnssaiInfo** | Pointer to [**SnssaiInfo**](SnssaiInfo.md) |  | [optional] 
**IsSubjectToNsac** | Pointer to **bool** |  | [optional] 
**MaxNumberofUEs** | Pointer to **int32** |  | [optional] 
**EACMode** | Pointer to **string** |  | [optional] 
**ActiveEacThreshhold** | Pointer to **int32** |  | [optional] 
**DeactiveEacThreshhold** | Pointer to **int32** |  | [optional] 
**NumberofUEs** | Pointer to **int32** |  | [optional] 
**UEIdList** | Pointer to **[]string** |  | [optional] 
**MaxNumberofPDUSessions** | Pointer to **int32** |  | [optional] 

## Methods

### NewNsacfInfoSnssai

`func NewNsacfInfoSnssai() *NsacfInfoSnssai`

NewNsacfInfoSnssai instantiates a new NsacfInfoSnssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsacfInfoSnssaiWithDefaults

`func NewNsacfInfoSnssaiWithDefaults() *NsacfInfoSnssai`

NewNsacfInfoSnssaiWithDefaults instantiates a new NsacfInfoSnssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssaiInfo

`func (o *NsacfInfoSnssai) GetSnssaiInfo() SnssaiInfo`

GetSnssaiInfo returns the SnssaiInfo field if non-nil, zero value otherwise.

### GetSnssaiInfoOk

`func (o *NsacfInfoSnssai) GetSnssaiInfoOk() (*SnssaiInfo, bool)`

GetSnssaiInfoOk returns a tuple with the SnssaiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiInfo

`func (o *NsacfInfoSnssai) SetSnssaiInfo(v SnssaiInfo)`

SetSnssaiInfo sets SnssaiInfo field to given value.

### HasSnssaiInfo

`func (o *NsacfInfoSnssai) HasSnssaiInfo() bool`

HasSnssaiInfo returns a boolean if a field has been set.

### GetIsSubjectToNsac

`func (o *NsacfInfoSnssai) GetIsSubjectToNsac() bool`

GetIsSubjectToNsac returns the IsSubjectToNsac field if non-nil, zero value otherwise.

### GetIsSubjectToNsacOk

`func (o *NsacfInfoSnssai) GetIsSubjectToNsacOk() (*bool, bool)`

GetIsSubjectToNsacOk returns a tuple with the IsSubjectToNsac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubjectToNsac

`func (o *NsacfInfoSnssai) SetIsSubjectToNsac(v bool)`

SetIsSubjectToNsac sets IsSubjectToNsac field to given value.

### HasIsSubjectToNsac

`func (o *NsacfInfoSnssai) HasIsSubjectToNsac() bool`

HasIsSubjectToNsac returns a boolean if a field has been set.

### GetMaxNumberofUEs

`func (o *NsacfInfoSnssai) GetMaxNumberofUEs() int32`

GetMaxNumberofUEs returns the MaxNumberofUEs field if non-nil, zero value otherwise.

### GetMaxNumberofUEsOk

`func (o *NsacfInfoSnssai) GetMaxNumberofUEsOk() (*int32, bool)`

GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofUEs

`func (o *NsacfInfoSnssai) SetMaxNumberofUEs(v int32)`

SetMaxNumberofUEs sets MaxNumberofUEs field to given value.

### HasMaxNumberofUEs

`func (o *NsacfInfoSnssai) HasMaxNumberofUEs() bool`

HasMaxNumberofUEs returns a boolean if a field has been set.

### GetEACMode

`func (o *NsacfInfoSnssai) GetEACMode() string`

GetEACMode returns the EACMode field if non-nil, zero value otherwise.

### GetEACModeOk

`func (o *NsacfInfoSnssai) GetEACModeOk() (*string, bool)`

GetEACModeOk returns a tuple with the EACMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEACMode

`func (o *NsacfInfoSnssai) SetEACMode(v string)`

SetEACMode sets EACMode field to given value.

### HasEACMode

`func (o *NsacfInfoSnssai) HasEACMode() bool`

HasEACMode returns a boolean if a field has been set.

### GetActiveEacThreshhold

`func (o *NsacfInfoSnssai) GetActiveEacThreshhold() int32`

GetActiveEacThreshhold returns the ActiveEacThreshhold field if non-nil, zero value otherwise.

### GetActiveEacThreshholdOk

`func (o *NsacfInfoSnssai) GetActiveEacThreshholdOk() (*int32, bool)`

GetActiveEacThreshholdOk returns a tuple with the ActiveEacThreshhold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEacThreshhold

`func (o *NsacfInfoSnssai) SetActiveEacThreshhold(v int32)`

SetActiveEacThreshhold sets ActiveEacThreshhold field to given value.

### HasActiveEacThreshhold

`func (o *NsacfInfoSnssai) HasActiveEacThreshhold() bool`

HasActiveEacThreshhold returns a boolean if a field has been set.

### GetDeactiveEacThreshhold

`func (o *NsacfInfoSnssai) GetDeactiveEacThreshhold() int32`

GetDeactiveEacThreshhold returns the DeactiveEacThreshhold field if non-nil, zero value otherwise.

### GetDeactiveEacThreshholdOk

`func (o *NsacfInfoSnssai) GetDeactiveEacThreshholdOk() (*int32, bool)`

GetDeactiveEacThreshholdOk returns a tuple with the DeactiveEacThreshhold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactiveEacThreshhold

`func (o *NsacfInfoSnssai) SetDeactiveEacThreshhold(v int32)`

SetDeactiveEacThreshhold sets DeactiveEacThreshhold field to given value.

### HasDeactiveEacThreshhold

`func (o *NsacfInfoSnssai) HasDeactiveEacThreshhold() bool`

HasDeactiveEacThreshhold returns a boolean if a field has been set.

### GetNumberofUEs

`func (o *NsacfInfoSnssai) GetNumberofUEs() int32`

GetNumberofUEs returns the NumberofUEs field if non-nil, zero value otherwise.

### GetNumberofUEsOk

`func (o *NsacfInfoSnssai) GetNumberofUEsOk() (*int32, bool)`

GetNumberofUEsOk returns a tuple with the NumberofUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberofUEs

`func (o *NsacfInfoSnssai) SetNumberofUEs(v int32)`

SetNumberofUEs sets NumberofUEs field to given value.

### HasNumberofUEs

`func (o *NsacfInfoSnssai) HasNumberofUEs() bool`

HasNumberofUEs returns a boolean if a field has been set.

### GetUEIdList

`func (o *NsacfInfoSnssai) GetUEIdList() []string`

GetUEIdList returns the UEIdList field if non-nil, zero value otherwise.

### GetUEIdListOk

`func (o *NsacfInfoSnssai) GetUEIdListOk() (*[]string, bool)`

GetUEIdListOk returns a tuple with the UEIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUEIdList

`func (o *NsacfInfoSnssai) SetUEIdList(v []string)`

SetUEIdList sets UEIdList field to given value.

### HasUEIdList

`func (o *NsacfInfoSnssai) HasUEIdList() bool`

HasUEIdList returns a boolean if a field has been set.

### GetMaxNumberofPDUSessions

`func (o *NsacfInfoSnssai) GetMaxNumberofPDUSessions() int32`

GetMaxNumberofPDUSessions returns the MaxNumberofPDUSessions field if non-nil, zero value otherwise.

### GetMaxNumberofPDUSessionsOk

`func (o *NsacfInfoSnssai) GetMaxNumberofPDUSessionsOk() (*int32, bool)`

GetMaxNumberofPDUSessionsOk returns a tuple with the MaxNumberofPDUSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofPDUSessions

`func (o *NsacfInfoSnssai) SetMaxNumberofPDUSessions(v int32)`

SetMaxNumberofPDUSessions sets MaxNumberofPDUSessions field to given value.

### HasMaxNumberofPDUSessions

`func (o *NsacfInfoSnssai) HasMaxNumberofPDUSessions() bool`

HasMaxNumberofPDUSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


