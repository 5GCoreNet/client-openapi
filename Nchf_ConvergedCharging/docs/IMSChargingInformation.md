# IMSChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to [**SIPEventType**](SIPEventType.md) |  | [optional] 
**IMSNodeFunctionality** | Pointer to [**IMSNodeFunctionality**](IMSNodeFunctionality.md) |  | [optional] 
**RoleOfNode** | Pointer to [**RoleOfIMSNode**](RoleOfIMSNode.md) |  | [optional] 
**UserInformation** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**UserLocationInfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**Var3gppPSDataOffStatus** | Pointer to [**Model3GPPPSDataOffStatus**](Model3GPPPSDataOffStatus.md) |  | [optional] 
**IsupCause** | Pointer to [**ISUPCause**](ISUPCause.md) |  | [optional] 
**ControlPlaneAddress** | Pointer to [**IMSAddress**](IMSAddress.md) |  | [optional] 
**VlrNumber** | Pointer to **string** |  | [optional] 
**MscAddress** | Pointer to **string** |  | [optional] 
**UserSessionID** | Pointer to **string** |  | [optional] 
**OutgoingSessionID** | Pointer to **string** |  | [optional] 
**SessionPriority** | Pointer to [**IMSSessionPriority**](IMSSessionPriority.md) |  | [optional] 
**CallingPartyAddresses** | Pointer to **[]string** |  | [optional] 
**CalledPartyAddress** | Pointer to **string** |  | [optional] 
**NumberPortabilityRoutinginformation** | Pointer to **string** |  | [optional] 
**CarrierSelectRoutingInformation** | Pointer to **string** |  | [optional] 
**AlternateChargedPartyAddress** | Pointer to **string** |  | [optional] 
**RequestedPartyAddress** | Pointer to **[]string** |  | [optional] 
**CalledAssertedIdentities** | Pointer to **[]string** |  | [optional] 
**CalledIdentityChanges** | Pointer to [**[]CalledIdentityChange**](CalledIdentityChange.md) |  | [optional] 
**AssociatedURI** | Pointer to **[]string** |  | [optional] 
**TimeStamps** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ApplicationServerInformation** | Pointer to **[]string** |  | [optional] 
**InterOperatorIdentifier** | Pointer to [**[]InterOperatorIdentifier**](InterOperatorIdentifier.md) |  | [optional] 
**ImsChargingIdentifier** | Pointer to **string** |  | [optional] 
**RelatedICID** | Pointer to **string** |  | [optional] 
**RelatedICIDGenerationNode** | Pointer to **string** |  | [optional] 
**TransitIOIList** | Pointer to **[]string** |  | [optional] 
**EarlyMediaDescription** | Pointer to [**[]EarlyMediaDescription**](EarlyMediaDescription.md) |  | [optional] 
**SdpSessionDescription** | Pointer to **[]string** |  | [optional] 
**SdpMediaComponent** | Pointer to [**[]SDPMediaComponent**](SDPMediaComponent.md) |  | [optional] 
**ServedPartyIPAddress** | Pointer to [**IMSAddress**](IMSAddress.md) |  | [optional] 
**ServerCapabilities** | Pointer to [**ServerCapabilities**](ServerCapabilities.md) |  | [optional] 
**TrunkGroupID** | Pointer to [**TrunkGroupID**](TrunkGroupID.md) |  | [optional] 
**BearerService** | Pointer to **string** |  | [optional] 
**ImsServiceId** | Pointer to **string** |  | [optional] 
**MessageBodies** | Pointer to [**[]MessageBody**](MessageBody.md) |  | [optional] 
**AccessNetworkInformation** | Pointer to **[]string** |  | [optional] 
**AdditionalAccessNetworkInformation** | Pointer to **string** |  | [optional] 
**CellularNetworkInformation** | Pointer to **string** |  | [optional] 
**AccessTransferInformation** | Pointer to [**[]AccessTransferInformation**](AccessTransferInformation.md) |  | [optional] 
**AccessNetworkInfoChange** | Pointer to [**[]AccessNetworkInfoChange**](AccessNetworkInfoChange.md) |  | [optional] 
**ImsCommunicationServiceID** | Pointer to **string** |  | [optional] 
**ImsApplicationReferenceID** | Pointer to **string** |  | [optional] 
**CauseCode** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**ReasonHeader** | Pointer to **[]string** |  | [optional] 
**InitialIMSChargingIdentifier** | Pointer to **string** |  | [optional] 
**NniInformation** | Pointer to [**[]NNIInformation**](NNIInformation.md) |  | [optional] 
**FromAddress** | Pointer to **string** |  | [optional] 
**ImsEmergencyIndication** | Pointer to **bool** |  | [optional] 
**ImsVisitedNetworkIdentifier** | Pointer to **string** |  | [optional] 
**SipRouteHeaderReceived** | Pointer to **string** |  | [optional] 
**SipRouteHeaderTransmitted** | Pointer to **string** |  | [optional] 
**TadIdentifier** | Pointer to [**TADIdentifier**](TADIdentifier.md) |  | [optional] 
**FeIdentifierList** | Pointer to **string** |  | [optional] 

## Methods

### NewIMSChargingInformation

`func NewIMSChargingInformation() *IMSChargingInformation`

NewIMSChargingInformation instantiates a new IMSChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIMSChargingInformationWithDefaults

`func NewIMSChargingInformationWithDefaults() *IMSChargingInformation`

NewIMSChargingInformationWithDefaults instantiates a new IMSChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *IMSChargingInformation) GetEventType() SIPEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IMSChargingInformation) GetEventTypeOk() (*SIPEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IMSChargingInformation) SetEventType(v SIPEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *IMSChargingInformation) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetIMSNodeFunctionality

`func (o *IMSChargingInformation) GetIMSNodeFunctionality() IMSNodeFunctionality`

GetIMSNodeFunctionality returns the IMSNodeFunctionality field if non-nil, zero value otherwise.

### GetIMSNodeFunctionalityOk

`func (o *IMSChargingInformation) GetIMSNodeFunctionalityOk() (*IMSNodeFunctionality, bool)`

GetIMSNodeFunctionalityOk returns a tuple with the IMSNodeFunctionality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMSNodeFunctionality

`func (o *IMSChargingInformation) SetIMSNodeFunctionality(v IMSNodeFunctionality)`

SetIMSNodeFunctionality sets IMSNodeFunctionality field to given value.

### HasIMSNodeFunctionality

`func (o *IMSChargingInformation) HasIMSNodeFunctionality() bool`

HasIMSNodeFunctionality returns a boolean if a field has been set.

### GetRoleOfNode

`func (o *IMSChargingInformation) GetRoleOfNode() RoleOfIMSNode`

GetRoleOfNode returns the RoleOfNode field if non-nil, zero value otherwise.

### GetRoleOfNodeOk

`func (o *IMSChargingInformation) GetRoleOfNodeOk() (*RoleOfIMSNode, bool)`

GetRoleOfNodeOk returns a tuple with the RoleOfNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleOfNode

`func (o *IMSChargingInformation) SetRoleOfNode(v RoleOfIMSNode)`

SetRoleOfNode sets RoleOfNode field to given value.

### HasRoleOfNode

`func (o *IMSChargingInformation) HasRoleOfNode() bool`

HasRoleOfNode returns a boolean if a field has been set.

### GetUserInformation

`func (o *IMSChargingInformation) GetUserInformation() UserInformation`

GetUserInformation returns the UserInformation field if non-nil, zero value otherwise.

### GetUserInformationOk

`func (o *IMSChargingInformation) GetUserInformationOk() (*UserInformation, bool)`

GetUserInformationOk returns a tuple with the UserInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInformation

`func (o *IMSChargingInformation) SetUserInformation(v UserInformation)`

SetUserInformation sets UserInformation field to given value.

### HasUserInformation

`func (o *IMSChargingInformation) HasUserInformation() bool`

HasUserInformation returns a boolean if a field has been set.

### GetUserLocationInfo

`func (o *IMSChargingInformation) GetUserLocationInfo() UserLocation`

GetUserLocationInfo returns the UserLocationInfo field if non-nil, zero value otherwise.

### GetUserLocationInfoOk

`func (o *IMSChargingInformation) GetUserLocationInfoOk() (*UserLocation, bool)`

GetUserLocationInfoOk returns a tuple with the UserLocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInfo

`func (o *IMSChargingInformation) SetUserLocationInfo(v UserLocation)`

SetUserLocationInfo sets UserLocationInfo field to given value.

### HasUserLocationInfo

`func (o *IMSChargingInformation) HasUserLocationInfo() bool`

HasUserLocationInfo returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *IMSChargingInformation) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *IMSChargingInformation) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *IMSChargingInformation) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *IMSChargingInformation) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetVar3gppPSDataOffStatus

`func (o *IMSChargingInformation) GetVar3gppPSDataOffStatus() Model3GPPPSDataOffStatus`

GetVar3gppPSDataOffStatus returns the Var3gppPSDataOffStatus field if non-nil, zero value otherwise.

### GetVar3gppPSDataOffStatusOk

`func (o *IMSChargingInformation) GetVar3gppPSDataOffStatusOk() (*Model3GPPPSDataOffStatus, bool)`

GetVar3gppPSDataOffStatusOk returns a tuple with the Var3gppPSDataOffStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppPSDataOffStatus

`func (o *IMSChargingInformation) SetVar3gppPSDataOffStatus(v Model3GPPPSDataOffStatus)`

SetVar3gppPSDataOffStatus sets Var3gppPSDataOffStatus field to given value.

### HasVar3gppPSDataOffStatus

`func (o *IMSChargingInformation) HasVar3gppPSDataOffStatus() bool`

HasVar3gppPSDataOffStatus returns a boolean if a field has been set.

### GetIsupCause

`func (o *IMSChargingInformation) GetIsupCause() ISUPCause`

GetIsupCause returns the IsupCause field if non-nil, zero value otherwise.

### GetIsupCauseOk

`func (o *IMSChargingInformation) GetIsupCauseOk() (*ISUPCause, bool)`

GetIsupCauseOk returns a tuple with the IsupCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsupCause

`func (o *IMSChargingInformation) SetIsupCause(v ISUPCause)`

SetIsupCause sets IsupCause field to given value.

### HasIsupCause

`func (o *IMSChargingInformation) HasIsupCause() bool`

HasIsupCause returns a boolean if a field has been set.

### GetControlPlaneAddress

`func (o *IMSChargingInformation) GetControlPlaneAddress() IMSAddress`

GetControlPlaneAddress returns the ControlPlaneAddress field if non-nil, zero value otherwise.

### GetControlPlaneAddressOk

`func (o *IMSChargingInformation) GetControlPlaneAddressOk() (*IMSAddress, bool)`

GetControlPlaneAddressOk returns a tuple with the ControlPlaneAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneAddress

`func (o *IMSChargingInformation) SetControlPlaneAddress(v IMSAddress)`

SetControlPlaneAddress sets ControlPlaneAddress field to given value.

### HasControlPlaneAddress

`func (o *IMSChargingInformation) HasControlPlaneAddress() bool`

HasControlPlaneAddress returns a boolean if a field has been set.

### GetVlrNumber

`func (o *IMSChargingInformation) GetVlrNumber() string`

GetVlrNumber returns the VlrNumber field if non-nil, zero value otherwise.

### GetVlrNumberOk

`func (o *IMSChargingInformation) GetVlrNumberOk() (*string, bool)`

GetVlrNumberOk returns a tuple with the VlrNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlrNumber

`func (o *IMSChargingInformation) SetVlrNumber(v string)`

SetVlrNumber sets VlrNumber field to given value.

### HasVlrNumber

`func (o *IMSChargingInformation) HasVlrNumber() bool`

HasVlrNumber returns a boolean if a field has been set.

### GetMscAddress

`func (o *IMSChargingInformation) GetMscAddress() string`

GetMscAddress returns the MscAddress field if non-nil, zero value otherwise.

### GetMscAddressOk

`func (o *IMSChargingInformation) GetMscAddressOk() (*string, bool)`

GetMscAddressOk returns a tuple with the MscAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMscAddress

`func (o *IMSChargingInformation) SetMscAddress(v string)`

SetMscAddress sets MscAddress field to given value.

### HasMscAddress

`func (o *IMSChargingInformation) HasMscAddress() bool`

HasMscAddress returns a boolean if a field has been set.

### GetUserSessionID

`func (o *IMSChargingInformation) GetUserSessionID() string`

GetUserSessionID returns the UserSessionID field if non-nil, zero value otherwise.

### GetUserSessionIDOk

`func (o *IMSChargingInformation) GetUserSessionIDOk() (*string, bool)`

GetUserSessionIDOk returns a tuple with the UserSessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionID

`func (o *IMSChargingInformation) SetUserSessionID(v string)`

SetUserSessionID sets UserSessionID field to given value.

### HasUserSessionID

`func (o *IMSChargingInformation) HasUserSessionID() bool`

HasUserSessionID returns a boolean if a field has been set.

### GetOutgoingSessionID

`func (o *IMSChargingInformation) GetOutgoingSessionID() string`

GetOutgoingSessionID returns the OutgoingSessionID field if non-nil, zero value otherwise.

### GetOutgoingSessionIDOk

`func (o *IMSChargingInformation) GetOutgoingSessionIDOk() (*string, bool)`

GetOutgoingSessionIDOk returns a tuple with the OutgoingSessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingSessionID

`func (o *IMSChargingInformation) SetOutgoingSessionID(v string)`

SetOutgoingSessionID sets OutgoingSessionID field to given value.

### HasOutgoingSessionID

`func (o *IMSChargingInformation) HasOutgoingSessionID() bool`

HasOutgoingSessionID returns a boolean if a field has been set.

### GetSessionPriority

`func (o *IMSChargingInformation) GetSessionPriority() IMSSessionPriority`

GetSessionPriority returns the SessionPriority field if non-nil, zero value otherwise.

### GetSessionPriorityOk

`func (o *IMSChargingInformation) GetSessionPriorityOk() (*IMSSessionPriority, bool)`

GetSessionPriorityOk returns a tuple with the SessionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPriority

`func (o *IMSChargingInformation) SetSessionPriority(v IMSSessionPriority)`

SetSessionPriority sets SessionPriority field to given value.

### HasSessionPriority

`func (o *IMSChargingInformation) HasSessionPriority() bool`

HasSessionPriority returns a boolean if a field has been set.

### GetCallingPartyAddresses

`func (o *IMSChargingInformation) GetCallingPartyAddresses() []string`

GetCallingPartyAddresses returns the CallingPartyAddresses field if non-nil, zero value otherwise.

### GetCallingPartyAddressesOk

`func (o *IMSChargingInformation) GetCallingPartyAddressesOk() (*[]string, bool)`

GetCallingPartyAddressesOk returns a tuple with the CallingPartyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingPartyAddresses

`func (o *IMSChargingInformation) SetCallingPartyAddresses(v []string)`

SetCallingPartyAddresses sets CallingPartyAddresses field to given value.

### HasCallingPartyAddresses

`func (o *IMSChargingInformation) HasCallingPartyAddresses() bool`

HasCallingPartyAddresses returns a boolean if a field has been set.

### GetCalledPartyAddress

`func (o *IMSChargingInformation) GetCalledPartyAddress() string`

GetCalledPartyAddress returns the CalledPartyAddress field if non-nil, zero value otherwise.

### GetCalledPartyAddressOk

`func (o *IMSChargingInformation) GetCalledPartyAddressOk() (*string, bool)`

GetCalledPartyAddressOk returns a tuple with the CalledPartyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledPartyAddress

`func (o *IMSChargingInformation) SetCalledPartyAddress(v string)`

SetCalledPartyAddress sets CalledPartyAddress field to given value.

### HasCalledPartyAddress

`func (o *IMSChargingInformation) HasCalledPartyAddress() bool`

HasCalledPartyAddress returns a boolean if a field has been set.

### GetNumberPortabilityRoutinginformation

`func (o *IMSChargingInformation) GetNumberPortabilityRoutinginformation() string`

GetNumberPortabilityRoutinginformation returns the NumberPortabilityRoutinginformation field if non-nil, zero value otherwise.

### GetNumberPortabilityRoutinginformationOk

`func (o *IMSChargingInformation) GetNumberPortabilityRoutinginformationOk() (*string, bool)`

GetNumberPortabilityRoutinginformationOk returns a tuple with the NumberPortabilityRoutinginformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberPortabilityRoutinginformation

`func (o *IMSChargingInformation) SetNumberPortabilityRoutinginformation(v string)`

SetNumberPortabilityRoutinginformation sets NumberPortabilityRoutinginformation field to given value.

### HasNumberPortabilityRoutinginformation

`func (o *IMSChargingInformation) HasNumberPortabilityRoutinginformation() bool`

HasNumberPortabilityRoutinginformation returns a boolean if a field has been set.

### GetCarrierSelectRoutingInformation

`func (o *IMSChargingInformation) GetCarrierSelectRoutingInformation() string`

GetCarrierSelectRoutingInformation returns the CarrierSelectRoutingInformation field if non-nil, zero value otherwise.

### GetCarrierSelectRoutingInformationOk

`func (o *IMSChargingInformation) GetCarrierSelectRoutingInformationOk() (*string, bool)`

GetCarrierSelectRoutingInformationOk returns a tuple with the CarrierSelectRoutingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSelectRoutingInformation

`func (o *IMSChargingInformation) SetCarrierSelectRoutingInformation(v string)`

SetCarrierSelectRoutingInformation sets CarrierSelectRoutingInformation field to given value.

### HasCarrierSelectRoutingInformation

`func (o *IMSChargingInformation) HasCarrierSelectRoutingInformation() bool`

HasCarrierSelectRoutingInformation returns a boolean if a field has been set.

### GetAlternateChargedPartyAddress

`func (o *IMSChargingInformation) GetAlternateChargedPartyAddress() string`

GetAlternateChargedPartyAddress returns the AlternateChargedPartyAddress field if non-nil, zero value otherwise.

### GetAlternateChargedPartyAddressOk

`func (o *IMSChargingInformation) GetAlternateChargedPartyAddressOk() (*string, bool)`

GetAlternateChargedPartyAddressOk returns a tuple with the AlternateChargedPartyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateChargedPartyAddress

`func (o *IMSChargingInformation) SetAlternateChargedPartyAddress(v string)`

SetAlternateChargedPartyAddress sets AlternateChargedPartyAddress field to given value.

### HasAlternateChargedPartyAddress

`func (o *IMSChargingInformation) HasAlternateChargedPartyAddress() bool`

HasAlternateChargedPartyAddress returns a boolean if a field has been set.

### GetRequestedPartyAddress

`func (o *IMSChargingInformation) GetRequestedPartyAddress() []string`

GetRequestedPartyAddress returns the RequestedPartyAddress field if non-nil, zero value otherwise.

### GetRequestedPartyAddressOk

`func (o *IMSChargingInformation) GetRequestedPartyAddressOk() (*[]string, bool)`

GetRequestedPartyAddressOk returns a tuple with the RequestedPartyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPartyAddress

`func (o *IMSChargingInformation) SetRequestedPartyAddress(v []string)`

SetRequestedPartyAddress sets RequestedPartyAddress field to given value.

### HasRequestedPartyAddress

`func (o *IMSChargingInformation) HasRequestedPartyAddress() bool`

HasRequestedPartyAddress returns a boolean if a field has been set.

### GetCalledAssertedIdentities

`func (o *IMSChargingInformation) GetCalledAssertedIdentities() []string`

GetCalledAssertedIdentities returns the CalledAssertedIdentities field if non-nil, zero value otherwise.

### GetCalledAssertedIdentitiesOk

`func (o *IMSChargingInformation) GetCalledAssertedIdentitiesOk() (*[]string, bool)`

GetCalledAssertedIdentitiesOk returns a tuple with the CalledAssertedIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledAssertedIdentities

`func (o *IMSChargingInformation) SetCalledAssertedIdentities(v []string)`

SetCalledAssertedIdentities sets CalledAssertedIdentities field to given value.

### HasCalledAssertedIdentities

`func (o *IMSChargingInformation) HasCalledAssertedIdentities() bool`

HasCalledAssertedIdentities returns a boolean if a field has been set.

### GetCalledIdentityChanges

`func (o *IMSChargingInformation) GetCalledIdentityChanges() []CalledIdentityChange`

GetCalledIdentityChanges returns the CalledIdentityChanges field if non-nil, zero value otherwise.

### GetCalledIdentityChangesOk

`func (o *IMSChargingInformation) GetCalledIdentityChangesOk() (*[]CalledIdentityChange, bool)`

GetCalledIdentityChangesOk returns a tuple with the CalledIdentityChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledIdentityChanges

`func (o *IMSChargingInformation) SetCalledIdentityChanges(v []CalledIdentityChange)`

SetCalledIdentityChanges sets CalledIdentityChanges field to given value.

### HasCalledIdentityChanges

`func (o *IMSChargingInformation) HasCalledIdentityChanges() bool`

HasCalledIdentityChanges returns a boolean if a field has been set.

### GetAssociatedURI

`func (o *IMSChargingInformation) GetAssociatedURI() []string`

GetAssociatedURI returns the AssociatedURI field if non-nil, zero value otherwise.

### GetAssociatedURIOk

`func (o *IMSChargingInformation) GetAssociatedURIOk() (*[]string, bool)`

GetAssociatedURIOk returns a tuple with the AssociatedURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedURI

`func (o *IMSChargingInformation) SetAssociatedURI(v []string)`

SetAssociatedURI sets AssociatedURI field to given value.

### HasAssociatedURI

`func (o *IMSChargingInformation) HasAssociatedURI() bool`

HasAssociatedURI returns a boolean if a field has been set.

### GetTimeStamps

`func (o *IMSChargingInformation) GetTimeStamps() time.Time`

GetTimeStamps returns the TimeStamps field if non-nil, zero value otherwise.

### GetTimeStampsOk

`func (o *IMSChargingInformation) GetTimeStampsOk() (*time.Time, bool)`

GetTimeStampsOk returns a tuple with the TimeStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamps

`func (o *IMSChargingInformation) SetTimeStamps(v time.Time)`

SetTimeStamps sets TimeStamps field to given value.

### HasTimeStamps

`func (o *IMSChargingInformation) HasTimeStamps() bool`

HasTimeStamps returns a boolean if a field has been set.

### GetApplicationServerInformation

`func (o *IMSChargingInformation) GetApplicationServerInformation() []string`

GetApplicationServerInformation returns the ApplicationServerInformation field if non-nil, zero value otherwise.

### GetApplicationServerInformationOk

`func (o *IMSChargingInformation) GetApplicationServerInformationOk() (*[]string, bool)`

GetApplicationServerInformationOk returns a tuple with the ApplicationServerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationServerInformation

`func (o *IMSChargingInformation) SetApplicationServerInformation(v []string)`

SetApplicationServerInformation sets ApplicationServerInformation field to given value.

### HasApplicationServerInformation

`func (o *IMSChargingInformation) HasApplicationServerInformation() bool`

HasApplicationServerInformation returns a boolean if a field has been set.

### GetInterOperatorIdentifier

`func (o *IMSChargingInformation) GetInterOperatorIdentifier() []InterOperatorIdentifier`

GetInterOperatorIdentifier returns the InterOperatorIdentifier field if non-nil, zero value otherwise.

### GetInterOperatorIdentifierOk

`func (o *IMSChargingInformation) GetInterOperatorIdentifierOk() (*[]InterOperatorIdentifier, bool)`

GetInterOperatorIdentifierOk returns a tuple with the InterOperatorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterOperatorIdentifier

`func (o *IMSChargingInformation) SetInterOperatorIdentifier(v []InterOperatorIdentifier)`

SetInterOperatorIdentifier sets InterOperatorIdentifier field to given value.

### HasInterOperatorIdentifier

`func (o *IMSChargingInformation) HasInterOperatorIdentifier() bool`

HasInterOperatorIdentifier returns a boolean if a field has been set.

### GetImsChargingIdentifier

`func (o *IMSChargingInformation) GetImsChargingIdentifier() string`

GetImsChargingIdentifier returns the ImsChargingIdentifier field if non-nil, zero value otherwise.

### GetImsChargingIdentifierOk

`func (o *IMSChargingInformation) GetImsChargingIdentifierOk() (*string, bool)`

GetImsChargingIdentifierOk returns a tuple with the ImsChargingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsChargingIdentifier

`func (o *IMSChargingInformation) SetImsChargingIdentifier(v string)`

SetImsChargingIdentifier sets ImsChargingIdentifier field to given value.

### HasImsChargingIdentifier

`func (o *IMSChargingInformation) HasImsChargingIdentifier() bool`

HasImsChargingIdentifier returns a boolean if a field has been set.

### GetRelatedICID

`func (o *IMSChargingInformation) GetRelatedICID() string`

GetRelatedICID returns the RelatedICID field if non-nil, zero value otherwise.

### GetRelatedICIDOk

`func (o *IMSChargingInformation) GetRelatedICIDOk() (*string, bool)`

GetRelatedICIDOk returns a tuple with the RelatedICID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedICID

`func (o *IMSChargingInformation) SetRelatedICID(v string)`

SetRelatedICID sets RelatedICID field to given value.

### HasRelatedICID

`func (o *IMSChargingInformation) HasRelatedICID() bool`

HasRelatedICID returns a boolean if a field has been set.

### GetRelatedICIDGenerationNode

`func (o *IMSChargingInformation) GetRelatedICIDGenerationNode() string`

GetRelatedICIDGenerationNode returns the RelatedICIDGenerationNode field if non-nil, zero value otherwise.

### GetRelatedICIDGenerationNodeOk

`func (o *IMSChargingInformation) GetRelatedICIDGenerationNodeOk() (*string, bool)`

GetRelatedICIDGenerationNodeOk returns a tuple with the RelatedICIDGenerationNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedICIDGenerationNode

`func (o *IMSChargingInformation) SetRelatedICIDGenerationNode(v string)`

SetRelatedICIDGenerationNode sets RelatedICIDGenerationNode field to given value.

### HasRelatedICIDGenerationNode

`func (o *IMSChargingInformation) HasRelatedICIDGenerationNode() bool`

HasRelatedICIDGenerationNode returns a boolean if a field has been set.

### GetTransitIOIList

`func (o *IMSChargingInformation) GetTransitIOIList() []string`

GetTransitIOIList returns the TransitIOIList field if non-nil, zero value otherwise.

### GetTransitIOIListOk

`func (o *IMSChargingInformation) GetTransitIOIListOk() (*[]string, bool)`

GetTransitIOIListOk returns a tuple with the TransitIOIList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitIOIList

`func (o *IMSChargingInformation) SetTransitIOIList(v []string)`

SetTransitIOIList sets TransitIOIList field to given value.

### HasTransitIOIList

`func (o *IMSChargingInformation) HasTransitIOIList() bool`

HasTransitIOIList returns a boolean if a field has been set.

### GetEarlyMediaDescription

`func (o *IMSChargingInformation) GetEarlyMediaDescription() []EarlyMediaDescription`

GetEarlyMediaDescription returns the EarlyMediaDescription field if non-nil, zero value otherwise.

### GetEarlyMediaDescriptionOk

`func (o *IMSChargingInformation) GetEarlyMediaDescriptionOk() (*[]EarlyMediaDescription, bool)`

GetEarlyMediaDescriptionOk returns a tuple with the EarlyMediaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyMediaDescription

`func (o *IMSChargingInformation) SetEarlyMediaDescription(v []EarlyMediaDescription)`

SetEarlyMediaDescription sets EarlyMediaDescription field to given value.

### HasEarlyMediaDescription

`func (o *IMSChargingInformation) HasEarlyMediaDescription() bool`

HasEarlyMediaDescription returns a boolean if a field has been set.

### GetSdpSessionDescription

`func (o *IMSChargingInformation) GetSdpSessionDescription() []string`

GetSdpSessionDescription returns the SdpSessionDescription field if non-nil, zero value otherwise.

### GetSdpSessionDescriptionOk

`func (o *IMSChargingInformation) GetSdpSessionDescriptionOk() (*[]string, bool)`

GetSdpSessionDescriptionOk returns a tuple with the SdpSessionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdpSessionDescription

`func (o *IMSChargingInformation) SetSdpSessionDescription(v []string)`

SetSdpSessionDescription sets SdpSessionDescription field to given value.

### HasSdpSessionDescription

`func (o *IMSChargingInformation) HasSdpSessionDescription() bool`

HasSdpSessionDescription returns a boolean if a field has been set.

### GetSdpMediaComponent

`func (o *IMSChargingInformation) GetSdpMediaComponent() []SDPMediaComponent`

GetSdpMediaComponent returns the SdpMediaComponent field if non-nil, zero value otherwise.

### GetSdpMediaComponentOk

`func (o *IMSChargingInformation) GetSdpMediaComponentOk() (*[]SDPMediaComponent, bool)`

GetSdpMediaComponentOk returns a tuple with the SdpMediaComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdpMediaComponent

`func (o *IMSChargingInformation) SetSdpMediaComponent(v []SDPMediaComponent)`

SetSdpMediaComponent sets SdpMediaComponent field to given value.

### HasSdpMediaComponent

`func (o *IMSChargingInformation) HasSdpMediaComponent() bool`

HasSdpMediaComponent returns a boolean if a field has been set.

### GetServedPartyIPAddress

`func (o *IMSChargingInformation) GetServedPartyIPAddress() IMSAddress`

GetServedPartyIPAddress returns the ServedPartyIPAddress field if non-nil, zero value otherwise.

### GetServedPartyIPAddressOk

`func (o *IMSChargingInformation) GetServedPartyIPAddressOk() (*IMSAddress, bool)`

GetServedPartyIPAddressOk returns a tuple with the ServedPartyIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedPartyIPAddress

`func (o *IMSChargingInformation) SetServedPartyIPAddress(v IMSAddress)`

SetServedPartyIPAddress sets ServedPartyIPAddress field to given value.

### HasServedPartyIPAddress

`func (o *IMSChargingInformation) HasServedPartyIPAddress() bool`

HasServedPartyIPAddress returns a boolean if a field has been set.

### GetServerCapabilities

`func (o *IMSChargingInformation) GetServerCapabilities() ServerCapabilities`

GetServerCapabilities returns the ServerCapabilities field if non-nil, zero value otherwise.

### GetServerCapabilitiesOk

`func (o *IMSChargingInformation) GetServerCapabilitiesOk() (*ServerCapabilities, bool)`

GetServerCapabilitiesOk returns a tuple with the ServerCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCapabilities

`func (o *IMSChargingInformation) SetServerCapabilities(v ServerCapabilities)`

SetServerCapabilities sets ServerCapabilities field to given value.

### HasServerCapabilities

`func (o *IMSChargingInformation) HasServerCapabilities() bool`

HasServerCapabilities returns a boolean if a field has been set.

### GetTrunkGroupID

`func (o *IMSChargingInformation) GetTrunkGroupID() TrunkGroupID`

GetTrunkGroupID returns the TrunkGroupID field if non-nil, zero value otherwise.

### GetTrunkGroupIDOk

`func (o *IMSChargingInformation) GetTrunkGroupIDOk() (*TrunkGroupID, bool)`

GetTrunkGroupIDOk returns a tuple with the TrunkGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkGroupID

`func (o *IMSChargingInformation) SetTrunkGroupID(v TrunkGroupID)`

SetTrunkGroupID sets TrunkGroupID field to given value.

### HasTrunkGroupID

`func (o *IMSChargingInformation) HasTrunkGroupID() bool`

HasTrunkGroupID returns a boolean if a field has been set.

### GetBearerService

`func (o *IMSChargingInformation) GetBearerService() string`

GetBearerService returns the BearerService field if non-nil, zero value otherwise.

### GetBearerServiceOk

`func (o *IMSChargingInformation) GetBearerServiceOk() (*string, bool)`

GetBearerServiceOk returns a tuple with the BearerService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerService

`func (o *IMSChargingInformation) SetBearerService(v string)`

SetBearerService sets BearerService field to given value.

### HasBearerService

`func (o *IMSChargingInformation) HasBearerService() bool`

HasBearerService returns a boolean if a field has been set.

### GetImsServiceId

`func (o *IMSChargingInformation) GetImsServiceId() string`

GetImsServiceId returns the ImsServiceId field if non-nil, zero value otherwise.

### GetImsServiceIdOk

`func (o *IMSChargingInformation) GetImsServiceIdOk() (*string, bool)`

GetImsServiceIdOk returns a tuple with the ImsServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsServiceId

`func (o *IMSChargingInformation) SetImsServiceId(v string)`

SetImsServiceId sets ImsServiceId field to given value.

### HasImsServiceId

`func (o *IMSChargingInformation) HasImsServiceId() bool`

HasImsServiceId returns a boolean if a field has been set.

### GetMessageBodies

`func (o *IMSChargingInformation) GetMessageBodies() []MessageBody`

GetMessageBodies returns the MessageBodies field if non-nil, zero value otherwise.

### GetMessageBodiesOk

`func (o *IMSChargingInformation) GetMessageBodiesOk() (*[]MessageBody, bool)`

GetMessageBodiesOk returns a tuple with the MessageBodies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBodies

`func (o *IMSChargingInformation) SetMessageBodies(v []MessageBody)`

SetMessageBodies sets MessageBodies field to given value.

### HasMessageBodies

`func (o *IMSChargingInformation) HasMessageBodies() bool`

HasMessageBodies returns a boolean if a field has been set.

### GetAccessNetworkInformation

`func (o *IMSChargingInformation) GetAccessNetworkInformation() []string`

GetAccessNetworkInformation returns the AccessNetworkInformation field if non-nil, zero value otherwise.

### GetAccessNetworkInformationOk

`func (o *IMSChargingInformation) GetAccessNetworkInformationOk() (*[]string, bool)`

GetAccessNetworkInformationOk returns a tuple with the AccessNetworkInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNetworkInformation

`func (o *IMSChargingInformation) SetAccessNetworkInformation(v []string)`

SetAccessNetworkInformation sets AccessNetworkInformation field to given value.

### HasAccessNetworkInformation

`func (o *IMSChargingInformation) HasAccessNetworkInformation() bool`

HasAccessNetworkInformation returns a boolean if a field has been set.

### GetAdditionalAccessNetworkInformation

`func (o *IMSChargingInformation) GetAdditionalAccessNetworkInformation() string`

GetAdditionalAccessNetworkInformation returns the AdditionalAccessNetworkInformation field if non-nil, zero value otherwise.

### GetAdditionalAccessNetworkInformationOk

`func (o *IMSChargingInformation) GetAdditionalAccessNetworkInformationOk() (*string, bool)`

GetAdditionalAccessNetworkInformationOk returns a tuple with the AdditionalAccessNetworkInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAccessNetworkInformation

`func (o *IMSChargingInformation) SetAdditionalAccessNetworkInformation(v string)`

SetAdditionalAccessNetworkInformation sets AdditionalAccessNetworkInformation field to given value.

### HasAdditionalAccessNetworkInformation

`func (o *IMSChargingInformation) HasAdditionalAccessNetworkInformation() bool`

HasAdditionalAccessNetworkInformation returns a boolean if a field has been set.

### GetCellularNetworkInformation

`func (o *IMSChargingInformation) GetCellularNetworkInformation() string`

GetCellularNetworkInformation returns the CellularNetworkInformation field if non-nil, zero value otherwise.

### GetCellularNetworkInformationOk

`func (o *IMSChargingInformation) GetCellularNetworkInformationOk() (*string, bool)`

GetCellularNetworkInformationOk returns a tuple with the CellularNetworkInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellularNetworkInformation

`func (o *IMSChargingInformation) SetCellularNetworkInformation(v string)`

SetCellularNetworkInformation sets CellularNetworkInformation field to given value.

### HasCellularNetworkInformation

`func (o *IMSChargingInformation) HasCellularNetworkInformation() bool`

HasCellularNetworkInformation returns a boolean if a field has been set.

### GetAccessTransferInformation

`func (o *IMSChargingInformation) GetAccessTransferInformation() []AccessTransferInformation`

GetAccessTransferInformation returns the AccessTransferInformation field if non-nil, zero value otherwise.

### GetAccessTransferInformationOk

`func (o *IMSChargingInformation) GetAccessTransferInformationOk() (*[]AccessTransferInformation, bool)`

GetAccessTransferInformationOk returns a tuple with the AccessTransferInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTransferInformation

`func (o *IMSChargingInformation) SetAccessTransferInformation(v []AccessTransferInformation)`

SetAccessTransferInformation sets AccessTransferInformation field to given value.

### HasAccessTransferInformation

`func (o *IMSChargingInformation) HasAccessTransferInformation() bool`

HasAccessTransferInformation returns a boolean if a field has been set.

### GetAccessNetworkInfoChange

`func (o *IMSChargingInformation) GetAccessNetworkInfoChange() []AccessNetworkInfoChange`

GetAccessNetworkInfoChange returns the AccessNetworkInfoChange field if non-nil, zero value otherwise.

### GetAccessNetworkInfoChangeOk

`func (o *IMSChargingInformation) GetAccessNetworkInfoChangeOk() (*[]AccessNetworkInfoChange, bool)`

GetAccessNetworkInfoChangeOk returns a tuple with the AccessNetworkInfoChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNetworkInfoChange

`func (o *IMSChargingInformation) SetAccessNetworkInfoChange(v []AccessNetworkInfoChange)`

SetAccessNetworkInfoChange sets AccessNetworkInfoChange field to given value.

### HasAccessNetworkInfoChange

`func (o *IMSChargingInformation) HasAccessNetworkInfoChange() bool`

HasAccessNetworkInfoChange returns a boolean if a field has been set.

### GetImsCommunicationServiceID

`func (o *IMSChargingInformation) GetImsCommunicationServiceID() string`

GetImsCommunicationServiceID returns the ImsCommunicationServiceID field if non-nil, zero value otherwise.

### GetImsCommunicationServiceIDOk

`func (o *IMSChargingInformation) GetImsCommunicationServiceIDOk() (*string, bool)`

GetImsCommunicationServiceIDOk returns a tuple with the ImsCommunicationServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsCommunicationServiceID

`func (o *IMSChargingInformation) SetImsCommunicationServiceID(v string)`

SetImsCommunicationServiceID sets ImsCommunicationServiceID field to given value.

### HasImsCommunicationServiceID

`func (o *IMSChargingInformation) HasImsCommunicationServiceID() bool`

HasImsCommunicationServiceID returns a boolean if a field has been set.

### GetImsApplicationReferenceID

`func (o *IMSChargingInformation) GetImsApplicationReferenceID() string`

GetImsApplicationReferenceID returns the ImsApplicationReferenceID field if non-nil, zero value otherwise.

### GetImsApplicationReferenceIDOk

`func (o *IMSChargingInformation) GetImsApplicationReferenceIDOk() (*string, bool)`

GetImsApplicationReferenceIDOk returns a tuple with the ImsApplicationReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsApplicationReferenceID

`func (o *IMSChargingInformation) SetImsApplicationReferenceID(v string)`

SetImsApplicationReferenceID sets ImsApplicationReferenceID field to given value.

### HasImsApplicationReferenceID

`func (o *IMSChargingInformation) HasImsApplicationReferenceID() bool`

HasImsApplicationReferenceID returns a boolean if a field has been set.

### GetCauseCode

`func (o *IMSChargingInformation) GetCauseCode() int32`

GetCauseCode returns the CauseCode field if non-nil, zero value otherwise.

### GetCauseCodeOk

`func (o *IMSChargingInformation) GetCauseCodeOk() (*int32, bool)`

GetCauseCodeOk returns a tuple with the CauseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseCode

`func (o *IMSChargingInformation) SetCauseCode(v int32)`

SetCauseCode sets CauseCode field to given value.

### HasCauseCode

`func (o *IMSChargingInformation) HasCauseCode() bool`

HasCauseCode returns a boolean if a field has been set.

### GetReasonHeader

`func (o *IMSChargingInformation) GetReasonHeader() []string`

GetReasonHeader returns the ReasonHeader field if non-nil, zero value otherwise.

### GetReasonHeaderOk

`func (o *IMSChargingInformation) GetReasonHeaderOk() (*[]string, bool)`

GetReasonHeaderOk returns a tuple with the ReasonHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonHeader

`func (o *IMSChargingInformation) SetReasonHeader(v []string)`

SetReasonHeader sets ReasonHeader field to given value.

### HasReasonHeader

`func (o *IMSChargingInformation) HasReasonHeader() bool`

HasReasonHeader returns a boolean if a field has been set.

### GetInitialIMSChargingIdentifier

`func (o *IMSChargingInformation) GetInitialIMSChargingIdentifier() string`

GetInitialIMSChargingIdentifier returns the InitialIMSChargingIdentifier field if non-nil, zero value otherwise.

### GetInitialIMSChargingIdentifierOk

`func (o *IMSChargingInformation) GetInitialIMSChargingIdentifierOk() (*string, bool)`

GetInitialIMSChargingIdentifierOk returns a tuple with the InitialIMSChargingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialIMSChargingIdentifier

`func (o *IMSChargingInformation) SetInitialIMSChargingIdentifier(v string)`

SetInitialIMSChargingIdentifier sets InitialIMSChargingIdentifier field to given value.

### HasInitialIMSChargingIdentifier

`func (o *IMSChargingInformation) HasInitialIMSChargingIdentifier() bool`

HasInitialIMSChargingIdentifier returns a boolean if a field has been set.

### GetNniInformation

`func (o *IMSChargingInformation) GetNniInformation() []NNIInformation`

GetNniInformation returns the NniInformation field if non-nil, zero value otherwise.

### GetNniInformationOk

`func (o *IMSChargingInformation) GetNniInformationOk() (*[]NNIInformation, bool)`

GetNniInformationOk returns a tuple with the NniInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniInformation

`func (o *IMSChargingInformation) SetNniInformation(v []NNIInformation)`

SetNniInformation sets NniInformation field to given value.

### HasNniInformation

`func (o *IMSChargingInformation) HasNniInformation() bool`

HasNniInformation returns a boolean if a field has been set.

### GetFromAddress

`func (o *IMSChargingInformation) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *IMSChargingInformation) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *IMSChargingInformation) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *IMSChargingInformation) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetImsEmergencyIndication

`func (o *IMSChargingInformation) GetImsEmergencyIndication() bool`

GetImsEmergencyIndication returns the ImsEmergencyIndication field if non-nil, zero value otherwise.

### GetImsEmergencyIndicationOk

`func (o *IMSChargingInformation) GetImsEmergencyIndicationOk() (*bool, bool)`

GetImsEmergencyIndicationOk returns a tuple with the ImsEmergencyIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsEmergencyIndication

`func (o *IMSChargingInformation) SetImsEmergencyIndication(v bool)`

SetImsEmergencyIndication sets ImsEmergencyIndication field to given value.

### HasImsEmergencyIndication

`func (o *IMSChargingInformation) HasImsEmergencyIndication() bool`

HasImsEmergencyIndication returns a boolean if a field has been set.

### GetImsVisitedNetworkIdentifier

`func (o *IMSChargingInformation) GetImsVisitedNetworkIdentifier() string`

GetImsVisitedNetworkIdentifier returns the ImsVisitedNetworkIdentifier field if non-nil, zero value otherwise.

### GetImsVisitedNetworkIdentifierOk

`func (o *IMSChargingInformation) GetImsVisitedNetworkIdentifierOk() (*string, bool)`

GetImsVisitedNetworkIdentifierOk returns a tuple with the ImsVisitedNetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsVisitedNetworkIdentifier

`func (o *IMSChargingInformation) SetImsVisitedNetworkIdentifier(v string)`

SetImsVisitedNetworkIdentifier sets ImsVisitedNetworkIdentifier field to given value.

### HasImsVisitedNetworkIdentifier

`func (o *IMSChargingInformation) HasImsVisitedNetworkIdentifier() bool`

HasImsVisitedNetworkIdentifier returns a boolean if a field has been set.

### GetSipRouteHeaderReceived

`func (o *IMSChargingInformation) GetSipRouteHeaderReceived() string`

GetSipRouteHeaderReceived returns the SipRouteHeaderReceived field if non-nil, zero value otherwise.

### GetSipRouteHeaderReceivedOk

`func (o *IMSChargingInformation) GetSipRouteHeaderReceivedOk() (*string, bool)`

GetSipRouteHeaderReceivedOk returns a tuple with the SipRouteHeaderReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipRouteHeaderReceived

`func (o *IMSChargingInformation) SetSipRouteHeaderReceived(v string)`

SetSipRouteHeaderReceived sets SipRouteHeaderReceived field to given value.

### HasSipRouteHeaderReceived

`func (o *IMSChargingInformation) HasSipRouteHeaderReceived() bool`

HasSipRouteHeaderReceived returns a boolean if a field has been set.

### GetSipRouteHeaderTransmitted

`func (o *IMSChargingInformation) GetSipRouteHeaderTransmitted() string`

GetSipRouteHeaderTransmitted returns the SipRouteHeaderTransmitted field if non-nil, zero value otherwise.

### GetSipRouteHeaderTransmittedOk

`func (o *IMSChargingInformation) GetSipRouteHeaderTransmittedOk() (*string, bool)`

GetSipRouteHeaderTransmittedOk returns a tuple with the SipRouteHeaderTransmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipRouteHeaderTransmitted

`func (o *IMSChargingInformation) SetSipRouteHeaderTransmitted(v string)`

SetSipRouteHeaderTransmitted sets SipRouteHeaderTransmitted field to given value.

### HasSipRouteHeaderTransmitted

`func (o *IMSChargingInformation) HasSipRouteHeaderTransmitted() bool`

HasSipRouteHeaderTransmitted returns a boolean if a field has been set.

### GetTadIdentifier

`func (o *IMSChargingInformation) GetTadIdentifier() TADIdentifier`

GetTadIdentifier returns the TadIdentifier field if non-nil, zero value otherwise.

### GetTadIdentifierOk

`func (o *IMSChargingInformation) GetTadIdentifierOk() (*TADIdentifier, bool)`

GetTadIdentifierOk returns a tuple with the TadIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTadIdentifier

`func (o *IMSChargingInformation) SetTadIdentifier(v TADIdentifier)`

SetTadIdentifier sets TadIdentifier field to given value.

### HasTadIdentifier

`func (o *IMSChargingInformation) HasTadIdentifier() bool`

HasTadIdentifier returns a boolean if a field has been set.

### GetFeIdentifierList

`func (o *IMSChargingInformation) GetFeIdentifierList() string`

GetFeIdentifierList returns the FeIdentifierList field if non-nil, zero value otherwise.

### GetFeIdentifierListOk

`func (o *IMSChargingInformation) GetFeIdentifierListOk() (*string, bool)`

GetFeIdentifierListOk returns a tuple with the FeIdentifierList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeIdentifierList

`func (o *IMSChargingInformation) SetFeIdentifierList(v string)`

SetFeIdentifierList sets FeIdentifierList field to given value.

### HasFeIdentifierList

`func (o *IMSChargingInformation) HasFeIdentifierList() bool`

HasFeIdentifierList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


