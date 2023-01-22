# QosMonitoringMeasurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DlPacketDelay** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**UlPacketDelay** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**RtrPacketDelay** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**MeasureFailure** | Pointer to **bool** |  | [optional] 

## Methods

### NewQosMonitoringMeasurement

`func NewQosMonitoringMeasurement() *QosMonitoringMeasurement`

NewQosMonitoringMeasurement instantiates a new QosMonitoringMeasurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosMonitoringMeasurementWithDefaults

`func NewQosMonitoringMeasurementWithDefaults() *QosMonitoringMeasurement`

NewQosMonitoringMeasurementWithDefaults instantiates a new QosMonitoringMeasurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDlPacketDelay

`func (o *QosMonitoringMeasurement) GetDlPacketDelay() int32`

GetDlPacketDelay returns the DlPacketDelay field if non-nil, zero value otherwise.

### GetDlPacketDelayOk

`func (o *QosMonitoringMeasurement) GetDlPacketDelayOk() (*int32, bool)`

GetDlPacketDelayOk returns a tuple with the DlPacketDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlPacketDelay

`func (o *QosMonitoringMeasurement) SetDlPacketDelay(v int32)`

SetDlPacketDelay sets DlPacketDelay field to given value.

### HasDlPacketDelay

`func (o *QosMonitoringMeasurement) HasDlPacketDelay() bool`

HasDlPacketDelay returns a boolean if a field has been set.

### GetUlPacketDelay

`func (o *QosMonitoringMeasurement) GetUlPacketDelay() int32`

GetUlPacketDelay returns the UlPacketDelay field if non-nil, zero value otherwise.

### GetUlPacketDelayOk

`func (o *QosMonitoringMeasurement) GetUlPacketDelayOk() (*int32, bool)`

GetUlPacketDelayOk returns a tuple with the UlPacketDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlPacketDelay

`func (o *QosMonitoringMeasurement) SetUlPacketDelay(v int32)`

SetUlPacketDelay sets UlPacketDelay field to given value.

### HasUlPacketDelay

`func (o *QosMonitoringMeasurement) HasUlPacketDelay() bool`

HasUlPacketDelay returns a boolean if a field has been set.

### GetRtrPacketDelay

`func (o *QosMonitoringMeasurement) GetRtrPacketDelay() int32`

GetRtrPacketDelay returns the RtrPacketDelay field if non-nil, zero value otherwise.

### GetRtrPacketDelayOk

`func (o *QosMonitoringMeasurement) GetRtrPacketDelayOk() (*int32, bool)`

GetRtrPacketDelayOk returns a tuple with the RtrPacketDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtrPacketDelay

`func (o *QosMonitoringMeasurement) SetRtrPacketDelay(v int32)`

SetRtrPacketDelay sets RtrPacketDelay field to given value.

### HasRtrPacketDelay

`func (o *QosMonitoringMeasurement) HasRtrPacketDelay() bool`

HasRtrPacketDelay returns a boolean if a field has been set.

### GetMeasureFailure

`func (o *QosMonitoringMeasurement) GetMeasureFailure() bool`

GetMeasureFailure returns the MeasureFailure field if non-nil, zero value otherwise.

### GetMeasureFailureOk

`func (o *QosMonitoringMeasurement) GetMeasureFailureOk() (*bool, bool)`

GetMeasureFailureOk returns a tuple with the MeasureFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureFailure

`func (o *QosMonitoringMeasurement) SetMeasureFailure(v bool)`

SetMeasureFailure sets MeasureFailure field to given value.

### HasMeasureFailure

`func (o *QosMonitoringMeasurement) HasMeasureFailure() bool`

HasMeasureFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


