# EECContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EecId** | **string** | Unique idenitfier of the EEC. | 
**CntxId** | **string** | Unique idenitfier assigned to the EEC context. | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**E1Subs** | Pointer to **[]string** | List of subscription IDs for the capability expsoure for the EEC ID. | [optional] 
**UeLoc** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**AcProfs** | Pointer to [**[]ACProfile**](ACProfile.md) | List AC profiles. | [optional] 
**SessCntxs** | Pointer to [**SessionContexts**](SessionContexts.md) |  | [optional] 

## Methods

### NewEECContext

`func NewEECContext(eecId string, cntxId string, ) *EECContext`

NewEECContext instantiates a new EECContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEECContextWithDefaults

`func NewEECContextWithDefaults() *EECContext`

NewEECContextWithDefaults instantiates a new EECContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEecId

`func (o *EECContext) GetEecId() string`

GetEecId returns the EecId field if non-nil, zero value otherwise.

### GetEecIdOk

`func (o *EECContext) GetEecIdOk() (*string, bool)`

GetEecIdOk returns a tuple with the EecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecId

`func (o *EECContext) SetEecId(v string)`

SetEecId sets EecId field to given value.


### GetCntxId

`func (o *EECContext) GetCntxId() string`

GetCntxId returns the CntxId field if non-nil, zero value otherwise.

### GetCntxIdOk

`func (o *EECContext) GetCntxIdOk() (*string, bool)`

GetCntxIdOk returns a tuple with the CntxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntxId

`func (o *EECContext) SetCntxId(v string)`

SetCntxId sets CntxId field to given value.


### GetUeId

`func (o *EECContext) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *EECContext) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *EECContext) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *EECContext) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetE1Subs

`func (o *EECContext) GetE1Subs() []string`

GetE1Subs returns the E1Subs field if non-nil, zero value otherwise.

### GetE1SubsOk

`func (o *EECContext) GetE1SubsOk() (*[]string, bool)`

GetE1SubsOk returns a tuple with the E1Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE1Subs

`func (o *EECContext) SetE1Subs(v []string)`

SetE1Subs sets E1Subs field to given value.

### HasE1Subs

`func (o *EECContext) HasE1Subs() bool`

HasE1Subs returns a boolean if a field has been set.

### GetUeLoc

`func (o *EECContext) GetUeLoc() LocationArea5G`

GetUeLoc returns the UeLoc field if non-nil, zero value otherwise.

### GetUeLocOk

`func (o *EECContext) GetUeLocOk() (*LocationArea5G, bool)`

GetUeLocOk returns a tuple with the UeLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLoc

`func (o *EECContext) SetUeLoc(v LocationArea5G)`

SetUeLoc sets UeLoc field to given value.

### HasUeLoc

`func (o *EECContext) HasUeLoc() bool`

HasUeLoc returns a boolean if a field has been set.

### GetAcProfs

`func (o *EECContext) GetAcProfs() []ACProfile`

GetAcProfs returns the AcProfs field if non-nil, zero value otherwise.

### GetAcProfsOk

`func (o *EECContext) GetAcProfsOk() (*[]ACProfile, bool)`

GetAcProfsOk returns a tuple with the AcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcProfs

`func (o *EECContext) SetAcProfs(v []ACProfile)`

SetAcProfs sets AcProfs field to given value.

### HasAcProfs

`func (o *EECContext) HasAcProfs() bool`

HasAcProfs returns a boolean if a field has been set.

### GetSessCntxs

`func (o *EECContext) GetSessCntxs() SessionContexts`

GetSessCntxs returns the SessCntxs field if non-nil, zero value otherwise.

### GetSessCntxsOk

`func (o *EECContext) GetSessCntxsOk() (*SessionContexts, bool)`

GetSessCntxsOk returns a tuple with the SessCntxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessCntxs

`func (o *EECContext) SetSessCntxs(v SessionContexts)`

SetSessCntxs sets SessCntxs field to given value.

### HasSessCntxs

`func (o *EECContext) HasSessCntxs() bool`

HasSessCntxs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


