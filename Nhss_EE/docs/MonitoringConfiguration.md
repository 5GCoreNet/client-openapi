# MonitoringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventType**](EventType.md) |  | 
**ImmediateFlag** | Pointer to **bool** |  | [optional] 
**LocationReportingConfiguration** | Pointer to [**LocationReportingConfiguration**](LocationReportingConfiguration.md) |  | [optional] 
**LossConnectivityConfiguration** | Pointer to [**LossConnectivityConfiguration**](LossConnectivityConfiguration.md) |  | [optional] 
**ReachabilityForDataConfiguration** | Pointer to [**ReachabilityForDataConfiguration**](ReachabilityForDataConfiguration.md) |  | [optional] 
**PduSessionStatusCfg** | Pointer to [**PduSessionStatusCfg**](PduSessionStatusCfg.md) |  | [optional] 
**IdleStatusInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMonitoringConfiguration

`func NewMonitoringConfiguration(eventType EventType, ) *MonitoringConfiguration`

NewMonitoringConfiguration instantiates a new MonitoringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringConfigurationWithDefaults

`func NewMonitoringConfigurationWithDefaults() *MonitoringConfiguration`

NewMonitoringConfigurationWithDefaults instantiates a new MonitoringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *MonitoringConfiguration) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MonitoringConfiguration) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MonitoringConfiguration) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetImmediateFlag

`func (o *MonitoringConfiguration) GetImmediateFlag() bool`

GetImmediateFlag returns the ImmediateFlag field if non-nil, zero value otherwise.

### GetImmediateFlagOk

`func (o *MonitoringConfiguration) GetImmediateFlagOk() (*bool, bool)`

GetImmediateFlagOk returns a tuple with the ImmediateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFlag

`func (o *MonitoringConfiguration) SetImmediateFlag(v bool)`

SetImmediateFlag sets ImmediateFlag field to given value.

### HasImmediateFlag

`func (o *MonitoringConfiguration) HasImmediateFlag() bool`

HasImmediateFlag returns a boolean if a field has been set.

### GetLocationReportingConfiguration

`func (o *MonitoringConfiguration) GetLocationReportingConfiguration() LocationReportingConfiguration`

GetLocationReportingConfiguration returns the LocationReportingConfiguration field if non-nil, zero value otherwise.

### GetLocationReportingConfigurationOk

`func (o *MonitoringConfiguration) GetLocationReportingConfigurationOk() (*LocationReportingConfiguration, bool)`

GetLocationReportingConfigurationOk returns a tuple with the LocationReportingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReportingConfiguration

`func (o *MonitoringConfiguration) SetLocationReportingConfiguration(v LocationReportingConfiguration)`

SetLocationReportingConfiguration sets LocationReportingConfiguration field to given value.

### HasLocationReportingConfiguration

`func (o *MonitoringConfiguration) HasLocationReportingConfiguration() bool`

HasLocationReportingConfiguration returns a boolean if a field has been set.

### GetLossConnectivityConfiguration

`func (o *MonitoringConfiguration) GetLossConnectivityConfiguration() LossConnectivityConfiguration`

GetLossConnectivityConfiguration returns the LossConnectivityConfiguration field if non-nil, zero value otherwise.

### GetLossConnectivityConfigurationOk

`func (o *MonitoringConfiguration) GetLossConnectivityConfigurationOk() (*LossConnectivityConfiguration, bool)`

GetLossConnectivityConfigurationOk returns a tuple with the LossConnectivityConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossConnectivityConfiguration

`func (o *MonitoringConfiguration) SetLossConnectivityConfiguration(v LossConnectivityConfiguration)`

SetLossConnectivityConfiguration sets LossConnectivityConfiguration field to given value.

### HasLossConnectivityConfiguration

`func (o *MonitoringConfiguration) HasLossConnectivityConfiguration() bool`

HasLossConnectivityConfiguration returns a boolean if a field has been set.

### GetReachabilityForDataConfiguration

`func (o *MonitoringConfiguration) GetReachabilityForDataConfiguration() ReachabilityForDataConfiguration`

GetReachabilityForDataConfiguration returns the ReachabilityForDataConfiguration field if non-nil, zero value otherwise.

### GetReachabilityForDataConfigurationOk

`func (o *MonitoringConfiguration) GetReachabilityForDataConfigurationOk() (*ReachabilityForDataConfiguration, bool)`

GetReachabilityForDataConfigurationOk returns a tuple with the ReachabilityForDataConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityForDataConfiguration

`func (o *MonitoringConfiguration) SetReachabilityForDataConfiguration(v ReachabilityForDataConfiguration)`

SetReachabilityForDataConfiguration sets ReachabilityForDataConfiguration field to given value.

### HasReachabilityForDataConfiguration

`func (o *MonitoringConfiguration) HasReachabilityForDataConfiguration() bool`

HasReachabilityForDataConfiguration returns a boolean if a field has been set.

### GetPduSessionStatusCfg

`func (o *MonitoringConfiguration) GetPduSessionStatusCfg() PduSessionStatusCfg`

GetPduSessionStatusCfg returns the PduSessionStatusCfg field if non-nil, zero value otherwise.

### GetPduSessionStatusCfgOk

`func (o *MonitoringConfiguration) GetPduSessionStatusCfgOk() (*PduSessionStatusCfg, bool)`

GetPduSessionStatusCfgOk returns a tuple with the PduSessionStatusCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionStatusCfg

`func (o *MonitoringConfiguration) SetPduSessionStatusCfg(v PduSessionStatusCfg)`

SetPduSessionStatusCfg sets PduSessionStatusCfg field to given value.

### HasPduSessionStatusCfg

`func (o *MonitoringConfiguration) HasPduSessionStatusCfg() bool`

HasPduSessionStatusCfg returns a boolean if a field has been set.

### GetIdleStatusInd

`func (o *MonitoringConfiguration) GetIdleStatusInd() bool`

GetIdleStatusInd returns the IdleStatusInd field if non-nil, zero value otherwise.

### GetIdleStatusIndOk

`func (o *MonitoringConfiguration) GetIdleStatusIndOk() (*bool, bool)`

GetIdleStatusIndOk returns a tuple with the IdleStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusInd

`func (o *MonitoringConfiguration) SetIdleStatusInd(v bool)`

SetIdleStatusInd sets IdleStatusInd field to given value.

### HasIdleStatusInd

`func (o *MonitoringConfiguration) HasIdleStatusInd() bool`

HasIdleStatusInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


