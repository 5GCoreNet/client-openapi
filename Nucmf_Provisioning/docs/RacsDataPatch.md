# RacsDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RacsConfigs** | Pointer to [**map[string]RacsConfigurationRm**](RacsConfigurationRm.md) | Identifies the configuration related to manufacturer specific UE radio capability. Each element uniquely identifies an RACS configuration for an RACS ID and is identified in the map via the RACS ID as key.  | [optional] 

## Methods

### NewRacsDataPatch

`func NewRacsDataPatch() *RacsDataPatch`

NewRacsDataPatch instantiates a new RacsDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacsDataPatchWithDefaults

`func NewRacsDataPatchWithDefaults() *RacsDataPatch`

NewRacsDataPatchWithDefaults instantiates a new RacsDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRacsConfigs

`func (o *RacsDataPatch) GetRacsConfigs() map[string]RacsConfigurationRm`

GetRacsConfigs returns the RacsConfigs field if non-nil, zero value otherwise.

### GetRacsConfigsOk

`func (o *RacsDataPatch) GetRacsConfigsOk() (*map[string]RacsConfigurationRm, bool)`

GetRacsConfigsOk returns a tuple with the RacsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsConfigs

`func (o *RacsDataPatch) SetRacsConfigs(v map[string]RacsConfigurationRm)`

SetRacsConfigs sets RacsConfigs field to given value.

### HasRacsConfigs

`func (o *RacsDataPatch) HasRacsConfigs() bool`

HasRacsConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


