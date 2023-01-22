# MonitorAuthReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**OpenDiscData** | Pointer to [**MonitorDiscDataForOpen**](MonitorDiscDataForOpen.md) |  | [optional] 
**RestrictedDiscData** | Pointer to [**MonitorDiscDataForRestricted**](MonitorDiscDataForRestricted.md) |  | [optional] 

## Methods

### NewMonitorAuthReqData

`func NewMonitorAuthReqData(discType DiscoveryType, ) *MonitorAuthReqData`

NewMonitorAuthReqData instantiates a new MonitorAuthReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorAuthReqDataWithDefaults

`func NewMonitorAuthReqDataWithDefaults() *MonitorAuthReqData`

NewMonitorAuthReqDataWithDefaults instantiates a new MonitorAuthReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *MonitorAuthReqData) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *MonitorAuthReqData) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *MonitorAuthReqData) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetOpenDiscData

`func (o *MonitorAuthReqData) GetOpenDiscData() MonitorDiscDataForOpen`

GetOpenDiscData returns the OpenDiscData field if non-nil, zero value otherwise.

### GetOpenDiscDataOk

`func (o *MonitorAuthReqData) GetOpenDiscDataOk() (*MonitorDiscDataForOpen, bool)`

GetOpenDiscDataOk returns a tuple with the OpenDiscData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDiscData

`func (o *MonitorAuthReqData) SetOpenDiscData(v MonitorDiscDataForOpen)`

SetOpenDiscData sets OpenDiscData field to given value.

### HasOpenDiscData

`func (o *MonitorAuthReqData) HasOpenDiscData() bool`

HasOpenDiscData returns a boolean if a field has been set.

### GetRestrictedDiscData

`func (o *MonitorAuthReqData) GetRestrictedDiscData() MonitorDiscDataForRestricted`

GetRestrictedDiscData returns the RestrictedDiscData field if non-nil, zero value otherwise.

### GetRestrictedDiscDataOk

`func (o *MonitorAuthReqData) GetRestrictedDiscDataOk() (*MonitorDiscDataForRestricted, bool)`

GetRestrictedDiscDataOk returns a tuple with the RestrictedDiscData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDiscData

`func (o *MonitorAuthReqData) SetRestrictedDiscData(v MonitorDiscDataForRestricted)`

SetRestrictedDiscData sets RestrictedDiscData field to given value.

### HasRestrictedDiscData

`func (o *MonitorAuthReqData) HasRestrictedDiscData() bool`

HasRestrictedDiscData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


