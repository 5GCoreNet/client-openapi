# EASRegistrationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasProf** | Pointer to [**EASProfile**](EASProfile.md) |  | [optional] 
**ExpTime** | Pointer to **NullableTime** | string with format &#39;date-time&#39; as defined in OpenAPI with &#39;nullable:true&#39; property.   | [optional] 

## Methods

### NewEASRegistrationPatch

`func NewEASRegistrationPatch() *EASRegistrationPatch`

NewEASRegistrationPatch instantiates a new EASRegistrationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEASRegistrationPatchWithDefaults

`func NewEASRegistrationPatchWithDefaults() *EASRegistrationPatch`

NewEASRegistrationPatchWithDefaults instantiates a new EASRegistrationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasProf

`func (o *EASRegistrationPatch) GetEasProf() EASProfile`

GetEasProf returns the EasProf field if non-nil, zero value otherwise.

### GetEasProfOk

`func (o *EASRegistrationPatch) GetEasProfOk() (*EASProfile, bool)`

GetEasProfOk returns a tuple with the EasProf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasProf

`func (o *EASRegistrationPatch) SetEasProf(v EASProfile)`

SetEasProf sets EasProf field to given value.

### HasEasProf

`func (o *EASRegistrationPatch) HasEasProf() bool`

HasEasProf returns a boolean if a field has been set.

### GetExpTime

`func (o *EASRegistrationPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *EASRegistrationPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *EASRegistrationPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *EASRegistrationPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### SetExpTimeNil

`func (o *EASRegistrationPatch) SetExpTimeNil(b bool)`

 SetExpTimeNil sets the value for ExpTime to be an explicit nil

### UnsetExpTime
`func (o *EASRegistrationPatch) UnsetExpTime()`

UnsetExpTime ensures that no value is present for ExpTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


