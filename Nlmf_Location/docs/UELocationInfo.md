# UELocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationEstimate** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**AgeOfLocationEstimate** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**TimestampOfLocationEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**VelocityEstimate** | Pointer to [**VelocityEstimate**](VelocityEstimate.md) |  | [optional] 
**AgeOfVelocityEstimate** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**TimestampOfVelocityEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewUELocationInfo

`func NewUELocationInfo() *UELocationInfo`

NewUELocationInfo instantiates a new UELocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUELocationInfoWithDefaults

`func NewUELocationInfoWithDefaults() *UELocationInfo`

NewUELocationInfoWithDefaults instantiates a new UELocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationEstimate

`func (o *UELocationInfo) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *UELocationInfo) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *UELocationInfo) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.

### HasLocationEstimate

`func (o *UELocationInfo) HasLocationEstimate() bool`

HasLocationEstimate returns a boolean if a field has been set.

### GetAgeOfLocationEstimate

`func (o *UELocationInfo) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *UELocationInfo) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *UELocationInfo) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.

### HasAgeOfLocationEstimate

`func (o *UELocationInfo) HasAgeOfLocationEstimate() bool`

HasAgeOfLocationEstimate returns a boolean if a field has been set.

### GetTimestampOfLocationEstimate

`func (o *UELocationInfo) GetTimestampOfLocationEstimate() time.Time`

GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field if non-nil, zero value otherwise.

### GetTimestampOfLocationEstimateOk

`func (o *UELocationInfo) GetTimestampOfLocationEstimateOk() (*time.Time, bool)`

GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfLocationEstimate

`func (o *UELocationInfo) SetTimestampOfLocationEstimate(v time.Time)`

SetTimestampOfLocationEstimate sets TimestampOfLocationEstimate field to given value.

### HasTimestampOfLocationEstimate

`func (o *UELocationInfo) HasTimestampOfLocationEstimate() bool`

HasTimestampOfLocationEstimate returns a boolean if a field has been set.

### GetVelocityEstimate

`func (o *UELocationInfo) GetVelocityEstimate() VelocityEstimate`

GetVelocityEstimate returns the VelocityEstimate field if non-nil, zero value otherwise.

### GetVelocityEstimateOk

`func (o *UELocationInfo) GetVelocityEstimateOk() (*VelocityEstimate, bool)`

GetVelocityEstimateOk returns a tuple with the VelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityEstimate

`func (o *UELocationInfo) SetVelocityEstimate(v VelocityEstimate)`

SetVelocityEstimate sets VelocityEstimate field to given value.

### HasVelocityEstimate

`func (o *UELocationInfo) HasVelocityEstimate() bool`

HasVelocityEstimate returns a boolean if a field has been set.

### GetAgeOfVelocityEstimate

`func (o *UELocationInfo) GetAgeOfVelocityEstimate() int32`

GetAgeOfVelocityEstimate returns the AgeOfVelocityEstimate field if non-nil, zero value otherwise.

### GetAgeOfVelocityEstimateOk

`func (o *UELocationInfo) GetAgeOfVelocityEstimateOk() (*int32, bool)`

GetAgeOfVelocityEstimateOk returns a tuple with the AgeOfVelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfVelocityEstimate

`func (o *UELocationInfo) SetAgeOfVelocityEstimate(v int32)`

SetAgeOfVelocityEstimate sets AgeOfVelocityEstimate field to given value.

### HasAgeOfVelocityEstimate

`func (o *UELocationInfo) HasAgeOfVelocityEstimate() bool`

HasAgeOfVelocityEstimate returns a boolean if a field has been set.

### GetTimestampOfVelocityEstimate

`func (o *UELocationInfo) GetTimestampOfVelocityEstimate() time.Time`

GetTimestampOfVelocityEstimate returns the TimestampOfVelocityEstimate field if non-nil, zero value otherwise.

### GetTimestampOfVelocityEstimateOk

`func (o *UELocationInfo) GetTimestampOfVelocityEstimateOk() (*time.Time, bool)`

GetTimestampOfVelocityEstimateOk returns a tuple with the TimestampOfVelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfVelocityEstimate

`func (o *UELocationInfo) SetTimestampOfVelocityEstimate(v time.Time)`

SetTimestampOfVelocityEstimate sets TimestampOfVelocityEstimate field to given value.

### HasTimestampOfVelocityEstimate

`func (o *UELocationInfo) HasTimestampOfVelocityEstimate() bool`

HasTimestampOfVelocityEstimate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


