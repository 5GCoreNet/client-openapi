# TimeSyncCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpNodeId** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | 
**GmCapables** | Pointer to [**[]GmCapable**](GmCapable.md) |  | [optional] 
**AsTimeRes** | Pointer to [**AsTimeResource**](AsTimeResource.md) |  | [optional] 
**PtpCapForUes** | Pointer to [**map[string]PtpCapabilitiesPerUe**](PtpCapabilitiesPerUe.md) | Contains the PTP capabilities supported by each of the UE(s). The key of the map is the gpsi.  | [optional] 

## Methods

### NewTimeSyncCapability

`func NewTimeSyncCapability(upNodeId int32, ) *TimeSyncCapability`

NewTimeSyncCapability instantiates a new TimeSyncCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncCapabilityWithDefaults

`func NewTimeSyncCapabilityWithDefaults() *TimeSyncCapability`

NewTimeSyncCapabilityWithDefaults instantiates a new TimeSyncCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpNodeId

`func (o *TimeSyncCapability) GetUpNodeId() int32`

GetUpNodeId returns the UpNodeId field if non-nil, zero value otherwise.

### GetUpNodeIdOk

`func (o *TimeSyncCapability) GetUpNodeIdOk() (*int32, bool)`

GetUpNodeIdOk returns a tuple with the UpNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpNodeId

`func (o *TimeSyncCapability) SetUpNodeId(v int32)`

SetUpNodeId sets UpNodeId field to given value.


### GetGmCapables

`func (o *TimeSyncCapability) GetGmCapables() []GmCapable`

GetGmCapables returns the GmCapables field if non-nil, zero value otherwise.

### GetGmCapablesOk

`func (o *TimeSyncCapability) GetGmCapablesOk() (*[]GmCapable, bool)`

GetGmCapablesOk returns a tuple with the GmCapables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmCapables

`func (o *TimeSyncCapability) SetGmCapables(v []GmCapable)`

SetGmCapables sets GmCapables field to given value.

### HasGmCapables

`func (o *TimeSyncCapability) HasGmCapables() bool`

HasGmCapables returns a boolean if a field has been set.

### GetAsTimeRes

`func (o *TimeSyncCapability) GetAsTimeRes() AsTimeResource`

GetAsTimeRes returns the AsTimeRes field if non-nil, zero value otherwise.

### GetAsTimeResOk

`func (o *TimeSyncCapability) GetAsTimeResOk() (*AsTimeResource, bool)`

GetAsTimeResOk returns a tuple with the AsTimeRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTimeRes

`func (o *TimeSyncCapability) SetAsTimeRes(v AsTimeResource)`

SetAsTimeRes sets AsTimeRes field to given value.

### HasAsTimeRes

`func (o *TimeSyncCapability) HasAsTimeRes() bool`

HasAsTimeRes returns a boolean if a field has been set.

### GetPtpCapForUes

`func (o *TimeSyncCapability) GetPtpCapForUes() map[string]PtpCapabilitiesPerUe`

GetPtpCapForUes returns the PtpCapForUes field if non-nil, zero value otherwise.

### GetPtpCapForUesOk

`func (o *TimeSyncCapability) GetPtpCapForUesOk() (*map[string]PtpCapabilitiesPerUe, bool)`

GetPtpCapForUesOk returns a tuple with the PtpCapForUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCapForUes

`func (o *TimeSyncCapability) SetPtpCapForUes(v map[string]PtpCapabilitiesPerUe)`

SetPtpCapForUes sets PtpCapForUes field to given value.

### HasPtpCapForUes

`func (o *TimeSyncCapability) HasPtpCapForUes() bool`

HasPtpCapForUes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


