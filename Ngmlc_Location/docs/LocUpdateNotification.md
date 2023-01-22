# LocUpdateNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**LocationRequestType** | [**LocationRequestType**](LocationRequestType.md) |  | 
**LocationEstimate** | [**GeographicArea**](GeographicArea.md) |  | 
**AgeOfLocationEstimate** | **int32** | Indicates value of the age of the location estimate. | 
**TimestampOfLocationEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AccuracyFulfilmentIndicator** | [**AccuracyFulfilmentIndicator**](AccuracyFulfilmentIndicator.md) |  | 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**LcsQosClass** | [**LcsQosClass**](LcsQosClass.md) |  | 
**AfId** | Pointer to **string** |  | [optional] 
**ServiceIdentity** | Pointer to **string** | Contains the service identity | [optional] 

## Methods

### NewLocUpdateNotification

`func NewLocUpdateNotification(locationRequestType LocationRequestType, locationEstimate GeographicArea, ageOfLocationEstimate int32, accuracyFulfilmentIndicator AccuracyFulfilmentIndicator, lcsQosClass LcsQosClass, ) *LocUpdateNotification`

NewLocUpdateNotification instantiates a new LocUpdateNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocUpdateNotificationWithDefaults

`func NewLocUpdateNotificationWithDefaults() *LocUpdateNotification`

NewLocUpdateNotificationWithDefaults instantiates a new LocUpdateNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *LocUpdateNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *LocUpdateNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *LocUpdateNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *LocUpdateNotification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *LocUpdateNotification) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *LocUpdateNotification) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *LocUpdateNotification) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *LocUpdateNotification) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetLocationRequestType

`func (o *LocUpdateNotification) GetLocationRequestType() LocationRequestType`

GetLocationRequestType returns the LocationRequestType field if non-nil, zero value otherwise.

### GetLocationRequestTypeOk

`func (o *LocUpdateNotification) GetLocationRequestTypeOk() (*LocationRequestType, bool)`

GetLocationRequestTypeOk returns a tuple with the LocationRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationRequestType

`func (o *LocUpdateNotification) SetLocationRequestType(v LocationRequestType)`

SetLocationRequestType sets LocationRequestType field to given value.


### GetLocationEstimate

`func (o *LocUpdateNotification) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *LocUpdateNotification) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *LocUpdateNotification) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.


### GetAgeOfLocationEstimate

`func (o *LocUpdateNotification) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *LocUpdateNotification) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *LocUpdateNotification) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.


### GetTimestampOfLocationEstimate

`func (o *LocUpdateNotification) GetTimestampOfLocationEstimate() time.Time`

GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field if non-nil, zero value otherwise.

### GetTimestampOfLocationEstimateOk

`func (o *LocUpdateNotification) GetTimestampOfLocationEstimateOk() (*time.Time, bool)`

GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfLocationEstimate

`func (o *LocUpdateNotification) SetTimestampOfLocationEstimate(v time.Time)`

SetTimestampOfLocationEstimate sets TimestampOfLocationEstimate field to given value.

### HasTimestampOfLocationEstimate

`func (o *LocUpdateNotification) HasTimestampOfLocationEstimate() bool`

HasTimestampOfLocationEstimate returns a boolean if a field has been set.

### GetAccuracyFulfilmentIndicator

`func (o *LocUpdateNotification) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator`

GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field if non-nil, zero value otherwise.

### GetAccuracyFulfilmentIndicatorOk

`func (o *LocUpdateNotification) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool)`

GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyFulfilmentIndicator

`func (o *LocUpdateNotification) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator)`

SetAccuracyFulfilmentIndicator sets AccuracyFulfilmentIndicator field to given value.


### GetCivicAddress

`func (o *LocUpdateNotification) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *LocUpdateNotification) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *LocUpdateNotification) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *LocUpdateNotification) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetLcsQosClass

`func (o *LocUpdateNotification) GetLcsQosClass() LcsQosClass`

GetLcsQosClass returns the LcsQosClass field if non-nil, zero value otherwise.

### GetLcsQosClassOk

`func (o *LocUpdateNotification) GetLcsQosClassOk() (*LcsQosClass, bool)`

GetLcsQosClassOk returns a tuple with the LcsQosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsQosClass

`func (o *LocUpdateNotification) SetLcsQosClass(v LcsQosClass)`

SetLcsQosClass sets LcsQosClass field to given value.


### GetAfId

`func (o *LocUpdateNotification) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *LocUpdateNotification) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *LocUpdateNotification) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *LocUpdateNotification) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetServiceIdentity

`func (o *LocUpdateNotification) GetServiceIdentity() string`

GetServiceIdentity returns the ServiceIdentity field if non-nil, zero value otherwise.

### GetServiceIdentityOk

`func (o *LocUpdateNotification) GetServiceIdentityOk() (*string, bool)`

GetServiceIdentityOk returns a tuple with the ServiceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIdentity

`func (o *LocUpdateNotification) SetServiceIdentity(v string)`

SetServiceIdentity sets ServiceIdentity field to given value.

### HasServiceIdentity

`func (o *LocUpdateNotification) HasServiceIdentity() bool`

HasServiceIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


