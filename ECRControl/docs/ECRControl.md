# ECRControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 
**MtcProviderId** | Pointer to **string** | Identifies the MTC Service Provider and/or MTC Application. | [optional] 
**ScsAsId** | Pointer to **string** | Identifier of the SCS/AS. | [optional] 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**EcrDataWbs** | Pointer to [**[]PlmnEcRestrictionDataWb**](PlmnEcRestrictionDataWb.md) |  | [optional] 
**RestrictedPlmnIds** | Pointer to [**[]PlmnId**](PlmnId.md) | Indicates a complete list (and possibly empty) of serving PLMNs where Enhanced Coverage shall be restricted. This attribute shall not be present for the query custom operation. | [optional] 
**AllowedPlmnIds** | Pointer to [**[]PlmnId**](PlmnId.md) | Indicates a complete list (and possibly empty) of serving PLMNs where Enhanced Coverage shall be allowed. This attribute shall not be present for the query custom operation. | [optional] 

## Methods

### NewECRControl

`func NewECRControl(supportedFeatures string, ) *ECRControl`

NewECRControl instantiates a new ECRControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECRControlWithDefaults

`func NewECRControlWithDefaults() *ECRControl`

NewECRControlWithDefaults instantiates a new ECRControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *ECRControl) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ECRControl) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ECRControl) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.


### GetMtcProviderId

`func (o *ECRControl) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *ECRControl) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *ECRControl) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *ECRControl) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetScsAsId

`func (o *ECRControl) GetScsAsId() string`

GetScsAsId returns the ScsAsId field if non-nil, zero value otherwise.

### GetScsAsIdOk

`func (o *ECRControl) GetScsAsIdOk() (*string, bool)`

GetScsAsIdOk returns a tuple with the ScsAsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsAsId

`func (o *ECRControl) SetScsAsId(v string)`

SetScsAsId sets ScsAsId field to given value.

### HasScsAsId

`func (o *ECRControl) HasScsAsId() bool`

HasScsAsId returns a boolean if a field has been set.

### GetExternalId

`func (o *ECRControl) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ECRControl) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ECRControl) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ECRControl) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *ECRControl) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *ECRControl) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *ECRControl) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *ECRControl) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetEcrDataWbs

`func (o *ECRControl) GetEcrDataWbs() []PlmnEcRestrictionDataWb`

GetEcrDataWbs returns the EcrDataWbs field if non-nil, zero value otherwise.

### GetEcrDataWbsOk

`func (o *ECRControl) GetEcrDataWbsOk() (*[]PlmnEcRestrictionDataWb, bool)`

GetEcrDataWbsOk returns a tuple with the EcrDataWbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcrDataWbs

`func (o *ECRControl) SetEcrDataWbs(v []PlmnEcRestrictionDataWb)`

SetEcrDataWbs sets EcrDataWbs field to given value.

### HasEcrDataWbs

`func (o *ECRControl) HasEcrDataWbs() bool`

HasEcrDataWbs returns a boolean if a field has been set.

### GetRestrictedPlmnIds

`func (o *ECRControl) GetRestrictedPlmnIds() []PlmnId`

GetRestrictedPlmnIds returns the RestrictedPlmnIds field if non-nil, zero value otherwise.

### GetRestrictedPlmnIdsOk

`func (o *ECRControl) GetRestrictedPlmnIdsOk() (*[]PlmnId, bool)`

GetRestrictedPlmnIdsOk returns a tuple with the RestrictedPlmnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedPlmnIds

`func (o *ECRControl) SetRestrictedPlmnIds(v []PlmnId)`

SetRestrictedPlmnIds sets RestrictedPlmnIds field to given value.

### HasRestrictedPlmnIds

`func (o *ECRControl) HasRestrictedPlmnIds() bool`

HasRestrictedPlmnIds returns a boolean if a field has been set.

### GetAllowedPlmnIds

`func (o *ECRControl) GetAllowedPlmnIds() []PlmnId`

GetAllowedPlmnIds returns the AllowedPlmnIds field if non-nil, zero value otherwise.

### GetAllowedPlmnIdsOk

`func (o *ECRControl) GetAllowedPlmnIdsOk() (*[]PlmnId, bool)`

GetAllowedPlmnIdsOk returns a tuple with the AllowedPlmnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmnIds

`func (o *ECRControl) SetAllowedPlmnIds(v []PlmnId)`

SetAllowedPlmnIds sets AllowedPlmnIds field to given value.

### HasAllowedPlmnIds

`func (o *ECRControl) HasAllowedPlmnIds() bool`

HasAllowedPlmnIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


