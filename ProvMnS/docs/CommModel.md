# CommModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **int32** |  | [optional] 
**CommModelType** | Pointer to [**CommModelType**](CommModelType.md) |  | [optional] 
**TargetNFServiceList** | Pointer to **[]string** |  | [optional] 
**CommModelConfiguration** | Pointer to **string** |  | [optional] 

## Methods

### NewCommModel

`func NewCommModel() *CommModel`

NewCommModel instantiates a new CommModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommModelWithDefaults

`func NewCommModelWithDefaults() *CommModel`

NewCommModelWithDefaults instantiates a new CommModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CommModel) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CommModel) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CommModel) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CommModel) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetCommModelType

`func (o *CommModel) GetCommModelType() CommModelType`

GetCommModelType returns the CommModelType field if non-nil, zero value otherwise.

### GetCommModelTypeOk

`func (o *CommModel) GetCommModelTypeOk() (*CommModelType, bool)`

GetCommModelTypeOk returns a tuple with the CommModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommModelType

`func (o *CommModel) SetCommModelType(v CommModelType)`

SetCommModelType sets CommModelType field to given value.

### HasCommModelType

`func (o *CommModel) HasCommModelType() bool`

HasCommModelType returns a boolean if a field has been set.

### GetTargetNFServiceList

`func (o *CommModel) GetTargetNFServiceList() []string`

GetTargetNFServiceList returns the TargetNFServiceList field if non-nil, zero value otherwise.

### GetTargetNFServiceListOk

`func (o *CommModel) GetTargetNFServiceListOk() (*[]string, bool)`

GetTargetNFServiceListOk returns a tuple with the TargetNFServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNFServiceList

`func (o *CommModel) SetTargetNFServiceList(v []string)`

SetTargetNFServiceList sets TargetNFServiceList field to given value.

### HasTargetNFServiceList

`func (o *CommModel) HasTargetNFServiceList() bool`

HasTargetNFServiceList returns a boolean if a field has been set.

### GetCommModelConfiguration

`func (o *CommModel) GetCommModelConfiguration() string`

GetCommModelConfiguration returns the CommModelConfiguration field if non-nil, zero value otherwise.

### GetCommModelConfigurationOk

`func (o *CommModel) GetCommModelConfigurationOk() (*string, bool)`

GetCommModelConfigurationOk returns a tuple with the CommModelConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommModelConfiguration

`func (o *CommModel) SetCommModelConfiguration(v string)`

SetCommModelConfiguration sets CommModelConfiguration field to given value.

### HasCommModelConfiguration

`func (o *CommModel) HasCommModelConfiguration() bool`

HasCommModelConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


