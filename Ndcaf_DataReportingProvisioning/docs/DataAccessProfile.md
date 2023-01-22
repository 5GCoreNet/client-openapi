# DataAccessProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataAccessProfileId** | **string** |  | 
**TargetEventConsumerTypes** | [**[]EventConsumerType**](EventConsumerType.md) |  | 
**Parameters** | **[]string** |  | 
**TimeAccessRestrictions** | Pointer to [**DataAccessProfileTimeAccessRestrictions**](DataAccessProfileTimeAccessRestrictions.md) |  | [optional] 
**UserAccessRestrictions** | Pointer to [**DataAccessProfileUserAccessRestrictions**](DataAccessProfileUserAccessRestrictions.md) |  | [optional] 
**LocationAccessRestrictions** | Pointer to [**DataAccessProfileLocationAccessRestrictions**](DataAccessProfileLocationAccessRestrictions.md) |  | [optional] 

## Methods

### NewDataAccessProfile

`func NewDataAccessProfile(dataAccessProfileId string, targetEventConsumerTypes []EventConsumerType, parameters []string, ) *DataAccessProfile`

NewDataAccessProfile instantiates a new DataAccessProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataAccessProfileWithDefaults

`func NewDataAccessProfileWithDefaults() *DataAccessProfile`

NewDataAccessProfileWithDefaults instantiates a new DataAccessProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataAccessProfileId

`func (o *DataAccessProfile) GetDataAccessProfileId() string`

GetDataAccessProfileId returns the DataAccessProfileId field if non-nil, zero value otherwise.

### GetDataAccessProfileIdOk

`func (o *DataAccessProfile) GetDataAccessProfileIdOk() (*string, bool)`

GetDataAccessProfileIdOk returns a tuple with the DataAccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessProfileId

`func (o *DataAccessProfile) SetDataAccessProfileId(v string)`

SetDataAccessProfileId sets DataAccessProfileId field to given value.


### GetTargetEventConsumerTypes

`func (o *DataAccessProfile) GetTargetEventConsumerTypes() []EventConsumerType`

GetTargetEventConsumerTypes returns the TargetEventConsumerTypes field if non-nil, zero value otherwise.

### GetTargetEventConsumerTypesOk

`func (o *DataAccessProfile) GetTargetEventConsumerTypesOk() (*[]EventConsumerType, bool)`

GetTargetEventConsumerTypesOk returns a tuple with the TargetEventConsumerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEventConsumerTypes

`func (o *DataAccessProfile) SetTargetEventConsumerTypes(v []EventConsumerType)`

SetTargetEventConsumerTypes sets TargetEventConsumerTypes field to given value.


### GetParameters

`func (o *DataAccessProfile) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DataAccessProfile) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DataAccessProfile) SetParameters(v []string)`

SetParameters sets Parameters field to given value.


### GetTimeAccessRestrictions

`func (o *DataAccessProfile) GetTimeAccessRestrictions() DataAccessProfileTimeAccessRestrictions`

GetTimeAccessRestrictions returns the TimeAccessRestrictions field if non-nil, zero value otherwise.

### GetTimeAccessRestrictionsOk

`func (o *DataAccessProfile) GetTimeAccessRestrictionsOk() (*DataAccessProfileTimeAccessRestrictions, bool)`

GetTimeAccessRestrictionsOk returns a tuple with the TimeAccessRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAccessRestrictions

`func (o *DataAccessProfile) SetTimeAccessRestrictions(v DataAccessProfileTimeAccessRestrictions)`

SetTimeAccessRestrictions sets TimeAccessRestrictions field to given value.

### HasTimeAccessRestrictions

`func (o *DataAccessProfile) HasTimeAccessRestrictions() bool`

HasTimeAccessRestrictions returns a boolean if a field has been set.

### GetUserAccessRestrictions

`func (o *DataAccessProfile) GetUserAccessRestrictions() DataAccessProfileUserAccessRestrictions`

GetUserAccessRestrictions returns the UserAccessRestrictions field if non-nil, zero value otherwise.

### GetUserAccessRestrictionsOk

`func (o *DataAccessProfile) GetUserAccessRestrictionsOk() (*DataAccessProfileUserAccessRestrictions, bool)`

GetUserAccessRestrictionsOk returns a tuple with the UserAccessRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessRestrictions

`func (o *DataAccessProfile) SetUserAccessRestrictions(v DataAccessProfileUserAccessRestrictions)`

SetUserAccessRestrictions sets UserAccessRestrictions field to given value.

### HasUserAccessRestrictions

`func (o *DataAccessProfile) HasUserAccessRestrictions() bool`

HasUserAccessRestrictions returns a boolean if a field has been set.

### GetLocationAccessRestrictions

`func (o *DataAccessProfile) GetLocationAccessRestrictions() DataAccessProfileLocationAccessRestrictions`

GetLocationAccessRestrictions returns the LocationAccessRestrictions field if non-nil, zero value otherwise.

### GetLocationAccessRestrictionsOk

`func (o *DataAccessProfile) GetLocationAccessRestrictionsOk() (*DataAccessProfileLocationAccessRestrictions, bool)`

GetLocationAccessRestrictionsOk returns a tuple with the LocationAccessRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAccessRestrictions

`func (o *DataAccessProfile) SetLocationAccessRestrictions(v DataAccessProfileLocationAccessRestrictions)`

SetLocationAccessRestrictions sets LocationAccessRestrictions field to given value.

### HasLocationAccessRestrictions

`func (o *DataAccessProfile) HasLocationAccessRestrictions() bool`

HasLocationAccessRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


