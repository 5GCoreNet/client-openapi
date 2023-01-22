# UAVAuthFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ProblemDetails**](ProblemDetails.md) |  | 
**UasResourceRelease** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUAVAuthFailure

`func NewUAVAuthFailure(error_ ProblemDetails, ) *UAVAuthFailure`

NewUAVAuthFailure instantiates a new UAVAuthFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUAVAuthFailureWithDefaults

`func NewUAVAuthFailureWithDefaults() *UAVAuthFailure`

NewUAVAuthFailureWithDefaults instantiates a new UAVAuthFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UAVAuthFailure) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UAVAuthFailure) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UAVAuthFailure) SetError(v ProblemDetails)`

SetError sets Error field to given value.


### GetUasResourceRelease

`func (o *UAVAuthFailure) GetUasResourceRelease() bool`

GetUasResourceRelease returns the UasResourceRelease field if non-nil, zero value otherwise.

### GetUasResourceReleaseOk

`func (o *UAVAuthFailure) GetUasResourceReleaseOk() (*bool, bool)`

GetUasResourceReleaseOk returns a tuple with the UasResourceRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUasResourceRelease

`func (o *UAVAuthFailure) SetUasResourceRelease(v bool)`

SetUasResourceRelease sets UasResourceRelease field to given value.

### HasUasResourceRelease

`func (o *UAVAuthFailure) HasUasResourceRelease() bool`

HasUasResourceRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


