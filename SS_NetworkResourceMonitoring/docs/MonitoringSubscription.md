# MonitoringSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValUeIds** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of VAL UEs whose QoS monitoring data is requested. | [optional] 
**ValGroupId** | Pointer to **string** | The VAL Group Id which QoS monitoring data is requested. | [optional] 
**ValStreamIds** | Pointer to **[]string** | List of VAL streams for which QoS monitoring data is requested. | [optional] 
**MeasReqs** | Pointer to [**MeasurementRequirements**](MeasurementRequirements.md) |  | [optional] 
**MonRep** | Pointer to [**MonitoringReport**](MonitoringReport.md) |  | [optional] 
**ReportReqs** | Pointer to [**ReportingRequirements**](ReportingRequirements.md) |  | [optional] 
**NotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ReqTestNotif** | Pointer to **bool** |  | [optional] 
**WsNotifCfg** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMonitoringSubscription

`func NewMonitoringSubscription() *MonitoringSubscription`

NewMonitoringSubscription instantiates a new MonitoringSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSubscriptionWithDefaults

`func NewMonitoringSubscriptionWithDefaults() *MonitoringSubscription`

NewMonitoringSubscriptionWithDefaults instantiates a new MonitoringSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValUeIds

`func (o *MonitoringSubscription) GetValUeIds() []ValTargetUe`

GetValUeIds returns the ValUeIds field if non-nil, zero value otherwise.

### GetValUeIdsOk

`func (o *MonitoringSubscription) GetValUeIdsOk() (*[]ValTargetUe, bool)`

GetValUeIdsOk returns a tuple with the ValUeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValUeIds

`func (o *MonitoringSubscription) SetValUeIds(v []ValTargetUe)`

SetValUeIds sets ValUeIds field to given value.

### HasValUeIds

`func (o *MonitoringSubscription) HasValUeIds() bool`

HasValUeIds returns a boolean if a field has been set.

### GetValGroupId

`func (o *MonitoringSubscription) GetValGroupId() string`

GetValGroupId returns the ValGroupId field if non-nil, zero value otherwise.

### GetValGroupIdOk

`func (o *MonitoringSubscription) GetValGroupIdOk() (*string, bool)`

GetValGroupIdOk returns a tuple with the ValGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGroupId

`func (o *MonitoringSubscription) SetValGroupId(v string)`

SetValGroupId sets ValGroupId field to given value.

### HasValGroupId

`func (o *MonitoringSubscription) HasValGroupId() bool`

HasValGroupId returns a boolean if a field has been set.

### GetValStreamIds

`func (o *MonitoringSubscription) GetValStreamIds() []string`

GetValStreamIds returns the ValStreamIds field if non-nil, zero value otherwise.

### GetValStreamIdsOk

`func (o *MonitoringSubscription) GetValStreamIdsOk() (*[]string, bool)`

GetValStreamIdsOk returns a tuple with the ValStreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValStreamIds

`func (o *MonitoringSubscription) SetValStreamIds(v []string)`

SetValStreamIds sets ValStreamIds field to given value.

### HasValStreamIds

`func (o *MonitoringSubscription) HasValStreamIds() bool`

HasValStreamIds returns a boolean if a field has been set.

### GetMeasReqs

`func (o *MonitoringSubscription) GetMeasReqs() MeasurementRequirements`

GetMeasReqs returns the MeasReqs field if non-nil, zero value otherwise.

### GetMeasReqsOk

`func (o *MonitoringSubscription) GetMeasReqsOk() (*MeasurementRequirements, bool)`

GetMeasReqsOk returns a tuple with the MeasReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasReqs

`func (o *MonitoringSubscription) SetMeasReqs(v MeasurementRequirements)`

SetMeasReqs sets MeasReqs field to given value.

### HasMeasReqs

`func (o *MonitoringSubscription) HasMeasReqs() bool`

HasMeasReqs returns a boolean if a field has been set.

### GetMonRep

`func (o *MonitoringSubscription) GetMonRep() MonitoringReport`

GetMonRep returns the MonRep field if non-nil, zero value otherwise.

### GetMonRepOk

`func (o *MonitoringSubscription) GetMonRepOk() (*MonitoringReport, bool)`

GetMonRepOk returns a tuple with the MonRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonRep

`func (o *MonitoringSubscription) SetMonRep(v MonitoringReport)`

SetMonRep sets MonRep field to given value.

### HasMonRep

`func (o *MonitoringSubscription) HasMonRep() bool`

HasMonRep returns a boolean if a field has been set.

### GetReportReqs

`func (o *MonitoringSubscription) GetReportReqs() ReportingRequirements`

GetReportReqs returns the ReportReqs field if non-nil, zero value otherwise.

### GetReportReqsOk

`func (o *MonitoringSubscription) GetReportReqsOk() (*ReportingRequirements, bool)`

GetReportReqsOk returns a tuple with the ReportReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportReqs

`func (o *MonitoringSubscription) SetReportReqs(v ReportingRequirements)`

SetReportReqs sets ReportReqs field to given value.

### HasReportReqs

`func (o *MonitoringSubscription) HasReportReqs() bool`

HasReportReqs returns a boolean if a field has been set.

### GetNotifUri

`func (o *MonitoringSubscription) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *MonitoringSubscription) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *MonitoringSubscription) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *MonitoringSubscription) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### GetReqTestNotif

`func (o *MonitoringSubscription) GetReqTestNotif() bool`

GetReqTestNotif returns the ReqTestNotif field if non-nil, zero value otherwise.

### GetReqTestNotifOk

`func (o *MonitoringSubscription) GetReqTestNotifOk() (*bool, bool)`

GetReqTestNotifOk returns a tuple with the ReqTestNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTestNotif

`func (o *MonitoringSubscription) SetReqTestNotif(v bool)`

SetReqTestNotif sets ReqTestNotif field to given value.

### HasReqTestNotif

`func (o *MonitoringSubscription) HasReqTestNotif() bool`

HasReqTestNotif returns a boolean if a field has been set.

### GetWsNotifCfg

`func (o *MonitoringSubscription) GetWsNotifCfg() WebsockNotifConfig`

GetWsNotifCfg returns the WsNotifCfg field if non-nil, zero value otherwise.

### GetWsNotifCfgOk

`func (o *MonitoringSubscription) GetWsNotifCfgOk() (*WebsockNotifConfig, bool)`

GetWsNotifCfgOk returns a tuple with the WsNotifCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsNotifCfg

`func (o *MonitoringSubscription) SetWsNotifCfg(v WebsockNotifConfig)`

SetWsNotifCfg sets WsNotifCfg field to given value.

### HasWsNotifCfg

`func (o *MonitoringSubscription) HasWsNotifCfg() bool`

HasWsNotifCfg returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MonitoringSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MonitoringSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MonitoringSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MonitoringSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


