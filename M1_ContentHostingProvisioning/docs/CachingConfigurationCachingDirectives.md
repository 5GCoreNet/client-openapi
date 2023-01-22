# CachingConfigurationCachingDirectives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCodeFilters** | Pointer to **[]int32** |  | [optional] 
**NoCache** | **bool** |  | 
**MaxAge** | Pointer to **int32** |  | [optional] 

## Methods

### NewCachingConfigurationCachingDirectives

`func NewCachingConfigurationCachingDirectives(noCache bool, ) *CachingConfigurationCachingDirectives`

NewCachingConfigurationCachingDirectives instantiates a new CachingConfigurationCachingDirectives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCachingConfigurationCachingDirectivesWithDefaults

`func NewCachingConfigurationCachingDirectivesWithDefaults() *CachingConfigurationCachingDirectives`

NewCachingConfigurationCachingDirectivesWithDefaults instantiates a new CachingConfigurationCachingDirectives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCodeFilters

`func (o *CachingConfigurationCachingDirectives) GetStatusCodeFilters() []int32`

GetStatusCodeFilters returns the StatusCodeFilters field if non-nil, zero value otherwise.

### GetStatusCodeFiltersOk

`func (o *CachingConfigurationCachingDirectives) GetStatusCodeFiltersOk() (*[]int32, bool)`

GetStatusCodeFiltersOk returns a tuple with the StatusCodeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeFilters

`func (o *CachingConfigurationCachingDirectives) SetStatusCodeFilters(v []int32)`

SetStatusCodeFilters sets StatusCodeFilters field to given value.

### HasStatusCodeFilters

`func (o *CachingConfigurationCachingDirectives) HasStatusCodeFilters() bool`

HasStatusCodeFilters returns a boolean if a field has been set.

### GetNoCache

`func (o *CachingConfigurationCachingDirectives) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *CachingConfigurationCachingDirectives) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *CachingConfigurationCachingDirectives) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.


### GetMaxAge

`func (o *CachingConfigurationCachingDirectives) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CachingConfigurationCachingDirectives) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CachingConfigurationCachingDirectives) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CachingConfigurationCachingDirectives) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


