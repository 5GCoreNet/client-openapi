# SubscrCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NfInstanceIdList** | **[]string** |  | 
**NfType** | **string** |  | 
**ServiceName** | [**ServiceName**](ServiceName.md) |  | 
**ConditionType** | **string** |  | 
**ServiceNameList** | [**[]ServiceName**](ServiceName.md) |  | 
**GuamiList** | [**[]Guami**](Guami.md) |  | 
**SnssaiList** | [**[]Snssai**](Snssai.md) |  | 
**NsiList** | Pointer to **[]string** |  | [optional] 
**NfGroupId** | **string** | Identifier of a group of NFs. | 
**NfGroupIdList** | **[]string** |  | 
**NfSetId** | **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | 
**NfServiceSetId** | **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | 
**SmfServingArea** | Pointer to **[]string** |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**ScpDomains** | **[]string** |  | 
**NfTypeList** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**AnalyticsIds** | Pointer to **[]string** |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**ServingNfTypeList** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**ServingNfSetIdList** | Pointer to **[]string** |  | [optional] 
**MlAnalyticsList** | Pointer to [**[]MlAnalyticsInfo**](MlAnalyticsInfo.md) |  | [optional] 
**AfEvents** | Pointer to [**[]AfEvent**](AfEvent.md) |  | [optional] 
**PfdData** | Pointer to [**PfdData**](PfdData.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ServedFqdnList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSubscrCond

`func NewSubscrCond(nfInstanceId string, nfInstanceIdList []string, nfType string, serviceName ServiceName, conditionType string, serviceNameList []ServiceName, guamiList []Guami, snssaiList []Snssai, nfGroupId string, nfGroupIdList []string, nfSetId string, nfServiceSetId string, scpDomains []string, ) *SubscrCond`

NewSubscrCond instantiates a new SubscrCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscrCondWithDefaults

`func NewSubscrCondWithDefaults() *SubscrCond`

NewSubscrCondWithDefaults instantiates a new SubscrCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *SubscrCond) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *SubscrCond) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *SubscrCond) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfInstanceIdList

`func (o *SubscrCond) GetNfInstanceIdList() []string`

GetNfInstanceIdList returns the NfInstanceIdList field if non-nil, zero value otherwise.

### GetNfInstanceIdListOk

`func (o *SubscrCond) GetNfInstanceIdListOk() (*[]string, bool)`

GetNfInstanceIdListOk returns a tuple with the NfInstanceIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceIdList

`func (o *SubscrCond) SetNfInstanceIdList(v []string)`

SetNfInstanceIdList sets NfInstanceIdList field to given value.


### GetNfType

`func (o *SubscrCond) GetNfType() string`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *SubscrCond) GetNfTypeOk() (*string, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *SubscrCond) SetNfType(v string)`

SetNfType sets NfType field to given value.


### GetServiceName

`func (o *SubscrCond) GetServiceName() ServiceName`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SubscrCond) GetServiceNameOk() (*ServiceName, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SubscrCond) SetServiceName(v ServiceName)`

SetServiceName sets ServiceName field to given value.


### GetConditionType

`func (o *SubscrCond) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *SubscrCond) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *SubscrCond) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetServiceNameList

`func (o *SubscrCond) GetServiceNameList() []ServiceName`

GetServiceNameList returns the ServiceNameList field if non-nil, zero value otherwise.

### GetServiceNameListOk

`func (o *SubscrCond) GetServiceNameListOk() (*[]ServiceName, bool)`

GetServiceNameListOk returns a tuple with the ServiceNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNameList

`func (o *SubscrCond) SetServiceNameList(v []ServiceName)`

SetServiceNameList sets ServiceNameList field to given value.


### GetGuamiList

`func (o *SubscrCond) GetGuamiList() []Guami`

GetGuamiList returns the GuamiList field if non-nil, zero value otherwise.

### GetGuamiListOk

`func (o *SubscrCond) GetGuamiListOk() (*[]Guami, bool)`

GetGuamiListOk returns a tuple with the GuamiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamiList

`func (o *SubscrCond) SetGuamiList(v []Guami)`

SetGuamiList sets GuamiList field to given value.


### GetSnssaiList

`func (o *SubscrCond) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *SubscrCond) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *SubscrCond) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.


### GetNsiList

`func (o *SubscrCond) GetNsiList() []string`

GetNsiList returns the NsiList field if non-nil, zero value otherwise.

### GetNsiListOk

`func (o *SubscrCond) GetNsiListOk() (*[]string, bool)`

GetNsiListOk returns a tuple with the NsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiList

`func (o *SubscrCond) SetNsiList(v []string)`

SetNsiList sets NsiList field to given value.

### HasNsiList

`func (o *SubscrCond) HasNsiList() bool`

HasNsiList returns a boolean if a field has been set.

### GetNfGroupId

`func (o *SubscrCond) GetNfGroupId() string`

GetNfGroupId returns the NfGroupId field if non-nil, zero value otherwise.

### GetNfGroupIdOk

`func (o *SubscrCond) GetNfGroupIdOk() (*string, bool)`

GetNfGroupIdOk returns a tuple with the NfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfGroupId

`func (o *SubscrCond) SetNfGroupId(v string)`

SetNfGroupId sets NfGroupId field to given value.


### GetNfGroupIdList

`func (o *SubscrCond) GetNfGroupIdList() []string`

GetNfGroupIdList returns the NfGroupIdList field if non-nil, zero value otherwise.

### GetNfGroupIdListOk

`func (o *SubscrCond) GetNfGroupIdListOk() (*[]string, bool)`

GetNfGroupIdListOk returns a tuple with the NfGroupIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfGroupIdList

`func (o *SubscrCond) SetNfGroupIdList(v []string)`

SetNfGroupIdList sets NfGroupIdList field to given value.


### GetNfSetId

`func (o *SubscrCond) GetNfSetId() string`

GetNfSetId returns the NfSetId field if non-nil, zero value otherwise.

### GetNfSetIdOk

`func (o *SubscrCond) GetNfSetIdOk() (*string, bool)`

GetNfSetIdOk returns a tuple with the NfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetId

`func (o *SubscrCond) SetNfSetId(v string)`

SetNfSetId sets NfSetId field to given value.


### GetNfServiceSetId

`func (o *SubscrCond) GetNfServiceSetId() string`

GetNfServiceSetId returns the NfServiceSetId field if non-nil, zero value otherwise.

### GetNfServiceSetIdOk

`func (o *SubscrCond) GetNfServiceSetIdOk() (*string, bool)`

GetNfServiceSetIdOk returns a tuple with the NfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceSetId

`func (o *SubscrCond) SetNfServiceSetId(v string)`

SetNfServiceSetId sets NfServiceSetId field to given value.


### GetSmfServingArea

`func (o *SubscrCond) GetSmfServingArea() []string`

GetSmfServingArea returns the SmfServingArea field if non-nil, zero value otherwise.

### GetSmfServingAreaOk

`func (o *SubscrCond) GetSmfServingAreaOk() (*[]string, bool)`

GetSmfServingAreaOk returns a tuple with the SmfServingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServingArea

`func (o *SubscrCond) SetSmfServingArea(v []string)`

SetSmfServingArea sets SmfServingArea field to given value.

### HasSmfServingArea

`func (o *SubscrCond) HasSmfServingArea() bool`

HasSmfServingArea returns a boolean if a field has been set.

### GetTaiList

`func (o *SubscrCond) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *SubscrCond) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *SubscrCond) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *SubscrCond) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetScpDomains

`func (o *SubscrCond) GetScpDomains() []string`

GetScpDomains returns the ScpDomains field if non-nil, zero value otherwise.

### GetScpDomainsOk

`func (o *SubscrCond) GetScpDomainsOk() (*[]string, bool)`

GetScpDomainsOk returns a tuple with the ScpDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpDomains

`func (o *SubscrCond) SetScpDomains(v []string)`

SetScpDomains sets ScpDomains field to given value.


### GetNfTypeList

`func (o *SubscrCond) GetNfTypeList() []NFType`

GetNfTypeList returns the NfTypeList field if non-nil, zero value otherwise.

### GetNfTypeListOk

`func (o *SubscrCond) GetNfTypeListOk() (*[]NFType, bool)`

GetNfTypeListOk returns a tuple with the NfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypeList

`func (o *SubscrCond) SetNfTypeList(v []NFType)`

SetNfTypeList sets NfTypeList field to given value.

### HasNfTypeList

`func (o *SubscrCond) HasNfTypeList() bool`

HasNfTypeList returns a boolean if a field has been set.

### GetAnalyticsIds

`func (o *SubscrCond) GetAnalyticsIds() []string`

GetAnalyticsIds returns the AnalyticsIds field if non-nil, zero value otherwise.

### GetAnalyticsIdsOk

`func (o *SubscrCond) GetAnalyticsIdsOk() (*[]string, bool)`

GetAnalyticsIdsOk returns a tuple with the AnalyticsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsIds

`func (o *SubscrCond) SetAnalyticsIds(v []string)`

SetAnalyticsIds sets AnalyticsIds field to given value.

### HasAnalyticsIds

`func (o *SubscrCond) HasAnalyticsIds() bool`

HasAnalyticsIds returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *SubscrCond) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *SubscrCond) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *SubscrCond) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *SubscrCond) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetServingNfTypeList

`func (o *SubscrCond) GetServingNfTypeList() []NFType`

GetServingNfTypeList returns the ServingNfTypeList field if non-nil, zero value otherwise.

### GetServingNfTypeListOk

`func (o *SubscrCond) GetServingNfTypeListOk() (*[]NFType, bool)`

GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfTypeList

`func (o *SubscrCond) SetServingNfTypeList(v []NFType)`

SetServingNfTypeList sets ServingNfTypeList field to given value.

### HasServingNfTypeList

`func (o *SubscrCond) HasServingNfTypeList() bool`

HasServingNfTypeList returns a boolean if a field has been set.

### GetServingNfSetIdList

`func (o *SubscrCond) GetServingNfSetIdList() []string`

GetServingNfSetIdList returns the ServingNfSetIdList field if non-nil, zero value otherwise.

### GetServingNfSetIdListOk

`func (o *SubscrCond) GetServingNfSetIdListOk() (*[]string, bool)`

GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfSetIdList

`func (o *SubscrCond) SetServingNfSetIdList(v []string)`

SetServingNfSetIdList sets ServingNfSetIdList field to given value.

### HasServingNfSetIdList

`func (o *SubscrCond) HasServingNfSetIdList() bool`

HasServingNfSetIdList returns a boolean if a field has been set.

### GetMlAnalyticsList

`func (o *SubscrCond) GetMlAnalyticsList() []MlAnalyticsInfo`

GetMlAnalyticsList returns the MlAnalyticsList field if non-nil, zero value otherwise.

### GetMlAnalyticsListOk

`func (o *SubscrCond) GetMlAnalyticsListOk() (*[]MlAnalyticsInfo, bool)`

GetMlAnalyticsListOk returns a tuple with the MlAnalyticsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlAnalyticsList

`func (o *SubscrCond) SetMlAnalyticsList(v []MlAnalyticsInfo)`

SetMlAnalyticsList sets MlAnalyticsList field to given value.

### HasMlAnalyticsList

`func (o *SubscrCond) HasMlAnalyticsList() bool`

HasMlAnalyticsList returns a boolean if a field has been set.

### GetAfEvents

`func (o *SubscrCond) GetAfEvents() []AfEvent`

GetAfEvents returns the AfEvents field if non-nil, zero value otherwise.

### GetAfEventsOk

`func (o *SubscrCond) GetAfEventsOk() (*[]AfEvent, bool)`

GetAfEventsOk returns a tuple with the AfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEvents

`func (o *SubscrCond) SetAfEvents(v []AfEvent)`

SetAfEvents sets AfEvents field to given value.

### HasAfEvents

`func (o *SubscrCond) HasAfEvents() bool`

HasAfEvents returns a boolean if a field has been set.

### GetPfdData

`func (o *SubscrCond) GetPfdData() PfdData`

GetPfdData returns the PfdData field if non-nil, zero value otherwise.

### GetPfdDataOk

`func (o *SubscrCond) GetPfdDataOk() (*PfdData, bool)`

GetPfdDataOk returns a tuple with the PfdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdData

`func (o *SubscrCond) SetPfdData(v PfdData)`

SetPfdData sets PfdData field to given value.

### HasPfdData

`func (o *SubscrCond) HasPfdData() bool`

HasPfdData returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *SubscrCond) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *SubscrCond) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *SubscrCond) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *SubscrCond) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *SubscrCond) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *SubscrCond) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *SubscrCond) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *SubscrCond) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetServedFqdnList

`func (o *SubscrCond) GetServedFqdnList() []string`

GetServedFqdnList returns the ServedFqdnList field if non-nil, zero value otherwise.

### GetServedFqdnListOk

`func (o *SubscrCond) GetServedFqdnListOk() (*[]string, bool)`

GetServedFqdnListOk returns a tuple with the ServedFqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedFqdnList

`func (o *SubscrCond) SetServedFqdnList(v []string)`

SetServedFqdnList sets ServedFqdnList field to given value.

### HasServedFqdnList

`func (o *SubscrCond) HasServedFqdnList() bool`

HasServedFqdnList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


