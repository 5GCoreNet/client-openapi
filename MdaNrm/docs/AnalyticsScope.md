# AnalyticsScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedEntitiesScope** | Pointer to **[]string** |  | [optional] 
**AreaScope** | Pointer to [**[]GeoArea**](GeoArea.md) |  | [optional] 

## Methods

### NewAnalyticsScope

`func NewAnalyticsScope() *AnalyticsScope`

NewAnalyticsScope instantiates a new AnalyticsScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsScopeWithDefaults

`func NewAnalyticsScopeWithDefaults() *AnalyticsScope`

NewAnalyticsScopeWithDefaults instantiates a new AnalyticsScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagedEntitiesScope

`func (o *AnalyticsScope) GetManagedEntitiesScope() []string`

GetManagedEntitiesScope returns the ManagedEntitiesScope field if non-nil, zero value otherwise.

### GetManagedEntitiesScopeOk

`func (o *AnalyticsScope) GetManagedEntitiesScopeOk() (*[]string, bool)`

GetManagedEntitiesScopeOk returns a tuple with the ManagedEntitiesScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedEntitiesScope

`func (o *AnalyticsScope) SetManagedEntitiesScope(v []string)`

SetManagedEntitiesScope sets ManagedEntitiesScope field to given value.

### HasManagedEntitiesScope

`func (o *AnalyticsScope) HasManagedEntitiesScope() bool`

HasManagedEntitiesScope returns a boolean if a field has been set.

### GetAreaScope

`func (o *AnalyticsScope) GetAreaScope() []GeoArea`

GetAreaScope returns the AreaScope field if non-nil, zero value otherwise.

### GetAreaScopeOk

`func (o *AnalyticsScope) GetAreaScopeOk() (*[]GeoArea, bool)`

GetAreaScopeOk returns a tuple with the AreaScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaScope

`func (o *AnalyticsScope) SetAreaScope(v []GeoArea)`

SetAreaScope sets AreaScope field to given value.

### HasAreaScope

`func (o *AnalyticsScope) HasAreaScope() bool`

HasAreaScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


