# AefLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CivicAddr** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**GeoArea** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**DcId** | Pointer to **string** | Identifies the data center where the AEF providing the service API is located.  | [optional] 

## Methods

### NewAefLocation

`func NewAefLocation() *AefLocation`

NewAefLocation instantiates a new AefLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAefLocationWithDefaults

`func NewAefLocationWithDefaults() *AefLocation`

NewAefLocationWithDefaults instantiates a new AefLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCivicAddr

`func (o *AefLocation) GetCivicAddr() CivicAddress`

GetCivicAddr returns the CivicAddr field if non-nil, zero value otherwise.

### GetCivicAddrOk

`func (o *AefLocation) GetCivicAddrOk() (*CivicAddress, bool)`

GetCivicAddrOk returns a tuple with the CivicAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddr

`func (o *AefLocation) SetCivicAddr(v CivicAddress)`

SetCivicAddr sets CivicAddr field to given value.

### HasCivicAddr

`func (o *AefLocation) HasCivicAddr() bool`

HasCivicAddr returns a boolean if a field has been set.

### GetGeoArea

`func (o *AefLocation) GetGeoArea() GeographicArea`

GetGeoArea returns the GeoArea field if non-nil, zero value otherwise.

### GetGeoAreaOk

`func (o *AefLocation) GetGeoAreaOk() (*GeographicArea, bool)`

GetGeoAreaOk returns a tuple with the GeoArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoArea

`func (o *AefLocation) SetGeoArea(v GeographicArea)`

SetGeoArea sets GeoArea field to given value.

### HasGeoArea

`func (o *AefLocation) HasGeoArea() bool`

HasGeoArea returns a boolean if a field has been set.

### GetDcId

`func (o *AefLocation) GetDcId() string`

GetDcId returns the DcId field if non-nil, zero value otherwise.

### GetDcIdOk

`func (o *AefLocation) GetDcIdOk() (*string, bool)`

GetDcIdOk returns a tuple with the DcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcId

`func (o *AefLocation) SetDcId(v string)`

SetDcId sets DcId field to given value.

### HasDcId

`func (o *AefLocation) HasDcId() bool`

HasDcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


