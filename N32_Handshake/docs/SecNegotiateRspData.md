# SecNegotiateRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | Fully Qualified Domain Name | 
**SelectedSecCapability** | [**SecurityCapability**](SecurityCapability.md) |  | 
**Var3GppSbiTargetApiRootSupported** | Pointer to **bool** |  | [optional] [default to false]
**PlmnIdList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**SnpnIdList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**AllowedUsagePurpose** | Pointer to [**[]IntendedN32Purpose**](IntendedN32Purpose.md) |  | [optional] 
**RejectedUsagePurpose** | Pointer to [**[]IntendedN32Purpose**](IntendedN32Purpose.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SenderN32fFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**SenderN32fPort** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewSecNegotiateRspData

`func NewSecNegotiateRspData(sender string, selectedSecCapability SecurityCapability, ) *SecNegotiateRspData`

NewSecNegotiateRspData instantiates a new SecNegotiateRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecNegotiateRspDataWithDefaults

`func NewSecNegotiateRspDataWithDefaults() *SecNegotiateRspData`

NewSecNegotiateRspDataWithDefaults instantiates a new SecNegotiateRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *SecNegotiateRspData) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SecNegotiateRspData) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SecNegotiateRspData) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSelectedSecCapability

`func (o *SecNegotiateRspData) GetSelectedSecCapability() SecurityCapability`

GetSelectedSecCapability returns the SelectedSecCapability field if non-nil, zero value otherwise.

### GetSelectedSecCapabilityOk

`func (o *SecNegotiateRspData) GetSelectedSecCapabilityOk() (*SecurityCapability, bool)`

GetSelectedSecCapabilityOk returns a tuple with the SelectedSecCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSecCapability

`func (o *SecNegotiateRspData) SetSelectedSecCapability(v SecurityCapability)`

SetSelectedSecCapability sets SelectedSecCapability field to given value.


### GetVar3GppSbiTargetApiRootSupported

`func (o *SecNegotiateRspData) GetVar3GppSbiTargetApiRootSupported() bool`

GetVar3GppSbiTargetApiRootSupported returns the Var3GppSbiTargetApiRootSupported field if non-nil, zero value otherwise.

### GetVar3GppSbiTargetApiRootSupportedOk

`func (o *SecNegotiateRspData) GetVar3GppSbiTargetApiRootSupportedOk() (*bool, bool)`

GetVar3GppSbiTargetApiRootSupportedOk returns a tuple with the Var3GppSbiTargetApiRootSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3GppSbiTargetApiRootSupported

`func (o *SecNegotiateRspData) SetVar3GppSbiTargetApiRootSupported(v bool)`

SetVar3GppSbiTargetApiRootSupported sets Var3GppSbiTargetApiRootSupported field to given value.

### HasVar3GppSbiTargetApiRootSupported

`func (o *SecNegotiateRspData) HasVar3GppSbiTargetApiRootSupported() bool`

HasVar3GppSbiTargetApiRootSupported returns a boolean if a field has been set.

### GetPlmnIdList

`func (o *SecNegotiateRspData) GetPlmnIdList() []PlmnId`

GetPlmnIdList returns the PlmnIdList field if non-nil, zero value otherwise.

### GetPlmnIdListOk

`func (o *SecNegotiateRspData) GetPlmnIdListOk() (*[]PlmnId, bool)`

GetPlmnIdListOk returns a tuple with the PlmnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnIdList

`func (o *SecNegotiateRspData) SetPlmnIdList(v []PlmnId)`

SetPlmnIdList sets PlmnIdList field to given value.

### HasPlmnIdList

`func (o *SecNegotiateRspData) HasPlmnIdList() bool`

HasPlmnIdList returns a boolean if a field has been set.

### GetSnpnIdList

`func (o *SecNegotiateRspData) GetSnpnIdList() []PlmnIdNid`

GetSnpnIdList returns the SnpnIdList field if non-nil, zero value otherwise.

### GetSnpnIdListOk

`func (o *SecNegotiateRspData) GetSnpnIdListOk() (*[]PlmnIdNid, bool)`

GetSnpnIdListOk returns a tuple with the SnpnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpnIdList

`func (o *SecNegotiateRspData) SetSnpnIdList(v []PlmnIdNid)`

SetSnpnIdList sets SnpnIdList field to given value.

### HasSnpnIdList

`func (o *SecNegotiateRspData) HasSnpnIdList() bool`

HasSnpnIdList returns a boolean if a field has been set.

### GetAllowedUsagePurpose

`func (o *SecNegotiateRspData) GetAllowedUsagePurpose() []IntendedN32Purpose`

GetAllowedUsagePurpose returns the AllowedUsagePurpose field if non-nil, zero value otherwise.

### GetAllowedUsagePurposeOk

`func (o *SecNegotiateRspData) GetAllowedUsagePurposeOk() (*[]IntendedN32Purpose, bool)`

GetAllowedUsagePurposeOk returns a tuple with the AllowedUsagePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsagePurpose

`func (o *SecNegotiateRspData) SetAllowedUsagePurpose(v []IntendedN32Purpose)`

SetAllowedUsagePurpose sets AllowedUsagePurpose field to given value.

### HasAllowedUsagePurpose

`func (o *SecNegotiateRspData) HasAllowedUsagePurpose() bool`

HasAllowedUsagePurpose returns a boolean if a field has been set.

### GetRejectedUsagePurpose

`func (o *SecNegotiateRspData) GetRejectedUsagePurpose() []IntendedN32Purpose`

GetRejectedUsagePurpose returns the RejectedUsagePurpose field if non-nil, zero value otherwise.

### GetRejectedUsagePurposeOk

`func (o *SecNegotiateRspData) GetRejectedUsagePurposeOk() (*[]IntendedN32Purpose, bool)`

GetRejectedUsagePurposeOk returns a tuple with the RejectedUsagePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedUsagePurpose

`func (o *SecNegotiateRspData) SetRejectedUsagePurpose(v []IntendedN32Purpose)`

SetRejectedUsagePurpose sets RejectedUsagePurpose field to given value.

### HasRejectedUsagePurpose

`func (o *SecNegotiateRspData) HasRejectedUsagePurpose() bool`

HasRejectedUsagePurpose returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SecNegotiateRspData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SecNegotiateRspData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SecNegotiateRspData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SecNegotiateRspData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSenderN32fFqdn

`func (o *SecNegotiateRspData) GetSenderN32fFqdn() string`

GetSenderN32fFqdn returns the SenderN32fFqdn field if non-nil, zero value otherwise.

### GetSenderN32fFqdnOk

`func (o *SecNegotiateRspData) GetSenderN32fFqdnOk() (*string, bool)`

GetSenderN32fFqdnOk returns a tuple with the SenderN32fFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderN32fFqdn

`func (o *SecNegotiateRspData) SetSenderN32fFqdn(v string)`

SetSenderN32fFqdn sets SenderN32fFqdn field to given value.

### HasSenderN32fFqdn

`func (o *SecNegotiateRspData) HasSenderN32fFqdn() bool`

HasSenderN32fFqdn returns a boolean if a field has been set.

### GetSenderN32fPort

`func (o *SecNegotiateRspData) GetSenderN32fPort() int32`

GetSenderN32fPort returns the SenderN32fPort field if non-nil, zero value otherwise.

### GetSenderN32fPortOk

`func (o *SecNegotiateRspData) GetSenderN32fPortOk() (*int32, bool)`

GetSenderN32fPortOk returns a tuple with the SenderN32fPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderN32fPort

`func (o *SecNegotiateRspData) SetSenderN32fPort(v int32)`

SetSenderN32fPort sets SenderN32fPort field to given value.

### HasSenderN32fPort

`func (o *SecNegotiateRspData) HasSenderN32fPort() bool`

HasSenderN32fPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


