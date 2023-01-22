# TimeSyncExposureConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpNodeId** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | 
**ReqPtpIns** | [**PtpInstance**](PtpInstance.md) |  | 
**GmEnable** | Pointer to **bool** | Indicates that the AF requests 5GS to act as a grandmaster for PTP or gPTP if it is included and set to true.  | [optional] 
**GmPrio** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TimeDom** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**TimeSyncErrBdgt** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**ConfigNotifId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**ConfigNotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**TempValidity** | Pointer to [**TemporalValidity**](TemporalValidity.md) |  | [optional] 

## Methods

### NewTimeSyncExposureConfig1

`func NewTimeSyncExposureConfig1(upNodeId int32, reqPtpIns PtpInstance, timeDom int32, configNotifId string, configNotifUri string, ) *TimeSyncExposureConfig1`

NewTimeSyncExposureConfig1 instantiates a new TimeSyncExposureConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncExposureConfig1WithDefaults

`func NewTimeSyncExposureConfig1WithDefaults() *TimeSyncExposureConfig1`

NewTimeSyncExposureConfig1WithDefaults instantiates a new TimeSyncExposureConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpNodeId

`func (o *TimeSyncExposureConfig1) GetUpNodeId() int32`

GetUpNodeId returns the UpNodeId field if non-nil, zero value otherwise.

### GetUpNodeIdOk

`func (o *TimeSyncExposureConfig1) GetUpNodeIdOk() (*int32, bool)`

GetUpNodeIdOk returns a tuple with the UpNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpNodeId

`func (o *TimeSyncExposureConfig1) SetUpNodeId(v int32)`

SetUpNodeId sets UpNodeId field to given value.


### GetReqPtpIns

`func (o *TimeSyncExposureConfig1) GetReqPtpIns() PtpInstance`

GetReqPtpIns returns the ReqPtpIns field if non-nil, zero value otherwise.

### GetReqPtpInsOk

`func (o *TimeSyncExposureConfig1) GetReqPtpInsOk() (*PtpInstance, bool)`

GetReqPtpInsOk returns a tuple with the ReqPtpIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqPtpIns

`func (o *TimeSyncExposureConfig1) SetReqPtpIns(v PtpInstance)`

SetReqPtpIns sets ReqPtpIns field to given value.


### GetGmEnable

`func (o *TimeSyncExposureConfig1) GetGmEnable() bool`

GetGmEnable returns the GmEnable field if non-nil, zero value otherwise.

### GetGmEnableOk

`func (o *TimeSyncExposureConfig1) GetGmEnableOk() (*bool, bool)`

GetGmEnableOk returns a tuple with the GmEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmEnable

`func (o *TimeSyncExposureConfig1) SetGmEnable(v bool)`

SetGmEnable sets GmEnable field to given value.

### HasGmEnable

`func (o *TimeSyncExposureConfig1) HasGmEnable() bool`

HasGmEnable returns a boolean if a field has been set.

### GetGmPrio

`func (o *TimeSyncExposureConfig1) GetGmPrio() int32`

GetGmPrio returns the GmPrio field if non-nil, zero value otherwise.

### GetGmPrioOk

`func (o *TimeSyncExposureConfig1) GetGmPrioOk() (*int32, bool)`

GetGmPrioOk returns a tuple with the GmPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmPrio

`func (o *TimeSyncExposureConfig1) SetGmPrio(v int32)`

SetGmPrio sets GmPrio field to given value.

### HasGmPrio

`func (o *TimeSyncExposureConfig1) HasGmPrio() bool`

HasGmPrio returns a boolean if a field has been set.

### GetTimeDom

`func (o *TimeSyncExposureConfig1) GetTimeDom() int32`

GetTimeDom returns the TimeDom field if non-nil, zero value otherwise.

### GetTimeDomOk

`func (o *TimeSyncExposureConfig1) GetTimeDomOk() (*int32, bool)`

GetTimeDomOk returns a tuple with the TimeDom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDom

`func (o *TimeSyncExposureConfig1) SetTimeDom(v int32)`

SetTimeDom sets TimeDom field to given value.


### GetTimeSyncErrBdgt

`func (o *TimeSyncExposureConfig1) GetTimeSyncErrBdgt() int32`

GetTimeSyncErrBdgt returns the TimeSyncErrBdgt field if non-nil, zero value otherwise.

### GetTimeSyncErrBdgtOk

`func (o *TimeSyncExposureConfig1) GetTimeSyncErrBdgtOk() (*int32, bool)`

GetTimeSyncErrBdgtOk returns a tuple with the TimeSyncErrBdgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSyncErrBdgt

`func (o *TimeSyncExposureConfig1) SetTimeSyncErrBdgt(v int32)`

SetTimeSyncErrBdgt sets TimeSyncErrBdgt field to given value.

### HasTimeSyncErrBdgt

`func (o *TimeSyncExposureConfig1) HasTimeSyncErrBdgt() bool`

HasTimeSyncErrBdgt returns a boolean if a field has been set.

### GetConfigNotifId

`func (o *TimeSyncExposureConfig1) GetConfigNotifId() string`

GetConfigNotifId returns the ConfigNotifId field if non-nil, zero value otherwise.

### GetConfigNotifIdOk

`func (o *TimeSyncExposureConfig1) GetConfigNotifIdOk() (*string, bool)`

GetConfigNotifIdOk returns a tuple with the ConfigNotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigNotifId

`func (o *TimeSyncExposureConfig1) SetConfigNotifId(v string)`

SetConfigNotifId sets ConfigNotifId field to given value.


### GetConfigNotifUri

`func (o *TimeSyncExposureConfig1) GetConfigNotifUri() string`

GetConfigNotifUri returns the ConfigNotifUri field if non-nil, zero value otherwise.

### GetConfigNotifUriOk

`func (o *TimeSyncExposureConfig1) GetConfigNotifUriOk() (*string, bool)`

GetConfigNotifUriOk returns a tuple with the ConfigNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigNotifUri

`func (o *TimeSyncExposureConfig1) SetConfigNotifUri(v string)`

SetConfigNotifUri sets ConfigNotifUri field to given value.


### GetTempValidity

`func (o *TimeSyncExposureConfig1) GetTempValidity() TemporalValidity`

GetTempValidity returns the TempValidity field if non-nil, zero value otherwise.

### GetTempValidityOk

`func (o *TimeSyncExposureConfig1) GetTempValidityOk() (*TemporalValidity, bool)`

GetTempValidityOk returns a tuple with the TempValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempValidity

`func (o *TimeSyncExposureConfig1) SetTempValidity(v TemporalValidity)`

SetTempValidity sets TempValidity field to given value.

### HasTempValidity

`func (o *TimeSyncExposureConfig1) HasTempValidity() bool`

HasTempValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


