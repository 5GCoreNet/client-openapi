# FailureEventInfoForMLModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NwdafEvent**](NwdafEvent.md) |  | 
**FailureCode** | [**FailureCode**](FailureCode.md) |  | 

## Methods

### NewFailureEventInfoForMLModel

`func NewFailureEventInfoForMLModel(event NwdafEvent, failureCode FailureCode, ) *FailureEventInfoForMLModel`

NewFailureEventInfoForMLModel instantiates a new FailureEventInfoForMLModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureEventInfoForMLModelWithDefaults

`func NewFailureEventInfoForMLModelWithDefaults() *FailureEventInfoForMLModel`

NewFailureEventInfoForMLModelWithDefaults instantiates a new FailureEventInfoForMLModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *FailureEventInfoForMLModel) GetEvent() NwdafEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *FailureEventInfoForMLModel) GetEventOk() (*NwdafEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *FailureEventInfoForMLModel) SetEvent(v NwdafEvent)`

SetEvent sets Event field to given value.


### GetFailureCode

`func (o *FailureEventInfoForMLModel) GetFailureCode() FailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *FailureEventInfoForMLModel) GetFailureCodeOk() (*FailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *FailureEventInfoForMLModel) SetFailureCode(v FailureCode)`

SetFailureCode sets FailureCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


