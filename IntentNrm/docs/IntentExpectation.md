# IntentExpectation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectationId** | Pointer to **string** |  | [optional] 
**ExpectationVerb** | Pointer to [**ExpectationVerb**](ExpectationVerb.md) |  | [optional] 
**ExpectationObjects** | Pointer to [**[]ExpectationObject**](ExpectationObject.md) |  | [optional] 
**ExpectationTargets** | Pointer to [**[]ExpectationTarget**](ExpectationTarget.md) |  | [optional] 
**ExpectationContexts** | Pointer to [**[]ExpectationContext**](ExpectationContext.md) |  | [optional] 
**ExpectationfulfilmentInfo** | Pointer to [**FulfilmentInfo**](FulfilmentInfo.md) |  | [optional] 

## Methods

### NewIntentExpectation

`func NewIntentExpectation() *IntentExpectation`

NewIntentExpectation instantiates a new IntentExpectation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntentExpectationWithDefaults

`func NewIntentExpectationWithDefaults() *IntentExpectation`

NewIntentExpectationWithDefaults instantiates a new IntentExpectation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectationId

`func (o *IntentExpectation) GetExpectationId() string`

GetExpectationId returns the ExpectationId field if non-nil, zero value otherwise.

### GetExpectationIdOk

`func (o *IntentExpectation) GetExpectationIdOk() (*string, bool)`

GetExpectationIdOk returns a tuple with the ExpectationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationId

`func (o *IntentExpectation) SetExpectationId(v string)`

SetExpectationId sets ExpectationId field to given value.

### HasExpectationId

`func (o *IntentExpectation) HasExpectationId() bool`

HasExpectationId returns a boolean if a field has been set.

### GetExpectationVerb

`func (o *IntentExpectation) GetExpectationVerb() ExpectationVerb`

GetExpectationVerb returns the ExpectationVerb field if non-nil, zero value otherwise.

### GetExpectationVerbOk

`func (o *IntentExpectation) GetExpectationVerbOk() (*ExpectationVerb, bool)`

GetExpectationVerbOk returns a tuple with the ExpectationVerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationVerb

`func (o *IntentExpectation) SetExpectationVerb(v ExpectationVerb)`

SetExpectationVerb sets ExpectationVerb field to given value.

### HasExpectationVerb

`func (o *IntentExpectation) HasExpectationVerb() bool`

HasExpectationVerb returns a boolean if a field has been set.

### GetExpectationObjects

`func (o *IntentExpectation) GetExpectationObjects() []ExpectationObject`

GetExpectationObjects returns the ExpectationObjects field if non-nil, zero value otherwise.

### GetExpectationObjectsOk

`func (o *IntentExpectation) GetExpectationObjectsOk() (*[]ExpectationObject, bool)`

GetExpectationObjectsOk returns a tuple with the ExpectationObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationObjects

`func (o *IntentExpectation) SetExpectationObjects(v []ExpectationObject)`

SetExpectationObjects sets ExpectationObjects field to given value.

### HasExpectationObjects

`func (o *IntentExpectation) HasExpectationObjects() bool`

HasExpectationObjects returns a boolean if a field has been set.

### GetExpectationTargets

`func (o *IntentExpectation) GetExpectationTargets() []ExpectationTarget`

GetExpectationTargets returns the ExpectationTargets field if non-nil, zero value otherwise.

### GetExpectationTargetsOk

`func (o *IntentExpectation) GetExpectationTargetsOk() (*[]ExpectationTarget, bool)`

GetExpectationTargetsOk returns a tuple with the ExpectationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationTargets

`func (o *IntentExpectation) SetExpectationTargets(v []ExpectationTarget)`

SetExpectationTargets sets ExpectationTargets field to given value.

### HasExpectationTargets

`func (o *IntentExpectation) HasExpectationTargets() bool`

HasExpectationTargets returns a boolean if a field has been set.

### GetExpectationContexts

`func (o *IntentExpectation) GetExpectationContexts() []ExpectationContext`

GetExpectationContexts returns the ExpectationContexts field if non-nil, zero value otherwise.

### GetExpectationContextsOk

`func (o *IntentExpectation) GetExpectationContextsOk() (*[]ExpectationContext, bool)`

GetExpectationContextsOk returns a tuple with the ExpectationContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationContexts

`func (o *IntentExpectation) SetExpectationContexts(v []ExpectationContext)`

SetExpectationContexts sets ExpectationContexts field to given value.

### HasExpectationContexts

`func (o *IntentExpectation) HasExpectationContexts() bool`

HasExpectationContexts returns a boolean if a field has been set.

### GetExpectationfulfilmentInfo

`func (o *IntentExpectation) GetExpectationfulfilmentInfo() FulfilmentInfo`

GetExpectationfulfilmentInfo returns the ExpectationfulfilmentInfo field if non-nil, zero value otherwise.

### GetExpectationfulfilmentInfoOk

`func (o *IntentExpectation) GetExpectationfulfilmentInfoOk() (*FulfilmentInfo, bool)`

GetExpectationfulfilmentInfoOk returns a tuple with the ExpectationfulfilmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationfulfilmentInfo

`func (o *IntentExpectation) SetExpectationfulfilmentInfo(v FulfilmentInfo)`

SetExpectationfulfilmentInfo sets ExpectationfulfilmentInfo field to given value.

### HasExpectationfulfilmentInfo

`func (o *IntentExpectation) HasExpectationfulfilmentInfo() bool`

HasExpectationfulfilmentInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


