# ServiceAPIDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiName** | **string** | API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**ApiId** | Pointer to **string** | API identifier assigned by the CAPIF core function to the published service API. Shall not be present in the HTTP POST request from the API publishing function to the CAPIF core function. Shall be present in the HTTP POST response from the CAPIF core function to the API publishing function and in the HTTP GET response from the CAPIF core function to the API invoker (discovery API).  | [optional] 
**AefProfiles** | Pointer to [**[]AefProfile**](AefProfile.md) | AEF profile information, which includes the exposed API details (e.g. protocol).  | [optional] 
**Description** | Pointer to **string** | Text description of the API | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ShareableInfo** | Pointer to [**ShareableInformation**](ShareableInformation.md) |  | [optional] 
**ServiceAPICategory** | Pointer to **string** |  | [optional] 
**ApiSuppFeats** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PubApiPath** | Pointer to [**PublishedApiPath**](PublishedApiPath.md) |  | [optional] 
**CcfId** | Pointer to **string** | CAPIF core function identifier. | [optional] 

## Methods

### NewServiceAPIDescription

`func NewServiceAPIDescription(apiName string, ) *ServiceAPIDescription`

NewServiceAPIDescription instantiates a new ServiceAPIDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAPIDescriptionWithDefaults

`func NewServiceAPIDescriptionWithDefaults() *ServiceAPIDescription`

NewServiceAPIDescriptionWithDefaults instantiates a new ServiceAPIDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *ServiceAPIDescription) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *ServiceAPIDescription) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *ServiceAPIDescription) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetApiId

`func (o *ServiceAPIDescription) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *ServiceAPIDescription) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *ServiceAPIDescription) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *ServiceAPIDescription) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetAefProfiles

`func (o *ServiceAPIDescription) GetAefProfiles() []AefProfile`

GetAefProfiles returns the AefProfiles field if non-nil, zero value otherwise.

### GetAefProfilesOk

`func (o *ServiceAPIDescription) GetAefProfilesOk() (*[]AefProfile, bool)`

GetAefProfilesOk returns a tuple with the AefProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefProfiles

`func (o *ServiceAPIDescription) SetAefProfiles(v []AefProfile)`

SetAefProfiles sets AefProfiles field to given value.

### HasAefProfiles

`func (o *ServiceAPIDescription) HasAefProfiles() bool`

HasAefProfiles returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceAPIDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAPIDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAPIDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAPIDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ServiceAPIDescription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ServiceAPIDescription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ServiceAPIDescription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ServiceAPIDescription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetShareableInfo

`func (o *ServiceAPIDescription) GetShareableInfo() ShareableInformation`

GetShareableInfo returns the ShareableInfo field if non-nil, zero value otherwise.

### GetShareableInfoOk

`func (o *ServiceAPIDescription) GetShareableInfoOk() (*ShareableInformation, bool)`

GetShareableInfoOk returns a tuple with the ShareableInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareableInfo

`func (o *ServiceAPIDescription) SetShareableInfo(v ShareableInformation)`

SetShareableInfo sets ShareableInfo field to given value.

### HasShareableInfo

`func (o *ServiceAPIDescription) HasShareableInfo() bool`

HasShareableInfo returns a boolean if a field has been set.

### GetServiceAPICategory

`func (o *ServiceAPIDescription) GetServiceAPICategory() string`

GetServiceAPICategory returns the ServiceAPICategory field if non-nil, zero value otherwise.

### GetServiceAPICategoryOk

`func (o *ServiceAPIDescription) GetServiceAPICategoryOk() (*string, bool)`

GetServiceAPICategoryOk returns a tuple with the ServiceAPICategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPICategory

`func (o *ServiceAPIDescription) SetServiceAPICategory(v string)`

SetServiceAPICategory sets ServiceAPICategory field to given value.

### HasServiceAPICategory

`func (o *ServiceAPIDescription) HasServiceAPICategory() bool`

HasServiceAPICategory returns a boolean if a field has been set.

### GetApiSuppFeats

`func (o *ServiceAPIDescription) GetApiSuppFeats() string`

GetApiSuppFeats returns the ApiSuppFeats field if non-nil, zero value otherwise.

### GetApiSuppFeatsOk

`func (o *ServiceAPIDescription) GetApiSuppFeatsOk() (*string, bool)`

GetApiSuppFeatsOk returns a tuple with the ApiSuppFeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSuppFeats

`func (o *ServiceAPIDescription) SetApiSuppFeats(v string)`

SetApiSuppFeats sets ApiSuppFeats field to given value.

### HasApiSuppFeats

`func (o *ServiceAPIDescription) HasApiSuppFeats() bool`

HasApiSuppFeats returns a boolean if a field has been set.

### GetPubApiPath

`func (o *ServiceAPIDescription) GetPubApiPath() PublishedApiPath`

GetPubApiPath returns the PubApiPath field if non-nil, zero value otherwise.

### GetPubApiPathOk

`func (o *ServiceAPIDescription) GetPubApiPathOk() (*PublishedApiPath, bool)`

GetPubApiPathOk returns a tuple with the PubApiPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubApiPath

`func (o *ServiceAPIDescription) SetPubApiPath(v PublishedApiPath)`

SetPubApiPath sets PubApiPath field to given value.

### HasPubApiPath

`func (o *ServiceAPIDescription) HasPubApiPath() bool`

HasPubApiPath returns a boolean if a field has been set.

### GetCcfId

`func (o *ServiceAPIDescription) GetCcfId() string`

GetCcfId returns the CcfId field if non-nil, zero value otherwise.

### GetCcfIdOk

`func (o *ServiceAPIDescription) GetCcfIdOk() (*string, bool)`

GetCcfIdOk returns a tuple with the CcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcfId

`func (o *ServiceAPIDescription) SetCcfId(v string)`

SetCcfId sets CcfId field to given value.

### HasCcfId

`func (o *ServiceAPIDescription) HasCcfId() bool`

HasCcfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


