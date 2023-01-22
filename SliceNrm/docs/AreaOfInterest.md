# AreaOfInterest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvexGeoPolygon** | Pointer to [**[]GeoCoordinate**](GeoCoordinate.md) |  | [optional] 
**AssociationThreshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewAreaOfInterest

`func NewAreaOfInterest() *AreaOfInterest`

NewAreaOfInterest instantiates a new AreaOfInterest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAreaOfInterestWithDefaults

`func NewAreaOfInterestWithDefaults() *AreaOfInterest`

NewAreaOfInterestWithDefaults instantiates a new AreaOfInterest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvexGeoPolygon

`func (o *AreaOfInterest) GetConvexGeoPolygon() []GeoCoordinate`

GetConvexGeoPolygon returns the ConvexGeoPolygon field if non-nil, zero value otherwise.

### GetConvexGeoPolygonOk

`func (o *AreaOfInterest) GetConvexGeoPolygonOk() (*[]GeoCoordinate, bool)`

GetConvexGeoPolygonOk returns a tuple with the ConvexGeoPolygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvexGeoPolygon

`func (o *AreaOfInterest) SetConvexGeoPolygon(v []GeoCoordinate)`

SetConvexGeoPolygon sets ConvexGeoPolygon field to given value.

### HasConvexGeoPolygon

`func (o *AreaOfInterest) HasConvexGeoPolygon() bool`

HasConvexGeoPolygon returns a boolean if a field has been set.

### GetAssociationThreshold

`func (o *AreaOfInterest) GetAssociationThreshold() int32`

GetAssociationThreshold returns the AssociationThreshold field if non-nil, zero value otherwise.

### GetAssociationThresholdOk

`func (o *AreaOfInterest) GetAssociationThresholdOk() (*int32, bool)`

GetAssociationThresholdOk returns a tuple with the AssociationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationThreshold

`func (o *AreaOfInterest) SetAssociationThreshold(v int32)`

SetAssociationThreshold sets AssociationThreshold field to given value.

### HasAssociationThreshold

`func (o *AreaOfInterest) HasAssociationThreshold() bool`

HasAssociationThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


