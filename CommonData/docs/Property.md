# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the property | 
**Required** | Pointer to **bool** | Indicates whether the property is required – true&#x3D; required –  false(default)&#x3D; not required.  | [optional] 
**Regex** | Pointer to **string** | A regular expression string to be applied to the value of the property. | [optional] 
**Value** | Pointer to **string** | The property value. When present, it shall be a valid JSON string. | [optional] 

## Methods

### NewProperty

`func NewProperty(name string, ) *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Property) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Property) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Property) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *Property) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Property) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Property) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Property) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetRegex

`func (o *Property) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *Property) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *Property) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *Property) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetValue

`func (o *Property) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Property) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Property) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Property) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


