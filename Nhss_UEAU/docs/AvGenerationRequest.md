# AvGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imsi** | **string** |  | 
**AuthType** | [**AuthType**](AuthType.md) |  | 
**ServingNetworkName** | **string** |  | 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 

## Methods

### NewAvGenerationRequest

`func NewAvGenerationRequest(imsi string, authType AuthType, servingNetworkName string, ) *AvGenerationRequest`

NewAvGenerationRequest instantiates a new AvGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvGenerationRequestWithDefaults

`func NewAvGenerationRequestWithDefaults() *AvGenerationRequest`

NewAvGenerationRequestWithDefaults instantiates a new AvGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImsi

`func (o *AvGenerationRequest) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *AvGenerationRequest) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *AvGenerationRequest) SetImsi(v string)`

SetImsi sets Imsi field to given value.


### GetAuthType

`func (o *AvGenerationRequest) GetAuthType() AuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AvGenerationRequest) GetAuthTypeOk() (*AuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AvGenerationRequest) SetAuthType(v AuthType)`

SetAuthType sets AuthType field to given value.


### GetServingNetworkName

`func (o *AvGenerationRequest) GetServingNetworkName() string`

GetServingNetworkName returns the ServingNetworkName field if non-nil, zero value otherwise.

### GetServingNetworkNameOk

`func (o *AvGenerationRequest) GetServingNetworkNameOk() (*string, bool)`

GetServingNetworkNameOk returns a tuple with the ServingNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkName

`func (o *AvGenerationRequest) SetServingNetworkName(v string)`

SetServingNetworkName sets ServingNetworkName field to given value.


### GetResynchronizationInfo

`func (o *AvGenerationRequest) GetResynchronizationInfo() ResynchronizationInfo`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *AvGenerationRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *AvGenerationRequest) SetResynchronizationInfo(v ResynchronizationInfo)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *AvGenerationRequest) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


