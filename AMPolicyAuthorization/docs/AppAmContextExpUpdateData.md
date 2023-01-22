# AppAmContextExpUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvSubscs** | Pointer to [**NullableAmEventsSubscDataRm**](AmEventsSubscDataRm.md) |  | [optional] 
**HighThruInd** | Pointer to **bool** |  | [optional] 
**CovReqs** | Pointer to [**[]GeographicalArea**](GeographicalArea.md) |  | [optional] 
**PolicyDuration** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 

## Methods

### NewAppAmContextExpUpdateData

`func NewAppAmContextExpUpdateData() *AppAmContextExpUpdateData`

NewAppAmContextExpUpdateData instantiates a new AppAmContextExpUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAmContextExpUpdateDataWithDefaults

`func NewAppAmContextExpUpdateDataWithDefaults() *AppAmContextExpUpdateData`

NewAppAmContextExpUpdateDataWithDefaults instantiates a new AppAmContextExpUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvSubscs

`func (o *AppAmContextExpUpdateData) GetEvSubscs() AmEventsSubscDataRm`

GetEvSubscs returns the EvSubscs field if non-nil, zero value otherwise.

### GetEvSubscsOk

`func (o *AppAmContextExpUpdateData) GetEvSubscsOk() (*AmEventsSubscDataRm, bool)`

GetEvSubscsOk returns a tuple with the EvSubscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubscs

`func (o *AppAmContextExpUpdateData) SetEvSubscs(v AmEventsSubscDataRm)`

SetEvSubscs sets EvSubscs field to given value.

### HasEvSubscs

`func (o *AppAmContextExpUpdateData) HasEvSubscs() bool`

HasEvSubscs returns a boolean if a field has been set.

### SetEvSubscsNil

`func (o *AppAmContextExpUpdateData) SetEvSubscsNil(b bool)`

 SetEvSubscsNil sets the value for EvSubscs to be an explicit nil

### UnsetEvSubscs
`func (o *AppAmContextExpUpdateData) UnsetEvSubscs()`

UnsetEvSubscs ensures that no value is present for EvSubscs, not even an explicit nil
### GetHighThruInd

`func (o *AppAmContextExpUpdateData) GetHighThruInd() bool`

GetHighThruInd returns the HighThruInd field if non-nil, zero value otherwise.

### GetHighThruIndOk

`func (o *AppAmContextExpUpdateData) GetHighThruIndOk() (*bool, bool)`

GetHighThruIndOk returns a tuple with the HighThruInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighThruInd

`func (o *AppAmContextExpUpdateData) SetHighThruInd(v bool)`

SetHighThruInd sets HighThruInd field to given value.

### HasHighThruInd

`func (o *AppAmContextExpUpdateData) HasHighThruInd() bool`

HasHighThruInd returns a boolean if a field has been set.

### GetCovReqs

`func (o *AppAmContextExpUpdateData) GetCovReqs() []GeographicalArea`

GetCovReqs returns the CovReqs field if non-nil, zero value otherwise.

### GetCovReqsOk

`func (o *AppAmContextExpUpdateData) GetCovReqsOk() (*[]GeographicalArea, bool)`

GetCovReqsOk returns a tuple with the CovReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovReqs

`func (o *AppAmContextExpUpdateData) SetCovReqs(v []GeographicalArea)`

SetCovReqs sets CovReqs field to given value.

### HasCovReqs

`func (o *AppAmContextExpUpdateData) HasCovReqs() bool`

HasCovReqs returns a boolean if a field has been set.

### GetPolicyDuration

`func (o *AppAmContextExpUpdateData) GetPolicyDuration() int32`

GetPolicyDuration returns the PolicyDuration field if non-nil, zero value otherwise.

### GetPolicyDurationOk

`func (o *AppAmContextExpUpdateData) GetPolicyDurationOk() (*int32, bool)`

GetPolicyDurationOk returns a tuple with the PolicyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDuration

`func (o *AppAmContextExpUpdateData) SetPolicyDuration(v int32)`

SetPolicyDuration sets PolicyDuration field to given value.

### HasPolicyDuration

`func (o *AppAmContextExpUpdateData) HasPolicyDuration() bool`

HasPolicyDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


