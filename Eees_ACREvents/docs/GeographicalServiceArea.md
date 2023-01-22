# GeographicalServiceArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoArs** | Pointer to [**[]GeographicArea**](GeographicArea.md) | A list of geographic area information. | [optional] 
**CivicAddrs** | Pointer to [**[]CivicAddress**](CivicAddress.md) | A list of civic address information. | [optional] 

## Methods

### NewGeographicalServiceArea

`func NewGeographicalServiceArea() *GeographicalServiceArea`

NewGeographicalServiceArea instantiates a new GeographicalServiceArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeographicalServiceAreaWithDefaults

`func NewGeographicalServiceAreaWithDefaults() *GeographicalServiceArea`

NewGeographicalServiceAreaWithDefaults instantiates a new GeographicalServiceArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeoArs

`func (o *GeographicalServiceArea) GetGeoArs() []GeographicArea`

GetGeoArs returns the GeoArs field if non-nil, zero value otherwise.

### GetGeoArsOk

`func (o *GeographicalServiceArea) GetGeoArsOk() (*[]GeographicArea, bool)`

GetGeoArsOk returns a tuple with the GeoArs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoArs

`func (o *GeographicalServiceArea) SetGeoArs(v []GeographicArea)`

SetGeoArs sets GeoArs field to given value.

### HasGeoArs

`func (o *GeographicalServiceArea) HasGeoArs() bool`

HasGeoArs returns a boolean if a field has been set.

### GetCivicAddrs

`func (o *GeographicalServiceArea) GetCivicAddrs() []CivicAddress`

GetCivicAddrs returns the CivicAddrs field if non-nil, zero value otherwise.

### GetCivicAddrsOk

`func (o *GeographicalServiceArea) GetCivicAddrsOk() (*[]CivicAddress, bool)`

GetCivicAddrsOk returns a tuple with the CivicAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddrs

`func (o *GeographicalServiceArea) SetCivicAddrs(v []CivicAddress)`

SetCivicAddrs sets CivicAddrs field to given value.

### HasCivicAddrs

`func (o *GeographicalServiceArea) HasCivicAddrs() bool`

HasCivicAddrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


