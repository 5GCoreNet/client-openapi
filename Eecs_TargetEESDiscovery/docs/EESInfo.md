# EESInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EesId** | **string** | Identity of the EES | 
**EndPt** | Pointer to [**EndPoint**](EndPoint.md) |  | [optional] 
**EasIds** | Pointer to **[]string** | Application identities of the Edge Application Servers registered with the EES. | [optional] 
**EcspInfo** | Pointer to **string** | Represents an ECSP Information. | [optional] 
**SvcArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**Dnais** | Pointer to **[]string** | Represents list of Data network access identifier. | [optional] 
**EesSvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EES supports service continuity or not, also indicates which ACR scenarios are supported by the EES.  | [optional] 
**EecRegConf** | **bool** | Indicates whether the EEC is required to register on the EES to use edge services or not.  | 

## Methods

### NewEESInfo

`func NewEESInfo(eesId string, eecRegConf bool, ) *EESInfo`

NewEESInfo instantiates a new EESInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEESInfoWithDefaults

`func NewEESInfoWithDefaults() *EESInfo`

NewEESInfoWithDefaults instantiates a new EESInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEesId

`func (o *EESInfo) GetEesId() string`

GetEesId returns the EesId field if non-nil, zero value otherwise.

### GetEesIdOk

`func (o *EESInfo) GetEesIdOk() (*string, bool)`

GetEesIdOk returns a tuple with the EesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEesId

`func (o *EESInfo) SetEesId(v string)`

SetEesId sets EesId field to given value.


### GetEndPt

`func (o *EESInfo) GetEndPt() EndPoint`

GetEndPt returns the EndPt field if non-nil, zero value otherwise.

### GetEndPtOk

`func (o *EESInfo) GetEndPtOk() (*EndPoint, bool)`

GetEndPtOk returns a tuple with the EndPt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPt

`func (o *EESInfo) SetEndPt(v EndPoint)`

SetEndPt sets EndPt field to given value.

### HasEndPt

`func (o *EESInfo) HasEndPt() bool`

HasEndPt returns a boolean if a field has been set.

### GetEasIds

`func (o *EESInfo) GetEasIds() []string`

GetEasIds returns the EasIds field if non-nil, zero value otherwise.

### GetEasIdsOk

`func (o *EESInfo) GetEasIdsOk() (*[]string, bool)`

GetEasIdsOk returns a tuple with the EasIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIds

`func (o *EESInfo) SetEasIds(v []string)`

SetEasIds sets EasIds field to given value.

### HasEasIds

`func (o *EESInfo) HasEasIds() bool`

HasEasIds returns a boolean if a field has been set.

### GetEcspInfo

`func (o *EESInfo) GetEcspInfo() string`

GetEcspInfo returns the EcspInfo field if non-nil, zero value otherwise.

### GetEcspInfoOk

`func (o *EESInfo) GetEcspInfoOk() (*string, bool)`

GetEcspInfoOk returns a tuple with the EcspInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspInfo

`func (o *EESInfo) SetEcspInfo(v string)`

SetEcspInfo sets EcspInfo field to given value.

### HasEcspInfo

`func (o *EESInfo) HasEcspInfo() bool`

HasEcspInfo returns a boolean if a field has been set.

### GetSvcArea

`func (o *EESInfo) GetSvcArea() LocationArea5G`

GetSvcArea returns the SvcArea field if non-nil, zero value otherwise.

### GetSvcAreaOk

`func (o *EESInfo) GetSvcAreaOk() (*LocationArea5G, bool)`

GetSvcAreaOk returns a tuple with the SvcArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcArea

`func (o *EESInfo) SetSvcArea(v LocationArea5G)`

SetSvcArea sets SvcArea field to given value.

### HasSvcArea

`func (o *EESInfo) HasSvcArea() bool`

HasSvcArea returns a boolean if a field has been set.

### GetDnais

`func (o *EESInfo) GetDnais() []string`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *EESInfo) GetDnaisOk() (*[]string, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *EESInfo) SetDnais(v []string)`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *EESInfo) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetEesSvcContSupp

`func (o *EESInfo) GetEesSvcContSupp() []ACRScenario`

GetEesSvcContSupp returns the EesSvcContSupp field if non-nil, zero value otherwise.

### GetEesSvcContSuppOk

`func (o *EESInfo) GetEesSvcContSuppOk() (*[]ACRScenario, bool)`

GetEesSvcContSuppOk returns a tuple with the EesSvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEesSvcContSupp

`func (o *EESInfo) SetEesSvcContSupp(v []ACRScenario)`

SetEesSvcContSupp sets EesSvcContSupp field to given value.

### HasEesSvcContSupp

`func (o *EESInfo) HasEesSvcContSupp() bool`

HasEesSvcContSupp returns a boolean if a field has been set.

### GetEecRegConf

`func (o *EESInfo) GetEecRegConf() bool`

GetEecRegConf returns the EecRegConf field if non-nil, zero value otherwise.

### GetEecRegConfOk

`func (o *EESInfo) GetEecRegConfOk() (*bool, bool)`

GetEecRegConfOk returns a tuple with the EecRegConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecRegConf

`func (o *EESInfo) SetEecRegConf(v bool)`

SetEecRegConf sets EecRegConf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


