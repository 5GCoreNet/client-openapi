# ServerCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MandatoryCapability** | Pointer to **[]int32** |  | [optional] 
**OptionalCapability** | Pointer to **[]int32** |  | [optional] 
**ServerName** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServerCapabilities

`func NewServerCapabilities() *ServerCapabilities`

NewServerCapabilities instantiates a new ServerCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCapabilitiesWithDefaults

`func NewServerCapabilitiesWithDefaults() *ServerCapabilities`

NewServerCapabilitiesWithDefaults instantiates a new ServerCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMandatoryCapability

`func (o *ServerCapabilities) GetMandatoryCapability() []int32`

GetMandatoryCapability returns the MandatoryCapability field if non-nil, zero value otherwise.

### GetMandatoryCapabilityOk

`func (o *ServerCapabilities) GetMandatoryCapabilityOk() (*[]int32, bool)`

GetMandatoryCapabilityOk returns a tuple with the MandatoryCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryCapability

`func (o *ServerCapabilities) SetMandatoryCapability(v []int32)`

SetMandatoryCapability sets MandatoryCapability field to given value.

### HasMandatoryCapability

`func (o *ServerCapabilities) HasMandatoryCapability() bool`

HasMandatoryCapability returns a boolean if a field has been set.

### GetOptionalCapability

`func (o *ServerCapabilities) GetOptionalCapability() []int32`

GetOptionalCapability returns the OptionalCapability field if non-nil, zero value otherwise.

### GetOptionalCapabilityOk

`func (o *ServerCapabilities) GetOptionalCapabilityOk() (*[]int32, bool)`

GetOptionalCapabilityOk returns a tuple with the OptionalCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalCapability

`func (o *ServerCapabilities) SetOptionalCapability(v []int32)`

SetOptionalCapability sets OptionalCapability field to given value.

### HasOptionalCapability

`func (o *ServerCapabilities) HasOptionalCapability() bool`

HasOptionalCapability returns a boolean if a field has been set.

### GetServerName

`func (o *ServerCapabilities) GetServerName() []string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerCapabilities) GetServerNameOk() (*[]string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerCapabilities) SetServerName(v []string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ServerCapabilities) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


