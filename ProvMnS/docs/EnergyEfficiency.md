# EnergyEfficiency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServAttrCom** | Pointer to [**ServAttrCom**](ServAttrCom.md) |  | [optional] 
**Performance** | Pointer to [**EEPerfReq**](EEPerfReq.md) |  | [optional] 

## Methods

### NewEnergyEfficiency

`func NewEnergyEfficiency() *EnergyEfficiency`

NewEnergyEfficiency instantiates a new EnergyEfficiency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyEfficiencyWithDefaults

`func NewEnergyEfficiencyWithDefaults() *EnergyEfficiency`

NewEnergyEfficiencyWithDefaults instantiates a new EnergyEfficiency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServAttrCom

`func (o *EnergyEfficiency) GetServAttrCom() ServAttrCom`

GetServAttrCom returns the ServAttrCom field if non-nil, zero value otherwise.

### GetServAttrComOk

`func (o *EnergyEfficiency) GetServAttrComOk() (*ServAttrCom, bool)`

GetServAttrComOk returns a tuple with the ServAttrCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAttrCom

`func (o *EnergyEfficiency) SetServAttrCom(v ServAttrCom)`

SetServAttrCom sets ServAttrCom field to given value.

### HasServAttrCom

`func (o *EnergyEfficiency) HasServAttrCom() bool`

HasServAttrCom returns a boolean if a field has been set.

### GetPerformance

`func (o *EnergyEfficiency) GetPerformance() EEPerfReq`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *EnergyEfficiency) GetPerformanceOk() (*EEPerfReq, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *EnergyEfficiency) SetPerformance(v EEPerfReq)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *EnergyEfficiency) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


