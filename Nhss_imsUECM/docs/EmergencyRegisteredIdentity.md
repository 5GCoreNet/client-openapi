# EmergencyRegisteredIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impi** | **string** | IMS Private Identity of the UE | 
**Impu** | **string** | IMS Public Identity of the UE (sip URI or tel URI) | 

## Methods

### NewEmergencyRegisteredIdentity

`func NewEmergencyRegisteredIdentity(impi string, impu string, ) *EmergencyRegisteredIdentity`

NewEmergencyRegisteredIdentity instantiates a new EmergencyRegisteredIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmergencyRegisteredIdentityWithDefaults

`func NewEmergencyRegisteredIdentityWithDefaults() *EmergencyRegisteredIdentity`

NewEmergencyRegisteredIdentityWithDefaults instantiates a new EmergencyRegisteredIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpi

`func (o *EmergencyRegisteredIdentity) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *EmergencyRegisteredIdentity) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *EmergencyRegisteredIdentity) SetImpi(v string)`

SetImpi sets Impi field to given value.


### GetImpu

`func (o *EmergencyRegisteredIdentity) GetImpu() string`

GetImpu returns the Impu field if non-nil, zero value otherwise.

### GetImpuOk

`func (o *EmergencyRegisteredIdentity) GetImpuOk() (*string, bool)`

GetImpuOk returns a tuple with the Impu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpu

`func (o *EmergencyRegisteredIdentity) SetImpu(v string)`

SetImpu sets Impu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


