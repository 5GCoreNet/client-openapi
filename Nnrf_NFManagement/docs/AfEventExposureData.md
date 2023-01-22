# AfEventExposureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfEvents** | [**[]AfEvent**](AfEvent.md) |  | 
**AfIds** | Pointer to **[]string** |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAfEventExposureData

`func NewAfEventExposureData(afEvents []AfEvent, ) *AfEventExposureData`

NewAfEventExposureData instantiates a new AfEventExposureData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfEventExposureDataWithDefaults

`func NewAfEventExposureDataWithDefaults() *AfEventExposureData`

NewAfEventExposureDataWithDefaults instantiates a new AfEventExposureData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfEvents

`func (o *AfEventExposureData) GetAfEvents() []AfEvent`

GetAfEvents returns the AfEvents field if non-nil, zero value otherwise.

### GetAfEventsOk

`func (o *AfEventExposureData) GetAfEventsOk() (*[]AfEvent, bool)`

GetAfEventsOk returns a tuple with the AfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEvents

`func (o *AfEventExposureData) SetAfEvents(v []AfEvent)`

SetAfEvents sets AfEvents field to given value.


### GetAfIds

`func (o *AfEventExposureData) GetAfIds() []string`

GetAfIds returns the AfIds field if non-nil, zero value otherwise.

### GetAfIdsOk

`func (o *AfEventExposureData) GetAfIdsOk() (*[]string, bool)`

GetAfIdsOk returns a tuple with the AfIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfIds

`func (o *AfEventExposureData) SetAfIds(v []string)`

SetAfIds sets AfIds field to given value.

### HasAfIds

`func (o *AfEventExposureData) HasAfIds() bool`

HasAfIds returns a boolean if a field has been set.

### GetAppIds

`func (o *AfEventExposureData) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AfEventExposureData) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AfEventExposureData) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *AfEventExposureData) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


