# ScscfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impi** | Pointer to **string** | IMS Private Identity of the UE | [optional] 
**ImsRegistrationType** | [**ImsRegistrationType**](ImsRegistrationType.md) |  | 
**CscfServerName** | **string** |  | 
**ScscfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**DeregCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AssociatedImpis** | Pointer to **[]string** |  | [optional] 
**AssociatedRegisteredImpis** | Pointer to **[]string** |  | [optional] 
**IrsImpus** | Pointer to **[]string** |  | [optional] 
**WildcardedPui** | Pointer to **string** | IMS Public Identity of the UE (sip URI or tel URI) | [optional] 
**LooseRouteIndicator** | Pointer to [**LooseRouteIndication**](LooseRouteIndication.md) |  | [optional] 
**WildcardedPsi** | Pointer to **string** | IMS Public Identity of the UE (sip URI or tel URI) | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MultipleRegistrationIndicator** | Pointer to **bool** |  | [optional] 
**PcscfRestorationIndicator** | Pointer to **bool** |  | [optional] [default to false]
**ScscfReselectionIndicator** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewScscfRegistration

`func NewScscfRegistration(imsRegistrationType ImsRegistrationType, cscfServerName string, ) *ScscfRegistration`

NewScscfRegistration instantiates a new ScscfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScscfRegistrationWithDefaults

`func NewScscfRegistrationWithDefaults() *ScscfRegistration`

NewScscfRegistrationWithDefaults instantiates a new ScscfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpi

`func (o *ScscfRegistration) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *ScscfRegistration) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *ScscfRegistration) SetImpi(v string)`

SetImpi sets Impi field to given value.

### HasImpi

`func (o *ScscfRegistration) HasImpi() bool`

HasImpi returns a boolean if a field has been set.

### GetImsRegistrationType

`func (o *ScscfRegistration) GetImsRegistrationType() ImsRegistrationType`

GetImsRegistrationType returns the ImsRegistrationType field if non-nil, zero value otherwise.

### GetImsRegistrationTypeOk

`func (o *ScscfRegistration) GetImsRegistrationTypeOk() (*ImsRegistrationType, bool)`

GetImsRegistrationTypeOk returns a tuple with the ImsRegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsRegistrationType

`func (o *ScscfRegistration) SetImsRegistrationType(v ImsRegistrationType)`

SetImsRegistrationType sets ImsRegistrationType field to given value.


### GetCscfServerName

`func (o *ScscfRegistration) GetCscfServerName() string`

GetCscfServerName returns the CscfServerName field if non-nil, zero value otherwise.

### GetCscfServerNameOk

`func (o *ScscfRegistration) GetCscfServerNameOk() (*string, bool)`

GetCscfServerNameOk returns a tuple with the CscfServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscfServerName

`func (o *ScscfRegistration) SetCscfServerName(v string)`

SetCscfServerName sets CscfServerName field to given value.


### GetScscfInstanceId

`func (o *ScscfRegistration) GetScscfInstanceId() string`

GetScscfInstanceId returns the ScscfInstanceId field if non-nil, zero value otherwise.

### GetScscfInstanceIdOk

`func (o *ScscfRegistration) GetScscfInstanceIdOk() (*string, bool)`

GetScscfInstanceIdOk returns a tuple with the ScscfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScscfInstanceId

`func (o *ScscfRegistration) SetScscfInstanceId(v string)`

SetScscfInstanceId sets ScscfInstanceId field to given value.

### HasScscfInstanceId

`func (o *ScscfRegistration) HasScscfInstanceId() bool`

HasScscfInstanceId returns a boolean if a field has been set.

### GetDeregCallbackUri

`func (o *ScscfRegistration) GetDeregCallbackUri() string`

GetDeregCallbackUri returns the DeregCallbackUri field if non-nil, zero value otherwise.

### GetDeregCallbackUriOk

`func (o *ScscfRegistration) GetDeregCallbackUriOk() (*string, bool)`

GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregCallbackUri

`func (o *ScscfRegistration) SetDeregCallbackUri(v string)`

SetDeregCallbackUri sets DeregCallbackUri field to given value.

### HasDeregCallbackUri

`func (o *ScscfRegistration) HasDeregCallbackUri() bool`

HasDeregCallbackUri returns a boolean if a field has been set.

### GetAssociatedImpis

`func (o *ScscfRegistration) GetAssociatedImpis() []string`

GetAssociatedImpis returns the AssociatedImpis field if non-nil, zero value otherwise.

### GetAssociatedImpisOk

`func (o *ScscfRegistration) GetAssociatedImpisOk() (*[]string, bool)`

GetAssociatedImpisOk returns a tuple with the AssociatedImpis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedImpis

`func (o *ScscfRegistration) SetAssociatedImpis(v []string)`

SetAssociatedImpis sets AssociatedImpis field to given value.

### HasAssociatedImpis

`func (o *ScscfRegistration) HasAssociatedImpis() bool`

HasAssociatedImpis returns a boolean if a field has been set.

### GetAssociatedRegisteredImpis

`func (o *ScscfRegistration) GetAssociatedRegisteredImpis() []string`

GetAssociatedRegisteredImpis returns the AssociatedRegisteredImpis field if non-nil, zero value otherwise.

### GetAssociatedRegisteredImpisOk

`func (o *ScscfRegistration) GetAssociatedRegisteredImpisOk() (*[]string, bool)`

GetAssociatedRegisteredImpisOk returns a tuple with the AssociatedRegisteredImpis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRegisteredImpis

`func (o *ScscfRegistration) SetAssociatedRegisteredImpis(v []string)`

SetAssociatedRegisteredImpis sets AssociatedRegisteredImpis field to given value.

### HasAssociatedRegisteredImpis

`func (o *ScscfRegistration) HasAssociatedRegisteredImpis() bool`

HasAssociatedRegisteredImpis returns a boolean if a field has been set.

### GetIrsImpus

`func (o *ScscfRegistration) GetIrsImpus() []string`

GetIrsImpus returns the IrsImpus field if non-nil, zero value otherwise.

### GetIrsImpusOk

`func (o *ScscfRegistration) GetIrsImpusOk() (*[]string, bool)`

GetIrsImpusOk returns a tuple with the IrsImpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrsImpus

`func (o *ScscfRegistration) SetIrsImpus(v []string)`

SetIrsImpus sets IrsImpus field to given value.

### HasIrsImpus

`func (o *ScscfRegistration) HasIrsImpus() bool`

HasIrsImpus returns a boolean if a field has been set.

### GetWildcardedPui

`func (o *ScscfRegistration) GetWildcardedPui() string`

GetWildcardedPui returns the WildcardedPui field if non-nil, zero value otherwise.

### GetWildcardedPuiOk

`func (o *ScscfRegistration) GetWildcardedPuiOk() (*string, bool)`

GetWildcardedPuiOk returns a tuple with the WildcardedPui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardedPui

`func (o *ScscfRegistration) SetWildcardedPui(v string)`

SetWildcardedPui sets WildcardedPui field to given value.

### HasWildcardedPui

`func (o *ScscfRegistration) HasWildcardedPui() bool`

HasWildcardedPui returns a boolean if a field has been set.

### GetLooseRouteIndicator

`func (o *ScscfRegistration) GetLooseRouteIndicator() LooseRouteIndication`

GetLooseRouteIndicator returns the LooseRouteIndicator field if non-nil, zero value otherwise.

### GetLooseRouteIndicatorOk

`func (o *ScscfRegistration) GetLooseRouteIndicatorOk() (*LooseRouteIndication, bool)`

GetLooseRouteIndicatorOk returns a tuple with the LooseRouteIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLooseRouteIndicator

`func (o *ScscfRegistration) SetLooseRouteIndicator(v LooseRouteIndication)`

SetLooseRouteIndicator sets LooseRouteIndicator field to given value.

### HasLooseRouteIndicator

`func (o *ScscfRegistration) HasLooseRouteIndicator() bool`

HasLooseRouteIndicator returns a boolean if a field has been set.

### GetWildcardedPsi

`func (o *ScscfRegistration) GetWildcardedPsi() string`

GetWildcardedPsi returns the WildcardedPsi field if non-nil, zero value otherwise.

### GetWildcardedPsiOk

`func (o *ScscfRegistration) GetWildcardedPsiOk() (*string, bool)`

GetWildcardedPsiOk returns a tuple with the WildcardedPsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardedPsi

`func (o *ScscfRegistration) SetWildcardedPsi(v string)`

SetWildcardedPsi sets WildcardedPsi field to given value.

### HasWildcardedPsi

`func (o *ScscfRegistration) HasWildcardedPsi() bool`

HasWildcardedPsi returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ScscfRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ScscfRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ScscfRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ScscfRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMultipleRegistrationIndicator

`func (o *ScscfRegistration) GetMultipleRegistrationIndicator() bool`

GetMultipleRegistrationIndicator returns the MultipleRegistrationIndicator field if non-nil, zero value otherwise.

### GetMultipleRegistrationIndicatorOk

`func (o *ScscfRegistration) GetMultipleRegistrationIndicatorOk() (*bool, bool)`

GetMultipleRegistrationIndicatorOk returns a tuple with the MultipleRegistrationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleRegistrationIndicator

`func (o *ScscfRegistration) SetMultipleRegistrationIndicator(v bool)`

SetMultipleRegistrationIndicator sets MultipleRegistrationIndicator field to given value.

### HasMultipleRegistrationIndicator

`func (o *ScscfRegistration) HasMultipleRegistrationIndicator() bool`

HasMultipleRegistrationIndicator returns a boolean if a field has been set.

### GetPcscfRestorationIndicator

`func (o *ScscfRegistration) GetPcscfRestorationIndicator() bool`

GetPcscfRestorationIndicator returns the PcscfRestorationIndicator field if non-nil, zero value otherwise.

### GetPcscfRestorationIndicatorOk

`func (o *ScscfRegistration) GetPcscfRestorationIndicatorOk() (*bool, bool)`

GetPcscfRestorationIndicatorOk returns a tuple with the PcscfRestorationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfRestorationIndicator

`func (o *ScscfRegistration) SetPcscfRestorationIndicator(v bool)`

SetPcscfRestorationIndicator sets PcscfRestorationIndicator field to given value.

### HasPcscfRestorationIndicator

`func (o *ScscfRegistration) HasPcscfRestorationIndicator() bool`

HasPcscfRestorationIndicator returns a boolean if a field has been set.

### GetScscfReselectionIndicator

`func (o *ScscfRegistration) GetScscfReselectionIndicator() bool`

GetScscfReselectionIndicator returns the ScscfReselectionIndicator field if non-nil, zero value otherwise.

### GetScscfReselectionIndicatorOk

`func (o *ScscfRegistration) GetScscfReselectionIndicatorOk() (*bool, bool)`

GetScscfReselectionIndicatorOk returns a tuple with the ScscfReselectionIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScscfReselectionIndicator

`func (o *ScscfRegistration) SetScscfReselectionIndicator(v bool)`

SetScscfReselectionIndicator sets ScscfReselectionIndicator field to given value.

### HasScscfReselectionIndicator

`func (o *ScscfRegistration) HasScscfReselectionIndicator() bool`

HasScscfReselectionIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


