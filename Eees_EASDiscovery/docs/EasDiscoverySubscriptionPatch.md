# EasDiscoverySubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasDiscoveryFilter** | Pointer to [**EasDiscoveryFilter**](EasDiscoveryFilter.md) |  | [optional] 
**EasDynInfoFilter** | Pointer to [**EasDynamicInfoFilter**](EasDynamicInfoFilter.md) |  | [optional] 
**EasSvcContinuity** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewEasDiscoverySubscriptionPatch

`func NewEasDiscoverySubscriptionPatch() *EasDiscoverySubscriptionPatch`

NewEasDiscoverySubscriptionPatch instantiates a new EasDiscoverySubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDiscoverySubscriptionPatchWithDefaults

`func NewEasDiscoverySubscriptionPatchWithDefaults() *EasDiscoverySubscriptionPatch`

NewEasDiscoverySubscriptionPatchWithDefaults instantiates a new EasDiscoverySubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasDiscoveryFilter

`func (o *EasDiscoverySubscriptionPatch) GetEasDiscoveryFilter() EasDiscoveryFilter`

GetEasDiscoveryFilter returns the EasDiscoveryFilter field if non-nil, zero value otherwise.

### GetEasDiscoveryFilterOk

`func (o *EasDiscoverySubscriptionPatch) GetEasDiscoveryFilterOk() (*EasDiscoveryFilter, bool)`

GetEasDiscoveryFilterOk returns a tuple with the EasDiscoveryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDiscoveryFilter

`func (o *EasDiscoverySubscriptionPatch) SetEasDiscoveryFilter(v EasDiscoveryFilter)`

SetEasDiscoveryFilter sets EasDiscoveryFilter field to given value.

### HasEasDiscoveryFilter

`func (o *EasDiscoverySubscriptionPatch) HasEasDiscoveryFilter() bool`

HasEasDiscoveryFilter returns a boolean if a field has been set.

### GetEasDynInfoFilter

`func (o *EasDiscoverySubscriptionPatch) GetEasDynInfoFilter() EasDynamicInfoFilter`

GetEasDynInfoFilter returns the EasDynInfoFilter field if non-nil, zero value otherwise.

### GetEasDynInfoFilterOk

`func (o *EasDiscoverySubscriptionPatch) GetEasDynInfoFilterOk() (*EasDynamicInfoFilter, bool)`

GetEasDynInfoFilterOk returns a tuple with the EasDynInfoFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDynInfoFilter

`func (o *EasDiscoverySubscriptionPatch) SetEasDynInfoFilter(v EasDynamicInfoFilter)`

SetEasDynInfoFilter sets EasDynInfoFilter field to given value.

### HasEasDynInfoFilter

`func (o *EasDiscoverySubscriptionPatch) HasEasDynInfoFilter() bool`

HasEasDynInfoFilter returns a boolean if a field has been set.

### GetEasSvcContinuity

`func (o *EasDiscoverySubscriptionPatch) GetEasSvcContinuity() []ACRScenario`

GetEasSvcContinuity returns the EasSvcContinuity field if non-nil, zero value otherwise.

### GetEasSvcContinuityOk

`func (o *EasDiscoverySubscriptionPatch) GetEasSvcContinuityOk() (*[]ACRScenario, bool)`

GetEasSvcContinuityOk returns a tuple with the EasSvcContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasSvcContinuity

`func (o *EasDiscoverySubscriptionPatch) SetEasSvcContinuity(v []ACRScenario)`

SetEasSvcContinuity sets EasSvcContinuity field to given value.

### HasEasSvcContinuity

`func (o *EasDiscoverySubscriptionPatch) HasEasSvcContinuity() bool`

HasEasSvcContinuity returns a boolean if a field has been set.

### GetExpTime

`func (o *EasDiscoverySubscriptionPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *EasDiscoverySubscriptionPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *EasDiscoverySubscriptionPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *EasDiscoverySubscriptionPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


