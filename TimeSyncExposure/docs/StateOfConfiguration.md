# StateOfConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateOfNwtt** | Pointer to **bool** | When the PTP port state is Leader, Follower or Passive, it is included and set to true to indicate the state of configuration for NW-TT port is active; when PTP port state is in any other case, it is included and set to false to indicate the state of configuration for NW-TT port is inactive. Default value is false.  | [optional] 
**StateOfDstts** | Pointer to [**[]StateOfDstt**](StateOfDstt.md) | Contains the PTP port states of the DS-TT(s).  | [optional] 

## Methods

### NewStateOfConfiguration

`func NewStateOfConfiguration() *StateOfConfiguration`

NewStateOfConfiguration instantiates a new StateOfConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateOfConfigurationWithDefaults

`func NewStateOfConfigurationWithDefaults() *StateOfConfiguration`

NewStateOfConfigurationWithDefaults instantiates a new StateOfConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateOfNwtt

`func (o *StateOfConfiguration) GetStateOfNwtt() bool`

GetStateOfNwtt returns the StateOfNwtt field if non-nil, zero value otherwise.

### GetStateOfNwttOk

`func (o *StateOfConfiguration) GetStateOfNwttOk() (*bool, bool)`

GetStateOfNwttOk returns a tuple with the StateOfNwtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfNwtt

`func (o *StateOfConfiguration) SetStateOfNwtt(v bool)`

SetStateOfNwtt sets StateOfNwtt field to given value.

### HasStateOfNwtt

`func (o *StateOfConfiguration) HasStateOfNwtt() bool`

HasStateOfNwtt returns a boolean if a field has been set.

### GetStateOfDstts

`func (o *StateOfConfiguration) GetStateOfDstts() []StateOfDstt`

GetStateOfDstts returns the StateOfDstts field if non-nil, zero value otherwise.

### GetStateOfDsttsOk

`func (o *StateOfConfiguration) GetStateOfDsttsOk() (*[]StateOfDstt, bool)`

GetStateOfDsttsOk returns a tuple with the StateOfDstts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfDstts

`func (o *StateOfConfiguration) SetStateOfDstts(v []StateOfDstt)`

SetStateOfDstts sets StateOfDstts field to given value.

### HasStateOfDstts

`func (o *StateOfConfiguration) HasStateOfDstts() bool`

HasStateOfDstts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


