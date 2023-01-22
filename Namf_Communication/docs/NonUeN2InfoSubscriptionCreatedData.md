# NonUeN2InfoSubscriptionCreatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N2NotifySubscriptionId** | **string** |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**N2InformationClass** | Pointer to [**N2InformationClass**](N2InformationClass.md) |  | [optional] 

## Methods

### NewNonUeN2InfoSubscriptionCreatedData

`func NewNonUeN2InfoSubscriptionCreatedData(n2NotifySubscriptionId string, ) *NonUeN2InfoSubscriptionCreatedData`

NewNonUeN2InfoSubscriptionCreatedData instantiates a new NonUeN2InfoSubscriptionCreatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonUeN2InfoSubscriptionCreatedDataWithDefaults

`func NewNonUeN2InfoSubscriptionCreatedDataWithDefaults() *NonUeN2InfoSubscriptionCreatedData`

NewNonUeN2InfoSubscriptionCreatedDataWithDefaults instantiates a new NonUeN2InfoSubscriptionCreatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN2NotifySubscriptionId

`func (o *NonUeN2InfoSubscriptionCreatedData) GetN2NotifySubscriptionId() string`

GetN2NotifySubscriptionId returns the N2NotifySubscriptionId field if non-nil, zero value otherwise.

### GetN2NotifySubscriptionIdOk

`func (o *NonUeN2InfoSubscriptionCreatedData) GetN2NotifySubscriptionIdOk() (*string, bool)`

GetN2NotifySubscriptionIdOk returns a tuple with the N2NotifySubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2NotifySubscriptionId

`func (o *NonUeN2InfoSubscriptionCreatedData) SetN2NotifySubscriptionId(v string)`

SetN2NotifySubscriptionId sets N2NotifySubscriptionId field to given value.


### GetSupportedFeatures

`func (o *NonUeN2InfoSubscriptionCreatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NonUeN2InfoSubscriptionCreatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NonUeN2InfoSubscriptionCreatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NonUeN2InfoSubscriptionCreatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetN2InformationClass

`func (o *NonUeN2InfoSubscriptionCreatedData) GetN2InformationClass() N2InformationClass`

GetN2InformationClass returns the N2InformationClass field if non-nil, zero value otherwise.

### GetN2InformationClassOk

`func (o *NonUeN2InfoSubscriptionCreatedData) GetN2InformationClassOk() (*N2InformationClass, bool)`

GetN2InformationClassOk returns a tuple with the N2InformationClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InformationClass

`func (o *NonUeN2InfoSubscriptionCreatedData) SetN2InformationClass(v N2InformationClass)`

SetN2InformationClass sets N2InformationClass field to given value.

### HasN2InformationClass

`func (o *NonUeN2InfoSubscriptionCreatedData) HasN2InformationClass() bool`

HasN2InformationClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


