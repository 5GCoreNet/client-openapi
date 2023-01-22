# PduSessionStatusCfg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 

## Methods

### NewPduSessionStatusCfg

`func NewPduSessionStatusCfg() *PduSessionStatusCfg`

NewPduSessionStatusCfg instantiates a new PduSessionStatusCfg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionStatusCfgWithDefaults

`func NewPduSessionStatusCfgWithDefaults() *PduSessionStatusCfg`

NewPduSessionStatusCfgWithDefaults instantiates a new PduSessionStatusCfg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *PduSessionStatusCfg) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *PduSessionStatusCfg) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *PduSessionStatusCfg) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *PduSessionStatusCfg) HasApn() bool`

HasApn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


