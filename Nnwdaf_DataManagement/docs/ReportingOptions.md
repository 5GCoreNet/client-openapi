# ReportingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**NotifyPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NotifyPeriodInc** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**DepEventSubId** | Pointer to **string** | Notifications for the present subscription are sent only upon occurrence of events of the subscription with identifier that matches this attribute.  | [optional] 
**MinClubbedNotif** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MaxClubbedNotif** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewReportingOptions

`func NewReportingOptions() *ReportingOptions`

NewReportingOptions instantiates a new ReportingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingOptionsWithDefaults

`func NewReportingOptionsWithDefaults() *ReportingOptions`

NewReportingOptionsWithDefaults instantiates a new ReportingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyWindow

`func (o *ReportingOptions) GetNotifyWindow() TimeWindow`

GetNotifyWindow returns the NotifyWindow field if non-nil, zero value otherwise.

### GetNotifyWindowOk

`func (o *ReportingOptions) GetNotifyWindowOk() (*TimeWindow, bool)`

GetNotifyWindowOk returns a tuple with the NotifyWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWindow

`func (o *ReportingOptions) SetNotifyWindow(v TimeWindow)`

SetNotifyWindow sets NotifyWindow field to given value.

### HasNotifyWindow

`func (o *ReportingOptions) HasNotifyWindow() bool`

HasNotifyWindow returns a boolean if a field has been set.

### GetNotifyPeriod

`func (o *ReportingOptions) GetNotifyPeriod() int32`

GetNotifyPeriod returns the NotifyPeriod field if non-nil, zero value otherwise.

### GetNotifyPeriodOk

`func (o *ReportingOptions) GetNotifyPeriodOk() (*int32, bool)`

GetNotifyPeriodOk returns a tuple with the NotifyPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPeriod

`func (o *ReportingOptions) SetNotifyPeriod(v int32)`

SetNotifyPeriod sets NotifyPeriod field to given value.

### HasNotifyPeriod

`func (o *ReportingOptions) HasNotifyPeriod() bool`

HasNotifyPeriod returns a boolean if a field has been set.

### GetNotifyPeriodInc

`func (o *ReportingOptions) GetNotifyPeriodInc() int32`

GetNotifyPeriodInc returns the NotifyPeriodInc field if non-nil, zero value otherwise.

### GetNotifyPeriodIncOk

`func (o *ReportingOptions) GetNotifyPeriodIncOk() (*int32, bool)`

GetNotifyPeriodIncOk returns a tuple with the NotifyPeriodInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPeriodInc

`func (o *ReportingOptions) SetNotifyPeriodInc(v int32)`

SetNotifyPeriodInc sets NotifyPeriodInc field to given value.

### HasNotifyPeriodInc

`func (o *ReportingOptions) HasNotifyPeriodInc() bool`

HasNotifyPeriodInc returns a boolean if a field has been set.

### GetDepEventSubId

`func (o *ReportingOptions) GetDepEventSubId() string`

GetDepEventSubId returns the DepEventSubId field if non-nil, zero value otherwise.

### GetDepEventSubIdOk

`func (o *ReportingOptions) GetDepEventSubIdOk() (*string, bool)`

GetDepEventSubIdOk returns a tuple with the DepEventSubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepEventSubId

`func (o *ReportingOptions) SetDepEventSubId(v string)`

SetDepEventSubId sets DepEventSubId field to given value.

### HasDepEventSubId

`func (o *ReportingOptions) HasDepEventSubId() bool`

HasDepEventSubId returns a boolean if a field has been set.

### GetMinClubbedNotif

`func (o *ReportingOptions) GetMinClubbedNotif() int32`

GetMinClubbedNotif returns the MinClubbedNotif field if non-nil, zero value otherwise.

### GetMinClubbedNotifOk

`func (o *ReportingOptions) GetMinClubbedNotifOk() (*int32, bool)`

GetMinClubbedNotifOk returns a tuple with the MinClubbedNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinClubbedNotif

`func (o *ReportingOptions) SetMinClubbedNotif(v int32)`

SetMinClubbedNotif sets MinClubbedNotif field to given value.

### HasMinClubbedNotif

`func (o *ReportingOptions) HasMinClubbedNotif() bool`

HasMinClubbedNotif returns a boolean if a field has been set.

### GetMaxClubbedNotif

`func (o *ReportingOptions) GetMaxClubbedNotif() int32`

GetMaxClubbedNotif returns the MaxClubbedNotif field if non-nil, zero value otherwise.

### GetMaxClubbedNotifOk

`func (o *ReportingOptions) GetMaxClubbedNotifOk() (*int32, bool)`

GetMaxClubbedNotifOk returns a tuple with the MaxClubbedNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClubbedNotif

`func (o *ReportingOptions) SetMaxClubbedNotif(v int32)`

SetMaxClubbedNotif sets MaxClubbedNotif field to given value.

### HasMaxClubbedNotif

`func (o *ReportingOptions) HasMaxClubbedNotif() bool`

HasMaxClubbedNotif returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


