# ArpRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriorityLevel** | **NullableInt32** | nullable true shall not be used for this attribute. Unsigned integer indicating the ARP Priority Level (see clause 5.7.2.2 of 3GPP TS 23.501, within the range 1 to 15.Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.   | 
**PreemptCap** | [**PreemptionCapability**](PreemptionCapability.md) |  | 
**PreemptVuln** | [**PreemptionVulnerability**](PreemptionVulnerability.md) |  | 

## Methods

### NewArpRm

`func NewArpRm(priorityLevel NullableInt32, preemptCap PreemptionCapability, preemptVuln PreemptionVulnerability, ) *ArpRm`

NewArpRm instantiates a new ArpRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArpRmWithDefaults

`func NewArpRmWithDefaults() *ArpRm`

NewArpRmWithDefaults instantiates a new ArpRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriorityLevel

`func (o *ArpRm) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *ArpRm) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *ArpRm) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.


### SetPriorityLevelNil

`func (o *ArpRm) SetPriorityLevelNil(b bool)`

 SetPriorityLevelNil sets the value for PriorityLevel to be an explicit nil

### UnsetPriorityLevel
`func (o *ArpRm) UnsetPriorityLevel()`

UnsetPriorityLevel ensures that no value is present for PriorityLevel, not even an explicit nil
### GetPreemptCap

`func (o *ArpRm) GetPreemptCap() PreemptionCapability`

GetPreemptCap returns the PreemptCap field if non-nil, zero value otherwise.

### GetPreemptCapOk

`func (o *ArpRm) GetPreemptCapOk() (*PreemptionCapability, bool)`

GetPreemptCapOk returns a tuple with the PreemptCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptCap

`func (o *ArpRm) SetPreemptCap(v PreemptionCapability)`

SetPreemptCap sets PreemptCap field to given value.


### GetPreemptVuln

`func (o *ArpRm) GetPreemptVuln() PreemptionVulnerability`

GetPreemptVuln returns the PreemptVuln field if non-nil, zero value otherwise.

### GetPreemptVulnOk

`func (o *ArpRm) GetPreemptVulnOk() (*PreemptionVulnerability, bool)`

GetPreemptVulnOk returns a tuple with the PreemptVuln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptVuln

`func (o *ArpRm) SetPreemptVuln(v PreemptionVulnerability)`

SetPreemptVuln sets PreemptVuln field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


