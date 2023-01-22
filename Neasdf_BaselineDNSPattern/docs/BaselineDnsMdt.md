# BaselineDnsMdt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MdtId** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**DnsQueryMdtList** | Pointer to [**map[string]DnsQueryMdt**](DnsQueryMdt.md) | map of DNS query message detection templates where a valid JSON string serves as key | [optional] 
**DnsRspMdtList** | Pointer to [**map[string]DnsRspMdt**](DnsRspMdt.md) | map of DNS response message detection templates where a valid JSON string serves as key | [optional] 

## Methods

### NewBaselineDnsMdt

`func NewBaselineDnsMdt(mdtId string, ) *BaselineDnsMdt`

NewBaselineDnsMdt instantiates a new BaselineDnsMdt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaselineDnsMdtWithDefaults

`func NewBaselineDnsMdtWithDefaults() *BaselineDnsMdt`

NewBaselineDnsMdtWithDefaults instantiates a new BaselineDnsMdt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMdtId

`func (o *BaselineDnsMdt) GetMdtId() string`

GetMdtId returns the MdtId field if non-nil, zero value otherwise.

### GetMdtIdOk

`func (o *BaselineDnsMdt) GetMdtIdOk() (*string, bool)`

GetMdtIdOk returns a tuple with the MdtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtId

`func (o *BaselineDnsMdt) SetMdtId(v string)`

SetMdtId sets MdtId field to given value.


### GetLabel

`func (o *BaselineDnsMdt) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BaselineDnsMdt) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BaselineDnsMdt) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BaselineDnsMdt) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDnsQueryMdtList

`func (o *BaselineDnsMdt) GetDnsQueryMdtList() map[string]DnsQueryMdt`

GetDnsQueryMdtList returns the DnsQueryMdtList field if non-nil, zero value otherwise.

### GetDnsQueryMdtListOk

`func (o *BaselineDnsMdt) GetDnsQueryMdtListOk() (*map[string]DnsQueryMdt, bool)`

GetDnsQueryMdtListOk returns a tuple with the DnsQueryMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQueryMdtList

`func (o *BaselineDnsMdt) SetDnsQueryMdtList(v map[string]DnsQueryMdt)`

SetDnsQueryMdtList sets DnsQueryMdtList field to given value.

### HasDnsQueryMdtList

`func (o *BaselineDnsMdt) HasDnsQueryMdtList() bool`

HasDnsQueryMdtList returns a boolean if a field has been set.

### GetDnsRspMdtList

`func (o *BaselineDnsMdt) GetDnsRspMdtList() map[string]DnsRspMdt`

GetDnsRspMdtList returns the DnsRspMdtList field if non-nil, zero value otherwise.

### GetDnsRspMdtListOk

`func (o *BaselineDnsMdt) GetDnsRspMdtListOk() (*map[string]DnsRspMdt, bool)`

GetDnsRspMdtListOk returns a tuple with the DnsRspMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRspMdtList

`func (o *BaselineDnsMdt) SetDnsRspMdtList(v map[string]DnsRspMdt)`

SetDnsRspMdtList sets DnsRspMdtList field to given value.

### HasDnsRspMdtList

`func (o *BaselineDnsMdt) HasDnsRspMdtList() bool`

HasDnsRspMdtList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


