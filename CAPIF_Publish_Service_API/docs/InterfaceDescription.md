# InterfaceDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**Ipv6Addr** | Pointer to **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | [optional] 
**Port** | Pointer to **int32** | Unsigned integer with valid values between 0 and 65535. | [optional] 
**SecurityMethods** | Pointer to [**[]SecurityMethod**](SecurityMethod.md) | Security methods supported by the interface, it take precedence over the security methods provided in AefProfile, for this specific interface.  | [optional] 

## Methods

### NewInterfaceDescription

`func NewInterfaceDescription() *InterfaceDescription`

NewInterfaceDescription instantiates a new InterfaceDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceDescriptionWithDefaults

`func NewInterfaceDescriptionWithDefaults() *InterfaceDescription`

NewInterfaceDescriptionWithDefaults instantiates a new InterfaceDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addr

`func (o *InterfaceDescription) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *InterfaceDescription) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *InterfaceDescription) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *InterfaceDescription) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *InterfaceDescription) GetIpv6Addr() string`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *InterfaceDescription) GetIpv6AddrOk() (*string, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *InterfaceDescription) SetIpv6Addr(v string)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *InterfaceDescription) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.

### GetPort

`func (o *InterfaceDescription) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InterfaceDescription) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InterfaceDescription) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InterfaceDescription) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurityMethods

`func (o *InterfaceDescription) GetSecurityMethods() []SecurityMethod`

GetSecurityMethods returns the SecurityMethods field if non-nil, zero value otherwise.

### GetSecurityMethodsOk

`func (o *InterfaceDescription) GetSecurityMethodsOk() (*[]SecurityMethod, bool)`

GetSecurityMethodsOk returns a tuple with the SecurityMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMethods

`func (o *InterfaceDescription) SetSecurityMethods(v []SecurityMethod)`

SetSecurityMethods sets SecurityMethods field to given value.

### HasSecurityMethods

`func (o *InterfaceDescription) HasSecurityMethods() bool`

HasSecurityMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


