# ServiceAPIDescriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AefProfiles** | Pointer to [**[]AefProfile**](AefProfile.md) |  | [optional] 
**Description** | Pointer to **string** | Text description of the API | [optional] 
**ShareableInfo** | Pointer to [**ShareableInformation**](ShareableInformation.md) |  | [optional] 
**ServiceAPICategory** | Pointer to **string** |  | [optional] 
**ApiSuppFeats** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PubApiPath** | Pointer to [**PublishedApiPath**](PublishedApiPath.md) |  | [optional] 
**CcfId** | Pointer to **string** | CAPIF core function identifier. | [optional] 

## Methods

### NewServiceAPIDescriptionPatch

`func NewServiceAPIDescriptionPatch() *ServiceAPIDescriptionPatch`

NewServiceAPIDescriptionPatch instantiates a new ServiceAPIDescriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAPIDescriptionPatchWithDefaults

`func NewServiceAPIDescriptionPatchWithDefaults() *ServiceAPIDescriptionPatch`

NewServiceAPIDescriptionPatchWithDefaults instantiates a new ServiceAPIDescriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAefProfiles

`func (o *ServiceAPIDescriptionPatch) GetAefProfiles() []AefProfile`

GetAefProfiles returns the AefProfiles field if non-nil, zero value otherwise.

### GetAefProfilesOk

`func (o *ServiceAPIDescriptionPatch) GetAefProfilesOk() (*[]AefProfile, bool)`

GetAefProfilesOk returns a tuple with the AefProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefProfiles

`func (o *ServiceAPIDescriptionPatch) SetAefProfiles(v []AefProfile)`

SetAefProfiles sets AefProfiles field to given value.

### HasAefProfiles

`func (o *ServiceAPIDescriptionPatch) HasAefProfiles() bool`

HasAefProfiles returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceAPIDescriptionPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAPIDescriptionPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAPIDescriptionPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAPIDescriptionPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShareableInfo

`func (o *ServiceAPIDescriptionPatch) GetShareableInfo() ShareableInformation`

GetShareableInfo returns the ShareableInfo field if non-nil, zero value otherwise.

### GetShareableInfoOk

`func (o *ServiceAPIDescriptionPatch) GetShareableInfoOk() (*ShareableInformation, bool)`

GetShareableInfoOk returns a tuple with the ShareableInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareableInfo

`func (o *ServiceAPIDescriptionPatch) SetShareableInfo(v ShareableInformation)`

SetShareableInfo sets ShareableInfo field to given value.

### HasShareableInfo

`func (o *ServiceAPIDescriptionPatch) HasShareableInfo() bool`

HasShareableInfo returns a boolean if a field has been set.

### GetServiceAPICategory

`func (o *ServiceAPIDescriptionPatch) GetServiceAPICategory() string`

GetServiceAPICategory returns the ServiceAPICategory field if non-nil, zero value otherwise.

### GetServiceAPICategoryOk

`func (o *ServiceAPIDescriptionPatch) GetServiceAPICategoryOk() (*string, bool)`

GetServiceAPICategoryOk returns a tuple with the ServiceAPICategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPICategory

`func (o *ServiceAPIDescriptionPatch) SetServiceAPICategory(v string)`

SetServiceAPICategory sets ServiceAPICategory field to given value.

### HasServiceAPICategory

`func (o *ServiceAPIDescriptionPatch) HasServiceAPICategory() bool`

HasServiceAPICategory returns a boolean if a field has been set.

### GetApiSuppFeats

`func (o *ServiceAPIDescriptionPatch) GetApiSuppFeats() string`

GetApiSuppFeats returns the ApiSuppFeats field if non-nil, zero value otherwise.

### GetApiSuppFeatsOk

`func (o *ServiceAPIDescriptionPatch) GetApiSuppFeatsOk() (*string, bool)`

GetApiSuppFeatsOk returns a tuple with the ApiSuppFeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSuppFeats

`func (o *ServiceAPIDescriptionPatch) SetApiSuppFeats(v string)`

SetApiSuppFeats sets ApiSuppFeats field to given value.

### HasApiSuppFeats

`func (o *ServiceAPIDescriptionPatch) HasApiSuppFeats() bool`

HasApiSuppFeats returns a boolean if a field has been set.

### GetPubApiPath

`func (o *ServiceAPIDescriptionPatch) GetPubApiPath() PublishedApiPath`

GetPubApiPath returns the PubApiPath field if non-nil, zero value otherwise.

### GetPubApiPathOk

`func (o *ServiceAPIDescriptionPatch) GetPubApiPathOk() (*PublishedApiPath, bool)`

GetPubApiPathOk returns a tuple with the PubApiPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubApiPath

`func (o *ServiceAPIDescriptionPatch) SetPubApiPath(v PublishedApiPath)`

SetPubApiPath sets PubApiPath field to given value.

### HasPubApiPath

`func (o *ServiceAPIDescriptionPatch) HasPubApiPath() bool`

HasPubApiPath returns a boolean if a field has been set.

### GetCcfId

`func (o *ServiceAPIDescriptionPatch) GetCcfId() string`

GetCcfId returns the CcfId field if non-nil, zero value otherwise.

### GetCcfIdOk

`func (o *ServiceAPIDescriptionPatch) GetCcfIdOk() (*string, bool)`

GetCcfIdOk returns a tuple with the CcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcfId

`func (o *ServiceAPIDescriptionPatch) SetCcfId(v string)`

SetCcfId sets CcfId field to given value.

### HasCcfId

`func (o *ServiceAPIDescriptionPatch) HasCcfId() bool`

HasCcfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


