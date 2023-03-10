# NonUeN2InfoSubscriptionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalRanNodeList** | Pointer to [**[]GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**AnTypeList** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**N2InformationClass** | [**N2InformationClass**](N2InformationClass.md) |  | 
**N2NotifyCallbackUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewNonUeN2InfoSubscriptionCreateData

`func NewNonUeN2InfoSubscriptionCreateData(n2InformationClass N2InformationClass, n2NotifyCallbackUri string, ) *NonUeN2InfoSubscriptionCreateData`

NewNonUeN2InfoSubscriptionCreateData instantiates a new NonUeN2InfoSubscriptionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonUeN2InfoSubscriptionCreateDataWithDefaults

`func NewNonUeN2InfoSubscriptionCreateDataWithDefaults() *NonUeN2InfoSubscriptionCreateData`

NewNonUeN2InfoSubscriptionCreateDataWithDefaults instantiates a new NonUeN2InfoSubscriptionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalRanNodeList

`func (o *NonUeN2InfoSubscriptionCreateData) GetGlobalRanNodeList() []GlobalRanNodeId`

GetGlobalRanNodeList returns the GlobalRanNodeList field if non-nil, zero value otherwise.

### GetGlobalRanNodeListOk

`func (o *NonUeN2InfoSubscriptionCreateData) GetGlobalRanNodeListOk() (*[]GlobalRanNodeId, bool)`

GetGlobalRanNodeListOk returns a tuple with the GlobalRanNodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRanNodeList

`func (o *NonUeN2InfoSubscriptionCreateData) SetGlobalRanNodeList(v []GlobalRanNodeId)`

SetGlobalRanNodeList sets GlobalRanNodeList field to given value.

### HasGlobalRanNodeList

`func (o *NonUeN2InfoSubscriptionCreateData) HasGlobalRanNodeList() bool`

HasGlobalRanNodeList returns a boolean if a field has been set.

### GetAnTypeList

`func (o *NonUeN2InfoSubscriptionCreateData) GetAnTypeList() []AccessType`

GetAnTypeList returns the AnTypeList field if non-nil, zero value otherwise.

### GetAnTypeListOk

`func (o *NonUeN2InfoSubscriptionCreateData) GetAnTypeListOk() (*[]AccessType, bool)`

GetAnTypeListOk returns a tuple with the AnTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnTypeList

`func (o *NonUeN2InfoSubscriptionCreateData) SetAnTypeList(v []AccessType)`

SetAnTypeList sets AnTypeList field to given value.

### HasAnTypeList

`func (o *NonUeN2InfoSubscriptionCreateData) HasAnTypeList() bool`

HasAnTypeList returns a boolean if a field has been set.

### GetN2InformationClass

`func (o *NonUeN2InfoSubscriptionCreateData) GetN2InformationClass() N2InformationClass`

GetN2InformationClass returns the N2InformationClass field if non-nil, zero value otherwise.

### GetN2InformationClassOk

`func (o *NonUeN2InfoSubscriptionCreateData) GetN2InformationClassOk() (*N2InformationClass, bool)`

GetN2InformationClassOk returns a tuple with the N2InformationClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InformationClass

`func (o *NonUeN2InfoSubscriptionCreateData) SetN2InformationClass(v N2InformationClass)`

SetN2InformationClass sets N2InformationClass field to given value.


### GetN2NotifyCallbackUri

`func (o *NonUeN2InfoSubscriptionCreateData) GetN2NotifyCallbackUri() string`

GetN2NotifyCallbackUri returns the N2NotifyCallbackUri field if non-nil, zero value otherwise.

### GetN2NotifyCallbackUriOk

`func (o *NonUeN2InfoSubscriptionCreateData) GetN2NotifyCallbackUriOk() (*string, bool)`

GetN2NotifyCallbackUriOk returns a tuple with the N2NotifyCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2NotifyCallbackUri

`func (o *NonUeN2InfoSubscriptionCreateData) SetN2NotifyCallbackUri(v string)`

SetN2NotifyCallbackUri sets N2NotifyCallbackUri field to given value.


### GetNfId

`func (o *NonUeN2InfoSubscriptionCreateData) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *NonUeN2InfoSubscriptionCreateData) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *NonUeN2InfoSubscriptionCreateData) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *NonUeN2InfoSubscriptionCreateData) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NonUeN2InfoSubscriptionCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NonUeN2InfoSubscriptionCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NonUeN2InfoSubscriptionCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NonUeN2InfoSubscriptionCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


