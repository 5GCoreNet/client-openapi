# FECConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FecScheme** | **string** | String providing an URI formatted according to RFC 3986. | 
**FecOverHead** | **int32** |  | 
**AdditionalParams** | Pointer to [**[]AddFecParams**](AddFecParams.md) |  | [optional] 

## Methods

### NewFECConfig

`func NewFECConfig(fecScheme string, fecOverHead int32, ) *FECConfig`

NewFECConfig instantiates a new FECConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFECConfigWithDefaults

`func NewFECConfigWithDefaults() *FECConfig`

NewFECConfigWithDefaults instantiates a new FECConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFecScheme

`func (o *FECConfig) GetFecScheme() string`

GetFecScheme returns the FecScheme field if non-nil, zero value otherwise.

### GetFecSchemeOk

`func (o *FECConfig) GetFecSchemeOk() (*string, bool)`

GetFecSchemeOk returns a tuple with the FecScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecScheme

`func (o *FECConfig) SetFecScheme(v string)`

SetFecScheme sets FecScheme field to given value.


### GetFecOverHead

`func (o *FECConfig) GetFecOverHead() int32`

GetFecOverHead returns the FecOverHead field if non-nil, zero value otherwise.

### GetFecOverHeadOk

`func (o *FECConfig) GetFecOverHeadOk() (*int32, bool)`

GetFecOverHeadOk returns a tuple with the FecOverHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecOverHead

`func (o *FECConfig) SetFecOverHead(v int32)`

SetFecOverHead sets FecOverHead field to given value.


### GetAdditionalParams

`func (o *FECConfig) GetAdditionalParams() []AddFecParams`

GetAdditionalParams returns the AdditionalParams field if non-nil, zero value otherwise.

### GetAdditionalParamsOk

`func (o *FECConfig) GetAdditionalParamsOk() (*[]AddFecParams, bool)`

GetAdditionalParamsOk returns a tuple with the AdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParams

`func (o *FECConfig) SetAdditionalParams(v []AddFecParams)`

SetAdditionalParams sets AdditionalParams field to given value.

### HasAdditionalParams

`func (o *FECConfig) HasAdditionalParams() bool`

HasAdditionalParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


