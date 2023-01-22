# RacsConfigurationRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RacsParamEps** | Pointer to **NullableString** | The UE radio capability data in EPS. | [optional] 
**RacsParam5Gs** | Pointer to **NullableString** | The UE radio capability data in 5GS. | [optional] 
**ImeiTacs** | Pointer to **[]string** | Related UE model&#39;s IMEI-TAC values. | [optional] 

## Methods

### NewRacsConfigurationRm

`func NewRacsConfigurationRm() *RacsConfigurationRm`

NewRacsConfigurationRm instantiates a new RacsConfigurationRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacsConfigurationRmWithDefaults

`func NewRacsConfigurationRmWithDefaults() *RacsConfigurationRm`

NewRacsConfigurationRmWithDefaults instantiates a new RacsConfigurationRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRacsParamEps

`func (o *RacsConfigurationRm) GetRacsParamEps() string`

GetRacsParamEps returns the RacsParamEps field if non-nil, zero value otherwise.

### GetRacsParamEpsOk

`func (o *RacsConfigurationRm) GetRacsParamEpsOk() (*string, bool)`

GetRacsParamEpsOk returns a tuple with the RacsParamEps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsParamEps

`func (o *RacsConfigurationRm) SetRacsParamEps(v string)`

SetRacsParamEps sets RacsParamEps field to given value.

### HasRacsParamEps

`func (o *RacsConfigurationRm) HasRacsParamEps() bool`

HasRacsParamEps returns a boolean if a field has been set.

### SetRacsParamEpsNil

`func (o *RacsConfigurationRm) SetRacsParamEpsNil(b bool)`

 SetRacsParamEpsNil sets the value for RacsParamEps to be an explicit nil

### UnsetRacsParamEps
`func (o *RacsConfigurationRm) UnsetRacsParamEps()`

UnsetRacsParamEps ensures that no value is present for RacsParamEps, not even an explicit nil
### GetRacsParam5Gs

`func (o *RacsConfigurationRm) GetRacsParam5Gs() string`

GetRacsParam5Gs returns the RacsParam5Gs field if non-nil, zero value otherwise.

### GetRacsParam5GsOk

`func (o *RacsConfigurationRm) GetRacsParam5GsOk() (*string, bool)`

GetRacsParam5GsOk returns a tuple with the RacsParam5Gs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsParam5Gs

`func (o *RacsConfigurationRm) SetRacsParam5Gs(v string)`

SetRacsParam5Gs sets RacsParam5Gs field to given value.

### HasRacsParam5Gs

`func (o *RacsConfigurationRm) HasRacsParam5Gs() bool`

HasRacsParam5Gs returns a boolean if a field has been set.

### SetRacsParam5GsNil

`func (o *RacsConfigurationRm) SetRacsParam5GsNil(b bool)`

 SetRacsParam5GsNil sets the value for RacsParam5Gs to be an explicit nil

### UnsetRacsParam5Gs
`func (o *RacsConfigurationRm) UnsetRacsParam5Gs()`

UnsetRacsParam5Gs ensures that no value is present for RacsParam5Gs, not even an explicit nil
### GetImeiTacs

`func (o *RacsConfigurationRm) GetImeiTacs() []string`

GetImeiTacs returns the ImeiTacs field if non-nil, zero value otherwise.

### GetImeiTacsOk

`func (o *RacsConfigurationRm) GetImeiTacsOk() (*[]string, bool)`

GetImeiTacsOk returns a tuple with the ImeiTacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImeiTacs

`func (o *RacsConfigurationRm) SetImeiTacs(v []string)`

SetImeiTacs sets ImeiTacs field to given value.

### HasImeiTacs

`func (o *RacsConfigurationRm) HasImeiTacs() bool`

HasImeiTacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


