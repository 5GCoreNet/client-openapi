# GeographicalArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**Shapes** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 

## Methods

### NewGeographicalArea

`func NewGeographicalArea() *GeographicalArea`

NewGeographicalArea instantiates a new GeographicalArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeographicalAreaWithDefaults

`func NewGeographicalAreaWithDefaults() *GeographicalArea`

NewGeographicalAreaWithDefaults instantiates a new GeographicalArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCivicAddress

`func (o *GeographicalArea) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *GeographicalArea) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *GeographicalArea) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *GeographicalArea) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetShapes

`func (o *GeographicalArea) GetShapes() GeographicArea`

GetShapes returns the Shapes field if non-nil, zero value otherwise.

### GetShapesOk

`func (o *GeographicalArea) GetShapesOk() (*GeographicArea, bool)`

GetShapesOk returns a tuple with the Shapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapes

`func (o *GeographicalArea) SetShapes(v GeographicArea)`

SetShapes sets Shapes field to given value.

### HasShapes

`func (o *GeographicalArea) HasShapes() bool`

HasShapes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


