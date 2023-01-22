# ContextStatusNotifyReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportList** | [**[]ContextStatusEventReport**](ContextStatusEventReport.md) |  | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 

## Methods

### NewContextStatusNotifyReqData

`func NewContextStatusNotifyReqData(reportList []ContextStatusEventReport, ) *ContextStatusNotifyReqData`

NewContextStatusNotifyReqData instantiates a new ContextStatusNotifyReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextStatusNotifyReqDataWithDefaults

`func NewContextStatusNotifyReqDataWithDefaults() *ContextStatusNotifyReqData`

NewContextStatusNotifyReqDataWithDefaults instantiates a new ContextStatusNotifyReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportList

`func (o *ContextStatusNotifyReqData) GetReportList() []ContextStatusEventReport`

GetReportList returns the ReportList field if non-nil, zero value otherwise.

### GetReportListOk

`func (o *ContextStatusNotifyReqData) GetReportListOk() (*[]ContextStatusEventReport, bool)`

GetReportListOk returns a tuple with the ReportList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportList

`func (o *ContextStatusNotifyReqData) SetReportList(v []ContextStatusEventReport)`

SetReportList sets ReportList field to given value.


### GetNotifyCorrelationId

`func (o *ContextStatusNotifyReqData) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *ContextStatusNotifyReqData) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *ContextStatusNotifyReqData) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *ContextStatusNotifyReqData) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


