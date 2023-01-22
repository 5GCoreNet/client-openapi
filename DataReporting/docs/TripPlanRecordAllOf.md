# TripPlanRecordAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartingPoint** | [**LocationData**](LocationData.md) |  | 
**Waypoints** | Pointer to [**[]LocationData**](LocationData.md) |  | [optional] 
**Destination** | [**LocationData**](LocationData.md) |  | 
**EstimatedAverageSpeed** | Pointer to **float32** | Indicates value of horizontal speed. | [optional] 
**EstimatedArrivalTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewTripPlanRecordAllOf

`func NewTripPlanRecordAllOf(startingPoint LocationData, destination LocationData, ) *TripPlanRecordAllOf`

NewTripPlanRecordAllOf instantiates a new TripPlanRecordAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTripPlanRecordAllOfWithDefaults

`func NewTripPlanRecordAllOfWithDefaults() *TripPlanRecordAllOf`

NewTripPlanRecordAllOfWithDefaults instantiates a new TripPlanRecordAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartingPoint

`func (o *TripPlanRecordAllOf) GetStartingPoint() LocationData`

GetStartingPoint returns the StartingPoint field if non-nil, zero value otherwise.

### GetStartingPointOk

`func (o *TripPlanRecordAllOf) GetStartingPointOk() (*LocationData, bool)`

GetStartingPointOk returns a tuple with the StartingPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingPoint

`func (o *TripPlanRecordAllOf) SetStartingPoint(v LocationData)`

SetStartingPoint sets StartingPoint field to given value.


### GetWaypoints

`func (o *TripPlanRecordAllOf) GetWaypoints() []LocationData`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *TripPlanRecordAllOf) GetWaypointsOk() (*[]LocationData, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *TripPlanRecordAllOf) SetWaypoints(v []LocationData)`

SetWaypoints sets Waypoints field to given value.

### HasWaypoints

`func (o *TripPlanRecordAllOf) HasWaypoints() bool`

HasWaypoints returns a boolean if a field has been set.

### GetDestination

`func (o *TripPlanRecordAllOf) GetDestination() LocationData`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TripPlanRecordAllOf) GetDestinationOk() (*LocationData, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TripPlanRecordAllOf) SetDestination(v LocationData)`

SetDestination sets Destination field to given value.


### GetEstimatedAverageSpeed

`func (o *TripPlanRecordAllOf) GetEstimatedAverageSpeed() float32`

GetEstimatedAverageSpeed returns the EstimatedAverageSpeed field if non-nil, zero value otherwise.

### GetEstimatedAverageSpeedOk

`func (o *TripPlanRecordAllOf) GetEstimatedAverageSpeedOk() (*float32, bool)`

GetEstimatedAverageSpeedOk returns a tuple with the EstimatedAverageSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAverageSpeed

`func (o *TripPlanRecordAllOf) SetEstimatedAverageSpeed(v float32)`

SetEstimatedAverageSpeed sets EstimatedAverageSpeed field to given value.

### HasEstimatedAverageSpeed

`func (o *TripPlanRecordAllOf) HasEstimatedAverageSpeed() bool`

HasEstimatedAverageSpeed returns a boolean if a field has been set.

### GetEstimatedArrivalTime

`func (o *TripPlanRecordAllOf) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *TripPlanRecordAllOf) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *TripPlanRecordAllOf) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *TripPlanRecordAllOf) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


