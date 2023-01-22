# ConsumptionReportingUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaConsumed** | **string** |  | 
**MediaEndpointAddress** | Pointer to [**EndpointAddress**](EndpointAddress.md) |  | [optional] 
**StartTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**Duration** | **int32** | indicating a time in seconds. | 
**Locations** | Pointer to [**[]TypedLocation**](TypedLocation.md) |  | [optional] 

## Methods

### NewConsumptionReportingUnit

`func NewConsumptionReportingUnit(mediaConsumed string, startTime time.Time, duration int32, ) *ConsumptionReportingUnit`

NewConsumptionReportingUnit instantiates a new ConsumptionReportingUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumptionReportingUnitWithDefaults

`func NewConsumptionReportingUnitWithDefaults() *ConsumptionReportingUnit`

NewConsumptionReportingUnitWithDefaults instantiates a new ConsumptionReportingUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaConsumed

`func (o *ConsumptionReportingUnit) GetMediaConsumed() string`

GetMediaConsumed returns the MediaConsumed field if non-nil, zero value otherwise.

### GetMediaConsumedOk

`func (o *ConsumptionReportingUnit) GetMediaConsumedOk() (*string, bool)`

GetMediaConsumedOk returns a tuple with the MediaConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaConsumed

`func (o *ConsumptionReportingUnit) SetMediaConsumed(v string)`

SetMediaConsumed sets MediaConsumed field to given value.


### GetMediaEndpointAddress

`func (o *ConsumptionReportingUnit) GetMediaEndpointAddress() EndpointAddress`

GetMediaEndpointAddress returns the MediaEndpointAddress field if non-nil, zero value otherwise.

### GetMediaEndpointAddressOk

`func (o *ConsumptionReportingUnit) GetMediaEndpointAddressOk() (*EndpointAddress, bool)`

GetMediaEndpointAddressOk returns a tuple with the MediaEndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEndpointAddress

`func (o *ConsumptionReportingUnit) SetMediaEndpointAddress(v EndpointAddress)`

SetMediaEndpointAddress sets MediaEndpointAddress field to given value.

### HasMediaEndpointAddress

`func (o *ConsumptionReportingUnit) HasMediaEndpointAddress() bool`

HasMediaEndpointAddress returns a boolean if a field has been set.

### GetStartTime

`func (o *ConsumptionReportingUnit) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ConsumptionReportingUnit) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ConsumptionReportingUnit) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetDuration

`func (o *ConsumptionReportingUnit) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ConsumptionReportingUnit) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ConsumptionReportingUnit) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetLocations

`func (o *ConsumptionReportingUnit) GetLocations() []TypedLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ConsumptionReportingUnit) GetLocationsOk() (*[]TypedLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ConsumptionReportingUnit) SetLocations(v []TypedLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ConsumptionReportingUnit) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


