# AnnounceDiscDataForOpen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProseAppId** | **string** | Contains the ProSe Application ID. | 
**ValidityTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**ProseAppCode** | Pointer to **string** | Contains the ProSe Application Code. | [optional] 
**ProseAppCodePrefix** | Pointer to **string** | Contains the Prose Application Code Prefix. | [optional] 
**ProseAppCodeSuffixPool** | Pointer to [**ProseApplicationCodeSuffixPool**](ProseApplicationCodeSuffixPool.md) |  | [optional] 
**MetaData** | Pointer to **string** | Contains the metadata. | [optional] 

## Methods

### NewAnnounceDiscDataForOpen

`func NewAnnounceDiscDataForOpen(proseAppId string, validityTime time.Time, ) *AnnounceDiscDataForOpen`

NewAnnounceDiscDataForOpen instantiates a new AnnounceDiscDataForOpen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnounceDiscDataForOpenWithDefaults

`func NewAnnounceDiscDataForOpenWithDefaults() *AnnounceDiscDataForOpen`

NewAnnounceDiscDataForOpenWithDefaults instantiates a new AnnounceDiscDataForOpen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProseAppId

`func (o *AnnounceDiscDataForOpen) GetProseAppId() string`

GetProseAppId returns the ProseAppId field if non-nil, zero value otherwise.

### GetProseAppIdOk

`func (o *AnnounceDiscDataForOpen) GetProseAppIdOk() (*string, bool)`

GetProseAppIdOk returns a tuple with the ProseAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppId

`func (o *AnnounceDiscDataForOpen) SetProseAppId(v string)`

SetProseAppId sets ProseAppId field to given value.


### GetValidityTime

`func (o *AnnounceDiscDataForOpen) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *AnnounceDiscDataForOpen) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *AnnounceDiscDataForOpen) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.


### GetProseAppCode

`func (o *AnnounceDiscDataForOpen) GetProseAppCode() string`

GetProseAppCode returns the ProseAppCode field if non-nil, zero value otherwise.

### GetProseAppCodeOk

`func (o *AnnounceDiscDataForOpen) GetProseAppCodeOk() (*string, bool)`

GetProseAppCodeOk returns a tuple with the ProseAppCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppCode

`func (o *AnnounceDiscDataForOpen) SetProseAppCode(v string)`

SetProseAppCode sets ProseAppCode field to given value.

### HasProseAppCode

`func (o *AnnounceDiscDataForOpen) HasProseAppCode() bool`

HasProseAppCode returns a boolean if a field has been set.

### GetProseAppCodePrefix

`func (o *AnnounceDiscDataForOpen) GetProseAppCodePrefix() string`

GetProseAppCodePrefix returns the ProseAppCodePrefix field if non-nil, zero value otherwise.

### GetProseAppCodePrefixOk

`func (o *AnnounceDiscDataForOpen) GetProseAppCodePrefixOk() (*string, bool)`

GetProseAppCodePrefixOk returns a tuple with the ProseAppCodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppCodePrefix

`func (o *AnnounceDiscDataForOpen) SetProseAppCodePrefix(v string)`

SetProseAppCodePrefix sets ProseAppCodePrefix field to given value.

### HasProseAppCodePrefix

`func (o *AnnounceDiscDataForOpen) HasProseAppCodePrefix() bool`

HasProseAppCodePrefix returns a boolean if a field has been set.

### GetProseAppCodeSuffixPool

`func (o *AnnounceDiscDataForOpen) GetProseAppCodeSuffixPool() ProseApplicationCodeSuffixPool`

GetProseAppCodeSuffixPool returns the ProseAppCodeSuffixPool field if non-nil, zero value otherwise.

### GetProseAppCodeSuffixPoolOk

`func (o *AnnounceDiscDataForOpen) GetProseAppCodeSuffixPoolOk() (*ProseApplicationCodeSuffixPool, bool)`

GetProseAppCodeSuffixPoolOk returns a tuple with the ProseAppCodeSuffixPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppCodeSuffixPool

`func (o *AnnounceDiscDataForOpen) SetProseAppCodeSuffixPool(v ProseApplicationCodeSuffixPool)`

SetProseAppCodeSuffixPool sets ProseAppCodeSuffixPool field to given value.

### HasProseAppCodeSuffixPool

`func (o *AnnounceDiscDataForOpen) HasProseAppCodeSuffixPool() bool`

HasProseAppCodeSuffixPool returns a boolean if a field has been set.

### GetMetaData

`func (o *AnnounceDiscDataForOpen) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AnnounceDiscDataForOpen) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AnnounceDiscDataForOpen) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AnnounceDiscDataForOpen) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


