# M1EASRelocationRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tolerance** | [**EASRelocationTolerance**](EASRelocationTolerance.md) |  | 
**MaxInterruptionDuration** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**MaxResponseTimeDifference** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 

## Methods

### NewM1EASRelocationRequirements

`func NewM1EASRelocationRequirements(tolerance EASRelocationTolerance, ) *M1EASRelocationRequirements`

NewM1EASRelocationRequirements instantiates a new M1EASRelocationRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM1EASRelocationRequirementsWithDefaults

`func NewM1EASRelocationRequirementsWithDefaults() *M1EASRelocationRequirements`

NewM1EASRelocationRequirementsWithDefaults instantiates a new M1EASRelocationRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTolerance

`func (o *M1EASRelocationRequirements) GetTolerance() EASRelocationTolerance`

GetTolerance returns the Tolerance field if non-nil, zero value otherwise.

### GetToleranceOk

`func (o *M1EASRelocationRequirements) GetToleranceOk() (*EASRelocationTolerance, bool)`

GetToleranceOk returns a tuple with the Tolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerance

`func (o *M1EASRelocationRequirements) SetTolerance(v EASRelocationTolerance)`

SetTolerance sets Tolerance field to given value.


### GetMaxInterruptionDuration

`func (o *M1EASRelocationRequirements) GetMaxInterruptionDuration() int32`

GetMaxInterruptionDuration returns the MaxInterruptionDuration field if non-nil, zero value otherwise.

### GetMaxInterruptionDurationOk

`func (o *M1EASRelocationRequirements) GetMaxInterruptionDurationOk() (*int32, bool)`

GetMaxInterruptionDurationOk returns a tuple with the MaxInterruptionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInterruptionDuration

`func (o *M1EASRelocationRequirements) SetMaxInterruptionDuration(v int32)`

SetMaxInterruptionDuration sets MaxInterruptionDuration field to given value.

### HasMaxInterruptionDuration

`func (o *M1EASRelocationRequirements) HasMaxInterruptionDuration() bool`

HasMaxInterruptionDuration returns a boolean if a field has been set.

### SetMaxInterruptionDurationNil

`func (o *M1EASRelocationRequirements) SetMaxInterruptionDurationNil(b bool)`

 SetMaxInterruptionDurationNil sets the value for MaxInterruptionDuration to be an explicit nil

### UnsetMaxInterruptionDuration
`func (o *M1EASRelocationRequirements) UnsetMaxInterruptionDuration()`

UnsetMaxInterruptionDuration ensures that no value is present for MaxInterruptionDuration, not even an explicit nil
### GetMaxResponseTimeDifference

`func (o *M1EASRelocationRequirements) GetMaxResponseTimeDifference() int32`

GetMaxResponseTimeDifference returns the MaxResponseTimeDifference field if non-nil, zero value otherwise.

### GetMaxResponseTimeDifferenceOk

`func (o *M1EASRelocationRequirements) GetMaxResponseTimeDifferenceOk() (*int32, bool)`

GetMaxResponseTimeDifferenceOk returns a tuple with the MaxResponseTimeDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTimeDifference

`func (o *M1EASRelocationRequirements) SetMaxResponseTimeDifference(v int32)`

SetMaxResponseTimeDifference sets MaxResponseTimeDifference field to given value.

### HasMaxResponseTimeDifference

`func (o *M1EASRelocationRequirements) HasMaxResponseTimeDifference() bool`

HasMaxResponseTimeDifference returns a boolean if a field has been set.

### SetMaxResponseTimeDifferenceNil

`func (o *M1EASRelocationRequirements) SetMaxResponseTimeDifferenceNil(b bool)`

 SetMaxResponseTimeDifferenceNil sets the value for MaxResponseTimeDifference to be an explicit nil

### UnsetMaxResponseTimeDifference
`func (o *M1EASRelocationRequirements) UnsetMaxResponseTimeDifference()`

UnsetMaxResponseTimeDifference ensures that no value is present for MaxResponseTimeDifference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


