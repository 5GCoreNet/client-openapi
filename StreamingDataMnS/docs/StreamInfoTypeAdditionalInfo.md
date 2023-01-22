# StreamInfoTypeAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | Pointer to [**JobTypeType**](JobTypeType.md) |  | [optional] 
**ListOfInterfaces** | Pointer to [**ListOfInterfacesType**](ListOfInterfacesType.md) |  | [optional] 
**ListOfNETypes** | Pointer to **[]string** | The Network Element types where Trace Session activation is needed. See 3GPP TS 32.422 clause 5.4 for additional details. | [optional] 
**PLMNTarget** | Pointer to [**PLMNTargetType**](PLMNTargetType.md) |  | [optional] 
**TraceReportingConsumerUri** | Pointer to **string** |  | [optional] 
**TraceCollectionEntityIPAddress** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**TraceDepth** | Pointer to [**TraceDepthType**](TraceDepthType.md) |  | [optional] 
**TraceReference** | Pointer to [**TraceReferenceType1**](TraceReferenceType1.md) |  | [optional] 
**TraceRecordingSessionReference** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**TraceReportingFormat** | Pointer to [**TraceReportingFormatType**](TraceReportingFormatType.md) |  | [optional] 
**TraceTarget** | Pointer to [**TraceTargetType**](TraceTargetType.md) |  | [optional] 
**TriggeringEvents** | Pointer to [**TriggeringEventsType**](TriggeringEventsType.md) |  | [optional] 
**AnonymizationOfMDTData** | Pointer to [**AnonymizationOfMDTDataType**](AnonymizationOfMDTDataType.md) |  | [optional] 
**AreaConfigurationForNeighCell** | Pointer to [**AreaConfig**](AreaConfig.md) |  | [optional] 
**AreaScope** | Pointer to [**[]AreaScope**](AreaScope.md) |  | [optional] 
**BeamLevelMeasurement** | Pointer to **bool** | Determines whether beam level measurements shall be included in case of immediate MDT M1 measurement in NR. For additional details see 3GPP TS 32.422 clause 5.10.40. | [optional] 
**CollectionPeriodRRMLTE** | Pointer to [**CollectionPeriodRRMLTEType**](CollectionPeriodRRMLTEType.md) |  | [optional] 
**CollectionPeriodM6LTE** | Pointer to [**CollectionPeriodM6LTEType**](CollectionPeriodM6LTEType.md) |  | [optional] 
**CollectionPeriodM7LTE** | Pointer to **int32** | See details in 3GPP TS 32.422 clause 5.10.33. | [optional] 
**CollectionPeriodRRMUMTS** | Pointer to [**CollectionPeriodRRMUMTSType**](CollectionPeriodRRMUMTSType.md) |  | [optional] 
**CollectionPeriodRRMNR** | Pointer to [**CollectionPeriodRRMNRType**](CollectionPeriodRRMNRType.md) |  | [optional] 
**CollectionPeriodM6NR** | Pointer to [**CollectionPeriodM6NRType**](CollectionPeriodM6NRType.md) |  | [optional] 
**CollectionPeriodM7NR** | Pointer to **int32** | See details in 3GPP TS 32.422 clause 5.10.35. | [optional] 
**EventListForEventTriggeredMeasurement** | Pointer to [**EventListForEventTriggeredMeasurementType**](EventListForEventTriggeredMeasurementType.md) |  | [optional] 
**EventThreshold** | Pointer to [**EventThresholdType**](EventThresholdType.md) |  | [optional] 
**ListOfMeasurements** | Pointer to [**ListOfMeasurementsType**](ListOfMeasurementsType.md) |  | [optional] 
**LoggingDuration** | Pointer to [**LoggingDurationType**](LoggingDurationType.md) |  | [optional] 
**LoggingInterval** | Pointer to [**LoggingIntervalType**](LoggingIntervalType.md) |  | [optional] 
**EventThresholdL1** | Pointer to [**EventThresholdL1Type**](EventThresholdL1Type.md) |  | [optional] 
**HysteresisL1** | Pointer to **int32** | See details in 3GPP TS 32.422 clause 5.10.Y. | [optional] 
**TimeToTriggerL1** | Pointer to [**TimeToTriggerL1Type**](TimeToTriggerL1Type.md) |  | [optional] 
**MBSFNAreaList** | Pointer to [**[]MbsfnArea**](MbsfnArea.md) |  | [optional] 
**MeasurementPeriodLTE** | Pointer to [**MeasurementPeriodLTEType**](MeasurementPeriodLTEType.md) |  | [optional] 
**MeasurementPeriodUMTS** | Pointer to [**MeasurementPeriodUMTSType**](MeasurementPeriodUMTSType.md) |  | [optional] 
**MeasurementQuantity** | Pointer to [**MeasurementQuantityType**](MeasurementQuantityType.md) |  | [optional] 
**EventThresholdUphUMTS** | Pointer to **int32** | See details in 3GPP TS 32.422 clause 5.10.A. | [optional] 
**PLMNList** | Pointer to [**[]PLMNListTypeInner**](PLMNListTypeInner.md) | See details in 3GPP TS 32.422 clause 5.10.24. | [optional] 
**PositioningMethod** | Pointer to [**PositioningMethodType**](PositioningMethodType.md) |  | [optional] 
**ReportAmount** | Pointer to [**ReportAmountType**](ReportAmountType.md) |  | [optional] 
**ReportingTrigger** | Pointer to **[]string** | See details in 3GPP TS 32.422 clause 5.10.4. | [optional] 
**ReportInterval** | Pointer to [**ReportIntervalType**](ReportIntervalType.md) |  | [optional] 
**ReportType** | Pointer to [**ReportTypeType**](ReportTypeType.md) |  | [optional] 
**SensorInformation** | Pointer to **[]string** | See details in 3GPP TS 32.422 clause 5.10.29. | [optional] 
**TraceCollectionEntityId** | Pointer to **int32** | See details in 3GPP TS 32.422 clause 5.10.11. Only TCE Id value may be sent over the air to the UE being configured for Logged MDT. | [optional] 
**MeasObjDn** | [**MeasObjDnType**](MeasObjDnType.md) |  | 
**PerformanceMetrics** | **[]string** | an ordered list of performance metric names (see clause 4.4.1 of 3GPP TS 28.622[11]) whose values are to be reported by the Performance Data Stream Units (see Annex C of TS 28.550 [42]) via this stream. Performance metrics include measurement and KPI | 
**ActivityDetails** | Pointer to **string** |  | [optional] 
**VsDataType** | Pointer to **string** |  | [optional] 
**VsData** | Pointer to **string** |  | [optional] 
**VsDataFormatVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewStreamInfoTypeAdditionalInfo

`func NewStreamInfoTypeAdditionalInfo(measObjDn MeasObjDnType, performanceMetrics []string, ) *StreamInfoTypeAdditionalInfo`

NewStreamInfoTypeAdditionalInfo instantiates a new StreamInfoTypeAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamInfoTypeAdditionalInfoWithDefaults

`func NewStreamInfoTypeAdditionalInfoWithDefaults() *StreamInfoTypeAdditionalInfo`

NewStreamInfoTypeAdditionalInfoWithDefaults instantiates a new StreamInfoTypeAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *StreamInfoTypeAdditionalInfo) GetJobType() JobTypeType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *StreamInfoTypeAdditionalInfo) GetJobTypeOk() (*JobTypeType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *StreamInfoTypeAdditionalInfo) SetJobType(v JobTypeType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *StreamInfoTypeAdditionalInfo) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetListOfInterfaces

`func (o *StreamInfoTypeAdditionalInfo) GetListOfInterfaces() ListOfInterfacesType`

GetListOfInterfaces returns the ListOfInterfaces field if non-nil, zero value otherwise.

### GetListOfInterfacesOk

`func (o *StreamInfoTypeAdditionalInfo) GetListOfInterfacesOk() (*ListOfInterfacesType, bool)`

GetListOfInterfacesOk returns a tuple with the ListOfInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfInterfaces

`func (o *StreamInfoTypeAdditionalInfo) SetListOfInterfaces(v ListOfInterfacesType)`

SetListOfInterfaces sets ListOfInterfaces field to given value.

### HasListOfInterfaces

`func (o *StreamInfoTypeAdditionalInfo) HasListOfInterfaces() bool`

HasListOfInterfaces returns a boolean if a field has been set.

### GetListOfNETypes

`func (o *StreamInfoTypeAdditionalInfo) GetListOfNETypes() []string`

GetListOfNETypes returns the ListOfNETypes field if non-nil, zero value otherwise.

### GetListOfNETypesOk

`func (o *StreamInfoTypeAdditionalInfo) GetListOfNETypesOk() (*[]string, bool)`

GetListOfNETypesOk returns a tuple with the ListOfNETypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfNETypes

`func (o *StreamInfoTypeAdditionalInfo) SetListOfNETypes(v []string)`

SetListOfNETypes sets ListOfNETypes field to given value.

### HasListOfNETypes

`func (o *StreamInfoTypeAdditionalInfo) HasListOfNETypes() bool`

HasListOfNETypes returns a boolean if a field has been set.

### GetPLMNTarget

`func (o *StreamInfoTypeAdditionalInfo) GetPLMNTarget() PLMNTargetType`

GetPLMNTarget returns the PLMNTarget field if non-nil, zero value otherwise.

### GetPLMNTargetOk

`func (o *StreamInfoTypeAdditionalInfo) GetPLMNTargetOk() (*PLMNTargetType, bool)`

GetPLMNTargetOk returns a tuple with the PLMNTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPLMNTarget

`func (o *StreamInfoTypeAdditionalInfo) SetPLMNTarget(v PLMNTargetType)`

SetPLMNTarget sets PLMNTarget field to given value.

### HasPLMNTarget

`func (o *StreamInfoTypeAdditionalInfo) HasPLMNTarget() bool`

HasPLMNTarget returns a boolean if a field has been set.

### GetTraceReportingConsumerUri

`func (o *StreamInfoTypeAdditionalInfo) GetTraceReportingConsumerUri() string`

GetTraceReportingConsumerUri returns the TraceReportingConsumerUri field if non-nil, zero value otherwise.

### GetTraceReportingConsumerUriOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceReportingConsumerUriOk() (*string, bool)`

GetTraceReportingConsumerUriOk returns a tuple with the TraceReportingConsumerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceReportingConsumerUri

`func (o *StreamInfoTypeAdditionalInfo) SetTraceReportingConsumerUri(v string)`

SetTraceReportingConsumerUri sets TraceReportingConsumerUri field to given value.

### HasTraceReportingConsumerUri

`func (o *StreamInfoTypeAdditionalInfo) HasTraceReportingConsumerUri() bool`

HasTraceReportingConsumerUri returns a boolean if a field has been set.

### GetTraceCollectionEntityIPAddress

`func (o *StreamInfoTypeAdditionalInfo) GetTraceCollectionEntityIPAddress() IpAddr`

GetTraceCollectionEntityIPAddress returns the TraceCollectionEntityIPAddress field if non-nil, zero value otherwise.

### GetTraceCollectionEntityIPAddressOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceCollectionEntityIPAddressOk() (*IpAddr, bool)`

GetTraceCollectionEntityIPAddressOk returns a tuple with the TraceCollectionEntityIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceCollectionEntityIPAddress

`func (o *StreamInfoTypeAdditionalInfo) SetTraceCollectionEntityIPAddress(v IpAddr)`

SetTraceCollectionEntityIPAddress sets TraceCollectionEntityIPAddress field to given value.

### HasTraceCollectionEntityIPAddress

`func (o *StreamInfoTypeAdditionalInfo) HasTraceCollectionEntityIPAddress() bool`

HasTraceCollectionEntityIPAddress returns a boolean if a field has been set.

### GetTraceDepth

`func (o *StreamInfoTypeAdditionalInfo) GetTraceDepth() TraceDepthType`

GetTraceDepth returns the TraceDepth field if non-nil, zero value otherwise.

### GetTraceDepthOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceDepthOk() (*TraceDepthType, bool)`

GetTraceDepthOk returns a tuple with the TraceDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceDepth

`func (o *StreamInfoTypeAdditionalInfo) SetTraceDepth(v TraceDepthType)`

SetTraceDepth sets TraceDepth field to given value.

### HasTraceDepth

`func (o *StreamInfoTypeAdditionalInfo) HasTraceDepth() bool`

HasTraceDepth returns a boolean if a field has been set.

### GetTraceReference

`func (o *StreamInfoTypeAdditionalInfo) GetTraceReference() TraceReferenceType1`

GetTraceReference returns the TraceReference field if non-nil, zero value otherwise.

### GetTraceReferenceOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceReferenceOk() (*TraceReferenceType1, bool)`

GetTraceReferenceOk returns a tuple with the TraceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceReference

`func (o *StreamInfoTypeAdditionalInfo) SetTraceReference(v TraceReferenceType1)`

SetTraceReference sets TraceReference field to given value.

### HasTraceReference

`func (o *StreamInfoTypeAdditionalInfo) HasTraceReference() bool`

HasTraceReference returns a boolean if a field has been set.

### GetTraceRecordingSessionReference

`func (o *StreamInfoTypeAdditionalInfo) GetTraceRecordingSessionReference() string`

GetTraceRecordingSessionReference returns the TraceRecordingSessionReference field if non-nil, zero value otherwise.

### GetTraceRecordingSessionReferenceOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceRecordingSessionReferenceOk() (*string, bool)`

GetTraceRecordingSessionReferenceOk returns a tuple with the TraceRecordingSessionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRecordingSessionReference

`func (o *StreamInfoTypeAdditionalInfo) SetTraceRecordingSessionReference(v string)`

SetTraceRecordingSessionReference sets TraceRecordingSessionReference field to given value.

### HasTraceRecordingSessionReference

`func (o *StreamInfoTypeAdditionalInfo) HasTraceRecordingSessionReference() bool`

HasTraceRecordingSessionReference returns a boolean if a field has been set.

### GetJobId

`func (o *StreamInfoTypeAdditionalInfo) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *StreamInfoTypeAdditionalInfo) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *StreamInfoTypeAdditionalInfo) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *StreamInfoTypeAdditionalInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetTraceReportingFormat

`func (o *StreamInfoTypeAdditionalInfo) GetTraceReportingFormat() TraceReportingFormatType`

GetTraceReportingFormat returns the TraceReportingFormat field if non-nil, zero value otherwise.

### GetTraceReportingFormatOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceReportingFormatOk() (*TraceReportingFormatType, bool)`

GetTraceReportingFormatOk returns a tuple with the TraceReportingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceReportingFormat

`func (o *StreamInfoTypeAdditionalInfo) SetTraceReportingFormat(v TraceReportingFormatType)`

SetTraceReportingFormat sets TraceReportingFormat field to given value.

### HasTraceReportingFormat

`func (o *StreamInfoTypeAdditionalInfo) HasTraceReportingFormat() bool`

HasTraceReportingFormat returns a boolean if a field has been set.

### GetTraceTarget

`func (o *StreamInfoTypeAdditionalInfo) GetTraceTarget() TraceTargetType`

GetTraceTarget returns the TraceTarget field if non-nil, zero value otherwise.

### GetTraceTargetOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceTargetOk() (*TraceTargetType, bool)`

GetTraceTargetOk returns a tuple with the TraceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceTarget

`func (o *StreamInfoTypeAdditionalInfo) SetTraceTarget(v TraceTargetType)`

SetTraceTarget sets TraceTarget field to given value.

### HasTraceTarget

`func (o *StreamInfoTypeAdditionalInfo) HasTraceTarget() bool`

HasTraceTarget returns a boolean if a field has been set.

### GetTriggeringEvents

`func (o *StreamInfoTypeAdditionalInfo) GetTriggeringEvents() TriggeringEventsType`

GetTriggeringEvents returns the TriggeringEvents field if non-nil, zero value otherwise.

### GetTriggeringEventsOk

`func (o *StreamInfoTypeAdditionalInfo) GetTriggeringEventsOk() (*TriggeringEventsType, bool)`

GetTriggeringEventsOk returns a tuple with the TriggeringEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringEvents

`func (o *StreamInfoTypeAdditionalInfo) SetTriggeringEvents(v TriggeringEventsType)`

SetTriggeringEvents sets TriggeringEvents field to given value.

### HasTriggeringEvents

`func (o *StreamInfoTypeAdditionalInfo) HasTriggeringEvents() bool`

HasTriggeringEvents returns a boolean if a field has been set.

### GetAnonymizationOfMDTData

`func (o *StreamInfoTypeAdditionalInfo) GetAnonymizationOfMDTData() AnonymizationOfMDTDataType`

GetAnonymizationOfMDTData returns the AnonymizationOfMDTData field if non-nil, zero value otherwise.

### GetAnonymizationOfMDTDataOk

`func (o *StreamInfoTypeAdditionalInfo) GetAnonymizationOfMDTDataOk() (*AnonymizationOfMDTDataType, bool)`

GetAnonymizationOfMDTDataOk returns a tuple with the AnonymizationOfMDTData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizationOfMDTData

`func (o *StreamInfoTypeAdditionalInfo) SetAnonymizationOfMDTData(v AnonymizationOfMDTDataType)`

SetAnonymizationOfMDTData sets AnonymizationOfMDTData field to given value.

### HasAnonymizationOfMDTData

`func (o *StreamInfoTypeAdditionalInfo) HasAnonymizationOfMDTData() bool`

HasAnonymizationOfMDTData returns a boolean if a field has been set.

### GetAreaConfigurationForNeighCell

`func (o *StreamInfoTypeAdditionalInfo) GetAreaConfigurationForNeighCell() AreaConfig`

GetAreaConfigurationForNeighCell returns the AreaConfigurationForNeighCell field if non-nil, zero value otherwise.

### GetAreaConfigurationForNeighCellOk

`func (o *StreamInfoTypeAdditionalInfo) GetAreaConfigurationForNeighCellOk() (*AreaConfig, bool)`

GetAreaConfigurationForNeighCellOk returns a tuple with the AreaConfigurationForNeighCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaConfigurationForNeighCell

`func (o *StreamInfoTypeAdditionalInfo) SetAreaConfigurationForNeighCell(v AreaConfig)`

SetAreaConfigurationForNeighCell sets AreaConfigurationForNeighCell field to given value.

### HasAreaConfigurationForNeighCell

`func (o *StreamInfoTypeAdditionalInfo) HasAreaConfigurationForNeighCell() bool`

HasAreaConfigurationForNeighCell returns a boolean if a field has been set.

### GetAreaScope

`func (o *StreamInfoTypeAdditionalInfo) GetAreaScope() []AreaScope`

GetAreaScope returns the AreaScope field if non-nil, zero value otherwise.

### GetAreaScopeOk

`func (o *StreamInfoTypeAdditionalInfo) GetAreaScopeOk() (*[]AreaScope, bool)`

GetAreaScopeOk returns a tuple with the AreaScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaScope

`func (o *StreamInfoTypeAdditionalInfo) SetAreaScope(v []AreaScope)`

SetAreaScope sets AreaScope field to given value.

### HasAreaScope

`func (o *StreamInfoTypeAdditionalInfo) HasAreaScope() bool`

HasAreaScope returns a boolean if a field has been set.

### GetBeamLevelMeasurement

`func (o *StreamInfoTypeAdditionalInfo) GetBeamLevelMeasurement() bool`

GetBeamLevelMeasurement returns the BeamLevelMeasurement field if non-nil, zero value otherwise.

### GetBeamLevelMeasurementOk

`func (o *StreamInfoTypeAdditionalInfo) GetBeamLevelMeasurementOk() (*bool, bool)`

GetBeamLevelMeasurementOk returns a tuple with the BeamLevelMeasurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamLevelMeasurement

`func (o *StreamInfoTypeAdditionalInfo) SetBeamLevelMeasurement(v bool)`

SetBeamLevelMeasurement sets BeamLevelMeasurement field to given value.

### HasBeamLevelMeasurement

`func (o *StreamInfoTypeAdditionalInfo) HasBeamLevelMeasurement() bool`

HasBeamLevelMeasurement returns a boolean if a field has been set.

### GetCollectionPeriodRRMLTE

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodRRMLTE() CollectionPeriodRRMLTEType`

GetCollectionPeriodRRMLTE returns the CollectionPeriodRRMLTE field if non-nil, zero value otherwise.

### GetCollectionPeriodRRMLTEOk

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodRRMLTEOk() (*CollectionPeriodRRMLTEType, bool)`

GetCollectionPeriodRRMLTEOk returns a tuple with the CollectionPeriodRRMLTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRRMLTE

`func (o *StreamInfoTypeAdditionalInfo) SetCollectionPeriodRRMLTE(v CollectionPeriodRRMLTEType)`

SetCollectionPeriodRRMLTE sets CollectionPeriodRRMLTE field to given value.

### HasCollectionPeriodRRMLTE

`func (o *StreamInfoTypeAdditionalInfo) HasCollectionPeriodRRMLTE() bool`

HasCollectionPeriodRRMLTE returns a boolean if a field has been set.

### GetCollectionPeriodM6LTE

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM6LTE() CollectionPeriodM6LTEType`

GetCollectionPeriodM6LTE returns the CollectionPeriodM6LTE field if non-nil, zero value otherwise.

### GetCollectionPeriodM6LTEOk

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM6LTEOk() (*CollectionPeriodM6LTEType, bool)`

GetCollectionPeriodM6LTEOk returns a tuple with the CollectionPeriodM6LTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodM6LTE

`func (o *StreamInfoTypeAdditionalInfo) SetCollectionPeriodM6LTE(v CollectionPeriodM6LTEType)`

SetCollectionPeriodM6LTE sets CollectionPeriodM6LTE field to given value.

### HasCollectionPeriodM6LTE

`func (o *StreamInfoTypeAdditionalInfo) HasCollectionPeriodM6LTE() bool`

HasCollectionPeriodM6LTE returns a boolean if a field has been set.

### GetCollectionPeriodM7LTE

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM7LTE() int32`

GetCollectionPeriodM7LTE returns the CollectionPeriodM7LTE field if non-nil, zero value otherwise.

### GetCollectionPeriodM7LTEOk

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM7LTEOk() (*int32, bool)`

GetCollectionPeriodM7LTEOk returns a tuple with the CollectionPeriodM7LTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodM7LTE

`func (o *StreamInfoTypeAdditionalInfo) SetCollectionPeriodM7LTE(v int32)`

SetCollectionPeriodM7LTE sets CollectionPeriodM7LTE field to given value.

### HasCollectionPeriodM7LTE

`func (o *StreamInfoTypeAdditionalInfo) HasCollectionPeriodM7LTE() bool`

HasCollectionPeriodM7LTE returns a boolean if a field has been set.

### GetCollectionPeriodRRMUMTS

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodRRMUMTS() CollectionPeriodRRMUMTSType`

GetCollectionPeriodRRMUMTS returns the CollectionPeriodRRMUMTS field if non-nil, zero value otherwise.

### GetCollectionPeriodRRMUMTSOk

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodRRMUMTSOk() (*CollectionPeriodRRMUMTSType, bool)`

GetCollectionPeriodRRMUMTSOk returns a tuple with the CollectionPeriodRRMUMTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRRMUMTS

`func (o *StreamInfoTypeAdditionalInfo) SetCollectionPeriodRRMUMTS(v CollectionPeriodRRMUMTSType)`

SetCollectionPeriodRRMUMTS sets CollectionPeriodRRMUMTS field to given value.

### HasCollectionPeriodRRMUMTS

`func (o *StreamInfoTypeAdditionalInfo) HasCollectionPeriodRRMUMTS() bool`

HasCollectionPeriodRRMUMTS returns a boolean if a field has been set.

### GetCollectionPeriodRRMNR

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodRRMNR() CollectionPeriodRRMNRType`

GetCollectionPeriodRRMNR returns the CollectionPeriodRRMNR field if non-nil, zero value otherwise.

### GetCollectionPeriodRRMNROk

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodRRMNROk() (*CollectionPeriodRRMNRType, bool)`

GetCollectionPeriodRRMNROk returns a tuple with the CollectionPeriodRRMNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRRMNR

`func (o *StreamInfoTypeAdditionalInfo) SetCollectionPeriodRRMNR(v CollectionPeriodRRMNRType)`

SetCollectionPeriodRRMNR sets CollectionPeriodRRMNR field to given value.

### HasCollectionPeriodRRMNR

`func (o *StreamInfoTypeAdditionalInfo) HasCollectionPeriodRRMNR() bool`

HasCollectionPeriodRRMNR returns a boolean if a field has been set.

### GetCollectionPeriodM6NR

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM6NR() CollectionPeriodM6NRType`

GetCollectionPeriodM6NR returns the CollectionPeriodM6NR field if non-nil, zero value otherwise.

### GetCollectionPeriodM6NROk

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM6NROk() (*CollectionPeriodM6NRType, bool)`

GetCollectionPeriodM6NROk returns a tuple with the CollectionPeriodM6NR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodM6NR

`func (o *StreamInfoTypeAdditionalInfo) SetCollectionPeriodM6NR(v CollectionPeriodM6NRType)`

SetCollectionPeriodM6NR sets CollectionPeriodM6NR field to given value.

### HasCollectionPeriodM6NR

`func (o *StreamInfoTypeAdditionalInfo) HasCollectionPeriodM6NR() bool`

HasCollectionPeriodM6NR returns a boolean if a field has been set.

### GetCollectionPeriodM7NR

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM7NR() int32`

GetCollectionPeriodM7NR returns the CollectionPeriodM7NR field if non-nil, zero value otherwise.

### GetCollectionPeriodM7NROk

`func (o *StreamInfoTypeAdditionalInfo) GetCollectionPeriodM7NROk() (*int32, bool)`

GetCollectionPeriodM7NROk returns a tuple with the CollectionPeriodM7NR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodM7NR

`func (o *StreamInfoTypeAdditionalInfo) SetCollectionPeriodM7NR(v int32)`

SetCollectionPeriodM7NR sets CollectionPeriodM7NR field to given value.

### HasCollectionPeriodM7NR

`func (o *StreamInfoTypeAdditionalInfo) HasCollectionPeriodM7NR() bool`

HasCollectionPeriodM7NR returns a boolean if a field has been set.

### GetEventListForEventTriggeredMeasurement

`func (o *StreamInfoTypeAdditionalInfo) GetEventListForEventTriggeredMeasurement() EventListForEventTriggeredMeasurementType`

GetEventListForEventTriggeredMeasurement returns the EventListForEventTriggeredMeasurement field if non-nil, zero value otherwise.

### GetEventListForEventTriggeredMeasurementOk

`func (o *StreamInfoTypeAdditionalInfo) GetEventListForEventTriggeredMeasurementOk() (*EventListForEventTriggeredMeasurementType, bool)`

GetEventListForEventTriggeredMeasurementOk returns a tuple with the EventListForEventTriggeredMeasurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventListForEventTriggeredMeasurement

`func (o *StreamInfoTypeAdditionalInfo) SetEventListForEventTriggeredMeasurement(v EventListForEventTriggeredMeasurementType)`

SetEventListForEventTriggeredMeasurement sets EventListForEventTriggeredMeasurement field to given value.

### HasEventListForEventTriggeredMeasurement

`func (o *StreamInfoTypeAdditionalInfo) HasEventListForEventTriggeredMeasurement() bool`

HasEventListForEventTriggeredMeasurement returns a boolean if a field has been set.

### GetEventThreshold

`func (o *StreamInfoTypeAdditionalInfo) GetEventThreshold() EventThresholdType`

GetEventThreshold returns the EventThreshold field if non-nil, zero value otherwise.

### GetEventThresholdOk

`func (o *StreamInfoTypeAdditionalInfo) GetEventThresholdOk() (*EventThresholdType, bool)`

GetEventThresholdOk returns a tuple with the EventThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThreshold

`func (o *StreamInfoTypeAdditionalInfo) SetEventThreshold(v EventThresholdType)`

SetEventThreshold sets EventThreshold field to given value.

### HasEventThreshold

`func (o *StreamInfoTypeAdditionalInfo) HasEventThreshold() bool`

HasEventThreshold returns a boolean if a field has been set.

### GetListOfMeasurements

`func (o *StreamInfoTypeAdditionalInfo) GetListOfMeasurements() ListOfMeasurementsType`

GetListOfMeasurements returns the ListOfMeasurements field if non-nil, zero value otherwise.

### GetListOfMeasurementsOk

`func (o *StreamInfoTypeAdditionalInfo) GetListOfMeasurementsOk() (*ListOfMeasurementsType, bool)`

GetListOfMeasurementsOk returns a tuple with the ListOfMeasurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfMeasurements

`func (o *StreamInfoTypeAdditionalInfo) SetListOfMeasurements(v ListOfMeasurementsType)`

SetListOfMeasurements sets ListOfMeasurements field to given value.

### HasListOfMeasurements

`func (o *StreamInfoTypeAdditionalInfo) HasListOfMeasurements() bool`

HasListOfMeasurements returns a boolean if a field has been set.

### GetLoggingDuration

`func (o *StreamInfoTypeAdditionalInfo) GetLoggingDuration() LoggingDurationType`

GetLoggingDuration returns the LoggingDuration field if non-nil, zero value otherwise.

### GetLoggingDurationOk

`func (o *StreamInfoTypeAdditionalInfo) GetLoggingDurationOk() (*LoggingDurationType, bool)`

GetLoggingDurationOk returns a tuple with the LoggingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingDuration

`func (o *StreamInfoTypeAdditionalInfo) SetLoggingDuration(v LoggingDurationType)`

SetLoggingDuration sets LoggingDuration field to given value.

### HasLoggingDuration

`func (o *StreamInfoTypeAdditionalInfo) HasLoggingDuration() bool`

HasLoggingDuration returns a boolean if a field has been set.

### GetLoggingInterval

`func (o *StreamInfoTypeAdditionalInfo) GetLoggingInterval() LoggingIntervalType`

GetLoggingInterval returns the LoggingInterval field if non-nil, zero value otherwise.

### GetLoggingIntervalOk

`func (o *StreamInfoTypeAdditionalInfo) GetLoggingIntervalOk() (*LoggingIntervalType, bool)`

GetLoggingIntervalOk returns a tuple with the LoggingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingInterval

`func (o *StreamInfoTypeAdditionalInfo) SetLoggingInterval(v LoggingIntervalType)`

SetLoggingInterval sets LoggingInterval field to given value.

### HasLoggingInterval

`func (o *StreamInfoTypeAdditionalInfo) HasLoggingInterval() bool`

HasLoggingInterval returns a boolean if a field has been set.

### GetEventThresholdL1

`func (o *StreamInfoTypeAdditionalInfo) GetEventThresholdL1() EventThresholdL1Type`

GetEventThresholdL1 returns the EventThresholdL1 field if non-nil, zero value otherwise.

### GetEventThresholdL1Ok

`func (o *StreamInfoTypeAdditionalInfo) GetEventThresholdL1Ok() (*EventThresholdL1Type, bool)`

GetEventThresholdL1Ok returns a tuple with the EventThresholdL1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdL1

`func (o *StreamInfoTypeAdditionalInfo) SetEventThresholdL1(v EventThresholdL1Type)`

SetEventThresholdL1 sets EventThresholdL1 field to given value.

### HasEventThresholdL1

`func (o *StreamInfoTypeAdditionalInfo) HasEventThresholdL1() bool`

HasEventThresholdL1 returns a boolean if a field has been set.

### GetHysteresisL1

`func (o *StreamInfoTypeAdditionalInfo) GetHysteresisL1() int32`

GetHysteresisL1 returns the HysteresisL1 field if non-nil, zero value otherwise.

### GetHysteresisL1Ok

`func (o *StreamInfoTypeAdditionalInfo) GetHysteresisL1Ok() (*int32, bool)`

GetHysteresisL1Ok returns a tuple with the HysteresisL1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresisL1

`func (o *StreamInfoTypeAdditionalInfo) SetHysteresisL1(v int32)`

SetHysteresisL1 sets HysteresisL1 field to given value.

### HasHysteresisL1

`func (o *StreamInfoTypeAdditionalInfo) HasHysteresisL1() bool`

HasHysteresisL1 returns a boolean if a field has been set.

### GetTimeToTriggerL1

`func (o *StreamInfoTypeAdditionalInfo) GetTimeToTriggerL1() TimeToTriggerL1Type`

GetTimeToTriggerL1 returns the TimeToTriggerL1 field if non-nil, zero value otherwise.

### GetTimeToTriggerL1Ok

`func (o *StreamInfoTypeAdditionalInfo) GetTimeToTriggerL1Ok() (*TimeToTriggerL1Type, bool)`

GetTimeToTriggerL1Ok returns a tuple with the TimeToTriggerL1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToTriggerL1

`func (o *StreamInfoTypeAdditionalInfo) SetTimeToTriggerL1(v TimeToTriggerL1Type)`

SetTimeToTriggerL1 sets TimeToTriggerL1 field to given value.

### HasTimeToTriggerL1

`func (o *StreamInfoTypeAdditionalInfo) HasTimeToTriggerL1() bool`

HasTimeToTriggerL1 returns a boolean if a field has been set.

### GetMBSFNAreaList

`func (o *StreamInfoTypeAdditionalInfo) GetMBSFNAreaList() []MbsfnArea`

GetMBSFNAreaList returns the MBSFNAreaList field if non-nil, zero value otherwise.

### GetMBSFNAreaListOk

`func (o *StreamInfoTypeAdditionalInfo) GetMBSFNAreaListOk() (*[]MbsfnArea, bool)`

GetMBSFNAreaListOk returns a tuple with the MBSFNAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMBSFNAreaList

`func (o *StreamInfoTypeAdditionalInfo) SetMBSFNAreaList(v []MbsfnArea)`

SetMBSFNAreaList sets MBSFNAreaList field to given value.

### HasMBSFNAreaList

`func (o *StreamInfoTypeAdditionalInfo) HasMBSFNAreaList() bool`

HasMBSFNAreaList returns a boolean if a field has been set.

### GetMeasurementPeriodLTE

`func (o *StreamInfoTypeAdditionalInfo) GetMeasurementPeriodLTE() MeasurementPeriodLTEType`

GetMeasurementPeriodLTE returns the MeasurementPeriodLTE field if non-nil, zero value otherwise.

### GetMeasurementPeriodLTEOk

`func (o *StreamInfoTypeAdditionalInfo) GetMeasurementPeriodLTEOk() (*MeasurementPeriodLTEType, bool)`

GetMeasurementPeriodLTEOk returns a tuple with the MeasurementPeriodLTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriodLTE

`func (o *StreamInfoTypeAdditionalInfo) SetMeasurementPeriodLTE(v MeasurementPeriodLTEType)`

SetMeasurementPeriodLTE sets MeasurementPeriodLTE field to given value.

### HasMeasurementPeriodLTE

`func (o *StreamInfoTypeAdditionalInfo) HasMeasurementPeriodLTE() bool`

HasMeasurementPeriodLTE returns a boolean if a field has been set.

### GetMeasurementPeriodUMTS

`func (o *StreamInfoTypeAdditionalInfo) GetMeasurementPeriodUMTS() MeasurementPeriodUMTSType`

GetMeasurementPeriodUMTS returns the MeasurementPeriodUMTS field if non-nil, zero value otherwise.

### GetMeasurementPeriodUMTSOk

`func (o *StreamInfoTypeAdditionalInfo) GetMeasurementPeriodUMTSOk() (*MeasurementPeriodUMTSType, bool)`

GetMeasurementPeriodUMTSOk returns a tuple with the MeasurementPeriodUMTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriodUMTS

`func (o *StreamInfoTypeAdditionalInfo) SetMeasurementPeriodUMTS(v MeasurementPeriodUMTSType)`

SetMeasurementPeriodUMTS sets MeasurementPeriodUMTS field to given value.

### HasMeasurementPeriodUMTS

`func (o *StreamInfoTypeAdditionalInfo) HasMeasurementPeriodUMTS() bool`

HasMeasurementPeriodUMTS returns a boolean if a field has been set.

### GetMeasurementQuantity

`func (o *StreamInfoTypeAdditionalInfo) GetMeasurementQuantity() MeasurementQuantityType`

GetMeasurementQuantity returns the MeasurementQuantity field if non-nil, zero value otherwise.

### GetMeasurementQuantityOk

`func (o *StreamInfoTypeAdditionalInfo) GetMeasurementQuantityOk() (*MeasurementQuantityType, bool)`

GetMeasurementQuantityOk returns a tuple with the MeasurementQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementQuantity

`func (o *StreamInfoTypeAdditionalInfo) SetMeasurementQuantity(v MeasurementQuantityType)`

SetMeasurementQuantity sets MeasurementQuantity field to given value.

### HasMeasurementQuantity

`func (o *StreamInfoTypeAdditionalInfo) HasMeasurementQuantity() bool`

HasMeasurementQuantity returns a boolean if a field has been set.

### GetEventThresholdUphUMTS

`func (o *StreamInfoTypeAdditionalInfo) GetEventThresholdUphUMTS() int32`

GetEventThresholdUphUMTS returns the EventThresholdUphUMTS field if non-nil, zero value otherwise.

### GetEventThresholdUphUMTSOk

`func (o *StreamInfoTypeAdditionalInfo) GetEventThresholdUphUMTSOk() (*int32, bool)`

GetEventThresholdUphUMTSOk returns a tuple with the EventThresholdUphUMTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdUphUMTS

`func (o *StreamInfoTypeAdditionalInfo) SetEventThresholdUphUMTS(v int32)`

SetEventThresholdUphUMTS sets EventThresholdUphUMTS field to given value.

### HasEventThresholdUphUMTS

`func (o *StreamInfoTypeAdditionalInfo) HasEventThresholdUphUMTS() bool`

HasEventThresholdUphUMTS returns a boolean if a field has been set.

### GetPLMNList

`func (o *StreamInfoTypeAdditionalInfo) GetPLMNList() []PLMNListTypeInner`

GetPLMNList returns the PLMNList field if non-nil, zero value otherwise.

### GetPLMNListOk

`func (o *StreamInfoTypeAdditionalInfo) GetPLMNListOk() (*[]PLMNListTypeInner, bool)`

GetPLMNListOk returns a tuple with the PLMNList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPLMNList

`func (o *StreamInfoTypeAdditionalInfo) SetPLMNList(v []PLMNListTypeInner)`

SetPLMNList sets PLMNList field to given value.

### HasPLMNList

`func (o *StreamInfoTypeAdditionalInfo) HasPLMNList() bool`

HasPLMNList returns a boolean if a field has been set.

### GetPositioningMethod

`func (o *StreamInfoTypeAdditionalInfo) GetPositioningMethod() PositioningMethodType`

GetPositioningMethod returns the PositioningMethod field if non-nil, zero value otherwise.

### GetPositioningMethodOk

`func (o *StreamInfoTypeAdditionalInfo) GetPositioningMethodOk() (*PositioningMethodType, bool)`

GetPositioningMethodOk returns a tuple with the PositioningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningMethod

`func (o *StreamInfoTypeAdditionalInfo) SetPositioningMethod(v PositioningMethodType)`

SetPositioningMethod sets PositioningMethod field to given value.

### HasPositioningMethod

`func (o *StreamInfoTypeAdditionalInfo) HasPositioningMethod() bool`

HasPositioningMethod returns a boolean if a field has been set.

### GetReportAmount

`func (o *StreamInfoTypeAdditionalInfo) GetReportAmount() ReportAmountType`

GetReportAmount returns the ReportAmount field if non-nil, zero value otherwise.

### GetReportAmountOk

`func (o *StreamInfoTypeAdditionalInfo) GetReportAmountOk() (*ReportAmountType, bool)`

GetReportAmountOk returns a tuple with the ReportAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAmount

`func (o *StreamInfoTypeAdditionalInfo) SetReportAmount(v ReportAmountType)`

SetReportAmount sets ReportAmount field to given value.

### HasReportAmount

`func (o *StreamInfoTypeAdditionalInfo) HasReportAmount() bool`

HasReportAmount returns a boolean if a field has been set.

### GetReportingTrigger

`func (o *StreamInfoTypeAdditionalInfo) GetReportingTrigger() []string`

GetReportingTrigger returns the ReportingTrigger field if non-nil, zero value otherwise.

### GetReportingTriggerOk

`func (o *StreamInfoTypeAdditionalInfo) GetReportingTriggerOk() (*[]string, bool)`

GetReportingTriggerOk returns a tuple with the ReportingTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTrigger

`func (o *StreamInfoTypeAdditionalInfo) SetReportingTrigger(v []string)`

SetReportingTrigger sets ReportingTrigger field to given value.

### HasReportingTrigger

`func (o *StreamInfoTypeAdditionalInfo) HasReportingTrigger() bool`

HasReportingTrigger returns a boolean if a field has been set.

### GetReportInterval

`func (o *StreamInfoTypeAdditionalInfo) GetReportInterval() ReportIntervalType`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *StreamInfoTypeAdditionalInfo) GetReportIntervalOk() (*ReportIntervalType, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *StreamInfoTypeAdditionalInfo) SetReportInterval(v ReportIntervalType)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *StreamInfoTypeAdditionalInfo) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetReportType

`func (o *StreamInfoTypeAdditionalInfo) GetReportType() ReportTypeType`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *StreamInfoTypeAdditionalInfo) GetReportTypeOk() (*ReportTypeType, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *StreamInfoTypeAdditionalInfo) SetReportType(v ReportTypeType)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *StreamInfoTypeAdditionalInfo) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetSensorInformation

`func (o *StreamInfoTypeAdditionalInfo) GetSensorInformation() []string`

GetSensorInformation returns the SensorInformation field if non-nil, zero value otherwise.

### GetSensorInformationOk

`func (o *StreamInfoTypeAdditionalInfo) GetSensorInformationOk() (*[]string, bool)`

GetSensorInformationOk returns a tuple with the SensorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorInformation

`func (o *StreamInfoTypeAdditionalInfo) SetSensorInformation(v []string)`

SetSensorInformation sets SensorInformation field to given value.

### HasSensorInformation

`func (o *StreamInfoTypeAdditionalInfo) HasSensorInformation() bool`

HasSensorInformation returns a boolean if a field has been set.

### GetTraceCollectionEntityId

`func (o *StreamInfoTypeAdditionalInfo) GetTraceCollectionEntityId() int32`

GetTraceCollectionEntityId returns the TraceCollectionEntityId field if non-nil, zero value otherwise.

### GetTraceCollectionEntityIdOk

`func (o *StreamInfoTypeAdditionalInfo) GetTraceCollectionEntityIdOk() (*int32, bool)`

GetTraceCollectionEntityIdOk returns a tuple with the TraceCollectionEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceCollectionEntityId

`func (o *StreamInfoTypeAdditionalInfo) SetTraceCollectionEntityId(v int32)`

SetTraceCollectionEntityId sets TraceCollectionEntityId field to given value.

### HasTraceCollectionEntityId

`func (o *StreamInfoTypeAdditionalInfo) HasTraceCollectionEntityId() bool`

HasTraceCollectionEntityId returns a boolean if a field has been set.

### GetMeasObjDn

`func (o *StreamInfoTypeAdditionalInfo) GetMeasObjDn() MeasObjDnType`

GetMeasObjDn returns the MeasObjDn field if non-nil, zero value otherwise.

### GetMeasObjDnOk

`func (o *StreamInfoTypeAdditionalInfo) GetMeasObjDnOk() (*MeasObjDnType, bool)`

GetMeasObjDnOk returns a tuple with the MeasObjDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasObjDn

`func (o *StreamInfoTypeAdditionalInfo) SetMeasObjDn(v MeasObjDnType)`

SetMeasObjDn sets MeasObjDn field to given value.


### GetPerformanceMetrics

`func (o *StreamInfoTypeAdditionalInfo) GetPerformanceMetrics() []string`

GetPerformanceMetrics returns the PerformanceMetrics field if non-nil, zero value otherwise.

### GetPerformanceMetricsOk

`func (o *StreamInfoTypeAdditionalInfo) GetPerformanceMetricsOk() (*[]string, bool)`

GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceMetrics

`func (o *StreamInfoTypeAdditionalInfo) SetPerformanceMetrics(v []string)`

SetPerformanceMetrics sets PerformanceMetrics field to given value.


### GetActivityDetails

`func (o *StreamInfoTypeAdditionalInfo) GetActivityDetails() string`

GetActivityDetails returns the ActivityDetails field if non-nil, zero value otherwise.

### GetActivityDetailsOk

`func (o *StreamInfoTypeAdditionalInfo) GetActivityDetailsOk() (*string, bool)`

GetActivityDetailsOk returns a tuple with the ActivityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDetails

`func (o *StreamInfoTypeAdditionalInfo) SetActivityDetails(v string)`

SetActivityDetails sets ActivityDetails field to given value.

### HasActivityDetails

`func (o *StreamInfoTypeAdditionalInfo) HasActivityDetails() bool`

HasActivityDetails returns a boolean if a field has been set.

### GetVsDataType

`func (o *StreamInfoTypeAdditionalInfo) GetVsDataType() string`

GetVsDataType returns the VsDataType field if non-nil, zero value otherwise.

### GetVsDataTypeOk

`func (o *StreamInfoTypeAdditionalInfo) GetVsDataTypeOk() (*string, bool)`

GetVsDataTypeOk returns a tuple with the VsDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataType

`func (o *StreamInfoTypeAdditionalInfo) SetVsDataType(v string)`

SetVsDataType sets VsDataType field to given value.

### HasVsDataType

`func (o *StreamInfoTypeAdditionalInfo) HasVsDataType() bool`

HasVsDataType returns a boolean if a field has been set.

### GetVsData

`func (o *StreamInfoTypeAdditionalInfo) GetVsData() string`

GetVsData returns the VsData field if non-nil, zero value otherwise.

### GetVsDataOk

`func (o *StreamInfoTypeAdditionalInfo) GetVsDataOk() (*string, bool)`

GetVsDataOk returns a tuple with the VsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsData

`func (o *StreamInfoTypeAdditionalInfo) SetVsData(v string)`

SetVsData sets VsData field to given value.

### HasVsData

`func (o *StreamInfoTypeAdditionalInfo) HasVsData() bool`

HasVsData returns a boolean if a field has been set.

### GetVsDataFormatVersion

`func (o *StreamInfoTypeAdditionalInfo) GetVsDataFormatVersion() string`

GetVsDataFormatVersion returns the VsDataFormatVersion field if non-nil, zero value otherwise.

### GetVsDataFormatVersionOk

`func (o *StreamInfoTypeAdditionalInfo) GetVsDataFormatVersionOk() (*string, bool)`

GetVsDataFormatVersionOk returns a tuple with the VsDataFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataFormatVersion

`func (o *StreamInfoTypeAdditionalInfo) SetVsDataFormatVersion(v string)`

SetVsDataFormatVersion sets VsDataFormatVersion field to given value.

### HasVsDataFormatVersion

`func (o *StreamInfoTypeAdditionalInfo) HasVsDataFormatVersion() bool`

HasVsDataFormatVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


