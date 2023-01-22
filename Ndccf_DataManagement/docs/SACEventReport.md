# SACEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Report** | [**SACEventReportItem**](SACEventReportItem.md) |  | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 

## Methods

### NewSACEventReport

`func NewSACEventReport(report SACEventReportItem, ) *SACEventReport`

NewSACEventReport instantiates a new SACEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSACEventReportWithDefaults

`func NewSACEventReportWithDefaults() *SACEventReport`

NewSACEventReportWithDefaults instantiates a new SACEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReport

`func (o *SACEventReport) GetReport() SACEventReportItem`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *SACEventReport) GetReportOk() (*SACEventReportItem, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *SACEventReport) SetReport(v SACEventReportItem)`

SetReport sets Report field to given value.


### GetNotifyCorrelationId

`func (o *SACEventReport) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *SACEventReport) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *SACEventReport) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *SACEventReport) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


