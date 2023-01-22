# RequestedUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**TotalVolume** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**UplinkVolume** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**DownlinkVolume** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**ServiceSpecificUnits** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 

## Methods

### NewRequestedUnit

`func NewRequestedUnit() *RequestedUnit`

NewRequestedUnit instantiates a new RequestedUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedUnitWithDefaults

`func NewRequestedUnitWithDefaults() *RequestedUnit`

NewRequestedUnitWithDefaults instantiates a new RequestedUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *RequestedUnit) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RequestedUnit) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RequestedUnit) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *RequestedUnit) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotalVolume

`func (o *RequestedUnit) GetTotalVolume() int32`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *RequestedUnit) GetTotalVolumeOk() (*int32, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *RequestedUnit) SetTotalVolume(v int32)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *RequestedUnit) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetUplinkVolume

`func (o *RequestedUnit) GetUplinkVolume() int32`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *RequestedUnit) GetUplinkVolumeOk() (*int32, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *RequestedUnit) SetUplinkVolume(v int32)`

SetUplinkVolume sets UplinkVolume field to given value.

### HasUplinkVolume

`func (o *RequestedUnit) HasUplinkVolume() bool`

HasUplinkVolume returns a boolean if a field has been set.

### GetDownlinkVolume

`func (o *RequestedUnit) GetDownlinkVolume() int32`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *RequestedUnit) GetDownlinkVolumeOk() (*int32, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *RequestedUnit) SetDownlinkVolume(v int32)`

SetDownlinkVolume sets DownlinkVolume field to given value.

### HasDownlinkVolume

`func (o *RequestedUnit) HasDownlinkVolume() bool`

HasDownlinkVolume returns a boolean if a field has been set.

### GetServiceSpecificUnits

`func (o *RequestedUnit) GetServiceSpecificUnits() int32`

GetServiceSpecificUnits returns the ServiceSpecificUnits field if non-nil, zero value otherwise.

### GetServiceSpecificUnitsOk

`func (o *RequestedUnit) GetServiceSpecificUnitsOk() (*int32, bool)`

GetServiceSpecificUnitsOk returns a tuple with the ServiceSpecificUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSpecificUnits

`func (o *RequestedUnit) SetServiceSpecificUnits(v int32)`

SetServiceSpecificUnits sets ServiceSpecificUnits field to given value.

### HasServiceSpecificUnits

`func (o *RequestedUnit) HasServiceSpecificUnits() bool`

HasServiceSpecificUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


