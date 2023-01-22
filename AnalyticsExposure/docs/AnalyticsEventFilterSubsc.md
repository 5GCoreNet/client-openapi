# AnalyticsEventFilterSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NwPerfReqs** | Pointer to [**[]NetworkPerfRequirement**](NetworkPerfRequirement.md) |  | [optional] 
**LocArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Dnais** | Pointer to **[]string** |  | [optional] 
**ExcepRequs** | Pointer to [**[]Exception**](Exception.md) |  | [optional] 
**ExptAnaType** | Pointer to [**ExpectedAnalyticsType**](ExpectedAnalyticsType.md) |  | [optional] 
**ExptUeBehav** | Pointer to [**ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) |  | [optional] 
**MatchingDir** | Pointer to [**MatchingDirection**](MatchingDirection.md) |  | [optional] 
**ReptThlds** | Pointer to [**[]ThresholdLevel**](ThresholdLevel.md) |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**NsiIdInfos** | Pointer to [**[]NsiIdInfo**](NsiIdInfo.md) |  | [optional] 
**QosReq** | Pointer to [**QosRequirement**](QosRequirement.md) |  | [optional] 
**QosFlowRetThds** | Pointer to [**[]RetainabilityThreshold**](RetainabilityThreshold.md) |  | [optional] 
**RanUeThrouThds** | Pointer to **[]string** |  | [optional] 
**DisperReqs** | Pointer to [**[]DispersionRequirement**](DispersionRequirement.md) |  | [optional] 
**ListOfAnaSubsets** | Pointer to [**[]AnalyticsSubset**](AnalyticsSubset.md) |  | [optional] 
**DnPerfReqs** | Pointer to [**[]DnPerformanceReq**](DnPerformanceReq.md) |  | [optional] 
**BwRequs** | Pointer to [**[]BwRequirement**](BwRequirement.md) |  | [optional] 
**RatFreqs** | Pointer to [**[]RatFreqInformation**](RatFreqInformation.md) |  | [optional] 
**AppServerAddrs** | Pointer to [**[]AddrFqdn**](AddrFqdn.md) |  | [optional] 
**ExtraReportReq** | Pointer to [**EventReportingRequirement**](EventReportingRequirement.md) |  | [optional] 
**MaxNumOfTopAppUl** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MaxNumOfTopAppDl** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**VisitedLocAreas** | Pointer to [**[]LocationArea5G**](LocationArea5G.md) |  | [optional] 

## Methods

### NewAnalyticsEventFilterSubsc

`func NewAnalyticsEventFilterSubsc() *AnalyticsEventFilterSubsc`

NewAnalyticsEventFilterSubsc instantiates a new AnalyticsEventFilterSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsEventFilterSubscWithDefaults

`func NewAnalyticsEventFilterSubscWithDefaults() *AnalyticsEventFilterSubsc`

NewAnalyticsEventFilterSubscWithDefaults instantiates a new AnalyticsEventFilterSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNwPerfReqs

`func (o *AnalyticsEventFilterSubsc) GetNwPerfReqs() []NetworkPerfRequirement`

GetNwPerfReqs returns the NwPerfReqs field if non-nil, zero value otherwise.

### GetNwPerfReqsOk

`func (o *AnalyticsEventFilterSubsc) GetNwPerfReqsOk() (*[]NetworkPerfRequirement, bool)`

GetNwPerfReqsOk returns a tuple with the NwPerfReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfReqs

`func (o *AnalyticsEventFilterSubsc) SetNwPerfReqs(v []NetworkPerfRequirement)`

SetNwPerfReqs sets NwPerfReqs field to given value.

### HasNwPerfReqs

`func (o *AnalyticsEventFilterSubsc) HasNwPerfReqs() bool`

HasNwPerfReqs returns a boolean if a field has been set.

### GetLocArea

`func (o *AnalyticsEventFilterSubsc) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *AnalyticsEventFilterSubsc) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *AnalyticsEventFilterSubsc) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *AnalyticsEventFilterSubsc) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetAppIds

`func (o *AnalyticsEventFilterSubsc) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AnalyticsEventFilterSubsc) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AnalyticsEventFilterSubsc) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *AnalyticsEventFilterSubsc) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetDnn

`func (o *AnalyticsEventFilterSubsc) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *AnalyticsEventFilterSubsc) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *AnalyticsEventFilterSubsc) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *AnalyticsEventFilterSubsc) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetDnais

`func (o *AnalyticsEventFilterSubsc) GetDnais() []string`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *AnalyticsEventFilterSubsc) GetDnaisOk() (*[]string, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *AnalyticsEventFilterSubsc) SetDnais(v []string)`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *AnalyticsEventFilterSubsc) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetExcepRequs

`func (o *AnalyticsEventFilterSubsc) GetExcepRequs() []Exception`

GetExcepRequs returns the ExcepRequs field if non-nil, zero value otherwise.

### GetExcepRequsOk

`func (o *AnalyticsEventFilterSubsc) GetExcepRequsOk() (*[]Exception, bool)`

GetExcepRequsOk returns a tuple with the ExcepRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepRequs

`func (o *AnalyticsEventFilterSubsc) SetExcepRequs(v []Exception)`

SetExcepRequs sets ExcepRequs field to given value.

### HasExcepRequs

`func (o *AnalyticsEventFilterSubsc) HasExcepRequs() bool`

HasExcepRequs returns a boolean if a field has been set.

### GetExptAnaType

`func (o *AnalyticsEventFilterSubsc) GetExptAnaType() ExpectedAnalyticsType`

GetExptAnaType returns the ExptAnaType field if non-nil, zero value otherwise.

### GetExptAnaTypeOk

`func (o *AnalyticsEventFilterSubsc) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool)`

GetExptAnaTypeOk returns a tuple with the ExptAnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptAnaType

`func (o *AnalyticsEventFilterSubsc) SetExptAnaType(v ExpectedAnalyticsType)`

SetExptAnaType sets ExptAnaType field to given value.

### HasExptAnaType

`func (o *AnalyticsEventFilterSubsc) HasExptAnaType() bool`

HasExptAnaType returns a boolean if a field has been set.

### GetExptUeBehav

`func (o *AnalyticsEventFilterSubsc) GetExptUeBehav() ExpectedUeBehaviourData`

GetExptUeBehav returns the ExptUeBehav field if non-nil, zero value otherwise.

### GetExptUeBehavOk

`func (o *AnalyticsEventFilterSubsc) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool)`

GetExptUeBehavOk returns a tuple with the ExptUeBehav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptUeBehav

`func (o *AnalyticsEventFilterSubsc) SetExptUeBehav(v ExpectedUeBehaviourData)`

SetExptUeBehav sets ExptUeBehav field to given value.

### HasExptUeBehav

`func (o *AnalyticsEventFilterSubsc) HasExptUeBehav() bool`

HasExptUeBehav returns a boolean if a field has been set.

### GetMatchingDir

`func (o *AnalyticsEventFilterSubsc) GetMatchingDir() MatchingDirection`

GetMatchingDir returns the MatchingDir field if non-nil, zero value otherwise.

### GetMatchingDirOk

`func (o *AnalyticsEventFilterSubsc) GetMatchingDirOk() (*MatchingDirection, bool)`

GetMatchingDirOk returns a tuple with the MatchingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingDir

`func (o *AnalyticsEventFilterSubsc) SetMatchingDir(v MatchingDirection)`

SetMatchingDir sets MatchingDir field to given value.

### HasMatchingDir

`func (o *AnalyticsEventFilterSubsc) HasMatchingDir() bool`

HasMatchingDir returns a boolean if a field has been set.

### GetReptThlds

`func (o *AnalyticsEventFilterSubsc) GetReptThlds() []ThresholdLevel`

GetReptThlds returns the ReptThlds field if non-nil, zero value otherwise.

### GetReptThldsOk

`func (o *AnalyticsEventFilterSubsc) GetReptThldsOk() (*[]ThresholdLevel, bool)`

GetReptThldsOk returns a tuple with the ReptThlds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptThlds

`func (o *AnalyticsEventFilterSubsc) SetReptThlds(v []ThresholdLevel)`

SetReptThlds sets ReptThlds field to given value.

### HasReptThlds

`func (o *AnalyticsEventFilterSubsc) HasReptThlds() bool`

HasReptThlds returns a boolean if a field has been set.

### GetSnssai

`func (o *AnalyticsEventFilterSubsc) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AnalyticsEventFilterSubsc) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AnalyticsEventFilterSubsc) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *AnalyticsEventFilterSubsc) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetNsiIdInfos

`func (o *AnalyticsEventFilterSubsc) GetNsiIdInfos() []NsiIdInfo`

GetNsiIdInfos returns the NsiIdInfos field if non-nil, zero value otherwise.

### GetNsiIdInfosOk

`func (o *AnalyticsEventFilterSubsc) GetNsiIdInfosOk() (*[]NsiIdInfo, bool)`

GetNsiIdInfosOk returns a tuple with the NsiIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiIdInfos

`func (o *AnalyticsEventFilterSubsc) SetNsiIdInfos(v []NsiIdInfo)`

SetNsiIdInfos sets NsiIdInfos field to given value.

### HasNsiIdInfos

`func (o *AnalyticsEventFilterSubsc) HasNsiIdInfos() bool`

HasNsiIdInfos returns a boolean if a field has been set.

### GetQosReq

`func (o *AnalyticsEventFilterSubsc) GetQosReq() QosRequirement`

GetQosReq returns the QosReq field if non-nil, zero value otherwise.

### GetQosReqOk

`func (o *AnalyticsEventFilterSubsc) GetQosReqOk() (*QosRequirement, bool)`

GetQosReqOk returns a tuple with the QosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReq

`func (o *AnalyticsEventFilterSubsc) SetQosReq(v QosRequirement)`

SetQosReq sets QosReq field to given value.

### HasQosReq

`func (o *AnalyticsEventFilterSubsc) HasQosReq() bool`

HasQosReq returns a boolean if a field has been set.

### GetQosFlowRetThds

`func (o *AnalyticsEventFilterSubsc) GetQosFlowRetThds() []RetainabilityThreshold`

GetQosFlowRetThds returns the QosFlowRetThds field if non-nil, zero value otherwise.

### GetQosFlowRetThdsOk

`func (o *AnalyticsEventFilterSubsc) GetQosFlowRetThdsOk() (*[]RetainabilityThreshold, bool)`

GetQosFlowRetThdsOk returns a tuple with the QosFlowRetThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowRetThds

`func (o *AnalyticsEventFilterSubsc) SetQosFlowRetThds(v []RetainabilityThreshold)`

SetQosFlowRetThds sets QosFlowRetThds field to given value.

### HasQosFlowRetThds

`func (o *AnalyticsEventFilterSubsc) HasQosFlowRetThds() bool`

HasQosFlowRetThds returns a boolean if a field has been set.

### GetRanUeThrouThds

`func (o *AnalyticsEventFilterSubsc) GetRanUeThrouThds() []string`

GetRanUeThrouThds returns the RanUeThrouThds field if non-nil, zero value otherwise.

### GetRanUeThrouThdsOk

`func (o *AnalyticsEventFilterSubsc) GetRanUeThrouThdsOk() (*[]string, bool)`

GetRanUeThrouThdsOk returns a tuple with the RanUeThrouThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUeThrouThds

`func (o *AnalyticsEventFilterSubsc) SetRanUeThrouThds(v []string)`

SetRanUeThrouThds sets RanUeThrouThds field to given value.

### HasRanUeThrouThds

`func (o *AnalyticsEventFilterSubsc) HasRanUeThrouThds() bool`

HasRanUeThrouThds returns a boolean if a field has been set.

### GetDisperReqs

`func (o *AnalyticsEventFilterSubsc) GetDisperReqs() []DispersionRequirement`

GetDisperReqs returns the DisperReqs field if non-nil, zero value otherwise.

### GetDisperReqsOk

`func (o *AnalyticsEventFilterSubsc) GetDisperReqsOk() (*[]DispersionRequirement, bool)`

GetDisperReqsOk returns a tuple with the DisperReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperReqs

`func (o *AnalyticsEventFilterSubsc) SetDisperReqs(v []DispersionRequirement)`

SetDisperReqs sets DisperReqs field to given value.

### HasDisperReqs

`func (o *AnalyticsEventFilterSubsc) HasDisperReqs() bool`

HasDisperReqs returns a boolean if a field has been set.

### GetListOfAnaSubsets

`func (o *AnalyticsEventFilterSubsc) GetListOfAnaSubsets() []AnalyticsSubset`

GetListOfAnaSubsets returns the ListOfAnaSubsets field if non-nil, zero value otherwise.

### GetListOfAnaSubsetsOk

`func (o *AnalyticsEventFilterSubsc) GetListOfAnaSubsetsOk() (*[]AnalyticsSubset, bool)`

GetListOfAnaSubsetsOk returns a tuple with the ListOfAnaSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfAnaSubsets

`func (o *AnalyticsEventFilterSubsc) SetListOfAnaSubsets(v []AnalyticsSubset)`

SetListOfAnaSubsets sets ListOfAnaSubsets field to given value.

### HasListOfAnaSubsets

`func (o *AnalyticsEventFilterSubsc) HasListOfAnaSubsets() bool`

HasListOfAnaSubsets returns a boolean if a field has been set.

### GetDnPerfReqs

`func (o *AnalyticsEventFilterSubsc) GetDnPerfReqs() []DnPerformanceReq`

GetDnPerfReqs returns the DnPerfReqs field if non-nil, zero value otherwise.

### GetDnPerfReqsOk

`func (o *AnalyticsEventFilterSubsc) GetDnPerfReqsOk() (*[]DnPerformanceReq, bool)`

GetDnPerfReqsOk returns a tuple with the DnPerfReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfReqs

`func (o *AnalyticsEventFilterSubsc) SetDnPerfReqs(v []DnPerformanceReq)`

SetDnPerfReqs sets DnPerfReqs field to given value.

### HasDnPerfReqs

`func (o *AnalyticsEventFilterSubsc) HasDnPerfReqs() bool`

HasDnPerfReqs returns a boolean if a field has been set.

### GetBwRequs

`func (o *AnalyticsEventFilterSubsc) GetBwRequs() []BwRequirement`

GetBwRequs returns the BwRequs field if non-nil, zero value otherwise.

### GetBwRequsOk

`func (o *AnalyticsEventFilterSubsc) GetBwRequsOk() (*[]BwRequirement, bool)`

GetBwRequsOk returns a tuple with the BwRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwRequs

`func (o *AnalyticsEventFilterSubsc) SetBwRequs(v []BwRequirement)`

SetBwRequs sets BwRequs field to given value.

### HasBwRequs

`func (o *AnalyticsEventFilterSubsc) HasBwRequs() bool`

HasBwRequs returns a boolean if a field has been set.

### GetRatFreqs

`func (o *AnalyticsEventFilterSubsc) GetRatFreqs() []RatFreqInformation`

GetRatFreqs returns the RatFreqs field if non-nil, zero value otherwise.

### GetRatFreqsOk

`func (o *AnalyticsEventFilterSubsc) GetRatFreqsOk() (*[]RatFreqInformation, bool)`

GetRatFreqsOk returns a tuple with the RatFreqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatFreqs

`func (o *AnalyticsEventFilterSubsc) SetRatFreqs(v []RatFreqInformation)`

SetRatFreqs sets RatFreqs field to given value.

### HasRatFreqs

`func (o *AnalyticsEventFilterSubsc) HasRatFreqs() bool`

HasRatFreqs returns a boolean if a field has been set.

### GetAppServerAddrs

`func (o *AnalyticsEventFilterSubsc) GetAppServerAddrs() []AddrFqdn`

GetAppServerAddrs returns the AppServerAddrs field if non-nil, zero value otherwise.

### GetAppServerAddrsOk

`func (o *AnalyticsEventFilterSubsc) GetAppServerAddrsOk() (*[]AddrFqdn, bool)`

GetAppServerAddrsOk returns a tuple with the AppServerAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerAddrs

`func (o *AnalyticsEventFilterSubsc) SetAppServerAddrs(v []AddrFqdn)`

SetAppServerAddrs sets AppServerAddrs field to given value.

### HasAppServerAddrs

`func (o *AnalyticsEventFilterSubsc) HasAppServerAddrs() bool`

HasAppServerAddrs returns a boolean if a field has been set.

### GetExtraReportReq

`func (o *AnalyticsEventFilterSubsc) GetExtraReportReq() EventReportingRequirement`

GetExtraReportReq returns the ExtraReportReq field if non-nil, zero value otherwise.

### GetExtraReportReqOk

`func (o *AnalyticsEventFilterSubsc) GetExtraReportReqOk() (*EventReportingRequirement, bool)`

GetExtraReportReqOk returns a tuple with the ExtraReportReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraReportReq

`func (o *AnalyticsEventFilterSubsc) SetExtraReportReq(v EventReportingRequirement)`

SetExtraReportReq sets ExtraReportReq field to given value.

### HasExtraReportReq

`func (o *AnalyticsEventFilterSubsc) HasExtraReportReq() bool`

HasExtraReportReq returns a boolean if a field has been set.

### GetMaxNumOfTopAppUl

`func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppUl() int32`

GetMaxNumOfTopAppUl returns the MaxNumOfTopAppUl field if non-nil, zero value otherwise.

### GetMaxNumOfTopAppUlOk

`func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppUlOk() (*int32, bool)`

GetMaxNumOfTopAppUlOk returns a tuple with the MaxNumOfTopAppUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfTopAppUl

`func (o *AnalyticsEventFilterSubsc) SetMaxNumOfTopAppUl(v int32)`

SetMaxNumOfTopAppUl sets MaxNumOfTopAppUl field to given value.

### HasMaxNumOfTopAppUl

`func (o *AnalyticsEventFilterSubsc) HasMaxNumOfTopAppUl() bool`

HasMaxNumOfTopAppUl returns a boolean if a field has been set.

### GetMaxNumOfTopAppDl

`func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppDl() int32`

GetMaxNumOfTopAppDl returns the MaxNumOfTopAppDl field if non-nil, zero value otherwise.

### GetMaxNumOfTopAppDlOk

`func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppDlOk() (*int32, bool)`

GetMaxNumOfTopAppDlOk returns a tuple with the MaxNumOfTopAppDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfTopAppDl

`func (o *AnalyticsEventFilterSubsc) SetMaxNumOfTopAppDl(v int32)`

SetMaxNumOfTopAppDl sets MaxNumOfTopAppDl field to given value.

### HasMaxNumOfTopAppDl

`func (o *AnalyticsEventFilterSubsc) HasMaxNumOfTopAppDl() bool`

HasMaxNumOfTopAppDl returns a boolean if a field has been set.

### GetVisitedLocAreas

`func (o *AnalyticsEventFilterSubsc) GetVisitedLocAreas() []LocationArea5G`

GetVisitedLocAreas returns the VisitedLocAreas field if non-nil, zero value otherwise.

### GetVisitedLocAreasOk

`func (o *AnalyticsEventFilterSubsc) GetVisitedLocAreasOk() (*[]LocationArea5G, bool)`

GetVisitedLocAreasOk returns a tuple with the VisitedLocAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedLocAreas

`func (o *AnalyticsEventFilterSubsc) SetVisitedLocAreas(v []LocationArea5G)`

SetVisitedLocAreas sets VisitedLocAreas field to given value.

### HasVisitedLocAreas

`func (o *AnalyticsEventFilterSubsc) HasVisitedLocAreas() bool`

HasVisitedLocAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


