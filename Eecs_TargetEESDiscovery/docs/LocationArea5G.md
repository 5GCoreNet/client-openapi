# LocationArea5G

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeographicAreas** | Pointer to [**[]GeographicArea**](GeographicArea.md) | Identifies a list of geographic area of the user where the UE is located. | [optional] 
**CivicAddresses** | Pointer to [**[]CivicAddress**](CivicAddress.md) | Identifies a list of civic addresses of the user where the UE is located. | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 

## Methods

### NewLocationArea5G

`func NewLocationArea5G() *LocationArea5G`

NewLocationArea5G instantiates a new LocationArea5G object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationArea5GWithDefaults

`func NewLocationArea5GWithDefaults() *LocationArea5G`

NewLocationArea5GWithDefaults instantiates a new LocationArea5G object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeographicAreas

`func (o *LocationArea5G) GetGeographicAreas() []GeographicArea`

GetGeographicAreas returns the GeographicAreas field if non-nil, zero value otherwise.

### GetGeographicAreasOk

`func (o *LocationArea5G) GetGeographicAreasOk() (*[]GeographicArea, bool)`

GetGeographicAreasOk returns a tuple with the GeographicAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicAreas

`func (o *LocationArea5G) SetGeographicAreas(v []GeographicArea)`

SetGeographicAreas sets GeographicAreas field to given value.

### HasGeographicAreas

`func (o *LocationArea5G) HasGeographicAreas() bool`

HasGeographicAreas returns a boolean if a field has been set.

### GetCivicAddresses

`func (o *LocationArea5G) GetCivicAddresses() []CivicAddress`

GetCivicAddresses returns the CivicAddresses field if non-nil, zero value otherwise.

### GetCivicAddressesOk

`func (o *LocationArea5G) GetCivicAddressesOk() (*[]CivicAddress, bool)`

GetCivicAddressesOk returns a tuple with the CivicAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddresses

`func (o *LocationArea5G) SetCivicAddresses(v []CivicAddress)`

SetCivicAddresses sets CivicAddresses field to given value.

### HasCivicAddresses

`func (o *LocationArea5G) HasCivicAddresses() bool`

HasCivicAddresses returns a boolean if a field has been set.

### GetNwAreaInfo

`func (o *LocationArea5G) GetNwAreaInfo() NetworkAreaInfo`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *LocationArea5G) GetNwAreaInfoOk() (*NetworkAreaInfo, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *LocationArea5G) SetNwAreaInfo(v NetworkAreaInfo)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *LocationArea5G) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


