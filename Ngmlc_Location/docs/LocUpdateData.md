# LocUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**PseudonymIndicator** | Pointer to [**PseudonymIndicator**](PseudonymIndicator.md) |  | [optional] 
**LocationRequestType** | [**LocationRequestType**](LocationRequestType.md) |  | 
**LocationEstimate** | [**GeographicArea**](GeographicArea.md) |  | 
**AgeOfLocationEstimate** | **int32** | Indicates value of the age of the location estimate. | 
**TimestampOfLocationEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AccuracyFulfilmentIndicator** | [**AccuracyFulfilmentIndicator**](AccuracyFulfilmentIndicator.md) |  | 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**LcsQosClass** | [**LcsQosClass**](LcsQosClass.md) |  | 
**ExternalClientIdentification** | Pointer to **string** | Contains the external client identification | [optional] 
**AfId** | Pointer to **string** |  | [optional] 
**GmlcNumber** | Pointer to **string** |  | [optional] 
**LcsServiceType** | Pointer to **int32** | LCS Service Type Id. | [optional] 

## Methods

### NewLocUpdateData

`func NewLocUpdateData(locationRequestType LocationRequestType, locationEstimate GeographicArea, ageOfLocationEstimate int32, accuracyFulfilmentIndicator AccuracyFulfilmentIndicator, lcsQosClass LcsQosClass, ) *LocUpdateData`

NewLocUpdateData instantiates a new LocUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocUpdateDataWithDefaults

`func NewLocUpdateDataWithDefaults() *LocUpdateData`

NewLocUpdateDataWithDefaults instantiates a new LocUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *LocUpdateData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *LocUpdateData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *LocUpdateData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *LocUpdateData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *LocUpdateData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *LocUpdateData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *LocUpdateData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *LocUpdateData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetPseudonymIndicator

`func (o *LocUpdateData) GetPseudonymIndicator() PseudonymIndicator`

GetPseudonymIndicator returns the PseudonymIndicator field if non-nil, zero value otherwise.

### GetPseudonymIndicatorOk

`func (o *LocUpdateData) GetPseudonymIndicatorOk() (*PseudonymIndicator, bool)`

GetPseudonymIndicatorOk returns a tuple with the PseudonymIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonymIndicator

`func (o *LocUpdateData) SetPseudonymIndicator(v PseudonymIndicator)`

SetPseudonymIndicator sets PseudonymIndicator field to given value.

### HasPseudonymIndicator

`func (o *LocUpdateData) HasPseudonymIndicator() bool`

HasPseudonymIndicator returns a boolean if a field has been set.

### GetLocationRequestType

`func (o *LocUpdateData) GetLocationRequestType() LocationRequestType`

GetLocationRequestType returns the LocationRequestType field if non-nil, zero value otherwise.

### GetLocationRequestTypeOk

`func (o *LocUpdateData) GetLocationRequestTypeOk() (*LocationRequestType, bool)`

GetLocationRequestTypeOk returns a tuple with the LocationRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationRequestType

`func (o *LocUpdateData) SetLocationRequestType(v LocationRequestType)`

SetLocationRequestType sets LocationRequestType field to given value.


### GetLocationEstimate

`func (o *LocUpdateData) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *LocUpdateData) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *LocUpdateData) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.


### GetAgeOfLocationEstimate

`func (o *LocUpdateData) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *LocUpdateData) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *LocUpdateData) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.


### GetTimestampOfLocationEstimate

`func (o *LocUpdateData) GetTimestampOfLocationEstimate() time.Time`

GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field if non-nil, zero value otherwise.

### GetTimestampOfLocationEstimateOk

`func (o *LocUpdateData) GetTimestampOfLocationEstimateOk() (*time.Time, bool)`

GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfLocationEstimate

`func (o *LocUpdateData) SetTimestampOfLocationEstimate(v time.Time)`

SetTimestampOfLocationEstimate sets TimestampOfLocationEstimate field to given value.

### HasTimestampOfLocationEstimate

`func (o *LocUpdateData) HasTimestampOfLocationEstimate() bool`

HasTimestampOfLocationEstimate returns a boolean if a field has been set.

### GetAccuracyFulfilmentIndicator

`func (o *LocUpdateData) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator`

GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field if non-nil, zero value otherwise.

### GetAccuracyFulfilmentIndicatorOk

`func (o *LocUpdateData) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool)`

GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyFulfilmentIndicator

`func (o *LocUpdateData) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator)`

SetAccuracyFulfilmentIndicator sets AccuracyFulfilmentIndicator field to given value.


### GetCivicAddress

`func (o *LocUpdateData) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *LocUpdateData) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *LocUpdateData) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *LocUpdateData) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetLcsQosClass

`func (o *LocUpdateData) GetLcsQosClass() LcsQosClass`

GetLcsQosClass returns the LcsQosClass field if non-nil, zero value otherwise.

### GetLcsQosClassOk

`func (o *LocUpdateData) GetLcsQosClassOk() (*LcsQosClass, bool)`

GetLcsQosClassOk returns a tuple with the LcsQosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsQosClass

`func (o *LocUpdateData) SetLcsQosClass(v LcsQosClass)`

SetLcsQosClass sets LcsQosClass field to given value.


### GetExternalClientIdentification

`func (o *LocUpdateData) GetExternalClientIdentification() string`

GetExternalClientIdentification returns the ExternalClientIdentification field if non-nil, zero value otherwise.

### GetExternalClientIdentificationOk

`func (o *LocUpdateData) GetExternalClientIdentificationOk() (*string, bool)`

GetExternalClientIdentificationOk returns a tuple with the ExternalClientIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClientIdentification

`func (o *LocUpdateData) SetExternalClientIdentification(v string)`

SetExternalClientIdentification sets ExternalClientIdentification field to given value.

### HasExternalClientIdentification

`func (o *LocUpdateData) HasExternalClientIdentification() bool`

HasExternalClientIdentification returns a boolean if a field has been set.

### GetAfId

`func (o *LocUpdateData) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *LocUpdateData) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *LocUpdateData) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *LocUpdateData) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetGmlcNumber

`func (o *LocUpdateData) GetGmlcNumber() string`

GetGmlcNumber returns the GmlcNumber field if non-nil, zero value otherwise.

### GetGmlcNumberOk

`func (o *LocUpdateData) GetGmlcNumberOk() (*string, bool)`

GetGmlcNumberOk returns a tuple with the GmlcNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmlcNumber

`func (o *LocUpdateData) SetGmlcNumber(v string)`

SetGmlcNumber sets GmlcNumber field to given value.

### HasGmlcNumber

`func (o *LocUpdateData) HasGmlcNumber() bool`

HasGmlcNumber returns a boolean if a field has been set.

### GetLcsServiceType

`func (o *LocUpdateData) GetLcsServiceType() int32`

GetLcsServiceType returns the LcsServiceType field if non-nil, zero value otherwise.

### GetLcsServiceTypeOk

`func (o *LocUpdateData) GetLcsServiceTypeOk() (*int32, bool)`

GetLcsServiceTypeOk returns a tuple with the LcsServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsServiceType

`func (o *LocUpdateData) SetLcsServiceType(v int32)`

SetLcsServiceType sets LcsServiceType field to given value.

### HasLcsServiceType

`func (o *LocUpdateData) HasLcsServiceType() bool`

HasLcsServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


