# NiddDownlinkDataDeliveryFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProblemDetail** | [**ProblemDetails**](ProblemDetails.md) |  | 
**RequestedRetransmissionTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewNiddDownlinkDataDeliveryFailure

`func NewNiddDownlinkDataDeliveryFailure(problemDetail ProblemDetails, ) *NiddDownlinkDataDeliveryFailure`

NewNiddDownlinkDataDeliveryFailure instantiates a new NiddDownlinkDataDeliveryFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddDownlinkDataDeliveryFailureWithDefaults

`func NewNiddDownlinkDataDeliveryFailureWithDefaults() *NiddDownlinkDataDeliveryFailure`

NewNiddDownlinkDataDeliveryFailureWithDefaults instantiates a new NiddDownlinkDataDeliveryFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblemDetail

`func (o *NiddDownlinkDataDeliveryFailure) GetProblemDetail() ProblemDetails`

GetProblemDetail returns the ProblemDetail field if non-nil, zero value otherwise.

### GetProblemDetailOk

`func (o *NiddDownlinkDataDeliveryFailure) GetProblemDetailOk() (*ProblemDetails, bool)`

GetProblemDetailOk returns a tuple with the ProblemDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDetail

`func (o *NiddDownlinkDataDeliveryFailure) SetProblemDetail(v ProblemDetails)`

SetProblemDetail sets ProblemDetail field to given value.


### GetRequestedRetransmissionTime

`func (o *NiddDownlinkDataDeliveryFailure) GetRequestedRetransmissionTime() time.Time`

GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field if non-nil, zero value otherwise.

### GetRequestedRetransmissionTimeOk

`func (o *NiddDownlinkDataDeliveryFailure) GetRequestedRetransmissionTimeOk() (*time.Time, bool)`

GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRetransmissionTime

`func (o *NiddDownlinkDataDeliveryFailure) SetRequestedRetransmissionTime(v time.Time)`

SetRequestedRetransmissionTime sets RequestedRetransmissionTime field to given value.

### HasRequestedRetransmissionTime

`func (o *NiddDownlinkDataDeliveryFailure) HasRequestedRetransmissionTime() bool`

HasRequestedRetransmissionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


