# NEFChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIndividualIdentifier** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExternalGroupIdentifier** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**GroupIdentifier** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**APIDirection** | Pointer to [**APIDirection**](APIDirection.md) |  | [optional] 
**APITargetNetworkFunction** | Pointer to [**NFIdentification**](NFIdentification.md) |  | [optional] 
**APIResultCode** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**APIName** | **string** |  | 
**APIReference** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**APIContent** | Pointer to **string** |  | [optional] 

## Methods

### NewNEFChargingInformation

`func NewNEFChargingInformation(aPIName string, ) *NEFChargingInformation`

NewNEFChargingInformation instantiates a new NEFChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNEFChargingInformationWithDefaults

`func NewNEFChargingInformationWithDefaults() *NEFChargingInformation`

NewNEFChargingInformationWithDefaults instantiates a new NEFChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIndividualIdentifier

`func (o *NEFChargingInformation) GetExternalIndividualIdentifier() string`

GetExternalIndividualIdentifier returns the ExternalIndividualIdentifier field if non-nil, zero value otherwise.

### GetExternalIndividualIdentifierOk

`func (o *NEFChargingInformation) GetExternalIndividualIdentifierOk() (*string, bool)`

GetExternalIndividualIdentifierOk returns a tuple with the ExternalIndividualIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIndividualIdentifier

`func (o *NEFChargingInformation) SetExternalIndividualIdentifier(v string)`

SetExternalIndividualIdentifier sets ExternalIndividualIdentifier field to given value.

### HasExternalIndividualIdentifier

`func (o *NEFChargingInformation) HasExternalIndividualIdentifier() bool`

HasExternalIndividualIdentifier returns a boolean if a field has been set.

### GetExternalGroupIdentifier

`func (o *NEFChargingInformation) GetExternalGroupIdentifier() string`

GetExternalGroupIdentifier returns the ExternalGroupIdentifier field if non-nil, zero value otherwise.

### GetExternalGroupIdentifierOk

`func (o *NEFChargingInformation) GetExternalGroupIdentifierOk() (*string, bool)`

GetExternalGroupIdentifierOk returns a tuple with the ExternalGroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifier

`func (o *NEFChargingInformation) SetExternalGroupIdentifier(v string)`

SetExternalGroupIdentifier sets ExternalGroupIdentifier field to given value.

### HasExternalGroupIdentifier

`func (o *NEFChargingInformation) HasExternalGroupIdentifier() bool`

HasExternalGroupIdentifier returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *NEFChargingInformation) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *NEFChargingInformation) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *NEFChargingInformation) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *NEFChargingInformation) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.

### GetAPIDirection

`func (o *NEFChargingInformation) GetAPIDirection() APIDirection`

GetAPIDirection returns the APIDirection field if non-nil, zero value otherwise.

### GetAPIDirectionOk

`func (o *NEFChargingInformation) GetAPIDirectionOk() (*APIDirection, bool)`

GetAPIDirectionOk returns a tuple with the APIDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIDirection

`func (o *NEFChargingInformation) SetAPIDirection(v APIDirection)`

SetAPIDirection sets APIDirection field to given value.

### HasAPIDirection

`func (o *NEFChargingInformation) HasAPIDirection() bool`

HasAPIDirection returns a boolean if a field has been set.

### GetAPITargetNetworkFunction

`func (o *NEFChargingInformation) GetAPITargetNetworkFunction() NFIdentification`

GetAPITargetNetworkFunction returns the APITargetNetworkFunction field if non-nil, zero value otherwise.

### GetAPITargetNetworkFunctionOk

`func (o *NEFChargingInformation) GetAPITargetNetworkFunctionOk() (*NFIdentification, bool)`

GetAPITargetNetworkFunctionOk returns a tuple with the APITargetNetworkFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPITargetNetworkFunction

`func (o *NEFChargingInformation) SetAPITargetNetworkFunction(v NFIdentification)`

SetAPITargetNetworkFunction sets APITargetNetworkFunction field to given value.

### HasAPITargetNetworkFunction

`func (o *NEFChargingInformation) HasAPITargetNetworkFunction() bool`

HasAPITargetNetworkFunction returns a boolean if a field has been set.

### GetAPIResultCode

`func (o *NEFChargingInformation) GetAPIResultCode() int32`

GetAPIResultCode returns the APIResultCode field if non-nil, zero value otherwise.

### GetAPIResultCodeOk

`func (o *NEFChargingInformation) GetAPIResultCodeOk() (*int32, bool)`

GetAPIResultCodeOk returns a tuple with the APIResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIResultCode

`func (o *NEFChargingInformation) SetAPIResultCode(v int32)`

SetAPIResultCode sets APIResultCode field to given value.

### HasAPIResultCode

`func (o *NEFChargingInformation) HasAPIResultCode() bool`

HasAPIResultCode returns a boolean if a field has been set.

### GetAPIName

`func (o *NEFChargingInformation) GetAPIName() string`

GetAPIName returns the APIName field if non-nil, zero value otherwise.

### GetAPINameOk

`func (o *NEFChargingInformation) GetAPINameOk() (*string, bool)`

GetAPINameOk returns a tuple with the APIName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIName

`func (o *NEFChargingInformation) SetAPIName(v string)`

SetAPIName sets APIName field to given value.


### GetAPIReference

`func (o *NEFChargingInformation) GetAPIReference() string`

GetAPIReference returns the APIReference field if non-nil, zero value otherwise.

### GetAPIReferenceOk

`func (o *NEFChargingInformation) GetAPIReferenceOk() (*string, bool)`

GetAPIReferenceOk returns a tuple with the APIReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIReference

`func (o *NEFChargingInformation) SetAPIReference(v string)`

SetAPIReference sets APIReference field to given value.

### HasAPIReference

`func (o *NEFChargingInformation) HasAPIReference() bool`

HasAPIReference returns a boolean if a field has been set.

### GetAPIContent

`func (o *NEFChargingInformation) GetAPIContent() string`

GetAPIContent returns the APIContent field if non-nil, zero value otherwise.

### GetAPIContentOk

`func (o *NEFChargingInformation) GetAPIContentOk() (*string, bool)`

GetAPIContentOk returns a tuple with the APIContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIContent

`func (o *NEFChargingInformation) SetAPIContent(v string)`

SetAPIContent sets APIContent field to given value.

### HasAPIContent

`func (o *NEFChargingInformation) HasAPIContent() bool`

HasAPIContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


