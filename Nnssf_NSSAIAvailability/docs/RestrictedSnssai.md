# RestrictedSnssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HomePlmnId** | [**PlmnId**](PlmnId.md) |  | 
**SNssaiList** | [**[]ExtSnssai**](ExtSnssai.md) |  | 
**HomePlmnIdList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**RoamingRestriction** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRestrictedSnssai

`func NewRestrictedSnssai(homePlmnId PlmnId, sNssaiList []ExtSnssai, ) *RestrictedSnssai`

NewRestrictedSnssai instantiates a new RestrictedSnssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedSnssaiWithDefaults

`func NewRestrictedSnssaiWithDefaults() *RestrictedSnssai`

NewRestrictedSnssaiWithDefaults instantiates a new RestrictedSnssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHomePlmnId

`func (o *RestrictedSnssai) GetHomePlmnId() PlmnId`

GetHomePlmnId returns the HomePlmnId field if non-nil, zero value otherwise.

### GetHomePlmnIdOk

`func (o *RestrictedSnssai) GetHomePlmnIdOk() (*PlmnId, bool)`

GetHomePlmnIdOk returns a tuple with the HomePlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePlmnId

`func (o *RestrictedSnssai) SetHomePlmnId(v PlmnId)`

SetHomePlmnId sets HomePlmnId field to given value.


### GetSNssaiList

`func (o *RestrictedSnssai) GetSNssaiList() []ExtSnssai`

GetSNssaiList returns the SNssaiList field if non-nil, zero value otherwise.

### GetSNssaiListOk

`func (o *RestrictedSnssai) GetSNssaiListOk() (*[]ExtSnssai, bool)`

GetSNssaiListOk returns a tuple with the SNssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiList

`func (o *RestrictedSnssai) SetSNssaiList(v []ExtSnssai)`

SetSNssaiList sets SNssaiList field to given value.


### GetHomePlmnIdList

`func (o *RestrictedSnssai) GetHomePlmnIdList() []PlmnId`

GetHomePlmnIdList returns the HomePlmnIdList field if non-nil, zero value otherwise.

### GetHomePlmnIdListOk

`func (o *RestrictedSnssai) GetHomePlmnIdListOk() (*[]PlmnId, bool)`

GetHomePlmnIdListOk returns a tuple with the HomePlmnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePlmnIdList

`func (o *RestrictedSnssai) SetHomePlmnIdList(v []PlmnId)`

SetHomePlmnIdList sets HomePlmnIdList field to given value.

### HasHomePlmnIdList

`func (o *RestrictedSnssai) HasHomePlmnIdList() bool`

HasHomePlmnIdList returns a boolean if a field has been set.

### GetRoamingRestriction

`func (o *RestrictedSnssai) GetRoamingRestriction() bool`

GetRoamingRestriction returns the RoamingRestriction field if non-nil, zero value otherwise.

### GetRoamingRestrictionOk

`func (o *RestrictedSnssai) GetRoamingRestrictionOk() (*bool, bool)`

GetRoamingRestrictionOk returns a tuple with the RoamingRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingRestriction

`func (o *RestrictedSnssai) SetRoamingRestriction(v bool)`

SetRoamingRestriction sets RoamingRestriction field to given value.

### HasRoamingRestriction

`func (o *RestrictedSnssai) HasRoamingRestriction() bool`

HasRoamingRestriction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


