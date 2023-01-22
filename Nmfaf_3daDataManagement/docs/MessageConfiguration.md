# MessageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorreId** | **string** | If the configuration is used for mapping analytics or data collection, representing the Analytics Consumer Notification Correlation ID or Data Consumer Notification Correlation ID.  | 
**FormatInstruct** | Pointer to [**FormattingInstruction**](FormattingInstruction.md) |  | [optional] 
**MfafNotiInfo** | Pointer to [**MfafNotiInfo**](MfafNotiInfo.md) |  | [optional] 
**NotificationURI** | **string** | String providing an URI formatted according to RFC 3986. | 
**ProcInstruct** | Pointer to [**ProcessingInstruction**](ProcessingInstruction.md) |  | [optional] 
**AdrfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewMessageConfiguration

`func NewMessageConfiguration(correId string, notificationURI string, ) *MessageConfiguration`

NewMessageConfiguration instantiates a new MessageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageConfigurationWithDefaults

`func NewMessageConfigurationWithDefaults() *MessageConfiguration`

NewMessageConfigurationWithDefaults instantiates a new MessageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorreId

`func (o *MessageConfiguration) GetCorreId() string`

GetCorreId returns the CorreId field if non-nil, zero value otherwise.

### GetCorreIdOk

`func (o *MessageConfiguration) GetCorreIdOk() (*string, bool)`

GetCorreIdOk returns a tuple with the CorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorreId

`func (o *MessageConfiguration) SetCorreId(v string)`

SetCorreId sets CorreId field to given value.


### GetFormatInstruct

`func (o *MessageConfiguration) GetFormatInstruct() FormattingInstruction`

GetFormatInstruct returns the FormatInstruct field if non-nil, zero value otherwise.

### GetFormatInstructOk

`func (o *MessageConfiguration) GetFormatInstructOk() (*FormattingInstruction, bool)`

GetFormatInstructOk returns a tuple with the FormatInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatInstruct

`func (o *MessageConfiguration) SetFormatInstruct(v FormattingInstruction)`

SetFormatInstruct sets FormatInstruct field to given value.

### HasFormatInstruct

`func (o *MessageConfiguration) HasFormatInstruct() bool`

HasFormatInstruct returns a boolean if a field has been set.

### GetMfafNotiInfo

`func (o *MessageConfiguration) GetMfafNotiInfo() MfafNotiInfo`

GetMfafNotiInfo returns the MfafNotiInfo field if non-nil, zero value otherwise.

### GetMfafNotiInfoOk

`func (o *MessageConfiguration) GetMfafNotiInfoOk() (*MfafNotiInfo, bool)`

GetMfafNotiInfoOk returns a tuple with the MfafNotiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfafNotiInfo

`func (o *MessageConfiguration) SetMfafNotiInfo(v MfafNotiInfo)`

SetMfafNotiInfo sets MfafNotiInfo field to given value.

### HasMfafNotiInfo

`func (o *MessageConfiguration) HasMfafNotiInfo() bool`

HasMfafNotiInfo returns a boolean if a field has been set.

### GetNotificationURI

`func (o *MessageConfiguration) GetNotificationURI() string`

GetNotificationURI returns the NotificationURI field if non-nil, zero value otherwise.

### GetNotificationURIOk

`func (o *MessageConfiguration) GetNotificationURIOk() (*string, bool)`

GetNotificationURIOk returns a tuple with the NotificationURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationURI

`func (o *MessageConfiguration) SetNotificationURI(v string)`

SetNotificationURI sets NotificationURI field to given value.


### GetProcInstruct

`func (o *MessageConfiguration) GetProcInstruct() ProcessingInstruction`

GetProcInstruct returns the ProcInstruct field if non-nil, zero value otherwise.

### GetProcInstructOk

`func (o *MessageConfiguration) GetProcInstructOk() (*ProcessingInstruction, bool)`

GetProcInstructOk returns a tuple with the ProcInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcInstruct

`func (o *MessageConfiguration) SetProcInstruct(v ProcessingInstruction)`

SetProcInstruct sets ProcInstruct field to given value.

### HasProcInstruct

`func (o *MessageConfiguration) HasProcInstruct() bool`

HasProcInstruct returns a boolean if a field has been set.

### GetAdrfId

`func (o *MessageConfiguration) GetAdrfId() string`

GetAdrfId returns the AdrfId field if non-nil, zero value otherwise.

### GetAdrfIdOk

`func (o *MessageConfiguration) GetAdrfIdOk() (*string, bool)`

GetAdrfIdOk returns a tuple with the AdrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfId

`func (o *MessageConfiguration) SetAdrfId(v string)`

SetAdrfId sets AdrfId field to given value.

### HasAdrfId

`func (o *MessageConfiguration) HasAdrfId() bool`

HasAdrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


