# LmfInfo

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

### NewLmfInfo

`func NewLmfInfo() *LmfInfo`

NewLmfInfo instantiates a new LmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLmfInfoWithDefaults

`func NewLmfInfoWithDefaults() *LmfInfo`

NewLmfInfoWithDefaults instantiates a new LmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServingClientTypes

`func (o *LmfInfo) GetServingClientTypes() []ExternalClientType`

GetServingClientTypes returns the ServingClientTypes field if non-nil, zero value otherwise.

### GetServingClientTypesOk

`func (o *LmfInfo) GetServingClientTypesOk() (*[]ExternalClientType, bool)`

GetServingClientTypesOk returns a tuple with the ServingClientTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingClientTypes

`func (o *LmfInfo) SetServingClientTypes(v []ExternalClientType)`

SetServingClientTypes sets ServingClientTypes field to given value.

### HasServingClientTypes

`func (o *LmfInfo) HasServingClientTypes() bool`

HasServingClientTypes returns a boolean if a field has been set.

### GetLmfId

`func (o *LmfInfo) GetLmfId() string`

GetLmfId returns the LmfId field if non-nil, zero value otherwise.

### GetLmfIdOk

`func (o *LmfInfo) GetLmfIdOk() (*string, bool)`

GetLmfIdOk returns a tuple with the LmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfId

`func (o *LmfInfo) SetLmfId(v string)`

SetLmfId sets LmfId field to given value.

### HasLmfId

`func (o *LmfInfo) HasLmfId() bool`

HasLmfId returns a boolean if a field has been set.

### GetServingAccessTypes

`func (o *LmfInfo) GetServingAccessTypes() []AccessType`

GetServingAccessTypes returns the ServingAccessTypes field if non-nil, zero value otherwise.

### GetServingAccessTypesOk

`func (o *LmfInfo) GetServingAccessTypesOk() (*[]AccessType, bool)`

GetServingAccessTypesOk returns a tuple with the ServingAccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingAccessTypes

`func (o *LmfInfo) SetServingAccessTypes(v []AccessType)`

SetServingAccessTypes sets ServingAccessTypes field to given value.

### HasServingAccessTypes

`func (o *LmfInfo) HasServingAccessTypes() bool`

HasServingAccessTypes returns a boolean if a field has been set.

### GetServingAnNodeTypes

`func (o *LmfInfo) GetServingAnNodeTypes() []AnNodeType`

GetServingAnNodeTypes returns the ServingAnNodeTypes field if non-nil, zero value otherwise.

### GetServingAnNodeTypesOk

`func (o *LmfInfo) GetServingAnNodeTypesOk() (*[]AnNodeType, bool)`

GetServingAnNodeTypesOk returns a tuple with the ServingAnNodeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingAnNodeTypes

`func (o *LmfInfo) SetServingAnNodeTypes(v []AnNodeType)`

SetServingAnNodeTypes sets ServingAnNodeTypes field to given value.

### HasServingAnNodeTypes

`func (o *LmfInfo) HasServingAnNodeTypes() bool`

HasServingAnNodeTypes returns a boolean if a field has been set.

### GetServingRatTypes

`func (o *LmfInfo) GetServingRatTypes() []RatType`

GetServingRatTypes returns the ServingRatTypes field if non-nil, zero value otherwise.

### GetServingRatTypesOk

`func (o *LmfInfo) GetServingRatTypesOk() (*[]RatType, bool)`

GetServingRatTypesOk returns a tuple with the ServingRatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingRatTypes

`func (o *LmfInfo) SetServingRatTypes(v []RatType)`

SetServingRatTypes sets ServingRatTypes field to given value.

### HasServingRatTypes

`func (o *LmfInfo) HasServingRatTypes() bool`

HasServingRatTypes returns a boolean if a field has been set.

### GetTaiList

`func (o *LmfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *LmfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *LmfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *LmfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *LmfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *LmfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *LmfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *LmfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetSupportedGADShapes

`func (o *LmfInfo) GetSupportedGADShapes() []SupportedGADShapes`

GetSupportedGADShapes returns the SupportedGADShapes field if non-nil, zero value otherwise.

### GetSupportedGADShapesOk

`func (o *LmfInfo) GetSupportedGADShapesOk() (*[]SupportedGADShapes, bool)`

GetSupportedGADShapesOk returns a tuple with the SupportedGADShapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedGADShapes

`func (o *LmfInfo) SetSupportedGADShapes(v []SupportedGADShapes)`

SetSupportedGADShapes sets SupportedGADShapes field to given value.

### HasSupportedGADShapes

`func (o *LmfInfo) HasSupportedGADShapes() bool`

HasSupportedGADShapes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


