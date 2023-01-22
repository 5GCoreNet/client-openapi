# UmtLocationArea5G

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeographicAreas** | Pointer to [**[]GeographicArea**](GeographicArea.md) | Identifies a list of geographic area of the user where the UE is located. | [optional] 
**CivicAddresses** | Pointer to [**[]CivicAddress**](CivicAddress.md) | Identifies a list of civic addresses of the user where the UE is located. | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**UmtTime** | Pointer to **string** | String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC). | [optional] 
**UmtDuration** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 

## Methods

### NewUmtLocationArea5G

`func NewUmtLocationArea5G() *UmtLocationArea5G`

NewUmtLocationArea5G instantiates a new UmtLocationArea5G object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmtLocationArea5GWithDefaults

`func NewUmtLocationArea5GWithDefaults() *UmtLocationArea5G`

NewUmtLocationArea5GWithDefaults instantiates a new UmtLocationArea5G object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeographicAreas

`func (o *UmtLocationArea5G) GetGeographicAreas() []GeographicArea`

GetGeographicAreas returns the GeographicAreas field if non-nil, zero value otherwise.

### GetGeographicAreasOk

`func (o *UmtLocationArea5G) GetGeographicAreasOk() (*[]GeographicArea, bool)`

GetGeographicAreasOk returns a tuple with the GeographicAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicAreas

`func (o *UmtLocationArea5G) SetGeographicAreas(v []GeographicArea)`

SetGeographicAreas sets GeographicAreas field to given value.

### HasGeographicAreas

`func (o *UmtLocationArea5G) HasGeographicAreas() bool`

HasGeographicAreas returns a boolean if a field has been set.

### GetCivicAddresses

`func (o *UmtLocationArea5G) GetCivicAddresses() []CivicAddress`

GetCivicAddresses returns the CivicAddresses field if non-nil, zero value otherwise.

### GetCivicAddressesOk

`func (o *UmtLocationArea5G) GetCivicAddressesOk() (*[]CivicAddress, bool)`

GetCivicAddressesOk returns a tuple with the CivicAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddresses

`func (o *UmtLocationArea5G) SetCivicAddresses(v []CivicAddress)`

SetCivicAddresses sets CivicAddresses field to given value.

### HasCivicAddresses

`func (o *UmtLocationArea5G) HasCivicAddresses() bool`

HasCivicAddresses returns a boolean if a field has been set.

### GetNwAreaInfo

`func (o *UmtLocationArea5G) GetNwAreaInfo() NetworkAreaInfo`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *UmtLocationArea5G) GetNwAreaInfoOk() (*NetworkAreaInfo, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *UmtLocationArea5G) SetNwAreaInfo(v NetworkAreaInfo)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *UmtLocationArea5G) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetUmtTime

`func (o *UmtLocationArea5G) GetUmtTime() string`

GetUmtTime returns the UmtTime field if non-nil, zero value otherwise.

### GetUmtTimeOk

`func (o *UmtLocationArea5G) GetUmtTimeOk() (*string, bool)`

GetUmtTimeOk returns a tuple with the UmtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmtTime

`func (o *UmtLocationArea5G) SetUmtTime(v string)`

SetUmtTime sets UmtTime field to given value.

### HasUmtTime

`func (o *UmtLocationArea5G) HasUmtTime() bool`

HasUmtTime returns a boolean if a field has been set.

### GetUmtDuration

`func (o *UmtLocationArea5G) GetUmtDuration() int32`

GetUmtDuration returns the UmtDuration field if non-nil, zero value otherwise.

### GetUmtDurationOk

`func (o *UmtLocationArea5G) GetUmtDurationOk() (*int32, bool)`

GetUmtDurationOk returns a tuple with the UmtDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmtDuration

`func (o *UmtLocationArea5G) SetUmtDuration(v int32)`

SetUmtDuration sets UmtDuration field to given value.

### HasUmtDuration

`func (o *UmtLocationArea5G) HasUmtDuration() bool`

HasUmtDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


