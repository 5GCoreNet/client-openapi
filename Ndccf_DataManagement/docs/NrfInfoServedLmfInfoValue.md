# NrfInfoServedLmfInfoValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServingClientTypes** | Pointer to [**[]ExternalClientType**](ExternalClientType.md) |  | [optional] 
**LmfId** | Pointer to **string** | LMF identification. | [optional] 
**ServingAccessTypes** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**ServingAnNodeTypes** | Pointer to [**[]AnNodeType**](AnNodeType.md) |  | [optional] 
**ServingRatTypes** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**SupportedGADShapes** | Pointer to [**[]SupportedGADShapes**](SupportedGADShapes.md) |  | [optional] 

## Methods

### NewNrfInfoServedLmfInfoValue

`func NewNrfInfoServedLmfInfoValue() *NrfInfoServedLmfInfoValue`

NewNrfInfoServedLmfInfoValue instantiates a new NrfInfoServedLmfInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedLmfInfoValueWithDefaults

`func NewNrfInfoServedLmfInfoValueWithDefaults() *NrfInfoServedLmfInfoValue`

NewNrfInfoServedLmfInfoValueWithDefaults instantiates a new NrfInfoServedLmfInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServingClientTypes

`func (o *NrfInfoServedLmfInfoValue) GetServingClientTypes() []ExternalClientType`

GetServingClientTypes returns the ServingClientTypes field if non-nil, zero value otherwise.

### GetServingClientTypesOk

`func (o *NrfInfoServedLmfInfoValue) GetServingClientTypesOk() (*[]ExternalClientType, bool)`

GetServingClientTypesOk returns a tuple with the ServingClientTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingClientTypes

`func (o *NrfInfoServedLmfInfoValue) SetServingClientTypes(v []ExternalClientType)`

SetServingClientTypes sets ServingClientTypes field to given value.

### HasServingClientTypes

`func (o *NrfInfoServedLmfInfoValue) HasServingClientTypes() bool`

HasServingClientTypes returns a boolean if a field has been set.

### GetLmfId

`func (o *NrfInfoServedLmfInfoValue) GetLmfId() string`

GetLmfId returns the LmfId field if non-nil, zero value otherwise.

### GetLmfIdOk

`func (o *NrfInfoServedLmfInfoValue) GetLmfIdOk() (*string, bool)`

GetLmfIdOk returns a tuple with the LmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfId

`func (o *NrfInfoServedLmfInfoValue) SetLmfId(v string)`

SetLmfId sets LmfId field to given value.

### HasLmfId

`func (o *NrfInfoServedLmfInfoValue) HasLmfId() bool`

HasLmfId returns a boolean if a field has been set.

### GetServingAccessTypes

`func (o *NrfInfoServedLmfInfoValue) GetServingAccessTypes() []AccessType`

GetServingAccessTypes returns the ServingAccessTypes field if non-nil, zero value otherwise.

### GetServingAccessTypesOk

`func (o *NrfInfoServedLmfInfoValue) GetServingAccessTypesOk() (*[]AccessType, bool)`

GetServingAccessTypesOk returns a tuple with the ServingAccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingAccessTypes

`func (o *NrfInfoServedLmfInfoValue) SetServingAccessTypes(v []AccessType)`

SetServingAccessTypes sets ServingAccessTypes field to given value.

### HasServingAccessTypes

`func (o *NrfInfoServedLmfInfoValue) HasServingAccessTypes() bool`

HasServingAccessTypes returns a boolean if a field has been set.

### GetServingAnNodeTypes

`func (o *NrfInfoServedLmfInfoValue) GetServingAnNodeTypes() []AnNodeType`

GetServingAnNodeTypes returns the ServingAnNodeTypes field if non-nil, zero value otherwise.

### GetServingAnNodeTypesOk

`func (o *NrfInfoServedLmfInfoValue) GetServingAnNodeTypesOk() (*[]AnNodeType, bool)`

GetServingAnNodeTypesOk returns a tuple with the ServingAnNodeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingAnNodeTypes

`func (o *NrfInfoServedLmfInfoValue) SetServingAnNodeTypes(v []AnNodeType)`

SetServingAnNodeTypes sets ServingAnNodeTypes field to given value.

### HasServingAnNodeTypes

`func (o *NrfInfoServedLmfInfoValue) HasServingAnNodeTypes() bool`

HasServingAnNodeTypes returns a boolean if a field has been set.

### GetServingRatTypes

`func (o *NrfInfoServedLmfInfoValue) GetServingRatTypes() []RatType`

GetServingRatTypes returns the ServingRatTypes field if non-nil, zero value otherwise.

### GetServingRatTypesOk

`func (o *NrfInfoServedLmfInfoValue) GetServingRatTypesOk() (*[]RatType, bool)`

GetServingRatTypesOk returns a tuple with the ServingRatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingRatTypes

`func (o *NrfInfoServedLmfInfoValue) SetServingRatTypes(v []RatType)`

SetServingRatTypes sets ServingRatTypes field to given value.

### HasServingRatTypes

`func (o *NrfInfoServedLmfInfoValue) HasServingRatTypes() bool`

HasServingRatTypes returns a boolean if a field has been set.

### GetTaiList

`func (o *NrfInfoServedLmfInfoValue) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NrfInfoServedLmfInfoValue) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NrfInfoServedLmfInfoValue) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NrfInfoServedLmfInfoValue) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NrfInfoServedLmfInfoValue) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NrfInfoServedLmfInfoValue) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NrfInfoServedLmfInfoValue) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NrfInfoServedLmfInfoValue) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetSupportedGADShapes

`func (o *NrfInfoServedLmfInfoValue) GetSupportedGADShapes() []SupportedGADShapes`

GetSupportedGADShapes returns the SupportedGADShapes field if non-nil, zero value otherwise.

### GetSupportedGADShapesOk

`func (o *NrfInfoServedLmfInfoValue) GetSupportedGADShapesOk() (*[]SupportedGADShapes, bool)`

GetSupportedGADShapesOk returns a tuple with the SupportedGADShapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedGADShapes

`func (o *NrfInfoServedLmfInfoValue) SetSupportedGADShapes(v []SupportedGADShapes)`

SetSupportedGADShapes sets SupportedGADShapes field to given value.

### HasSupportedGADShapes

`func (o *NrfInfoServedLmfInfoValue) HasSupportedGADShapes() bool`

HasSupportedGADShapes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


