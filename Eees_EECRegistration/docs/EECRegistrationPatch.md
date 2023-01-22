# EECRegistrationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcProfs** | Pointer to [**[]ACProfile**](ACProfile.md) | Profiles of ACs for which the EEC provides edge enabling services. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**UnfulfilledAcProfs** | Pointer to [**UnfulfilledAcProfile**](UnfulfilledAcProfile.md) |  | [optional] 

## Methods

### NewEECRegistrationPatch

`func NewEECRegistrationPatch() *EECRegistrationPatch`

NewEECRegistrationPatch instantiates a new EECRegistrationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEECRegistrationPatchWithDefaults

`func NewEECRegistrationPatchWithDefaults() *EECRegistrationPatch`

NewEECRegistrationPatchWithDefaults instantiates a new EECRegistrationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcProfs

`func (o *EECRegistrationPatch) GetAcProfs() []ACProfile`

GetAcProfs returns the AcProfs field if non-nil, zero value otherwise.

### GetAcProfsOk

`func (o *EECRegistrationPatch) GetAcProfsOk() (*[]ACProfile, bool)`

GetAcProfsOk returns a tuple with the AcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcProfs

`func (o *EECRegistrationPatch) SetAcProfs(v []ACProfile)`

SetAcProfs sets AcProfs field to given value.

### HasAcProfs

`func (o *EECRegistrationPatch) HasAcProfs() bool`

HasAcProfs returns a boolean if a field has been set.

### GetExpTime

`func (o *EECRegistrationPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *EECRegistrationPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *EECRegistrationPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *EECRegistrationPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetUnfulfilledAcProfs

`func (o *EECRegistrationPatch) GetUnfulfilledAcProfs() UnfulfilledAcProfile`

GetUnfulfilledAcProfs returns the UnfulfilledAcProfs field if non-nil, zero value otherwise.

### GetUnfulfilledAcProfsOk

`func (o *EECRegistrationPatch) GetUnfulfilledAcProfsOk() (*UnfulfilledAcProfile, bool)`

GetUnfulfilledAcProfsOk returns a tuple with the UnfulfilledAcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfulfilledAcProfs

`func (o *EECRegistrationPatch) SetUnfulfilledAcProfs(v UnfulfilledAcProfile)`

SetUnfulfilledAcProfs sets UnfulfilledAcProfs field to given value.

### HasUnfulfilledAcProfs

`func (o *EECRegistrationPatch) HasUnfulfilledAcProfs() bool`

HasUnfulfilledAcProfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


