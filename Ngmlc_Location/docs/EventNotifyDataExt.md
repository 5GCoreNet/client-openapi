# EventNotifyDataExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**LdrReference** | **string** | LDR Reference. | 
**EventNotifyDataType** | [**EventNotifyDataType**](EventNotifyDataType.md) |  | 
**LocationEstimate** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**LocalLocationEstimate** | Pointer to [**LocalArea**](LocalArea.md) |  | [optional] 
**AgeOfLocationEstimate** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**TimestampOfLocationEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PositioningDataList** | Pointer to [**[]PositioningMethodAndUsage**](PositioningMethodAndUsage.md) |  | [optional] 
**GnssPositioningDataList** | Pointer to [**[]GnssPositioningMethodAndUsage**](GnssPositioningMethodAndUsage.md) |  | [optional] 
**LmfIdentification** | Pointer to **string** | LMF identification. | [optional] 
**AmfId** | Pointer to **string** | String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).   | [optional] 
**TerminationCause** | Pointer to [**TerminationCause**](TerminationCause.md) |  | [optional] 
**VelocityEstimate** | Pointer to [**VelocityEstimate**](VelocityEstimate.md) |  | [optional] 
**Altitude** | Pointer to **float64** | Indicates value of altitude. | [optional] 
**TargetNode** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**AccuracyFulfilmentIndicator** | Pointer to [**AccuracyFulfilmentIndicator**](AccuracyFulfilmentIndicator.md) |  | [optional] 
**FailureCause** | Pointer to [**FailureCause**](FailureCause.md) |  | [optional] 
**AchievedQos** | Pointer to [**MinorLocationQoS**](MinorLocationQoS.md) |  | [optional] 
**AddEventDataList** | Pointer to [**[]EventNotifyData**](EventNotifyData.md) |  | [optional] 

## Methods

### NewEventNotifyDataExt

`func NewEventNotifyDataExt(ldrReference string, eventNotifyDataType EventNotifyDataType, ) *EventNotifyDataExt`

NewEventNotifyDataExt instantiates a new EventNotifyDataExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotifyDataExtWithDefaults

`func NewEventNotifyDataExtWithDefaults() *EventNotifyDataExt`

NewEventNotifyDataExtWithDefaults instantiates a new EventNotifyDataExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *EventNotifyDataExt) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *EventNotifyDataExt) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *EventNotifyDataExt) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *EventNotifyDataExt) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *EventNotifyDataExt) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *EventNotifyDataExt) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *EventNotifyDataExt) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *EventNotifyDataExt) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetLdrReference

`func (o *EventNotifyDataExt) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *EventNotifyDataExt) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *EventNotifyDataExt) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.


### GetEventNotifyDataType

`func (o *EventNotifyDataExt) GetEventNotifyDataType() EventNotifyDataType`

GetEventNotifyDataType returns the EventNotifyDataType field if non-nil, zero value otherwise.

### GetEventNotifyDataTypeOk

`func (o *EventNotifyDataExt) GetEventNotifyDataTypeOk() (*EventNotifyDataType, bool)`

GetEventNotifyDataTypeOk returns a tuple with the EventNotifyDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyDataType

`func (o *EventNotifyDataExt) SetEventNotifyDataType(v EventNotifyDataType)`

SetEventNotifyDataType sets EventNotifyDataType field to given value.


### GetLocationEstimate

`func (o *EventNotifyDataExt) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *EventNotifyDataExt) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *EventNotifyDataExt) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.

### HasLocationEstimate

`func (o *EventNotifyDataExt) HasLocationEstimate() bool`

HasLocationEstimate returns a boolean if a field has been set.

### GetCivicAddress

`func (o *EventNotifyDataExt) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *EventNotifyDataExt) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *EventNotifyDataExt) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *EventNotifyDataExt) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetLocalLocationEstimate

`func (o *EventNotifyDataExt) GetLocalLocationEstimate() LocalArea`

GetLocalLocationEstimate returns the LocalLocationEstimate field if non-nil, zero value otherwise.

### GetLocalLocationEstimateOk

`func (o *EventNotifyDataExt) GetLocalLocationEstimateOk() (*LocalArea, bool)`

GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLocationEstimate

`func (o *EventNotifyDataExt) SetLocalLocationEstimate(v LocalArea)`

SetLocalLocationEstimate sets LocalLocationEstimate field to given value.

### HasLocalLocationEstimate

`func (o *EventNotifyDataExt) HasLocalLocationEstimate() bool`

HasLocalLocationEstimate returns a boolean if a field has been set.

### GetAgeOfLocationEstimate

`func (o *EventNotifyDataExt) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *EventNotifyDataExt) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *EventNotifyDataExt) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.

### HasAgeOfLocationEstimate

`func (o *EventNotifyDataExt) HasAgeOfLocationEstimate() bool`

HasAgeOfLocationEstimate returns a boolean if a field has been set.

### GetTimestampOfLocationEstimate

`func (o *EventNotifyDataExt) GetTimestampOfLocationEstimate() time.Time`

GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field if non-nil, zero value otherwise.

### GetTimestampOfLocationEstimateOk

`func (o *EventNotifyDataExt) GetTimestampOfLocationEstimateOk() (*time.Time, bool)`

GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfLocationEstimate

`func (o *EventNotifyDataExt) SetTimestampOfLocationEstimate(v time.Time)`

SetTimestampOfLocationEstimate sets TimestampOfLocationEstimate field to given value.

### HasTimestampOfLocationEstimate

`func (o *EventNotifyDataExt) HasTimestampOfLocationEstimate() bool`

HasTimestampOfLocationEstimate returns a boolean if a field has been set.

### GetPositioningDataList

`func (o *EventNotifyDataExt) GetPositioningDataList() []PositioningMethodAndUsage`

GetPositioningDataList returns the PositioningDataList field if non-nil, zero value otherwise.

### GetPositioningDataListOk

`func (o *EventNotifyDataExt) GetPositioningDataListOk() (*[]PositioningMethodAndUsage, bool)`

GetPositioningDataListOk returns a tuple with the PositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningDataList

`func (o *EventNotifyDataExt) SetPositioningDataList(v []PositioningMethodAndUsage)`

SetPositioningDataList sets PositioningDataList field to given value.

### HasPositioningDataList

`func (o *EventNotifyDataExt) HasPositioningDataList() bool`

HasPositioningDataList returns a boolean if a field has been set.

### GetGnssPositioningDataList

`func (o *EventNotifyDataExt) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage`

GetGnssPositioningDataList returns the GnssPositioningDataList field if non-nil, zero value otherwise.

### GetGnssPositioningDataListOk

`func (o *EventNotifyDataExt) GetGnssPositioningDataListOk() (*[]GnssPositioningMethodAndUsage, bool)`

GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnssPositioningDataList

`func (o *EventNotifyDataExt) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage)`

SetGnssPositioningDataList sets GnssPositioningDataList field to given value.

### HasGnssPositioningDataList

`func (o *EventNotifyDataExt) HasGnssPositioningDataList() bool`

HasGnssPositioningDataList returns a boolean if a field has been set.

### GetLmfIdentification

`func (o *EventNotifyDataExt) GetLmfIdentification() string`

GetLmfIdentification returns the LmfIdentification field if non-nil, zero value otherwise.

### GetLmfIdentificationOk

`func (o *EventNotifyDataExt) GetLmfIdentificationOk() (*string, bool)`

GetLmfIdentificationOk returns a tuple with the LmfIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfIdentification

`func (o *EventNotifyDataExt) SetLmfIdentification(v string)`

SetLmfIdentification sets LmfIdentification field to given value.

### HasLmfIdentification

`func (o *EventNotifyDataExt) HasLmfIdentification() bool`

HasLmfIdentification returns a boolean if a field has been set.

### GetAmfId

`func (o *EventNotifyDataExt) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *EventNotifyDataExt) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *EventNotifyDataExt) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.

### HasAmfId

`func (o *EventNotifyDataExt) HasAmfId() bool`

HasAmfId returns a boolean if a field has been set.

### GetTerminationCause

`func (o *EventNotifyDataExt) GetTerminationCause() TerminationCause`

GetTerminationCause returns the TerminationCause field if non-nil, zero value otherwise.

### GetTerminationCauseOk

`func (o *EventNotifyDataExt) GetTerminationCauseOk() (*TerminationCause, bool)`

GetTerminationCauseOk returns a tuple with the TerminationCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationCause

`func (o *EventNotifyDataExt) SetTerminationCause(v TerminationCause)`

SetTerminationCause sets TerminationCause field to given value.

### HasTerminationCause

`func (o *EventNotifyDataExt) HasTerminationCause() bool`

HasTerminationCause returns a boolean if a field has been set.

### GetVelocityEstimate

`func (o *EventNotifyDataExt) GetVelocityEstimate() VelocityEstimate`

GetVelocityEstimate returns the VelocityEstimate field if non-nil, zero value otherwise.

### GetVelocityEstimateOk

`func (o *EventNotifyDataExt) GetVelocityEstimateOk() (*VelocityEstimate, bool)`

GetVelocityEstimateOk returns a tuple with the VelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityEstimate

`func (o *EventNotifyDataExt) SetVelocityEstimate(v VelocityEstimate)`

SetVelocityEstimate sets VelocityEstimate field to given value.

### HasVelocityEstimate

`func (o *EventNotifyDataExt) HasVelocityEstimate() bool`

HasVelocityEstimate returns a boolean if a field has been set.

### GetAltitude

`func (o *EventNotifyDataExt) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *EventNotifyDataExt) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *EventNotifyDataExt) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *EventNotifyDataExt) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetTargetNode

`func (o *EventNotifyDataExt) GetTargetNode() string`

GetTargetNode returns the TargetNode field if non-nil, zero value otherwise.

### GetTargetNodeOk

`func (o *EventNotifyDataExt) GetTargetNodeOk() (*string, bool)`

GetTargetNodeOk returns a tuple with the TargetNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNode

`func (o *EventNotifyDataExt) SetTargetNode(v string)`

SetTargetNode sets TargetNode field to given value.

### HasTargetNode

`func (o *EventNotifyDataExt) HasTargetNode() bool`

HasTargetNode returns a boolean if a field has been set.

### GetAccuracyFulfilmentIndicator

`func (o *EventNotifyDataExt) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator`

GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field if non-nil, zero value otherwise.

### GetAccuracyFulfilmentIndicatorOk

`func (o *EventNotifyDataExt) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool)`

GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyFulfilmentIndicator

`func (o *EventNotifyDataExt) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator)`

SetAccuracyFulfilmentIndicator sets AccuracyFulfilmentIndicator field to given value.

### HasAccuracyFulfilmentIndicator

`func (o *EventNotifyDataExt) HasAccuracyFulfilmentIndicator() bool`

HasAccuracyFulfilmentIndicator returns a boolean if a field has been set.

### GetFailureCause

`func (o *EventNotifyDataExt) GetFailureCause() FailureCause`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *EventNotifyDataExt) GetFailureCauseOk() (*FailureCause, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *EventNotifyDataExt) SetFailureCause(v FailureCause)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *EventNotifyDataExt) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetAchievedQos

`func (o *EventNotifyDataExt) GetAchievedQos() MinorLocationQoS`

GetAchievedQos returns the AchievedQos field if non-nil, zero value otherwise.

### GetAchievedQosOk

`func (o *EventNotifyDataExt) GetAchievedQosOk() (*MinorLocationQoS, bool)`

GetAchievedQosOk returns a tuple with the AchievedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedQos

`func (o *EventNotifyDataExt) SetAchievedQos(v MinorLocationQoS)`

SetAchievedQos sets AchievedQos field to given value.

### HasAchievedQos

`func (o *EventNotifyDataExt) HasAchievedQos() bool`

HasAchievedQos returns a boolean if a field has been set.

### GetAddEventDataList

`func (o *EventNotifyDataExt) GetAddEventDataList() []EventNotifyData`

GetAddEventDataList returns the AddEventDataList field if non-nil, zero value otherwise.

### GetAddEventDataListOk

`func (o *EventNotifyDataExt) GetAddEventDataListOk() (*[]EventNotifyData, bool)`

GetAddEventDataListOk returns a tuple with the AddEventDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEventDataList

`func (o *EventNotifyDataExt) SetAddEventDataList(v []EventNotifyData)`

SetAddEventDataList sets AddEventDataList field to given value.

### HasAddEventDataList

`func (o *EventNotifyDataExt) HasAddEventDataList() bool`

HasAddEventDataList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


