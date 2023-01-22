# MonitorAuthDataForOpen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProseAppCodes** | Pointer to **[]string** |  | [optional] 
**ProseAppPrefix** | Pointer to **string** | Contains the Prose Application Code Prefix. | [optional] 
**ProseAppMasks** | **[]string** |  | 
**Ttl** | **int32** |  | 

## Methods

### NewMonitorAuthDataForOpen

`func NewMonitorAuthDataForOpen(proseAppMasks []string, ttl int32, ) *MonitorAuthDataForOpen`

NewMonitorAuthDataForOpen instantiates a new MonitorAuthDataForOpen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorAuthDataForOpenWithDefaults

`func NewMonitorAuthDataForOpenWithDefaults() *MonitorAuthDataForOpen`

NewMonitorAuthDataForOpenWithDefaults instantiates a new MonitorAuthDataForOpen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProseAppCodes

`func (o *MonitorAuthDataForOpen) GetProseAppCodes() []string`

GetProseAppCodes returns the ProseAppCodes field if non-nil, zero value otherwise.

### GetProseAppCodesOk

`func (o *MonitorAuthDataForOpen) GetProseAppCodesOk() (*[]string, bool)`

GetProseAppCodesOk returns a tuple with the ProseAppCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppCodes

`func (o *MonitorAuthDataForOpen) SetProseAppCodes(v []string)`

SetProseAppCodes sets ProseAppCodes field to given value.

### HasProseAppCodes

`func (o *MonitorAuthDataForOpen) HasProseAppCodes() bool`

HasProseAppCodes returns a boolean if a field has been set.

### GetProseAppPrefix

`func (o *MonitorAuthDataForOpen) GetProseAppPrefix() string`

GetProseAppPrefix returns the ProseAppPrefix field if non-nil, zero value otherwise.

### GetProseAppPrefixOk

`func (o *MonitorAuthDataForOpen) GetProseAppPrefixOk() (*string, bool)`

GetProseAppPrefixOk returns a tuple with the ProseAppPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppPrefix

`func (o *MonitorAuthDataForOpen) SetProseAppPrefix(v string)`

SetProseAppPrefix sets ProseAppPrefix field to given value.

### HasProseAppPrefix

`func (o *MonitorAuthDataForOpen) HasProseAppPrefix() bool`

HasProseAppPrefix returns a boolean if a field has been set.

### GetProseAppMasks

`func (o *MonitorAuthDataForOpen) GetProseAppMasks() []string`

GetProseAppMasks returns the ProseAppMasks field if non-nil, zero value otherwise.

### GetProseAppMasksOk

`func (o *MonitorAuthDataForOpen) GetProseAppMasksOk() (*[]string, bool)`

GetProseAppMasksOk returns a tuple with the ProseAppMasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppMasks

`func (o *MonitorAuthDataForOpen) SetProseAppMasks(v []string)`

SetProseAppMasks sets ProseAppMasks field to given value.


### GetTtl

`func (o *MonitorAuthDataForOpen) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MonitorAuthDataForOpen) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *MonitorAuthDataForOpen) SetTtl(v int32)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


