# ECRData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 
**VisitedPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**EcrDataWbs** | Pointer to [**[]PlmnEcRestrictionDataWb**](PlmnEcRestrictionDataWb.md) |  | [optional] 
**RestrictedPlmnIds** | Pointer to [**[]PlmnId**](PlmnId.md) | Indicates a complete list (and possibly empty) of serving PLMNs where Enhanced Coverage shall be restricted. | [optional] 
**AllowedPlmnIds** | Pointer to [**[]PlmnId**](PlmnId.md) | Indicates a complete list (and possibly empty) of serving PLMNs where Enhanced Coverage shall be allowed. | [optional] 

## Methods

### NewECRData

`func NewECRData(supportedFeatures string, ) *ECRData`

NewECRData instantiates a new ECRData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECRDataWithDefaults

`func NewECRDataWithDefaults() *ECRData`

NewECRDataWithDefaults instantiates a new ECRData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *ECRData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ECRData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ECRData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.


### GetVisitedPlmnId

`func (o *ECRData) GetVisitedPlmnId() PlmnId`

GetVisitedPlmnId returns the VisitedPlmnId field if non-nil, zero value otherwise.

### GetVisitedPlmnIdOk

`func (o *ECRData) GetVisitedPlmnIdOk() (*PlmnId, bool)`

GetVisitedPlmnIdOk returns a tuple with the VisitedPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedPlmnId

`func (o *ECRData) SetVisitedPlmnId(v PlmnId)`

SetVisitedPlmnId sets VisitedPlmnId field to given value.

### HasVisitedPlmnId

`func (o *ECRData) HasVisitedPlmnId() bool`

HasVisitedPlmnId returns a boolean if a field has been set.

### GetEcrDataWbs

`func (o *ECRData) GetEcrDataWbs() []PlmnEcRestrictionDataWb`

GetEcrDataWbs returns the EcrDataWbs field if non-nil, zero value otherwise.

### GetEcrDataWbsOk

`func (o *ECRData) GetEcrDataWbsOk() (*[]PlmnEcRestrictionDataWb, bool)`

GetEcrDataWbsOk returns a tuple with the EcrDataWbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcrDataWbs

`func (o *ECRData) SetEcrDataWbs(v []PlmnEcRestrictionDataWb)`

SetEcrDataWbs sets EcrDataWbs field to given value.

### HasEcrDataWbs

`func (o *ECRData) HasEcrDataWbs() bool`

HasEcrDataWbs returns a boolean if a field has been set.

### GetRestrictedPlmnIds

`func (o *ECRData) GetRestrictedPlmnIds() []PlmnId`

GetRestrictedPlmnIds returns the RestrictedPlmnIds field if non-nil, zero value otherwise.

### GetRestrictedPlmnIdsOk

`func (o *ECRData) GetRestrictedPlmnIdsOk() (*[]PlmnId, bool)`

GetRestrictedPlmnIdsOk returns a tuple with the RestrictedPlmnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedPlmnIds

`func (o *ECRData) SetRestrictedPlmnIds(v []PlmnId)`

SetRestrictedPlmnIds sets RestrictedPlmnIds field to given value.

### HasRestrictedPlmnIds

`func (o *ECRData) HasRestrictedPlmnIds() bool`

HasRestrictedPlmnIds returns a boolean if a field has been set.

### GetAllowedPlmnIds

`func (o *ECRData) GetAllowedPlmnIds() []PlmnId`

GetAllowedPlmnIds returns the AllowedPlmnIds field if non-nil, zero value otherwise.

### GetAllowedPlmnIdsOk

`func (o *ECRData) GetAllowedPlmnIdsOk() (*[]PlmnId, bool)`

GetAllowedPlmnIdsOk returns a tuple with the AllowedPlmnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmnIds

`func (o *ECRData) SetAllowedPlmnIds(v []PlmnId)`

SetAllowedPlmnIds sets AllowedPlmnIds field to given value.

### HasAllowedPlmnIds

`func (o *ECRData) HasAllowedPlmnIds() bool`

HasAllowedPlmnIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


