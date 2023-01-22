# PathRewriteRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestPattern** | **string** |  | 
**MappedPath** | **string** |  | 

## Methods

### NewPathRewriteRule

`func NewPathRewriteRule(requestPattern string, mappedPath string, ) *PathRewriteRule`

NewPathRewriteRule instantiates a new PathRewriteRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathRewriteRuleWithDefaults

`func NewPathRewriteRuleWithDefaults() *PathRewriteRule`

NewPathRewriteRuleWithDefaults instantiates a new PathRewriteRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestPattern

`func (o *PathRewriteRule) GetRequestPattern() string`

GetRequestPattern returns the RequestPattern field if non-nil, zero value otherwise.

### GetRequestPatternOk

`func (o *PathRewriteRule) GetRequestPatternOk() (*string, bool)`

GetRequestPatternOk returns a tuple with the RequestPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPattern

`func (o *PathRewriteRule) SetRequestPattern(v string)`

SetRequestPattern sets RequestPattern field to given value.


### GetMappedPath

`func (o *PathRewriteRule) GetMappedPath() string`

GetMappedPath returns the MappedPath field if non-nil, zero value otherwise.

### GetMappedPathOk

`func (o *PathRewriteRule) GetMappedPathOk() (*string, bool)`

GetMappedPathOk returns a tuple with the MappedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedPath

`func (o *PathRewriteRule) SetMappedPath(v string)`

SetMappedPath sets MappedPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


