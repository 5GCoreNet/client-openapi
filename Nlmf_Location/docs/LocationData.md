# LocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationEstimate** | [**GeographicArea**](GeographicArea.md) |  | 
**AccuracyFulfilmentIndicator** | Pointer to [**AccuracyFulfilmentIndicator**](AccuracyFulfilmentIndicator.md) |  | [optional] 
**AgeOfLocationEstimate** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**TimestampOfLocationEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**VelocityEstimate** | Pointer to [**VelocityEstimate**](VelocityEstimate.md) |  | [optional] 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**LocalLocationEstimate** | Pointer to [**LocalArea**](LocalArea.md) |  | [optional] 
**PositioningDataList** | Pointer to [**[]PositioningMethodAndUsage**](PositioningMethodAndUsage.md) |  | [optional] 
**GnssPositioningDataList** | Pointer to [**[]GnssPositioningMethodAndUsage**](GnssPositioningMethodAndUsage.md) |  | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**Altitude** | Pointer to **float64** | Indicates value of altitude. | [optional] 
**BarometricPressure** | Pointer to **int32** | Specifies the measured uncompensated atmospheric pressure. | [optional] 
**ServingLMFIdentification** | Pointer to **string** | LMF identification. | [optional] 
**UePositioningCap** | Pointer to **string** | Positioning capabilities supported by the UE. A string encoding the \&quot;ProvideCapabilities-r9-IEs\&quot; IE as specified in clause 6.3 of 3GPP TS 37.355 (start from octet 1). | [optional] 
**UeAreaInd** | Pointer to [**UeAreaIndication**](UeAreaIndication.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AchievedQos** | Pointer to [**MinorLocationQoS**](MinorLocationQoS.md) |  | [optional] 

## Methods

### NewLocationData

`func NewLocationData(locationEstimate GeographicArea, ) *LocationData`

NewLocationData instantiates a new LocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationDataWithDefaults

`func NewLocationDataWithDefaults() *LocationData`

NewLocationDataWithDefaults instantiates a new LocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetVelocityEstimate

`func (o *LocationData) GetVelocityEstimate() VelocityEstimate`

GetVelocityEstimate returns the VelocityEstimate field if non-nil, zero value otherwise.

### GetVelocityEstimateOk

`func (o *LocationData) GetVelocityEstimateOk() (*VelocityEstimate, bool)`

GetVelocityEstimateOk returns a tuple with the VelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityEstimate

`func (o *LocationData) SetVelocityEstimate(v VelocityEstimate)`

SetVelocityEstimate sets VelocityEstimate field to given value.

### HasVelocityEstimate

`func (o *LocationData) HasVelocityEstimate() bool`

HasVelocityEstimate returns a boolean if a field has been set.

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

### GetEcgi

`func (o *LocationData) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *LocationData) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *LocationData) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *LocationData) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetNcgi

`func (o *LocationData) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *LocationData) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *LocationData) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *LocationData) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

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

### GetBarometricPressure

`func (o *LocationData) GetBarometricPressure() int32`

GetBarometricPressure returns the BarometricPressure field if non-nil, zero value otherwise.

### GetBarometricPressureOk

`func (o *LocationData) GetBarometricPressureOk() (*int32, bool)`

GetBarometricPressureOk returns a tuple with the BarometricPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarometricPressure

`func (o *LocationData) SetBarometricPressure(v int32)`

SetBarometricPressure sets BarometricPressure field to given value.

### HasBarometricPressure

`func (o *LocationData) HasBarometricPressure() bool`

HasBarometricPressure returns a boolean if a field has been set.

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

### GetUePositioningCap

`func (o *LocationData) GetUePositioningCap() string`

GetUePositioningCap returns the UePositioningCap field if non-nil, zero value otherwise.

### GetUePositioningCapOk

`func (o *LocationData) GetUePositioningCapOk() (*string, bool)`

GetUePositioningCapOk returns a tuple with the UePositioningCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePositioningCap

`func (o *LocationData) SetUePositioningCap(v string)`

SetUePositioningCap sets UePositioningCap field to given value.

### HasUePositioningCap

`func (o *LocationData) HasUePositioningCap() bool`

HasUePositioningCap returns a boolean if a field has been set.

### GetUeAreaInd

`func (o *LocationData) GetUeAreaInd() UeAreaIndication`

GetUeAreaInd returns the UeAreaInd field if non-nil, zero value otherwise.

### GetUeAreaIndOk

`func (o *LocationData) GetUeAreaIndOk() (*UeAreaIndication, bool)`

GetUeAreaIndOk returns a tuple with the UeAreaInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAreaInd

`func (o *LocationData) SetUeAreaInd(v UeAreaIndication)`

SetUeAreaInd sets UeAreaInd field to given value.

### HasUeAreaInd

`func (o *LocationData) HasUeAreaInd() bool`

HasUeAreaInd returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *LocationData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *LocationData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *LocationData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *LocationData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

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


