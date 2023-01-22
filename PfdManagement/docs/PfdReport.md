# PfdReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAppIds** | **[]string** | Identifies the external application identifier(s) which PFD(s) are not added or modified successfully | 
**FailureCode** | [**FailureCode**](FailureCode.md) |  | 
**CachingTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**LocationArea** | Pointer to [**UserPlaneLocationArea**](UserPlaneLocationArea.md) |  | [optional] 

## Methods

### NewPfdReport

`func NewPfdReport(externalAppIds []string, failureCode FailureCode, ) *PfdReport`

NewPfdReport instantiates a new PfdReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdReportWithDefaults

`func NewPfdReportWithDefaults() *PfdReport`

NewPfdReportWithDefaults instantiates a new PfdReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAppIds

`func (o *PfdReport) GetExternalAppIds() []string`

GetExternalAppIds returns the ExternalAppIds field if non-nil, zero value otherwise.

### GetExternalAppIdsOk

`func (o *PfdReport) GetExternalAppIdsOk() (*[]string, bool)`

GetExternalAppIdsOk returns a tuple with the ExternalAppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAppIds

`func (o *PfdReport) SetExternalAppIds(v []string)`

SetExternalAppIds sets ExternalAppIds field to given value.


### GetFailureCode

`func (o *PfdReport) GetFailureCode() FailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PfdReport) GetFailureCodeOk() (*FailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PfdReport) SetFailureCode(v FailureCode)`

SetFailureCode sets FailureCode field to given value.


### GetCachingTime

`func (o *PfdReport) GetCachingTime() int32`

GetCachingTime returns the CachingTime field if non-nil, zero value otherwise.

### GetCachingTimeOk

`func (o *PfdReport) GetCachingTimeOk() (*int32, bool)`

GetCachingTimeOk returns a tuple with the CachingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTime

`func (o *PfdReport) SetCachingTime(v int32)`

SetCachingTime sets CachingTime field to given value.

### HasCachingTime

`func (o *PfdReport) HasCachingTime() bool`

HasCachingTime returns a boolean if a field has been set.

### GetLocationArea

`func (o *PfdReport) GetLocationArea() UserPlaneLocationArea`

GetLocationArea returns the LocationArea field if non-nil, zero value otherwise.

### GetLocationAreaOk

`func (o *PfdReport) GetLocationAreaOk() (*UserPlaneLocationArea, bool)`

GetLocationAreaOk returns a tuple with the LocationArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea

`func (o *PfdReport) SetLocationArea(v UserPlaneLocationArea)`

SetLocationArea sets LocationArea field to given value.

### HasLocationArea

`func (o *PfdReport) HasLocationArea() bool`

HasLocationArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


