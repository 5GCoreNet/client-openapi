# ConfigureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UassId** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**UasId** | [**UasId**](UasId.md) |  | 
**AllowedC2CommModes** | [**[]C2CommMode**](C2CommMode.md) |  | 
**C2CommModeSwitchTypes** | [**[]C2CommModeSwitching**](C2CommModeSwitching.md) |  | 
**NotificationUri** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**PrimaryC2CommMode** | [**C2CommMode**](C2CommMode.md) |  | 
**SecondaryC2CommMode** | Pointer to [**C2CommMode**](C2CommMode.md) |  | [optional] 
**C2SwitchPolicies** | [**C2SwitchPolicies**](C2SwitchPolicies.md) |  | 
**C2ServiceArea** | Pointer to [**C2ServiceArea**](C2ServiceArea.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewConfigureData

`func NewConfigureData(uassId string, uasId UasId, allowedC2CommModes []C2CommMode, c2CommModeSwitchTypes []C2CommModeSwitching, notificationUri string, primaryC2CommMode C2CommMode, c2SwitchPolicies C2SwitchPolicies, ) *ConfigureData`

NewConfigureData instantiates a new ConfigureData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureDataWithDefaults

`func NewConfigureDataWithDefaults() *ConfigureData`

NewConfigureDataWithDefaults instantiates a new ConfigureData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUassId

`func (o *ConfigureData) GetUassId() string`

GetUassId returns the UassId field if non-nil, zero value otherwise.

### GetUassIdOk

`func (o *ConfigureData) GetUassIdOk() (*string, bool)`

GetUassIdOk returns a tuple with the UassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUassId

`func (o *ConfigureData) SetUassId(v string)`

SetUassId sets UassId field to given value.


### GetUasId

`func (o *ConfigureData) GetUasId() UasId`

GetUasId returns the UasId field if non-nil, zero value otherwise.

### GetUasIdOk

`func (o *ConfigureData) GetUasIdOk() (*UasId, bool)`

GetUasIdOk returns a tuple with the UasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUasId

`func (o *ConfigureData) SetUasId(v UasId)`

SetUasId sets UasId field to given value.


### GetAllowedC2CommModes

`func (o *ConfigureData) GetAllowedC2CommModes() []C2CommMode`

GetAllowedC2CommModes returns the AllowedC2CommModes field if non-nil, zero value otherwise.

### GetAllowedC2CommModesOk

`func (o *ConfigureData) GetAllowedC2CommModesOk() (*[]C2CommMode, bool)`

GetAllowedC2CommModesOk returns a tuple with the AllowedC2CommModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedC2CommModes

`func (o *ConfigureData) SetAllowedC2CommModes(v []C2CommMode)`

SetAllowedC2CommModes sets AllowedC2CommModes field to given value.


### GetC2CommModeSwitchTypes

`func (o *ConfigureData) GetC2CommModeSwitchTypes() []C2CommModeSwitching`

GetC2CommModeSwitchTypes returns the C2CommModeSwitchTypes field if non-nil, zero value otherwise.

### GetC2CommModeSwitchTypesOk

`func (o *ConfigureData) GetC2CommModeSwitchTypesOk() (*[]C2CommModeSwitching, bool)`

GetC2CommModeSwitchTypesOk returns a tuple with the C2CommModeSwitchTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2CommModeSwitchTypes

`func (o *ConfigureData) SetC2CommModeSwitchTypes(v []C2CommModeSwitching)`

SetC2CommModeSwitchTypes sets C2CommModeSwitchTypes field to given value.


### GetNotificationUri

`func (o *ConfigureData) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *ConfigureData) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *ConfigureData) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetPrimaryC2CommMode

`func (o *ConfigureData) GetPrimaryC2CommMode() C2CommMode`

GetPrimaryC2CommMode returns the PrimaryC2CommMode field if non-nil, zero value otherwise.

### GetPrimaryC2CommModeOk

`func (o *ConfigureData) GetPrimaryC2CommModeOk() (*C2CommMode, bool)`

GetPrimaryC2CommModeOk returns a tuple with the PrimaryC2CommMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryC2CommMode

`func (o *ConfigureData) SetPrimaryC2CommMode(v C2CommMode)`

SetPrimaryC2CommMode sets PrimaryC2CommMode field to given value.


### GetSecondaryC2CommMode

`func (o *ConfigureData) GetSecondaryC2CommMode() C2CommMode`

GetSecondaryC2CommMode returns the SecondaryC2CommMode field if non-nil, zero value otherwise.

### GetSecondaryC2CommModeOk

`func (o *ConfigureData) GetSecondaryC2CommModeOk() (*C2CommMode, bool)`

GetSecondaryC2CommModeOk returns a tuple with the SecondaryC2CommMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryC2CommMode

`func (o *ConfigureData) SetSecondaryC2CommMode(v C2CommMode)`

SetSecondaryC2CommMode sets SecondaryC2CommMode field to given value.

### HasSecondaryC2CommMode

`func (o *ConfigureData) HasSecondaryC2CommMode() bool`

HasSecondaryC2CommMode returns a boolean if a field has been set.

### GetC2SwitchPolicies

`func (o *ConfigureData) GetC2SwitchPolicies() C2SwitchPolicies`

GetC2SwitchPolicies returns the C2SwitchPolicies field if non-nil, zero value otherwise.

### GetC2SwitchPoliciesOk

`func (o *ConfigureData) GetC2SwitchPoliciesOk() (*C2SwitchPolicies, bool)`

GetC2SwitchPoliciesOk returns a tuple with the C2SwitchPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2SwitchPolicies

`func (o *ConfigureData) SetC2SwitchPolicies(v C2SwitchPolicies)`

SetC2SwitchPolicies sets C2SwitchPolicies field to given value.


### GetC2ServiceArea

`func (o *ConfigureData) GetC2ServiceArea() C2ServiceArea`

GetC2ServiceArea returns the C2ServiceArea field if non-nil, zero value otherwise.

### GetC2ServiceAreaOk

`func (o *ConfigureData) GetC2ServiceAreaOk() (*C2ServiceArea, bool)`

GetC2ServiceAreaOk returns a tuple with the C2ServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2ServiceArea

`func (o *ConfigureData) SetC2ServiceArea(v C2ServiceArea)`

SetC2ServiceArea sets C2ServiceArea field to given value.

### HasC2ServiceArea

`func (o *ConfigureData) HasC2ServiceArea() bool`

HasC2ServiceArea returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ConfigureData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ConfigureData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ConfigureData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ConfigureData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


