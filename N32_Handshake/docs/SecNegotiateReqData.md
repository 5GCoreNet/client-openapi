# SecNegotiateReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | Fully Qualified Domain Name | 
**SupportedSecCapabilityList** | [**[]SecurityCapability**](SecurityCapability.md) |  | 
**Var3GppSbiTargetApiRootSupported** | Pointer to **bool** |  | [optional] [default to false]
**PlmnIdList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**SnpnIdList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**TargetPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**TargetSnpnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**IntendedUsagePurpose** | Pointer to [**[]IntendedN32Purpose**](IntendedN32Purpose.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SenderN32fFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**SenderN32fPort** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewSecNegotiateReqData

`func NewSecNegotiateReqData(sender string, supportedSecCapabilityList []SecurityCapability, ) *SecNegotiateReqData`

NewSecNegotiateReqData instantiates a new SecNegotiateReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecNegotiateReqDataWithDefaults

`func NewSecNegotiateReqDataWithDefaults() *SecNegotiateReqData`

NewSecNegotiateReqDataWithDefaults instantiates a new SecNegotiateReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *SecNegotiateReqData) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SecNegotiateReqData) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SecNegotiateReqData) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSupportedSecCapabilityList

`func (o *SecNegotiateReqData) GetSupportedSecCapabilityList() []SecurityCapability`

GetSupportedSecCapabilityList returns the SupportedSecCapabilityList field if non-nil, zero value otherwise.

### GetSupportedSecCapabilityListOk

`func (o *SecNegotiateReqData) GetSupportedSecCapabilityListOk() (*[]SecurityCapability, bool)`

GetSupportedSecCapabilityListOk returns a tuple with the SupportedSecCapabilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSecCapabilityList

`func (o *SecNegotiateReqData) SetSupportedSecCapabilityList(v []SecurityCapability)`

SetSupportedSecCapabilityList sets SupportedSecCapabilityList field to given value.


### GetVar3GppSbiTargetApiRootSupported

`func (o *SecNegotiateReqData) GetVar3GppSbiTargetApiRootSupported() bool`

GetVar3GppSbiTargetApiRootSupported returns the Var3GppSbiTargetApiRootSupported field if non-nil, zero value otherwise.

### GetVar3GppSbiTargetApiRootSupportedOk

`func (o *SecNegotiateReqData) GetVar3GppSbiTargetApiRootSupportedOk() (*bool, bool)`

GetVar3GppSbiTargetApiRootSupportedOk returns a tuple with the Var3GppSbiTargetApiRootSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3GppSbiTargetApiRootSupported

`func (o *SecNegotiateReqData) SetVar3GppSbiTargetApiRootSupported(v bool)`

SetVar3GppSbiTargetApiRootSupported sets Var3GppSbiTargetApiRootSupported field to given value.

### HasVar3GppSbiTargetApiRootSupported

`func (o *SecNegotiateReqData) HasVar3GppSbiTargetApiRootSupported() bool`

HasVar3GppSbiTargetApiRootSupported returns a boolean if a field has been set.

### GetPlmnIdList

`func (o *SecNegotiateReqData) GetPlmnIdList() []PlmnId`

GetPlmnIdList returns the PlmnIdList field if non-nil, zero value otherwise.

### GetPlmnIdListOk

`func (o *SecNegotiateReqData) GetPlmnIdListOk() (*[]PlmnId, bool)`

GetPlmnIdListOk returns a tuple with the PlmnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnIdList

`func (o *SecNegotiateReqData) SetPlmnIdList(v []PlmnId)`

SetPlmnIdList sets PlmnIdList field to given value.

### HasPlmnIdList

`func (o *SecNegotiateReqData) HasPlmnIdList() bool`

HasPlmnIdList returns a boolean if a field has been set.

### GetSnpnIdList

`func (o *SecNegotiateReqData) GetSnpnIdList() []PlmnIdNid`

GetSnpnIdList returns the SnpnIdList field if non-nil, zero value otherwise.

### GetSnpnIdListOk

`func (o *SecNegotiateReqData) GetSnpnIdListOk() (*[]PlmnIdNid, bool)`

GetSnpnIdListOk returns a tuple with the SnpnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpnIdList

`func (o *SecNegotiateReqData) SetSnpnIdList(v []PlmnIdNid)`

SetSnpnIdList sets SnpnIdList field to given value.

### HasSnpnIdList

`func (o *SecNegotiateReqData) HasSnpnIdList() bool`

HasSnpnIdList returns a boolean if a field has been set.

### GetTargetPlmnId

`func (o *SecNegotiateReqData) GetTargetPlmnId() PlmnId`

GetTargetPlmnId returns the TargetPlmnId field if non-nil, zero value otherwise.

### GetTargetPlmnIdOk

`func (o *SecNegotiateReqData) GetTargetPlmnIdOk() (*PlmnId, bool)`

GetTargetPlmnIdOk returns a tuple with the TargetPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlmnId

`func (o *SecNegotiateReqData) SetTargetPlmnId(v PlmnId)`

SetTargetPlmnId sets TargetPlmnId field to given value.

### HasTargetPlmnId

`func (o *SecNegotiateReqData) HasTargetPlmnId() bool`

HasTargetPlmnId returns a boolean if a field has been set.

### GetTargetSnpnId

`func (o *SecNegotiateReqData) GetTargetSnpnId() PlmnIdNid`

GetTargetSnpnId returns the TargetSnpnId field if non-nil, zero value otherwise.

### GetTargetSnpnIdOk

`func (o *SecNegotiateReqData) GetTargetSnpnIdOk() (*PlmnIdNid, bool)`

GetTargetSnpnIdOk returns a tuple with the TargetSnpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSnpnId

`func (o *SecNegotiateReqData) SetTargetSnpnId(v PlmnIdNid)`

SetTargetSnpnId sets TargetSnpnId field to given value.

### HasTargetSnpnId

`func (o *SecNegotiateReqData) HasTargetSnpnId() bool`

HasTargetSnpnId returns a boolean if a field has been set.

### GetIntendedUsagePurpose

`func (o *SecNegotiateReqData) GetIntendedUsagePurpose() []IntendedN32Purpose`

GetIntendedUsagePurpose returns the IntendedUsagePurpose field if non-nil, zero value otherwise.

### GetIntendedUsagePurposeOk

`func (o *SecNegotiateReqData) GetIntendedUsagePurposeOk() (*[]IntendedN32Purpose, bool)`

GetIntendedUsagePurposeOk returns a tuple with the IntendedUsagePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedUsagePurpose

`func (o *SecNegotiateReqData) SetIntendedUsagePurpose(v []IntendedN32Purpose)`

SetIntendedUsagePurpose sets IntendedUsagePurpose field to given value.

### HasIntendedUsagePurpose

`func (o *SecNegotiateReqData) HasIntendedUsagePurpose() bool`

HasIntendedUsagePurpose returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SecNegotiateReqData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SecNegotiateReqData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SecNegotiateReqData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SecNegotiateReqData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSenderN32fFqdn

`func (o *SecNegotiateReqData) GetSenderN32fFqdn() string`

GetSenderN32fFqdn returns the SenderN32fFqdn field if non-nil, zero value otherwise.

### GetSenderN32fFqdnOk

`func (o *SecNegotiateReqData) GetSenderN32fFqdnOk() (*string, bool)`

GetSenderN32fFqdnOk returns a tuple with the SenderN32fFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderN32fFqdn

`func (o *SecNegotiateReqData) SetSenderN32fFqdn(v string)`

SetSenderN32fFqdn sets SenderN32fFqdn field to given value.

### HasSenderN32fFqdn

`func (o *SecNegotiateReqData) HasSenderN32fFqdn() bool`

HasSenderN32fFqdn returns a boolean if a field has been set.

### GetSenderN32fPort

`func (o *SecNegotiateReqData) GetSenderN32fPort() int32`

GetSenderN32fPort returns the SenderN32fPort field if non-nil, zero value otherwise.

### GetSenderN32fPortOk

`func (o *SecNegotiateReqData) GetSenderN32fPortOk() (*int32, bool)`

GetSenderN32fPortOk returns a tuple with the SenderN32fPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderN32fPort

`func (o *SecNegotiateReqData) SetSenderN32fPort(v int32)`

SetSenderN32fPort sets SenderN32fPort field to given value.

### HasSenderN32fPort

`func (o *SecNegotiateReqData) HasSenderN32fPort() bool`

HasSenderN32fPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


