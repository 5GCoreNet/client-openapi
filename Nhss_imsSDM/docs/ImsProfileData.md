# ImsProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImsServiceProfiles** | [**[]ImsServiceProfile**](ImsServiceProfile.md) |  | 
**ChargingInfo** | Pointer to [**ChargingInfo**](ChargingInfo.md) |  | [optional] 
**ServiceLevelTraceInfo** | Pointer to [**ServiceLevelTraceInformation**](ServiceLevelTraceInformation.md) |  | [optional] 
**ServicePriorityLevelList** | Pointer to **[]string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MaxAllowedSimulReg** | Pointer to **int32** |  | [optional] 
**ServicePriorityLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewImsProfileData

`func NewImsProfileData(imsServiceProfiles []ImsServiceProfile, ) *ImsProfileData`

NewImsProfileData instantiates a new ImsProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImsProfileDataWithDefaults

`func NewImsProfileDataWithDefaults() *ImsProfileData`

NewImsProfileDataWithDefaults instantiates a new ImsProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImsServiceProfiles

`func (o *ImsProfileData) GetImsServiceProfiles() []ImsServiceProfile`

GetImsServiceProfiles returns the ImsServiceProfiles field if non-nil, zero value otherwise.

### GetImsServiceProfilesOk

`func (o *ImsProfileData) GetImsServiceProfilesOk() (*[]ImsServiceProfile, bool)`

GetImsServiceProfilesOk returns a tuple with the ImsServiceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsServiceProfiles

`func (o *ImsProfileData) SetImsServiceProfiles(v []ImsServiceProfile)`

SetImsServiceProfiles sets ImsServiceProfiles field to given value.


### GetChargingInfo

`func (o *ImsProfileData) GetChargingInfo() ChargingInfo`

GetChargingInfo returns the ChargingInfo field if non-nil, zero value otherwise.

### GetChargingInfoOk

`func (o *ImsProfileData) GetChargingInfoOk() (*ChargingInfo, bool)`

GetChargingInfoOk returns a tuple with the ChargingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingInfo

`func (o *ImsProfileData) SetChargingInfo(v ChargingInfo)`

SetChargingInfo sets ChargingInfo field to given value.

### HasChargingInfo

`func (o *ImsProfileData) HasChargingInfo() bool`

HasChargingInfo returns a boolean if a field has been set.

### GetServiceLevelTraceInfo

`func (o *ImsProfileData) GetServiceLevelTraceInfo() ServiceLevelTraceInformation`

GetServiceLevelTraceInfo returns the ServiceLevelTraceInfo field if non-nil, zero value otherwise.

### GetServiceLevelTraceInfoOk

`func (o *ImsProfileData) GetServiceLevelTraceInfoOk() (*ServiceLevelTraceInformation, bool)`

GetServiceLevelTraceInfoOk returns a tuple with the ServiceLevelTraceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelTraceInfo

`func (o *ImsProfileData) SetServiceLevelTraceInfo(v ServiceLevelTraceInformation)`

SetServiceLevelTraceInfo sets ServiceLevelTraceInfo field to given value.

### HasServiceLevelTraceInfo

`func (o *ImsProfileData) HasServiceLevelTraceInfo() bool`

HasServiceLevelTraceInfo returns a boolean if a field has been set.

### GetServicePriorityLevelList

`func (o *ImsProfileData) GetServicePriorityLevelList() []string`

GetServicePriorityLevelList returns the ServicePriorityLevelList field if non-nil, zero value otherwise.

### GetServicePriorityLevelListOk

`func (o *ImsProfileData) GetServicePriorityLevelListOk() (*[]string, bool)`

GetServicePriorityLevelListOk returns a tuple with the ServicePriorityLevelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePriorityLevelList

`func (o *ImsProfileData) SetServicePriorityLevelList(v []string)`

SetServicePriorityLevelList sets ServicePriorityLevelList field to given value.

### HasServicePriorityLevelList

`func (o *ImsProfileData) HasServicePriorityLevelList() bool`

HasServicePriorityLevelList returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ImsProfileData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ImsProfileData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ImsProfileData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ImsProfileData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMaxAllowedSimulReg

`func (o *ImsProfileData) GetMaxAllowedSimulReg() int32`

GetMaxAllowedSimulReg returns the MaxAllowedSimulReg field if non-nil, zero value otherwise.

### GetMaxAllowedSimulRegOk

`func (o *ImsProfileData) GetMaxAllowedSimulRegOk() (*int32, bool)`

GetMaxAllowedSimulRegOk returns a tuple with the MaxAllowedSimulReg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedSimulReg

`func (o *ImsProfileData) SetMaxAllowedSimulReg(v int32)`

SetMaxAllowedSimulReg sets MaxAllowedSimulReg field to given value.

### HasMaxAllowedSimulReg

`func (o *ImsProfileData) HasMaxAllowedSimulReg() bool`

HasMaxAllowedSimulReg returns a boolean if a field has been set.

### GetServicePriorityLevel

`func (o *ImsProfileData) GetServicePriorityLevel() int32`

GetServicePriorityLevel returns the ServicePriorityLevel field if non-nil, zero value otherwise.

### GetServicePriorityLevelOk

`func (o *ImsProfileData) GetServicePriorityLevelOk() (*int32, bool)`

GetServicePriorityLevelOk returns a tuple with the ServicePriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePriorityLevel

`func (o *ImsProfileData) SetServicePriorityLevel(v int32)`

SetServicePriorityLevel sets ServicePriorityLevel field to given value.

### HasServicePriorityLevel

`func (o *ImsProfileData) HasServicePriorityLevel() bool`

HasServicePriorityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


