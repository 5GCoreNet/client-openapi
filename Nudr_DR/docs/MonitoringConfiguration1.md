# MonitoringConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventType**](EventType.md) |  | 
**ImmediateFlag** | Pointer to **bool** |  | [optional] 
**LocationReportingConfiguration** | Pointer to [**LocationReportingConfiguration1**](LocationReportingConfiguration1.md) |  | [optional] 
**AssociationType** | Pointer to [**AssociationType**](AssociationType.md) |  | [optional] 
**DatalinkReportCfg** | Pointer to [**DatalinkReportingConfiguration1**](DatalinkReportingConfiguration1.md) |  | [optional] 
**LossConnectivityCfg** | Pointer to [**LossConnectivityCfg1**](LossConnectivityCfg1.md) |  | [optional] 
**MaximumLatency** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**MaximumResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SuggestedPacketNumDl** | Pointer to **int32** |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**PduSessionStatusCfg** | Pointer to [**PduSessionStatusCfg1**](PduSessionStatusCfg1.md) |  | [optional] 
**ReachabilityForSmsCfg** | Pointer to [**ReachabilityForSmsConfiguration**](ReachabilityForSmsConfiguration.md) |  | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**AfId** | Pointer to **string** |  | [optional] 
**ReachabilityForDataCfg** | Pointer to [**ReachabilityForDataConfiguration1**](ReachabilityForDataConfiguration1.md) |  | [optional] 
**IdleStatusInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMonitoringConfiguration1

`func NewMonitoringConfiguration1(eventType EventType, ) *MonitoringConfiguration1`

NewMonitoringConfiguration1 instantiates a new MonitoringConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringConfiguration1WithDefaults

`func NewMonitoringConfiguration1WithDefaults() *MonitoringConfiguration1`

NewMonitoringConfiguration1WithDefaults instantiates a new MonitoringConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *MonitoringConfiguration1) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MonitoringConfiguration1) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MonitoringConfiguration1) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetImmediateFlag

`func (o *MonitoringConfiguration1) GetImmediateFlag() bool`

GetImmediateFlag returns the ImmediateFlag field if non-nil, zero value otherwise.

### GetImmediateFlagOk

`func (o *MonitoringConfiguration1) GetImmediateFlagOk() (*bool, bool)`

GetImmediateFlagOk returns a tuple with the ImmediateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFlag

`func (o *MonitoringConfiguration1) SetImmediateFlag(v bool)`

SetImmediateFlag sets ImmediateFlag field to given value.

### HasImmediateFlag

`func (o *MonitoringConfiguration1) HasImmediateFlag() bool`

HasImmediateFlag returns a boolean if a field has been set.

### GetLocationReportingConfiguration

`func (o *MonitoringConfiguration1) GetLocationReportingConfiguration() LocationReportingConfiguration1`

GetLocationReportingConfiguration returns the LocationReportingConfiguration field if non-nil, zero value otherwise.

### GetLocationReportingConfigurationOk

`func (o *MonitoringConfiguration1) GetLocationReportingConfigurationOk() (*LocationReportingConfiguration1, bool)`

GetLocationReportingConfigurationOk returns a tuple with the LocationReportingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReportingConfiguration

`func (o *MonitoringConfiguration1) SetLocationReportingConfiguration(v LocationReportingConfiguration1)`

SetLocationReportingConfiguration sets LocationReportingConfiguration field to given value.

### HasLocationReportingConfiguration

`func (o *MonitoringConfiguration1) HasLocationReportingConfiguration() bool`

HasLocationReportingConfiguration returns a boolean if a field has been set.

### GetAssociationType

`func (o *MonitoringConfiguration1) GetAssociationType() AssociationType`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *MonitoringConfiguration1) GetAssociationTypeOk() (*AssociationType, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *MonitoringConfiguration1) SetAssociationType(v AssociationType)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *MonitoringConfiguration1) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### GetDatalinkReportCfg

`func (o *MonitoringConfiguration1) GetDatalinkReportCfg() DatalinkReportingConfiguration1`

GetDatalinkReportCfg returns the DatalinkReportCfg field if non-nil, zero value otherwise.

### GetDatalinkReportCfgOk

`func (o *MonitoringConfiguration1) GetDatalinkReportCfgOk() (*DatalinkReportingConfiguration1, bool)`

GetDatalinkReportCfgOk returns a tuple with the DatalinkReportCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalinkReportCfg

`func (o *MonitoringConfiguration1) SetDatalinkReportCfg(v DatalinkReportingConfiguration1)`

SetDatalinkReportCfg sets DatalinkReportCfg field to given value.

### HasDatalinkReportCfg

`func (o *MonitoringConfiguration1) HasDatalinkReportCfg() bool`

HasDatalinkReportCfg returns a boolean if a field has been set.

### GetLossConnectivityCfg

`func (o *MonitoringConfiguration1) GetLossConnectivityCfg() LossConnectivityCfg1`

GetLossConnectivityCfg returns the LossConnectivityCfg field if non-nil, zero value otherwise.

### GetLossConnectivityCfgOk

`func (o *MonitoringConfiguration1) GetLossConnectivityCfgOk() (*LossConnectivityCfg1, bool)`

GetLossConnectivityCfgOk returns a tuple with the LossConnectivityCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossConnectivityCfg

`func (o *MonitoringConfiguration1) SetLossConnectivityCfg(v LossConnectivityCfg1)`

SetLossConnectivityCfg sets LossConnectivityCfg field to given value.

### HasLossConnectivityCfg

`func (o *MonitoringConfiguration1) HasLossConnectivityCfg() bool`

HasLossConnectivityCfg returns a boolean if a field has been set.

### GetMaximumLatency

`func (o *MonitoringConfiguration1) GetMaximumLatency() int32`

GetMaximumLatency returns the MaximumLatency field if non-nil, zero value otherwise.

### GetMaximumLatencyOk

`func (o *MonitoringConfiguration1) GetMaximumLatencyOk() (*int32, bool)`

GetMaximumLatencyOk returns a tuple with the MaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLatency

`func (o *MonitoringConfiguration1) SetMaximumLatency(v int32)`

SetMaximumLatency sets MaximumLatency field to given value.

### HasMaximumLatency

`func (o *MonitoringConfiguration1) HasMaximumLatency() bool`

HasMaximumLatency returns a boolean if a field has been set.

### GetMaximumResponseTime

`func (o *MonitoringConfiguration1) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *MonitoringConfiguration1) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *MonitoringConfiguration1) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.

### HasMaximumResponseTime

`func (o *MonitoringConfiguration1) HasMaximumResponseTime() bool`

HasMaximumResponseTime returns a boolean if a field has been set.

### GetSuggestedPacketNumDl

`func (o *MonitoringConfiguration1) GetSuggestedPacketNumDl() int32`

GetSuggestedPacketNumDl returns the SuggestedPacketNumDl field if non-nil, zero value otherwise.

### GetSuggestedPacketNumDlOk

`func (o *MonitoringConfiguration1) GetSuggestedPacketNumDlOk() (*int32, bool)`

GetSuggestedPacketNumDlOk returns a tuple with the SuggestedPacketNumDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPacketNumDl

`func (o *MonitoringConfiguration1) SetSuggestedPacketNumDl(v int32)`

SetSuggestedPacketNumDl sets SuggestedPacketNumDl field to given value.

### HasSuggestedPacketNumDl

`func (o *MonitoringConfiguration1) HasSuggestedPacketNumDl() bool`

HasSuggestedPacketNumDl returns a boolean if a field has been set.

### GetDnn

`func (o *MonitoringConfiguration1) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *MonitoringConfiguration1) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *MonitoringConfiguration1) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *MonitoringConfiguration1) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSingleNssai

`func (o *MonitoringConfiguration1) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *MonitoringConfiguration1) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *MonitoringConfiguration1) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *MonitoringConfiguration1) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetPduSessionStatusCfg

`func (o *MonitoringConfiguration1) GetPduSessionStatusCfg() PduSessionStatusCfg1`

GetPduSessionStatusCfg returns the PduSessionStatusCfg field if non-nil, zero value otherwise.

### GetPduSessionStatusCfgOk

`func (o *MonitoringConfiguration1) GetPduSessionStatusCfgOk() (*PduSessionStatusCfg1, bool)`

GetPduSessionStatusCfgOk returns a tuple with the PduSessionStatusCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionStatusCfg

`func (o *MonitoringConfiguration1) SetPduSessionStatusCfg(v PduSessionStatusCfg1)`

SetPduSessionStatusCfg sets PduSessionStatusCfg field to given value.

### HasPduSessionStatusCfg

`func (o *MonitoringConfiguration1) HasPduSessionStatusCfg() bool`

HasPduSessionStatusCfg returns a boolean if a field has been set.

### GetReachabilityForSmsCfg

`func (o *MonitoringConfiguration1) GetReachabilityForSmsCfg() ReachabilityForSmsConfiguration`

GetReachabilityForSmsCfg returns the ReachabilityForSmsCfg field if non-nil, zero value otherwise.

### GetReachabilityForSmsCfgOk

`func (o *MonitoringConfiguration1) GetReachabilityForSmsCfgOk() (*ReachabilityForSmsConfiguration, bool)`

GetReachabilityForSmsCfgOk returns a tuple with the ReachabilityForSmsCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityForSmsCfg

`func (o *MonitoringConfiguration1) SetReachabilityForSmsCfg(v ReachabilityForSmsConfiguration)`

SetReachabilityForSmsCfg sets ReachabilityForSmsCfg field to given value.

### HasReachabilityForSmsCfg

`func (o *MonitoringConfiguration1) HasReachabilityForSmsCfg() bool`

HasReachabilityForSmsCfg returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *MonitoringConfiguration1) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *MonitoringConfiguration1) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *MonitoringConfiguration1) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *MonitoringConfiguration1) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetAfId

`func (o *MonitoringConfiguration1) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *MonitoringConfiguration1) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *MonitoringConfiguration1) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *MonitoringConfiguration1) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetReachabilityForDataCfg

`func (o *MonitoringConfiguration1) GetReachabilityForDataCfg() ReachabilityForDataConfiguration1`

GetReachabilityForDataCfg returns the ReachabilityForDataCfg field if non-nil, zero value otherwise.

### GetReachabilityForDataCfgOk

`func (o *MonitoringConfiguration1) GetReachabilityForDataCfgOk() (*ReachabilityForDataConfiguration1, bool)`

GetReachabilityForDataCfgOk returns a tuple with the ReachabilityForDataCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityForDataCfg

`func (o *MonitoringConfiguration1) SetReachabilityForDataCfg(v ReachabilityForDataConfiguration1)`

SetReachabilityForDataCfg sets ReachabilityForDataCfg field to given value.

### HasReachabilityForDataCfg

`func (o *MonitoringConfiguration1) HasReachabilityForDataCfg() bool`

HasReachabilityForDataCfg returns a boolean if a field has been set.

### GetIdleStatusInd

`func (o *MonitoringConfiguration1) GetIdleStatusInd() bool`

GetIdleStatusInd returns the IdleStatusInd field if non-nil, zero value otherwise.

### GetIdleStatusIndOk

`func (o *MonitoringConfiguration1) GetIdleStatusIndOk() (*bool, bool)`

GetIdleStatusIndOk returns a tuple with the IdleStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusInd

`func (o *MonitoringConfiguration1) SetIdleStatusInd(v bool)`

SetIdleStatusInd sets IdleStatusInd field to given value.

### HasIdleStatusInd

`func (o *MonitoringConfiguration1) HasIdleStatusInd() bool`

HasIdleStatusInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


