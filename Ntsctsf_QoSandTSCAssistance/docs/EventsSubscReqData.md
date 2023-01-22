# EventsSubscReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]TscEvent**](TscEvent.md) |  | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**QosMon** | Pointer to [**QosMonitoringInformation**](QosMonitoringInformation.md) |  | [optional] 
**UsgThres** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**NotifCorreId** | **string** |  | 

## Methods

### NewEventsSubscReqData

`func NewEventsSubscReqData(events []TscEvent, notifUri string, notifCorreId string, ) *EventsSubscReqData`

NewEventsSubscReqData instantiates a new EventsSubscReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSubscReqDataWithDefaults

`func NewEventsSubscReqDataWithDefaults() *EventsSubscReqData`

NewEventsSubscReqDataWithDefaults instantiates a new EventsSubscReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsSubscReqData) GetEvents() []TscEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsSubscReqData) GetEventsOk() (*[]TscEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsSubscReqData) SetEvents(v []TscEvent)`

SetEvents sets Events field to given value.


### GetNotifUri

`func (o *EventsSubscReqData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *EventsSubscReqData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *EventsSubscReqData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetQosMon

`func (o *EventsSubscReqData) GetQosMon() QosMonitoringInformation`

GetQosMon returns the QosMon field if non-nil, zero value otherwise.

### GetQosMonOk

`func (o *EventsSubscReqData) GetQosMonOk() (*QosMonitoringInformation, bool)`

GetQosMonOk returns a tuple with the QosMon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMon

`func (o *EventsSubscReqData) SetQosMon(v QosMonitoringInformation)`

SetQosMon sets QosMon field to given value.

### HasQosMon

`func (o *EventsSubscReqData) HasQosMon() bool`

HasQosMon returns a boolean if a field has been set.

### GetUsgThres

`func (o *EventsSubscReqData) GetUsgThres() UsageThreshold`

GetUsgThres returns the UsgThres field if non-nil, zero value otherwise.

### GetUsgThresOk

`func (o *EventsSubscReqData) GetUsgThresOk() (*UsageThreshold, bool)`

GetUsgThresOk returns a tuple with the UsgThres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsgThres

`func (o *EventsSubscReqData) SetUsgThres(v UsageThreshold)`

SetUsgThres sets UsgThres field to given value.

### HasUsgThres

`func (o *EventsSubscReqData) HasUsgThres() bool`

HasUsgThres returns a boolean if a field has been set.

### GetNotifCorreId

`func (o *EventsSubscReqData) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *EventsSubscReqData) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *EventsSubscReqData) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


