# ECSServProvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EecId** | **string** | Represents a unique identifier of the EEC. | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AcProfs** | Pointer to [**[]ACProfile**](ACProfile.md) | Information about services the EEC wants to connect to. | [optional] 
**EecSvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.  | [optional] 
**ConnInfo** | Pointer to [**[]ConnectivityInfo**](ConnectivityInfo.md) | List of connectivity information for the UE. | [optional] 
**LocInf** | Pointer to [**LocationInfo**](LocationInfo.md) |  | [optional] 

## Methods

### NewECSServProvReq

`func NewECSServProvReq(eecId string, ) *ECSServProvReq`

NewECSServProvReq instantiates a new ECSServProvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSServProvReqWithDefaults

`func NewECSServProvReqWithDefaults() *ECSServProvReq`

NewECSServProvReqWithDefaults instantiates a new ECSServProvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEecId

`func (o *ECSServProvReq) GetEecId() string`

GetEecId returns the EecId field if non-nil, zero value otherwise.

### GetEecIdOk

`func (o *ECSServProvReq) GetEecIdOk() (*string, bool)`

GetEecIdOk returns a tuple with the EecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecId

`func (o *ECSServProvReq) SetEecId(v string)`

SetEecId sets EecId field to given value.


### GetUeId

`func (o *ECSServProvReq) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ECSServProvReq) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ECSServProvReq) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *ECSServProvReq) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetAcProfs

`func (o *ECSServProvReq) GetAcProfs() []ACProfile`

GetAcProfs returns the AcProfs field if non-nil, zero value otherwise.

### GetAcProfsOk

`func (o *ECSServProvReq) GetAcProfsOk() (*[]ACProfile, bool)`

GetAcProfsOk returns a tuple with the AcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcProfs

`func (o *ECSServProvReq) SetAcProfs(v []ACProfile)`

SetAcProfs sets AcProfs field to given value.

### HasAcProfs

`func (o *ECSServProvReq) HasAcProfs() bool`

HasAcProfs returns a boolean if a field has been set.

### GetEecSvcContSupp

`func (o *ECSServProvReq) GetEecSvcContSupp() []ACRScenario`

GetEecSvcContSupp returns the EecSvcContSupp field if non-nil, zero value otherwise.

### GetEecSvcContSuppOk

`func (o *ECSServProvReq) GetEecSvcContSuppOk() (*[]ACRScenario, bool)`

GetEecSvcContSuppOk returns a tuple with the EecSvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecSvcContSupp

`func (o *ECSServProvReq) SetEecSvcContSupp(v []ACRScenario)`

SetEecSvcContSupp sets EecSvcContSupp field to given value.

### HasEecSvcContSupp

`func (o *ECSServProvReq) HasEecSvcContSupp() bool`

HasEecSvcContSupp returns a boolean if a field has been set.

### GetConnInfo

`func (o *ECSServProvReq) GetConnInfo() []ConnectivityInfo`

GetConnInfo returns the ConnInfo field if non-nil, zero value otherwise.

### GetConnInfoOk

`func (o *ECSServProvReq) GetConnInfoOk() (*[]ConnectivityInfo, bool)`

GetConnInfoOk returns a tuple with the ConnInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnInfo

`func (o *ECSServProvReq) SetConnInfo(v []ConnectivityInfo)`

SetConnInfo sets ConnInfo field to given value.

### HasConnInfo

`func (o *ECSServProvReq) HasConnInfo() bool`

HasConnInfo returns a boolean if a field has been set.

### GetLocInf

`func (o *ECSServProvReq) GetLocInf() LocationInfo`

GetLocInf returns the LocInf field if non-nil, zero value otherwise.

### GetLocInfOk

`func (o *ECSServProvReq) GetLocInfOk() (*LocationInfo, bool)`

GetLocInfOk returns a tuple with the LocInf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInf

`func (o *ECSServProvReq) SetLocInf(v LocationInfo)`

SetLocInf sets LocInf field to given value.

### HasLocInf

`func (o *ECSServProvReq) HasLocInf() bool`

HasLocInf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


