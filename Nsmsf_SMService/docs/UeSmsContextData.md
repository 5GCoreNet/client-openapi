# UeSmsContextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**AmfId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**Guamis** | Pointer to [**[]Guami**](Guami.md) |  | [optional] 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**AdditionalAccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**UdmGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**RoutingIndicator** | Pointer to **string** |  | [optional] 
**HNwPubKeyId** | Pointer to **int32** |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**AdditionalRatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewUeSmsContextData

`func NewUeSmsContextData(supi string, amfId string, accessType AccessType, ) *UeSmsContextData`

NewUeSmsContextData instantiates a new UeSmsContextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeSmsContextDataWithDefaults

`func NewUeSmsContextDataWithDefaults() *UeSmsContextData`

NewUeSmsContextDataWithDefaults instantiates a new UeSmsContextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeSmsContextData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeSmsContextData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeSmsContextData) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetPei

`func (o *UeSmsContextData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *UeSmsContextData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *UeSmsContextData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *UeSmsContextData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetAmfId

`func (o *UeSmsContextData) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *UeSmsContextData) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *UeSmsContextData) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.


### GetGuamis

`func (o *UeSmsContextData) GetGuamis() []Guami`

GetGuamis returns the Guamis field if non-nil, zero value otherwise.

### GetGuamisOk

`func (o *UeSmsContextData) GetGuamisOk() (*[]Guami, bool)`

GetGuamisOk returns a tuple with the Guamis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamis

`func (o *UeSmsContextData) SetGuamis(v []Guami)`

SetGuamis sets Guamis field to given value.

### HasGuamis

`func (o *UeSmsContextData) HasGuamis() bool`

HasGuamis returns a boolean if a field has been set.

### GetAccessType

`func (o *UeSmsContextData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UeSmsContextData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UeSmsContextData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetAdditionalAccessType

`func (o *UeSmsContextData) GetAdditionalAccessType() AccessType`

GetAdditionalAccessType returns the AdditionalAccessType field if non-nil, zero value otherwise.

### GetAdditionalAccessTypeOk

`func (o *UeSmsContextData) GetAdditionalAccessTypeOk() (*AccessType, bool)`

GetAdditionalAccessTypeOk returns a tuple with the AdditionalAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAccessType

`func (o *UeSmsContextData) SetAdditionalAccessType(v AccessType)`

SetAdditionalAccessType sets AdditionalAccessType field to given value.

### HasAdditionalAccessType

`func (o *UeSmsContextData) HasAdditionalAccessType() bool`

HasAdditionalAccessType returns a boolean if a field has been set.

### GetGpsi

`func (o *UeSmsContextData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UeSmsContextData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UeSmsContextData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *UeSmsContextData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetUeLocation

`func (o *UeSmsContextData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *UeSmsContextData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *UeSmsContextData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *UeSmsContextData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *UeSmsContextData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *UeSmsContextData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *UeSmsContextData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *UeSmsContextData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetTraceData

`func (o *UeSmsContextData) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *UeSmsContextData) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *UeSmsContextData) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *UeSmsContextData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *UeSmsContextData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *UeSmsContextData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetBackupAmfInfo

`func (o *UeSmsContextData) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *UeSmsContextData) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *UeSmsContextData) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *UeSmsContextData) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### GetUdmGroupId

`func (o *UeSmsContextData) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *UeSmsContextData) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *UeSmsContextData) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *UeSmsContextData) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.

### GetRoutingIndicator

`func (o *UeSmsContextData) GetRoutingIndicator() string`

GetRoutingIndicator returns the RoutingIndicator field if non-nil, zero value otherwise.

### GetRoutingIndicatorOk

`func (o *UeSmsContextData) GetRoutingIndicatorOk() (*string, bool)`

GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicator

`func (o *UeSmsContextData) SetRoutingIndicator(v string)`

SetRoutingIndicator sets RoutingIndicator field to given value.

### HasRoutingIndicator

`func (o *UeSmsContextData) HasRoutingIndicator() bool`

HasRoutingIndicator returns a boolean if a field has been set.

### GetHNwPubKeyId

`func (o *UeSmsContextData) GetHNwPubKeyId() int32`

GetHNwPubKeyId returns the HNwPubKeyId field if non-nil, zero value otherwise.

### GetHNwPubKeyIdOk

`func (o *UeSmsContextData) GetHNwPubKeyIdOk() (*int32, bool)`

GetHNwPubKeyIdOk returns a tuple with the HNwPubKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHNwPubKeyId

`func (o *UeSmsContextData) SetHNwPubKeyId(v int32)`

SetHNwPubKeyId sets HNwPubKeyId field to given value.

### HasHNwPubKeyId

`func (o *UeSmsContextData) HasHNwPubKeyId() bool`

HasHNwPubKeyId returns a boolean if a field has been set.

### GetRatType

`func (o *UeSmsContextData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *UeSmsContextData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *UeSmsContextData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *UeSmsContextData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetAdditionalRatType

`func (o *UeSmsContextData) GetAdditionalRatType() RatType`

GetAdditionalRatType returns the AdditionalRatType field if non-nil, zero value otherwise.

### GetAdditionalRatTypeOk

`func (o *UeSmsContextData) GetAdditionalRatTypeOk() (*RatType, bool)`

GetAdditionalRatTypeOk returns a tuple with the AdditionalRatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRatType

`func (o *UeSmsContextData) SetAdditionalRatType(v RatType)`

SetAdditionalRatType sets AdditionalRatType field to given value.

### HasAdditionalRatType

`func (o *UeSmsContextData) HasAdditionalRatType() bool`

HasAdditionalRatType returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeSmsContextData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeSmsContextData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeSmsContextData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeSmsContextData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


