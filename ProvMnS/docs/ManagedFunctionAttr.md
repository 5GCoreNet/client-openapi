# ManagedFunctionAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLabel** | Pointer to **string** |  | [optional] 
**VnfParametersList** | Pointer to [**[]VnfParameter**](VnfParameter.md) |  | [optional] 
**PeeParametersList** | Pointer to [**[]PeeParameter**](PeeParameter.md) |  | [optional] 
**PriorityLabel** | Pointer to **int32** |  | [optional] 
**SupportedPerfMetricGroups** | Pointer to [**[]SupportedPerfMetricGroup**](SupportedPerfMetricGroup.md) |  | [optional] 
**SupportedTraceMetrics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewManagedFunctionAttr

`func NewManagedFunctionAttr() *ManagedFunctionAttr`

NewManagedFunctionAttr instantiates a new ManagedFunctionAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedFunctionAttrWithDefaults

`func NewManagedFunctionAttrWithDefaults() *ManagedFunctionAttr`

NewManagedFunctionAttrWithDefaults instantiates a new ManagedFunctionAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLabel

`func (o *ManagedFunctionAttr) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ManagedFunctionAttr) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ManagedFunctionAttr) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ManagedFunctionAttr) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetVnfParametersList

`func (o *ManagedFunctionAttr) GetVnfParametersList() []VnfParameter`

GetVnfParametersList returns the VnfParametersList field if non-nil, zero value otherwise.

### GetVnfParametersListOk

`func (o *ManagedFunctionAttr) GetVnfParametersListOk() (*[]VnfParameter, bool)`

GetVnfParametersListOk returns a tuple with the VnfParametersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnfParametersList

`func (o *ManagedFunctionAttr) SetVnfParametersList(v []VnfParameter)`

SetVnfParametersList sets VnfParametersList field to given value.

### HasVnfParametersList

`func (o *ManagedFunctionAttr) HasVnfParametersList() bool`

HasVnfParametersList returns a boolean if a field has been set.

### GetPeeParametersList

`func (o *ManagedFunctionAttr) GetPeeParametersList() []PeeParameter`

GetPeeParametersList returns the PeeParametersList field if non-nil, zero value otherwise.

### GetPeeParametersListOk

`func (o *ManagedFunctionAttr) GetPeeParametersListOk() (*[]PeeParameter, bool)`

GetPeeParametersListOk returns a tuple with the PeeParametersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeParametersList

`func (o *ManagedFunctionAttr) SetPeeParametersList(v []PeeParameter)`

SetPeeParametersList sets PeeParametersList field to given value.

### HasPeeParametersList

`func (o *ManagedFunctionAttr) HasPeeParametersList() bool`

HasPeeParametersList returns a boolean if a field has been set.

### GetPriorityLabel

`func (o *ManagedFunctionAttr) GetPriorityLabel() int32`

GetPriorityLabel returns the PriorityLabel field if non-nil, zero value otherwise.

### GetPriorityLabelOk

`func (o *ManagedFunctionAttr) GetPriorityLabelOk() (*int32, bool)`

GetPriorityLabelOk returns a tuple with the PriorityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLabel

`func (o *ManagedFunctionAttr) SetPriorityLabel(v int32)`

SetPriorityLabel sets PriorityLabel field to given value.

### HasPriorityLabel

`func (o *ManagedFunctionAttr) HasPriorityLabel() bool`

HasPriorityLabel returns a boolean if a field has been set.

### GetSupportedPerfMetricGroups

`func (o *ManagedFunctionAttr) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup`

GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field if non-nil, zero value otherwise.

### GetSupportedPerfMetricGroupsOk

`func (o *ManagedFunctionAttr) GetSupportedPerfMetricGroupsOk() (*[]SupportedPerfMetricGroup, bool)`

GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPerfMetricGroups

`func (o *ManagedFunctionAttr) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup)`

SetSupportedPerfMetricGroups sets SupportedPerfMetricGroups field to given value.

### HasSupportedPerfMetricGroups

`func (o *ManagedFunctionAttr) HasSupportedPerfMetricGroups() bool`

HasSupportedPerfMetricGroups returns a boolean if a field has been set.

### GetSupportedTraceMetrics

`func (o *ManagedFunctionAttr) GetSupportedTraceMetrics() []string`

GetSupportedTraceMetrics returns the SupportedTraceMetrics field if non-nil, zero value otherwise.

### GetSupportedTraceMetricsOk

`func (o *ManagedFunctionAttr) GetSupportedTraceMetricsOk() (*[]string, bool)`

GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTraceMetrics

`func (o *ManagedFunctionAttr) SetSupportedTraceMetrics(v []string)`

SetSupportedTraceMetrics sets SupportedTraceMetrics field to given value.

### HasSupportedTraceMetrics

`func (o *ManagedFunctionAttr) HasSupportedTraceMetrics() bool`

HasSupportedTraceMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


