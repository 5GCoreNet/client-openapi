# EeGroupProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedEventTypes** | Pointer to [**[]EventType**](EventType.md) |  | [optional] 
**AllowedMtcProvider** | Pointer to [**map[string][]MtcProvider**](array.md) | A map (list of key-value pairs where EventType serves as key) of MTC provider lists. In addition to defined EventTypes, the key value \&quot;ALL\&quot; may be used to identify a map entry which contains a list of MtcProviders that are allowed monitoring all Event Types. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**IwkEpcRestricted** | Pointer to **bool** |  | [optional] [default to false]
**ExtGroupId** | Pointer to **string** |  | [optional] 
**HssGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 

## Methods

### NewEeGroupProfileData

`func NewEeGroupProfileData() *EeGroupProfileData`

NewEeGroupProfileData instantiates a new EeGroupProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeGroupProfileDataWithDefaults

`func NewEeGroupProfileDataWithDefaults() *EeGroupProfileData`

NewEeGroupProfileDataWithDefaults instantiates a new EeGroupProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedEventTypes

`func (o *EeGroupProfileData) GetRestrictedEventTypes() []EventType`

GetRestrictedEventTypes returns the RestrictedEventTypes field if non-nil, zero value otherwise.

### GetRestrictedEventTypesOk

`func (o *EeGroupProfileData) GetRestrictedEventTypesOk() (*[]EventType, bool)`

GetRestrictedEventTypesOk returns a tuple with the RestrictedEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedEventTypes

`func (o *EeGroupProfileData) SetRestrictedEventTypes(v []EventType)`

SetRestrictedEventTypes sets RestrictedEventTypes field to given value.

### HasRestrictedEventTypes

`func (o *EeGroupProfileData) HasRestrictedEventTypes() bool`

HasRestrictedEventTypes returns a boolean if a field has been set.

### GetAllowedMtcProvider

`func (o *EeGroupProfileData) GetAllowedMtcProvider() map[string][]MtcProvider`

GetAllowedMtcProvider returns the AllowedMtcProvider field if non-nil, zero value otherwise.

### GetAllowedMtcProviderOk

`func (o *EeGroupProfileData) GetAllowedMtcProviderOk() (*map[string][]MtcProvider, bool)`

GetAllowedMtcProviderOk returns a tuple with the AllowedMtcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMtcProvider

`func (o *EeGroupProfileData) SetAllowedMtcProvider(v map[string][]MtcProvider)`

SetAllowedMtcProvider sets AllowedMtcProvider field to given value.

### HasAllowedMtcProvider

`func (o *EeGroupProfileData) HasAllowedMtcProvider() bool`

HasAllowedMtcProvider returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeGroupProfileData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeGroupProfileData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeGroupProfileData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeGroupProfileData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetIwkEpcRestricted

`func (o *EeGroupProfileData) GetIwkEpcRestricted() bool`

GetIwkEpcRestricted returns the IwkEpcRestricted field if non-nil, zero value otherwise.

### GetIwkEpcRestrictedOk

`func (o *EeGroupProfileData) GetIwkEpcRestrictedOk() (*bool, bool)`

GetIwkEpcRestrictedOk returns a tuple with the IwkEpcRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpcRestricted

`func (o *EeGroupProfileData) SetIwkEpcRestricted(v bool)`

SetIwkEpcRestricted sets IwkEpcRestricted field to given value.

### HasIwkEpcRestricted

`func (o *EeGroupProfileData) HasIwkEpcRestricted() bool`

HasIwkEpcRestricted returns a boolean if a field has been set.

### GetExtGroupId

`func (o *EeGroupProfileData) GetExtGroupId() string`

GetExtGroupId returns the ExtGroupId field if non-nil, zero value otherwise.

### GetExtGroupIdOk

`func (o *EeGroupProfileData) GetExtGroupIdOk() (*string, bool)`

GetExtGroupIdOk returns a tuple with the ExtGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGroupId

`func (o *EeGroupProfileData) SetExtGroupId(v string)`

SetExtGroupId sets ExtGroupId field to given value.

### HasExtGroupId

`func (o *EeGroupProfileData) HasExtGroupId() bool`

HasExtGroupId returns a boolean if a field has been set.

### GetHssGroupId

`func (o *EeGroupProfileData) GetHssGroupId() string`

GetHssGroupId returns the HssGroupId field if non-nil, zero value otherwise.

### GetHssGroupIdOk

`func (o *EeGroupProfileData) GetHssGroupIdOk() (*string, bool)`

GetHssGroupIdOk returns a tuple with the HssGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssGroupId

`func (o *EeGroupProfileData) SetHssGroupId(v string)`

SetHssGroupId sets HssGroupId field to given value.

### HasHssGroupId

`func (o *EeGroupProfileData) HasHssGroupId() bool`

HasHssGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


