# EPRPAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLabel** | Pointer to **string** |  | [optional] 
**FarEndEntity** | Pointer to **string** |  | [optional] 
**SupportedPerfMetricGroups** | Pointer to [**[]SupportedPerfMetricGroup**](SupportedPerfMetricGroup.md) |  | [optional] 

## Methods

### NewEPRPAttr

`func NewEPRPAttr() *EPRPAttr`

NewEPRPAttr instantiates a new EPRPAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEPRPAttrWithDefaults

`func NewEPRPAttrWithDefaults() *EPRPAttr`

NewEPRPAttrWithDefaults instantiates a new EPRPAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLabel

`func (o *EPRPAttr) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *EPRPAttr) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *EPRPAttr) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *EPRPAttr) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetFarEndEntity

`func (o *EPRPAttr) GetFarEndEntity() string`

GetFarEndEntity returns the FarEndEntity field if non-nil, zero value otherwise.

### GetFarEndEntityOk

`func (o *EPRPAttr) GetFarEndEntityOk() (*string, bool)`

GetFarEndEntityOk returns a tuple with the FarEndEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarEndEntity

`func (o *EPRPAttr) SetFarEndEntity(v string)`

SetFarEndEntity sets FarEndEntity field to given value.

### HasFarEndEntity

`func (o *EPRPAttr) HasFarEndEntity() bool`

HasFarEndEntity returns a boolean if a field has been set.

### GetSupportedPerfMetricGroups

`func (o *EPRPAttr) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup`

GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field if non-nil, zero value otherwise.

### GetSupportedPerfMetricGroupsOk

`func (o *EPRPAttr) GetSupportedPerfMetricGroupsOk() (*[]SupportedPerfMetricGroup, bool)`

GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPerfMetricGroups

`func (o *EPRPAttr) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup)`

SetSupportedPerfMetricGroups sets SupportedPerfMetricGroups field to given value.

### HasSupportedPerfMetricGroups

`func (o *EPRPAttr) HasSupportedPerfMetricGroups() bool`

HasSupportedPerfMetricGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


