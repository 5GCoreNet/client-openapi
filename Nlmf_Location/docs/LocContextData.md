# LocContextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**LocationQoS** | Pointer to [**LocationQoS**](LocationQoS.md) |  | [optional] 
**SupportedGADShapes** | Pointer to [**[]SupportedGADShapes**](SupportedGADShapes.md) |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**LdrType** | [**LdrType**](LdrType.md) |  | 
**HgmlcCallBackURI** | **string** | String providing an URI formatted according to RFC 3986. | 
**LdrReference** | **string** | LDR Reference. | 
**PeriodicEventInfo** | Pointer to [**PeriodicEventInfo**](PeriodicEventInfo.md) |  | [optional] 
**AreaEventInfo** | Pointer to [**AreaEventInfo**](AreaEventInfo.md) |  | [optional] 
**MotionEventInfo** | Pointer to [**MotionEventInfo**](MotionEventInfo.md) |  | [optional] 
**EventReportMessage** | [**EventReportMessage**](EventReportMessage.md) |  | 
**EventReportingStatus** | Pointer to [**EventReportingStatus**](EventReportingStatus.md) |  | [optional] 
**UeLocationInfo** | Pointer to [**UELocationInfo**](UELocationInfo.md) |  | [optional] 
**CIoT5GSOptimisation** | Pointer to **bool** |  | [optional] [default to false]
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**UePositioningCap** | Pointer to **string** | Positioning capabilities supported by the UE. A string encoding the \&quot;ProvideCapabilities-r9-IEs\&quot; IE as specified in clause 6.3 of 3GPP TS 37.355 (start from octet 1). | [optional] 
**ScheduledLocTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewLocContextData

`func NewLocContextData(amfId string, ldrType LdrType, hgmlcCallBackURI string, ldrReference string, eventReportMessage EventReportMessage, ) *LocContextData`

NewLocContextData instantiates a new LocContextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocContextDataWithDefaults

`func NewLocContextDataWithDefaults() *LocContextData`

NewLocContextDataWithDefaults instantiates a new LocContextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfId

`func (o *LocContextData) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *LocContextData) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *LocContextData) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.


### GetLocationQoS

`func (o *LocContextData) GetLocationQoS() LocationQoS`

GetLocationQoS returns the LocationQoS field if non-nil, zero value otherwise.

### GetLocationQoSOk

`func (o *LocContextData) GetLocationQoSOk() (*LocationQoS, bool)`

GetLocationQoSOk returns a tuple with the LocationQoS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationQoS

`func (o *LocContextData) SetLocationQoS(v LocationQoS)`

SetLocationQoS sets LocationQoS field to given value.

### HasLocationQoS

`func (o *LocContextData) HasLocationQoS() bool`

HasLocationQoS returns a boolean if a field has been set.

### GetSupportedGADShapes

`func (o *LocContextData) GetSupportedGADShapes() []SupportedGADShapes`

GetSupportedGADShapes returns the SupportedGADShapes field if non-nil, zero value otherwise.

### GetSupportedGADShapesOk

`func (o *LocContextData) GetSupportedGADShapesOk() (*[]SupportedGADShapes, bool)`

GetSupportedGADShapesOk returns a tuple with the SupportedGADShapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedGADShapes

`func (o *LocContextData) SetSupportedGADShapes(v []SupportedGADShapes)`

SetSupportedGADShapes sets SupportedGADShapes field to given value.

### HasSupportedGADShapes

`func (o *LocContextData) HasSupportedGADShapes() bool`

HasSupportedGADShapes returns a boolean if a field has been set.

### GetSupi

`func (o *LocContextData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *LocContextData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *LocContextData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *LocContextData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *LocContextData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *LocContextData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *LocContextData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *LocContextData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetLdrType

`func (o *LocContextData) GetLdrType() LdrType`

GetLdrType returns the LdrType field if non-nil, zero value otherwise.

### GetLdrTypeOk

`func (o *LocContextData) GetLdrTypeOk() (*LdrType, bool)`

GetLdrTypeOk returns a tuple with the LdrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrType

`func (o *LocContextData) SetLdrType(v LdrType)`

SetLdrType sets LdrType field to given value.


### GetHgmlcCallBackURI

`func (o *LocContextData) GetHgmlcCallBackURI() string`

GetHgmlcCallBackURI returns the HgmlcCallBackURI field if non-nil, zero value otherwise.

### GetHgmlcCallBackURIOk

`func (o *LocContextData) GetHgmlcCallBackURIOk() (*string, bool)`

GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgmlcCallBackURI

`func (o *LocContextData) SetHgmlcCallBackURI(v string)`

SetHgmlcCallBackURI sets HgmlcCallBackURI field to given value.


### GetLdrReference

`func (o *LocContextData) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *LocContextData) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *LocContextData) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.


### GetPeriodicEventInfo

`func (o *LocContextData) GetPeriodicEventInfo() PeriodicEventInfo`

GetPeriodicEventInfo returns the PeriodicEventInfo field if non-nil, zero value otherwise.

### GetPeriodicEventInfoOk

`func (o *LocContextData) GetPeriodicEventInfoOk() (*PeriodicEventInfo, bool)`

GetPeriodicEventInfoOk returns a tuple with the PeriodicEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicEventInfo

`func (o *LocContextData) SetPeriodicEventInfo(v PeriodicEventInfo)`

SetPeriodicEventInfo sets PeriodicEventInfo field to given value.

### HasPeriodicEventInfo

`func (o *LocContextData) HasPeriodicEventInfo() bool`

HasPeriodicEventInfo returns a boolean if a field has been set.

### GetAreaEventInfo

`func (o *LocContextData) GetAreaEventInfo() AreaEventInfo`

GetAreaEventInfo returns the AreaEventInfo field if non-nil, zero value otherwise.

### GetAreaEventInfoOk

`func (o *LocContextData) GetAreaEventInfoOk() (*AreaEventInfo, bool)`

GetAreaEventInfoOk returns a tuple with the AreaEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaEventInfo

`func (o *LocContextData) SetAreaEventInfo(v AreaEventInfo)`

SetAreaEventInfo sets AreaEventInfo field to given value.

### HasAreaEventInfo

`func (o *LocContextData) HasAreaEventInfo() bool`

HasAreaEventInfo returns a boolean if a field has been set.

### GetMotionEventInfo

`func (o *LocContextData) GetMotionEventInfo() MotionEventInfo`

GetMotionEventInfo returns the MotionEventInfo field if non-nil, zero value otherwise.

### GetMotionEventInfoOk

`func (o *LocContextData) GetMotionEventInfoOk() (*MotionEventInfo, bool)`

GetMotionEventInfoOk returns a tuple with the MotionEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionEventInfo

`func (o *LocContextData) SetMotionEventInfo(v MotionEventInfo)`

SetMotionEventInfo sets MotionEventInfo field to given value.

### HasMotionEventInfo

`func (o *LocContextData) HasMotionEventInfo() bool`

HasMotionEventInfo returns a boolean if a field has been set.

### GetEventReportMessage

`func (o *LocContextData) GetEventReportMessage() EventReportMessage`

GetEventReportMessage returns the EventReportMessage field if non-nil, zero value otherwise.

### GetEventReportMessageOk

`func (o *LocContextData) GetEventReportMessageOk() (*EventReportMessage, bool)`

GetEventReportMessageOk returns a tuple with the EventReportMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReportMessage

`func (o *LocContextData) SetEventReportMessage(v EventReportMessage)`

SetEventReportMessage sets EventReportMessage field to given value.


### GetEventReportingStatus

`func (o *LocContextData) GetEventReportingStatus() EventReportingStatus`

GetEventReportingStatus returns the EventReportingStatus field if non-nil, zero value otherwise.

### GetEventReportingStatusOk

`func (o *LocContextData) GetEventReportingStatusOk() (*EventReportingStatus, bool)`

GetEventReportingStatusOk returns a tuple with the EventReportingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReportingStatus

`func (o *LocContextData) SetEventReportingStatus(v EventReportingStatus)`

SetEventReportingStatus sets EventReportingStatus field to given value.

### HasEventReportingStatus

`func (o *LocContextData) HasEventReportingStatus() bool`

HasEventReportingStatus returns a boolean if a field has been set.

### GetUeLocationInfo

`func (o *LocContextData) GetUeLocationInfo() UELocationInfo`

GetUeLocationInfo returns the UeLocationInfo field if non-nil, zero value otherwise.

### GetUeLocationInfoOk

`func (o *LocContextData) GetUeLocationInfoOk() (*UELocationInfo, bool)`

GetUeLocationInfoOk returns a tuple with the UeLocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocationInfo

`func (o *LocContextData) SetUeLocationInfo(v UELocationInfo)`

SetUeLocationInfo sets UeLocationInfo field to given value.

### HasUeLocationInfo

`func (o *LocContextData) HasUeLocationInfo() bool`

HasUeLocationInfo returns a boolean if a field has been set.

### GetCIoT5GSOptimisation

`func (o *LocContextData) GetCIoT5GSOptimisation() bool`

GetCIoT5GSOptimisation returns the CIoT5GSOptimisation field if non-nil, zero value otherwise.

### GetCIoT5GSOptimisationOk

`func (o *LocContextData) GetCIoT5GSOptimisationOk() (*bool, bool)`

GetCIoT5GSOptimisationOk returns a tuple with the CIoT5GSOptimisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIoT5GSOptimisation

`func (o *LocContextData) SetCIoT5GSOptimisation(v bool)`

SetCIoT5GSOptimisation sets CIoT5GSOptimisation field to given value.

### HasCIoT5GSOptimisation

`func (o *LocContextData) HasCIoT5GSOptimisation() bool`

HasCIoT5GSOptimisation returns a boolean if a field has been set.

### GetEcgi

`func (o *LocContextData) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *LocContextData) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *LocContextData) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *LocContextData) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetNcgi

`func (o *LocContextData) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *LocContextData) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *LocContextData) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *LocContextData) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

### GetGuami

`func (o *LocContextData) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *LocContextData) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *LocContextData) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *LocContextData) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *LocContextData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *LocContextData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *LocContextData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *LocContextData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetUePositioningCap

`func (o *LocContextData) GetUePositioningCap() string`

GetUePositioningCap returns the UePositioningCap field if non-nil, zero value otherwise.

### GetUePositioningCapOk

`func (o *LocContextData) GetUePositioningCapOk() (*string, bool)`

GetUePositioningCapOk returns a tuple with the UePositioningCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePositioningCap

`func (o *LocContextData) SetUePositioningCap(v string)`

SetUePositioningCap sets UePositioningCap field to given value.

### HasUePositioningCap

`func (o *LocContextData) HasUePositioningCap() bool`

HasUePositioningCap returns a boolean if a field has been set.

### GetScheduledLocTime

`func (o *LocContextData) GetScheduledLocTime() time.Time`

GetScheduledLocTime returns the ScheduledLocTime field if non-nil, zero value otherwise.

### GetScheduledLocTimeOk

`func (o *LocContextData) GetScheduledLocTimeOk() (*time.Time, bool)`

GetScheduledLocTimeOk returns a tuple with the ScheduledLocTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledLocTime

`func (o *LocContextData) SetScheduledLocTime(v time.Time)`

SetScheduledLocTime sets ScheduledLocTime field to given value.

### HasScheduledLocTime

`func (o *LocContextData) HasScheduledLocTime() bool`

HasScheduledLocTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


