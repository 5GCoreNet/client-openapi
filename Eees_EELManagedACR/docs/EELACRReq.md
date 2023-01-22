# EELACRReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**EasCharacs** | [**[]EasCharacteristics**](EasCharacteristics.md) |  | 
**AppCtxtStoreAddr** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewEELACRReq

`func NewEELACRReq(ueId string, easCharacs []EasCharacteristics, ) *EELACRReq`

NewEELACRReq instantiates a new EELACRReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEELACRReqWithDefaults

`func NewEELACRReqWithDefaults() *EELACRReq`

NewEELACRReqWithDefaults instantiates a new EELACRReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *EELACRReq) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *EELACRReq) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *EELACRReq) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetEasCharacs

`func (o *EELACRReq) GetEasCharacs() []EasCharacteristics`

GetEasCharacs returns the EasCharacs field if non-nil, zero value otherwise.

### GetEasCharacsOk

`func (o *EELACRReq) GetEasCharacsOk() (*[]EasCharacteristics, bool)`

GetEasCharacsOk returns a tuple with the EasCharacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasCharacs

`func (o *EELACRReq) SetEasCharacs(v []EasCharacteristics)`

SetEasCharacs sets EasCharacs field to given value.


### GetAppCtxtStoreAddr

`func (o *EELACRReq) GetAppCtxtStoreAddr() string`

GetAppCtxtStoreAddr returns the AppCtxtStoreAddr field if non-nil, zero value otherwise.

### GetAppCtxtStoreAddrOk

`func (o *EELACRReq) GetAppCtxtStoreAddrOk() (*string, bool)`

GetAppCtxtStoreAddrOk returns a tuple with the AppCtxtStoreAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCtxtStoreAddr

`func (o *EELACRReq) SetAppCtxtStoreAddr(v string)`

SetAppCtxtStoreAddr sets AppCtxtStoreAddr field to given value.

### HasAppCtxtStoreAddr

`func (o *EELACRReq) HasAppCtxtStoreAddr() bool`

HasAppCtxtStoreAddr returns a boolean if a field has been set.

### GetSuppFeat

`func (o *EELACRReq) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *EELACRReq) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *EELACRReq) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *EELACRReq) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


