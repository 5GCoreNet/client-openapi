# MnsInfoSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MnsLabel** | Pointer to **string** |  | [optional] 
**MnsType** | Pointer to **string** |  | [optional] 
**MnsVersion** | Pointer to **string** |  | [optional] 
**MnsAddress** | Pointer to **string** |  | [optional] 
**MnsScope** | Pointer to **[]string** | List of the managed object instances that can be accessed using the MnS. If a complete SubNetwork can be accessed using the MnS, this attribute may contain the DN of the SubNetwork instead of the DNs of the individual managed entities within the SubNetwork. | [optional] 

## Methods

### NewMnsInfoSingle

`func NewMnsInfoSingle() *MnsInfoSingle`

NewMnsInfoSingle instantiates a new MnsInfoSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMnsInfoSingleWithDefaults

`func NewMnsInfoSingleWithDefaults() *MnsInfoSingle`

NewMnsInfoSingleWithDefaults instantiates a new MnsInfoSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMnsLabel

`func (o *MnsInfoSingle) GetMnsLabel() string`

GetMnsLabel returns the MnsLabel field if non-nil, zero value otherwise.

### GetMnsLabelOk

`func (o *MnsInfoSingle) GetMnsLabelOk() (*string, bool)`

GetMnsLabelOk returns a tuple with the MnsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsLabel

`func (o *MnsInfoSingle) SetMnsLabel(v string)`

SetMnsLabel sets MnsLabel field to given value.

### HasMnsLabel

`func (o *MnsInfoSingle) HasMnsLabel() bool`

HasMnsLabel returns a boolean if a field has been set.

### GetMnsType

`func (o *MnsInfoSingle) GetMnsType() string`

GetMnsType returns the MnsType field if non-nil, zero value otherwise.

### GetMnsTypeOk

`func (o *MnsInfoSingle) GetMnsTypeOk() (*string, bool)`

GetMnsTypeOk returns a tuple with the MnsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsType

`func (o *MnsInfoSingle) SetMnsType(v string)`

SetMnsType sets MnsType field to given value.

### HasMnsType

`func (o *MnsInfoSingle) HasMnsType() bool`

HasMnsType returns a boolean if a field has been set.

### GetMnsVersion

`func (o *MnsInfoSingle) GetMnsVersion() string`

GetMnsVersion returns the MnsVersion field if non-nil, zero value otherwise.

### GetMnsVersionOk

`func (o *MnsInfoSingle) GetMnsVersionOk() (*string, bool)`

GetMnsVersionOk returns a tuple with the MnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsVersion

`func (o *MnsInfoSingle) SetMnsVersion(v string)`

SetMnsVersion sets MnsVersion field to given value.

### HasMnsVersion

`func (o *MnsInfoSingle) HasMnsVersion() bool`

HasMnsVersion returns a boolean if a field has been set.

### GetMnsAddress

`func (o *MnsInfoSingle) GetMnsAddress() string`

GetMnsAddress returns the MnsAddress field if non-nil, zero value otherwise.

### GetMnsAddressOk

`func (o *MnsInfoSingle) GetMnsAddressOk() (*string, bool)`

GetMnsAddressOk returns a tuple with the MnsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAddress

`func (o *MnsInfoSingle) SetMnsAddress(v string)`

SetMnsAddress sets MnsAddress field to given value.

### HasMnsAddress

`func (o *MnsInfoSingle) HasMnsAddress() bool`

HasMnsAddress returns a boolean if a field has been set.

### GetMnsScope

`func (o *MnsInfoSingle) GetMnsScope() []string`

GetMnsScope returns the MnsScope field if non-nil, zero value otherwise.

### GetMnsScopeOk

`func (o *MnsInfoSingle) GetMnsScopeOk() (*[]string, bool)`

GetMnsScopeOk returns a tuple with the MnsScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsScope

`func (o *MnsInfoSingle) SetMnsScope(v []string)`

SetMnsScope sets MnsScope field to given value.

### HasMnsScope

`func (o *MnsInfoSingle) HasMnsScope() bool`

HasMnsScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


