# EasDiscoveryReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestorId** | [**RequestorId**](RequestorId.md) |  | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**EasDiscoveryFilter** | Pointer to [**EasDiscoveryFilter**](EasDiscoveryFilter.md) |  | [optional] 
**EecSvcContinuity** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC. | [optional] 
**EesSvcContinuity** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC. | [optional] 
**EasSvcContinuity** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC. | [optional] 
**LocInf** | Pointer to [**LocationInfo**](LocationInfo.md) |  | [optional] 
**EasTDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 

## Methods

### NewEasDiscoveryReq

`func NewEasDiscoveryReq(requestorId RequestorId, ) *EasDiscoveryReq`

NewEasDiscoveryReq instantiates a new EasDiscoveryReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDiscoveryReqWithDefaults

`func NewEasDiscoveryReqWithDefaults() *EasDiscoveryReq`

NewEasDiscoveryReqWithDefaults instantiates a new EasDiscoveryReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestorId

`func (o *EasDiscoveryReq) GetRequestorId() RequestorId`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *EasDiscoveryReq) GetRequestorIdOk() (*RequestorId, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *EasDiscoveryReq) SetRequestorId(v RequestorId)`

SetRequestorId sets RequestorId field to given value.


### GetUeId

`func (o *EasDiscoveryReq) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *EasDiscoveryReq) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *EasDiscoveryReq) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *EasDiscoveryReq) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetEasDiscoveryFilter

`func (o *EasDiscoveryReq) GetEasDiscoveryFilter() EasDiscoveryFilter`

GetEasDiscoveryFilter returns the EasDiscoveryFilter field if non-nil, zero value otherwise.

### GetEasDiscoveryFilterOk

`func (o *EasDiscoveryReq) GetEasDiscoveryFilterOk() (*EasDiscoveryFilter, bool)`

GetEasDiscoveryFilterOk returns a tuple with the EasDiscoveryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDiscoveryFilter

`func (o *EasDiscoveryReq) SetEasDiscoveryFilter(v EasDiscoveryFilter)`

SetEasDiscoveryFilter sets EasDiscoveryFilter field to given value.

### HasEasDiscoveryFilter

`func (o *EasDiscoveryReq) HasEasDiscoveryFilter() bool`

HasEasDiscoveryFilter returns a boolean if a field has been set.

### GetEecSvcContinuity

`func (o *EasDiscoveryReq) GetEecSvcContinuity() []ACRScenario`

GetEecSvcContinuity returns the EecSvcContinuity field if non-nil, zero value otherwise.

### GetEecSvcContinuityOk

`func (o *EasDiscoveryReq) GetEecSvcContinuityOk() (*[]ACRScenario, bool)`

GetEecSvcContinuityOk returns a tuple with the EecSvcContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecSvcContinuity

`func (o *EasDiscoveryReq) SetEecSvcContinuity(v []ACRScenario)`

SetEecSvcContinuity sets EecSvcContinuity field to given value.

### HasEecSvcContinuity

`func (o *EasDiscoveryReq) HasEecSvcContinuity() bool`

HasEecSvcContinuity returns a boolean if a field has been set.

### GetEesSvcContinuity

`func (o *EasDiscoveryReq) GetEesSvcContinuity() []ACRScenario`

GetEesSvcContinuity returns the EesSvcContinuity field if non-nil, zero value otherwise.

### GetEesSvcContinuityOk

`func (o *EasDiscoveryReq) GetEesSvcContinuityOk() (*[]ACRScenario, bool)`

GetEesSvcContinuityOk returns a tuple with the EesSvcContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEesSvcContinuity

`func (o *EasDiscoveryReq) SetEesSvcContinuity(v []ACRScenario)`

SetEesSvcContinuity sets EesSvcContinuity field to given value.

### HasEesSvcContinuity

`func (o *EasDiscoveryReq) HasEesSvcContinuity() bool`

HasEesSvcContinuity returns a boolean if a field has been set.

### GetEasSvcContinuity

`func (o *EasDiscoveryReq) GetEasSvcContinuity() []ACRScenario`

GetEasSvcContinuity returns the EasSvcContinuity field if non-nil, zero value otherwise.

### GetEasSvcContinuityOk

`func (o *EasDiscoveryReq) GetEasSvcContinuityOk() (*[]ACRScenario, bool)`

GetEasSvcContinuityOk returns a tuple with the EasSvcContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasSvcContinuity

`func (o *EasDiscoveryReq) SetEasSvcContinuity(v []ACRScenario)`

SetEasSvcContinuity sets EasSvcContinuity field to given value.

### HasEasSvcContinuity

`func (o *EasDiscoveryReq) HasEasSvcContinuity() bool`

HasEasSvcContinuity returns a boolean if a field has been set.

### GetLocInf

`func (o *EasDiscoveryReq) GetLocInf() LocationInfo`

GetLocInf returns the LocInf field if non-nil, zero value otherwise.

### GetLocInfOk

`func (o *EasDiscoveryReq) GetLocInfOk() (*LocationInfo, bool)`

GetLocInfOk returns a tuple with the LocInf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInf

`func (o *EasDiscoveryReq) SetLocInf(v LocationInfo)`

SetLocInf sets LocInf field to given value.

### HasLocInf

`func (o *EasDiscoveryReq) HasLocInf() bool`

HasLocInf returns a boolean if a field has been set.

### GetEasTDnai

`func (o *EasDiscoveryReq) GetEasTDnai() string`

GetEasTDnai returns the EasTDnai field if non-nil, zero value otherwise.

### GetEasTDnaiOk

`func (o *EasDiscoveryReq) GetEasTDnaiOk() (*string, bool)`

GetEasTDnaiOk returns a tuple with the EasTDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasTDnai

`func (o *EasDiscoveryReq) SetEasTDnai(v string)`

SetEasTDnai sets EasTDnai field to given value.

### HasEasTDnai

`func (o *EasDiscoveryReq) HasEasTDnai() bool`

HasEasTDnai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


