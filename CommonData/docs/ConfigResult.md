# ConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIds** | Pointer to **[]string** | Each element indicates an external identifier of the UE. | [optional] 
**Msisdns** | Pointer to **[]string** | Each element identifies the MS internal PSTN/ISDN number allocated for the UE. | [optional] 
**ResultReason** | [**ResultReason**](ResultReason.md) |  | 

## Methods

### NewConfigResult

`func NewConfigResult(resultReason ResultReason, ) *ConfigResult`

NewConfigResult instantiates a new ConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigResultWithDefaults

`func NewConfigResultWithDefaults() *ConfigResult`

NewConfigResultWithDefaults instantiates a new ConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIds

`func (o *ConfigResult) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *ConfigResult) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *ConfigResult) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *ConfigResult) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetMsisdns

`func (o *ConfigResult) GetMsisdns() []string`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *ConfigResult) GetMsisdnsOk() (*[]string, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *ConfigResult) SetMsisdns(v []string)`

SetMsisdns sets Msisdns field to given value.

### HasMsisdns

`func (o *ConfigResult) HasMsisdns() bool`

HasMsisdns returns a boolean if a field has been set.

### GetResultReason

`func (o *ConfigResult) GetResultReason() ResultReason`

GetResultReason returns the ResultReason field if non-nil, zero value otherwise.

### GetResultReasonOk

`func (o *ConfigResult) GetResultReasonOk() (*ResultReason, bool)`

GetResultReasonOk returns a tuple with the ResultReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReason

`func (o *ConfigResult) SetResultReason(v ResultReason)`

SetResultReason sets ResultReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


