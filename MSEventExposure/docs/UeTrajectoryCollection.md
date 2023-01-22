# UeTrajectoryCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**LocArea** | [**LocationArea5G**](LocationArea5G.md) |  | 

## Methods

### NewUeTrajectoryCollection

`func NewUeTrajectoryCollection(ts time.Time, locArea LocationArea5G, ) *UeTrajectoryCollection`

NewUeTrajectoryCollection instantiates a new UeTrajectoryCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeTrajectoryCollectionWithDefaults

`func NewUeTrajectoryCollectionWithDefaults() *UeTrajectoryCollection`

NewUeTrajectoryCollectionWithDefaults instantiates a new UeTrajectoryCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *UeTrajectoryCollection) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UeTrajectoryCollection) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UeTrajectoryCollection) SetTs(v time.Time)`

SetTs sets Ts field to given value.


### GetLocArea

`func (o *UeTrajectoryCollection) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *UeTrajectoryCollection) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *UeTrajectoryCollection) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


