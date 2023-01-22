# MonitorEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnEvnts** | Pointer to [**[]MonitoringType**](MonitoringType.md) | List of monitoring events related to VAL UE. | [optional] 
**AnlEvnts** | Pointer to [**[]AnalyticsEvent**](AnalyticsEvent.md) | List of analytics events related to VAL UE. | [optional] 

## Methods

### NewMonitorEvents

`func NewMonitorEvents() *MonitorEvents`

NewMonitorEvents instantiates a new MonitorEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorEventsWithDefaults

`func NewMonitorEventsWithDefaults() *MonitorEvents`

NewMonitorEventsWithDefaults instantiates a new MonitorEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnEvnts

`func (o *MonitorEvents) GetCnEvnts() []MonitoringType`

GetCnEvnts returns the CnEvnts field if non-nil, zero value otherwise.

### GetCnEvntsOk

`func (o *MonitorEvents) GetCnEvntsOk() (*[]MonitoringType, bool)`

GetCnEvntsOk returns a tuple with the CnEvnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnEvnts

`func (o *MonitorEvents) SetCnEvnts(v []MonitoringType)`

SetCnEvnts sets CnEvnts field to given value.

### HasCnEvnts

`func (o *MonitorEvents) HasCnEvnts() bool`

HasCnEvnts returns a boolean if a field has been set.

### GetAnlEvnts

`func (o *MonitorEvents) GetAnlEvnts() []AnalyticsEvent`

GetAnlEvnts returns the AnlEvnts field if non-nil, zero value otherwise.

### GetAnlEvntsOk

`func (o *MonitorEvents) GetAnlEvntsOk() (*[]AnalyticsEvent, bool)`

GetAnlEvntsOk returns a tuple with the AnlEvnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnlEvnts

`func (o *MonitorEvents) SetAnlEvnts(v []AnalyticsEvent)`

SetAnlEvnts sets AnlEvnts field to given value.

### HasAnlEvnts

`func (o *MonitorEvents) HasAnlEvnts() bool`

HasAnlEvnts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


