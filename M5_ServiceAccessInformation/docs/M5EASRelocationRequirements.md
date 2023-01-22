# M5EASRelocationRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tolerance** | [**EASRelocationTolerance**](EASRelocationTolerance.md) |  | 
**MaxInterruptionDuration** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 

## Methods

### NewM5EASRelocationRequirements

`func NewM5EASRelocationRequirements(tolerance EASRelocationTolerance, ) *M5EASRelocationRequirements`

NewM5EASRelocationRequirements instantiates a new M5EASRelocationRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM5EASRelocationRequirementsWithDefaults

`func NewM5EASRelocationRequirementsWithDefaults() *M5EASRelocationRequirements`

NewM5EASRelocationRequirementsWithDefaults instantiates a new M5EASRelocationRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTolerance

`func (o *M5EASRelocationRequirements) GetTolerance() EASRelocationTolerance`

GetTolerance returns the Tolerance field if non-nil, zero value otherwise.

### GetToleranceOk

`func (o *M5EASRelocationRequirements) GetToleranceOk() (*EASRelocationTolerance, bool)`

GetToleranceOk returns a tuple with the Tolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerance

`func (o *M5EASRelocationRequirements) SetTolerance(v EASRelocationTolerance)`

SetTolerance sets Tolerance field to given value.


### GetMaxInterruptionDuration

`func (o *M5EASRelocationRequirements) GetMaxInterruptionDuration() int32`

GetMaxInterruptionDuration returns the MaxInterruptionDuration field if non-nil, zero value otherwise.

### GetMaxInterruptionDurationOk

`func (o *M5EASRelocationRequirements) GetMaxInterruptionDurationOk() (*int32, bool)`

GetMaxInterruptionDurationOk returns a tuple with the MaxInterruptionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInterruptionDuration

`func (o *M5EASRelocationRequirements) SetMaxInterruptionDuration(v int32)`

SetMaxInterruptionDuration sets MaxInterruptionDuration field to given value.

### HasMaxInterruptionDuration

`func (o *M5EASRelocationRequirements) HasMaxInterruptionDuration() bool`

HasMaxInterruptionDuration returns a boolean if a field has been set.

### SetMaxInterruptionDurationNil

`func (o *M5EASRelocationRequirements) SetMaxInterruptionDurationNil(b bool)`

 SetMaxInterruptionDurationNil sets the value for MaxInterruptionDuration to be an explicit nil

### UnsetMaxInterruptionDuration
`func (o *M5EASRelocationRequirements) UnsetMaxInterruptionDuration()`

UnsetMaxInterruptionDuration ensures that no value is present for MaxInterruptionDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


