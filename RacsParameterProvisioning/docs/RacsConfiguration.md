# RacsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RacsId** | **string** | The UE radio capability ID provided by the SCS/AS to identify the UE radio capability data. See 3GPP TS 23.003 for the encoding. | 
**RacsParamEps** | Pointer to **string** | The UE radio capability data in EPS. | [optional] 
**RacsParam5Gs** | Pointer to **string** | The UE radio capability data in 5GS. | [optional] 
**ImeiTacs** | **[]string** | Related UE model&#39;s IMEI-TAC values. | 

## Methods

### NewRacsConfiguration

`func NewRacsConfiguration(racsId string, imeiTacs []string, ) *RacsConfiguration`

NewRacsConfiguration instantiates a new RacsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacsConfigurationWithDefaults

`func NewRacsConfigurationWithDefaults() *RacsConfiguration`

NewRacsConfigurationWithDefaults instantiates a new RacsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRacsId

`func (o *RacsConfiguration) GetRacsId() string`

GetRacsId returns the RacsId field if non-nil, zero value otherwise.

### GetRacsIdOk

`func (o *RacsConfiguration) GetRacsIdOk() (*string, bool)`

GetRacsIdOk returns a tuple with the RacsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsId

`func (o *RacsConfiguration) SetRacsId(v string)`

SetRacsId sets RacsId field to given value.


### GetRacsParamEps

`func (o *RacsConfiguration) GetRacsParamEps() string`

GetRacsParamEps returns the RacsParamEps field if non-nil, zero value otherwise.

### GetRacsParamEpsOk

`func (o *RacsConfiguration) GetRacsParamEpsOk() (*string, bool)`

GetRacsParamEpsOk returns a tuple with the RacsParamEps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsParamEps

`func (o *RacsConfiguration) SetRacsParamEps(v string)`

SetRacsParamEps sets RacsParamEps field to given value.

### HasRacsParamEps

`func (o *RacsConfiguration) HasRacsParamEps() bool`

HasRacsParamEps returns a boolean if a field has been set.

### GetRacsParam5Gs

`func (o *RacsConfiguration) GetRacsParam5Gs() string`

GetRacsParam5Gs returns the RacsParam5Gs field if non-nil, zero value otherwise.

### GetRacsParam5GsOk

`func (o *RacsConfiguration) GetRacsParam5GsOk() (*string, bool)`

GetRacsParam5GsOk returns a tuple with the RacsParam5Gs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsParam5Gs

`func (o *RacsConfiguration) SetRacsParam5Gs(v string)`

SetRacsParam5Gs sets RacsParam5Gs field to given value.

### HasRacsParam5Gs

`func (o *RacsConfiguration) HasRacsParam5Gs() bool`

HasRacsParam5Gs returns a boolean if a field has been set.

### GetImeiTacs

`func (o *RacsConfiguration) GetImeiTacs() []string`

GetImeiTacs returns the ImeiTacs field if non-nil, zero value otherwise.

### GetImeiTacsOk

`func (o *RacsConfiguration) GetImeiTacsOk() (*[]string, bool)`

GetImeiTacsOk returns a tuple with the ImeiTacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImeiTacs

`func (o *RacsConfiguration) SetImeiTacs(v []string)`

SetImeiTacs sets ImeiTacs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


