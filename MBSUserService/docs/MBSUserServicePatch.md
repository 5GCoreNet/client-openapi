# MBSUserServicePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtServiceIds** | Pointer to **[]string** |  | [optional] 
**ServClass** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ServAnnModes** | Pointer to [**[]ServiceAnnouncementMode**](ServiceAnnouncementMode.md) |  | [optional] 
**ServNameDescs** | Pointer to [**[]ServiceNameDescription**](ServiceNameDescription.md) |  | [optional] 
**MainServLang** | Pointer to **string** |  | [optional] 

## Methods

### NewMBSUserServicePatch

`func NewMBSUserServicePatch() *MBSUserServicePatch`

NewMBSUserServicePatch instantiates a new MBSUserServicePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserServicePatchWithDefaults

`func NewMBSUserServicePatchWithDefaults() *MBSUserServicePatch`

NewMBSUserServicePatchWithDefaults instantiates a new MBSUserServicePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtServiceIds

`func (o *MBSUserServicePatch) GetExtServiceIds() []string`

GetExtServiceIds returns the ExtServiceIds field if non-nil, zero value otherwise.

### GetExtServiceIdsOk

`func (o *MBSUserServicePatch) GetExtServiceIdsOk() (*[]string, bool)`

GetExtServiceIdsOk returns a tuple with the ExtServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtServiceIds

`func (o *MBSUserServicePatch) SetExtServiceIds(v []string)`

SetExtServiceIds sets ExtServiceIds field to given value.

### HasExtServiceIds

`func (o *MBSUserServicePatch) HasExtServiceIds() bool`

HasExtServiceIds returns a boolean if a field has been set.

### GetServClass

`func (o *MBSUserServicePatch) GetServClass() string`

GetServClass returns the ServClass field if non-nil, zero value otherwise.

### GetServClassOk

`func (o *MBSUserServicePatch) GetServClassOk() (*string, bool)`

GetServClassOk returns a tuple with the ServClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServClass

`func (o *MBSUserServicePatch) SetServClass(v string)`

SetServClass sets ServClass field to given value.

### HasServClass

`func (o *MBSUserServicePatch) HasServClass() bool`

HasServClass returns a boolean if a field has been set.

### GetServAnnModes

`func (o *MBSUserServicePatch) GetServAnnModes() []ServiceAnnouncementMode`

GetServAnnModes returns the ServAnnModes field if non-nil, zero value otherwise.

### GetServAnnModesOk

`func (o *MBSUserServicePatch) GetServAnnModesOk() (*[]ServiceAnnouncementMode, bool)`

GetServAnnModesOk returns a tuple with the ServAnnModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAnnModes

`func (o *MBSUserServicePatch) SetServAnnModes(v []ServiceAnnouncementMode)`

SetServAnnModes sets ServAnnModes field to given value.

### HasServAnnModes

`func (o *MBSUserServicePatch) HasServAnnModes() bool`

HasServAnnModes returns a boolean if a field has been set.

### GetServNameDescs

`func (o *MBSUserServicePatch) GetServNameDescs() []ServiceNameDescription`

GetServNameDescs returns the ServNameDescs field if non-nil, zero value otherwise.

### GetServNameDescsOk

`func (o *MBSUserServicePatch) GetServNameDescsOk() (*[]ServiceNameDescription, bool)`

GetServNameDescsOk returns a tuple with the ServNameDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServNameDescs

`func (o *MBSUserServicePatch) SetServNameDescs(v []ServiceNameDescription)`

SetServNameDescs sets ServNameDescs field to given value.

### HasServNameDescs

`func (o *MBSUserServicePatch) HasServNameDescs() bool`

HasServNameDescs returns a boolean if a field has been set.

### GetMainServLang

`func (o *MBSUserServicePatch) GetMainServLang() string`

GetMainServLang returns the MainServLang field if non-nil, zero value otherwise.

### GetMainServLangOk

`func (o *MBSUserServicePatch) GetMainServLangOk() (*string, bool)`

GetMainServLangOk returns a tuple with the MainServLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainServLang

`func (o *MBSUserServicePatch) SetMainServLang(v string)`

SetMainServLang sets MainServLang field to given value.

### HasMainServLang

`func (o *MBSUserServicePatch) HasMainServLang() bool`

HasMainServLang returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


