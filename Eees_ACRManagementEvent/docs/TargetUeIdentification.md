# TargetUeIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**IntGrpId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**ExtGrpId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 

## Methods

### NewTargetUeIdentification

`func NewTargetUeIdentification() *TargetUeIdentification`

NewTargetUeIdentification instantiates a new TargetUeIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUeIdentificationWithDefaults

`func NewTargetUeIdentificationWithDefaults() *TargetUeIdentification`

NewTargetUeIdentificationWithDefaults instantiates a new TargetUeIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *TargetUeIdentification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *TargetUeIdentification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *TargetUeIdentification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *TargetUeIdentification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetIntGrpId

`func (o *TargetUeIdentification) GetIntGrpId() string`

GetIntGrpId returns the IntGrpId field if non-nil, zero value otherwise.

### GetIntGrpIdOk

`func (o *TargetUeIdentification) GetIntGrpIdOk() (*string, bool)`

GetIntGrpIdOk returns a tuple with the IntGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGrpId

`func (o *TargetUeIdentification) SetIntGrpId(v string)`

SetIntGrpId sets IntGrpId field to given value.

### HasIntGrpId

`func (o *TargetUeIdentification) HasIntGrpId() bool`

HasIntGrpId returns a boolean if a field has been set.

### GetExtGrpId

`func (o *TargetUeIdentification) GetExtGrpId() string`

GetExtGrpId returns the ExtGrpId field if non-nil, zero value otherwise.

### GetExtGrpIdOk

`func (o *TargetUeIdentification) GetExtGrpIdOk() (*string, bool)`

GetExtGrpIdOk returns a tuple with the ExtGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGrpId

`func (o *TargetUeIdentification) SetExtGrpId(v string)`

SetExtGrpId sets ExtGrpId field to given value.

### HasExtGrpId

`func (o *TargetUeIdentification) HasExtGrpId() bool`

HasExtGrpId returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *TargetUeIdentification) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *TargetUeIdentification) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *TargetUeIdentification) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *TargetUeIdentification) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


