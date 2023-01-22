# Pfd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PfdId** | **string** | Identifies a PDF of an application identifier. | 
**FlowDescriptions** | Pointer to **[]string** | Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic. The content of the string has the same encoding as the IPFilterRule AVP value as defined in IETF RFC 6733. | [optional] 
**Urls** | Pointer to **[]string** | Indicates a URL or a regular expression which is used to match the significant parts of the URL. | [optional] 
**DomainNames** | Pointer to **[]string** | Indicates an FQDN or a regular expression as a domain name matching criteria. | [optional] 
**DnProtocol** | Pointer to [**DomainNameProtocol**](DomainNameProtocol.md) |  | [optional] 

## Methods

### NewPfd

`func NewPfd(pfdId string, ) *Pfd`

NewPfd instantiates a new Pfd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdWithDefaults

`func NewPfdWithDefaults() *Pfd`

NewPfdWithDefaults instantiates a new Pfd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPfdId

`func (o *Pfd) GetPfdId() string`

GetPfdId returns the PfdId field if non-nil, zero value otherwise.

### GetPfdIdOk

`func (o *Pfd) GetPfdIdOk() (*string, bool)`

GetPfdIdOk returns a tuple with the PfdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdId

`func (o *Pfd) SetPfdId(v string)`

SetPfdId sets PfdId field to given value.


### GetFlowDescriptions

`func (o *Pfd) GetFlowDescriptions() []string`

GetFlowDescriptions returns the FlowDescriptions field if non-nil, zero value otherwise.

### GetFlowDescriptionsOk

`func (o *Pfd) GetFlowDescriptionsOk() (*[]string, bool)`

GetFlowDescriptionsOk returns a tuple with the FlowDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescriptions

`func (o *Pfd) SetFlowDescriptions(v []string)`

SetFlowDescriptions sets FlowDescriptions field to given value.

### HasFlowDescriptions

`func (o *Pfd) HasFlowDescriptions() bool`

HasFlowDescriptions returns a boolean if a field has been set.

### GetUrls

`func (o *Pfd) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Pfd) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Pfd) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *Pfd) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetDomainNames

`func (o *Pfd) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *Pfd) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *Pfd) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *Pfd) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetDnProtocol

`func (o *Pfd) GetDnProtocol() DomainNameProtocol`

GetDnProtocol returns the DnProtocol field if non-nil, zero value otherwise.

### GetDnProtocolOk

`func (o *Pfd) GetDnProtocolOk() (*DomainNameProtocol, bool)`

GetDnProtocolOk returns a tuple with the DnProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnProtocol

`func (o *Pfd) SetDnProtocol(v DomainNameProtocol)`

SetDnProtocol sets DnProtocol field to given value.

### HasDnProtocol

`func (o *Pfd) HasDnProtocol() bool`

HasDnProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


