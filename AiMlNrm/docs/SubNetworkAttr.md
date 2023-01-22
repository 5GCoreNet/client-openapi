# SubNetworkAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnPrefix** | Pointer to **string** |  | [optional] 
**UserLabel** | Pointer to **string** |  | [optional] 
**UserDefinedNetworkType** | Pointer to **string** |  | [optional] 
**SetOfMcc** | Pointer to **[]string** |  | [optional] 
**PriorityLabel** | Pointer to **int32** |  | [optional] 
**SupportedPerfMetricGroups** | Pointer to [**[]SupportedPerfMetricGroup**](SupportedPerfMetricGroup.md) |  | [optional] 
**SupportedTraceMetrics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSubNetworkAttr

`func NewSubNetworkAttr() *SubNetworkAttr`

NewSubNetworkAttr instantiates a new SubNetworkAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkAttrWithDefaults

`func NewSubNetworkAttrWithDefaults() *SubNetworkAttr`

NewSubNetworkAttrWithDefaults instantiates a new SubNetworkAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnPrefix

`func (o *SubNetworkAttr) GetDnPrefix() string`

GetDnPrefix returns the DnPrefix field if non-nil, zero value otherwise.

### GetDnPrefixOk

`func (o *SubNetworkAttr) GetDnPrefixOk() (*string, bool)`

GetDnPrefixOk returns a tuple with the DnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPrefix

`func (o *SubNetworkAttr) SetDnPrefix(v string)`

SetDnPrefix sets DnPrefix field to given value.

### HasDnPrefix

`func (o *SubNetworkAttr) HasDnPrefix() bool`

HasDnPrefix returns a boolean if a field has been set.

### GetUserLabel

`func (o *SubNetworkAttr) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *SubNetworkAttr) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *SubNetworkAttr) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *SubNetworkAttr) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetUserDefinedNetworkType

`func (o *SubNetworkAttr) GetUserDefinedNetworkType() string`

GetUserDefinedNetworkType returns the UserDefinedNetworkType field if non-nil, zero value otherwise.

### GetUserDefinedNetworkTypeOk

`func (o *SubNetworkAttr) GetUserDefinedNetworkTypeOk() (*string, bool)`

GetUserDefinedNetworkTypeOk returns a tuple with the UserDefinedNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedNetworkType

`func (o *SubNetworkAttr) SetUserDefinedNetworkType(v string)`

SetUserDefinedNetworkType sets UserDefinedNetworkType field to given value.

### HasUserDefinedNetworkType

`func (o *SubNetworkAttr) HasUserDefinedNetworkType() bool`

HasUserDefinedNetworkType returns a boolean if a field has been set.

### GetSetOfMcc

`func (o *SubNetworkAttr) GetSetOfMcc() []string`

GetSetOfMcc returns the SetOfMcc field if non-nil, zero value otherwise.

### GetSetOfMccOk

`func (o *SubNetworkAttr) GetSetOfMccOk() (*[]string, bool)`

GetSetOfMccOk returns a tuple with the SetOfMcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOfMcc

`func (o *SubNetworkAttr) SetSetOfMcc(v []string)`

SetSetOfMcc sets SetOfMcc field to given value.

### HasSetOfMcc

`func (o *SubNetworkAttr) HasSetOfMcc() bool`

HasSetOfMcc returns a boolean if a field has been set.

### GetPriorityLabel

`func (o *SubNetworkAttr) GetPriorityLabel() int32`

GetPriorityLabel returns the PriorityLabel field if non-nil, zero value otherwise.

### GetPriorityLabelOk

`func (o *SubNetworkAttr) GetPriorityLabelOk() (*int32, bool)`

GetPriorityLabelOk returns a tuple with the PriorityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLabel

`func (o *SubNetworkAttr) SetPriorityLabel(v int32)`

SetPriorityLabel sets PriorityLabel field to given value.

### HasPriorityLabel

`func (o *SubNetworkAttr) HasPriorityLabel() bool`

HasPriorityLabel returns a boolean if a field has been set.

### GetSupportedPerfMetricGroups

`func (o *SubNetworkAttr) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup`

GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field if non-nil, zero value otherwise.

### GetSupportedPerfMetricGroupsOk

`func (o *SubNetworkAttr) GetSupportedPerfMetricGroupsOk() (*[]SupportedPerfMetricGroup, bool)`

GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPerfMetricGroups

`func (o *SubNetworkAttr) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup)`

SetSupportedPerfMetricGroups sets SupportedPerfMetricGroups field to given value.

### HasSupportedPerfMetricGroups

`func (o *SubNetworkAttr) HasSupportedPerfMetricGroups() bool`

HasSupportedPerfMetricGroups returns a boolean if a field has been set.

### GetSupportedTraceMetrics

`func (o *SubNetworkAttr) GetSupportedTraceMetrics() []string`

GetSupportedTraceMetrics returns the SupportedTraceMetrics field if non-nil, zero value otherwise.

### GetSupportedTraceMetricsOk

`func (o *SubNetworkAttr) GetSupportedTraceMetricsOk() (*[]string, bool)`

GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTraceMetrics

`func (o *SubNetworkAttr) SetSupportedTraceMetrics(v []string)`

SetSupportedTraceMetrics sets SupportedTraceMetrics field to given value.

### HasSupportedTraceMetrics

`func (o *SubNetworkAttr) HasSupportedTraceMetrics() bool`

HasSupportedTraceMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


