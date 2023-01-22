# PfdDataForAppExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | String providing an application identifier. | 
**Pfds** | [**[]PfdContent**](PfdContent.md) |  | 
**CachingTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**AllowedDelay** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewPfdDataForAppExt

`func NewPfdDataForAppExt(applicationId string, pfds []PfdContent, ) *PfdDataForAppExt`

NewPfdDataForAppExt instantiates a new PfdDataForAppExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdDataForAppExtWithDefaults

`func NewPfdDataForAppExtWithDefaults() *PfdDataForAppExt`

NewPfdDataForAppExtWithDefaults instantiates a new PfdDataForAppExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PfdDataForAppExt) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PfdDataForAppExt) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PfdDataForAppExt) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetPfds

`func (o *PfdDataForAppExt) GetPfds() []PfdContent`

GetPfds returns the Pfds field if non-nil, zero value otherwise.

### GetPfdsOk

`func (o *PfdDataForAppExt) GetPfdsOk() (*[]PfdContent, bool)`

GetPfdsOk returns a tuple with the Pfds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfds

`func (o *PfdDataForAppExt) SetPfds(v []PfdContent)`

SetPfds sets Pfds field to given value.


### GetCachingTime

`func (o *PfdDataForAppExt) GetCachingTime() time.Time`

GetCachingTime returns the CachingTime field if non-nil, zero value otherwise.

### GetCachingTimeOk

`func (o *PfdDataForAppExt) GetCachingTimeOk() (*time.Time, bool)`

GetCachingTimeOk returns a tuple with the CachingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTime

`func (o *PfdDataForAppExt) SetCachingTime(v time.Time)`

SetCachingTime sets CachingTime field to given value.

### HasCachingTime

`func (o *PfdDataForAppExt) HasCachingTime() bool`

HasCachingTime returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PfdDataForAppExt) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PfdDataForAppExt) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PfdDataForAppExt) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *PfdDataForAppExt) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResetIds

`func (o *PfdDataForAppExt) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *PfdDataForAppExt) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *PfdDataForAppExt) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *PfdDataForAppExt) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetAllowedDelay

`func (o *PfdDataForAppExt) GetAllowedDelay() int32`

GetAllowedDelay returns the AllowedDelay field if non-nil, zero value otherwise.

### GetAllowedDelayOk

`func (o *PfdDataForAppExt) GetAllowedDelayOk() (*int32, bool)`

GetAllowedDelayOk returns a tuple with the AllowedDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDelay

`func (o *PfdDataForAppExt) SetAllowedDelay(v int32)`

SetAllowedDelay sets AllowedDelay field to given value.

### HasAllowedDelay

`func (o *PfdDataForAppExt) HasAllowedDelay() bool`

HasAllowedDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


