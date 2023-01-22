# PendingPolicyCounterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyCounterStatus** | **string** | Identifies the policy counter status applicable for a specific policy counter identified by the policyCounterId. The values (e.g. valid, invalid or any other status) are not specified. The interpretation and actions related to the defined values are out of scope of 3GPP.  | 
**ActivationTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewPendingPolicyCounterStatus

`func NewPendingPolicyCounterStatus(policyCounterStatus string, activationTime time.Time, ) *PendingPolicyCounterStatus`

NewPendingPolicyCounterStatus instantiates a new PendingPolicyCounterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingPolicyCounterStatusWithDefaults

`func NewPendingPolicyCounterStatusWithDefaults() *PendingPolicyCounterStatus`

NewPendingPolicyCounterStatusWithDefaults instantiates a new PendingPolicyCounterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyCounterStatus

`func (o *PendingPolicyCounterStatus) GetPolicyCounterStatus() string`

GetPolicyCounterStatus returns the PolicyCounterStatus field if non-nil, zero value otherwise.

### GetPolicyCounterStatusOk

`func (o *PendingPolicyCounterStatus) GetPolicyCounterStatusOk() (*string, bool)`

GetPolicyCounterStatusOk returns a tuple with the PolicyCounterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCounterStatus

`func (o *PendingPolicyCounterStatus) SetPolicyCounterStatus(v string)`

SetPolicyCounterStatus sets PolicyCounterStatus field to given value.


### GetActivationTime

`func (o *PendingPolicyCounterStatus) GetActivationTime() time.Time`

GetActivationTime returns the ActivationTime field if non-nil, zero value otherwise.

### GetActivationTimeOk

`func (o *PendingPolicyCounterStatus) GetActivationTimeOk() (*time.Time, bool)`

GetActivationTimeOk returns a tuple with the ActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTime

`func (o *PendingPolicyCounterStatus) SetActivationTime(v time.Time)`

SetActivationTime sets ActivationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


