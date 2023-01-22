# IndUeIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 

## Methods

### NewIndUeIdentification

`func NewIndUeIdentification() *IndUeIdentification`

NewIndUeIdentification instantiates a new IndUeIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndUeIdentificationWithDefaults

`func NewIndUeIdentificationWithDefaults() *IndUeIdentification`

NewIndUeIdentificationWithDefaults instantiates a new IndUeIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *IndUeIdentification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *IndUeIdentification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *IndUeIdentification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *IndUeIdentification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetExternalId

`func (o *IndUeIdentification) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IndUeIdentification) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IndUeIdentification) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IndUeIdentification) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *IndUeIdentification) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *IndUeIdentification) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *IndUeIdentification) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *IndUeIdentification) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


