# PolicyCounterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyCounterId** | **string** | Identifies a policy counter. | 
**CurrentStatus** | **string** | Identifies the policy counter status applicable for a specific policy counter identified by the policyCounterId. The values (e.g. valid, invalid or any other status) are not specified. The interpretation and actions related to the defined values are out of scope of 3GPP.  | 
**PenPolCounterStatuses** | Pointer to [**[]PendingPolicyCounterStatus**](PendingPolicyCounterStatus.md) | Provides the pending policy counter status. | [optional] 

## Methods

### NewPolicyCounterInfo

`func NewPolicyCounterInfo(policyCounterId string, currentStatus string, ) *PolicyCounterInfo`

NewPolicyCounterInfo instantiates a new PolicyCounterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCounterInfoWithDefaults

`func NewPolicyCounterInfoWithDefaults() *PolicyCounterInfo`

NewPolicyCounterInfoWithDefaults instantiates a new PolicyCounterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyCounterId

`func (o *PolicyCounterInfo) GetPolicyCounterId() string`

GetPolicyCounterId returns the PolicyCounterId field if non-nil, zero value otherwise.

### GetPolicyCounterIdOk

`func (o *PolicyCounterInfo) GetPolicyCounterIdOk() (*string, bool)`

GetPolicyCounterIdOk returns a tuple with the PolicyCounterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCounterId

`func (o *PolicyCounterInfo) SetPolicyCounterId(v string)`

SetPolicyCounterId sets PolicyCounterId field to given value.


### GetCurrentStatus

`func (o *PolicyCounterInfo) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *PolicyCounterInfo) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *PolicyCounterInfo) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.


### GetPenPolCounterStatuses

`func (o *PolicyCounterInfo) GetPenPolCounterStatuses() []PendingPolicyCounterStatus`

GetPenPolCounterStatuses returns the PenPolCounterStatuses field if non-nil, zero value otherwise.

### GetPenPolCounterStatusesOk

`func (o *PolicyCounterInfo) GetPenPolCounterStatusesOk() (*[]PendingPolicyCounterStatus, bool)`

GetPenPolCounterStatusesOk returns a tuple with the PenPolCounterStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenPolCounterStatuses

`func (o *PolicyCounterInfo) SetPenPolCounterStatuses(v []PendingPolicyCounterStatus)`

SetPenPolCounterStatuses sets PenPolCounterStatuses field to given value.

### HasPenPolCounterStatuses

`func (o *PolicyCounterInfo) HasPenPolCounterStatuses() bool`

HasPenPolCounterStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


