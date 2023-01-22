# AmInfluEventNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfTransId** | **string** |  | 
**Event** | [**AmInfluEvent**](AmInfluEvent.md) |  | 
**GeoAreas** | Pointer to [**[]GeographicalArea**](GeographicalArea.md) | Identifies geographic areas of the user where the request is applicable. | [optional] 

## Methods

### NewAmInfluEventNotif

`func NewAmInfluEventNotif(afTransId string, event AmInfluEvent, ) *AmInfluEventNotif`

NewAmInfluEventNotif instantiates a new AmInfluEventNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmInfluEventNotifWithDefaults

`func NewAmInfluEventNotifWithDefaults() *AmInfluEventNotif`

NewAmInfluEventNotifWithDefaults instantiates a new AmInfluEventNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfTransId

`func (o *AmInfluEventNotif) GetAfTransId() string`

GetAfTransId returns the AfTransId field if non-nil, zero value otherwise.

### GetAfTransIdOk

`func (o *AmInfluEventNotif) GetAfTransIdOk() (*string, bool)`

GetAfTransIdOk returns a tuple with the AfTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfTransId

`func (o *AmInfluEventNotif) SetAfTransId(v string)`

SetAfTransId sets AfTransId field to given value.


### GetEvent

`func (o *AmInfluEventNotif) GetEvent() AmInfluEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AmInfluEventNotif) GetEventOk() (*AmInfluEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AmInfluEventNotif) SetEvent(v AmInfluEvent)`

SetEvent sets Event field to given value.


### GetGeoAreas

`func (o *AmInfluEventNotif) GetGeoAreas() []GeographicalArea`

GetGeoAreas returns the GeoAreas field if non-nil, zero value otherwise.

### GetGeoAreasOk

`func (o *AmInfluEventNotif) GetGeoAreasOk() (*[]GeographicalArea, bool)`

GetGeoAreasOk returns a tuple with the GeoAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoAreas

`func (o *AmInfluEventNotif) SetGeoAreas(v []GeographicalArea)`

SetGeoAreas sets GeoAreas field to given value.

### HasGeoAreas

`func (o *AmInfluEventNotif) HasGeoAreas() bool`

HasGeoAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


