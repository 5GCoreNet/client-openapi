# ServAttrCom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**Tagging** | Pointer to **[]string** |  | [optional] 
**Exposure** | Pointer to [**Exposure**](Exposure.md) |  | [optional] 

## Methods

### NewServAttrCom

`func NewServAttrCom() *ServAttrCom`

NewServAttrCom instantiates a new ServAttrCom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServAttrComWithDefaults

`func NewServAttrComWithDefaults() *ServAttrCom`

NewServAttrComWithDefaults instantiates a new ServAttrCom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ServAttrCom) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ServAttrCom) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ServAttrCom) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ServAttrCom) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTagging

`func (o *ServAttrCom) GetTagging() []string`

GetTagging returns the Tagging field if non-nil, zero value otherwise.

### GetTaggingOk

`func (o *ServAttrCom) GetTaggingOk() (*[]string, bool)`

GetTaggingOk returns a tuple with the Tagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagging

`func (o *ServAttrCom) SetTagging(v []string)`

SetTagging sets Tagging field to given value.

### HasTagging

`func (o *ServAttrCom) HasTagging() bool`

HasTagging returns a boolean if a field has been set.

### GetExposure

`func (o *ServAttrCom) GetExposure() Exposure`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *ServAttrCom) GetExposureOk() (*Exposure, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *ServAttrCom) SetExposure(v Exposure)`

SetExposure sets Exposure field to given value.

### HasExposure

`func (o *ServAttrCom) HasExposure() bool`

HasExposure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


