# HalTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A human-readable string that can be used to identify this template | [optional] 
**Method** | [**HttpMethod**](HttpMethod.md) |  | 
**ContentType** | Pointer to **string** | The media type that should be used for the corresponding request. If the attribute is missing, or contains an unrecognized value, the client should act as if the  contentType is set to \&quot;application/json\&quot;.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The properties that should be included in the body of the corresponding request.  If the contentType attribute is set to \&quot;application/json\&quot;, then this attribute  describes the attributes of the JSON object of the body.  | [optional] 

## Methods

### NewHalTemplate

`func NewHalTemplate(method HttpMethod, ) *HalTemplate`

NewHalTemplate instantiates a new HalTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHalTemplateWithDefaults

`func NewHalTemplateWithDefaults() *HalTemplate`

NewHalTemplateWithDefaults instantiates a new HalTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *HalTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HalTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HalTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HalTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMethod

`func (o *HalTemplate) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HalTemplate) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HalTemplate) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.


### GetContentType

`func (o *HalTemplate) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HalTemplate) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HalTemplate) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HalTemplate) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetProperties

`func (o *HalTemplate) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *HalTemplate) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *HalTemplate) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *HalTemplate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


