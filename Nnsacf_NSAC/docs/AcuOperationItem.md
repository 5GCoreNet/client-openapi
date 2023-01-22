# AcuOperationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateFlag** | [**AcuFlag**](AcuFlag.md) |  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 

## Methods

### NewAcuOperationItem

`func NewAcuOperationItem(updateFlag AcuFlag, snssai Snssai, ) *AcuOperationItem`

NewAcuOperationItem instantiates a new AcuOperationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcuOperationItemWithDefaults

`func NewAcuOperationItemWithDefaults() *AcuOperationItem`

NewAcuOperationItemWithDefaults instantiates a new AcuOperationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateFlag

`func (o *AcuOperationItem) GetUpdateFlag() AcuFlag`

GetUpdateFlag returns the UpdateFlag field if non-nil, zero value otherwise.

### GetUpdateFlagOk

`func (o *AcuOperationItem) GetUpdateFlagOk() (*AcuFlag, bool)`

GetUpdateFlagOk returns a tuple with the UpdateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFlag

`func (o *AcuOperationItem) SetUpdateFlag(v AcuFlag)`

SetUpdateFlag sets UpdateFlag field to given value.


### GetSnssai

`func (o *AcuOperationItem) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AcuOperationItem) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AcuOperationItem) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetPlmnId

`func (o *AcuOperationItem) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *AcuOperationItem) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *AcuOperationItem) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *AcuOperationItem) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


