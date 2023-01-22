# PfdDataForApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | String providing an application identifier. | 
**Pfds** | Pointer to [**[]PfdContent**](PfdContent.md) |  | [optional] 
**CachingTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**CachingTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**PfdTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PartialFlag** | Pointer to **bool** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewPfdDataForApp

`func NewPfdDataForApp(applicationId string, ) *PfdDataForApp`

NewPfdDataForApp instantiates a new PfdDataForApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdDataForAppWithDefaults

`func NewPfdDataForAppWithDefaults() *PfdDataForApp`

NewPfdDataForAppWithDefaults instantiates a new PfdDataForApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PfdDataForApp) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PfdDataForApp) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PfdDataForApp) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetPfds

`func (o *PfdDataForApp) GetPfds() []PfdContent`

GetPfds returns the Pfds field if non-nil, zero value otherwise.

### GetPfdsOk

`func (o *PfdDataForApp) GetPfdsOk() (*[]PfdContent, bool)`

GetPfdsOk returns a tuple with the Pfds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfds

`func (o *PfdDataForApp) SetPfds(v []PfdContent)`

SetPfds sets Pfds field to given value.

### HasPfds

`func (o *PfdDataForApp) HasPfds() bool`

HasPfds returns a boolean if a field has been set.

### GetCachingTime

`func (o *PfdDataForApp) GetCachingTime() time.Time`

GetCachingTime returns the CachingTime field if non-nil, zero value otherwise.

### GetCachingTimeOk

`func (o *PfdDataForApp) GetCachingTimeOk() (*time.Time, bool)`

GetCachingTimeOk returns a tuple with the CachingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTime

`func (o *PfdDataForApp) SetCachingTime(v time.Time)`

SetCachingTime sets CachingTime field to given value.

### HasCachingTime

`func (o *PfdDataForApp) HasCachingTime() bool`

HasCachingTime returns a boolean if a field has been set.

### GetCachingTimer

`func (o *PfdDataForApp) GetCachingTimer() int32`

GetCachingTimer returns the CachingTimer field if non-nil, zero value otherwise.

### GetCachingTimerOk

`func (o *PfdDataForApp) GetCachingTimerOk() (*int32, bool)`

GetCachingTimerOk returns a tuple with the CachingTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTimer

`func (o *PfdDataForApp) SetCachingTimer(v int32)`

SetCachingTimer sets CachingTimer field to given value.

### HasCachingTimer

`func (o *PfdDataForApp) HasCachingTimer() bool`

HasCachingTimer returns a boolean if a field has been set.

### GetPfdTimestamp

`func (o *PfdDataForApp) GetPfdTimestamp() time.Time`

GetPfdTimestamp returns the PfdTimestamp field if non-nil, zero value otherwise.

### GetPfdTimestampOk

`func (o *PfdDataForApp) GetPfdTimestampOk() (*time.Time, bool)`

GetPfdTimestampOk returns a tuple with the PfdTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdTimestamp

`func (o *PfdDataForApp) SetPfdTimestamp(v time.Time)`

SetPfdTimestamp sets PfdTimestamp field to given value.

### HasPfdTimestamp

`func (o *PfdDataForApp) HasPfdTimestamp() bool`

HasPfdTimestamp returns a boolean if a field has been set.

### GetPartialFlag

`func (o *PfdDataForApp) GetPartialFlag() bool`

GetPartialFlag returns the PartialFlag field if non-nil, zero value otherwise.

### GetPartialFlagOk

`func (o *PfdDataForApp) GetPartialFlagOk() (*bool, bool)`

GetPartialFlagOk returns a tuple with the PartialFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFlag

`func (o *PfdDataForApp) SetPartialFlag(v bool)`

SetPartialFlag sets PartialFlag field to given value.

### HasPartialFlag

`func (o *PfdDataForApp) HasPartialFlag() bool`

HasPartialFlag returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PfdDataForApp) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PfdDataForApp) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PfdDataForApp) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PfdDataForApp) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


