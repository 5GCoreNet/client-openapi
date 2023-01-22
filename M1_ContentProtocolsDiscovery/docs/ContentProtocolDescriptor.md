# ContentProtocolDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermIdentifier** | **string** | String providing an URI formatted according to RFC 3986. | 
**DescriptionLocator** | Pointer to **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | [optional] 

## Methods

### NewContentProtocolDescriptor

`func NewContentProtocolDescriptor(termIdentifier string, ) *ContentProtocolDescriptor`

NewContentProtocolDescriptor instantiates a new ContentProtocolDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentProtocolDescriptorWithDefaults

`func NewContentProtocolDescriptorWithDefaults() *ContentProtocolDescriptor`

NewContentProtocolDescriptorWithDefaults instantiates a new ContentProtocolDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermIdentifier

`func (o *ContentProtocolDescriptor) GetTermIdentifier() string`

GetTermIdentifier returns the TermIdentifier field if non-nil, zero value otherwise.

### GetTermIdentifierOk

`func (o *ContentProtocolDescriptor) GetTermIdentifierOk() (*string, bool)`

GetTermIdentifierOk returns a tuple with the TermIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermIdentifier

`func (o *ContentProtocolDescriptor) SetTermIdentifier(v string)`

SetTermIdentifier sets TermIdentifier field to given value.


### GetDescriptionLocator

`func (o *ContentProtocolDescriptor) GetDescriptionLocator() string`

GetDescriptionLocator returns the DescriptionLocator field if non-nil, zero value otherwise.

### GetDescriptionLocatorOk

`func (o *ContentProtocolDescriptor) GetDescriptionLocatorOk() (*string, bool)`

GetDescriptionLocatorOk returns a tuple with the DescriptionLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocator

`func (o *ContentProtocolDescriptor) SetDescriptionLocator(v string)`

SetDescriptionLocator sets DescriptionLocator field to given value.

### HasDescriptionLocator

`func (o *ContentProtocolDescriptor) HasDescriptionLocator() bool`

HasDescriptionLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


