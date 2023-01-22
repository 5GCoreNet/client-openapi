# LocationArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellIds** | Pointer to **[]string** | Indicates a list of Cell Global Identities of the user which identifies the cell the UE is registered. | [optional] 
**EnodeBIds** | Pointer to **[]string** | Indicates a list of eNodeB identities in which the UE is currently located. | [optional] 
**RoutingAreaIds** | Pointer to **[]string** | Identifies a list of Routing Area Identities of the user where the UE is located. | [optional] 
**TrackingAreaIds** | Pointer to **[]string** | Identifies a list of Tracking Area Identities of the user where the UE is located. | [optional] 
**GeographicAreas** | Pointer to [**[]GeographicArea**](GeographicArea.md) | Identifies a list of geographic area of the user where the UE is located. | [optional] 
**CivicAddresses** | Pointer to [**[]CivicAddress**](CivicAddress.md) | Identifies a list of civic addresses of the user where the UE is located. | [optional] 

## Methods

### NewLocationArea

`func NewLocationArea() *LocationArea`

NewLocationArea instantiates a new LocationArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAreaWithDefaults

`func NewLocationAreaWithDefaults() *LocationArea`

NewLocationAreaWithDefaults instantiates a new LocationArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellIds

`func (o *LocationArea) GetCellIds() []string`

GetCellIds returns the CellIds field if non-nil, zero value otherwise.

### GetCellIdsOk

`func (o *LocationArea) GetCellIdsOk() (*[]string, bool)`

GetCellIdsOk returns a tuple with the CellIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellIds

`func (o *LocationArea) SetCellIds(v []string)`

SetCellIds sets CellIds field to given value.

### HasCellIds

`func (o *LocationArea) HasCellIds() bool`

HasCellIds returns a boolean if a field has been set.

### GetEnodeBIds

`func (o *LocationArea) GetEnodeBIds() []string`

GetEnodeBIds returns the EnodeBIds field if non-nil, zero value otherwise.

### GetEnodeBIdsOk

`func (o *LocationArea) GetEnodeBIdsOk() (*[]string, bool)`

GetEnodeBIdsOk returns a tuple with the EnodeBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnodeBIds

`func (o *LocationArea) SetEnodeBIds(v []string)`

SetEnodeBIds sets EnodeBIds field to given value.

### HasEnodeBIds

`func (o *LocationArea) HasEnodeBIds() bool`

HasEnodeBIds returns a boolean if a field has been set.

### GetRoutingAreaIds

`func (o *LocationArea) GetRoutingAreaIds() []string`

GetRoutingAreaIds returns the RoutingAreaIds field if non-nil, zero value otherwise.

### GetRoutingAreaIdsOk

`func (o *LocationArea) GetRoutingAreaIdsOk() (*[]string, bool)`

GetRoutingAreaIdsOk returns a tuple with the RoutingAreaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingAreaIds

`func (o *LocationArea) SetRoutingAreaIds(v []string)`

SetRoutingAreaIds sets RoutingAreaIds field to given value.

### HasRoutingAreaIds

`func (o *LocationArea) HasRoutingAreaIds() bool`

HasRoutingAreaIds returns a boolean if a field has been set.

### GetTrackingAreaIds

`func (o *LocationArea) GetTrackingAreaIds() []string`

GetTrackingAreaIds returns the TrackingAreaIds field if non-nil, zero value otherwise.

### GetTrackingAreaIdsOk

`func (o *LocationArea) GetTrackingAreaIdsOk() (*[]string, bool)`

GetTrackingAreaIdsOk returns a tuple with the TrackingAreaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingAreaIds

`func (o *LocationArea) SetTrackingAreaIds(v []string)`

SetTrackingAreaIds sets TrackingAreaIds field to given value.

### HasTrackingAreaIds

`func (o *LocationArea) HasTrackingAreaIds() bool`

HasTrackingAreaIds returns a boolean if a field has been set.

### GetGeographicAreas

`func (o *LocationArea) GetGeographicAreas() []GeographicArea`

GetGeographicAreas returns the GeographicAreas field if non-nil, zero value otherwise.

### GetGeographicAreasOk

`func (o *LocationArea) GetGeographicAreasOk() (*[]GeographicArea, bool)`

GetGeographicAreasOk returns a tuple with the GeographicAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicAreas

`func (o *LocationArea) SetGeographicAreas(v []GeographicArea)`

SetGeographicAreas sets GeographicAreas field to given value.

### HasGeographicAreas

`func (o *LocationArea) HasGeographicAreas() bool`

HasGeographicAreas returns a boolean if a field has been set.

### GetCivicAddresses

`func (o *LocationArea) GetCivicAddresses() []CivicAddress`

GetCivicAddresses returns the CivicAddresses field if non-nil, zero value otherwise.

### GetCivicAddressesOk

`func (o *LocationArea) GetCivicAddressesOk() (*[]CivicAddress, bool)`

GetCivicAddressesOk returns a tuple with the CivicAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddresses

`func (o *LocationArea) SetCivicAddresses(v []CivicAddress)`

SetCivicAddresses sets CivicAddresses field to given value.

### HasCivicAddresses

`func (o *LocationArea) HasCivicAddresses() bool`

HasCivicAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


