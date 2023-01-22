# GeoLoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeographicalCoordinates** | Pointer to [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | [optional] 
**CivicLocation** | Pointer to **string** |  | [optional] 

## Methods

### NewGeoLoc

`func NewGeoLoc() *GeoLoc`

NewGeoLoc instantiates a new GeoLoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoLocWithDefaults

`func NewGeoLocWithDefaults() *GeoLoc`

NewGeoLocWithDefaults instantiates a new GeoLoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeographicalCoordinates

`func (o *GeoLoc) GetGeographicalCoordinates() GeographicalCoordinates`

GetGeographicalCoordinates returns the GeographicalCoordinates field if non-nil, zero value otherwise.

### GetGeographicalCoordinatesOk

`func (o *GeoLoc) GetGeographicalCoordinatesOk() (*GeographicalCoordinates, bool)`

GetGeographicalCoordinatesOk returns a tuple with the GeographicalCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalCoordinates

`func (o *GeoLoc) SetGeographicalCoordinates(v GeographicalCoordinates)`

SetGeographicalCoordinates sets GeographicalCoordinates field to given value.

### HasGeographicalCoordinates

`func (o *GeoLoc) HasGeographicalCoordinates() bool`

HasGeographicalCoordinates returns a boolean if a field has been set.

### GetCivicLocation

`func (o *GeoLoc) GetCivicLocation() string`

GetCivicLocation returns the CivicLocation field if non-nil, zero value otherwise.

### GetCivicLocationOk

`func (o *GeoLoc) GetCivicLocationOk() (*string, bool)`

GetCivicLocationOk returns a tuple with the CivicLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicLocation

`func (o *GeoLoc) SetCivicLocation(v string)`

SetCivicLocation sets CivicLocation field to given value.

### HasCivicLocation

`func (o *GeoLoc) HasCivicLocation() bool`

HasCivicLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


