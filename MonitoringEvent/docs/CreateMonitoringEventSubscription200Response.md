# CreateMonitoringEventSubscription200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImeiChange** | Pointer to [**AssociationType**](AssociationType.md) |  | [optional] 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**IdleStatusInfo** | Pointer to [**IdleStatusInfo**](IdleStatusInfo.md) |  | [optional] 
**LocationInfo** | Pointer to [**LocationInfo**](LocationInfo.md) |  | [optional] 
**LocFailureCause** | Pointer to [**LocationFailureCause**](LocationFailureCause.md) |  | [optional] 
**LossOfConnectReason** | Pointer to **int32** | If \&quot;monitoringType\&quot; is \&quot;LOSS_OF_CONNECTIVITY\&quot;, this parameter shall be included if available to identify the reason why loss of connectivity is reported. Refer to 3GPP TS 29.336 clause 8.4.58. | [optional] 
**MaxUEAvailabilityTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**MonitoringType** | [**MonitoringType**](MonitoringType.md) |  | 
**UePerLocationReport** | Pointer to [**UePerLocationReport**](UePerLocationReport.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ReachabilityType** | Pointer to [**ReachabilityType**](ReachabilityType.md) |  | [optional] 
**RoamingStatus** | Pointer to **bool** | If \&quot;monitoringType\&quot; is \&quot;ROAMING_STATUS\&quot;, this parameter shall be set to \&quot;true\&quot; if the UE is on roaming status. Set to false or omitted otherwise. | [optional] 
**FailureCause** | Pointer to [**FailureCause**](FailureCause.md) |  | [optional] 
**EventTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**PdnConnInfoList** | Pointer to [**[]PdnConnectionInformation**](PdnConnectionInformation.md) |  | [optional] 
**DddStatus** | Pointer to [**DlDataDeliveryStatus**](DlDataDeliveryStatus.md) |  | [optional] 
**DddTrafDescriptor** | Pointer to [**DddTrafficDescriptor**](DddTrafficDescriptor.md) |  | [optional] 
**MaxWaitTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**ApiCaps** | Pointer to [**[]ApiCapabilityInfo**](ApiCapabilityInfo.md) |  | [optional] 
**NSStatusInfo** | Pointer to [**SACEventStatus**](SACEventStatus.md) |  | [optional] 
**AfServiceId** | Pointer to **string** |  | [optional] 
**ServLevelDevId** | Pointer to **string** | If \&quot;monitoringType\&quot; is \&quot;AREA_OF_INTEREST\&quot;, this parameter may be included to identify the UAV. | [optional] 
**UavPresInd** | Pointer to **bool** | If \&quot;monitoringType\&quot; is \&quot;AREA_OF_INTEREST\&quot;, this parameter shall be set to true if the specified UAV is in the monitoring area. Set to false or omitted otherwise. | [optional] 
**MonitoringEventReports** | [**[]MonitoringEventReport**](MonitoringEventReport.md) |  | 

## Methods

### NewCreateMonitoringEventSubscription200Response

`func NewCreateMonitoringEventSubscription200Response(monitoringType MonitoringType, monitoringEventReports []MonitoringEventReport, ) *CreateMonitoringEventSubscription200Response`

NewCreateMonitoringEventSubscription200Response instantiates a new CreateMonitoringEventSubscription200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMonitoringEventSubscription200ResponseWithDefaults

`func NewCreateMonitoringEventSubscription200ResponseWithDefaults() *CreateMonitoringEventSubscription200Response`

NewCreateMonitoringEventSubscription200ResponseWithDefaults instantiates a new CreateMonitoringEventSubscription200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImeiChange

`func (o *CreateMonitoringEventSubscription200Response) GetImeiChange() AssociationType`

GetImeiChange returns the ImeiChange field if non-nil, zero value otherwise.

### GetImeiChangeOk

`func (o *CreateMonitoringEventSubscription200Response) GetImeiChangeOk() (*AssociationType, bool)`

GetImeiChangeOk returns a tuple with the ImeiChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImeiChange

`func (o *CreateMonitoringEventSubscription200Response) SetImeiChange(v AssociationType)`

SetImeiChange sets ImeiChange field to given value.

### HasImeiChange

`func (o *CreateMonitoringEventSubscription200Response) HasImeiChange() bool`

HasImeiChange returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateMonitoringEventSubscription200Response) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateMonitoringEventSubscription200Response) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateMonitoringEventSubscription200Response) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateMonitoringEventSubscription200Response) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdleStatusInfo

`func (o *CreateMonitoringEventSubscription200Response) GetIdleStatusInfo() IdleStatusInfo`

GetIdleStatusInfo returns the IdleStatusInfo field if non-nil, zero value otherwise.

### GetIdleStatusInfoOk

`func (o *CreateMonitoringEventSubscription200Response) GetIdleStatusInfoOk() (*IdleStatusInfo, bool)`

GetIdleStatusInfoOk returns a tuple with the IdleStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusInfo

`func (o *CreateMonitoringEventSubscription200Response) SetIdleStatusInfo(v IdleStatusInfo)`

SetIdleStatusInfo sets IdleStatusInfo field to given value.

### HasIdleStatusInfo

`func (o *CreateMonitoringEventSubscription200Response) HasIdleStatusInfo() bool`

HasIdleStatusInfo returns a boolean if a field has been set.

### GetLocationInfo

`func (o *CreateMonitoringEventSubscription200Response) GetLocationInfo() LocationInfo`

GetLocationInfo returns the LocationInfo field if non-nil, zero value otherwise.

### GetLocationInfoOk

`func (o *CreateMonitoringEventSubscription200Response) GetLocationInfoOk() (*LocationInfo, bool)`

GetLocationInfoOk returns a tuple with the LocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInfo

`func (o *CreateMonitoringEventSubscription200Response) SetLocationInfo(v LocationInfo)`

SetLocationInfo sets LocationInfo field to given value.

### HasLocationInfo

`func (o *CreateMonitoringEventSubscription200Response) HasLocationInfo() bool`

HasLocationInfo returns a boolean if a field has been set.

### GetLocFailureCause

`func (o *CreateMonitoringEventSubscription200Response) GetLocFailureCause() LocationFailureCause`

GetLocFailureCause returns the LocFailureCause field if non-nil, zero value otherwise.

### GetLocFailureCauseOk

`func (o *CreateMonitoringEventSubscription200Response) GetLocFailureCauseOk() (*LocationFailureCause, bool)`

GetLocFailureCauseOk returns a tuple with the LocFailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocFailureCause

`func (o *CreateMonitoringEventSubscription200Response) SetLocFailureCause(v LocationFailureCause)`

SetLocFailureCause sets LocFailureCause field to given value.

### HasLocFailureCause

`func (o *CreateMonitoringEventSubscription200Response) HasLocFailureCause() bool`

HasLocFailureCause returns a boolean if a field has been set.

### GetLossOfConnectReason

`func (o *CreateMonitoringEventSubscription200Response) GetLossOfConnectReason() int32`

GetLossOfConnectReason returns the LossOfConnectReason field if non-nil, zero value otherwise.

### GetLossOfConnectReasonOk

`func (o *CreateMonitoringEventSubscription200Response) GetLossOfConnectReasonOk() (*int32, bool)`

GetLossOfConnectReasonOk returns a tuple with the LossOfConnectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossOfConnectReason

`func (o *CreateMonitoringEventSubscription200Response) SetLossOfConnectReason(v int32)`

SetLossOfConnectReason sets LossOfConnectReason field to given value.

### HasLossOfConnectReason

`func (o *CreateMonitoringEventSubscription200Response) HasLossOfConnectReason() bool`

HasLossOfConnectReason returns a boolean if a field has been set.

### GetMaxUEAvailabilityTime

`func (o *CreateMonitoringEventSubscription200Response) GetMaxUEAvailabilityTime() time.Time`

GetMaxUEAvailabilityTime returns the MaxUEAvailabilityTime field if non-nil, zero value otherwise.

### GetMaxUEAvailabilityTimeOk

`func (o *CreateMonitoringEventSubscription200Response) GetMaxUEAvailabilityTimeOk() (*time.Time, bool)`

GetMaxUEAvailabilityTimeOk returns a tuple with the MaxUEAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUEAvailabilityTime

`func (o *CreateMonitoringEventSubscription200Response) SetMaxUEAvailabilityTime(v time.Time)`

SetMaxUEAvailabilityTime sets MaxUEAvailabilityTime field to given value.

### HasMaxUEAvailabilityTime

`func (o *CreateMonitoringEventSubscription200Response) HasMaxUEAvailabilityTime() bool`

HasMaxUEAvailabilityTime returns a boolean if a field has been set.

### GetMsisdn

`func (o *CreateMonitoringEventSubscription200Response) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *CreateMonitoringEventSubscription200Response) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *CreateMonitoringEventSubscription200Response) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *CreateMonitoringEventSubscription200Response) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetMonitoringType

`func (o *CreateMonitoringEventSubscription200Response) GetMonitoringType() MonitoringType`

GetMonitoringType returns the MonitoringType field if non-nil, zero value otherwise.

### GetMonitoringTypeOk

`func (o *CreateMonitoringEventSubscription200Response) GetMonitoringTypeOk() (*MonitoringType, bool)`

GetMonitoringTypeOk returns a tuple with the MonitoringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringType

`func (o *CreateMonitoringEventSubscription200Response) SetMonitoringType(v MonitoringType)`

SetMonitoringType sets MonitoringType field to given value.


### GetUePerLocationReport

`func (o *CreateMonitoringEventSubscription200Response) GetUePerLocationReport() UePerLocationReport`

GetUePerLocationReport returns the UePerLocationReport field if non-nil, zero value otherwise.

### GetUePerLocationReportOk

`func (o *CreateMonitoringEventSubscription200Response) GetUePerLocationReportOk() (*UePerLocationReport, bool)`

GetUePerLocationReportOk returns a tuple with the UePerLocationReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePerLocationReport

`func (o *CreateMonitoringEventSubscription200Response) SetUePerLocationReport(v UePerLocationReport)`

SetUePerLocationReport sets UePerLocationReport field to given value.

### HasUePerLocationReport

`func (o *CreateMonitoringEventSubscription200Response) HasUePerLocationReport() bool`

HasUePerLocationReport returns a boolean if a field has been set.

### GetPlmnId

`func (o *CreateMonitoringEventSubscription200Response) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *CreateMonitoringEventSubscription200Response) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *CreateMonitoringEventSubscription200Response) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *CreateMonitoringEventSubscription200Response) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetReachabilityType

`func (o *CreateMonitoringEventSubscription200Response) GetReachabilityType() ReachabilityType`

GetReachabilityType returns the ReachabilityType field if non-nil, zero value otherwise.

### GetReachabilityTypeOk

`func (o *CreateMonitoringEventSubscription200Response) GetReachabilityTypeOk() (*ReachabilityType, bool)`

GetReachabilityTypeOk returns a tuple with the ReachabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityType

`func (o *CreateMonitoringEventSubscription200Response) SetReachabilityType(v ReachabilityType)`

SetReachabilityType sets ReachabilityType field to given value.

### HasReachabilityType

`func (o *CreateMonitoringEventSubscription200Response) HasReachabilityType() bool`

HasReachabilityType returns a boolean if a field has been set.

### GetRoamingStatus

`func (o *CreateMonitoringEventSubscription200Response) GetRoamingStatus() bool`

GetRoamingStatus returns the RoamingStatus field if non-nil, zero value otherwise.

### GetRoamingStatusOk

`func (o *CreateMonitoringEventSubscription200Response) GetRoamingStatusOk() (*bool, bool)`

GetRoamingStatusOk returns a tuple with the RoamingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingStatus

`func (o *CreateMonitoringEventSubscription200Response) SetRoamingStatus(v bool)`

SetRoamingStatus sets RoamingStatus field to given value.

### HasRoamingStatus

`func (o *CreateMonitoringEventSubscription200Response) HasRoamingStatus() bool`

HasRoamingStatus returns a boolean if a field has been set.

### GetFailureCause

`func (o *CreateMonitoringEventSubscription200Response) GetFailureCause() FailureCause`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *CreateMonitoringEventSubscription200Response) GetFailureCauseOk() (*FailureCause, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *CreateMonitoringEventSubscription200Response) SetFailureCause(v FailureCause)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *CreateMonitoringEventSubscription200Response) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetEventTime

`func (o *CreateMonitoringEventSubscription200Response) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *CreateMonitoringEventSubscription200Response) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *CreateMonitoringEventSubscription200Response) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *CreateMonitoringEventSubscription200Response) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetPdnConnInfoList

`func (o *CreateMonitoringEventSubscription200Response) GetPdnConnInfoList() []PdnConnectionInformation`

GetPdnConnInfoList returns the PdnConnInfoList field if non-nil, zero value otherwise.

### GetPdnConnInfoListOk

`func (o *CreateMonitoringEventSubscription200Response) GetPdnConnInfoListOk() (*[]PdnConnectionInformation, bool)`

GetPdnConnInfoListOk returns a tuple with the PdnConnInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnConnInfoList

`func (o *CreateMonitoringEventSubscription200Response) SetPdnConnInfoList(v []PdnConnectionInformation)`

SetPdnConnInfoList sets PdnConnInfoList field to given value.

### HasPdnConnInfoList

`func (o *CreateMonitoringEventSubscription200Response) HasPdnConnInfoList() bool`

HasPdnConnInfoList returns a boolean if a field has been set.

### GetDddStatus

`func (o *CreateMonitoringEventSubscription200Response) GetDddStatus() DlDataDeliveryStatus`

GetDddStatus returns the DddStatus field if non-nil, zero value otherwise.

### GetDddStatusOk

`func (o *CreateMonitoringEventSubscription200Response) GetDddStatusOk() (*DlDataDeliveryStatus, bool)`

GetDddStatusOk returns a tuple with the DddStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddStatus

`func (o *CreateMonitoringEventSubscription200Response) SetDddStatus(v DlDataDeliveryStatus)`

SetDddStatus sets DddStatus field to given value.

### HasDddStatus

`func (o *CreateMonitoringEventSubscription200Response) HasDddStatus() bool`

HasDddStatus returns a boolean if a field has been set.

### GetDddTrafDescriptor

`func (o *CreateMonitoringEventSubscription200Response) GetDddTrafDescriptor() DddTrafficDescriptor`

GetDddTrafDescriptor returns the DddTrafDescriptor field if non-nil, zero value otherwise.

### GetDddTrafDescriptorOk

`func (o *CreateMonitoringEventSubscription200Response) GetDddTrafDescriptorOk() (*DddTrafficDescriptor, bool)`

GetDddTrafDescriptorOk returns a tuple with the DddTrafDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddTrafDescriptor

`func (o *CreateMonitoringEventSubscription200Response) SetDddTrafDescriptor(v DddTrafficDescriptor)`

SetDddTrafDescriptor sets DddTrafDescriptor field to given value.

### HasDddTrafDescriptor

`func (o *CreateMonitoringEventSubscription200Response) HasDddTrafDescriptor() bool`

HasDddTrafDescriptor returns a boolean if a field has been set.

### GetMaxWaitTime

`func (o *CreateMonitoringEventSubscription200Response) GetMaxWaitTime() time.Time`

GetMaxWaitTime returns the MaxWaitTime field if non-nil, zero value otherwise.

### GetMaxWaitTimeOk

`func (o *CreateMonitoringEventSubscription200Response) GetMaxWaitTimeOk() (*time.Time, bool)`

GetMaxWaitTimeOk returns a tuple with the MaxWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitTime

`func (o *CreateMonitoringEventSubscription200Response) SetMaxWaitTime(v time.Time)`

SetMaxWaitTime sets MaxWaitTime field to given value.

### HasMaxWaitTime

`func (o *CreateMonitoringEventSubscription200Response) HasMaxWaitTime() bool`

HasMaxWaitTime returns a boolean if a field has been set.

### GetApiCaps

`func (o *CreateMonitoringEventSubscription200Response) GetApiCaps() []ApiCapabilityInfo`

GetApiCaps returns the ApiCaps field if non-nil, zero value otherwise.

### GetApiCapsOk

`func (o *CreateMonitoringEventSubscription200Response) GetApiCapsOk() (*[]ApiCapabilityInfo, bool)`

GetApiCapsOk returns a tuple with the ApiCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCaps

`func (o *CreateMonitoringEventSubscription200Response) SetApiCaps(v []ApiCapabilityInfo)`

SetApiCaps sets ApiCaps field to given value.

### HasApiCaps

`func (o *CreateMonitoringEventSubscription200Response) HasApiCaps() bool`

HasApiCaps returns a boolean if a field has been set.

### GetNSStatusInfo

`func (o *CreateMonitoringEventSubscription200Response) GetNSStatusInfo() SACEventStatus`

GetNSStatusInfo returns the NSStatusInfo field if non-nil, zero value otherwise.

### GetNSStatusInfoOk

`func (o *CreateMonitoringEventSubscription200Response) GetNSStatusInfoOk() (*SACEventStatus, bool)`

GetNSStatusInfoOk returns a tuple with the NSStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSStatusInfo

`func (o *CreateMonitoringEventSubscription200Response) SetNSStatusInfo(v SACEventStatus)`

SetNSStatusInfo sets NSStatusInfo field to given value.

### HasNSStatusInfo

`func (o *CreateMonitoringEventSubscription200Response) HasNSStatusInfo() bool`

HasNSStatusInfo returns a boolean if a field has been set.

### GetAfServiceId

`func (o *CreateMonitoringEventSubscription200Response) GetAfServiceId() string`

GetAfServiceId returns the AfServiceId field if non-nil, zero value otherwise.

### GetAfServiceIdOk

`func (o *CreateMonitoringEventSubscription200Response) GetAfServiceIdOk() (*string, bool)`

GetAfServiceIdOk returns a tuple with the AfServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfServiceId

`func (o *CreateMonitoringEventSubscription200Response) SetAfServiceId(v string)`

SetAfServiceId sets AfServiceId field to given value.

### HasAfServiceId

`func (o *CreateMonitoringEventSubscription200Response) HasAfServiceId() bool`

HasAfServiceId returns a boolean if a field has been set.

### GetServLevelDevId

`func (o *CreateMonitoringEventSubscription200Response) GetServLevelDevId() string`

GetServLevelDevId returns the ServLevelDevId field if non-nil, zero value otherwise.

### GetServLevelDevIdOk

`func (o *CreateMonitoringEventSubscription200Response) GetServLevelDevIdOk() (*string, bool)`

GetServLevelDevIdOk returns a tuple with the ServLevelDevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServLevelDevId

`func (o *CreateMonitoringEventSubscription200Response) SetServLevelDevId(v string)`

SetServLevelDevId sets ServLevelDevId field to given value.

### HasServLevelDevId

`func (o *CreateMonitoringEventSubscription200Response) HasServLevelDevId() bool`

HasServLevelDevId returns a boolean if a field has been set.

### GetUavPresInd

`func (o *CreateMonitoringEventSubscription200Response) GetUavPresInd() bool`

GetUavPresInd returns the UavPresInd field if non-nil, zero value otherwise.

### GetUavPresIndOk

`func (o *CreateMonitoringEventSubscription200Response) GetUavPresIndOk() (*bool, bool)`

GetUavPresIndOk returns a tuple with the UavPresInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavPresInd

`func (o *CreateMonitoringEventSubscription200Response) SetUavPresInd(v bool)`

SetUavPresInd sets UavPresInd field to given value.

### HasUavPresInd

`func (o *CreateMonitoringEventSubscription200Response) HasUavPresInd() bool`

HasUavPresInd returns a boolean if a field has been set.

### GetMonitoringEventReports

`func (o *CreateMonitoringEventSubscription200Response) GetMonitoringEventReports() []MonitoringEventReport`

GetMonitoringEventReports returns the MonitoringEventReports field if non-nil, zero value otherwise.

### GetMonitoringEventReportsOk

`func (o *CreateMonitoringEventSubscription200Response) GetMonitoringEventReportsOk() (*[]MonitoringEventReport, bool)`

GetMonitoringEventReportsOk returns a tuple with the MonitoringEventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEventReports

`func (o *CreateMonitoringEventSubscription200Response) SetMonitoringEventReports(v []MonitoringEventReport)`

SetMonitoringEventReports sets MonitoringEventReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


