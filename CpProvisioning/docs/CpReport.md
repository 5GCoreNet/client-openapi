# CpReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetIds** | Pointer to **[]string** | Identifies the CP set identifier(s) which CP parameter(s) are not added or modified successfully | [optional] 
**FailureCode** | [**CpFailureCode**](CpFailureCode.md) |  | 

## Methods

### NewCpReport

`func NewCpReport(failureCode CpFailureCode, ) *CpReport`

NewCpReport instantiates a new CpReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpReportWithDefaults

`func NewCpReportWithDefaults() *CpReport`

NewCpReportWithDefaults instantiates a new CpReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetIds

`func (o *CpReport) GetSetIds() []string`

GetSetIds returns the SetIds field if non-nil, zero value otherwise.

### GetSetIdsOk

`func (o *CpReport) GetSetIdsOk() (*[]string, bool)`

GetSetIdsOk returns a tuple with the SetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetIds

`func (o *CpReport) SetSetIds(v []string)`

SetSetIds sets SetIds field to given value.

### HasSetIds

`func (o *CpReport) HasSetIds() bool`

HasSetIds returns a boolean if a field has been set.

### GetFailureCode

`func (o *CpReport) GetFailureCode() CpFailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *CpReport) GetFailureCodeOk() (*CpFailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *CpReport) SetFailureCode(v CpFailureCode)`

SetFailureCode sets FailureCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


