# AcrDetermReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestorId** | **string** |  | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AcId** | Pointer to **string** |  | [optional] 
**EasId** | Pointer to **string** |  | [optional] 
**SEasEndpoint** | [**EndPoint**](EndPoint.md) |  | 

## Methods

### NewAcrDetermReq

`func NewAcrDetermReq(requestorId string, sEasEndpoint EndPoint, ) *AcrDetermReq`

NewAcrDetermReq instantiates a new AcrDetermReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrDetermReqWithDefaults

`func NewAcrDetermReqWithDefaults() *AcrDetermReq`

NewAcrDetermReqWithDefaults instantiates a new AcrDetermReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestorId

`func (o *AcrDetermReq) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *AcrDetermReq) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *AcrDetermReq) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.


### GetUeId

`func (o *AcrDetermReq) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *AcrDetermReq) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *AcrDetermReq) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *AcrDetermReq) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetAcId

`func (o *AcrDetermReq) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *AcrDetermReq) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *AcrDetermReq) SetAcId(v string)`

SetAcId sets AcId field to given value.

### HasAcId

`func (o *AcrDetermReq) HasAcId() bool`

HasAcId returns a boolean if a field has been set.

### GetEasId

`func (o *AcrDetermReq) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *AcrDetermReq) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *AcrDetermReq) SetEasId(v string)`

SetEasId sets EasId field to given value.

### HasEasId

`func (o *AcrDetermReq) HasEasId() bool`

HasEasId returns a boolean if a field has been set.

### GetSEasEndpoint

`func (o *AcrDetermReq) GetSEasEndpoint() EndPoint`

GetSEasEndpoint returns the SEasEndpoint field if non-nil, zero value otherwise.

### GetSEasEndpointOk

`func (o *AcrDetermReq) GetSEasEndpointOk() (*EndPoint, bool)`

GetSEasEndpointOk returns a tuple with the SEasEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEasEndpoint

`func (o *AcrDetermReq) SetSEasEndpoint(v EndPoint)`

SetSEasEndpoint sets SEasEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


