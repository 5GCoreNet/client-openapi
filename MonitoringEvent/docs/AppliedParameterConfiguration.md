# AppliedParameterConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIds** | Pointer to **[]string** | Each element uniquely identifies a user. | [optional] 
**Msisdns** | Pointer to **[]string** | Each element identifies the MS internal PSTN/ISDN number allocated for a UE. | [optional] 
**MaximumLatency** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**MaximumResponseTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**MaximumDetectionTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 

## Methods

### NewAppliedParameterConfiguration

`func NewAppliedParameterConfiguration() *AppliedParameterConfiguration`

NewAppliedParameterConfiguration instantiates a new AppliedParameterConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedParameterConfigurationWithDefaults

`func NewAppliedParameterConfigurationWithDefaults() *AppliedParameterConfiguration`

NewAppliedParameterConfigurationWithDefaults instantiates a new AppliedParameterConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIds

`func (o *AppliedParameterConfiguration) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AppliedParameterConfiguration) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AppliedParameterConfiguration) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *AppliedParameterConfiguration) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetMsisdns

`func (o *AppliedParameterConfiguration) GetMsisdns() []string`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *AppliedParameterConfiguration) GetMsisdnsOk() (*[]string, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *AppliedParameterConfiguration) SetMsisdns(v []string)`

SetMsisdns sets Msisdns field to given value.

### HasMsisdns

`func (o *AppliedParameterConfiguration) HasMsisdns() bool`

HasMsisdns returns a boolean if a field has been set.

### GetMaximumLatency

`func (o *AppliedParameterConfiguration) GetMaximumLatency() int32`

GetMaximumLatency returns the MaximumLatency field if non-nil, zero value otherwise.

### GetMaximumLatencyOk

`func (o *AppliedParameterConfiguration) GetMaximumLatencyOk() (*int32, bool)`

GetMaximumLatencyOk returns a tuple with the MaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLatency

`func (o *AppliedParameterConfiguration) SetMaximumLatency(v int32)`

SetMaximumLatency sets MaximumLatency field to given value.

### HasMaximumLatency

`func (o *AppliedParameterConfiguration) HasMaximumLatency() bool`

HasMaximumLatency returns a boolean if a field has been set.

### GetMaximumResponseTime

`func (o *AppliedParameterConfiguration) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *AppliedParameterConfiguration) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *AppliedParameterConfiguration) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.

### HasMaximumResponseTime

`func (o *AppliedParameterConfiguration) HasMaximumResponseTime() bool`

HasMaximumResponseTime returns a boolean if a field has been set.

### GetMaximumDetectionTime

`func (o *AppliedParameterConfiguration) GetMaximumDetectionTime() int32`

GetMaximumDetectionTime returns the MaximumDetectionTime field if non-nil, zero value otherwise.

### GetMaximumDetectionTimeOk

`func (o *AppliedParameterConfiguration) GetMaximumDetectionTimeOk() (*int32, bool)`

GetMaximumDetectionTimeOk returns a tuple with the MaximumDetectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDetectionTime

`func (o *AppliedParameterConfiguration) SetMaximumDetectionTime(v int32)`

SetMaximumDetectionTime sets MaximumDetectionTime field to given value.

### HasMaximumDetectionTime

`func (o *AppliedParameterConfiguration) HasMaximumDetectionTime() bool`

HasMaximumDetectionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


