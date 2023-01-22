# DeterministicComm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServAttrCom** | Pointer to [**ServAttrCom**](ServAttrCom.md) |  | [optional] 
**Availability** | Pointer to [**Support**](Support.md) |  | [optional] 
**PeriodicityList** | Pointer to **string** |  | [optional] 

## Methods

### NewDeterministicComm

`func NewDeterministicComm() *DeterministicComm`

NewDeterministicComm instantiates a new DeterministicComm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeterministicCommWithDefaults

`func NewDeterministicCommWithDefaults() *DeterministicComm`

NewDeterministicCommWithDefaults instantiates a new DeterministicComm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServAttrCom

`func (o *DeterministicComm) GetServAttrCom() ServAttrCom`

GetServAttrCom returns the ServAttrCom field if non-nil, zero value otherwise.

### GetServAttrComOk

`func (o *DeterministicComm) GetServAttrComOk() (*ServAttrCom, bool)`

GetServAttrComOk returns a tuple with the ServAttrCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAttrCom

`func (o *DeterministicComm) SetServAttrCom(v ServAttrCom)`

SetServAttrCom sets ServAttrCom field to given value.

### HasServAttrCom

`func (o *DeterministicComm) HasServAttrCom() bool`

HasServAttrCom returns a boolean if a field has been set.

### GetAvailability

`func (o *DeterministicComm) GetAvailability() Support`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *DeterministicComm) GetAvailabilityOk() (*Support, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *DeterministicComm) SetAvailability(v Support)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *DeterministicComm) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetPeriodicityList

`func (o *DeterministicComm) GetPeriodicityList() string`

GetPeriodicityList returns the PeriodicityList field if non-nil, zero value otherwise.

### GetPeriodicityListOk

`func (o *DeterministicComm) GetPeriodicityListOk() (*string, bool)`

GetPeriodicityListOk returns a tuple with the PeriodicityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicityList

`func (o *DeterministicComm) SetPeriodicityList(v string)`

SetPeriodicityList sets PeriodicityList field to given value.

### HasPeriodicityList

`func (o *DeterministicComm) HasPeriodicityList() bool`

HasPeriodicityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


