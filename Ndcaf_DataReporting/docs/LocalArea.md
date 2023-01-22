# LocalArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shape** | [**SupportedGADShapes**](SupportedGADShapes.md) |  | 
**LocalOrigin** | [**LocalOrigin**](LocalOrigin.md) |  | 
**Point** | [**RelativeCartesianLocation**](RelativeCartesianLocation.md) |  | 
**UncertaintyEllipsoid** | [**UncertaintyEllipsoid**](UncertaintyEllipsoid.md) |  | 
**Confidence** | **int32** | Indicates value of confidence. | 

## Methods

### NewLocalArea

`func NewLocalArea(shape SupportedGADShapes, localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipsoid UncertaintyEllipsoid, confidence int32, ) *LocalArea`

NewLocalArea instantiates a new LocalArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalAreaWithDefaults

`func NewLocalAreaWithDefaults() *LocalArea`

NewLocalAreaWithDefaults instantiates a new LocalArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShape

`func (o *LocalArea) GetShape() SupportedGADShapes`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *LocalArea) GetShapeOk() (*SupportedGADShapes, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *LocalArea) SetShape(v SupportedGADShapes)`

SetShape sets Shape field to given value.


### GetLocalOrigin

`func (o *LocalArea) GetLocalOrigin() LocalOrigin`

GetLocalOrigin returns the LocalOrigin field if non-nil, zero value otherwise.

### GetLocalOriginOk

`func (o *LocalArea) GetLocalOriginOk() (*LocalOrigin, bool)`

GetLocalOriginOk returns a tuple with the LocalOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalOrigin

`func (o *LocalArea) SetLocalOrigin(v LocalOrigin)`

SetLocalOrigin sets LocalOrigin field to given value.


### GetPoint

`func (o *LocalArea) GetPoint() RelativeCartesianLocation`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *LocalArea) GetPointOk() (*RelativeCartesianLocation, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *LocalArea) SetPoint(v RelativeCartesianLocation)`

SetPoint sets Point field to given value.


### GetUncertaintyEllipsoid

`func (o *LocalArea) GetUncertaintyEllipsoid() UncertaintyEllipsoid`

GetUncertaintyEllipsoid returns the UncertaintyEllipsoid field if non-nil, zero value otherwise.

### GetUncertaintyEllipsoidOk

`func (o *LocalArea) GetUncertaintyEllipsoidOk() (*UncertaintyEllipsoid, bool)`

GetUncertaintyEllipsoidOk returns a tuple with the UncertaintyEllipsoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyEllipsoid

`func (o *LocalArea) SetUncertaintyEllipsoid(v UncertaintyEllipsoid)`

SetUncertaintyEllipsoid sets UncertaintyEllipsoid field to given value.


### GetConfidence

`func (o *LocalArea) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *LocalArea) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *LocalArea) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


