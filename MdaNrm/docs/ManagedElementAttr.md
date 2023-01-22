# ManagedElementAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnPrefix** | Pointer to **string** |  | [optional] 
**ManagedElementTypeList** | Pointer to **[]string** |  | [optional] 
**UserLabel** | Pointer to **string** |  | [optional] 
**LocationName** | Pointer to **string** |  | [optional] 
**ManagedBy** | Pointer to **[]string** |  | [optional] 
**VendorName** | Pointer to **string** |  | [optional] 
**UserDefinedState** | Pointer to **string** |  | [optional] 
**SwVersion** | Pointer to **string** |  | [optional] 
**PriorityLabel** | Pointer to **int32** |  | [optional] 
**SupportedPerfMetricGroups** | Pointer to [**[]SupportedPerfMetricGroup**](SupportedPerfMetricGroup.md) |  | [optional] 
**SupportedTraceMetrics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewManagedElementAttr

`func NewManagedElementAttr() *ManagedElementAttr`

NewManagedElementAttr instantiates a new ManagedElementAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedElementAttrWithDefaults

`func NewManagedElementAttrWithDefaults() *ManagedElementAttr`

NewManagedElementAttrWithDefaults instantiates a new ManagedElementAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnPrefix

`func (o *ManagedElementAttr) GetDnPrefix() string`

GetDnPrefix returns the DnPrefix field if non-nil, zero value otherwise.

### GetDnPrefixOk

`func (o *ManagedElementAttr) GetDnPrefixOk() (*string, bool)`

GetDnPrefixOk returns a tuple with the DnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPrefix

`func (o *ManagedElementAttr) SetDnPrefix(v string)`

SetDnPrefix sets DnPrefix field to given value.

### HasDnPrefix

`func (o *ManagedElementAttr) HasDnPrefix() bool`

HasDnPrefix returns a boolean if a field has been set.

### GetManagedElementTypeList

`func (o *ManagedElementAttr) GetManagedElementTypeList() []string`

GetManagedElementTypeList returns the ManagedElementTypeList field if non-nil, zero value otherwise.

### GetManagedElementTypeListOk

`func (o *ManagedElementAttr) GetManagedElementTypeListOk() (*[]string, bool)`

GetManagedElementTypeListOk returns a tuple with the ManagedElementTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedElementTypeList

`func (o *ManagedElementAttr) SetManagedElementTypeList(v []string)`

SetManagedElementTypeList sets ManagedElementTypeList field to given value.

### HasManagedElementTypeList

`func (o *ManagedElementAttr) HasManagedElementTypeList() bool`

HasManagedElementTypeList returns a boolean if a field has been set.

### GetUserLabel

`func (o *ManagedElementAttr) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ManagedElementAttr) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ManagedElementAttr) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ManagedElementAttr) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetLocationName

`func (o *ManagedElementAttr) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *ManagedElementAttr) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *ManagedElementAttr) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *ManagedElementAttr) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetManagedBy

`func (o *ManagedElementAttr) GetManagedBy() []string`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *ManagedElementAttr) GetManagedByOk() (*[]string, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *ManagedElementAttr) SetManagedBy(v []string)`

SetManagedBy sets ManagedBy field to given value.

### HasManagedBy

`func (o *ManagedElementAttr) HasManagedBy() bool`

HasManagedBy returns a boolean if a field has been set.

### GetVendorName

`func (o *ManagedElementAttr) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *ManagedElementAttr) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *ManagedElementAttr) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *ManagedElementAttr) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetUserDefinedState

`func (o *ManagedElementAttr) GetUserDefinedState() string`

GetUserDefinedState returns the UserDefinedState field if non-nil, zero value otherwise.

### GetUserDefinedStateOk

`func (o *ManagedElementAttr) GetUserDefinedStateOk() (*string, bool)`

GetUserDefinedStateOk returns a tuple with the UserDefinedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedState

`func (o *ManagedElementAttr) SetUserDefinedState(v string)`

SetUserDefinedState sets UserDefinedState field to given value.

### HasUserDefinedState

`func (o *ManagedElementAttr) HasUserDefinedState() bool`

HasUserDefinedState returns a boolean if a field has been set.

### GetSwVersion

`func (o *ManagedElementAttr) GetSwVersion() string`

GetSwVersion returns the SwVersion field if non-nil, zero value otherwise.

### GetSwVersionOk

`func (o *ManagedElementAttr) GetSwVersionOk() (*string, bool)`

GetSwVersionOk returns a tuple with the SwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwVersion

`func (o *ManagedElementAttr) SetSwVersion(v string)`

SetSwVersion sets SwVersion field to given value.

### HasSwVersion

`func (o *ManagedElementAttr) HasSwVersion() bool`

HasSwVersion returns a boolean if a field has been set.

### GetPriorityLabel

`func (o *ManagedElementAttr) GetPriorityLabel() int32`

GetPriorityLabel returns the PriorityLabel field if non-nil, zero value otherwise.

### GetPriorityLabelOk

`func (o *ManagedElementAttr) GetPriorityLabelOk() (*int32, bool)`

GetPriorityLabelOk returns a tuple with the PriorityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLabel

`func (o *ManagedElementAttr) SetPriorityLabel(v int32)`

SetPriorityLabel sets PriorityLabel field to given value.

### HasPriorityLabel

`func (o *ManagedElementAttr) HasPriorityLabel() bool`

HasPriorityLabel returns a boolean if a field has been set.

### GetSupportedPerfMetricGroups

`func (o *ManagedElementAttr) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup`

GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field if non-nil, zero value otherwise.

### GetSupportedPerfMetricGroupsOk

`func (o *ManagedElementAttr) GetSupportedPerfMetricGroupsOk() (*[]SupportedPerfMetricGroup, bool)`

GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPerfMetricGroups

`func (o *ManagedElementAttr) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup)`

SetSupportedPerfMetricGroups sets SupportedPerfMetricGroups field to given value.

### HasSupportedPerfMetricGroups

`func (o *ManagedElementAttr) HasSupportedPerfMetricGroups() bool`

HasSupportedPerfMetricGroups returns a boolean if a field has been set.

### GetSupportedTraceMetrics

`func (o *ManagedElementAttr) GetSupportedTraceMetrics() []string`

GetSupportedTraceMetrics returns the SupportedTraceMetrics field if non-nil, zero value otherwise.

### GetSupportedTraceMetricsOk

`func (o *ManagedElementAttr) GetSupportedTraceMetricsOk() (*[]string, bool)`

GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTraceMetrics

`func (o *ManagedElementAttr) SetSupportedTraceMetrics(v []string)`

SetSupportedTraceMetrics sets SupportedTraceMetrics field to given value.

### HasSupportedTraceMetrics

`func (o *ManagedElementAttr) HasSupportedTraceMetrics() bool`

HasSupportedTraceMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


