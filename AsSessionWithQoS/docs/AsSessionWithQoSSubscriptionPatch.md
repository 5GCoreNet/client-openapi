# AsSessionWithQoSSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExterAppId** | Pointer to **string** | Identifies the external Application Identifier. | [optional] 
**FlowInfo** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Describe the IP data flow which requires QoS. | [optional] 
**EthFlowInfo** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet flows. | [optional] 
**EnEthFlowInfo** | Pointer to [**[]EthFlowInfo**](EthFlowInfo.md) | Identifies the Ethernet flows which require QoS. Each Ethernet flow consists of a flow idenifer and the corresponding UL and/or DL flows.  | [optional] 
**QosReference** | Pointer to **string** | Pre-defined QoS reference | [optional] 
**AltQoSReferences** | Pointer to **[]string** | Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority. | [optional] 
**AltQosReqs** | Pointer to [**[]AlternativeServiceRequirementsData**](AlternativeServiceRequirementsData.md) | Identifies an ordered list of alternative service requirements that include individual QoS parameter sets. The lower the index of the array for a given entry, the higher the priority. | [optional] 
**DisUeNotif** | Pointer to **bool** | Indicates whether the QoS flow parameters signalling to the UE when the SMF is notified by the NG-RAN of changes in the fulfilled QoS situation is disabled (true) or not (false). The fulfilled situation is either the QoS profile or an Alternative QoS Profile.  | [optional] 
**UsageThreshold** | Pointer to [**NullableUsageThresholdRm**](UsageThresholdRm.md) |  | [optional] 
**QosMonInfo** | Pointer to [**QosMonitoringInformationRm**](QosMonitoringInformationRm.md) |  | [optional] 
**DirectNotifInd** | Pointer to **bool** | Indicates whether the direct event notification is requested (true) or not (false).  | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**TscQosReq** | Pointer to [**TscQosRequirementRm**](TscQosRequirementRm.md) |  | [optional] 
**Events** | Pointer to [**[]UserPlaneEvent**](UserPlaneEvent.md) | Represents the updated list of user plane event(s) to which the SCS/AS requests to subscribe to. | [optional] 

## Methods

### NewAsSessionWithQoSSubscriptionPatch

`func NewAsSessionWithQoSSubscriptionPatch() *AsSessionWithQoSSubscriptionPatch`

NewAsSessionWithQoSSubscriptionPatch instantiates a new AsSessionWithQoSSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsSessionWithQoSSubscriptionPatchWithDefaults

`func NewAsSessionWithQoSSubscriptionPatchWithDefaults() *AsSessionWithQoSSubscriptionPatch`

NewAsSessionWithQoSSubscriptionPatchWithDefaults instantiates a new AsSessionWithQoSSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExterAppId

`func (o *AsSessionWithQoSSubscriptionPatch) GetExterAppId() string`

GetExterAppId returns the ExterAppId field if non-nil, zero value otherwise.

### GetExterAppIdOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetExterAppIdOk() (*string, bool)`

GetExterAppIdOk returns a tuple with the ExterAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterAppId

`func (o *AsSessionWithQoSSubscriptionPatch) SetExterAppId(v string)`

SetExterAppId sets ExterAppId field to given value.

### HasExterAppId

`func (o *AsSessionWithQoSSubscriptionPatch) HasExterAppId() bool`

HasExterAppId returns a boolean if a field has been set.

### GetFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) GetFlowInfo() []FlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetFlowInfoOk() (*[]FlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) SetFlowInfo(v []FlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) GetEthFlowInfo() []EthFlowDescription`

GetEthFlowInfo returns the EthFlowInfo field if non-nil, zero value otherwise.

### GetEthFlowInfoOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetEthFlowInfoOk() (*[]EthFlowDescription, bool)`

GetEthFlowInfoOk returns a tuple with the EthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) SetEthFlowInfo(v []EthFlowDescription)`

SetEthFlowInfo sets EthFlowInfo field to given value.

### HasEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) HasEthFlowInfo() bool`

HasEthFlowInfo returns a boolean if a field has been set.

### GetEnEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) GetEnEthFlowInfo() []EthFlowInfo`

GetEnEthFlowInfo returns the EnEthFlowInfo field if non-nil, zero value otherwise.

### GetEnEthFlowInfoOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetEnEthFlowInfoOk() (*[]EthFlowInfo, bool)`

GetEnEthFlowInfoOk returns a tuple with the EnEthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) SetEnEthFlowInfo(v []EthFlowInfo)`

SetEnEthFlowInfo sets EnEthFlowInfo field to given value.

### HasEnEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) HasEnEthFlowInfo() bool`

HasEnEthFlowInfo returns a boolean if a field has been set.

### GetQosReference

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *AsSessionWithQoSSubscriptionPatch) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *AsSessionWithQoSSubscriptionPatch) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetAltQoSReferences

`func (o *AsSessionWithQoSSubscriptionPatch) GetAltQoSReferences() []string`

GetAltQoSReferences returns the AltQoSReferences field if non-nil, zero value otherwise.

### GetAltQoSReferencesOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetAltQoSReferencesOk() (*[]string, bool)`

GetAltQoSReferencesOk returns a tuple with the AltQoSReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQoSReferences

`func (o *AsSessionWithQoSSubscriptionPatch) SetAltQoSReferences(v []string)`

SetAltQoSReferences sets AltQoSReferences field to given value.

### HasAltQoSReferences

`func (o *AsSessionWithQoSSubscriptionPatch) HasAltQoSReferences() bool`

HasAltQoSReferences returns a boolean if a field has been set.

### GetAltQosReqs

`func (o *AsSessionWithQoSSubscriptionPatch) GetAltQosReqs() []AlternativeServiceRequirementsData`

GetAltQosReqs returns the AltQosReqs field if non-nil, zero value otherwise.

### GetAltQosReqsOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetAltQosReqsOk() (*[]AlternativeServiceRequirementsData, bool)`

GetAltQosReqsOk returns a tuple with the AltQosReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosReqs

`func (o *AsSessionWithQoSSubscriptionPatch) SetAltQosReqs(v []AlternativeServiceRequirementsData)`

SetAltQosReqs sets AltQosReqs field to given value.

### HasAltQosReqs

`func (o *AsSessionWithQoSSubscriptionPatch) HasAltQosReqs() bool`

HasAltQosReqs returns a boolean if a field has been set.

### GetDisUeNotif

`func (o *AsSessionWithQoSSubscriptionPatch) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *AsSessionWithQoSSubscriptionPatch) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *AsSessionWithQoSSubscriptionPatch) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *AsSessionWithQoSSubscriptionPatch) GetUsageThreshold() UsageThresholdRm`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetUsageThresholdOk() (*UsageThresholdRm, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *AsSessionWithQoSSubscriptionPatch) SetUsageThreshold(v UsageThresholdRm)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *AsSessionWithQoSSubscriptionPatch) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### SetUsageThresholdNil

`func (o *AsSessionWithQoSSubscriptionPatch) SetUsageThresholdNil(b bool)`

 SetUsageThresholdNil sets the value for UsageThreshold to be an explicit nil

### UnsetUsageThreshold
`func (o *AsSessionWithQoSSubscriptionPatch) UnsetUsageThreshold()`

UnsetUsageThreshold ensures that no value is present for UsageThreshold, not even an explicit nil
### GetQosMonInfo

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosMonInfo() QosMonitoringInformationRm`

GetQosMonInfo returns the QosMonInfo field if non-nil, zero value otherwise.

### GetQosMonInfoOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosMonInfoOk() (*QosMonitoringInformationRm, bool)`

GetQosMonInfoOk returns a tuple with the QosMonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonInfo

`func (o *AsSessionWithQoSSubscriptionPatch) SetQosMonInfo(v QosMonitoringInformationRm)`

SetQosMonInfo sets QosMonInfo field to given value.

### HasQosMonInfo

`func (o *AsSessionWithQoSSubscriptionPatch) HasQosMonInfo() bool`

HasQosMonInfo returns a boolean if a field has been set.

### GetDirectNotifInd

`func (o *AsSessionWithQoSSubscriptionPatch) GetDirectNotifInd() bool`

GetDirectNotifInd returns the DirectNotifInd field if non-nil, zero value otherwise.

### GetDirectNotifIndOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetDirectNotifIndOk() (*bool, bool)`

GetDirectNotifIndOk returns a tuple with the DirectNotifInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectNotifInd

`func (o *AsSessionWithQoSSubscriptionPatch) SetDirectNotifInd(v bool)`

SetDirectNotifInd sets DirectNotifInd field to given value.

### HasDirectNotifInd

`func (o *AsSessionWithQoSSubscriptionPatch) HasDirectNotifInd() bool`

HasDirectNotifInd returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *AsSessionWithQoSSubscriptionPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *AsSessionWithQoSSubscriptionPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *AsSessionWithQoSSubscriptionPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetTscQosReq

`func (o *AsSessionWithQoSSubscriptionPatch) GetTscQosReq() TscQosRequirementRm`

GetTscQosReq returns the TscQosReq field if non-nil, zero value otherwise.

### GetTscQosReqOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetTscQosReqOk() (*TscQosRequirementRm, bool)`

GetTscQosReqOk returns a tuple with the TscQosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscQosReq

`func (o *AsSessionWithQoSSubscriptionPatch) SetTscQosReq(v TscQosRequirementRm)`

SetTscQosReq sets TscQosReq field to given value.

### HasTscQosReq

`func (o *AsSessionWithQoSSubscriptionPatch) HasTscQosReq() bool`

HasTscQosReq returns a boolean if a field has been set.

### GetEvents

`func (o *AsSessionWithQoSSubscriptionPatch) GetEvents() []UserPlaneEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetEventsOk() (*[]UserPlaneEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AsSessionWithQoSSubscriptionPatch) SetEvents(v []UserPlaneEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AsSessionWithQoSSubscriptionPatch) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


