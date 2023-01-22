# DistSessionEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**DistSessionEventType**](DistSessionEventType.md) |  | 
**TimeStamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**MbsSecurityContext** | Pointer to [**MbsSecurityContext**](MbsSecurityContext.md) |  | [optional] 

## Methods

### NewDistSessionEventReport

`func NewDistSessionEventReport(eventType DistSessionEventType, ) *DistSessionEventReport`

NewDistSessionEventReport instantiates a new DistSessionEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistSessionEventReportWithDefaults

`func NewDistSessionEventReportWithDefaults() *DistSessionEventReport`

NewDistSessionEventReportWithDefaults instantiates a new DistSessionEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *DistSessionEventReport) GetEventType() DistSessionEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *DistSessionEventReport) GetEventTypeOk() (*DistSessionEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *DistSessionEventReport) SetEventType(v DistSessionEventType)`

SetEventType sets EventType field to given value.


### GetTimeStamp

`func (o *DistSessionEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *DistSessionEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *DistSessionEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *DistSessionEventReport) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetMbsSecurityContext

`func (o *DistSessionEventReport) GetMbsSecurityContext() MbsSecurityContext`

GetMbsSecurityContext returns the MbsSecurityContext field if non-nil, zero value otherwise.

### GetMbsSecurityContextOk

`func (o *DistSessionEventReport) GetMbsSecurityContextOk() (*MbsSecurityContext, bool)`

GetMbsSecurityContextOk returns a tuple with the MbsSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSecurityContext

`func (o *DistSessionEventReport) SetMbsSecurityContext(v MbsSecurityContext)`

SetMbsSecurityContext sets MbsSecurityContext field to given value.

### HasMbsSecurityContext

`func (o *DistSessionEventReport) HasMbsSecurityContext() bool`

HasMbsSecurityContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


