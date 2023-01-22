# LocationInfoCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoArea** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**RefUe** | Pointer to [**ReferenceUEDetail**](ReferenceUEDetail.md) |  | [optional] 

## Methods

### NewLocationInfoCriteria

`func NewLocationInfoCriteria() *LocationInfoCriteria`

NewLocationInfoCriteria instantiates a new LocationInfoCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationInfoCriteriaWithDefaults

`func NewLocationInfoCriteriaWithDefaults() *LocationInfoCriteria`

NewLocationInfoCriteriaWithDefaults instantiates a new LocationInfoCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeoArea

`func (o *LocationInfoCriteria) GetGeoArea() GeographicArea`

GetGeoArea returns the GeoArea field if non-nil, zero value otherwise.

### GetGeoAreaOk

`func (o *LocationInfoCriteria) GetGeoAreaOk() (*GeographicArea, bool)`

GetGeoAreaOk returns a tuple with the GeoArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoArea

`func (o *LocationInfoCriteria) SetGeoArea(v GeographicArea)`

SetGeoArea sets GeoArea field to given value.

### HasGeoArea

`func (o *LocationInfoCriteria) HasGeoArea() bool`

HasGeoArea returns a boolean if a field has been set.

### GetRefUe

`func (o *LocationInfoCriteria) GetRefUe() ReferenceUEDetail`

GetRefUe returns the RefUe field if non-nil, zero value otherwise.

### GetRefUeOk

`func (o *LocationInfoCriteria) GetRefUeOk() (*ReferenceUEDetail, bool)`

GetRefUeOk returns a tuple with the RefUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUe

`func (o *LocationInfoCriteria) SetRefUe(v ReferenceUEDetail)`

SetRefUe sets RefUe field to given value.

### HasRefUe

`func (o *LocationInfoCriteria) HasRefUe() bool`

HasRefUe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


