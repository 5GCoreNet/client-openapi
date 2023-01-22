# UavId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**CaaId** | Pointer to **string** |  | [optional] 

## Methods

### NewUavId

`func NewUavId() *UavId`

NewUavId instantiates a new UavId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUavIdWithDefaults

`func NewUavIdWithDefaults() *UavId`

NewUavIdWithDefaults instantiates a new UavId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *UavId) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UavId) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UavId) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *UavId) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetCaaId

`func (o *UavId) GetCaaId() string`

GetCaaId returns the CaaId field if non-nil, zero value otherwise.

### GetCaaIdOk

`func (o *UavId) GetCaaIdOk() (*string, bool)`

GetCaaIdOk returns a tuple with the CaaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaaId

`func (o *UavId) SetCaaId(v string)`

SetCaaId sets CaaId field to given value.

### HasCaaId

`func (o *UavId) HasCaaId() bool`

HasCaaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


