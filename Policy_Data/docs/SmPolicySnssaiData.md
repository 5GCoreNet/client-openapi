# SmPolicySnssaiData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**SmPolicyDnnData** | Pointer to [**map[string]SmPolicyDnnData**](SmPolicyDnnData.md) | Session Management Policy data per DNN for all the DNNs of the indicated S-NSSAI. The key of the map is the DNN.  | [optional] 
**UeSliceMbr** | Pointer to [**SliceMbr**](SliceMbr.md) |  | [optional] 

## Methods

### NewSmPolicySnssaiData

`func NewSmPolicySnssaiData(snssai Snssai, ) *SmPolicySnssaiData`

NewSmPolicySnssaiData instantiates a new SmPolicySnssaiData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicySnssaiDataWithDefaults

`func NewSmPolicySnssaiDataWithDefaults() *SmPolicySnssaiData`

NewSmPolicySnssaiDataWithDefaults instantiates a new SmPolicySnssaiData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *SmPolicySnssaiData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SmPolicySnssaiData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SmPolicySnssaiData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetSmPolicyDnnData

`func (o *SmPolicySnssaiData) GetSmPolicyDnnData() map[string]SmPolicyDnnData`

GetSmPolicyDnnData returns the SmPolicyDnnData field if non-nil, zero value otherwise.

### GetSmPolicyDnnDataOk

`func (o *SmPolicySnssaiData) GetSmPolicyDnnDataOk() (*map[string]SmPolicyDnnData, bool)`

GetSmPolicyDnnDataOk returns a tuple with the SmPolicyDnnData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyDnnData

`func (o *SmPolicySnssaiData) SetSmPolicyDnnData(v map[string]SmPolicyDnnData)`

SetSmPolicyDnnData sets SmPolicyDnnData field to given value.

### HasSmPolicyDnnData

`func (o *SmPolicySnssaiData) HasSmPolicyDnnData() bool`

HasSmPolicyDnnData returns a boolean if a field has been set.

### GetUeSliceMbr

`func (o *SmPolicySnssaiData) GetUeSliceMbr() SliceMbr`

GetUeSliceMbr returns the UeSliceMbr field if non-nil, zero value otherwise.

### GetUeSliceMbrOk

`func (o *SmPolicySnssaiData) GetUeSliceMbrOk() (*SliceMbr, bool)`

GetUeSliceMbrOk returns a tuple with the UeSliceMbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSliceMbr

`func (o *SmPolicySnssaiData) SetUeSliceMbr(v SliceMbr)`

SetUeSliceMbr sets UeSliceMbr field to given value.

### HasUeSliceMbr

`func (o *SmPolicySnssaiData) HasUeSliceMbr() bool`

HasUeSliceMbr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


