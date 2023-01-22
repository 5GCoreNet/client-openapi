# AnnounceDiscDataForRestricted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rpauid** | **string** | Contains the RPAUID. | 
**AppId** | **string** | Contains the Application ID. | 
**ValidityTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**ProseRestrictedCode** | Pointer to **string** | Contains the ProSe Restricted Code. | [optional] 
**ProseRestrictedPrefix** | Pointer to **string** | Contains the ProSe Restricted Code Prefix. | [optional] 
**CodeSuffixPool** | Pointer to [**RestrictedCodeSuffixPool**](RestrictedCodeSuffixPool.md) |  | [optional] 

## Methods

### NewAnnounceDiscDataForRestricted

`func NewAnnounceDiscDataForRestricted(rpauid string, appId string, validityTime time.Time, ) *AnnounceDiscDataForRestricted`

NewAnnounceDiscDataForRestricted instantiates a new AnnounceDiscDataForRestricted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnounceDiscDataForRestrictedWithDefaults

`func NewAnnounceDiscDataForRestrictedWithDefaults() *AnnounceDiscDataForRestricted`

NewAnnounceDiscDataForRestrictedWithDefaults instantiates a new AnnounceDiscDataForRestricted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRpauid

`func (o *AnnounceDiscDataForRestricted) GetRpauid() string`

GetRpauid returns the Rpauid field if non-nil, zero value otherwise.

### GetRpauidOk

`func (o *AnnounceDiscDataForRestricted) GetRpauidOk() (*string, bool)`

GetRpauidOk returns a tuple with the Rpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpauid

`func (o *AnnounceDiscDataForRestricted) SetRpauid(v string)`

SetRpauid sets Rpauid field to given value.


### GetAppId

`func (o *AnnounceDiscDataForRestricted) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AnnounceDiscDataForRestricted) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AnnounceDiscDataForRestricted) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetValidityTime

`func (o *AnnounceDiscDataForRestricted) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *AnnounceDiscDataForRestricted) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *AnnounceDiscDataForRestricted) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.


### GetProseRestrictedCode

`func (o *AnnounceDiscDataForRestricted) GetProseRestrictedCode() string`

GetProseRestrictedCode returns the ProseRestrictedCode field if non-nil, zero value otherwise.

### GetProseRestrictedCodeOk

`func (o *AnnounceDiscDataForRestricted) GetProseRestrictedCodeOk() (*string, bool)`

GetProseRestrictedCodeOk returns a tuple with the ProseRestrictedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseRestrictedCode

`func (o *AnnounceDiscDataForRestricted) SetProseRestrictedCode(v string)`

SetProseRestrictedCode sets ProseRestrictedCode field to given value.

### HasProseRestrictedCode

`func (o *AnnounceDiscDataForRestricted) HasProseRestrictedCode() bool`

HasProseRestrictedCode returns a boolean if a field has been set.

### GetProseRestrictedPrefix

`func (o *AnnounceDiscDataForRestricted) GetProseRestrictedPrefix() string`

GetProseRestrictedPrefix returns the ProseRestrictedPrefix field if non-nil, zero value otherwise.

### GetProseRestrictedPrefixOk

`func (o *AnnounceDiscDataForRestricted) GetProseRestrictedPrefixOk() (*string, bool)`

GetProseRestrictedPrefixOk returns a tuple with the ProseRestrictedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseRestrictedPrefix

`func (o *AnnounceDiscDataForRestricted) SetProseRestrictedPrefix(v string)`

SetProseRestrictedPrefix sets ProseRestrictedPrefix field to given value.

### HasProseRestrictedPrefix

`func (o *AnnounceDiscDataForRestricted) HasProseRestrictedPrefix() bool`

HasProseRestrictedPrefix returns a boolean if a field has been set.

### GetCodeSuffixPool

`func (o *AnnounceDiscDataForRestricted) GetCodeSuffixPool() RestrictedCodeSuffixPool`

GetCodeSuffixPool returns the CodeSuffixPool field if non-nil, zero value otherwise.

### GetCodeSuffixPoolOk

`func (o *AnnounceDiscDataForRestricted) GetCodeSuffixPoolOk() (*RestrictedCodeSuffixPool, bool)`

GetCodeSuffixPoolOk returns a tuple with the CodeSuffixPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSuffixPool

`func (o *AnnounceDiscDataForRestricted) SetCodeSuffixPool(v RestrictedCodeSuffixPool)`

SetCodeSuffixPool sets CodeSuffixPool field to given value.

### HasCodeSuffixPool

`func (o *AnnounceDiscDataForRestricted) HasCodeSuffixPool() bool`

HasCodeSuffixPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


