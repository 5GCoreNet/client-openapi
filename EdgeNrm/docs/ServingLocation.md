# ServingLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeographicalLocation** | Pointer to [**GeoLoc**](GeoLoc.md) |  | [optional] 
**TopologicalLocation** | Pointer to [**TopologicalServiceArea**](TopologicalServiceArea.md) |  | [optional] 

## Methods

### NewServingLocation

`func NewServingLocation() *ServingLocation`

NewServingLocation instantiates a new ServingLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServingLocationWithDefaults

`func NewServingLocationWithDefaults() *ServingLocation`

NewServingLocationWithDefaults instantiates a new ServingLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeographicalLocation

`func (o *ServingLocation) GetGeographicalLocation() GeoLoc`

GetGeographicalLocation returns the GeographicalLocation field if non-nil, zero value otherwise.

### GetGeographicalLocationOk

`func (o *ServingLocation) GetGeographicalLocationOk() (*GeoLoc, bool)`

GetGeographicalLocationOk returns a tuple with the GeographicalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalLocation

`func (o *ServingLocation) SetGeographicalLocation(v GeoLoc)`

SetGeographicalLocation sets GeographicalLocation field to given value.

### HasGeographicalLocation

`func (o *ServingLocation) HasGeographicalLocation() bool`

HasGeographicalLocation returns a boolean if a field has been set.

### GetTopologicalLocation

`func (o *ServingLocation) GetTopologicalLocation() TopologicalServiceArea`

GetTopologicalLocation returns the TopologicalLocation field if non-nil, zero value otherwise.

### GetTopologicalLocationOk

`func (o *ServingLocation) GetTopologicalLocationOk() (*TopologicalServiceArea, bool)`

GetTopologicalLocationOk returns a tuple with the TopologicalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologicalLocation

`func (o *ServingLocation) SetTopologicalLocation(v TopologicalServiceArea)`

SetTopologicalLocation sets TopologicalLocation field to given value.

### HasTopologicalLocation

`func (o *ServingLocation) HasTopologicalLocation() bool`

HasTopologicalLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


