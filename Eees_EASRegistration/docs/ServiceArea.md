# ServiceArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopServAr** | Pointer to [**TopologicalServiceArea**](TopologicalServiceArea.md) |  | [optional] 
**GeoServAr** | Pointer to [**GeographicalServiceArea**](GeographicalServiceArea.md) |  | [optional] 

## Methods

### NewServiceArea

`func NewServiceArea() *ServiceArea`

NewServiceArea instantiates a new ServiceArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAreaWithDefaults

`func NewServiceAreaWithDefaults() *ServiceArea`

NewServiceAreaWithDefaults instantiates a new ServiceArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopServAr

`func (o *ServiceArea) GetTopServAr() TopologicalServiceArea`

GetTopServAr returns the TopServAr field if non-nil, zero value otherwise.

### GetTopServArOk

`func (o *ServiceArea) GetTopServArOk() (*TopologicalServiceArea, bool)`

GetTopServArOk returns a tuple with the TopServAr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopServAr

`func (o *ServiceArea) SetTopServAr(v TopologicalServiceArea)`

SetTopServAr sets TopServAr field to given value.

### HasTopServAr

`func (o *ServiceArea) HasTopServAr() bool`

HasTopServAr returns a boolean if a field has been set.

### GetGeoServAr

`func (o *ServiceArea) GetGeoServAr() GeographicalServiceArea`

GetGeoServAr returns the GeoServAr field if non-nil, zero value otherwise.

### GetGeoServArOk

`func (o *ServiceArea) GetGeoServArOk() (*GeographicalServiceArea, bool)`

GetGeoServArOk returns a tuple with the GeoServAr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoServAr

`func (o *ServiceArea) SetGeoServAr(v GeographicalServiceArea)`

SetGeoServAr sets GeoServAr field to given value.

### HasGeoServAr

`func (o *ServiceArea) HasGeoServAr() bool`

HasGeoServAr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


