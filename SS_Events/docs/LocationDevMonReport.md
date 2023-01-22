# LocationDevMonReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TgtUes** | [**[]ValTargetUe**](ValTargetUe.md) | List of VAL Users or UE IDs for which report is related to. | 
**LocInfo** | [**LocationInfo**](LocationInfo.md) |  | 
**NotifType** | [**LocDevNotification**](LocDevNotification.md) |  | 

## Methods

### NewLocationDevMonReport

`func NewLocationDevMonReport(tgtUes []ValTargetUe, locInfo LocationInfo, notifType LocDevNotification, ) *LocationDevMonReport`

NewLocationDevMonReport instantiates a new LocationDevMonReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationDevMonReportWithDefaults

`func NewLocationDevMonReportWithDefaults() *LocationDevMonReport`

NewLocationDevMonReportWithDefaults instantiates a new LocationDevMonReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgtUes

`func (o *LocationDevMonReport) GetTgtUes() []ValTargetUe`

GetTgtUes returns the TgtUes field if non-nil, zero value otherwise.

### GetTgtUesOk

`func (o *LocationDevMonReport) GetTgtUesOk() (*[]ValTargetUe, bool)`

GetTgtUesOk returns a tuple with the TgtUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUes

`func (o *LocationDevMonReport) SetTgtUes(v []ValTargetUe)`

SetTgtUes sets TgtUes field to given value.


### GetLocInfo

`func (o *LocationDevMonReport) GetLocInfo() LocationInfo`

GetLocInfo returns the LocInfo field if non-nil, zero value otherwise.

### GetLocInfoOk

`func (o *LocationDevMonReport) GetLocInfoOk() (*LocationInfo, bool)`

GetLocInfoOk returns a tuple with the LocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfo

`func (o *LocationDevMonReport) SetLocInfo(v LocationInfo)`

SetLocInfo sets LocInfo field to given value.


### GetNotifType

`func (o *LocationDevMonReport) GetNotifType() LocDevNotification`

GetNotifType returns the NotifType field if non-nil, zero value otherwise.

### GetNotifTypeOk

`func (o *LocationDevMonReport) GetNotifTypeOk() (*LocDevNotification, bool)`

GetNotifTypeOk returns a tuple with the NotifType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifType

`func (o *LocationDevMonReport) SetNotifType(v LocDevNotification)`

SetNotifType sets NotifType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


