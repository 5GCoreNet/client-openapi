# ReachableUeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeList** | **[]string** |  | 
**UserLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 

## Methods

### NewReachableUeInfo

`func NewReachableUeInfo(ueList []string, ) *ReachableUeInfo`

NewReachableUeInfo instantiates a new ReachableUeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachableUeInfoWithDefaults

`func NewReachableUeInfoWithDefaults() *ReachableUeInfo`

NewReachableUeInfoWithDefaults instantiates a new ReachableUeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeList

`func (o *ReachableUeInfo) GetUeList() []string`

GetUeList returns the UeList field if non-nil, zero value otherwise.

### GetUeListOk

`func (o *ReachableUeInfo) GetUeListOk() (*[]string, bool)`

GetUeListOk returns a tuple with the UeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeList

`func (o *ReachableUeInfo) SetUeList(v []string)`

SetUeList sets UeList field to given value.


### GetUserLocation

`func (o *ReachableUeInfo) GetUserLocation() UserLocation`

GetUserLocation returns the UserLocation field if non-nil, zero value otherwise.

### GetUserLocationOk

`func (o *ReachableUeInfo) GetUserLocationOk() (*UserLocation, bool)`

GetUserLocationOk returns a tuple with the UserLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocation

`func (o *ReachableUeInfo) SetUserLocation(v UserLocation)`

SetUserLocation sets UserLocation field to given value.

### HasUserLocation

`func (o *ReachableUeInfo) HasUserLocation() bool`

HasUserLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


