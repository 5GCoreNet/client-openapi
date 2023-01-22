# EESProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EesId** | **string** | Identifier of the EES. | 
**EndPt** | [**EndPoint**](EndPoint.md) |  | 
**EasIds** | Pointer to **[]string** | Application identifiers of EASs that are registered with EES. | [optional] 
**ProvId** | Pointer to **string** | Identifier of the ECSP that provides the EES provider. | [optional] 
**SvcArea** | Pointer to [**ServiceArea**](ServiceArea.md) |  | [optional] 
**AppLocs** | Pointer to **[]string** | List of DNAI(s) associated with the EES. | [optional] 
**SvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | The ACR scenarios supported by the EES for service continuity. | [optional] 
**EecRegConf** | **bool** | Set to true if the EEC is required to register to the EES to use edge service. Set to false if the EEC is not required to register to use edge services.  | 

## Methods

### NewEESProfile

`func NewEESProfile(eesId string, endPt EndPoint, eecRegConf bool, ) *EESProfile`

NewEESProfile instantiates a new EESProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEESProfileWithDefaults

`func NewEESProfileWithDefaults() *EESProfile`

NewEESProfileWithDefaults instantiates a new EESProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEesId

`func (o *EESProfile) GetEesId() string`

GetEesId returns the EesId field if non-nil, zero value otherwise.

### GetEesIdOk

`func (o *EESProfile) GetEesIdOk() (*string, bool)`

GetEesIdOk returns a tuple with the EesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEesId

`func (o *EESProfile) SetEesId(v string)`

SetEesId sets EesId field to given value.


### GetEndPt

`func (o *EESProfile) GetEndPt() EndPoint`

GetEndPt returns the EndPt field if non-nil, zero value otherwise.

### GetEndPtOk

`func (o *EESProfile) GetEndPtOk() (*EndPoint, bool)`

GetEndPtOk returns a tuple with the EndPt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPt

`func (o *EESProfile) SetEndPt(v EndPoint)`

SetEndPt sets EndPt field to given value.


### GetEasIds

`func (o *EESProfile) GetEasIds() []string`

GetEasIds returns the EasIds field if non-nil, zero value otherwise.

### GetEasIdsOk

`func (o *EESProfile) GetEasIdsOk() (*[]string, bool)`

GetEasIdsOk returns a tuple with the EasIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIds

`func (o *EESProfile) SetEasIds(v []string)`

SetEasIds sets EasIds field to given value.

### HasEasIds

`func (o *EESProfile) HasEasIds() bool`

HasEasIds returns a boolean if a field has been set.

### GetProvId

`func (o *EESProfile) GetProvId() string`

GetProvId returns the ProvId field if non-nil, zero value otherwise.

### GetProvIdOk

`func (o *EESProfile) GetProvIdOk() (*string, bool)`

GetProvIdOk returns a tuple with the ProvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvId

`func (o *EESProfile) SetProvId(v string)`

SetProvId sets ProvId field to given value.

### HasProvId

`func (o *EESProfile) HasProvId() bool`

HasProvId returns a boolean if a field has been set.

### GetSvcArea

`func (o *EESProfile) GetSvcArea() ServiceArea`

GetSvcArea returns the SvcArea field if non-nil, zero value otherwise.

### GetSvcAreaOk

`func (o *EESProfile) GetSvcAreaOk() (*ServiceArea, bool)`

GetSvcAreaOk returns a tuple with the SvcArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcArea

`func (o *EESProfile) SetSvcArea(v ServiceArea)`

SetSvcArea sets SvcArea field to given value.

### HasSvcArea

`func (o *EESProfile) HasSvcArea() bool`

HasSvcArea returns a boolean if a field has been set.

### GetAppLocs

`func (o *EESProfile) GetAppLocs() []string`

GetAppLocs returns the AppLocs field if non-nil, zero value otherwise.

### GetAppLocsOk

`func (o *EESProfile) GetAppLocsOk() (*[]string, bool)`

GetAppLocsOk returns a tuple with the AppLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLocs

`func (o *EESProfile) SetAppLocs(v []string)`

SetAppLocs sets AppLocs field to given value.

### HasAppLocs

`func (o *EESProfile) HasAppLocs() bool`

HasAppLocs returns a boolean if a field has been set.

### GetSvcContSupp

`func (o *EESProfile) GetSvcContSupp() []ACRScenario`

GetSvcContSupp returns the SvcContSupp field if non-nil, zero value otherwise.

### GetSvcContSuppOk

`func (o *EESProfile) GetSvcContSuppOk() (*[]ACRScenario, bool)`

GetSvcContSuppOk returns a tuple with the SvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcContSupp

`func (o *EESProfile) SetSvcContSupp(v []ACRScenario)`

SetSvcContSupp sets SvcContSupp field to given value.

### HasSvcContSupp

`func (o *EESProfile) HasSvcContSupp() bool`

HasSvcContSupp returns a boolean if a field has been set.

### GetEecRegConf

`func (o *EESProfile) GetEecRegConf() bool`

GetEecRegConf returns the EecRegConf field if non-nil, zero value otherwise.

### GetEecRegConfOk

`func (o *EESProfile) GetEecRegConfOk() (*bool, bool)`

GetEecRegConfOk returns a tuple with the EecRegConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecRegConf

`func (o *EESProfile) SetEecRegConf(v bool)`

SetEecRegConf sets EecRegConf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


