# PcEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**PcEvent**](PcEvent.md) |  | 
**AccType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AddAccessInfo** | Pointer to [**AdditionalAccessInfo**](AdditionalAccessInfo.md) |  | [optional] 
**RelAccessInfo** | Pointer to [**AdditionalAccessInfo**](AdditionalAccessInfo.md) |  | [optional] 
**AnGwAddr** | Pointer to [**AnGwAddress**](AnGwAddress.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**SatBackhaulCategory** | Pointer to [**SatelliteBackhaulCategory**](SatelliteBackhaulCategory.md) |  | [optional] 
**AppliedCov** | Pointer to [**ServiceAreaCoverageInfo**](ServiceAreaCoverageInfo.md) |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**PduSessionInfo** | Pointer to [**PduSessionInformation**](PduSessionInformation.md) |  | [optional] 
**RepServices** | Pointer to [**ServiceIdentification**](ServiceIdentification.md) |  | [optional] 
**DelivFailure** | Pointer to [**Failure**](Failure.md) |  | [optional] 

## Methods

### NewPcEventNotification

`func NewPcEventNotification(event PcEvent, timeStamp time.Time, ) *PcEventNotification`

NewPcEventNotification instantiates a new PcEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcEventNotificationWithDefaults

`func NewPcEventNotificationWithDefaults() *PcEventNotification`

NewPcEventNotificationWithDefaults instantiates a new PcEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PcEventNotification) GetEvent() PcEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PcEventNotification) GetEventOk() (*PcEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PcEventNotification) SetEvent(v PcEvent)`

SetEvent sets Event field to given value.


### GetAccType

`func (o *PcEventNotification) GetAccType() AccessType`

GetAccType returns the AccType field if non-nil, zero value otherwise.

### GetAccTypeOk

`func (o *PcEventNotification) GetAccTypeOk() (*AccessType, bool)`

GetAccTypeOk returns a tuple with the AccType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccType

`func (o *PcEventNotification) SetAccType(v AccessType)`

SetAccType sets AccType field to given value.

### HasAccType

`func (o *PcEventNotification) HasAccType() bool`

HasAccType returns a boolean if a field has been set.

### GetAddAccessInfo

`func (o *PcEventNotification) GetAddAccessInfo() AdditionalAccessInfo`

GetAddAccessInfo returns the AddAccessInfo field if non-nil, zero value otherwise.

### GetAddAccessInfoOk

`func (o *PcEventNotification) GetAddAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetAddAccessInfoOk returns a tuple with the AddAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessInfo

`func (o *PcEventNotification) SetAddAccessInfo(v AdditionalAccessInfo)`

SetAddAccessInfo sets AddAccessInfo field to given value.

### HasAddAccessInfo

`func (o *PcEventNotification) HasAddAccessInfo() bool`

HasAddAccessInfo returns a boolean if a field has been set.

### GetRelAccessInfo

`func (o *PcEventNotification) GetRelAccessInfo() AdditionalAccessInfo`

GetRelAccessInfo returns the RelAccessInfo field if non-nil, zero value otherwise.

### GetRelAccessInfoOk

`func (o *PcEventNotification) GetRelAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetRelAccessInfoOk returns a tuple with the RelAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelAccessInfo

`func (o *PcEventNotification) SetRelAccessInfo(v AdditionalAccessInfo)`

SetRelAccessInfo sets RelAccessInfo field to given value.

### HasRelAccessInfo

`func (o *PcEventNotification) HasRelAccessInfo() bool`

HasRelAccessInfo returns a boolean if a field has been set.

### GetAnGwAddr

`func (o *PcEventNotification) GetAnGwAddr() AnGwAddress`

GetAnGwAddr returns the AnGwAddr field if non-nil, zero value otherwise.

### GetAnGwAddrOk

`func (o *PcEventNotification) GetAnGwAddrOk() (*AnGwAddress, bool)`

GetAnGwAddrOk returns a tuple with the AnGwAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnGwAddr

`func (o *PcEventNotification) SetAnGwAddr(v AnGwAddress)`

SetAnGwAddr sets AnGwAddr field to given value.

### HasAnGwAddr

`func (o *PcEventNotification) HasAnGwAddr() bool`

HasAnGwAddr returns a boolean if a field has been set.

### GetRatType

`func (o *PcEventNotification) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *PcEventNotification) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *PcEventNotification) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *PcEventNotification) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetPlmnId

`func (o *PcEventNotification) GetPlmnId() PlmnIdNid`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PcEventNotification) GetPlmnIdOk() (*PlmnIdNid, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PcEventNotification) SetPlmnId(v PlmnIdNid)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PcEventNotification) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetSatBackhaulCategory

`func (o *PcEventNotification) GetSatBackhaulCategory() SatelliteBackhaulCategory`

GetSatBackhaulCategory returns the SatBackhaulCategory field if non-nil, zero value otherwise.

### GetSatBackhaulCategoryOk

`func (o *PcEventNotification) GetSatBackhaulCategoryOk() (*SatelliteBackhaulCategory, bool)`

GetSatBackhaulCategoryOk returns a tuple with the SatBackhaulCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatBackhaulCategory

`func (o *PcEventNotification) SetSatBackhaulCategory(v SatelliteBackhaulCategory)`

SetSatBackhaulCategory sets SatBackhaulCategory field to given value.

### HasSatBackhaulCategory

`func (o *PcEventNotification) HasSatBackhaulCategory() bool`

HasSatBackhaulCategory returns a boolean if a field has been set.

### GetAppliedCov

`func (o *PcEventNotification) GetAppliedCov() ServiceAreaCoverageInfo`

GetAppliedCov returns the AppliedCov field if non-nil, zero value otherwise.

### GetAppliedCovOk

`func (o *PcEventNotification) GetAppliedCovOk() (*ServiceAreaCoverageInfo, bool)`

GetAppliedCovOk returns a tuple with the AppliedCov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCov

`func (o *PcEventNotification) SetAppliedCov(v ServiceAreaCoverageInfo)`

SetAppliedCov sets AppliedCov field to given value.

### HasAppliedCov

`func (o *PcEventNotification) HasAppliedCov() bool`

HasAppliedCov returns a boolean if a field has been set.

### GetSupi

`func (o *PcEventNotification) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PcEventNotification) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PcEventNotification) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PcEventNotification) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *PcEventNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PcEventNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PcEventNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PcEventNotification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetTimeStamp

`func (o *PcEventNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *PcEventNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *PcEventNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetPduSessionInfo

`func (o *PcEventNotification) GetPduSessionInfo() PduSessionInformation`

GetPduSessionInfo returns the PduSessionInfo field if non-nil, zero value otherwise.

### GetPduSessionInfoOk

`func (o *PcEventNotification) GetPduSessionInfoOk() (*PduSessionInformation, bool)`

GetPduSessionInfoOk returns a tuple with the PduSessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionInfo

`func (o *PcEventNotification) SetPduSessionInfo(v PduSessionInformation)`

SetPduSessionInfo sets PduSessionInfo field to given value.

### HasPduSessionInfo

`func (o *PcEventNotification) HasPduSessionInfo() bool`

HasPduSessionInfo returns a boolean if a field has been set.

### GetRepServices

`func (o *PcEventNotification) GetRepServices() ServiceIdentification`

GetRepServices returns the RepServices field if non-nil, zero value otherwise.

### GetRepServicesOk

`func (o *PcEventNotification) GetRepServicesOk() (*ServiceIdentification, bool)`

GetRepServicesOk returns a tuple with the RepServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepServices

`func (o *PcEventNotification) SetRepServices(v ServiceIdentification)`

SetRepServices sets RepServices field to given value.

### HasRepServices

`func (o *PcEventNotification) HasRepServices() bool`

HasRepServices returns a boolean if a field has been set.

### GetDelivFailure

`func (o *PcEventNotification) GetDelivFailure() Failure`

GetDelivFailure returns the DelivFailure field if non-nil, zero value otherwise.

### GetDelivFailureOk

`func (o *PcEventNotification) GetDelivFailureOk() (*Failure, bool)`

GetDelivFailureOk returns a tuple with the DelivFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivFailure

`func (o *PcEventNotification) SetDelivFailure(v Failure)`

SetDelivFailure sets DelivFailure field to given value.

### HasDelivFailure

`func (o *PcEventNotification) HasDelivFailure() bool`

HasDelivFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


