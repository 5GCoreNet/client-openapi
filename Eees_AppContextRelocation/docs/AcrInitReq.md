# AcrInitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestorId** | **string** |  | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AcId** | Pointer to **string** |  | [optional] 
**EasId** | Pointer to **string** |  | [optional] 
**TEasEndpoint** | [**EndPoint**](EndPoint.md) |  | 
**SEasEndpoint** | Pointer to [**EndPoint**](EndPoint.md) |  | [optional] 
**PrevTEasEndpoint** | Pointer to [**EndPoint**](EndPoint.md) |  | [optional] 
**RouteReq** | Pointer to [**NullableRouteToLocation**](RouteToLocation.md) |  | [optional] 
**EasNotifInd** | **bool** |  | [default to false]
**PrevEasNotifInd** | Pointer to **bool** |  | [optional] [default to false]
**EecCtxtReloc** | Pointer to [**EecCtxtReloc**](EecCtxtReloc.md) |  | [optional] 

## Methods

### NewAcrInitReq

`func NewAcrInitReq(requestorId string, tEasEndpoint EndPoint, easNotifInd bool, ) *AcrInitReq`

NewAcrInitReq instantiates a new AcrInitReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrInitReqWithDefaults

`func NewAcrInitReqWithDefaults() *AcrInitReq`

NewAcrInitReqWithDefaults instantiates a new AcrInitReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestorId

`func (o *AcrInitReq) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *AcrInitReq) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *AcrInitReq) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.


### GetUeId

`func (o *AcrInitReq) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *AcrInitReq) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *AcrInitReq) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *AcrInitReq) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetAcId

`func (o *AcrInitReq) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *AcrInitReq) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *AcrInitReq) SetAcId(v string)`

SetAcId sets AcId field to given value.

### HasAcId

`func (o *AcrInitReq) HasAcId() bool`

HasAcId returns a boolean if a field has been set.

### GetEasId

`func (o *AcrInitReq) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *AcrInitReq) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *AcrInitReq) SetEasId(v string)`

SetEasId sets EasId field to given value.

### HasEasId

`func (o *AcrInitReq) HasEasId() bool`

HasEasId returns a boolean if a field has been set.

### GetTEasEndpoint

`func (o *AcrInitReq) GetTEasEndpoint() EndPoint`

GetTEasEndpoint returns the TEasEndpoint field if non-nil, zero value otherwise.

### GetTEasEndpointOk

`func (o *AcrInitReq) GetTEasEndpointOk() (*EndPoint, bool)`

GetTEasEndpointOk returns a tuple with the TEasEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEasEndpoint

`func (o *AcrInitReq) SetTEasEndpoint(v EndPoint)`

SetTEasEndpoint sets TEasEndpoint field to given value.


### GetSEasEndpoint

`func (o *AcrInitReq) GetSEasEndpoint() EndPoint`

GetSEasEndpoint returns the SEasEndpoint field if non-nil, zero value otherwise.

### GetSEasEndpointOk

`func (o *AcrInitReq) GetSEasEndpointOk() (*EndPoint, bool)`

GetSEasEndpointOk returns a tuple with the SEasEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEasEndpoint

`func (o *AcrInitReq) SetSEasEndpoint(v EndPoint)`

SetSEasEndpoint sets SEasEndpoint field to given value.

### HasSEasEndpoint

`func (o *AcrInitReq) HasSEasEndpoint() bool`

HasSEasEndpoint returns a boolean if a field has been set.

### GetPrevTEasEndpoint

`func (o *AcrInitReq) GetPrevTEasEndpoint() EndPoint`

GetPrevTEasEndpoint returns the PrevTEasEndpoint field if non-nil, zero value otherwise.

### GetPrevTEasEndpointOk

`func (o *AcrInitReq) GetPrevTEasEndpointOk() (*EndPoint, bool)`

GetPrevTEasEndpointOk returns a tuple with the PrevTEasEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTEasEndpoint

`func (o *AcrInitReq) SetPrevTEasEndpoint(v EndPoint)`

SetPrevTEasEndpoint sets PrevTEasEndpoint field to given value.

### HasPrevTEasEndpoint

`func (o *AcrInitReq) HasPrevTEasEndpoint() bool`

HasPrevTEasEndpoint returns a boolean if a field has been set.

### GetRouteReq

`func (o *AcrInitReq) GetRouteReq() RouteToLocation`

GetRouteReq returns the RouteReq field if non-nil, zero value otherwise.

### GetRouteReqOk

`func (o *AcrInitReq) GetRouteReqOk() (*RouteToLocation, bool)`

GetRouteReqOk returns a tuple with the RouteReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteReq

`func (o *AcrInitReq) SetRouteReq(v RouteToLocation)`

SetRouteReq sets RouteReq field to given value.

### HasRouteReq

`func (o *AcrInitReq) HasRouteReq() bool`

HasRouteReq returns a boolean if a field has been set.

### SetRouteReqNil

`func (o *AcrInitReq) SetRouteReqNil(b bool)`

 SetRouteReqNil sets the value for RouteReq to be an explicit nil

### UnsetRouteReq
`func (o *AcrInitReq) UnsetRouteReq()`

UnsetRouteReq ensures that no value is present for RouteReq, not even an explicit nil
### GetEasNotifInd

`func (o *AcrInitReq) GetEasNotifInd() bool`

GetEasNotifInd returns the EasNotifInd field if non-nil, zero value otherwise.

### GetEasNotifIndOk

`func (o *AcrInitReq) GetEasNotifIndOk() (*bool, bool)`

GetEasNotifIndOk returns a tuple with the EasNotifInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasNotifInd

`func (o *AcrInitReq) SetEasNotifInd(v bool)`

SetEasNotifInd sets EasNotifInd field to given value.


### GetPrevEasNotifInd

`func (o *AcrInitReq) GetPrevEasNotifInd() bool`

GetPrevEasNotifInd returns the PrevEasNotifInd field if non-nil, zero value otherwise.

### GetPrevEasNotifIndOk

`func (o *AcrInitReq) GetPrevEasNotifIndOk() (*bool, bool)`

GetPrevEasNotifIndOk returns a tuple with the PrevEasNotifInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevEasNotifInd

`func (o *AcrInitReq) SetPrevEasNotifInd(v bool)`

SetPrevEasNotifInd sets PrevEasNotifInd field to given value.

### HasPrevEasNotifInd

`func (o *AcrInitReq) HasPrevEasNotifInd() bool`

HasPrevEasNotifInd returns a boolean if a field has been set.

### GetEecCtxtReloc

`func (o *AcrInitReq) GetEecCtxtReloc() EecCtxtReloc`

GetEecCtxtReloc returns the EecCtxtReloc field if non-nil, zero value otherwise.

### GetEecCtxtRelocOk

`func (o *AcrInitReq) GetEecCtxtRelocOk() (*EecCtxtReloc, bool)`

GetEecCtxtRelocOk returns a tuple with the EecCtxtReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecCtxtReloc

`func (o *AcrInitReq) SetEecCtxtReloc(v EecCtxtReloc)`

SetEecCtxtReloc sets EecCtxtReloc field to given value.

### HasEecCtxtReloc

`func (o *AcrInitReq) HasEecCtxtReloc() bool`

HasEecCtxtReloc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


