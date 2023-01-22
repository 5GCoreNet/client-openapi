# MBSUserService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtServiceIds** | **[]string** |  | 
**ServType** | [**MbsServiceType**](MbsServiceType.md) |  | 
**ServClass** | **string** | String providing an URI formatted according to RFC 3986. | 
**ServAnnModes** | [**[]ServiceAnnouncementMode**](ServiceAnnouncementMode.md) |  | 
**ServNameDescs** | Pointer to [**[]ServiceNameDescription**](ServiceNameDescription.md) |  | [optional] 
**MainServLang** | Pointer to **string** |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMBSUserService

`func NewMBSUserService(extServiceIds []string, servType MbsServiceType, servClass string, servAnnModes []ServiceAnnouncementMode, ) *MBSUserService`

NewMBSUserService instantiates a new MBSUserService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserServiceWithDefaults

`func NewMBSUserServiceWithDefaults() *MBSUserService`

NewMBSUserServiceWithDefaults instantiates a new MBSUserService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtServiceIds

`func (o *MBSUserService) GetExtServiceIds() []string`

GetExtServiceIds returns the ExtServiceIds field if non-nil, zero value otherwise.

### GetExtServiceIdsOk

`func (o *MBSUserService) GetExtServiceIdsOk() (*[]string, bool)`

GetExtServiceIdsOk returns a tuple with the ExtServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtServiceIds

`func (o *MBSUserService) SetExtServiceIds(v []string)`

SetExtServiceIds sets ExtServiceIds field to given value.


### GetServType

`func (o *MBSUserService) GetServType() MbsServiceType`

GetServType returns the ServType field if non-nil, zero value otherwise.

### GetServTypeOk

`func (o *MBSUserService) GetServTypeOk() (*MbsServiceType, bool)`

GetServTypeOk returns a tuple with the ServType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServType

`func (o *MBSUserService) SetServType(v MbsServiceType)`

SetServType sets ServType field to given value.


### GetServClass

`func (o *MBSUserService) GetServClass() string`

GetServClass returns the ServClass field if non-nil, zero value otherwise.

### GetServClassOk

`func (o *MBSUserService) GetServClassOk() (*string, bool)`

GetServClassOk returns a tuple with the ServClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServClass

`func (o *MBSUserService) SetServClass(v string)`

SetServClass sets ServClass field to given value.


### GetServAnnModes

`func (o *MBSUserService) GetServAnnModes() []ServiceAnnouncementMode`

GetServAnnModes returns the ServAnnModes field if non-nil, zero value otherwise.

### GetServAnnModesOk

`func (o *MBSUserService) GetServAnnModesOk() (*[]ServiceAnnouncementMode, bool)`

GetServAnnModesOk returns a tuple with the ServAnnModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAnnModes

`func (o *MBSUserService) SetServAnnModes(v []ServiceAnnouncementMode)`

SetServAnnModes sets ServAnnModes field to given value.


### GetServNameDescs

`func (o *MBSUserService) GetServNameDescs() []ServiceNameDescription`

GetServNameDescs returns the ServNameDescs field if non-nil, zero value otherwise.

### GetServNameDescsOk

`func (o *MBSUserService) GetServNameDescsOk() (*[]ServiceNameDescription, bool)`

GetServNameDescsOk returns a tuple with the ServNameDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServNameDescs

`func (o *MBSUserService) SetServNameDescs(v []ServiceNameDescription)`

SetServNameDescs sets ServNameDescs field to given value.

### HasServNameDescs

`func (o *MBSUserService) HasServNameDescs() bool`

HasServNameDescs returns a boolean if a field has been set.

### GetMainServLang

`func (o *MBSUserService) GetMainServLang() string`

GetMainServLang returns the MainServLang field if non-nil, zero value otherwise.

### GetMainServLangOk

`func (o *MBSUserService) GetMainServLangOk() (*string, bool)`

GetMainServLangOk returns a tuple with the MainServLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainServLang

`func (o *MBSUserService) SetMainServLang(v string)`

SetMainServLang sets MainServLang field to given value.

### HasMainServLang

`func (o *MBSUserService) HasMainServLang() bool`

HasMainServLang returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MBSUserService) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MBSUserService) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MBSUserService) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MBSUserService) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


