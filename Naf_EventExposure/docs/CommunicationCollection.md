# CommunicationCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**EndTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**UlVol** | **int64** | Unsigned integer identifying a volume in units of bytes. | 
**DlVol** | **int64** | Unsigned integer identifying a volume in units of bytes. | 

## Methods

### NewCommunicationCollection

`func NewCommunicationCollection(startTime time.Time, endTime time.Time, ulVol int64, dlVol int64, ) *CommunicationCollection`

NewCommunicationCollection instantiates a new CommunicationCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationCollectionWithDefaults

`func NewCommunicationCollectionWithDefaults() *CommunicationCollection`

NewCommunicationCollectionWithDefaults instantiates a new CommunicationCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *CommunicationCollection) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CommunicationCollection) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CommunicationCollection) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CommunicationCollection) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CommunicationCollection) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CommunicationCollection) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetUlVol

`func (o *CommunicationCollection) GetUlVol() int64`

GetUlVol returns the UlVol field if non-nil, zero value otherwise.

### GetUlVolOk

`func (o *CommunicationCollection) GetUlVolOk() (*int64, bool)`

GetUlVolOk returns a tuple with the UlVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlVol

`func (o *CommunicationCollection) SetUlVol(v int64)`

SetUlVol sets UlVol field to given value.


### GetDlVol

`func (o *CommunicationCollection) GetDlVol() int64`

GetDlVol returns the DlVol field if non-nil, zero value otherwise.

### GetDlVolOk

`func (o *CommunicationCollection) GetDlVolOk() (*int64, bool)`

GetDlVolOk returns a tuple with the DlVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlVol

`func (o *CommunicationCollection) SetDlVol(v int64)`

SetDlVol sets DlVol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


