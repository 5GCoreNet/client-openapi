# TargetDnaiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**SmfSelectionType** | [**SmfSelectionType**](SmfSelectionType.md) |  | 

## Methods

### NewTargetDnaiInfo

`func NewTargetDnaiInfo(smfSelectionType SmfSelectionType, ) *TargetDnaiInfo`

NewTargetDnaiInfo instantiates a new TargetDnaiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetDnaiInfoWithDefaults

`func NewTargetDnaiInfoWithDefaults() *TargetDnaiInfo`

NewTargetDnaiInfoWithDefaults instantiates a new TargetDnaiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetDnai

`func (o *TargetDnaiInfo) GetTargetDnai() string`

GetTargetDnai returns the TargetDnai field if non-nil, zero value otherwise.

### GetTargetDnaiOk

`func (o *TargetDnaiInfo) GetTargetDnaiOk() (*string, bool)`

GetTargetDnaiOk returns a tuple with the TargetDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnai

`func (o *TargetDnaiInfo) SetTargetDnai(v string)`

SetTargetDnai sets TargetDnai field to given value.

### HasTargetDnai

`func (o *TargetDnaiInfo) HasTargetDnai() bool`

HasTargetDnai returns a boolean if a field has been set.

### GetSmfSelectionType

`func (o *TargetDnaiInfo) GetSmfSelectionType() SmfSelectionType`

GetSmfSelectionType returns the SmfSelectionType field if non-nil, zero value otherwise.

### GetSmfSelectionTypeOk

`func (o *TargetDnaiInfo) GetSmfSelectionTypeOk() (*SmfSelectionType, bool)`

GetSmfSelectionTypeOk returns a tuple with the SmfSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelectionType

`func (o *TargetDnaiInfo) SetSmfSelectionType(v SmfSelectionType)`

SetSmfSelectionType sets SmfSelectionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


