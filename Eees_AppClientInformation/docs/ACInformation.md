# ACInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcProfs** | [**[]ACProfile**](ACProfile.md) | List of profile information of ACs. | 
**UeIds** | Pointer to **[]string** | List of UE identifiers hosting the AC. | [optional] 
**UeLocInfs** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 

## Methods

### NewACInformation

`func NewACInformation(acProfs []ACProfile, ) *ACInformation`

NewACInformation instantiates a new ACInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACInformationWithDefaults

`func NewACInformationWithDefaults() *ACInformation`

NewACInformationWithDefaults instantiates a new ACInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcProfs

`func (o *ACInformation) GetAcProfs() []ACProfile`

GetAcProfs returns the AcProfs field if non-nil, zero value otherwise.

### GetAcProfsOk

`func (o *ACInformation) GetAcProfsOk() (*[]ACProfile, bool)`

GetAcProfsOk returns a tuple with the AcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcProfs

`func (o *ACInformation) SetAcProfs(v []ACProfile)`

SetAcProfs sets AcProfs field to given value.


### GetUeIds

`func (o *ACInformation) GetUeIds() []string`

GetUeIds returns the UeIds field if non-nil, zero value otherwise.

### GetUeIdsOk

`func (o *ACInformation) GetUeIdsOk() (*[]string, bool)`

GetUeIdsOk returns a tuple with the UeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIds

`func (o *ACInformation) SetUeIds(v []string)`

SetUeIds sets UeIds field to given value.

### HasUeIds

`func (o *ACInformation) HasUeIds() bool`

HasUeIds returns a boolean if a field has been set.

### GetUeLocInfs

`func (o *ACInformation) GetUeLocInfs() LocationArea5G`

GetUeLocInfs returns the UeLocInfs field if non-nil, zero value otherwise.

### GetUeLocInfsOk

`func (o *ACInformation) GetUeLocInfsOk() (*LocationArea5G, bool)`

GetUeLocInfsOk returns a tuple with the UeLocInfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocInfs

`func (o *ACInformation) SetUeLocInfs(v LocationArea5G)`

SetUeLocInfs sets UeLocInfs field to given value.

### HasUeLocInfs

`func (o *ACInformation) HasUeLocInfs() bool`

HasUeLocInfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


