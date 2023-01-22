# MonitorLocationInterestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TgtUes** | [**[]ValTargetUe**](ValTargetUe.md) | List of VAL Users or UE IDs for which location monitoring is requested. | 
**LocInt** | [**LocationInfo**](LocationInfo.md) |  | 
**NotInt** | **int32** | indicating a time in seconds. | 

## Methods

### NewMonitorLocationInterestFilter

`func NewMonitorLocationInterestFilter(tgtUes []ValTargetUe, locInt LocationInfo, notInt int32, ) *MonitorLocationInterestFilter`

NewMonitorLocationInterestFilter instantiates a new MonitorLocationInterestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorLocationInterestFilterWithDefaults

`func NewMonitorLocationInterestFilterWithDefaults() *MonitorLocationInterestFilter`

NewMonitorLocationInterestFilterWithDefaults instantiates a new MonitorLocationInterestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgtUes

`func (o *MonitorLocationInterestFilter) GetTgtUes() []ValTargetUe`

GetTgtUes returns the TgtUes field if non-nil, zero value otherwise.

### GetTgtUesOk

`func (o *MonitorLocationInterestFilter) GetTgtUesOk() (*[]ValTargetUe, bool)`

GetTgtUesOk returns a tuple with the TgtUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUes

`func (o *MonitorLocationInterestFilter) SetTgtUes(v []ValTargetUe)`

SetTgtUes sets TgtUes field to given value.


### GetLocInt

`func (o *MonitorLocationInterestFilter) GetLocInt() LocationInfo`

GetLocInt returns the LocInt field if non-nil, zero value otherwise.

### GetLocIntOk

`func (o *MonitorLocationInterestFilter) GetLocIntOk() (*LocationInfo, bool)`

GetLocIntOk returns a tuple with the LocInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInt

`func (o *MonitorLocationInterestFilter) SetLocInt(v LocationInfo)`

SetLocInt sets LocInt field to given value.


### GetNotInt

`func (o *MonitorLocationInterestFilter) GetNotInt() int32`

GetNotInt returns the NotInt field if non-nil, zero value otherwise.

### GetNotIntOk

`func (o *MonitorLocationInterestFilter) GetNotIntOk() (*int32, bool)`

GetNotIntOk returns a tuple with the NotInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInt

`func (o *MonitorLocationInterestFilter) SetNotInt(v int32)`

SetNotInt sets NotInt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


