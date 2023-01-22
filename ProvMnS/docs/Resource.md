# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**MnsAgent** | Pointer to [**[]MnsAgentSingle**](MnsAgentSingle.md) |  | [optional] 
**Files** | Pointer to [**[]FilesSingle**](FilesSingle.md) |  | [optional] 
**HeartbeatControl** | Pointer to [**HeartbeatControlSingle**](HeartbeatControlSingle.md) |  | [optional] 
**MnsInfo** | Pointer to [**[]MnsInfoSingle**](MnsInfoSingle.md) |  | [optional] 
**MnsLabel** | Pointer to **string** |  | [optional] 
**MnsType** | Pointer to **string** |  | [optional] 
**MnsVersion** | Pointer to **string** |  | [optional] 
**MnsAddress** | Pointer to **string** |  | [optional] 
**MnsScope** | Pointer to **[]string** | List of the managed object instances that can be accessed using the MnS. If a complete SubNetwork can be accessed using the MnS, this attribute may contain the DN of the SubNetwork instead of the DNs of the individual managed entities within the SubNetwork. | [optional] 
**SubNetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**ManagedElement** | Pointer to [**[]ManagedElementSingle**](ManagedElementSingle.md) |  | [optional] 
**ManagementNode** | Pointer to [**[]ManagementNodeSingle**](ManagementNodeSingle.md) |  | [optional] 
**MeContext** | Pointer to [**[]MeContextSingle**](MeContextSingle.md) |  | [optional] 
**PerfMetricJob** | Pointer to [**[]PerfMetricJobSingle**](PerfMetricJobSingle.md) |  | [optional] 
**ThresholdMonitor** | Pointer to [**[]ThresholdMonitorSingle**](ThresholdMonitorSingle.md) |  | [optional] 
**TraceJob** | Pointer to [**[]TraceJobSingle**](TraceJobSingle.md) |  | [optional] 
**ManagementDataCollection** | Pointer to [**[]ManagementDataCollectionSingle**](ManagementDataCollectionSingle.md) |  | [optional] 
**NtfSubscriptionControl** | Pointer to [**[]NtfSubscriptionControlSingle**](NtfSubscriptionControlSingle.md) |  | [optional] 
**AlarmList** | Pointer to [**AlarmListSingle**](AlarmListSingle.md) |  | [optional] 
**FileDownloadJob** | Pointer to [**[]FileDownloadJobSingle**](FileDownloadJobSingle.md) |  | [optional] 
**MnsRegistry** | Pointer to [**MnsRegistrySingle**](MnsRegistrySingle.md) |  | [optional] 
**NRFrequency** | Pointer to [**[]NRFrequencySingle**](NRFrequencySingle.md) |  | [optional] 
**ExternalGnbCuCpFunction** | Pointer to [**[]ExternalGnbCuCpFunctionSingle**](ExternalGnbCuCpFunctionSingle.md) |  | [optional] 
**ExternalENBFunction** | Pointer to [**[]ExternalENBFunctionSingle**](ExternalENBFunctionSingle.md) |  | [optional] 
**EUtranFrequency** | Pointer to [**[]EUtranFrequencySingle**](EUtranFrequencySingle.md) |  | [optional] 
**DESManagementFunction** | Pointer to [**DESManagementFunctionSingle**](DESManagementFunctionSingle.md) |  | [optional] 
**DRACHOptimizationFunction** | Pointer to [**DRACHOptimizationFunctionSingle**](DRACHOptimizationFunctionSingle.md) |  | [optional] 
**DMROFunction** | Pointer to [**DMROFunctionSingle**](DMROFunctionSingle.md) |  | [optional] 
**DLBOFunction** | Pointer to [**DLBOFunctionSingle**](DLBOFunctionSingle.md) |  | [optional] 
**DPCIConfigurationFunction** | Pointer to [**DPCIConfigurationFunctionSingle**](DPCIConfigurationFunctionSingle.md) |  | [optional] 
**CPCIConfigurationFunction** | Pointer to [**CPCIConfigurationFunctionSingle**](CPCIConfigurationFunctionSingle.md) |  | [optional] 
**CESManagementFunction** | Pointer to [**CESManagementFunctionSingle**](CESManagementFunctionSingle.md) |  | [optional] 
**Configurable5QISet** | Pointer to [**[]Configurable5QISetSingle**](Configurable5QISetSingle.md) |  | [optional] 
**RimRSGlobal** | Pointer to [**RimRSGlobalSingle**](RimRSGlobalSingle.md) |  | [optional] 
**Dynamic5QISet** | Pointer to [**[]Dynamic5QISetSingle**](Dynamic5QISetSingle.md) |  | [optional] 
**CCOFunction** | Pointer to [**CCOFunctionSingle**](CCOFunctionSingle.md) |  | [optional] 
**GnbDuFunction** | Pointer to [**[]GnbDuFunctionSingle**](GnbDuFunctionSingle.md) |  | [optional] 
**GnbCuUpFunction** | Pointer to [**[]GnbCuUpFunctionSingle**](GnbCuUpFunctionSingle.md) |  | [optional] 
**GnbCuCpFunction** | Pointer to [**[]GnbCuCpFunctionSingle**](GnbCuCpFunctionSingle.md) |  | [optional] 
**ManagedNFService** | Pointer to [**[]ManagedNFServiceSingle**](ManagedNFServiceSingle.md) |  | [optional] 
**RRMPolicyRatio** | Pointer to [**[]RRMPolicyRatioSingle**](RRMPolicyRatioSingle.md) |  | [optional] 
**NrCellDu** | Pointer to [**[]NrCellDuSingle**](NrCellDuSingle.md) |  | [optional] 
**BwpMultiple** | Pointer to [**[]BwpSingle**](BwpSingle.md) |  | [optional] 
**NrSectorCarrierMultiple** | Pointer to [**[]NrSectorCarrierSingle**](NrSectorCarrierSingle.md) |  | [optional] 
**EPF1C** | Pointer to [**[]EPF1CSingle**](EPF1CSingle.md) |  | [optional] 
**EPF1U** | Pointer to [**[]EPF1USingle**](EPF1USingle.md) |  | [optional] 
**OperatorDU** | Pointer to [**[]OperatorDuSingle**](OperatorDuSingle.md) |  | [optional] 
**BWPSet** | Pointer to [**[]BWPSetSingle**](BWPSetSingle.md) |  | [optional] 
**EPE1** | Pointer to [**[]EPE1Single**](EPE1Single.md) |  | [optional] 
**EPXnU** | Pointer to [**[]EPXnUSingle**](EPXnUSingle.md) |  | [optional] 
**EPNgU** | Pointer to [**[]EPNgUSingle**](EPNgUSingle.md) |  | [optional] 
**EPX2U** | Pointer to [**[]EPX2USingle**](EPX2USingle.md) |  | [optional] 
**EPS1U** | Pointer to [**[]EPS1USingle**](EPS1USingle.md) |  | [optional] 
**NrCellCu** | Pointer to [**[]NrCellCuSingle**](NrCellCuSingle.md) |  | [optional] 
**EPXnC** | Pointer to [**[]EPXnCSingle**](EPXnCSingle.md) |  | [optional] 
**EPNgC** | Pointer to [**[]EPNgCSingle**](EPNgCSingle.md) |  | [optional] 
**EPX2C** | Pointer to [**[]EPX2CSingle**](EPX2CSingle.md) |  | [optional] 
**DANRManagementFunction** | Pointer to [**DANRManagementFunctionSingle**](DANRManagementFunctionSingle.md) |  | [optional] 
**GnbId** | Pointer to **string** |  | [optional] 
**GnbIdLength** | Pointer to **int32** |  | [optional] 
**NRCellRelation** | Pointer to [**[]NRCellRelationSingle**](NRCellRelationSingle.md) |  | [optional] 
**EUtranCellRelation** | Pointer to [**[]EUtranCellRelationSingle**](EUtranCellRelationSingle.md) |  | [optional] 
**NRFreqRelation** | Pointer to [**[]NRFreqRelationSingle**](NRFreqRelationSingle.md) |  | [optional] 
**EUtranFreqRelation** | Pointer to [**[]EUtranFreqRelationSingle**](EUtranFreqRelationSingle.md) |  | [optional] 
**NrOperatorCellDu** | Pointer to [**[]NrOperatorCellDuSingle**](NrOperatorCellDuSingle.md) |  | [optional] 
**CellLocalId** | Pointer to **int32** |  | [optional] 
**AdministrativeState** | Pointer to [**AdministrativeState**](AdministrativeState.md) |  | [optional] 
**PlmnInfoList** | Pointer to [**[]PlmnInfo**](PlmnInfo.md) |  | [optional] 
**NrTac** | Pointer to **int32** |  | [optional] 
**CommonBeamformingFunction** | Pointer to [**CommonBeamformingFunctionSingle**](CommonBeamformingFunctionSingle.md) |  | [optional] 
**BWPlist** | Pointer to **[]string** |  | [optional] 
**Beam** | Pointer to [**[]BeamSingle**](BeamSingle.md) |  | [optional] 
**RimRSSet** | Pointer to [**[]RimRSSetSingle**](RimRSSetSingle.md) |  | [optional] 
**ExternalNrCellCu** | Pointer to [**[]ExternalNrCellCuSingle**](ExternalNrCellCuSingle.md) |  | [optional] 
**ExternalEUTranCell** | Pointer to [**[]ExternalEUTranCellSingle**](ExternalEUTranCellSingle.md) |  | [optional] 
**ExternalAmfFunction** | Pointer to [**[]ExternalAmfFunctionSingle**](ExternalAmfFunctionSingle.md) |  | [optional] 
**ExternalNrfFunction** | Pointer to [**[]ExternalNrfFunctionSingle**](ExternalNrfFunctionSingle.md) |  | [optional] 
**ExternalNssfFunction** | Pointer to [**[]ExternalNssfFunctionSingle**](ExternalNssfFunctionSingle.md) |  | [optional] 
**AmfSet** | Pointer to [**[]AmfSetSingle**](AmfSetSingle.md) |  | [optional] 
**AmfRegion** | Pointer to [**[]AmfRegionSingle**](AmfRegionSingle.md) |  | [optional] 
**EcmConnectionInfo** | Pointer to [**[]EcmConnectionInfoSingle**](EcmConnectionInfoSingle.md) |  | [optional] 
**AmfFunction** | Pointer to [**[]AmfFunctionSingle**](AmfFunctionSingle.md) |  | [optional] 
**SmfFunction** | Pointer to [**[]SmfFunctionSingle**](SmfFunctionSingle.md) |  | [optional] 
**UpfFunction** | Pointer to [**[]UpfFunctionSingle**](UpfFunctionSingle.md) |  | [optional] 
**N3iwfFunction** | Pointer to [**[]N3iwfFunctionSingle**](N3iwfFunctionSingle.md) |  | [optional] 
**PcfFunction** | Pointer to [**[]PcfFunctionSingle**](PcfFunctionSingle.md) |  | [optional] 
**AusfFunction** | Pointer to [**[]AusfFunctionSingle**](AusfFunctionSingle.md) |  | [optional] 
**UdmFunction** | Pointer to [**[]UdmFunctionSingle**](UdmFunctionSingle.md) |  | [optional] 
**UdrFunction** | Pointer to [**[]UdrFunctionSingle**](UdrFunctionSingle.md) |  | [optional] 
**UdsfFunction** | Pointer to [**[]UdsfFunctionSingle**](UdsfFunctionSingle.md) |  | [optional] 
**NrfFunction** | Pointer to [**[]NrfFunctionSingle**](NrfFunctionSingle.md) |  | [optional] 
**NssfFunction** | Pointer to [**[]NssfFunctionSingle**](NssfFunctionSingle.md) |  | [optional] 
**SmsfFunction** | Pointer to [**[]SmsfFunctionSingle**](SmsfFunctionSingle.md) |  | [optional] 
**LmfFunction** | Pointer to [**[]LmfFunctionSingle**](LmfFunctionSingle.md) |  | [optional] 
**NgeirFunction** | Pointer to [**[]NgeirFunctionSingle**](NgeirFunctionSingle.md) |  | [optional] 
**SeppFunction** | Pointer to [**[]SeppFunctionSingle**](SeppFunctionSingle.md) |  | [optional] 
**NwdafFunction** | Pointer to [**[]NwdafFunctionSingle**](NwdafFunctionSingle.md) |  | [optional] 
**ScpFunction** | Pointer to [**[]ScpFunctionSingle**](ScpFunctionSingle.md) |  | [optional] 
**NefFunction** | Pointer to [**[]NefFunctionSingle**](NefFunctionSingle.md) |  | [optional] 
**EASDFFunction** | Pointer to [**[]EASDFFunctionSingle**](EASDFFunctionSingle.md) |  | [optional] 
**EPN2** | Pointer to [**[]EPN2Single**](EPN2Single.md) |  | [optional] 
**EPN8** | Pointer to [**[]EPN8Single**](EPN8Single.md) |  | [optional] 
**EPN11** | Pointer to [**[]EPN11Single**](EPN11Single.md) |  | [optional] 
**EPN12** | Pointer to [**[]EPN12Single**](EPN12Single.md) |  | [optional] 
**EPN14** | Pointer to [**[]EPN14Single**](EPN14Single.md) |  | [optional] 
**EPN15** | Pointer to [**[]EPN15Single**](EPN15Single.md) |  | [optional] 
**EPN17** | Pointer to [**[]EPN17Single**](EPN17Single.md) |  | [optional] 
**EPN20** | Pointer to [**[]EPN20Single**](EPN20Single.md) |  | [optional] 
**EPN22** | Pointer to [**[]EPN22Single**](EPN22Single.md) |  | [optional] 
**EPN26** | Pointer to [**[]EPN26Single**](EPN26Single.md) |  | [optional] 
**EP_NLS** | Pointer to [**[]EPNLSSingle**](EPNLSSingle.md) |  | [optional] 
**EP_NLG** | Pointer to [**[]EPNLGSingle**](EPNLGSingle.md) |  | [optional] 
**EPN4** | Pointer to [**[]EPN4Single**](EPN4Single.md) |  | [optional] 
**EPN7** | Pointer to [**[]EPN7Single**](EPN7Single.md) |  | [optional] 
**EPN10** | Pointer to [**[]EPN10Single**](EPN10Single.md) |  | [optional] 
**EPN16** | Pointer to [**[]EPN16Single**](EPN16Single.md) |  | [optional] 
**EPS5C** | Pointer to [**[]EPS5CSingle**](EPS5CSingle.md) |  | [optional] 
**FiveQiDscpMappingSet** | Pointer to [**FiveQiDscpMappingSetSingle**](FiveQiDscpMappingSetSingle.md) |  | [optional] 
**GtpUPathQoSMonitoringControl** | Pointer to [**GtpUPathQoSMonitoringControlSingle**](GtpUPathQoSMonitoringControlSingle.md) |  | [optional] 
**QFQoSMonitoringControl** | Pointer to [**QFQoSMonitoringControlSingle**](QFQoSMonitoringControlSingle.md) |  | [optional] 
**PredefinedPccRuleSet** | Pointer to [**PredefinedPccRuleSetSingle**](PredefinedPccRuleSetSingle.md) |  | [optional] 
**EPN3** | Pointer to [**[]EPN3Single**](EPN3Single.md) |  | [optional] 
**EPN6** | Pointer to [**[]EPN6Single**](EPN6Single.md) |  | [optional] 
**EPN9** | Pointer to [**[]EPN9Single**](EPN9Single.md) |  | [optional] 
**EPS5U** | Pointer to [**[]EPS5USingle**](EPS5USingle.md) |  | [optional] 
**EPN5** | Pointer to [**[]EPN5Single**](EPN5Single.md) |  | [optional] 
**EPRx** | Pointer to [**[]EPRxSingle**](EPRxSingle.md) |  | [optional] 
**EPN13** | Pointer to [**[]EPN13Single**](EPN13Single.md) |  | [optional] 
**EPN27** | Pointer to [**[]EPN27Single**](EPN27Single.md) |  | [optional] 
**EPN31** | Pointer to [**[]EPN31Single**](EPN31Single.md) |  | [optional] 
**EPN21** | Pointer to [**[]EPN21Single**](EPN21Single.md) |  | [optional] 
**EP_MAP_SMSC** | Pointer to [**[]EPMAPSMSCSingle**](EPMAPSMSCSingle.md) |  | [optional] 
**EPN32** | Pointer to [**[]EPN32Single**](EPN32Single.md) |  | [optional] 
**EPN33** | Pointer to [**[]EPN33Single**](EPN33Single.md) |  | [optional] 
**EPN60** | Pointer to [**[]EPN60Single**](EPN60Single.md) |  | [optional] 
**EPNpc4** | Pointer to [**[]EPNpc4Single**](EPNpc4Single.md) |  | [optional] 
**EPNpc6** | Pointer to [**[]EPNpc6Single**](EPNpc6Single.md) |  | [optional] 
**EPNpc7** | Pointer to [**[]EPNpc7Single**](EPNpc7Single.md) |  | [optional] 
**EPNpc8** | Pointer to [**[]EPNpc8Single**](EPNpc8Single.md) |  | [optional] 
**EPN88** | Pointer to [**[]EPN88Single**](EPN88Single.md) |  | [optional] 
**NetworkSlice** | Pointer to [**[]NetworkSliceSingle**](NetworkSliceSingle.md) |  | [optional] 
**NetworkSliceSubnet** | Pointer to [**[]NetworkSliceSubnetSingle**](NetworkSliceSubnetSingle.md) |  | [optional] 
**EPTransport** | Pointer to [**[]EPTransportSingle**](EPTransportSingle.md) |  | [optional] 
**NetworkSliceSubnetProviderCapabilities** | Pointer to [**[]NetworkSliceSubnetProviderCapabilitiesSingle**](NetworkSliceSubnetProviderCapabilitiesSingle.md) |  | [optional] 
**FeasibilityCheckAndReservationJob** | Pointer to [**[]FeasibilityCheckAndReservationJobSingle**](FeasibilityCheckAndReservationJobSingle.md) |  | [optional] 
**AssuranceGoal** | Pointer to [**[]AssuranceGoalSingle**](AssuranceGoalSingle.md) |  | [optional] 
**AssuranceClosedControlLoop** | Pointer to [**[]AssuranceClosedControlLoopSingle**](AssuranceClosedControlLoopSingle.md) |  | [optional] 
**UserLabel** | Pointer to **string** |  | [optional] 
**IntentExpectations** | Pointer to [**[]OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation**](OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation.md) |  | [optional] 
**IntentContexts** | Pointer to [**[]IntentContext**](IntentContext.md) |  | [optional] 
**IntentFulfilmentInfo** | Pointer to [**FulfilmentInfo**](FulfilmentInfo.md) |  | [optional] 
**MDAFunction** | Pointer to [**[]MDAFunctionSingle**](MDAFunctionSingle.md) |  | [optional] 
**MDAReport** | Pointer to [**[]MDAReport**](MDAReport.md) |  | [optional] 
**MDARequest** | Pointer to [**[]MDARequestSingle**](MDARequestSingle.md) |  | [optional] 
**MLTrainingFunction** | Pointer to [**[]MLTrainingFunctionSingle**](MLTrainingFunctionSingle.md) |  | [optional] 
**MLTrainingRequest** | Pointer to [**[]MLTrainingRequestSingle**](MLTrainingRequestSingle.md) |  | [optional] 
**MLTrainingProcess** | Pointer to [**[]MLTrainingProcessSingle**](MLTrainingProcessSingle.md) |  | [optional] 
**MLTrainingReport** | Pointer to [**[]MLTrainingReportSingle**](MLTrainingReportSingle.md) |  | [optional] 

## Methods

### NewResource

`func NewResource(id NullableString, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Resource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Resource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *Resource) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *Resource) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *Resource) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *Resource) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *Resource) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *Resource) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *Resource) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *Resource) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetAttributes

`func (o *Resource) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Resource) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Resource) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Resource) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *Resource) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *Resource) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *Resource) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *Resource) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetMnsAgent

`func (o *Resource) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *Resource) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *Resource) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *Resource) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetFiles

`func (o *Resource) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Resource) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Resource) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Resource) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetHeartbeatControl

`func (o *Resource) GetHeartbeatControl() HeartbeatControlSingle`

GetHeartbeatControl returns the HeartbeatControl field if non-nil, zero value otherwise.

### GetHeartbeatControlOk

`func (o *Resource) GetHeartbeatControlOk() (*HeartbeatControlSingle, bool)`

GetHeartbeatControlOk returns a tuple with the HeartbeatControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatControl

`func (o *Resource) SetHeartbeatControl(v HeartbeatControlSingle)`

SetHeartbeatControl sets HeartbeatControl field to given value.

### HasHeartbeatControl

`func (o *Resource) HasHeartbeatControl() bool`

HasHeartbeatControl returns a boolean if a field has been set.

### GetMnsInfo

`func (o *Resource) GetMnsInfo() []MnsInfoSingle`

GetMnsInfo returns the MnsInfo field if non-nil, zero value otherwise.

### GetMnsInfoOk

`func (o *Resource) GetMnsInfoOk() (*[]MnsInfoSingle, bool)`

GetMnsInfoOk returns a tuple with the MnsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsInfo

`func (o *Resource) SetMnsInfo(v []MnsInfoSingle)`

SetMnsInfo sets MnsInfo field to given value.

### HasMnsInfo

`func (o *Resource) HasMnsInfo() bool`

HasMnsInfo returns a boolean if a field has been set.

### GetMnsLabel

`func (o *Resource) GetMnsLabel() string`

GetMnsLabel returns the MnsLabel field if non-nil, zero value otherwise.

### GetMnsLabelOk

`func (o *Resource) GetMnsLabelOk() (*string, bool)`

GetMnsLabelOk returns a tuple with the MnsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsLabel

`func (o *Resource) SetMnsLabel(v string)`

SetMnsLabel sets MnsLabel field to given value.

### HasMnsLabel

`func (o *Resource) HasMnsLabel() bool`

HasMnsLabel returns a boolean if a field has been set.

### GetMnsType

`func (o *Resource) GetMnsType() string`

GetMnsType returns the MnsType field if non-nil, zero value otherwise.

### GetMnsTypeOk

`func (o *Resource) GetMnsTypeOk() (*string, bool)`

GetMnsTypeOk returns a tuple with the MnsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsType

`func (o *Resource) SetMnsType(v string)`

SetMnsType sets MnsType field to given value.

### HasMnsType

`func (o *Resource) HasMnsType() bool`

HasMnsType returns a boolean if a field has been set.

### GetMnsVersion

`func (o *Resource) GetMnsVersion() string`

GetMnsVersion returns the MnsVersion field if non-nil, zero value otherwise.

### GetMnsVersionOk

`func (o *Resource) GetMnsVersionOk() (*string, bool)`

GetMnsVersionOk returns a tuple with the MnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsVersion

`func (o *Resource) SetMnsVersion(v string)`

SetMnsVersion sets MnsVersion field to given value.

### HasMnsVersion

`func (o *Resource) HasMnsVersion() bool`

HasMnsVersion returns a boolean if a field has been set.

### GetMnsAddress

`func (o *Resource) GetMnsAddress() string`

GetMnsAddress returns the MnsAddress field if non-nil, zero value otherwise.

### GetMnsAddressOk

`func (o *Resource) GetMnsAddressOk() (*string, bool)`

GetMnsAddressOk returns a tuple with the MnsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAddress

`func (o *Resource) SetMnsAddress(v string)`

SetMnsAddress sets MnsAddress field to given value.

### HasMnsAddress

`func (o *Resource) HasMnsAddress() bool`

HasMnsAddress returns a boolean if a field has been set.

### GetMnsScope

`func (o *Resource) GetMnsScope() []string`

GetMnsScope returns the MnsScope field if non-nil, zero value otherwise.

### GetMnsScopeOk

`func (o *Resource) GetMnsScopeOk() (*[]string, bool)`

GetMnsScopeOk returns a tuple with the MnsScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsScope

`func (o *Resource) SetMnsScope(v []string)`

SetMnsScope sets MnsScope field to given value.

### HasMnsScope

`func (o *Resource) HasMnsScope() bool`

HasMnsScope returns a boolean if a field has been set.

### GetSubNetwork

`func (o *Resource) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *Resource) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *Resource) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *Resource) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetManagedElement

`func (o *Resource) GetManagedElement() []ManagedElementSingle`

GetManagedElement returns the ManagedElement field if non-nil, zero value otherwise.

### GetManagedElementOk

`func (o *Resource) GetManagedElementOk() (*[]ManagedElementSingle, bool)`

GetManagedElementOk returns a tuple with the ManagedElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedElement

`func (o *Resource) SetManagedElement(v []ManagedElementSingle)`

SetManagedElement sets ManagedElement field to given value.

### HasManagedElement

`func (o *Resource) HasManagedElement() bool`

HasManagedElement returns a boolean if a field has been set.

### GetManagementNode

`func (o *Resource) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *Resource) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *Resource) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *Resource) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMeContext

`func (o *Resource) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *Resource) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *Resource) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *Resource) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *Resource) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *Resource) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *Resource) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *Resource) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *Resource) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *Resource) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *Resource) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *Resource) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *Resource) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *Resource) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *Resource) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *Resource) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *Resource) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *Resource) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *Resource) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *Resource) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *Resource) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *Resource) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *Resource) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *Resource) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *Resource) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *Resource) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *Resource) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *Resource) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *Resource) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *Resource) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *Resource) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *Resource) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *Resource) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *Resource) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *Resource) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *Resource) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.

### GetNRFrequency

`func (o *Resource) GetNRFrequency() []NRFrequencySingle`

GetNRFrequency returns the NRFrequency field if non-nil, zero value otherwise.

### GetNRFrequencyOk

`func (o *Resource) GetNRFrequencyOk() (*[]NRFrequencySingle, bool)`

GetNRFrequencyOk returns a tuple with the NRFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRFrequency

`func (o *Resource) SetNRFrequency(v []NRFrequencySingle)`

SetNRFrequency sets NRFrequency field to given value.

### HasNRFrequency

`func (o *Resource) HasNRFrequency() bool`

HasNRFrequency returns a boolean if a field has been set.

### GetExternalGnbCuCpFunction

`func (o *Resource) GetExternalGnbCuCpFunction() []ExternalGnbCuCpFunctionSingle`

GetExternalGnbCuCpFunction returns the ExternalGnbCuCpFunction field if non-nil, zero value otherwise.

### GetExternalGnbCuCpFunctionOk

`func (o *Resource) GetExternalGnbCuCpFunctionOk() (*[]ExternalGnbCuCpFunctionSingle, bool)`

GetExternalGnbCuCpFunctionOk returns a tuple with the ExternalGnbCuCpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGnbCuCpFunction

`func (o *Resource) SetExternalGnbCuCpFunction(v []ExternalGnbCuCpFunctionSingle)`

SetExternalGnbCuCpFunction sets ExternalGnbCuCpFunction field to given value.

### HasExternalGnbCuCpFunction

`func (o *Resource) HasExternalGnbCuCpFunction() bool`

HasExternalGnbCuCpFunction returns a boolean if a field has been set.

### GetExternalENBFunction

`func (o *Resource) GetExternalENBFunction() []ExternalENBFunctionSingle`

GetExternalENBFunction returns the ExternalENBFunction field if non-nil, zero value otherwise.

### GetExternalENBFunctionOk

`func (o *Resource) GetExternalENBFunctionOk() (*[]ExternalENBFunctionSingle, bool)`

GetExternalENBFunctionOk returns a tuple with the ExternalENBFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalENBFunction

`func (o *Resource) SetExternalENBFunction(v []ExternalENBFunctionSingle)`

SetExternalENBFunction sets ExternalENBFunction field to given value.

### HasExternalENBFunction

`func (o *Resource) HasExternalENBFunction() bool`

HasExternalENBFunction returns a boolean if a field has been set.

### GetEUtranFrequency

`func (o *Resource) GetEUtranFrequency() []EUtranFrequencySingle`

GetEUtranFrequency returns the EUtranFrequency field if non-nil, zero value otherwise.

### GetEUtranFrequencyOk

`func (o *Resource) GetEUtranFrequencyOk() (*[]EUtranFrequencySingle, bool)`

GetEUtranFrequencyOk returns a tuple with the EUtranFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUtranFrequency

`func (o *Resource) SetEUtranFrequency(v []EUtranFrequencySingle)`

SetEUtranFrequency sets EUtranFrequency field to given value.

### HasEUtranFrequency

`func (o *Resource) HasEUtranFrequency() bool`

HasEUtranFrequency returns a boolean if a field has been set.

### GetDESManagementFunction

`func (o *Resource) GetDESManagementFunction() DESManagementFunctionSingle`

GetDESManagementFunction returns the DESManagementFunction field if non-nil, zero value otherwise.

### GetDESManagementFunctionOk

`func (o *Resource) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool)`

GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDESManagementFunction

`func (o *Resource) SetDESManagementFunction(v DESManagementFunctionSingle)`

SetDESManagementFunction sets DESManagementFunction field to given value.

### HasDESManagementFunction

`func (o *Resource) HasDESManagementFunction() bool`

HasDESManagementFunction returns a boolean if a field has been set.

### GetDRACHOptimizationFunction

`func (o *Resource) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle`

GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field if non-nil, zero value otherwise.

### GetDRACHOptimizationFunctionOk

`func (o *Resource) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool)`

GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRACHOptimizationFunction

`func (o *Resource) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle)`

SetDRACHOptimizationFunction sets DRACHOptimizationFunction field to given value.

### HasDRACHOptimizationFunction

`func (o *Resource) HasDRACHOptimizationFunction() bool`

HasDRACHOptimizationFunction returns a boolean if a field has been set.

### GetDMROFunction

`func (o *Resource) GetDMROFunction() DMROFunctionSingle`

GetDMROFunction returns the DMROFunction field if non-nil, zero value otherwise.

### GetDMROFunctionOk

`func (o *Resource) GetDMROFunctionOk() (*DMROFunctionSingle, bool)`

GetDMROFunctionOk returns a tuple with the DMROFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMROFunction

`func (o *Resource) SetDMROFunction(v DMROFunctionSingle)`

SetDMROFunction sets DMROFunction field to given value.

### HasDMROFunction

`func (o *Resource) HasDMROFunction() bool`

HasDMROFunction returns a boolean if a field has been set.

### GetDLBOFunction

`func (o *Resource) GetDLBOFunction() DLBOFunctionSingle`

GetDLBOFunction returns the DLBOFunction field if non-nil, zero value otherwise.

### GetDLBOFunctionOk

`func (o *Resource) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool)`

GetDLBOFunctionOk returns a tuple with the DLBOFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLBOFunction

`func (o *Resource) SetDLBOFunction(v DLBOFunctionSingle)`

SetDLBOFunction sets DLBOFunction field to given value.

### HasDLBOFunction

`func (o *Resource) HasDLBOFunction() bool`

HasDLBOFunction returns a boolean if a field has been set.

### GetDPCIConfigurationFunction

`func (o *Resource) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle`

GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetDPCIConfigurationFunctionOk

`func (o *Resource) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool)`

GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPCIConfigurationFunction

`func (o *Resource) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle)`

SetDPCIConfigurationFunction sets DPCIConfigurationFunction field to given value.

### HasDPCIConfigurationFunction

`func (o *Resource) HasDPCIConfigurationFunction() bool`

HasDPCIConfigurationFunction returns a boolean if a field has been set.

### GetCPCIConfigurationFunction

`func (o *Resource) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle`

GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetCPCIConfigurationFunctionOk

`func (o *Resource) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool)`

GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPCIConfigurationFunction

`func (o *Resource) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle)`

SetCPCIConfigurationFunction sets CPCIConfigurationFunction field to given value.

### HasCPCIConfigurationFunction

`func (o *Resource) HasCPCIConfigurationFunction() bool`

HasCPCIConfigurationFunction returns a boolean if a field has been set.

### GetCESManagementFunction

`func (o *Resource) GetCESManagementFunction() CESManagementFunctionSingle`

GetCESManagementFunction returns the CESManagementFunction field if non-nil, zero value otherwise.

### GetCESManagementFunctionOk

`func (o *Resource) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool)`

GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCESManagementFunction

`func (o *Resource) SetCESManagementFunction(v CESManagementFunctionSingle)`

SetCESManagementFunction sets CESManagementFunction field to given value.

### HasCESManagementFunction

`func (o *Resource) HasCESManagementFunction() bool`

HasCESManagementFunction returns a boolean if a field has been set.

### GetConfigurable5QISet

`func (o *Resource) GetConfigurable5QISet() []Configurable5QISetSingle`

GetConfigurable5QISet returns the Configurable5QISet field if non-nil, zero value otherwise.

### GetConfigurable5QISetOk

`func (o *Resource) GetConfigurable5QISetOk() (*[]Configurable5QISetSingle, bool)`

GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable5QISet

`func (o *Resource) SetConfigurable5QISet(v []Configurable5QISetSingle)`

SetConfigurable5QISet sets Configurable5QISet field to given value.

### HasConfigurable5QISet

`func (o *Resource) HasConfigurable5QISet() bool`

HasConfigurable5QISet returns a boolean if a field has been set.

### GetRimRSGlobal

`func (o *Resource) GetRimRSGlobal() RimRSGlobalSingle`

GetRimRSGlobal returns the RimRSGlobal field if non-nil, zero value otherwise.

### GetRimRSGlobalOk

`func (o *Resource) GetRimRSGlobalOk() (*RimRSGlobalSingle, bool)`

GetRimRSGlobalOk returns a tuple with the RimRSGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRimRSGlobal

`func (o *Resource) SetRimRSGlobal(v RimRSGlobalSingle)`

SetRimRSGlobal sets RimRSGlobal field to given value.

### HasRimRSGlobal

`func (o *Resource) HasRimRSGlobal() bool`

HasRimRSGlobal returns a boolean if a field has been set.

### GetDynamic5QISet

`func (o *Resource) GetDynamic5QISet() []Dynamic5QISetSingle`

GetDynamic5QISet returns the Dynamic5QISet field if non-nil, zero value otherwise.

### GetDynamic5QISetOk

`func (o *Resource) GetDynamic5QISetOk() (*[]Dynamic5QISetSingle, bool)`

GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic5QISet

`func (o *Resource) SetDynamic5QISet(v []Dynamic5QISetSingle)`

SetDynamic5QISet sets Dynamic5QISet field to given value.

### HasDynamic5QISet

`func (o *Resource) HasDynamic5QISet() bool`

HasDynamic5QISet returns a boolean if a field has been set.

### GetCCOFunction

`func (o *Resource) GetCCOFunction() CCOFunctionSingle`

GetCCOFunction returns the CCOFunction field if non-nil, zero value otherwise.

### GetCCOFunctionOk

`func (o *Resource) GetCCOFunctionOk() (*CCOFunctionSingle, bool)`

GetCCOFunctionOk returns a tuple with the CCOFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCCOFunction

`func (o *Resource) SetCCOFunction(v CCOFunctionSingle)`

SetCCOFunction sets CCOFunction field to given value.

### HasCCOFunction

`func (o *Resource) HasCCOFunction() bool`

HasCCOFunction returns a boolean if a field has been set.

### GetGnbDuFunction

`func (o *Resource) GetGnbDuFunction() []GnbDuFunctionSingle`

GetGnbDuFunction returns the GnbDuFunction field if non-nil, zero value otherwise.

### GetGnbDuFunctionOk

`func (o *Resource) GetGnbDuFunctionOk() (*[]GnbDuFunctionSingle, bool)`

GetGnbDuFunctionOk returns a tuple with the GnbDuFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbDuFunction

`func (o *Resource) SetGnbDuFunction(v []GnbDuFunctionSingle)`

SetGnbDuFunction sets GnbDuFunction field to given value.

### HasGnbDuFunction

`func (o *Resource) HasGnbDuFunction() bool`

HasGnbDuFunction returns a boolean if a field has been set.

### GetGnbCuUpFunction

`func (o *Resource) GetGnbCuUpFunction() []GnbCuUpFunctionSingle`

GetGnbCuUpFunction returns the GnbCuUpFunction field if non-nil, zero value otherwise.

### GetGnbCuUpFunctionOk

`func (o *Resource) GetGnbCuUpFunctionOk() (*[]GnbCuUpFunctionSingle, bool)`

GetGnbCuUpFunctionOk returns a tuple with the GnbCuUpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbCuUpFunction

`func (o *Resource) SetGnbCuUpFunction(v []GnbCuUpFunctionSingle)`

SetGnbCuUpFunction sets GnbCuUpFunction field to given value.

### HasGnbCuUpFunction

`func (o *Resource) HasGnbCuUpFunction() bool`

HasGnbCuUpFunction returns a boolean if a field has been set.

### GetGnbCuCpFunction

`func (o *Resource) GetGnbCuCpFunction() []GnbCuCpFunctionSingle`

GetGnbCuCpFunction returns the GnbCuCpFunction field if non-nil, zero value otherwise.

### GetGnbCuCpFunctionOk

`func (o *Resource) GetGnbCuCpFunctionOk() (*[]GnbCuCpFunctionSingle, bool)`

GetGnbCuCpFunctionOk returns a tuple with the GnbCuCpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbCuCpFunction

`func (o *Resource) SetGnbCuCpFunction(v []GnbCuCpFunctionSingle)`

SetGnbCuCpFunction sets GnbCuCpFunction field to given value.

### HasGnbCuCpFunction

`func (o *Resource) HasGnbCuCpFunction() bool`

HasGnbCuCpFunction returns a boolean if a field has been set.

### GetManagedNFService

`func (o *Resource) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *Resource) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *Resource) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *Resource) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetRRMPolicyRatio

`func (o *Resource) GetRRMPolicyRatio() []RRMPolicyRatioSingle`

GetRRMPolicyRatio returns the RRMPolicyRatio field if non-nil, zero value otherwise.

### GetRRMPolicyRatioOk

`func (o *Resource) GetRRMPolicyRatioOk() (*[]RRMPolicyRatioSingle, bool)`

GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRRMPolicyRatio

`func (o *Resource) SetRRMPolicyRatio(v []RRMPolicyRatioSingle)`

SetRRMPolicyRatio sets RRMPolicyRatio field to given value.

### HasRRMPolicyRatio

`func (o *Resource) HasRRMPolicyRatio() bool`

HasRRMPolicyRatio returns a boolean if a field has been set.

### GetNrCellDu

`func (o *Resource) GetNrCellDu() []NrCellDuSingle`

GetNrCellDu returns the NrCellDu field if non-nil, zero value otherwise.

### GetNrCellDuOk

`func (o *Resource) GetNrCellDuOk() (*[]NrCellDuSingle, bool)`

GetNrCellDuOk returns a tuple with the NrCellDu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrCellDu

`func (o *Resource) SetNrCellDu(v []NrCellDuSingle)`

SetNrCellDu sets NrCellDu field to given value.

### HasNrCellDu

`func (o *Resource) HasNrCellDu() bool`

HasNrCellDu returns a boolean if a field has been set.

### GetBwpMultiple

`func (o *Resource) GetBwpMultiple() []BwpSingle`

GetBwpMultiple returns the BwpMultiple field if non-nil, zero value otherwise.

### GetBwpMultipleOk

`func (o *Resource) GetBwpMultipleOk() (*[]BwpSingle, bool)`

GetBwpMultipleOk returns a tuple with the BwpMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwpMultiple

`func (o *Resource) SetBwpMultiple(v []BwpSingle)`

SetBwpMultiple sets BwpMultiple field to given value.

### HasBwpMultiple

`func (o *Resource) HasBwpMultiple() bool`

HasBwpMultiple returns a boolean if a field has been set.

### GetNrSectorCarrierMultiple

`func (o *Resource) GetNrSectorCarrierMultiple() []NrSectorCarrierSingle`

GetNrSectorCarrierMultiple returns the NrSectorCarrierMultiple field if non-nil, zero value otherwise.

### GetNrSectorCarrierMultipleOk

`func (o *Resource) GetNrSectorCarrierMultipleOk() (*[]NrSectorCarrierSingle, bool)`

GetNrSectorCarrierMultipleOk returns a tuple with the NrSectorCarrierMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrSectorCarrierMultiple

`func (o *Resource) SetNrSectorCarrierMultiple(v []NrSectorCarrierSingle)`

SetNrSectorCarrierMultiple sets NrSectorCarrierMultiple field to given value.

### HasNrSectorCarrierMultiple

`func (o *Resource) HasNrSectorCarrierMultiple() bool`

HasNrSectorCarrierMultiple returns a boolean if a field has been set.

### GetEPF1C

`func (o *Resource) GetEPF1C() []EPF1CSingle`

GetEPF1C returns the EPF1C field if non-nil, zero value otherwise.

### GetEPF1COk

`func (o *Resource) GetEPF1COk() (*[]EPF1CSingle, bool)`

GetEPF1COk returns a tuple with the EPF1C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1C

`func (o *Resource) SetEPF1C(v []EPF1CSingle)`

SetEPF1C sets EPF1C field to given value.

### HasEPF1C

`func (o *Resource) HasEPF1C() bool`

HasEPF1C returns a boolean if a field has been set.

### GetEPF1U

`func (o *Resource) GetEPF1U() []EPF1USingle`

GetEPF1U returns the EPF1U field if non-nil, zero value otherwise.

### GetEPF1UOk

`func (o *Resource) GetEPF1UOk() (*[]EPF1USingle, bool)`

GetEPF1UOk returns a tuple with the EPF1U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1U

`func (o *Resource) SetEPF1U(v []EPF1USingle)`

SetEPF1U sets EPF1U field to given value.

### HasEPF1U

`func (o *Resource) HasEPF1U() bool`

HasEPF1U returns a boolean if a field has been set.

### GetOperatorDU

`func (o *Resource) GetOperatorDU() []OperatorDuSingle`

GetOperatorDU returns the OperatorDU field if non-nil, zero value otherwise.

### GetOperatorDUOk

`func (o *Resource) GetOperatorDUOk() (*[]OperatorDuSingle, bool)`

GetOperatorDUOk returns a tuple with the OperatorDU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorDU

`func (o *Resource) SetOperatorDU(v []OperatorDuSingle)`

SetOperatorDU sets OperatorDU field to given value.

### HasOperatorDU

`func (o *Resource) HasOperatorDU() bool`

HasOperatorDU returns a boolean if a field has been set.

### GetBWPSet

`func (o *Resource) GetBWPSet() []BWPSetSingle`

GetBWPSet returns the BWPSet field if non-nil, zero value otherwise.

### GetBWPSetOk

`func (o *Resource) GetBWPSetOk() (*[]BWPSetSingle, bool)`

GetBWPSetOk returns a tuple with the BWPSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWPSet

`func (o *Resource) SetBWPSet(v []BWPSetSingle)`

SetBWPSet sets BWPSet field to given value.

### HasBWPSet

`func (o *Resource) HasBWPSet() bool`

HasBWPSet returns a boolean if a field has been set.

### GetEPE1

`func (o *Resource) GetEPE1() []EPE1Single`

GetEPE1 returns the EPE1 field if non-nil, zero value otherwise.

### GetEPE1Ok

`func (o *Resource) GetEPE1Ok() (*[]EPE1Single, bool)`

GetEPE1Ok returns a tuple with the EPE1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPE1

`func (o *Resource) SetEPE1(v []EPE1Single)`

SetEPE1 sets EPE1 field to given value.

### HasEPE1

`func (o *Resource) HasEPE1() bool`

HasEPE1 returns a boolean if a field has been set.

### GetEPXnU

`func (o *Resource) GetEPXnU() []EPXnUSingle`

GetEPXnU returns the EPXnU field if non-nil, zero value otherwise.

### GetEPXnUOk

`func (o *Resource) GetEPXnUOk() (*[]EPXnUSingle, bool)`

GetEPXnUOk returns a tuple with the EPXnU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPXnU

`func (o *Resource) SetEPXnU(v []EPXnUSingle)`

SetEPXnU sets EPXnU field to given value.

### HasEPXnU

`func (o *Resource) HasEPXnU() bool`

HasEPXnU returns a boolean if a field has been set.

### GetEPNgU

`func (o *Resource) GetEPNgU() []EPNgUSingle`

GetEPNgU returns the EPNgU field if non-nil, zero value otherwise.

### GetEPNgUOk

`func (o *Resource) GetEPNgUOk() (*[]EPNgUSingle, bool)`

GetEPNgUOk returns a tuple with the EPNgU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNgU

`func (o *Resource) SetEPNgU(v []EPNgUSingle)`

SetEPNgU sets EPNgU field to given value.

### HasEPNgU

`func (o *Resource) HasEPNgU() bool`

HasEPNgU returns a boolean if a field has been set.

### GetEPX2U

`func (o *Resource) GetEPX2U() []EPX2USingle`

GetEPX2U returns the EPX2U field if non-nil, zero value otherwise.

### GetEPX2UOk

`func (o *Resource) GetEPX2UOk() (*[]EPX2USingle, bool)`

GetEPX2UOk returns a tuple with the EPX2U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPX2U

`func (o *Resource) SetEPX2U(v []EPX2USingle)`

SetEPX2U sets EPX2U field to given value.

### HasEPX2U

`func (o *Resource) HasEPX2U() bool`

HasEPX2U returns a boolean if a field has been set.

### GetEPS1U

`func (o *Resource) GetEPS1U() []EPS1USingle`

GetEPS1U returns the EPS1U field if non-nil, zero value otherwise.

### GetEPS1UOk

`func (o *Resource) GetEPS1UOk() (*[]EPS1USingle, bool)`

GetEPS1UOk returns a tuple with the EPS1U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPS1U

`func (o *Resource) SetEPS1U(v []EPS1USingle)`

SetEPS1U sets EPS1U field to given value.

### HasEPS1U

`func (o *Resource) HasEPS1U() bool`

HasEPS1U returns a boolean if a field has been set.

### GetNrCellCu

`func (o *Resource) GetNrCellCu() []NrCellCuSingle`

GetNrCellCu returns the NrCellCu field if non-nil, zero value otherwise.

### GetNrCellCuOk

`func (o *Resource) GetNrCellCuOk() (*[]NrCellCuSingle, bool)`

GetNrCellCuOk returns a tuple with the NrCellCu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrCellCu

`func (o *Resource) SetNrCellCu(v []NrCellCuSingle)`

SetNrCellCu sets NrCellCu field to given value.

### HasNrCellCu

`func (o *Resource) HasNrCellCu() bool`

HasNrCellCu returns a boolean if a field has been set.

### GetEPXnC

`func (o *Resource) GetEPXnC() []EPXnCSingle`

GetEPXnC returns the EPXnC field if non-nil, zero value otherwise.

### GetEPXnCOk

`func (o *Resource) GetEPXnCOk() (*[]EPXnCSingle, bool)`

GetEPXnCOk returns a tuple with the EPXnC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPXnC

`func (o *Resource) SetEPXnC(v []EPXnCSingle)`

SetEPXnC sets EPXnC field to given value.

### HasEPXnC

`func (o *Resource) HasEPXnC() bool`

HasEPXnC returns a boolean if a field has been set.

### GetEPNgC

`func (o *Resource) GetEPNgC() []EPNgCSingle`

GetEPNgC returns the EPNgC field if non-nil, zero value otherwise.

### GetEPNgCOk

`func (o *Resource) GetEPNgCOk() (*[]EPNgCSingle, bool)`

GetEPNgCOk returns a tuple with the EPNgC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNgC

`func (o *Resource) SetEPNgC(v []EPNgCSingle)`

SetEPNgC sets EPNgC field to given value.

### HasEPNgC

`func (o *Resource) HasEPNgC() bool`

HasEPNgC returns a boolean if a field has been set.

### GetEPX2C

`func (o *Resource) GetEPX2C() []EPX2CSingle`

GetEPX2C returns the EPX2C field if non-nil, zero value otherwise.

### GetEPX2COk

`func (o *Resource) GetEPX2COk() (*[]EPX2CSingle, bool)`

GetEPX2COk returns a tuple with the EPX2C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPX2C

`func (o *Resource) SetEPX2C(v []EPX2CSingle)`

SetEPX2C sets EPX2C field to given value.

### HasEPX2C

`func (o *Resource) HasEPX2C() bool`

HasEPX2C returns a boolean if a field has been set.

### GetDANRManagementFunction

`func (o *Resource) GetDANRManagementFunction() DANRManagementFunctionSingle`

GetDANRManagementFunction returns the DANRManagementFunction field if non-nil, zero value otherwise.

### GetDANRManagementFunctionOk

`func (o *Resource) GetDANRManagementFunctionOk() (*DANRManagementFunctionSingle, bool)`

GetDANRManagementFunctionOk returns a tuple with the DANRManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDANRManagementFunction

`func (o *Resource) SetDANRManagementFunction(v DANRManagementFunctionSingle)`

SetDANRManagementFunction sets DANRManagementFunction field to given value.

### HasDANRManagementFunction

`func (o *Resource) HasDANRManagementFunction() bool`

HasDANRManagementFunction returns a boolean if a field has been set.

### GetGnbId

`func (o *Resource) GetGnbId() string`

GetGnbId returns the GnbId field if non-nil, zero value otherwise.

### GetGnbIdOk

`func (o *Resource) GetGnbIdOk() (*string, bool)`

GetGnbIdOk returns a tuple with the GnbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbId

`func (o *Resource) SetGnbId(v string)`

SetGnbId sets GnbId field to given value.

### HasGnbId

`func (o *Resource) HasGnbId() bool`

HasGnbId returns a boolean if a field has been set.

### GetGnbIdLength

`func (o *Resource) GetGnbIdLength() int32`

GetGnbIdLength returns the GnbIdLength field if non-nil, zero value otherwise.

### GetGnbIdLengthOk

`func (o *Resource) GetGnbIdLengthOk() (*int32, bool)`

GetGnbIdLengthOk returns a tuple with the GnbIdLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbIdLength

`func (o *Resource) SetGnbIdLength(v int32)`

SetGnbIdLength sets GnbIdLength field to given value.

### HasGnbIdLength

`func (o *Resource) HasGnbIdLength() bool`

HasGnbIdLength returns a boolean if a field has been set.

### GetNRCellRelation

`func (o *Resource) GetNRCellRelation() []NRCellRelationSingle`

GetNRCellRelation returns the NRCellRelation field if non-nil, zero value otherwise.

### GetNRCellRelationOk

`func (o *Resource) GetNRCellRelationOk() (*[]NRCellRelationSingle, bool)`

GetNRCellRelationOk returns a tuple with the NRCellRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRCellRelation

`func (o *Resource) SetNRCellRelation(v []NRCellRelationSingle)`

SetNRCellRelation sets NRCellRelation field to given value.

### HasNRCellRelation

`func (o *Resource) HasNRCellRelation() bool`

HasNRCellRelation returns a boolean if a field has been set.

### GetEUtranCellRelation

`func (o *Resource) GetEUtranCellRelation() []EUtranCellRelationSingle`

GetEUtranCellRelation returns the EUtranCellRelation field if non-nil, zero value otherwise.

### GetEUtranCellRelationOk

`func (o *Resource) GetEUtranCellRelationOk() (*[]EUtranCellRelationSingle, bool)`

GetEUtranCellRelationOk returns a tuple with the EUtranCellRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUtranCellRelation

`func (o *Resource) SetEUtranCellRelation(v []EUtranCellRelationSingle)`

SetEUtranCellRelation sets EUtranCellRelation field to given value.

### HasEUtranCellRelation

`func (o *Resource) HasEUtranCellRelation() bool`

HasEUtranCellRelation returns a boolean if a field has been set.

### GetNRFreqRelation

`func (o *Resource) GetNRFreqRelation() []NRFreqRelationSingle`

GetNRFreqRelation returns the NRFreqRelation field if non-nil, zero value otherwise.

### GetNRFreqRelationOk

`func (o *Resource) GetNRFreqRelationOk() (*[]NRFreqRelationSingle, bool)`

GetNRFreqRelationOk returns a tuple with the NRFreqRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRFreqRelation

`func (o *Resource) SetNRFreqRelation(v []NRFreqRelationSingle)`

SetNRFreqRelation sets NRFreqRelation field to given value.

### HasNRFreqRelation

`func (o *Resource) HasNRFreqRelation() bool`

HasNRFreqRelation returns a boolean if a field has been set.

### GetEUtranFreqRelation

`func (o *Resource) GetEUtranFreqRelation() []EUtranFreqRelationSingle`

GetEUtranFreqRelation returns the EUtranFreqRelation field if non-nil, zero value otherwise.

### GetEUtranFreqRelationOk

`func (o *Resource) GetEUtranFreqRelationOk() (*[]EUtranFreqRelationSingle, bool)`

GetEUtranFreqRelationOk returns a tuple with the EUtranFreqRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUtranFreqRelation

`func (o *Resource) SetEUtranFreqRelation(v []EUtranFreqRelationSingle)`

SetEUtranFreqRelation sets EUtranFreqRelation field to given value.

### HasEUtranFreqRelation

`func (o *Resource) HasEUtranFreqRelation() bool`

HasEUtranFreqRelation returns a boolean if a field has been set.

### GetNrOperatorCellDu

`func (o *Resource) GetNrOperatorCellDu() []NrOperatorCellDuSingle`

GetNrOperatorCellDu returns the NrOperatorCellDu field if non-nil, zero value otherwise.

### GetNrOperatorCellDuOk

`func (o *Resource) GetNrOperatorCellDuOk() (*[]NrOperatorCellDuSingle, bool)`

GetNrOperatorCellDuOk returns a tuple with the NrOperatorCellDu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOperatorCellDu

`func (o *Resource) SetNrOperatorCellDu(v []NrOperatorCellDuSingle)`

SetNrOperatorCellDu sets NrOperatorCellDu field to given value.

### HasNrOperatorCellDu

`func (o *Resource) HasNrOperatorCellDu() bool`

HasNrOperatorCellDu returns a boolean if a field has been set.

### GetCellLocalId

`func (o *Resource) GetCellLocalId() int32`

GetCellLocalId returns the CellLocalId field if non-nil, zero value otherwise.

### GetCellLocalIdOk

`func (o *Resource) GetCellLocalIdOk() (*int32, bool)`

GetCellLocalIdOk returns a tuple with the CellLocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellLocalId

`func (o *Resource) SetCellLocalId(v int32)`

SetCellLocalId sets CellLocalId field to given value.

### HasCellLocalId

`func (o *Resource) HasCellLocalId() bool`

HasCellLocalId returns a boolean if a field has been set.

### GetAdministrativeState

`func (o *Resource) GetAdministrativeState() AdministrativeState`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *Resource) GetAdministrativeStateOk() (*AdministrativeState, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *Resource) SetAdministrativeState(v AdministrativeState)`

SetAdministrativeState sets AdministrativeState field to given value.

### HasAdministrativeState

`func (o *Resource) HasAdministrativeState() bool`

HasAdministrativeState returns a boolean if a field has been set.

### GetPlmnInfoList

`func (o *Resource) GetPlmnInfoList() []PlmnInfo`

GetPlmnInfoList returns the PlmnInfoList field if non-nil, zero value otherwise.

### GetPlmnInfoListOk

`func (o *Resource) GetPlmnInfoListOk() (*[]PlmnInfo, bool)`

GetPlmnInfoListOk returns a tuple with the PlmnInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnInfoList

`func (o *Resource) SetPlmnInfoList(v []PlmnInfo)`

SetPlmnInfoList sets PlmnInfoList field to given value.

### HasPlmnInfoList

`func (o *Resource) HasPlmnInfoList() bool`

HasPlmnInfoList returns a boolean if a field has been set.

### GetNrTac

`func (o *Resource) GetNrTac() int32`

GetNrTac returns the NrTac field if non-nil, zero value otherwise.

### GetNrTacOk

`func (o *Resource) GetNrTacOk() (*int32, bool)`

GetNrTacOk returns a tuple with the NrTac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrTac

`func (o *Resource) SetNrTac(v int32)`

SetNrTac sets NrTac field to given value.

### HasNrTac

`func (o *Resource) HasNrTac() bool`

HasNrTac returns a boolean if a field has been set.

### GetCommonBeamformingFunction

`func (o *Resource) GetCommonBeamformingFunction() CommonBeamformingFunctionSingle`

GetCommonBeamformingFunction returns the CommonBeamformingFunction field if non-nil, zero value otherwise.

### GetCommonBeamformingFunctionOk

`func (o *Resource) GetCommonBeamformingFunctionOk() (*CommonBeamformingFunctionSingle, bool)`

GetCommonBeamformingFunctionOk returns a tuple with the CommonBeamformingFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonBeamformingFunction

`func (o *Resource) SetCommonBeamformingFunction(v CommonBeamformingFunctionSingle)`

SetCommonBeamformingFunction sets CommonBeamformingFunction field to given value.

### HasCommonBeamformingFunction

`func (o *Resource) HasCommonBeamformingFunction() bool`

HasCommonBeamformingFunction returns a boolean if a field has been set.

### GetBWPlist

`func (o *Resource) GetBWPlist() []string`

GetBWPlist returns the BWPlist field if non-nil, zero value otherwise.

### GetBWPlistOk

`func (o *Resource) GetBWPlistOk() (*[]string, bool)`

GetBWPlistOk returns a tuple with the BWPlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWPlist

`func (o *Resource) SetBWPlist(v []string)`

SetBWPlist sets BWPlist field to given value.

### HasBWPlist

`func (o *Resource) HasBWPlist() bool`

HasBWPlist returns a boolean if a field has been set.

### GetBeam

`func (o *Resource) GetBeam() []BeamSingle`

GetBeam returns the Beam field if non-nil, zero value otherwise.

### GetBeamOk

`func (o *Resource) GetBeamOk() (*[]BeamSingle, bool)`

GetBeamOk returns a tuple with the Beam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeam

`func (o *Resource) SetBeam(v []BeamSingle)`

SetBeam sets Beam field to given value.

### HasBeam

`func (o *Resource) HasBeam() bool`

HasBeam returns a boolean if a field has been set.

### GetRimRSSet

`func (o *Resource) GetRimRSSet() []RimRSSetSingle`

GetRimRSSet returns the RimRSSet field if non-nil, zero value otherwise.

### GetRimRSSetOk

`func (o *Resource) GetRimRSSetOk() (*[]RimRSSetSingle, bool)`

GetRimRSSetOk returns a tuple with the RimRSSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRimRSSet

`func (o *Resource) SetRimRSSet(v []RimRSSetSingle)`

SetRimRSSet sets RimRSSet field to given value.

### HasRimRSSet

`func (o *Resource) HasRimRSSet() bool`

HasRimRSSet returns a boolean if a field has been set.

### GetExternalNrCellCu

`func (o *Resource) GetExternalNrCellCu() []ExternalNrCellCuSingle`

GetExternalNrCellCu returns the ExternalNrCellCu field if non-nil, zero value otherwise.

### GetExternalNrCellCuOk

`func (o *Resource) GetExternalNrCellCuOk() (*[]ExternalNrCellCuSingle, bool)`

GetExternalNrCellCuOk returns a tuple with the ExternalNrCellCu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNrCellCu

`func (o *Resource) SetExternalNrCellCu(v []ExternalNrCellCuSingle)`

SetExternalNrCellCu sets ExternalNrCellCu field to given value.

### HasExternalNrCellCu

`func (o *Resource) HasExternalNrCellCu() bool`

HasExternalNrCellCu returns a boolean if a field has been set.

### GetExternalEUTranCell

`func (o *Resource) GetExternalEUTranCell() []ExternalEUTranCellSingle`

GetExternalEUTranCell returns the ExternalEUTranCell field if non-nil, zero value otherwise.

### GetExternalEUTranCellOk

`func (o *Resource) GetExternalEUTranCellOk() (*[]ExternalEUTranCellSingle, bool)`

GetExternalEUTranCellOk returns a tuple with the ExternalEUTranCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEUTranCell

`func (o *Resource) SetExternalEUTranCell(v []ExternalEUTranCellSingle)`

SetExternalEUTranCell sets ExternalEUTranCell field to given value.

### HasExternalEUTranCell

`func (o *Resource) HasExternalEUTranCell() bool`

HasExternalEUTranCell returns a boolean if a field has been set.

### GetExternalAmfFunction

`func (o *Resource) GetExternalAmfFunction() []ExternalAmfFunctionSingle`

GetExternalAmfFunction returns the ExternalAmfFunction field if non-nil, zero value otherwise.

### GetExternalAmfFunctionOk

`func (o *Resource) GetExternalAmfFunctionOk() (*[]ExternalAmfFunctionSingle, bool)`

GetExternalAmfFunctionOk returns a tuple with the ExternalAmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAmfFunction

`func (o *Resource) SetExternalAmfFunction(v []ExternalAmfFunctionSingle)`

SetExternalAmfFunction sets ExternalAmfFunction field to given value.

### HasExternalAmfFunction

`func (o *Resource) HasExternalAmfFunction() bool`

HasExternalAmfFunction returns a boolean if a field has been set.

### GetExternalNrfFunction

`func (o *Resource) GetExternalNrfFunction() []ExternalNrfFunctionSingle`

GetExternalNrfFunction returns the ExternalNrfFunction field if non-nil, zero value otherwise.

### GetExternalNrfFunctionOk

`func (o *Resource) GetExternalNrfFunctionOk() (*[]ExternalNrfFunctionSingle, bool)`

GetExternalNrfFunctionOk returns a tuple with the ExternalNrfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNrfFunction

`func (o *Resource) SetExternalNrfFunction(v []ExternalNrfFunctionSingle)`

SetExternalNrfFunction sets ExternalNrfFunction field to given value.

### HasExternalNrfFunction

`func (o *Resource) HasExternalNrfFunction() bool`

HasExternalNrfFunction returns a boolean if a field has been set.

### GetExternalNssfFunction

`func (o *Resource) GetExternalNssfFunction() []ExternalNssfFunctionSingle`

GetExternalNssfFunction returns the ExternalNssfFunction field if non-nil, zero value otherwise.

### GetExternalNssfFunctionOk

`func (o *Resource) GetExternalNssfFunctionOk() (*[]ExternalNssfFunctionSingle, bool)`

GetExternalNssfFunctionOk returns a tuple with the ExternalNssfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNssfFunction

`func (o *Resource) SetExternalNssfFunction(v []ExternalNssfFunctionSingle)`

SetExternalNssfFunction sets ExternalNssfFunction field to given value.

### HasExternalNssfFunction

`func (o *Resource) HasExternalNssfFunction() bool`

HasExternalNssfFunction returns a boolean if a field has been set.

### GetAmfSet

`func (o *Resource) GetAmfSet() []AmfSetSingle`

GetAmfSet returns the AmfSet field if non-nil, zero value otherwise.

### GetAmfSetOk

`func (o *Resource) GetAmfSetOk() (*[]AmfSetSingle, bool)`

GetAmfSetOk returns a tuple with the AmfSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSet

`func (o *Resource) SetAmfSet(v []AmfSetSingle)`

SetAmfSet sets AmfSet field to given value.

### HasAmfSet

`func (o *Resource) HasAmfSet() bool`

HasAmfSet returns a boolean if a field has been set.

### GetAmfRegion

`func (o *Resource) GetAmfRegion() []AmfRegionSingle`

GetAmfRegion returns the AmfRegion field if non-nil, zero value otherwise.

### GetAmfRegionOk

`func (o *Resource) GetAmfRegionOk() (*[]AmfRegionSingle, bool)`

GetAmfRegionOk returns a tuple with the AmfRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfRegion

`func (o *Resource) SetAmfRegion(v []AmfRegionSingle)`

SetAmfRegion sets AmfRegion field to given value.

### HasAmfRegion

`func (o *Resource) HasAmfRegion() bool`

HasAmfRegion returns a boolean if a field has been set.

### GetEcmConnectionInfo

`func (o *Resource) GetEcmConnectionInfo() []EcmConnectionInfoSingle`

GetEcmConnectionInfo returns the EcmConnectionInfo field if non-nil, zero value otherwise.

### GetEcmConnectionInfoOk

`func (o *Resource) GetEcmConnectionInfoOk() (*[]EcmConnectionInfoSingle, bool)`

GetEcmConnectionInfoOk returns a tuple with the EcmConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmConnectionInfo

`func (o *Resource) SetEcmConnectionInfo(v []EcmConnectionInfoSingle)`

SetEcmConnectionInfo sets EcmConnectionInfo field to given value.

### HasEcmConnectionInfo

`func (o *Resource) HasEcmConnectionInfo() bool`

HasEcmConnectionInfo returns a boolean if a field has been set.

### GetAmfFunction

`func (o *Resource) GetAmfFunction() []AmfFunctionSingle`

GetAmfFunction returns the AmfFunction field if non-nil, zero value otherwise.

### GetAmfFunctionOk

`func (o *Resource) GetAmfFunctionOk() (*[]AmfFunctionSingle, bool)`

GetAmfFunctionOk returns a tuple with the AmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfFunction

`func (o *Resource) SetAmfFunction(v []AmfFunctionSingle)`

SetAmfFunction sets AmfFunction field to given value.

### HasAmfFunction

`func (o *Resource) HasAmfFunction() bool`

HasAmfFunction returns a boolean if a field has been set.

### GetSmfFunction

`func (o *Resource) GetSmfFunction() []SmfFunctionSingle`

GetSmfFunction returns the SmfFunction field if non-nil, zero value otherwise.

### GetSmfFunctionOk

`func (o *Resource) GetSmfFunctionOk() (*[]SmfFunctionSingle, bool)`

GetSmfFunctionOk returns a tuple with the SmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfFunction

`func (o *Resource) SetSmfFunction(v []SmfFunctionSingle)`

SetSmfFunction sets SmfFunction field to given value.

### HasSmfFunction

`func (o *Resource) HasSmfFunction() bool`

HasSmfFunction returns a boolean if a field has been set.

### GetUpfFunction

`func (o *Resource) GetUpfFunction() []UpfFunctionSingle`

GetUpfFunction returns the UpfFunction field if non-nil, zero value otherwise.

### GetUpfFunctionOk

`func (o *Resource) GetUpfFunctionOk() (*[]UpfFunctionSingle, bool)`

GetUpfFunctionOk returns a tuple with the UpfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfFunction

`func (o *Resource) SetUpfFunction(v []UpfFunctionSingle)`

SetUpfFunction sets UpfFunction field to given value.

### HasUpfFunction

`func (o *Resource) HasUpfFunction() bool`

HasUpfFunction returns a boolean if a field has been set.

### GetN3iwfFunction

`func (o *Resource) GetN3iwfFunction() []N3iwfFunctionSingle`

GetN3iwfFunction returns the N3iwfFunction field if non-nil, zero value otherwise.

### GetN3iwfFunctionOk

`func (o *Resource) GetN3iwfFunctionOk() (*[]N3iwfFunctionSingle, bool)`

GetN3iwfFunctionOk returns a tuple with the N3iwfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3iwfFunction

`func (o *Resource) SetN3iwfFunction(v []N3iwfFunctionSingle)`

SetN3iwfFunction sets N3iwfFunction field to given value.

### HasN3iwfFunction

`func (o *Resource) HasN3iwfFunction() bool`

HasN3iwfFunction returns a boolean if a field has been set.

### GetPcfFunction

`func (o *Resource) GetPcfFunction() []PcfFunctionSingle`

GetPcfFunction returns the PcfFunction field if non-nil, zero value otherwise.

### GetPcfFunctionOk

`func (o *Resource) GetPcfFunctionOk() (*[]PcfFunctionSingle, bool)`

GetPcfFunctionOk returns a tuple with the PcfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFunction

`func (o *Resource) SetPcfFunction(v []PcfFunctionSingle)`

SetPcfFunction sets PcfFunction field to given value.

### HasPcfFunction

`func (o *Resource) HasPcfFunction() bool`

HasPcfFunction returns a boolean if a field has been set.

### GetAusfFunction

`func (o *Resource) GetAusfFunction() []AusfFunctionSingle`

GetAusfFunction returns the AusfFunction field if non-nil, zero value otherwise.

### GetAusfFunctionOk

`func (o *Resource) GetAusfFunctionOk() (*[]AusfFunctionSingle, bool)`

GetAusfFunctionOk returns a tuple with the AusfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfFunction

`func (o *Resource) SetAusfFunction(v []AusfFunctionSingle)`

SetAusfFunction sets AusfFunction field to given value.

### HasAusfFunction

`func (o *Resource) HasAusfFunction() bool`

HasAusfFunction returns a boolean if a field has been set.

### GetUdmFunction

`func (o *Resource) GetUdmFunction() []UdmFunctionSingle`

GetUdmFunction returns the UdmFunction field if non-nil, zero value otherwise.

### GetUdmFunctionOk

`func (o *Resource) GetUdmFunctionOk() (*[]UdmFunctionSingle, bool)`

GetUdmFunctionOk returns a tuple with the UdmFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmFunction

`func (o *Resource) SetUdmFunction(v []UdmFunctionSingle)`

SetUdmFunction sets UdmFunction field to given value.

### HasUdmFunction

`func (o *Resource) HasUdmFunction() bool`

HasUdmFunction returns a boolean if a field has been set.

### GetUdrFunction

`func (o *Resource) GetUdrFunction() []UdrFunctionSingle`

GetUdrFunction returns the UdrFunction field if non-nil, zero value otherwise.

### GetUdrFunctionOk

`func (o *Resource) GetUdrFunctionOk() (*[]UdrFunctionSingle, bool)`

GetUdrFunctionOk returns a tuple with the UdrFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrFunction

`func (o *Resource) SetUdrFunction(v []UdrFunctionSingle)`

SetUdrFunction sets UdrFunction field to given value.

### HasUdrFunction

`func (o *Resource) HasUdrFunction() bool`

HasUdrFunction returns a boolean if a field has been set.

### GetUdsfFunction

`func (o *Resource) GetUdsfFunction() []UdsfFunctionSingle`

GetUdsfFunction returns the UdsfFunction field if non-nil, zero value otherwise.

### GetUdsfFunctionOk

`func (o *Resource) GetUdsfFunctionOk() (*[]UdsfFunctionSingle, bool)`

GetUdsfFunctionOk returns a tuple with the UdsfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsfFunction

`func (o *Resource) SetUdsfFunction(v []UdsfFunctionSingle)`

SetUdsfFunction sets UdsfFunction field to given value.

### HasUdsfFunction

`func (o *Resource) HasUdsfFunction() bool`

HasUdsfFunction returns a boolean if a field has been set.

### GetNrfFunction

`func (o *Resource) GetNrfFunction() []NrfFunctionSingle`

GetNrfFunction returns the NrfFunction field if non-nil, zero value otherwise.

### GetNrfFunctionOk

`func (o *Resource) GetNrfFunctionOk() (*[]NrfFunctionSingle, bool)`

GetNrfFunctionOk returns a tuple with the NrfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfFunction

`func (o *Resource) SetNrfFunction(v []NrfFunctionSingle)`

SetNrfFunction sets NrfFunction field to given value.

### HasNrfFunction

`func (o *Resource) HasNrfFunction() bool`

HasNrfFunction returns a boolean if a field has been set.

### GetNssfFunction

`func (o *Resource) GetNssfFunction() []NssfFunctionSingle`

GetNssfFunction returns the NssfFunction field if non-nil, zero value otherwise.

### GetNssfFunctionOk

`func (o *Resource) GetNssfFunctionOk() (*[]NssfFunctionSingle, bool)`

GetNssfFunctionOk returns a tuple with the NssfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssfFunction

`func (o *Resource) SetNssfFunction(v []NssfFunctionSingle)`

SetNssfFunction sets NssfFunction field to given value.

### HasNssfFunction

`func (o *Resource) HasNssfFunction() bool`

HasNssfFunction returns a boolean if a field has been set.

### GetSmsfFunction

`func (o *Resource) GetSmsfFunction() []SmsfFunctionSingle`

GetSmsfFunction returns the SmsfFunction field if non-nil, zero value otherwise.

### GetSmsfFunctionOk

`func (o *Resource) GetSmsfFunctionOk() (*[]SmsfFunctionSingle, bool)`

GetSmsfFunctionOk returns a tuple with the SmsfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfFunction

`func (o *Resource) SetSmsfFunction(v []SmsfFunctionSingle)`

SetSmsfFunction sets SmsfFunction field to given value.

### HasSmsfFunction

`func (o *Resource) HasSmsfFunction() bool`

HasSmsfFunction returns a boolean if a field has been set.

### GetLmfFunction

`func (o *Resource) GetLmfFunction() []LmfFunctionSingle`

GetLmfFunction returns the LmfFunction field if non-nil, zero value otherwise.

### GetLmfFunctionOk

`func (o *Resource) GetLmfFunctionOk() (*[]LmfFunctionSingle, bool)`

GetLmfFunctionOk returns a tuple with the LmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfFunction

`func (o *Resource) SetLmfFunction(v []LmfFunctionSingle)`

SetLmfFunction sets LmfFunction field to given value.

### HasLmfFunction

`func (o *Resource) HasLmfFunction() bool`

HasLmfFunction returns a boolean if a field has been set.

### GetNgeirFunction

`func (o *Resource) GetNgeirFunction() []NgeirFunctionSingle`

GetNgeirFunction returns the NgeirFunction field if non-nil, zero value otherwise.

### GetNgeirFunctionOk

`func (o *Resource) GetNgeirFunctionOk() (*[]NgeirFunctionSingle, bool)`

GetNgeirFunctionOk returns a tuple with the NgeirFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgeirFunction

`func (o *Resource) SetNgeirFunction(v []NgeirFunctionSingle)`

SetNgeirFunction sets NgeirFunction field to given value.

### HasNgeirFunction

`func (o *Resource) HasNgeirFunction() bool`

HasNgeirFunction returns a boolean if a field has been set.

### GetSeppFunction

`func (o *Resource) GetSeppFunction() []SeppFunctionSingle`

GetSeppFunction returns the SeppFunction field if non-nil, zero value otherwise.

### GetSeppFunctionOk

`func (o *Resource) GetSeppFunctionOk() (*[]SeppFunctionSingle, bool)`

GetSeppFunctionOk returns a tuple with the SeppFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppFunction

`func (o *Resource) SetSeppFunction(v []SeppFunctionSingle)`

SetSeppFunction sets SeppFunction field to given value.

### HasSeppFunction

`func (o *Resource) HasSeppFunction() bool`

HasSeppFunction returns a boolean if a field has been set.

### GetNwdafFunction

`func (o *Resource) GetNwdafFunction() []NwdafFunctionSingle`

GetNwdafFunction returns the NwdafFunction field if non-nil, zero value otherwise.

### GetNwdafFunctionOk

`func (o *Resource) GetNwdafFunctionOk() (*[]NwdafFunctionSingle, bool)`

GetNwdafFunctionOk returns a tuple with the NwdafFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafFunction

`func (o *Resource) SetNwdafFunction(v []NwdafFunctionSingle)`

SetNwdafFunction sets NwdafFunction field to given value.

### HasNwdafFunction

`func (o *Resource) HasNwdafFunction() bool`

HasNwdafFunction returns a boolean if a field has been set.

### GetScpFunction

`func (o *Resource) GetScpFunction() []ScpFunctionSingle`

GetScpFunction returns the ScpFunction field if non-nil, zero value otherwise.

### GetScpFunctionOk

`func (o *Resource) GetScpFunctionOk() (*[]ScpFunctionSingle, bool)`

GetScpFunctionOk returns a tuple with the ScpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpFunction

`func (o *Resource) SetScpFunction(v []ScpFunctionSingle)`

SetScpFunction sets ScpFunction field to given value.

### HasScpFunction

`func (o *Resource) HasScpFunction() bool`

HasScpFunction returns a boolean if a field has been set.

### GetNefFunction

`func (o *Resource) GetNefFunction() []NefFunctionSingle`

GetNefFunction returns the NefFunction field if non-nil, zero value otherwise.

### GetNefFunctionOk

`func (o *Resource) GetNefFunctionOk() (*[]NefFunctionSingle, bool)`

GetNefFunctionOk returns a tuple with the NefFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefFunction

`func (o *Resource) SetNefFunction(v []NefFunctionSingle)`

SetNefFunction sets NefFunction field to given value.

### HasNefFunction

`func (o *Resource) HasNefFunction() bool`

HasNefFunction returns a boolean if a field has been set.

### GetEASDFFunction

`func (o *Resource) GetEASDFFunction() []EASDFFunctionSingle`

GetEASDFFunction returns the EASDFFunction field if non-nil, zero value otherwise.

### GetEASDFFunctionOk

`func (o *Resource) GetEASDFFunctionOk() (*[]EASDFFunctionSingle, bool)`

GetEASDFFunctionOk returns a tuple with the EASDFFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEASDFFunction

`func (o *Resource) SetEASDFFunction(v []EASDFFunctionSingle)`

SetEASDFFunction sets EASDFFunction field to given value.

### HasEASDFFunction

`func (o *Resource) HasEASDFFunction() bool`

HasEASDFFunction returns a boolean if a field has been set.

### GetEPN2

`func (o *Resource) GetEPN2() []EPN2Single`

GetEPN2 returns the EPN2 field if non-nil, zero value otherwise.

### GetEPN2Ok

`func (o *Resource) GetEPN2Ok() (*[]EPN2Single, bool)`

GetEPN2Ok returns a tuple with the EPN2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN2

`func (o *Resource) SetEPN2(v []EPN2Single)`

SetEPN2 sets EPN2 field to given value.

### HasEPN2

`func (o *Resource) HasEPN2() bool`

HasEPN2 returns a boolean if a field has been set.

### GetEPN8

`func (o *Resource) GetEPN8() []EPN8Single`

GetEPN8 returns the EPN8 field if non-nil, zero value otherwise.

### GetEPN8Ok

`func (o *Resource) GetEPN8Ok() (*[]EPN8Single, bool)`

GetEPN8Ok returns a tuple with the EPN8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN8

`func (o *Resource) SetEPN8(v []EPN8Single)`

SetEPN8 sets EPN8 field to given value.

### HasEPN8

`func (o *Resource) HasEPN8() bool`

HasEPN8 returns a boolean if a field has been set.

### GetEPN11

`func (o *Resource) GetEPN11() []EPN11Single`

GetEPN11 returns the EPN11 field if non-nil, zero value otherwise.

### GetEPN11Ok

`func (o *Resource) GetEPN11Ok() (*[]EPN11Single, bool)`

GetEPN11Ok returns a tuple with the EPN11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN11

`func (o *Resource) SetEPN11(v []EPN11Single)`

SetEPN11 sets EPN11 field to given value.

### HasEPN11

`func (o *Resource) HasEPN11() bool`

HasEPN11 returns a boolean if a field has been set.

### GetEPN12

`func (o *Resource) GetEPN12() []EPN12Single`

GetEPN12 returns the EPN12 field if non-nil, zero value otherwise.

### GetEPN12Ok

`func (o *Resource) GetEPN12Ok() (*[]EPN12Single, bool)`

GetEPN12Ok returns a tuple with the EPN12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN12

`func (o *Resource) SetEPN12(v []EPN12Single)`

SetEPN12 sets EPN12 field to given value.

### HasEPN12

`func (o *Resource) HasEPN12() bool`

HasEPN12 returns a boolean if a field has been set.

### GetEPN14

`func (o *Resource) GetEPN14() []EPN14Single`

GetEPN14 returns the EPN14 field if non-nil, zero value otherwise.

### GetEPN14Ok

`func (o *Resource) GetEPN14Ok() (*[]EPN14Single, bool)`

GetEPN14Ok returns a tuple with the EPN14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN14

`func (o *Resource) SetEPN14(v []EPN14Single)`

SetEPN14 sets EPN14 field to given value.

### HasEPN14

`func (o *Resource) HasEPN14() bool`

HasEPN14 returns a boolean if a field has been set.

### GetEPN15

`func (o *Resource) GetEPN15() []EPN15Single`

GetEPN15 returns the EPN15 field if non-nil, zero value otherwise.

### GetEPN15Ok

`func (o *Resource) GetEPN15Ok() (*[]EPN15Single, bool)`

GetEPN15Ok returns a tuple with the EPN15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN15

`func (o *Resource) SetEPN15(v []EPN15Single)`

SetEPN15 sets EPN15 field to given value.

### HasEPN15

`func (o *Resource) HasEPN15() bool`

HasEPN15 returns a boolean if a field has been set.

### GetEPN17

`func (o *Resource) GetEPN17() []EPN17Single`

GetEPN17 returns the EPN17 field if non-nil, zero value otherwise.

### GetEPN17Ok

`func (o *Resource) GetEPN17Ok() (*[]EPN17Single, bool)`

GetEPN17Ok returns a tuple with the EPN17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN17

`func (o *Resource) SetEPN17(v []EPN17Single)`

SetEPN17 sets EPN17 field to given value.

### HasEPN17

`func (o *Resource) HasEPN17() bool`

HasEPN17 returns a boolean if a field has been set.

### GetEPN20

`func (o *Resource) GetEPN20() []EPN20Single`

GetEPN20 returns the EPN20 field if non-nil, zero value otherwise.

### GetEPN20Ok

`func (o *Resource) GetEPN20Ok() (*[]EPN20Single, bool)`

GetEPN20Ok returns a tuple with the EPN20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN20

`func (o *Resource) SetEPN20(v []EPN20Single)`

SetEPN20 sets EPN20 field to given value.

### HasEPN20

`func (o *Resource) HasEPN20() bool`

HasEPN20 returns a boolean if a field has been set.

### GetEPN22

`func (o *Resource) GetEPN22() []EPN22Single`

GetEPN22 returns the EPN22 field if non-nil, zero value otherwise.

### GetEPN22Ok

`func (o *Resource) GetEPN22Ok() (*[]EPN22Single, bool)`

GetEPN22Ok returns a tuple with the EPN22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN22

`func (o *Resource) SetEPN22(v []EPN22Single)`

SetEPN22 sets EPN22 field to given value.

### HasEPN22

`func (o *Resource) HasEPN22() bool`

HasEPN22 returns a boolean if a field has been set.

### GetEPN26

`func (o *Resource) GetEPN26() []EPN26Single`

GetEPN26 returns the EPN26 field if non-nil, zero value otherwise.

### GetEPN26Ok

`func (o *Resource) GetEPN26Ok() (*[]EPN26Single, bool)`

GetEPN26Ok returns a tuple with the EPN26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN26

`func (o *Resource) SetEPN26(v []EPN26Single)`

SetEPN26 sets EPN26 field to given value.

### HasEPN26

`func (o *Resource) HasEPN26() bool`

HasEPN26 returns a boolean if a field has been set.

### GetEP_NLS

`func (o *Resource) GetEP_NLS() []EPNLSSingle`

GetEP_NLS returns the EP_NLS field if non-nil, zero value otherwise.

### GetEP_NLSOk

`func (o *Resource) GetEP_NLSOk() (*[]EPNLSSingle, bool)`

GetEP_NLSOk returns a tuple with the EP_NLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEP_NLS

`func (o *Resource) SetEP_NLS(v []EPNLSSingle)`

SetEP_NLS sets EP_NLS field to given value.

### HasEP_NLS

`func (o *Resource) HasEP_NLS() bool`

HasEP_NLS returns a boolean if a field has been set.

### GetEP_NLG

`func (o *Resource) GetEP_NLG() []EPNLGSingle`

GetEP_NLG returns the EP_NLG field if non-nil, zero value otherwise.

### GetEP_NLGOk

`func (o *Resource) GetEP_NLGOk() (*[]EPNLGSingle, bool)`

GetEP_NLGOk returns a tuple with the EP_NLG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEP_NLG

`func (o *Resource) SetEP_NLG(v []EPNLGSingle)`

SetEP_NLG sets EP_NLG field to given value.

### HasEP_NLG

`func (o *Resource) HasEP_NLG() bool`

HasEP_NLG returns a boolean if a field has been set.

### GetEPN4

`func (o *Resource) GetEPN4() []EPN4Single`

GetEPN4 returns the EPN4 field if non-nil, zero value otherwise.

### GetEPN4Ok

`func (o *Resource) GetEPN4Ok() (*[]EPN4Single, bool)`

GetEPN4Ok returns a tuple with the EPN4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN4

`func (o *Resource) SetEPN4(v []EPN4Single)`

SetEPN4 sets EPN4 field to given value.

### HasEPN4

`func (o *Resource) HasEPN4() bool`

HasEPN4 returns a boolean if a field has been set.

### GetEPN7

`func (o *Resource) GetEPN7() []EPN7Single`

GetEPN7 returns the EPN7 field if non-nil, zero value otherwise.

### GetEPN7Ok

`func (o *Resource) GetEPN7Ok() (*[]EPN7Single, bool)`

GetEPN7Ok returns a tuple with the EPN7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN7

`func (o *Resource) SetEPN7(v []EPN7Single)`

SetEPN7 sets EPN7 field to given value.

### HasEPN7

`func (o *Resource) HasEPN7() bool`

HasEPN7 returns a boolean if a field has been set.

### GetEPN10

`func (o *Resource) GetEPN10() []EPN10Single`

GetEPN10 returns the EPN10 field if non-nil, zero value otherwise.

### GetEPN10Ok

`func (o *Resource) GetEPN10Ok() (*[]EPN10Single, bool)`

GetEPN10Ok returns a tuple with the EPN10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN10

`func (o *Resource) SetEPN10(v []EPN10Single)`

SetEPN10 sets EPN10 field to given value.

### HasEPN10

`func (o *Resource) HasEPN10() bool`

HasEPN10 returns a boolean if a field has been set.

### GetEPN16

`func (o *Resource) GetEPN16() []EPN16Single`

GetEPN16 returns the EPN16 field if non-nil, zero value otherwise.

### GetEPN16Ok

`func (o *Resource) GetEPN16Ok() (*[]EPN16Single, bool)`

GetEPN16Ok returns a tuple with the EPN16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN16

`func (o *Resource) SetEPN16(v []EPN16Single)`

SetEPN16 sets EPN16 field to given value.

### HasEPN16

`func (o *Resource) HasEPN16() bool`

HasEPN16 returns a boolean if a field has been set.

### GetEPS5C

`func (o *Resource) GetEPS5C() []EPS5CSingle`

GetEPS5C returns the EPS5C field if non-nil, zero value otherwise.

### GetEPS5COk

`func (o *Resource) GetEPS5COk() (*[]EPS5CSingle, bool)`

GetEPS5COk returns a tuple with the EPS5C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPS5C

`func (o *Resource) SetEPS5C(v []EPS5CSingle)`

SetEPS5C sets EPS5C field to given value.

### HasEPS5C

`func (o *Resource) HasEPS5C() bool`

HasEPS5C returns a boolean if a field has been set.

### GetFiveQiDscpMappingSet

`func (o *Resource) GetFiveQiDscpMappingSet() FiveQiDscpMappingSetSingle`

GetFiveQiDscpMappingSet returns the FiveQiDscpMappingSet field if non-nil, zero value otherwise.

### GetFiveQiDscpMappingSetOk

`func (o *Resource) GetFiveQiDscpMappingSetOk() (*FiveQiDscpMappingSetSingle, bool)`

GetFiveQiDscpMappingSetOk returns a tuple with the FiveQiDscpMappingSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveQiDscpMappingSet

`func (o *Resource) SetFiveQiDscpMappingSet(v FiveQiDscpMappingSetSingle)`

SetFiveQiDscpMappingSet sets FiveQiDscpMappingSet field to given value.

### HasFiveQiDscpMappingSet

`func (o *Resource) HasFiveQiDscpMappingSet() bool`

HasFiveQiDscpMappingSet returns a boolean if a field has been set.

### GetGtpUPathQoSMonitoringControl

`func (o *Resource) GetGtpUPathQoSMonitoringControl() GtpUPathQoSMonitoringControlSingle`

GetGtpUPathQoSMonitoringControl returns the GtpUPathQoSMonitoringControl field if non-nil, zero value otherwise.

### GetGtpUPathQoSMonitoringControlOk

`func (o *Resource) GetGtpUPathQoSMonitoringControlOk() (*GtpUPathQoSMonitoringControlSingle, bool)`

GetGtpUPathQoSMonitoringControlOk returns a tuple with the GtpUPathQoSMonitoringControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtpUPathQoSMonitoringControl

`func (o *Resource) SetGtpUPathQoSMonitoringControl(v GtpUPathQoSMonitoringControlSingle)`

SetGtpUPathQoSMonitoringControl sets GtpUPathQoSMonitoringControl field to given value.

### HasGtpUPathQoSMonitoringControl

`func (o *Resource) HasGtpUPathQoSMonitoringControl() bool`

HasGtpUPathQoSMonitoringControl returns a boolean if a field has been set.

### GetQFQoSMonitoringControl

`func (o *Resource) GetQFQoSMonitoringControl() QFQoSMonitoringControlSingle`

GetQFQoSMonitoringControl returns the QFQoSMonitoringControl field if non-nil, zero value otherwise.

### GetQFQoSMonitoringControlOk

`func (o *Resource) GetQFQoSMonitoringControlOk() (*QFQoSMonitoringControlSingle, bool)`

GetQFQoSMonitoringControlOk returns a tuple with the QFQoSMonitoringControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQFQoSMonitoringControl

`func (o *Resource) SetQFQoSMonitoringControl(v QFQoSMonitoringControlSingle)`

SetQFQoSMonitoringControl sets QFQoSMonitoringControl field to given value.

### HasQFQoSMonitoringControl

`func (o *Resource) HasQFQoSMonitoringControl() bool`

HasQFQoSMonitoringControl returns a boolean if a field has been set.

### GetPredefinedPccRuleSet

`func (o *Resource) GetPredefinedPccRuleSet() PredefinedPccRuleSetSingle`

GetPredefinedPccRuleSet returns the PredefinedPccRuleSet field if non-nil, zero value otherwise.

### GetPredefinedPccRuleSetOk

`func (o *Resource) GetPredefinedPccRuleSetOk() (*PredefinedPccRuleSetSingle, bool)`

GetPredefinedPccRuleSetOk returns a tuple with the PredefinedPccRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedPccRuleSet

`func (o *Resource) SetPredefinedPccRuleSet(v PredefinedPccRuleSetSingle)`

SetPredefinedPccRuleSet sets PredefinedPccRuleSet field to given value.

### HasPredefinedPccRuleSet

`func (o *Resource) HasPredefinedPccRuleSet() bool`

HasPredefinedPccRuleSet returns a boolean if a field has been set.

### GetEPN3

`func (o *Resource) GetEPN3() []EPN3Single`

GetEPN3 returns the EPN3 field if non-nil, zero value otherwise.

### GetEPN3Ok

`func (o *Resource) GetEPN3Ok() (*[]EPN3Single, bool)`

GetEPN3Ok returns a tuple with the EPN3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN3

`func (o *Resource) SetEPN3(v []EPN3Single)`

SetEPN3 sets EPN3 field to given value.

### HasEPN3

`func (o *Resource) HasEPN3() bool`

HasEPN3 returns a boolean if a field has been set.

### GetEPN6

`func (o *Resource) GetEPN6() []EPN6Single`

GetEPN6 returns the EPN6 field if non-nil, zero value otherwise.

### GetEPN6Ok

`func (o *Resource) GetEPN6Ok() (*[]EPN6Single, bool)`

GetEPN6Ok returns a tuple with the EPN6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN6

`func (o *Resource) SetEPN6(v []EPN6Single)`

SetEPN6 sets EPN6 field to given value.

### HasEPN6

`func (o *Resource) HasEPN6() bool`

HasEPN6 returns a boolean if a field has been set.

### GetEPN9

`func (o *Resource) GetEPN9() []EPN9Single`

GetEPN9 returns the EPN9 field if non-nil, zero value otherwise.

### GetEPN9Ok

`func (o *Resource) GetEPN9Ok() (*[]EPN9Single, bool)`

GetEPN9Ok returns a tuple with the EPN9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN9

`func (o *Resource) SetEPN9(v []EPN9Single)`

SetEPN9 sets EPN9 field to given value.

### HasEPN9

`func (o *Resource) HasEPN9() bool`

HasEPN9 returns a boolean if a field has been set.

### GetEPS5U

`func (o *Resource) GetEPS5U() []EPS5USingle`

GetEPS5U returns the EPS5U field if non-nil, zero value otherwise.

### GetEPS5UOk

`func (o *Resource) GetEPS5UOk() (*[]EPS5USingle, bool)`

GetEPS5UOk returns a tuple with the EPS5U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPS5U

`func (o *Resource) SetEPS5U(v []EPS5USingle)`

SetEPS5U sets EPS5U field to given value.

### HasEPS5U

`func (o *Resource) HasEPS5U() bool`

HasEPS5U returns a boolean if a field has been set.

### GetEPN5

`func (o *Resource) GetEPN5() []EPN5Single`

GetEPN5 returns the EPN5 field if non-nil, zero value otherwise.

### GetEPN5Ok

`func (o *Resource) GetEPN5Ok() (*[]EPN5Single, bool)`

GetEPN5Ok returns a tuple with the EPN5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN5

`func (o *Resource) SetEPN5(v []EPN5Single)`

SetEPN5 sets EPN5 field to given value.

### HasEPN5

`func (o *Resource) HasEPN5() bool`

HasEPN5 returns a boolean if a field has been set.

### GetEPRx

`func (o *Resource) GetEPRx() []EPRxSingle`

GetEPRx returns the EPRx field if non-nil, zero value otherwise.

### GetEPRxOk

`func (o *Resource) GetEPRxOk() (*[]EPRxSingle, bool)`

GetEPRxOk returns a tuple with the EPRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPRx

`func (o *Resource) SetEPRx(v []EPRxSingle)`

SetEPRx sets EPRx field to given value.

### HasEPRx

`func (o *Resource) HasEPRx() bool`

HasEPRx returns a boolean if a field has been set.

### GetEPN13

`func (o *Resource) GetEPN13() []EPN13Single`

GetEPN13 returns the EPN13 field if non-nil, zero value otherwise.

### GetEPN13Ok

`func (o *Resource) GetEPN13Ok() (*[]EPN13Single, bool)`

GetEPN13Ok returns a tuple with the EPN13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN13

`func (o *Resource) SetEPN13(v []EPN13Single)`

SetEPN13 sets EPN13 field to given value.

### HasEPN13

`func (o *Resource) HasEPN13() bool`

HasEPN13 returns a boolean if a field has been set.

### GetEPN27

`func (o *Resource) GetEPN27() []EPN27Single`

GetEPN27 returns the EPN27 field if non-nil, zero value otherwise.

### GetEPN27Ok

`func (o *Resource) GetEPN27Ok() (*[]EPN27Single, bool)`

GetEPN27Ok returns a tuple with the EPN27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN27

`func (o *Resource) SetEPN27(v []EPN27Single)`

SetEPN27 sets EPN27 field to given value.

### HasEPN27

`func (o *Resource) HasEPN27() bool`

HasEPN27 returns a boolean if a field has been set.

### GetEPN31

`func (o *Resource) GetEPN31() []EPN31Single`

GetEPN31 returns the EPN31 field if non-nil, zero value otherwise.

### GetEPN31Ok

`func (o *Resource) GetEPN31Ok() (*[]EPN31Single, bool)`

GetEPN31Ok returns a tuple with the EPN31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN31

`func (o *Resource) SetEPN31(v []EPN31Single)`

SetEPN31 sets EPN31 field to given value.

### HasEPN31

`func (o *Resource) HasEPN31() bool`

HasEPN31 returns a boolean if a field has been set.

### GetEPN21

`func (o *Resource) GetEPN21() []EPN21Single`

GetEPN21 returns the EPN21 field if non-nil, zero value otherwise.

### GetEPN21Ok

`func (o *Resource) GetEPN21Ok() (*[]EPN21Single, bool)`

GetEPN21Ok returns a tuple with the EPN21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN21

`func (o *Resource) SetEPN21(v []EPN21Single)`

SetEPN21 sets EPN21 field to given value.

### HasEPN21

`func (o *Resource) HasEPN21() bool`

HasEPN21 returns a boolean if a field has been set.

### GetEP_MAP_SMSC

`func (o *Resource) GetEP_MAP_SMSC() []EPMAPSMSCSingle`

GetEP_MAP_SMSC returns the EP_MAP_SMSC field if non-nil, zero value otherwise.

### GetEP_MAP_SMSCOk

`func (o *Resource) GetEP_MAP_SMSCOk() (*[]EPMAPSMSCSingle, bool)`

GetEP_MAP_SMSCOk returns a tuple with the EP_MAP_SMSC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEP_MAP_SMSC

`func (o *Resource) SetEP_MAP_SMSC(v []EPMAPSMSCSingle)`

SetEP_MAP_SMSC sets EP_MAP_SMSC field to given value.

### HasEP_MAP_SMSC

`func (o *Resource) HasEP_MAP_SMSC() bool`

HasEP_MAP_SMSC returns a boolean if a field has been set.

### GetEPN32

`func (o *Resource) GetEPN32() []EPN32Single`

GetEPN32 returns the EPN32 field if non-nil, zero value otherwise.

### GetEPN32Ok

`func (o *Resource) GetEPN32Ok() (*[]EPN32Single, bool)`

GetEPN32Ok returns a tuple with the EPN32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN32

`func (o *Resource) SetEPN32(v []EPN32Single)`

SetEPN32 sets EPN32 field to given value.

### HasEPN32

`func (o *Resource) HasEPN32() bool`

HasEPN32 returns a boolean if a field has been set.

### GetEPN33

`func (o *Resource) GetEPN33() []EPN33Single`

GetEPN33 returns the EPN33 field if non-nil, zero value otherwise.

### GetEPN33Ok

`func (o *Resource) GetEPN33Ok() (*[]EPN33Single, bool)`

GetEPN33Ok returns a tuple with the EPN33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN33

`func (o *Resource) SetEPN33(v []EPN33Single)`

SetEPN33 sets EPN33 field to given value.

### HasEPN33

`func (o *Resource) HasEPN33() bool`

HasEPN33 returns a boolean if a field has been set.

### GetEPN60

`func (o *Resource) GetEPN60() []EPN60Single`

GetEPN60 returns the EPN60 field if non-nil, zero value otherwise.

### GetEPN60Ok

`func (o *Resource) GetEPN60Ok() (*[]EPN60Single, bool)`

GetEPN60Ok returns a tuple with the EPN60 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN60

`func (o *Resource) SetEPN60(v []EPN60Single)`

SetEPN60 sets EPN60 field to given value.

### HasEPN60

`func (o *Resource) HasEPN60() bool`

HasEPN60 returns a boolean if a field has been set.

### GetEPNpc4

`func (o *Resource) GetEPNpc4() []EPNpc4Single`

GetEPNpc4 returns the EPNpc4 field if non-nil, zero value otherwise.

### GetEPNpc4Ok

`func (o *Resource) GetEPNpc4Ok() (*[]EPNpc4Single, bool)`

GetEPNpc4Ok returns a tuple with the EPNpc4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc4

`func (o *Resource) SetEPNpc4(v []EPNpc4Single)`

SetEPNpc4 sets EPNpc4 field to given value.

### HasEPNpc4

`func (o *Resource) HasEPNpc4() bool`

HasEPNpc4 returns a boolean if a field has been set.

### GetEPNpc6

`func (o *Resource) GetEPNpc6() []EPNpc6Single`

GetEPNpc6 returns the EPNpc6 field if non-nil, zero value otherwise.

### GetEPNpc6Ok

`func (o *Resource) GetEPNpc6Ok() (*[]EPNpc6Single, bool)`

GetEPNpc6Ok returns a tuple with the EPNpc6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc6

`func (o *Resource) SetEPNpc6(v []EPNpc6Single)`

SetEPNpc6 sets EPNpc6 field to given value.

### HasEPNpc6

`func (o *Resource) HasEPNpc6() bool`

HasEPNpc6 returns a boolean if a field has been set.

### GetEPNpc7

`func (o *Resource) GetEPNpc7() []EPNpc7Single`

GetEPNpc7 returns the EPNpc7 field if non-nil, zero value otherwise.

### GetEPNpc7Ok

`func (o *Resource) GetEPNpc7Ok() (*[]EPNpc7Single, bool)`

GetEPNpc7Ok returns a tuple with the EPNpc7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc7

`func (o *Resource) SetEPNpc7(v []EPNpc7Single)`

SetEPNpc7 sets EPNpc7 field to given value.

### HasEPNpc7

`func (o *Resource) HasEPNpc7() bool`

HasEPNpc7 returns a boolean if a field has been set.

### GetEPNpc8

`func (o *Resource) GetEPNpc8() []EPNpc8Single`

GetEPNpc8 returns the EPNpc8 field if non-nil, zero value otherwise.

### GetEPNpc8Ok

`func (o *Resource) GetEPNpc8Ok() (*[]EPNpc8Single, bool)`

GetEPNpc8Ok returns a tuple with the EPNpc8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc8

`func (o *Resource) SetEPNpc8(v []EPNpc8Single)`

SetEPNpc8 sets EPNpc8 field to given value.

### HasEPNpc8

`func (o *Resource) HasEPNpc8() bool`

HasEPNpc8 returns a boolean if a field has been set.

### GetEPN88

`func (o *Resource) GetEPN88() []EPN88Single`

GetEPN88 returns the EPN88 field if non-nil, zero value otherwise.

### GetEPN88Ok

`func (o *Resource) GetEPN88Ok() (*[]EPN88Single, bool)`

GetEPN88Ok returns a tuple with the EPN88 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN88

`func (o *Resource) SetEPN88(v []EPN88Single)`

SetEPN88 sets EPN88 field to given value.

### HasEPN88

`func (o *Resource) HasEPN88() bool`

HasEPN88 returns a boolean if a field has been set.

### GetNetworkSlice

`func (o *Resource) GetNetworkSlice() []NetworkSliceSingle`

GetNetworkSlice returns the NetworkSlice field if non-nil, zero value otherwise.

### GetNetworkSliceOk

`func (o *Resource) GetNetworkSliceOk() (*[]NetworkSliceSingle, bool)`

GetNetworkSliceOk returns a tuple with the NetworkSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSlice

`func (o *Resource) SetNetworkSlice(v []NetworkSliceSingle)`

SetNetworkSlice sets NetworkSlice field to given value.

### HasNetworkSlice

`func (o *Resource) HasNetworkSlice() bool`

HasNetworkSlice returns a boolean if a field has been set.

### GetNetworkSliceSubnet

`func (o *Resource) GetNetworkSliceSubnet() []NetworkSliceSubnetSingle`

GetNetworkSliceSubnet returns the NetworkSliceSubnet field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetOk

`func (o *Resource) GetNetworkSliceSubnetOk() (*[]NetworkSliceSubnetSingle, bool)`

GetNetworkSliceSubnetOk returns a tuple with the NetworkSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnet

`func (o *Resource) SetNetworkSliceSubnet(v []NetworkSliceSubnetSingle)`

SetNetworkSliceSubnet sets NetworkSliceSubnet field to given value.

### HasNetworkSliceSubnet

`func (o *Resource) HasNetworkSliceSubnet() bool`

HasNetworkSliceSubnet returns a boolean if a field has been set.

### GetEPTransport

`func (o *Resource) GetEPTransport() []EPTransportSingle`

GetEPTransport returns the EPTransport field if non-nil, zero value otherwise.

### GetEPTransportOk

`func (o *Resource) GetEPTransportOk() (*[]EPTransportSingle, bool)`

GetEPTransportOk returns a tuple with the EPTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPTransport

`func (o *Resource) SetEPTransport(v []EPTransportSingle)`

SetEPTransport sets EPTransport field to given value.

### HasEPTransport

`func (o *Resource) HasEPTransport() bool`

HasEPTransport returns a boolean if a field has been set.

### GetNetworkSliceSubnetProviderCapabilities

`func (o *Resource) GetNetworkSliceSubnetProviderCapabilities() []NetworkSliceSubnetProviderCapabilitiesSingle`

GetNetworkSliceSubnetProviderCapabilities returns the NetworkSliceSubnetProviderCapabilities field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetProviderCapabilitiesOk

`func (o *Resource) GetNetworkSliceSubnetProviderCapabilitiesOk() (*[]NetworkSliceSubnetProviderCapabilitiesSingle, bool)`

GetNetworkSliceSubnetProviderCapabilitiesOk returns a tuple with the NetworkSliceSubnetProviderCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnetProviderCapabilities

`func (o *Resource) SetNetworkSliceSubnetProviderCapabilities(v []NetworkSliceSubnetProviderCapabilitiesSingle)`

SetNetworkSliceSubnetProviderCapabilities sets NetworkSliceSubnetProviderCapabilities field to given value.

### HasNetworkSliceSubnetProviderCapabilities

`func (o *Resource) HasNetworkSliceSubnetProviderCapabilities() bool`

HasNetworkSliceSubnetProviderCapabilities returns a boolean if a field has been set.

### GetFeasibilityCheckAndReservationJob

`func (o *Resource) GetFeasibilityCheckAndReservationJob() []FeasibilityCheckAndReservationJobSingle`

GetFeasibilityCheckAndReservationJob returns the FeasibilityCheckAndReservationJob field if non-nil, zero value otherwise.

### GetFeasibilityCheckAndReservationJobOk

`func (o *Resource) GetFeasibilityCheckAndReservationJobOk() (*[]FeasibilityCheckAndReservationJobSingle, bool)`

GetFeasibilityCheckAndReservationJobOk returns a tuple with the FeasibilityCheckAndReservationJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeasibilityCheckAndReservationJob

`func (o *Resource) SetFeasibilityCheckAndReservationJob(v []FeasibilityCheckAndReservationJobSingle)`

SetFeasibilityCheckAndReservationJob sets FeasibilityCheckAndReservationJob field to given value.

### HasFeasibilityCheckAndReservationJob

`func (o *Resource) HasFeasibilityCheckAndReservationJob() bool`

HasFeasibilityCheckAndReservationJob returns a boolean if a field has been set.

### GetAssuranceGoal

`func (o *Resource) GetAssuranceGoal() []AssuranceGoalSingle`

GetAssuranceGoal returns the AssuranceGoal field if non-nil, zero value otherwise.

### GetAssuranceGoalOk

`func (o *Resource) GetAssuranceGoalOk() (*[]AssuranceGoalSingle, bool)`

GetAssuranceGoalOk returns a tuple with the AssuranceGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceGoal

`func (o *Resource) SetAssuranceGoal(v []AssuranceGoalSingle)`

SetAssuranceGoal sets AssuranceGoal field to given value.

### HasAssuranceGoal

`func (o *Resource) HasAssuranceGoal() bool`

HasAssuranceGoal returns a boolean if a field has been set.

### GetAssuranceClosedControlLoop

`func (o *Resource) GetAssuranceClosedControlLoop() []AssuranceClosedControlLoopSingle`

GetAssuranceClosedControlLoop returns the AssuranceClosedControlLoop field if non-nil, zero value otherwise.

### GetAssuranceClosedControlLoopOk

`func (o *Resource) GetAssuranceClosedControlLoopOk() (*[]AssuranceClosedControlLoopSingle, bool)`

GetAssuranceClosedControlLoopOk returns a tuple with the AssuranceClosedControlLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceClosedControlLoop

`func (o *Resource) SetAssuranceClosedControlLoop(v []AssuranceClosedControlLoopSingle)`

SetAssuranceClosedControlLoop sets AssuranceClosedControlLoop field to given value.

### HasAssuranceClosedControlLoop

`func (o *Resource) HasAssuranceClosedControlLoop() bool`

HasAssuranceClosedControlLoop returns a boolean if a field has been set.

### GetUserLabel

`func (o *Resource) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *Resource) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *Resource) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *Resource) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetIntentExpectations

`func (o *Resource) GetIntentExpectations() []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation`

GetIntentExpectations returns the IntentExpectations field if non-nil, zero value otherwise.

### GetIntentExpectationsOk

`func (o *Resource) GetIntentExpectationsOk() (*[]OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation, bool)`

GetIntentExpectationsOk returns a tuple with the IntentExpectations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentExpectations

`func (o *Resource) SetIntentExpectations(v []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation)`

SetIntentExpectations sets IntentExpectations field to given value.

### HasIntentExpectations

`func (o *Resource) HasIntentExpectations() bool`

HasIntentExpectations returns a boolean if a field has been set.

### GetIntentContexts

`func (o *Resource) GetIntentContexts() []IntentContext`

GetIntentContexts returns the IntentContexts field if non-nil, zero value otherwise.

### GetIntentContextsOk

`func (o *Resource) GetIntentContextsOk() (*[]IntentContext, bool)`

GetIntentContextsOk returns a tuple with the IntentContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentContexts

`func (o *Resource) SetIntentContexts(v []IntentContext)`

SetIntentContexts sets IntentContexts field to given value.

### HasIntentContexts

`func (o *Resource) HasIntentContexts() bool`

HasIntentContexts returns a boolean if a field has been set.

### GetIntentFulfilmentInfo

`func (o *Resource) GetIntentFulfilmentInfo() FulfilmentInfo`

GetIntentFulfilmentInfo returns the IntentFulfilmentInfo field if non-nil, zero value otherwise.

### GetIntentFulfilmentInfoOk

`func (o *Resource) GetIntentFulfilmentInfoOk() (*FulfilmentInfo, bool)`

GetIntentFulfilmentInfoOk returns a tuple with the IntentFulfilmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentFulfilmentInfo

`func (o *Resource) SetIntentFulfilmentInfo(v FulfilmentInfo)`

SetIntentFulfilmentInfo sets IntentFulfilmentInfo field to given value.

### HasIntentFulfilmentInfo

`func (o *Resource) HasIntentFulfilmentInfo() bool`

HasIntentFulfilmentInfo returns a boolean if a field has been set.

### GetMDAFunction

`func (o *Resource) GetMDAFunction() []MDAFunctionSingle`

GetMDAFunction returns the MDAFunction field if non-nil, zero value otherwise.

### GetMDAFunctionOk

`func (o *Resource) GetMDAFunctionOk() (*[]MDAFunctionSingle, bool)`

GetMDAFunctionOk returns a tuple with the MDAFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDAFunction

`func (o *Resource) SetMDAFunction(v []MDAFunctionSingle)`

SetMDAFunction sets MDAFunction field to given value.

### HasMDAFunction

`func (o *Resource) HasMDAFunction() bool`

HasMDAFunction returns a boolean if a field has been set.

### GetMDAReport

`func (o *Resource) GetMDAReport() []MDAReport`

GetMDAReport returns the MDAReport field if non-nil, zero value otherwise.

### GetMDAReportOk

`func (o *Resource) GetMDAReportOk() (*[]MDAReport, bool)`

GetMDAReportOk returns a tuple with the MDAReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDAReport

`func (o *Resource) SetMDAReport(v []MDAReport)`

SetMDAReport sets MDAReport field to given value.

### HasMDAReport

`func (o *Resource) HasMDAReport() bool`

HasMDAReport returns a boolean if a field has been set.

### GetMDARequest

`func (o *Resource) GetMDARequest() []MDARequestSingle`

GetMDARequest returns the MDARequest field if non-nil, zero value otherwise.

### GetMDARequestOk

`func (o *Resource) GetMDARequestOk() (*[]MDARequestSingle, bool)`

GetMDARequestOk returns a tuple with the MDARequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDARequest

`func (o *Resource) SetMDARequest(v []MDARequestSingle)`

SetMDARequest sets MDARequest field to given value.

### HasMDARequest

`func (o *Resource) HasMDARequest() bool`

HasMDARequest returns a boolean if a field has been set.

### GetMLTrainingFunction

`func (o *Resource) GetMLTrainingFunction() []MLTrainingFunctionSingle`

GetMLTrainingFunction returns the MLTrainingFunction field if non-nil, zero value otherwise.

### GetMLTrainingFunctionOk

`func (o *Resource) GetMLTrainingFunctionOk() (*[]MLTrainingFunctionSingle, bool)`

GetMLTrainingFunctionOk returns a tuple with the MLTrainingFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingFunction

`func (o *Resource) SetMLTrainingFunction(v []MLTrainingFunctionSingle)`

SetMLTrainingFunction sets MLTrainingFunction field to given value.

### HasMLTrainingFunction

`func (o *Resource) HasMLTrainingFunction() bool`

HasMLTrainingFunction returns a boolean if a field has been set.

### GetMLTrainingRequest

`func (o *Resource) GetMLTrainingRequest() []MLTrainingRequestSingle`

GetMLTrainingRequest returns the MLTrainingRequest field if non-nil, zero value otherwise.

### GetMLTrainingRequestOk

`func (o *Resource) GetMLTrainingRequestOk() (*[]MLTrainingRequestSingle, bool)`

GetMLTrainingRequestOk returns a tuple with the MLTrainingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingRequest

`func (o *Resource) SetMLTrainingRequest(v []MLTrainingRequestSingle)`

SetMLTrainingRequest sets MLTrainingRequest field to given value.

### HasMLTrainingRequest

`func (o *Resource) HasMLTrainingRequest() bool`

HasMLTrainingRequest returns a boolean if a field has been set.

### GetMLTrainingProcess

`func (o *Resource) GetMLTrainingProcess() []MLTrainingProcessSingle`

GetMLTrainingProcess returns the MLTrainingProcess field if non-nil, zero value otherwise.

### GetMLTrainingProcessOk

`func (o *Resource) GetMLTrainingProcessOk() (*[]MLTrainingProcessSingle, bool)`

GetMLTrainingProcessOk returns a tuple with the MLTrainingProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingProcess

`func (o *Resource) SetMLTrainingProcess(v []MLTrainingProcessSingle)`

SetMLTrainingProcess sets MLTrainingProcess field to given value.

### HasMLTrainingProcess

`func (o *Resource) HasMLTrainingProcess() bool`

HasMLTrainingProcess returns a boolean if a field has been set.

### GetMLTrainingReport

`func (o *Resource) GetMLTrainingReport() []MLTrainingReportSingle`

GetMLTrainingReport returns the MLTrainingReport field if non-nil, zero value otherwise.

### GetMLTrainingReportOk

`func (o *Resource) GetMLTrainingReportOk() (*[]MLTrainingReportSingle, bool)`

GetMLTrainingReportOk returns a tuple with the MLTrainingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingReport

`func (o *Resource) SetMLTrainingReport(v []MLTrainingReportSingle)`

SetMLTrainingReport sets MLTrainingReport field to given value.

### HasMLTrainingReport

`func (o *Resource) HasMLTrainingReport() bool`

HasMLTrainingReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


