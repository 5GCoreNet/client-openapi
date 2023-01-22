# MeasurementData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DlDelay** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**UlDelay** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**RtDelay** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AvgPlr** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 
**AvgDataRate** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxDataRate** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**AvrDlTrafficVol** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AvrUlTrafficVol** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewMeasurementData

`func NewMeasurementData() *MeasurementData`

NewMeasurementData instantiates a new MeasurementData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementDataWithDefaults

`func NewMeasurementDataWithDefaults() *MeasurementData`

NewMeasurementDataWithDefaults instantiates a new MeasurementData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDlDelay

`func (o *MeasurementData) GetDlDelay() int32`

GetDlDelay returns the DlDelay field if non-nil, zero value otherwise.

### GetDlDelayOk

`func (o *MeasurementData) GetDlDelayOk() (*int32, bool)`

GetDlDelayOk returns a tuple with the DlDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlDelay

`func (o *MeasurementData) SetDlDelay(v int32)`

SetDlDelay sets DlDelay field to given value.

### HasDlDelay

`func (o *MeasurementData) HasDlDelay() bool`

HasDlDelay returns a boolean if a field has been set.

### GetUlDelay

`func (o *MeasurementData) GetUlDelay() int32`

GetUlDelay returns the UlDelay field if non-nil, zero value otherwise.

### GetUlDelayOk

`func (o *MeasurementData) GetUlDelayOk() (*int32, bool)`

GetUlDelayOk returns a tuple with the UlDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlDelay

`func (o *MeasurementData) SetUlDelay(v int32)`

SetUlDelay sets UlDelay field to given value.

### HasUlDelay

`func (o *MeasurementData) HasUlDelay() bool`

HasUlDelay returns a boolean if a field has been set.

### GetRtDelay

`func (o *MeasurementData) GetRtDelay() int32`

GetRtDelay returns the RtDelay field if non-nil, zero value otherwise.

### GetRtDelayOk

`func (o *MeasurementData) GetRtDelayOk() (*int32, bool)`

GetRtDelayOk returns a tuple with the RtDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtDelay

`func (o *MeasurementData) SetRtDelay(v int32)`

SetRtDelay sets RtDelay field to given value.

### HasRtDelay

`func (o *MeasurementData) HasRtDelay() bool`

HasRtDelay returns a boolean if a field has been set.

### GetAvgPlr

`func (o *MeasurementData) GetAvgPlr() int32`

GetAvgPlr returns the AvgPlr field if non-nil, zero value otherwise.

### GetAvgPlrOk

`func (o *MeasurementData) GetAvgPlrOk() (*int32, bool)`

GetAvgPlrOk returns a tuple with the AvgPlr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPlr

`func (o *MeasurementData) SetAvgPlr(v int32)`

SetAvgPlr sets AvgPlr field to given value.

### HasAvgPlr

`func (o *MeasurementData) HasAvgPlr() bool`

HasAvgPlr returns a boolean if a field has been set.

### GetAvgDataRate

`func (o *MeasurementData) GetAvgDataRate() string`

GetAvgDataRate returns the AvgDataRate field if non-nil, zero value otherwise.

### GetAvgDataRateOk

`func (o *MeasurementData) GetAvgDataRateOk() (*string, bool)`

GetAvgDataRateOk returns a tuple with the AvgDataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDataRate

`func (o *MeasurementData) SetAvgDataRate(v string)`

SetAvgDataRate sets AvgDataRate field to given value.

### HasAvgDataRate

`func (o *MeasurementData) HasAvgDataRate() bool`

HasAvgDataRate returns a boolean if a field has been set.

### GetMaxDataRate

`func (o *MeasurementData) GetMaxDataRate() string`

GetMaxDataRate returns the MaxDataRate field if non-nil, zero value otherwise.

### GetMaxDataRateOk

`func (o *MeasurementData) GetMaxDataRateOk() (*string, bool)`

GetMaxDataRateOk returns a tuple with the MaxDataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRate

`func (o *MeasurementData) SetMaxDataRate(v string)`

SetMaxDataRate sets MaxDataRate field to given value.

### HasMaxDataRate

`func (o *MeasurementData) HasMaxDataRate() bool`

HasMaxDataRate returns a boolean if a field has been set.

### GetAvrDlTrafficVol

`func (o *MeasurementData) GetAvrDlTrafficVol() int32`

GetAvrDlTrafficVol returns the AvrDlTrafficVol field if non-nil, zero value otherwise.

### GetAvrDlTrafficVolOk

`func (o *MeasurementData) GetAvrDlTrafficVolOk() (*int32, bool)`

GetAvrDlTrafficVolOk returns a tuple with the AvrDlTrafficVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvrDlTrafficVol

`func (o *MeasurementData) SetAvrDlTrafficVol(v int32)`

SetAvrDlTrafficVol sets AvrDlTrafficVol field to given value.

### HasAvrDlTrafficVol

`func (o *MeasurementData) HasAvrDlTrafficVol() bool`

HasAvrDlTrafficVol returns a boolean if a field has been set.

### GetAvrUlTrafficVol

`func (o *MeasurementData) GetAvrUlTrafficVol() int32`

GetAvrUlTrafficVol returns the AvrUlTrafficVol field if non-nil, zero value otherwise.

### GetAvrUlTrafficVolOk

`func (o *MeasurementData) GetAvrUlTrafficVolOk() (*int32, bool)`

GetAvrUlTrafficVolOk returns a tuple with the AvrUlTrafficVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvrUlTrafficVol

`func (o *MeasurementData) SetAvrUlTrafficVol(v int32)`

SetAvrUlTrafficVol sets AvrUlTrafficVol field to given value.

### HasAvrUlTrafficVol

`func (o *MeasurementData) HasAvrUlTrafficVol() bool`

HasAvrUlTrafficVol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


