# PfdChangeReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PfdError** | [**ProblemDetails**](ProblemDetails.md) |  | 
**ApplicationId** | **[]string** |  | 

## Methods

### NewPfdChangeReport

`func NewPfdChangeReport(pfdError ProblemDetails, applicationId []string, ) *PfdChangeReport`

NewPfdChangeReport instantiates a new PfdChangeReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdChangeReportWithDefaults

`func NewPfdChangeReportWithDefaults() *PfdChangeReport`

NewPfdChangeReportWithDefaults instantiates a new PfdChangeReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPfdError

`func (o *PfdChangeReport) GetPfdError() ProblemDetails`

GetPfdError returns the PfdError field if non-nil, zero value otherwise.

### GetPfdErrorOk

`func (o *PfdChangeReport) GetPfdErrorOk() (*ProblemDetails, bool)`

GetPfdErrorOk returns a tuple with the PfdError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdError

`func (o *PfdChangeReport) SetPfdError(v ProblemDetails)`

SetPfdError sets PfdError field to given value.


### GetApplicationId

`func (o *PfdChangeReport) GetApplicationId() []string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PfdChangeReport) GetApplicationIdOk() (*[]string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PfdChangeReport) SetApplicationId(v []string)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


