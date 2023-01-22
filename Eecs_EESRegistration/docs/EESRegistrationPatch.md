# EESRegistrationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EesProf** | Pointer to [**EESProfile**](EESProfile.md) |  | [optional] 
**ExpTime** | Pointer to **NullableTime** | string with format &#39;date-time&#39; as defined in OpenAPI with &#39;nullable:true&#39; property.   | [optional] 

## Methods

### NewEESRegistrationPatch

`func NewEESRegistrationPatch() *EESRegistrationPatch`

NewEESRegistrationPatch instantiates a new EESRegistrationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEESRegistrationPatchWithDefaults

`func NewEESRegistrationPatchWithDefaults() *EESRegistrationPatch`

NewEESRegistrationPatchWithDefaults instantiates a new EESRegistrationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEesProf

`func (o *EESRegistrationPatch) GetEesProf() EESProfile`

GetEesProf returns the EesProf field if non-nil, zero value otherwise.

### GetEesProfOk

`func (o *EESRegistrationPatch) GetEesProfOk() (*EESProfile, bool)`

GetEesProfOk returns a tuple with the EesProf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEesProf

`func (o *EESRegistrationPatch) SetEesProf(v EESProfile)`

SetEesProf sets EesProf field to given value.

### HasEesProf

`func (o *EESRegistrationPatch) HasEesProf() bool`

HasEesProf returns a boolean if a field has been set.

### GetExpTime

`func (o *EESRegistrationPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *EESRegistrationPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *EESRegistrationPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *EESRegistrationPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### SetExpTimeNil

`func (o *EESRegistrationPatch) SetExpTimeNil(b bool)`

 SetExpTimeNil sets the value for ExpTime to be an explicit nil

### UnsetExpTime
`func (o *EESRegistrationPatch) UnsetExpTime()`

UnsetExpTime ensures that no value is present for ExpTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


