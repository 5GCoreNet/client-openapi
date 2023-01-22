# LocationReportingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentLocation** | **bool** |  | 
**Accuracy** | Pointer to [**LocationAccuracy**](LocationAccuracy.md) |  | [optional] 

## Methods

### NewLocationReportingConfiguration

`func NewLocationReportingConfiguration(currentLocation bool, ) *LocationReportingConfiguration`

NewLocationReportingConfiguration instantiates a new LocationReportingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationReportingConfigurationWithDefaults

`func NewLocationReportingConfigurationWithDefaults() *LocationReportingConfiguration`

NewLocationReportingConfigurationWithDefaults instantiates a new LocationReportingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentLocation

`func (o *LocationReportingConfiguration) GetCurrentLocation() bool`

GetCurrentLocation returns the CurrentLocation field if non-nil, zero value otherwise.

### GetCurrentLocationOk

`func (o *LocationReportingConfiguration) GetCurrentLocationOk() (*bool, bool)`

GetCurrentLocationOk returns a tuple with the CurrentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLocation

`func (o *LocationReportingConfiguration) SetCurrentLocation(v bool)`

SetCurrentLocation sets CurrentLocation field to given value.


### GetAccuracy

`func (o *LocationReportingConfiguration) GetAccuracy() LocationAccuracy`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LocationReportingConfiguration) GetAccuracyOk() (*LocationAccuracy, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LocationReportingConfiguration) SetAccuracy(v LocationAccuracy)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LocationReportingConfiguration) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


