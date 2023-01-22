# NearbyUeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NearbyUeId** | **string** | Represents the identifier of the V2X UE. | 
**Location** | [**UserLocation**](UserLocation.md) |  | 
**Distance** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 

## Methods

### NewNearbyUeInfo

`func NewNearbyUeInfo(nearbyUeId string, location UserLocation, distance int32, ) *NearbyUeInfo`

NewNearbyUeInfo instantiates a new NearbyUeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNearbyUeInfoWithDefaults

`func NewNearbyUeInfoWithDefaults() *NearbyUeInfo`

NewNearbyUeInfoWithDefaults instantiates a new NearbyUeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNearbyUeId

`func (o *NearbyUeInfo) GetNearbyUeId() string`

GetNearbyUeId returns the NearbyUeId field if non-nil, zero value otherwise.

### GetNearbyUeIdOk

`func (o *NearbyUeInfo) GetNearbyUeIdOk() (*string, bool)`

GetNearbyUeIdOk returns a tuple with the NearbyUeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearbyUeId

`func (o *NearbyUeInfo) SetNearbyUeId(v string)`

SetNearbyUeId sets NearbyUeId field to given value.


### GetLocation

`func (o *NearbyUeInfo) GetLocation() UserLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NearbyUeInfo) GetLocationOk() (*UserLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NearbyUeInfo) SetLocation(v UserLocation)`

SetLocation sets Location field to given value.


### GetDistance

`func (o *NearbyUeInfo) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *NearbyUeInfo) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *NearbyUeInfo) SetDistance(v int32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


