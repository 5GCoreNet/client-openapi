# IdentityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupiList** | Pointer to **[]string** |  | [optional] 
**GpsiList** | Pointer to **[]string** |  | [optional] 
**AllowedAfIds** | Pointer to **[]string** |  | [optional] 
**ApplicationPortIds** | Pointer to **map[string]string** | A map (list of key-value pairs where AppPortId serves as key) of GPSIs. | [optional] 

## Methods

### NewIdentityData

`func NewIdentityData() *IdentityData`

NewIdentityData instantiates a new IdentityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDataWithDefaults

`func NewIdentityDataWithDefaults() *IdentityData`

NewIdentityDataWithDefaults instantiates a new IdentityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupiList

`func (o *IdentityData) GetSupiList() []string`

GetSupiList returns the SupiList field if non-nil, zero value otherwise.

### GetSupiListOk

`func (o *IdentityData) GetSupiListOk() (*[]string, bool)`

GetSupiListOk returns a tuple with the SupiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiList

`func (o *IdentityData) SetSupiList(v []string)`

SetSupiList sets SupiList field to given value.

### HasSupiList

`func (o *IdentityData) HasSupiList() bool`

HasSupiList returns a boolean if a field has been set.

### GetGpsiList

`func (o *IdentityData) GetGpsiList() []string`

GetGpsiList returns the GpsiList field if non-nil, zero value otherwise.

### GetGpsiListOk

`func (o *IdentityData) GetGpsiListOk() (*[]string, bool)`

GetGpsiListOk returns a tuple with the GpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiList

`func (o *IdentityData) SetGpsiList(v []string)`

SetGpsiList sets GpsiList field to given value.

### HasGpsiList

`func (o *IdentityData) HasGpsiList() bool`

HasGpsiList returns a boolean if a field has been set.

### GetAllowedAfIds

`func (o *IdentityData) GetAllowedAfIds() []string`

GetAllowedAfIds returns the AllowedAfIds field if non-nil, zero value otherwise.

### GetAllowedAfIdsOk

`func (o *IdentityData) GetAllowedAfIdsOk() (*[]string, bool)`

GetAllowedAfIdsOk returns a tuple with the AllowedAfIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAfIds

`func (o *IdentityData) SetAllowedAfIds(v []string)`

SetAllowedAfIds sets AllowedAfIds field to given value.

### HasAllowedAfIds

`func (o *IdentityData) HasAllowedAfIds() bool`

HasAllowedAfIds returns a boolean if a field has been set.

### GetApplicationPortIds

`func (o *IdentityData) GetApplicationPortIds() map[string]string`

GetApplicationPortIds returns the ApplicationPortIds field if non-nil, zero value otherwise.

### GetApplicationPortIdsOk

`func (o *IdentityData) GetApplicationPortIdsOk() (*map[string]string, bool)`

GetApplicationPortIdsOk returns a tuple with the ApplicationPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPortIds

`func (o *IdentityData) SetApplicationPortIds(v map[string]string)`

SetApplicationPortIds sets ApplicationPortIds field to given value.

### HasApplicationPortIds

`func (o *IdentityData) HasApplicationPortIds() bool`

HasApplicationPortIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


