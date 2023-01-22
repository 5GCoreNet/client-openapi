# BsfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UiccType** | Pointer to [**UiccType**](UiccType.md) |  | [optional] 
**LifeTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SecurityFeatures** | Pointer to [**[]SecFeature**](SecFeature.md) |  | [optional] 

## Methods

### NewBsfInfo

`func NewBsfInfo() *BsfInfo`

NewBsfInfo instantiates a new BsfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBsfInfoWithDefaults

`func NewBsfInfoWithDefaults() *BsfInfo`

NewBsfInfoWithDefaults instantiates a new BsfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUiccType

`func (o *BsfInfo) GetUiccType() UiccType`

GetUiccType returns the UiccType field if non-nil, zero value otherwise.

### GetUiccTypeOk

`func (o *BsfInfo) GetUiccTypeOk() (*UiccType, bool)`

GetUiccTypeOk returns a tuple with the UiccType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiccType

`func (o *BsfInfo) SetUiccType(v UiccType)`

SetUiccType sets UiccType field to given value.

### HasUiccType

`func (o *BsfInfo) HasUiccType() bool`

HasUiccType returns a boolean if a field has been set.

### GetLifeTime

`func (o *BsfInfo) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *BsfInfo) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *BsfInfo) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.

### HasLifeTime

`func (o *BsfInfo) HasLifeTime() bool`

HasLifeTime returns a boolean if a field has been set.

### GetSecurityFeatures

`func (o *BsfInfo) GetSecurityFeatures() []SecFeature`

GetSecurityFeatures returns the SecurityFeatures field if non-nil, zero value otherwise.

### GetSecurityFeaturesOk

`func (o *BsfInfo) GetSecurityFeaturesOk() (*[]SecFeature, bool)`

GetSecurityFeaturesOk returns a tuple with the SecurityFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityFeatures

`func (o *BsfInfo) SetSecurityFeatures(v []SecFeature)`

SetSecurityFeatures sets SecurityFeatures field to given value.

### HasSecurityFeatures

`func (o *BsfInfo) HasSecurityFeatures() bool`

HasSecurityFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


