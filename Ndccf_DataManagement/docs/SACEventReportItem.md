# SACEventReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**SACEventType**](SACEventType.md) |  | 
**EventState** | [**SACEventState**](SACEventState.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**EventFilter** | [**Snssai**](Snssai.md) |  | 
**SliceStautsInfo** | Pointer to [**SACEventStatus**](SACEventStatus.md) |  | [optional] 

## Methods

### NewSACEventReportItem

`func NewSACEventReportItem(eventType SACEventType, eventState SACEventState, timeStamp time.Time, eventFilter Snssai, ) *SACEventReportItem`

NewSACEventReportItem instantiates a new SACEventReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSACEventReportItemWithDefaults

`func NewSACEventReportItemWithDefaults() *SACEventReportItem`

NewSACEventReportItemWithDefaults instantiates a new SACEventReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *SACEventReportItem) GetEventType() SACEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SACEventReportItem) GetEventTypeOk() (*SACEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SACEventReportItem) SetEventType(v SACEventType)`

SetEventType sets EventType field to given value.


### GetEventState

`func (o *SACEventReportItem) GetEventState() SACEventState`

GetEventState returns the EventState field if non-nil, zero value otherwise.

### GetEventStateOk

`func (o *SACEventReportItem) GetEventStateOk() (*SACEventState, bool)`

GetEventStateOk returns a tuple with the EventState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventState

`func (o *SACEventReportItem) SetEventState(v SACEventState)`

SetEventState sets EventState field to given value.


### GetTimeStamp

`func (o *SACEventReportItem) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *SACEventReportItem) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *SACEventReportItem) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetEventFilter

`func (o *SACEventReportItem) GetEventFilter() Snssai`

GetEventFilter returns the EventFilter field if non-nil, zero value otherwise.

### GetEventFilterOk

`func (o *SACEventReportItem) GetEventFilterOk() (*Snssai, bool)`

GetEventFilterOk returns a tuple with the EventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilter

`func (o *SACEventReportItem) SetEventFilter(v Snssai)`

SetEventFilter sets EventFilter field to given value.


### GetSliceStautsInfo

`func (o *SACEventReportItem) GetSliceStautsInfo() SACEventStatus`

GetSliceStautsInfo returns the SliceStautsInfo field if non-nil, zero value otherwise.

### GetSliceStautsInfoOk

`func (o *SACEventReportItem) GetSliceStautsInfoOk() (*SACEventStatus, bool)`

GetSliceStautsInfoOk returns a tuple with the SliceStautsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceStautsInfo

`func (o *SACEventReportItem) SetSliceStautsInfo(v SACEventStatus)`

SetSliceStautsInfo sets SliceStautsInfo field to given value.

### HasSliceStautsInfo

`func (o *SACEventReportItem) HasSliceStautsInfo() bool`

HasSliceStautsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


