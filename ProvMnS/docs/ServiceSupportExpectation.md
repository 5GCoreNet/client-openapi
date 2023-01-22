# ServiceSupportExpectation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectationId** | Pointer to **string** |  | [optional] 
**ExpectationVerb** | Pointer to [**ExpectationVerb**](ExpectationVerb.md) |  | [optional] 
**ExpectationObjects** | Pointer to [**[]ServiceSupportExpectationObject**](ServiceSupportExpectationObject.md) |  | [optional] 
**ExpectationTargets** | Pointer to [**[]ServiceSupportExpectationExpectationTargetsInner**](ServiceSupportExpectationExpectationTargetsInner.md) |  | [optional] 
**ExpectationContexts** | Pointer to [**[]ServiceSupportExpectationExpectationContextsInner**](ServiceSupportExpectationExpectationContextsInner.md) |  | [optional] 
**ExpectationfulfilmentInfo** | Pointer to [**FulfilmentInfo**](FulfilmentInfo.md) |  | [optional] 

## Methods

### NewServiceSupportExpectation

`func NewServiceSupportExpectation() *ServiceSupportExpectation`

NewServiceSupportExpectation instantiates a new ServiceSupportExpectation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSupportExpectationWithDefaults

`func NewServiceSupportExpectationWithDefaults() *ServiceSupportExpectation`

NewServiceSupportExpectationWithDefaults instantiates a new ServiceSupportExpectation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectationId

`func (o *ServiceSupportExpectation) GetExpectationId() string`

GetExpectationId returns the ExpectationId field if non-nil, zero value otherwise.

### GetExpectationIdOk

`func (o *ServiceSupportExpectation) GetExpectationIdOk() (*string, bool)`

GetExpectationIdOk returns a tuple with the ExpectationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationId

`func (o *ServiceSupportExpectation) SetExpectationId(v string)`

SetExpectationId sets ExpectationId field to given value.

### HasExpectationId

`func (o *ServiceSupportExpectation) HasExpectationId() bool`

HasExpectationId returns a boolean if a field has been set.

### GetExpectationVerb

`func (o *ServiceSupportExpectation) GetExpectationVerb() ExpectationVerb`

GetExpectationVerb returns the ExpectationVerb field if non-nil, zero value otherwise.

### GetExpectationVerbOk

`func (o *ServiceSupportExpectation) GetExpectationVerbOk() (*ExpectationVerb, bool)`

GetExpectationVerbOk returns a tuple with the ExpectationVerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationVerb

`func (o *ServiceSupportExpectation) SetExpectationVerb(v ExpectationVerb)`

SetExpectationVerb sets ExpectationVerb field to given value.

### HasExpectationVerb

`func (o *ServiceSupportExpectation) HasExpectationVerb() bool`

HasExpectationVerb returns a boolean if a field has been set.

### GetExpectationObjects

`func (o *ServiceSupportExpectation) GetExpectationObjects() []ServiceSupportExpectationObject`

GetExpectationObjects returns the ExpectationObjects field if non-nil, zero value otherwise.

### GetExpectationObjectsOk

`func (o *ServiceSupportExpectation) GetExpectationObjectsOk() (*[]ServiceSupportExpectationObject, bool)`

GetExpectationObjectsOk returns a tuple with the ExpectationObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationObjects

`func (o *ServiceSupportExpectation) SetExpectationObjects(v []ServiceSupportExpectationObject)`

SetExpectationObjects sets ExpectationObjects field to given value.

### HasExpectationObjects

`func (o *ServiceSupportExpectation) HasExpectationObjects() bool`

HasExpectationObjects returns a boolean if a field has been set.

### GetExpectationTargets

`func (o *ServiceSupportExpectation) GetExpectationTargets() []ServiceSupportExpectationExpectationTargetsInner`

GetExpectationTargets returns the ExpectationTargets field if non-nil, zero value otherwise.

### GetExpectationTargetsOk

`func (o *ServiceSupportExpectation) GetExpectationTargetsOk() (*[]ServiceSupportExpectationExpectationTargetsInner, bool)`

GetExpectationTargetsOk returns a tuple with the ExpectationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationTargets

`func (o *ServiceSupportExpectation) SetExpectationTargets(v []ServiceSupportExpectationExpectationTargetsInner)`

SetExpectationTargets sets ExpectationTargets field to given value.

### HasExpectationTargets

`func (o *ServiceSupportExpectation) HasExpectationTargets() bool`

HasExpectationTargets returns a boolean if a field has been set.

### GetExpectationContexts

`func (o *ServiceSupportExpectation) GetExpectationContexts() []ServiceSupportExpectationExpectationContextsInner`

GetExpectationContexts returns the ExpectationContexts field if non-nil, zero value otherwise.

### GetExpectationContextsOk

`func (o *ServiceSupportExpectation) GetExpectationContextsOk() (*[]ServiceSupportExpectationExpectationContextsInner, bool)`

GetExpectationContextsOk returns a tuple with the ExpectationContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationContexts

`func (o *ServiceSupportExpectation) SetExpectationContexts(v []ServiceSupportExpectationExpectationContextsInner)`

SetExpectationContexts sets ExpectationContexts field to given value.

### HasExpectationContexts

`func (o *ServiceSupportExpectation) HasExpectationContexts() bool`

HasExpectationContexts returns a boolean if a field has been set.

### GetExpectationfulfilmentInfo

`func (o *ServiceSupportExpectation) GetExpectationfulfilmentInfo() FulfilmentInfo`

GetExpectationfulfilmentInfo returns the ExpectationfulfilmentInfo field if non-nil, zero value otherwise.

### GetExpectationfulfilmentInfoOk

`func (o *ServiceSupportExpectation) GetExpectationfulfilmentInfoOk() (*FulfilmentInfo, bool)`

GetExpectationfulfilmentInfoOk returns a tuple with the ExpectationfulfilmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectationfulfilmentInfo

`func (o *ServiceSupportExpectation) SetExpectationfulfilmentInfo(v FulfilmentInfo)`

SetExpectationfulfilmentInfo sets ExpectationfulfilmentInfo field to given value.

### HasExpectationfulfilmentInfo

`func (o *ServiceSupportExpectation) HasExpectationfulfilmentInfo() bool`

HasExpectationfulfilmentInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


