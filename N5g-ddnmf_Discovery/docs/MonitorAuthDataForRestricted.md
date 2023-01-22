# MonitorAuthDataForRestricted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProseRestrictedCode** | **string** | Contains the ProSe Restricted Code. | 
**ValidityTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewMonitorAuthDataForRestricted

`func NewMonitorAuthDataForRestricted(proseRestrictedCode string, validityTime time.Time, ) *MonitorAuthDataForRestricted`

NewMonitorAuthDataForRestricted instantiates a new MonitorAuthDataForRestricted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorAuthDataForRestrictedWithDefaults

`func NewMonitorAuthDataForRestrictedWithDefaults() *MonitorAuthDataForRestricted`

NewMonitorAuthDataForRestrictedWithDefaults instantiates a new MonitorAuthDataForRestricted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProseRestrictedCode

`func (o *MonitorAuthDataForRestricted) GetProseRestrictedCode() string`

GetProseRestrictedCode returns the ProseRestrictedCode field if non-nil, zero value otherwise.

### GetProseRestrictedCodeOk

`func (o *MonitorAuthDataForRestricted) GetProseRestrictedCodeOk() (*string, bool)`

GetProseRestrictedCodeOk returns a tuple with the ProseRestrictedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseRestrictedCode

`func (o *MonitorAuthDataForRestricted) SetProseRestrictedCode(v string)`

SetProseRestrictedCode sets ProseRestrictedCode field to given value.


### GetValidityTime

`func (o *MonitorAuthDataForRestricted) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *MonitorAuthDataForRestricted) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *MonitorAuthDataForRestricted) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


