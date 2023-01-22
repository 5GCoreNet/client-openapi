# EasCharacteristics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasId** | Pointer to **string** | EAS application identifier. | [optional] 
**EasProvId** | Pointer to **string** | EAS provider identifier. | [optional] 
**EasType** | Pointer to **string** | EAS type. | [optional] 
**EasSched** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**SvcArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**EasSvcContinuity** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC. | [optional] 
**SvcPermLevel** | Pointer to **string** | Service permissions level. | [optional] 
**SvcFeats** | Pointer to **[]string** | Service features. | [optional] 

## Methods

### NewEasCharacteristics

`func NewEasCharacteristics() *EasCharacteristics`

NewEasCharacteristics instantiates a new EasCharacteristics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasCharacteristicsWithDefaults

`func NewEasCharacteristicsWithDefaults() *EasCharacteristics`

NewEasCharacteristicsWithDefaults instantiates a new EasCharacteristics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasId

`func (o *EasCharacteristics) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *EasCharacteristics) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *EasCharacteristics) SetEasId(v string)`

SetEasId sets EasId field to given value.

### HasEasId

`func (o *EasCharacteristics) HasEasId() bool`

HasEasId returns a boolean if a field has been set.

### GetEasProvId

`func (o *EasCharacteristics) GetEasProvId() string`

GetEasProvId returns the EasProvId field if non-nil, zero value otherwise.

### GetEasProvIdOk

`func (o *EasCharacteristics) GetEasProvIdOk() (*string, bool)`

GetEasProvIdOk returns a tuple with the EasProvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasProvId

`func (o *EasCharacteristics) SetEasProvId(v string)`

SetEasProvId sets EasProvId field to given value.

### HasEasProvId

`func (o *EasCharacteristics) HasEasProvId() bool`

HasEasProvId returns a boolean if a field has been set.

### GetEasType

`func (o *EasCharacteristics) GetEasType() string`

GetEasType returns the EasType field if non-nil, zero value otherwise.

### GetEasTypeOk

`func (o *EasCharacteristics) GetEasTypeOk() (*string, bool)`

GetEasTypeOk returns a tuple with the EasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasType

`func (o *EasCharacteristics) SetEasType(v string)`

SetEasType sets EasType field to given value.

### HasEasType

`func (o *EasCharacteristics) HasEasType() bool`

HasEasType returns a boolean if a field has been set.

### GetEasSched

`func (o *EasCharacteristics) GetEasSched() TimeWindow`

GetEasSched returns the EasSched field if non-nil, zero value otherwise.

### GetEasSchedOk

`func (o *EasCharacteristics) GetEasSchedOk() (*TimeWindow, bool)`

GetEasSchedOk returns a tuple with the EasSched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasSched

`func (o *EasCharacteristics) SetEasSched(v TimeWindow)`

SetEasSched sets EasSched field to given value.

### HasEasSched

`func (o *EasCharacteristics) HasEasSched() bool`

HasEasSched returns a boolean if a field has been set.

### GetSvcArea

`func (o *EasCharacteristics) GetSvcArea() LocationArea5G`

GetSvcArea returns the SvcArea field if non-nil, zero value otherwise.

### GetSvcAreaOk

`func (o *EasCharacteristics) GetSvcAreaOk() (*LocationArea5G, bool)`

GetSvcAreaOk returns a tuple with the SvcArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcArea

`func (o *EasCharacteristics) SetSvcArea(v LocationArea5G)`

SetSvcArea sets SvcArea field to given value.

### HasSvcArea

`func (o *EasCharacteristics) HasSvcArea() bool`

HasSvcArea returns a boolean if a field has been set.

### GetEasSvcContinuity

`func (o *EasCharacteristics) GetEasSvcContinuity() []ACRScenario`

GetEasSvcContinuity returns the EasSvcContinuity field if non-nil, zero value otherwise.

### GetEasSvcContinuityOk

`func (o *EasCharacteristics) GetEasSvcContinuityOk() (*[]ACRScenario, bool)`

GetEasSvcContinuityOk returns a tuple with the EasSvcContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasSvcContinuity

`func (o *EasCharacteristics) SetEasSvcContinuity(v []ACRScenario)`

SetEasSvcContinuity sets EasSvcContinuity field to given value.

### HasEasSvcContinuity

`func (o *EasCharacteristics) HasEasSvcContinuity() bool`

HasEasSvcContinuity returns a boolean if a field has been set.

### GetSvcPermLevel

`func (o *EasCharacteristics) GetSvcPermLevel() string`

GetSvcPermLevel returns the SvcPermLevel field if non-nil, zero value otherwise.

### GetSvcPermLevelOk

`func (o *EasCharacteristics) GetSvcPermLevelOk() (*string, bool)`

GetSvcPermLevelOk returns a tuple with the SvcPermLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcPermLevel

`func (o *EasCharacteristics) SetSvcPermLevel(v string)`

SetSvcPermLevel sets SvcPermLevel field to given value.

### HasSvcPermLevel

`func (o *EasCharacteristics) HasSvcPermLevel() bool`

HasSvcPermLevel returns a boolean if a field has been set.

### GetSvcFeats

`func (o *EasCharacteristics) GetSvcFeats() []string`

GetSvcFeats returns the SvcFeats field if non-nil, zero value otherwise.

### GetSvcFeatsOk

`func (o *EasCharacteristics) GetSvcFeatsOk() (*[]string, bool)`

GetSvcFeatsOk returns a tuple with the SvcFeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcFeats

`func (o *EasCharacteristics) SetSvcFeats(v []string)`

SetSvcFeats sets SvcFeats field to given value.

### HasSvcFeats

`func (o *EasCharacteristics) HasSvcFeats() bool`

HasSvcFeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


