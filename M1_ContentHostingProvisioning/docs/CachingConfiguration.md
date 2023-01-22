# CachingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlPatternFilter** | **string** |  | 
**CachingDirectives** | Pointer to [**CachingConfigurationCachingDirectives**](CachingConfigurationCachingDirectives.md) |  | [optional] 

## Methods

### NewCachingConfiguration

`func NewCachingConfiguration(urlPatternFilter string, ) *CachingConfiguration`

NewCachingConfiguration instantiates a new CachingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCachingConfigurationWithDefaults

`func NewCachingConfigurationWithDefaults() *CachingConfiguration`

NewCachingConfigurationWithDefaults instantiates a new CachingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlPatternFilter

`func (o *CachingConfiguration) GetUrlPatternFilter() string`

GetUrlPatternFilter returns the UrlPatternFilter field if non-nil, zero value otherwise.

### GetUrlPatternFilterOk

`func (o *CachingConfiguration) GetUrlPatternFilterOk() (*string, bool)`

GetUrlPatternFilterOk returns a tuple with the UrlPatternFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPatternFilter

`func (o *CachingConfiguration) SetUrlPatternFilter(v string)`

SetUrlPatternFilter sets UrlPatternFilter field to given value.


### GetCachingDirectives

`func (o *CachingConfiguration) GetCachingDirectives() CachingConfigurationCachingDirectives`

GetCachingDirectives returns the CachingDirectives field if non-nil, zero value otherwise.

### GetCachingDirectivesOk

`func (o *CachingConfiguration) GetCachingDirectivesOk() (*CachingConfigurationCachingDirectives, bool)`

GetCachingDirectivesOk returns a tuple with the CachingDirectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingDirectives

`func (o *CachingConfiguration) SetCachingDirectives(v CachingConfigurationCachingDirectives)`

SetCachingDirectives sets CachingDirectives field to given value.

### HasCachingDirectives

`func (o *CachingConfiguration) HasCachingDirectives() bool`

HasCachingDirectives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


