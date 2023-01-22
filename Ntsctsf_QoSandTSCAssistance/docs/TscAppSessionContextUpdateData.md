# TscAppSessionContextUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AppId** | Pointer to **string** | Identifies the Application Identifier. | [optional] 
**EthFlowInfo** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) |  | [optional] 
**FlowInfo** | Pointer to [**[]FlowInfo**](FlowInfo.md) |  | [optional] 
**TscQosReq** | Pointer to [**TscQosRequirementRm**](TscQosRequirementRm.md) |  | [optional] 
**QosReference** | Pointer to **string** | Identifies a pre-defined QoS information. | [optional] 
**AltQosReferences** | Pointer to **[]string** | Identifies an ordered list of pre-defined QoS information. | [optional] 
**AltQosReqs** | Pointer to [**[]AlternativeServiceRequirementsData**](AlternativeServiceRequirementsData.md) | Identifies an ordered list of alternative service requirements that include individual QoS parameter sets. The lower the index of the array for a given entry, the higher the priority.  | [optional] 
**AspId** | Pointer to **string** | Contains an identity of an application service provider. | [optional] 
**SponId** | Pointer to **string** | Contains an identity of a sponsor. | [optional] 
**SponStatus** | Pointer to [**SponsoringStatus**](SponsoringStatus.md) |  | [optional] 
**EvSubsc** | Pointer to [**NullableEventsSubscReqDataRm**](EventsSubscReqDataRm.md) |  | [optional] 

## Methods

### NewTscAppSessionContextUpdateData

`func NewTscAppSessionContextUpdateData() *TscAppSessionContextUpdateData`

NewTscAppSessionContextUpdateData instantiates a new TscAppSessionContextUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscAppSessionContextUpdateDataWithDefaults

`func NewTscAppSessionContextUpdateDataWithDefaults() *TscAppSessionContextUpdateData`

NewTscAppSessionContextUpdateDataWithDefaults instantiates a new TscAppSessionContextUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifUri

`func (o *TscAppSessionContextUpdateData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *TscAppSessionContextUpdateData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *TscAppSessionContextUpdateData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *TscAppSessionContextUpdateData) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### GetAppId

`func (o *TscAppSessionContextUpdateData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *TscAppSessionContextUpdateData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *TscAppSessionContextUpdateData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *TscAppSessionContextUpdateData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetEthFlowInfo

`func (o *TscAppSessionContextUpdateData) GetEthFlowInfo() []EthFlowDescription`

GetEthFlowInfo returns the EthFlowInfo field if non-nil, zero value otherwise.

### GetEthFlowInfoOk

`func (o *TscAppSessionContextUpdateData) GetEthFlowInfoOk() (*[]EthFlowDescription, bool)`

GetEthFlowInfoOk returns a tuple with the EthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowInfo

`func (o *TscAppSessionContextUpdateData) SetEthFlowInfo(v []EthFlowDescription)`

SetEthFlowInfo sets EthFlowInfo field to given value.

### HasEthFlowInfo

`func (o *TscAppSessionContextUpdateData) HasEthFlowInfo() bool`

HasEthFlowInfo returns a boolean if a field has been set.

### GetFlowInfo

`func (o *TscAppSessionContextUpdateData) GetFlowInfo() []FlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *TscAppSessionContextUpdateData) GetFlowInfoOk() (*[]FlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *TscAppSessionContextUpdateData) SetFlowInfo(v []FlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *TscAppSessionContextUpdateData) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetTscQosReq

`func (o *TscAppSessionContextUpdateData) GetTscQosReq() TscQosRequirementRm`

GetTscQosReq returns the TscQosReq field if non-nil, zero value otherwise.

### GetTscQosReqOk

`func (o *TscAppSessionContextUpdateData) GetTscQosReqOk() (*TscQosRequirementRm, bool)`

GetTscQosReqOk returns a tuple with the TscQosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscQosReq

`func (o *TscAppSessionContextUpdateData) SetTscQosReq(v TscQosRequirementRm)`

SetTscQosReq sets TscQosReq field to given value.

### HasTscQosReq

`func (o *TscAppSessionContextUpdateData) HasTscQosReq() bool`

HasTscQosReq returns a boolean if a field has been set.

### GetQosReference

`func (o *TscAppSessionContextUpdateData) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *TscAppSessionContextUpdateData) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *TscAppSessionContextUpdateData) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *TscAppSessionContextUpdateData) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetAltQosReferences

`func (o *TscAppSessionContextUpdateData) GetAltQosReferences() []string`

GetAltQosReferences returns the AltQosReferences field if non-nil, zero value otherwise.

### GetAltQosReferencesOk

`func (o *TscAppSessionContextUpdateData) GetAltQosReferencesOk() (*[]string, bool)`

GetAltQosReferencesOk returns a tuple with the AltQosReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosReferences

`func (o *TscAppSessionContextUpdateData) SetAltQosReferences(v []string)`

SetAltQosReferences sets AltQosReferences field to given value.

### HasAltQosReferences

`func (o *TscAppSessionContextUpdateData) HasAltQosReferences() bool`

HasAltQosReferences returns a boolean if a field has been set.

### GetAltQosReqs

`func (o *TscAppSessionContextUpdateData) GetAltQosReqs() []AlternativeServiceRequirementsData`

GetAltQosReqs returns the AltQosReqs field if non-nil, zero value otherwise.

### GetAltQosReqsOk

`func (o *TscAppSessionContextUpdateData) GetAltQosReqsOk() (*[]AlternativeServiceRequirementsData, bool)`

GetAltQosReqsOk returns a tuple with the AltQosReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosReqs

`func (o *TscAppSessionContextUpdateData) SetAltQosReqs(v []AlternativeServiceRequirementsData)`

SetAltQosReqs sets AltQosReqs field to given value.

### HasAltQosReqs

`func (o *TscAppSessionContextUpdateData) HasAltQosReqs() bool`

HasAltQosReqs returns a boolean if a field has been set.

### GetAspId

`func (o *TscAppSessionContextUpdateData) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *TscAppSessionContextUpdateData) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *TscAppSessionContextUpdateData) SetAspId(v string)`

SetAspId sets AspId field to given value.

### HasAspId

`func (o *TscAppSessionContextUpdateData) HasAspId() bool`

HasAspId returns a boolean if a field has been set.

### GetSponId

`func (o *TscAppSessionContextUpdateData) GetSponId() string`

GetSponId returns the SponId field if non-nil, zero value otherwise.

### GetSponIdOk

`func (o *TscAppSessionContextUpdateData) GetSponIdOk() (*string, bool)`

GetSponIdOk returns a tuple with the SponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponId

`func (o *TscAppSessionContextUpdateData) SetSponId(v string)`

SetSponId sets SponId field to given value.

### HasSponId

`func (o *TscAppSessionContextUpdateData) HasSponId() bool`

HasSponId returns a boolean if a field has been set.

### GetSponStatus

`func (o *TscAppSessionContextUpdateData) GetSponStatus() SponsoringStatus`

GetSponStatus returns the SponStatus field if non-nil, zero value otherwise.

### GetSponStatusOk

`func (o *TscAppSessionContextUpdateData) GetSponStatusOk() (*SponsoringStatus, bool)`

GetSponStatusOk returns a tuple with the SponStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponStatus

`func (o *TscAppSessionContextUpdateData) SetSponStatus(v SponsoringStatus)`

SetSponStatus sets SponStatus field to given value.

### HasSponStatus

`func (o *TscAppSessionContextUpdateData) HasSponStatus() bool`

HasSponStatus returns a boolean if a field has been set.

### GetEvSubsc

`func (o *TscAppSessionContextUpdateData) GetEvSubsc() EventsSubscReqDataRm`

GetEvSubsc returns the EvSubsc field if non-nil, zero value otherwise.

### GetEvSubscOk

`func (o *TscAppSessionContextUpdateData) GetEvSubscOk() (*EventsSubscReqDataRm, bool)`

GetEvSubscOk returns a tuple with the EvSubsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsc

`func (o *TscAppSessionContextUpdateData) SetEvSubsc(v EventsSubscReqDataRm)`

SetEvSubsc sets EvSubsc field to given value.

### HasEvSubsc

`func (o *TscAppSessionContextUpdateData) HasEvSubsc() bool`

HasEvSubsc returns a boolean if a field has been set.

### SetEvSubscNil

`func (o *TscAppSessionContextUpdateData) SetEvSubscNil(b bool)`

 SetEvSubscNil sets the value for EvSubsc to be an explicit nil

### UnsetEvSubsc
`func (o *TscAppSessionContextUpdateData) UnsetEvSubsc()`

UnsetEvSubsc ensures that no value is present for EvSubsc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


