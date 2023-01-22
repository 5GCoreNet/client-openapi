# ChargeablePartyPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Describes the IP flows. | [optional] 
**ExterAppId** | Pointer to **string** | Identifies the external Application Identifier. | [optional] 
**EthFlowInfo** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet flows. | [optional] 
**SponsoringEnabled** | Pointer to **bool** | Indicates whether the sponsoring data connectivity is enabled (true) or not (false).  | [optional] 
**ReferenceId** | Pointer to **string** | string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154. | [optional] 
**UsageThreshold** | Pointer to [**NullableUsageThresholdRm**](UsageThresholdRm.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**Events** | Pointer to [**[]Event**](Event.md) | Represents the list of event(s) to which the SCS/AS requests to subscribe to. | [optional] 

## Methods

### NewChargeablePartyPatch

`func NewChargeablePartyPatch() *ChargeablePartyPatch`

NewChargeablePartyPatch instantiates a new ChargeablePartyPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeablePartyPatchWithDefaults

`func NewChargeablePartyPatchWithDefaults() *ChargeablePartyPatch`

NewChargeablePartyPatchWithDefaults instantiates a new ChargeablePartyPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *ChargeablePartyPatch) GetFlowInfo() []FlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *ChargeablePartyPatch) GetFlowInfoOk() (*[]FlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *ChargeablePartyPatch) SetFlowInfo(v []FlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *ChargeablePartyPatch) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetExterAppId

`func (o *ChargeablePartyPatch) GetExterAppId() string`

GetExterAppId returns the ExterAppId field if non-nil, zero value otherwise.

### GetExterAppIdOk

`func (o *ChargeablePartyPatch) GetExterAppIdOk() (*string, bool)`

GetExterAppIdOk returns a tuple with the ExterAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterAppId

`func (o *ChargeablePartyPatch) SetExterAppId(v string)`

SetExterAppId sets ExterAppId field to given value.

### HasExterAppId

`func (o *ChargeablePartyPatch) HasExterAppId() bool`

HasExterAppId returns a boolean if a field has been set.

### GetEthFlowInfo

`func (o *ChargeablePartyPatch) GetEthFlowInfo() []EthFlowDescription`

GetEthFlowInfo returns the EthFlowInfo field if non-nil, zero value otherwise.

### GetEthFlowInfoOk

`func (o *ChargeablePartyPatch) GetEthFlowInfoOk() (*[]EthFlowDescription, bool)`

GetEthFlowInfoOk returns a tuple with the EthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowInfo

`func (o *ChargeablePartyPatch) SetEthFlowInfo(v []EthFlowDescription)`

SetEthFlowInfo sets EthFlowInfo field to given value.

### HasEthFlowInfo

`func (o *ChargeablePartyPatch) HasEthFlowInfo() bool`

HasEthFlowInfo returns a boolean if a field has been set.

### GetSponsoringEnabled

`func (o *ChargeablePartyPatch) GetSponsoringEnabled() bool`

GetSponsoringEnabled returns the SponsoringEnabled field if non-nil, zero value otherwise.

### GetSponsoringEnabledOk

`func (o *ChargeablePartyPatch) GetSponsoringEnabledOk() (*bool, bool)`

GetSponsoringEnabledOk returns a tuple with the SponsoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoringEnabled

`func (o *ChargeablePartyPatch) SetSponsoringEnabled(v bool)`

SetSponsoringEnabled sets SponsoringEnabled field to given value.

### HasSponsoringEnabled

`func (o *ChargeablePartyPatch) HasSponsoringEnabled() bool`

HasSponsoringEnabled returns a boolean if a field has been set.

### GetReferenceId

`func (o *ChargeablePartyPatch) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ChargeablePartyPatch) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ChargeablePartyPatch) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ChargeablePartyPatch) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *ChargeablePartyPatch) GetUsageThreshold() UsageThresholdRm`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *ChargeablePartyPatch) GetUsageThresholdOk() (*UsageThresholdRm, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *ChargeablePartyPatch) SetUsageThreshold(v UsageThresholdRm)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *ChargeablePartyPatch) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### SetUsageThresholdNil

`func (o *ChargeablePartyPatch) SetUsageThresholdNil(b bool)`

 SetUsageThresholdNil sets the value for UsageThreshold to be an explicit nil

### UnsetUsageThreshold
`func (o *ChargeablePartyPatch) UnsetUsageThreshold()`

UnsetUsageThreshold ensures that no value is present for UsageThreshold, not even an explicit nil
### GetNotificationDestination

`func (o *ChargeablePartyPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ChargeablePartyPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ChargeablePartyPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *ChargeablePartyPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetEvents

`func (o *ChargeablePartyPatch) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ChargeablePartyPatch) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ChargeablePartyPatch) SetEvents(v []Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ChargeablePartyPatch) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


