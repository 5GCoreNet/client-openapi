# AnnouncementInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnouncementIdentifier** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**AnnouncementReference** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**VariableParts** | Pointer to [**[]VariablePart**](VariablePart.md) |  | [optional] 
**TimeToPlay** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**QuotaConsumptionIndicator** | Pointer to [**QuotaConsumptionIndicator**](QuotaConsumptionIndicator.md) |  | [optional] 
**AnnouncementPriority** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**PlayToParty** | Pointer to [**PlayToParty**](PlayToParty.md) |  | [optional] 
**AnnouncementPrivacyIndicator** | Pointer to [**AnnouncementPrivacyIndicator**](AnnouncementPrivacyIndicator.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewAnnouncementInformation

`func NewAnnouncementInformation() *AnnouncementInformation`

NewAnnouncementInformation instantiates a new AnnouncementInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementInformationWithDefaults

`func NewAnnouncementInformationWithDefaults() *AnnouncementInformation`

NewAnnouncementInformationWithDefaults instantiates a new AnnouncementInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncementIdentifier

`func (o *AnnouncementInformation) GetAnnouncementIdentifier() int32`

GetAnnouncementIdentifier returns the AnnouncementIdentifier field if non-nil, zero value otherwise.

### GetAnnouncementIdentifierOk

`func (o *AnnouncementInformation) GetAnnouncementIdentifierOk() (*int32, bool)`

GetAnnouncementIdentifierOk returns a tuple with the AnnouncementIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementIdentifier

`func (o *AnnouncementInformation) SetAnnouncementIdentifier(v int32)`

SetAnnouncementIdentifier sets AnnouncementIdentifier field to given value.

### HasAnnouncementIdentifier

`func (o *AnnouncementInformation) HasAnnouncementIdentifier() bool`

HasAnnouncementIdentifier returns a boolean if a field has been set.

### GetAnnouncementReference

`func (o *AnnouncementInformation) GetAnnouncementReference() string`

GetAnnouncementReference returns the AnnouncementReference field if non-nil, zero value otherwise.

### GetAnnouncementReferenceOk

`func (o *AnnouncementInformation) GetAnnouncementReferenceOk() (*string, bool)`

GetAnnouncementReferenceOk returns a tuple with the AnnouncementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementReference

`func (o *AnnouncementInformation) SetAnnouncementReference(v string)`

SetAnnouncementReference sets AnnouncementReference field to given value.

### HasAnnouncementReference

`func (o *AnnouncementInformation) HasAnnouncementReference() bool`

HasAnnouncementReference returns a boolean if a field has been set.

### GetVariableParts

`func (o *AnnouncementInformation) GetVariableParts() []VariablePart`

GetVariableParts returns the VariableParts field if non-nil, zero value otherwise.

### GetVariablePartsOk

`func (o *AnnouncementInformation) GetVariablePartsOk() (*[]VariablePart, bool)`

GetVariablePartsOk returns a tuple with the VariableParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableParts

`func (o *AnnouncementInformation) SetVariableParts(v []VariablePart)`

SetVariableParts sets VariableParts field to given value.

### HasVariableParts

`func (o *AnnouncementInformation) HasVariableParts() bool`

HasVariableParts returns a boolean if a field has been set.

### GetTimeToPlay

`func (o *AnnouncementInformation) GetTimeToPlay() int32`

GetTimeToPlay returns the TimeToPlay field if non-nil, zero value otherwise.

### GetTimeToPlayOk

`func (o *AnnouncementInformation) GetTimeToPlayOk() (*int32, bool)`

GetTimeToPlayOk returns a tuple with the TimeToPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToPlay

`func (o *AnnouncementInformation) SetTimeToPlay(v int32)`

SetTimeToPlay sets TimeToPlay field to given value.

### HasTimeToPlay

`func (o *AnnouncementInformation) HasTimeToPlay() bool`

HasTimeToPlay returns a boolean if a field has been set.

### GetQuotaConsumptionIndicator

`func (o *AnnouncementInformation) GetQuotaConsumptionIndicator() QuotaConsumptionIndicator`

GetQuotaConsumptionIndicator returns the QuotaConsumptionIndicator field if non-nil, zero value otherwise.

### GetQuotaConsumptionIndicatorOk

`func (o *AnnouncementInformation) GetQuotaConsumptionIndicatorOk() (*QuotaConsumptionIndicator, bool)`

GetQuotaConsumptionIndicatorOk returns a tuple with the QuotaConsumptionIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaConsumptionIndicator

`func (o *AnnouncementInformation) SetQuotaConsumptionIndicator(v QuotaConsumptionIndicator)`

SetQuotaConsumptionIndicator sets QuotaConsumptionIndicator field to given value.

### HasQuotaConsumptionIndicator

`func (o *AnnouncementInformation) HasQuotaConsumptionIndicator() bool`

HasQuotaConsumptionIndicator returns a boolean if a field has been set.

### GetAnnouncementPriority

`func (o *AnnouncementInformation) GetAnnouncementPriority() int32`

GetAnnouncementPriority returns the AnnouncementPriority field if non-nil, zero value otherwise.

### GetAnnouncementPriorityOk

`func (o *AnnouncementInformation) GetAnnouncementPriorityOk() (*int32, bool)`

GetAnnouncementPriorityOk returns a tuple with the AnnouncementPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementPriority

`func (o *AnnouncementInformation) SetAnnouncementPriority(v int32)`

SetAnnouncementPriority sets AnnouncementPriority field to given value.

### HasAnnouncementPriority

`func (o *AnnouncementInformation) HasAnnouncementPriority() bool`

HasAnnouncementPriority returns a boolean if a field has been set.

### GetPlayToParty

`func (o *AnnouncementInformation) GetPlayToParty() PlayToParty`

GetPlayToParty returns the PlayToParty field if non-nil, zero value otherwise.

### GetPlayToPartyOk

`func (o *AnnouncementInformation) GetPlayToPartyOk() (*PlayToParty, bool)`

GetPlayToPartyOk returns a tuple with the PlayToParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayToParty

`func (o *AnnouncementInformation) SetPlayToParty(v PlayToParty)`

SetPlayToParty sets PlayToParty field to given value.

### HasPlayToParty

`func (o *AnnouncementInformation) HasPlayToParty() bool`

HasPlayToParty returns a boolean if a field has been set.

### GetAnnouncementPrivacyIndicator

`func (o *AnnouncementInformation) GetAnnouncementPrivacyIndicator() AnnouncementPrivacyIndicator`

GetAnnouncementPrivacyIndicator returns the AnnouncementPrivacyIndicator field if non-nil, zero value otherwise.

### GetAnnouncementPrivacyIndicatorOk

`func (o *AnnouncementInformation) GetAnnouncementPrivacyIndicatorOk() (*AnnouncementPrivacyIndicator, bool)`

GetAnnouncementPrivacyIndicatorOk returns a tuple with the AnnouncementPrivacyIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementPrivacyIndicator

`func (o *AnnouncementInformation) SetAnnouncementPrivacyIndicator(v AnnouncementPrivacyIndicator)`

SetAnnouncementPrivacyIndicator sets AnnouncementPrivacyIndicator field to given value.

### HasAnnouncementPrivacyIndicator

`func (o *AnnouncementInformation) HasAnnouncementPrivacyIndicator() bool`

HasAnnouncementPrivacyIndicator returns a boolean if a field has been set.

### GetLanguage

`func (o *AnnouncementInformation) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AnnouncementInformation) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AnnouncementInformation) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AnnouncementInformation) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


