# IngestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Pull** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**EntryPoint** | Pointer to **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | [optional] 

## Methods

### NewIngestConfiguration

`func NewIngestConfiguration() *IngestConfiguration`

NewIngestConfiguration instantiates a new IngestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestConfigurationWithDefaults

`func NewIngestConfigurationWithDefaults() *IngestConfiguration`

NewIngestConfigurationWithDefaults instantiates a new IngestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *IngestConfiguration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IngestConfiguration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IngestConfiguration) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IngestConfiguration) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPull

`func (o *IngestConfiguration) GetPull() bool`

GetPull returns the Pull field if non-nil, zero value otherwise.

### GetPullOk

`func (o *IngestConfiguration) GetPullOk() (*bool, bool)`

GetPullOk returns a tuple with the Pull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPull

`func (o *IngestConfiguration) SetPull(v bool)`

SetPull sets Pull field to given value.

### HasPull

`func (o *IngestConfiguration) HasPull() bool`

HasPull returns a boolean if a field has been set.

### GetProtocol

`func (o *IngestConfiguration) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IngestConfiguration) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IngestConfiguration) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IngestConfiguration) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEntryPoint

`func (o *IngestConfiguration) GetEntryPoint() string`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *IngestConfiguration) GetEntryPointOk() (*string, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *IngestConfiguration) SetEntryPoint(v string)`

SetEntryPoint sets EntryPoint field to given value.

### HasEntryPoint

`func (o *IngestConfiguration) HasEntryPoint() bool`

HasEntryPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


