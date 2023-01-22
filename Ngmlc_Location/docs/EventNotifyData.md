# EventNotifyData

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

## Methods

### NewEventNotifyData

`func NewEventNotifyData(ldrReference string, eventNotifyDataType EventNotifyDataType, ) *EventNotifyData`

NewEventNotifyData instantiates a new EventNotifyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotifyDataWithDefaults

`func NewEventNotifyDataWithDefaults() *EventNotifyData`

NewEventNotifyDataWithDefaults instantiates a new EventNotifyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *EventNotifyData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *EventNotifyData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *EventNotifyData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *EventNotifyData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *EventNotifyData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *EventNotifyData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *EventNotifyData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *EventNotifyData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetLdrReference

`func (o *EventNotifyData) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *EventNotifyData) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *EventNotifyData) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.


### GetEventNotifyDataType

`func (o *EventNotifyData) GetEventNotifyDataType() EventNotifyDataType`

GetEventNotifyDataType returns the EventNotifyDataType field if non-nil, zero value otherwise.

### GetEventNotifyDataTypeOk

`func (o *EventNotifyData) GetEventNotifyDataTypeOk() (*EventNotifyDataType, bool)`

GetEventNotifyDataTypeOk returns a tuple with the EventNotifyDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyDataType

`func (o *EventNotifyData) SetEventNotifyDataType(v EventNotifyDataType)`

SetEventNotifyDataType sets EventNotifyDataType field to given value.


### GetLocationEstimate

`func (o *EventNotifyData) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *EventNotifyData) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *EventNotifyData) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.

### HasLocationEstimate

`func (o *EventNotifyData) HasLocationEstimate() bool`

HasLocationEstimate returns a boolean if a field has been set.

### GetCivicAddress

`func (o *EventNotifyData) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *EventNotifyData) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *EventNotifyData) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *EventNotifyData) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetLocalLocationEstimate

`func (o *EventNotifyData) GetLocalLocationEstimate() LocalArea`

GetLocalLocationEstimate returns the LocalLocationEstimate field if non-nil, zero value otherwise.

### GetLocalLocationEstimateOk

`func (o *EventNotifyData) GetLocalLocationEstimateOk() (*LocalArea, bool)`

GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLocationEstimate

`func (o *EventNotifyData) SetLocalLocationEstimate(v LocalArea)`

SetLocalLocationEstimate sets LocalLocationEstimate field to given value.

### HasLocalLocationEstimate

`func (o *EventNotifyData) HasLocalLocationEstimate() bool`

HasLocalLocationEstimate returns a boolean if a field has been set.

### GetAgeOfLocationEstimate

`func (o *EventNotifyData) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *EventNotifyData) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *EventNotifyData) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.

### HasAgeOfLocationEstimate

`func (o *EventNotifyData) HasAgeOfLocationEstimate() bool`

HasAgeOfLocationEstimate returns a boolean if a field has been set.

### GetTimestampOfLocationEstimate

`func (o *EventNotifyData) GetTimestampOfLocationEstimate() time.Time`

GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field if non-nil, zero value otherwise.

### GetTimestampOfLocationEstimateOk

`func (o *EventNotifyData) GetTimestampOfLocationEstimateOk() (*time.Time, bool)`

GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfLocationEstimate

`func (o *EventNotifyData) SetTimestampOfLocationEstimate(v time.Time)`

SetTimestampOfLocationEstimate sets TimestampOfLocationEstimate field to given value.

### HasTimestampOfLocationEstimate

`func (o *EventNotifyData) HasTimestampOfLocationEstimate() bool`

HasTimestampOfLocationEstimate returns a boolean if a field has been set.

### GetPositioningDataList

`func (o *EventNotifyData) GetPositioningDataList() []PositioningMethodAndUsage`

GetPositioningDataList returns the PositioningDataList field if non-nil, zero value otherwise.

### GetPositioningDataListOk

`func (o *EventNotifyData) GetPositioningDataListOk() (*[]PositioningMethodAndUsage, bool)`

GetPositioningDataListOk returns a tuple with the PositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningDataList

`func (o *EventNotifyData) SetPositioningDataList(v []PositioningMethodAndUsage)`

SetPositioningDataList sets PositioningDataList field to given value.

### HasPositioningDataList

`func (o *EventNotifyData) HasPositioningDataList() bool`

HasPositioningDataList returns a boolean if a field has been set.

### GetGnssPositioningDataList

`func (o *EventNotifyData) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage`

GetGnssPositioningDataList returns the GnssPositioningDataList field if non-nil, zero value otherwise.

### GetGnssPositioningDataListOk

`func (o *EventNotifyData) GetGnssPositioningDataListOk() (*[]GnssPositioningMethodAndUsage, bool)`

GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnssPositioningDataList

`func (o *EventNotifyData) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage)`

SetGnssPositioningDataList sets GnssPositioningDataList field to given value.

### HasGnssPositioningDataList

`func (o *EventNotifyData) HasGnssPositioningDataList() bool`

HasGnssPositioningDataList returns a boolean if a field has been set.

### GetLmfIdentification

`func (o *EventNotifyData) GetLmfIdentification() string`

GetLmfIdentification returns the LmfIdentification field if non-nil, zero value otherwise.

### GetLmfIdentificationOk

`func (o *EventNotifyData) GetLmfIdentificationOk() (*string, bool)`

GetLmfIdentificationOk returns a tuple with the LmfIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfIdentification

`func (o *EventNotifyData) SetLmfIdentification(v string)`

SetLmfIdentification sets LmfIdentification field to given value.

### HasLmfIdentification

`func (o *EventNotifyData) HasLmfIdentification() bool`

HasLmfIdentification returns a boolean if a field has been set.

### GetAmfId

`func (o *EventNotifyData) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *EventNotifyData) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *EventNotifyData) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.

### HasAmfId

`func (o *EventNotifyData) HasAmfId() bool`

HasAmfId returns a boolean if a field has been set.

### GetTerminationCause

`func (o *EventNotifyData) GetTerminationCause() TerminationCause`

GetTerminationCause returns the TerminationCause field if non-nil, zero value otherwise.

### GetTerminationCauseOk

`func (o *EventNotifyData) GetTerminationCauseOk() (*TerminationCause, bool)`

GetTerminationCauseOk returns a tuple with the TerminationCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationCause

`func (o *EventNotifyData) SetTerminationCause(v TerminationCause)`

SetTerminationCause sets TerminationCause field to given value.

### HasTerminationCause

`func (o *EventNotifyData) HasTerminationCause() bool`

HasTerminationCause returns a boolean if a field has been set.

### GetVelocityEstimate

`func (o *EventNotifyData) GetVelocityEstimate() VelocityEstimate`

GetVelocityEstimate returns the VelocityEstimate field if non-nil, zero value otherwise.

### GetVelocityEstimateOk

`func (o *EventNotifyData) GetVelocityEstimateOk() (*VelocityEstimate, bool)`

GetVelocityEstimateOk returns a tuple with the VelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityEstimate

`func (o *EventNotifyData) SetVelocityEstimate(v VelocityEstimate)`

SetVelocityEstimate sets VelocityEstimate field to given value.

### HasVelocityEstimate

`func (o *EventNotifyData) HasVelocityEstimate() bool`

HasVelocityEstimate returns a boolean if a field has been set.

### GetAltitude

`func (o *EventNotifyData) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *EventNotifyData) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *EventNotifyData) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *EventNotifyData) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetTargetNode

`func (o *EventNotifyData) GetTargetNode() string`

GetTargetNode returns the TargetNode field if non-nil, zero value otherwise.

### GetTargetNodeOk

`func (o *EventNotifyData) GetTargetNodeOk() (*string, bool)`

GetTargetNodeOk returns a tuple with the TargetNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNode

`func (o *EventNotifyData) SetTargetNode(v string)`

SetTargetNode sets TargetNode field to given value.

### HasTargetNode

`func (o *EventNotifyData) HasTargetNode() bool`

HasTargetNode returns a boolean if a field has been set.

### GetAccuracyFulfilmentIndicator

`func (o *EventNotifyData) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator`

GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field if non-nil, zero value otherwise.

### GetAccuracyFulfilmentIndicatorOk

`func (o *EventNotifyData) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool)`

GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyFulfilmentIndicator

`func (o *EventNotifyData) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator)`

SetAccuracyFulfilmentIndicator sets AccuracyFulfilmentIndicator field to given value.

### HasAccuracyFulfilmentIndicator

`func (o *EventNotifyData) HasAccuracyFulfilmentIndicator() bool`

HasAccuracyFulfilmentIndicator returns a boolean if a field has been set.

### GetFailureCause

`func (o *EventNotifyData) GetFailureCause() FailureCause`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *EventNotifyData) GetFailureCauseOk() (*FailureCause, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *EventNotifyData) SetFailureCause(v FailureCause)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *EventNotifyData) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetAchievedQos

`func (o *EventNotifyData) GetAchievedQos() MinorLocationQoS`

GetAchievedQos returns the AchievedQos field if non-nil, zero value otherwise.

### GetAchievedQosOk

`func (o *EventNotifyData) GetAchievedQosOk() (*MinorLocationQoS, bool)`

GetAchievedQosOk returns a tuple with the AchievedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedQos

`func (o *EventNotifyData) SetAchievedQos(v MinorLocationQoS)`

SetAchievedQos sets AchievedQos field to given value.

### HasAchievedQos

`func (o *EventNotifyData) HasAchievedQos() bool`

HasAchievedQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


