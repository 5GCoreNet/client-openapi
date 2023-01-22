# ManagePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**AppId** | **string** | Identifies the application. | 
**ManageEntity** | Pointer to [**ManageEntity**](ManageEntity.md) |  | [optional] 
**SkipUeInquiry** | Pointer to **bool** | Indicate whether to skip UE inquiry. | [optional] 
**SupportedFormats** | Pointer to [**[]SerializationFormat**](SerializationFormat.md) | Indicates the serialization format(s) that are supported by the SCS/AS on the associated RDS port. | [optional] 
**ConfiguredFormat** | Pointer to [**SerializationFormat**](SerializationFormat.md) |  | [optional] 

## Methods

### NewManagePort

`func NewManagePort(appId string, ) *ManagePort`

NewManagePort instantiates a new ManagePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagePortWithDefaults

`func NewManagePortWithDefaults() *ManagePort`

NewManagePortWithDefaults instantiates a new ManagePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ManagePort) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ManagePort) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ManagePort) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ManagePort) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetAppId

`func (o *ManagePort) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ManagePort) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ManagePort) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetManageEntity

`func (o *ManagePort) GetManageEntity() ManageEntity`

GetManageEntity returns the ManageEntity field if non-nil, zero value otherwise.

### GetManageEntityOk

`func (o *ManagePort) GetManageEntityOk() (*ManageEntity, bool)`

GetManageEntityOk returns a tuple with the ManageEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageEntity

`func (o *ManagePort) SetManageEntity(v ManageEntity)`

SetManageEntity sets ManageEntity field to given value.

### HasManageEntity

`func (o *ManagePort) HasManageEntity() bool`

HasManageEntity returns a boolean if a field has been set.

### GetSkipUeInquiry

`func (o *ManagePort) GetSkipUeInquiry() bool`

GetSkipUeInquiry returns the SkipUeInquiry field if non-nil, zero value otherwise.

### GetSkipUeInquiryOk

`func (o *ManagePort) GetSkipUeInquiryOk() (*bool, bool)`

GetSkipUeInquiryOk returns a tuple with the SkipUeInquiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUeInquiry

`func (o *ManagePort) SetSkipUeInquiry(v bool)`

SetSkipUeInquiry sets SkipUeInquiry field to given value.

### HasSkipUeInquiry

`func (o *ManagePort) HasSkipUeInquiry() bool`

HasSkipUeInquiry returns a boolean if a field has been set.

### GetSupportedFormats

`func (o *ManagePort) GetSupportedFormats() []SerializationFormat`

GetSupportedFormats returns the SupportedFormats field if non-nil, zero value otherwise.

### GetSupportedFormatsOk

`func (o *ManagePort) GetSupportedFormatsOk() (*[]SerializationFormat, bool)`

GetSupportedFormatsOk returns a tuple with the SupportedFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFormats

`func (o *ManagePort) SetSupportedFormats(v []SerializationFormat)`

SetSupportedFormats sets SupportedFormats field to given value.

### HasSupportedFormats

`func (o *ManagePort) HasSupportedFormats() bool`

HasSupportedFormats returns a boolean if a field has been set.

### GetConfiguredFormat

`func (o *ManagePort) GetConfiguredFormat() SerializationFormat`

GetConfiguredFormat returns the ConfiguredFormat field if non-nil, zero value otherwise.

### GetConfiguredFormatOk

`func (o *ManagePort) GetConfiguredFormatOk() (*SerializationFormat, bool)`

GetConfiguredFormatOk returns a tuple with the ConfiguredFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredFormat

`func (o *ManagePort) SetConfiguredFormat(v SerializationFormat)`

SetConfiguredFormat sets ConfiguredFormat field to given value.

### HasConfiguredFormat

`func (o *ManagePort) HasConfiguredFormat() bool`

HasConfiguredFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


