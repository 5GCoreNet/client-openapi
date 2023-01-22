# SmContextCreatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HSmfUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SmfUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**PduSessionId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**UpCnxState** | Pointer to [**UpCnxState**](UpCnxState.md) |  | [optional] 
**N2SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoType** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**AllocatedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 
**HoState** | Pointer to [**HoState**](HoState.md) |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**SmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SelectedSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SelectedOldSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**InterPlmnApiRoot** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewSmContextCreatedData

`func NewSmContextCreatedData() *SmContextCreatedData`

NewSmContextCreatedData instantiates a new SmContextCreatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextCreatedDataWithDefaults

`func NewSmContextCreatedDataWithDefaults() *SmContextCreatedData`

NewSmContextCreatedDataWithDefaults instantiates a new SmContextCreatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHSmfUri

`func (o *SmContextCreatedData) GetHSmfUri() string`

GetHSmfUri returns the HSmfUri field if non-nil, zero value otherwise.

### GetHSmfUriOk

`func (o *SmContextCreatedData) GetHSmfUriOk() (*string, bool)`

GetHSmfUriOk returns a tuple with the HSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfUri

`func (o *SmContextCreatedData) SetHSmfUri(v string)`

SetHSmfUri sets HSmfUri field to given value.

### HasHSmfUri

`func (o *SmContextCreatedData) HasHSmfUri() bool`

HasHSmfUri returns a boolean if a field has been set.

### GetSmfUri

`func (o *SmContextCreatedData) GetSmfUri() string`

GetSmfUri returns the SmfUri field if non-nil, zero value otherwise.

### GetSmfUriOk

`func (o *SmContextCreatedData) GetSmfUriOk() (*string, bool)`

GetSmfUriOk returns a tuple with the SmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfUri

`func (o *SmContextCreatedData) SetSmfUri(v string)`

SetSmfUri sets SmfUri field to given value.

### HasSmfUri

`func (o *SmContextCreatedData) HasSmfUri() bool`

HasSmfUri returns a boolean if a field has been set.

### GetPduSessionId

`func (o *SmContextCreatedData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmContextCreatedData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmContextCreatedData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *SmContextCreatedData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetSNssai

`func (o *SmContextCreatedData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SmContextCreatedData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SmContextCreatedData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *SmContextCreatedData) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetUpCnxState

`func (o *SmContextCreatedData) GetUpCnxState() UpCnxState`

GetUpCnxState returns the UpCnxState field if non-nil, zero value otherwise.

### GetUpCnxStateOk

`func (o *SmContextCreatedData) GetUpCnxStateOk() (*UpCnxState, bool)`

GetUpCnxStateOk returns a tuple with the UpCnxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpCnxState

`func (o *SmContextCreatedData) SetUpCnxState(v UpCnxState)`

SetUpCnxState sets UpCnxState field to given value.

### HasUpCnxState

`func (o *SmContextCreatedData) HasUpCnxState() bool`

HasUpCnxState returns a boolean if a field has been set.

### GetN2SmInfo

`func (o *SmContextCreatedData) GetN2SmInfo() RefToBinaryData`

GetN2SmInfo returns the N2SmInfo field if non-nil, zero value otherwise.

### GetN2SmInfoOk

`func (o *SmContextCreatedData) GetN2SmInfoOk() (*RefToBinaryData, bool)`

GetN2SmInfoOk returns a tuple with the N2SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfo

`func (o *SmContextCreatedData) SetN2SmInfo(v RefToBinaryData)`

SetN2SmInfo sets N2SmInfo field to given value.

### HasN2SmInfo

`func (o *SmContextCreatedData) HasN2SmInfo() bool`

HasN2SmInfo returns a boolean if a field has been set.

### GetN2SmInfoType

`func (o *SmContextCreatedData) GetN2SmInfoType() N2SmInfoType`

GetN2SmInfoType returns the N2SmInfoType field if non-nil, zero value otherwise.

### GetN2SmInfoTypeOk

`func (o *SmContextCreatedData) GetN2SmInfoTypeOk() (*N2SmInfoType, bool)`

GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoType

`func (o *SmContextCreatedData) SetN2SmInfoType(v N2SmInfoType)`

SetN2SmInfoType sets N2SmInfoType field to given value.

### HasN2SmInfoType

`func (o *SmContextCreatedData) HasN2SmInfoType() bool`

HasN2SmInfoType returns a boolean if a field has been set.

### GetAllocatedEbiList

`func (o *SmContextCreatedData) GetAllocatedEbiList() []EbiArpMapping`

GetAllocatedEbiList returns the AllocatedEbiList field if non-nil, zero value otherwise.

### GetAllocatedEbiListOk

`func (o *SmContextCreatedData) GetAllocatedEbiListOk() (*[]EbiArpMapping, bool)`

GetAllocatedEbiListOk returns a tuple with the AllocatedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedEbiList

`func (o *SmContextCreatedData) SetAllocatedEbiList(v []EbiArpMapping)`

SetAllocatedEbiList sets AllocatedEbiList field to given value.

### HasAllocatedEbiList

`func (o *SmContextCreatedData) HasAllocatedEbiList() bool`

HasAllocatedEbiList returns a boolean if a field has been set.

### GetHoState

`func (o *SmContextCreatedData) GetHoState() HoState`

GetHoState returns the HoState field if non-nil, zero value otherwise.

### GetHoStateOk

`func (o *SmContextCreatedData) GetHoStateOk() (*HoState, bool)`

GetHoStateOk returns a tuple with the HoState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoState

`func (o *SmContextCreatedData) SetHoState(v HoState)`

SetHoState sets HoState field to given value.

### HasHoState

`func (o *SmContextCreatedData) HasHoState() bool`

HasHoState returns a boolean if a field has been set.

### GetGpsi

`func (o *SmContextCreatedData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SmContextCreatedData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SmContextCreatedData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *SmContextCreatedData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSmfServiceInstanceId

`func (o *SmContextCreatedData) GetSmfServiceInstanceId() string`

GetSmfServiceInstanceId returns the SmfServiceInstanceId field if non-nil, zero value otherwise.

### GetSmfServiceInstanceIdOk

`func (o *SmContextCreatedData) GetSmfServiceInstanceIdOk() (*string, bool)`

GetSmfServiceInstanceIdOk returns a tuple with the SmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServiceInstanceId

`func (o *SmContextCreatedData) SetSmfServiceInstanceId(v string)`

SetSmfServiceInstanceId sets SmfServiceInstanceId field to given value.

### HasSmfServiceInstanceId

`func (o *SmContextCreatedData) HasSmfServiceInstanceId() bool`

HasSmfServiceInstanceId returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *SmContextCreatedData) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *SmContextCreatedData) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *SmContextCreatedData) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *SmContextCreatedData) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmContextCreatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmContextCreatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmContextCreatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmContextCreatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSelectedSmfId

`func (o *SmContextCreatedData) GetSelectedSmfId() string`

GetSelectedSmfId returns the SelectedSmfId field if non-nil, zero value otherwise.

### GetSelectedSmfIdOk

`func (o *SmContextCreatedData) GetSelectedSmfIdOk() (*string, bool)`

GetSelectedSmfIdOk returns a tuple with the SelectedSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSmfId

`func (o *SmContextCreatedData) SetSelectedSmfId(v string)`

SetSelectedSmfId sets SelectedSmfId field to given value.

### HasSelectedSmfId

`func (o *SmContextCreatedData) HasSelectedSmfId() bool`

HasSelectedSmfId returns a boolean if a field has been set.

### GetSelectedOldSmfId

`func (o *SmContextCreatedData) GetSelectedOldSmfId() string`

GetSelectedOldSmfId returns the SelectedOldSmfId field if non-nil, zero value otherwise.

### GetSelectedOldSmfIdOk

`func (o *SmContextCreatedData) GetSelectedOldSmfIdOk() (*string, bool)`

GetSelectedOldSmfIdOk returns a tuple with the SelectedOldSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOldSmfId

`func (o *SmContextCreatedData) SetSelectedOldSmfId(v string)`

SetSelectedOldSmfId sets SelectedOldSmfId field to given value.

### HasSelectedOldSmfId

`func (o *SmContextCreatedData) HasSelectedOldSmfId() bool`

HasSelectedOldSmfId returns a boolean if a field has been set.

### GetInterPlmnApiRoot

`func (o *SmContextCreatedData) GetInterPlmnApiRoot() string`

GetInterPlmnApiRoot returns the InterPlmnApiRoot field if non-nil, zero value otherwise.

### GetInterPlmnApiRootOk

`func (o *SmContextCreatedData) GetInterPlmnApiRootOk() (*string, bool)`

GetInterPlmnApiRootOk returns a tuple with the InterPlmnApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnApiRoot

`func (o *SmContextCreatedData) SetInterPlmnApiRoot(v string)`

SetInterPlmnApiRoot sets InterPlmnApiRoot field to given value.

### HasInterPlmnApiRoot

`func (o *SmContextCreatedData) HasInterPlmnApiRoot() bool`

HasInterPlmnApiRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


