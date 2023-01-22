# RacsFailureReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RacsIds** | **[]string** | Identifies the RACS ID(s) for which the RACS data are not provisioned successfully. | 
**FailureCode** | [**RacsFailureCode**](RacsFailureCode.md) |  | 

## Methods

### NewRacsFailureReport

`func NewRacsFailureReport(racsIds []string, failureCode RacsFailureCode, ) *RacsFailureReport`

NewRacsFailureReport instantiates a new RacsFailureReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacsFailureReportWithDefaults

`func NewRacsFailureReportWithDefaults() *RacsFailureReport`

NewRacsFailureReportWithDefaults instantiates a new RacsFailureReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRacsIds

`func (o *RacsFailureReport) GetRacsIds() []string`

GetRacsIds returns the RacsIds field if non-nil, zero value otherwise.

### GetRacsIdsOk

`func (o *RacsFailureReport) GetRacsIdsOk() (*[]string, bool)`

GetRacsIdsOk returns a tuple with the RacsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsIds

`func (o *RacsFailureReport) SetRacsIds(v []string)`

SetRacsIds sets RacsIds field to given value.


### GetFailureCode

`func (o *RacsFailureReport) GetFailureCode() RacsFailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *RacsFailureReport) GetFailureCodeOk() (*RacsFailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *RacsFailureReport) SetFailureCode(v RacsFailureCode)`

SetFailureCode sets FailureCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


