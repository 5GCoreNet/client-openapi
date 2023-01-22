# TerminationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Cause** | [**SmPolicyAssociationReleaseCause**](SmPolicyAssociationReleaseCause.md) |  | 

## Methods

### NewTerminationNotification

`func NewTerminationNotification(resourceUri string, cause SmPolicyAssociationReleaseCause, ) *TerminationNotification`

NewTerminationNotification instantiates a new TerminationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminationNotificationWithDefaults

`func NewTerminationNotificationWithDefaults() *TerminationNotification`

NewTerminationNotificationWithDefaults instantiates a new TerminationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *TerminationNotification) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *TerminationNotification) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *TerminationNotification) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetCause

`func (o *TerminationNotification) GetCause() SmPolicyAssociationReleaseCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *TerminationNotification) GetCauseOk() (*SmPolicyAssociationReleaseCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *TerminationNotification) SetCause(v SmPolicyAssociationReleaseCause)`

SetCause sets Cause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


