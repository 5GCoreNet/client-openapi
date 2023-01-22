# DnsContextEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**DnsRuleId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**DnsQueryReport** | Pointer to [**DnsQueryReport**](DnsQueryReport.md) |  | [optional] 
**DnsRspReport** | Pointer to [**DnsRspReport**](DnsRspReport.md) |  | [optional] 
**DnsMsgId** | Pointer to **string** |  | [optional] 

## Methods

### NewDnsContextEventReport

`func NewDnsContextEventReport() *DnsContextEventReport`

NewDnsContextEventReport instantiates a new DnsContextEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsContextEventReportWithDefaults

`func NewDnsContextEventReportWithDefaults() *DnsContextEventReport`

NewDnsContextEventReportWithDefaults instantiates a new DnsContextEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *DnsContextEventReport) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DnsContextEventReport) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DnsContextEventReport) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DnsContextEventReport) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDnsRuleId

`func (o *DnsContextEventReport) GetDnsRuleId() int32`

GetDnsRuleId returns the DnsRuleId field if non-nil, zero value otherwise.

### GetDnsRuleIdOk

`func (o *DnsContextEventReport) GetDnsRuleIdOk() (*int32, bool)`

GetDnsRuleIdOk returns a tuple with the DnsRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRuleId

`func (o *DnsContextEventReport) SetDnsRuleId(v int32)`

SetDnsRuleId sets DnsRuleId field to given value.

### HasDnsRuleId

`func (o *DnsContextEventReport) HasDnsRuleId() bool`

HasDnsRuleId returns a boolean if a field has been set.

### GetDnsQueryReport

`func (o *DnsContextEventReport) GetDnsQueryReport() DnsQueryReport`

GetDnsQueryReport returns the DnsQueryReport field if non-nil, zero value otherwise.

### GetDnsQueryReportOk

`func (o *DnsContextEventReport) GetDnsQueryReportOk() (*DnsQueryReport, bool)`

GetDnsQueryReportOk returns a tuple with the DnsQueryReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQueryReport

`func (o *DnsContextEventReport) SetDnsQueryReport(v DnsQueryReport)`

SetDnsQueryReport sets DnsQueryReport field to given value.

### HasDnsQueryReport

`func (o *DnsContextEventReport) HasDnsQueryReport() bool`

HasDnsQueryReport returns a boolean if a field has been set.

### GetDnsRspReport

`func (o *DnsContextEventReport) GetDnsRspReport() DnsRspReport`

GetDnsRspReport returns the DnsRspReport field if non-nil, zero value otherwise.

### GetDnsRspReportOk

`func (o *DnsContextEventReport) GetDnsRspReportOk() (*DnsRspReport, bool)`

GetDnsRspReportOk returns a tuple with the DnsRspReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRspReport

`func (o *DnsContextEventReport) SetDnsRspReport(v DnsRspReport)`

SetDnsRspReport sets DnsRspReport field to given value.

### HasDnsRspReport

`func (o *DnsContextEventReport) HasDnsRspReport() bool`

HasDnsRspReport returns a boolean if a field has been set.

### GetDnsMsgId

`func (o *DnsContextEventReport) GetDnsMsgId() string`

GetDnsMsgId returns the DnsMsgId field if non-nil, zero value otherwise.

### GetDnsMsgIdOk

`func (o *DnsContextEventReport) GetDnsMsgIdOk() (*string, bool)`

GetDnsMsgIdOk returns a tuple with the DnsMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMsgId

`func (o *DnsContextEventReport) SetDnsMsgId(v string)`

SetDnsMsgId sets DnsMsgId field to given value.

### HasDnsMsgId

`func (o *DnsContextEventReport) HasDnsMsgId() bool`

HasDnsMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


