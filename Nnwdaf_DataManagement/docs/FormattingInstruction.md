# FormattingInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsTrigNotif** | Pointer to **bool** | Indicates that notifications shall be buffered until the NF service consumer requests their delivery.  | [optional] 
**ReportingOptions** | Pointer to [**ReportingOptions**](ReportingOptions.md) |  | [optional] 

## Methods

### NewFormattingInstruction

`func NewFormattingInstruction() *FormattingInstruction`

NewFormattingInstruction instantiates a new FormattingInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattingInstructionWithDefaults

`func NewFormattingInstructionWithDefaults() *FormattingInstruction`

NewFormattingInstructionWithDefaults instantiates a new FormattingInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsTrigNotif

`func (o *FormattingInstruction) GetConsTrigNotif() bool`

GetConsTrigNotif returns the ConsTrigNotif field if non-nil, zero value otherwise.

### GetConsTrigNotifOk

`func (o *FormattingInstruction) GetConsTrigNotifOk() (*bool, bool)`

GetConsTrigNotifOk returns a tuple with the ConsTrigNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsTrigNotif

`func (o *FormattingInstruction) SetConsTrigNotif(v bool)`

SetConsTrigNotif sets ConsTrigNotif field to given value.

### HasConsTrigNotif

`func (o *FormattingInstruction) HasConsTrigNotif() bool`

HasConsTrigNotif returns a boolean if a field has been set.

### GetReportingOptions

`func (o *FormattingInstruction) GetReportingOptions() ReportingOptions`

GetReportingOptions returns the ReportingOptions field if non-nil, zero value otherwise.

### GetReportingOptionsOk

`func (o *FormattingInstruction) GetReportingOptionsOk() (*ReportingOptions, bool)`

GetReportingOptionsOk returns a tuple with the ReportingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOptions

`func (o *FormattingInstruction) SetReportingOptions(v ReportingOptions)`

SetReportingOptions sets ReportingOptions field to given value.

### HasReportingOptions

`func (o *FormattingInstruction) HasReportingOptions() bool`

HasReportingOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


