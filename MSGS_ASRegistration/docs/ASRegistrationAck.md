# ASRegistrationAck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsSvcId** | **string** |  | 
**Result** | [**ProblemDetails**](ProblemDetails.md) |  | 

## Methods

### NewASRegistrationAck

`func NewASRegistrationAck(asSvcId string, result ProblemDetails, ) *ASRegistrationAck`

NewASRegistrationAck instantiates a new ASRegistrationAck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASRegistrationAckWithDefaults

`func NewASRegistrationAckWithDefaults() *ASRegistrationAck`

NewASRegistrationAckWithDefaults instantiates a new ASRegistrationAck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsSvcId

`func (o *ASRegistrationAck) GetAsSvcId() string`

GetAsSvcId returns the AsSvcId field if non-nil, zero value otherwise.

### GetAsSvcIdOk

`func (o *ASRegistrationAck) GetAsSvcIdOk() (*string, bool)`

GetAsSvcIdOk returns a tuple with the AsSvcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSvcId

`func (o *ASRegistrationAck) SetAsSvcId(v string)`

SetAsSvcId sets AsSvcId field to given value.


### GetResult

`func (o *ASRegistrationAck) GetResult() ProblemDetails`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ASRegistrationAck) GetResultOk() (*ProblemDetails, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ASRegistrationAck) SetResult(v ProblemDetails)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


