# TargetUeId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnyUeInd** | Pointer to **bool** |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExterGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 

## Methods

### NewTargetUeId

`func NewTargetUeId() *TargetUeId`

NewTargetUeId instantiates a new TargetUeId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUeIdWithDefaults

`func NewTargetUeIdWithDefaults() *TargetUeId`

NewTargetUeIdWithDefaults instantiates a new TargetUeId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnyUeInd

`func (o *TargetUeId) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *TargetUeId) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *TargetUeId) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *TargetUeId) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetGpsi

`func (o *TargetUeId) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *TargetUeId) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *TargetUeId) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *TargetUeId) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetExterGroupId

`func (o *TargetUeId) GetExterGroupId() string`

GetExterGroupId returns the ExterGroupId field if non-nil, zero value otherwise.

### GetExterGroupIdOk

`func (o *TargetUeId) GetExterGroupIdOk() (*string, bool)`

GetExterGroupIdOk returns a tuple with the ExterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupId

`func (o *TargetUeId) SetExterGroupId(v string)`

SetExterGroupId sets ExterGroupId field to given value.

### HasExterGroupId

`func (o *TargetUeId) HasExterGroupId() bool`

HasExterGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


