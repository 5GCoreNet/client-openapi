# Go API client for NrNrm

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 17.7.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import NrNrm "//"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), NrNrm.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), NrNrm.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), NrNrm.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), NrNrm.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------


## Documentation For Models

 - [AckState](docs/AckState.md)
 - [AddressWithVlan](docs/AddressWithVlan.md)
 - [AdministrativeState](docs/AdministrativeState.md)
 - [AlarmListSingle](docs/AlarmListSingle.md)
 - [AlarmListSingleAllOf](docs/AlarmListSingleAllOf.md)
 - [AlarmListSingleAllOfAttributes](docs/AlarmListSingleAllOfAttributes.md)
 - [AlarmNotificationTypes](docs/AlarmNotificationTypes.md)
 - [AlarmRecord](docs/AlarmRecord.md)
 - [AlarmType](docs/AlarmType.md)
 - [AnonymizationOfMDTDataType](docs/AnonymizationOfMDTDataType.md)
 - [AreaConfig](docs/AreaConfig.md)
 - [AreaOfInterest](docs/AreaOfInterest.md)
 - [AreaScope](docs/AreaScope.md)
 - [BWPSetSingle](docs/BWPSetSingle.md)
 - [BWPSetSingleAllOf](docs/BWPSetSingleAllOf.md)
 - [BackhaulAddress](docs/BackhaulAddress.md)
 - [BeamSingle](docs/BeamSingle.md)
 - [BeamSingleAllOf](docs/BeamSingleAllOf.md)
 - [BwpContext](docs/BwpContext.md)
 - [BwpSingle](docs/BwpSingle.md)
 - [BwpSingleAllOf](docs/BwpSingleAllOf.md)
 - [CCOFunctionSingle](docs/CCOFunctionSingle.md)
 - [CCOFunctionSingleAllOf](docs/CCOFunctionSingleAllOf.md)
 - [CCOFunctionSingleAllOfAttributes](docs/CCOFunctionSingleAllOfAttributes.md)
 - [CCOOvershootCoverageParametersSingle](docs/CCOOvershootCoverageParametersSingle.md)
 - [CCOParametersAttr](docs/CCOParametersAttr.md)
 - [CCOParametersAttrAllOf](docs/CCOParametersAttrAllOf.md)
 - [CCOParametersAttrAllOfAttributes](docs/CCOParametersAttrAllOfAttributes.md)
 - [CCOPilotPollutionParametersSingle](docs/CCOPilotPollutionParametersSingle.md)
 - [CCOWeakCoverageParametersSingle](docs/CCOWeakCoverageParametersSingle.md)
 - [CESManagementFunctionSingle](docs/CESManagementFunctionSingle.md)
 - [CESManagementFunctionSingleAllOf](docs/CESManagementFunctionSingleAllOf.md)
 - [CESManagementFunctionSingleAllOfAttributes](docs/CESManagementFunctionSingleAllOfAttributes.md)
 - [CPCIConfigurationFunctionSingle](docs/CPCIConfigurationFunctionSingle.md)
 - [CPCIConfigurationFunctionSingleAllOf](docs/CPCIConfigurationFunctionSingleAllOf.md)
 - [CPCIConfigurationFunctionSingleAllOfAttributes](docs/CPCIConfigurationFunctionSingleAllOfAttributes.md)
 - [CSonPciList](docs/CSonPciList.md)
 - [CellIndividualOffset](docs/CellIndividualOffset.md)
 - [CellState](docs/CellState.md)
 - [CmNotificationTypes](docs/CmNotificationTypes.md)
 - [CollectionPeriodM6LTEType](docs/CollectionPeriodM6LTEType.md)
 - [CollectionPeriodM6NRType](docs/CollectionPeriodM6NRType.md)
 - [CollectionPeriodRRMLTEType](docs/CollectionPeriodRRMLTEType.md)
 - [CollectionPeriodRRMNRType](docs/CollectionPeriodRRMNRType.md)
 - [CollectionPeriodRRMUMTSType](docs/CollectionPeriodRRMUMTSType.md)
 - [CommonBeamformingFunctionSingle](docs/CommonBeamformingFunctionSingle.md)
 - [CommonBeamformingFunctionSingleAllOf](docs/CommonBeamformingFunctionSingleAllOf.md)
 - [CommonBeamformingFunctionSingleAllOf1](docs/CommonBeamformingFunctionSingleAllOf1.md)
 - [Configurable5QISetSingle](docs/Configurable5QISetSingle.md)
 - [Configurable5QISetSingleAllOf](docs/Configurable5QISetSingleAllOf.md)
 - [CorrelatedNotification](docs/CorrelatedNotification.md)
 - [CyclicPrefix](docs/CyclicPrefix.md)
 - [DANRManagementFunctionSingle](docs/DANRManagementFunctionSingle.md)
 - [DANRManagementFunctionSingleAllOf](docs/DANRManagementFunctionSingleAllOf.md)
 - [DANRManagementFunctionSingleAllOfAttributes](docs/DANRManagementFunctionSingleAllOfAttributes.md)
 - [DESManagementFunctionSingle](docs/DESManagementFunctionSingle.md)
 - [DESManagementFunctionSingleAllOf](docs/DESManagementFunctionSingleAllOf.md)
 - [DESManagementFunctionSingleAllOfAttributes](docs/DESManagementFunctionSingleAllOfAttributes.md)
 - [DLBOFunctionSingle](docs/DLBOFunctionSingle.md)
 - [DLBOFunctionSingleAllOf](docs/DLBOFunctionSingleAllOf.md)
 - [DLBOFunctionSingleAllOfAttributes](docs/DLBOFunctionSingleAllOfAttributes.md)
 - [DMROFunctionSingle](docs/DMROFunctionSingle.md)
 - [DMROFunctionSingleAllOf](docs/DMROFunctionSingleAllOf.md)
 - [DMROFunctionSingleAllOfAttributes](docs/DMROFunctionSingleAllOfAttributes.md)
 - [DPCIConfigurationFunctionSingle](docs/DPCIConfigurationFunctionSingle.md)
 - [DPCIConfigurationFunctionSingleAllOf](docs/DPCIConfigurationFunctionSingleAllOf.md)
 - [DPCIConfigurationFunctionSingleAllOfAttributes](docs/DPCIConfigurationFunctionSingleAllOfAttributes.md)
 - [DRACHOptimizationFunctionSingle](docs/DRACHOptimizationFunctionSingle.md)
 - [DRACHOptimizationFunctionSingleAllOf](docs/DRACHOptimizationFunctionSingleAllOf.md)
 - [DRACHOptimizationFunctionSingleAllOfAttributes](docs/DRACHOptimizationFunctionSingleAllOfAttributes.md)
 - [Dynamic5QISetSingle](docs/Dynamic5QISetSingle.md)
 - [Dynamic5QISetSingleAllOf](docs/Dynamic5QISetSingleAllOf.md)
 - [EPE1Single](docs/EPE1Single.md)
 - [EPF1CSingle](docs/EPF1CSingle.md)
 - [EPF1USingle](docs/EPF1USingle.md)
 - [EPF1USingleAllOf](docs/EPF1USingleAllOf.md)
 - [EPNgCSingle](docs/EPNgCSingle.md)
 - [EPNgUSingle](docs/EPNgUSingle.md)
 - [EPRPAttr](docs/EPRPAttr.md)
 - [EPS1USingle](docs/EPS1USingle.md)
 - [EPX2CSingle](docs/EPX2CSingle.md)
 - [EPX2USingle](docs/EPX2USingle.md)
 - [EPXnCSingle](docs/EPXnCSingle.md)
 - [EPXnCSingleAllOf](docs/EPXnCSingleAllOf.md)
 - [EPXnUSingle](docs/EPXnUSingle.md)
 - [EUtranCellRelationSingle](docs/EUtranCellRelationSingle.md)
 - [EUtranCellRelationSingleAllOf](docs/EUtranCellRelationSingleAllOf.md)
 - [EUtranFreqRelationSingle](docs/EUtranFreqRelationSingle.md)
 - [EUtranFreqRelationSingleAllOf](docs/EUtranFreqRelationSingleAllOf.md)
 - [EUtranFreqRelationSingleAllOfAttributes](docs/EUtranFreqRelationSingleAllOfAttributes.md)
 - [EUtranFrequencySingle](docs/EUtranFrequencySingle.md)
 - [EUtranFrequencySingleAllOf](docs/EUtranFrequencySingleAllOf.md)
 - [EUtranFrequencySingleAllOfAttributes](docs/EUtranFrequencySingleAllOfAttributes.md)
 - [EsNotAllowedTimePeriod](docs/EsNotAllowedTimePeriod.md)
 - [EventListForEventTriggeredMeasurementType](docs/EventListForEventTriggeredMeasurementType.md)
 - [EventThresholdL1Type](docs/EventThresholdL1Type.md)
 - [EventThresholdType](docs/EventThresholdType.md)
 - [EventThresholdTypeEventThreshold1F](docs/EventThresholdTypeEventThreshold1F.md)
 - [EventThresholdTypeEventThresholdRSRP](docs/EventThresholdTypeEventThresholdRSRP.md)
 - [EventThresholdTypeEventThresholdRSRQ](docs/EventThresholdTypeEventThresholdRSRQ.md)
 - [ExternalENBFunctionSingle](docs/ExternalENBFunctionSingle.md)
 - [ExternalENBFunctionSingleAllOf](docs/ExternalENBFunctionSingleAllOf.md)
 - [ExternalENBFunctionSingleAllOf1](docs/ExternalENBFunctionSingleAllOf1.md)
 - [ExternalEUTranCellSingle](docs/ExternalEUTranCellSingle.md)
 - [ExternalEUTranCellSingleAllOf](docs/ExternalEUTranCellSingleAllOf.md)
 - [ExternalGnbCuCpFunctionSingle](docs/ExternalGnbCuCpFunctionSingle.md)
 - [ExternalGnbCuCpFunctionSingleAllOf](docs/ExternalGnbCuCpFunctionSingleAllOf.md)
 - [ExternalGnbCuCpFunctionSingleAllOf1](docs/ExternalGnbCuCpFunctionSingleAllOf1.md)
 - [ExternalGnbCuUpFunctionSingle](docs/ExternalGnbCuUpFunctionSingle.md)
 - [ExternalGnbCuUpFunctionSingleAllOf](docs/ExternalGnbCuUpFunctionSingleAllOf.md)
 - [ExternalGnbDuFunctionSingle](docs/ExternalGnbDuFunctionSingle.md)
 - [ExternalGnbDuFunctionSingleAllOf](docs/ExternalGnbDuFunctionSingleAllOf.md)
 - [ExternalGnbDuFunctionSingleAllOf1](docs/ExternalGnbDuFunctionSingleAllOf1.md)
 - [ExternalNrCellCuSingle](docs/ExternalNrCellCuSingle.md)
 - [ExternalNrCellCuSingleAllOf](docs/ExternalNrCellCuSingleAllOf.md)
 - [FileDownloadJobProcessMonitor](docs/FileDownloadJobProcessMonitor.md)
 - [FileDownloadJobProcessMonitorResultStateInfo](docs/FileDownloadJobProcessMonitorResultStateInfo.md)
 - [FileDownloadJobSingle](docs/FileDownloadJobSingle.md)
 - [FileDownloadJobSingleAllOf](docs/FileDownloadJobSingleAllOf.md)
 - [FileDownloadJobSingleAllOfAttributes](docs/FileDownloadJobSingleAllOfAttributes.md)
 - [FileNotificationTypes](docs/FileNotificationTypes.md)
 - [FileSingle](docs/FileSingle.md)
 - [FileSingleAllOf](docs/FileSingleAllOf.md)
 - [FileSingleAllOfAttributes](docs/FileSingleAllOfAttributes.md)
 - [FilesSingle](docs/FilesSingle.md)
 - [FilesSingleAllOf](docs/FilesSingleAllOf.md)
 - [FilesSingleAllOfAttributes](docs/FilesSingleAllOfAttributes.md)
 - [FiveQICharacteristicsSingle](docs/FiveQICharacteristicsSingle.md)
 - [FiveQICharacteristicsSingleAllOf](docs/FiveQICharacteristicsSingleAllOf.md)
 - [FreqInfo](docs/FreqInfo.md)
 - [FrequencyDomainPara](docs/FrequencyDomainPara.md)
 - [GeoAreaToCellMapping](docs/GeoAreaToCellMapping.md)
 - [GeoCoordinate](docs/GeoCoordinate.md)
 - [GnbCuCpFunctionSingle](docs/GnbCuCpFunctionSingle.md)
 - [GnbCuCpFunctionSingleAllOf](docs/GnbCuCpFunctionSingleAllOf.md)
 - [GnbCuCpFunctionSingleAllOf1](docs/GnbCuCpFunctionSingleAllOf1.md)
 - [GnbCuUpFunctionSingle](docs/GnbCuUpFunctionSingle.md)
 - [GnbCuUpFunctionSingleAllOf](docs/GnbCuUpFunctionSingleAllOf.md)
 - [GnbCuUpFunctionSingleAllOf1](docs/GnbCuUpFunctionSingleAllOf1.md)
 - [GnbDuFunctionSingle](docs/GnbDuFunctionSingle.md)
 - [GnbDuFunctionSingleAllOf](docs/GnbDuFunctionSingleAllOf.md)
 - [GnbDuFunctionSingleAllOf1](docs/GnbDuFunctionSingleAllOf1.md)
 - [HeartbeatControlSingle](docs/HeartbeatControlSingle.md)
 - [HeartbeatControlSingleAllOf](docs/HeartbeatControlSingleAllOf.md)
 - [HeartbeatControlSingleAllOfAttributes](docs/HeartbeatControlSingleAllOfAttributes.md)
 - [HeartbeatNotificationTypes](docs/HeartbeatNotificationTypes.md)
 - [HostAddr](docs/HostAddr.md)
 - [InterRatEsActivationCandidateCellParameters](docs/InterRatEsActivationCandidateCellParameters.md)
 - [InterRatEsActivationOriginalCellParameters](docs/InterRatEsActivationOriginalCellParameters.md)
 - [InterRatEsDeactivationCandidateCellParameters](docs/InterRatEsDeactivationCandidateCellParameters.md)
 - [IntraRatEsActivationCandidateCellsLoadParameters](docs/IntraRatEsActivationCandidateCellsLoadParameters.md)
 - [IntraRatEsActivationOriginalCellLoadParameters](docs/IntraRatEsActivationOriginalCellLoadParameters.md)
 - [IntraRatEsDeactivationCandidateCellsLoadParameters](docs/IntraRatEsDeactivationCandidateCellsLoadParameters.md)
 - [IpAddr](docs/IpAddr.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [IsESCoveredBy](docs/IsESCoveredBy.md)
 - [IsInitialBwp](docs/IsInitialBwp.md)
 - [JobTypeType](docs/JobTypeType.md)
 - [ListOfInterfacesType](docs/ListOfInterfacesType.md)
 - [ListOfMeasurementsType](docs/ListOfMeasurementsType.md)
 - [LocalAddress](docs/LocalAddress.md)
 - [LoggingDurationType](docs/LoggingDurationType.md)
 - [LoggingIntervalType](docs/LoggingIntervalType.md)
 - [ManagedElementAttr](docs/ManagedElementAttr.md)
 - [ManagedElementNcO](docs/ManagedElementNcO.md)
 - [ManagedElementSingle](docs/ManagedElementSingle.md)
 - [ManagedElementSingleAllOf](docs/ManagedElementSingleAllOf.md)
 - [ManagedElementSingleAllOf1](docs/ManagedElementSingleAllOf1.md)
 - [ManagedFunctionAttr](docs/ManagedFunctionAttr.md)
 - [ManagedFunctionNcO](docs/ManagedFunctionNcO.md)
 - [ManagedNFServiceSingle](docs/ManagedNFServiceSingle.md)
 - [ManagedNFServiceSingleAllOf](docs/ManagedNFServiceSingleAllOf.md)
 - [ManagedNFServiceSingleAllOfAttributes](docs/ManagedNFServiceSingleAllOfAttributes.md)
 - [ManagementData](docs/ManagementData.md)
 - [ManagementDataCollectionSingle](docs/ManagementDataCollectionSingle.md)
 - [ManagementDataCollectionSingleAllOf](docs/ManagementDataCollectionSingleAllOf.md)
 - [ManagementDataCollectionSingleAllOfAttributes](docs/ManagementDataCollectionSingleAllOfAttributes.md)
 - [ManagementNodeSingle](docs/ManagementNodeSingle.md)
 - [ManagementNodeSingleAllOf](docs/ManagementNodeSingleAllOf.md)
 - [ManagementNodeSingleAllOfAttributes](docs/ManagementNodeSingleAllOfAttributes.md)
 - [MappingSetIDBackhaulAddress](docs/MappingSetIDBackhaulAddress.md)
 - [MbsfnArea](docs/MbsfnArea.md)
 - [MeContextSingle](docs/MeContextSingle.md)
 - [MeContextSingleAllOf](docs/MeContextSingleAllOf.md)
 - [MeContextSingleAllOfAttributes](docs/MeContextSingleAllOfAttributes.md)
 - [MeasurementPeriodLTEType](docs/MeasurementPeriodLTEType.md)
 - [MeasurementPeriodUMTSType](docs/MeasurementPeriodUMTSType.md)
 - [MeasurementQuantityType](docs/MeasurementQuantityType.md)
 - [MnS](docs/MnS.md)
 - [MnSOneOf](docs/MnSOneOf.md)
 - [MnSOneOf1](docs/MnSOneOf1.md)
 - [MnsAgentSingle](docs/MnsAgentSingle.md)
 - [MnsAgentSingleAllOf](docs/MnsAgentSingleAllOf.md)
 - [MnsAgentSingleAllOfAttributes](docs/MnsAgentSingleAllOfAttributes.md)
 - [MnsInfoSingle](docs/MnsInfoSingle.md)
 - [MnsRegistrySingle](docs/MnsRegistrySingle.md)
 - [NFServiceType](docs/NFServiceType.md)
 - [NFType](docs/NFType.md)
 - [NRCellRelationSingle](docs/NRCellRelationSingle.md)
 - [NRCellRelationSingleAllOf](docs/NRCellRelationSingleAllOf.md)
 - [NRCellRelationSingleAllOfAttributes](docs/NRCellRelationSingleAllOfAttributes.md)
 - [NRFreqRelationSingle](docs/NRFreqRelationSingle.md)
 - [NRFreqRelationSingleAllOf](docs/NRFreqRelationSingleAllOf.md)
 - [NRFreqRelationSingleAllOfAttributes](docs/NRFreqRelationSingleAllOfAttributes.md)
 - [NRFrequencySingle](docs/NRFrequencySingle.md)
 - [NRFrequencySingleAllOf](docs/NRFrequencySingleAllOf.md)
 - [NRFrequencySingleAllOfAttributes](docs/NRFrequencySingleAllOfAttributes.md)
 - [NRPciList](docs/NRPciList.md)
 - [NodeFilter](docs/NodeFilter.md)
 - [NotificationType](docs/NotificationType.md)
 - [NpnIdentity](docs/NpnIdentity.md)
 - [NrCellCuSingle](docs/NrCellCuSingle.md)
 - [NrCellCuSingleAllOf](docs/NrCellCuSingleAllOf.md)
 - [NrCellCuSingleAllOf1](docs/NrCellCuSingleAllOf1.md)
 - [NrCellDuSingle](docs/NrCellDuSingle.md)
 - [NrCellDuSingleAllOf](docs/NrCellDuSingleAllOf.md)
 - [NrCellDuSingleAllOf1](docs/NrCellDuSingleAllOf1.md)
 - [NrOperatorCellDuSingle](docs/NrOperatorCellDuSingle.md)
 - [NrOperatorCellDuSingleAllOf](docs/NrOperatorCellDuSingleAllOf.md)
 - [NrSectorCarrierSingle](docs/NrSectorCarrierSingle.md)
 - [NrSectorCarrierSingleAllOf](docs/NrSectorCarrierSingleAllOf.md)
 - [NrSectorCarrierSingleAllOf1](docs/NrSectorCarrierSingleAllOf1.md)
 - [NtfSubscriptionControlSingle](docs/NtfSubscriptionControlSingle.md)
 - [NtfSubscriptionControlSingleAllOf](docs/NtfSubscriptionControlSingleAllOf.md)
 - [NtfSubscriptionControlSingleAllOfAttributes](docs/NtfSubscriptionControlSingleAllOfAttributes.md)
 - [Operation](docs/Operation.md)
 - [OperationSemantics](docs/OperationSemantics.md)
 - [OperationalState](docs/OperationalState.md)
 - [OperatorDuSingle](docs/OperatorDuSingle.md)
 - [OperatorDuSingleAllOf](docs/OperatorDuSingleAllOf.md)
 - [OperatorDuSingleAllOf1](docs/OperatorDuSingleAllOf1.md)
 - [PLMNListTypeInner](docs/PLMNListTypeInner.md)
 - [PLMNTargetType](docs/PLMNTargetType.md)
 - [PacketErrorRate](docs/PacketErrorRate.md)
 - [ParameterRange](docs/ParameterRange.md)
 - [PeeParameter](docs/PeeParameter.md)
 - [PerceivedSeverity](docs/PerceivedSeverity.md)
 - [PerfMetricJobSingle](docs/PerfMetricJobSingle.md)
 - [PerfMetricJobSingleAllOf](docs/PerfMetricJobSingleAllOf.md)
 - [PerfMetricJobSingleAllOfAttributes](docs/PerfMetricJobSingleAllOfAttributes.md)
 - [PerfNotificationTypes](docs/PerfNotificationTypes.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnInfo](docs/PlmnInfo.md)
 - [PositioningMethodType](docs/PositioningMethodType.md)
 - [ProbableCause](docs/ProbableCause.md)
 - [QOffsetRange](docs/QOffsetRange.md)
 - [QOffsetRangeList](docs/QOffsetRangeList.md)
 - [RRMPolicyRatioSingle](docs/RRMPolicyRatioSingle.md)
 - [RRMPolicyRatioSingleAllOf](docs/RRMPolicyRatioSingleAllOf.md)
 - [RSSetType](docs/RSSetType.md)
 - [RegistrationState](docs/RegistrationState.md)
 - [RemoteAddress](docs/RemoteAddress.md)
 - [ReportAmountType](docs/ReportAmountType.md)
 - [ReportIntervalType](docs/ReportIntervalType.md)
 - [ReportTypeType](docs/ReportTypeType.md)
 - [ReportingCtrl](docs/ReportingCtrl.md)
 - [ReportingCtrlOneOf](docs/ReportingCtrlOneOf.md)
 - [ReportingCtrlOneOf1](docs/ReportingCtrlOneOf1.md)
 - [ReportingCtrlOneOf2](docs/ReportingCtrlOneOf2.md)
 - [ResourceType](docs/ResourceType.md)
 - [ResourcesNrNrm](docs/ResourcesNrNrm.md)
 - [RimRSGlobalSingle](docs/RimRSGlobalSingle.md)
 - [RimRSGlobalSingleAllOf](docs/RimRSGlobalSingleAllOf.md)
 - [RimRSGlobalSingleAllOfAttributes](docs/RimRSGlobalSingleAllOfAttributes.md)
 - [RimRSReportConf](docs/RimRSReportConf.md)
 - [RimRSReportInfo](docs/RimRSReportInfo.md)
 - [RimRSSetSingle](docs/RimRSSetSingle.md)
 - [RimRSSetSingleAllOf](docs/RimRSSetSingleAllOf.md)
 - [RimRSSetSingleAllOfAttributes](docs/RimRSSetSingleAllOfAttributes.md)
 - [RrmPolicyAttr](docs/RrmPolicyAttr.md)
 - [RrmPolicyMember](docs/RrmPolicyMember.md)
 - [SAP](docs/SAP.md)
 - [Scope](docs/Scope.md)
 - [SequenceDomainPara](docs/SequenceDomainPara.md)
 - [Snssai](docs/Snssai.md)
 - [SpecificProblem](docs/SpecificProblem.md)
 - [SsbDuration](docs/SsbDuration.md)
 - [SsbPeriodicity](docs/SsbPeriodicity.md)
 - [SsbSubCarrierSpacing](docs/SsbSubCarrierSpacing.md)
 - [SubNetworkAttr](docs/SubNetworkAttr.md)
 - [SubNetworkNcO](docs/SubNetworkNcO.md)
 - [SubNetworkSingle](docs/SubNetworkSingle.md)
 - [SubNetworkSingleAllOf](docs/SubNetworkSingleAllOf.md)
 - [SubNetworkSingleAllOf1](docs/SubNetworkSingleAllOf1.md)
 - [SupportedPerfMetricGroup](docs/SupportedPerfMetricGroup.md)
 - [TReselectionNRSf](docs/TReselectionNRSf.md)
 - [Tai](docs/Tai.md)
 - [Tai1](docs/Tai1.md)
 - [TceMappingInfo](docs/TceMappingInfo.md)
 - [TceMappingInfoTceIPAddress](docs/TceMappingInfoTceIPAddress.md)
 - [ThresholdInfo](docs/ThresholdInfo.md)
 - [ThresholdInfo1](docs/ThresholdInfo1.md)
 - [ThresholdInfoHysteresis](docs/ThresholdInfoHysteresis.md)
 - [ThresholdInfoThresholdValue](docs/ThresholdInfoThresholdValue.md)
 - [ThresholdLevelInd](docs/ThresholdLevelInd.md)
 - [ThresholdLevelIndOneOf](docs/ThresholdLevelIndOneOf.md)
 - [ThresholdLevelIndOneOf1](docs/ThresholdLevelIndOneOf1.md)
 - [ThresholdMonitorSingle](docs/ThresholdMonitorSingle.md)
 - [ThresholdMonitorSingleAllOf](docs/ThresholdMonitorSingleAllOf.md)
 - [ThresholdMonitorSingleAllOfAttributes](docs/ThresholdMonitorSingleAllOfAttributes.md)
 - [TimeDomainPara](docs/TimeDomainPara.md)
 - [TimeToTriggerL1Type](docs/TimeToTriggerL1Type.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [Top](docs/Top.md)
 - [TraceDepthType](docs/TraceDepthType.md)
 - [TraceJobAttr](docs/TraceJobAttr.md)
 - [TraceJobSingle](docs/TraceJobSingle.md)
 - [TraceJobSingleAllOf](docs/TraceJobSingleAllOf.md)
 - [TraceReferenceType](docs/TraceReferenceType.md)
 - [TraceReportingFormatType](docs/TraceReportingFormatType.md)
 - [TraceTargetType](docs/TraceTargetType.md)
 - [TrendIndication](docs/TrendIndication.md)
 - [TriggeringEventsType](docs/TriggeringEventsType.md)
 - [TxDirection](docs/TxDirection.md)
 - [UeAccDelayProbilityDist](docs/UeAccDelayProbilityDist.md)
 - [UeAccProbilityDist](docs/UeAccProbilityDist.md)
 - [UsageState](docs/UsageState.md)
 - [VnfParameter](docs/VnfParameter.md)
 - [VsDataContainerSingle](docs/VsDataContainerSingle.md)
 - [VsDataContainerSingleAttributes](docs/VsDataContainerSingleAttributes.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

