# UeContextInPgwData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgwInfo** | Pointer to [**[]PgwInfo**](PgwInfo.md) |  | [optional] 
**EmergencyFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**EmergencyPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**EmergencyIpAddr** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**EmergencyRegistrationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewUeContextInPgwData

`func NewUeContextInPgwData() *UeContextInPgwData`

NewUeContextInPgwData instantiates a new UeContextInPgwData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextInPgwDataWithDefaults

`func NewUeContextInPgwDataWithDefaults() *UeContextInPgwData`

NewUeContextInPgwDataWithDefaults instantiates a new UeContextInPgwData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPgwInfo

`func (o *UeContextInPgwData) GetPgwInfo() []PgwInfo`

GetPgwInfo returns the PgwInfo field if non-nil, zero value otherwise.

### GetPgwInfoOk

`func (o *UeContextInPgwData) GetPgwInfoOk() (*[]PgwInfo, bool)`

GetPgwInfoOk returns a tuple with the PgwInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwInfo

`func (o *UeContextInPgwData) SetPgwInfo(v []PgwInfo)`

SetPgwInfo sets PgwInfo field to given value.

### HasPgwInfo

`func (o *UeContextInPgwData) HasPgwInfo() bool`

HasPgwInfo returns a boolean if a field has been set.

### GetEmergencyFqdn

`func (o *UeContextInPgwData) GetEmergencyFqdn() string`

GetEmergencyFqdn returns the EmergencyFqdn field if non-nil, zero value otherwise.

### GetEmergencyFqdnOk

`func (o *UeContextInPgwData) GetEmergencyFqdnOk() (*string, bool)`

GetEmergencyFqdnOk returns a tuple with the EmergencyFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyFqdn

`func (o *UeContextInPgwData) SetEmergencyFqdn(v string)`

SetEmergencyFqdn sets EmergencyFqdn field to given value.

### HasEmergencyFqdn

`func (o *UeContextInPgwData) HasEmergencyFqdn() bool`

HasEmergencyFqdn returns a boolean if a field has been set.

### GetEmergencyPlmnId

`func (o *UeContextInPgwData) GetEmergencyPlmnId() PlmnId`

GetEmergencyPlmnId returns the EmergencyPlmnId field if non-nil, zero value otherwise.

### GetEmergencyPlmnIdOk

`func (o *UeContextInPgwData) GetEmergencyPlmnIdOk() (*PlmnId, bool)`

GetEmergencyPlmnIdOk returns a tuple with the EmergencyPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyPlmnId

`func (o *UeContextInPgwData) SetEmergencyPlmnId(v PlmnId)`

SetEmergencyPlmnId sets EmergencyPlmnId field to given value.

### HasEmergencyPlmnId

`func (o *UeContextInPgwData) HasEmergencyPlmnId() bool`

HasEmergencyPlmnId returns a boolean if a field has been set.

### GetEmergencyIpAddr

`func (o *UeContextInPgwData) GetEmergencyIpAddr() IpAddress`

GetEmergencyIpAddr returns the EmergencyIpAddr field if non-nil, zero value otherwise.

### GetEmergencyIpAddrOk

`func (o *UeContextInPgwData) GetEmergencyIpAddrOk() (*IpAddress, bool)`

GetEmergencyIpAddrOk returns a tuple with the EmergencyIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyIpAddr

`func (o *UeContextInPgwData) SetEmergencyIpAddr(v IpAddress)`

SetEmergencyIpAddr sets EmergencyIpAddr field to given value.

### HasEmergencyIpAddr

`func (o *UeContextInPgwData) HasEmergencyIpAddr() bool`

HasEmergencyIpAddr returns a boolean if a field has been set.

### GetEmergencyRegistrationTime

`func (o *UeContextInPgwData) GetEmergencyRegistrationTime() time.Time`

GetEmergencyRegistrationTime returns the EmergencyRegistrationTime field if non-nil, zero value otherwise.

### GetEmergencyRegistrationTimeOk

`func (o *UeContextInPgwData) GetEmergencyRegistrationTimeOk() (*time.Time, bool)`

GetEmergencyRegistrationTimeOk returns a tuple with the EmergencyRegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyRegistrationTime

`func (o *UeContextInPgwData) SetEmergencyRegistrationTime(v time.Time)`

SetEmergencyRegistrationTime sets EmergencyRegistrationTime field to given value.

### HasEmergencyRegistrationTime

`func (o *UeContextInPgwData) HasEmergencyRegistrationTime() bool`

HasEmergencyRegistrationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


