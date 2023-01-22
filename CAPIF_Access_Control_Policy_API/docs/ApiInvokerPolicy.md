# ApiInvokerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInvokerId** | **string** | API invoker ID assigned by the CAPIF core function | 
**AllowedTotalInvocations** | Pointer to **int32** | Total number of invocations allowed on the service API by the API invoker. | [optional] 
**AllowedInvocationsPerSecond** | Pointer to **int32** | Invocations per second allowed on the service API by the API invoker. | [optional] 
**AllowedInvocationTimeRangeList** | Pointer to [**[]TimeRangeList**](TimeRangeList.md) | The time ranges during which the invocations are allowed on the service API by the API invoker.  | [optional] 

## Methods

### NewApiInvokerPolicy

`func NewApiInvokerPolicy(apiInvokerId string, ) *ApiInvokerPolicy`

NewApiInvokerPolicy instantiates a new ApiInvokerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInvokerPolicyWithDefaults

`func NewApiInvokerPolicyWithDefaults() *ApiInvokerPolicy`

NewApiInvokerPolicyWithDefaults instantiates a new ApiInvokerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInvokerId

`func (o *ApiInvokerPolicy) GetApiInvokerId() string`

GetApiInvokerId returns the ApiInvokerId field if non-nil, zero value otherwise.

### GetApiInvokerIdOk

`func (o *ApiInvokerPolicy) GetApiInvokerIdOk() (*string, bool)`

GetApiInvokerIdOk returns a tuple with the ApiInvokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerId

`func (o *ApiInvokerPolicy) SetApiInvokerId(v string)`

SetApiInvokerId sets ApiInvokerId field to given value.


### GetAllowedTotalInvocations

`func (o *ApiInvokerPolicy) GetAllowedTotalInvocations() int32`

GetAllowedTotalInvocations returns the AllowedTotalInvocations field if non-nil, zero value otherwise.

### GetAllowedTotalInvocationsOk

`func (o *ApiInvokerPolicy) GetAllowedTotalInvocationsOk() (*int32, bool)`

GetAllowedTotalInvocationsOk returns a tuple with the AllowedTotalInvocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTotalInvocations

`func (o *ApiInvokerPolicy) SetAllowedTotalInvocations(v int32)`

SetAllowedTotalInvocations sets AllowedTotalInvocations field to given value.

### HasAllowedTotalInvocations

`func (o *ApiInvokerPolicy) HasAllowedTotalInvocations() bool`

HasAllowedTotalInvocations returns a boolean if a field has been set.

### GetAllowedInvocationsPerSecond

`func (o *ApiInvokerPolicy) GetAllowedInvocationsPerSecond() int32`

GetAllowedInvocationsPerSecond returns the AllowedInvocationsPerSecond field if non-nil, zero value otherwise.

### GetAllowedInvocationsPerSecondOk

`func (o *ApiInvokerPolicy) GetAllowedInvocationsPerSecondOk() (*int32, bool)`

GetAllowedInvocationsPerSecondOk returns a tuple with the AllowedInvocationsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInvocationsPerSecond

`func (o *ApiInvokerPolicy) SetAllowedInvocationsPerSecond(v int32)`

SetAllowedInvocationsPerSecond sets AllowedInvocationsPerSecond field to given value.

### HasAllowedInvocationsPerSecond

`func (o *ApiInvokerPolicy) HasAllowedInvocationsPerSecond() bool`

HasAllowedInvocationsPerSecond returns a boolean if a field has been set.

### GetAllowedInvocationTimeRangeList

`func (o *ApiInvokerPolicy) GetAllowedInvocationTimeRangeList() []TimeRangeList`

GetAllowedInvocationTimeRangeList returns the AllowedInvocationTimeRangeList field if non-nil, zero value otherwise.

### GetAllowedInvocationTimeRangeListOk

`func (o *ApiInvokerPolicy) GetAllowedInvocationTimeRangeListOk() (*[]TimeRangeList, bool)`

GetAllowedInvocationTimeRangeListOk returns a tuple with the AllowedInvocationTimeRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInvocationTimeRangeList

`func (o *ApiInvokerPolicy) SetAllowedInvocationTimeRangeList(v []TimeRangeList)`

SetAllowedInvocationTimeRangeList sets AllowedInvocationTimeRangeList field to given value.

### HasAllowedInvocationTimeRangeList

`func (o *ApiInvokerPolicy) HasAllowedInvocationTimeRangeList() bool`

HasAllowedInvocationTimeRangeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


