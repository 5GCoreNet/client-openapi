# TimeSyncExposureConfigNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigNotifId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**StateOfConfig** | [**StateOfConfiguration**](StateOfConfiguration.md) |  | 

## Methods

### NewTimeSyncExposureConfigNotif

`func NewTimeSyncExposureConfigNotif(configNotifId string, stateOfConfig StateOfConfiguration, ) *TimeSyncExposureConfigNotif`

NewTimeSyncExposureConfigNotif instantiates a new TimeSyncExposureConfigNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncExposureConfigNotifWithDefaults

`func NewTimeSyncExposureConfigNotifWithDefaults() *TimeSyncExposureConfigNotif`

NewTimeSyncExposureConfigNotifWithDefaults instantiates a new TimeSyncExposureConfigNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigNotifId

`func (o *TimeSyncExposureConfigNotif) GetConfigNotifId() string`

GetConfigNotifId returns the ConfigNotifId field if non-nil, zero value otherwise.

### GetConfigNotifIdOk

`func (o *TimeSyncExposureConfigNotif) GetConfigNotifIdOk() (*string, bool)`

GetConfigNotifIdOk returns a tuple with the ConfigNotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigNotifId

`func (o *TimeSyncExposureConfigNotif) SetConfigNotifId(v string)`

SetConfigNotifId sets ConfigNotifId field to given value.


### GetStateOfConfig

`func (o *TimeSyncExposureConfigNotif) GetStateOfConfig() StateOfConfiguration`

GetStateOfConfig returns the StateOfConfig field if non-nil, zero value otherwise.

### GetStateOfConfigOk

`func (o *TimeSyncExposureConfigNotif) GetStateOfConfigOk() (*StateOfConfiguration, bool)`

GetStateOfConfigOk returns a tuple with the StateOfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfConfig

`func (o *TimeSyncExposureConfigNotif) SetStateOfConfig(v StateOfConfiguration)`

SetStateOfConfig sets StateOfConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


