# DnPerf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServerInsAddr** | Pointer to [**AddrFqdn**](AddrFqdn.md) |  | [optional] 
**UpfInfo** | Pointer to [**UpfInformation**](UpfInformation.md) |  | [optional] 
**Dnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**PerfData** | [**PerfData**](PerfData.md) |  | 
**SpatialValidCon** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**TemporalValidCon** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 

## Methods

### NewDnPerf

`func NewDnPerf(perfData PerfData, ) *DnPerf`

NewDnPerf instantiates a new DnPerf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnPerfWithDefaults

`func NewDnPerfWithDefaults() *DnPerf`

NewDnPerfWithDefaults instantiates a new DnPerf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServerInsAddr

`func (o *DnPerf) GetAppServerInsAddr() AddrFqdn`

GetAppServerInsAddr returns the AppServerInsAddr field if non-nil, zero value otherwise.

### GetAppServerInsAddrOk

`func (o *DnPerf) GetAppServerInsAddrOk() (*AddrFqdn, bool)`

GetAppServerInsAddrOk returns a tuple with the AppServerInsAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerInsAddr

`func (o *DnPerf) SetAppServerInsAddr(v AddrFqdn)`

SetAppServerInsAddr sets AppServerInsAddr field to given value.

### HasAppServerInsAddr

`func (o *DnPerf) HasAppServerInsAddr() bool`

HasAppServerInsAddr returns a boolean if a field has been set.

### GetUpfInfo

`func (o *DnPerf) GetUpfInfo() UpfInformation`

GetUpfInfo returns the UpfInfo field if non-nil, zero value otherwise.

### GetUpfInfoOk

`func (o *DnPerf) GetUpfInfoOk() (*UpfInformation, bool)`

GetUpfInfoOk returns a tuple with the UpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfo

`func (o *DnPerf) SetUpfInfo(v UpfInformation)`

SetUpfInfo sets UpfInfo field to given value.

### HasUpfInfo

`func (o *DnPerf) HasUpfInfo() bool`

HasUpfInfo returns a boolean if a field has been set.

### GetDnai

`func (o *DnPerf) GetDnai() string`

GetDnai returns the Dnai field if non-nil, zero value otherwise.

### GetDnaiOk

`func (o *DnPerf) GetDnaiOk() (*string, bool)`

GetDnaiOk returns a tuple with the Dnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnai

`func (o *DnPerf) SetDnai(v string)`

SetDnai sets Dnai field to given value.

### HasDnai

`func (o *DnPerf) HasDnai() bool`

HasDnai returns a boolean if a field has been set.

### GetPerfData

`func (o *DnPerf) GetPerfData() PerfData`

GetPerfData returns the PerfData field if non-nil, zero value otherwise.

### GetPerfDataOk

`func (o *DnPerf) GetPerfDataOk() (*PerfData, bool)`

GetPerfDataOk returns a tuple with the PerfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfData

`func (o *DnPerf) SetPerfData(v PerfData)`

SetPerfData sets PerfData field to given value.


### GetSpatialValidCon

`func (o *DnPerf) GetSpatialValidCon() NetworkAreaInfo`

GetSpatialValidCon returns the SpatialValidCon field if non-nil, zero value otherwise.

### GetSpatialValidConOk

`func (o *DnPerf) GetSpatialValidConOk() (*NetworkAreaInfo, bool)`

GetSpatialValidConOk returns a tuple with the SpatialValidCon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidCon

`func (o *DnPerf) SetSpatialValidCon(v NetworkAreaInfo)`

SetSpatialValidCon sets SpatialValidCon field to given value.

### HasSpatialValidCon

`func (o *DnPerf) HasSpatialValidCon() bool`

HasSpatialValidCon returns a boolean if a field has been set.

### GetTemporalValidCon

`func (o *DnPerf) GetTemporalValidCon() TimeWindow`

GetTemporalValidCon returns the TemporalValidCon field if non-nil, zero value otherwise.

### GetTemporalValidConOk

`func (o *DnPerf) GetTemporalValidConOk() (*TimeWindow, bool)`

GetTemporalValidConOk returns a tuple with the TemporalValidCon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalValidCon

`func (o *DnPerf) SetTemporalValidCon(v TimeWindow)`

SetTemporalValidCon sets TemporalValidCon field to given value.

### HasTemporalValidCon

`func (o *DnPerf) HasTemporalValidCon() bool`

HasTemporalValidCon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


