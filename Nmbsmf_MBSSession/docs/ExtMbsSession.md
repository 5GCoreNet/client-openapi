# ExtMbsSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyAuthInd** | Pointer to **bool** |  | [optional] [default to false]
**MbsSecurityContext** | Pointer to [**MbsSecurityContext**](MbsSecurityContext.md) |  | [optional] 
**ContactPcfInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewExtMbsSession

`func NewExtMbsSession() *ExtMbsSession`

NewExtMbsSession instantiates a new ExtMbsSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtMbsSessionWithDefaults

`func NewExtMbsSessionWithDefaults() *ExtMbsSession`

NewExtMbsSessionWithDefaults instantiates a new ExtMbsSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyAuthInd

`func (o *ExtMbsSession) GetPolicyAuthInd() bool`

GetPolicyAuthInd returns the PolicyAuthInd field if non-nil, zero value otherwise.

### GetPolicyAuthIndOk

`func (o *ExtMbsSession) GetPolicyAuthIndOk() (*bool, bool)`

GetPolicyAuthIndOk returns a tuple with the PolicyAuthInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAuthInd

`func (o *ExtMbsSession) SetPolicyAuthInd(v bool)`

SetPolicyAuthInd sets PolicyAuthInd field to given value.

### HasPolicyAuthInd

`func (o *ExtMbsSession) HasPolicyAuthInd() bool`

HasPolicyAuthInd returns a boolean if a field has been set.

### GetMbsSecurityContext

`func (o *ExtMbsSession) GetMbsSecurityContext() MbsSecurityContext`

GetMbsSecurityContext returns the MbsSecurityContext field if non-nil, zero value otherwise.

### GetMbsSecurityContextOk

`func (o *ExtMbsSession) GetMbsSecurityContextOk() (*MbsSecurityContext, bool)`

GetMbsSecurityContextOk returns a tuple with the MbsSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSecurityContext

`func (o *ExtMbsSession) SetMbsSecurityContext(v MbsSecurityContext)`

SetMbsSecurityContext sets MbsSecurityContext field to given value.

### HasMbsSecurityContext

`func (o *ExtMbsSession) HasMbsSecurityContext() bool`

HasMbsSecurityContext returns a boolean if a field has been set.

### GetContactPcfInd

`func (o *ExtMbsSession) GetContactPcfInd() bool`

GetContactPcfInd returns the ContactPcfInd field if non-nil, zero value otherwise.

### GetContactPcfIndOk

`func (o *ExtMbsSession) GetContactPcfIndOk() (*bool, bool)`

GetContactPcfIndOk returns a tuple with the ContactPcfInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPcfInd

`func (o *ExtMbsSession) SetContactPcfInd(v bool)`

SetContactPcfInd sets ContactPcfInd field to given value.

### HasContactPcfInd

`func (o *ExtMbsSession) HasContactPcfInd() bool`

HasContactPcfInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


