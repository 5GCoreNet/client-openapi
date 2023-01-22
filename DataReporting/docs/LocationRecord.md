# LocationRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**Location** | [**LocationData**](LocationData.md) |  | 

## Methods

### NewLocationRecord

`func NewLocationRecord(timestamp time.Time, location LocationData, ) *LocationRecord`

NewLocationRecord instantiates a new LocationRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationRecordWithDefaults

`func NewLocationRecordWithDefaults() *LocationRecord`

NewLocationRecordWithDefaults instantiates a new LocationRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *LocationRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LocationRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LocationRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetLocation

`func (o *LocationRecord) GetLocation() LocationData`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationRecord) GetLocationOk() (*LocationData, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationRecord) SetLocation(v LocationData)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


