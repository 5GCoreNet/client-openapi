# AnalyticsEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Dnais** | Pointer to **[]string** |  | [optional] 
**NwPerfTypes** | Pointer to [**[]NetworkPerfType**](NetworkPerfType.md) |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**ExcepIds** | Pointer to [**[]ExceptionId**](ExceptionId.md) |  | [optional] 
**ExptAnaType** | Pointer to [**ExpectedAnalyticsType**](ExpectedAnalyticsType.md) |  | [optional] 
**ExptUeBehav** | Pointer to [**ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**NsiIdInfos** | Pointer to [**[]NsiIdInfo**](NsiIdInfo.md) |  | [optional] 
**QosReq** | Pointer to [**QosRequirement**](QosRequirement.md) |  | [optional] 
**ListOfAnaSubsets** | Pointer to [**[]AnalyticsSubset**](AnalyticsSubset.md) |  | [optional] 
**DnPerfReqs** | Pointer to [**[]DnPerformanceReq**](DnPerformanceReq.md) |  | [optional] 
**BwRequs** | Pointer to [**[]BwRequirement**](BwRequirement.md) |  | [optional] 
**RatFreqs** | Pointer to [**[]RatFreqInformation**](RatFreqInformation.md) |  | [optional] 
**AppServerAddrs** | Pointer to [**[]AddrFqdn**](AddrFqdn.md) |  | [optional] 
**MaxNumOfTopAppUl** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MaxNumOfTopAppDl** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**VisitedLocAreas** | Pointer to [**[]LocationArea5G**](LocationArea5G.md) |  | [optional] 

## Methods

### NewAnalyticsEventFilter

`func NewAnalyticsEventFilter() *AnalyticsEventFilter`

NewAnalyticsEventFilter instantiates a new AnalyticsEventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsEventFilterWithDefaults

`func NewAnalyticsEventFilterWithDefaults() *AnalyticsEventFilter`

NewAnalyticsEventFilterWithDefaults instantiates a new AnalyticsEventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocArea

`func (o *AnalyticsEventFilter) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *AnalyticsEventFilter) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *AnalyticsEventFilter) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *AnalyticsEventFilter) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetDnn

`func (o *AnalyticsEventFilter) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *AnalyticsEventFilter) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *AnalyticsEventFilter) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *AnalyticsEventFilter) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetDnais

`func (o *AnalyticsEventFilter) GetDnais() []string`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *AnalyticsEventFilter) GetDnaisOk() (*[]string, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *AnalyticsEventFilter) SetDnais(v []string)`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *AnalyticsEventFilter) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetNwPerfTypes

`func (o *AnalyticsEventFilter) GetNwPerfTypes() []NetworkPerfType`

GetNwPerfTypes returns the NwPerfTypes field if non-nil, zero value otherwise.

### GetNwPerfTypesOk

`func (o *AnalyticsEventFilter) GetNwPerfTypesOk() (*[]NetworkPerfType, bool)`

GetNwPerfTypesOk returns a tuple with the NwPerfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfTypes

`func (o *AnalyticsEventFilter) SetNwPerfTypes(v []NetworkPerfType)`

SetNwPerfTypes sets NwPerfTypes field to given value.

### HasNwPerfTypes

`func (o *AnalyticsEventFilter) HasNwPerfTypes() bool`

HasNwPerfTypes returns a boolean if a field has been set.

### GetAppIds

`func (o *AnalyticsEventFilter) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AnalyticsEventFilter) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AnalyticsEventFilter) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *AnalyticsEventFilter) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetExcepIds

`func (o *AnalyticsEventFilter) GetExcepIds() []ExceptionId`

GetExcepIds returns the ExcepIds field if non-nil, zero value otherwise.

### GetExcepIdsOk

`func (o *AnalyticsEventFilter) GetExcepIdsOk() (*[]ExceptionId, bool)`

GetExcepIdsOk returns a tuple with the ExcepIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepIds

`func (o *AnalyticsEventFilter) SetExcepIds(v []ExceptionId)`

SetExcepIds sets ExcepIds field to given value.

### HasExcepIds

`func (o *AnalyticsEventFilter) HasExcepIds() bool`

HasExcepIds returns a boolean if a field has been set.

### GetExptAnaType

`func (o *AnalyticsEventFilter) GetExptAnaType() ExpectedAnalyticsType`

GetExptAnaType returns the ExptAnaType field if non-nil, zero value otherwise.

### GetExptAnaTypeOk

`func (o *AnalyticsEventFilter) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool)`

GetExptAnaTypeOk returns a tuple with the ExptAnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptAnaType

`func (o *AnalyticsEventFilter) SetExptAnaType(v ExpectedAnalyticsType)`

SetExptAnaType sets ExptAnaType field to given value.

### HasExptAnaType

`func (o *AnalyticsEventFilter) HasExptAnaType() bool`

HasExptAnaType returns a boolean if a field has been set.

### GetExptUeBehav

`func (o *AnalyticsEventFilter) GetExptUeBehav() ExpectedUeBehaviourData`

GetExptUeBehav returns the ExptUeBehav field if non-nil, zero value otherwise.

### GetExptUeBehavOk

`func (o *AnalyticsEventFilter) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool)`

GetExptUeBehavOk returns a tuple with the ExptUeBehav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptUeBehav

`func (o *AnalyticsEventFilter) SetExptUeBehav(v ExpectedUeBehaviourData)`

SetExptUeBehav sets ExptUeBehav field to given value.

### HasExptUeBehav

`func (o *AnalyticsEventFilter) HasExptUeBehav() bool`

HasExptUeBehav returns a boolean if a field has been set.

### GetSnssai

`func (o *AnalyticsEventFilter) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AnalyticsEventFilter) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AnalyticsEventFilter) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *AnalyticsEventFilter) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetNsiIdInfos

`func (o *AnalyticsEventFilter) GetNsiIdInfos() []NsiIdInfo`

GetNsiIdInfos returns the NsiIdInfos field if non-nil, zero value otherwise.

### GetNsiIdInfosOk

`func (o *AnalyticsEventFilter) GetNsiIdInfosOk() (*[]NsiIdInfo, bool)`

GetNsiIdInfosOk returns a tuple with the NsiIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiIdInfos

`func (o *AnalyticsEventFilter) SetNsiIdInfos(v []NsiIdInfo)`

SetNsiIdInfos sets NsiIdInfos field to given value.

### HasNsiIdInfos

`func (o *AnalyticsEventFilter) HasNsiIdInfos() bool`

HasNsiIdInfos returns a boolean if a field has been set.

### GetQosReq

`func (o *AnalyticsEventFilter) GetQosReq() QosRequirement`

GetQosReq returns the QosReq field if non-nil, zero value otherwise.

### GetQosReqOk

`func (o *AnalyticsEventFilter) GetQosReqOk() (*QosRequirement, bool)`

GetQosReqOk returns a tuple with the QosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReq

`func (o *AnalyticsEventFilter) SetQosReq(v QosRequirement)`

SetQosReq sets QosReq field to given value.

### HasQosReq

`func (o *AnalyticsEventFilter) HasQosReq() bool`

HasQosReq returns a boolean if a field has been set.

### GetListOfAnaSubsets

`func (o *AnalyticsEventFilter) GetListOfAnaSubsets() []AnalyticsSubset`

GetListOfAnaSubsets returns the ListOfAnaSubsets field if non-nil, zero value otherwise.

### GetListOfAnaSubsetsOk

`func (o *AnalyticsEventFilter) GetListOfAnaSubsetsOk() (*[]AnalyticsSubset, bool)`

GetListOfAnaSubsetsOk returns a tuple with the ListOfAnaSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfAnaSubsets

`func (o *AnalyticsEventFilter) SetListOfAnaSubsets(v []AnalyticsSubset)`

SetListOfAnaSubsets sets ListOfAnaSubsets field to given value.

### HasListOfAnaSubsets

`func (o *AnalyticsEventFilter) HasListOfAnaSubsets() bool`

HasListOfAnaSubsets returns a boolean if a field has been set.

### GetDnPerfReqs

`func (o *AnalyticsEventFilter) GetDnPerfReqs() []DnPerformanceReq`

GetDnPerfReqs returns the DnPerfReqs field if non-nil, zero value otherwise.

### GetDnPerfReqsOk

`func (o *AnalyticsEventFilter) GetDnPerfReqsOk() (*[]DnPerformanceReq, bool)`

GetDnPerfReqsOk returns a tuple with the DnPerfReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfReqs

`func (o *AnalyticsEventFilter) SetDnPerfReqs(v []DnPerformanceReq)`

SetDnPerfReqs sets DnPerfReqs field to given value.

### HasDnPerfReqs

`func (o *AnalyticsEventFilter) HasDnPerfReqs() bool`

HasDnPerfReqs returns a boolean if a field has been set.

### GetBwRequs

`func (o *AnalyticsEventFilter) GetBwRequs() []BwRequirement`

GetBwRequs returns the BwRequs field if non-nil, zero value otherwise.

### GetBwRequsOk

`func (o *AnalyticsEventFilter) GetBwRequsOk() (*[]BwRequirement, bool)`

GetBwRequsOk returns a tuple with the BwRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwRequs

`func (o *AnalyticsEventFilter) SetBwRequs(v []BwRequirement)`

SetBwRequs sets BwRequs field to given value.

### HasBwRequs

`func (o *AnalyticsEventFilter) HasBwRequs() bool`

HasBwRequs returns a boolean if a field has been set.

### GetRatFreqs

`func (o *AnalyticsEventFilter) GetRatFreqs() []RatFreqInformation`

GetRatFreqs returns the RatFreqs field if non-nil, zero value otherwise.

### GetRatFreqsOk

`func (o *AnalyticsEventFilter) GetRatFreqsOk() (*[]RatFreqInformation, bool)`

GetRatFreqsOk returns a tuple with the RatFreqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatFreqs

`func (o *AnalyticsEventFilter) SetRatFreqs(v []RatFreqInformation)`

SetRatFreqs sets RatFreqs field to given value.

### HasRatFreqs

`func (o *AnalyticsEventFilter) HasRatFreqs() bool`

HasRatFreqs returns a boolean if a field has been set.

### GetAppServerAddrs

`func (o *AnalyticsEventFilter) GetAppServerAddrs() []AddrFqdn`

GetAppServerAddrs returns the AppServerAddrs field if non-nil, zero value otherwise.

### GetAppServerAddrsOk

`func (o *AnalyticsEventFilter) GetAppServerAddrsOk() (*[]AddrFqdn, bool)`

GetAppServerAddrsOk returns a tuple with the AppServerAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerAddrs

`func (o *AnalyticsEventFilter) SetAppServerAddrs(v []AddrFqdn)`

SetAppServerAddrs sets AppServerAddrs field to given value.

### HasAppServerAddrs

`func (o *AnalyticsEventFilter) HasAppServerAddrs() bool`

HasAppServerAddrs returns a boolean if a field has been set.

### GetMaxNumOfTopAppUl

`func (o *AnalyticsEventFilter) GetMaxNumOfTopAppUl() int32`

GetMaxNumOfTopAppUl returns the MaxNumOfTopAppUl field if non-nil, zero value otherwise.

### GetMaxNumOfTopAppUlOk

`func (o *AnalyticsEventFilter) GetMaxNumOfTopAppUlOk() (*int32, bool)`

GetMaxNumOfTopAppUlOk returns a tuple with the MaxNumOfTopAppUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfTopAppUl

`func (o *AnalyticsEventFilter) SetMaxNumOfTopAppUl(v int32)`

SetMaxNumOfTopAppUl sets MaxNumOfTopAppUl field to given value.

### HasMaxNumOfTopAppUl

`func (o *AnalyticsEventFilter) HasMaxNumOfTopAppUl() bool`

HasMaxNumOfTopAppUl returns a boolean if a field has been set.

### GetMaxNumOfTopAppDl

`func (o *AnalyticsEventFilter) GetMaxNumOfTopAppDl() int32`

GetMaxNumOfTopAppDl returns the MaxNumOfTopAppDl field if non-nil, zero value otherwise.

### GetMaxNumOfTopAppDlOk

`func (o *AnalyticsEventFilter) GetMaxNumOfTopAppDlOk() (*int32, bool)`

GetMaxNumOfTopAppDlOk returns a tuple with the MaxNumOfTopAppDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfTopAppDl

`func (o *AnalyticsEventFilter) SetMaxNumOfTopAppDl(v int32)`

SetMaxNumOfTopAppDl sets MaxNumOfTopAppDl field to given value.

### HasMaxNumOfTopAppDl

`func (o *AnalyticsEventFilter) HasMaxNumOfTopAppDl() bool`

HasMaxNumOfTopAppDl returns a boolean if a field has been set.

### GetVisitedLocAreas

`func (o *AnalyticsEventFilter) GetVisitedLocAreas() []LocationArea5G`

GetVisitedLocAreas returns the VisitedLocAreas field if non-nil, zero value otherwise.

### GetVisitedLocAreasOk

`func (o *AnalyticsEventFilter) GetVisitedLocAreasOk() (*[]LocationArea5G, bool)`

GetVisitedLocAreasOk returns a tuple with the VisitedLocAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedLocAreas

`func (o *AnalyticsEventFilter) SetVisitedLocAreas(v []LocationArea5G)`

SetVisitedLocAreas sets VisitedLocAreas field to given value.

### HasVisitedLocAreas

`func (o *AnalyticsEventFilter) HasVisitedLocAreas() bool`

HasVisitedLocAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


