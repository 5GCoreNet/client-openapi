# AefProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AefId** | **string** | Identifier of the API exposing function | 
**Versions** | [**[]Version**](Version.md) | API version | 
**Protocol** | Pointer to [**Protocol**](Protocol.md) |  | [optional] 
**DataFormat** | Pointer to [**DataFormat**](DataFormat.md) |  | [optional] 
**SecurityMethods** | Pointer to [**[]SecurityMethod**](SecurityMethod.md) | Security methods supported by the AEF | [optional] 
**DomainName** | Pointer to **string** | Domain to which API belongs to | [optional] 
**InterfaceDescriptions** | Pointer to [**[]InterfaceDescription**](InterfaceDescription.md) | Interface details | [optional] 
**AefLocation** | Pointer to [**AefLocation**](AefLocation.md) |  | [optional] 

## Methods

### NewAefProfile

`func NewAefProfile(aefId string, versions []Version, ) *AefProfile`

NewAefProfile instantiates a new AefProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAefProfileWithDefaults

`func NewAefProfileWithDefaults() *AefProfile`

NewAefProfileWithDefaults instantiates a new AefProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAefId

`func (o *AefProfile) GetAefId() string`

GetAefId returns the AefId field if non-nil, zero value otherwise.

### GetAefIdOk

`func (o *AefProfile) GetAefIdOk() (*string, bool)`

GetAefIdOk returns a tuple with the AefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefId

`func (o *AefProfile) SetAefId(v string)`

SetAefId sets AefId field to given value.


### GetVersions

`func (o *AefProfile) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AefProfile) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AefProfile) SetVersions(v []Version)`

SetVersions sets Versions field to given value.


### GetProtocol

`func (o *AefProfile) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AefProfile) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AefProfile) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AefProfile) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDataFormat

`func (o *AefProfile) GetDataFormat() DataFormat`

GetDataFormat returns the DataFormat field if non-nil, zero value otherwise.

### GetDataFormatOk

`func (o *AefProfile) GetDataFormatOk() (*DataFormat, bool)`

GetDataFormatOk returns a tuple with the DataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFormat

`func (o *AefProfile) SetDataFormat(v DataFormat)`

SetDataFormat sets DataFormat field to given value.

### HasDataFormat

`func (o *AefProfile) HasDataFormat() bool`

HasDataFormat returns a boolean if a field has been set.

### GetSecurityMethods

`func (o *AefProfile) GetSecurityMethods() []SecurityMethod`

GetSecurityMethods returns the SecurityMethods field if non-nil, zero value otherwise.

### GetSecurityMethodsOk

`func (o *AefProfile) GetSecurityMethodsOk() (*[]SecurityMethod, bool)`

GetSecurityMethodsOk returns a tuple with the SecurityMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMethods

`func (o *AefProfile) SetSecurityMethods(v []SecurityMethod)`

SetSecurityMethods sets SecurityMethods field to given value.

### HasSecurityMethods

`func (o *AefProfile) HasSecurityMethods() bool`

HasSecurityMethods returns a boolean if a field has been set.

### GetDomainName

`func (o *AefProfile) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AefProfile) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AefProfile) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AefProfile) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetInterfaceDescriptions

`func (o *AefProfile) GetInterfaceDescriptions() []InterfaceDescription`

GetInterfaceDescriptions returns the InterfaceDescriptions field if non-nil, zero value otherwise.

### GetInterfaceDescriptionsOk

`func (o *AefProfile) GetInterfaceDescriptionsOk() (*[]InterfaceDescription, bool)`

GetInterfaceDescriptionsOk returns a tuple with the InterfaceDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceDescriptions

`func (o *AefProfile) SetInterfaceDescriptions(v []InterfaceDescription)`

SetInterfaceDescriptions sets InterfaceDescriptions field to given value.

### HasInterfaceDescriptions

`func (o *AefProfile) HasInterfaceDescriptions() bool`

HasInterfaceDescriptions returns a boolean if a field has been set.

### GetAefLocation

`func (o *AefProfile) GetAefLocation() AefLocation`

GetAefLocation returns the AefLocation field if non-nil, zero value otherwise.

### GetAefLocationOk

`func (o *AefProfile) GetAefLocationOk() (*AefLocation, bool)`

GetAefLocationOk returns a tuple with the AefLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefLocation

`func (o *AefProfile) SetAefLocation(v AefLocation)`

SetAefLocation sets AefLocation field to given value.

### HasAefLocation

`func (o *AefProfile) HasAefLocation() bool`

HasAefLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


