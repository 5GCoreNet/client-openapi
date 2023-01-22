# MonitorAuthRespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthDataOpen** | Pointer to [**MonitorAuthDataForOpen**](MonitorAuthDataForOpen.md) |  | [optional] 
**AuthDataRestricted** | Pointer to [**MonitorAuthDataForRestricted**](MonitorAuthDataForRestricted.md) |  | [optional] 

## Methods

### NewMonitorAuthRespData

`func NewMonitorAuthRespData() *MonitorAuthRespData`

NewMonitorAuthRespData instantiates a new MonitorAuthRespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorAuthRespDataWithDefaults

`func NewMonitorAuthRespDataWithDefaults() *MonitorAuthRespData`

NewMonitorAuthRespDataWithDefaults instantiates a new MonitorAuthRespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthDataOpen

`func (o *MonitorAuthRespData) GetAuthDataOpen() MonitorAuthDataForOpen`

GetAuthDataOpen returns the AuthDataOpen field if non-nil, zero value otherwise.

### GetAuthDataOpenOk

`func (o *MonitorAuthRespData) GetAuthDataOpenOk() (*MonitorAuthDataForOpen, bool)`

GetAuthDataOpenOk returns a tuple with the AuthDataOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDataOpen

`func (o *MonitorAuthRespData) SetAuthDataOpen(v MonitorAuthDataForOpen)`

SetAuthDataOpen sets AuthDataOpen field to given value.

### HasAuthDataOpen

`func (o *MonitorAuthRespData) HasAuthDataOpen() bool`

HasAuthDataOpen returns a boolean if a field has been set.

### GetAuthDataRestricted

`func (o *MonitorAuthRespData) GetAuthDataRestricted() MonitorAuthDataForRestricted`

GetAuthDataRestricted returns the AuthDataRestricted field if non-nil, zero value otherwise.

### GetAuthDataRestrictedOk

`func (o *MonitorAuthRespData) GetAuthDataRestrictedOk() (*MonitorAuthDataForRestricted, bool)`

GetAuthDataRestrictedOk returns a tuple with the AuthDataRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDataRestricted

`func (o *MonitorAuthRespData) SetAuthDataRestricted(v MonitorAuthDataForRestricted)`

SetAuthDataRestricted sets AuthDataRestricted field to given value.

### HasAuthDataRestricted

`func (o *MonitorAuthRespData) HasAuthDataRestricted() bool`

HasAuthDataRestricted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


