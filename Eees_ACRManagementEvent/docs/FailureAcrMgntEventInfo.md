# FailureAcrMgntEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AcrMgntEvent**](AcrMgntEvent.md) |  | 
**FailureCode** | [**AcrMgntEventFailureCode**](AcrMgntEventFailureCode.md) |  | 

## Methods

### NewFailureAcrMgntEventInfo

`func NewFailureAcrMgntEventInfo(event AcrMgntEvent, failureCode AcrMgntEventFailureCode, ) *FailureAcrMgntEventInfo`

NewFailureAcrMgntEventInfo instantiates a new FailureAcrMgntEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureAcrMgntEventInfoWithDefaults

`func NewFailureAcrMgntEventInfoWithDefaults() *FailureAcrMgntEventInfo`

NewFailureAcrMgntEventInfoWithDefaults instantiates a new FailureAcrMgntEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *FailureAcrMgntEventInfo) GetEvent() AcrMgntEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *FailureAcrMgntEventInfo) GetEventOk() (*AcrMgntEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *FailureAcrMgntEventInfo) SetEvent(v AcrMgntEvent)`

SetEvent sets Event field to given value.


### GetFailureCode

`func (o *FailureAcrMgntEventInfo) GetFailureCode() AcrMgntEventFailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *FailureAcrMgntEventInfo) GetFailureCodeOk() (*AcrMgntEventFailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *FailureAcrMgntEventInfo) SetFailureCode(v AcrMgntEventFailureCode)`

SetFailureCode sets FailureCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


