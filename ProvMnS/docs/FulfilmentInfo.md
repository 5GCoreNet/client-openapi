# FulfilmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfilStatus** | Pointer to [**FulfilStatus**](FulfilStatus.md) |  | [optional] 
**NotFullfilledState** | Pointer to [**NotFulfilledState**](NotFulfilledState.md) |  | [optional] 
**NotFulfilledReasons** | Pointer to **string** | -&gt; An attribute which is used when FulfilmentInfo is implemented for IntentFulfilmentInfo | [optional] 

## Methods

### NewFulfilmentInfo

`func NewFulfilmentInfo() *FulfilmentInfo`

NewFulfilmentInfo instantiates a new FulfilmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfilmentInfoWithDefaults

`func NewFulfilmentInfoWithDefaults() *FulfilmentInfo`

NewFulfilmentInfoWithDefaults instantiates a new FulfilmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfilStatus

`func (o *FulfilmentInfo) GetFulfilStatus() FulfilStatus`

GetFulfilStatus returns the FulfilStatus field if non-nil, zero value otherwise.

### GetFulfilStatusOk

`func (o *FulfilmentInfo) GetFulfilStatusOk() (*FulfilStatus, bool)`

GetFulfilStatusOk returns a tuple with the FulfilStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilStatus

`func (o *FulfilmentInfo) SetFulfilStatus(v FulfilStatus)`

SetFulfilStatus sets FulfilStatus field to given value.

### HasFulfilStatus

`func (o *FulfilmentInfo) HasFulfilStatus() bool`

HasFulfilStatus returns a boolean if a field has been set.

### GetNotFullfilledState

`func (o *FulfilmentInfo) GetNotFullfilledState() NotFulfilledState`

GetNotFullfilledState returns the NotFullfilledState field if non-nil, zero value otherwise.

### GetNotFullfilledStateOk

`func (o *FulfilmentInfo) GetNotFullfilledStateOk() (*NotFulfilledState, bool)`

GetNotFullfilledStateOk returns a tuple with the NotFullfilledState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFullfilledState

`func (o *FulfilmentInfo) SetNotFullfilledState(v NotFulfilledState)`

SetNotFullfilledState sets NotFullfilledState field to given value.

### HasNotFullfilledState

`func (o *FulfilmentInfo) HasNotFullfilledState() bool`

HasNotFullfilledState returns a boolean if a field has been set.

### GetNotFulfilledReasons

`func (o *FulfilmentInfo) GetNotFulfilledReasons() string`

GetNotFulfilledReasons returns the NotFulfilledReasons field if non-nil, zero value otherwise.

### GetNotFulfilledReasonsOk

`func (o *FulfilmentInfo) GetNotFulfilledReasonsOk() (*string, bool)`

GetNotFulfilledReasonsOk returns a tuple with the NotFulfilledReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFulfilledReasons

`func (o *FulfilmentInfo) SetNotFulfilledReasons(v string)`

SetNotFulfilledReasons sets NotFulfilledReasons field to given value.

### HasNotFulfilledReasons

`func (o *FulfilmentInfo) HasNotFulfilledReasons() bool`

HasNotFulfilledReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


