# DistSessionEventReportList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventReportList** | [**[]DistSessionEventReport**](DistSessionEventReport.md) |  | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 

## Methods

### NewDistSessionEventReportList

`func NewDistSessionEventReportList(eventReportList []DistSessionEventReport, ) *DistSessionEventReportList`

NewDistSessionEventReportList instantiates a new DistSessionEventReportList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistSessionEventReportListWithDefaults

`func NewDistSessionEventReportListWithDefaults() *DistSessionEventReportList`

NewDistSessionEventReportListWithDefaults instantiates a new DistSessionEventReportList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventReportList

`func (o *DistSessionEventReportList) GetEventReportList() []DistSessionEventReport`

GetEventReportList returns the EventReportList field if non-nil, zero value otherwise.

### GetEventReportListOk

`func (o *DistSessionEventReportList) GetEventReportListOk() (*[]DistSessionEventReport, bool)`

GetEventReportListOk returns a tuple with the EventReportList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReportList

`func (o *DistSessionEventReportList) SetEventReportList(v []DistSessionEventReport)`

SetEventReportList sets EventReportList field to given value.


### GetNotifyCorrelationId

`func (o *DistSessionEventReportList) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *DistSessionEventReportList) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *DistSessionEventReportList) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *DistSessionEventReportList) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


