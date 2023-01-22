# LocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**LocationEstimate** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**LocalLocationEstimate** | Pointer to [**LocalArea**](LocalArea.md) |  | [optional] 
**AgeOfLocationEstimate** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**TimestampOfLocationEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PositioningDataList** | Pointer to [**[]PositioningMethodAndUsage**](PositioningMethodAndUsage.md) |  | [optional] 
**GnssPositioningDataList** | Pointer to [**[]GnssPositioningMethodAndUsage**](GnssPositioningMethodAndUsage.md) |  | [optional] 
**AccuracyFulfilmentIndicator** | Pointer to [**AccuracyFulfilmentIndicator**](AccuracyFulfilmentIndicator.md) |  | [optional] 
**UeVelocity** | Pointer to [**VelocityEstimate**](VelocityEstimate.md) |  | [optional] 
**LdrReference** | Pointer to **string** | LDR Reference. | [optional] 
**Altitude** | Pointer to **float64** | Indicates value of altitude. | [optional] 
**ServingLMFIdentification** | Pointer to **string** | LMF identification. | [optional] 
**LocationPrivacyVerResult** | Pointer to [**LocationPrivacyVerResult**](LocationPrivacyVerResult.md) |  | [optional] 
**SuccessType** | Pointer to [**SuccessType**](SuccessType.md) |  | [optional] 
**AchievedQos** | Pointer to [**MinorLocationQoS**](MinorLocationQoS.md) |  | [optional] 

## Methods

### NewLocationData

`func NewLocationData() *LocationData`

NewLocationData instantiates a new LocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationDataWithDefaults

`func NewLocationDataWithDefaults() *LocationData`

NewLocationDataWithDefaults instantiates a new LocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *LocationData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *LocationData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *LocationData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *LocationData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *LocationData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *LocationData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *LocationData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *LocationData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetLocationEstimate

`func (o *LocationData) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *LocationData) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *LocationData) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.

### HasLocationEstimate

`func (o *LocationData) HasLocationEstimate() bool`

HasLocationEstimate returns a boolean if a field has been set.

### GetCivicAddress

`func (o *LocationData) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *LocationData) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *LocationData) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *LocationData) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetLocalLocationEstimate

`func (o *LocationData) GetLocalLocationEstimate() LocalArea`

GetLocalLocationEstimate returns the LocalLocationEstimate field if non-nil, zero value otherwise.

### GetLocalLocationEstimateOk

`func (o *LocationData) GetLocalLocationEstimateOk() (*LocalArea, bool)`

GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLocationEstimate

`func (o *LocationData) SetLocalLocationEstimate(v LocalArea)`

SetLocalLocationEstimate sets LocalLocationEstimate field to given value.

### HasLocalLocationEstimate

`func (o *LocationData) HasLocalLocationEstimate() bool`

HasLocalLocationEstimate returns a boolean if a field has been set.

### GetAgeOfLocationEstimate

`func (o *LocationData) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *LocationData) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *LocationData) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.

### HasAgeOfLocationEstimate

`func (o *LocationData) HasAgeOfLocationEstimate() bool`

HasAgeOfLocationEstimate returns a boolean if a field has been set.

### GetTimestampOfLocationEstimate

`func (o *LocationData) GetTimestampOfLocationEstimate() time.Time`

GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field if non-nil, zero value otherwise.

### GetTimestampOfLocationEstimateOk

`func (o *LocationData) GetTimestampOfLocationEstimateOk() (*time.Time, bool)`

GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfLocationEstimate

`func (o *LocationData) SetTimestampOfLocationEstimate(v time.Time)`

SetTimestampOfLocationEstimate sets TimestampOfLocationEstimate field to given value.

### HasTimestampOfLocationEstimate

`func (o *LocationData) HasTimestampOfLocationEstimate() bool`

HasTimestampOfLocationEstimate returns a boolean if a field has been set.

### GetPositioningDataList

`func (o *LocationData) GetPositioningDataList() []PositioningMethodAndUsage`

GetPositioningDataList returns the PositioningDataList field if non-nil, zero value otherwise.

### GetPositioningDataListOk

`func (o *LocationData) GetPositioningDataListOk() (*[]PositioningMethodAndUsage, bool)`

GetPositioningDataListOk returns a tuple with the PositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningDataList

`func (o *LocationData) SetPositioningDataList(v []PositioningMethodAndUsage)`

SetPositioningDataList sets PositioningDataList field to given value.

### HasPositioningDataList

`func (o *LocationData) HasPositioningDataList() bool`

HasPositioningDataList returns a boolean if a field has been set.

### GetGnssPositioningDataList

`func (o *LocationData) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage`

GetGnssPositioningDataList returns the GnssPositioningDataList field if non-nil, zero value otherwise.

### GetGnssPositioningDataListOk

`func (o *LocationData) GetGnssPositioningDataListOk() (*[]GnssPositioningMethodAndUsage, bool)`

GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnssPositioningDataList

`func (o *LocationData) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage)`

SetGnssPositioningDataList sets GnssPositioningDataList field to given value.

### HasGnssPositioningDataList

`func (o *LocationData) HasGnssPositioningDataList() bool`

HasGnssPositioningDataList returns a boolean if a field has been set.

### GetAccuracyFulfilmentIndicator

`func (o *LocationData) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator`

GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field if non-nil, zero value otherwise.

### GetAccuracyFulfilmentIndicatorOk

`func (o *LocationData) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool)`

GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyFulfilmentIndicator

`func (o *LocationData) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator)`

SetAccuracyFulfilmentIndicator sets AccuracyFulfilmentIndicator field to given value.

### HasAccuracyFulfilmentIndicator

`func (o *LocationData) HasAccuracyFulfilmentIndicator() bool`

HasAccuracyFulfilmentIndicator returns a boolean if a field has been set.

### GetUeVelocity

`func (o *LocationData) GetUeVelocity() VelocityEstimate`

GetUeVelocity returns the UeVelocity field if non-nil, zero value otherwise.

### GetUeVelocityOk

`func (o *LocationData) GetUeVelocityOk() (*VelocityEstimate, bool)`

GetUeVelocityOk returns a tuple with the UeVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeVelocity

`func (o *LocationData) SetUeVelocity(v VelocityEstimate)`

SetUeVelocity sets UeVelocity field to given value.

### HasUeVelocity

`func (o *LocationData) HasUeVelocity() bool`

HasUeVelocity returns a boolean if a field has been set.

### GetLdrReference

`func (o *LocationData) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *LocationData) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *LocationData) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.

### HasLdrReference

`func (o *LocationData) HasLdrReference() bool`

HasLdrReference returns a boolean if a field has been set.

### GetAltitude

`func (o *LocationData) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *LocationData) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *LocationData) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *LocationData) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetServingLMFIdentification

`func (o *LocationData) GetServingLMFIdentification() string`

GetServingLMFIdentification returns the ServingLMFIdentification field if non-nil, zero value otherwise.

### GetServingLMFIdentificationOk

`func (o *LocationData) GetServingLMFIdentificationOk() (*string, bool)`

GetServingLMFIdentificationOk returns a tuple with the ServingLMFIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingLMFIdentification

`func (o *LocationData) SetServingLMFIdentification(v string)`

SetServingLMFIdentification sets ServingLMFIdentification field to given value.

### HasServingLMFIdentification

`func (o *LocationData) HasServingLMFIdentification() bool`

HasServingLMFIdentification returns a boolean if a field has been set.

### GetLocationPrivacyVerResult

`func (o *LocationData) GetLocationPrivacyVerResult() LocationPrivacyVerResult`

GetLocationPrivacyVerResult returns the LocationPrivacyVerResult field if non-nil, zero value otherwise.

### GetLocationPrivacyVerResultOk

`func (o *LocationData) GetLocationPrivacyVerResultOk() (*LocationPrivacyVerResult, bool)`

GetLocationPrivacyVerResultOk returns a tuple with the LocationPrivacyVerResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationPrivacyVerResult

`func (o *LocationData) SetLocationPrivacyVerResult(v LocationPrivacyVerResult)`

SetLocationPrivacyVerResult sets LocationPrivacyVerResult field to given value.

### HasLocationPrivacyVerResult

`func (o *LocationData) HasLocationPrivacyVerResult() bool`

HasLocationPrivacyVerResult returns a boolean if a field has been set.

### GetSuccessType

`func (o *LocationData) GetSuccessType() SuccessType`

GetSuccessType returns the SuccessType field if non-nil, zero value otherwise.

### GetSuccessTypeOk

`func (o *LocationData) GetSuccessTypeOk() (*SuccessType, bool)`

GetSuccessTypeOk returns a tuple with the SuccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessType

`func (o *LocationData) SetSuccessType(v SuccessType)`

SetSuccessType sets SuccessType field to given value.

### HasSuccessType

`func (o *LocationData) HasSuccessType() bool`

HasSuccessType returns a boolean if a field has been set.

### GetAchievedQos

`func (o *LocationData) GetAchievedQos() MinorLocationQoS`

GetAchievedQos returns the AchievedQos field if non-nil, zero value otherwise.

### GetAchievedQosOk

`func (o *LocationData) GetAchievedQosOk() (*MinorLocationQoS, bool)`

GetAchievedQosOk returns a tuple with the AchievedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedQos

`func (o *LocationData) SetAchievedQos(v MinorLocationQoS)`

SetAchievedQos sets AchievedQos field to given value.

### HasAchievedQos

`func (o *LocationData) HasAchievedQos() bool`

HasAchievedQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


