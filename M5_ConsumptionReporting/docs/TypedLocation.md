# TypedLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationIdentifierType** | [**CellIdentifierType**](CellIdentifierType.md) |  | 
**Location** | **string** |  | 

## Methods

### NewTypedLocation

`func NewTypedLocation(locationIdentifierType CellIdentifierType, location string, ) *TypedLocation`

NewTypedLocation instantiates a new TypedLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedLocationWithDefaults

`func NewTypedLocationWithDefaults() *TypedLocation`

NewTypedLocationWithDefaults instantiates a new TypedLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationIdentifierType

`func (o *TypedLocation) GetLocationIdentifierType() CellIdentifierType`

GetLocationIdentifierType returns the LocationIdentifierType field if non-nil, zero value otherwise.

### GetLocationIdentifierTypeOk

`func (o *TypedLocation) GetLocationIdentifierTypeOk() (*CellIdentifierType, bool)`

GetLocationIdentifierTypeOk returns a tuple with the LocationIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIdentifierType

`func (o *TypedLocation) SetLocationIdentifierType(v CellIdentifierType)`

SetLocationIdentifierType sets LocationIdentifierType field to given value.


### GetLocation

`func (o *TypedLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TypedLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TypedLocation) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


