# EECRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EecId** | **string** | Represents a unique identifier of the EEC. | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AcProfs** | Pointer to [**[]ACProfile**](ACProfile.md) | Profiles of ACs for which the EEC provides edge enabling services. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EecSvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Profiles of ACs for which the EEC provides edge enabling services. | [optional] 
**EecCntxId** | Pointer to **string** | Identifier of the EEC context obtained from a previous registration. | [optional] 
**SrcEesId** | Pointer to **string** | Identifier of the EES that provided EEC context ID. | [optional] 
**EndPt** | Pointer to [**EndPoint**](EndPoint.md) |  | [optional] 
**UnfulfilledAcProfs** | Pointer to [**UnfulfilledAcProfile**](UnfulfilledAcProfile.md) |  | [optional] 

## Methods

### NewEECRegistration

`func NewEECRegistration(eecId string, ) *EECRegistration`

NewEECRegistration instantiates a new EECRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEECRegistrationWithDefaults

`func NewEECRegistrationWithDefaults() *EECRegistration`

NewEECRegistrationWithDefaults instantiates a new EECRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEecId

`func (o *EECRegistration) GetEecId() string`

GetEecId returns the EecId field if non-nil, zero value otherwise.

### GetEecIdOk

`func (o *EECRegistration) GetEecIdOk() (*string, bool)`

GetEecIdOk returns a tuple with the EecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecId

`func (o *EECRegistration) SetEecId(v string)`

SetEecId sets EecId field to given value.


### GetUeId

`func (o *EECRegistration) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *EECRegistration) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *EECRegistration) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *EECRegistration) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetAcProfs

`func (o *EECRegistration) GetAcProfs() []ACProfile`

GetAcProfs returns the AcProfs field if non-nil, zero value otherwise.

### GetAcProfsOk

`func (o *EECRegistration) GetAcProfsOk() (*[]ACProfile, bool)`

GetAcProfsOk returns a tuple with the AcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcProfs

`func (o *EECRegistration) SetAcProfs(v []ACProfile)`

SetAcProfs sets AcProfs field to given value.

### HasAcProfs

`func (o *EECRegistration) HasAcProfs() bool`

HasAcProfs returns a boolean if a field has been set.

### GetExpTime

`func (o *EECRegistration) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *EECRegistration) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *EECRegistration) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *EECRegistration) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetEecSvcContSupp

`func (o *EECRegistration) GetEecSvcContSupp() []ACRScenario`

GetEecSvcContSupp returns the EecSvcContSupp field if non-nil, zero value otherwise.

### GetEecSvcContSuppOk

`func (o *EECRegistration) GetEecSvcContSuppOk() (*[]ACRScenario, bool)`

GetEecSvcContSuppOk returns a tuple with the EecSvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecSvcContSupp

`func (o *EECRegistration) SetEecSvcContSupp(v []ACRScenario)`

SetEecSvcContSupp sets EecSvcContSupp field to given value.

### HasEecSvcContSupp

`func (o *EECRegistration) HasEecSvcContSupp() bool`

HasEecSvcContSupp returns a boolean if a field has been set.

### GetEecCntxId

`func (o *EECRegistration) GetEecCntxId() string`

GetEecCntxId returns the EecCntxId field if non-nil, zero value otherwise.

### GetEecCntxIdOk

`func (o *EECRegistration) GetEecCntxIdOk() (*string, bool)`

GetEecCntxIdOk returns a tuple with the EecCntxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecCntxId

`func (o *EECRegistration) SetEecCntxId(v string)`

SetEecCntxId sets EecCntxId field to given value.

### HasEecCntxId

`func (o *EECRegistration) HasEecCntxId() bool`

HasEecCntxId returns a boolean if a field has been set.

### GetSrcEesId

`func (o *EECRegistration) GetSrcEesId() string`

GetSrcEesId returns the SrcEesId field if non-nil, zero value otherwise.

### GetSrcEesIdOk

`func (o *EECRegistration) GetSrcEesIdOk() (*string, bool)`

GetSrcEesIdOk returns a tuple with the SrcEesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcEesId

`func (o *EECRegistration) SetSrcEesId(v string)`

SetSrcEesId sets SrcEesId field to given value.

### HasSrcEesId

`func (o *EECRegistration) HasSrcEesId() bool`

HasSrcEesId returns a boolean if a field has been set.

### GetEndPt

`func (o *EECRegistration) GetEndPt() EndPoint`

GetEndPt returns the EndPt field if non-nil, zero value otherwise.

### GetEndPtOk

`func (o *EECRegistration) GetEndPtOk() (*EndPoint, bool)`

GetEndPtOk returns a tuple with the EndPt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPt

`func (o *EECRegistration) SetEndPt(v EndPoint)`

SetEndPt sets EndPt field to given value.

### HasEndPt

`func (o *EECRegistration) HasEndPt() bool`

HasEndPt returns a boolean if a field has been set.

### GetUnfulfilledAcProfs

`func (o *EECRegistration) GetUnfulfilledAcProfs() UnfulfilledAcProfile`

GetUnfulfilledAcProfs returns the UnfulfilledAcProfs field if non-nil, zero value otherwise.

### GetUnfulfilledAcProfsOk

`func (o *EECRegistration) GetUnfulfilledAcProfsOk() (*UnfulfilledAcProfile, bool)`

GetUnfulfilledAcProfsOk returns a tuple with the UnfulfilledAcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfulfilledAcProfs

`func (o *EECRegistration) SetUnfulfilledAcProfs(v UnfulfilledAcProfile)`

SetUnfulfilledAcProfs sets UnfulfilledAcProfs field to given value.

### HasUnfulfilledAcProfs

`func (o *EECRegistration) HasUnfulfilledAcProfs() bool`

HasUnfulfilledAcProfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


