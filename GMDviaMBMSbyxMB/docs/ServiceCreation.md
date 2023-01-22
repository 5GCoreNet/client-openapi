# ServiceCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**UserServiceId** | Pointer to **string** | Identifies the MBMS User Service supplied by the SCEF. | [optional] [readonly] 
**ServiceClass** | Pointer to **string** | The service class that service belongs to supplied by the SCEF. | [optional] [readonly] 
**ServiceLanguages** | Pointer to **[]string** | List of language of the service content supplied by the SCEF. | [optional] [readonly] 
**ServiceNames** | Pointer to **[]string** | List of Service Names supplied by the SCEF. | [optional] [readonly] 
**ReceiveOnlyMode** | Pointer to **bool** | When set to &#39;true&#39;, the Content Provider indicates that the service is a Receive Only Mode service. This parameter is supplied by the SCEF. | [optional] [readonly] 
**ServiceAnnouncementMode** | Pointer to [**ServiceAnnouncementMode**](ServiceAnnouncementMode.md) |  | [optional] 

## Methods

### NewServiceCreation

`func NewServiceCreation() *ServiceCreation`

NewServiceCreation instantiates a new ServiceCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreationWithDefaults

`func NewServiceCreationWithDefaults() *ServiceCreation`

NewServiceCreationWithDefaults instantiates a new ServiceCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ServiceCreation) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ServiceCreation) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ServiceCreation) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ServiceCreation) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ServiceCreation) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ServiceCreation) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ServiceCreation) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ServiceCreation) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *ServiceCreation) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *ServiceCreation) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *ServiceCreation) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *ServiceCreation) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetUserServiceId

`func (o *ServiceCreation) GetUserServiceId() string`

GetUserServiceId returns the UserServiceId field if non-nil, zero value otherwise.

### GetUserServiceIdOk

`func (o *ServiceCreation) GetUserServiceIdOk() (*string, bool)`

GetUserServiceIdOk returns a tuple with the UserServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserServiceId

`func (o *ServiceCreation) SetUserServiceId(v string)`

SetUserServiceId sets UserServiceId field to given value.

### HasUserServiceId

`func (o *ServiceCreation) HasUserServiceId() bool`

HasUserServiceId returns a boolean if a field has been set.

### GetServiceClass

`func (o *ServiceCreation) GetServiceClass() string`

GetServiceClass returns the ServiceClass field if non-nil, zero value otherwise.

### GetServiceClassOk

`func (o *ServiceCreation) GetServiceClassOk() (*string, bool)`

GetServiceClassOk returns a tuple with the ServiceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClass

`func (o *ServiceCreation) SetServiceClass(v string)`

SetServiceClass sets ServiceClass field to given value.

### HasServiceClass

`func (o *ServiceCreation) HasServiceClass() bool`

HasServiceClass returns a boolean if a field has been set.

### GetServiceLanguages

`func (o *ServiceCreation) GetServiceLanguages() []string`

GetServiceLanguages returns the ServiceLanguages field if non-nil, zero value otherwise.

### GetServiceLanguagesOk

`func (o *ServiceCreation) GetServiceLanguagesOk() (*[]string, bool)`

GetServiceLanguagesOk returns a tuple with the ServiceLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanguages

`func (o *ServiceCreation) SetServiceLanguages(v []string)`

SetServiceLanguages sets ServiceLanguages field to given value.

### HasServiceLanguages

`func (o *ServiceCreation) HasServiceLanguages() bool`

HasServiceLanguages returns a boolean if a field has been set.

### GetServiceNames

`func (o *ServiceCreation) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *ServiceCreation) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *ServiceCreation) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *ServiceCreation) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.

### GetReceiveOnlyMode

`func (o *ServiceCreation) GetReceiveOnlyMode() bool`

GetReceiveOnlyMode returns the ReceiveOnlyMode field if non-nil, zero value otherwise.

### GetReceiveOnlyModeOk

`func (o *ServiceCreation) GetReceiveOnlyModeOk() (*bool, bool)`

GetReceiveOnlyModeOk returns a tuple with the ReceiveOnlyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveOnlyMode

`func (o *ServiceCreation) SetReceiveOnlyMode(v bool)`

SetReceiveOnlyMode sets ReceiveOnlyMode field to given value.

### HasReceiveOnlyMode

`func (o *ServiceCreation) HasReceiveOnlyMode() bool`

HasReceiveOnlyMode returns a boolean if a field has been set.

### GetServiceAnnouncementMode

`func (o *ServiceCreation) GetServiceAnnouncementMode() ServiceAnnouncementMode`

GetServiceAnnouncementMode returns the ServiceAnnouncementMode field if non-nil, zero value otherwise.

### GetServiceAnnouncementModeOk

`func (o *ServiceCreation) GetServiceAnnouncementModeOk() (*ServiceAnnouncementMode, bool)`

GetServiceAnnouncementModeOk returns a tuple with the ServiceAnnouncementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAnnouncementMode

`func (o *ServiceCreation) SetServiceAnnouncementMode(v ServiceAnnouncementMode)`

SetServiceAnnouncementMode sets ServiceAnnouncementMode field to given value.

### HasServiceAnnouncementMode

`func (o *ServiceCreation) HasServiceAnnouncementMode() bool`

HasServiceAnnouncementMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


