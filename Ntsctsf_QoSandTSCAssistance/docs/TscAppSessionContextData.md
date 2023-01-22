# TscAppSessionContextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**IpDomain** | Pointer to **string** | The IPv4 address domain identifier. | [optional] 
**UeMac** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**AppId** | Pointer to **string** | Identifies the Application Identifier. | [optional] 
**EthFlowInfo** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) |  | [optional] 
**FlowInfo** | Pointer to [**[]FlowInfo**](FlowInfo.md) |  | [optional] 
**AfId** | **string** | Identifies the AF identifier. | 
**TscQosReq** | Pointer to [**TscQosRequirement**](TscQosRequirement.md) |  | [optional] 
**QosReference** | **string** | Identifies a pre-defined QoS information. | 
**AltQosReferences** | Pointer to **[]string** | Identifies an ordered list of pre-defined QoS information. | [optional] 
**AltQosReqs** | Pointer to [**[]AlternativeServiceRequirementsData**](AlternativeServiceRequirementsData.md) | Identifies an ordered list of alternative service requirements that include individual QoS parameter sets. The lower the index of the array for a given entry, the higher the priority.  | [optional] 
**AspId** | Pointer to **string** | Contains an identity of an application service provider. | [optional] 
**SponId** | Pointer to **string** | Contains an identity of a sponsor. | [optional] 
**SponStatus** | Pointer to [**SponsoringStatus**](SponsoringStatus.md) |  | [optional] 
**EvSubsc** | Pointer to [**EventsSubscReqData**](EventsSubscReqData.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewTscAppSessionContextData

`func NewTscAppSessionContextData(notifUri string, afId string, qosReference string, ) *TscAppSessionContextData`

NewTscAppSessionContextData instantiates a new TscAppSessionContextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscAppSessionContextDataWithDefaults

`func NewTscAppSessionContextDataWithDefaults() *TscAppSessionContextData`

NewTscAppSessionContextDataWithDefaults instantiates a new TscAppSessionContextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeIpAddr

`func (o *TscAppSessionContextData) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *TscAppSessionContextData) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *TscAppSessionContextData) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *TscAppSessionContextData) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.

### GetIpDomain

`func (o *TscAppSessionContextData) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *TscAppSessionContextData) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *TscAppSessionContextData) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *TscAppSessionContextData) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetUeMac

`func (o *TscAppSessionContextData) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *TscAppSessionContextData) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *TscAppSessionContextData) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *TscAppSessionContextData) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetDnn

`func (o *TscAppSessionContextData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TscAppSessionContextData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TscAppSessionContextData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *TscAppSessionContextData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *TscAppSessionContextData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TscAppSessionContextData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TscAppSessionContextData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *TscAppSessionContextData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetNotifUri

`func (o *TscAppSessionContextData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *TscAppSessionContextData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *TscAppSessionContextData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetAppId

`func (o *TscAppSessionContextData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *TscAppSessionContextData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *TscAppSessionContextData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *TscAppSessionContextData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetEthFlowInfo

`func (o *TscAppSessionContextData) GetEthFlowInfo() []EthFlowDescription`

GetEthFlowInfo returns the EthFlowInfo field if non-nil, zero value otherwise.

### GetEthFlowInfoOk

`func (o *TscAppSessionContextData) GetEthFlowInfoOk() (*[]EthFlowDescription, bool)`

GetEthFlowInfoOk returns a tuple with the EthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowInfo

`func (o *TscAppSessionContextData) SetEthFlowInfo(v []EthFlowDescription)`

SetEthFlowInfo sets EthFlowInfo field to given value.

### HasEthFlowInfo

`func (o *TscAppSessionContextData) HasEthFlowInfo() bool`

HasEthFlowInfo returns a boolean if a field has been set.

### GetFlowInfo

`func (o *TscAppSessionContextData) GetFlowInfo() []FlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *TscAppSessionContextData) GetFlowInfoOk() (*[]FlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *TscAppSessionContextData) SetFlowInfo(v []FlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *TscAppSessionContextData) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetAfId

`func (o *TscAppSessionContextData) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *TscAppSessionContextData) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *TscAppSessionContextData) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetTscQosReq

`func (o *TscAppSessionContextData) GetTscQosReq() TscQosRequirement`

GetTscQosReq returns the TscQosReq field if non-nil, zero value otherwise.

### GetTscQosReqOk

`func (o *TscAppSessionContextData) GetTscQosReqOk() (*TscQosRequirement, bool)`

GetTscQosReqOk returns a tuple with the TscQosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscQosReq

`func (o *TscAppSessionContextData) SetTscQosReq(v TscQosRequirement)`

SetTscQosReq sets TscQosReq field to given value.

### HasTscQosReq

`func (o *TscAppSessionContextData) HasTscQosReq() bool`

HasTscQosReq returns a boolean if a field has been set.

### GetQosReference

`func (o *TscAppSessionContextData) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *TscAppSessionContextData) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *TscAppSessionContextData) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.


### GetAltQosReferences

`func (o *TscAppSessionContextData) GetAltQosReferences() []string`

GetAltQosReferences returns the AltQosReferences field if non-nil, zero value otherwise.

### GetAltQosReferencesOk

`func (o *TscAppSessionContextData) GetAltQosReferencesOk() (*[]string, bool)`

GetAltQosReferencesOk returns a tuple with the AltQosReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosReferences

`func (o *TscAppSessionContextData) SetAltQosReferences(v []string)`

SetAltQosReferences sets AltQosReferences field to given value.

### HasAltQosReferences

`func (o *TscAppSessionContextData) HasAltQosReferences() bool`

HasAltQosReferences returns a boolean if a field has been set.

### GetAltQosReqs

`func (o *TscAppSessionContextData) GetAltQosReqs() []AlternativeServiceRequirementsData`

GetAltQosReqs returns the AltQosReqs field if non-nil, zero value otherwise.

### GetAltQosReqsOk

`func (o *TscAppSessionContextData) GetAltQosReqsOk() (*[]AlternativeServiceRequirementsData, bool)`

GetAltQosReqsOk returns a tuple with the AltQosReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosReqs

`func (o *TscAppSessionContextData) SetAltQosReqs(v []AlternativeServiceRequirementsData)`

SetAltQosReqs sets AltQosReqs field to given value.

### HasAltQosReqs

`func (o *TscAppSessionContextData) HasAltQosReqs() bool`

HasAltQosReqs returns a boolean if a field has been set.

### GetAspId

`func (o *TscAppSessionContextData) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *TscAppSessionContextData) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *TscAppSessionContextData) SetAspId(v string)`

SetAspId sets AspId field to given value.

### HasAspId

`func (o *TscAppSessionContextData) HasAspId() bool`

HasAspId returns a boolean if a field has been set.

### GetSponId

`func (o *TscAppSessionContextData) GetSponId() string`

GetSponId returns the SponId field if non-nil, zero value otherwise.

### GetSponIdOk

`func (o *TscAppSessionContextData) GetSponIdOk() (*string, bool)`

GetSponIdOk returns a tuple with the SponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponId

`func (o *TscAppSessionContextData) SetSponId(v string)`

SetSponId sets SponId field to given value.

### HasSponId

`func (o *TscAppSessionContextData) HasSponId() bool`

HasSponId returns a boolean if a field has been set.

### GetSponStatus

`func (o *TscAppSessionContextData) GetSponStatus() SponsoringStatus`

GetSponStatus returns the SponStatus field if non-nil, zero value otherwise.

### GetSponStatusOk

`func (o *TscAppSessionContextData) GetSponStatusOk() (*SponsoringStatus, bool)`

GetSponStatusOk returns a tuple with the SponStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponStatus

`func (o *TscAppSessionContextData) SetSponStatus(v SponsoringStatus)`

SetSponStatus sets SponStatus field to given value.

### HasSponStatus

`func (o *TscAppSessionContextData) HasSponStatus() bool`

HasSponStatus returns a boolean if a field has been set.

### GetEvSubsc

`func (o *TscAppSessionContextData) GetEvSubsc() EventsSubscReqData`

GetEvSubsc returns the EvSubsc field if non-nil, zero value otherwise.

### GetEvSubscOk

`func (o *TscAppSessionContextData) GetEvSubscOk() (*EventsSubscReqData, bool)`

GetEvSubscOk returns a tuple with the EvSubsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsc

`func (o *TscAppSessionContextData) SetEvSubsc(v EventsSubscReqData)`

SetEvSubsc sets EvSubsc field to given value.

### HasEvSubsc

`func (o *TscAppSessionContextData) HasEvSubsc() bool`

HasEvSubsc returns a boolean if a field has been set.

### GetSuppFeat

`func (o *TscAppSessionContextData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *TscAppSessionContextData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *TscAppSessionContextData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *TscAppSessionContextData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


