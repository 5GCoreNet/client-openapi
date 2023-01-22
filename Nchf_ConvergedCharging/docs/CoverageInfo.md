# CoverageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverageStatus** | Pointer to **bool** |  | [optional] 
**ChangeTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**LocationInfo** | Pointer to [**[]UserLocation**](UserLocation.md) |  | [optional] 

## Methods

### NewCoverageInfo

`func NewCoverageInfo() *CoverageInfo`

NewCoverageInfo instantiates a new CoverageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoverageInfoWithDefaults

`func NewCoverageInfoWithDefaults() *CoverageInfo`

NewCoverageInfoWithDefaults instantiates a new CoverageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverageStatus

`func (o *CoverageInfo) GetCoverageStatus() bool`

GetCoverageStatus returns the CoverageStatus field if non-nil, zero value otherwise.

### GetCoverageStatusOk

`func (o *CoverageInfo) GetCoverageStatusOk() (*bool, bool)`

GetCoverageStatusOk returns a tuple with the CoverageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageStatus

`func (o *CoverageInfo) SetCoverageStatus(v bool)`

SetCoverageStatus sets CoverageStatus field to given value.

### HasCoverageStatus

`func (o *CoverageInfo) HasCoverageStatus() bool`

HasCoverageStatus returns a boolean if a field has been set.

### GetChangeTime

`func (o *CoverageInfo) GetChangeTime() time.Time`

GetChangeTime returns the ChangeTime field if non-nil, zero value otherwise.

### GetChangeTimeOk

`func (o *CoverageInfo) GetChangeTimeOk() (*time.Time, bool)`

GetChangeTimeOk returns a tuple with the ChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTime

`func (o *CoverageInfo) SetChangeTime(v time.Time)`

SetChangeTime sets ChangeTime field to given value.

### HasChangeTime

`func (o *CoverageInfo) HasChangeTime() bool`

HasChangeTime returns a boolean if a field has been set.

### GetLocationInfo

`func (o *CoverageInfo) GetLocationInfo() []UserLocation`

GetLocationInfo returns the LocationInfo field if non-nil, zero value otherwise.

### GetLocationInfoOk

`func (o *CoverageInfo) GetLocationInfoOk() (*[]UserLocation, bool)`

GetLocationInfoOk returns a tuple with the LocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInfo

`func (o *CoverageInfo) SetLocationInfo(v []UserLocation)`

SetLocationInfo sets LocationInfo field to given value.

### HasLocationInfo

`func (o *CoverageInfo) HasLocationInfo() bool`

HasLocationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


