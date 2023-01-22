# MonitorUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**OpenUpdateData** | Pointer to [**MonitorUpdateDataForOpen**](MonitorUpdateDataForOpen.md) |  | [optional] 
**RestrictedUpdateData** | Pointer to [**MonitorUpdateDataForRestricted**](MonitorUpdateDataForRestricted.md) |  | [optional] 

## Methods

### NewMonitorUpdateData

`func NewMonitorUpdateData(discType DiscoveryType, ) *MonitorUpdateData`

NewMonitorUpdateData instantiates a new MonitorUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorUpdateDataWithDefaults

`func NewMonitorUpdateDataWithDefaults() *MonitorUpdateData`

NewMonitorUpdateDataWithDefaults instantiates a new MonitorUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *MonitorUpdateData) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *MonitorUpdateData) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *MonitorUpdateData) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetOpenUpdateData

`func (o *MonitorUpdateData) GetOpenUpdateData() MonitorUpdateDataForOpen`

GetOpenUpdateData returns the OpenUpdateData field if non-nil, zero value otherwise.

### GetOpenUpdateDataOk

`func (o *MonitorUpdateData) GetOpenUpdateDataOk() (*MonitorUpdateDataForOpen, bool)`

GetOpenUpdateDataOk returns a tuple with the OpenUpdateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenUpdateData

`func (o *MonitorUpdateData) SetOpenUpdateData(v MonitorUpdateDataForOpen)`

SetOpenUpdateData sets OpenUpdateData field to given value.

### HasOpenUpdateData

`func (o *MonitorUpdateData) HasOpenUpdateData() bool`

HasOpenUpdateData returns a boolean if a field has been set.

### GetRestrictedUpdateData

`func (o *MonitorUpdateData) GetRestrictedUpdateData() MonitorUpdateDataForRestricted`

GetRestrictedUpdateData returns the RestrictedUpdateData field if non-nil, zero value otherwise.

### GetRestrictedUpdateDataOk

`func (o *MonitorUpdateData) GetRestrictedUpdateDataOk() (*MonitorUpdateDataForRestricted, bool)`

GetRestrictedUpdateDataOk returns a tuple with the RestrictedUpdateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedUpdateData

`func (o *MonitorUpdateData) SetRestrictedUpdateData(v MonitorUpdateDataForRestricted)`

SetRestrictedUpdateData sets RestrictedUpdateData field to given value.

### HasRestrictedUpdateData

`func (o *MonitorUpdateData) HasRestrictedUpdateData() bool`

HasRestrictedUpdateData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


