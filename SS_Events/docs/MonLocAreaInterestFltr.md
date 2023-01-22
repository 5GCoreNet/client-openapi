# MonLocAreaInterestFltr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocInfoCri** | [**LocationInfoCriteria**](LocationInfoCriteria.md) |  | 
**TrigEvnts** | Pointer to [**[]MonLocTriggerEvent**](MonLocTriggerEvent.md) | Triggering events when to send information. | [optional] 

## Methods

### NewMonLocAreaInterestFltr

`func NewMonLocAreaInterestFltr(locInfoCri LocationInfoCriteria, ) *MonLocAreaInterestFltr`

NewMonLocAreaInterestFltr instantiates a new MonLocAreaInterestFltr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonLocAreaInterestFltrWithDefaults

`func NewMonLocAreaInterestFltrWithDefaults() *MonLocAreaInterestFltr`

NewMonLocAreaInterestFltrWithDefaults instantiates a new MonLocAreaInterestFltr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocInfoCri

`func (o *MonLocAreaInterestFltr) GetLocInfoCri() LocationInfoCriteria`

GetLocInfoCri returns the LocInfoCri field if non-nil, zero value otherwise.

### GetLocInfoCriOk

`func (o *MonLocAreaInterestFltr) GetLocInfoCriOk() (*LocationInfoCriteria, bool)`

GetLocInfoCriOk returns a tuple with the LocInfoCri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfoCri

`func (o *MonLocAreaInterestFltr) SetLocInfoCri(v LocationInfoCriteria)`

SetLocInfoCri sets LocInfoCri field to given value.


### GetTrigEvnts

`func (o *MonLocAreaInterestFltr) GetTrigEvnts() []MonLocTriggerEvent`

GetTrigEvnts returns the TrigEvnts field if non-nil, zero value otherwise.

### GetTrigEvntsOk

`func (o *MonLocAreaInterestFltr) GetTrigEvntsOk() (*[]MonLocTriggerEvent, bool)`

GetTrigEvntsOk returns a tuple with the TrigEvnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigEvnts

`func (o *MonLocAreaInterestFltr) SetTrigEvnts(v []MonLocTriggerEvent)`

SetTrigEvnts sets TrigEvnts field to given value.

### HasTrigEvnts

`func (o *MonLocAreaInterestFltr) HasTrigEvnts() bool`

HasTrigEvnts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


