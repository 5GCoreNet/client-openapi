# EeProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedEventTypes** | Pointer to [**[]EventType**](EventType.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AllowedMtcProvider** | Pointer to [**map[string][]MtcProvider**](array.md) | A map (list of key-value pairs where EventType serves as key) of MTC provider lists. In addition to defined EventTypes, the key value \&quot;ALL\&quot; may be used to identify a map entry which contains a list of MtcProviders that are allowed monitoring all Event Types. | [optional] 
**IwkEpcRestricted** | Pointer to **bool** |  | [optional] [default to false]
**Imsi** | Pointer to **string** |  | [optional] 
**HssGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 

## Methods

### NewEeProfileData

`func NewEeProfileData() *EeProfileData`

NewEeProfileData instantiates a new EeProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeProfileDataWithDefaults

`func NewEeProfileDataWithDefaults() *EeProfileData`

NewEeProfileDataWithDefaults instantiates a new EeProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedEventTypes

`func (o *EeProfileData) GetRestrictedEventTypes() []EventType`

GetRestrictedEventTypes returns the RestrictedEventTypes field if non-nil, zero value otherwise.

### GetRestrictedEventTypesOk

`func (o *EeProfileData) GetRestrictedEventTypesOk() (*[]EventType, bool)`

GetRestrictedEventTypesOk returns a tuple with the RestrictedEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedEventTypes

`func (o *EeProfileData) SetRestrictedEventTypes(v []EventType)`

SetRestrictedEventTypes sets RestrictedEventTypes field to given value.

### HasRestrictedEventTypes

`func (o *EeProfileData) HasRestrictedEventTypes() bool`

HasRestrictedEventTypes returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeProfileData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeProfileData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeProfileData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeProfileData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAllowedMtcProvider

`func (o *EeProfileData) GetAllowedMtcProvider() map[string][]MtcProvider`

GetAllowedMtcProvider returns the AllowedMtcProvider field if non-nil, zero value otherwise.

### GetAllowedMtcProviderOk

`func (o *EeProfileData) GetAllowedMtcProviderOk() (*map[string][]MtcProvider, bool)`

GetAllowedMtcProviderOk returns a tuple with the AllowedMtcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMtcProvider

`func (o *EeProfileData) SetAllowedMtcProvider(v map[string][]MtcProvider)`

SetAllowedMtcProvider sets AllowedMtcProvider field to given value.

### HasAllowedMtcProvider

`func (o *EeProfileData) HasAllowedMtcProvider() bool`

HasAllowedMtcProvider returns a boolean if a field has been set.

### GetIwkEpcRestricted

`func (o *EeProfileData) GetIwkEpcRestricted() bool`

GetIwkEpcRestricted returns the IwkEpcRestricted field if non-nil, zero value otherwise.

### GetIwkEpcRestrictedOk

`func (o *EeProfileData) GetIwkEpcRestrictedOk() (*bool, bool)`

GetIwkEpcRestrictedOk returns a tuple with the IwkEpcRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpcRestricted

`func (o *EeProfileData) SetIwkEpcRestricted(v bool)`

SetIwkEpcRestricted sets IwkEpcRestricted field to given value.

### HasIwkEpcRestricted

`func (o *EeProfileData) HasIwkEpcRestricted() bool`

HasIwkEpcRestricted returns a boolean if a field has been set.

### GetImsi

`func (o *EeProfileData) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *EeProfileData) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *EeProfileData) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *EeProfileData) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetHssGroupId

`func (o *EeProfileData) GetHssGroupId() string`

GetHssGroupId returns the HssGroupId field if non-nil, zero value otherwise.

### GetHssGroupIdOk

`func (o *EeProfileData) GetHssGroupIdOk() (*string, bool)`

GetHssGroupIdOk returns a tuple with the HssGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssGroupId

`func (o *EeProfileData) SetHssGroupId(v string)`

SetHssGroupId sets HssGroupId field to given value.

### HasHssGroupId

`func (o *EeProfileData) HasHssGroupId() bool`

HasHssGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


