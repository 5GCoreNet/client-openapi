# SdpDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | **string** |  | 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewSdpDescription

`func NewSdpDescription(line string, ) *SdpDescription`

NewSdpDescription instantiates a new SdpDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdpDescriptionWithDefaults

`func NewSdpDescriptionWithDefaults() *SdpDescription`

NewSdpDescriptionWithDefaults instantiates a new SdpDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *SdpDescription) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SdpDescription) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SdpDescription) SetLine(v string)`

SetLine sets Line field to given value.


### GetContent

`func (o *SdpDescription) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SdpDescription) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SdpDescription) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SdpDescription) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


