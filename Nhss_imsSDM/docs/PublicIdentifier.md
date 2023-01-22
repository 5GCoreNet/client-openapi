# PublicIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIdentity** | [**PublicIdentity**](PublicIdentity.md) |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**ImsServicePriority** | Pointer to [**PriorityLevels**](PriorityLevels.md) |  | [optional] 
**ServiceLevelTraceInfo** | Pointer to [**ServiceLevelTraceInformation**](ServiceLevelTraceInformation.md) |  | [optional] 
**BarringIndicator** | Pointer to **bool** |  | [optional] 
**WildcardedImpu** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicIdentifier

`func NewPublicIdentifier(publicIdentity PublicIdentity, ) *PublicIdentifier`

NewPublicIdentifier instantiates a new PublicIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIdentifierWithDefaults

`func NewPublicIdentifierWithDefaults() *PublicIdentifier`

NewPublicIdentifierWithDefaults instantiates a new PublicIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIdentity

`func (o *PublicIdentifier) GetPublicIdentity() PublicIdentity`

GetPublicIdentity returns the PublicIdentity field if non-nil, zero value otherwise.

### GetPublicIdentityOk

`func (o *PublicIdentifier) GetPublicIdentityOk() (*PublicIdentity, bool)`

GetPublicIdentityOk returns a tuple with the PublicIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdentity

`func (o *PublicIdentifier) SetPublicIdentity(v PublicIdentity)`

SetPublicIdentity sets PublicIdentity field to given value.


### GetDisplayName

`func (o *PublicIdentifier) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PublicIdentifier) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PublicIdentifier) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PublicIdentifier) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetImsServicePriority

`func (o *PublicIdentifier) GetImsServicePriority() PriorityLevels`

GetImsServicePriority returns the ImsServicePriority field if non-nil, zero value otherwise.

### GetImsServicePriorityOk

`func (o *PublicIdentifier) GetImsServicePriorityOk() (*PriorityLevels, bool)`

GetImsServicePriorityOk returns a tuple with the ImsServicePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsServicePriority

`func (o *PublicIdentifier) SetImsServicePriority(v PriorityLevels)`

SetImsServicePriority sets ImsServicePriority field to given value.

### HasImsServicePriority

`func (o *PublicIdentifier) HasImsServicePriority() bool`

HasImsServicePriority returns a boolean if a field has been set.

### GetServiceLevelTraceInfo

`func (o *PublicIdentifier) GetServiceLevelTraceInfo() ServiceLevelTraceInformation`

GetServiceLevelTraceInfo returns the ServiceLevelTraceInfo field if non-nil, zero value otherwise.

### GetServiceLevelTraceInfoOk

`func (o *PublicIdentifier) GetServiceLevelTraceInfoOk() (*ServiceLevelTraceInformation, bool)`

GetServiceLevelTraceInfoOk returns a tuple with the ServiceLevelTraceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelTraceInfo

`func (o *PublicIdentifier) SetServiceLevelTraceInfo(v ServiceLevelTraceInformation)`

SetServiceLevelTraceInfo sets ServiceLevelTraceInfo field to given value.

### HasServiceLevelTraceInfo

`func (o *PublicIdentifier) HasServiceLevelTraceInfo() bool`

HasServiceLevelTraceInfo returns a boolean if a field has been set.

### GetBarringIndicator

`func (o *PublicIdentifier) GetBarringIndicator() bool`

GetBarringIndicator returns the BarringIndicator field if non-nil, zero value otherwise.

### GetBarringIndicatorOk

`func (o *PublicIdentifier) GetBarringIndicatorOk() (*bool, bool)`

GetBarringIndicatorOk returns a tuple with the BarringIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarringIndicator

`func (o *PublicIdentifier) SetBarringIndicator(v bool)`

SetBarringIndicator sets BarringIndicator field to given value.

### HasBarringIndicator

`func (o *PublicIdentifier) HasBarringIndicator() bool`

HasBarringIndicator returns a boolean if a field has been set.

### GetWildcardedImpu

`func (o *PublicIdentifier) GetWildcardedImpu() string`

GetWildcardedImpu returns the WildcardedImpu field if non-nil, zero value otherwise.

### GetWildcardedImpuOk

`func (o *PublicIdentifier) GetWildcardedImpuOk() (*string, bool)`

GetWildcardedImpuOk returns a tuple with the WildcardedImpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardedImpu

`func (o *PublicIdentifier) SetWildcardedImpu(v string)`

SetWildcardedImpu sets WildcardedImpu field to given value.

### HasWildcardedImpu

`func (o *PublicIdentifier) HasWildcardedImpu() bool`

HasWildcardedImpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


