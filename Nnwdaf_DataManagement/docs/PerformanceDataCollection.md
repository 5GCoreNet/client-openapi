# PerformanceDataCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**IpTrafficFilter** | Pointer to [**FlowInfo**](FlowInfo.md) |  | [optional] 
**UeLoc** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**AppLocs** | Pointer to **[]string** |  | [optional] 
**AsAddr** | Pointer to [**AddrFqdn**](AddrFqdn.md) |  | [optional] 
**PerfData** | [**PerformanceData**](PerformanceData.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewPerformanceDataCollection

`func NewPerformanceDataCollection(perfData PerformanceData, timeStamp time.Time, ) *PerformanceDataCollection`

NewPerformanceDataCollection instantiates a new PerformanceDataCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceDataCollectionWithDefaults

`func NewPerformanceDataCollectionWithDefaults() *PerformanceDataCollection`

NewPerformanceDataCollectionWithDefaults instantiates a new PerformanceDataCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *PerformanceDataCollection) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PerformanceDataCollection) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PerformanceDataCollection) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *PerformanceDataCollection) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *PerformanceDataCollection) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *PerformanceDataCollection) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *PerformanceDataCollection) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *PerformanceDataCollection) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.

### GetIpTrafficFilter

`func (o *PerformanceDataCollection) GetIpTrafficFilter() FlowInfo`

GetIpTrafficFilter returns the IpTrafficFilter field if non-nil, zero value otherwise.

### GetIpTrafficFilterOk

`func (o *PerformanceDataCollection) GetIpTrafficFilterOk() (*FlowInfo, bool)`

GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTrafficFilter

`func (o *PerformanceDataCollection) SetIpTrafficFilter(v FlowInfo)`

SetIpTrafficFilter sets IpTrafficFilter field to given value.

### HasIpTrafficFilter

`func (o *PerformanceDataCollection) HasIpTrafficFilter() bool`

HasIpTrafficFilter returns a boolean if a field has been set.

### GetUeLoc

`func (o *PerformanceDataCollection) GetUeLoc() LocationArea5G`

GetUeLoc returns the UeLoc field if non-nil, zero value otherwise.

### GetUeLocOk

`func (o *PerformanceDataCollection) GetUeLocOk() (*LocationArea5G, bool)`

GetUeLocOk returns a tuple with the UeLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLoc

`func (o *PerformanceDataCollection) SetUeLoc(v LocationArea5G)`

SetUeLoc sets UeLoc field to given value.

### HasUeLoc

`func (o *PerformanceDataCollection) HasUeLoc() bool`

HasUeLoc returns a boolean if a field has been set.

### GetAppLocs

`func (o *PerformanceDataCollection) GetAppLocs() []string`

GetAppLocs returns the AppLocs field if non-nil, zero value otherwise.

### GetAppLocsOk

`func (o *PerformanceDataCollection) GetAppLocsOk() (*[]string, bool)`

GetAppLocsOk returns a tuple with the AppLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLocs

`func (o *PerformanceDataCollection) SetAppLocs(v []string)`

SetAppLocs sets AppLocs field to given value.

### HasAppLocs

`func (o *PerformanceDataCollection) HasAppLocs() bool`

HasAppLocs returns a boolean if a field has been set.

### GetAsAddr

`func (o *PerformanceDataCollection) GetAsAddr() AddrFqdn`

GetAsAddr returns the AsAddr field if non-nil, zero value otherwise.

### GetAsAddrOk

`func (o *PerformanceDataCollection) GetAsAddrOk() (*AddrFqdn, bool)`

GetAsAddrOk returns a tuple with the AsAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAddr

`func (o *PerformanceDataCollection) SetAsAddr(v AddrFqdn)`

SetAsAddr sets AsAddr field to given value.

### HasAsAddr

`func (o *PerformanceDataCollection) HasAsAddr() bool`

HasAsAddr returns a boolean if a field has been set.

### GetPerfData

`func (o *PerformanceDataCollection) GetPerfData() PerformanceData`

GetPerfData returns the PerfData field if non-nil, zero value otherwise.

### GetPerfDataOk

`func (o *PerformanceDataCollection) GetPerfDataOk() (*PerformanceData, bool)`

GetPerfDataOk returns a tuple with the PerfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfData

`func (o *PerformanceDataCollection) SetPerfData(v PerformanceData)`

SetPerfData sets PerfData field to given value.


### GetTimeStamp

`func (o *PerformanceDataCollection) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *PerformanceDataCollection) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *PerformanceDataCollection) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


