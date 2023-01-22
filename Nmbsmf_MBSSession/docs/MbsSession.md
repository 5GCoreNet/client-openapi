# MbsSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | Pointer to [**MbsSessionId**](MbsSessionId.md) |  | [optional] 
**TmgiAllocReq** | Pointer to **bool** |  | [optional] [default to false]
**Tmgi** | Pointer to [**Tmgi**](Tmgi.md) |  | [optional] 
**ExpirationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ServiceType** | [**MbsServiceType**](MbsServiceType.md) |  | 
**LocationDependent** | Pointer to **bool** |  | [optional] [default to false]
**AreaSessionId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | [optional] 
**IngressTunAddrReq** | Pointer to **bool** |  | [optional] [default to false]
**IngressTunAddr** | Pointer to [**[]TunnelAddress**](TunnelAddress.md) |  | [optional] [readonly] 
**Ssm** | Pointer to [**Ssm**](Ssm.md) |  | [optional] 
**MbsServiceArea** | Pointer to [**MbsServiceArea**](MbsServiceArea.md) |  | [optional] 
**ExtMbsServiceArea** | Pointer to [**ExternalMbsServiceArea**](ExternalMbsServiceArea.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ActivationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TerminationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**MbsServInfo** | Pointer to [**MbsServiceInfo**](MbsServiceInfo.md) |  | [optional] 
**MbsSessionSubsc** | Pointer to [**MbsSessionSubscription**](MbsSessionSubscription.md) |  | [optional] 
**ActivityStatus** | Pointer to [**MbsSessionActivityStatus**](MbsSessionActivityStatus.md) |  | [optional] 
**AnyUeInd** | Pointer to **bool** |  | [optional] [default to false]
**MbsFsaIdList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMbsSession

`func NewMbsSession(serviceType MbsServiceType, ) *MbsSession`

NewMbsSession instantiates a new MbsSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionWithDefaults

`func NewMbsSessionWithDefaults() *MbsSession`

NewMbsSessionWithDefaults instantiates a new MbsSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *MbsSession) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MbsSession) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MbsSession) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.

### HasMbsSessionId

`func (o *MbsSession) HasMbsSessionId() bool`

HasMbsSessionId returns a boolean if a field has been set.

### GetTmgiAllocReq

`func (o *MbsSession) GetTmgiAllocReq() bool`

GetTmgiAllocReq returns the TmgiAllocReq field if non-nil, zero value otherwise.

### GetTmgiAllocReqOk

`func (o *MbsSession) GetTmgiAllocReqOk() (*bool, bool)`

GetTmgiAllocReqOk returns a tuple with the TmgiAllocReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgiAllocReq

`func (o *MbsSession) SetTmgiAllocReq(v bool)`

SetTmgiAllocReq sets TmgiAllocReq field to given value.

### HasTmgiAllocReq

`func (o *MbsSession) HasTmgiAllocReq() bool`

HasTmgiAllocReq returns a boolean if a field has been set.

### GetTmgi

`func (o *MbsSession) GetTmgi() Tmgi`

GetTmgi returns the Tmgi field if non-nil, zero value otherwise.

### GetTmgiOk

`func (o *MbsSession) GetTmgiOk() (*Tmgi, bool)`

GetTmgiOk returns a tuple with the Tmgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgi

`func (o *MbsSession) SetTmgi(v Tmgi)`

SetTmgi sets Tmgi field to given value.

### HasTmgi

`func (o *MbsSession) HasTmgi() bool`

HasTmgi returns a boolean if a field has been set.

### GetExpirationTime

`func (o *MbsSession) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *MbsSession) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *MbsSession) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *MbsSession) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetServiceType

`func (o *MbsSession) GetServiceType() MbsServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *MbsSession) GetServiceTypeOk() (*MbsServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *MbsSession) SetServiceType(v MbsServiceType)`

SetServiceType sets ServiceType field to given value.


### GetLocationDependent

`func (o *MbsSession) GetLocationDependent() bool`

GetLocationDependent returns the LocationDependent field if non-nil, zero value otherwise.

### GetLocationDependentOk

`func (o *MbsSession) GetLocationDependentOk() (*bool, bool)`

GetLocationDependentOk returns a tuple with the LocationDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDependent

`func (o *MbsSession) SetLocationDependent(v bool)`

SetLocationDependent sets LocationDependent field to given value.

### HasLocationDependent

`func (o *MbsSession) HasLocationDependent() bool`

HasLocationDependent returns a boolean if a field has been set.

### GetAreaSessionId

`func (o *MbsSession) GetAreaSessionId() int32`

GetAreaSessionId returns the AreaSessionId field if non-nil, zero value otherwise.

### GetAreaSessionIdOk

`func (o *MbsSession) GetAreaSessionIdOk() (*int32, bool)`

GetAreaSessionIdOk returns a tuple with the AreaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaSessionId

`func (o *MbsSession) SetAreaSessionId(v int32)`

SetAreaSessionId sets AreaSessionId field to given value.

### HasAreaSessionId

`func (o *MbsSession) HasAreaSessionId() bool`

HasAreaSessionId returns a boolean if a field has been set.

### GetIngressTunAddrReq

`func (o *MbsSession) GetIngressTunAddrReq() bool`

GetIngressTunAddrReq returns the IngressTunAddrReq field if non-nil, zero value otherwise.

### GetIngressTunAddrReqOk

`func (o *MbsSession) GetIngressTunAddrReqOk() (*bool, bool)`

GetIngressTunAddrReqOk returns a tuple with the IngressTunAddrReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressTunAddrReq

`func (o *MbsSession) SetIngressTunAddrReq(v bool)`

SetIngressTunAddrReq sets IngressTunAddrReq field to given value.

### HasIngressTunAddrReq

`func (o *MbsSession) HasIngressTunAddrReq() bool`

HasIngressTunAddrReq returns a boolean if a field has been set.

### GetIngressTunAddr

`func (o *MbsSession) GetIngressTunAddr() []TunnelAddress`

GetIngressTunAddr returns the IngressTunAddr field if non-nil, zero value otherwise.

### GetIngressTunAddrOk

`func (o *MbsSession) GetIngressTunAddrOk() (*[]TunnelAddress, bool)`

GetIngressTunAddrOk returns a tuple with the IngressTunAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressTunAddr

`func (o *MbsSession) SetIngressTunAddr(v []TunnelAddress)`

SetIngressTunAddr sets IngressTunAddr field to given value.

### HasIngressTunAddr

`func (o *MbsSession) HasIngressTunAddr() bool`

HasIngressTunAddr returns a boolean if a field has been set.

### GetSsm

`func (o *MbsSession) GetSsm() Ssm`

GetSsm returns the Ssm field if non-nil, zero value otherwise.

### GetSsmOk

`func (o *MbsSession) GetSsmOk() (*Ssm, bool)`

GetSsmOk returns a tuple with the Ssm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsm

`func (o *MbsSession) SetSsm(v Ssm)`

SetSsm sets Ssm field to given value.

### HasSsm

`func (o *MbsSession) HasSsm() bool`

HasSsm returns a boolean if a field has been set.

### GetMbsServiceArea

`func (o *MbsSession) GetMbsServiceArea() MbsServiceArea`

GetMbsServiceArea returns the MbsServiceArea field if non-nil, zero value otherwise.

### GetMbsServiceAreaOk

`func (o *MbsSession) GetMbsServiceAreaOk() (*MbsServiceArea, bool)`

GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceArea

`func (o *MbsSession) SetMbsServiceArea(v MbsServiceArea)`

SetMbsServiceArea sets MbsServiceArea field to given value.

### HasMbsServiceArea

`func (o *MbsSession) HasMbsServiceArea() bool`

HasMbsServiceArea returns a boolean if a field has been set.

### GetExtMbsServiceArea

`func (o *MbsSession) GetExtMbsServiceArea() ExternalMbsServiceArea`

GetExtMbsServiceArea returns the ExtMbsServiceArea field if non-nil, zero value otherwise.

### GetExtMbsServiceAreaOk

`func (o *MbsSession) GetExtMbsServiceAreaOk() (*ExternalMbsServiceArea, bool)`

GetExtMbsServiceAreaOk returns a tuple with the ExtMbsServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMbsServiceArea

`func (o *MbsSession) SetExtMbsServiceArea(v ExternalMbsServiceArea)`

SetExtMbsServiceArea sets ExtMbsServiceArea field to given value.

### HasExtMbsServiceArea

`func (o *MbsSession) HasExtMbsServiceArea() bool`

HasExtMbsServiceArea returns a boolean if a field has been set.

### GetDnn

`func (o *MbsSession) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *MbsSession) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *MbsSession) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *MbsSession) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *MbsSession) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *MbsSession) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *MbsSession) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *MbsSession) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetActivationTime

`func (o *MbsSession) GetActivationTime() time.Time`

GetActivationTime returns the ActivationTime field if non-nil, zero value otherwise.

### GetActivationTimeOk

`func (o *MbsSession) GetActivationTimeOk() (*time.Time, bool)`

GetActivationTimeOk returns a tuple with the ActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTime

`func (o *MbsSession) SetActivationTime(v time.Time)`

SetActivationTime sets ActivationTime field to given value.

### HasActivationTime

`func (o *MbsSession) HasActivationTime() bool`

HasActivationTime returns a boolean if a field has been set.

### GetTerminationTime

`func (o *MbsSession) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *MbsSession) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *MbsSession) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *MbsSession) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetMbsServInfo

`func (o *MbsSession) GetMbsServInfo() MbsServiceInfo`

GetMbsServInfo returns the MbsServInfo field if non-nil, zero value otherwise.

### GetMbsServInfoOk

`func (o *MbsSession) GetMbsServInfoOk() (*MbsServiceInfo, bool)`

GetMbsServInfoOk returns a tuple with the MbsServInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServInfo

`func (o *MbsSession) SetMbsServInfo(v MbsServiceInfo)`

SetMbsServInfo sets MbsServInfo field to given value.

### HasMbsServInfo

`func (o *MbsSession) HasMbsServInfo() bool`

HasMbsServInfo returns a boolean if a field has been set.

### GetMbsSessionSubsc

`func (o *MbsSession) GetMbsSessionSubsc() MbsSessionSubscription`

GetMbsSessionSubsc returns the MbsSessionSubsc field if non-nil, zero value otherwise.

### GetMbsSessionSubscOk

`func (o *MbsSession) GetMbsSessionSubscOk() (*MbsSessionSubscription, bool)`

GetMbsSessionSubscOk returns a tuple with the MbsSessionSubsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionSubsc

`func (o *MbsSession) SetMbsSessionSubsc(v MbsSessionSubscription)`

SetMbsSessionSubsc sets MbsSessionSubsc field to given value.

### HasMbsSessionSubsc

`func (o *MbsSession) HasMbsSessionSubsc() bool`

HasMbsSessionSubsc returns a boolean if a field has been set.

### GetActivityStatus

`func (o *MbsSession) GetActivityStatus() MbsSessionActivityStatus`

GetActivityStatus returns the ActivityStatus field if non-nil, zero value otherwise.

### GetActivityStatusOk

`func (o *MbsSession) GetActivityStatusOk() (*MbsSessionActivityStatus, bool)`

GetActivityStatusOk returns a tuple with the ActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatus

`func (o *MbsSession) SetActivityStatus(v MbsSessionActivityStatus)`

SetActivityStatus sets ActivityStatus field to given value.

### HasActivityStatus

`func (o *MbsSession) HasActivityStatus() bool`

HasActivityStatus returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *MbsSession) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *MbsSession) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *MbsSession) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *MbsSession) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetMbsFsaIdList

`func (o *MbsSession) GetMbsFsaIdList() []string`

GetMbsFsaIdList returns the MbsFsaIdList field if non-nil, zero value otherwise.

### GetMbsFsaIdListOk

`func (o *MbsSession) GetMbsFsaIdListOk() (*[]string, bool)`

GetMbsFsaIdListOk returns a tuple with the MbsFsaIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsFsaIdList

`func (o *MbsSession) SetMbsFsaIdList(v []string)`

SetMbsFsaIdList sets MbsFsaIdList field to given value.

### HasMbsFsaIdList

`func (o *MbsSession) HasMbsFsaIdList() bool`

HasMbsFsaIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


