# AmRequestedValueRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**PraStatuses** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | Contains the UE presence statuses for tracking areas. The praId attribute within the PresenceInfo data type is the key of the map.  | [optional] 
**AccessTypes** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**RatTypes** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**AllowedSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of allowed S-NSSAIs for the 3GPP access. | [optional] 
**N3gAllowedSnssais** | Pointer to [**[]Snssai**](Snssai.md) | array of allowed S-NSSAIs for the Non-3GPP access. | [optional] 

## Methods

### NewAmRequestedValueRep

`func NewAmRequestedValueRep() *AmRequestedValueRep`

NewAmRequestedValueRep instantiates a new AmRequestedValueRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmRequestedValueRepWithDefaults

`func NewAmRequestedValueRepWithDefaults() *AmRequestedValueRep`

NewAmRequestedValueRepWithDefaults instantiates a new AmRequestedValueRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLoc

`func (o *AmRequestedValueRep) GetUserLoc() UserLocation`

GetUserLoc returns the UserLoc field if non-nil, zero value otherwise.

### GetUserLocOk

`func (o *AmRequestedValueRep) GetUserLocOk() (*UserLocation, bool)`

GetUserLocOk returns a tuple with the UserLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoc

`func (o *AmRequestedValueRep) SetUserLoc(v UserLocation)`

SetUserLoc sets UserLoc field to given value.

### HasUserLoc

`func (o *AmRequestedValueRep) HasUserLoc() bool`

HasUserLoc returns a boolean if a field has been set.

### GetPraStatuses

`func (o *AmRequestedValueRep) GetPraStatuses() map[string]PresenceInfo`

GetPraStatuses returns the PraStatuses field if non-nil, zero value otherwise.

### GetPraStatusesOk

`func (o *AmRequestedValueRep) GetPraStatusesOk() (*map[string]PresenceInfo, bool)`

GetPraStatusesOk returns a tuple with the PraStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraStatuses

`func (o *AmRequestedValueRep) SetPraStatuses(v map[string]PresenceInfo)`

SetPraStatuses sets PraStatuses field to given value.

### HasPraStatuses

`func (o *AmRequestedValueRep) HasPraStatuses() bool`

HasPraStatuses returns a boolean if a field has been set.

### GetAccessTypes

`func (o *AmRequestedValueRep) GetAccessTypes() []AccessType`

GetAccessTypes returns the AccessTypes field if non-nil, zero value otherwise.

### GetAccessTypesOk

`func (o *AmRequestedValueRep) GetAccessTypesOk() (*[]AccessType, bool)`

GetAccessTypesOk returns a tuple with the AccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypes

`func (o *AmRequestedValueRep) SetAccessTypes(v []AccessType)`

SetAccessTypes sets AccessTypes field to given value.

### HasAccessTypes

`func (o *AmRequestedValueRep) HasAccessTypes() bool`

HasAccessTypes returns a boolean if a field has been set.

### GetRatTypes

`func (o *AmRequestedValueRep) GetRatTypes() []RatType`

GetRatTypes returns the RatTypes field if non-nil, zero value otherwise.

### GetRatTypesOk

`func (o *AmRequestedValueRep) GetRatTypesOk() (*[]RatType, bool)`

GetRatTypesOk returns a tuple with the RatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatTypes

`func (o *AmRequestedValueRep) SetRatTypes(v []RatType)`

SetRatTypes sets RatTypes field to given value.

### HasRatTypes

`func (o *AmRequestedValueRep) HasRatTypes() bool`

HasRatTypes returns a boolean if a field has been set.

### GetAllowedSnssais

`func (o *AmRequestedValueRep) GetAllowedSnssais() []Snssai`

GetAllowedSnssais returns the AllowedSnssais field if non-nil, zero value otherwise.

### GetAllowedSnssaisOk

`func (o *AmRequestedValueRep) GetAllowedSnssaisOk() (*[]Snssai, bool)`

GetAllowedSnssaisOk returns a tuple with the AllowedSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnssais

`func (o *AmRequestedValueRep) SetAllowedSnssais(v []Snssai)`

SetAllowedSnssais sets AllowedSnssais field to given value.

### HasAllowedSnssais

`func (o *AmRequestedValueRep) HasAllowedSnssais() bool`

HasAllowedSnssais returns a boolean if a field has been set.

### GetN3gAllowedSnssais

`func (o *AmRequestedValueRep) GetN3gAllowedSnssais() []Snssai`

GetN3gAllowedSnssais returns the N3gAllowedSnssais field if non-nil, zero value otherwise.

### GetN3gAllowedSnssaisOk

`func (o *AmRequestedValueRep) GetN3gAllowedSnssaisOk() (*[]Snssai, bool)`

GetN3gAllowedSnssaisOk returns a tuple with the N3gAllowedSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gAllowedSnssais

`func (o *AmRequestedValueRep) SetN3gAllowedSnssais(v []Snssai)`

SetN3gAllowedSnssais sets N3gAllowedSnssais field to given value.

### HasN3gAllowedSnssais

`func (o *AmRequestedValueRep) HasN3gAllowedSnssais() bool`

HasN3gAllowedSnssais returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


