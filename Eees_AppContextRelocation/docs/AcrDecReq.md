# AcrDecReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**AcId** | Pointer to **string** |  | [optional] 
**TEasId** | **string** |  | 
**TEasEndpoint** | [**EndPoint**](EndPoint.md) |  | 

## Methods

### NewAcrDecReq

`func NewAcrDecReq(ueId string, tEasId string, tEasEndpoint EndPoint, ) *AcrDecReq`

NewAcrDecReq instantiates a new AcrDecReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrDecReqWithDefaults

`func NewAcrDecReqWithDefaults() *AcrDecReq`

NewAcrDecReqWithDefaults instantiates a new AcrDecReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *AcrDecReq) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *AcrDecReq) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *AcrDecReq) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetAcId

`func (o *AcrDecReq) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *AcrDecReq) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *AcrDecReq) SetAcId(v string)`

SetAcId sets AcId field to given value.

### HasAcId

`func (o *AcrDecReq) HasAcId() bool`

HasAcId returns a boolean if a field has been set.

### GetTEasId

`func (o *AcrDecReq) GetTEasId() string`

GetTEasId returns the TEasId field if non-nil, zero value otherwise.

### GetTEasIdOk

`func (o *AcrDecReq) GetTEasIdOk() (*string, bool)`

GetTEasIdOk returns a tuple with the TEasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEasId

`func (o *AcrDecReq) SetTEasId(v string)`

SetTEasId sets TEasId field to given value.


### GetTEasEndpoint

`func (o *AcrDecReq) GetTEasEndpoint() EndPoint`

GetTEasEndpoint returns the TEasEndpoint field if non-nil, zero value otherwise.

### GetTEasEndpointOk

`func (o *AcrDecReq) GetTEasEndpointOk() (*EndPoint, bool)`

GetTEasEndpointOk returns a tuple with the TEasEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEasEndpoint

`func (o *AcrDecReq) SetTEasEndpoint(v EndPoint)`

SetTEasEndpoint sets TEasEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


