# GeoArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to [**[]Coordinate**](Coordinate.md) |  | [optional] 
**Altitude** | Pointer to **float32** |  | [optional] 

## Methods

### NewGeoArea

`func NewGeoArea() *GeoArea`

NewGeoArea instantiates a new GeoArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoAreaWithDefaults

`func NewGeoAreaWithDefaults() *GeoArea`

NewGeoAreaWithDefaults instantiates a new GeoArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *GeoArea) GetCoordinates() []Coordinate`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoArea) GetCoordinatesOk() (*[]Coordinate, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoArea) SetCoordinates(v []Coordinate)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *GeoArea) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetAltitude

`func (o *GeoArea) GetAltitude() float32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *GeoArea) GetAltitudeOk() (*float32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *GeoArea) SetAltitude(v float32)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *GeoArea) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


