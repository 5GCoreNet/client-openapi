# MonitorUpdateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**ProseRestrictedCode** | **string** | Contains the ProSe Restricted Code. | 
**AppId** | **string** | Contains the Application ID. | 
**BannedRpauid** | **string** | Contains the RPAUID. | 
**BannedPduid** | **string** | Contains the PDUID. | 
**RevocationResult** | [**RevocationResult**](RevocationResult.md) |  | 

## Methods

### NewMonitorUpdateResult

`func NewMonitorUpdateResult(discType DiscoveryType, proseRestrictedCode string, appId string, bannedRpauid string, bannedPduid string, revocationResult RevocationResult, ) *MonitorUpdateResult`

NewMonitorUpdateResult instantiates a new MonitorUpdateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorUpdateResultWithDefaults

`func NewMonitorUpdateResultWithDefaults() *MonitorUpdateResult`

NewMonitorUpdateResultWithDefaults instantiates a new MonitorUpdateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *MonitorUpdateResult) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *MonitorUpdateResult) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *MonitorUpdateResult) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetProseRestrictedCode

`func (o *MonitorUpdateResult) GetProseRestrictedCode() string`

GetProseRestrictedCode returns the ProseRestrictedCode field if non-nil, zero value otherwise.

### GetProseRestrictedCodeOk

`func (o *MonitorUpdateResult) GetProseRestrictedCodeOk() (*string, bool)`

GetProseRestrictedCodeOk returns a tuple with the ProseRestrictedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseRestrictedCode

`func (o *MonitorUpdateResult) SetProseRestrictedCode(v string)`

SetProseRestrictedCode sets ProseRestrictedCode field to given value.


### GetAppId

`func (o *MonitorUpdateResult) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MonitorUpdateResult) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MonitorUpdateResult) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetBannedRpauid

`func (o *MonitorUpdateResult) GetBannedRpauid() string`

GetBannedRpauid returns the BannedRpauid field if non-nil, zero value otherwise.

### GetBannedRpauidOk

`func (o *MonitorUpdateResult) GetBannedRpauidOk() (*string, bool)`

GetBannedRpauidOk returns a tuple with the BannedRpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedRpauid

`func (o *MonitorUpdateResult) SetBannedRpauid(v string)`

SetBannedRpauid sets BannedRpauid field to given value.


### GetBannedPduid

`func (o *MonitorUpdateResult) GetBannedPduid() string`

GetBannedPduid returns the BannedPduid field if non-nil, zero value otherwise.

### GetBannedPduidOk

`func (o *MonitorUpdateResult) GetBannedPduidOk() (*string, bool)`

GetBannedPduidOk returns a tuple with the BannedPduid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedPduid

`func (o *MonitorUpdateResult) SetBannedPduid(v string)`

SetBannedPduid sets BannedPduid field to given value.


### GetRevocationResult

`func (o *MonitorUpdateResult) GetRevocationResult() RevocationResult`

GetRevocationResult returns the RevocationResult field if non-nil, zero value otherwise.

### GetRevocationResultOk

`func (o *MonitorUpdateResult) GetRevocationResultOk() (*RevocationResult, bool)`

GetRevocationResultOk returns a tuple with the RevocationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationResult

`func (o *MonitorUpdateResult) SetRevocationResult(v RevocationResult)`

SetRevocationResult sets RevocationResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


