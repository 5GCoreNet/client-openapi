# ProseChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnouncingPlmnID** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AnnouncingUeHplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AnnouncingUeVplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**MonitoringUeHplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**MonitoringUeVplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**DiscovererUeHplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**DiscovererUeVplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**DiscovereeUeHplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**DiscovereeUeVplmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**MonitoredPlmnIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ProseApplicationID** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationSpecificDataList** | Pointer to **[]string** |  | [optional] 
**ProseFunctionality** | Pointer to [**ProseFunctionality**](ProseFunctionality.md) |  | [optional] 
**ProseEventType** | Pointer to [**ProseEventType**](ProseEventType.md) |  | [optional] 
**DirectDiscoveryModel** | Pointer to [**DirectDiscoveryModel**](DirectDiscoveryModel.md) |  | [optional] 
**ValidityPeriod** | Pointer to **int32** |  | [optional] 
**RoleOfUE** | Pointer to [**RoleOfUE**](RoleOfUE.md) |  | [optional] 
**ProseRequestTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PC3ProtocolCause** | Pointer to **int32** |  | [optional] 
**MonitoringUEIdentifier** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**RequestedPLMNIdentifier** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**TimeWindow** | Pointer to **int32** |  | [optional] 
**RangeClass** | Pointer to [**RangeClass**](RangeClass.md) |  | [optional] 
**ProximityAlertIndication** | Pointer to **bool** |  | [optional] 
**ProximityAlertTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ProximityCancellationTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RelayIPAddress** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**ProseUEToNetworkRelayUEID** | Pointer to **string** |  | [optional] 
**ProseDestinationLayer2ID** | Pointer to **string** |  | [optional] 
**PFIContainerInformation** | Pointer to [**[]PFIContainerInformation**](PFIContainerInformation.md) |  | [optional] 
**TransmissionDataContainer** | Pointer to [**[]PC5DataContainer**](PC5DataContainer.md) |  | [optional] 
**ReceptionDataContainer** | Pointer to [**[]PC5DataContainer**](PC5DataContainer.md) |  | [optional] 

## Methods

### NewProseChargingInformation

`func NewProseChargingInformation() *ProseChargingInformation`

NewProseChargingInformation instantiates a new ProseChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProseChargingInformationWithDefaults

`func NewProseChargingInformationWithDefaults() *ProseChargingInformation`

NewProseChargingInformationWithDefaults instantiates a new ProseChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncingPlmnID

`func (o *ProseChargingInformation) GetAnnouncingPlmnID() PlmnId`

GetAnnouncingPlmnID returns the AnnouncingPlmnID field if non-nil, zero value otherwise.

### GetAnnouncingPlmnIDOk

`func (o *ProseChargingInformation) GetAnnouncingPlmnIDOk() (*PlmnId, bool)`

GetAnnouncingPlmnIDOk returns a tuple with the AnnouncingPlmnID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncingPlmnID

`func (o *ProseChargingInformation) SetAnnouncingPlmnID(v PlmnId)`

SetAnnouncingPlmnID sets AnnouncingPlmnID field to given value.

### HasAnnouncingPlmnID

`func (o *ProseChargingInformation) HasAnnouncingPlmnID() bool`

HasAnnouncingPlmnID returns a boolean if a field has been set.

### GetAnnouncingUeHplmnIdentifier

`func (o *ProseChargingInformation) GetAnnouncingUeHplmnIdentifier() PlmnId`

GetAnnouncingUeHplmnIdentifier returns the AnnouncingUeHplmnIdentifier field if non-nil, zero value otherwise.

### GetAnnouncingUeHplmnIdentifierOk

`func (o *ProseChargingInformation) GetAnnouncingUeHplmnIdentifierOk() (*PlmnId, bool)`

GetAnnouncingUeHplmnIdentifierOk returns a tuple with the AnnouncingUeHplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncingUeHplmnIdentifier

`func (o *ProseChargingInformation) SetAnnouncingUeHplmnIdentifier(v PlmnId)`

SetAnnouncingUeHplmnIdentifier sets AnnouncingUeHplmnIdentifier field to given value.

### HasAnnouncingUeHplmnIdentifier

`func (o *ProseChargingInformation) HasAnnouncingUeHplmnIdentifier() bool`

HasAnnouncingUeHplmnIdentifier returns a boolean if a field has been set.

### GetAnnouncingUeVplmnIdentifier

`func (o *ProseChargingInformation) GetAnnouncingUeVplmnIdentifier() PlmnId`

GetAnnouncingUeVplmnIdentifier returns the AnnouncingUeVplmnIdentifier field if non-nil, zero value otherwise.

### GetAnnouncingUeVplmnIdentifierOk

`func (o *ProseChargingInformation) GetAnnouncingUeVplmnIdentifierOk() (*PlmnId, bool)`

GetAnnouncingUeVplmnIdentifierOk returns a tuple with the AnnouncingUeVplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncingUeVplmnIdentifier

`func (o *ProseChargingInformation) SetAnnouncingUeVplmnIdentifier(v PlmnId)`

SetAnnouncingUeVplmnIdentifier sets AnnouncingUeVplmnIdentifier field to given value.

### HasAnnouncingUeVplmnIdentifier

`func (o *ProseChargingInformation) HasAnnouncingUeVplmnIdentifier() bool`

HasAnnouncingUeVplmnIdentifier returns a boolean if a field has been set.

### GetMonitoringUeHplmnIdentifier

`func (o *ProseChargingInformation) GetMonitoringUeHplmnIdentifier() PlmnId`

GetMonitoringUeHplmnIdentifier returns the MonitoringUeHplmnIdentifier field if non-nil, zero value otherwise.

### GetMonitoringUeHplmnIdentifierOk

`func (o *ProseChargingInformation) GetMonitoringUeHplmnIdentifierOk() (*PlmnId, bool)`

GetMonitoringUeHplmnIdentifierOk returns a tuple with the MonitoringUeHplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUeHplmnIdentifier

`func (o *ProseChargingInformation) SetMonitoringUeHplmnIdentifier(v PlmnId)`

SetMonitoringUeHplmnIdentifier sets MonitoringUeHplmnIdentifier field to given value.

### HasMonitoringUeHplmnIdentifier

`func (o *ProseChargingInformation) HasMonitoringUeHplmnIdentifier() bool`

HasMonitoringUeHplmnIdentifier returns a boolean if a field has been set.

### GetMonitoringUeVplmnIdentifier

`func (o *ProseChargingInformation) GetMonitoringUeVplmnIdentifier() PlmnId`

GetMonitoringUeVplmnIdentifier returns the MonitoringUeVplmnIdentifier field if non-nil, zero value otherwise.

### GetMonitoringUeVplmnIdentifierOk

`func (o *ProseChargingInformation) GetMonitoringUeVplmnIdentifierOk() (*PlmnId, bool)`

GetMonitoringUeVplmnIdentifierOk returns a tuple with the MonitoringUeVplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUeVplmnIdentifier

`func (o *ProseChargingInformation) SetMonitoringUeVplmnIdentifier(v PlmnId)`

SetMonitoringUeVplmnIdentifier sets MonitoringUeVplmnIdentifier field to given value.

### HasMonitoringUeVplmnIdentifier

`func (o *ProseChargingInformation) HasMonitoringUeVplmnIdentifier() bool`

HasMonitoringUeVplmnIdentifier returns a boolean if a field has been set.

### GetDiscovererUeHplmnIdentifier

`func (o *ProseChargingInformation) GetDiscovererUeHplmnIdentifier() PlmnId`

GetDiscovererUeHplmnIdentifier returns the DiscovererUeHplmnIdentifier field if non-nil, zero value otherwise.

### GetDiscovererUeHplmnIdentifierOk

`func (o *ProseChargingInformation) GetDiscovererUeHplmnIdentifierOk() (*PlmnId, bool)`

GetDiscovererUeHplmnIdentifierOk returns a tuple with the DiscovererUeHplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovererUeHplmnIdentifier

`func (o *ProseChargingInformation) SetDiscovererUeHplmnIdentifier(v PlmnId)`

SetDiscovererUeHplmnIdentifier sets DiscovererUeHplmnIdentifier field to given value.

### HasDiscovererUeHplmnIdentifier

`func (o *ProseChargingInformation) HasDiscovererUeHplmnIdentifier() bool`

HasDiscovererUeHplmnIdentifier returns a boolean if a field has been set.

### GetDiscovererUeVplmnIdentifier

`func (o *ProseChargingInformation) GetDiscovererUeVplmnIdentifier() PlmnId`

GetDiscovererUeVplmnIdentifier returns the DiscovererUeVplmnIdentifier field if non-nil, zero value otherwise.

### GetDiscovererUeVplmnIdentifierOk

`func (o *ProseChargingInformation) GetDiscovererUeVplmnIdentifierOk() (*PlmnId, bool)`

GetDiscovererUeVplmnIdentifierOk returns a tuple with the DiscovererUeVplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovererUeVplmnIdentifier

`func (o *ProseChargingInformation) SetDiscovererUeVplmnIdentifier(v PlmnId)`

SetDiscovererUeVplmnIdentifier sets DiscovererUeVplmnIdentifier field to given value.

### HasDiscovererUeVplmnIdentifier

`func (o *ProseChargingInformation) HasDiscovererUeVplmnIdentifier() bool`

HasDiscovererUeVplmnIdentifier returns a boolean if a field has been set.

### GetDiscovereeUeHplmnIdentifier

`func (o *ProseChargingInformation) GetDiscovereeUeHplmnIdentifier() PlmnId`

GetDiscovereeUeHplmnIdentifier returns the DiscovereeUeHplmnIdentifier field if non-nil, zero value otherwise.

### GetDiscovereeUeHplmnIdentifierOk

`func (o *ProseChargingInformation) GetDiscovereeUeHplmnIdentifierOk() (*PlmnId, bool)`

GetDiscovereeUeHplmnIdentifierOk returns a tuple with the DiscovereeUeHplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovereeUeHplmnIdentifier

`func (o *ProseChargingInformation) SetDiscovereeUeHplmnIdentifier(v PlmnId)`

SetDiscovereeUeHplmnIdentifier sets DiscovereeUeHplmnIdentifier field to given value.

### HasDiscovereeUeHplmnIdentifier

`func (o *ProseChargingInformation) HasDiscovereeUeHplmnIdentifier() bool`

HasDiscovereeUeHplmnIdentifier returns a boolean if a field has been set.

### GetDiscovereeUeVplmnIdentifier

`func (o *ProseChargingInformation) GetDiscovereeUeVplmnIdentifier() PlmnId`

GetDiscovereeUeVplmnIdentifier returns the DiscovereeUeVplmnIdentifier field if non-nil, zero value otherwise.

### GetDiscovereeUeVplmnIdentifierOk

`func (o *ProseChargingInformation) GetDiscovereeUeVplmnIdentifierOk() (*PlmnId, bool)`

GetDiscovereeUeVplmnIdentifierOk returns a tuple with the DiscovereeUeVplmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovereeUeVplmnIdentifier

`func (o *ProseChargingInformation) SetDiscovereeUeVplmnIdentifier(v PlmnId)`

SetDiscovereeUeVplmnIdentifier sets DiscovereeUeVplmnIdentifier field to given value.

### HasDiscovereeUeVplmnIdentifier

`func (o *ProseChargingInformation) HasDiscovereeUeVplmnIdentifier() bool`

HasDiscovereeUeVplmnIdentifier returns a boolean if a field has been set.

### GetMonitoredPlmnIdentifier

`func (o *ProseChargingInformation) GetMonitoredPlmnIdentifier() PlmnId`

GetMonitoredPlmnIdentifier returns the MonitoredPlmnIdentifier field if non-nil, zero value otherwise.

### GetMonitoredPlmnIdentifierOk

`func (o *ProseChargingInformation) GetMonitoredPlmnIdentifierOk() (*PlmnId, bool)`

GetMonitoredPlmnIdentifierOk returns a tuple with the MonitoredPlmnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredPlmnIdentifier

`func (o *ProseChargingInformation) SetMonitoredPlmnIdentifier(v PlmnId)`

SetMonitoredPlmnIdentifier sets MonitoredPlmnIdentifier field to given value.

### HasMonitoredPlmnIdentifier

`func (o *ProseChargingInformation) HasMonitoredPlmnIdentifier() bool`

HasMonitoredPlmnIdentifier returns a boolean if a field has been set.

### GetProseApplicationID

`func (o *ProseChargingInformation) GetProseApplicationID() string`

GetProseApplicationID returns the ProseApplicationID field if non-nil, zero value otherwise.

### GetProseApplicationIDOk

`func (o *ProseChargingInformation) GetProseApplicationIDOk() (*string, bool)`

GetProseApplicationIDOk returns a tuple with the ProseApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseApplicationID

`func (o *ProseChargingInformation) SetProseApplicationID(v string)`

SetProseApplicationID sets ProseApplicationID field to given value.

### HasProseApplicationID

`func (o *ProseChargingInformation) HasProseApplicationID() bool`

HasProseApplicationID returns a boolean if a field has been set.

### GetApplicationId

`func (o *ProseChargingInformation) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ProseChargingInformation) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ProseChargingInformation) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ProseChargingInformation) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationSpecificDataList

`func (o *ProseChargingInformation) GetApplicationSpecificDataList() []string`

GetApplicationSpecificDataList returns the ApplicationSpecificDataList field if non-nil, zero value otherwise.

### GetApplicationSpecificDataListOk

`func (o *ProseChargingInformation) GetApplicationSpecificDataListOk() (*[]string, bool)`

GetApplicationSpecificDataListOk returns a tuple with the ApplicationSpecificDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSpecificDataList

`func (o *ProseChargingInformation) SetApplicationSpecificDataList(v []string)`

SetApplicationSpecificDataList sets ApplicationSpecificDataList field to given value.

### HasApplicationSpecificDataList

`func (o *ProseChargingInformation) HasApplicationSpecificDataList() bool`

HasApplicationSpecificDataList returns a boolean if a field has been set.

### GetProseFunctionality

`func (o *ProseChargingInformation) GetProseFunctionality() ProseFunctionality`

GetProseFunctionality returns the ProseFunctionality field if non-nil, zero value otherwise.

### GetProseFunctionalityOk

`func (o *ProseChargingInformation) GetProseFunctionalityOk() (*ProseFunctionality, bool)`

GetProseFunctionalityOk returns a tuple with the ProseFunctionality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseFunctionality

`func (o *ProseChargingInformation) SetProseFunctionality(v ProseFunctionality)`

SetProseFunctionality sets ProseFunctionality field to given value.

### HasProseFunctionality

`func (o *ProseChargingInformation) HasProseFunctionality() bool`

HasProseFunctionality returns a boolean if a field has been set.

### GetProseEventType

`func (o *ProseChargingInformation) GetProseEventType() ProseEventType`

GetProseEventType returns the ProseEventType field if non-nil, zero value otherwise.

### GetProseEventTypeOk

`func (o *ProseChargingInformation) GetProseEventTypeOk() (*ProseEventType, bool)`

GetProseEventTypeOk returns a tuple with the ProseEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseEventType

`func (o *ProseChargingInformation) SetProseEventType(v ProseEventType)`

SetProseEventType sets ProseEventType field to given value.

### HasProseEventType

`func (o *ProseChargingInformation) HasProseEventType() bool`

HasProseEventType returns a boolean if a field has been set.

### GetDirectDiscoveryModel

`func (o *ProseChargingInformation) GetDirectDiscoveryModel() DirectDiscoveryModel`

GetDirectDiscoveryModel returns the DirectDiscoveryModel field if non-nil, zero value otherwise.

### GetDirectDiscoveryModelOk

`func (o *ProseChargingInformation) GetDirectDiscoveryModelOk() (*DirectDiscoveryModel, bool)`

GetDirectDiscoveryModelOk returns a tuple with the DirectDiscoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDiscoveryModel

`func (o *ProseChargingInformation) SetDirectDiscoveryModel(v DirectDiscoveryModel)`

SetDirectDiscoveryModel sets DirectDiscoveryModel field to given value.

### HasDirectDiscoveryModel

`func (o *ProseChargingInformation) HasDirectDiscoveryModel() bool`

HasDirectDiscoveryModel returns a boolean if a field has been set.

### GetValidityPeriod

`func (o *ProseChargingInformation) GetValidityPeriod() int32`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *ProseChargingInformation) GetValidityPeriodOk() (*int32, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *ProseChargingInformation) SetValidityPeriod(v int32)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *ProseChargingInformation) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### GetRoleOfUE

`func (o *ProseChargingInformation) GetRoleOfUE() RoleOfUE`

GetRoleOfUE returns the RoleOfUE field if non-nil, zero value otherwise.

### GetRoleOfUEOk

`func (o *ProseChargingInformation) GetRoleOfUEOk() (*RoleOfUE, bool)`

GetRoleOfUEOk returns a tuple with the RoleOfUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleOfUE

`func (o *ProseChargingInformation) SetRoleOfUE(v RoleOfUE)`

SetRoleOfUE sets RoleOfUE field to given value.

### HasRoleOfUE

`func (o *ProseChargingInformation) HasRoleOfUE() bool`

HasRoleOfUE returns a boolean if a field has been set.

### GetProseRequestTimestamp

`func (o *ProseChargingInformation) GetProseRequestTimestamp() time.Time`

GetProseRequestTimestamp returns the ProseRequestTimestamp field if non-nil, zero value otherwise.

### GetProseRequestTimestampOk

`func (o *ProseChargingInformation) GetProseRequestTimestampOk() (*time.Time, bool)`

GetProseRequestTimestampOk returns a tuple with the ProseRequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseRequestTimestamp

`func (o *ProseChargingInformation) SetProseRequestTimestamp(v time.Time)`

SetProseRequestTimestamp sets ProseRequestTimestamp field to given value.

### HasProseRequestTimestamp

`func (o *ProseChargingInformation) HasProseRequestTimestamp() bool`

HasProseRequestTimestamp returns a boolean if a field has been set.

### GetPC3ProtocolCause

`func (o *ProseChargingInformation) GetPC3ProtocolCause() int32`

GetPC3ProtocolCause returns the PC3ProtocolCause field if non-nil, zero value otherwise.

### GetPC3ProtocolCauseOk

`func (o *ProseChargingInformation) GetPC3ProtocolCauseOk() (*int32, bool)`

GetPC3ProtocolCauseOk returns a tuple with the PC3ProtocolCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPC3ProtocolCause

`func (o *ProseChargingInformation) SetPC3ProtocolCause(v int32)`

SetPC3ProtocolCause sets PC3ProtocolCause field to given value.

### HasPC3ProtocolCause

`func (o *ProseChargingInformation) HasPC3ProtocolCause() bool`

HasPC3ProtocolCause returns a boolean if a field has been set.

### GetMonitoringUEIdentifier

`func (o *ProseChargingInformation) GetMonitoringUEIdentifier() string`

GetMonitoringUEIdentifier returns the MonitoringUEIdentifier field if non-nil, zero value otherwise.

### GetMonitoringUEIdentifierOk

`func (o *ProseChargingInformation) GetMonitoringUEIdentifierOk() (*string, bool)`

GetMonitoringUEIdentifierOk returns a tuple with the MonitoringUEIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUEIdentifier

`func (o *ProseChargingInformation) SetMonitoringUEIdentifier(v string)`

SetMonitoringUEIdentifier sets MonitoringUEIdentifier field to given value.

### HasMonitoringUEIdentifier

`func (o *ProseChargingInformation) HasMonitoringUEIdentifier() bool`

HasMonitoringUEIdentifier returns a boolean if a field has been set.

### GetRequestedPLMNIdentifier

`func (o *ProseChargingInformation) GetRequestedPLMNIdentifier() PlmnId`

GetRequestedPLMNIdentifier returns the RequestedPLMNIdentifier field if non-nil, zero value otherwise.

### GetRequestedPLMNIdentifierOk

`func (o *ProseChargingInformation) GetRequestedPLMNIdentifierOk() (*PlmnId, bool)`

GetRequestedPLMNIdentifierOk returns a tuple with the RequestedPLMNIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPLMNIdentifier

`func (o *ProseChargingInformation) SetRequestedPLMNIdentifier(v PlmnId)`

SetRequestedPLMNIdentifier sets RequestedPLMNIdentifier field to given value.

### HasRequestedPLMNIdentifier

`func (o *ProseChargingInformation) HasRequestedPLMNIdentifier() bool`

HasRequestedPLMNIdentifier returns a boolean if a field has been set.

### GetTimeWindow

`func (o *ProseChargingInformation) GetTimeWindow() int32`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *ProseChargingInformation) GetTimeWindowOk() (*int32, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *ProseChargingInformation) SetTimeWindow(v int32)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *ProseChargingInformation) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetRangeClass

`func (o *ProseChargingInformation) GetRangeClass() RangeClass`

GetRangeClass returns the RangeClass field if non-nil, zero value otherwise.

### GetRangeClassOk

`func (o *ProseChargingInformation) GetRangeClassOk() (*RangeClass, bool)`

GetRangeClassOk returns a tuple with the RangeClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClass

`func (o *ProseChargingInformation) SetRangeClass(v RangeClass)`

SetRangeClass sets RangeClass field to given value.

### HasRangeClass

`func (o *ProseChargingInformation) HasRangeClass() bool`

HasRangeClass returns a boolean if a field has been set.

### GetProximityAlertIndication

`func (o *ProseChargingInformation) GetProximityAlertIndication() bool`

GetProximityAlertIndication returns the ProximityAlertIndication field if non-nil, zero value otherwise.

### GetProximityAlertIndicationOk

`func (o *ProseChargingInformation) GetProximityAlertIndicationOk() (*bool, bool)`

GetProximityAlertIndicationOk returns a tuple with the ProximityAlertIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityAlertIndication

`func (o *ProseChargingInformation) SetProximityAlertIndication(v bool)`

SetProximityAlertIndication sets ProximityAlertIndication field to given value.

### HasProximityAlertIndication

`func (o *ProseChargingInformation) HasProximityAlertIndication() bool`

HasProximityAlertIndication returns a boolean if a field has been set.

### GetProximityAlertTimestamp

`func (o *ProseChargingInformation) GetProximityAlertTimestamp() time.Time`

GetProximityAlertTimestamp returns the ProximityAlertTimestamp field if non-nil, zero value otherwise.

### GetProximityAlertTimestampOk

`func (o *ProseChargingInformation) GetProximityAlertTimestampOk() (*time.Time, bool)`

GetProximityAlertTimestampOk returns a tuple with the ProximityAlertTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityAlertTimestamp

`func (o *ProseChargingInformation) SetProximityAlertTimestamp(v time.Time)`

SetProximityAlertTimestamp sets ProximityAlertTimestamp field to given value.

### HasProximityAlertTimestamp

`func (o *ProseChargingInformation) HasProximityAlertTimestamp() bool`

HasProximityAlertTimestamp returns a boolean if a field has been set.

### GetProximityCancellationTimestamp

`func (o *ProseChargingInformation) GetProximityCancellationTimestamp() time.Time`

GetProximityCancellationTimestamp returns the ProximityCancellationTimestamp field if non-nil, zero value otherwise.

### GetProximityCancellationTimestampOk

`func (o *ProseChargingInformation) GetProximityCancellationTimestampOk() (*time.Time, bool)`

GetProximityCancellationTimestampOk returns a tuple with the ProximityCancellationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityCancellationTimestamp

`func (o *ProseChargingInformation) SetProximityCancellationTimestamp(v time.Time)`

SetProximityCancellationTimestamp sets ProximityCancellationTimestamp field to given value.

### HasProximityCancellationTimestamp

`func (o *ProseChargingInformation) HasProximityCancellationTimestamp() bool`

HasProximityCancellationTimestamp returns a boolean if a field has been set.

### GetRelayIPAddress

`func (o *ProseChargingInformation) GetRelayIPAddress() IpAddr`

GetRelayIPAddress returns the RelayIPAddress field if non-nil, zero value otherwise.

### GetRelayIPAddressOk

`func (o *ProseChargingInformation) GetRelayIPAddressOk() (*IpAddr, bool)`

GetRelayIPAddressOk returns a tuple with the RelayIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayIPAddress

`func (o *ProseChargingInformation) SetRelayIPAddress(v IpAddr)`

SetRelayIPAddress sets RelayIPAddress field to given value.

### HasRelayIPAddress

`func (o *ProseChargingInformation) HasRelayIPAddress() bool`

HasRelayIPAddress returns a boolean if a field has been set.

### GetProseUEToNetworkRelayUEID

`func (o *ProseChargingInformation) GetProseUEToNetworkRelayUEID() string`

GetProseUEToNetworkRelayUEID returns the ProseUEToNetworkRelayUEID field if non-nil, zero value otherwise.

### GetProseUEToNetworkRelayUEIDOk

`func (o *ProseChargingInformation) GetProseUEToNetworkRelayUEIDOk() (*string, bool)`

GetProseUEToNetworkRelayUEIDOk returns a tuple with the ProseUEToNetworkRelayUEID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseUEToNetworkRelayUEID

`func (o *ProseChargingInformation) SetProseUEToNetworkRelayUEID(v string)`

SetProseUEToNetworkRelayUEID sets ProseUEToNetworkRelayUEID field to given value.

### HasProseUEToNetworkRelayUEID

`func (o *ProseChargingInformation) HasProseUEToNetworkRelayUEID() bool`

HasProseUEToNetworkRelayUEID returns a boolean if a field has been set.

### GetProseDestinationLayer2ID

`func (o *ProseChargingInformation) GetProseDestinationLayer2ID() string`

GetProseDestinationLayer2ID returns the ProseDestinationLayer2ID field if non-nil, zero value otherwise.

### GetProseDestinationLayer2IDOk

`func (o *ProseChargingInformation) GetProseDestinationLayer2IDOk() (*string, bool)`

GetProseDestinationLayer2IDOk returns a tuple with the ProseDestinationLayer2ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseDestinationLayer2ID

`func (o *ProseChargingInformation) SetProseDestinationLayer2ID(v string)`

SetProseDestinationLayer2ID sets ProseDestinationLayer2ID field to given value.

### HasProseDestinationLayer2ID

`func (o *ProseChargingInformation) HasProseDestinationLayer2ID() bool`

HasProseDestinationLayer2ID returns a boolean if a field has been set.

### GetPFIContainerInformation

`func (o *ProseChargingInformation) GetPFIContainerInformation() []PFIContainerInformation`

GetPFIContainerInformation returns the PFIContainerInformation field if non-nil, zero value otherwise.

### GetPFIContainerInformationOk

`func (o *ProseChargingInformation) GetPFIContainerInformationOk() (*[]PFIContainerInformation, bool)`

GetPFIContainerInformationOk returns a tuple with the PFIContainerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPFIContainerInformation

`func (o *ProseChargingInformation) SetPFIContainerInformation(v []PFIContainerInformation)`

SetPFIContainerInformation sets PFIContainerInformation field to given value.

### HasPFIContainerInformation

`func (o *ProseChargingInformation) HasPFIContainerInformation() bool`

HasPFIContainerInformation returns a boolean if a field has been set.

### GetTransmissionDataContainer

`func (o *ProseChargingInformation) GetTransmissionDataContainer() []PC5DataContainer`

GetTransmissionDataContainer returns the TransmissionDataContainer field if non-nil, zero value otherwise.

### GetTransmissionDataContainerOk

`func (o *ProseChargingInformation) GetTransmissionDataContainerOk() (*[]PC5DataContainer, bool)`

GetTransmissionDataContainerOk returns a tuple with the TransmissionDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionDataContainer

`func (o *ProseChargingInformation) SetTransmissionDataContainer(v []PC5DataContainer)`

SetTransmissionDataContainer sets TransmissionDataContainer field to given value.

### HasTransmissionDataContainer

`func (o *ProseChargingInformation) HasTransmissionDataContainer() bool`

HasTransmissionDataContainer returns a boolean if a field has been set.

### GetReceptionDataContainer

`func (o *ProseChargingInformation) GetReceptionDataContainer() []PC5DataContainer`

GetReceptionDataContainer returns the ReceptionDataContainer field if non-nil, zero value otherwise.

### GetReceptionDataContainerOk

`func (o *ProseChargingInformation) GetReceptionDataContainerOk() (*[]PC5DataContainer, bool)`

GetReceptionDataContainerOk returns a tuple with the ReceptionDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptionDataContainer

`func (o *ProseChargingInformation) SetReceptionDataContainer(v []PC5DataContainer)`

SetReceptionDataContainer sets ReceptionDataContainer field to given value.

### HasReceptionDataContainer

`func (o *ProseChargingInformation) HasReceptionDataContainer() bool`

HasReceptionDataContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


