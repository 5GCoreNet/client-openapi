# UpfCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | **string** |  | 
**SmfServingArea** | Pointer to **[]string** |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 

## Methods

### NewUpfCond

`func NewUpfCond(conditionType string, ) *UpfCond`

NewUpfCond instantiates a new UpfCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpfCondWithDefaults

`func NewUpfCondWithDefaults() *UpfCond`

NewUpfCondWithDefaults instantiates a new UpfCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *UpfCond) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *UpfCond) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *UpfCond) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetSmfServingArea

`func (o *UpfCond) GetSmfServingArea() []string`

GetSmfServingArea returns the SmfServingArea field if non-nil, zero value otherwise.

### GetSmfServingAreaOk

`func (o *UpfCond) GetSmfServingAreaOk() (*[]string, bool)`

GetSmfServingAreaOk returns a tuple with the SmfServingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServingArea

`func (o *UpfCond) SetSmfServingArea(v []string)`

SetSmfServingArea sets SmfServingArea field to given value.

### HasSmfServingArea

`func (o *UpfCond) HasSmfServingArea() bool`

HasSmfServingArea returns a boolean if a field has been set.

### GetTaiList

`func (o *UpfCond) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *UpfCond) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *UpfCond) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *UpfCond) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


