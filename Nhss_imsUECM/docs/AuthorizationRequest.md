# AuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impi** | Pointer to **string** | IMS Private Identity of the UE | [optional] 
**AuthorizationType** | [**AuthorizationType**](AuthorizationType.md) |  | 
**VisitedNetworkIdentifier** | Pointer to **string** |  | [optional] 
**EmergencyIndicator** | Pointer to **bool** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAuthorizationRequest

`func NewAuthorizationRequest(authorizationType AuthorizationType, ) *AuthorizationRequest`

NewAuthorizationRequest instantiates a new AuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationRequestWithDefaults

`func NewAuthorizationRequestWithDefaults() *AuthorizationRequest`

NewAuthorizationRequestWithDefaults instantiates a new AuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpi

`func (o *AuthorizationRequest) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *AuthorizationRequest) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *AuthorizationRequest) SetImpi(v string)`

SetImpi sets Impi field to given value.

### HasImpi

`func (o *AuthorizationRequest) HasImpi() bool`

HasImpi returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *AuthorizationRequest) GetAuthorizationType() AuthorizationType`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *AuthorizationRequest) GetAuthorizationTypeOk() (*AuthorizationType, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *AuthorizationRequest) SetAuthorizationType(v AuthorizationType)`

SetAuthorizationType sets AuthorizationType field to given value.


### GetVisitedNetworkIdentifier

`func (o *AuthorizationRequest) GetVisitedNetworkIdentifier() string`

GetVisitedNetworkIdentifier returns the VisitedNetworkIdentifier field if non-nil, zero value otherwise.

### GetVisitedNetworkIdentifierOk

`func (o *AuthorizationRequest) GetVisitedNetworkIdentifierOk() (*string, bool)`

GetVisitedNetworkIdentifierOk returns a tuple with the VisitedNetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedNetworkIdentifier

`func (o *AuthorizationRequest) SetVisitedNetworkIdentifier(v string)`

SetVisitedNetworkIdentifier sets VisitedNetworkIdentifier field to given value.

### HasVisitedNetworkIdentifier

`func (o *AuthorizationRequest) HasVisitedNetworkIdentifier() bool`

HasVisitedNetworkIdentifier returns a boolean if a field has been set.

### GetEmergencyIndicator

`func (o *AuthorizationRequest) GetEmergencyIndicator() bool`

GetEmergencyIndicator returns the EmergencyIndicator field if non-nil, zero value otherwise.

### GetEmergencyIndicatorOk

`func (o *AuthorizationRequest) GetEmergencyIndicatorOk() (*bool, bool)`

GetEmergencyIndicatorOk returns a tuple with the EmergencyIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyIndicator

`func (o *AuthorizationRequest) SetEmergencyIndicator(v bool)`

SetEmergencyIndicator sets EmergencyIndicator field to given value.

### HasEmergencyIndicator

`func (o *AuthorizationRequest) HasEmergencyIndicator() bool`

HasEmergencyIndicator returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AuthorizationRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthorizationRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthorizationRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthorizationRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


