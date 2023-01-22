# ReportingCtrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileReportingPeriod** | Pointer to **int32** |  | [optional] 
**FileLocation** | Pointer to **string** |  | [optional] 
**StreamTarget** | Pointer to **string** |  | [optional] 

## Methods

### NewReportingCtrl

`func NewReportingCtrl() *ReportingCtrl`

NewReportingCtrl instantiates a new ReportingCtrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingCtrlWithDefaults

`func NewReportingCtrlWithDefaults() *ReportingCtrl`

NewReportingCtrlWithDefaults instantiates a new ReportingCtrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileReportingPeriod

`func (o *ReportingCtrl) GetFileReportingPeriod() int32`

GetFileReportingPeriod returns the FileReportingPeriod field if non-nil, zero value otherwise.

### GetFileReportingPeriodOk

`func (o *ReportingCtrl) GetFileReportingPeriodOk() (*int32, bool)`

GetFileReportingPeriodOk returns a tuple with the FileReportingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileReportingPeriod

`func (o *ReportingCtrl) SetFileReportingPeriod(v int32)`

SetFileReportingPeriod sets FileReportingPeriod field to given value.

### HasFileReportingPeriod

`func (o *ReportingCtrl) HasFileReportingPeriod() bool`

HasFileReportingPeriod returns a boolean if a field has been set.

### GetFileLocation

`func (o *ReportingCtrl) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *ReportingCtrl) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *ReportingCtrl) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *ReportingCtrl) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetStreamTarget

`func (o *ReportingCtrl) GetStreamTarget() string`

GetStreamTarget returns the StreamTarget field if non-nil, zero value otherwise.

### GetStreamTargetOk

`func (o *ReportingCtrl) GetStreamTargetOk() (*string, bool)`

GetStreamTargetOk returns a tuple with the StreamTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamTarget

`func (o *ReportingCtrl) SetStreamTarget(v string)`

SetStreamTarget sets StreamTarget field to given value.

### HasStreamTarget

`func (o *ReportingCtrl) HasStreamTarget() bool`

HasStreamTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


