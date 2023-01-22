# GeoAreaToCellMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvexGeoPolygon** | Pointer to [**[]GeoCoordinate**](GeoCoordinate.md) |  | [optional] 
**AssociationThreshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewGeoAreaToCellMapping

`func NewGeoAreaToCellMapping() *GeoAreaToCellMapping`

NewGeoAreaToCellMapping instantiates a new GeoAreaToCellMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoAreaToCellMappingWithDefaults

`func NewGeoAreaToCellMappingWithDefaults() *GeoAreaToCellMapping`

NewGeoAreaToCellMappingWithDefaults instantiates a new GeoAreaToCellMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvexGeoPolygon

`func (o *GeoAreaToCellMapping) GetConvexGeoPolygon() []GeoCoordinate`

GetConvexGeoPolygon returns the ConvexGeoPolygon field if non-nil, zero value otherwise.

### GetConvexGeoPolygonOk

`func (o *GeoAreaToCellMapping) GetConvexGeoPolygonOk() (*[]GeoCoordinate, bool)`

GetConvexGeoPolygonOk returns a tuple with the ConvexGeoPolygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvexGeoPolygon

`func (o *GeoAreaToCellMapping) SetConvexGeoPolygon(v []GeoCoordinate)`

SetConvexGeoPolygon sets ConvexGeoPolygon field to given value.

### HasConvexGeoPolygon

`func (o *GeoAreaToCellMapping) HasConvexGeoPolygon() bool`

HasConvexGeoPolygon returns a boolean if a field has been set.

### GetAssociationThreshold

`func (o *GeoAreaToCellMapping) GetAssociationThreshold() int32`

GetAssociationThreshold returns the AssociationThreshold field if non-nil, zero value otherwise.

### GetAssociationThresholdOk

`func (o *GeoAreaToCellMapping) GetAssociationThresholdOk() (*int32, bool)`

GetAssociationThresholdOk returns a tuple with the AssociationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationThreshold

`func (o *GeoAreaToCellMapping) SetAssociationThreshold(v int32)`

SetAssociationThreshold sets AssociationThreshold field to given value.

### HasAssociationThreshold

`func (o *GeoAreaToCellMapping) HasAssociationThreshold() bool`

HasAssociationThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


