# AnalyticsFailureEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AnalyticsEvent**](AnalyticsEvent.md) |  | 
**FailureCode** | [**AnalyticsFailureCode**](AnalyticsFailureCode.md) |  | 

## Methods

### NewAnalyticsFailureEventInfo

`func NewAnalyticsFailureEventInfo(event AnalyticsEvent, failureCode AnalyticsFailureCode, ) *AnalyticsFailureEventInfo`

NewAnalyticsFailureEventInfo instantiates a new AnalyticsFailureEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsFailureEventInfoWithDefaults

`func NewAnalyticsFailureEventInfoWithDefaults() *AnalyticsFailureEventInfo`

NewAnalyticsFailureEventInfoWithDefaults instantiates a new AnalyticsFailureEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AnalyticsFailureEventInfo) GetEvent() AnalyticsEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AnalyticsFailureEventInfo) GetEventOk() (*AnalyticsEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AnalyticsFailureEventInfo) SetEvent(v AnalyticsEvent)`

SetEvent sets Event field to given value.


### GetFailureCode

`func (o *AnalyticsFailureEventInfo) GetFailureCode() AnalyticsFailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *AnalyticsFailureEventInfo) GetFailureCodeOk() (*AnalyticsFailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *AnalyticsFailureEventInfo) SetFailureCode(v AnalyticsFailureCode)`

SetFailureCode sets FailureCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


