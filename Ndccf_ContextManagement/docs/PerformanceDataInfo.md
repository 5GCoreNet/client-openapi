# PerformanceDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**IpTrafficFilter** | Pointer to [**FlowInfo**](FlowInfo.md) |  | [optional] 
**UserLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**AppLocs** | Pointer to **[]string** |  | [optional] 
**AsAddr** | Pointer to [**AddrFqdn**](AddrFqdn.md) |  | [optional] 
**PerfData** | [**PerformanceData**](PerformanceData.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewPerformanceDataInfo

`func NewPerformanceDataInfo(perfData PerformanceData, timeStamp time.Time, ) *PerformanceDataInfo`

NewPerformanceDataInfo instantiates a new PerformanceDataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceDataInfoWithDefaults

`func NewPerformanceDataInfoWithDefaults() *PerformanceDataInfo`

NewPerformanceDataInfoWithDefaults instantiates a new PerformanceDataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *PerformanceDataInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PerformanceDataInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PerformanceDataInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *PerformanceDataInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *PerformanceDataInfo) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *PerformanceDataInfo) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *PerformanceDataInfo) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *PerformanceDataInfo) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.

### GetIpTrafficFilter

`func (o *PerformanceDataInfo) GetIpTrafficFilter() FlowInfo`

GetIpTrafficFilter returns the IpTrafficFilter field if non-nil, zero value otherwise.

### GetIpTrafficFilterOk

`func (o *PerformanceDataInfo) GetIpTrafficFilterOk() (*FlowInfo, bool)`

GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTrafficFilter

`func (o *PerformanceDataInfo) SetIpTrafficFilter(v FlowInfo)`

SetIpTrafficFilter sets IpTrafficFilter field to given value.

### HasIpTrafficFilter

`func (o *PerformanceDataInfo) HasIpTrafficFilter() bool`

HasIpTrafficFilter returns a boolean if a field has been set.

### GetUserLoc

`func (o *PerformanceDataInfo) GetUserLoc() UserLocation`

GetUserLoc returns the UserLoc field if non-nil, zero value otherwise.

### GetUserLocOk

`func (o *PerformanceDataInfo) GetUserLocOk() (*UserLocation, bool)`

GetUserLocOk returns a tuple with the UserLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoc

`func (o *PerformanceDataInfo) SetUserLoc(v UserLocation)`

SetUserLoc sets UserLoc field to given value.

### HasUserLoc

`func (o *PerformanceDataInfo) HasUserLoc() bool`

HasUserLoc returns a boolean if a field has been set.

### GetAppLocs

`func (o *PerformanceDataInfo) GetAppLocs() []string`

GetAppLocs returns the AppLocs field if non-nil, zero value otherwise.

### GetAppLocsOk

`func (o *PerformanceDataInfo) GetAppLocsOk() (*[]string, bool)`

GetAppLocsOk returns a tuple with the AppLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLocs

`func (o *PerformanceDataInfo) SetAppLocs(v []string)`

SetAppLocs sets AppLocs field to given value.

### HasAppLocs

`func (o *PerformanceDataInfo) HasAppLocs() bool`

HasAppLocs returns a boolean if a field has been set.

### GetAsAddr

`func (o *PerformanceDataInfo) GetAsAddr() AddrFqdn`

GetAsAddr returns the AsAddr field if non-nil, zero value otherwise.

### GetAsAddrOk

`func (o *PerformanceDataInfo) GetAsAddrOk() (*AddrFqdn, bool)`

GetAsAddrOk returns a tuple with the AsAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAddr

`func (o *PerformanceDataInfo) SetAsAddr(v AddrFqdn)`

SetAsAddr sets AsAddr field to given value.

### HasAsAddr

`func (o *PerformanceDataInfo) HasAsAddr() bool`

HasAsAddr returns a boolean if a field has been set.

### GetPerfData

`func (o *PerformanceDataInfo) GetPerfData() PerformanceData`

GetPerfData returns the PerfData field if non-nil, zero value otherwise.

### GetPerfDataOk

`func (o *PerformanceDataInfo) GetPerfDataOk() (*PerformanceData, bool)`

GetPerfDataOk returns a tuple with the PerfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfData

`func (o *PerformanceDataInfo) SetPerfData(v PerformanceData)`

SetPerfData sets PerfData field to given value.


### GetTimeStamp

`func (o *PerformanceDataInfo) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *PerformanceDataInfo) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *PerformanceDataInfo) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


