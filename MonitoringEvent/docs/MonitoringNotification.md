# MonitoringNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**ConfigResults** | Pointer to [**[]ConfigResult**](ConfigResult.md) | Each element identifies a notification of grouping configuration result. | [optional] 
**MonitoringEventReports** | Pointer to [**[]MonitoringEventReport**](MonitoringEventReport.md) | Monitoring event reports. | [optional] 
**AddedExternalIds** | Pointer to **[]string** | Identifies the added external Identifier(s) within the active group via the \&quot;externalGroupId\&quot; attribute within the MonitoringEventSubscription data type. | [optional] 
**AddedMsisdns** | Pointer to **[]string** | Identifies the added MSISDN(s) within the active group via the \&quot;externalGroupId\&quot; attribute within the MonitoringEventSubscription data type. | [optional] 
**CancelExternalIds** | Pointer to **[]string** | Identifies the cancelled external Identifier(s) within the active group via the \&quot;externalGroupId\&quot; attribute within the MonitoringEventSubscription data type. | [optional] 
**CancelMsisdns** | Pointer to **[]string** | Identifies the cancelled MSISDN(s) within the active group via the \&quot;externalGroupId\&quot; attribute within the MonitoringEventSubscription data type. | [optional] 
**CancelInd** | Pointer to **bool** | Indicates whether to request to cancel the corresponding monitoring subscription. Set to false or omitted otherwise.  | [optional] 
**AppliedParam** | Pointer to [**AppliedParameterConfiguration**](AppliedParameterConfiguration.md) |  | [optional] 

## Methods

### NewMonitoringNotification

`func NewMonitoringNotification(subscription string, ) *MonitoringNotification`

NewMonitoringNotification instantiates a new MonitoringNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringNotificationWithDefaults

`func NewMonitoringNotificationWithDefaults() *MonitoringNotification`

NewMonitoringNotificationWithDefaults instantiates a new MonitoringNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *MonitoringNotification) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MonitoringNotification) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MonitoringNotification) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetConfigResults

`func (o *MonitoringNotification) GetConfigResults() []ConfigResult`

GetConfigResults returns the ConfigResults field if non-nil, zero value otherwise.

### GetConfigResultsOk

`func (o *MonitoringNotification) GetConfigResultsOk() (*[]ConfigResult, bool)`

GetConfigResultsOk returns a tuple with the ConfigResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResults

`func (o *MonitoringNotification) SetConfigResults(v []ConfigResult)`

SetConfigResults sets ConfigResults field to given value.

### HasConfigResults

`func (o *MonitoringNotification) HasConfigResults() bool`

HasConfigResults returns a boolean if a field has been set.

### GetMonitoringEventReports

`func (o *MonitoringNotification) GetMonitoringEventReports() []MonitoringEventReport`

GetMonitoringEventReports returns the MonitoringEventReports field if non-nil, zero value otherwise.

### GetMonitoringEventReportsOk

`func (o *MonitoringNotification) GetMonitoringEventReportsOk() (*[]MonitoringEventReport, bool)`

GetMonitoringEventReportsOk returns a tuple with the MonitoringEventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEventReports

`func (o *MonitoringNotification) SetMonitoringEventReports(v []MonitoringEventReport)`

SetMonitoringEventReports sets MonitoringEventReports field to given value.

### HasMonitoringEventReports

`func (o *MonitoringNotification) HasMonitoringEventReports() bool`

HasMonitoringEventReports returns a boolean if a field has been set.

### GetAddedExternalIds

`func (o *MonitoringNotification) GetAddedExternalIds() []string`

GetAddedExternalIds returns the AddedExternalIds field if non-nil, zero value otherwise.

### GetAddedExternalIdsOk

`func (o *MonitoringNotification) GetAddedExternalIdsOk() (*[]string, bool)`

GetAddedExternalIdsOk returns a tuple with the AddedExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedExternalIds

`func (o *MonitoringNotification) SetAddedExternalIds(v []string)`

SetAddedExternalIds sets AddedExternalIds field to given value.

### HasAddedExternalIds

`func (o *MonitoringNotification) HasAddedExternalIds() bool`

HasAddedExternalIds returns a boolean if a field has been set.

### GetAddedMsisdns

`func (o *MonitoringNotification) GetAddedMsisdns() []string`

GetAddedMsisdns returns the AddedMsisdns field if non-nil, zero value otherwise.

### GetAddedMsisdnsOk

`func (o *MonitoringNotification) GetAddedMsisdnsOk() (*[]string, bool)`

GetAddedMsisdnsOk returns a tuple with the AddedMsisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedMsisdns

`func (o *MonitoringNotification) SetAddedMsisdns(v []string)`

SetAddedMsisdns sets AddedMsisdns field to given value.

### HasAddedMsisdns

`func (o *MonitoringNotification) HasAddedMsisdns() bool`

HasAddedMsisdns returns a boolean if a field has been set.

### GetCancelExternalIds

`func (o *MonitoringNotification) GetCancelExternalIds() []string`

GetCancelExternalIds returns the CancelExternalIds field if non-nil, zero value otherwise.

### GetCancelExternalIdsOk

`func (o *MonitoringNotification) GetCancelExternalIdsOk() (*[]string, bool)`

GetCancelExternalIdsOk returns a tuple with the CancelExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelExternalIds

`func (o *MonitoringNotification) SetCancelExternalIds(v []string)`

SetCancelExternalIds sets CancelExternalIds field to given value.

### HasCancelExternalIds

`func (o *MonitoringNotification) HasCancelExternalIds() bool`

HasCancelExternalIds returns a boolean if a field has been set.

### GetCancelMsisdns

`func (o *MonitoringNotification) GetCancelMsisdns() []string`

GetCancelMsisdns returns the CancelMsisdns field if non-nil, zero value otherwise.

### GetCancelMsisdnsOk

`func (o *MonitoringNotification) GetCancelMsisdnsOk() (*[]string, bool)`

GetCancelMsisdnsOk returns a tuple with the CancelMsisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelMsisdns

`func (o *MonitoringNotification) SetCancelMsisdns(v []string)`

SetCancelMsisdns sets CancelMsisdns field to given value.

### HasCancelMsisdns

`func (o *MonitoringNotification) HasCancelMsisdns() bool`

HasCancelMsisdns returns a boolean if a field has been set.

### GetCancelInd

`func (o *MonitoringNotification) GetCancelInd() bool`

GetCancelInd returns the CancelInd field if non-nil, zero value otherwise.

### GetCancelIndOk

`func (o *MonitoringNotification) GetCancelIndOk() (*bool, bool)`

GetCancelIndOk returns a tuple with the CancelInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelInd

`func (o *MonitoringNotification) SetCancelInd(v bool)`

SetCancelInd sets CancelInd field to given value.

### HasCancelInd

`func (o *MonitoringNotification) HasCancelInd() bool`

HasCancelInd returns a boolean if a field has been set.

### GetAppliedParam

`func (o *MonitoringNotification) GetAppliedParam() AppliedParameterConfiguration`

GetAppliedParam returns the AppliedParam field if non-nil, zero value otherwise.

### GetAppliedParamOk

`func (o *MonitoringNotification) GetAppliedParamOk() (*AppliedParameterConfiguration, bool)`

GetAppliedParamOk returns a tuple with the AppliedParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedParam

`func (o *MonitoringNotification) SetAppliedParam(v AppliedParameterConfiguration)`

SetAppliedParam sets AppliedParam field to given value.

### HasAppliedParam

`func (o *MonitoringNotification) HasAppliedParam() bool`

HasAppliedParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


